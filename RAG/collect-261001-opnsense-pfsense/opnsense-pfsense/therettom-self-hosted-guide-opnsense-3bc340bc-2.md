---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/therettom-self-hosted-guide-opnsense-3bc340bc-2
title: "Prune the oldest historical environment to control storage footprint"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-opnsense-pfsense/therettom-self-hosted-guide-opnsense-3bc340bc.md
source_anchor: ""
source_lines: [42, 159]
sha256: afd50e9dbed25a67f0c24755a4ceadf21eec4a4a8e218e2b82b7a54ae4603694
---

# Prune the oldest historical environment to control storage footprint

Once the assignments return you to the main console home layout, you must explicitly bind a static IP network to your new virtual LAN interface so your browser can find it later: 1: Type 2 and hitEnter to select2) Set interface IP address .2: Select the number corresponding to your LAN interface. 3: Configure IPv4 address via DHCP? Typen and hitEnter .4: Enter the new LAN IPv4 address: Type a gateway address (e.g.,192.168.1.1 ) and hitEnter .5: Enter the new LAN IPv4 subnet bit count: Type24 and hitEnter . This just means it creates the subnet for192.168.1.1 -192.168.1.2556: For WAN upstream settings, hit Enter to skip/leave blank. 7: Do you want to enable the DHCP server on LAN? Typey if you want the device to hand out IPs locally over tag 20 immediately, then set your pool ranges when prompted.
- 
If you haven't configured your 802.1Q VLAN capable switch to pass VLAN tags, you just need to do access the device's GUI by typing in its IP address, then configure as referenced in the table below. This is assuming it's a 5-port switch.
| VLAN ID | VLAN NAME | Member Ports | Tagged Ports | Untagged Ports | 
|---|---|---|---|---|
| 1 | Default | 1 | None | 1 | 
| 10 | N/A | 1-2 | 1 | 2 | 
| 20 | N/A | 1, 3-5 | 1 | 3-5 | 
Port 1 is the Trunk Port: It is a tagged member of both VLAN 10 and VLAN 20, meaning this is the physical cable that plugs directly into your single-port OPNsense NUC appliance to pass trunked traffic.
Port 2 is the WAN Access Port: It is an untagged member of VLAN 10. Anything you plug into Port 2 will instantly drop straight into the OPNsense WAN subnet.
Ports 3, 4, and 5 are the LAN Ports: They are untagged members of VLAN 20 (with 4 and 5 also sharing space on the default management network VLAN 1). This is where the local devices will plug in.
In your 802.1Q PVID settings, leave or configure Port 1 as 1, Port 2 as 10, and the rest as 20.
Save the settings, plug the OPNsense device into port 1, and any LAN device into 3 or greater.
- Ensure in the GUI that you have something listed like below within the basic configuration of Interfaces: [LAN] :
Enable:        ☑️ Enable Interface
Identifier:    lan
Device:        vlan20
Description:   LAN
- Ensure in the GUI that you have something listed like below within the basic configuration of Interfaces: [WAN] :
Enable:        ☑️ Enable Interface
Identifier:    wan
Device:        vlan10
Description:   WAN
- If something seems off, go to Interfaces: Assignments and change the settings as listed:
| Interface | Identifier | Device | 
|---|---|---|
| [LAN] | lan | vlan20 VLAN_LAN (Parent: em0, Tag: 20) | 
| [WAN] | wan | vlan10 VLAN_WAN (Parent: em0, Tag: 10) | 
Two Ethernet Ports
Yay. This is very simple.
OPNsense requires at least one interface to be assigned to the LAN role so that an administrator can access the management plane. It will scan your motherboard, find the first two active network controllers (e.g., em0 and em1), and assign them to roles. By default, it assigns the lower-numbered hardware adapter (e.g., em0) to WAN and the second hardware adapter (e.g., em1) to LAN.
To summarize, you shouldn't have to do anything. But worst case scenario, go to the Interfaces: Assignments settings in the GUI and change the device bound to the specific interfaces to match what you want to do.
I highly recommend you set up static IPs for any devices that need it. I'll give an example below.
- Go to Services: Dnsmasq DNS & DHCP: General and select and apply the following settings:
| Option | Value | 
|---|---|
| Enable | ☑️ | 
| Interface | LAN | 
| Listen port | 53053 | 
| DNSSEC | 🔳 | 
| No hosts lookup | 🔳 | 
| Query DNS servers sequentially | 🔳 | 
| Require domain | 🔳 | 
| Do not forward to system defined DNS servers | 🔳 | 
| DHCP FQDN | ☑️ | 
| DHCP default domain | internal | 
| DHCP authoritative | 🔳 | 
| DHCP reply delay |  | 
| DHCP register firewall rules | ☑️ | 
| Router advertisements | ☑️ | 
| Disable HA sync | 🔳 | 
The remaining options can be left alone.
- 
Go to Services: Dnsmasq DNS & DHCP: Hosts and click the orange+ button to create an entry. Clicking on the orange ℹ️ will give you more information, but I will be generalizing for our use case.
  - 
