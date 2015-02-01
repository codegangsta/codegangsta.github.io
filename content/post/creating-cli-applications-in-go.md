+++
title = "On Distributing Command line Applications: Why I switched from Ruby to Go"
date = "2013-07-21T16:28:53-08:00"
slug = "creating-cli-applications-in-go"
keywords = ["go", "golang", "cli", "ruby"]
categories = ["go", "golang", "cli", "ruby"]
+++

For as long as I have been using it, my go-to language for creating CLI applications was Ruby. After all, Ruby is a robust, mature language with many things to make us developers want to gyrate with joy (okay maybe that is just me). 

To this day I still hold the opinion that Ruby provides the most elegant amount of expression when it comes to my everyday programming activities. Whether it is the language or the community that fosters clean code, I feel like an artisan every time I sit down to write an application or library in Ruby.

There is something so elegant and beautiful about creating a command line application with the following code:

``` ruby
desc "A todo application"
app :todo do
  desc "add a todo"
  command :add do
    puts "Your todo has been added"
  end
  desc "remove a todo"
  command :remove do
    puts "your todo has been removed"
  end
end
```

## Distributing Ruby
While the code is beautiful, distributing Ruby programs is not. Now - I don't want to make you believe that program distribution should be an easy thing, because it definitely is not. There are so many things that can possibly go wrong with an application when you expect it to leave your programming and test environments to run perfectly on your end users machines. Let me say it again; *Distributing programs is **hard work!***

I'm sure you are thinking, "Why not distribute your app via Rubygems?". After attempting the distribution of a command line application via Rubygems I came to the following conclusion: **Any application not intended exclusively for Rubyists should not be distributed via Rubygems.** Luckily [I am not the only one](http://mitchellh.com/abandoning-rubygems) who holds this position on distribution, so that leads those of us that wish to use Ruby on a long road to properly distribute a command line application.

## Custom Ruby and Vendoring
The next natural option with Ruby would be to distribute your entire application independent of RubyGems. So you *really* only have one good option here, and that is to [compile a custom version of Ruby](http://yehudakatz.com/2012/06/05/tokaido-status-update-implementation-details/) to distribute with each OS installer (Yes, you must have an installer) of your application. Just so you don't mess with a users current installation you must compile under a custom prefix; something along the lines of `/usr/local/<yourapp>/ruby`. 

Once this is done it is important to make sure you vendor and distribute all of your gems that you use for your application. This can get nasty when it comes to native gems that rely on local dynamic libraries. Anyway, I hope you get the picture that *distributing a ruby app to end users can be a hassle.*

## Building it in Go
If you are adventurous you have probably played with, or heard of the [Go programming language](http://golang.org/). In fact I have started to blog about it for the past couple of weeks. My recent love affair with Go has led me to believe that distributing applications to end users in Go is a much simpler process than languages like Ruby.

Since Go is a compiled language, binaries of your app can be precompiled for each platform that you wish to distribute for. While not as portable as C, (and who can be, C compilers are everywhere!) Go provides a lightweight set of tools that can compile on many platforms for those who wish their end users to compile from source.

## Introducing cli.go
While I loved the simplicity in distributing an application in Go, I missed that expressiveness that I got from Ruby, so (shameless plug) I rolled my own CLI library called [cli.go](https://github.com/codegangsta/cli).  

[cli.go](https://github.com/codegangsta/cli) provides that expressive, "Document as you code" model for building command line applications in Go. It is easily distributable and super fast! Building real command line apps is as easy as this:

``` go
package main

import "os"
import "github.com/codegangsta/cli"

func main() {
  app := cli.NewApp()
  app.Name = "boom"
  app.Usage = "make an explosive entrance"
  app.Action = func(c *cli.Context) {
    println("boom! I say!")
  }

  app.Run(os.Args)
}
```

Go check out the [cli.go](https://github.com/codegangsta/cli) repo. **Star** it, **fork** it, **Contribute** to it. Let's make Go the go-to language for command line tools!

Keep writing beautiful code.
