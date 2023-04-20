package main

import (
	"github.com/ag9920/db-demo/gendemo/dal/model"
	"gorm.io/gen"
)

/*
 * 生成代码逻辑
 * https://juejin.cn/post/7196511671626940476
 */
func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithDefaultQuery,
	})

	g.ApplyBasic(model.Passport{}, model.User{})

	g.ApplyInterface(func(model.Method) {}, model.User{})
	g.ApplyInterface(func(model.UserMethod) {}, model.User{})

	g.Execute()
}
