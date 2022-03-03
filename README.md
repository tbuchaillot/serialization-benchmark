# Tyk analytics serialization benchmark


## Description
Checking serialization performance + forward compatibility using [Tyk Pump analytics struct](https://github.com/TykTechnologies/tyk-pump/blob/ff09a45ae92d3c68f35cdd4ebb8e36d02e93760f/analytics/analytics.go#L31)


At this moment, msgpack/v2 is used and encoding/json for Tyk Analytics.

To test:
- gotiny
- bebop
- protobuf
- gogo/proto
- flatbuffers
- colfer?
- musgo?
- ?

It also utilizes randomized data generated in the demo pkg.

## How to run it?

```
make bench
```

## Latest
benchmark                                     | iter       | time/iter    | bytes/serial  | allocs/op  |     ns/alloc|
----------------------------------------------|------------|--------------|-----------|------------|--------|
Benchmark_JSON_Marshal-8                      | 748794     | 3225 ns/op   | 988.9 B/serial |      1637 B/op  |        4 allocs/op|
Benchmark_Msgpack_Marshal-8           |    757999        |      3336 ns/op      |         795.4 B/serial  |     2901 B/op      |   11 allocs/op|
Benchmark_GoTiny_Marshal-8            |   2090042          |    1148 ns/op        |       372.2 B/serial  |     1792 B/op     |     8 allocs/op|
Benchmark_NormalProto_Marshal-8        |  1690639          |    1416 ns/op         |      402.3 B/serial   |     834 B/op     |     2 allocs/op|
Benchmark_Bebop_Marshal-8              |    3170756            |     750.4 ns/op      |         550.7 B/serial      |    739 B/op     |       3 allocs/op |
Benchmark_MsgpGen_Marshal-8             |  3757410               | 647.1 ns/op           |   809.4 B/serial     |   1039 B/op        |   1 allocs/op |