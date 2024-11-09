#! /bin/bash

#this is how you write a comment
echo "hello world"

#system variables
echo our shell name is $BASH
echo our shell version name is $BASH_VERSION
echo $HOME
echo $PWD

#user defined varliable
name=void # if you do name = void its a error no space between name and =
echo my name is $name
# which is same as
echo "my name is $name"

val10=10 #note that variable name shoudnt start with a number
echo value $val10
