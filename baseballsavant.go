package baseballsavant

import (
	"fmt"
	"log"
	"reflect"
)

// Client is the base type for retrieving and ingesting baseball savant data
type Client struct {
}

// Filters that can be used in the URL
// arrays  are seperated by | (%7C) in the url
type Filter struct {
	All             bool      `filter:"all"`  // I don't really know what the all is but in the python project I found it was true
	HfPT            *[]string `filter:"hfPT"` // Pitch Type
	HfAB            *string   `filter:"hfAB"`
	HfBBT           *string   `filter:"hfBBT"`
	HfPR            *[]string `filter:"hfPR"` // Pitch Result
	HfZ             *string   `filter:"hfZ"`
	Stadium         *string   `filter:"stadium"`
	HfBBL           *[]string `filter:"hfBBL"` // Batted Ball Location
	HfNewZones      *string   `filter:"hfNewZones"`
	HfGT            *string   `filter:"hfGT"`
	HfC             []string  `filter:"hfC"`   // Count 0-0 is 00
	HfSea           []string  `filter:"hfSea"` // Season
	HfSit           *string   `filter:"hfSit"`
	PlayerType      *string   `filter:"player_type"`
	HfOuts          *string   `filter:"hfOuts"`
	Opponent        *string   `filter:"opponent"`
	PitcherThrows   *string   `filter:"pitcher_throws"`
	BatterStands    *string   `filter:"batter_stands"`
	HfSA            *string   `filter:"hfSA"`
	GameDateGt      *string   `filter:"game_date_gt"` // Game Date Greater than
	GameDateLt      *string   `filter:"game_date_lt"` // Game Date Less than
	Team            *string   `filter:"team"`         // Team
	Position        *string   `filter:"position"`     // Position
	HfRO            *[]string `filter:"hfRO"`         // Runners on
	HomeRoad        string    `filter:"home_road"`
	HfFlag          *[]string `filter:"hfFlag"`   // Flags
	Metric1         *string   `filter:"metric_1"` // Metric Range?
	HfInn           *[]string `filter:"hfInn"`    // Inning
	MinPitches      *string   `filter:"min_pitches"`
	MinResults      *string   `filter:"min_results"`
	GroupBy         *string   `filter:"group_by"` // Group By
	SortCol         *string   `filter:"sort_col"`
	PlayerEventSort *string   `filter:"player_event_sort"`
	SortOrder       *string   `filter:"sort_order"`
	MinAbs          *string   `filter:"min_abs"`
	Type            string    `filter:"type"` // Player Type
}

// Get a csv from baseball savant
func (c *Client) Get() (records [][]string) {

	// 	url := fmt.Sprintf("https://usr53-services.dayforcehcm.com/api/pccmarkets/v1/Employees?EmploymentStatusXRefCode=Active&FilterUpdatedStartDate=%s&FilterUpdatedEndDate=%s", formattedStartDate, formattedEndDate)

	// // create the http client
	// client := &http.Client{}

	// // create the request
	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	return employees, err
	// }

	return
}

// URL builds and returns the URL from a filters
func (f Filter) URL() string {

	ty := reflect.TypeOf(f)

	va := reflect.ValueOf(f)

	length := ty.NumField()

	var urlParams []string

	for i := 0; i < length; i++ {
		filter := ty.Field(i).Tag.Get("filter")
		name := ty.Field(i).Name
		switch va.FieldByName(name).Kind().String() {
		case "ptr":
			continue
		case "slice":

		default:
			part := fmt.Sprintf("%s=%s", filter, va.FieldByName(name).String())
			urlParams = append(urlParams, part)
			log.Println(filter, va.FieldByName(name).String())
		}
		// if inter != "" && inter != "<nil>" {
		// 	log.Println(name, filter, va.FieldByName(name), va.FieldByName(name).Kind())
		// }

	}

	return "something"
}
