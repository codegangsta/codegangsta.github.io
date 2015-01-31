+++
title = "My Thoughts on Martini"
categories = ["go", "golang", "web", "martini", "negroni"]
date = "2014-05-19T10:51:49-08:00"
keywords = []
+++
 
It's never easy to take honest criticism, especially when the target of the criticism is considered your 'baby'. Earlier this week, [Stephen Searles posted an entirely honest review](http://stephensearles.com/?p=254) of Martini and how, despite the popularity and hype, there are numerous reasons why Martini should not be used.

I have been asked by many people what my thoughts are regarding the reasons not to use Martini. I figure that it would be best to compile these thoughts into a blog post so I can shed some light on my opinions and how they may have changed over time. So here we go, my reflections on Martini...

## Martini's Popularity was a HUGE surprise

>Arguably, many people at the time were young Go developers.

This may be an obvious one. I was not expecting the positive attention that Martini received when I made the initial announcement. I was extremely surprised to see my little experiment of doing dependency injection in Go become popular so quickly. It was so well received that my inbox was flooded with emails for weeks, Martini blog posts were all over the internet, and hundreds of contributions were added to both the Martini and martini-contrib GitHub repos.

Not only was this unexpected, I was still very young as a Go developer. Arguably, many people at the time were young Go developers. I doubted some of the decisions I made with Martini, but soon those doubts were drowned out by the positive reenforcement relayed by the Go community. In my mind Martini was a great fit for the Go community because people liked it! It's popular, so it must be the best choice right? It's completely idiomatic right? Right?!

## Martini is not Idiomatic

As I grew from a baby gopher to a well respected Go developer in the community, this was the toughest pill for me to swallow. *Martini, and it's design, is simply not idiomatic Go.* This is not to say that Martini is not well designed, I feel like it is one of my better demonstrations of API design that I've had in my career. There have been many excellent patterns for web development that have been executed in the Martini codebase and many very high quality web packages found in martini-contrib. The contributions can be attributed to the great community surrounding Martini and the design of the framework itself.

Despite all of these facts, Martini is still not idiomatic Go. The mantras surrounding to Go community are simplicity, familiarity with the stdlib, and explicit interactions with the type system. Martini does not line up with the way the stdlib was designed, and it therefore can never be considered idiomatic. This doesn't make Martini wrong, it is just not going in the direction that the Go community as a whole is going.


## Martini reflection is flawed

>The real tradeoff with reflection is one of complexity.

One of the crowning features of Martini is it's reflective dependency injection. From a modularity perspective it sounds awesome! Martini handlers only get the dependencies they need, services can be injected, mocked, swapped out, and fulfilled by this method of dependency provisioning. The problem is that the reflection comes at a cost, both in performance and complexity.

In most situations, the performance overhead is negligible compared to the other components required in building a web application/api. The real tradeoff with reflection is one of complexity. Rather than have a strict interface for implementing and extending a web application, Martini allows you to *inject all the things* and that leads to a level of indirection that, while modular, requires a certain amount of cognitive overhead to fully understand what is actually going on.


## So... Now What?


{% img http://makeameme.org/media/created/who-am-i-bmup6m.png %}


If you are using Martini, please continue using it. If you are interested in using Martini, it is still a fantastic framework and arguably the most productive set of tools for building web applications in Go. It is well designed and has a great community around it. Martini is not going away, and I'm not going to stop supporting Martini.

All that said, I still want to accomplish the goal I was originally set out to accomplish when I first created Martini; promote the creation of awesome reusable web components for Go. I recently have been thinking about how I could take the good parts of Martini and combine them into a package that is simple, beautiful, non-intrusive, and most of all *idiomatic* Go.


## Introducing Negroni

[Negroni](http://negroni.codegangsta.io) is a idiomatic approach to middleware in `net/http`. I wrote it as an alternative to Martini and it attempts to accomplish similar goals:

* Non-intrusive design

* Compatibility with net/http

* Super easy to use

* Promote good practices

[Negroni](http://negroni.codegangsta.io) does one thing well, net/http middleware. It comes with the same default middleware that Martini comes with; Logging, Panic Recovery, and Static file serving. The API is simple and intuitive. Most of all, this middleware stack does not use reflection or dependency injection, and building middleware for Negroni requires no dependencies outside of net/http. Here is a quick example of how to use Negroni to compose a feature full middleware stack:

``` go
package main

import (
  "github.com/codegangsta/negroni"
  "net/http"
  "fmt"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page!")
  })

  // Use the default middleware.
  n := negroni.Classic()
  // ... Add any other middlware here
  
  // add the router as the last handler in the stack
  n.UseHandler(mux)
  n.Run(":3000")
}
```

[Negroni](http://negroni.codegangsta.io) is not considered a framework as it does not have a built in router. There are many great http routers already available to to the Go community already. The goal of Negroni is to be focused on the problem that Martini was originally set out to resolve.

I'm interested in hearing your thoughts on the package, if you are interested in hearing more from me on the subject, feel free to comment below. Also, [be sure to give Negroni a star on GitHub](http://negroni.codegangsta.io) so we can create a vibrant community around reusable net/http handlers for Go!

Keep on building awesome things.
