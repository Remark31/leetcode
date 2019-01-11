package pprofuse

import (
	"fmt"
	"time"
)

var pp Property

func init() {
	pp = NewProperty(0)
}

func produce() {
	// 20秒内循环，超过后结束
	t := time.NewTicker(10 * time.Second)
	i := 0
	for {
		select {
		case <-t.C:
			return
		default:
			time.Sleep(1 * time.Second)
			i += 1
			pp.Update(i)
		}
	}
}

func consume(name string) {
	// 消费队列中内容并且输出
	n := pp.Observe()
	for {
		select {
		case <-n.Changes():
			n.Next()
			v := n.Value()
			fmt.Println("name: ", name, "v: ", v)
		}
	}

}

func Use() {
	go produce()

	for i := 0; i < 5; i++ {
		go consume(string('a' + i))
		time.Sleep(1 * time.Second)
	}

	for i := 0; i < 5; i++ {
		go func() {
			n := pp.Observe()
			for {
				select {
				case <-n.Changes():
					n.Next()
					v := n.Value()
					fmt.Println("name: ", string('A'+i), "v: ", v)
				}
			}
		}()
	}

	notend := make(chan struct{})
	<-notend
}
