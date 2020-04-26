package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/saurabhjangir/bookstore_oauth-api/src/domain/access_token"
	"github.com/saurabhjangir/bookstore_userapi/utils/errors"
	"io/ioutil"
	"net/http"
)

type IaccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	atService access_token.IatService
}

func NewAccessTokenHandler(service access_token.IatService) IaccessTokenHandler {
	return &accessTokenHandler{
		atService: service,
	}
}
func (handler *accessTokenHandler) GetByID(c *gin.Context){
	at, Resterr := handler.atService.GetByID(c.Param("access_token"))
	if Resterr != nil {
		c.JSON(http.StatusBadRequest, Resterr)
		return
	}
	c.JSON(http.StatusOK, at)
}

func (handler *accessTokenHandler)Create(c *gin.Context){
	var atRequest access_token.CreateTokenRequest
	bytes, err  := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewRestErrBadRequest(err.Error()))
		return
	}
	err = json.Unmarshal(bytes, &atRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewRestErrBadRequest(err.Error()))
		return
	}

	at, createErr := handler.atService.Create(atRequest)
	if createErr != nil {
		c.JSON(createErr.Status, createErr)
		return
	}
	c.JSON(http.StatusCreated, at)
	return
}
