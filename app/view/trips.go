package view

import (
	"go-gtfs-server/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTripByAgencyAndTripId(c echo.Context) error {
	agencyId := c.Param("agencyId")
	tripId := c.Param("tripId")
	gtfs, found := db.GetGtfsByAgencyId(agencyId)
	if found {
		trip, found := gtfs.GetTripByid(tripId)
		if found {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"trip_id":   tripId,
				"agency_id": agencyId,
				"trip":      trip,
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Stop Not found",
	})
}

func GetTripsByAgencyAndRouteId(c echo.Context) error {
	agencyId := c.Param("agencyId")
	routeId := c.Param("routeId")
	gtfs, found := db.GetGtfsByAgencyId(agencyId)
	if found {
		trips, found := gtfs.GetTripsByRouteId(routeId)
		if found {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"route_id":  routeId,
				"agency_id": agencyId,
				"trips":     trips,
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Stop Not found",
	})
}
