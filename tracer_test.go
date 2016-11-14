package gotee

import (
	"log"
	"testing"

	"github.com/opentracing/basictracer-go"
	"github.com/opentracing/basictracer-go/examples/dapperish"
	"github.com/opentracing/opentracing-go"
)

func TestTracer(t *testing.T) {
	recorders := []basictracer.SpanRecorder{
		dapperish.NewTrivialRecorder("testDapper1"),
		dapperish.NewTrivialRecorder("testDapper2"),
	}

	teeTracer := New(recorders)

	opentracing.InitGlobalTracer(teeTracer)

	opentracing.StartSpan("test").Finish()
}
