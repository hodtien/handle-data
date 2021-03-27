package main

import (
	"fmt"

	"github.com/mixi-gaming/handle-data/convert"
	"github.com/mixi-gaming/handle-data/mapstring"
)

func main() {
	data := map[string]interface{}{
		"string": nil,
		"list":   []interface{}{"1234", 3124},
	}

	var a interface{} = 1234
	var b = "1234"
	var c = "1234a"
	var d interface{} = nil
	var e = ""
	var f = 10.11111111


	fmt.Println(mapstring.GetString(data, "string.int.cd", false))
	fmt.Println(convert.ToInt(a))
	fmt.Println(convert.ToInt(b))
	fmt.Println(convert.ToInt(c))
	fmt.Println(convert.ToInt(d))
	fmt.Println(convert.ToInt(e))
	fmt.Println(convert.ToInt(f))
}
