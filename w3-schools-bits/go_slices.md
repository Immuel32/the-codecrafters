***Slices***


Slices
Slices are similar to arrays,.but the are more flexible meaning the length of slices can grow and shrink.


There are several ways to create a slice:
- Using the []datatype{values} format
- Create a slice from an array
- Using the make() function

***Creating a slice with data type***
The common way to declare and initialize a slice
    myslice := []int{1,2,3} 


***Function to return the length of a slice***
- len() :  Returns the length of the slice (the number of elements in the slice)
- cap() :  Returns the capacity of the slice (the number of elements the slice can grow or shrink to)