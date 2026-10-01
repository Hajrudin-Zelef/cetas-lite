---
id: collect-261001-general-networking/general-networking/manual-how-tos-wireguard-client-html-1429b58c-1
title: "manual-how-tos-wireguard-client-html-1429b58c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-how-tos-wireguard-client-html-1429b58c.md
source_anchor: ""
source_lines: [1, 66]
sha256: 1d13f2db96af0907bedfdfe7626eb921cbd07c40ec38deeee180719a15310505
---

# manual-how-tos-wireguard-client-html-1429b58c

WireGuard Road Warrior Setup
Introduction
WireGuard is a simple, fast VPN protocol using modern cryptography. It aims to be faster and less complex than IPsec whilst also being a considerably more performant alternative to OpenVPN. Initially released for the Linux kernel, it is now cross-platform and widely deployable.
This how-to describes setting up a central WireGuard Instance (server) on OPNsense and configuring one or more client peers to create a tunnel to it.
Step 1 - Configure the Wireguard Instance
- Go to
- Click + to add a new Instance configuration
- Configure the Instance configuration as follows (if an option is not mentioned below, leave it as the default): Enabled Checked Name Call it whatever you want (eg HomeWireGuard )Public Key This will initially be blank; Press the cogwheel to auto-generate new keys. Private Key This will initially be blank; Press the cogwheel to auto-generate new keys. Listen Port 51820 or a higher numbered unique port MTU 1420 (default) or 1412 if you use PPPoE; it’s 80 bytes less than your WAN MTU Tunnel Address For example, 10.10.10.1/24. See note below Peers The (client) peers will be specified here; leave it blank initially until the Peer configuration is created in Step 2 Disable Routes Unchecked
Note
The tunnel address must be in CIDR notation and must be a unique IP and subnet for your network, such as if it was on a physically different routed interface. The subnet should be an appropriate size that includes all the client peers that will use the tunnel. For IPv4 it should be a private (RFC1918) address, for example 10.10.10.1/24. For IPv6, it could either be a unique ULA /64 address, or a unique GUA /64 address derived from your prefix delegation. Do not use a tunnel address that is a /32 (IPv4) or a /128 (IPv6)
Note
Leave the DNS Server field (which appears if advanced mode is selected) blank. Otherwise WireGuard will overwrite OPNsense’s DNS configuration
- Save the Instance configuration, and then click Save again
- Re-open the Instance configuration
- Copy the public key that has been generated in the configuration. This will be needed for the client device - see Step 6
- Save or Cancel to exit the configuration
Step 2 - Configure the client peer
Tip
Peers can be generated using the new peer generator feature under .
If using the peer generator and require Unbound DNS to serve names, fill the DNS server with the tunnel address (eg 10.10.10.1 ).
- Go to
- Click + to add a new Peer
- Configure the Peer as follows (if an option is not mentioned below, leave it as the default): Enabled Checked Name Call it whatever you want (eg Phone )Public Key Insert the public key from the client; if needed skip ahead and start Step 6 to generate the client public key Allowed IPs Unique tunnel IP address (IPv4 and/or IPv6) of client - it should be a /32 or /128 (as applicable) within the subnet configured on the WireGuard Instance. For example, 10.10.10.2/32
- Save the Peer configuration, and then click Apply
- Now go back to
- Open the Instance configuration that was created in Step 1 (eg HomeWireGuard )
- In the Peers dropdown, select the newly created Peer (eg Phone )
- Save the Instance configuration again, and then click Apply
- Repeat this Step 2 for as many clients as you wish to configure
Step 3 - Turn on/restart WireGuard
- Turn on WireGuard under if it is not already on (click Apply after checking the checkbox)
- Otherwise, restart WireGuard - you can do this by turning it off and on under (click Apply after both unchecking and checking the checkbox)
Step 4 - Assignments and routing
Note
The steps outlined in Steps 4(a) and 4(b) below may not be required at all in your circumstances. Strictly speaking, if you only intend for your clients to use the tunnel to access local IPs/subnets behind OPNsense, then neither step is actually necessary. If you intend to use the WireGuard tunnel to also access IPs outside of the local network, for example the public internet, then at least one, and perhaps both, of the steps will be required. This is explained below
However, it is useful to complete Step 4(a) anyway, for the reasons explained in that step
Step 4(a) - Assign an interface to WireGuard (recommended)
Hint
This step is not strictly necessary in any circumstances for a road warrior setup. However, it is useful to implement, for several reasons:
First, it generates an alias for the tunnel subnet(s) that can be used in firewall rules. Otherwise you will need to define your own alias or at least manually specify the subnet(s)
Second, it automatically adds an IPv4 Source NAT rule, which will allow the tunnel to access IPv4 IPs outside of the local network (if that is desired), without needing to manually add a rule
Finally, it allows separation of the firewall rules of each WireGuard instance (each wgX device). Otherwise they all need to be configured on the default WireGuard group that OPNsense creates. This is more an organisational aesthetic, rather than an issue of substance
- Go to
- In the dropdown next to “New interface:”, select the WireGuard device ( wg1 if this is your first one)
- Add a description (eg HomeWireGuard )
- Click + to add it, then click Save
- Then select your new interface under the Interfaces menu
- Configure it as follows (if an option is not mentioned below, leave it as the default): Enable Checked Lock Checked Description Same as under Assignments, if this box is not already populated IPv4 Configuration Type None IPv6 Configuration Type None
Note
There is no need to configure IPs on the interface. The tunnel address(es) specified in the Instance configuration for your server will be automatically assigned to the interface once WireGuard is restarted
- Save the interface configuration and then click Apply changes
- Restart WireGuard - you can do this by turning it off and on under (click Apply after both unchecking and checking the checkbox)
Tip
When assigning interfaces, gateways can be added to them. This is useful if balancing traffic across multiple tunnels is required or in more complex routing scenarios. To do this, go to and add a new gateway. Choose the relevant WireGuard interface under and check the checkbox Dynamic gateway policy. These scenarios are otherwise beyond the scope of this how-to.
Tip
If Unbound DNS is configured with all interfaces registered it requires a reload of Unbound DNS to get the new Wireguard interface added. This is necessary to get DNS working through the VPN tunnel.
Step 4(b) - Create a Source NAT rule
Hint
This step is only necessary (if at all) to allow client peers to access IPs outside of the local IPs/subnets behind OPNsense - see the note under Step 4. If an interface has already been assigned under Step 4(a), then it is not necessary for IPv4 traffic, and is only necessary for IPv6 traffic if the tunnel uses IPv6 ULAs (IPv6 GUAs don’t need NAT). So in many use cases this step can be skipped
- Go to
- Select “Hybrid Source NAT rule generation” if it is not already selected, and click Save and then Apply changes
- Click Add to add a new rule
- Configure the rule as follows (if an option is not mentioned below, leave it as the default): Interface WAN TCP/IP Version IPv4 or IPv6 (as applicable) Protocol any Source invert Unchecked Source address If you assigned an interface under Step 4(a), select the generated alias for the interface subnet(s) (eg HomeWireGuard net ) - see note below if you didn’t assign this interfaceSource port any Destination invert Unchecked Destination address any Destination port any Translation / target Interface address Description Add one if you wish to
- Save the rule, and then click Apply changes
- Restart WireGuard - you can do this by turning it off and on under (click Apply after both unchecking and checking the checkbox)
Hint
