---
id: collect-261001-general-networking/general-networking/manual-how-tos-wireguard-client-html-1429b58c-2
title: "manual-how-tos-wireguard-client-html-1429b58c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-how-tos-wireguard-client-html-1429b58c.md
source_anchor: ""
source_lines: [67, 132]
sha256: 231318d201a045997e2ba8466aa47ca6329cc9dbe0174654ff5aa52cad0ab34b
---

# manual-how-tos-wireguard-client-html-1429b58c

If you didn’t assign an interface as suggested in Step 4(a), then you will need to manually specify the source IPs/subnet(s) for the tunnel (for example, 10.10.10.0/24). It’s probably easiest to define an alias (via ) for those IPs/subnet(s) and use that. If you have only one WireGuard Instance and only one WireGuard Peer configured, you can use the default WireGuard net, although this is generally not recommended due to unexpected behaviour
Step 5 - Create firewall rules
This will involve two steps - first creating a firewall rule on the WAN interface to allow clients to connect to the OPNsense WireGuard server, and then creating a firewall rule to allow access by the clients to whatever IPs they are intended to have access to.
- Go to
- Click Add to add a new rule
- Configure the rule as follows (if an option is not mentioned below, leave it as the default): Action Pass Quick Checked Interface WAN Direction in TCP/IP Version IPv4 or IPv4+IPv6 (as desired, depending on how you want clients to connect to the server; note this is distinct from what type of traffic is allowed in the tunnel once established) Protocol UDP Source / Invert Unchecked Source any Destination / Invert Unchecked Destination WAN address Destination port range The WireGuard port specified in the Instance configuration in Step 1 Description Add one if you wish to
- Save the rule, and then click Apply Changes
- Then go to - see note below if you didn’t assign this interface
- Click Add to add a new rule
- Configure the rule as follows (if an option is not mentioned below, leave it as the default): Action Pass Quick Checked Interface Whatever interface you are configuring the rule on (eg HomeWireGuard ) - see note belowDirection in TCP/IP Version IPv4 or IPv4+IPv6 (as applicable) Protocol any Source / Invert Unchecked Source If you assigned an interface under Step 4(a), select the generated alias for the interface subnet(s) (eg HomeWireGuard net ) - see note below if you didn’t assign this interfaceDestination / Invert Unchecked Destination Specify the IPs that client peers should be able to access, eg “any” or specific IPs/subnets Destination port range any Description Add one if you wish to
- Save the rule, and then click Apply Changes
Note
If you didn’t assign an interface as suggested in Step 4(a), then the second firewall rule outlined above will need to be configured on the automatically created WireGuard group that appears once the Instance configuration is enabled and WireGuard is started. You will also need to manually specify the source IPs/subnet(s) for the tunnel. It’s probably easiest to define an alias (via ) for those IPs/subnet(s) and use that. If you have only one WireGuard Instance and only one WireGuard Peer configured, you can use the default WireGuard net, although this is generally not recommended due to unexpected behaviour
Step 5a - Create normalization rules
- Go to and press + to create one new normalization rule.
  - If you only pass IPv4 traffic through the wireguard tunnel, create the following rule:
  - Interface WireGuard (Group) Direction Any Protocol any Source any Destination any Destination port any Description Wireguard MSS Clamping IPv4 Max mss 1380 (default) or 1372 if you use PPPoE; it’s 40 bytes less than your Wireguard MTU
- Save the rule
  - If you pass IPv4+IPv6 - or only IPv6 traffic - through the wireguard tunnel, create the following rule:
  - Interface WireGuard (Group) Direction Any Protocol any Source any Destination any Destination port any Description Wireguard MSS Clamping IPv6 Max mss 1360 (default) or 1352 if you use PPPoE; it’s 60 bytes less than your Wireguard MTU
- Save the rule
Tip
- The header size for IPv4 is usually 20 bytes, and for TCP 20 bytes. In total that is 40 bytes for IPv4 TCP.
- IPv6 has a larger header size with 40 bytes. That encreases the total to 60 bytes for IPv6 TCP.
Note
By creating the normalization rules, you ensure that IPv4 TCP and IPv6 TCP can pass through the Wireguard tunnel without being fragmented. Otherwise you could get working ICMP and UDP, but some encrypted TCP sessions will refuse to work.
Step 6 - Configure the WireGuard client
Tip
Key generation can be performed on an appropriate device with WireGuard client tools installed. A one-liner for generating a matching private and public keypair is wg genkey | tee private.key | wg pubkey > public.key. Alternatively, WireGuard apps that can be used on some devices can automate key generation for you
Client configuration is largely beyond the scope of this how-to since there is such a wide array of possible targets (and corresponding configuration methods). An example client (and server) configuration is in the Appendix. The key pieces of information required to configure a client are described below:
[Interface]
Address
Refers to the IP(s) specified as Allowed IPs in the Peer configuration on OPNsense. For example, 10.10.10.2/32
PrivateKey
Refers to the private key that (along with a public key) needs to be manually or automatically generated on the client. The corresponding public key must then be copied into the Peer configuration on OPNsense for the relevant client peer - see Step 2
DNS
Refers to the DNS servers that the client should use for the tunnel (see note below). For example, 10.10.10.1
[Peer]
PublicKey
Refers to the public key that is generated on OPNsense. Copy the public key from the Instance configuration on OPNsense - see Step 1
Endpoint
Refers to the public IP address or publicly resolvable domain name of your OPNsense host, and the port specified in the Instance configuration on OPNsense
AllowedIPs
Refers to the traffic (by destination IPs/subnets) that is to be sent via the tunnel. For example, if all traffic on the client is to be sent through the tunnel, specify 0.0.0.0/0 (IPv4) and/or ::/0 (IPv6)
Note
If the DNS server(s) specified are only accessible over the tunnel, or you want them to be accessed over the tunnel, make sure they are covered by the AllowedIPs
Appendix - Example configurations
Warning
Do not reuse these example keys!
An example client configuration file:
[Interface]
PrivateKey = 8GboYh0YF3q/hJhoPFoL3HM/ObgOuC8YI6UXWsgWL2M=
Address = 10.10.10.2/32, fd00:1234:abcd:ef09:10:2/128
DNS = 192.168.1.254, fd00:1234:abcd:ef09:1:254
[Peer]
PublicKey = OwdegSTyhlpw7Dbpg8VSUBKXF9CxoQp2gAOdwgqtPVI=
AllowedIPs = 0.0.0.0/0, ::/0
Endpoint = opnsense.example.com:51820
An example server configuration file:
[Interface]
Address = 10.10.10.1/24, fd00:1234:abcd:ef09:10:1/64
ListenPort = 51820
PrivateKey = YNqHwpcAmVj0lVzPSt3oUnL7cRPKB/geVxccs0C0kk0=
[Peer]
PublicKey = CLnGaiAfyf6kTBJKh0M529MnlqfFqoWJ5K4IAJ2+X08=
AllowedIPs = 10.10.10.2/32, fd00:1234:abcd:ef09:10:2/128
