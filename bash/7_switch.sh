#! /bin/bash

echo -e "Enter sone character :\c"
read value
case $value in
  [a-z] ) #regular exp
    echo "User entered $value a to z" ;;
   [A-Z] )
    echo "User entered $value A to Z" ;;
   [0-9] )
    echo "User entered $value 0 to 9" ;;
   ? )
    echo "User entered $value special character" ;;
     * )
    echo "unknown input" ;;
esac
