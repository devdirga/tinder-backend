package util

import "time"

func GetNow() time.Time {
	// utc := time.FixedZone("UTC+7", 7*60*60)
	return time.Now().In(time.FixedZone("UTC", 7*60*60)).UTC()
}
