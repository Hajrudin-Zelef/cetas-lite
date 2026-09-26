---
id: collect-260926-mikrotik/mikrotik/questions-919478-routeros-push-static-route-for-ipsec-ikev2-vpn-9b7e3dfc
title: "questions-919478-routeros-push-static-route-for-ipsec-ikev2-vpn-9b7e3dfc"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-919478-routeros-push-static-route-for-ipsec-ikev2-vpn-9b7e3dfc.md
source_anchor: ""
source_lines: [1, 12]
sha256: 770d5fde9fd2cea9d5ff358c7927ddefb1b491401fd672aaeb145d2cc8d7131f
---

# questions-919478-routeros-push-static-route-for-ipsec-ikev2-vpn-9b7e3dfc

Can't find a solution how actually I can push a static routes to VPN clients, when they connect via VPN.
Configuration from here
RouterOS, IPSec, IKEv2. Clients mainly macOS users via standard soft.
Any ideas with examples are appreciate.
The ipsec mode-config is supposed to provide such functionality.
For example if you have the following settings in RouterOS:
/ip pool add name="ipsec_pool" ranges=192.168.50.2-192.168.50.6
/ip ipsec mode-config add name="windows" system-dns=no static-dns=192.168.88.1 address-pool=ipsec_pool address-prefix-length=29 split-include=192.168.88.0/24
/ip ipsec peer add address=0.0.0.0/0 passive=yes auth-method=rsa-signature certificate=ipsec-server-03 generate-policy=port-strict policy-template-group=win-ikev2 exchange-mode=ike2 mode-config=windows send-initial-contact=no hash-algorithm=sha1 enc-algorithm=aes-256,aes-128 dh-group=ecp256,ecp384,modp2048,modp1024 lifetime=2h dpd-interval=2m
then RouterOS will assign a virtual ip from the address pool called "ipsec_pool" to the client and tells it to set the DNS server address to 192.168.88.1 (static-dns) and create policies to route all traffic to 192.168.88.0/24 thru IPSec (split-include) tunnel.
However it's up to the client implementation what is going to happen with these mode-config requests. For example: because Windows is setting up a point-to-point VPN connection when using IKEv2, it processes the DNS server setting request but ignores the split-include settings. But it creates a new routing rule to subnet 192.168.50.0/24 regardless what the server is telling.
Unfortunately I'm not familiar with the IPSec client in MacOS so I cannot help you with that. It might be much more flexible with mode-configs.
