# Simple Transit API
A simple transit api written in Go with [gtfsparser](https://www.github.com/geops) and [Echo](https://echo.labstack.com/)

## Quickstart
There are 2 agencies already setup out of the box that you can take a look at in the providers folder. Each agency will need to be setup with the unique id, name and the gtfs url.

```golang
package providers

import "go-gtfs-server/app/model"

var StLouisConfig = model.AgencyConfig{
	Id:             "stlouis",
	Name:           "Saint Louis Metro",
	Url:            "https://www.metrostlouis.org/Transit/google_transit.zip",
}
```
### On Start
When the server is first started it will load all of the gtfs files that are in the listed in the `providers/all.go` file. Once those are loaded they are stored with their id's as unique identifiers in memory and accessible via the rest api


	// Start of Routes

	// Start of Trips

### On Update 
***IN PROGRESS***

### Stops
- `/stops/:agencyId/` - Returns all of the stops by agency id
- `/stops/:agencyId/?stop_name=:<query>`- Returns all of the stops by agency id that start with the query
- `/stops/:agencyId/:stopId/` - The stop from the agency with that stop_id

### Routes
- `/routes/:agencyId/` - Returns all of the routes by agency id
- `/routes/:agencyId/?route_long_name=:<query>`- Returns all of the routes by agency id that start with the query
- `/routes/:agencyId/:routeId/` - The route from the agency with that route_id

### Trips
- `/trips/:agencyId/:tripId/` - Returns all of the trips by agency id
- `/trips/:agencyId/route/:routeId/`  - The trips from the agency with that trip_id 

