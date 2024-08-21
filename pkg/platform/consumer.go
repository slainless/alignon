package platform

import "github.com/slainless/my-alignon/pkg/internal/artifact/database/my_alignon/public/model"

type Consumer struct {
	model.Consumers
	Nik         string `form:"nik" validate:"required,len=16"`
	FullName    string `form:"full_name" validate:"required,max=255"`
	LegalName   string `form:"legal_name" validate:"required,max=255"`
	BirthPlace  string `form:"birth_place" validate:"required,max=255"`
	P_BirthDate string `form:"birth_date" validate:"required,datetime=2006-01-02"`
	Salary      int64  `form:"salary"`
}

type Limit struct {
	model.Limits
	Tenor1 int64 `json:"tenor_1"`
	Tenor2 int64 `json:"tenor_2"`
	Tenor3 int64 `json:"tenor_3"`
	Tenor4 int64 `json:"tenor_4"`
}
