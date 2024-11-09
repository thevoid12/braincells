#! /bin/bash
echo "Enter Name:"
read name #read command takes input from the user
echo "Entered Name is $name"

#multiple input
echo "Enter 3 names"
read name1 name2 name3
echo "the entered names are: $name1,$name2,$name3"

#enter input in the same line (using p flag)
read -p "Enter your name: " name
echo "Name is $name"

#enter input in the same line and hide the user input (using s flag)
read -sp "Enter your password: " password
echo "your password is $password"

# get input as a array (using a flag)
echo "Enter two name"
read -a names
echo "the entered 2 names are: Name1 -> ${names[0]}, Name2 -> ${names[1]}"

# the value goes and stores in a default system variable called reply 
# when no variable is assigned
echo "Enter a text"
read #not giving any variable to store the text that has been read
echo " the enetered text is $REPLY" 

#passing arguments
#$0 prints the file name itself, and the value starts from $1
 echo $0 $1 $2 $3  '->this is a test' # this denotes that we are going to print 3 values which we take as argument when the script is run

#passing argument as array
args=("$@") # the arguments get stored in the array args
#here args[0] wont print the file name but it prints the arg at index 0
echo ${args[0]} ${args[1]} ${args[2]} ${args[3]} '-> this is the result'
echo $@ '-> printing all the inputs in one shot'

#printing the number of arguments passed
args=("$@") # the arguments get stored in the array args
echo $@ '-> printing all the inputs in one shot'
echo 'the number of arguments passed is:' $#
