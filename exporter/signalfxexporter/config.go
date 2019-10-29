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
	"github.com/open-telemetry/opentelemetry-collector/config/configmodels"
)

// Config defines configuration for logging exporter.
type Config struct {
	configmodels.ExporterSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.
	// AuthToken is the authentication token used to authorize with signalfx
	AuthToken string `mapstructure:"authtoken"`
	// DatapointEndpoint is url for sending metrics to signalfx
	DatapointEndpoint string `mapstructure:"DatapointEndpoint"`
	// TraceEndpoint is url for sending traces to signalfx
	TraceEndpoint string `mapstructure:"TraceEndpoint"`
}
