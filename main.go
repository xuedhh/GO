package main

import (
	_ "user/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"user/internal/cmd"
	_ "user/internal/logic/enterprise"
	_ "user/internal/logic/user"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())

	//type User struct {
	//	Uid  int    `c:"uid"`
	//	Name string `c:"name"`
	//}
	//
	//g.Dump(gconv.Map(User{
	//	Uid:  1,
	//	Name: "john",
	//}))
	//g.Dump(gconv.Map(&User{
	//	Uid:  1,
	//	Name: "john",
	//}))
	//
	//g.Dump(gconv.Map(map[int]int{
	//	100: 1000,
	//}))
}
