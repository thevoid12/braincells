# SCP
- scp stands for secure copy
- SCP (secure copy) is a command-line utility that allows you to securely copy files and directories between two locations.
- When transferring data with scp, both the files and password are encrypted so that anyone snooping on the traffic doesn’t get anything sensitive. SCP uses the **SSH protocol** for both authentication and encryption.
```bash
scp [OPTION] [user@]SRC_HOST:]file1 [user@]DEST_HOST:]file2
```
https://linuxize.com/post/how-to-use-scp-command-to-securely-transfer-files/

```bash
scp file.txt remote_username@10.10.0.2:/remote/directory
```
 #####  Scp from mac data from pi to mac:
 ```bash
 scp -r void@192.168.0.112:/home/void/void-lapse/timelapse_photos ./timelapse_photos/
 ```
source is the rasp pi. since it is remote we give the user and ip. dest is the folder in mac

---
### Use SCP to Transfer Files Between Your Local System and EC2
Now that the EC2 instance accepts SSH connections with your key, you can use SCP to transfer files between your local system and EC2.

- **we are currently at the local machine**
#### files From Local Machine to EC2:

```bash
scp -i ~/.ssh/my-ec2-key.pem /path/to/local/file ec2-user@<EC2_PUBLIC_IP>:/path/to/remote/directory
```
### files From EC2 to Local Machine:

```bash
scp -i ~/.ssh/my-ec2-key.pem ec2-user@<EC2_PUBLIC_IP>:/path/to/remote/file /path/to/local/directory
```

Replace:

my-ec2-key.pem with your private key file.
ec2-user with the appropriate username (e.g., ubuntu for Ubuntu instances).
<EC2_PUBLIC_IP> with the public IP of your EC2 instance.
/path/to/local/file and /path/to/remote/file with the actual paths of the file you wish to transfer.
