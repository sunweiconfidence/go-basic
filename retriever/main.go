package main

import (
	"fmt"
	"github.com/sunweiconfidence/learngo/retriever/mock"
	"github.com/sunweiconfidence/learngo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is fake api test"}
    r = real.Retriever{}
	fmt.Println(download(r))
}
