---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/therettom-self-hosted-guide-opnsense-3bc340bc-3
title: "Prune the oldest historical environment to control storage footprint"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/therettom-self-hosted-guide-opnsense-3bc340bc.md
source_anchor: ""
source_lines: [160, 284]
sha256: e39625229e6835c539350c7721fbec2d34c8bc29af9e5bb1e2c2ec670f99ce85
---

# Prune the oldest historical environment to control storage footprint

3: Result:True . The destination is inverted, so it matches the rule.
4: The NAT rule intercepts the packet, rewrites the destination to theLoopback network (127.0.0.1 ), and forces it to AdGuard Home/Unbound.
- 
Scenario 2: A standard client cleanly queries the OPNsense Gateway IP (192.168.1.1) 1: The packet enters the LAN interface with a destination of192.168.1.1 .
2: The firewall checks the rule: Is192.168.1.1 NOT the LAN address (192.168.1.1 )?
3: Result:False . The packet is destined for theLAN address , failing the inversion match.
4: The firewall skips the NAT rule, allowing the packet to pass directly to your local DNS listener without altering the data headers. This is also the correct, desired behavior.What happens if you uncheck Invert Destination ? The rule would only trigger if a client explicitly targeted 192.168.1.1. It would completely ignore rogue traffic bound for 8.8.8.8, allowing devices to bypass your local DNS entirely.
We didn't go over configuring a VLAN interface for IoT devices or something else, but figured I'd include this just in case. This rule is only applicable if, for example, you have IoT devices you don't communicating outside the LAN. We're creating a standard pass rule on the LAN (but it can be applied to others you might configure like an IoT VLAN or a local server VLAN). If a device on that network tries to talk, we configure OPNsense to stamp that packet with a local tag: NO_WAN_EGRESS, and prevent it from talking outside the LAN. It references an alias we didn't create, Isolated_Devices, that would cover IoT devices.
If you had a smart switch or even a regular switch hooked up to a smart switch, you'd give the Port itself a unique 801.1Q PVID (like 50), and make Port 1 a tagged member (if using Router-on-a-Stick method as configured earlier in the guide) and the VLAN 50 ports untagged.
| Option | Value | 
|---|---|
| Action | Pass | 
| Disabled | 🔳 Disable this rule | 
| Quick | ☑️ Apply the action immediately on match | 
| Interface | VLAN_IoT | 
| Direction | in | 
| TCP/IP Version | IPv4 | 
| Protocol | any | 
| Source / Invert | 🔳 Use this option to invert the sense of the match. | 
| Source | Isolated_Devices | 
| Destination / Invert | 🔳 Use this option to invert the sense of the match. | 
| Destination | any | 
| Destination port range | from / to: any | 
| Log | 🔳 Log packets that are handled by this rule | 
| Category |  | 
| Description | Allow isolated devices local access but tag for WAN block | 
Advanced options:
| Option | Value | 
|---|---|
| Set local tag | NO_WAN_EGRESS | 
- Go to Firewall: Rules: LAN and click the orange+ button to create rules.
| Option | Value | 
|---|---|
| Action | Pass | 
| Disabled | 🔳 Disable this rule | 
| Quick | ☑️ Apply the action immediately on match | 
| Interface | LAN | 
| Direction | in | 
| TCP/IP Version | IPv4 | 
| Protocol | any | 
| Source / Invert | 🔳 Use this option to invert the sense of the match. | 
| Source | LAN network | 
| Destination / Invert | 🔳 Use this option to invert the sense of the match. | 
| Destination | any | 
| Destination port range | from / to: any | 
| Log | 🔳 Log packets that are handled by this rule | 
| Category |  | 
| Description | Default rule, allow LAN to any | 
This rule is used to enforce tight security boundaries, like keeping isolated IoT devices, local media servers, or secure subnets completely locked inside the local network, even if someone accidentally changes a rule later. Because this floating rule sits at the very edge of your network (the WAN interface), we make a bulletproof catch-all. No matter what happens on the internal interfaces, and no matter what other rules are written later, if a packet carrying that tag tries to sneak out to the open internet, this floating rule kills it instantly.
This means we need to create a specific rule to help the floating rule work. But we already did that with Rule: Tag LAN Packets in LAN.
Why use a Floating rule with a Local Tag instead of a normal rule? It eliminates rule duplication. We write the blocking logic once on the WAN interface, rather than repeating blocks on every single VLAN tab. And if you accidentally add a Pass All rule on a local interface down the road while troubleshooting, that local interface rule might override a standard block rule. However, because this floating rule uses the Quick option at the WAN egress point, it acts as the absolute final authority. If the packet is tagged, it dies at the border, saving you from accidental internet leaks.
| Option | Value | 
|---|---|
| Action | Block | 
| Disabled | 🔳 Disable this rule | 
| Quick | ☑️ Apply the action immediately on match | 
| Interface | WAN | 
| Direction | out | 
| TCP/IP Version | IPv4 | 
| Protocol | any | 
| Source / Invert | 🔳 Use this option to invert the sense of the match. | 
| Source | any | 
| Destination / Invert | 🔳 Use this option to invert the sense of the match. | 
| Destination | any | 
| Destination port range | from / to: any | 
| Log | 🔳 Log packets that are handled by this rule | 
| Category |  | 
| Description |  | 
For Advanced Features:
| Option | Value | 
|---|---|
| allow options | 🔳 | 
| reply-to | default | 
| Set priority | All packets: Keep current priority and Low Delay/TCP ACK: Use main priority | 
| Match priority | Any priority | 
| Match TOS / DSCP | Any | 
| Set local tag |  | 
| Match local tag | NO_WAN_EGRESS | 
The rest of the options are default values or unset.
You can likely leave this as default. However, adding AdGuard Home is highly beneficial for DNS blocking, whether for stopping ads or protecting against certain domains from loading. I recommend you follow the AdGuard Home guide after reading the rest of this guide. As the overhead is minimal and it isn't complex to get running, I'm going to continue this section assuming you will be using AdGuard Home.
- Go to Services: Unbound DNS: General :
| Option | Value | 
|---|---|
| Enable Unbound | ☑️ | 
| Listen Port | 5335 | 
| Network Interfaces | All (recommended) | 
| Enable DNSSEC Support | ☑️ | 
| Enable DNS64 Support | 🔳 | 
| DNS64 Prefix |  | 
| Enable AAAA-only mode | 🔳 | 
| Register ISC DHCP4 Leases | ☑️ | 
| DHCP Domain Override |  | 
| Register DHCP Static Mappings | ☑️ | 
| Do not register IPv6 Link-Local addresses | 🔳 | 
| Do not register system A/AAAA records | 🔳 | 
| TXT Comment Support | 🔳 | 
| Flush DNS Cache during reload | 🔳 | 
| Local Zone Type | transparent | 
Apply the settings.
- Go to Services: Unbound DNS: Advanced . The first seven checkboxes are for security and privacy. The rest are mostly for performance. You should figure out what you're able to do based on your hardware by researching it. As there is a very broad range of hardware that can run OPNsense, I won't cover that. But here are the settings to check except one:
| Option | Value | 
|---|---|
| Hide Identity | ☑️ | 
| Hide Version | ☑️ | 
| Prefetch DNS Key Support | ☑️ | 
| Harden DNSSEC Data | ☑️ | 
| Harden Below NXDOMAIN | ☑️ | 
| Aggressive NSEC | ☑️ | 
| Strict QNAME Minimisation | 🔳 | 
- Go to Services: Unbound DNS: Access Lists and ensure you have LAN in there. If not, click the orange+ button to create one.
| Option | Value | 
|---|---|
| Enabled | ☑️ | 
| Access List Name | LAN | 
| Action | Allow | 
| Networks | 192.168.1.0/24 | 
| Description |  | 
- Go to Services: Unbound DNS: Query Forwarding . This can be redundant with ourRegister ISC DHCP4 Leases andRegister DHCP Static Mappings settings checked, but creating an entry here ensures Unbound knows about our DHCP clients. If you find that the logs for Unbound (or AdGuard Home) do not list clients in the query logs, then click the orange+ button to create one:
| Option | Value | 
|---|---|
| Domain | 1.168.192.in-addr.arpa | 
| Server IP | 127.0.0.1 | 
| Server Port | 53053 | 
| Forward first | 🔳 | 
| Description |  | 
