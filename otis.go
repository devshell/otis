package main

import (
	"fmt"
	"net/http"
	"net/url"
)

/**************************************************
        Set up the Handler function
**************************************************/
type Request interface {
	In() (u *url.URL, h *http.Header, req *http.Request, err error)
}

type Response interface {
	Out() (status int, head *http.Header, resp *http.Response, err error)
}

type Handler func(r Request) Response

/**************************************************
        Set up the Otis object
**************************************************/
type Otis struct {
	Current        int
	DefaltHandlers map[string]Handler
	UserHandlers   map[string]Handler

	//  Check err and determine which handler to use using convention "handlerName_error"
	//  with the handler called "error" handling all errors not caught by a specific
	//  "handlerName_error" Handler
	ErrHandlers map[string]Handler
}

func New() *Otis {
	return &Otis{0, make(map[string]Handler), make(map[string]Handler)}
}

/**************************************************
        Set up the Otis methods
**************************************************/
func (o *Otis) Before(handlerName string) {
	// First establish where t

	//o.Handlers = append(o.Handlers, h)

	//debug
	fmt.Println("[", len(o.Handlers), "]")
}

/**************************************************
        Test if we made it this far
**************************************************/
func main() {
	fmt.Println("All systems are Go!")
}

/*
func (o *Otis) Walk() {

	//debug
	fmt.Println("[", len(h.Strtr), "]")
	//fmt.Println(h.strtr[0]("World").store)

	for _, fn := range h.Strtr {
		fmt.Println(fn(s).Store)
	}

}

func (h *Handler) Pop(n int64) (s strt) {

	//debug
	fmt.Println("[", len(h.Strtr), "]")
	//fmt.Println(h.strtr[0]("World").store)

	return h.Strtr[n]

}

func Crazy(s string) (r str) {
	r.Store = fmt.Sprintf("Crazy %s", s)
	return r
}

func main() {

	h := New()

	//debug
	fmt.Println("[", len(h.Strtr), "]")

	h.Add(func(s string) (r str) {
		r.Store = fmt.Sprintf("Goodbye %s", s)
		return r
	})

	h.Add(func(s string) (r str) {
		r.Store = fmt.Sprintf("Hello %s", s)
		return r
	})

	h.Add(Crazy)

	fmt.Println(h.Pop(2)("World").Store)

	//debug
	fmt.Println("[", len(h.Strtr), "]")

	h.Walk("World")

		for _, fn := range t {
			fmt.Println(fn("World").store)
		}

}
*/
