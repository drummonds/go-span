package span_test

import (
	"fmt"

	span "git.bytestone.uk/hum3/go-span"
)

func Example() {
	fmt.Println(span.Format(6411360895289))
	fmt.Println(span.Format(162345678))
	fmt.Println(span.Days(388))
	// Output:
	// 1h46m51s
	// 162.3ms
	// 1y23d
}
