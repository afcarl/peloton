// @generated AUTO GENERATED - DO NOT EDIT!
// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package requirements

import (
	"fmt"
	"strings"
)

// Transcript represents a transcript of which requirements passed and failed when evaluating groups for an entity.
type Transcript struct {
	Requirement  string
	GroupsPassed int
	GroupsFailed int
	Subscripts   map[Transcriptable]*Transcript
}

// NewTranscript creates a new transcript with a description.
func NewTranscript(name string) *Transcript {
	return &Transcript{
		Requirement: name,
		Subscripts:  map[Transcriptable]*Transcript{},
	}
}

// Subscript will create a sub transcript with the given description if one does not exist.
func (transcript *Transcript) Subscript(transcriptable Transcriptable) *Transcript {
	if transcript == nil {
		return nil
	}
	subscript, exists := transcript.Subscripts[transcriptable]
	if !exists {
		composite, name := transcriptable.Composite()
		if !composite {
			name = transcriptable.String()
		}
		subscript = NewTranscript(name)
		transcript.Subscripts[transcriptable] = subscript
	}
	return subscript
}

// IncPassed will increment the number of groups that passed the requirements.
func (transcript *Transcript) IncPassed() {
	if transcript == nil {
		return
	}
	transcript.GroupsPassed++
}

// IncFailed will increment the number of groups that failed the requirements.
func (transcript *Transcript) IncFailed() {
	if transcript == nil {
		return
	}
	transcript.GroupsFailed++
}

func (transcript *Transcript) string(indent int) string {
	space := strings.Repeat(" ", indent)
	result := fmt.Sprintf("%v%v passed %v times and failed %v times\n",
		space, transcript.Requirement, transcript.GroupsPassed, transcript.GroupsFailed)
	for _, subTranscript := range transcript.Subscripts {
		result += subTranscript.string(indent + 1)
	}
	return result
}

// String will create a human readable representation of the transcript.
func (transcript *Transcript) String() string {
	return transcript.string(0)
}
