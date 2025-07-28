package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/kyle/stumblestatus/internal/battery"
	"github.com/kyle/stumblestatus/internal/cpustatus"
)

func main() {
	var statusElements []string
	includeBattery := battery.IncludeBattery()
	ip := GetOutboundIP()

	cpuload, _ := cpustatus.GetCPULoad()
	for {
		date := getDateAndTime()
		if includeBattery {
			battery := battery.BatterySection()

			statusElements = append(statusElements, battery)

		}
		statusElements = append(statusElements, fmt.Sprintf("%.2f", cpuload))
		statusElements = append(statusElements, ip.String())
		statusElements = append(statusElements, date)
		statusLine := strings.Join(statusElements, " 󰤃 ")

		fmt.Printf("%s ", statusLine)
		time.Sleep(10 * time.Second)
		statusElements = nil
	}
}

//	func main() {
//		load, err := GetCPULoad()
//		if err != nil {
//			fmt.Println("Error:", err)
//			return
//		}
//		fmt.Printf("Current total CPU load: %.2f%%\n", load)
//	}
func cpu() {

	cpustatus.GetCPULoad()
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
