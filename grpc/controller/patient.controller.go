package controller

import (
	"context"
	grpcservice "patient/grpc/service"
	util "patient/grpc/util"
	"patient/proto/gcommon"
	"patient/proto/patient"
)

type PatientGrpcController struct {
	patient.UnimplementedPatientServiceServer
	patientService grpcservice.PatientGrpcService
}

func NewPatientGrpcController(
	patientService grpcservice.PatientGrpcService,
) *PatientGrpcController {
	return &PatientGrpcController{
		patientService: patientService,
	}
}

func (c *PatientGrpcController) UpsertPatient(ctx context.Context, req *patient.UpsertPatientRequest) (*gcommon.EmptyResponse, error) {
	return util.WithSafeErrBiFunction(ctx, req, c.patientService.UpsertPatient)
}

func (c *PatientGrpcController) GetPatient(ctx context.Context, req *gcommon.IdRequest) (*patient.GetPatientResponse, error) {
	return util.WithSafeErrBiFunction(ctx, req, c.patientService.GetPatient)
}

func (c *PatientGrpcController) ListPatient(ctx context.Context, req *patient.ListPatientRequest) (*patient.ListPatientResponse, error) {
	return util.WithSafeErrBiFunction(ctx, req, c.patientService.ListPatient)
}

func (c *PatientGrpcController) GetMedicalHistory(ctx context.Context, req *patient.GetMedicalHistoryRequest) (*patient.GetMedicalHistoryResponse, error) {
	return util.WithSafeErrBiFunction(ctx, req, c.patientService.GetMedicalHistory)
}

func (c *PatientGrpcController) UpsertMedicalRecord(ctx context.Context, req *patient.UpsertMedicalRecordRequest) (*gcommon.EmptyResponse, error) {
	return util.WithSafeErrBiFunction(ctx, req, c.patientService.UpsertMedicalRecord)
}

func (c *PatientGrpcController) UpsertMedicalTreatment(ctx context.Context, req *patient.UpsertMedicalTreatmentRequest) (*gcommon.EmptyResponse, error) {
	return util.WithSafeErrBiFunction(ctx, req, c.patientService.UpsertMedicalTreatment)
}

func (c *PatientGrpcController) UpsertMedicalSurgery(ctx context.Context, req *patient.UpsertMedicalSurgeryRequest) (*gcommon.EmptyResponse, error) {
	return util.WithSafeErrBiFunction(ctx, req, c.patientService.UpsertMedicalSurgery)
}

func (c *PatientGrpcController) UpsertMedicalPrescription(ctx context.Context, req *patient.UpsertMedicalPrescriptionRequest) (*gcommon.EmptyResponse, error) {
	return util.WithSafeErrBiFunction(ctx, req, c.patientService.UpsertMedicalPrescription)
}

func (c *PatientGrpcController) GetMedicalHistoryDetail(ctx context.Context, req *gcommon.IdRequest) (*patient.GetMedicalHistoryDetailResponse, error) {
	return util.WithSafeErrBiFunction(ctx, req, c.patientService.GetMedicalHistoryDetail)
}
