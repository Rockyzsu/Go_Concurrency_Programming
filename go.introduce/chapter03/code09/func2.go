//模拟事件绑定机制
package main

type OnSumBefore func(int) int
type OnSum func(int, int) int
type OnSumEnd func(string)

var SumBeforeEvent OnSumBefore
var SumEvent OnSum
var SumEndEvent OnSumEnd

//StartSum 启动计算
func StartSum(a, b int, c string) int {
	t, f := 0, 0
	//判断释放绑定实现，并按事件执行顺序执行
	if SumBeforeEvent != nil {
		t = SumBeforeEvent(a)
	}
	if SumEvent != nil {
		f = SumEvent(t, b)
	}
	if SumEndEvent != nil {
		SumEndEvent(c)
	}

	return f
}

//RegEvent 注册事件实现
func RegEvent(f1 OnSumBefore, f2 OnSum, f3 OnSumEnd) {
	SumBeforeEvent = f1
	SumEvent = f2
	SumEndEvent = f3
}
func main() {

	f1 := func(a int) int {
		println("====SumBeforeEvent====")
		return a + 1
	}
	f2 := func(a int, b int) int {
		println("====SumEvent====")
		return a + b
	}
	f3 := func(c string) {
		println("====SumEndEvent====")
		println(c)
	}
	RegEvent(f1, f2, f3)
	f := StartSum(3, 7, "End")
	println(f) //11->3+1+7
}
