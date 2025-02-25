## security headers

https://securityheaders.com/
- Security headers can usually be applied on either the web server or your code. To apply them you generally will add one line of code or check a box to configure your web server
#### lets discuss some security headers:
- ##### Content-Security-Policy (CSP)
  goated video: https://www.youtube.com/watch?v=J90t0h0AP1U
  **before we get into csp we will look into xss (cross site scripting attack):**
  -  one among the mitigation for xss attack
  - suppose you have a html template as output and a input form that will go to the backend and put the result back to the template.
  - if you enter a js script in the input form and click submit, the output html template will have the script tag in it thus it will run the js code.
  ```html
  <p>{{input}}</p>
  ```
  after inputint <script>alert(1)</script> the template will becode
  ```html
  <p><script>alert(1)</script> </p>
  ```
 which will get executed in the browser.
  - one way to prevent is validating the input and escaping the output before parsing the html template(by adding escape characters) 
  - this will work well with html but escaping the output wont work in case of html attributed or suppose we have templating in css and js
  eg:
  ```html
    <p><a href="{{input}}">link</a></p>
    ```
  - we give a url style js as input  
  ```html
    <p><a href="javascript:alert(1)">link</a></p>
    ```
    this holds true for even script tag which takes input as template like onclick ={{input}}
  - despite escaping when we click the link the js file loads so xss attack happens
- **csp** is meant and enforced by the browser. its saying browser that just dont render every payloads and just render payload which i allow you to render and block payloads which i dont allow.
- it is a set via **http header**, by the **web server** for the browser to enforce.
- **default-src: self**. by giving this header we say that we will block everything which arent coming from our source we server. this will break the ui or anything that renders from someother place like a google font, third party css,some ext js which takes client side metrics etc. all those things need to be carefully checked with the developer team and then they should be allowed
 > you can generate csp header through this website https://report-uri.com/home/generate

- **eg of csp policy:**
**Content-Security-Policy: default-src 'self'; script-src 'self'; font-src 'none'**
1. **default-src 'self'**
- This is a fallback directive that applies to all types of resources unless another more specific directive (like script-src, img-src, style-src, etc.) is defined.
- The value 'self' means that resources can only be loaded from the same origin (i.e., the domain serving the page).
- This applies to resources such as:
Images (img-src)
Scripts (script-src)
Styles (style-src)
Fonts (font-src)
Frames (frame-src)
Media files like videos and audio (media-src)
AJAX requests / Fetch / XHR (connect-src)
Web Workers (worker-src)
Any other types of resources not explicitly defined in other directives.
Example of Behavior:
- If an image <img src="https://example.com/logo.png"> is loaded from a different domain (e.g., example.com while your site is mysite.com), the image will not load.
- If a CSS file <link rel="stylesheet" href="https://cdn.example.com/style.css"> is linked from an external domain, it will not be allowed unless style-src is defined separately.
- However, if you have a more specific directive, such as script-src 'self', then scripts will override script-src instead of default-src.
2. **script-src 'self'**
- This directive overrides default-src for JavaScript files (.js).
- The value 'self' means that scripts must be loaded only from the same origin.
- If script-src is defined, it completely replaces default-src for scripts.
Example of Behavior:
Allowed:
<script src="/static/main.js"></script> <!-- Allowed if hosted on the same domain -->
Blocked:
<script src="https://cdn.example.com/main.js"></script> <!-- Blocked -->
Blocked:
<script>console.log("Inline script");</script> <!-- Blocked (because 'unsafe-inline' is not allowed) -->
- To allow inline scripts, you’d need 'unsafe-inline' (which is not recommended for security reasons).

**Why Override default-src?**
- Even though default-src 'self' already applies to scripts, explicitly defining script-src 'self' gives more granular control.
- If default-src 'self' were used alone, it would allow all resource types except those explicitly overridden.
- By specifying script-src, you ensure that scripts follow their own specific rule, independent of other resource types.
3. **font-src 'none'**
- This directive blocks the loading of all fonts.
- The value 'none' means no fonts are allowed, not even from the same domain.
Fonts include:
- Web fonts (.woff, .woff2, .ttf, .otf) loaded via CSS @font-face.
- Font files referenced in <link> elements.
- Any fonts fetched from CDNs or external sources.
Example of Behavior:
```css
@font-face {
  font-family: "MyFont";
  src: url("/fonts/myfont.woff2"); /* Blocked, even though it's hosted locally */
}
```
Blocked:
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto">

**Will font-src 'none' Break the UI?**
- If the website relies on web fonts, the UI will break (or look significantly different).
- If the website uses system fonts (e.g., Arial, Times New Roman, sans-serif), it will fall back to system defaults, which may look different but will not necessarily break.
- thats the reason why you should always consult a developer before using csp

- there are lot more tags which will give granular control to csp.
