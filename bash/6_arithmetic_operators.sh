#! /bin/bash

#$(( ... )) is used for arithmetic evaluation in bash
num1=20
num2=5
echo $(( $num1 +$num2 ))
num3=$num1+$num2 #this still comcatinates the result 20+5
echo $num3 # op:20+5
num3=$(( $num1+$num2 )) 
echo $num3
#an alernate way to do arthmetic is to use expr command
echo $(expr $num1 + $num2 )
echo $(expr $num1 - $num2 )
echo $(expr $num1 \* $num2 )
echo $(expr $num1 / $num2 )
echo $(expr $num1 % $num2 )

num1=20.5
num2=5
echo "20.5+5"|bc 
echo $num1+$num2 |bc

echo $num1-$num2 |bc
echo $num1\*$num2 |bc
echo "scale=4;$num1/$num2" |bc # scale helps us to set the number of decimal part we want in our result
echo $num1%$num2 |bc
echo "scale=3;sqrt($num2)" |bc -l #-l calls the math library which has sqrt for bc to use
echo "scale=3;($num2)^($num2)" |bc -l  # 5 power 5
