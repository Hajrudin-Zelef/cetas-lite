---
id: collect-261001-general-networking/general-networking/manual-how-tos-ipsec-swanctl-rw-ikev2-eap-mschapv2-html-81aa2fcc-1
title: "Add IPv4 route"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "parameters"]
source: docs/RAG/collect-261001-general-networking/manual-how-tos-ipsec-swanctl-rw-ikev2-eap-mschapv2-html-81aa2fcc.md
source_anchor: ""
source_lines: [1, 211]
sha256: 68ed8d41b6cd2106e45b74ee7515159232d5ecd3521508060ac32a4de257b33e
---

# Add IPv4 route

IPsec - Roadwarriors IKEv2
Networks used in this How-To section
| Interface | Network IPv4 | Network IPv6 | 
|---|---|---|
| WAN | 203.0.113.0/24 | 2001:db8:1234::/48 | 
| LAN | 192.168.1.0/24 | 2001:db8:1234:1::/64 | 
| IPsec | 172.16.203.0/24 | 2001:db8:1234:ec::/64 | 
The example users are John and Laura. The example FQDN is vpn1.example.com.
Hint
Any IPv6 functionality is optional. If you do not want to use IPv4+IPv6 dual stack, just skip all IPv6 addresses/networks and focus on IPv4.
Warning
- Do not copy security relevant configuration parameters like passwords into your configuration. Create your own!
- Change all IP addresses, usernames and DNS Records to your own usecase.
Prerequisites
Attention
In all following examples, parameters that should be empty or at default are omitted. Do not change them without a good reason.
External DNS Records
Your OPNsense Firewall has the example IP Subnets 203.0.113.0/24 and 2001:db8:1234::/48. The FQDN can point to any bindable IPv4 and IPv6 address in those subnets. It will be used by clients to connect to the IPsec VPN Server - and by the OPNsense to bind the local listen address.
- Create an A-Record with your external DNS provider, for example vpn1.example.com in A 203.0.113.1
- Create an AAAA-Record, for example vpn1.example.com in AAAA 2001:db8:1234::1
The DNS records must be resolvable from the internet, and they should point to the public IP address of your OPNsense Firewall.
Firewall: Aliases
Create an alias for the IP addresses of your FQDN. That way you can create a combined IPv4/IPv6 rule to allow incoming connections to your IPsec VPN server.
Name:
host_vpn1_example_com
Type:
Host(s)
Content:
203.0.113.1 2001:db8:1234::1
Description:
Host vpn1.example.com
Create an alias for the UDP ports used by IPsec. Port 500 is ISAKMP and port 4500 is IPsec NAT-T.
Name:
port_ipsec_500_4500
Type:
Port(s)
Content:
500 4500
Description:
Ports IPsec 500 and 4500
Firewall: Rules: WAN
Since this roadwarrior configuration will use UDP encapsulation, the ESP packets will be encapsulated inside UDP packets. That is why you do not need a rule to allow the ESP protocol. You only need a firewall rule to allow UDP 500 and UDP 4500. Use the aliases you created in the prior step.
Action
Pass
Interface
WAN
Direction
In
TCP/IP Version
IPv4+IPv6
Protocol
UDP
Source
Any
Source port
Any
Destination
host_vpn1_example_com
Destination port
port_ipsec_500_4500
Description
Allow IPsec UDP ports from ANY source to this firewall
Note
- Now that the Prerequisites have been met, you can choose where to continue:
- Method 1 - Shared IP pool for all roadwarriors
- Method 2 - Static IP address per roadwarrior
Attention
- Do not create both methods at the same time, since authentication between these methods would overlap.
- Only create one connection where you use EAP id: %any (Method 1). If you create multiples of these connections, any roadwarrior can connect to any of them.
EAP-MSCHAPv2
The following roadwarrior configuration is universally usable for many different clients and easy to setup.
EAP-MSCHAPv2 via IKEv2 is based on a server certificate and an EAP Pre-Shared Key (username + password). The CA certificate must be installed on the users device.
Before continuing: Prerequisites
Method 1 - Shared IP pool for all roadwarriors
- Benefit: Easy configuration and works with most clients out of the box.
- Drawback: All configured EAP Identities can authenticate with this connection, so you cannot have tight access control. Roadwarriors do not have unique IP addresses.
Method 2 - Static IP address per roadwarrior
- Benefit: Tight security because every user can be controlled individually with firewall rules.
- Drawback: Configuration needs more time and might not scale with large user counts. Windows native VPN client does not like this configuration since it demands the eap identity exchange method eap id = %any .
Method 2 - Static IP address per roadwarrior
2.1 - VPN: IPsec: Connections: Pools
Create an individual IPv4 pool for each roadwarrior. This configuration will result in 1 usable IPv4 address. The DNS Server(s) will be pushed as Configuration Payload (RFC4306 and RFC7296 3.15). In this example they represent the Unbound Server of the OPNsense.
Name:
pool-roadwarrior-john-ipv4
Network:
172.16.203.1/32
DNS:
192.168.1.1
Name:
pool-roadwarrior-laura-ipv4
Network:
172.16.203.2/32
DNS:
192.168.1.1
Create an individual IPv6 pool for each roadwarrior. This configuration will result in 1 usable IPv6 address.
Name:
pool-roadwarrior-john-ipv6
Network:
2001:db8:1234:ec::1/128
DNS:
2001:db8:1234:1::1
Name:
pool-roadwarrior-laura-ipv6
Network:
2001:db8:1234:ec::2/128
DNS:
2001:db8:1234:1::1
Note
If a roadwarrior has more than one device, you can provide them a larger pool. For example /31 would result in 2 IPv4 addresses, and /127 in 2 IPv6 addresses. You will have to keep track of this yourself though, do not configure pools that overlap.
Note
You can skip the DNS field if you do not want to push DNS Servers to your clients.
2.2 - VPN: IPsec: Pre-Shared Keys
Create EAP Pre-Shared Keys. The local identifier is the username, and the Pre-Shared Key is the password for the VPN connection.
Local Identifier:
john@vpn1.example.com
Remote Identifier:
vpn1.example.com (optional, needed for native android client, if it causes issues remove it)
Pre-Shared Key:
48o72g3h4ro8123g8r
Type:
EAP
Local Identifier:
laura@vpn1.example.com
Remote Identifier:
vpn1.example.com
Pre-Shared Key:
LIUAHSDq2nak!12
Type:
EAP
Note
Instead of john@vpn1.example.com you can use any string as local identifier, for example only john. If you have multiple VPN servers, the FQDN makes it easier to know which one the user is assigned to.
2.3 - VPN: IPsec: Connections
- Enable IPsec with the checkbox at the bottom right and apply.
2.3.1 Create connection for john@vpn1.example.com:
- Press + to add a new Connection, enable advanced mode with the toggle.
General Settings:
Proposals:
aes256-sha256-modp2048 (Disable default!)
Version:
IKEv2
Local addresses:
vpn1.example.com
UDP encapsulation:
X
Rekey time:
2400 for most clients - Or 86400 when using Windows native VPN client
DPD delay:
30
Pools:
pool-roadwarrior-john-ipv4 pool-roadwarrior-john-ipv6
Keyingtries:
0
Description:
roadwarrior-john-eap-mschapv2-p1
Save to reveal the next options:
Local Authentication:
Round:
0
Authentication:
Public Key
Id:
vpn1.example.com
Certificates:
vpn1.example.com
Description:
local-vpn1.example.com
Remote Authentication:
Round:
0
Authentication:
EAP-MSCHAPv2
EAP Id:
john@vpn1.example.com
Description:
remote-john-eap-mschapv2
Children:
Press + to add a new Child, enable advanced mode with the toggle.
Start action:
None
ESP proposals:
aes256-sha256-modp2048 (Disable default!)
Local:
0.0.0.0/0 ::/0
Rekey time (s):
600 for most clients - Or 0 when using Windows native VPN client
Description:
roadwarrior-john-eap-mschapv2-p2
Save and Apply the configuration.
Note
With children you select the networks your roadwarrior should be able to access. In a split tunnel scenario, you would specify the example LAN nets 192.168.1.0/24 and  2001:db8:1234:1::/64 as local traffic selectors. In a full tunnel scenario (all traffic forced through the tunnel) you would specify 0.0.0.0/0 and ::/0 as local traffic selectors. The following example child will use the full tunnel method. A full tunnel is generally more secure - especially with IPv6 involved - since no traffic can leak.
2.3.2 Create connection for laura@vpn1.example.com:
- Press + to add a new Connection, enable advanced mode with the toggle. You could also clone the connection you already configured.
General Settings:
Proposals:
aes256-sha256-modp2048 (Disable default!)
Version:
IKEv2
Local addresses:
vpn1.example.com
UDP encapsulation:
X
Rekey time:
2400 for most clients - Or 86400 when using Windows native VPN client
DPD delay:
30
Pools:
pool-roadwarrior-laura-ipv4 pool-roadwarrior-laura-ipv6
Keyingtries:
0
Description:
roadwarrior-laura-eap-mschapv2-p1
