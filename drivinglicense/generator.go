package drivinglicense

import "errors"

type Applicant interface {
	IsOver18() bool
	HoldsLicense() bool
}

type Logger interface {
	LogStuff(v string)
}

type NumberGenerator struct {
	l Logger
}

func (g NumberGenerator) Generate(a Applicant) (string, error) {

	if a.HoldsLicense() {
		return "", errors.New("applicant holds license")
	}

	g.l.LogStuff("Underaged")

	return "", errors.New("Underaged Applicant")
}

func NewNumberGenerator(l Logger) NumberGenerator {
	return NumberGenerator{l: l}
}
