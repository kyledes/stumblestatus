package cpustatus

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// Define a type for the function signature
type cpuPercentFuncType func(interval time.Duration, percpu bool) ([]float64, error)

// Assign the actual function by default
var cpuPercentFunc cpuPercentFuncType = cpu.Percent

// GetCPULoad retrieves the current total CPU load percentage averaged over 1 second.
func GetCPULoad() (float64, error) {
	percentages, err := cpuPercentFunc(time.Second, false)
	if err != nil {
		return 0, err
	}
	if len(percentages) == 0 {
		return 0, fmt.Errorf("no CPU usage info retrieved")
	}
	return percentages[0], nil
}
