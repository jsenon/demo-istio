//go:generate swagger generate spec
// Package main demoserver.
//
// the purpose of this application is to provide an CMDB application
// that will store information in mongodb backend
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
package main

import (
	"net/http"
	"os"
	"time"

	"github.com/jsenon/demo-istio/api"
	"github.com/jsenon/demo-istio/web"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/zipkin"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		logger.Error("Failed to create zap logger",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
	}
	defer logger.Sync() // nolint: errcheck

	// Web Part
	http.HandleFunc("/", web.Index)

	// API Part
	http.HandleFunc("/healthz", api.Health)
	http.HandleFunc("/.well-known", api.Wellknown)
	http.HandleFunc("/play", api.Play)
	http.HandleFunc("/ping", api.Pong)

	// Block used to propagate span
	myjaeger := os.Getenv("MY_JAEGER_AGENT")
	// Check if Jaeger variable has been set
	if myjaeger != "" {
		spanname := os.Getenv("MY_SVC_SPAN_NAME")
		zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
		injector := jaeger.TracerOptions.Injector(opentracing.HTTPHeaders, zipkinPropagator)
		extractor := jaeger.TracerOptions.Extractor(opentracing.HTTPHeaders, zipkinPropagator)
		// Zipkin shares span ID between client and server spans; it must be enabled via the following option.
		zipkinSharedRPCSpan := jaeger.TracerOptions.ZipkinSharedRPCSpan(true)
		// sender, _ := jaeger.NewUDPTransport("jaeger-agent.istio-system:5775", 0)
		sender, err2 := jaeger.NewUDPTransport(myjaeger, 0)
		if err2 != nil {
			logger.Error("Failed to start jaeger udp transport",
				zap.String("status", "ERROR"),
				zap.Int("statusCode", 500),
				zap.Duration("backoff", time.Second),
				zap.Error(err2),
			)
		}
		tracer, closer := jaeger.NewTracer(
			spanname,
			jaeger.NewConstSampler(true),
			jaeger.NewRemoteReporter(
				sender,
				jaeger.ReporterOptions.BufferFlushInterval(1*time.Second)),
			injector,
			extractor,
			zipkinSharedRPCSpan,
		)
		defer closer.Close() // nolint: errcheck
		err = http.ListenAndServe(":9010", nethttp.Middleware(tracer, http.DefaultServeMux))
		if err != nil {
			logger.Error("Failed to start web server",
				zap.String("status", "ERROR"),
				zap.Int("statusCode", 500),
				zap.Duration("backoff", time.Second),
				zap.Error(err),
			)
		}
	} else {
		// If no Jaeger variable set we don't propagate header
		err = http.ListenAndServe(":9010", nil)
		if err != nil {
			logger.Error("Failed to start web server",
				zap.String("status", "ERROR"),
				zap.Int("statusCode", 500),
				zap.Duration("backoff", time.Second),
				zap.Error(err),
			)
		}
	}

}
