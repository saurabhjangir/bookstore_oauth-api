package http

import (
	"github.com/gin-gonic/gin"
	"github.com/saurabhjangir/bookstore_oauth-api/src/domain/access_token"
	"github.com/saurabhjangir/bookstore_oauth-api/src/domain/errors"
	"net/http"
	"strconv"
)

type AccessTokenHandler interface {
	GetByID(*gin.Context)
}

type accessTokenHandler struct {
	atService access_token.AtService
}

func NewAccessTokenHandler(service access_token.AtService) AccessTokenHandler {
	return &accessTokenHandler{
		atService: service,
	}
}
func (handler *accessTokenHandler) GetByID(c *gin.Context){
	id, true := c.Params.Get("access_token_id")
	if !true {
		c.JSON(http.StatusBadRequest, errors.NewRestErrBadRequest("http bad request"))
		return
	}
	Id, err := strconv.ParseInt(id, 10 , 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewRestErrBadRequest("http bad request"))
		return
	}
	at, Resterr := handler.atService.GetByID(Id)
	if Resterr != nil {
		c.JSON(http.StatusBadRequest, Resterr)
		return
	}
	c.JSON(http.StatusOK, at)
}