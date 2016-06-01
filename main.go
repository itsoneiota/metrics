package metrics // import "github.com/itsoneiota/metrics"

import "github.com/cactus/go-statsd-client/statsd"

type metricPublisher interface {
	Inc(statName string, value int64)
}

// MetricPublisher -
type MetricPublisher struct {
	Client MetricClient
}

// MetricClient -
type MetricClient interface {
	Inc(statName string, value int64)
}

// StatsdMetricClient - wrapper
type StatsdMetricClient struct {
	Client statsd.Statter
}

// MockMetricClient - wrapper
type MockMetricClient struct {
	MetricMap map[string]int64
}

// NewMetricPublisher -
func NewMetricPublisher(mc MetricClient) *MetricPublisher {
	return &MetricPublisher{Client: mc}
}

// NewStatsdMetricsClient - prefx should be service name
func NewStatsdMetricsClient(host string, prefix string) *StatsdMetricClient {
	mtrcs, err := statsd.NewClient(host, prefix)
	if err != nil {
		panic("Error creating Metrics client.")
	}
	mclient := StatsdMetricClient{Client: mtrcs}
	return &mclient
}

// NewMockMetricsClient -
func NewMockMetricsClient() *MockMetricClient {
	mockClient := MockMetricClient{MetricMap: make(map[string]int64)}
	return &mockClient
}

// Inc - increment metric
func (mClient StatsdMetricClient) Inc(metricName string, value int64) {
	mClient.Client.Inc(metricName, value, 1.0)
}

// Inc - increment mock metric
func (mClient MockMetricClient) Inc(metricName string, value int64) {
	mClient.MetricMap[metricName] += value
}
