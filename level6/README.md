# Exercise #1
1. Create a func with the identifier foo that returns an int
2. Create a func with the identifier bar that returns an int and a string
3. Call both funcs
4. Print out their results

# Exercise #2
1. Create a func with the identifier foo that
    1. takes in a variadic parameter of type int
    2. pass in a value of type []int into your func (unfurl the []int)
    3. returns the sum of all values of type int passed in
2. Create a func with the identifier bar that
    1. takes in a parameter of type []int
    2. returns the sum of all values of type int passed in

# Exercise #3
Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

# Exercise #4
1. Create a user defined struct with
    1. the identifier “person”
    2. the fields:
        1. first
        2. last
        3. age
2. Attach a method to type person with
    1. the identifier “speak”
    2. the method should have the person say their name and age
3. Create a value of type person
4. Call the method from the value of type person

# Exercise #5
1. Create a type SQUARE
2. Create a type CIRCLE
3. Attach a method to each that calculates AREA and returns it
    1. circle area= π r 2
    2. square area = L * W
4. Create a type SHAPE that defines an interface as anything that has the AREA method
5. Create a func INFO which takes type shape and then prints the area
6. Create a value of type square
7. Create a value of type circle
8. Use func info to print the area of square
9. Use func info to print the area of circle

# Exercise #6
Build and use an anonymous func

# Exercise #7
Assign a func to a variable, then call that func

# Exercise #8
1. Create a func which returns a func
2. Assign the returned func to a variable
3. Call the returned func

# Exercise #9
A “callback” is when we pass a func into a func as an argument. For this exercise,
1. Pass a func into a func as an argument

# Exercise #10
Closure is when we have “enclosed” the scope of a variable in some code block. For this hands-on exercise, create a func which “encloses” the scope of a variable.
