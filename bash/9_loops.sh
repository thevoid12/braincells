#! /bin/bash

#while loops
#syntax
# while [ condition ]
# do  
#   command 1
#   command 2
#   command 3
# done

#printing from 1 to n using while loop
n=1
while (($n < 10 )) #[ $n -lt 10 ] this also works  we use the typical comparison operator 
do 
 echo "the value of n is: $n"
 n=$(( n+1 ))
done


#until loop
# in until loop if the condition is false the commands are executed unitl it becomes true
# but in while loop command is executed only if it is true 

# until [ condition ]
# do  
#   command 1
#   command 2
#   command 3
# done

# for loops
# for val in  somelist
# do
# command 1
# command 2
# command 3
# done

# for (( i=0;i<10;i++ ))
# do 
# command 1
# command 2
# command 3
# done

for val in 1 2 3 4 5
do 
echo $val
done

for val in {1..10} # 1..10 means iterate from 1 to 10
do 
echo $val
done

for val in {1..10..2} # {start..end..increment by 2}1..10 means iterate from 1 to 10
do 
echo $val
done

for (( i=0; i<5; i++ ))
do
echo $i
done
