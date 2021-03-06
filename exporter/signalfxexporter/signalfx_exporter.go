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
	"fmt"
	"github.com/open-telemetry/opentelemetry-collector/config/configmodels"
	"github.com/open-telemetry/opentelemetry-collector/consumer/consumerdata"
	"github.com/open-telemetry/opentelemetry-collector/exporter"
	"github.com/open-telemetry/opentelemetry-collector/exporter/exporterhelper"
	"github.com/signalfx/golib/sfxclient"
)

type signalFxExporter struct {
	client* sfxclient.HTTPSink
}

func (*signalFxExporter) Name() string {
	return "signalfx"
}

func (sfxe *signalFxExporter) Shutdown() error {
	return nil
}

func NewTraceExporter(config configmodels.Exporter, client* sfxclient.HTTPSink) (exporter.TraceExporter, error) {
	sfxe := &signalFxExporter{client: client}

	exp , err := exporterhelper.NewTraceExporter(
		config,
		sfxe.PushTraceData,
		exporterhelper.WithTracing(true),
		exporterhelper.WithMetrics(true),
		exporterhelper.WithShutdown(sfxe.Shutdown),
	)

	if err != nil {
		return nil, err
	}

	return exp, nil
}

func (sfxe *signalFxExporter) PushTraceData(ctx context.Context, td consumerdata.TraceData) (int, error) {
	fmt.Println("PushTraceData")
	//var errs []error
	//goodSpans := 0
	//for _, span := range td.Spans {
	//	sd, err := spandatatranslator.ProtoSpanToOCSpanData(span)
	//	if err == nil {
	//		se.exporter.ExportSpan(sd)
	//		goodSpans++
	//	} else {
	//		errs = append(errs, err)
	//	}
	//}
	//
	//return len(td.Spans) - goodSpans, oterr.CombineErrors(errs)
	return 0, nil
}

func (sfxe *signalFxExporter) PushMetricsData(ctx context.Context, md consumerdata.MetricsData) (int, error) {
	for _, metric := range md.Metrics {
		fmt.Println("PushMetricsData: " + metric.String() + " " + md.Node.String() + " " + md.Resource.String())
	}
	return 0, nil
}

func NewMetricsExporter(config configmodels.Exporter, client* sfxclient.HTTPSink) (exporter.MetricsExporter, error) {
	sfxe := &signalFxExporter{client: client}

	exp , err := exporterhelper.NewMetricsExporter(
		config,
		sfxe.PushMetricsData,
		exporterhelper.WithTracing(true),
		exporterhelper.WithMetrics(true),
		exporterhelper.WithShutdown(sfxe.Shutdown),
	)

	if err != nil {
		return nil, err
	}

	return exp, nil
}