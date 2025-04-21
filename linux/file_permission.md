# Linux file permissions

-rwxrwxrwx
The permission consist of 3 group Owner group Others(same other)
The order is always rwx. Think this as an octal number so r=2^2=4 , w=2^1=2, x=2^0=1.  So 4+2+1=7
0751 means 7 permission to owner(all readwrite and execute), 5 to group(read and write only),1 to others (write only)
And gate with whatever permission we want to check if we have the required permission to access the directory
The letter r is read, w is write and x is executable permission 
The Owner is usually the root user
