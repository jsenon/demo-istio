// Package api demoserver.
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
package api

import (
	"net/http"
	"time"

	"encoding/json"
	"fmt"
	"math/rand"

	"go.uber.org/zap"
)

// statusjson struct
type statusjson struct {
	// The Status message
	// in: body
	Code    int32  `json:"statuscode"`
	Message string `json:"statusmessage"`
}

// Health func that send 200
func Health(w http.ResponseWriter, req *http.Request) {
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
	myhead := req.Header.Get("X-Custom-Error")
	if myhead == "true" {
		myrand := rand.Intn(6)
		if myrand <= 0 {
			logger.Error("mydemo",
				zap.String("status", "ERROR"),
				zap.Int("statusCode", 500),
				zap.Duration("backoff", time.Second),
				zap.String("error", "random less than 0"),
			)
		}
		fmt.Println("Rand:", myrand)
		if myrand == 5 {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("503 - Now, young Skywalker, you will die"))
			logger.Error("mydemo",
				zap.String("status", "ERROR"),
				zap.Int("statusCode", 503),
				zap.Duration("backoff", time.Second),
			)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Chewie, we’re home"))
		logger.Info("mydemo",
			zap.String("status", "INFO"),
			zap.Int("statusCode", 200),
			zap.Duration("backoff", time.Second),
		)

	}
}

// Wellknown func that send information of service
func Wellknown(w http.ResponseWriter, req *http.Request) {
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
	answerjson := statusjson{
		Code:    200,
		Message: "Made by Somebody",
	}
	b, err := json.Marshal(answerjson)
	if err != nil {
		logger.Error("mydemo",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
	}
	w.Write(b)
}
