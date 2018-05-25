package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"time"
)

// App ...
type App struct {
}

// String ...
func String(n int64) string {
	buf := [11]byte{}
	pos := len(buf)
	i := n
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	start := int64(time.Now().UnixNano())
	time.Sleep(100 * time.Millisecond)
	res.Write([]byte("{\"ServerSideNano\":\"" + String(int64(time.Now().UnixNano())-start) + "\"}"))
}

func main() {
	numbPtr := flag.Int("p", 3000, "tcp port for listening")
	fmt.Println("listening on port:", *numbPtr)
	var buffer bytes.Buffer
	buffer.WriteString(":")
	buffer.WriteString(String(int64(*numbPtr)))
	flag.Parse()
	server := &App{}
	fmt.Println(buffer.String())
	http.ListenAndServe(buffer.String(), server)
}
