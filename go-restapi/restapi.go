package gorestapi

import (
	"net/http"
	"time"

	// "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	swagger "github.com/swaggo/echo-swagger"
)

type (
	RestApiResponse struct {
		Status  string      `json:"status"`
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}
	GoApi struct{}
)

func init() {
	InitConfig()
}

func NewRestApi() *GoApi {
	return &GoApi{}
}

func (x *GoApi) ServeHttp(e *echo.Echo) {
	prepare(e)
	log.Panic(e.Start(HttpEndpoint()))
}

func (x *GoApi) Prepare(e *echo.Echo) *echo.Echo {
	p := prepare(e)
	if ErrorLoadConfig != nil {
		log.WithFields(log.Fields{
			"at":      time.Now().Format(time.RFC3339),
			"comment": "something wrong when load config",
		}).Panic(ErrorLoadConfig)
	}
	return p
}

func prepare(e *echo.Echo) *echo.Echo {
	e.GET("/swagger/*", swagger.WrapHandler)
	e.GET("/swagger", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	// e.Pre(MiddlewarePreProcces) -> any procces from other method or class

	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	// e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.BodyDump(Audit))

	// error handler
	e.HTTPErrorHandler = ErrorHandler

	return e

}
