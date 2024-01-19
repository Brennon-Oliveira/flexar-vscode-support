package teste

import (
	"fmt"
	"runtime"
)


func main() {
	fmt.Println("Hello World")
	fmt.Println(runtime.NumCPU())
}