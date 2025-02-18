## Security fundamentals
#### Insider Threats
- someone has permission to that system,network or data but performs mallicious activity with the given permission
- eg: An employee downloading intellectual property onto a portable drive, leaving the
building, and then selling the information to your competitors
#### Defense in depth
- it  is the idea of having multiple layers of security in case one is not enough
#### Least Privilege
- Giving users exactly how much access and control they need to do their jobs, but nothing
more, is the concept of least privilege.
#### supply chain security
- multiple pieces together form a software. these are part of supply chain. similarly lot of external code (dependency) joins together to form software. these dependency needs to be throughly checked before using them.
- When you plug dependencies into our applications, you are accepting the risks of the code they contain that your application uses. 
#### Security by Obscurity
- The concept of security by obscurity means that if something is hidden it will be “more
secure,” as potential attackers will not notice it. The most common implementation of this
is software companies that hide their source code, rather than putting it open on the
internet 
#### Attack Surface Reduction
- Attack surface reduction means removing anything from your application that is
unrequired
- For instance, a feature that is not fully implemented but you have the button
grayed out, would be an ideal place to start for a malicious actor because it’s not fully tested or hardened yet. Instead, you should remove this code before publishing it to production
and wait until it’s finished to publish it
- Even if it’s hidden, that’s not enough; reduce your attack surface by removing that part of your code.
#### Hard Coding
- Hard coding means programming values into the code, rather than getting the values
organically (from the user, database, an API, etc.).
- if anyone can access the source code they will get the confidential data
- Hard coding is generally considered a symptom of poor software development 
#### Never Trust, Always Verify
- anything that comes from outside verify and validate. even valdate data from your own databases
- #### Usable Security
- If security features make your application difficult to use, users will find a way around it or go to your competitor. There are countless examples online of users creatively
circumventing inconvenient security features; humans are very good at solving problems,
and we don’t want security to be the problem.
- The answer to this is creating usable security features.
- eg: Allowing a fingerprint, facial recognition, or pattern to unlock your personal device
instead of a long and complicated password.
