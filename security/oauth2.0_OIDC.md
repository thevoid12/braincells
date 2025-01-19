# oauth2.0 and open id connect (OIDC) 


> goated video : https://www.youtube.com/watch?v=996OiexHze0
>ppt: https://drive.google.com/file/d/1UyPqnrGnCCJ7PeIY-rDV-3tRprIEprBB/view?usp=sharing


- oauth2.0 for authorization
- openid connect for authentication

#### authorization vs Authentication:

##### Authentication:
- Authentication is the process of verifying the identity of a user, system, or entity. It ensures that the person or system is who they claim to be.

##### Authorization:
- what are the permission that an authenticated user has
- authorization is the process of determining what an authenticated user or system is allowed to access or do. It defines permissions and access levels.
##### Analogy:
Authentication: Showing your ID to prove who you are.

Authorization: Being allowed into a restricted area based on your ID.
| **Aspect**            | **Authentication**                          | **Authorization**                          |
|------------------------|---------------------------------------------|--------------------------------------------|
| **Purpose**            | Verifies identity                           | Grants or restricts access                 |
| **When it happens**    | Before authorization                        | After authentication                       |
| **Mechanisms**         | Passwords, biometrics, MFA                  | Role-based access control (RBAC), permissions |
| **Example**            | Logging into a system                       | Accessing specific files or features       |



![authorization flow](img/2.png)

- takes advantage of both back and front channel
    - ##### backchannel and front channel:
      - The forward channel is used for user-facing interactions and operates through the browser, while the back channel is used for secure server-to-server communication.
      - While the back channel uses POST requests, it is still more secure than the forward channel because:
        - All communication is encrypted with HTTPS.
        - The client application authenticates itself using a client secret.
        - Sensitive data (e.g., authorization code, access token) is never exposed to the browser.
        - Authorization codes are short-lived, reducing the risk of misuse
      - In contrast, the forward channel is less secure because:
        - It relies on the browser to transmit sensitive data (e.g., authorization code) via redirects.
        - The browser is a public environment and is more vulnerable to attacks (e.g., man-in-the-middle, phishing, or malicious extensions).
        - The client application has no control over the security of the browser or the network.

![back channel auth flow](img/3.png)

- the user intraction part is done using a front channel(browser) but since we cant fully trust the browser, we use a back channel to do the final part of the flow acess token exchange and get the the resources part  

#### OAuth 2.0 Authorization Code Flow
OAuth 2.0 Authorization Code Flow to highlight the security differences:

##### Forward Channel (Less Secure)
- The user clicks "Sign in with Google" on the client application.
- The client app redirects the user's browser to Google's authorization server.
- The user logs in and consents to the requested permissions.
- Google redirects the user's browser back to the client app with an authorization code in the URL.(redirect uri or callback)

**Risk:** The authorization code is exposed to the browser and could be intercepted. but thats of no use because an authorization code alone means nothing because access token is more important and we get access token only when we provide authorization code + a secret

##### Back Channel (More Secure)
- The client app takes the authorization code and sends it directly to Google's token endpoint via a POST request (over HTTPS).

- **Security:** The POST request is encrypted with HTTPS, and the client app authenticates itself using a client secret.

- Google responds with an access token.
**Security:** The access token is transmitted securely over HTTPS and is never exposed to the browser.
```bash
POST /token HTTP/1.1
Host: authorization-server.com
Content-Type: application/x-www-form-urlencoded

grant_type=authorization_code
&code=AUTH_CODE
&redirect_uri=https://client-app.com/callback
&client_id=CLIENT_ID
&client_secret=CLIENT_SECRET
```
---

- oauth was never defined to be used for **authentication**. it was desgined for **authorization**. what oauth does is checks the scope that the authenticated user has (which is  authroization ).
- OAuth 2.0 is designed primarily for **authorization**, not authentication, because its core purpose is to allow third-party applications to access resources on behalf of a user without exposing the user's credentials (e.g., passwords).
- OAuth 2.0 uses **access tokens** to grant access to resources. These tokens are short-lived and scoped to specific permissions (e.g., read-only access to emails).
- It does not inherently verify the user's identity; it only ensures that the application has permission to access the requested resources.

##### Why OAuth 2.0 is Not for Authentication:
###### 1. Lack of Identity Verification:
- OAuth 2.0 does not provide a standard way to verify the user's identity. It only provides an access token, which does not inherently contain identity information.

