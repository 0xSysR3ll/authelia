package suites

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MSSQLSuite struct {
	*RodSuite
}

func NewMSSQLSuite() *MSSQLSuite {
	return &MSSQLSuite{
		RodSuite: NewRodSuite(mssqlSuiteName),
	}
}

func (s *MSSQLSuite) Test1FAScenario() {
	suite.Run(s.T(), New1FAScenario())
}

func (s *MSSQLSuite) TestTwoFactorTOTPScenario() {
	suite.Run(s.T(), NewTwoFactorTOTPScenario())
}

func TestMSSQLSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping suite test in short mode")
	}

	suite.Run(t, NewMSSQLSuite())
}
