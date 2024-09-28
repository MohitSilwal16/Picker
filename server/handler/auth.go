package handler

import (
	"context"
	"log"
	"regexp"

	"github.com/MohitSilwal16/Picker/server/db"
	"github.com/MohitSilwal16/Picker/server/errs"
	"github.com/MohitSilwal16/Picker/server/pb"
	"github.com/MohitSilwal16/Picker/server/utils"
)

type AuthServer struct {
	pb.UnimplementedAuthServer
}

func (s *AuthServer) Register(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	// Errors:
	// USERNAME IS ALREADY USED
	// USER REGISTERED BUT FAILED TO INIT DIR
	// INTERNAL SERVER ERROR

	if !regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`).MatchString(req.Name) {
		return nil, errs.ErrInvalidUsernameFormat
	}
	if !utils.IsPasswordInFormat(req.Pass) {
		return nil, errs.ErrInvalidPasswordFormat
	}

	sessionToken, err := db.Register(req)
	if err != nil {
		if err.Error() == "USERNAME IS ALREADY USED" {
			return nil, errs.ErrUsernameAlreadyUsedError
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot Register User")
		log.Println("Source: Register()")
		return nil, errs.ErrInternalServerError
	}

	return &pb.AuthResponse{SessionToken: sessionToken}, nil
}

func (s *AuthServer) Login(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	// Errors:
	// INVALID CREDENTIALS
	// INTERNAL SERVER ERROR

	if !regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`).MatchString(req.Name) {
		return nil, errs.ErrInvalidUsernameFormat
	} else if !utils.IsPasswordInFormat(req.Pass) {
		return nil, errs.ErrInvalidPasswordFormat
	}

	sessionToken, err := db.Login(req)
	if err != nil {
		if err.Error() == "INVALID CREDENTIALS" {
			return nil, errs.ErrInvalidCredentials
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot Login User")
		log.Println("Source: Login()")
		return nil, errs.ErrInternalServerError
	}

	return &pb.AuthResponse{SessionToken: sessionToken}, nil
}

func (s *AuthServer) VerifySessionToken(ctx context.Context, req *pb.VerifySessionTokenRequest) (*pb.VerifySessionTokenResponse, error) {
	// Errors:
	// INTERNAL SERVER ERROR

	if req.SessionToken == "" {
		return &pb.VerifySessionTokenResponse{IsSessionTokenValid: false}, nil
	}

	isSessionTokenValid, err := db.IsSessionTokenValid(req.SessionToken)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Verify Session Token")
		log.Println("Source: VerifySessionToken()")

		return nil, errs.ErrInternalServerError
	}

	return &pb.VerifySessionTokenResponse{IsSessionTokenValid: isSessionTokenValid}, nil
}

func (s *AuthServer) Logout(ctx context.Context, req *pb.LogOutRequest) (*pb.LogOutResponse, error) {
	if req.SessionToken == "" {
		return &pb.LogOutResponse{IsUserLoggedOut: false}, nil
	}

	isUserLoggedOut, err := db.Logout(req.SessionToken)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Log Out User")
		log.Println("Source: LogOut()")

		return nil, errs.ErrInternalServerError
	}
	return &pb.LogOutResponse{IsUserLoggedOut: isUserLoggedOut}, nil
}
