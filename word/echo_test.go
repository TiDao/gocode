/***********************************************************************
# File Name: echo_test.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-12-25 10:21:33
*********************************************************************/
package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		newline bool
		sep     string
		args    []string
		want    string
	}{
		{true, "", []string{}, "\n"},
		{false, "", []string{}, ""},
		{true, "\t", []string{"one", "two", "three"}, "one\ttwo\tthree\n"},
		{true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
		{false, ":", []string{"1", "2", "3"}, "1:2:3"},
	}
    for _,test := range tests{
        descr := fmt.Sprintf("echo(%v,%q,%q)",
            test.newline,test.sep,test.args)

        fmt.Println(descr)

        out = new(bytes.Buffer)
        if err := echo(test.newline,test.sep,test.args); err != nil{
            t.Errorf("%s failed: %v",descr,err)
            continue
        }
        
        got := out.(*bytes.Buffer).String()
        if got != test.want{
            t.Errorf("%s = %q,want %q",descr,got,test.want)
        }
    }
}


