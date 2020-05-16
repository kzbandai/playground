package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// Goならわかるシステムプログラミング（PDF版）.pdf page 125

func isGZipAcceptable(r *http.Request) bool {
	return strings.Index(strings.Join(r.Header["Accept-Encoding"], ","), "gzip") != -1
}

func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()

	for {
		conn.SetDeadline(time.Now().Add(5 * time.Second))
		req, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}
		dump, err := httputil.DumpRequest(req, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		res := http.Response{
			StatusCode: 200,
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header:     make(http.Header),
		}

		if isGZipAcceptable(req) {
			content := "hello gzip\n"
			var buf bytes.Buffer
			writer := gzip.NewWriter(&buf)
			io.WriteString(writer, content)
			writer.Close()

			res.Body = ioutil.NopCloser(&buf)
			res.ContentLength = int64(buf.Len())
			res.Header.Set("Content-Encoding", "gzip")
		} else {
			content := "hello"
			res.Body = ioutil.NopCloser(strings.NewReader(content))
			res.ContentLength = int64(len(content))
		}

		res.Write(conn)
	}
}

func main() {
	listner, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("server is running at localhost:8888")

	for {
		conn, err := listner.Accept()
		if err != nil {
			panic(err)
		}

		go processSession(conn)
	}
}
