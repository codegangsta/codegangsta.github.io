---
layout: post
title: "Practical Go: Build Tags"
date: 2013-07-11 07:42
comments: true
categories: cgo go golang libmockspotify libspotify practicalgo spotify buildtags
---
In our [last post](blog/2013/07/08/c-bindings-in-go-a-practical-example/) we wrote a simple set of bindings for libspotify in Go. By the end of the post we had an example compiling, but we had a bad API key for our spotify application. One obvious way to recify this would be to grab an API key if you are a Spotfy Premium user. Another workaround is to use a *mock* library to make sure the code is working the way we want. We will link this mock library with **Build tags.**

## Mocking Spotify
Start by installing libmockspotify:
``` text
git clone git@github.com:mopidy/libmockspotify.git
cd libmockspotify
./autogen.sh
./configure
make
sudo make install
```

Once libmockspotify is installed, we can link to that instead of the real libspotify:


``` go
package main

/*
#include <libspotify/api.h>
#cgo LDFLAGS: -lmockspotify.2
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
  println("Session created!")
}
```

The following code should output:
``` text
Session created!
```

## Using build tags
It's great that we have libmockspotify for testing purposes, but it doesn't make sense to always mock spotify in our project. This is where **Build Tags** come into play. We can use build tags to conditionally choose which library to link to. 

Let's link both libspotify and libmockspotify in our file:
``` go
/*
#include <libspotify/api.h>
#cgo LDFLAGS: -lmockspotify.2
#cgo LDFLAGS: -lspotify.12
*/
```

This approach obviously will not work as is, so we will instead link each library conditionally based on the *mock* build take that we will define:
``` go
/*
#include <libspotify/api.h>
#cgo mock LDFLAGS: -lmockspotify.2
#cgo !mock LDFLAGS: -lspotify.12
*/
```

This will link libmockspotify when the *mock* build tag is defined, otherwise we will use the real libspotify. So then how do we define build tags? Simple:

``` text
go run -tags mock myfile.go
```

Thats enough for today, have fun with build tags! I will leave you with one more pro tip!

In my `~/.vim` folder I have a ftplugin for Go that includes a binding for running tests:
``` text
nmap <Leader>t :!go test -tags test <cr>
```

You see that I includ the *test* tag in all of my test running? This makes it convenient when linking mock libraries specific for running tests!

Until next time!
