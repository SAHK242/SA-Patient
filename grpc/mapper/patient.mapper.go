package grpcmapper

import (
	"patient/ent"
	"patient/proto/gcommon"
	"patient/proto/patient"
)

type patientGrpcMapper struct{}

func (p patientGrpcMapper) ConvertPatient(from *ent.Patient) *patient.Patient {
	if from == nil {
		return nil
	}

	to := &patient.Patient{
		Id:          from.ID.String(),
		FirstName:   from.FirstName,
		LastName:    from.LastName,
		PhoneNumber: from.Phone,
		Address:     from.Address,
		DateOfBirth: from.DateOfBirth.UnixMilli(),
		Gender:      gcommon.Gender(from.Gender),
		CreatedAt:   from.CreatedAt.UnixMilli(),
		UpdatedAt:   from.UpdatedAt.UnixMilli(),
		CreatedBy:   &patient.User{Id: from.CreatedBy.String()},
		UpdatedBy:   &patient.User{Id: from.UpdatedBy.String()},
	}
	return to
}

func (p patientGrpcMapper) ConvertPatientSlice(from []*ent.Patient) []*patient.Patient {
	to := make([]*patient.Patient, len(from))
	for i, v := range from {
		to[i] = p.ConvertPatient(v)
	}
	return to
}

func NewPatientGrpcMapper() PatientGrpcMapper {
	return &patientGrpcMapper{}
}
