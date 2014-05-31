otis
====
Hard to Handle

```
// Create a handler chain of common handlers that others can inherit
CommonHandlers := Otis.New()

// Assign a base set of handlers that will be inherited by all other handlers
CommonHandlers.First("name1", functioncallB(args))
CommonHandlers.Next("name2", functioncallC(args))
CommonHandlers.Next("name3", functioncallD(args))


// Create custom handler
CustomHandlers := Otis.New()

// INHERITANCE
// Insert another Otis chain starting at index 0
CustomHandlers.First("Inherited", CommonHandlers.Handle())


CustomHandlers.Next("name4", functioncall2(args))
CustomHandlers.Last("name10", functioncall10(args))
CustomHandlers.After("name4").Insert("name7", functioncall3(args))
CustomHandlers.Before("name7").Insert("name6", functioncall4(args))


// Output current list in a formatted string obj
CustomHandlers.Inspect()



mux = NewMux()
mux.Get("/hello", CustomHandlers.Handle())

http.ListenAndServe(":8080", mux)
```
