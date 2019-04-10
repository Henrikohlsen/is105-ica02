package main

import (
	"fmt"
<<<<<<< HEAD
	"github.com/Henrikohlsen/Is-105/is105-ica02/slice"
=======
	"./slice"
>>>>>>> ecbb1c18f3aa735881bc43637b9eb9394d958062
)

func main() {

	var byteSlice1 = slice.AllocateVar(100)
	fmt.Println(" &byteslice1[0]")
	fmt.Println(&byteSlice1[0])
	aslice := slice.Reslice(byteSlice1, 50, 100)
	fmt.Println(" &aslice[0]")
	fmt.Println(&aslice[0])
	fmt.Println("&byteslice1[50]")
	fmt.Println(&byteSlice1[50])
}
