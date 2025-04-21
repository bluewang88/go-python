# go map

Go语言中提供的映射关系容器为map，其内部使用散列表（hash）实现。

map是一种无序的基于`key-value`的数据结构，Go语言中的map是`引用类型`，必须初始化才能使用。

## map的定义

```go
var mapName map[keyType]valueType
```
- `keyType`：键的类型，可以是任何可比较的类型（如字符串、整数、布尔值等）,不能是空类型（nil）、函数类型、指针类型、切片类型、通道类型、接口类型、map类型。
- `valueType`：值的类型，可以是任何类型，包括自定义类型、结构体等。
- `mapName`：映射的变量名称。
- `var`：声明变量的关键字。
- `map`：关键字，表示声明一个映射类型的变量。
- `keyType`和`valueType`：分别表示键和值的类型。

> map类型的变量是引用类型，，默认初始值为 nil，需要使用make()函数初始化分配内存。

##  map的基本使用

map中的数据都是成对出现的，map的基本使用示例代码如下：

```go
package main
import (
    "fmt"
)
func main() {}
    // 声明一个map变量
    var m map[string]int
    fmt.Println(m) // 输出：map[]
    // 初始化map
    m = make(map[string]int)
    fmt.Println(m) // 输出：map[]
    // 添加元素
    m["a"] = 1
    m["b"] = 2
    fmt.Println(m) // 输出：map[a:1 b:2]
}
```
- `make(map[string]int)`：使用`make`函数初始化一个空的map，键为字符串类型，值为整数类型。
- `m["a"] = 1`：向map中添加键值对，键为`"a"`，值为`1`。
- `m["b"] = 2`：向map中添加键值对，键为`"b"`，值为`2`。
- `fmt.Println(m)`：打印map的内容，输出`map[a:1 b:2]`，表示map中有两个键值对，键为`"a"`和`"b"`，对应的值分别为`1`和`2`。

## 判断某个键是否存在

在Go语言中，可以使用`value, ok := m[key]`的方式来判断某个键是否存在于map中。

- `value`：表示键对应的值，如果键存在，则为对应的值；如果键不存在，则为值类型的零值。
- `ok`：表示键是否存在于map中，如果存在，则为`true`；如果不存在，则为`false`。
- `m[key]`：表示获取map中键为`key`的值，如果键不存在，则返回值类型的零值。

## map的遍历
在Go语言中，可以使用`for range`语句来遍历map中的键值对。

```go
package main
import (
    "fmt"
)
func main() {
    // 声明并初始化一个map
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }
    // 遍历map
    for key, value := range m {
        fmt.Printf("key: %s, value: %d\n", key, value)
    }
}
```
- `for key, value := range m`：使用`for range`语句遍历map中的键值对，`key`表示键，`value`表示值。
- `fmt.Printf("key: %s, value: %d\n", key, value)`：打印每个键值对，`%s`表示字符串格式，`%d`表示整数格式。


> 注意： 遍历map时的元素顺序与添加键值对的顺序无关。

## 按照指定顺序遍历 map
在Go语言中，可以使用`for range`语句来按照指定顺序遍历map中的键值对。

```go
package main
import (
    "fmt"
    "sort"
)
func main() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
```
- `rand.Seed(time.Now().UnixNano())`：初始化随机数种子。
- `var scoreMap = make(map[string]int, 200)`：声明并初始化一个map，键为字符串类型，值为整数类型。
- `for i := 0; i < 100; i++`：循环100次，生成100个随机数。
- `key := fmt.Sprintf("stu%02d", i)`：生成`"stu"`开头的字符串，后面跟上两位数字。
- `value := rand.Intn(100)`：生成0~99的随机整数。
- `scoreMap[key] = value`：将生成的键值对添加到map中。
- `var keys = make([]string, 0, 200)`：声明一个切片，用于存储map中的所有键。
- `for key := range scoreMap`：遍历map中的所有键，将键添加到切片中。
- `sort.Strings(keys)`：对切片中的键进行排序。
- `for _, key := range keys`：按照排序后的键遍历map，打印每个键值对。
- `fmt.Println(key, scoreMap[key])`：打印每个键值对，`key`为键，`scoreMap[key]`为对应的值。
- `fmt.Println(key, scoreMap[key])`：打印每个键值对，`key`为键，`scoreMap[key]`为对应的值。


## map的删除
在Go语言中，可以使用`delete(m, key)`函数来删除map中的键值对。

```go
package main
import (
    "fmt"
)
func main() {
    // 声明并初始化一个map
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }
    // 删除键为"b"的键值对
    delete(m, "b")
    fmt.Println(m) // 输出：map[a:1 c:3]
}
```
- `delete(m, "b")`：删除map中键为`"b"`的键值对。
- `fmt.Println(m)`：打印map的内容，输出`map[a:1 c:3]`，表示map中只剩下键为`"a"`和`"c"`的键值对。


## map的长度
在Go语言中，可以使用`len(m)`函数来获取map的长度，即map中键值对的数量。

```go
package main
import (
    "fmt"
)
func main() {
    // 声明并初始化一个map
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }
    // 获取map的长度
    length := len(m)
    fmt.Println(length) // 输出：3
}
```
- `len(m)`：获取map的长度，即map中键值对的数量。
- `fmt.Println(length)`：打印map的长度，输出`3`，表示map中有三个键值对。
  
## map的拷贝
在Go语言中，map是引用类型，直接赋值会导致两个变量指向同一个底层数据结构，因此修改一个变量的值会影响另一个变量的值。

```go
package main

import (
    "fmt"
    "copy" // 添加拷贝包
    "reflect" // 添加反射包 
)
func main() {
    // 声明并初始化一个map
    m1 := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }
    // 拷贝map
    m2 := copy.Copy(m1)
    fmt.Println(m2) // 输出：map[a:1 b:2 c:3]
}
```
- `copy.Copy(m1)`：使用`copy`包中的`Copy`函数拷贝map，返回一个新的map。
- `fmt.Println(m2)`：打印拷贝后的map的内容，输出`map[a:1 b:2 c:3]`，表示拷

