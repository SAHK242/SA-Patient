package grpcmapper

import (
	"patient/ent"
	"patient/proto/patient"
)

type (
	PatientGrpcMapper interface {
		ConvertPatient(from *ent.Patient) *patient.PatientDetail
		ConvertPatientSlice(from []*ent.Patient) []*patient.PatientDetail
	}
)
