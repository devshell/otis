Otis.
====
!So Hard to Handle

http://gofiddle.net/#N0GI87Ju

http://gofiddle.net/#ooYxsoGe

The goal is to create a stack of middleware that when called will iterate from the first to the last
and correct errors along the way. At the end, the output should be dependent on the adapter being used
including support out of the box for the http.Handler interface, and the go-tigertonic return handler.

Slicing notes: https://code.google.com/p/go-wiki/wiki/SliceTricks

```
// Create a handler chain of common handlers that others can inherit
CommonHandlers := Otis.New()

// Assign a base set of handlers that will be inherited by all other handlers
_ := CommonHandlers.Append("name1", FunctioncallB(args))


// Create custom handler
CustomHandlers := Otis.New()

// INHERITANCE
// Insert another Otis chain starting at cursor location
_ := CustomHandlers.InjectAppend(CommonHandlers.Handlers)


_:= CustomHandlers.Append("name10", Functioncall10(args))   // Add after last item == Firt()
_:= CustomHandlers.Insert("name10", "name7", Functioncall3(args))
_:= CustomHandlers.Insert("name7", "name6", Functioncall4(args))
_:= CustomHandlers.Delete("name7")

// INHERIT AFTER ITEM
// Insert Otis chain before index location
_ := CustomHandlers.InjectInsert("name6", CommonHandlers.Handlers)

// The following is basically the same as delete "name10", and insert
_ := CustomHandlers.Overwrite("name10", "name20", Functioncall20(args))


// Handle errors

//Check for errors on each handler return, and if there is an error, run the middleware's
//Error function, otherwise, run the error in middleware at stack index 0.



// Output current list in a formatted string obj
CustomHandlers.Inspect()



mux = NewMux()
mux.Get("/hello", CustomHandlers.Handle())

http.ListenAndServe(":8080", mux)
```
