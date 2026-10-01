---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1172890-openvpn-on-opnsense-everything-routed-through-tunnel-but-shoul-c89ff15d
title: "OpenVPN on OPNsense, everything routed through tunnel but shouldn&#39;t"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1172890-openvpn-on-opnsense-everything-routed-through-tunnel-but-shoul-c89ff15d.md
source_anchor: ""
source_lines: [1, 93]
sha256: 966b6bfc7b484fe6df854f564a1289a9925f07a62de6011284763937bade29c3
---

# OpenVPN on OPNsense, everything routed through tunnel but shouldn&#39;t

*Score : 1 | Source : https://serverfault.com/questions/1172890/openvpn-on-opnsense-everything-routed-through-tunnel-but-shouldnt*

I have the following Setup An OPNsense on which OPNVPN is running. I created an VPN instance and also some users. As the "Server (IPv4)" I have set 10.123.0.0/24. Under Routing I have set nothing. Then I use the "Client Specific Overrides" feature. There I set for demouser the "Local Network" to 10.10.75.0/24, so that this routing information gets forwarded to the Client.
My Problem now is, that on my Linux Mont Machine, all traffic is forwarded through this tunnel although I did NOT define 0.0.0.0
On an Windows PC it is working correctly with the same config
with VPN enabled:
user@pc:~$netstat -r
Kernel IP routing table
Destination     Gateway         Genmask         Flags   MSS Window  irtt Iface
default         _gateway        0.0.0.0         UG        0 0          0 tun0
default         _gateway        0.0.0.0         UG        0 0          0 enxa02919c350d5
default         _gateway        0.0.0.0         UG        0 0          0 wlp0s20f3
10.10.2.0       0.0.0.0         255.255.255.0   U         0 0          0 wlp0s20f3
10.10.3.0       0.0.0.0         255.255.255.0   U         0 0          0 enxa02919c350d5
_gateway        0.0.0.0         255.255.255.255 UH        0 0          0 enxa02919c350d5
10.10.75.0      _gateway        255.255.255.0   UG        0 0          0 tun0
10.123.0.0      0.0.0.0         255.255.255.0   U         0 0          0 tun0
<provider_DNS_something> _gateway        255.255.255.255 UGH       0 0          0 enxa02919c350d5
user@pc:~$ route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         10.123.0.1      0.0.0.0         UG    50     0        0 tun0
0.0.0.0         10.10.3.1       0.0.0.0         UG    100    0        0 enxa02919c350d5
0.0.0.0         10.10.2.1       0.0.0.0         UG    600    0        0 wlp0s20f3
10.10.2.0       0.0.0.0         255.255.255.0   U     600    0        0 wlp0s20f3
10.10.3.0       0.0.0.0         255.255.255.0   U     100    0        0 enxa02919c350d5
10.10.3.1       0.0.0.0         255.255.255.255 UH    50     0        0 enxa02919c350d5
10.10.75.0      10.123.0.1      255.255.255.0   UG    50     0        0 tun0
10.123.0.0      0.0.0.0         255.255.255.0   U     50     0        0 tun0
<Public IP of router>  10.10.3.1       255.255.255.255 UGH   50     0        0 enxa02919c350d5
without VPN:
user@pc:~$ netstat -r
Kernel IP routing table
Destination     Gateway         Genmask         Flags   MSS Window  irtt Iface
default         _gateway        0.0.0.0         UG        0 0          0 enxa02919c350d5
default         _gateway        0.0.0.0         UG        0 0          0 wlp0s20f3
10.10.2.0       0.0.0.0         255.255.255.0   U         0 0          0 wlp0s20f3
10.10.3.0       0.0.0.0         255.255.255.0   U         0 0          0 enxa02919c350d5
user@pc:~$ route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         10.10.3.1       0.0.0.0         UG    100    0        0 enxa02919c350d5
0.0.0.0         10.10.2.1       0.0.0.0         UG    600    0        0 wlp0s20f3
10.10.2.0       0.0.0.0         255.255.255.0   U     600    0        0 wlp0s20f3
10.10.3.0       0.0.0.0         255.255.255.0   U     100    0        0 enxa02919c350d5
When I manually add routing entries to the VPN on the Client. not everything get's routed through the tunnel.
Can you help me figuring out, which Settings need to be changed on the OPNsense?
PS: The "Redirect Gateway" dropdown menü is left empty.
Append: Server-config (pictures, because done in GUI: Instance: Server-config-01 Server-config-02 Server-config-03
Client Specific Overrides: CSO-01 CSO-02
Client-config: (exported from the OPNsense)
dev tun
persist-tun
persist-key
client
resolv-retry infinite
remote <insert_DNS> 1196 udp
lport 0
verify-x509-name "C=DE, CN=OpenVPN-Server" subject
remote-cert-tls server
auth-user-pass
<ca>
-----BEGIN CERTIFICATE-----
<insert_CERT>
-----END CERTIFICATE-----
</ca>
<cert>
-----BEGIN CERTIFICATE-----
<insert_Cert>
-----END CERTIFICATE-----
</cert>
<key>
-----BEGIN PRIVATE KEY-----
<insert_Cert priv key>
-----END PRIVATE KEY-----
</key>
<tls-crypt>
#
# 2048 bit OpenVPN static key
#
-----BEGIN OpenVPN Static key V1-----
<insert_OpenVPN static key>
-----END OpenVPN Static key V1-----
</tls-crypt>

---

### Reponse — score 1

I found the Answer. It was related to an unchecked box in the Linux Mint Networkmanager: Under the Setting of the specified VPN connection there is an checkbox which says "Use this connection only for resources on its network". After I checked it everything is working right.
This checkbox changes the priorities in the routing table.
