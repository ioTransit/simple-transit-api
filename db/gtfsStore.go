package db

import (
	"go-gtfs-server/app/model"
	"strings"

	"github.com/geops/gtfsparser"
	"github.com/geops/gtfsparser/gtfs"
)

type AgencyStore struct {
	Gtfs   *gtfsparser.Feed
	Config model.AgencyConfig
}

var GtfsStore = make(map[string]AgencyStore)

func GetGtfsByAgencyId(agencyId string) (agencyStory AgencyStore, found bool) {
	gtfs, found := GtfsStore[agencyId]
	return gtfs, found
}

func (agencyStore *AgencyStore) GetStopById(stopId string) (stop *gtfs.Stop, found bool) {
	stop, found = agencyStore.Gtfs.Stops[stopId]
	if found {
		return stop, found
	}
	return nil, false
}

func (agencyStore *AgencyStore) GetRouteByid(routeId string) (route *gtfs.Route, found bool) {
	route, found = agencyStore.Gtfs.Routes[routeId]
	if found {
		return route, found
	}
	return nil, false
}

func (agencyStore *AgencyStore) GetTripByid(tripId string) (trip *gtfs.Trip, found bool) {
	trip, found = agencyStore.Gtfs.Trips[tripId]
	if found {
		return trip, found
	}
	return nil, false
}

func (agencyStore *AgencyStore) GetTripsByRouteId(routeId string) (trip []*gtfs.Trip, found bool) {
	var trips []*gtfs.Trip
	for _, trip := range agencyStore.Gtfs.Trips {
		if trip.Route.Id == routeId {
			trips = append(trips, trip)
		}
	}
	return trips, true
}

func (agencyStore *AgencyStore) FilterStopsByName(name string) map[string]*gtfs.Stop {
	filtered := make(map[string]*gtfs.Stop)
	for id, stop := range agencyStore.Gtfs.Stops {
		if strings.HasPrefix(stop.Name, name) {
			filtered[id] = stop
		}
	}
	return filtered
}

func (agencyStore *AgencyStore) FilterRoutesByName(name string) map[string]*gtfs.Route {
	filtered := make(map[string]*gtfs.Route)
	for id, route := range agencyStore.Gtfs.Routes {
		if strings.HasPrefix(route.Long_name, name) {
			filtered[id] = route
		}
	}
	return filtered
}
