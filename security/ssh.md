# Secure Shell (SSH) Key-Based Authentication  

This guide does not explain what SSH is; rather, it provides a straightforward set of steps to set up key-based SSH access to an EC2 instance.  

#### How to Set Up Key-Based SSH Access to an EC2 Instance  

#### Step 1: Generate Your Key Pair (If You Haven't Yet)  

If you haven't already created a key pair on your local machine, generate one using the following command:  

```bash
ssh-keygen -t rsa -b 2048 -f ~/.ssh/my-ec2-key
```  

This will generate two files:  

- `my-ec2-key`: Your private key (keep this safe).  
- `my-ec2-key.pub`: Your public key.  

#### Step 2: Copy Your Public Key to the EC2 Instance  

Since you are already logged into the EC2 instance via **EC2 Instance Connect**, the next step is to add your public key to the **~/.ssh/authorized_keys** file for the `ec2-user` (or whichever user you are using).  

1. Open the public key file on your local machine:  

    ```bash
    cat ~/.ssh/my-ec2-key.pub
    ```  

2. Copy the output and paste it into the EC2 instance.  

3. On your EC2 instance (while logged in through EC2 Instance Connect), create or edit the **~/.ssh/authorized_keys** file:  

    ```bash
    mkdir -p ~/.ssh
    nano ~/.ssh/authorized_keys
    ```  

4. Paste your public key (from the `my-ec2-key.pub` file) into the **authorized_keys** file and save the changes.  

5. Set the correct permissions for the `.ssh` directory and `authorized_keys` file:  

    ```bash
    chmod 700 ~/.ssh
    chmod 600 ~/.ssh/authorized_keys
    ```  

6. Exit EC2 Instance Connect:  

    ```bash
    exit
    ```  

#### SSH into the EC2 Instance Using the Key  

Ensure the private key has the correct permissions:  

```bash
chmod 400 ~/.ssh/my-ec2-key
```  

Now, connect to your EC2 instance using:  

```bash
ssh -i ~/.ssh/my-ec2-key ec2-user@<EC2_PUBLIC_IP>
```  

#### SSH Using a Certificate  

If you are using a certificate for authentication, use the following command:  

```bash
ssh -i ~/.ssh/my-ec2-key -o CertificateFile=/path/to/ssh-certificate-cert.pub ec2-user@<EC2_PUBLIC_IP>
```  

#### Example:  

```bash
ssh -i void-test-ssh-key -o CertificateFile=void-test-ssh-user-cert.pub void@192.168.0.112
```  

---
