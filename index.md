# go-span

How long something took, only as precisely as a reader needs.
`time.Duration.String()` prints every nanosecond, `1h46m51.360895289s`,
which is right for a timestamp and noise for a run that took the best
part of two hours. `span.Format` keeps about four significant figures and
lets each larger unit that comes in push a smaller one out.

```go
import span "git.bytestone.uk/hum3/go-span"

span.Format(6411360895289) // "1h46m51s"
span.Format(162345678)     // "162.3ms"
span.Days(388)             // "1y23d"
```

Pure Go, standard library only.

## The rule

| Span | Precision | Example |
|---|---|---|
| below 1µs | whole nanoseconds, the finest unit | `162ns` |
| 1µs to 1m | 4 significant figures | `16.23µs`, `162.3ms`, `1.623s`, `16.23s` |
| 1m to 1h | seconds to one decimal | `2m3.5s` |
| 1h to 1d | whole seconds | `1h46m51s` |
| 1d to 10d | days, hours, minutes | `1d4h33m` |
| 10d to 1y | days, hours | `12d4h` |
| 1y and up (a year is 365 days) | years, days | `1y23d` |

Rounding that carries a span into the next tier is shown at that tier's
precision: 59.996s is `1m0s`, not `60s`. Below a day the output is a Go
duration string, so `time.ParseDuration` reads it back; days and years are
not Go units. Trailing zeros go as they do in Go (`1.6s`, not `1.600s`).
A negative span is the positive one with a minus sign.

`span.Days(n)` is for a count whose unit is the day, such as simulated
days run: `340d` below a year and `1y23d` from 365 on, never hours.

## Compared with go-humanize

[dustin/go-humanize](https://github.com/dustin/go-humanize) answers a
different question. Its `Time` and `RelTime` say roughly *when*, relative
to now, in one English unit; go-span says *how long*, compactly and with
the precision kept. go-humanize has no formatter for a duration's length.
Here `humanize.RelTime(t, t.Add(d), "", "")` (v1.1.0) beside
`span.Format(d)`:

| `d.String()` | `span.Format(d)` | go-humanize |
|---|---|---|
| `162ns` | `162ns` | now |
| `162.345ms` | `162.3ms` | now |
| `16.234567s` | `16.23s` | 16 seconds |
| `2m3.456s` | `2m3.5s` | 2 minutes |
| `1h46m51.360895289s` | `1h46m51s` | 1 hour |
| `28h32m40s` | `1d4h33m` | 1 day |
| `292h20m0s` | `12d4h` | 1 week |
| `9317h0m0s` | `1y23d` | 1 year |

Use go-humanize for "3 days ago" on a timeline, and its `Bytes` and
`Comma` for sizes and counts; go-span for a column of how long runs,
requests or stages took, where 1 hour against 1 hour 59 minutes matters.
