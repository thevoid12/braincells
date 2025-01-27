# Messaging 
### SQS:
![sqs](img/32.png)
- Default retention of messages: 4 days, maximum of 14 days
-  Limitation of 256KB per message sent
- Encryption:
  - In-flight encryption using HTTPS API
  - At-rest encryption using KMS keys
  - Client-side encryption if the client wants to perform encryption/decryption itself
- **SQS – Message Visibility Timeout:**
  -  After a message is polled by a consumer, it becomes invisible to other consumers
  - By default, the “message visibility timeout” is 30 seconds
  - That means the message has 30 seconds to be processed
  - After the message visibility timeout is over, the message is “visible” in SQS
