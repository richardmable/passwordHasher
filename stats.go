package main

import (
	"time"
)

type Stats struct {
	TotalRequests       int64         `json:"total"`
	AverageProccessTime int64         `json:"average"`
	TotalProccessTime   time.Duration `json:"-"`
}
