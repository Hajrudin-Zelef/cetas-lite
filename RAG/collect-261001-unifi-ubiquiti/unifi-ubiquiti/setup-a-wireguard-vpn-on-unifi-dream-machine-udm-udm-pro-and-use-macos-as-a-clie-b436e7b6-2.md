---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/setup-a-wireguard-vpn-on-unifi-dream-machine-udm-udm-pro-and-use-macos-as-a-clie-b436e7b6-2
title: "setup-a-wireguard-vpn-on-unifi-dream-machine-udm-udm-pro-and-use-macos-as-a-clie-b436e7b6"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/setup-a-wireguard-vpn-on-unifi-dream-machine-udm-udm-pro-and-use-macos-as-a-clie-b436e7b6.md
source_anchor: ""
source_lines: [142, 151]
sha256: 2f553fac79c53392150127b76995ad9b846047337ba9ce5266d7cbf2017511ac
---

# setup-a-wireguard-vpn-on-unifi-dream-machine-udm-udm-pro-and-use-macos-as-a-clie-b436e7b6

Finally, you are all setup and your WireGuard server is up and running!
Appendix
If you have connection issues with your server because of the internal routing, try adding these rules to your UDM wireguard config:
PrivateKey = YOUR-UDMP-WIREGUARD-PRIVATE-KEY
ListenPort = 51820
PostUp = iptables -A FORWARD -i %i -j ACCEPT; iptables -A FORWARD -o %i -j ACCEPT; iptables -t nat -A POSTROUTING -o br0 -j MASQUERADE
PostDown = iptables -D FORWARD -i %i -j ACCEPT; iptables -D FORWARD -o %i -j ACCEPT; iptables -t nat -D POSTROUTING -o bro -j MASQUERADE
[Peer]
PublicKey = YOUR-CLIENT-PUBLIC-KEY
AllowedIPs = 10.0.0.3/32
