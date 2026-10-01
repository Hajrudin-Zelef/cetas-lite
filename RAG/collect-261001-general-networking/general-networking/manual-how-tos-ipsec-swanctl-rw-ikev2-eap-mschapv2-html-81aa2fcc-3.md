---
id: collect-261001-general-networking/general-networking/manual-how-tos-ipsec-swanctl-rw-ikev2-eap-mschapv2-html-81aa2fcc-3
title: "Add IPv4 route"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft", "Nvidia"]
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/manual-how-tos-ipsec-swanctl-rw-ikev2-eap-mschapv2-html-81aa2fcc.md
source_anchor: ""
source_lines: [415, 612]
sha256: f41b2127164ad4316d31e7644065e3fba3777985479a540fb2fd9d9c908dd895
---

# Add IPv4 route

CertReqWithData=0
IpAddrAssign=0
IPAddress=
SubnetMask=
DNS1=
DNS2=
DomainName=
DomainInTunnel=
SubjectCert=
IssuerCert=
FingerPrint=
UseSHA1=0
Firewall=0
OnlyTunnel=0
RasOnlyTunnel=0
DNSActiv=1
DNS1Tmp=
DNS2Tmp=
[IKEV2POLICY1]
Ikev2Name=aes256-sha256
Ikev2Crypt=6
Ikev2PRF=5
Ikev2IntAlgo=12
[IPSECPOLICY1]
IPSecName=aes256-sha256
IpsecCrypt=6
IpsecAuth=5
- For other users edit IkeIdStr=john@vpn1.example.com . ChangeName=vpn1.example.com andGateway=vpn1.example.com to your vpn gateway.
- Import the example.ini Profile: 
  - Launch the NCP Secure Entry Client.
  - Navigate to the Profile menu.
  - Select the option to Import Profile.
  - Browse to the location where your example.ini profile is saved.
  - Select the profile and click Open or Import (whichever option appears).
  - You can enter the username and password of the user when importing the profile. 
    - Username: john@vpn1.example.com
    - Password: 48o72g3h4ro8123g8r
