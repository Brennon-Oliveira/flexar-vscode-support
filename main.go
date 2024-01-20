package teste

import (
	"fmt"
	"runtime"
)

type Pessoa struct {
	teste int
}


func main() {
	fmt.Println("Hello World")
	fmt.Println(runtime.NumCPU())
}
