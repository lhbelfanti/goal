# Exercise #1
Starting with [this code](https://play.golang.org/p/dJlIVxLfPkS), marshal the []user to JSON. There is a little bit of a curve ball in this hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a package.

# Exercise #2
Starting with [this code](https://play.golang.org/p/b_UuCcZag9), unmarshal the JSON into a Go data structure. Hint: https://mholt.github.io/json-to-go/

# Exercise #3
Starting with [this code](https://play.golang.org/p/BVRZTdlUZ_), encode to JSON the []user sending the results to Stdout. Hint: you will need to use json.NewEncoder(os.Stdout).encode(v interface{})

# Exercise #4
Starting with [this code](https://play.golang.org/p/H_q75mpmHW), sort the []int and []string for each person.

# Exercise #5
Starting with [this code](https://play.golang.org/p/BVRZTdlUZ_), sort the []user by
1. age
2. last

Also sort each []string “Sayings” for each user
1. print everything in a way that is pleasant
