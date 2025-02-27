# Security Headers

## Overview
Security headers can usually be applied on either the web server or your code. To apply them, you generally will add one line of code or check a box to configure your web server.

Check security headers at: https://securityheaders.com/

## 1.Content-Security-Policy (CSP)

goated video: https://www.youtube.com/watch?v=J90t0h0AP1U

### Understanding XSS (Cross-Site Scripting Attack)

CSP is one among the mitigations for XSS attacks. Let's understand XSS first:

- Suppose you have an HTML template as output and an input form that sends data to the backend and puts the result back in the template
- If you enter a JS script in the input form and click submit, the output HTML template will have the script tag in it, thus executing the JS code

**Example:**
```html
<p>{{input}}</p>
```

After inputting `<script>alert(1)</script>`, the template becomes:
```html
<p><script>alert(1)</script></p>
```
This script will execute in the browser.

#### Mitigation Attempts
- One way to prevent this is validating the input and escaping the output before parsing the HTML template (by adding escape characters)
- This works well with HTML content, but escaping the output won't work with HTML attributes or when templating in CSS and JS

**Example of attribute vulnerability:**
```html
<p><a href="{{input}}">link</a></p>
```

If we provide a JavaScript URL as input:
```html
<p><a href="javascript:alert(1)">link</a></p>
```

Despite escaping, when we click the link, the JS code executes, allowing the XSS attack. This also applies to event handlers like `onclick={{input}}`.

### Content Security Policy (CSP)

- CSP is meant to be enforced by the browser
- It tells the browser which payloads to render and which to block
- It is set via **HTTP header** by the **web server** for the browser to enforce

#### Basic CSP Settings
- **`default-src: 'self'`**: Blocks everything not coming from your own domain
  - This may break UI elements that load resources from other domains like Google Fonts, third-party CSS, external JS for metrics, etc.
  - These external resources need to be carefully checked with the development team before allowing them

> You can generate CSP headers through this website: https://report-uri.com/home/generate

### Example CSP Policy

**`Content-Security-Policy: default-src 'self'; script-src 'self'; font-src 'none'`**

#### 1. `default-src 'self'`
- A fallback directive that applies to all resource types unless overridden by a more specific directive
- The value `'self'` means resources can only be loaded from the same origin
- Applies to resources such as:
  - Images (`img-src`)
  - Scripts (`script-src`)
  - Styles (`style-src`)
  - Fonts (`font-src`)
  - Frames (`frame-src`)
  - Media files (`media-src`)
  - AJAX/Fetch/XHR requests (`connect-src`)
  - Web Workers (`worker-src`)
  - Any other resource types not explicitly defined

**Example Behavior:**
- If an image `<img src="https://example.com/logo.png">` is loaded from a different domain, it will not load
- If a CSS file `<link rel="stylesheet" href="https://cdn.example.com/style.css">` is linked from an external domain, it will not be allowed

#### 2. `script-src 'self'`
- This directive overrides `default-src` for JavaScript files
- Scripts must be loaded only from the same origin
- If `script-src` is defined, it completely replaces `default-src` for scripts

**Example Behavior:**
- Allowed: `<script src="/static/main.js"></script>` (if hosted on the same domain)
- Blocked: `<script src="https://cdn.example.com/main.js"></script>`
- Blocked: `<script>console.log("Inline script");</script>` (because `'unsafe-inline'` is not allowed)

**Why Override `default-src`?**
- Explicitly defining `script-src 'self'` gives more granular control
- By specifying `script-src`, you ensure scripts follow their own specific rule, independent of other resource types

#### 3. `font-src 'none'`
- This directive blocks the loading of all fonts
- The value `'none'` means no fonts are allowed, not even from the same domain

**Fonts include:**
- Web fonts (.woff, .woff2, .ttf, .otf) loaded via CSS @font-face
- Font files referenced in `<link>` elements
- Any fonts fetched from CDNs or external sources

**Example Behavior:**
```css
@font-face {
  font-family: "MyFont";
  src: url("/fonts/myfont.woff2"); /* Blocked, even though it's hosted locally */
}
```

Also blocked: `<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto">`

