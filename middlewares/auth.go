package middlewares

import (
	"context"
	"fmt"
	"log"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

var (
	client *auth.Client
	ctx    context.Context
	u_info *auth.UserInfo
	err    error
)

func InitClient() error {
	ctx = context.Background()
	opt := option.WithCredentialsFile("middlewares/firebase.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}
	client, err = app.Auth(ctx)

	if err != nil {
		return err
	}
	return nil
}

func getFirebaseUser(token string) (*auth.UserInfo, error) {
	t, err := client.VerifyIDToken(ctx, token)
	fmt.Println("Check Token")
	if err != nil {
		return nil, err
	}
	u, err := client.GetUser(ctx, t.UID)
	fmt.Println("Search User")
	if err != nil {
		return nil, err
	}
	log.Printf("Successfully fetched user data: %#v\n", u.UserInfo)

	return u.UserInfo, nil
}

func FirebaseAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.String(403, "No authorization header provided")
			c.Abort()
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		if token == auth {
			c.String(403, "Could not find bearer token in Authorization header")
			c.Abort()
			return
		}
		u_info, err = getFirebaseUser(token)
		if err != nil {
			c.String(403, "This user is not found in firebase")
			c.Abort()
			return
		}
		c.Next()
		u_info = nil
	}
}

func GetCurrentUserInfo() *auth.UserInfo {
	return u_info
}

func GetClient() *auth.Client {
	return client
}
