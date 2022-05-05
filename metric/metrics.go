package metric

import (
	"github.com/appleboy/gorush/status"

	"github.com/golang-queue/queue"
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
	BusyWorkers    *prometheus.Desc
	SuccessTasks   *prometheus.Desc
	FailureTasks   *prometheus.Desc
	SubmittedTasks *prometheus.Desc
	q              *queue.Queue
}

// NewMetrics returns a new Metrics with all prometheus.Desc initialized
func NewMetrics(q *queue.Queue) Metrics {
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
		BusyWorkers: prometheus.NewDesc(
			namespace+"busy_workers",
			"Length of busy workers",
			nil, nil,
		),
		FailureTasks: prometheus.NewDesc(
			namespace+"failure_tasks",
			"Length of Failure Tasks",
			nil, nil,
		),
		SuccessTasks: prometheus.NewDesc(
			namespace+"success_tasks",
			"Length of Success Tasks",
			nil, nil,
		),
		SubmittedTasks: prometheus.NewDesc(
			namespace+"submitted_tasks",
			"Length of Submitted Tasks",
			nil, nil,
		),
		q: q,
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
	ch <- c.BusyWorkers
	ch <- c.SuccessTasks
	ch <- c.FailureTasks
	ch <- c.SubmittedTasks
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
		c.BusyWorkers,
		prometheus.GaugeValue,
		float64(c.q.BusyWorkers()),
	)
	ch <- prometheus.MustNewConstMetric(
		c.SuccessTasks,
		prometheus.CounterValue,
		float64(c.q.SuccessTasks()),
	)
	ch <- prometheus.MustNewConstMetric(
		c.FailureTasks,
		prometheus.CounterValue,
		float64(c.q.FailureTasks()),
	)
	ch <- prometheus.MustNewConstMetric(
		c.SubmittedTasks,
		prometheus.CounterValue,
		float64(c.q.SubmittedTasks()),
	)
}
