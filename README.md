Otis.
====
Hard to Handle

```
// Create a handler chain of common handlers that others can inherit
CommonHandlers := Otis.New()

// Assign a base set of handlers that will be inherited by all other handlers
CommonHandlers.First("name1", FunctioncallB(args))
CommonHandlers.Next("name2", FunctioncallC(args))
CommonHandlers.Next("name3", FunctioncallD(args))


// Create custom handler
CustomHandlers := Otis.New()

// INHERITANCE
// Insert another Otis chain starting at index 0
CustomHandlers.Inject(CommonHandlers.Handlers)


CustomHandlers.Next("name4", Functioncall2(args))

// INHERIT AFTER ITEM
// Insert Otis chain starting at index returned by (After)
CustomHandlers.After("name4").Inject(CommonHandlers.Handlers)


CustomHandlers.Last("name10", Functioncall10(args))
CustomHandlers.After("name4").Insert("name7", Functioncall3(args))
CustomHandlers.Before("name7").Insert("name6", Functioncall4(args))


// Output current list in a formatted string obj
CustomHandlers.Inspect()



mux = NewMux()
mux.Get("/hello", CustomHandlers.Handle())

http.ListenAndServe(":8080", mux)
```
