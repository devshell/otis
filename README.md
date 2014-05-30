otis
====
Hard to Handle

```
// Create a base handler that is initially empty 
CommonHandlers := Otis.New()

// Assign a base set of handlers that will be inherited by all other handlers
CommonHandlers.First("name1", functioncallB(args))
CommonHandlers.Next("name2", functioncallC(args))
CommonHandlers.Next("name3", functioncallD(args))


CustomHandlers := Otis.New()
CustomHandlers.First("Common", CommonHandlers)
CustomHandlers.Next("name4", functioncall2(args))
CustomHandlers.Last("name10", functioncall10(args))
CustomHandlers.After("name4").Insert("name7", functioncall3(args))
CustomHandlers.Before("name7").Insert("name6", functioncall4(args))

mux = NewMux()
mux.Get("/hello", httphandler)

http.ListenAndServe(":8080", mux)
```
