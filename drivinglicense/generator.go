package drivinglicense

import "errors"

type Applicant interface {
	IsOver18() bool
	HoldsLicense() bool
}

type NumberGenerator struct{}

func (g NumberGenerator) Generate(a Applicant) (string, error) {

	if a.HoldsLicense() == true {
		return "", errors.New("applicant holds license")
	}

	return "", errors.New("Underaged Applicant")
}

func NewNumberGenerator() NumberGenerator {
	return NumberGenerator{}
}
