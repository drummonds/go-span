# Changelog

All notable changes to this project are documented here. The format is
based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

## [0.2.0] - 2026-10-17

 - Fix fatal stack overflow on the minimum signed value and day rounding of the largest spans

### Fixed

- `Format(math.MinInt64)` and `Days(math.MinInt)` no longer recurse until
  a fatal stack overflow (#2).
- `Format` no longer rounds the largest spans down a day: the maximum
  Duration is `292y172d`, not `292y171d` (#3).

## [0.1.0] - 2026-10-17

 - First release: span.Format and span.Days, a duration only as precise as a reader needs

### Added

- `Format`: a duration only as precise as a reader needs, about four
  significant figures, from `162ns` through `16.23s`, `2m3.5s`,
  `1h46m51s` and `1d4h33m` to `1y23d`.
- `Days`: a count of days in days, or years and days from 365 on.
