// Package web demoserver.
//
// the purpose of this package is to provide Web HTML Interface
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Julien SENON <julien.senon@gmail.com>
package web

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
)

// Present Information on Dedicated WebPortal

// Index func to display all server on table view
func Index(res http.ResponseWriter, req *http.Request) {
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
	// var rs Server
	ip, err := externalIP()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(ip)
	myversion := os.Getenv("MY_VERSION")
	if myversion == "" {
		myversion = "0.0.1"
	}
	_, err = io.WriteString(res, "Hello, Im Service version: "+myversion+"\n"+"My IP is: "+ip+"\n")
	logger.Info("mydemo",
		zap.String("status", "INFO"),
		zap.Int("statusCode", 200),
		zap.Duration("backoff", time.Second),
	)
	if err != nil {
		logger.Error("mydemo",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
	}

}

// externalIP func that display ip of pod
func externalIP() (string, error) { // nolint: gocyclo
	logger, err := zap.NewProduction()
	if err != nil {
		logger.Error("mydemo",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
		return "", err
	}

	defer logger.Sync() // nolint: errcheck
	ifaces, err := net.Interfaces()
	if err != nil {
		logger.Error("mydemo",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			logger.Error("mydemo",
				zap.String("status", "ERROR"),
				zap.Int("statusCode", 500),
				zap.Duration("backoff", time.Second),
				zap.Error(err),
			)
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("Are you connected to the network?")
}
