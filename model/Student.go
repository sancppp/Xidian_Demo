package model

type Student struct {
	ID         int    `gorm:"primaryKey" json:"student_id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Major      string `json:"major"`
}
