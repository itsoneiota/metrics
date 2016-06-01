package metrics

import "testing"

func TestMockMetricTypeEquals(t *testing.T) {
	client := NewMockMetricsClient()
	mPub := NewMetricPublisher(client)

	mPub.Client.Inc("TestMetric", 1)
	assertEquals(1, client.MetricMap["TestMetric"], t)
	mPub.Client.Inc("TestMetric", 1)
	assertEquals(2, client.MetricMap["TestMetric"], t)
	mPub.Client.Inc("TestMetric", 5)
	assertEquals(7, client.MetricMap["TestMetric"], t)
}

func assertEquals(expected int64, have int64, t *testing.T) {
	if expected != have {
		t.Errorf("Expected metric value of %x, was %x", expected, have)
	}
}
