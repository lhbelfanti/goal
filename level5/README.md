# Exercise #1
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
1. first name
2. last name
3. favorite ice cream flavors

Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

# Exercise #2
Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.

# Exercise #3
1. Create a new type: vehicle.
    1. The underlying type is a struct.
    2. The fields:
        1. doors
        2. color
2. Create two new types: truck & sedan.
    1. The underlying type of each of these new types is a struct.
    2. Embed the “vehicle” type in both truck & sedan.
    3. Give truck the field “fourWheel” which will be set to bool.
    4. Give sedan the field “luxury” which will be set to bool. solution
3. Using the vehicle, truck, and sedan structs:
    1. using a composite literal, create a value of type truck and assign values to the fields;
    2. using a composite literal, create a value of type sedan and assign values to the fields.
4. Print out each of these values.

# Exercise #4
Create and use an anonymous struct.
