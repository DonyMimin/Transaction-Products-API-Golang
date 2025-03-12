package utils

import "time"

// GetCurrentLocalTime generate current local time (Asia/Jakarta)
func GetCurrentLocalTime() (now time.Time, err error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return
	}
	now = time.Now().In(location)
	return
}
