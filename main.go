package main

import (
	jobs "go-gtfs-server/app/job"
	"go-gtfs-server/app/view"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error with .env")
	}

	jobs.Load()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome to the gtfs server")
	})

	// Start of Stops
	e.GET("/stops/:agencyId/", view.GetStopsByAgency)
	e.GET("/stops/:agencyId/:stopId/", view.GetStopByAgencyAndStopId)

	// Start of Routes
	e.GET("/routes/:agencyId/", view.GetRoutesByAgency)
	e.GET("/routes/:agencyId/:routeId/", view.GetRouteByAgencyAndRouteId)

	// Start of Trips
	e.GET("/trips/:agencyId/:tripId/", view.GetTripByAgencyAndTripId)
	e.GET("/trips/:agencyId/route/:routeId/", view.GetTripsByAgencyAndRouteId)

	e.Start(":1080")
}
