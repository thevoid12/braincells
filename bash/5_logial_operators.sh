#! /bin/bash

# and operator
age=25
 if [ $age -gt 10 ] && [ $age -lt 30 ]
 then
 echo "Its a valid age"
 else 
 echo "Not a valid age"
 fi

# or logical operator
 if [ $age -gt 30 ] || [ $age -lt 30 ]
 then
 echo "Its a valid age"
 else 
 echo "Not a valid age"
 fi
