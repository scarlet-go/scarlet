package scarlet

import (
	"errors"
	"fmt"
	"net/http"
)

type ScarletContext struct {
	Request   http.Request
	inherited map[string]interface{}
}

type IScarletContext interface {
	From(key string) (value interface{}, err error)
	To(key string, value interface{}) interface{}
	GetHeader(key string) string
	GetParam(key string) string
}

func (ctx *ScarletContext) From(key string) (value interface{}, err error) {
	if _, ok := ctx.inherited[key]; ok {
		return ctx.inherited[key], nil
	}

	errorMessage := fmt.Sprintf("key %s not found in inherited context", key)

	return nil, errors.New(errorMessage)
}

func (ctx *ScarletContext) To(key string, value interface{}) interface{} {
	if _, ok := ctx.inherited[key]; !ok {
		ctx.inherited[key] = make(map[string]interface{})
	}

	ctx.inherited[key] = value

	return value
}

func (ctx *ScarletContext) GetHeader(key string) string {
	return ctx.Request.Header.Get(key)
}

func (ctx *ScarletContext) GetParam(key string) string {
	return ctx.Request.PathValue(key)
}

func checkInterface() IScarletContext {
	return &ScarletContext{}
}
