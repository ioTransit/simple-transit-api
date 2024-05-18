package main

import (
	"context"
	jobs "go-gtfs-server/app/job"
	"go-gtfs-server/app/pages"
	"go-gtfs-server/app/view"
	"go-gtfs-server/cli"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error with .env")
	}

	cli.CliRouter()
	c := jobs.UpdateGtfs()

	e := echo.New()
	e.Static("/js", "js")
	e.Static("/styles", "styles")

	e.GET("/", helloRender)

	// Start of Stops
	e.GET("/api/stops/:agencyId/", view.GetStopsByAgency)
	e.GET("/api/stops/:agencyId/:stopId/", view.GetStopByAgencyAndStopId)
	e.GET("/stops/:agencyId/:stopId/", view.GetStopByAgencyAndStopId)

	// Start of Routes
	e.GET("/api/routes/:agencyId/", view.GetRoutesByAgency)
	e.GET("/api/routes/:agencyId/:routeId/", view.GetRouteByAgencyAndRouteId)

	// Start of Trips
	e.GET("/api/trips/:agencyId/:tripId/", view.GetTripByAgencyAndTripId)
	e.GET("/api/trips/:agencyId/route/:routeId/", view.GetTripsByAgencyAndRouteId)

	// Start server in a goroutine
	go func() {
		if err := e.Start(":1080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	// Context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	c.Stop()
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func helloRender(c echo.Context) error {
	return Render(c, http.StatusOK, pages.Hello())
}
