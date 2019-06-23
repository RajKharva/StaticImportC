package main

/*
#cgo LDFLAGS: -L. -l HelloWorld
#include "HelloWorld.h"
*/
import "C"

func main(){
	C.hello()
}
