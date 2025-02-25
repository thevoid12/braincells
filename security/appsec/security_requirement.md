# Security Requirement

### encryption
- Encryption is two-way, in that you can jumble up the information into an
unreadable mess, and then “decrypt” it back into its original form. Hashing is one-way; the original value can never be recovered.
- the data as it flows between the application and the API should be encrypted to protect the privacy of your user.
- any data that gets transmitted should be encrypted
### never trust system input
- Input to your application means literally anything and everything that is not a part of your application or that could have been manipulated outside of your application.
eg: User input on the screen (for instance, entering search phrases into a field),
Information from a database (even the database you designed for your app) 
- anything that comes from outside validate and sanitize
- if we write an application in a non memory safe language do a bounds checking to make sure it doesnt overflow. attackers can buffer overflow (overwrite parts of memory) which not handled would crash our application
- There is no purpose in performing validation of the data after you have used it. It must be the very first thing you do after receiving input into your application.
- When issuing an error message to the screen to reject user input,
if you decide to show the user’s input, be aware that it may be malicious
and therefore potentially cause your program to malfunction. Always
encode the output using HTML encoding
- ndefault administrator group
- we create sub org org admin group
- that obj should be in obj table
- while sharing we will 
- obj table new record
- find out risk preset we are doing (owner
find what all the child arg, create stsme name, group id find., use that )

- Verifying that all third-party components are not known to be vulnerable is a quick-and-easy win in
regard to understanding how secure your application is
- **various statergies to verify known vulnerabilites in third-party components:**
  - use more than 1 tool to check verify. different tools check different stuff so more than 1 tool is preferable
  - regularly scan our code repository  (daily, or at least weekly), as well as scan every time
you release code to production. try building this pipeline into your ci/cd

### Security Headers: Seatbelts for Web Apps
- [security headers](security_headers.md)
