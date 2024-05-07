package providers

import "go-gtfs-server/app/model"

var TriangleTransitConfig = model.AgencyConfig{
	Id:             "triangletransit",
	Name:           "Triangle Transit",
	Url:            "https://gotriangle.org/gtfs",
	TripUpdates:    "https://gotriangle.tripsparkhost.com/gtfs/Realtime/GTFS_TripUpdates.pb",
	AlertUpdates:   "https://gotriangle.tripsparkhost.com/gtfs/Realtime/GTFS_ServiceAlerts.pb",
	VehicleUpdates: "https://gotriangle.tripsparkhost.com/gtfs/Realtime/GTFS_VehiclePositions.pb",
	ExcludeList:    []string{},
	Exclude:        []string{},
	Updates: model.Updates{
		Stops:        func() {},
		MissingTrips: func() {},
		Alerts:       func() {},
	},
}