- Import the self-signed CA certificate into the NCP certificate store. Go to C:\ProgramData\NCP\SecureClient\cacerts and copy your the .pem file in there.
- The profile should now be loaded into the NCP Secure Entry Client. You can start it and it should connect. If not, check the Logfile in “Help” for the error message.
Note
There is also a version for macOS, which works with the same configuration as above. The only challenge is finding the right folder for the cacerts. You can find it by going into the terminal and using the command sudo find / -name cacerts. Then you can pinpoint the path and copy the CA certificates there.
Postrequisites
Firewall rules, Source NAT and DNS
Now that you have configured split or full tunnel mode, you need rules to allow the traffic into your LAN and to the WAN (Internet). For IPv4 connection to the WAN (Internet) you need a Source NAT rule for IP-Masquerading. If you want the OPNsense to handle DNS, you can to configure Unbound so your roadwarriors use it as DNS server to prevent DNS leaks.
Tip
If you have internal IPv4 services (like a mailserver) that have external IPs in their DNS A-Records, you should configure Reflection NAT. There is a tutorial in the How-To section of Network Address Translation. If you follow it, add the ipsec interface in the Destination NAT (Port Forward) rules you create.
Firewall: Aliases
Create the following aliases:
Name:
InternetIPv4
Type:
Network(s)
Content:
10.0.0.0/8  172.16.0.0/12  192.168.0.0/16  127.0.0.0/8
Description:
Internet IPv4 - use inverted
Note
The InternetIPv6 alias needs to be your own IPv6 network.
Name:
InternetIPv6
Type:
Network(s)
Content:
2001:db8:1234::/48
Description:
Internet IPv6 - use inverted
Name:
net_pool_roadwarrior
Type:
Network(s)
Content:
172.16.203.0/24  2001:db8:1234:ec::/64
Description:
Network pool-roadwarrior-ipv4 and ipv6
Additionally, if you created separate IP pools for individual roadwarriors (Method 2), create the following aliases so you are able to create individual firewall rules per roadwarrior:
Name:
host_pool_roadwarrior_john
Type:
Host(s)
Content:
172.16.203.1/32  2001:db8:1234:ec::1/128
Description:
john@vpn1.example.com
Name:
host_pool_roadwarrior_laura
Type:
Host(s)
Content:
172.16.203.2/32  2001:db8:1234:ec::2/128
Description:
laura@vpn1.example.com
Firewall: Rules: IPsec
Here you use the aliases you created in the prior step in order to create firewall rules on the IPsec interface in order to allow traffic from the roadwarrior networks to your LAN and to the WAN (Internet).
As first rule it is a good idea to allow ICMP for troubleshooting purposes. With that rule, roadwarriors can ping the OPNsense firewall. Please note that they can only ping those IPs that are included in the local traffic selectors of the children.
Action
Pass
Interface
IPsec
Direction
In
TCP/IP Version
IPv4+IPv6
Protocol
ICMP
Source
Any
Source port
Any
Destination
This Firewall
Destination port
Any
Description
Allow ICMP to this firewall
As second rule, you should allow LAN access from the IPsec roadwarrior networks. If you created individual aliases, you can create multiples of those rules with the aliases of the individuals added instead of the whole network.
- Example for a rule that allows the whole IPsec roadwarrior network to the LAN. LAN net is a predefined alias if you have an interface called LAN:Action Pass Interface IPsec Direction In TCP/IP Version IPv4+IPv6 Protocol TCP/UDP Source net_pool_roadwarriorSource port Any Destination LAN netDestination port Any Description Allow ICMP to this firewall
- Example for an individual allow rule to the LAN: Action Pass Interface IPsec Direction In TCP/IP Version IPv4+IPv6 Protocol TCP/UDP Source host_pool_roadwarrior_johnSource port Any Destination LAN netDestination port Any Description Allow john@vpn1.example.com access to LAN net
The last matching rules can allow Internet access if you have configured a full tunnel. Just as the example above, you can also create individual rules to restrict Internet access to some roadwarriors:
Action
Pass
Interface
IPsec
Direction
In
TCP/IP Version
IPv4
Protocol
Any
Source
net_pool_roadwarrior
Source port
Any
Destination / Invert
X
Destination
InternetIPv4
Destination port
Any
Description
Allow Internet Access IPv4
Action
Pass
Interface
IPsec
Direction
In
TCP/IP Version
IPv6
Protocol
Any
Source
net_pool_roadwarrior
Source port
Any
Destination / Invert
X
Destination
InternetIPv6
Destination port
Any
Description
Allow Internet Access IPv6
Note
By setting Destination / Invert you invert the match of the alias. Do not use “Any” as Destination to the Internet, since it also includes all networks that are locally attached to your firewall.
Firewall: NAT: Source NAT
For IPv4 Internet access to work, you need to set up a Source NAT rule for IP-Masquerading. Start by enabling at least Hybrid Source NAT rule generation and Save. Otherwise you cannot add your new manual NAT rule.
Interface
WAN
Direction
In
TCP/IP Version
IPv4
Protocol
any
Source
net_pool_roadwarrior
Source port
any
Destination
any
Destination port
any
Translation / target
WAN address
Description
IPsec MASQ
Services: Unbound DNS
Note
If you do not serve internal DNS records (Split DNS) or do not use an Active Directory you can skip the DNS configuration.
For full control over DNS, you should either use Unbound on the OPNsense or the DNS servers in your own network. If you provide your roadwarriors with external DNS servers (like 8.8.8.8), they cannot resolve your internal resources and will send those requests to external DNS servers, thus exposing your internal DNS records. (DNS Leak)
Attention
If you created a full tunnel for IPv4 only (0.0.0.0/0 without ::/0), and your roadwarriors are in IPv4+IPv6 dual stack networks, their devices will prefer the link local IPv6 DNS servers provided by SLAAC or DHCPv6 over your IPv4 VPN DNS server.
Enable Unbound and leave the Network Interfaces on All (recommended). Next go to Query Forwarding and input your Custom forwarding servers. For example your Samba or Microsoft Active Directory Domain Controllers.
Unbound listens on port 53 UDP/TCP on all network interfaces of the OPNsense. If you followed all prior steps, access to your LAN is already permitted from the IPsec Network. You can use the IP addresses of the OPNsense in that network as target for the DNS queries.
In this example they are: 192.168.1.1 and 2001:db8:1234:1::1.
Troubleshooting
If the VPN connection does not establish right away there are several steps you can take to troubleshoot the connection. Here is a short summary where to start. Debugging an IPsec connection takes time, do not get discouraged if you can not solve the problem right away.
- If it is your first IPsec connection, do not forget to enable IPsec and apply.
