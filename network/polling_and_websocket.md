### Short Polling:
**How it works:** The client repeatedly sends requests to the server at regular intervals to check for updates.

**Advantages:**
1. **Simplicity:** Easy to implement.
2. **Compatibility:** Works well with existing HTTP protocols and infrastructure.

**Disadvantages:**
1. **High Latency:** Delays in receiving updates due to fixed intervals.
2. **Inefficient:** Frequent requests can waste bandwidth and server resources, especially when no updates are available.

---

### Long Polling:
**How it works:** The client sends a request to the server, and the server holds the request open until new data is available or a timeout occurs.

**Advantages:**
1. **Reduced Latency:** Faster updates since the server responds as soon as data is available.
2. **Efficiency:** Fewer requests compared to short polling, reducing bandwidth and server load.

**Disadvantages:**
1. **Resource Intensive:** Server resources are tied up while waiting for updates.
2. **Complexity:** Harder to implement and manage compared to short polling.
3. **Timeout Issues:** Requires handling of timeouts and reconnections.

---
- **Short Polling:** Simple but inefficient for real-time updates.
- **Long Polling:** More efficient and faster for real-time updates but harder to implement and manage.

---
### Web socket:
good video: https://www.youtube.com/watch?v=xTR5OflgwgU
good video in go implementation: https://programmingpercy.tech/blog/mastering-websockets-with-go/

- very common in realtime system
- stateful 
- persistent connection between client and the server
- since there is a persistent connection, server can send data anytime it receives back to client
- but in http think of it like a letter, you send a letter to server and the server immideately responds back with another letter as response but http server cant automatically respond when there is no request
- working:
step 1: enables tcp connection between client and server(tcp handshake)
step 2: send a http req
step3: the connection gets updated to websocket
step 4: persistent web socket connection is established
- the drawback of websocket is horizontal scaling. this is because of the stateful nature of the connection. when you horizontally scale ie add new servers the existing state wont transient to the new machine
- difference between **webrtc** and **web socket** is webrtc is udp so some data might be lost but websocket is tcp no data loss
- room 
