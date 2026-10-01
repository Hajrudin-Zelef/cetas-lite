---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-51355-unifi-usg-gateway-de6488a1
title: "questions-51355-unifi-usg-gateway-de6488a1"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-51355-unifi-usg-gateway-de6488a1.md
source_anchor: ""
source_lines: [1, 12]
sha256: b5a39a103c4769968b418d064b11f30149defafb41ea1d5f8c397aedead584aa
---

# questions-51355-unifi-usg-gateway-de6488a1

On my firewall "gate protect", I have set one network to use DHCP with the subnet 172.25.112.0/24.
After I connect it on my USG pro as WAN and the LAN port goes on the ubnt POE switch there is connect my cloud-key and some APs.
FW x.x.112.1/24 "DHCP-Server" 
USG Pro x.x.112.5 as router "x.x.112.1"
Switch POW x.x.112.3
Cloudkex x.x.112.5
AP- x.x.112.6-15
Wificlients <15
I have always error in my network the clients finish in a services zone "169.x.x.x"
USG get 172.25.112.4/24 and as router my FW "172.25.112.1"
On my controller I set a IP range " 172.25.112.1/24 but I'm sure about DHCP options "server, relay, or no"
What I need is my clients on the 112.x network.
