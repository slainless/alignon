package platform

import "github.com/slainless/my-alignon/pkg/internal/artifact/database/my_alignon/public/model"

type ConsumerRegisterInput struct {
	model.Consumers
	Nik         string `form:"nik" json:"nik" validate:"required,len=16"`
	FullName    string `form:"full_name" json:"full_name" validate:"required,max=255"`
	LegalName   string `form:"legal_name" json:"legal_name" validate:"required,max=255"`
	BirthPlace  string `form:"birth_place" json:"birth_place" validate:"required,max=255"`
	P_BirthDate string `form:"birth_date" json:"birth_date" validate:"required,datetime=2006-01-02"`
	Salary      int64  `form:"salary" json:"salary" validate:"required,min=1,max=1000000000000"`
}

type Consumer struct {
	model.Consumers
	Nik         string `json:"nik"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	LegalName   string `json:"legal_name"`
	BirthPlace  string `json:"birth_place"`
	Salary      int64  `json:"salary"`
	KtpPhoto    string `json:"ktp_photo"`
	SelfiePhoto string `json:"selfie_photo"`
	P_BirthDate string `json:"birth_date"`
}

type Limit struct {
	model.Limits
	Tenor1 int64 `json:"tenor_1"`
	Tenor2 int64 `json:"tenor_2"`
	Tenor3 int64 `json:"tenor_3"`
	Tenor4 int64 `json:"tenor_4"`
}
