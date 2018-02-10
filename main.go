package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"fmt"
	"strings"
	//"github.com/robertkrimen/otto"
)

type EndPoint struct {
	Name string
	FileName string
}

type Response struct {
	StatusCode int
}

func main() {
	r := gin.Default()

	f, e := ioutil.ReadDir("./js")
	if e != nil {
		fmt.Errorf("%v", e)
	} else {
		for _, _f := range f {
			if _f.IsDir() {
				path := "./js/" + _f.Name()
				files := getFiles(path)

				for _, file := range files {
					r.Handle(strings.ToUpper(_f.Name()), file.Name, getCallback(path + "/" + file.FileName))
				}
			}
		}
	}

	r.Run(":8080")
}

func getFiles(dirname string) []EndPoint {
	var result []EndPoint
	files, _ := ioutil.ReadDir(dirname)
	for _, file := range files {
		result = append(result, EndPoint{
			Name: strings.Replace(file.Name(), ".js", "", -1),
			FileName: file.Name(),
		})
	}
	return result
}

func getCallback(name string) func(*gin.Context) {
	res := &Response{
		StatusCode: 404,
	}
	return func(ctx *gin.Context) {
		ctx.String(res.StatusCode, name)
	}
}