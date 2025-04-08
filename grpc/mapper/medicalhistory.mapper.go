package grpcmapper

import (
	"patient/ent"
	"patient/proto/patient"
)

type medicalHistoryGrpcMapper struct{}

func (p medicalHistoryGrpcMapper) ConvertMedicalHistory(from *ent.MedicalHistories) *patient.MedicalHistory {
	if from == nil {
		return nil
	}
	to := patient.MedicalHistory{
		Id:             from.ID.String(),
		PatientId:      from.PatientID.String(),
		Reason:         from.Reason,
		Diagnosis:      from.Diagnosis,
		DoctorNotes:    from.DoctorNotes,
		MedicalEndDate: from.MedicalEndDate.UnixMilli(),
		CreatedAt:      from.CreatedAt.UnixMilli(),
		UpdatedAt:      from.UpdatedAt.UnixMilli(),
		CreatedBy:      &patient.User{Id: from.CreatedBy.String()},
		UpdatedBy:      &patient.User{Id: from.UpdatedBy.String()},
	}

	if len(from.Edges.MedicalTreatment) > 0 {
		to.HasTreatment = true
		for _, v := range from.Edges.MedicalTreatment {
			to.TotalFee += v.Fee
		}
	}

	if len(from.Edges.MedicalSurgery) > 0 {
		to.HasSurgery = true
		for _, v := range from.Edges.MedicalSurgery {
			to.TotalFee += v.Fee
		}
	}

	if len(from.Edges.MedicalPrescription) > 0 {
		to.HasPrescription = true
		for _, v := range from.Edges.MedicalPrescription {
			to.TotalFee += v.Fee
		}
	}

	return &to
}

func (p medicalHistoryGrpcMapper) ConvertMedicalHistorySlice(from []*ent.MedicalHistories) []*patient.MedicalHistory {
	to := make([]*patient.MedicalHistory, len(from))
	for i, v := range from {
		to[i] = p.ConvertMedicalHistory(v)
	}
	return to
}

func NewMedicalHistoryGrpcMapper() MedicalHistoryGrpcMapper {
	return &medicalHistoryGrpcMapper{}
}
