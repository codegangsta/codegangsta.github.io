---
layout: post
title: "Dependency Injection in Ruby"
date: 2013-04-04 08:26
comments: true
categories: 
---

Dynamically typed languages like Ruby are an interesting beast when it comes to dependency injection. The topic itself has been debated in the Ruby community every once in a while:

<blockquote class="twitter-tweet"><p>Dependency injection is not a virtue in Ruby <a href="http://t.co/4w2qSCfo" title="http://bit.ly/ZtzSz9">bit.ly/ZtzSz9</a>, and then there's the immediate criticism: <a href="http://t.co/Q1O9TvGm" title="http://bit.ly/ZtzQHw">bit.ly/ZtzQHw</a></p>&mdash; Nicola Iarocci (@nicolaiarocci) <a href="https://twitter.com/nicolaiarocci/status/288215967622914048">January 7, 2013</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

Let me start off by first saying that dependency injection is not a catch-all for managing dependencies in Ruby. There are many other ways to abstract hard dependencies within your code. Today I want to specifically talk about programming to interfaces in Ruby and how to solve the “dependencies” problem in that space.

## Programming to Interfaces
A common argument is that Ruby does not need dependency injection because Module mixins exist in the language. Module mixins surely allow implementation details to be abstracted and shared so that change can happen mostly in one place. A Ruby developer must think of the intensions of the Class/Implementation they are composing to correctly identify the solution to managing dependencies.

## The Module Approach
Modules provide a clean solution to resolving dependencies, the basic premise looks like this:
{% codeblock lang:ruby %}
module Downloadable
  def download(url)
    # download the url and location
  end
end
{% endcodeblock %}
{% codeblock lang:ruby %}
class SDK
  include Downloadable

  def install
    download(url)
    ...
  end 
end
{% endcodeblock %} 

This is a great example of how module mixins can be an excellent solution to our problem. `Downloadable` is completely reusable and can be safely mixed into any other class we want. 

### The Danger
The dangerous part about all of this is the *abuse* cases with mixins. When we type:
{% codeblock lang:ruby %}
include Downloadable
{% endcodeblock %} 
    
We are essentially allowing the `Downloadable` module to add any number of methods it desires to our class. This pattern can easily veer codebases out of control if it is not monitored properly, so it is not quite a *silver bullet* for dependency resolution in Ruby.

## The attr_inject Approach
Finding out the Modules was a solution to a particular problem and not an solution to the entire problem space encouraged me to go out an write a elegant and Rubyesque dependency injection framework. I call it [attr_inject](https://github.com/jeremysaenz/attr_inject).

Before I go into details I want to spend a little time sharing what I think is a great explanation of the problem we are seeking to solve: 

>Using dependency injection to shape code relies on your ability to recognize that the responsibility for knowing the name of a class and the responsibility for knowing the name of a message to send to that class may belong in different objects. -Sandi Metz. Practical Object-Oriented Design in Ruby

I love this statement because it essentially says that when I explicitly want to inject something I have to stop thinking of the dependency as a class and start thinking of it as an interface. For instance:
{% codeblock lang:ruby %}
class User
  def name
    "foobar"
  end

  def age
    22
  end
end
{% endcodeblock %}
{% codeblock lang:ruby %}    
class Project
  def initialize(user)
    @user = user
  end

  def username
    @user.name
  end
end
{% endcodeblock %}
As far as `Project` is concerned, `@user` is an object that responds to the `name` message. This makes dependency injection shine when  used properly.

### Flexible DI
In *Practical Object-Oriented Design in Ruby*, Sandi walks the reader through a real world example, progressively improving how dependencies are managed through injection. Here is a sample of the ultimate solution:

{% codeblock lang:ruby %}
def initialize(args)
  @foo = args.fetch(:foo, 40)
  @bar = args.fetch(:bar, 18)
  @baz = args[:baz]
end
{% endcodeblock %}
    
Injecting through the constructor with a hash provides immense flexibility and allows the class to not care about the Module or Class name of the dependency. It also solves the problem of *argument order dependencies*.

### One Step Further
The [attr_inject](https://github.com/jeremysaenz/attr_inject) gem takes this pattern one step further. Examine the code below:
{% codeblock lang:ruby %}
def Project
  attr_inject :user
  attr_inject :sdk

  def initialize(args)
    inject_attributes(args)
  end
end
{% endcodeblock %}
    
This snippet of code injects the `args` hash into the `Project`’s `user` and `sdk` attributes. If the `user` and `sdk` attributes do not exist, `inject_attributes` will throw an exception by default. This of course can be configured:
{% codeblock lang:ruby %}
attr_inject :user, :required => false
{% endcodeblock %}    
or
{% codeblock lang:ruby %}
attr_inject :user, default => some_default
{% endcodeblock %}
    
### Even Further...
If you are like me and you don’t always want to pass arguments into the object initializer or you want to inject dependencies whenever you want, you can use the `Injector` class.

{% codeblock lang:ruby %}
class Main
  include Inject

  def initialize()
    injector = Injector.new
    injector.map :user, User.new
    injector.map :sdk, SDK.new  
    
    injector.apply Project.new
  end
end
{% endcodeblock %}
{% codeblock lang:ruby %}
class Project
  attr_inject :user
  attr_inject :sdk
end
{% endcodeblock %}
    
### Factories
Sometimes the dependency would like to know some information about the object it is injected into, or the dependency needs to be created specially for each injection. This is where factories come into play.

{% codeblock lang:ruby %}
injector = Injector.new
injector.factory :logger do |target|
  Logger.new target
end
{% endcodeblock %}
    
Our logger object will be created upon every injection. You can request the logger object on your class the same way as any other dependency:

{% codeblock lang:ruby %}
attr_inject :logger
{% endcodeblock %}

##Conclusion    
[attr_inject](https://github.com/jeremysaenz/attr_inject) is an elegant and scalable way to manage dependencies. Documentation is a bit sparse at the moment and I hope to get a tutorial up and going very soon. Please feel free to fork it on Github an play around with it. Let me know what you think!
