package commands

import (
	"docomo-bike/internal/services/auth"
	"docomo-bike/internal/services/booking"
	"docomo-bike/internal/services/listing"
	"log"
	"time"

	"github.com/spf13/cobra"
)

func StationBook(
	authServ auth.JWTService,
	listingServ listing.Service,
	bookingServ booking.Service,
) *cobra.Command {
	c := &cobra.Command{
		Use:   "book [station name in which you want to book a bike]",
		Short: "Book a bike.",
		Long:  `book any bike in the station.`,
		Args:  cobra.ExactArgs(1),
	}
	var (
		username    string
		password    string
		stationName string
	)
	c.Flags().StringVarP(&username, "username", "u", "", "username for login Docomo")
	c.MarkFlagRequired("username")
	c.Flags().StringVarP(&password, "password", "p", "", "password for login Docomo")
	c.MarkFlagRequired("password")
	c.Flags().StringVarP(&stationName, "station", "s", "", "station to book a bike")
	c.MarkFlagRequired("station")

	c.Run = func(cmd *cobra.Command, args []string) {
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

		for {
			station, err := listingServ.GetStation(auth, stationName)
			if err != nil {
				log.Fatalf("Failed to get station [error=%s]", err.Error())
				return
			}
			if len(station.Bikes) == 0 {
				time.Sleep(3 * time.Second)
				continue
			}

			bikeID := station.Bikes[0].ID
			bookingResult, err := bookingServ.BookBike(bikeID)
			if err != nil {
				log.Fatalf("Failed to book [bikeID=%d]", bikeID)
				return
			}
			printBookingResult(bookingResult)
		}
	}

	return c
}
func printBookingResult(result booking.BookingResult) {
	panic("Not implemented")
}
