package objectgotest

// Go接口设计的关键规范
// 接口命名：

// 使用"er"后缀（如Reader, Writer, Stringer）
// 表达能力或行为，而非身份
// 单一方法接口通常命名为"方法名+er"
// 接口大小：

// 倾向于小接口，最好是单一方法接口
// 大接口通过组合小接口形成
// 遵循"接口隔离原则"
// 方法命名：

// 使用简短、清晰的动词或动词短语
// 不需要在方法名中重复接收者类型
// 保持一致性（如Read/Write, Get/Set）
// 实现规范：

// 隐式实现，无需声明
// 接口应该由调用者定义，而非实现者
// 优先使用值接收者，除非需要修改接收者或有大型结构体
// 文档：

// 每个接口都应有清晰的文档说明
// 说明接口意图和预期行为
// 注明任何特殊约定（如错误处理）
import (
	"fmt"
	"io"
	"time"
)

// 1. 接口命名规范：
// - 通常使用"er"后缀（表示"做某事的人/物"）
// - 单一职责，每个接口应该只关注一个功能点
// - 名称应该表达"做什么"而不是"是什么"

// Reader 是读取数据的接口
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Writer 是写入数据的接口
type Writer interface {
	Write(p []byte) (n int, err error)
}

// 2. 接口组合：通过嵌入其他接口来创建更大的接口

// ReadWriter 组合了Reader和Writer功能
type ReadWriter interface {
	Reader
	Writer
}

// 3. 方法命名：使用动词或动词短语，表明行为

// Closer 是一个可关闭资源的接口
type Closer interface {
	Close() error
}

// 4. 单一方法接口：通常命名为"方法名+er"

// Stringer 接口用于获取对象的字符串表示
type Stringer interface {
	String() string
}

// 5. 实现接口的结构体：不需要显式声明实现了哪个接口

// FileHandler 实现了ReadWriter和Closer接口
type FileHandler struct {
	data []byte
	pos  int
	name string
}

func (f *FileHandler) Read(p []byte) (n int, err error) {
	if f.pos >= len(f.data) {
		return 0, io.EOF
	}

	n = copy(p, f.data[f.pos:])
	f.pos += n
	return n, nil
}

func (f *FileHandler) Write(p []byte) (n int, err error) {
	f.data = append(f.data, p...)
	return len(p), nil
}

func (f *FileHandler) Close() error {
	f.data = nil
	f.pos = 0
	return nil
}

// 6. 接口参数：优先使用小接口作为函数参数

// CopyData 函数接受最小接口需求
func CopyData(r Reader, w Writer) error {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		_, err = w.Write(buf[:n])
		if err != nil {
			return err
		}
	}
	return nil
}

// 7. 空接口：用于表示"任何类型"，现在推荐使用any替代interface{}

// Process 处理任何类型的数据
func Process(data any) {
	// 使用类型断言处理不同类型
	switch v := data.(type) {
	case string:
		fmt.Println("字符串:", v)
	case int:
		fmt.Println("整数:", v)
	case time.Time:
		fmt.Println("时间:", v.Format("2006-01-02"))
	default:
		fmt.Println("未知类型:", v)
	}
}

// 8. 实现示例函数
func ExampleUsage() {
	// 创建实现接口的对象
	file := &FileHandler{name: "test.txt"}

	// 写入数据
	file.Write([]byte("Hello, Interface!"))

	// 重置读取位置
	file.pos = 0

	// 读取数据
	buf := make([]byte, 100)
	n, _ := file.Read(buf)
	fmt.Println(string(buf[:n]))

	// 通过接口引用使用
	var rw ReadWriter = file
	var c Closer = file

	// 使用接口方法
	rw.Write([]byte(" More data."))
	c.Close()
}

// Go语言接口命名规范和编码示例
// 在Go语言中，接口的命名和实现有一些约定俗成的规范。以下是一个示例代码，展示了Go语言中接口的命名规范和编码规范：

// Go接口设计的关键规范
// 接口命名：

// 使用"er"后缀（如Reader, Writer, Stringer）
// 表达能力或行为，而非身份
// 单一方法接口通常命名为"方法名+er"
// 接口大小：

// 倾向于小接口，最好是单一方法接口
// 大接口通过组合小接口形成
// 遵循"接口隔离原则"
// 方法命名：

// 使用简短、清晰的动词或动词短语
// 不需要在方法名中重复接收者类型
// 保持一致性（如Read/Write, Get/Set）
// 实现规范：

// 隐式实现，无需声明
// 接口应该由调用者定义，而非实现者
// 优先使用值接收者，除非需要修改接收者或有大型结构体
// 文档：

// 每个接口都应有清晰的文档说明
// 说明接口意图和预期行为
// 注明任何特殊约定（如错误处理）