Host: Enter your device name here, one that would be recognizable. If you've configured hostnames for a device using Linux, this should be similar. Example: homeserver .
  - 
Domain: Since my example are for LAN devices, use something like internal .
  - 
Local: Since we used internal above, select this.
  - 
IP addresses: This is where you can assign it a static IP address. Example: 192.168.1.50 .
  - 
Alias records: This gives it essentially another hostname. You shouldn't need this.
  - 
CNAME records: This gives it essentially another hostname when not assigning an IP. You shouldn't need this.
  - 
Client identifier: You shouldn't need this.
  - 
Hardware addresses: This is where you will input a device's network MAC address.
  - 
Lease time: Leave this as default.
- 
This is to make things easier while creating rules, for example, if you have multiple devices that need an exclusion rather than a single device. You can also reference single devices as well, and I'll include an example for one.
- 
Go to Firewall: Aliases and click the orange+ button to create an alias.
  - 
Name: Something recognizable, like homeserver .
  - 
Type: Host(s)
  - 
Content: This is where you'd put the previously configured static IP.
- 
- 
For a group of devices, create a single alias.
I'm going to give you some basic firewall rules. I'm aware that there are automatically generated rules, however, this will ensure you will have high security unless you explicitly delete your created rule. If you want to exempt a specific device or devices from the rule, you would select Source Invert and specify the source(s), which would be our previously created alias. I'll include an example only on the first rule and explain it.
OPNsense processes rules in order, which really matters because of match detection.
We want to force all DNS traffic to use OPNsense's DNS instead of bypassing it, right? All devices should use your secure and private DNS configuration. Some devices, like phones, smart TVs, and IoT devices hardcode their DNS and will bypass your configuration unless you do this. In this example, I invert the sources to exclude a specific alias from this force-redirection rule; homeserver in this case.
- Go to Firewall: NAT: Destination NAT and click the orange+ button to create a rule. Expand all the tabs so we can see all options.
| Option | Value | 
|---|---|
| Disabled | 🔳 | 
| Categories | Nothing selected | 
| Description | Force DNS to OPNsense (except homeserver) | 
| Interface | LAN | 
| Version | IPv4 | 
| Protocol | TCP/UDP | 
| Invert Source | ☑️ | 
| Source Address | homeserver | 
| Source Port | any | 
| Invert Destination | ☑️ | 
| Destination Address | LAN address | 
| Destination Port | DOMAIN (53) | 
| Redirect Target IP | LAN address | 
| Redirect Target Port | DOMAIN (53) | 
| Log | ☑️ | 
| Firewall rule | Register rule | 
Let me explain what exactly the invert and redirect settings mean.
| Option | Value | Operational Logic | 
|---|---|---|
| Invert Source | ☑️ | Apply this rule to any device except homeserver. | 
| Source Address | homeserver | Specifies the excluded alias. | 
| Invert Destination | ☑️ | Apply this rule to any destination except the firewall itself. | 
| Destination Address | LAN address | Specifies the local gateway IP to be excluded from redirection. | 
| Redirect Target IP | Loopback network (or 127.0.0.1) | The internal destination where the hijacked packets are sent. | 
- 
Scenario 1: A smart TV tries to bypass your network and queries a rogue DNS (8.8.8.8) 1: The packet enters the LAN interface with a destination of 8.8.8.8 .
2: The firewall checks the rule: Is8.8.8.8 NOT the LAN address (192.168.1.1 )?
