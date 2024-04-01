package types

type Course struct {
	CourseCode  string
	Title       string
	Department  string
	Description *string
	Rating      float32
	Offered     bool
}
