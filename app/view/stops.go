package view

import (
	"go-gtfs-server/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetStopsByAgency(c echo.Context) error {
	agencyId := c.Param("agencyId")
	gtfs, found := db.GetGtfsByAgencyId(agencyId)
	if found {
		name := c.QueryParam("stop_name")
		if name == "" {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"name_query":  nil,
				"stops":       gtfs.Gtfs.Stops,
				"stops_count": len(gtfs.Gtfs.Stops),
				"agency_id":   agencyId,
			})
		} else {
			filteredStops := gtfs.FilterStopsByName(name)
			if len(filteredStops) > 0 {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"name_query":  name,
					"stops":       filteredStops,
					"stops_count": len(filteredStops),
					"agency_id":   agencyId,
				})
			} else {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"name_query":  name,
					"stops":       nil,
					"stops_count": 0,
					"agency_id":   agencyId,
				})
			}

		}
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Agency Not found",
	})
}

func GetStopByAgencyAndStopId(c echo.Context) error {
	agencyId := c.Param("agencyId")
	stopId := c.Param("stopId")
	gtfs, found := db.GetGtfsByAgencyId(agencyId)
	if found {
		stop, found := gtfs.GetStopById(stopId)
		if found {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"stop_id":   stopId,
				"agency_id": agencyId,
				"stop":      stop,
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Stop Not found",
	})
}
