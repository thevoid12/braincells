#! /bin/bash

# this is a jq tutorial https://jqlang.org/tutorial/

#$() captures the output of commands
OP=$(curl 'https://api.github.com/repos/jqlang/jq/commits?per_page=5')
# echo ${OP} 
# # pretty print the json
# echo ${OP} | jq .

# gets the first index of the json array
echo ${OP} | jq '.[0]'

# create a new json where message is commit key message key's value and so on
echo ${OP} | jq '.[0] | {message: .commit.message, name: .commit.committer.name}'

# for every index one after the other (just like looping through)
echo ${OP} | jq '.[] | {message: .commit.message, name: .commit.committer.name}'

# placing all the values inside a array so the result becomes array of json
echo ${OP} | jq '[.[] | {message: .commit.message, name: .commit.committer.name}]'

