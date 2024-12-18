# IAM
- Identity and access manager

## Roles vs policy:
- ### roles:
  - An IAM role is an identity that you can assume to gain temporary access to AWS resources.
  - Roles do not have permanent credentials (such as a username and password). Instead, they provide temporary security credentials (access keys) for a session.
  - Typically used for applications, AWS services, or users who need to perform specific actions without creating dedicated IAM users.
  - Facilitates cross-account access or access to AWS services without embedding credentials.  
  - role can only be assumed meaning they do not provide permanent credentials.
  
