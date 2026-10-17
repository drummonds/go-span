# go-span

How long something took, only as precisely as a reader needs:
`1h46m51s`, not `1h46m51.360895289s`; `162.3ms`, `16.23s`, `1d4h33m`,
`1y23d`. About four significant figures, standard library only.

```go
import span "git.bytestone.uk/hum3/go-span"

span.Format(d) // a time.Duration
span.Days(n)   // a count of days
```

The precision rule, and a comparison with
[dustin/go-humanize](https://github.com/dustin/go-humanize), are in the
[documentation](https://go-span.docs.bytestone.uk/).

## Links

- Documentation: https://go-span.docs.bytestone.uk/
- Source: https://git.bytestone.uk/hum3/go-span
- Mirror: https://github.com/drummonds/go-span
