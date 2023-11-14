package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//type MyInterface interface {}
//
//type Example struct {
//	Value string
//}
//
//func example1() MyInterface {
//	var e *int
//	//e.Value = "1"
//	return e
//}
//
//func example2() MyInterface {
//	var i MyInterface
//	return i
//}

//func main() {
//	fmt.Println(example1() == example2())
//}

type ss1 struct{}
type ss2 struct{}

func (s1 ss1) f() {
	fmt.Println("1")
}
func (s2 ss2) f() {
	fmt.Println("2")
}

type s3 struct {
	ss1
	ss2
}

func newVar(i []string) {
	q := "3"
	fmt.Println(len(i), cap(i))
	i = append(i, q)
	fmt.Println(len(i), cap(i))
}
func main() {
	//q := test_inc.TestStruct{TestField: 1}
	//fmt.Println("start main", q.TestField)
	//type Shape interface {
	//	Area() float64
	//}
	//
	//// Значение интерфейса
	//var s Shape
	//fmt.Println(s == nil)
	//fmt.Println(nil)
	//
	//fmt.Println(example1())
	//fmt.Println(example1() == example2())

	//qwe := s3{}
	//qwe.f()
	//i := make([]string, 0, 10)
	//newVar(i)
	//fmt.Println(len(i), cap(i))
	//fmt.Println(i)
	//var i int
	//if i == 0 {
	//	fmt.Println(i)
	//}

	//appendToSliceWithoutPointer()
	//appendToSliceSlice()
	//multipleAppend()
	//changeVar()
	//waitGroupWithArr()
	//multipleSleepGoroutine()
	//interfaceType()
	//channels()
	//moreChannels()
	//mutexLockUnlock()
	//contextWithCancel()
	//contextWithTimeout()
	//interfacesNumbers()
	//panicRecover()
}

func appendToSliceWithoutPointer() {
	qwe := make([]int, 1, 10)
	qwe[0] = 1
	func(qwe []int) {
		qwe = append(qwe, 2)
		fmt.Println(len(qwe), cap(qwe), qwe)

	}(qwe)

	// вернет [1] тк длина увеличивается в функции но не передается обратно (return либо &)
	// и будет ссылаться на qwe := make([]int, 1, 10) где len == 1
	fmt.Println(len(qwe), cap(qwe), qwe)
}

func appendToSliceSlice() {
	bibi := []int{1, 2, 3, 4, 5, 6}

	//qwe := bibi[1:2]
	//fmt.Println(len(qwe), cap(qwe), qwe)

	qwe := bibi[3:4]
	fmt.Println(len(qwe), cap(qwe), qwe)

	qwe = append(qwe, 7, 7, 7)

	fmt.Println(len(qwe), cap(qwe), qwe)
	fmt.Println(len(bibi), cap(bibi), bibi)
}

func multipleAppend() {
	var a []int
	var b []int
	var c []int
	a = append(a, 1)
	a = append(a, 2)
	b = append(a, 4)
	c = append(a, 5)
	c = append(b, 6)
	fmt.Println(a, b, c)

}

func changeVar() {
	a := 1
	func(a *int) {
		b := 2
		a = &b
	}(&a)
	fmt.Println(a)
}

func waitGroupWithArr() {
	var wg sync.WaitGroup
	var arr [3]int

	for i := 0; i != 3; i++ {
		wg.Add(1)

		time.Sleep(1 * time.Second)

		go func(wg *sync.WaitGroup, i int, arr *[3]int) {
			fmt.Printf("work %d\n", i)
			arr[i] = i
			defer wg.Done()
		}(&wg, i, &arr)
	}

	wg.Wait()
	fmt.Println("done")
	fmt.Println(arr)
}

func multipleSleepGoroutine() {
	tStart := time.Now().Unix()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(wg *sync.WaitGroup) {
			time.Sleep(2 * time.Second)
			defer wg.Done()
		}(&wg)
	}
	wg.Wait()

	tStop := time.Now().Unix()
	fmt.Println(tStop - tStart)
}

func interfaceType() {
	m := make(map[string]interface{})
	m["one"] = 1
	m["two"] = "bibi"
	m["thee"] = true

	for key, val := range m {
		switch val.(type) {
		case int:
			fmt.Printf("key %s has value of type int\n", key)
		case string:
			fmt.Printf("key %s has value of type string\n", key)
		case bool:
			fmt.Printf("key %s has value of type bool\n", key)
		}
	}
}

