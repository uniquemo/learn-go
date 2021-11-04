package obj_pool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10) // 生产者生产10个对象

	// if err := pool.PutObj(&ReusableObj{}); err != nil { // 尝试放置超出池大小的对象
	// 	t.Error(err)
	// }

	for i := 1; i <= 11; i++ { // 消费者消费11个对象
		if v, err := pool.GetObj(time.Second * 1); err != nil { // 消费者
			fmt.Println("111111")
			t.Error(err)
		} else {
			fmt.Printf("%T-%d\n", v, i)
			// 生产者，池子空了，放置对象进对象池以供消费者使用
			if err := pool.PutObj(v); err != nil {
				t.Error(err)
			}
		}
	}

	fmt.Println("Done")
}
