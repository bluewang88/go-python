package Go_base

import (
	"fmt"
	"reflect"
	"strconv"
)

func HelloStrconv() {
	// strconv包提供了将基本数据类型转换为字符串以及将字符串转换为基本数据类型的函数。
	// strconv.Itoa()将整数转换为字符串
	// strconv.Atoi()将字符串转换为整数
	// strconv.FormatBool()将布尔值转换为字符串
	// strconv.ParseBool()将字符串转换为布尔值
	// strconv.ParseInt()将字符串转换为整数
	// strconv.ParseUint()将字符串转换为无符号整数
	// strconv.ParseFloat()将字符串转换为浮点数
	fmt.Println("---------------------------------strconv包类型转换开始---------------------------------")

	fmt.Println("itoa函数将整数转换为字符串")
	var num int = 10
	fmt.Println("类型转换var num int = 10输出num的值为", num, "num的类型为", reflect.TypeOf(num))
	strNum := strconv.Itoa(num) //将整数转换为字符串
	fmt.Println("类型转换var strNum = strconv.Itoa(num)输出strNum的值为", strNum, "strNum的类型为", reflect.TypeOf(strNum))

	fmt.Println("atoi函数将字符串转换为整数")
	parsedNum, err := strconv.Atoi(strNum) //将字符串转换为整数
	if err == nil {                        //如果转换成功，err为nil
		fmt.Println("类型转换var parsedNum, err = strconv.Atoi(strNum)输出parsedNum的值为", parsedNum, "parsedNum的类型为", reflect.TypeOf(parsedNum))
	} else {
		fmt.Println("Error converting string to int:", err)
	}

	fmt.Println("formatbool函数将布尔值转换为字符串")
	boolVal := true
	fmt.Println("类型转换var boolVal = true输出boolVal的值为", boolVal, "boolVal的类型为", reflect.TypeOf(boolVal))
	strBool := strconv.FormatBool(boolVal)
	fmt.Println("类型转换var strBool = strconv.FormatBool(boolVal)输出strBool的值为", strBool, "strBool的类型为", reflect.TypeOf(strBool))

	parsedBool, err := strconv.ParseBool(strBool)
	if err == nil {
		fmt.Println("类型转换var parsedBool, err = strconv.ParseBool(strBool)输出parsedBool的值为", parsedBool, "parsedBool的类型为", reflect.TypeOf(parsedBool))
	}

	fmt.Println("parseint函数将字符串转换为整数")
	strInt := "10"
	parsedInt, err := strconv.ParseInt(strInt, 10, 64)
	if err == nil {
		fmt.Println("类型转换var parsedInt, err = strconv.ParseInt(strInt, 10, 64)输出parsedInt的值为", parsedInt, "parsedInt的类型为", reflect.TypeOf(parsedInt))
	} else {
		fmt.Println("Error converting string to int64:", err)
	}

	// parseint 和 atio的区别
	// strconv.Atoi()函数是strconv.ParseInt()函数的一个简单封装，它的base参数默认为10，bitSize参数默认为0。
	// strconv.Atoi()函数返回的是int类型的整数，而strconv.ParseInt()函数返回的是int64类型的整数。
	// strconv.Atoi()函数只能将字符串转换为整数，而strconv.ParseInt()函数可以将字符串转换为int64类型的整数。

	fmt.Println("formatfloat函数将浮点数转换为字符串")
	floatVal := 3.14159890188
	fmt.Println("类型转换var floatVal = 3.14159输出floatVal的值为", floatVal, "floatVal的类型为", reflect.TypeOf(floatVal))
	strFloat := strconv.FormatFloat(floatVal, 'f', 6, 64) //将浮点数转换为字符串, 'f'表示格式化为小数点格式，6表示小数点后保留6位，64表示float64类型
	fmt.Println("类型转换var strFloat = strconv.FormatFloat(floatVal, 'f', 6, 64)输出strFloat的值为", strFloat, "strFloat的类型为", reflect.TypeOf(strFloat))

	fmt.Println("---------------------------------strconv包类型转换结束---------------------------------")

}
