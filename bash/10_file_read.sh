#! /bin/bash 

#method 1
while read p
do 
  echo $p
done < 9_loops.sh #the file content is redirected to the while loop and each and every line is echoed

# method 2
cat 9_looops.sh | while read p
do
  echo $p
done

# method 3

while IFS='' read -r p
do 
  echo $p
done <  9_loops.sh
