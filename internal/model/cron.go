package model

// cron entry struct
type CronEntry struct {
	ID       string `json:"id"`
	Next     string `json:"next"`
	Prev     string `json:"prev"`
	Schedule string `json:"schedule"`
}
