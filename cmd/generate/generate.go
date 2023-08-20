package main

import (
	"flag"
	"fmt"

	"gorm.io/gen"
	"gorm.io/gen/examples/dal"
)

var dataMap = map[string]func(detailType string) (dataType string){
	"int":  func(detailType string) (dataType string) { return "int64" },
	"json": func(string) string { return "json.RawMessage" },
}

func main() {
	dsn := flag.String("dsn", "root:123456@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True", "consult[https://gorm.io/docs/connecting_to_the_database.html]")
	flag.Parse()

	fmt.Println("dd", *dsn)

	g := gen.NewGenerator(gen.Config{
		OutPath:      "../../biz/dal/query",
		ModelPkgPath: "../../biz/dal/model",
		Mode:         gen.WithDefaultQuery,

		WithUnitTest: true,

		FieldNullable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	g.UseDB(dal.ConnectDB(*dsn))

	g.WithDataTypeMap(dataMap)
	g.WithJSONTagNameStrategy(func(c string) string { return c })

	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
