package methods

import (
	"hospital-management/config"
	"hospital-management/models"
)

func CreateDoctor(doctor *models.Doctor) error {
	return config.DB.Create(doctor).Error
}

func GetDoctorByID(id string) (*models.Doctor, error) {
	var doctor models.Doctor
	err := config.DB.First(&doctor, "id = ?", id).Error
	return &doctor, err
}

func UpdateDoctorContact(id string, contactNo string) error {
	return config.DB.Model(&models.Doctor{}).Where("id = ?", id).Update("contact_no", contactNo).Error
}
