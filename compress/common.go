package compress

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"github.com/itchio/lzma"
	"io"
)

func ZlibCompress(b []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(b)
	w.Close()
	return in.Bytes()
}

func ZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

func GzipCompress(b []byte) []byte {
	var in bytes.Buffer
	w := gzip.NewWriter(&in)
	w.Write(b)
	w.Close()
	return in.Bytes()
}

func GzipUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := gzip.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

func LzmaCompress(b []byte) []byte {
	var in bytes.Buffer
	w := lzma.NewWriter(&in)
	w.Write(b)
	w.Close()
	return in.Bytes()
}

func LzmaUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r := lzma.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}