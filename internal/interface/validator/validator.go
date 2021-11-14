package validator

import (
	"github.com/go-playground/validator/v10"

	"github.com/mfsyahrz/bareksaproject/internal/shared/constant"
)

type (
	StructLevelValidator interface {
		Validate() error
	}

	DataValidator struct {
		ValidatorData *validator.Validate
	}
)

func NewValidator() *DataValidator {
	v := &DataValidator{
		validator.New(),
	}
	v.registerCustomValidation()

	return v
}

func (v *DataValidator) Validate(i interface{}) (err error) {
	if err = v.ValidatorData.Struct(i); err != nil {
		return
	}
	return
}

func (v *DataValidator) registerCustomValidation() {
	_ = v.ValidatorData.RegisterValidation("news_status", validateNewsStatus)
}

var (
	newsStatusMap = map[string]bool{
		constant.TypeNewsStatusDraft.String():     true,
		constant.TypeNewsStatusDeleted.String():   true,
		constant.TypeNewsStatusPublished.String(): true,
	}
)

func validateNewsStatus(fl validator.FieldLevel) bool {
	return newsStatusMap[fl.Field().String()]
}
