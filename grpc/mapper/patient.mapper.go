package grpcmapper

import (
	"patient/ent"
	"patient/proto/gcommon"
	"patient/proto/patient"
)

type patientGrpcMapper struct{}

func (p patientGrpcMapper) ConvertPatient(from *ent.Patient) *patient.PatientDetail {
	to := &patient.PatientDetail{
		Patient: &patient.Patient{
			Id:          from.ID.String(),
			FirstName:   from.FirstName,
			LastName:    from.LastName,
			PhoneNumber: from.PhoneNumber,
			Address:     from.Address,
			DateOfBirth: from.DateOfBirth.Unix(),
			Gender:      gcommon.Gender(from.Gender),
		},
		InpatientDetail:  nil,
		OutpatientDetail: nil,
	}
	return to
}

func (p patientGrpcMapper) ConvertPatientSlice(from []*ent.Patient) []*patient.PatientDetail {
	to := make([]*patient.PatientDetail, len(from))
	for i, v := range from {
		to[i] = p.ConvertPatient(v)
	}
	return to
}

func NewPatientGrpcMapper() PatientGrpcMapper {
	return &patientGrpcMapper{}
}
