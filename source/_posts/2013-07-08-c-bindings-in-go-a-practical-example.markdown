---
layout: post
title: "C Bindings in Go: A Practical Example Part 1"
date: 2013-07-08 06:40
comments: true
categories: golang cgo go spotify libspotify libmockspotify
---
I have been playing with [Go](http://golang.org/) lately. The language itself has a lot going for it, one of which is a decent set of interop with existing C code. Today I am going to walk you through a practical example of how this is done by showing some code that I have been working on lately with Go.

## The Simplest Hello World
Assuming you have your [Go development environment all set up,](http://golang.org/doc/code.html) go ahead and create a new go file with the following contents:

``` go
package main

func main() {
  println("Hello world")
}
```

Running this should obviously print "Hello world". Let's move on!

## Using libspotify
Lets start creating some C bindings for libspotify! libspotify is a C api distributed by Spotify that basically allows us to do everything a typical spotify app can do; log in, play, music, manipulate playlists - all available to us with libspotify.

For the sake of keeping this post short and to the point, we will just be creating a `sp_session` object. This session will lay some groundwork for us to log in and do the fun stuff libspotify allows us to do. Why don't we dive in:

``` go
package main

type Session struct {
  session *C.sp_session
}

func main() {
  println("Hello world")
}
```

Running this will result in an error
``` text
./bindings.go:4: undefined: C
```

This is because we actually need to use some special magic that Go provides for us, the C package.

``` go
package main

import "C"

type Session struct {
  session *C.sp_session
}

func main() {
  println("Hello world")
}
```

This gives us a different error:
``` text
error: 'sp_session' undeclared (first use in this function)
error: (Each undeclared identifier is reported only once
```

Cool, so it sees the C package, but sp_session doesn't exist. This is because we need to include libspotify using the `#include` directive for the C preprocessor. First make sure you have libspotify installed. On a Mac you can run `brew install libspotify`. Now try to run the following code:

``` go
package main

/*
#include <libspotify/api.h>
*/
import "C"

type Session struct {
  session *C.sp_session
}

func main() {
  println("Hello world")
}
```

This should compile and run since `sp_session` now resides in the global C namespace! This works great for things that exist in the header, but it breaks down when we try to initialize our session:

``` go
package main

/*
#include <libspotify/api.h>
*/
import "C"
import "unsafe"

type Session struct {
  session *C.sp_session
}

func main() {
  key := "appkey_good"
  session := Session{}
  appkey := C.CString(key)
  appkey_size := len(key)

  var config = C.sp_session_config {
    api_version:          C.SPOTIFY_API_VERSION,
    cache_location:       C.CString(".spotify/"),
    settings_location:    C.CString(".spotify/"),
    user_agent:           C.CString("spotify for go"),
    application_key:      unsafe.Pointer(appkey),
    application_key_size: C.size_t(appkey_size)}

  C.sp_session_create(&config, &session.session)
}
```

Will result in the following error:
``` text
Undefined symbols for architecture x86_64:
  "_sp_session_create", referenced from:
      __cgo_90280c6a3021_Cfunc_sp_session_create in bindings.cgo2.o
     (maybe you meant: __cgo_90280c6a3021_Cfunc_sp_session_create)
ld: symbol(s) not found for architecture x86_64
collect2: ld returned 1 exit status
```
This is because we didn't link libspotify. This is an easy fix, add this in the comment above `import "C"`:
``` text
#cgo LDFLAGS: -lspotify.12
```

Your code should compile and link fine. You will get a SIGSEGV in your application, this is becuase we specified a bogus app key. We will cover this in detail during Part 2 of *C Bindings in Go: A Practical Example*.

Until next time!
