# secure shell hash
## how to key based ssh into ec2:
**step 1:**
1. Generate Your Key Pair (If You Haven't Yet)
If you haven't already created a key pair for your local system, you can generate it by running the following command in your local machine terminal:
```bash
ssh-keygen -t rsa -b 2048 -f ~/.ssh/my-ec2-key.pem
```
This will generate two files:
my-ec2-key.pem: Your private key (kept safe).
my-ec2-key.pem.pub: Your public key.
**step 2:**
- copy Your Public Key to the EC2 Instance
Since you're already logged into the EC2 instance via EC2 Instance Connect, the next step is to add your public key to the **~/.ssh/authorized_keys** file for the ec2-user (or whichever user you're using).

Copy the public key to the EC2 instance. First, open the public key file on your local machine:
```bash
cat ~/.ssh/my-ec2-key.pem.pub
```
- Paste the public key into the EC2 instance:
- On your EC2 instance (while logged in through EC2 Instance Connect), create or edit the **~/.ssh/authorized_keys** file for the ec2-user:
```bash
mkdir -p ~/.ssh
nano ~/.ssh/authorized_keys
```
- Paste your public key (from the my-ec2-key.pem.pub file) into the **authorized_keys** file and save the changes.
- Set correct permissions for the .ssh directory and authorized_keys file:
```bash
chmod 700 ~/.ssh
chmod 600 ~/.ssh/authorized_keys
```

- Exit EC2 Instance Connect
```bash
exit
```
## ssh using key:
```bash
chmod 400 ~/.ssh/my-ec2-key.pem
ssh -i /path/to/private-key.pem ec2-user@<EC2_PUBLIC_IP>
```
## ssh using certificate:
```bash
ssh -i /path/to/private-key.pem -o CertificateFile=/path/to/ssh-certificate-cert.pub ec2-user@<EC2_PUBLIC_IP>
```
eg:
```bash
ssh -i void-test-ssh-key -o CertificateFile=void-test-ssh-user-cert.pub void@192.168.0.112
```
