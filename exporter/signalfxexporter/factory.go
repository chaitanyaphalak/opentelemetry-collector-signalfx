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
	"errors"
	"github.com/signalfx/golib/sfxclient"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector/config/configmodels"
	"github.com/open-telemetry/opentelemetry-collector/exporter"
)

const (
	// The value of "type" key in configuration.
	typeStr = "signalfx"
)

// Factory is the factory for logging exporter.
type Factory struct {
}

// Type gets the type of the Exporter config created by this factory.
func (f *Factory) Type() string {
	return typeStr
}

// CreateDefaultConfig creates the default configuration for exporter.
func (f *Factory) CreateDefaultConfig() configmodels.Exporter {
	return &Config{
		ExporterSettings: configmodels.ExporterSettings{
			TypeVal: typeStr,
			NameVal: typeStr,
		},
		AuthToken: "rhy8R-BiL_GbMZUgCenVwg",
		DatapointEndpoint: "https://ingest.us1.signalfx.com/v2/datapoint",
		TraceEndpoint: "https://ingest.us1.signalfx.com/v2/trace",
	}
}

func (f *Factory) createSignalFxClient(config configmodels.Exporter) (*sfxclient.HTTPSink, error) {
	cfg := config.(*Config)
	var signalFxClient = sfxclient.NewHTTPSink()
	// modify endpoints if needed
	signalFxClient.DatapointEndpoint = cfg.DatapointEndpoint
	signalFxClient.TraceEndpoint = cfg.TraceEndpoint
	if cfg.AuthToken == "" {
		return nil, errors.New("exporter config requires a non-empty 'AuthToken'")
	}
	signalFxClient.AuthToken = cfg.AuthToken
	return signalFxClient, nil
}

// CreateTraceExporter creates a trace exporter based on this config.
func (f *Factory) CreateTraceExporter(logger *zap.Logger, config configmodels.Exporter) (exporter.TraceExporter, error) {
	cfg := config.(*Config)
	client, err := f.createSignalFxClient(cfg)
	if err != nil {
		return nil, err
	}

	traceExp, err := NewTraceExporter(config, client)
	if err != nil {
		return nil, err
	}

	return traceExp, err
}

// CreateMetricsExporter creates a metrics exporter based on this config.
func (f *Factory) CreateMetricsExporter(logger *zap.Logger, config configmodels.Exporter) (exporter.MetricsExporter, error) {
	cfg := config.(*Config)
	client, err := f.createSignalFxClient(cfg)
	if err != nil {
		return nil, err
	}

	metricsExp, err := NewMetricsExporter(config, client)
	if err != nil {
		return nil, err
	}

	return metricsExp, err
}

