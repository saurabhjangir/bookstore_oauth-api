package access_token

import "github.com/saurabhjangir/bookstore_oauth-api/src/domain/errors"

type Repository interface {
	GetByID(int64 )(*AccessToken, *errors.RestErr)
}

type AtService interface {
	GetByID(int64 )(*AccessToken, *errors.RestErr)
}

type atService struct{
	repo Repository
}

func NewAccessTokenService(repo Repository) *atService{
	return &atService{
		repo: repo,
	}
}

func (s *atService)GetByID(accessTokenID int64 ) (*AccessToken, *errors.RestErr ){
	at, err := s.repo.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return at, nil
}