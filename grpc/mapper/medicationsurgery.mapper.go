package grpcmapper

import (
	"patient/ent"
	"patient/proto/patient"
)

type medicalSurgeryGrpcMapper struct{}

func (p medicalSurgeryGrpcMapper) ConvertMedicalSurgery(from *ent.MedicalSurgery) *patient.MedicalSurgery {
	if from == nil {
		return nil
	}
	to := &patient.MedicalSurgery{
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

func (p medicalSurgeryGrpcMapper) ConvertMedicalSurgerySlice(from []*ent.MedicalSurgery) []*patient.MedicalSurgery {
	to := make([]*patient.MedicalSurgery, len(from))
	for i, v := range from {
		to[i] = p.ConvertMedicalSurgery(v)
	}
	return to
}

func NewMedicalSurgeryGrpcMapper() MedicalSurgeryGrpcMapper {
	return &medicalSurgeryGrpcMapper{}
}
