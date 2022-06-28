package investapi

import (
	"crypto/tls"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func initConnection() (conn *grpc.ClientConn, err error) {
	conn, err = grpc.Dial("invest-public-api.tinkoff.ru:443",
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithPerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: os.Getenv("TINK_TOKEN"),
		})))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return
}
