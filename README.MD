# Interpreter for Monkey Language

This is my implementation of the Monkey Language Interpreter, I built this as I worked through the book "Writing An Interpreter in Go" by Thorsten Ball (https://interpreterbook.com/). 

I made some small extentions just to challenge myself to add a few features:
- Parsing support for escape sequences: \n \t \"
- Support for SETs and some set-related functions
```
Hello csueiras! Type in your Monkey Language commands
>> set(1,2,2,2,2,3)
set(1,2,3)
>> let setA = set(1,2,2,2,3)
>> let setB = set(1,2,3,4,4,4)
>> union(setA, setB)
set(2,1,3,4)
```

The project was fun to work on, I used it as a means to learn Go Lang, and I highly recommend for any newbie to Go!
