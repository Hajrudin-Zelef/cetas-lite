---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/posts-opnsense-wireguard-vpn-846f08f8-4
title: "posts-opnsense-wireguard-vpn-846f08f8"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/posts-opnsense-wireguard-vpn-846f08f8.md
source_anchor: ""
source_lines: [306, 440]
sha256: 4d97def6855d80401800632a08aeecfba360f71e12ae7ff76daf0557ef3bd74f
---

# posts-opnsense-wireguard-vpn-846f08f8

Finish WireGuard interface by resetting WireGuard services
The save and apply are meaningless, as WireGuard never resets the service to load the new configuration. You must be sure to either check and uncheck the Enable WireGuard button in Settings ▸ General or go to the Dashboard and reset services from there.
Unbound DNS requires a reload of Unbound DNS’s services to get the new WireGuard interface added.
PART VI - OPNsense rules for WireGuard
This area is where the networking configuration begins.
This may be a good time to go make a tea ☕ and grab a snack.
We are starting to shape where our traffic can and cannot go.
Please note this can be changed for preference and is not ridged.
Create a WireGuard outbound NAT rule
Detailing outbound NAT changes
This step is only necessary if you intend to allow client peers to access IPs outside of the local IPs/subnets behind OPNsense.
Think VPN provider, do you want to allow wgopn1memestor clients to forward their Internet traffic over your network?
WireGuard subnets that you want to have internet access need their subnets in Firewall ▸ NAT ▸ Outbound.
Edits made to NAT in brief:
- Interface : WAN
- Source : whatever theLocal Tunnel subnet is set to.
Details of outbound NAT rule
Note: Changes made to the rule are highlighted as code blocks, areas that were not modified are written as standard text.
- Go to Firewall ▸ NAT ▸ Outbound
- Select Hybrid outbound NAT rule generation at the top, if it is not already selected.
- Save and thenApply changes
- Click Add to add a new rule (Click on the + symbol)
- Interface - WAN
- TCP/IP Version - IPv4 or IPv6 (as applicable)
- Protocol - any
- Source invert - Unchecked
- Source address - Select the network of our new interface: wgopn1memestor net
- Source port - any
- Destination invert - Unchecked
- Destination address - any
- Destination port - any
- Translation / target - Interface address
- Description - Allow traffic from wgopn1memestor to outbound LAN/Internet
- Save
- Apply
Setup firewall rules on OPNsense for WireGuard
This will involve two steps.
- Creating a firewall rule on the WAN interface to allow clients to connect to the OPNsense WireGuard server.
- Creating a firewall rule to allow access by the clients to whatever IPs on the local network they are intended to have access to.
Create firewall rules on - WAN
Letting in your port you made for WireGuard opens your firewall up. You now have a hole in your network, on the port you choose.
Be aware, UDP hole-punching is available in Wireguard, but it is one way. See the PersistentKeepalive section. VPNs like Tailscale, Zerotier, Netbird all work with this method.
Now that you’re aware of the risks and alternatives, let’s begin:
- Go to Firewall ▸ Rules ▸ WAN
- Click Add to add a new rule (Click on the + symbol)
- Action - Pass
- Quick - Checked
- Interface - WAN
- Direction - in
- Protocol - UDP
- Source / Invert - Unchecked
- Source - any
- Destination / Invert - Unchecked
- Destination - WAN address
- Destination port range - Select (other) . The number to enter is probably the default,51820 , but check the WireGuard port you set in the Instance configuration on an earlier step.
- Description - WireGuard in WAN allow
- Save
- Apply
Create firewall rules on - WireGuard
The firewall rule outlined below will need to be configured on the automatically created WireGuard group that appears once the Instance configuration is enabled and WireGuard is started.
You will also need to manually specify the subnet for the tunnel.
You can also define an alias (via Firewall ▸ Aliases) for any IPs/subnet that you want to use.
- Go to Firewall ▸ Rules ▸ WireGuard (Group)
- Click Add (Click on the + symbol) to add a new rule.
- Action - Pass
- Quick - Checked
- Interface - WireGuard (Group) or an alias
- Direction - in
- TCP/IP Version - IPv4 or IPv4+IPv6 (as applicable)
- Protocol - any
- Source / Invert - Unchecked
- Source - WireGuard (Group) net or an alias
- Destination / Invert - Unchecked
- Destination - Specify the IPs or subnet that client peers should be able toaccess . You can use an alias here too.
- Destination port range - any
- Description - Any WireGuard interface can access these networks
- Save
- Apply
COMPLETE
Working WireGuard Site-to-Site VPN with OPNsense
You should now have a working WireGuard VPN.
- The instance has been made, this created ourinterfaces , and we enabled them.
- The WireGuard peer is connecting to anEndpoint addresss and port – which is an open WAN port set in the firewall.
- The server’s instance Public Key is set for other machine’s connecting ‘peer’.
- Allowed IPs reflect the traffic that is allowed to pass through the tunnel.
- Firewall rules are in place to allow the WireGuard interface onto the router.
This traffic can come from anywhere and go anywhere. We should further restrict this…
PART VII - Site-to-Site VPN - OPNsense firewall configuration
You may skip this section if you do not require detailed firewalling for a site-to-site WireGuard VPN. Please note, this is not required, most can be changed for preference, and should not be considered ridged.
If you are strictly using this as a way to remote devices into your OPNsense router’s subnet, you will not need to concern yourself with the details below.
Site-to-Site WireGuard WAN connection
Think of this section like two peers connecting to each other.
You will be editing the firewall settings for the WAN connection between the two WireGuard clients.
For reference, link to table of all addresses used in this writeup.
Site A - WAN Firewall setup - Meme Storage Bunker HQ’s Server
Back at the bunker, we need to add a new rule to allow incoming WireGuard traffic from Site B (Sarah’s Flower Shop).
- Go to Firewall ▸ Rules ▸ WAN
- Click Add (Click on the + symbol) to add a new rule.
- Action - Pass
- Interface - WAN
- Direction - In
- TCP/IP Version - IPv4
- Protocol - UDP
- Source - Single host or Network and set this to the remote site’s WAN address:203.0.113.2
- Destination - WAN address set this to the WAN address to allow on WAN from our remote source.
- Destination port - 51820
- Description - Allow WireGuard from remote Site B to this Site A
Site B - WAN Firewall setup - Sarah’s Flower Shop Server
The same step needs to be taken, but with the WAN addresses reversed - allow incoming WireGuard traffic from Site A (Meme Storage Bunker HQ).
- Go to Firewall ▸ Rules ▸ WAN
- Click Add (Click on the + symbol) to add a new rule.
- Action - Pass
- Interface - WAN
- Direction - In
- TCP/IP Version - IPv4
- Protocol - UDP
- Source - Single host or Network and set this to the remote site’s WAN address:203.0.113.1
- Destination - WAN address set this to the WAN address to allow a WAN connection from our remote source.
- Destination port - 51820
- Description - Allow WireGuard from remote Site A to this Site B
- Press Save and Apply.
Verify WireGuard connection on Site A and Site B
Load new WireGuard configuration
Ensure you are able to reset the WireGuard service to load the new configuration.
- Go to VPN ▸ WireGuard ▸ Settings on both sites andcheck anduncheck theEnable WireGuard and pressApply .
- Go to the Dashboard on both sites andreset services from there.
Check WireGuard Logs
To verify any of this is working correctly, go to VPN ▸ WireGuard ▸ Diagnostics.
You should see Send and Received traffic and Handshake should be populated by a number. This happens as soon as the first traffic flows between the sites.
If you see this, your tunnel is now up and running.
PART VIII - Site-to-Site VPN - OPNsense router configuration
Routing different subnets across WireGuard
Different subnets separate networks from communicating with each other. The firewall also stops these networks.
Currently, your two LANs cannot see each other.
We’re going to add firewall rules to allow these two sites to communicate like they were in the same room.
