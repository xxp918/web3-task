package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)
func main() {
	// number := MyInt{value: 10}
	// number.Increment()
	// fmt.Println(number.value)

	// slice := []int{1, 2, 3, 4, 5}
	// multiplyByTwo(slice)
	// fmt.Println(slice)
	// var wg sync.WaitGroup
	// ji_ou(&wg)
	// //等待所有协程完成
	// wg.Wait()
	//任务调度模拟
	// tasks := []Task{
	// 	func() {
	// 	fmt.Println("task0调用中")
	// }, func() {	
	// 	fmt.Println("task1调用中")
	// },func ()  {
	// 	fmt.Println("task2调用中")
	// }}
	// results :=taskScheduler(tasks)
	// for _, result := range results {
	// 	fmt.Println(result)
	// }
	//实例化长方体长宽，和圆形半径
	// rectangle := Rectangle{Width: 5, Height: 3}
	// circle := Circle{Radius: 3}
	// //计算长方体面积
	// rectangleArea := rectangle.Area()
	// fmt.Println("长方体面积:", rectangleArea)
	// //计算长方体周长
	// rectanglePerimeter := rectangle.Perimeter()
	// fmt.Println("长方体周长:", rectanglePerimeter)
	// //计算圆形面积
	// circleArea := circle.Area()
	// fmt.Println("圆形面积:", circleArea)
	// //计算圆形周长
	// circlePerimeter := circle.Perimeter()
	// fmt.Println("圆形周长：", circlePerimeter)
	// //实例接口
	// var s Shape
	// s = Rectangle{Width: 5, Height: 3}
	// fmt.Printf("长方体面积:%.2f, 周长:%.2f\n", s.Area(),s.Perimeter())
	// s = Circle{Radius: 3}
	// fmt.Printf("圆形面积:%.2f, 周长:%.2f\n", s.Area(),s.Perimeter())

	//输出员工的信息
	// employee := Employee{
	// 	Person: Person{
	// 		Name: "张三",
	// 		Age:  25,
	// 	},
	// 	EmployeeID: 1001,
	// }
	// employee.PrintInfo()
	//channle
	//channelFunc()
	
	// counter()
	atomicCounter()

}
//锁机制
//编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func counter() {
	var count int
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(10)
	for i := 0; i < 10; i++ {		
		go func ()  {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
		}()
	}
	//等待
	wg.Wait()
	fmt.Println("计数器值：", count)
}


//使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

func atomicCounter() {
	var count int64
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func ()  {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("计数器值：", count)
}

//Channel
//编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
//实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
func channelFunc(){
	var wg sync.WaitGroup
	//创建一个缓冲为10的通道
	ch := make(chan int,10)
	wg.Add(2)
	//启动一个协程生成整数 生产者
	go func ()  {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	//启动另一个协程接收整数并打印 消费者
	go func ()  {
		defer wg.Done()
		for num := range ch{
			fmt.Println(num)
		}
	}()
	//等待所有协程完成
	wg.Wait()

}
//面向对象
//使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeID int
}
func (e Employee) PrintInfo() {
	fmt.Printf("员工姓名：%s, 年龄：%d, 员工ID：%d\n", e.Name, e.Age, e.EmployeeID)
}

//定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}
//长方体面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
//长方体周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
//圆面积
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
//圆周长
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}
//指针
//编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
type MyInt struct {
	value int
}
func (mi *MyInt) Increment() {
	mi.value += 10
}

//实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func multiplyByTwo(slice []int) {
	for i := 0; i < len(slice); i++ {
		(slice)[i] *= 2
	}
}


//Goroutine
//编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func ji_ou(wg *sync.WaitGroup){
	wg.Add(2)
	fmt.Println("hello world")
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println("输出奇数",i)
		}
	}()
	go func ()  {
		defer wg.Done()
		for i := 2; i <= 10; i+=2 {
			fmt.Println("输出偶数",i)
		}
	}()
}
//设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 定义任务类型
type Task func()
// 定义任务结构体
type TaskResult struct {
	ID int
	Duration time.Duration
}
func taskScheduler(tasks []Task) []TaskResult{
	// 创建一个通道用于接收任务结果
	resultChan := make(chan TaskResult,len(tasks))
	results := make([]TaskResult,len(tasks))
	var wg sync.WaitGroup
	for id, task := range tasks {
		wg.Add(1)
		// 启动一个协程执行任务
		go func(i int , t Task){
			defer wg.Done()
			start := time.Now()
			//调用任务函数
			t()
			/// 计算任务执行时间间隔 并发送结果到通道
			resultChan <- TaskResult{ID: i, Duration: time.Since(start)}
			// 输出任务执行完成信息
			fmt.Printf("id=%d任务执行完成\n", i)
		}(id,task)	}
	//设置结束条件关闭通道
	go func() {
		wg.Wait()
		close(resultChan)
	}()
	//results 收集结果 等待所有任务完成
	for result := range resultChan {
		results[result.ID] = result
	}
	return results
}