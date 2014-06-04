package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

/**************************************************
        Set up the Handler function
**************************************************/
type Request interface {
	OReq(u *url.URL, h *http.Header, req *http.Request, err error) (u *url.URL, h *http.Header, req *http.Request, err error)
}

type Response interface {
	OResp(u *url.URL, h *http.Header, resp *http.Response, err error) (status int, head *http.Header, resp *http.Response, err error)
}

type Handler func(req *Request) resp* Response

/**************************************************
        Set up the Otis object
**************************************************/
type Otis struct {
	ucursor uint // Insert User-defined handlers after/before this position
	ecursor uint // Insert Error handlers after/before this position

	// This is the section for user generated handler stacking
	Handlers      map[string]Handler // Use this to stack handlers
	HandlersIndex map[uint]string    // Use this to look up a UserHandler using Cursor

	//  Check err and determine which handler to use using convention "handlerName_error"
	//  with the handler called "error" handling all errors not caught by a specific
	//  "handlerName_error" Handler
	ErrHandlers    map[string]Handler // Used to stack error handlers
	eHandlersIndex map[uint]string    // Use this to look up an ErrHandler using Cursor

}

/**************************************************
        SPECIFICATION
***************************************************


// Create a handler chain of common handlers that others can inherit
CommonHandlers := Otis.New()

// Assign a base set of handlers that will be inherited by all other handlers
_ := CommonHandlers.Append("name1", FunctioncallB(args))


// Create custom handler
CustomHandlers := Otis.New()

// INHERITANCE
// Insert another Otis chain starting at index 0
_ := CustomHandlers.Inject(CommonHandlers.Handlers)


// INHERIT AFTER ITEM
// Insert Otis chain starting at index returned by (After)
_ := CustomHandlers.After("name4").Inject(CommonHandlers.Handlers)


_:= CustomHandlers.Append("name10", Functioncall10(args))   // Add after last item == Firt()
_:= CustomHandlers.After("name4").Append("name7", Functioncall3(args))
_:= CustomHandlers.Before("name7").Append("name6", Functioncall4(args))
_:= CustomHandlers.Delete("name7")
_ := CustomHandlers.Overwrite("name10").Insert("name20", Functioncall20(args))


// Handle errors

Check for errors on each handler return, and if there is an error check the ErrHandlers map for a specific
handlerName_error entry, and if there isn't, then check for an "error" entry, and if there isn't one,
go to the defaults map, and check the "error" value.



// Output current list in a formatted string obj
CustomHandlers.Inspect()



mux = NewMux()
mux.Get("/hello", CustomHandlers.Handle())

http.ListenAndServe(":8080", mux)

**************************************************/

func New() *Otis {
	return &Otis{
		0, // Current cursor position for user-defined handlers
		0, // Current cursor position for error handlers
		make(map[string]Handler), // User-defined Handlers
		make(map[uint]string),    // Index of user-defined handlers
		make(map[string]Handler), // Error Handlers
		make(map[uint]string)}    // Index of error handlers
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
	er := errors.New("All systems are Go!")
	fmt.Println(er)
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
