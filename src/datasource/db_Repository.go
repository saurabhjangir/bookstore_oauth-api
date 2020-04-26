package datasource

import (
	"github.com/saurabhjangir/bookstore_oauth-api/src/clients/cassandra"
	"github.com/saurabhjangir/bookstore_oauth-api/src/domain/access_token"
	"github.com/saurabhjangir/bookstore_oauth-api/src/utils/errors"
)

const (
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, client_id, user_id, expires) VALUES(?,?,?,?);"
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type DbRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpires(access_token.AccessToken) *errors.RestErr
}
type dbRepository struct{}

func NewRepository() *dbRepository {
	return &dbRepository{}
}

func (s *dbRepository) GetByID(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	var at access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(
		&at.Token,
		&at.UserID,
		&at.ClientID,
		&at.Expires,
	); err != nil {
		return nil, errors.NewRestErrInteralServer(err.Error())
	}
	return &at, nil
}

func (s *dbRepository) Create(token access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return err
	}
	defer session.Close()
	if err := session.Query(queryCreateAccessToken,
		token.Token,
		token.ClientID,
		token.UserID,
		token.Expires).Exec(); err != nil {
		return errors.NewRestErrInteralServer(err.Error())
	}
	return nil
}

func (s *dbRepository) UpdateExpires(token access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return err
	}
	defer session.Close()
	if err := session.Query(queryUpdateExpires,
		token.Expires,
		token.Token,
	).Exec(); err != nil {
		return errors.NewRestErrInteralServer(err.Error())
	}
	return nil
}
