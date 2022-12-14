package main

// "syscall/js"
// "syscall/js"
// "syscall/js"

type State struct {
	state    string
	setState func(string)
}

var x int
var strings []string
var hooks []interface{}
var idx int

func main() {
	// done := make(chan struct{}, 0)
	// _ = done
	strings = make([]string, 0)

	// js.Global().Set("wasmHash", js.FuncOf(save))
	// js.Global().Set("wasmHash", js.FuncOf(b))
	// a := js.Global().Get("inputField").String()
	// select {}
	// x,y := new(State) = useState("HI")
	var x string
	var y func(string)
	x, y = useState("HI", &x)
	y("LLL")
}

// func save(this js.Value, args []js.Value) interface{} {
// 	strings = append(strings, args[0].String())
// 	x++
// 	if len(strings) > 1 {
// 		return strings[0] + strings[1]
// 	}
// 	return strings[len(strings)-1]
// }

func useState(initialValue string, ptr *string) (string, func(string)) {
	var _idx = idx

	hooks = append(hooks, initialValue)
	hooks[_idx] = initialValue

	setState := func(newValue string) {
		*ptr = newValue
		hooks[_idx] = newValue

	}

	idx++

	return initialValue, setState
}
