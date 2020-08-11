# Exercise #1
Using a COMPOSITE LITERAL:
1. create an ARRAY which holds 5 VALUES of TYPE int
2. assign VALUES to each index position.
3. Range over the array and print the values out.

Using format printing:
1. print out the TYPE of the array

# Exercise #2
Using a COMPOSITE LITERAL:
1. create a SLICE of TYPE int
2. assign 10 VALUES
3. Range over the slice and print the values out.

Using format printing:
1. print out the TYPE of the slice

# Exercise #3
Using the code from the previous example, use SLICING to create the following new slices which are then printed:
1. [42 43 44 45 46]
2. [47 48 49 50 51]
3. [44 45 46 47 48]
4. [43 44 45 46 47]

# Exercise #4
Follow these steps:
1. start with this slice
    1. x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
2. append to that slice this value
    1. 52
3. print out the slice
4. in ONE STATEMENT append to that slice these values
    1. 53
    2. 54
    3. 55
5. print out the slice
6. append to the slice this slice
    1. y := []int{56, 57, 58, 59, 60}
7. print out the slice

# Exercise #5
To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:
1. start with this slice
    1. x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
2. use APPEND & SLICING to get these values here which you should ASSIGN to a variable “y” and then print:
    1. [42, 43, 44, 48, 49, 50, 51]

# Exercise #6
Create a slice to store the names of all of the states in the United States of America. What is the length of your slice? What is the capacity? Print out all of the values, along with their index position in the slice, without using the range clause. Here is a list of the states:

` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,

# Exercise #7
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

"James", "Bond", "Shaken, not stirred"

"Miss", "Moneypenny", "Helloooooo, James."

Range over the records, then range over the data in each record.

# Exercise #8
Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.

`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

# Exercise #9
Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop

# Exercise #10
Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop
