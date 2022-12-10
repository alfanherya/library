package gorestapi

import (
	"regexp"

	"github.com/labstack/echo/v4"
)

func Audit(c echo.Context, reqBody, resBody []byte) {
	skip, err := regexp.MatchString(SkipAuditLoggerURIPattern(), c.Request().URL.String())

	if err != nil {
		skip = false
	}
	if skip {
		return
	}

}
