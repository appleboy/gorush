golang-stats-api-handler
========================

Golang cpu, memory, gc, etc information api handler.

## Install

    go get github.com/fukata/golang-stats-api-handler

## Example

```go
import (
    "net/http"
    "log"
    "github.com/fukata/golang-stats-api-handler"
)
func main() {
    http.HandleFunc("/api/stats", stats_api.Handler)
    log.Fatal( http.ListenAndServe(":8080", nil) )
}
```

## Response

    $ curl -i http://localhost:8080/api/stats/
    HTTP/1.1 200 OK
    Content-Length: 712
    Content-Type: application/json
    Date: Sun, 23 Aug 2015 16:52:13 GMT
    
    {
        "time": 1440348733548339479,
        "go_version": "go1.5",
        "go_os": "darwin",
        "go_arch": "amd64",
        "cpu_num": 8,
        "goroutine_num": 24,
        "gomaxprocs": 8,
        "cgo_call_num": 9,
        "memory_alloc": 3974536,
        "memory_total_alloc": 12857888,
        "memory_sys": 12871928,
        "memory_lookups": 52,
        "memory_mallocs": 144922,
        "memory_frees": 118936,
        "memory_stack": 688128,
        "heap_alloc": 3974536,
        "heap_sys": 8028160,
        "heap_idle": 2170880,
        "heap_inuse": 5857280,
        "heap_released": 0,
        "heap_objects": 25986,
        "gc_next": 4833706,
        "gc_last": 1440348732827834419,
        "gc_num": 4,
        "gc_per_second": 0,
        "gc_pause_per_second": 0,
        "gc_pause": [
            0.196828,
            2.027442,
            0.181887,
            0.312866
        ]
    }

|Key                |Value|
|-------------------|----------------|
|time               |unix timestamp as nano-seconds|
|go_version         |runtime.Version()|
|go_os              |runtime.GOOS|
|go_arch            |runtime.GOARCH|
|cpu_num            |number of CPUs|
|goroutine_num      |number of goroutines|
|gomaxprocs         |runtime.GOMAXRPOCS(0)|
|cgo_call_num       |runtime.NumCgoCall()|
|memory_alloc       |bytes allocated and not yet freed|
|memory_total_alloc |bytes allocated (even if freed)|
|memory_sys         |bytes obtained from system|
|memory_lookups     |number of pointer lookups|
|memory_mallocs     |number of mallocs|
|memory_frees       |number of frees|
|memory_stack       |bytes used by stack allocator|
|heap_alloc         |bytes allocated and not yet freed (same as memory_alloc above)|
|heap_sys           |bytes obtained from system (not same as memory_sys)|
|heap_idle          |bytes in idle spans|
|heap_inuse         |bytes in non-idle span|
|heap_released      |bytes released to the OS|
|heap_objects       |total number of allocated objects|
|gc_next            |next collection will happen when HeapAlloc ≥ this amount|
|gc_last            |end time of last collection|
|gc_num             |number of GC-run|
|gc_per_second      |number of GC-run per second|
|gc_pause_per_second|pause duration by GC per seconds|
|gc_pause           |pause durations by GC|

## Plugins

- Zabbix: [fukata/golang-stats-api-handler-zabbix-userparameter](https://github.com/fukata/golang-stats-api-handler-zabbix-userparameter)
- Munin: [fukata/golang-stats-api-handler-munin-plugin](https://github.com/fukata/golang-stats-api-handler-munin-plugin) 
