package drivinglicense_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"testing-all-the-things/drivinglicense"
)

func (s *DrivingLicenseSuite) TestUnderageApplicant() {
	a := UnderageApplicant{}

	lg := drivinglicense.NewNumberGenerator()
	_, err := lg.Generate(a)

	s.Error(err)
	s.Contains(err.Error(), "Underaged")

}

func (s *DrivingLicenseSuite) TestNoSecondLicense() {
	a := LicenseHolderApplicant{}

	lg := drivinglicense.NewNumberGenerator()
	_, err := lg.Generate(a)

	s.Error(err)
	s.Contains(err.Error(), "applicant holds license")

}

type DrivingLicenseSuite struct {
	suite.Suite
}

func TestDrivingLicenseSuite(t *testing.T) {
	suite.Run(t, new(DrivingLicenseSuite))
}

type UnderageApplicant struct {
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
