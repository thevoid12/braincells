## secure cookies
- When you are setting cookies in your application, there are settings that you will need to
use in order to ensure the data within your cookie remains safe.
- Also, your session should always be passed in a session
cookie (not a persistent cookie, and never saved into local storage). Always use a session
cookie for your session.

### secure flag
- always set secure flag so that we use https. 
-The secure flag ensures that your cookie will only be sent over encrypted (HTTPS)
channels. If an attacker attempts to downgrade your session to HTTP, your web application
will refuse to send the cookie.
- **Set-Cookie: Secure;**
### The HttpOnly Flag
- The naming of this flag is not intuitive, as it has nothing to do with forcing the connection
to be unencrypted (HTTP), and thus makes it confusing for programmers. 
- When this flag is
set on a cookie it means that the cookie cannot be accessed via **JavaScript**; it can only be
changed on the server side.
-  The reason for using this setting on your cookie is to protect
against XSS attacks attempting to access the values in your cookie. 
- You always want to set
this value, in all cookies, as another layer of defense against XSS attacks against the
confidentiality of the data in your cookie:
- **Set-Cookie: HttpOnly;**
### persistence
- If you are collecting sensitive user data or managing your session, your cookie should not
be persistent
- it should self-destruct at the end of the session to protect that data. 
- A cookie that self-destructs at the end of your session is a called a **session cookie**. If the cookie
does not self-destruct at the end of the session, it is called a **persistent cookie** or a **tracking cookie**.
- Set-Cookie: Max-Age=3600;
 ### Domain
 - If you want your cookie to be able to be accessed by domains outside of your own, you must
explicitly list them as trusted domains using the “domain” attribute. Otherwise browsers
will assume your cookies are “host-only,” meaning only your domain can access them and
it will block all other access. This type of built-in protection is considered “secure by
default”; if only all software settings worked this way!
- Set-Cookie: Domain=app.thisisvoid.com;
- If you do not set the subdomain (thisisvoid.com instead of
app.thisisvoid.com), every application and page hosted within that domain will be able to access your cookie. You probably don’t want that.
### Same-Site
- created by google to Google to combat the **cross-site request forgery** (CSRF) vulnerability.
- The SameSite attribute in cookies is a security feature that helps prevent Cross-Site Request Forgery (CSRF) attacks by controlling how cookies are sent along with cross-site requests.

-  ### csrf:
    https://www.youtube.com/watch?v=vRBihr41JTo
    https://www.youtube.com/watch?v=80S8h5hEwTY&t=162s

    - User Login: The user logs into bank.com and gets an authentication session cookie.
    - Malicious Site: The user visits a malicious website (attacker.com) while still logged into bank.com.
    - Forged Request: The malicious site automatically sends a request to bank.com (without the user's knowledge)
    - The browser automatically attaches the user's session cookies to the request.
    - If bank.com does not verify the request origin, the request will be processed as if the user initiated it.
    - User Makes Request to GET /form
    - Server generates a random CSRF token.
    - Token is stored in the server-side session (memory, Redis, or DB).
    - Token is sent to the client as:
        - Hidden form field (<input type="hidden" name="_csrf" value="...">)
        - OR as a custom response header like X-CSRF-Token.
        User Submits Form (POST /submit)
    - Browser automatically attaches session cookies.
    Form body contains the CSRF token.
    - Server checks:
    - Is the session valid?
    - Does the CSRF token in the request match the token in the session?
    - The CSRF token must never be stored in cookies because:
        - Cookies are automatically attached by the browser to every request.
        - An attacker can forge requests with the cookie, but they cannot read hidden form fields or headers due to Same-Origin Policy.
        - remember that the same cookie is shared between two different tabs in the same browser if both tabs access the same origin.
        - How Cookies Work Across Tabs
        - Cookies are always shared across tabs/windows if:
        - The domain, path, and protocol are the same.
        - The cookie is not marked as HttpOnly + Secure (for client-side access).
        - The browser process is the same (not in incognito or different user profiles).
        - CSRF tokens break this assumption
- **Same-site** cookie attribute enforces the rule that cookies can only come from
within the same site.
- Cookies cannot come from cross-site (not your site) origins. The
options are **None** (cookies can be sent from anywhere), **Strict** (only from your domain),
and **Lax** (if you want cookies to be sent when coming from a paticular link or another page, only
send non-sensitive information in these cookies). 
- If you don’t set this value in your cookies,
the default on modern browsers is **lax:/** and on older browsers it’s **none**