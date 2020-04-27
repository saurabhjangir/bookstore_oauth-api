package access_token

import (
	"github.com/saurabhjangir/utils-lib-golang/errors"
	"github.com/saurabhjangir/utils-lib-golang/crypto_utils"
	"strings"
	"time"
	"fmt"
)

const (
	expirationtime = 24
	granttypePassword = "password"
	granttypeClientcredentials = "clientcredentials"
)

type CreateTokenRequest struct{
	GrantType string `json:"grant_type"`
	//grant type password
	Password string `json:"password"`
	Email string `json:"email"`
	// grant type client_credentials
	ClientID string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (r *CreateTokenRequest) Validate() *errors.RestErr{
	var err *errors.RestErr
	switch r.GrantType {
	// TODO:
	case granttypePassword:
		err = nil
		break
	case granttypeClientcredentials:
		err = nil
		break
	default:
		err = errors.NewRestErrBadRequest("grant type not supported")
	}
	return err
}

type AccessToken struct{
	Token string `json:"access_token"`
	ClientID int64 `json:"client_id"`
	UserID int64 `json:"user_id"`
	Expires int64 `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.Token = strings.TrimSpace(at.Token)
	if at.Token == "" {
		return errors.NewRestErrBadRequest("invalid token")
	}
	if at.UserID <= 0 {
		return errors.NewRestErrBadRequest("invalid user id")
	}
	if at.ClientID <= 0 {
		return errors.NewRestErrBadRequest("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewRestErrBadRequest("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken(userID int64) (AccessToken) {
	return AccessToken{
		UserID: userID,
		Expires: time.Now().Add(expirationtime * time.Hour).Unix(),
	}
}

func (at AccessToken)IsExpired() bool {
	return time.Unix(at.Expires,0).Before(time.Now().UTC())
}

func (at *AccessToken)Generate(){
	s := fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires)
	at.Token = crypto_utils.GetMd5(s)
}