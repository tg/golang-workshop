package main

import (
	"fmt"
	"strings"
)

func main() {
	strings.ToUpper("i love my bot")    // I LOVE MY BOT
	strings.Split("i love my bot", " ") // ["i" "love" "my" "bot"]
	strings.Count("i love my bot", "o") // 2
	fmt.Println(strings.Join([]string{"go", "lang"}, "_"))
	// strings.ToU
	//
	// s := strings.Join(
	// 	strings.Split(
	//
	// 		" ",
	// 	),
	// 	"_",
	// )
	//

}
