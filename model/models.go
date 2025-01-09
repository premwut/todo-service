package model

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"size:50"`
	LastName  string `gorm:"size:50"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Project struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	UserID    uint `gorm:"foreignKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Status string

const (
	Created    Status = "created"
	InProgress Status = "in_progress"
	Completed  Status = "completed"
	Rejected   Status = "rejected"
)

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255"`
	Description string `gorm:"size:255"`
	ProjectID   uint   `gorm:"foreignKey"`
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
