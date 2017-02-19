package stats_api

import (
	"encoding/json"
	"io"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// Stats represents activity status of Go.
type Stats struct {
	Time int64 `json:"time"`
	// runtime
	GoVersion    string `json:"go_version"`
	GoOs         string `json:"go_os"`
	GoArch       string `json:"go_arch"`
	CpuNum       int    `json:"cpu_num"`
	GoroutineNum int    `json:"goroutine_num"`
	Gomaxprocs   int    `json:"gomaxprocs"`
	CgoCallNum   int64  `json:"cgo_call_num"`
	// memory
	MemoryAlloc      uint64 `json:"memory_alloc"`
	MemoryTotalAlloc uint64 `json:"memory_total_alloc"`
	MemorySys        uint64 `json:"memory_sys"`
	MemoryLookups    uint64 `json:"memory_lookups"`
	MemoryMallocs    uint64 `json:"memory_mallocs"`
	MemoryFrees      uint64 `json:"memory_frees"`
	// stack
	StackInUse uint64 `json:"memory_stack"`
	// heap
	HeapAlloc    uint64 `json:"heap_alloc"`
	HeapSys      uint64 `json:"heap_sys"`
	HeapIdle     uint64 `json:"heap_idle"`
	HeapInuse    uint64 `json:"heap_inuse"`
	HeapReleased uint64 `json:"heap_released"`
	HeapObjects  uint64 `json:"heap_objects"`
	// garbage collection
	GcNext           uint64    `json:"gc_next"`
	GcLast           uint64    `json:"gc_last"`
	GcNum            uint32    `json:"gc_num"`
	GcPerSecond      float64   `json:"gc_per_second"`
	GcPausePerSecond float64   `json:"gc_pause_per_second"`
	GcPause          []float64 `json:"gc_pause"`
}

var lastSampleTime time.Time
var lastPauseNs uint64 = 0
var lastNumGc uint32 = 0

var nsInMs float64 = float64(time.Millisecond)

var statsMux sync.Mutex

func GetStats() *Stats {
	statsMux.Lock()
	defer statsMux.Unlock()

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	now := time.Now()

	var gcPausePerSecond float64

	if lastPauseNs > 0 {
		pauseSinceLastSample := mem.PauseTotalNs - lastPauseNs
		gcPausePerSecond = float64(pauseSinceLastSample) / nsInMs
	}

	lastPauseNs = mem.PauseTotalNs

	countGc := int(mem.NumGC - lastNumGc)

	var gcPerSecond float64

	if lastNumGc > 0 {
		diff := float64(countGc)
		diffTime := now.Sub(lastSampleTime).Seconds()
		gcPerSecond = diff / diffTime
	}

	if countGc > 256 {
		// lagging GC pause times
		countGc = 256
	}

	gcPause := make([]float64, countGc)

	for i := 0; i < countGc; i++ {
		idx := int((mem.NumGC-uint32(i))+255) % 256
		pause := float64(mem.PauseNs[idx])
		gcPause[i] = pause / nsInMs
	}

	lastNumGc = mem.NumGC
	lastSampleTime = time.Now()

	return &Stats{
		Time:         now.UnixNano(),
		GoVersion:    runtime.Version(),
		GoOs:         runtime.GOOS,
		GoArch:       runtime.GOARCH,
		CpuNum:       runtime.NumCPU(),
		GoroutineNum: runtime.NumGoroutine(),
		Gomaxprocs:   runtime.GOMAXPROCS(0),
		CgoCallNum:   runtime.NumCgoCall(),
		// memory
		MemoryAlloc:      mem.Alloc,
		MemoryTotalAlloc: mem.TotalAlloc,
		MemorySys:        mem.Sys,
		MemoryLookups:    mem.Lookups,
		MemoryMallocs:    mem.Mallocs,
		MemoryFrees:      mem.Frees,
		// stack
		StackInUse: mem.StackInuse,
		// heap
		HeapAlloc:    mem.HeapAlloc,
		HeapSys:      mem.HeapSys,
		HeapIdle:     mem.HeapIdle,
		HeapInuse:    mem.HeapInuse,
		HeapReleased: mem.HeapReleased,
		HeapObjects:  mem.HeapObjects,
		// garbage collection
		GcNext:           mem.NextGC,
		GcLast:           mem.LastGC,
		GcNum:            mem.NumGC,
		GcPerSecond:      gcPerSecond,
		GcPausePerSecond: gcPausePerSecond,
		GcPause:          gcPause,
	}
}

var newLineTerm bool = false
var prettyPrint bool = false

// NewLineTermEnabled enable termination with newline for response body.
func NewLineTermEnabled() {
	newLineTerm = true
}

// NewLineTermDisabled disable termination with newline for response body.
func NewLineTermDisabled() {
	newLineTerm = false
}

// PrettyPrintEnabled enable pretty-print for response body.
func PrettyPrintEnabled() {
	prettyPrint = true
}

// PrettyPrintDisabled disable pritty-print for response body.
func PrettyPrintDisabled() {
	prettyPrint = false
}

// Handler returns activity status of Go.
func Handler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	for _, c := range []string{"1", "true"} {
		if values.Get("pp") == c {
			prettyPrint = true
		}
	}

	var jsonBytes []byte
	var jsonErr error
	if prettyPrint {
		jsonBytes, jsonErr = json.MarshalIndent(GetStats(), "", "  ")
	} else {
		jsonBytes, jsonErr = json.Marshal(GetStats())
	}
	var body string
	if jsonErr != nil {
		body = jsonErr.Error()
	} else {
		body = string(jsonBytes)
	}

	if newLineTerm {
		body += "\n"
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Content-Length"] = strconv.Itoa(len(body))
	for name, value := range headers {
		w.Header().Set(name, value)
	}

	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	io.WriteString(w, body)
}
