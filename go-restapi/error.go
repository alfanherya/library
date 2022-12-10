package gorestapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func ErrorHandler(err error, c echo.Context) {
	if !c.Response().Committed {
		restResp := RestApiResponse{}
		restResp.Status = "error"
		restResp.Code = "99"
		report, ok := err.(*echo.HTTPError)
		if ok {
			res, yap := report.Message.(RestApiResponse)
			if yap {
				restResp = res
			} else {
				restResp.Message = report.Message.(string)
			}
		} else {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			restResp.Message = report.Message.(string)
		}

		log.WithFields(log.Fields{
			"at":      time.Now().Format(time.RFC3339),
			"comment": fmt.Sprintf("rest api errot with code:  %d amd error: %s", report.Code, report.Message),
		}).Warn()
		c.Set("error", err)
		_ = c.JSON(report.Code, restResp)
	}
}
