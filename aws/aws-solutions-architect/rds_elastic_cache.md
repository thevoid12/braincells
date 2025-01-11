# RDS:
- In AWS there’s a network cost when data goes from one AZ to another
- For RDS Read Replicas within the same region, you don’t pay that fee
- The Read Replicas be setup as Multi AZ for Disaster Recovery (DR)
- The following happens internally for multi az replication:
• A snapshot is taken
• A new DB is restored from the
snapshot in a new AZ
• Synchronization is established
between the two databases
- RDS Custom is Managed Oracle and Microsoft SQL Server Database with OS and database customization
- the difference between custom RDS and RDS is:
  - RDS: entire database and the OS to be managed by AWS. we cannot ssh into the rds servers as it is fully managed. But it gives us all the required functionalities needed 
  - RDS Custom: full admin access to the underlying OS and the database 
- Read Replicas add new endpoints with their own DNS name. We need to change our application to reference them individually to balance the read load.
- Multi-AZ keeps the same connection string regardless of which database is up.
- Storing Session Data in ElastiCache is a common pattern to ensuring different EC2 instances can retrieve your user's state if needed.
- **Aurora Global Databases** allows you to have an Aurora Replica in another AWS Region, with up to 5 secondary regions.
- RDS has **IAM database authentication** to authenticate people based on IAM 
- **Read replica** uses **asyncronous Replication** and **multi AZ** is **syncronous replication**
- to **encrypt RDS** : take a **snapshot**, encrypt using kms and then unencrypt the snapshot when needed
- you can not create encrypted Read Replicas from an unencrypted RDS DB instance.
- 15 Aurora Read Replicas can you have in a single Aurora DB Cluster
- Amazon Aurora supports both **mysql** and **posgress**
- The maximum retention of aurora automatic backup is only **35 days**
