package main

import (
	"fmt"
	"math"
	"unicode/utf16"

	"github.com/distatus/battery"
)

func main() {
	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Could not get battery info!")
		return
	}
	for i, battery := range batteries {
		// 󰚥
		percent, icon := calculatePercent(battery)
		fmt.Print("%\v+", battery.State)
		fmt.Printf("Bat%d: ", i)
		fmt.Printf("state: %s, ", battery.State.String())
		fmt.Printf("current capacity: %f mWh, ", battery.Current)
		fmt.Printf("last full capacity: %f mWh, ", battery.Full)
		fmt.Printf("percent %f %U", percent, icon)
		fmt.Printf("design capacity: %f mWh, ", battery.Design)
		fmt.Printf("charge rate: %f mW, ", battery.ChargeRate)
		fmt.Printf("voltage: %f V, ", battery.Voltage)
		fmt.Printf("design voltage: %f V\n", battery.DesignVoltage)
	}
}

func calculatePercent(b *battery.Battery) (float64, []uint16) {
	percent := b.Current / b.Full
	rounded := roundToNearestTen(percent)
	icon := mapIcon(rounded)

	return percent, icon
}

func mapIcon(percent int) []uint16 {

	//󰂄 charging \udb80\udc84
	bstatus := map[int]map[string][]rune{
		90: {"icon": []rune("󰂂")},
		80: {"icon": []rune("󰂁")},
		70: {"icon": []rune("󰂀")},
		60: {"icon": []rune("󰁿")},
		50: {"icon": []rune("󰁾")},
		40: {"icon": []rune("󰁽")},
		30: {"icon": []rune("󰁼")},
		20: {"icon": []rune("󰁻")},
		10: {"icon": []rune("󰁺")},
	}

	fmt.Printf("rounded %d", percent)
	icon := utf16.Encode(bstatus[percent]["icon"])

	fmt.Printf("icon: %U", icon)
	return icon

}

func roundToNearestTen(n float64) int {
	divided := n / 10.0
	rounded := math.Round(divided)
	result := rounded * 10.0
	return int(result)
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
