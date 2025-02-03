# IAM
- Identity and access manager

## Roles vs policy:
- ### roles:
  - An IAM role is an identity that you can assume to gain temporary access to AWS resources.
  - Roles do not have permanent credentials (such as a username and password). Instead, they provide temporary security credentials (access keys) for a session.
  - Typically used for applications, AWS services, or users who need to perform specific actions without creating dedicated IAM users.
  - Facilitates cross-account access or access to AWS services without embedding credentials.  
  - role can only be assumed meaning they do not provide permanent credentials.
  
## IAM security tool
 ### IAM creadential report:
 - this is at **account level**
 - a report that lists all your account's users and the status of their various credentials
 ### IAM Access Advisor/ Last Access:
 - this is at **user level**
 - Access advisor shows the service permissions granted to a user and when those services were last accessed. this info can be used to revise the policies (what are the the services he might need based on his usage pattern from access advisor)
  - this shows which services are accessed by that user and when
  - go to user-> select the user -> choose access advisor
## IAM policy Json:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:ListAllMyBuckets",
        "s3:GetBucketLocation"
      ],
      "Resource": "arn:aws:s3:::*"
    },
    {
      "Effect": "Allow",
      "Action": "s3:ListBucket",
      "Resource": "arn:aws:s3:::amzn-s3-demo-bucket",
      "Condition": {"StringLike": {"s3:prefix": [
        "",
        "home/",
        "home/${aws:username}/"
      ]}}
    },
    {
      "Effect": "Allow",
      "Action": "s3:*",
      "Resource": [
        "arn:aws:s3:::amzn-s3-demo-bucket/home/${aws:username}",
        "arn:aws:s3:::amzn-s3-demo-bucket/home/${aws:username}/*"
      ]
    }
  ]
}
```
- **version:** just a version for the policy json
- **id:** The Id element specifies an optional identifier for the policy. can be some uuid or anything.
- **Statement:** The Statement element is the main element for a policy. This element is required. The Statement element can contain a single statement or an array of individual statements. Each individual statement block must be enclosed in curly braces { }. For multiple statements, the array must be enclosed in square brackets [ ].
  - statement consists of the following:
      - sid
      - effect
      - principle
      - not Principal
      - action
      - not action
      - resource
      - not resource
      - condition
         - **sid:** You can provide a Sid (statement ID) as an optional identifier for the policy statement. You can assign a Sid value to each statement in a statement array. You can use the Sid value as a description for the policy statement. IAM does not expose the Sid in the IAM API. You can't retrieve a particular statement based on this ID.
         - **effect:** The Effect element is required and specifies whether the statement results in an allow or an explicit deny. Valid values for Effect are Allow and Deny. The Effect value is case sensitive.
         - **Principal:** usual principal which we use in certificates.
         Use the Principal element in a resource-based JSON policy to specify the principal that is allowed or denied access to a resource.
         - **action:** The Action element describes the specific action or actions that will be allowed or denied. Statements must include either an Action or NotAction element. Each AWS service has its own set of actions that describe tasks that you can perform with that service. 
         - **resource:** The Resource element in an IAM policy statement defines the object or objects that the statement applies to. Statements must include either a Resource or a NotResource element.You specify a resource using an Amazon Resource Name (ARN). 
# IAM DB AUTHENTICATION:
- you can authenticate to your DB instance using AWS Identity and Access Management (IAM) database authentication. IAM database authentication works with MySQL and PostgreSQL. With this authentication method, you don’t need to use a password when you connect to a DB instance. Instead, you use an authentication token.

- An authentication token is a unique string of characters that Amazon RDS generates on request. Authentication tokens are generated using AWS Signature Version 4. Each token has a lifetime of 15 minutes. You don’t need to store user credentials in the database, because authentication is managed externally using IAM. You can also still use standard database authentication.
