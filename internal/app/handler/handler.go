package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/go-bootstrap/config"
	"github.com/kitabisa/go-bootstrap/internal/app/service"
	"github.com/kitabisa/go-bootstrap/version"
	"gopkg.in/gorp.v2"
)

// Response our standar response object
type Response struct {
	Code int          `json:"response_code"`
	Desc ResponseDesc `json:"response_desc"`
	Next *string      `json:"next,omitempty"`
	Data interface{}  `json:"data,omitempty"`
	Meta ResponseMeta `json:"meta"`
}

// ResponseDesc response descriptio in multi language
type ResponseDesc struct {
	ID string `json:"id"`
	EN string `json:"en"`
}

// ResponseMeta our meta for the services
type ResponseMeta struct {
	Version string `json:"version"`
	Name    string `json:"api_name"`
	Env     string `json:"api_env"`
}

type Handler struct {
	config    config.Provider
	services  *service.Service
	dbMysql   *gorp.DbMap
	dbPostgre *gorp.DbMap
	cachePool *redis.Pool
}

func NewHandler(svc *service.Service, dbMysql *gorp.DbMap, dbPostgre *gorp.DbMap, cachePool *redis.Pool) *Handler {
	cfg := config.Config()
	return &Handler{
		config:    cfg,
		services:  svc,
		dbMysql:   dbMysql,
		dbPostgre: dbPostgre,
		cachePool: cachePool,
	}
}

// WriteResponse writing response to client
// httpStatus as the HTTP Status
// respCode as the response code
// data is interface: struct or array of struct
func (h *Handler) WriteResponse(w http.ResponseWriter, httpStatus int, respCode int, data interface{}, next *string) {
	w.WriteHeader(httpStatus)

	resp := Response{}
	resp.Code = respCode
	resp.Desc = h.getRespDesc(respCode)
	resp.Next = next

	// resp.Data
	voData := reflect.ValueOf(data)
	arrayData := []interface{}{}
	if voData.Kind() != reflect.Slice {
		if voData.IsValid() {
			arrayData = []interface{}{data}
		}
		resp.Data = arrayData
	} else {
		if voData.Len() != 0 {
			resp.Data = data
		} else {
			resp.Data = arrayData
		}
	}

	//resp.Meta
	resp.Meta = ResponseMeta{
		Version: version.Version,
		Name:    h.config.GetString("app.name"),
		Env:     version.Environment,
	}

	respJSON, _ := json.Marshal(resp)
	w.Write(respJSON)
}

func (h *Handler) getRespDesc(respCode int) ResponseDesc {
	respCodeStr := fmt.Sprintf("%d", respCode)
	return ResponseDesc{
		ID: h.config.GetString(fmt.Sprintf("%s%s", "response_code.ID.", respCodeStr)),
		EN: h.config.GetString(fmt.Sprintf("%s%s", "response_code.EN.", respCodeStr)),
	}
}
