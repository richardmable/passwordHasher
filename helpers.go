package main

import (
	"fmt"
	"os"
	"time"
)

// error helper
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// time how long a function takes
func timeToProcess(start time.Time, stats *Stats) time.Duration {
	elapsed := time.Since(start)
	fmt.Printf("took %s\n", elapsed)
	return elapsed
}

// figure out the avg time by dividing total time of all requests by total request number
func avgTimeToProcess(timeTotal time.Duration, reqTotal int64) int64 {
	// divide and convert to microseconds
	averageTimeUs := int64(time.Duration(int64(timeTotal)/reqTotal) / time.Microsecond)
	return averageTimeUs
}

func statsGenerator(startTime time.Time, stats *Stats) *Stats {
	// check how long process took
	elaspsed := timeToProcess(startTime, stats)
	// add one to the stats count
	stats.TotalRequests += 1
	// add time elapsed for this most recent request to total time
	stats.TotalProccessTime += elaspsed
	// caculate new average time
	stats.AverageProccessTime = avgTimeToProcess(stats.TotalProccessTime, stats.TotalRequests)
	return stats
}
