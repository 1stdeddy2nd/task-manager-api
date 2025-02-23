package models

type Task struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Status string `json:"status" gorm:"default:pending;check:status IN ('pending', 'done')"`
	UserID uint   `json:"user_id"`
}
