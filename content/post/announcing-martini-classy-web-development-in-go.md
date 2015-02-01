+++
title = "Announcing Martini: Classy Web Development in Go"
date = "2013-11-14T16:52:36-08:00"
keywords = ["golang", "web", "martini"]
categories = ["golang", "web", "martini"]
+++

Last night I tweeted about a web framework for Go called Martini.
Martini has a non-intrusive design, and has awesome routing and middleware support.

[http://martini.codegangsta.io/](http://martini.codegangsta.io/)

Below is a basic "Hello world!":

``` go
package main

import "github.com/codegangsta/martini"

func main() {
  m := martini.Classic()
  
  m.Get("/", func() string {
    return "Hello world!"
  })

  m.Run()
}
```

So far the response has been fantastic, I am excited to see which ways the Golang community can come together to make something awesome:

<blockquote class="twitter-tweet"><p><a href="https://twitter.com/codegangsta">@codegangsta</a> damn! that is really impressive. This has to be the best lightweight <a href="https://twitter.com/search?q=%23golang&amp;src=hash">#golang</a> web framework yet! Kudos.</p>&mdash; Dave Cheney (@davecheney) <a href="https://twitter.com/davecheney/statuses/400941765076611072">November 14, 2013</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

Go check it out, you won't regret it. [https://github.com/codegangsta/martini](https://github.com/codegangsta/martini)