- Without additional mechanisms, OAuth 2.0 cannot confirm who the user is—it only confirms that the application has been granted access.
What Happens in OAuth 2.0 with Google
Access Token:

- eg: When you use OAuth 2.0 to access Google APIs (e.g., Google Drive, Gmail), Google issues an access token to the third-party application.

- This access token is a string that represents the authorization granted to the app to access specific resources (e.g., read emails, view files).

- **What the Access Token Represents:**
  - The access token is a proof that the app has been granted permission to access certain resources on behalf of the user.

  - It does not contain detailed information about the user (e.g., name, email, profile picture). Instead, it is just a key that allows the app to make API calls to Google's services.

  - OAuth 2.0 does not define how to verify the user's identity. The access token alone does not tell you who the user is. it only tells you that the app has permission to access certain resources.
##### 2. Misuse for Authentication:

- Many developers misuse OAuth 2.0 for authentication by assuming that the presence of an access token implies the user's identity has been verified. This can lead to security vulnerabilities.

- For example, if an attacker steals an access token, they could impersonate the user without ever verifying their identity.

##### 3. Inconsitency in authentication implementation:
- OAuth 2.0 does not define how user information (e.g., name, email) should be retrieved. This is left to the implementation, which can lead to inconsistencies. everybody implement their own way on top of oauth to use oauth to authenticate. which is not the standard practice

---
## openID connect:
- to solve the above problem of needing user info to authenticate a layer is written on top of oauth2.0 which is OIDC.
- OIDC adds an identity layer to OAuth 2.0, enabling authentication:

**ID Tokens:**

- OIDC introduces ID tokens, which are JSON Web Tokens (JWTs) containing information about the user's identity (e.g., name, email, etc.).

- These tokens are signed and can be verified to ensure the user's identity.

**Standardized User Info:**

OIDC provides a standard way to retrieve user information using the /userinfo endpoint.

**Authentication Flow:**

OIDC uses OAuth 2.0 flows but adds steps to verify the user's identity and return an ID token alongside the access token.

User Info Endpoint:

Google provides a /userinfo endpoint that the app can call to retrieve basic information about the user (e.g., name, email, profile picture).

This requires the app to make a separate API call using the access token.

Scopes:

When requesting authorization, the app specifies scopes (e.g., https://www.googleapis.com/auth/userinfo.profile) to indicate what user information it needs.
- we can add openid in scope which will return us the openid jwt
- Google then grants access to this information, but it is not part of the access token itself.

![OIDC](img/4.png)
| **Step**               | **OAuth 2.0**                              | **OpenID Connect (OIDC)**                  |
|-------------------------|--------------------------------------------|--------------------------------------------|
| **Token Issued**        | Access token                              | Access token + ID token                    |
| **Identity Information**| Not included in the access token          | Included in the ID token                   |
| **User Info Retrieval** | Requires additional API call to `/userinfo`| ID token contains user info; `/userinfo` is optional |
| **Use Case**            | Accessing Google Drive or Gmail           | Logging into a website using "Sign in with Google" |

---
## refresh token:
####  Refresh Tokens are used to:
- Get a new access token when the current one expires
- Maintain long-term access to Google APIs
**Important points about refresh tokens:**
  - They don't expire (but can be revoked)
  - Should be stored securely (database/encrypted storage)
  - You only get them once during the first authorization
  - You only get them when requesting "offline" access
  - One refresh token can be used multiple times to get new access tokens
**when is refresh token used:**
- Google's access tokens typically expire after 1 hour
- When you make an API request and get a 401 (Unauthorized) response
- Before making an API call if you know the token has expired (by tracking expiry time)

**benefits of including refresh tokens in authentication:**
- Users stay logged in longer
- More secure than storing access tokens long-term
- Better user experience (no frequent re-logins)
- Required for offline access to user data
- Follows OAuth 2.0 best practices

**the access token expiring doesn't automatically "log out" the user in the way you might think. Here's what actually happens:**
At 10:00 AM when the access token expires:
The user still sees your website/app
Their browser session is still active
They haven't been kicked out
BUT, when your application tries to:
Access Google APIs
Get user information
Validate the user's identity
The request will fail because the access token is expired.
