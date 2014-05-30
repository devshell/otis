otis
====
Hard to Handle

```
// Create a base handler that is initially empty 
httphandlerCommon := Otis.New()
httphandlerCommon.First("name1", functioncallB(args))
httphandlerCommon.Next("name2", functioncallB(args))
httphandlerCommon.Next("name2", functioncallB(args))

htthandler := Otis.New()

// Assign a base set of handlers




htthandler0 := Wrapper.New(httphandler0) // sets these to be at top of stack (FIFO)
httphandler1.First("name1", functioncall(args))
httphandler1.Next("name2", functioncall2(args))
httphandler1.Last("name4", functioncall10(args))
httphandler1.After("name2").Insert("name3", functioncall3(args))

mux = NewMux()
mux.Get("/hello", httphandler)

http.ListenAndServe(":8080", mux)
```
