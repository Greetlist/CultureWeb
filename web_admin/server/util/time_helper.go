package util

import (
    "time"
)

// return today str as %Y-%M-%D
func GetTodayStr() string {
    now := time.Now()
    return now.Format("2006-01-02")
}
