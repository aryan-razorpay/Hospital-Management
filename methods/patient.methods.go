package methods

import (
	"hospital-management/config"
	"hospital-management/models"
)

type UserError struct {
    Message string
}

func (e *UserError) Error() string {
    return e.Message
}
func CreatePatient(patient *models.Patient) error {
	return config.DB.Create(patient).Error
}

func GetPatientByID(id string) (*models.Patient, error) {
	var patient models.Patient
	err := config.DB.First(&patient, "id = ?", id).Error
	return &patient, err
}

func UpdatePatientDetails(id string, contactNo, address, doctorID string) error {
	var doctor models.Doctor
	err := config.DB.First(&doctor, "id = ?", doctorID).Error
	if err != nil {
		return &UserError{Message: "Please enter a valid doctor id"}
	}
	return config.DB.Model(&models.Patient{}).Where("id = ?", id).Updates(models.Patient{
		ContactNo: contactNo,
		Address:   address,
		DoctorID:  doctorID,
	}).Error
}

func GetPatientsByDoctorID(doctorID string) ([]models.Patient, error) {
	var patients []models.Patient
	err := config.DB.Where("doctor_id = ?", doctorID).Find(&patients).Error
	return patients, err
}
