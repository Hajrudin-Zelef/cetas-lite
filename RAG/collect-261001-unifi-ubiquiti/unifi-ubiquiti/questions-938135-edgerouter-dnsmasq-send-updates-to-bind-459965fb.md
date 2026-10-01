---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-938135-edgerouter-dnsmasq-send-updates-to-bind-459965fb
title: "questions-938135-edgerouter-dnsmasq-send-updates-to-bind-459965fb"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-938135-edgerouter-dnsmasq-send-updates-to-bind-459965fb.md
source_anchor: ""
source_lines: [1, 20]
sha256: 7af3cfa323044f4a215af098b9524bc27eb79ba118a4cd378eb859776e41eed5
---

# questions-938135-edgerouter-dnsmasq-send-updates-to-bind-459965fb

TL;DR: How do I configure the Ubiquiti EdgeRouter, running DNSmasq or DHCPd, and NSupdate, to send its DHCP leases to a central BIND server?
I have a system comprised of a few Ubiquiti EdgeRouter X's all acting as DHCP and DNS servers for a local subnet. Inter-subnet routing between ERXs is handled by a core router. What I would like to do is configure the ERX gateways to update a master DNS server, running BIND, whenever they update their DHCP leases.
Referencing the diagram below: With each ERX serving DHCP and DNS for its local subnet, the clients can only resolve DNS names of other clients behind their local ERX. However, if each ERX sent its records to dns-master whenver a DHCP lease update occurred, and each client used dns-master as its DNS server, then any client could resolve the name of any other client. This is what I am trying to achieve.
        router-core 
       /      |     \
      /  dns-master  \
     /                \
 erx-1                erx-2
   |                    |
   |-- client-1a        |-- client-2a
   |-- client-1b        |-- client-2b
What I've done so far:
- BIND is configured to accept updates via nsupdate and I confirmed that it works from my local workstation
- I've installed dnsutils on the ERX's using Apt after activating the Debian repo
- I've explored the dhcp-script option from the Dnsmasq config file (better description of that option in the man page), but I can't get my script to execute
  - Dnsmasq on the EdgeRouter seems to use both /etc/dnsmasq.conf and/etc/dnsmasq.d/dnsmasq-dhcp-config.conf , but adding thedhcp-script option to the former had no effect and adding it to the latter generated a config syntax error
  - I have tried both modifying the Dnsmasq config files directly and setting options in the ERX config under set service dhcp-server global-parameters and neither seemed to work (but I'm also not sure if I did them right)
- Dnsmasq on the EdgeRouter seems to use both 
[EDIT]: The Edgerouter also supports using Dhcpd as a DHCP server, so I'd accept an answer describing how to configure DHCPd on the ER to do this as well; I'm not married to using DNSmasq
router-coreare given a DNS name in the root zone (*.local) while each ERX serves a subdomain (*.sub.local). The BIND server is configured to be authoritative for the entire (*.local) domain
