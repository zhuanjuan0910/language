package main

import "fmt"

type user struct {
	id   int
	name string
}
type managr struct {
	user
	title string
}

func (self *user) ToString() string {
	return fmt.Sprintf("user:%p,%v", self, self)
}
func (self *managr) ToString() string {
	return fmt.Sprintf("managr:%p,%v", self, self)
}
func main() {
	u := managr{user{1, "tom"}, "author"}
	fmt.Println(u.ToString())
	fmt.Println(u.user.ToString())
}

//  输出结果为 managr:0xc000076150,&{{1 tom} author}
//user:0xc000076150,&{1 tom}
