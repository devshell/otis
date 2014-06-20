package main

import (
	"errors"
	"fmt"
	"net/http"
)

/**************************************************
        TODO LIST


-	Need to be able to add to Request & Response interfaces
	so that there is consistency as the data flows down the stack
	from top to bottom (same data from start to finish).

	Some middleware will need additional data about the http request
	than other middleware.

	These additional data should be injected into the Req/Resp interfaces
	at the very top of the stack, and they should continue to run all the
	way down the stack.

	->  	Is this really needed?  The inteface can be modified at the pt
		where the middleware creates it, and it will continue to flow
		down the stack from that point forward, with each middleware
		adding to the interface if/as necessary.

**************************************************/

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

/**************************************************
        Set up the Handler function
**************************************************/
type Middleware interface {
	http.Handler // Make this middleware an Http Handler interface

	/************************************************
	    All middleware should have a New() constructor function
	    that returns an instance of the middleware object
	    which will then be used by Otis, the New() constructor function
	    call is where the options are specified for the configuration
	    of the middleware for that particular chain


	    func New(...args) *Middleware



		This is where we define our special middleware interface
		We want all middleware to store the request & response for itself
		and allow access to it from outside, including passing them to
		other middleware
		************************************************/
	Request() (*http.Request, error)
	Response() (*http.Response, error)
	Error(err error) // Build on basic http Handler to add error handling
}

/**************************************************
        Set up the Otis object
**************************************************/
type Otis struct {
	cursor int // Insert User-defined handlers after/before this position

	// This is the section for user generated handler stacking
	stack        []Middleware   // Use this to stack Middleware
	indexInt2Str map[int]string // Use this to look up a Middleware using Cursor
	indexStr2Int map[string]int // Use this to look up a Cursor using Middleware
}

/**************************************************
        Otis Constructor
**************************************************/
func New() *Otis {
	return &Otis{
		0, // Current cursor position for user-defined handlers
		//0, // Current cursor position for error handlers
		make([]Middleware, 0), //
		make(map[int]string),  // Index of Middleware
		make(map[string]int)}  // Index of Middleware in reverse of above
}

/**************************************************
        Set up the Otis methods
**************************************************/

//  Create an empty spot for a middleware to be inserted into
//  the stack. Set the Cursor position of this Otis
//  object to the position of that empty space.
// TODO:
// 1. Add duplicate checking and throw error on finding duplicate name
func (o *Otis) Append(name string, mw Middleware) error {
	// update stack
	o.stack = append(o.stack, mw)

	// update cursor location
	o.cursor = len(o.stack) - 1

	// update index(int2str)
	o.indexInt2Str[o.cursor] = name

	// update index(str2int)
	o.indexStr2Int[name] = o.cursor

	return nil
}

/************************************************************************/
//  Create an empty spot for a middleware to be inserted into
//  the stack. Set the Cursor position of this Otis
//  object to the position of that empty space.
// TODO:
// 1. Add duplicate checking and throw error on finding duplicate name
func (o *Otis) Insert(before string, name string, mw Middleware) error {
	tmp := make([]Middleware, len(o.stack), (cap(o.stack) + 1))
	copy(tmp, o.stack)
	o.stack = tmp

	//      2. Set cursor to position in slice of handlerNamr
	o.cursor = o.indexStr2Int[before]

	copy(o.stack[o.cursor+1:], o.stack[o.cursor:])

	//      3. Use copy to move the upper part of the slice
	//          out of the way and open a hole at o.cursor as set above
	copy(o.stack[o.cursor+1:], o.stack[o.cursor:])

	// insert the new middleware into empty spot
	o.stack[o.cursor] = mw

	// update index(int2str)
	// update all numbers in sequence starting from current cursor position
	// to map names to new index numbers due to insertion

	for i := len(o.indexInt2Str); i > o.cursor; i-- {
		str := o.indexInt2Str[i-1]
		o.indexInt2Str[i], o.indexStr2Int[str] = str, i
	}

	o.indexInt2Str[o.cursor] = name

	// update index(str2int)

	o.indexStr2Int[name] = o.cursor

	// update cursor location
	o.cursor = len(o.stack) - 1

	return nil
}

/************************************************************************/

func (o *Otis) NameIndex(name string) int {
	return o.indexStr2Int[name]
}

/************************************************************************/

func (o *Otis) IndexName(index int) string {
	return o.indexInt2Str[index]
}

/**************************************************
        Test if we made it this far
**************************************************/

type Mid struct{}

func NewMid() *Mid {
	return &Mid{}
}

func (m *Mid) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, r)
}

func (m *Mid) Request() (req *http.Request, err error) {
	return nil, nil
}

func (m *Mid) Response() (res *http.Response, err error) {
	return nil, nil
}

func (m *Mid) Error(err error) {
	fmt.Println(err)
}

func main() {

	s := New()
	s.Append("first_middleware", NewMid())
	s.Append("mid_middleware", NewMid())
	s.Append("last_middleware", NewMid())

	// 	fmt.Println(s.NameIndex("first_middleware"))

	s.Insert("first_middleware", "before_first_mw", NewMid())

	s.Append("lastlast_middleware", NewMid())

	s.Insert("mid_middleware", "newmid_mw", NewMid())

	// 	fmt.Println(s.NameIndex("first_middleware"))

	er := errors.New("\n\nAll systems are Go!")
	fmt.Println(er)
}
