package gee

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// 重定义类型

type H map[string]interface{}

// Context 原有对象的封装
type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
	// middleware
	handlers []HandlerFunc
	index    int
	//engine pointer
	engine *Engine
}

func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(500, err.Error())
	}
}
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// newContext return an instance of Context
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}
func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

// PostForm get the form data
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query get the url query
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status setting the code
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 设置响应头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String string类型的返回
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 返回json类型参数
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) Fail(code int, err string) {
	c.StatusCode = code
	log.Fatal(err)
}
