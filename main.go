package main

import (
	"encoding/json"
	"fmt"

	"github.com/mixi-gaming/handle-data/convert"
	"github.com/mixi-gaming/handle-data/mapstring"
)

func main() {
	data := map[string]interface{}{
		"string": nil,
		"list":   []interface{}{"1234", 3124},
		"mapstring": map[string]interface{}{
			"test": "134",
		},
	}

	bodyBytes := []byte(`{
		"string": null,
		"list":   [1234, 3124]
	}`)

	data2 := make(map[string]interface{})
	var data3 interface{}
	_ = json.Unmarshal(bodyBytes, &data2)
	_ = json.Unmarshal(bodyBytes, &data3)

	var a interface{} = 1234
	var b = "1234"
	var c = "1234a"
	var d interface{} = nil
	var e = ""
	var f = 10.11111111

	fmt.Println(mapstring.GetString(data, "string.int.cd", false))
	fmt.Println(convert.ToMapString(data["mapstring"]))
	fmt.Println(convert.ToMapString(data3))

	fmt.Println(convert.ToSliceString(data2["list"]))
	fmt.Println(convert.ToSliceMapString(data2["list"]))

	fmt.Println(convert.ToInt(a))
	fmt.Println(convert.ToInt(b))
	fmt.Println(convert.ToInt(c))
	fmt.Println(convert.ToInt(d))
	fmt.Println(convert.ToInt(e))
	fmt.Println(convert.ToInt(f))
}
