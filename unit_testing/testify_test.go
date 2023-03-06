package unit_testing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	assertions := assert.New(t)

	// assert equality
	assertions.Equal(123, 123, "they should be equal")

	// assert inequality
	assertions.NotEqual(123, 456, "they should not be equal")

}

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *ExampleTestSuite) SetupTest() {
	fmt.Println("SetupTest")
	suite.VariableThatShouldStartAtFive = 5
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleTestSuite) TestExample() {
	fmt.Println("TestExample")
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	fmt.Println("TestExampleTestSuite")
	suite.Run(t, new(ExampleTestSuite))
}

func TestExampleTestSuite2(t *testing.T) {
	fmt.Println("TestExampleTestSuite2")
	suite.Run(t, new(ExampleTestSuite))
}
