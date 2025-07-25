package main

import (
	"fmt"
	"log"
	"math"
	"net"
	"time"

	"github.com/distatus/battery"
)

func main() {

	ip := GetOutboundIP()
	for {
		date := getDateAndTime()
		battery := batterySection()

		fmt.Printf("%s 󰤃 %s 󰤃 %s ", battery, ip, date)
		time.Sleep(10 * time.Second)

	}
}

func batterySection() string {
	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Could not get battery info!")
		return ""
	}
	var percent float64
	var icon string
	var state string
	var state_ico string

	for _, battery := range batteries {
		// 󰚥
		percent, icon = calculatePercent(battery)
		state = battery.State.String()
		if state == "Discharging" {
			state_ico = "󱐤"

		} else if state == "Charging" {
			state_ico = "󰚥"

		}
		// fmt.Print("%\v+", battery.State)
		// fmt.Printf("Bat%d: ", i)
		// fmt.Printf("state: %s, ", battery.State.String())
		// fmt.Printf("current capacity: %f mWh, ", battery.Current)
		// fmt.Printf("last full capacity: %f mWh, ", battery.Full)
		// fmt.Printf("percent %f %s", percent, icon)
		// fmt.Printf("design capacity: %f mWh, ", battery.Design)
		// fmt.Printf("charge rate: %f mW, ", battery.ChargeRate)
		// fmt.Printf("voltage: %f V, ", battery.Voltage)
		// fmt.Printf("design voltage: %f V\n", battery.DesignVoltage)

	}
	return fmt.Sprintf(" %d %s %s ", int(percent*100), icon, state_ico)
}

func calculatePercent(b *battery.Battery) (float64, string) {
	percent := b.Current / b.Full
	rounded := roundToNearestTen(percent)
	icon := mapIcon(rounded)

	return percent, icon
}

func mapIcon(percent int) string {

	//󰂄 charging \udb80\udc84
	bstatus := map[int]map[string]string{
		90: {"icon": "󰂂"},
		80: {"icon": "󰂁"},
		70: {"icon": "󰂀"},
		60: {"icon": "󰁿"},
		50: {"icon": "󰁾"},
		40: {"icon": "󰁽"},
		30: {"icon": "󰁼"},
		20: {"icon": "󰁻"},
		10: {"icon": "󰁺"},
	}

	icon := bstatus[percent]["icon"]

	return icon

}

func roundToNearestTen(n float64) int {
	divided := n * 10.0
	rounded := math.Round(divided)
	result := rounded * 10.0

	return int(result)
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func getDateAndTime() string {
	t := time.Now()
	return t.Format(time.Stamp)
}

// \\udb80\\udc82
// \\udb80\\udc81
// \\udb80\\udc80
// \\udb80\\udc7f
// \\udb80\\udc7e
// \\udb80\\udc7d
// \\udb80\\udc7c
// \\udb80\\udc7b
// \\udb80\\udc7a
