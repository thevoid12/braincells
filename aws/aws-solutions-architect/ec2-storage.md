# EC2 instance storage
- EBS Volumes are created for a specific AZ. It is possible to migrate them between different AZs using EBS Snapshots.
- AMIs are built for a specific AWS Region, they're unique for each AWS Region. You can't launch an EC2 instance using an AMI in another AWS Region, but you can copy the AMI to the target AWS Region and then use it to create your EC2 instances.
- EFS is a network file system (NFS) that allows you to mount the same file system on EC2 instances that are in different AZs.
- EC2 Instance Store provides the best disk I/O performance but will diappear on termination
### EBS Volume types:
  - general Purpose ssd volumes
    - GP2
    - GP3
  - Provisioned IOPS ssd volumes
    - io1
    - io2 block express
  - throughput optimized ssd volume
    - st1
  - cold HDD volume
    - sc1
##### 1.general Purpose ssd volumes
- balances price and performance, normal,balanced good volume type
- cost effective
- gp3 is newer generation as compared to gp2 so higher iops
- useful for boot volume, test and development env, wherever things are balanced  

##### 2.Provisioned IOPS ssd volumes
- critical business needs where iops are more important
- best for database related works
- this type of volume supports ebs multiattach

##### 3. HDD
- cannot be used as boot volume
- st1 is throughput optimized which can be used for log processing,data warehousing etc
- sc1 is cold hdd which is slowest of all so can be used for infrequent access volumes

### EBS multiattach:
  - attach the same ebs volume for multiple ec2 but in the **same AZ** only  
  - certain types of ebs volumes are only allowed to do this (provisioned iops volumes io1 and io2)
  - can do read and write at the same time
  - can attach maximum of **16 ec2**
  - the ec2's must use a file system which is cluster aware ie xfs etc
