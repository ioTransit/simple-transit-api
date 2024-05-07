package view

import (
	"go-gtfs-server/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetShapesByAgency(c echo.Context) error {
	agencyId := c.Param("agencyId")
	gtfs, found := db.GetGtfsByAgencyId(agencyId)
	if found {
		name := c.QueryParam("route_long_name")
		if name == "" {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"name_query":   nil,
				"shapes":       gtfs.Gtfs.Shapes,
				"shapes_count": len(gtfs.Gtfs.Shapes),
				"agency_id":    agencyId,
			})
		} else {
			filteredRoutes := gtfs.FilterRoutesByName(name)
			if len(filteredRoutes) > 0 {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"name_query":   name,
					"routes":       filteredRoutes,
					"routes_count": len(filteredRoutes),
					"agency_id":    agencyId,
				})
			} else {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"name_query":   name,
					"routes":       nil,
					"routes_count": 0,
					"agency_id":    agencyId,
				})
			}

		}
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Agency Not found",
	})
}

func GetShapeByAgencyAndRouteId(c echo.Context) error {
	agencyId := c.Param("agencyId")
	routeId := c.Param("routeId")
	gtfs, found := db.GetGtfsByAgencyId(agencyId)
	if found {
		route, found := gtfs.GetRouteByid(routeId)
		if found {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"route_id":  routeId,
				"agency_id": agencyId,
				"route":     route,
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Stop Not found",
	})
}
