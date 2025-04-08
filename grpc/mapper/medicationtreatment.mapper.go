package grpcmapper

import (
	"patient/ent"
	"patient/proto/patient"
)

type medicalTreatmentGrpcMapper struct{}

func (p medicalTreatmentGrpcMapper) ConvertMedicalTreatment(from *ent.MedicalTreatment) *patient.MedicalTreatment {
	if from == nil {
		return nil
	}
	to := &patient.MedicalTreatment{
		Id:               from.ID.String(),
		MedicalHistoryId: from.MedicalHistoryID.String(),
		StartDate:        from.StartDate.UnixMilli(),
		EndDate:          from.EndDate.UnixMilli(),
		Name:             from.Name,
		Result:           from.Result,
		Description:      from.Description,
		Fee:              from.Fee,
		MainDoctorId:     from.MainDoctorID.String(),
		CreatedAt:        from.CreatedAt.UnixMilli(),
		UpdatedAt:        from.UpdatedAt.UnixMilli(),
		CreatedBy:        &patient.User{Id: from.CreatedBy.String()},
		UpdatedBy:        &patient.User{Id: from.UpdatedBy.String()},
	}

	if from.SupportDoctorIds != nil {
		to.SupportDoctorIds = *from.SupportDoctorIds
	}

	if from.SupportNurseIds != nil {
		to.SupportNurseIds = *from.SupportNurseIds
	}

	return to
}

func (p medicalTreatmentGrpcMapper) ConvertMedicalTreatmentSlice(from []*ent.MedicalTreatment) []*patient.MedicalTreatment {
	to := make([]*patient.MedicalTreatment, len(from))
	for i, v := range from {
		to[i] = p.ConvertMedicalTreatment(v)
	}
	return to
}

func NewMedicalTreatmentGrpcMapper() MedicalTreatmentGrpcMapper {
	return &medicalTreatmentGrpcMapper{}
}
