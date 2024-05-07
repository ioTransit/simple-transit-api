package model

type Updates struct {
	Stops        func()
	MissingTrips func()
	Alerts       func()
}

type AgencyConfig struct {
	Id   string
	Name string
	Url  string
}
