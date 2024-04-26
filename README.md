# Backend Service
This service uses GoLang + gRPC as the backend service.
A dockerfile is provided in the root folder to be used in GCP's Cloud Build + Run product to be automatically deployed + built. No test case added for now, will be added at a later stage.

Server reflection is on currently.

## Makefile
The Makefile consists of `gen-proto` command to help generate gRPC code based on the protobuf in another repo. A `protobuf` folder must need to be created in the outer directory and clone [this repo](https://github.com/ManeHKU/protobuf).
The generated proto file will be stored in the `pb` folder.

## Running the service
MAKE SURE you add the environment variables in your local environment or in the Dockerfile.
The 2 environment variables are: DB_URI and JWT_SECRET

To run this file, just use the Dockerfile provided in the root folder.

Make sure you have chromium installed in your local machine first (or go rod will install it for you).
If you need to run this locally, just run `make run` in the root folder.