package models

type JobLevel struct {
	JobLevelID          int    `json:"job_level_id"`
	JobLevelDescription string `json:"job_level_desc"`
	IsActive            bool   `json:"is_active"`
}
