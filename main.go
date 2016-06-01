package metrics // import "github.com/itsoneiota/metrics"

import (
	"time"

	"github.com/cactus/go-statsd-client/statsd"
)

// Metrics interface
type Metrics interface {
	Inc(statName string, value int64)
}

// MetricClient - wrapper
type MetricClient struct {
	Client statsd.Statter
}

// MockMetricClient - wrapper
type MockMetricClient struct {
	MetricMap map[string]int64
}

// NewMetricsClient - prefx should be service name
func NewMetricsClient(host string, prefix string) *MetricClient {
	mtrcs, err := statsd.NewBufferedClient(host, prefix, 300*time.Millisecond, 0)
	if err != nil {
		panic("Error creating Metrics client.")
	}
	mclient := MetricClient{Client: mtrcs}
	return &mclient
}

// NewMockMetricsClient - creates mock client
func NewMockMetricsClient() *MockMetricClient {
	mockClient := MockMetricClient{MetricMap: make(map[string]int64)}
	return &mockClient
}

// Inc - increment metric
func (mClient MetricClient) Inc(metricName string, value int64) {
	mClient.Client.Inc(metricName, value, 1.0)
}

// Inc - increment mock metric
func (mClient MockMetricClient) Inc(metricName string, value int64) {
	mClient.MetricMap[metricName] += value
}
