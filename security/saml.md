# SAML
- goated video: https://www.youtube.com/watch?v=l-6QSEqDJPo&t=30s
- Security Assertion Markup Language (SAML)
- SAML is an open standard used for **authentication**.
- Single sign on is achieved through saml. It achieves this objective by centralizing user authentication with an identity provider(IP). 
- Web applications can then leverage SAML via the identity provider to grant access to their users. 
- This SAML authentication approach means users do not need to remember multiple usernames and passwords. Our application becomes more secure because we arent storing passwords anymore (our identity provider stores all the password on our behalf) or forgotten our password.
- identity providers scale on their own, and they have resources to enable multiple layer of security like MFA.
- we can use it anywhere but it is widely used at enterprises or organisation authentication
- when the user tries to access a site, the identity provider passes the SAML authentication to the service provider, who then grants the user entry.
## SAML terminologies:
- **Identity Provider (IDP):**
the guys who manage our security credentials
- **Service Provider(SP):**
the application which wants to sso
- **SAML Request:**
the request service provider (our app) sends to identity provider
- **SAML Response:**
the response that identity provider sends to service provider (our app)
- **Assertion:**

- **XML Signature(Dsig):**

- **Assertion Consumer Service:**

- **Attributes:**
- **relay state:**
where the user was before the authentication
- **SAML Trust:**
- **Metadata:**
   