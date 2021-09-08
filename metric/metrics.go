package metric

import (
	"github.com/appleboy/gorush/status"

	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "gorush_"

// Metrics implements the prometheus.Metrics interface and
// exposes gorush metrics for prometheus
type Metrics struct {
	TotalPushCount *prometheus.Desc
	IosSuccess     *prometheus.Desc
	IosError       *prometheus.Desc
	AndroidSuccess *prometheus.Desc
	AndroidError   *prometheus.Desc
	HuaweiSuccess  *prometheus.Desc
	HuaweiError    *prometheus.Desc
	QueueUsage     *prometheus.Desc
	GetQueueUsage  func() int
}

var getGetQueueUsage = func() int { return 0 }

// NewMetrics returns a new Metrics with all prometheus.Desc initialized
func NewMetrics(c ...func() int) Metrics {
	m := Metrics{
		TotalPushCount: prometheus.NewDesc(
			namespace+"total_push_count",
			"Number of push count",
			nil, nil,
		),
		IosSuccess: prometheus.NewDesc(
			namespace+"ios_success",
			"Number of iOS success count",
			nil, nil,
		),
		IosError: prometheus.NewDesc(
			namespace+"ios_error",
			"Number of iOS fail count",
			nil, nil,
		),
		AndroidSuccess: prometheus.NewDesc(
			namespace+"android_success",
			"Number of android success count",
			nil, nil,
		),
		AndroidError: prometheus.NewDesc(
			namespace+"android_fail",
			"Number of android fail count",
			nil, nil,
		),
		HuaweiSuccess: prometheus.NewDesc(
			namespace+"huawei_success",
			"Number of huawei success count",
			nil, nil,
		),
		HuaweiError: prometheus.NewDesc(
			namespace+"huawei_fail",
			"Number of huawei fail count",
			nil, nil,
		),
		QueueUsage: prometheus.NewDesc(
			namespace+"queue_usage",
			"Length of internal queue",
			nil, nil,
		),
		GetQueueUsage: getGetQueueUsage,
	}

	if len(c) > 0 {
		m.GetQueueUsage = c[0]
	}

	return m
}

// Describe returns all possible prometheus.Desc
func (c Metrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.TotalPushCount
	ch <- c.IosSuccess
	ch <- c.IosError
	ch <- c.AndroidSuccess
	ch <- c.AndroidError
	ch <- c.HuaweiSuccess
	ch <- c.HuaweiError
	ch <- c.QueueUsage
}

// Collect returns the metrics with values
func (c Metrics) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(
		c.TotalPushCount,
		prometheus.CounterValue,
		float64(status.StatStorage.GetTotalCount()),
	)
	ch <- prometheus.MustNewConstMetric(
		c.IosSuccess,
		prometheus.CounterValue,
		float64(status.StatStorage.GetIosSuccess()),
	)
	ch <- prometheus.MustNewConstMetric(
		c.IosError,
		prometheus.CounterValue,
		float64(status.StatStorage.GetIosError()),
	)
	ch <- prometheus.MustNewConstMetric(
		c.AndroidSuccess,
		prometheus.CounterValue,
		float64(status.StatStorage.GetAndroidSuccess()),
	)
	ch <- prometheus.MustNewConstMetric(
		c.AndroidError,
		prometheus.CounterValue,
		float64(status.StatStorage.GetAndroidError()),
	)
	ch <- prometheus.MustNewConstMetric(
		c.HuaweiSuccess,
		prometheus.CounterValue,
		float64(status.StatStorage.GetHuaweiSuccess()),
	)
	ch <- prometheus.MustNewConstMetric(
		c.HuaweiError,
		prometheus.CounterValue,
		float64(status.StatStorage.GetHuaweiError()),
	)
	ch <- prometheus.MustNewConstMetric(
		c.QueueUsage,
		prometheus.GaugeValue,
		float64(c.GetQueueUsage()),
	)
}
