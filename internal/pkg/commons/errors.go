package commons

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/kitabisa/go-bootstrap/config"

	phttp "github.com/kitabisa/perkakas/v2/http"
	"github.com/kitabisa/perkakas/v2/structs"
)

var cfg = config.Config()

// InjectErrors injecting all error response to the handler context
func InjectErrors(handlerCtx *phttp.HttpHandlerContext) {
	handlerCtx.AddError(ErrDBConn, ErrDBConnResp)
	handlerCtx.AddError(ErrCacheConn, ErrCacheConnResp)
	// etc...
}

// ErrDBConn error type for Error DB Connection
var ErrDBConn = errors.New("ErrDBConn")

// ErrDBConnResp ErrDBConn's response
var ErrDBConnResp *structs.ErrorResponse = &structs.ErrorResponse{
	Response: structs.Response{
		ResponseCode: "101001",
		ResponseDesc: structs.ResponseDesc{
			ID: cfg.GetString(fmt.Sprintf("%s%s", "response_code.ID.", "101001")),
			EN: cfg.GetString(fmt.Sprintf("%s%s", "response_code.EN.", "101001")),
		},
	},
	HttpStatus: http.StatusInternalServerError,
}

// ErrCacheConn error type for Error Cache Connection
var ErrCacheConn = errors.New("ErrCacheConn")

// ErrCacheConnResp ErrCacheConn's response
var ErrCacheConnResp *structs.ErrorResponse = &structs.ErrorResponse{
	Response: structs.Response{
		ResponseCode: "101002",
		ResponseDesc: structs.ResponseDesc{
			ID: cfg.GetString(fmt.Sprintf("%s%s", "response_code.ID.", "101002")),
			EN: cfg.GetString(fmt.Sprintf("%s%s", "response_code.EN.", "101002")),
		},
	},
	HttpStatus: http.StatusInternalServerError,
}
