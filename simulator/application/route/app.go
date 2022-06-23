package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

type Route struct {
	ID        string     `json:"routeID"`
	ClientID  string     `json:"clientID"`
	Positions []Position `json:"positions"`
}

type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PartialRoutePosition struct {
	ID       string    `json:"routeID"`
	ClientID string    `json:"clientID"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

func (r *Route) parseLocationDataToFloat(line []string) (lat, long float64, err error) {
	lat, err = strconv.ParseFloat(line[0], 64)
	if err != nil {
		return lat, long, errors.New("failed to parse latitude: " + err.Error())
	}

	long, err = strconv.ParseFloat(line[1], 64)
	if err != nil {
		return lat, long, errors.New("failed to parse longitude: " + err.Error())
	}

	return
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("route ID was not supplied")
	}

	var f, err = os.Open(path.Join("simulator", "destinations", r.ID+".txt"))
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	var fileScanner = bufio.NewScanner(f)

	const positionSeparator = ","
	for fileScanner.Scan() {
		var (
			positionsLine               = strings.Split(strings.Trim(fileScanner.Text(), " "), positionSeparator)
			latitudeLine, longitudeLine float64
		)

		latitudeLine, longitudeLine, err = r.parseLocationDataToFloat(positionsLine)
		if err != nil {
			return err
		}

		r.Positions = append(r.Positions, Position{
			Latitude:  latitudeLine,
			Longitude: longitudeLine,
		})
	}

	return nil
}

func (r *Route) ExportDataToJSON() ([]string, error) {
	var (
		partialRoutePosition = PartialRoutePosition{
			ID:       r.ID,
			ClientID: r.ClientID,
		}

		positions = make([]string, 0, len(r.Positions))
	)

	for i := range r.Positions {
		if i == (len(r.Positions) - 1) {
			partialRoutePosition.Finished = true
		}

		partialRoutePosition.Position = []float64{
			r.Positions[i].Latitude,
			r.Positions[i].Longitude,
		}

		var routePartialJSON, err = json.Marshal(&partialRoutePosition)
		if err != nil {
			log.Fatalln("failed to parse partial route data to JSON:", err.Error())
		}

		positions = append(positions, string(routePartialJSON))
	}

	return positions, nil
}
