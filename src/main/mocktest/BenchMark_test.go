package mocktest

import (
	"fmt"
	"github.com/hashicorp/go-version"
	"testing"
	"time"
)

func Test_test(t *testing.T) {
	fmt.Println("start")
	fmt.Println(time.Now())
	for i := 0; i < 10000; i++ {
		x, _ := version.NewVersion("9.3.1")
		y, _ := version.NewVersion("9.3.1")
		x.Equal(y)
	}
	fmt.Println("end")
	fmt.Println(time.Now())
}
