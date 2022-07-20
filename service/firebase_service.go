package service

import (
	"context"

	"firebase.google.com/go/auth"
)

type FirebaseService struct {
	*auth.Client
}

func NewFirebaseService(fb *auth.Client) FirebaseService {
	return FirebaseService{
		fb,
	}
}

func (fb *FirebaseService) VerifyToken(idToken string) (*auth.Token, error) {
	token, err := fb.VerifyIDToken(context.Background(), idToken)
	return token, err
}
