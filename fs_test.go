package appstore

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FilesystemTestSuite struct {
	suite.Suite
}

func (suite *FilesystemTestSuite) TestFileExistsSuccess() {
	exists := fileExists(StubAuthKeyPath)
	assert.True(suite.T(), exists)
}

func (suite *FilesystemTestSuite) TestReadFileSuccess() {
	exists, err := readFile(StubAuthKeyPath)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), exists)
}

func TestFilesystemTestSuite(t *testing.T) {
	suite.Run(t, new(FilesystemTestSuite))
}
