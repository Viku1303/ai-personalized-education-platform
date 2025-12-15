package models

import (
	"time"
)

// Course represents an educational course
type Course struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Title        string    `json:"title" gorm:"not null"`
	Description  string    `json:"description"`
	InstructorID uint      `json:"instructor_id"`
	Instructor   User      `json:"instructor" gorm:"foreignKey:InstructorID"`
	Lessons      []Lesson  `json:"lessons" gorm:"foreignKey:CourseID"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Lesson represents a lesson within a course
type Lesson struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CourseID  uint      `json:"course_id" gorm:"not null"`
	Course    Course    `json:"course" gorm:"foreignKey:CourseID"`
	Title     string    `json:"title" gorm:"not null"`
	Content   string    `json:"content"`
	Order     int       `json:"order" gorm:"default:0"`
	Duration  int       `json:"duration"` // Duration in minutes
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Enrollment represents a student's enrollment in a course
type Enrollment struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	UserID      uint       `json:"user_id" gorm:"not null"`
	User        User       `json:"user" gorm:"foreignKey:UserID"`
	CourseID    uint       `json:"course_id" gorm:"not null"`
	Course      Course     `json:"course" gorm:"foreignKey:CourseID"`
	Progress    int        `json:"progress" gorm:"default:0"` // Progress percentage
	EnrolledAt  time.Time  `json:"enrolled_at"`
	CompletedAt *time.Time `json:"completed_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
