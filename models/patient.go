package models

import "time"

type Patient struct {
	ID        string    `gorm:"primaryKey;type:char(5);unique" json:"id"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
	Name      string    `gorm:"type:varchar(255); not null" json:"name"`
	ContactNo string    `gorm:"type:char(10) not null" json:"contact_no"`
	Address   string    `gorm:"type:varchar(255) not null" json:"address"`
	DoctorID  string    `gorm:"type:char(5) not null" json:"doctor_id"`
}
