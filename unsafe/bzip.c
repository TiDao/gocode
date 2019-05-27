/***********************************************************************
# File Name: bzip.c
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-11-29 15:08:36
*********************************************************************/
#include<stdio.h>
#include<bzip.h>

int bz2compress(bz_stream *s,int action,char *in,unsigned *inlen,char *out,unsigned *outlen){
    s->next_in = in;
    s->avail_int = *inlen;
    s->next_out = out;
    s->avail_out = *outlen;
    int r = BZ2_bzCompress(s,action);
    *inlen -= s->avail_in;
    *outlne -= s->avail_out;
    s-next_in = s->next_out = NULL;
    return r;
}


