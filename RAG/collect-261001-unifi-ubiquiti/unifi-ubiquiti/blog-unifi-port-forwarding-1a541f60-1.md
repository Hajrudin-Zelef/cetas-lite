---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-unifi-port-forwarding-1a541f60-1
title: "blog-unifi-port-forwarding-1a541f60"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "latency", "parameters", "throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-unifi-port-forwarding-1a541f60.md
source_anchor: ""
source_lines: [1, 84]
sha256: a8f01720fd493eed2f329fd938513011ab7c410a686aedf7baddc097c1d5ef90
---

# blog-unifi-port-forwarding-1a541f60

How to set up Port Forwarding in UniFi
In this guide I'll show you how to setup port forwarding on a UniFi Coud Gateway, you can use this with any other UniFi gateway. The version of the network application running on this UCG is 10.3.58.
1. Prerequisites and ISP Limitations
Before configuring inbound NAT rules (Port Forwarding), it is imperative to verify the WAN connection properties.
- 
Static IP / DHCP Reservation: The internal target device must have a static IP address or a configured DHCP reservation in the UniFi Network Application. If the device dynamically changes its IP, the port forwarding rule will fail.
- 
Public IPv4 Address: Port forwarding requires a routable public IPv4 address directly assigned to the UniFi Gateway's WAN interface. (Note: IPv6 operates differently, relying on firewall rules to permit traffic to globally routable addresses rather than NAT/port forwarding).
- 
Carrier-Grade NAT (CGNAT): If the ISP uses CGNAT, inbound connections will fail because the WAN IP is shared among multiple ISP customers. Port forwarding cannot bypass CGNAT.
- 
Common CGNAT/Private Ranges: 100.64.0.0/10 (specifically reserved for CGNAT),10.0.0.0/8 ,172.16.0.0/12 , and192.168.0.0/16 .
- 
Verification Step: To verify if you are behind CGNAT, check your WAN IP in the UniFi Network Application (Settings > Internet). Then, from a device on your network, visit a site like ifconfig.me orwhatismyip.com . If the external IP shown on the website does not match the WAN IP shown in your UniFi console, your ISP is using CGNAT or upstream routing.
- 
Double NAT: If the UniFi Gateway is downstream of an ISP router/modem acting as a router (not in Bridge Mode), inbound traffic will be dropped by the upstream device. Mitigation: This can be resolved if the upstream ISP router supports configuring a Demilitarized Zone (DMZ) pointing to the UniFi Gateway's WAN IP, or by configuring a secondary Port Forwarding rule on the ISP router targeting the UniFi Gateway.
2. Security Assessment: The Risks of Port Forwarding
Opening a port on the WAN interface exposes internal assets directly to the public internet. From a security engineering perspective, this violates the principle of zero trust and expands the external attack surface.
- 
Automated Exploitation: Services exposed on common ports (e.g., SMB/445, HTTP/80, RDP/3389) are actively scanned by botnets and threat actors.
- 
Vulnerability Exposure: A vulnerability in the exposed internal service can lead to immediate remote code execution (RCE) or lateral movement within the LAN.
- 
DDoS Amplification: Unfiltered inbound UDP traffic can be leveraged for DDoS reflection attacks.
- 
UPnP Vulnerabilities: Disable Universal Plug and Play (UPnP) on the UniFi Gateway. UPnP allows internal devices to dynamically open firewall ports without administrator oversight, creating a significant, unmonitored attack surface.
- 
Lack of Network Segmentation: If an exposed device is compromised, attackers can pivot to other devices on the same subnet. Mitigate this by placing exposed devices in an isolated VLAN (DMZ) with strict firewall rules preventing inter-VLAN routing to trusted networks.
Directive: Only proceed with port forwarding if a Reverse Proxy (e.g., NGINX, HAProxy, Cloudflare Tunnels), CDN, or Zero Trust Network Access (ZTNA) solution is technologically infeasible. Reverse proxies are highly recommended for hosting multiple web services, allowing you to expose only port 443 and route traffic internally based on domain names. Never expose administrative interfaces or unencrypted protocols (Telnet, HTTP) directly to the WAN.
3. VPN: The Enterprise Alternative to Port Forwarding
Rather than mapping public ports to internal IP addresses, secure remote access should be brokered through encrypted tunnels. UniFi supports robust VPN protocols that eliminate the need for traditional port forwarding.
- 
WireGuard: Recommended for modern deployments. Offers high throughput, low latency, and uses state-of-the-art cryptography. It silently drops unauthenticated packets, rendering it invisible to port scanners.
- 
Teleport: Ubiquiti's proprietary zero-configuration VPN, ideal for rapid deployment to end-user devices via the WiFiman app. Uses WireGuard under the hood.
- 
OpenVPN: Less performant than WireGuard, still a solid choice to use.
- 
Site-to-Site VPN: Preferred for connecting branch offices to a central hub without exposing resources to the public internet.
Recommendation: Always default to configuring a VPN server on the UniFi Gateway for remote management access instead of port forwarding.
4. Procedure: Configuring Port Forwarding in UniFi Network Application
If port forwarding is absolutely required, execute the following steps within the UniFi Network Application.
- Navigate to the UniFi Network Application.
- Depending on your controller version, navigate to the Port Forwarding section:
- 
Newer Versions: Go to Settings > Policy Engine > Policy Table and click on Create new policy, or go to the Port Forwarding section, and then click on Create new.
- 
Older versions: Go to Settings > Routing > Port Forwarding (or Settings > Security > Port Forwarding).
- Configure the following parameters:
- 
Checkbox: Set to Port Forwarding .
- 
Name: Assign a descriptive identifier (e.g., HTTPS-Web-Server-VLAN20 ).
- 
Interface: Select the active WAN interface (e.g., WAN1 orWAN2 ).
- 
WAN IP Address: Determines which IP address is used to forward the traffic, here you will enter the WAN IP configured on your UniFi firewall for the WAN interface you selected before, most likely will be automatically configured.
- 
WAN Port: Specify the incoming external port. Best Practice: Use a non-standard external port (e.g., 45892 ) to reduce automated botnet scanning noise.
- 
From: Limit the source IP if possible. Change from Any to a specific IP or IP Group to enforce strict access control.
- 
Forward IP Address: Enter the internal, static IP address or DHCP reservation of the target host.
- 
Forwarding Port: Specify the standard internal port the service is actually listening on (e.g., 443 or22 ).
- 
Protocol: Select TCP ,UDP , orBoth . Adhere to the principle of least privilege—do not selectBoth if the service only requires TCP.
- 
Logging: Enable Syslog logging to track and receive alerts whenever traffic hits the rule for SIEM ingestion and anomaly detection.
- 
Click Add or Apply Changes.
Zone-Based Firewall Automation
When you create a port forwarding rule, you do not need to manually create a corresponding firewall rule to permit the traffic. UniFi's Zone-Based Firewall (ZBF) automates this process.
- 
Automatic Rule Generation: The system automatically generates a firewall policy allowing external access to the specific local device and port.
- 
Firewall Matrix Placement: This auto-generated rule appears in the firewall matrix under External to DMZ (or your respective local network zone), rather than under External to Gateway.
Validation & Troubleshooting
- Test the external connection using an external vantage point (e.g., a mobile device on cellular data).
- 
Check Firewall States: Verify the internal host receives the traffic and that the UniFi Gateway firewall states reflect the connection.
- 
Host-Based Firewalls: If the connection fails, verify the target device's local OS firewall (e.g., Windows Defender Firewall, Linux UFW/iptables). Host firewalls frequently drop connections originating from outside their local subnet unless explicitly allowed.
Real-World Threat Showcase: The Cost of Exposure
