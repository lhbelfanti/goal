# Exercise #1
In addition to the main goroutine, launch two additional goroutines
1. Each additional goroutine should print something out
2. Use waitgroups to make sure each goroutine finishes before your program exists


# Exercise #2
This exercise will reinforce our understanding of method sets:
1. Create a type person struct
2. Attach a method speak to type person using a pointer receiver
    1. *person
3. Create a type human interface
    1. to implicitly implement the interface, a human must have the speak method
4. Create func “saySomething”
    1. have it take in a human as a parameter
    2. have it call the speak method
5. Show the following in your code
    1. you CAN pass a value of type *person into saySomething
    2. you CANNOT pass a value of type person into saySomething
6. Here is a hint if you need some help
    1. https://play.golang.org/p/FAwcQbNtMG

    Receivers       Values
    ------------------------
    (t T)           T and *T
    (t *T)          *T

# Exercise #3
1. Using goroutines, create an incrementer program
    1. have a variable to hold the incrementer value
    2. launch a bunch of goroutines
        1. each goroutine should
            1. read the incrementer value
                1. store it in a new variable
            2. yield the processor with runtime.Gosched()
            3. increment the new variable
            4. write the value in the new variable back to the incrementer variable
2. Use waitgroups to wait for all of your goroutines to finish
3. The above will create a race condition.
4. Prove that it is a race condition by using the -race flag
5. If you need help, here is a hint: https://play.golang.org/p/FYGoflKQej

# Exercise #4
Fix the race condition you created in the previous exercise by using a mutex
1. It makes sense to remove runtime.Gosched()

# Exercise #5
Fix the race condition you created in exercise #3 by using package atomic

# Exercise #6
Create a program that prints out your OS and ARCH. Use the following commands to run it
1. go run
2. go build
3. go install
