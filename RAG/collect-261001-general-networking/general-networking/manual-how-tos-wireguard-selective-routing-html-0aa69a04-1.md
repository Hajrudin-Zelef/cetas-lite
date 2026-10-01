---
id: collect-261001-general-networking/general-networking/manual-how-tos-wireguard-selective-routing-html-0aa69a04-1
title: "manual-how-tos-wireguard-selective-routing-html-0aa69a04"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-how-tos-wireguard-selective-routing-html-0aa69a04.md
source_anchor: ""
source_lines: [1, 52]
sha256: 280ce5679948d5e9ff8223c5b9aa81377bf896c46a8b3992b8bfb929710e8332
---

# manual-how-tos-wireguard-selective-routing-html-0aa69a04

WireGuard Selective Routing to External VPN Endpoint
Introduction
This how-to is designed to assist with setting up WireGuard on OPNsense to use selective routing to an external VPN peer - most commonly to an external VPN provider.
These circumstances may apply where only certain local hosts are intended to use the VPN tunnel. Or it could apply where multiple connections to the VPN provider are desired, with each connection intended to be used by different specific local hosts.
This how-to focuses on the configuration of OPNsense. You will also have to configure the peer at your VPN provider - consult your VPN provider’s documentation as to how to do that.
Your OPNsense WireGuard Instance public key will need to be registered with your VPN provider, and you will need to get your VPN provider’s endpoint public key and the VPN tunnel IP provided for your WireGuard Instance by your VPN provider. In some cases, you will not be able to get the Peer public key and VPN tunnel IP until you register your WireGuard Instance public key. In that case, create the OPNsense Instance configuration first, using a dummy tunnel IP and no peer selected, so that the public key is generated, and then update the configuration later once the other information is known.
For an example of configuring the peer at a VPN provider (Mullvad), see Step 1 of the how-to WireGuard MullvadVPN Road Warrior Setup.
This how-to primarily focuses on IPv4 configuration. It can be readily adapted for IPv6 as well. See Configuring IPv6 below.
Step 1 - Configure the peer
- Go to
- Click + to add a new Peer
- Configure the Peer as follows (if an option is not mentioned below, leave it as the default): Enabled Checked Name Call it whatever you want (eg VPNProviderName_Location )Public Key Insert the public key from your VPN provider Allowed IPs 0.0.0.0/0 Endpoint Address Insert the public IP address (desirably) or domain name of your VPN provider, as provided by it Endpoint Port Insert the port of your VPN provider, as provided by it Keepalive 25
- Save the Peer configuration, and then click Save again
Step 2 - Configure the WireGuard Instance
- Go to
- Click + to add a new Instance configuration
- Turn on “advanced mode”
- Configure the Instance configuration as follows (if an option is not mentioned below, leave it as the default): Enabled Checked Name Call it whatever you want (eg VPNProviderName )Public Key This will initially be blank; it will be populated once the configuration is saved Private Key This will initially be blank; it will be populated once the configuration is saved Listen Port 51820 or a higher numbered unique port DNS Server Leave this blank, otherwise WireGuard will overwrite OPNsense’s DNS configuration Tunnel Address Insert the WireGuard Instance VPN tunnel IP provided by your VPN provider, in CIDR format, eg 10.24.24.10/32 Peers In the dropdown, select the Peer you configured above Disable Routes Checked Gateway Specify an IP that is 1 number below your VPN tunnel IP, eg 10.24.24.9 - see note below
Note
The IP you choose for the Gateway is essentially arbitrary; pretty much any unique IP will do. The suggestion here is for convenience and to avoid conflicts
- Save the Instance configuration, and then click Save again
Step 3 - Turn on WireGuard
Turn on WireGuard under if it is not already on
Step 4 - Assign an interface to WireGuard and enable it
- Go to
- In the dropdown next to “New interface:”, select the WireGuard device ( wg0 if this is your first one)
- Add a description (eg WAN_VPNProviderName )
- Click + to add it, then click Save
- Then select your new interface under the Interfaces menu
- Configure it as follows (if an option is not mentioned below, leave it as the default): Enable Checked Lock Checked if you wish to Description Same as under Assignments, if this box is not already populated IPv4 Configuration Type None IPv6 Configuration Type None
- Save the interface configuration and then click Apply changes
Step 5 - Restart WireGuard
Now restart WireGuard - you can do this from the Dashboard (if you have the services widget) or by turning it off and on under
Step 6 - Create a gateway
- Go to
- Click Add
- Configure the gateway as follows (if an option is not mentioned below, leave it as the default): Name Call it whatever you want, easiest to name it the same as the interface Description Add one if you wish to Interface Select your newly created interface in the dropdown Address Family Select IPv4 in the dropdown IP address Insert the gateway IP that you configured under the WireGuard Instance configuration Far Gateway Checked Disable Gateway Monitoring Unchecked Monitor IP Insert the endpoint VPN tunnel IP (NOT the public IP) of your VPN provider - see note below
Note
Specifying the endpoint VPN tunnel IP is preferable. As an alternative, you could include an external IP such as 1.1.1.1 or 8.8.8.8, but be aware that this IP will only be accessible through the VPN tunnel (OPNsense creates a static route for it), and therefore will not be accessible from local hosts that are not using the tunnel. Use the Disable Host Route check box if you wish to use an external IP AND it still be accessible by everything.
Some VPN providers will include the VPN tunnel IP of the endpoint in the configuration data they provide. For others (such as Mullvad), you can get the IP by running a traceroute from a host that is using the tunnel - the first hop after OPNsense is the VPN provider’s tunnel IP. Please note that this IP must be pingable. If the IP does not respond to ping (as is the case for Mullvad), use a different IP in the traceroute chain or disable the gateway monitoring altogether by checking the box Disable Gateway Monitoring.
- Save the gateway configuration and then click Apply changes
Step 7 - Create an Alias for the relevant local hosts that will access the tunnel
- Go to
- Click + to add a new Alias
- Configure the Alias as follows (if an option is not mentioned below, leave it as the default): Enabled Checked Name Call it whatever your want, eg WG_VPN_HostsType Select either Host(s) or Network(s) in the dropdown, depending on whether you want specific host IPs to use the tunnel, or an entire local network (such as a VLAN) Content Enter the host IPs, or the network in CIDR format Description Add one if you wish to
- Save the Alias, and then click Apply
Step 8 - Create a firewall rule
The purpose of this step is to create a firewall rule to allow the relevant hosts to access the tunnel. At the same time, it also ensures that the relevant hosts using the tunnel can still access local resources as necessary - such as a local DNS server, or file storage
The step has two parts - first creating a second Alias for all local (private) networks, and then creating the firewall rule itself. The ultimate effect of these two steps is that only traffic from the relevant hosts that is destined for non-local destinations will be sent down the tunnel
Note
The rule below will mean that no local (private) IPs can be accessed over the tunnel. You may have a need however to access certain IPs or networks at the VPN endpoint, such as a DNS server or monitor IP. In that case, you will need to create an additional firewall rule in OPNsense to ensure that requests to those IPs/networks use the tunnel gateway rather than the normal WAN gateway. This rule would be similar to that created below, except that the destination would be the relevant IPs/networks (or a new Alias for them) and the destination invert box would be unchecked. This rule would also need to be placed above the rule created below
Warning
