package grpcservice

import (
	"context"
	"patient/proto/gcommon"
	"patient/proto/patient"
)

type (
	PatientGrpcService interface {
		GetPatient(ctx context.Context, request *gcommon.IdRequest) (*patient.GetPatientResponse, error)
		ListPatient(ctx context.Context, request *patient.ListPatientRequest) (*patient.ListPatientResponse, error)
		UpsertPatient(ctx context.Context, request *patient.UpsertPatientRequest) (*gcommon.EmptyResponse, error)
		UpsertMedicalRecord(ctx context.Context, request *patient.UpsertMedicalRecordRequest) (*gcommon.EmptyResponse, error)
		UpsertMedicalTreatment(ctx context.Context, request *patient.UpsertMedicalTreatmentRequest) (*gcommon.EmptyResponse, error)
		UpsertMedicalSurgery(ctx context.Context, request *patient.UpsertMedicalSurgeryRequest) (*gcommon.EmptyResponse, error)
		UpsertMedicalPrescription(ctx context.Context, request *patient.UpsertMedicalPrescriptionRequest) (*gcommon.EmptyResponse, error)
		GetMedicalHistory(ctx context.Context, request *patient.GetMedicalHistoryRequest) (*patient.GetMedicalHistoryResponse, error)
		GetMedicalHistoryDetail(ctx context.Context, request *gcommon.IdRequest) (*patient.GetMedicalHistoryDetailResponse, error)
	}

	MedicationGrpcService interface {
		UpsertMedication(ctx context.Context, request *patient.UpsertMedicationRequest) (*gcommon.EmptyResponse, error)
		ListMedication(ctx context.Context, request *patient.ListMedicationRequest) (*patient.ListMedicationResponse, error)
	}
)
