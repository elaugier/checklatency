package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
)

//Options ...
type Options struct {
	//IpAddress ...
	IpAddress string `short:"a" long:"ipaddress" default:"" description:"Ip Address"`
	//Port ...
	Port *int `short:"p" default:"3000" description:"Tcp port"`
}

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
	var opts Options
	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		os.Exit(1)
	}

	var buffer bytes.Buffer
	buffer.WriteString(opts.IpAddress)
	buffer.WriteString(":")
	buffer.WriteString(String(int64(*opts.Port)))
	server := &App{}
	fmt.Println("listening on ", buffer.String())
	http.ListenAndServe(buffer.String(), server)
}
