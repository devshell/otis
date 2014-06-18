Otis.
====
!So Hard to Handle

http://gofiddle.net/#dnS0djdV
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
// Insert another Otis chain starting at index 0
_ := CustomHandlers.Inject(CommonHandlers.Handlers)


// INHERIT AFTER ITEM
// Insert Otis chain starting at index returned by (After)
_ := CustomHandlers.After("name4").Inject(CommonHandlers.Handlers)


_:= CustomHandlers.Append("name10", Functioncall10(args))   // Add after last item == Firt()
_:= CustomHandlers.After("name4").Append("name7", Functioncall3(args))
_:= CustomHandlers.Before("name7").Append("name6", Functioncall4(args))
_:= CustomHandlers.Remove("name7")
_ := CustomHandlers.Replace("name10").Append("name20", Functioncall20(args))


// Handle errors
/*
Check for errors on each handler return, and if there is an error check the ErrHandlers map for a specific
handlerName_error entry, and if there isn't, then check for an "error" entry, and if there isn't one,
go to the defaults map, and check the "error" value.
*/


// Output current stack in a formatted string obj
CustomHandlers.Inspect()



mux = NewMux()
mux.Get("/hello", CustomHandlers.Handle())

http.ListenAndServe(":8080", mux)
```
