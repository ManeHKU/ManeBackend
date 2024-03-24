package service

import (
	"ManeBackend/pb"
	"context"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strings"
	"time"
)

type InitService struct {
	browser *rod.Browser
	pb.UnimplementedInitServiceServer
}

func (s *InitService) getFreshBrowserWithCookies(cookies []*pb.Cookie) (*rod.Browser, error) {
	browser := s.browser
	browser.MustSetCookies()
	if cookies != nil && len(cookies) > 0 {
		browserCookies := make([]*proto.NetworkCookieParam, len(cookies))
		for i, cookie := range cookies {
			sourcePort := int(cookie.SourcePort)
			cookie := &proto.NetworkCookieParam{
				Name:         cookie.Name,
				Value:        cookie.Value,
				Domain:       cookie.Domain,
				Path:         cookie.Path,
				Expires:      proto.TimeSinceEpoch(cookie.Expires),
				HTTPOnly:     cookie.HTTPOnly,
				Secure:       cookie.Secure,
				SameSite:     proto.NetworkCookieSameSite(cookie.GetSameSite()),
				Priority:     proto.NetworkCookiePriority(cookie.GetPriority()),
				SameParty:    cookie.SameParty,
				SourceScheme: proto.NetworkCookieSourceScheme(cookie.SourceScheme),
				SourcePort:   &sourcePort,
				PartitionKey: cookie.GetPartitionKey(),
			}
			browserCookies[i] = cookie
		}
		err := browser.SetCookies(browserCookies)
		if err != nil {
			return browser, err
		}
	}
	return browser, nil
}

func (s *InitService) GetSISTicket(_ context.Context, request *pb.UserSignInRequest) (*pb.UserSignInResponse, error) {
	log.Printf("Received GetSISTicket request")
	portalId := request.GetUserId()
	password := request.GetPassword()
	clientCookies := request.GetCookies()
	if portalId == "" || strings.Contains(portalId, "@") {
		log.Printf("Invalid portal id, rejecting request")
		return nil, status.Errorf(codes.InvalidArgument, "Invalid portal id format")
	}
	if password == "" {
		log.Printf("Empty password, rejecting request")
		return nil, status.Errorf(codes.InvalidArgument, "Empty password")
	}

	browser, err := s.getFreshBrowserWithCookies(clientCookies)
	if err != nil {
		log.Printf("Error setting cookies: %v", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	page := browser.MustPage("https://hkuportal.hku.hk/cas/login?service=HKUPORTAL")
	defer func() {
		log.Print("Ending GetSISTicket")
		page.MustClose()
	}()

	hkuEmailInput, err := page.Element("#email")
	if err != nil {
		log.Printf("Error finding hku email input: %v", err)
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}

	err = rod.Try(func() {
		hkuEmailInput.MustInput(fmt.Sprintf("%v@connect.hku.hk", portalId)).MustType(input.Enter)
	})
	if err != nil {
		log.Printf("Error inputting email: %v", err)
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}
	log.Printf("Input email success")

	err = rod.Try(func() {
		page.MustElement("#passwordInput").MustInput(password).MustType(input.Enter)
	})
	if err != nil {
		log.Printf("Error inputting password: %v", err)
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}
	log.Printf("Entered password")

	loginResult, err := page.Timeout(2*time.Second).Race().ElementR("#errorText", "Incorrect user ID or password.").Element("input[value='Yes']").Element("input[value='Continue']").Do()
	if err != nil {
		log.Printf("Error finding login result: %v", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if result, _ := loginResult.Matches("input[value='Continue']"); result {
		log.Printf("Trust connect or not?")
		loginResult.MustClick()

		loginResult, _ = page.Element("input[value='Yes']")
		if loginResult == nil {
			log.Printf("Failed to trust website")
			response := &pb.UserSignInResponse{
				CanLogInToSIS: false,
			}
			return response, nil
		}
	}

	if result, _ := loginResult.Matches("input[value='Yes']"); !result {
		log.Printf("Login failed")
		response := &pb.UserSignInResponse{
			CanLogInToSIS: false,
		}
		return response, nil
	}

	ticketUrlChannel := make(chan string)
	router := page.HijackRequests()
	defer router.MustStop()
	router.MustAdd("*sis-eportal.hku.hk/ptlprod/z_signon.jsp?ticket*", func(ctx *rod.Hijack) {
		ticketUrl := ctx.Request.URL().String()
		log.Printf("Got full ticket url: %v", ctx.Request.URL())
		ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
		ticketUrlChannel <- ticketUrl
		err := router.Stop()
		if err != nil {
			log.Printf("Error stopping router: %v", err)
		}
		return
	})

	go router.Run()
	loginResult.MustClick()
	log.Printf("Pressed rmb button to login")

	var ticketUrl string
	select {
	case url := <-ticketUrlChannel:
		ticketUrl = url
	case <-time.After(3 * time.Second):
		log.Printf("Timeout waiting for ticket url")
		response := &pb.UserSignInResponse{
			CanLogInToSIS: false,
		}
		return response, nil
	}
	log.Printf("Ticket url: %v, getting cookies now", ticketUrl)

	cookies, _ := browser.GetCookies()
	pbCookies := make([]*pb.Cookie, len(cookies))
	for i, cookie := range cookies {
		sameSite := string(cookie.SameSite)
		pbCookie := &pb.Cookie{
			Name:         cookie.Name,
			Value:        cookie.Value,
			Domain:       cookie.Domain,
			Path:         cookie.Path,
			Expires:      float32(cookie.Expires),
			HTTPOnly:     cookie.HTTPOnly,
			Secure:       cookie.Secure,
			SameSite:     &sameSite,
			Priority:     string(cookie.Priority),
			SameParty:    cookie.SameParty,
			SourceScheme: string(cookie.SourceScheme),
			SourcePort:   int32(cookie.SourcePort),
			PartitionKey: &cookie.PartitionKey,
		}
		pbCookies[i] = pbCookie
	}
	response := &pb.UserSignInResponse{
		CanLogInToSIS: true,
		TicketURL:     &ticketUrl,
		Cookies:       pbCookies,
	}
	return response, nil
}

func NewInitService(browser *rod.Browser) *InitService {
	return &InitService{browser: browser}
}
