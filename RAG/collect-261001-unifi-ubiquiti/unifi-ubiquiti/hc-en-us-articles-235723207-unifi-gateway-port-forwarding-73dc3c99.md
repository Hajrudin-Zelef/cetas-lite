---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-235723207-unifi-gateway-port-forwarding-73dc3c99
title: "hc-en-us-articles-235723207-unifi-gateway-port-forwarding-73dc3c99"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["cyber"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-235723207-unifi-gateway-port-forwarding-73dc3c99.md
source_anchor: ""
source_lines: [1, 40]
sha256: 02f7a19c85456d86a3d627cdb24d3f892ef96f4733131e058d0b77a1351de93a
---

# hc-en-us-articles-235723207-unifi-gateway-port-forwarding-73dc3c99

UniFi Gateway - Port Forwarding
Port Forwarding allows external devices or services to access specific resources within your UniFi network—such as a web server, security camera, or gaming console—by forwarding incoming traffic from your public IP to a designated internal IP and port.
For highly customizable NAT configuration, see here.
For a full overview of UniFi’s Traffic and Policy Management capabilities, see here.
For a full overview of UniFi's Network and Cyber Security capabilities, see here.
Configuring a Port Forward
- 
Navigate to Policy-Based Routing Rules: Follow the path depending on your UniFi Network version:
  - Network 9.4: Settings > Policy Engine > Port Forwarding
  - Network 9.3: Settings > Policy Table > Create New Policy > Port Forwarding
- Create Name: Assign a name to the rule.
- Select WAN Interface: You can choose to use only one, or all WANs.
- 
Specify Incoming Traffic
  - IP Address and Port: These should correspond to the IP address and port that traffic associated with incoming traffic.
  - From: Choose all incoming traffic, or incoming traffic originating from specified IP addresses.
- Forward IP Address and Port: This should correspond to the device you are forwarding traffic to.
- Specify Protocol: Choose between TCP, UDP, or Both.
- (Optional) Enable Syslog to send logs to an external SIEM or log collector.
Troubleshooting
There are a few common reasons why a Port Forwarding rule may not seem to work as expected.
Your UniFi Gateway does not have a public IP address (Double NAT or CGNAT).
This happens if your UniFi Gateway is located behind another router/modem that uses NAT. You are likely affected by this if your UniFi Gateway has a WAN IP address in one of the following ranges:
- 10.0.0.0/8 (10.0.0.0 - 10.255.255.255)
- 172.16.0.0/12 (172.16.0.0 - 172.31.255.255)
- 192.168.0.0/16 (192.168.0.0 - 192.168.255.255)
- 100.64.0.0/10 (100.64.0.0 - 100.127.255.255)
To fix this issue, try to re-configure your ISP modem/router into bridge mode so that your UniFi Gateway can obtain a public IP address on the WAN interface.
If that is not supported, you will need to first forward the port(s) on the upstream router/modem to the WAN address of your UniFi Gateway in addition to forwarding them from your UniFi Gateway to the desired location. You may wish to contact your ISP to assist with port forwarding or providing a DMZ option that allows you to automatically forward the ports.
Note: When behind NAT and forwarding ports on the upstream router/modem, accessing the upstream router/modem's public IP from a client on the LAN (Hairpin NAT) will not work.
Your UniFi Gateway is already forwarding the port to another device or has UPnP enabled.
A given WAN port can only be forwarded to a single device within your network. For example, TCP port 443 can only be forwarded to one LAN port.
Note: It is possible to forward multiple WAN ports to the same LAN port.
Another possible cause is that UPnP is enabled and is already using the port. Try disabling UPnP in your UniFi Network Application’s Internet Settings.
Incoming traffic is not reaching the WAN interface of your UniFi Gateway.
In this case, the traffic is most likely blocked somewhere upstream, such as at the ISP modem/router, or a third party firewall. We recommend disabling any upstream firewalls for testing, and then contacting your ISP for more details.
The LAN host is blocking the port with a local firewall, or does not have the correct route configured.
In this case, the host/server on the LAN is not allowing outside connections to access the port. On Windows computers, this may be a result of the Windows Firewall rules. On Linux machines, this could be a result of the connection not being allowed in the iptables firewall. We recommend consulting with the particular client’s manufacturer for more information.
There is an incorrect route configured on the LAN host.
It is possible that the LAN host does not know how to reach the IP address of the Internet client. This can result if the default gateway is not configured correctly. You should verify your routing settings on the local host to resolve this situation.
