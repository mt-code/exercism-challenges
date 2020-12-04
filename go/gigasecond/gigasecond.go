package gigasecond

import "time"

// Added a gigasecond (1,000,000,000 seconds) to the time provided.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
