# Roadmap

Each item is an open issue labelled `roadmap` on the Forgejo, grouped
into named epics (task-plus ADR 0004); `tp roadmap check` keeps file and
forge in step.

## Adopt go-span where spans are formatted by hand

Five hand-rolled duration formatters across four projects each round
differently. Replace each with `span.Format` (or `span.Days`), one
project per item, the change made in that project. Where a page
deliberately shows less precision than the rule (a coarse estimate of time
remaining, say), the item decides whether the rule or the page is right
before swapping.
