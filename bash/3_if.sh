#! /bin/bash

# syntax
# if [condition]
# then 
#   statement
# fi

#integer comparison 
# integer comparison
# -eq - is equal to  - if [ $a -eq $b ]
# -ne - not equal to - if [ $a -nq $b ]
# -gt - is greater than - if [ $a -gt $b ]
# -ge - greater than or equal to - if [ $a -ge $b ]
# -lt - is less than - if [ $a -lt $b ]
# -le - less than or equal to - if [ $a -le $b ]
#  to use symbols we need to wrap them around (()) as well as []
# < -is less than - (( $a<$b ))
# <= - is less than or equal to - (( $a<=$b ))
# > -greater than -  (( $a>$b ))
# >= - greater than or equal to (( $a>=$b ))

count=100
if [ $count -gt 50 ] # leave space between [ count ]
then
echo "the statement is true"
fi


if [ $count -ne 50 ]
then
echo "the if statement is true"
fi

if (( $count  > 50 ))
then 
echo "symbolic comparision is true"
fi

if [ $count  > 50 ]
then 
echo "symbolic comparision is true"
fi


# #string comparision- for string comparision directly use [[ ]] for everything for uniformity
#  = and == are same - is equal to - if [ $a == $b ]  or [[ ]]
#  != - not equal to - if [ $a != $b ] or [[ ]]
# to compare using symbols >,< use double [[ ]]
#  < - is less than, in ascii alphabetical order - if [[ $a < $b ]]  
#  > - is greater than, in ascii alphabetical order - if [[ $a > $b ]]  
#  -z - string is null, that is, has zero length

word1="void"
word2="void"
if [ $word1 == $word2 ]
then 
echo "true"
fi


word1="void"
word2="p"
if [[ $word1 > $word2 ]]
then 
echo "comparision true"
else echo "comparision false"
fi

wordA="hello"
wordB="hello"
if [[ $wordA != $wordB ]] #leave gap between word and !=
then
echo "both words are not equal"
elif [[ $wordA == $wordB ]]
then
echo "both words are equal"
else
echo "cant judge"
fi
