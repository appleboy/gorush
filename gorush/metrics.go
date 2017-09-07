package gorush

import (
	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "gorush_"

// Metrics implements the prometheus.Metrics interface and
// exposes gorush metrics for prometheus
type Metrics struct {
	TotalPushCount      *prometheus.Desc
	TotalIosSuccess     *prometheus.Desc
	TotalIosError       *prometheus.Desc
	TotalAndroidSuccess *prometheus.Desc
	TotalAndroidError   *prometheus.Desc
	TotalWebSuccess     *prometheus.Desc
	TotalWebError       *prometheus.Desc
	PushCount           *prometheus.Desc
	IosSuccess          *prometheus.Desc
	IosError            *prometheus.Desc
	AndroidSuccess      *prometheus.Desc
	AndroidError        *prometheus.Desc
	WebSuccess          *prometheus.Desc
	WebError            *prometheus.Desc
}

var lastTotalPushCount int64 = 0
var lastTotalIosSuccess int64 = 0
var lastTotalIosError int64 = 0
var lastTotalAndroidSuccess int64 = 0
var lastTotalAndroidError int64 = 0
var lastTotalWebSuccess int64 = 0
var lastTotalWebError int64 = 0

// NewMetrics returns a new Metrics with all prometheus.Desc initialized
func NewMetrics() Metrics {
	return Metrics{
		TotalPushCount: prometheus.NewDesc(
			namespace+"total_push_count",
			"Number of total push count",
			nil, nil,
		),
		TotalIosSuccess: prometheus.NewDesc(
			namespace+"total_ios_success",
			"Number of total iOS success count",
			nil, nil,
		),
		TotalIosError: prometheus.NewDesc(
			namespace+"total_ios_fail",
			"Number of total iOS fail count",
			nil, nil,
		),
		TotalAndroidSuccess: prometheus.NewDesc(
			namespace+"total_android_success",
			"Number of total android success count",
			nil, nil,
		),
		TotalAndroidError: prometheus.NewDesc(
			namespace+"total_android_fail",
			"Number of total android fail count",
			nil, nil,
		),
		TotalWebSuccess: prometheus.NewDesc(
			namespace+"total_web_success",
			"Number of total web success count",
			nil, nil,
		),
		TotalWebError: prometheus.NewDesc(
			namespace+"total_web_fail",
			"Number of total web fail count",
			nil, nil,
		),
		PushCount: prometheus.NewDesc(
			namespace+"push_count",
			"Number of push count",
			nil, nil,
		),
		IosSuccess: prometheus.NewDesc(
			namespace+"ios_success",
			"Number of iOS success count",
			nil, nil,
		),
		IosError: prometheus.NewDesc(
			namespace+"ios_fail",
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
		WebSuccess: prometheus.NewDesc(
			namespace+"web_success",
			"Number of web success count",
			nil, nil,
		),
		WebError: prometheus.NewDesc(
			namespace+"web_fail",
			"Number of web fail count",
			nil, nil,
		),
	}
}

// Describe returns all possible prometheus.Desc
func (c Metrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.TotalPushCount
	ch <- c.TotalIosSuccess
	ch <- c.TotalIosError
	ch <- c.TotalAndroidSuccess
	ch <- c.TotalAndroidError
	ch <- c.TotalWebSuccess
	ch <- c.TotalWebError
	ch <- c.PushCount
	ch <- c.IosSuccess
	ch <- c.IosError
	ch <- c.AndroidSuccess
	ch <- c.AndroidError
	ch <- c.WebSuccess
	ch <- c.WebError
}

// Collect returns the metrics with values
func (c Metrics) Collect(ch chan<- prometheus.Metric) {
	var newTotalPushCount = StatStorage.GetTotalCount()
	var newTotalIosSuccess = StatStorage.GetIosSuccess()
	var newTotalIosError = StatStorage.GetIosError()
	var newTotalAndroidSuccess = StatStorage.GetAndroidSuccess()
	var newTotalAndroidError = StatStorage.GetAndroidError()
	var newTotalWebSuccess = StatStorage.GetWebSuccess()
	var newTotalWebError = StatStorage.GetWebError()
	var pushCount = newTotalPushCount - lastTotalPushCount
	var iosSuccess = newTotalIosSuccess - lastTotalIosSuccess
	var iosError = newTotalIosError - lastTotalIosError
	var androidSuccess = newTotalAndroidSuccess - lastTotalAndroidSuccess
	var androidError = newTotalAndroidError - lastTotalAndroidError
	var webSuccess = newTotalWebSuccess - lastTotalWebSuccess
	var webError = newTotalWebError - lastTotalWebError

	lastTotalPushCount = newTotalPushCount
	lastTotalIosSuccess = newTotalIosSuccess
	lastTotalIosError = newTotalIosError
	lastTotalAndroidSuccess = newTotalAndroidSuccess
	lastTotalAndroidError = newTotalAndroidError
	lastTotalWebSuccess = newTotalWebSuccess
	lastTotalWebError = newTotalWebError

	ch <- prometheus.MustNewConstMetric(
		c.TotalPushCount,
		prometheus.CounterValue,
		float64(newTotalPushCount),
	)
	ch <- prometheus.MustNewConstMetric(
		c.TotalIosSuccess,
		prometheus.CounterValue,
		float64(newTotalIosSuccess),
	)
	ch <- prometheus.MustNewConstMetric(
		c.TotalIosError,
		prometheus.CounterValue,
		float64(newTotalIosError),
	)
	ch <- prometheus.MustNewConstMetric(
		c.TotalAndroidSuccess,
		prometheus.CounterValue,
		float64(newTotalAndroidSuccess),
	)
	ch <- prometheus.MustNewConstMetric(
		c.TotalAndroidError,
		prometheus.CounterValue,
		float64(newTotalAndroidError),
	)
	ch <- prometheus.MustNewConstMetric(
		c.TotalWebSuccess,
		prometheus.CounterValue,
		float64(newTotalWebSuccess),
	)
	ch <- prometheus.MustNewConstMetric(
		c.TotalWebError,
		prometheus.CounterValue,
		float64(newTotalWebError),
	)
	ch <- prometheus.MustNewConstMetric(
		c.PushCount,
		prometheus.GaugeValue,
		float64(pushCount),
	)
	ch <- prometheus.MustNewConstMetric(
		c.IosSuccess,
		prometheus.GaugeValue,
		float64(iosSuccess),
	)
	ch <- prometheus.MustNewConstMetric(
		c.IosError,
		prometheus.GaugeValue,
		float64(iosError),
	)
	ch <- prometheus.MustNewConstMetric(
		c.AndroidSuccess,
		prometheus.GaugeValue,
		float64(androidSuccess),
	)
	ch <- prometheus.MustNewConstMetric(
		c.AndroidError,
		prometheus.GaugeValue,
		float64(androidError),
	)
	ch <- prometheus.MustNewConstMetric(
		c.WebSuccess,
		prometheus.GaugeValue,
		float64(webSuccess),
	)
	ch <- prometheus.MustNewConstMetric(
		c.WebError,
		prometheus.GaugeValue,
		float64(webError),
	)
}
