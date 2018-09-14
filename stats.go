package main

import (
	"time"
)

type Stats struct {
	TotalRequests       int64         `json:"totalRequests"`
	AverageProccessTime int64         `json:"averageProccessTime"`
	TotalProccessTime   time.Duration `json:"totalProcessTime"`
}
