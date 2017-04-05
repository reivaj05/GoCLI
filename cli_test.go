package GoCLI

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CLITestSuite struct {
	suite.Suite
	assert *assert.Assertions
}

func (suite *CLITestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
}

func (suite *CLITestSuite) TearDownSuite() {
}

func (suite *CLITestSuite) TestCreateCLIAppWithDefaultActionAndCommands() {
	options := createCLIMockOptions()
	app := createCLIApp(options)
	suite.assert.NotNil(app)
	suite.assert.Equal("mockName", app.Name)
	suite.assert.Equal("mockUsage", app.Usage)
	suite.assert.Equal(len(options.Commands), len(app.Commands))
}

func createCLIMockOptions() *Options {
	return &Options{
		AppName:       "mockName",
		AppUsage:      "mockUsage",
		DefaultAction: mockAction,
		Commands:      createMockCommands(),
	}
}

func createMockCommands() []*Command {
	return []*Command{
		&Command{
			Name:   "mockNameCommand",
			Usage:  "mockUsageCommand",
			Action: mockAction,
		},
	}
}

func mockAction(args ...string) error {
	return nil
}

func TestCLI(t *testing.T) {
	suite.Run(t, new(CLITestSuite))
}
