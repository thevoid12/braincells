#! /bin/bash

#array
array=( 'hi' 'hello' 'how are you' )
echo "${array[@]}"
echo "${!array[@]}" #print the indices of array
echo "${#array[@]}" #print the length of the array
#add element to array
array[3]='testing'

#remove a element from the array
unset array[2] #removes element at index 2

string="void"
echo "${string[@]}"
echo "${string[0]}" # the whole string void is assigned at 0th index
