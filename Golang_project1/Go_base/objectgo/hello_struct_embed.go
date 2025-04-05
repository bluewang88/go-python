package objectgotest

// go语言不支持继承

// 但是可以通过组合compose的方式实现类似的功能

type PersonSchool struct {
	Name       string
	Age        int
	Gender     string
	SchoolName string
}

type StudentSchool struct {
	PersonSchool //匿名嵌套
	Grade        int
	Score        float64
	Course       string
}

func (Student *StudentSchool) GetName() string {
	return Student.Name
}
func (Student *StudentSchool) GetAge() int {
	return Student.Age
}
