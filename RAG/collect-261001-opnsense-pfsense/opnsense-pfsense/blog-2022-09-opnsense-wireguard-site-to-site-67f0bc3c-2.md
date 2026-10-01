---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/blog-2022-09-opnsense-wireguard-site-to-site-67f0bc3c-2
title: "/etc/wireguard/wg0.conf"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/blog-2022-09-opnsense-wireguard-site-to-site-67f0bc3c.md
source_anchor: ""
source_lines: [98, 199]
sha256: 16a6b151110a62245e85d0c8e172c969ed69bd6bf4cb51bc73184e0da2985e59
---

# /etc/wireguard/wg0.conf

define lan_iface = "eth1"
define wg_iface = "wg0"
define wg_port = 51821
table inet filter {
    chain input {
        type filter hook input priority 0; policy drop;
        # accept all loopback packets
        iif "lo" accept
        # accept all icmp/icmpv6 packets
        meta l4proto { icmp, ipv6-icmp } accept
        # accept all packets that are part of an already-established connection
        ct state vmap { invalid : drop, established : accept, related : accept }
        # drop new connections over rate limit
        ct state new limit rate over 1/second burst 10 packets drop
        # accept all DNS/NTP/DHCPv6 packets received at a link-local address
        ip6 daddr fe80::/64 udp dport { domain, ntp, dhcpv6-server, dhcpv6-client } accept
        # accept all SSH packets received on the LAN interface
        iifname $lan_iface tcp dport ssh accept
        # accept all WireGuard packets received on the WAN interface
        iifname $wan_iface udp dport $wg_port accept
        # reject with polite "port unreachable" icmp response
        reject
    }
    chain forward {
        type filter hook forward priority 0; policy drop;
        # accept all packets that are part of an already-established connection
        ct state vmap { invalid : drop, established : accept, related : accept }
        # allow all packets outbound from Site A
        iifname $lan_iface accept
        # allow all packets inbound from Site B
        iifname $wg_iface accept
        # reject with polite "host unreachable" icmp response
        reject with icmpx type host-unreachable
    }
}
table inet nat {
    chain postrouting {
        type nat hook postrouting priority 100; policy accept;
        oifname $wan_iface masquerade
    }
}
See the Site to Site section of the WireGuard With Nftables guide for an example of how to restrict inbound access to Site A from Site B through WireGuard.
Configure Firewall on OPNsense
On Router B, you usually will need to make to changes to your firewall. First, you need to add a firewall rule that allows UDP port 51822 access from Router A.
In the OPNsense GUI, navigate to the Firewall > Rules > WAN page. Click the Add icon on it:
In the resulting firewall-rule edit page, set the TCP/IP Version field to IPv4/IPv6, the Protocol field to UDP, and the Destination field to WAN address. Set the Destination port range to (other) 51822:
Then click the Save button on that page, and click the Apply changes button on the resulting page:
If you’re only going to use the WireGuard tunnel to connect outbound from Site B to Site A, you don’t need to make any more changes. But if you want to allow inbound connections — like we do in our example scenario, where Endpoint A in Site A initiates connections to Endpoint B in Site B — you usually will need to add additional firewall rules to allow this access.
To do this, navigate to the Firewall > Rules > WireGuard page. Click the Add icon on it:
If you want to allow unrestricted inbound access from Site A to Site B, you only need to set the following fields in the resulting firewall-rule edit page:
- Source
- 
Set this to the network (or networks) used by Site A. In our example, this is just 192.168.1.0/24 .
- Destination
- 
Set this to the network (or networks) used by Site B. In our example, this is just 192.168.200.0/24 .
However, if you want to restrict access to allow only the single access scenario in our example, HTTP access from Endpoint A to Endpoint B, set these fields:
- Protocol
- 
In our example, we’re allowing just HTTP access, so select TCP .
- Source
- 
Endpoint A’s IP address on the Site A LAN is 192.168.1.11 , so select this single address.
- Destination
- 
Endpoint B’s IP address on the Site B LAN is 192.168.200.22 , so select this single address.
- Destination port range
- 
We’re allowing just HTTP access, so select the to HTTP option (port 80).
Click the Save button at the bottom of this page, and then click the Apply Changes button on the resulting page:
Test the Tunnel
Make sure you have a webserver on Endpoint B running on port 80. A simple substitute for a full-fledged webserver is to run Python with the http.server module:
$ sudo python3 -m http.server 80
This will serve the current directory via HTTP on port 80.
On Endpoint A, try to access the webserver on Endpoint B using Endpoint B’s LAN address (192.168.200.22):
$ curl 192.168.200.22
If you see any HTML output from this, then your WireGuard tunnel works!
Testing From the Routers Themselves
Note that if you try to test out connectivity by running the ping or curl commands from one of the routers themselves, it won’t work, since we haven’t included the IP address of the WireGuard interface from either router in the AllowedIPs setting of the other router.
However, assuming that the firewall rules you’ve set up on both routers allow unrestricted access between the Site A and Site B LANs, you can use each router’s own LAN address to test connectivity with the other site. (So you may want to start out with unrestricted access first, and then add more restrictive firewall rules once you know the tunnel is working.)
For example, to ping Endpoint B (192.168.200.22) from Router A (192.168.1.1), run this command:
$ ping -nc1 -I 192.168.1.1 192.168.200.22
The -n flag in the above command directs ping not to try to lookup hostnames, and the -c1 flag directs it to send just 1 packet. On Linux, the -I flag specifies the local source address to use.
On FreeBSD (the OS used by OPNsense), the -S flag specifies the local source address to use; so to ping Endpoint A (192.168.1.11) from Router B (192.168.200.1), run this command:
$ ping -nc1 -S 192.168.200.1 192.168.1.11
And to try to access the webserver running on Endpoint B (192.168.200.22) from Router A (192.168.1.1), run this command:
$ curl --interface 192.168.1.1 192.168.200.22
If you do want to allow regular access to Site A from Router B itself, without always having to specify which source address to use, add Router B’s WireGuard address (10.0.0.2) to Router A’s AllowedIPs setting:
# /etc/wireguard/wg0.conf
# local settings for Router A
[Interface]
PrivateKey = AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEE=
Address = 10.0.0.1/32
ListenPort = 51821
# remote settings for Router B
[Peer]
PublicKey = fE/wdxzl0klVp/IR8UcaoGUMjqaWi3jAd7KzHKFS6Ds=
Endpoint = 203.0.113.2:51822
AllowedIPs = 192.168.200.0/24
AllowedIPs = 10.0.0.2
And to allow regular access to Site B from Router A itself, add Router A’s WireGuard address (10.0.0.1) to the Allowed IPs setting of the endpoint for Router A on Router B:
Also make sure you adjust your firewall rules to allow access from these WireGuard addresses.
