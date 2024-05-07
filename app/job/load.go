package jobs

import (
	"database/sql"
	"fmt"
	"go-gtfs-server/app/model"
	"go-gtfs-server/db"
	"go-gtfs-server/providers"
	"go-gtfs-server/utils"
	"io"
	"net/http"
	"os"

	"github.com/geops/gtfsparser"
	_ "github.com/lib/pq"
)

var PostGres *sql.DB

var insertStatement = "INSERT INTO $1 ($2) VALUES ($3) ON CONFLICT DO NOTHING"

func downloadFile(url string, filePath string) error {
	// Create the file
	output, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer output.Close()

	// Get the data from the URL
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error getting response: %v", err)
	}
	defer response.Body.Close()

	// Check if response is successful
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %s", response.Status)
	}

	// Write the body to file
	_, err = io.Copy(output, response.Body)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

type AgencyRow struct {
	TcAgencyId string
	Phone      string
	Url        string
	Lang       string
	Name       string
	Timezone   string
	FareUrl    string
	AgencyId   string
}

func calendar_dates(agencyConfig model.AgencyConfig) {
}

func routes(agencyConfig model.AgencyConfig) {}

func shapes(agencyConfig model.AgencyConfig) {}

func trips(agencyConfig model.AgencyConfig) {}

func stop_times(agencyConfig model.AgencyConfig) {}

func arrival_times(agencyConfig model.AgencyConfig) {}

func departure_times(agencyConfig model.AgencyConfig) {}

func end_time(agencyConfig model.AgencyConfig) {}

func stops(agencyConfig model.AgencyConfig) {}

func transfers(agencyConfig model.AgencyConfig) {}

func initFeed(agencyConfig model.AgencyConfig) *gtfsparser.Feed {
	file := fmt.Sprintf("gtfs/%s.zip", agencyConfig.Id)
	Args := os.Args[1:]
	update := utils.ContainsString(Args, "--update")
	if update {
		downloadFile(agencyConfig.Url, file)
		fmt.Printf("Downloaded %s gtfs", agencyConfig.Name)
	}

	Feed := gtfsparser.NewFeed()
	Feed.Parse(file)

	fmt.Printf("%s's GTFS Loaded 🚂 \n", agencyConfig.Name)
	fmt.Printf("Done, parsed %d agencies, %d stops, %d routes, %d trips, %d fare attributes\n\n",
		len(Feed.Agencies), len(Feed.Stops), len(Feed.Routes), len(Feed.Trips), len(Feed.FareAttributes))

	return Feed
}

func Load() {
	for _, Agency := range providers.ProvidersArray {
		Feed := initFeed(Agency)
		db.GtfsStore[Agency.Id] = db.AgencyStore{Gtfs: Feed, Config: Agency}
	}
}
