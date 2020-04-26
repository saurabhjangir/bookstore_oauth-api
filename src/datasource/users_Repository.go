package datasource

import (
	"encoding/json"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/saurabhjangir/bookstore_oauth-api/src/domain/rest/user"
	"github.com/saurabhjangir/bookstore_oauth-api/src/utils/errors"
	"time"
)
const (
	UserLoginAPIEndpoint = "https://127.0.0.1:3300/user/login"
)
var (
	userRepo IuserRepository = &userRepository{}
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:3300",
		Timeout: 100 * time.Millisecond,
	}
)


type IuserRepository interface {
	LoginUser(string, string) (*user.UserLoginResponse, *errors.RestErr)
}

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (s *userRepository) LoginUser(email string, password string) (*user.UserLoginResponse, *errors.RestErr) {
	request := user.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	var result user.UserLoginResponse
	resp := usersRestClient.Post("/user/login", request)
	if resp == nil || resp.Response == nil{
		return nil, errors.NewRestErrBadRequest("Error connecting to user service")
	}

	if resp.StatusCode > 299 {
		var err errors.RestErr
		marshErr := json.Unmarshal(resp.Bytes(), &err)
		if marshErr != nil {
			return nil, &err
		}
		return nil, &err
	}
	marshErr := json.Unmarshal(resp.Bytes(), &result)
	if marshErr != nil {
		return nil, errors.NewRestErrBadRequest("Error processing user response")
	}
	return &result, nil
}
