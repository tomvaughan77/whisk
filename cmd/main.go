package main

import "whisk/pkg/schedule"

func main() {
	schedule.Schedule(schedule.Year{}, 0)
}
