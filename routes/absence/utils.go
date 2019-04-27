package absence

import "time"

func dbTime(data []byte) string {
	dataTime, err := time.Parse(time.RFC3339, string(data))
	if err != nil {
		return ""
	}
	return dataTime.Format("2006-01-02 15:04:05")
}
