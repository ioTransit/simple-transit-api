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

