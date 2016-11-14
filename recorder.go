// Package gotee is influenced by the basictracing-go package examples
// https://github.com/opentracing/basictracer-go/tree/master/examples
package gotee

import (
	"fmt"

	basictracer "github.com/opentracing/basictracer-go"
)

// Recorder implements the opentracing.Recorder interface.
type Recorder struct {
	recorders []basictracer.SpanRecorder
	tags      map[string]string
}

// NewRecorder returns a new tee recorder.
func NewRecorder(recorders []basictracer.SpanRecorder) *Recorder {
	return &Recorder{
		tags:      make(map[string]string),
		recorders: recorders,
	}
}

// SetTag sets a tag
func (r *Recorder) SetTag(key string, val interface{}) *Recorder {
	r.tags[key] = fmt.Sprint(val)
	return r
}

// RecordSpan implements the basictracer.Recorder interface.
func (r *Recorder) RecordSpan(span basictracer.RawSpan) {
	r.recorders[0].RecordSpan(span)
	r.recorders[1].RecordSpan(span)
}
