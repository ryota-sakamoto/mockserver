package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"fmt"
	"strings"
	"github.com/robertkrimen/otto"
	"flag"
	"strconv"
)

var (
	port = flag.Int("p", 8080, "port")
)

type EndPoint struct {
	Name string
	FileName string
}

type Response struct {
	StatusCode int
	ContentType string
	Body string
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

	r.Run(":" + strconv.Itoa(*port))
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
	vm := otto.New()
	return func(c *gin.Context) {
		buffer, _ := ioutil.ReadFile(name)
		vm.Set("response", res)
		vm.Run(string(buffer))
		switch res.ContentType {
		case "json":
			c.JSON(res.StatusCode, res.Body)
		default:
			c.String(res.StatusCode, res.Body)
		}
	}
}