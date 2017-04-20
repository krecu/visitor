package processor

import (
	"visitor/model"
)

type PersonalProcessor struct {
}

func (r *PersonalProcessor) Process(param string) (model.Personal, error) {

	return model.Personal{
		Ua: param,
	}, nil
}