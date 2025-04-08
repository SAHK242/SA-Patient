package grpcmapper

import "go.uber.org/fx"

var Module = fx.Provide(
	NewPatientGrpcMapper,
	NewMedicationGrpcMapper,
	NewMedicalTreatmentGrpcMapper,
	NewMedicalSurgeryGrpcMapper,
	NewMedicalPrescriptionGrpcMapper,
	NewMedicalHistoryGrpcMapper,
)
