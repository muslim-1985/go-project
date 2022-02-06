package unit

import (
	"github.com/stretchr/testify/suite"
	"go_project/src/users/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
)

type AuthTestSuite struct {
	suite.Suite
	User models.User
}

func (suite *AuthTestSuite) SetupTest() {
	suite.User.Password = "123456"
}

func (suite *AuthTestSuite) TestComparePassword() {

	bytePassword := []byte(suite.User.Password)
	hash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost)
	err := suite.User.IsPasswordValid(string(hash))
	if err != nil {
		log.Fatal(err)
	}
	suite.Equal(suite.User.Password, "123456")
}

func (suite *AuthTestSuite) TestCreateHashPassword()  {
	cleanPass := []byte(suite.User.Password)

	err := suite.User.CreatePasswordHash()
	if err != nil {
		log.Fatal(err)
	}
	hashPass := []byte(suite.User.Password)

	result := bcrypt.CompareHashAndPassword(hashPass, cleanPass)

	suite.Equal(result, nil)
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}
