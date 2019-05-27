/***********************************************************************
# File Name: interface.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-05-14 14:36:25
*********************************************************************/
package main
type File struct{
    ...
}

func (f *File) Read(buf []byte) (n int, err error)
func (f *File) Write(buf []byte) (n int,err error)
func (f *File) Seek(off int64,whence int) (n int,err error)
func (f *File) Close() error

type Integer int

func (a Interger) Less(b Integer) bool{
    return a < b
}

func (a *Interger) Add(b Integer) {
    *a += b
}

type LessAdder interface{
    Less(b Integer) bool
    Add(b Integer)
}

package one
type ReadWriter interface{
    Read(buf []byte) (n int,err error)
    Write(buf []byte) (n int,err error)
}
package two
type IStream interface{
    Write(buf []byte) (n int,err error)
    Read(buf []byte) (n int,err error)
}
var file1 two.IStream = new(File)
var file2 one.ReadWriter = file1
var file3 two.IStream = file2

type Writer interface{
    Write(buf []byte) (n int,err error)
}


