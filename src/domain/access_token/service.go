package access_token

import (
	"fmt"
	"github.com/saurabhjangir/bookstore_oauth-api/src/domain/rest/user"
	"github.com/saurabhjangir/bookstore_oauth-api/src/utils/errors"
	"strings"
)

type IatService interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(request CreateTokenRequest) (*AccessToken, *errors.RestErr)
}

type DbRepository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpires(AccessToken) *errors.RestErr
}

type IuserRepository interface {
	LoginUser(string, string) (*user.UserLoginResponse, *errors.RestErr)
}

type atService struct {
	dbrepo   DbRepository
	userRepo IuserRepository
}

func NewService(dbrepo DbRepository, userRepo IuserRepository ) *atService {
	return &atService{
		dbrepo:   dbrepo,
		userRepo: userRepo,
	}
}

func (s *atService) GetByID(accessToken string) (*AccessToken, *errors.RestErr) {
	accessToken = strings.TrimSpace(accessToken)
	if accessToken == "" {
		return nil, errors.NewRestErrBadRequest("empty access token")
	}
	at, err := s.dbrepo.GetByID(accessToken)
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (s *atService) Create(request CreateTokenRequest) (*AccessToken, *errors.RestErr) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	user, err := s.userRepo.LoginUser(request.Email, request.Password)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	at := GetNewAccessToken(user.Id)
	at.Generate()
	err = s.dbrepo.Create(at)
	if err != nil {
		return nil, err
	}
	return &at, nil
}

