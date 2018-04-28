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
	"io/ioutil"
	"net/http"
	"os"
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
		if myrand < 0 {
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
			_, err := w.Write([]byte("503 - Now, young Skywalker, you will die"))
			if err != nil {
				logger.Error("mydemo",
					zap.String("status", "ERROR"),
					zap.Int("statusCode", 500),
					zap.Duration("backoff", time.Second),
					zap.Error(err),
				)
			}
			logger.Error("mydemo",
				zap.String("status", "ERROR"),
				zap.Int("statusCode", 503),
				zap.Duration("backoff", time.Second),
			)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("200 - Chewie, we’re home"))
		if err != nil {
			logger.Error("mydemo",
				zap.String("status", "ERROR"),
				zap.Int("statusCode", 500),
				zap.Duration("backoff", time.Second),
				zap.Error(err),
			)
		}
		logger.Info("Success send Healthz",
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
	_, err = w.Write(b)
	if err != nil {
		logger.Error("mydemo",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
	}
}

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
	port := os.Getenv("MY_TARGET_PING_PORT")
	url := "http://" + svc + ":" + port + "/ping"
	client := &http.Client{}
	post, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error("mydemo",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
	}
	for header, values := range req.Header {
		for _, value := range values {
			post.Header.Set(header, value)
			logger.Info("Send Header: "+header+" "+value,
				zap.String("status", "INFO"),
				zap.Duration("backoff", time.Second),
			)
		}
	}
	resp, err := client.Do(post)
	if err != nil {
		logger.Error("mydemo",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
	}
	defer resp.Body.Close() // nolint: errcheck
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
	logger.Info("Success send ping",
		zap.String("status", "INFO"),
		zap.Int("statusCode", 200),
		zap.Duration("backoff", time.Second),
	)
}

// Pong func reply to api ping
func Pong(w http.ResponseWriter, req *http.Request) {
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
	for header, values := range req.Header {
		for _, value := range values {
			req.Header.Get(header)
			logger.Info("Received Header: "+header+" "+value,
				zap.String("status", "INFO"),
				zap.Duration("backoff", time.Second),
			)
		}
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("200 - Pong"))
	if err != nil {
		logger.Error("mydemo",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
	}
	logger.Info("Success reply",
		zap.String("status", "INFO"),
		zap.Int("statusCode", 200),
		zap.Duration("backoff", time.Second),
	)
}
