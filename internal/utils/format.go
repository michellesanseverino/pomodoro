package format

import (
	"fmt"
	"time"
)

// FormatTime formats a time.Time object to a string
func formatTime(t time.Time, layout string) string {
	return t.Format(layout)
}

// FormatDate formats a time.Time object to a string with the layout "2025-01-02"
func formatDate(t time.Time) string {
	return formatTime(t, "2025-01-02")
}

func formatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
