package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"googlemaps.github.io/maps"
)

func main() {
	api_key := "AIzaSyBC6tf0jSZTjZ7BbJWWy8ZOzrKFraReqkk"
	c, err := maps.NewClient(maps.WithAPIKey(api_key))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	//	r := &maps.DirectionsRequest{
	r := &maps.DistanceMatrixRequest{
		Origins:      []string{"Dammam"},
		Destinations: []string{"Riyadh"},
		ArrivalTime:  (time.Date(1, time.September, 1, 0, 0, 0, 0, time.UTC)).String(),
	}

	//	Language:      *language,
	//	DepartureTime: *departureTime,
	//	ArrivalTime:   *arrivalTime,

	resp, err := c.DistanceMatrix(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	//pretty.Println(resp)
	x, y := resp.Rows[0].Elements[0].Duration, resp.Rows[0].Elements[0].Distance.HumanReadable
	fmt.Println("The total time is:", x, "and the total distance is:", y)

	/*	route, _, err := c.Directions(context.Background(), r)
		if err != nil {
			log.Fatalf("fatal error: %s", err)
		}

		pretty.Println(route) */
}
