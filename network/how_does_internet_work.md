# how does internet work

https://developer.mozilla.org/en-US/docs/Learn_web_development/Howto/Web_mechanics/How_does_the_Internet_work

## difference between switch and router
| Feature             | **Switch**                                        | **Router**                                                      |
| ------------------- | ------------------------------------------------- | --------------------------------------------------------------- |
| **Function**        | Connects devices **within** a local network (LAN) | Connects **multiple networks**, including LAN to WAN (internet) |
| **Works on Layer**  | Layer 2 (Data Link Layer – MAC addresses)         | Layer 3 (Network Layer – IP addresses)                          |
| **Address Used**    | MAC address                                       | IP address                                                      |
| **Primary Use**     | Forwards data between devices in the same network | Routes data between **different networks**                      |
| **NAT Support**     | No                                                | Yes (performs Network Address Translation)                      |
| **DHCP**            | No (usually)                                      | Yes (often assigns IPs to devices)                              |
| **Internet Access** | Cannot provide internet access                    | Can route to the internet                                       |

-Switch is for local device-to-device communication (e.g., PC to printer).
- Router is for device-to-internet communication and inter-network routing.
## When You Type google.com, What the Router and Switch Do
Assume you're connected over Ethernet (wired) and both router and switch exist. If you're on Wi-Fi, the switch is skipped.

### Switch (Layer 2 – MAC based)
**What it does:**
- Receives the packet from your computer.
- Reads the destination MAC address.
- Looks up its MAC address table to decide which port to forward it to (likely the router's port).
- Forwards the packet within your LAN to the router.
- The switch just forwards frames inside your local network.
#### How does the switch know the MAC address of the router?
- Switches build their MAC address table dynamically.
- When any device (including the router) sends a frame, the switch records the source MAC address and which port it came from.
- So when the router sends any packet (e.g., DHCP reply, DNS response), the switch learns:
- MAC A is reachable via port X.
- Later, when a device sends a packet to that MAC, the switch knows which port to forward it to.
- If the switch doesn’t yet know the destination MAC, it floods the frame out to all ports (except the one it came from) and waits to learn.

### Router (Layer 3 – IP based)
**What it does:**

- Receives the IP packet from the switch.
- Sees the destination is an external IP (e.g., Google’s).
- NAT: Translates your private IP (e.g., 192.168.1.5) to the router's public IP.
- Routes the packet to the modem/ISP via its WAN interface.
- Keeps track of this flow in its NAT table so it can map responses back to your machine.

High-level Flow
You type google.com.

DNS query sent → via switch → to router → to DNS server (via modem and ISP).

DNS response returns with Google’s IP.

You now send an IP packet to Google’s IP.

Switch forwards frame to router using MAC table.

Router:

Sees the destination is external,

Applies NAT,

Sends it to the modem, then ISP, then internet.