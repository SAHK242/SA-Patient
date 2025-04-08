package grpcvalidator

import "go.uber.org/fx"

var Module = fx.Provide(
	NewPatientGrpcValidator,
	NewMedicationValidator,
	NewMedicalSurgeryValidator,
	NewMedicalTreatmentValidator,
	NewMedicalPrescriptionValidator,
	NewMedicalHistoryValidator,
)
