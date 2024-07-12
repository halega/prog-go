package parsedate

import "time"

func parse(date string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05.999Z", date)
}

func duration(start, end time.Time) time.Duration {
	return end.Sub(start)
}
