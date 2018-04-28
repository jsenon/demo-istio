// Package apitrace demoserver
//
// the purpose of this package is to provide Api Interface
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost:9010
//     BasePath: /api
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Julien SENON <julien.senon@gmail.com>
package apitrace

import (
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
)

// Play func launch ping to target answer server
func Play(w http.ResponseWriter, req *http.Request) {
	logger, err := zap.NewProduction()
	if err != nil {
		logger.Error("mydemo",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
	}
	defer logger.Sync() // nolint: errcheck

	svc := os.Getenv("MY_TARGET_PING_SVC")
	port := os.Getenv("MY_TARGET_PING_SPANPORT")
	url := "http://" + svc + ":" + port + "/ping"
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("mydemo",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
	}
	defer resp.Body.Close() // nolint: errcheck
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("mydemo",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		logger.Error("mydemo",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
	}
	logger.Info("mydemo",
		zap.String("status", "INFO"),
		zap.Int("statusCode", 200),
		zap.Duration("backoff", time.Second),
	)
}
