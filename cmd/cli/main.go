package main

import (
	"docomo-bike/cmd/cli/commands"
	"docomo-bike/internal/libs/docomo/getstation"
	"docomo-bike/internal/libs/docomo/login"
	"docomo-bike/internal/libs/logging"
	"docomo-bike/internal/services/auth"
	"docomo-bike/internal/services/listing"
	"os"
	"time"

	"github.com/gojektech/heimdall/httpclient"

	"github.com/dgrijalva/jwt-go"

	"github.com/spf13/cobra"
)

func main() {
	logger := logging.New("App", false, false, os.Stdout, false)

	httpClient := httpclient.NewClient()
	loginClient := &login.ScrappingClient{
		HTTPClient: httpClient,
		Logger:     logger,
	}
	getstationClient := &getstation.ScrappingClient{
		HTTPClient: httpClient,
		Logger:     logger,
	}

	authServ := auth.NewService(auth.JWTConfig{
		ExpiresIn:     60 * time.Second,
		Issuer:        "docomo",
		Secret:        []byte("docomo"),
		SigningMethod: jwt.SigningMethodHS512,
	}, loginClient)
	listingServ := listing.NewService(getstationClient)

	var rootCmd = &cobra.Command{Use: "docomo"}
	rootCmd.AddCommand(commands.Station(authServ, listingServ))
	rootCmd.Execute()
}