**Will `font-src 'none'` Break the UI?**
- If the website relies on web fonts, the UI will look significantly different
- The site will fall back to system fonts, which may change appearance but won't necessarily break functionality
- This is why you should always consult developers before implementing CSP

### Additional Security for CSP Using Nonce

- A nonce (number used once) is a randomly generated, unique token added to a CSP policy
- It allows specific inline scripts or styles to execute while blocking all other unauthorized scripts
- Helps protect against XSS attacks without using the insecure `'unsafe-inline'` directive

**How it works:**
1. Generate a random nonce value on the server for each request
2. Include the same nonce value in:
   - The Content-Security-Policy header
   - The `<script>` or `<style>` tags that need to be allowed
3. The browser will only execute inline scripts/styles with a matching nonce

**Example:**
```
Content-Security-Policy: script-src 'self' 'nonce-abc123';

<script nonce="abc123">
    console.log("This inline script is allowed!");
</script>
```

Since the CSP script-src nonce and the script's nonce match, the browser allows the script to execute.

- **report-uri:**
- CSP will make a report for you about what it has blocked and other
helpful information. This URI is to tell it where to send the report.
- but this uri is public so beware of ddos

- eg: Content-Security-Policy: default-src 'self'; style-src https://*.jquery.com;
script-src https://*.google.com;
allows the browser to load styles from jquery.com and scripts from google.com.
---

> note: there are 2 types of headers standard and non standard.
standard headers are defined by official http specification (eg Content-Type,Content-Length)
non standard headers are custom headers specific to applications so all browsers and server may not recognize. but latest browsers will recognize. they start with X or X-
header like x-content-type-options were widely accepted and they brought into standard header. so this created a confusing between standard and non standard thus they rejected standard and non standard concept in 2012.

---

## 2. X-Frame-Options:
- X-Frame-Options is deprecated and Content-Security-Policy
(CSP) is used in its place for modern browsers. X-Frame-Options is used
for backward compatibility of older browsers and will hopefully be phased
out slowly from active use.

---
## 3. X-Content-Type-Options:
- content type header is used by client and server to specify the MIME type of data being transmitted.
- eg contentType: application/json
- contentType: image/png
- This security header only has one possible setting: X-Content-Type-Options: nosniff 
this tells the browser to not sniff the content type and respect the response content-type header value provided 
- ### mime sniffing/content type sniffing:
https://www.youtube.com/watch?v=eq6R6dxRuiU
  - MIME sniffing (or content type sniffing) is when a web browser tries to guess the type of content being served, rather than relying on the Content-Type header provided by the server. 
  - This can be a security risk if a malicious file (e.g., a disguised script) is executed in an unintended way.
  - eg:Imagine a website allows users to upload profile pictures, but it does not strictly enforce the Content-Type header.
  - An attacker uploads a file called malicious.html, but renames it to profile.jpg.
  - The server incorrectly sets Content-Type: image/jpeg or doesnt set at all, but the file actually contains:
<script>alert('Hacked!');</script>
  - If the browser does MIME sniffing, it may detect the actual content as an HTML file and execute the script, leading to an XSS attack.
  - setting the X-Content-Type-Options header to nosniff tells the browser not to guess the file type and strictly follow the Content-Type provided by the server.
  - If a file is served with Content-Type: image/jpeg, the browser will only treat it as an image, even if the content looks like HTML or JavaScript. This prevents malicious files from being executed improperly.
---

## 4. Referrer-Policy:
- When you surf from site to site on the internet, each site sends the next site a value called
the “referrer,” which means a link to the page you came from. This is incredibly helpful for
people analyzing their traffic, so they know where people are coming from.
- However, if you are visiting a site of a sensitive nature (a mortgage application, or a site detailing a specific
medical condition, for example), you may not want the specifics sent to the next page you visit.
- In order to protect your user’s privacy, as the website creator you can set the referrer
value to only pass the domain and not which page the user was on (thisisvoid.in
versus thisisvoid.in/embarrassing-blog-post-title), or to not pass any value at all.
- In order to only pass on the protocol and domain information, set the referrer to ”origin.” No other context can change this setting.
**Referrer-Policy: origin**
- For example, a document at https://thisisvoid.com/page.html will send the referrer
https://thisisvoid.com/.
