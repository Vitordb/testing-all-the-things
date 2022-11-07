package drivinglicense_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"testing-all-the-things/drivinglicense"
)

type DrivingLicenseSuite struct {
	suite.Suite
}

func TestDrivingLicenseSuite(t *testing.T) {
	suite.Run(t, new(DrivingLicenseSuite))
}

type UnderageApplicant struct {
}

func (s *DrivingLicenseSuite) TestNoSecondLicense() {
	a := LicenseHolderApplicant{}
	l := &SpyLogger{}

	lg := drivinglicense.NewNumberGenerator(l)
	_, err := lg.Generate(a)

	s.Error(err)
	s.Contains(err.Error(), "applicant holds license")

	s.Equal(1, l.callCount)
	s.Contains(l.lastMessage, "applicant holds license")

}

func (s *DrivingLicenseSuite) TestUnderageApplicant() {
	a := UnderageApplicant{}
	l := &SpyLogger{}

	lg := drivinglicense.NewNumberGenerator(l)
	_, err := lg.Generate(a)

	s.Error(err)
	s.Contains(err.Error(), "Underaged")

	s.Equal(1, l.callCount)
	s.Contains(l.lastMessage, "Underaged")
}

func (u UnderageApplicant) IsOver18() bool {
	return false
}

func (u UnderageApplicant) HoldsLicense() bool {
	return false
}

type LicenseHolderApplicant struct{}

func (l LicenseHolderApplicant) IsOver18() bool {
	return true
}

func (l LicenseHolderApplicant) HoldsLicense() bool {
	return true
}

type SpyLogger struct {
	callCount   int
	lastMessage string
}

func (s *SpyLogger) LogStuff(v string) {
	s.callCount++
	s.lastMessage = v
}
