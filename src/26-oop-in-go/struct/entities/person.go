package entities

// Person 实体
type Person struct {
	Name string
	Age  int
}

// Older 比较哪个人更老. 返回较老的人，年龄差
func Older(p1, p2 Person) (Person, int) {
	if p1.Age > p2.Age {
		return p1, p1.Age - p2.Age
	}
	return p2, p2.Age - p1.Age
}
