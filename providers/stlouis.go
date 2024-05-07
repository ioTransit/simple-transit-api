package providers

import "go-gtfs-server/app/model"

var StLouisConfig = model.AgencyConfig{
	Id:   "stlouis",
	Name: "Saint Louis Metro",
	Url:  "https://www.metrostlouis.org/Transit/google_transit.zip",
}
