# panic/recover

Go语言中目前（Go1.12）是没有异常机制，但是使用panic/recover模式来处理错误。 panic可以在任何地方引发，但recover只有在defer调用的函数中有效.

# panic
## panic的使用
panic是Go语言中的一个内置函数，用于引发运行时错误。它会立即停止当前函数的执行，并开始向上层函数传递错误信息，直到找到一个处理该错误的defer函数为止。如果没有找到处理该错误的defer函数，程序将终止并输出错误信息。
panic的语法如下：
```go
panic(v interface{})
```
- `v`：要传递的错误信息，可以是任何类型的值。
- `panic`函数会引发一个运行时错误，并停止当前函数的执行。

## panic的使用示例
```go
func main() {}
    // 引发一个运行时错误
    panic("something went wrong")
}
```
- 在上面的示例中，调用`panic`函数引发了一个运行时错误，并传递了错误信息"something went wrong"。
- 当程序运行到`panic`函数时，当前函数的执行将立即停止，并开始向上层函数传递错误信息。
- 如果没有找到处理该错误的`defer`函数，程序将终止并输出错误信息。
- `panic`函数可以用于处理一些不可恢复的错误，例如数组越界、空指针引用等。


# recover
程序运行期间funcB中引发了panic导致程序崩溃，异常退出了。这个时候我们就可以通过recover将程序恢复回来，继续往后执行。

## recover的使用
recover是Go语言中的一个内置函数，用于恢复从panic中引发的错误。它只能在defer函数中使用，并且只能恢复当前goroutine中的panic。
recover的语法如下：
```go
recover() interface{}
```
- `recover`函数会返回引发panic时传递的错误信息，如果没有panic发生，则返回nil。
- `recover`函数只能在defer函数中使用，否则会返回nil。
- `recover`函数可以用于处理一些可恢复的错误，例如网络请求失败、文件读取失败等。
- `recover`函数可以用于在panic发生时恢复程序的执行，并继续往后执行。
## recover的使用示例
```go
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	funcA()
	funcB()
	funcC()
}
```