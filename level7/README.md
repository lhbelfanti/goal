# Exercise #1
1. Create a value and assign it to a variable.
2. Print the address of that value.

# Exercise #2
1. Create a person struct
2. Create a func called “changeMe” with a *person as a parameter
    1. change the value stored at the *person address
        1. important
            1. to dereference a struct, use (*value).field
                1. p1.first
                2. (*p1).first
            2. “As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), x.f is shorthand for (*x).f.”
                1. https://golang.org/ref/spec#Selectors
3. Create a value of type person
    1. print out the value
4. In func main
    1. call “changeMe”
5. In func main
    1. print out the value
