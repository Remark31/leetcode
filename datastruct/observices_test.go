package datastruct

import (
	"fmt"
	"testing"
	"time"
)

var pp Property

func init() {
	pp = NewProperty(0)
}

func produce(t *testing.T) {
	tick := time.NewTicker(20 * time.Second)
	i := 0
	for i < 100 {
		select {
		case <-tick.C:
			return
		default:
			time.Sleep(1 * time.Second)
			i += 1
			pp.Update(i)
		}
	}
}

func consume(s string, t *testing.T) {
	a := pp.Observe()
	if s == "C" {
		time.Sleep(3 * time.Second)
	}
	for {
		select {
		case <-a.Changes():
			a.Next()
			v := a.Value()
			fmt.Println("name", s, "value", v)
		}
	}
}

func TestConsume(t *testing.T) {

	go produce(t)

	go consume("A", t)

	time.Sleep(3 * time.Second)
	go consume("B", t)
	go consume("C", t)

	for i := 0; i < 1000; i++ {
		time.Sleep(1 * time.Second)
	}
}
