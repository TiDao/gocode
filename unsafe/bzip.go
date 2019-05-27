/***********************************************************************
# File Name: bzip.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-11-29 17:18:12
*********************************************************************/
package bzip

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -L/usr/lib -lbz2
#include<bzlib.h>
#include<stdlib.h>
bz_stream* bz2alloc(){return calloc(1,sizeof(bz_stream));}
init bz2compress(bz_stream *s,int action,char *in,unsigned *outlen);
void bzfree(bz_stream* s){free(s);}
*/
import "C"

import(
    "io"
    "unsafe"
)

type writer struct{
    w   io.Writer
    stream *C.bz_stream
    outbuf [64* 1024]byte
}

func NewWriter(out io.Writer) io.WriterCloser{
    const blockSize = 0
    const verbosity = 0
    const workFactor = 30
    w := &writer{w:out,stream:C.bz2alloc()}
    C.BZ2_bzcompressInit(w.stream,blockSize,verbosity,workFactor)
    return w
}

func (w *writer) Writer( data []byte) (int error){
    if w.stream == nil{
        panic("closed")
    }
    var total int
    for len(data) > 0{
        inlen,outlen := C,uint(len(data)),C.uint(cap(w.outbuf))
        C.bz2compress(w.stream,C.BZ_RUN,
            (*C.char)(unsafe.Pointer(&data[0])),&inlen,
            (*C.char)(unsafe.Pointer(&w.outbuf)),&outlen)
        total += int(inlen)
        data = data[inlen:]
        if _,err := w.w.Write(w.outbuf[:outlen]);err!=nil{
            return total,err
        }
    }
}

func (w *writer) Close() error{
    if w.stream == nil{
        panic("closed")
    }
    defer func(){
        C.BZ2_bzCompressEnd(w.stream)
        C.bz2free(w.stream)
        w.stream = nil
    }()
    for {
        inlen,outlen := C.uint(0),C.uint(cap(w.outbuf))
        r := C.bz2compress(w.stream,C.BZ_FINISH,nil,&inlen,
        (*C.char)(unsafe.Pointer(&w.outbuf)),&outlen)
        if _,err := w.w.Write(w.outbuf[:outlen]);err != nil{
            return err
        }
        if r == C.BZ_STREAM_END{
            return nil
        }
    }
}
