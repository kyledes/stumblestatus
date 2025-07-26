package cpustatus

import (
	"errors"
	"testing"
	"time"
)

// Mock function to replace cpuPercentFunc during tests
func mockCPUPercent(interval time.Duration, percpu bool) ([]float64, error) {
	if interval != time.Second {
		return nil, errors.New("unexpected interval")
	}
	if percpu {
		return nil, errors.New("unexpected percpu true")
	}
	// Return a fixed CPU load value for testing
	return []float64{42.5}, nil
}

func TestGetCPULoad(t *testing.T) {
	// Save original function and restore after test
	origFunc := cpuPercentFunc
	defer func() { cpuPercentFunc = origFunc }()

	// Replace with mock
	cpuPercentFunc = mockCPUPercent

	load, err := GetCPULoad()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	expected := 42.5
	if load != expected {
		t.Errorf("Expected CPU load %.1f, got %.1f", expected, load)
	}
}

func TestGetCPULoad_Error(t *testing.T) {
	origFunc := cpuPercentFunc
	defer func() { cpuPercentFunc = origFunc }()

	// Mock returns error
	cpuPercentFunc = func(interval time.Duration, percpu bool) ([]float64, error) {
		return nil, errors.New("mock error")
	}

	_, err := GetCPULoad()
	if err == nil {
		t.Fatal("Expected error but got nil")
	}
}
