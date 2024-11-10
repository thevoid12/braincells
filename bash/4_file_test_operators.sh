#! /bin/bash
#  echo -e "Enter the name of the file: \c" #-e interprets the back slash
#  read file_name
#  if [ -e $file_name ]  # -e command checks if file exists or not
#  then 
#  echo "$file_name found"
#  else
#  echo "$file_name not found"
#  fi


#  echo -e "Enter the name of the regular file: \c" #-e in echo interprets the back slash
#  read file_name
#  if [ -e $file_name ]  # -f command checks if file exists or not and it is a regular file
#  then 
#  echo "$file_name regualar file found"
#  else
#  echo "$file_name regular file not found"
#  fi


#  echo -e "Enter the name of the directory: \c" #-e in echo interprets the back slash
#  read dir_name
#  if [ -d $dir_name ]  # -d command checks if directory exists 
#  then 
#  echo "$dir_name dir found"
#  else
#  echo "$dir_name dir not found"
#  fi


# # check if the file is empty or not
#  echo -e "Enter the name of the  file: \c" #-e in echo interprets the back slash
#  read file_name
#  if [ -s $file_name ]  # -s flag checks if the file is empty or not
#  then 
#  echo "$file_name file not empty"
#  else
#  echo "$file_name file empty"
#  fi

# #append text at the end of the file using cat
# echo -e "Enter the fileName for the text to be appended:\c"
# read filename
# file_name="${file_name}.txt" #appending .txt at the end of the filename
# if [ ! -f $file_name ] #check if the file doesnt exists
# then 
# touch $file_name
# chmod +w $file_name #giving write permission
# fi
# echo "Enter text that will be appended in the file use ctrl+d to exit:"
# #read $text
# cat >> $file_name #appending text at the end
# echo "the text has been successfully added into the file $filename"


#append text at the end of the file without using cat
echo -e "Enter the fileName for the text to be appended:\c"
read filename
file_name="${file_name}.txt" #appending .txt at the end of the filename
if [ ! -f $file_name ] #check if the file doesnt exists
then 
touch $file_name
chmod +w $file_name #giving write permission
fi
echo "Enter text that will be appended in the file:"
read text
echo $text >> $file_name #appending text at the end of the file
echo "the text has been successfully added into the file $filename"
