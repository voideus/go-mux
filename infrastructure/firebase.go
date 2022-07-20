package infrastructure

import (
	"context"
	"fmt"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

// Creates new firebase app instance
func NewFBApp() *firebase.App {
	ctx := context.Background()

	serviceAccountKeyFilePath, err := filepath.Abs("./firebase/token.json")
	if err != nil {
		fmt.Println("Unable to load serviceAccountKey.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		fmt.Printf("Firebase NewApp: %v", err)
	}
	fmt.Println("✅ Firebase app initialized.")
	return app
}

// Creates new firebase auth client
func NewFBAuth(app *firebase.App) *auth.Client {
	ctx := context.Background()

	firebaseAuth, err := app.Auth(ctx)
	if err != nil {
		fmt.Printf("Firebase Authentication: %v", err)
	}

	return firebaseAuth
}
