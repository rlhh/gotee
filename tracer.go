package gotee

import (
	basictracer "github.com/opentracing/basictracer-go"
	"github.com/opentracing/opentracing-go"
)

// New returns a bascitracer Tracer with a gotee recorder instance
func New(recorders []basictracer.SpanRecorder) opentracing.Tracer {
	return basictracer.New(NewRecorder(recorders))
}
