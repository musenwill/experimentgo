package sample

/*
#include <stdio.h>
static void myprint(char* s) {
  printf("%s\n", s);
}
*/
import "C"

func helloCgo() {
	cs := C.CString("hello cgo")
	C.myprint(cs)
}
