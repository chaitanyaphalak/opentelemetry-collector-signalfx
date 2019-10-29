// Copyright 2019, OpenTelemetry Authors
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

package signalfxexporter

import (
	"context"
	"github.com/open-telemetry/opentelemetry-collector/consumer/consumerdata"
	"github.com/open-telemetry/opentelemetry-collector/exporter"
	"github.com/signalfx/golib/sfxclient"
)

type signalExporter struct {
	name     string
	client* sfxclient.HTTPSink
}

const (
	DefaultDatapointEndpointURL      = "https://ingest.signalfx.com/v2/datapoint"
	DefaultTracesEndpointURL      = "https://ingest.signalfx.com/v2/traces"
)

func NewSignalfxMetricsExporter(datapointEndpoint string, authToken string, name string) *signalExporter {
	var signalfxmetrics* signalExporter = new(signalExporter)
	signalfxmetrics.client = sfxclient.NewHTTPSink()
	// modify endpoints if needed
	signalfxmetrics.client.DatapointEndpoint = datapointEndpoint
	signalfxmetrics.client.AuthToken = authToken
	signalfxmetrics.name = name

	return signalfxmetrics
}

func (se *signalExporter) Start(host exporter.Host) error {
	return nil
}

func (se *signalExporter) ConsumeMetricsData(ctx context.Context, md consumerdata.MetricsData) error {
	return nil
}

func (se *signalExporter) Name() string {
	return se.name
}

// Shutdown stops the exporter and is invoked during shutdown.
func (se *signalExporter) Shutdown() error {
	return se.shutdown()
}

func NewSignalfxTraceExporter(traceEndpoint string, authToken string, name string) *signalExporter {
	var signalfxtraces* signalExporter = new(signalExporter)
	signalfxtraces.client = sfxclient.NewHTTPSink()
	// modify endpoints if needed
	signalfxtraces.client.TraceEndpoint = traceEndpoint
	signalfxtraces.client.AuthToken = authToken
	signalfxtraces.name = name

	return signalfxtraces
}

func (ze *signalExporter) ConsumeTraceData(ctx context.Context, td consumerdata.TraceData) (zerr error) {
	return nil
}