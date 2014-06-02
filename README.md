Otis.
====
!So Hard to Handle

The goal is to create a slice of handlers that when called will iterate from the first to the last
and correct errors along the way.



```
// Create a handler chain of common handlers that others can inherit
CommonHandlers := Otis.New()

// Assign a base set of handlers that will be inherited by all other handlers
_ := CommonHandlers.First("name1", FunctioncallB(args))
_ := CommonHandlers.Next("name2", FunctioncallC(args))
_ := CommonHandlers.Next("name3", FunctioncallD(args))


// Create custom handler
CustomHandlers := Otis.New()

// INHERITANCE
// Insert another Otis chain starting at index 0
_ := CustomHandlers.Inject(CommonHandlers.Handlers)


// INHERIT AFTER ITEM
// Insert Otis chain starting at index returned by (After)
_ := CustomHandlers.After("name4").Inject(CommonHandlers.Handlers)


_:= CustomHandlers.Preppend("name3", Functioncall3(args))   // Add before first item == Last()
_:= CustomHandlers.Append("name10", Functioncall10(args))   // Add after last item == Firt()
_:= CustomHandlers.After("name4").Insert("name7", Functioncall3(args))
_:= CustomHandlers.Before("name7").Insert("name6", Functioncall4(args))
_:= CustomHandlers.Delete("name7")
_ := CustomHandlers.Overwrite("name10").Insert("name20", Functioncall20(args))


// Handle errors
// What's a good way?


// Output current list in a formatted string obj
CustomHandlers.Inspect()



mux = NewMux()
mux.Get("/hello", CustomHandlers.Handle())

http.ListenAndServe(":8080", mux)
```
