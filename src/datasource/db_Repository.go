package datasource

import (
	"github.com/saurabhjangir/bookstore_oauth-api/src/domain/access_token"
	"github.com/saurabhjangir/bookstore_oauth-api/src/domain/errors"
)

type DbRepository interface {
	GetByID(int64 )(*access_token.AccessToken, *errors.RestErr )
}
type dbRepository struct {}

func NewRepository() *dbRepository{
	return &dbRepository{}
}

func (s *dbRepository)GetByID(int64 )(*access_token.AccessToken, *errors.RestErr ){
	//database logic to get by ID
	return nil, errors.NewRestErrInteralServer("database down !!")
}