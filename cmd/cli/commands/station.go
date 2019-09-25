package commands

import (
	"docomo-bike/internal/services/auth"
	"docomo-bike/internal/services/listing"
	"log"

	"github.com/spf13/cobra"
)

func Station(
	authServ auth.JWTService,
	listingServ listing.Service,
) *cobra.Command {
	c := &cobra.Command{
		Use:   "station [station name to list]",
		Short: "List a station.",
		Long:  `list a station.`,
		Args:  cobra.ExactArgs(1),
	}
	var (
		username string
		password string
	)
	c.Flags().StringVarP(&username, "username", "u", "", "username for login Docomo")
	c.MarkFlagRequired("username")
	c.Flags().StringVarP(&password, "password", "p", "", "password for login Docomo")
	c.MarkFlagRequired("password")

	c.Run = func(cmd *cobra.Command, args []string) {
		stationName := args[0]
		jwt, err := authServ.Authorize(username, password)
		if err != nil {
			log.Fatalf("Failed to login [error=%s]", err.Error())
			return
		}
		auth, err := authServ.AuthFromToken(jwt.TokenString)
		if err != nil {
			log.Fatalf("Token is invalid [error=%s]", err.Error())
			return
		}
		station, err := listingServ.GetStation(auth, stationName)
		if err != nil {
			log.Fatalf("Failed to get station [error=%s]", err.Error())
			return
		}
		if station == nil {
			printNoStation(stationName)
			return
		}
		printStation(station)
	}

	return c
}
func printNoStation(stationName string) {
	log.Printf("No such station [stationName=%s]\n", stationName)
}
func printStation(station *listing.Station) {
	log.Printf("Station ID: %s\n", station.ID)
	log.Printf("Station name: %s\n", station.Name)
	for idx, b := range station.Bikes {
		printBike(idx+1, b)
	}
}
func printBike(num int, bike *listing.Bike) {
	log.Printf("Bike ID: %s\n", bike.ID)
}