func channels() {
	//ch := make(chan int, 1)
	//ch <- 1
	//
	//select {
	//case val := <-ch:
	//	fmt.Println(val)
	//default:
	//	fmt.Println("channel is closed")
	//}
	//
	//ch = nil // nil channel
	////close(ch) // close channel
	//
	//select {
	//case ch <- 2:
	//default:
	//	fmt.Println("handle panic")
	//}
	//
	//select {
	//case val2 := <-ch:
	//	fmt.Println(val2)
	//default:
	//	fmt.Println("channel is closed")
	//}

	chann := make(chan int, 20)

	go func() {
		for i := 0; i < 30; i++ {
			chann <- i
		}
		close(chann)
	}()

	go func() {
		time.Sleep(2 * time.Second)
	loop:
		for {
			select {
			case val, ok := <-chann:
				if ok {
					fmt.Println(val)
				}
			default:
				if len(chann) == 0 {
					fmt.Println("channel is empty")
					time.Sleep(1 * time.Second)
				} else {
					fmt.Println("other reason")
					break loop
				}
			}
		}
	}()

	time.Sleep(7 * time.Second)
}

func moreChannels() {
	ch := make(chan int, 3)

	go func(ch chan int) {
		for i := 0; i < 3; i++ {
			select {
			case ch <- i:
			default:
				fmt.Println("channel is closed")
			}
		}
	}(ch)

	var count int
	for {
		if count == 3 {
			break
		}
		time.Sleep(1 * time.Second)

		select {
		case val := <-ch:
			fmt.Println(val)
			count++
		default:
			fmt.Println("channel is closed")
		}
	}
}

func mutexLockUnlock() {
	slc := make([]int, 0, 3)
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(wg *sync.WaitGroup, i int, slc *[]int) {
			var mutex sync.Mutex
			mutex.Lock()
			*slc = append(*slc, i)
			mutex.Unlock()

			defer wg.Done()
		}(&wg, i, &slc)
	}
	wg.Wait()
	fmt.Println(slc)
}

func contextWithCancel() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	ch := make(chan int, 3)

	go func(ctx context.Context, ch chan int) {
	loop:
		for {
			time.Sleep(1 * time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("context done")
				break loop
			case val := <-ch:
				fmt.Println(val)
			default:
				if len(ch) == 0 {
					fmt.Println("channel is empty")
				}
			}
		}
	}(ctx, ch)

	//ch = nil

	for i := 0; i < 3; i++ {
		select {
		case ch <- i:
			fmt.Printf("i - %d\n", i)
		default:
			fmt.Println("channel is closed")
		}
	}

	time.Sleep(5 * time.Second)

	cancel()

	time.Sleep(2 * time.Second)
}

func contextWithTimeout() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	ch := make(chan int, 0)

	go func(ctx context.Context, ch chan int) {
		for {
			time.Sleep(1 * time.Second)

			select {
			case <-ctx.Done():
				fmt.Println("context done")
			case val := <-ch:
				fmt.Println(val)
			default:
				if len(ch) == 0 {
					fmt.Println("channel is empty")
				}
			}
		}
	}(ctx, ch)

	ch <- 1
	ch <- 2

	time.Sleep(5 * time.Second)
}

type DBConn interface {
	NewConn() string
}

type DB1 struct {
	desc string
}

type DB2 struct {
	desc string
}

func (db1 DB1) NewConn() string {
	return fmt.Sprintf("connected to: %s", db1.desc)
}

func (db2 DB2) NewConn() string {
	return fmt.Sprintf("connected to: %s", db2.desc)
}

func doConn(db DBConn) {
	fmt.Println(db.NewConn())
}

func interfacesNumbers() {
	base1 := DB1{
		desc: "postgres",
	}
	base2 := DB2{
		desc: "mongo",
	}

	var connection DBConn
	if 2/2 == 1 {
		connection = base1
	} else {
		connection = base2
	}

	fmt.Println(connection.NewConn())

	// or
	var connectionOr DBConn
	if 2/2 == 1 {
		connectionOr = base1
	} else {
		connectionOr = base2
	}
	doConn(connectionOr)
}

func panicRecover() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recovered")
		}
	}()

	ch := make(chan int)
	close(ch)
	ch <- 1
}
