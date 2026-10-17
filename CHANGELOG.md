# Changelog

All notable changes to this project are documented here. The format is
based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

### Fixed

- `Format(math.MinInt64)` and `Days(math.MinInt)` no longer recurse until
  a fatal stack overflow. `Days` gives the exact magnitude; `Format` gives
  the maximum Duration's, the nearest that fits (#2).

## [0.1.0] - 2026-10-17

 - First release: span.Format and span.Days, a duration only as precise as a reader needs

### Added

- `Format`: a duration only as precise as a reader needs, about four
  significant figures, from `162ns` through `16.23s`, `2m3.5s`,
  `1h46m51s` and `1d4h33m` to `1y23d`.
- `Days`: a count of days in days, or years and days from 365 on.
