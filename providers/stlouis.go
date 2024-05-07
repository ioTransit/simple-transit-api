package providers

import "go-gtfs-server/app/model"

var StLouisConfig = model.AgencyConfig{
	Id:             "stlouis",
	Name:           "Saint Louis Metro",
	Url:            "https://www.metrostlouis.org/Transit/google_transit.zip",
	TripUpdates:    "https://www.metrostlouis.org/RealTimeData/StlRealTimeTrips.pb",
	AlertUpdates:   "https://www.metrostlouis.org/RealTimeData/StlRealTimeAlerts.pb",
	VehicleUpdates: "https://www.metrostlouis.org/RealTimeData/StlRealTimeVehicles.pb",
	ExcludeList:    []string{},
	Exclude:        []string{},
	Updates: model.Updates{
		Stops:        func() {},
		MissingTrips: func() {},
		Alerts:       func() {},
	},
}
