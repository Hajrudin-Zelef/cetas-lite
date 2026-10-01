---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9-2
title: "blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark"]
source: docs/RAG/collect-261001-opnsense-pfsense/blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9.md
source_anchor: ""
source_lines: [105, 267]
sha256: ff7edc9247ab3e883717ee08765d04f26fdc39c4055f4a200def3c3e25812c26
---

# blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9

Navigate to Interfaces → Settings.
| Hardware CRC | unchecked | Disable hardware checksum offload | 
VLANs
Switch Choice
A 802.1Q-capable switch with properly configured VLANs is required. Check my router on a stick VLAN configuration guide to see an example setup with a Mikrotik switch.
VLAN Definitions
Typically, the LAN port also carries the VLAN traffic and functions as
trunk port. For me,
the default is the igb0 port. I chose it as the parent interface for all VLANs
in the following steps.
Navigate to Interfaces → Other Types → VLAN and add the VLANs.
Management VLAN
| Parent | igb0 | 
| VLAN tag | 10 | 
| Description | VLAN10_MANAGE | 
VPN VLAN
| Parent | igb0 | 
| VLAN tag | 20 | 
| Description | VLAN20_VPN | 
Clear VLAN
| Parent | igb0 | 
| VLAN tag | 30 | 
| Description | VLAN30_CLEAR | 
Guest VLAN
| Parent | igb0 | 
| VLAN tag | 40 | 
| Description | VLAN40_GUEST | 
VLAN Interfaces
We add an interface for each VLAN. Navigate to Interfaces → Assignments.
- Select vlan 10 , enter the descriptionVLAN10_MANAGE , and click+
- Select vlan 20 , enter the descriptionVLAN20_VPN , and click+
- Select vlan 30 , enter the descriptionVLAN30_CLEAR , and click+
- Select vlan 40 , enter the descriptionVLAN40_GUEST , and click+
Click Save.
VLAN Interface IPs
To easier remember which IP range belongs to which VLAN, I like the convention of matching the third octet of the IP with the VLAN ID. I.e., assigning the VLAN with the ID 10 the address 192.168.10.0/24.
Interface: VLAN10_MANAGE
Select the VLAN10_MANAGE interface.
| Enable Interface | checked | 
| IPv4 Configuration Type | Static IPv4 | 
| IPv4 Address | 192.168.10.1/24 | 
Click Save.
Interface: VLAN20_VPN
| Enable Interface | checked | 
| IPv4 Configuration Type | Static IPv4 | 
| IPv4 Address | 192.168.20.1/24 | 
Interface: VLAN30_CLEAR
| Enable Interface | checked | 
| IPv4 Configuration Type | Static IPv4 | 
| IPv4 Address | 192.168.30.1/24 | 
Interface: VLAN40_GUEST
| Enable Interface | checked | 
| IPv4 Configuration Type | Static IPv4 | 
| IPv4 Address | 192.168.40.1/24 | 
VLAN Interface DHCP
We need to configure DHCP for each VLAN we created. I use x.x.x.100-199 for
dynamic and x.x.x.10.10-99 for static IP address assignments. You might want
to amend these ranges to your requirements.
Navigate to Services → DHCPv4.
DHCP: VLAN10_MANAGE
Select VLAN10_MANAGE.
| Enable | checked | 
| Range | from 192.168.10.100 to192.168.10.199 | 
Click Save.
DHCP: VLAN20_VPN
| Enable | checked | 
| Range | from 192.168.20.100 to192.168.20.199 | 
DHCP: VLAN30_CLEAR
| Enable | checked | 
| Range | from 192.168.30.100 to192.168.30.199 | 
DHCP: VLAN40_GUEST
| Enable | checked | 
| Range | from 192.168.40.100 to192.168.40.199 | 
| DNS servers | 1.1.1.11.0.0.1 | 
DHCP: LAN
| Range | from 192.168.1.100 to192.168.1.199 | 
WireGuard VPN with Mullvad
In recent years, Mullvad has been my VPN provider of choice. When That One Privacy Site was still a thing, Mullvad was one of the top recommendations there. After reading the review, I decided to try it out and haven’t looked back since. No personally identifiable information is required to register, and paying cash via mail works perfectly.
I decided to go with WireGuard because I’m fine riding the bleeding edge. 😎 For more detailed steps, check the official OPNsense documentation on setting up WireGuard with Mullvad and WireGuard selective routing.
Please note that the FreeBSD kernel does not (yet) natively support WireGuard, so you must install it as a plugin. Possibly, this doesn’t meet your stability, security, or performance requirements.
Navigate to System → Firmware → Plugins and install
os-wireguard. Refresh the browser and navigate to
VPN → WireGuard.
Remote Peers
Select your preferred WireGuard servers from the Mullvad’s server list and take note of their names and public keys. It’s worth spending some time to benchmark server performance before making a choice.
Select the Endpoints tab and click Add. Here is the configuration for
the remote ch5-wireguard Mullvad endpoint.
| Name | mullvad-ch5-wireguard | 
| Public Key | /iivwlyqWqxQ0BVWmJRhcXIFdJeo0WbHQ/hZwuXaN3g= | 
| Allowed IPs | 0.0.0.0/0 | 
| Endpoint Address | 193.32.127.66 | 
| Endpoint Port | 51820 | 
| Keepalive | 25 | 
To mitigate risks against DNS poisoning, resolve the server’s hostname and enter
its IP as Endpoint Address. You can do this by running
nslookup ch5-wireguard.mullvad.net in a shell. Make sure to not confuse this
address with the SOCKS5 Proxy Address from Mullvad’s server list!
Repeat the steps above to add another server, e.g., ch6-wireguard. Note that
all endpoint configurations use the Endpoint Port 51820.
Local Peers
Select the Local tab, click Add, and enable the advanced mode.
| Name | mullvad0 | 
| Listen Port | 51820 | 
| Tunnel Address | <LEAVE EMPTY> | 
| Peers | ch5-wireguard | 
| Disable Routes | checked | 
| Gateway | <LEAVE EMPTY> | 
Click Save to generate the WireGuard key pair. Click Edit and copy the
generated Public Key.
Next, run the following shell command to get a Mullvad access token:
Then run the following command to create a Mullvad Device with DNS hijacking disabled:
I cover the snippet above and Mullvad’s DNS hijacking in another post: Use Custom DNS Servers With Mullvad And Any WireGuard Client.
Copy the IPv4 IP address to the Tunnel Address field of the Wireguard local
peer. Subtract one from the Tunnel Address and enter the result as
Gateway IP. E.g., 10.105.248.50 for the example above. It’s just a
convention I like, but you can use any arbitrary, unused
private RFC1918 IP.
Repeat the steps above to create a second local peer named mullvad1. Remember
to use a different Listen Port (e.g., 51821).
When you finish, select the General tab. Check Enable WireGuard. You
should see a handshake for the wg0 and wg1 tunnels on the Handshakes
tab.
WireGuard Interfaces
Navigate to Interfaces → Assignments.
- Select wg0 , add the descriptionWAN_VPN0 , and click+
- Select wg1 , add the descriptionWAN_VPN1 , and click+
Enable the newly created interfaces and restart the WireGuard service after. It ensures the interfaces get an IP address from WireGuard.
VPN Gateways
Navigate to System → Gateways → Single and add the VPN gateways.
WAN_VPN0
| Name | WAN_VPN0 | 
| Interface | WAN_VPN0 | 
| Address Family | IPv4 | 
| IP Address | 10.105.248.50 | 
| Far Gateway | checked | 
| Disable Gateway Monitoring | unchecked | 
| Monitor IP | 100.64.0.1 | 
WAN_VPN1
| Name | WAN_VPN1 | 
| Interface | WAN_VPN1 | 
| Address Family | IPv4 | 
| IP Address | 10.109.231.89 | 
| Far Gateway | checked | 
| Disable Gateway Monitoring | unchecked | 
| Monitor IP | 100.64.0.2 | 
Monitoring IPs
Each VPN gateway requires a unique monitoring IP because setting a monitoring IP installs a static route. Optimally, the monitoring IP should be the least possible amount of hops away from the gateway. For Mullvad specifically, we can “abuse” the local infrastructure that’s available through a Mullvad connection. Any of the following IPs are only one hop away from the tunnel exit.
- 100.64.0.1 to100.64.0.3 are
Mullvad’s ad-blocking and tracker-blocking DNS service servers
- 10.64.0.1 is the local Mullvad gateway
You can easily verify the above by running traceroute 100.64.0.1 from a host
connected to Mullvad.
Add Static IPv4 Configuration to the WireGuard Interfaces
OPNsense versions newer than 21.7.3 require adding static IPv4 configuration
to the WireGuard interface. Otherwise, Unbound will use the default route
despite setting the Outgoing Network Interfaces option. Other solutions
exist, but I’m not sure which the “best” or most logical one is. As WireGuard
integration matures, this section hopefully becomes obsolete.
You can find more information regarding this issue on GitHub.
Navigate to Interfaces and edit the WireGuard interfaces.
IP Configuration: WAN_VPN0
| IPv4 Configuration Type | Static IPv4 | 
| IPv4 address | 10.105.248.51/32 | 
