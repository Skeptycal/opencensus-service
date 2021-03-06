// Copyright 2018, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Sample contains a program that exports to the OpenCensus service.
package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"contrib.go.opencensus.io/exporter/ocagent"
	"go.opencensus.io/trace"
)

func main() {
	oce, err := ocagent.NewExporter(ocagent.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to create ocagent-exporter: %v", err)
	}
	trace.RegisterExporter(oce)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		_, span := trace.StartSpan(context.Background(), "foo")
		time.Sleep(time.Duration(rng.Int63n(1000)) * time.Millisecond)
		span.End()
	}
}
