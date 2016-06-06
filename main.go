package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/huin/goserial"
	"github.com/sent-hil/caltrain-realtime"
)

const (
	// default settings for Arduino UNO
	deviceName = "/dev/tty.usbmodem1421"
	deviceBaud = 9600

	interval = 10 * time.Second
)

var (
	// certain characters aren't printed properly, so we replace them with
	// printable characters
	unprintables = map[string]string{
		"o": "0",
		"O": "0",
	}
)

func main() {
	for _ = range time.Tick(interval) {
		msg, err := getTimings()
		if err != nil {
			sendMessage(fmt.Sprintf("Error: %s", err))
			log.Fatal(err)
		}

		if err := sendMessage(msg); err != nil {
			log.Fatal(err)
		}
	}
}

func getTimings() (string, error) {
	timings, err := caltrain.GetRealTimings(caltrain.PaloAlto, caltrain.NorthBound)
	if err != nil {
		return "", err
	}

	o := []string{}
	for _, t := range timings {
		o = append(o, t.String())
	}

	if len(o) == 0 {
		o = append(o, "Empty")
	}

	return strings.Join(o, ","), nil
}

func sendMessage(msg string) error {
	c := &goserial.Config{Name: deviceName, Baud: deviceBaud}
	s, err := goserial.OpenPort(c)
	if err != nil {
		return err
	}

	defer s.Close()

	// !important! wait for bootstraper to initialize
	time.Sleep(2 * time.Second)

	_, err = s.Write([]byte(removeUnprintables(msg)))
	return err
}

func removeUnprintables(s string) string {
	for k, v := range unprintables {
		s = strings.Replace(s, k, v, -1)
	}

	return s
}
