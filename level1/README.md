# Exercise #1

1. Using the short declaration operator, **ASSIGN** these **VALUES** to **VARIABLES** with the **IDENTIFIERS** “x” and “y” and “z”
    1. 42
    2. “James Bond”
    3. true
2. Now print the values stored in those variables using
    1. a single print statement
    2. multiple print statements

# Exercise #2
1. Use var to **DECLARE** three **VARIABLES**. The variables should have package level scope. Do not assign **VALUES** to the variables. Use the following **IDENTIFIERS** for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that **TYPE**).
    1. identifier “x” type int
    2. identifier “y” type string
    3. identifier “z” type bool
2. In func main
    1. print out the values for each identifier
    2. The compiler assigned values to the variables. What are these values called? Answer: zero value

# Exercise #3
Using the code from the previous exercise,
1. At the package level scope, assign the following values to the three variables
    1. for x assign 42
    2. for y assign “James Bond”
    3. for z assign true
2. In func main
    1. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
    2. print out the value stored by variable “s”

# Exercise #4
For this exercise
1. Create your own type. Have the underlying type be an int.
2. Create a **VARIABLE** of your new **TYPE** with the **IDENTIFIER** “x” using the “VAR” keyword
3. In func main
    1. print out the value of the variable “x”
    2. print out the type of the variable “x”
    3. assign 42 to the **VARIABLE** “x” using the “=” **OPERATOR**
    4. print out the value of the variable “x”

# Exercise #5
Building on the code from the previous example
1. At the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
    1. eg:<br>
    type hotdog int<br>
    var x hotdog<br>
    var y int
2. In func main
    1. this should already be done
        1. print out the value of the variable “x”
        2. print out the type of the variable “x”
        3. assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
        4. print out the value of the variable “x”
    2. now do this
        1. now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
            1. then use the “=” operator to ASSIGN that value to “y”
            2. print out the value stored in “y”
            3. print out the type of “y”

# Exercise #6
1. Take this quiz: https://goo.gl/forms/dfwmTuYTe5ox8nyt1 (There is not code for this).
