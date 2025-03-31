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
