# Proxy:
- A proxy in computing acts as an intermediary between a client and another server. 
- It allows indirect access to services, enabling features like enhanced security, privacy, or performance.
- There are various types of proxies, each serving different purposes, but the core function is to forward requests and responses between a client and the target server.

# Types of Proxies and How They Work

### **1. Forward Proxy**
A forward proxy acts as an intermediary for a client to access resources on a server.

#### **Steps**:
1. **Client Request**:  
   The client sends a request to the forward proxy instead of the destination server (e.g., requesting a webpage).  
2. **Proxy Processes the Request**:  
   The proxy server examines the request, logs it, and applies rules or filters (e.g., checking if access to the requested site is allowed).  
3. **Forwarding the Request**:  
   The proxy forwards the client’s request to the destination server.  
4. **Server Response**:  
   The destination server processes the request and sends the response back to the proxy.  
5. **Returning the Response**:  
   The proxy receives the server’s response, optionally caches it, and forwards it to the client.  

### **Use Case**:  
- **Bypassing Restrictions**: Accessing blocked websites in restricted regions.

---

### **2. Reverse Proxy**
A reverse proxy is positioned in front of servers and handles client requests on their behalf.

#### **Steps**:
1. **Client Request**:  
   A client sends a request to a website, which is routed to the reverse proxy instead of the server.  
2. **Proxy Receives the Request**:  
   The reverse proxy determines which backend server to forward the request to (e.g., based on load-balancing rules).  
3. **Forwarding to the Server**:  
   The proxy forwards the client’s request to the chosen backend server.  
4. **Server Response**:  
   The backend server processes the request and sends the response back to the proxy.  
5. **Returning the Response**:  
   The reverse proxy optionally performs caching or encryption and forwards the response to the client.  

### **Use Case**:  
- **Load Balancing**: Distributing traffic across multiple servers to ensure availability.

##### Difference Between Forward Proxy and Reverse Proxy

| **Aspect**            | **Forward Proxy**                                               | **Reverse Proxy**                                               |
|------------------------|-----------------------------------------------------------------|-----------------------------------------------------------------|
| **Purpose**           | Used by clients to access servers.                              | Used by servers to manage client requests.                     |
| **Position**          | Sits between the client and the destination server.             | Sits between the client and the backend servers.               |
| **Direction of Proxying** | Proxies client requests to a server.                          | Proxies client requests to backend servers.                    |
| **Client Awareness**  | The client is aware of the forward proxy and sends requests to it directly. | The client is not aware of the reverse proxy; it appears as the destination server. |
| **Hides Information** | Hides the client’s identity (e.g., IP address).                 | Hides the backend server details and architecture.             |
| **Caching**           | Can cache client-requested data to improve response time.       | Can cache server responses to improve scalability and performance. |
| **Use Cases**         | - Bypassing geographic restrictions. <br> - Anonymizing browsing. | - Load balancing.<br> - Enhancing security.<br> - Caching server responses. |

---

---

## **3. Transparent Proxy**
A transparent proxy intercepts communication between the client and server without modifying requests or responses.

### **Steps**:
1. **Interception**:  
   The client’s request is intercepted by the transparent proxy without the client explicitly knowing.  
2. **Filtering or Caching**:  
   The proxy applies filters (e.g., blocking URLs) or checks if the requested content is cached.  
3. **Forwarding to Server**:  
   If allowed, the proxy forwards the request to the destination server.  
4. **Server Response**:  
   The destination server sends the response back to the proxy, which forwards it to the client.  

### **Use Case**:  
- **Public Networks**: Monitoring and restricting internet usage in schools or libraries.

---

## **4. Anonymous Proxy**
An anonymous proxy hides the client’s IP address when making requests.

### **Steps**:
1. **Client Request**:  
   The client sends a request to the proxy server instead of the destination server.  
2. **Anonymization**:  
   The proxy replaces the client’s IP address with its own before forwarding the request.  
3. **Forwarding to Server**:  
   The proxy forwards the anonymized request to the destination server.  
4. **Server Response**:  
   The server sends the response back to the proxy.  
5. **Returning the Response**:  
   The proxy relays the response to the client.  

### **Use Case**:  
- **Privacy**: Hiding the client’s IP to access websites anonymously.

---

## **5. High-Anonymity Proxy**
A high-anonymity proxy hides both the client’s IP and the fact that a proxy is being used.

### **Steps**:
1. **Client Request**:  
   The client sends a request to the proxy.  
2. **Complete Anonymization**:  
   The proxy removes identifying headers and avoids including any proxy-related information.  
3. **Forwarding to Server**:  
   The proxy forwards the request to the destination server as if it originated directly.  
4. **Server Response**:  
   The server processes the request and sends the response to the proxy.  
5. **Returning the Response**:  
   The proxy relays the response to the client.  

### **Use Case**:  
- **Secure Browsing**: Ensuring complete anonymity during sensitive activities.

---

## **6. Web Proxy**
A web proxy provides a browser interface to access blocked or restricted websites.

### **Steps**:
1. **Access the Web Proxy**:  
   The user navigates to the web proxy website and enters the destination URL.  
2. **Proxy Fetches the Page**:  
   The web proxy sends a request to the destination website.  
3. **Server Response**:  
   The destination server sends the response back to the proxy.  
4. **Display to the User**:  
   The proxy relays the website content to the user’s browser.  

### **Use Case**:  
- **Bypassing Firewalls**: Accessing restricted websites through public web proxies.

---

## **7. SOCKS Proxy**
A SOCKS proxy is protocol-agnostic and operates at the transport layer.

### **Steps**:
1. **Client Connection**:  
   The client establishes a connection to the SOCKS proxy.  
2. **Authentication (Optional)**:  
   If required, the proxy authenticates the client.  
3. **Traffic Forwarding**:  
   The proxy forwards the client’s traffic (e.g., HTTP, FTP) to the destination server.  
4. **Server Response**:  
   The server sends the response back to the SOCKS proxy.  
5. **Returning the Response**:  
   The proxy relays the server’s response to the client.  

### **Use Case**:  
- **Gaming**: Reducing latency in online games.

---

## **8. Caching Proxy**
A caching proxy stores frequently accessed resources to improve speed.

### **Steps**:
1. **Client Request**:  
   The client sends a request to the proxy.  
2. **Cache Check**:  
   The proxy checks if the resource is cached.  
3. **Serve from Cache**:  
   If cached, the proxy serves the content directly to the client.  
4. **Forwarding Request (if not cached)**:  
   If not cached, the proxy forwards the request to the server.  
5. **Server Response**:  
   The server sends the response, which the proxy caches for future use.  
6. **Return Response**:  
   The proxy sends the response to the client.  

### **Use Case**:  
- **Performance**: Speeding up browsing by caching popular resources.

---

## **9. Content Filtering Proxy**
A content filtering proxy blocks or restricts specific content based on rules.

### **Steps**:
1. **Client Request**:  
   The client sends a request to the proxy.  
2. **Filter Application**:  
   The proxy checks the request against filtering rules (e.g., blocking specific URLs).  
3. **Decision**:  
   - If allowed, the proxy forwards the request to the destination server.  
   - If blocked, the proxy denies the request and optionally displays a warning.  
4. **Server Response**:  
   Allowed requests are processed by the server and sent back to the proxy.  
5. **Return Response**:  
   The proxy relays the response to the client.  

### **Use Case**:  
- **Parental Controls**: Restricting access to inappropriate content.
