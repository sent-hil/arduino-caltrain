package caltrain

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Station string
type Direction int

const (
	// direction of trains
	SouthBound Direction = 0
	NorthBound Direction = 1

	// list of train strations as of June 1, 2016
	SanFrancisco          Station = "sanfrancisco"
	TwentySecondStreet    Station = "22ndstreet"
	Bayshore              Station = "bayshore"
	SouthSanFrancisco     Station = "southsanfrancisco"
	SanBruno              Station = "sanbruno"
	MillbraeTransitCenter Station = "millbraetransitcenter"
	Broadway              Station = "broadway"
	Burlingame            Station = "burlingame"
	SanMateo              Station = "sanmateo"
	HaywardPark           Station = "haywardpark"
	Hillsdale             Station = "hillsdale"
	Belmont               Station = "belmont"
	SanCarlos             Station = "sancarlos"
	RedwoodCity           Station = "redwoodcity"
	Atherton              Station = "atherton"
	MenloPark             Station = "menlopark"
	PaloAlto              Station = "paloalto"
	CaliforniaAve         Station = "californiaave"
	SanAntonio            Station = "sanantonio"
	MountainView          Station = "mountainview"
	Sunnyvale             Station = "sunnyvale"
	Lawrence              Station = "lawrence"
	SantaClara            Station = "santaclara"
	CollegePark           Station = "collegepark"
	SanJoseDiridon        Station = "sanjosediridon"
	Tamien                Station = "tamien"
	Capitol               Station = "capitol"
	BlossomHill           Station = "blossomhill"
	MorganHill            Station = "morganhill"
	SanMartin             Station = "sanmartin"
	Gilroy                Station = "gilroy"

	// baseURL is the common url between station realtime web pages.
	baseURL = "http://www.caltrain.com/schedules/realtime/stations/%sstation-mobile.html"

	timingSuffix = " min."

	directionSelector = ".ipf-st-ip-trains-subtable"
	timingSelector    = ".ipf-st-ip-trains-subtable-td-arrivaltime"
)

// GetRealTimings returns duration of latest trains arriving at the specified
// station and going the specified direction. The Caltrain realtime page
// provides maximum of 3 upcoming trains.
//
// See http://www.caltrain.com/schedules/realtime/stations.html for list of
// stations.
//
// See http://www.caltrain.com/schedules.html if you want scheduled timetables.
//
// TODO: incorporate alerts which are embedded in the page.
func GetRealTimings(s Station, d Direction) (timings []time.Duration, err error) {
	doc, err := goquery.NewDocument(fmt.Sprintf(baseURL, s))
	if err != nil {
		return nil, err
	}

	// Caltrain realtime page has the data we want in tables:
	//
	//    <table class="ipf-st-ip-trains-subtable">
	//       <tr class="ipf-st-ip-trains-subtable-tr">
	//          ...
	//          <td class="ipf-st-ip-trains-subtable-td-arrivaltime">30 min.</td>
	//       </tr>
	//    </table>
	//    <table class="ipf-st-ip-trains-subtable">
	//       <tr class="ipf-st-ip-trains-subtable-tr">
	//          ...
	//          <td class="ipf-st-ip-trains-subtable-td-arrivaltime">10 min.</td>
	//       </tr>
	//    </table>
	//
	// Since the headers are in a different row than the data, we assume the first
	// table with class "ipf-st-ip-trains-subtable" contains SouthBound timings
	// and the last one contains NorthBound.
	//
	i := int(d)

	// get the approriate (0 or 1) table
	doc.Find(directionSelector).Eq(i).Each(func(_ int, s1 *goquery.Selection) {

		// iterate through arrival timings
		s1.Find(timingSelector).Each(func(_ int, s2 *goquery.Selection) {
			if t, err := parseStrIntoTime(s2.Text()); err == nil {
				timings = append(timings, t)
			}
		})
	})

	return timings, nil
}

func parseStrIntoTime(str string) (time.Duration, error) {
	raw := strings.TrimSuffix(str, timingSuffix)
	min, err := strconv.Atoi(raw)
	if err != nil {
		return 0 * time.Minute, err
	}

	return time.Duration(min) * time.Minute, nil
}
