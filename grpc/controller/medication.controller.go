package controller

import (
	"context"
	grpcservice "patient/grpc/service"
	util "patient/grpc/util"
	"patient/proto/gcommon"
	"patient/proto/patient"
)

type MedicationGrpcController struct {
	patient.UnimplementedMedicationServiceServer
	medicationService grpcservice.MedicationGrpcService
}

func NewMedicationGrpcController(
	medicationService grpcservice.MedicationGrpcService,
) *MedicationGrpcController {
	return &MedicationGrpcController{
		medicationService: medicationService,
	}
}

func (c *MedicationGrpcController) UpsertMedication(ctx context.Context, req *patient.UpsertMedicationRequest) (*gcommon.EmptyResponse, error) {
	return util.WithSafeErrBiFunction(ctx, req, c.medicationService.UpsertMedication)
}

func (c *MedicationGrpcController) ListMedication(ctx context.Context, req *patient.ListMedicationRequest) (*patient.ListMedicationResponse, error) {
	return util.WithSafeErrBiFunction(ctx, req, c.medicationService.ListMedication)
}
