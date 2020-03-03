package main

import "fmt"

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("%d, %s", self.id, self.name)
}
func main() {
	var o interface{} = &User{1, "Tom"}
	if i, ok := o.(fmt.Stringer); ok { //接口内的具体类型判断
		fmt.Println(i)
	}
	u := o.(*User)
	fmt.Println(u)

}
