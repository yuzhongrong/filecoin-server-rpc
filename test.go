package main

import (
	"fmt"
	"github.com/myxtype/filecoin-client/util"
)

func main1() {
	testStr:="t1otzxmzqqrlnpfzczhb4tx6c4glwzlozqx5lalja"
	enstr:=util.Encrypt(testStr)
	fmt.Println(enstr)
	destr:=util.Decrypt(enstr)
	fmt.Println(destr)
}


