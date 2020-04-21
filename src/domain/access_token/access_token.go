package access_token

import (
	"github.com/saurabhjangir/bookstore_oauth-api/src/domain/errors"
	"github.com/saurabhjangir/bookstore_oauth-api/src/utils/crypto_utils"
	"strings"
	"time"
	"fmt"
)

const (
	expirationtime = 24
)

type AccessToken struct{
	ID int64`json:"access_token_id"`
	ClientID int64 `json:"client_id"`
	UserID int64 `json:"user_id"`
	Token string `json:"token"`
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