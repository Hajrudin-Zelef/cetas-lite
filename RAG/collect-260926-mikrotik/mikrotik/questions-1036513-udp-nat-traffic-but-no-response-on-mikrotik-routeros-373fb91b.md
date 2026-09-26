---
id: collect-260926-mikrotik/mikrotik/questions-1036513-udp-nat-traffic-but-no-response-on-mikrotik-routeros-373fb91b
title: "questions-1036513-udp-nat-traffic-but-no-response-on-mikrotik-routeros-373fb91b"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-1036513-udp-nat-traffic-but-no-response-on-mikrotik-routeros-373fb91b.md
source_anchor: ""
source_lines: [1, 31]
sha256: fe07584f350c4e4b24c3d09f04ac53efdc9b8ad27a381ce26fe21beee9a3b24a
---

# questions-1036513-udp-nat-traffic-but-no-response-on-mikrotik-routeros-373fb91b

I have a MikroTik router with v7.1beta2 firmware installed
It's WAN (eth1) has an IP address of 192.168.7.122
There are two devices connected to its LAN
- Device #1 is a webserver, communicating on port 80TCP192.168.88.254
- Device #2 is a PLC that communicates on port 9999UDP192.168.88.250
I've successfully setup a dst-nat to exposes Device #1 webserver on port 8080 from the WAN.
I cannot, however, get the PLC to communicate through the NAT. I've configured the dstnat similar to the webserver, changing only the port, address and protocol. Here's what I have configured right now:
Chain: dstnat
Protocol: 17 (udp)
Dst. Port: 9999
Action: dst-nat
Log: x
To Addressess: 192.168.88.250
To Ports: 9999
I've disabled all drops on the Firewall.
When I use the communication utility I point it to the WAN address and configured port: 192.168.7.122:9999 and search for the device, the MikroTik RateGraph shows a spike (so it's coming in) but the utility reports the device as 'Missing' (e.g. it's not getting a response).
When I connect to the LAN directly and point to 192.168.88.250:9999 directly the device shows up instantly as 'Available'.
To the best of my knowledge the PLC device doesn't care whether or not the src address is from the local network, as we've had the same model communicating via a NAT in the past (and I don't believe any special treatment was done). Other hardware in the field currently uses a socat of UDP 9999 through a Linux box (not NAT) and that works perfectly fine, so I'd be open to figuring out how to configure a socat-like NAT for testing.
I have also tried to configure a srcnat in case the dstnat wasn't reversing the traffic back through. Here's that:
Chain: srcnat
Src. Address: 192.168.88.250
Protocol: 17 (udp)
Dst. Port: 9999
Action: src-nat
To Addresses: 192.168.7.122
To Ports: 9999
Which, this also doesn't work, and this srcnat does not show any traffic on the Rate Graph.
I'm new to RouterOS, and networking has never been a particular strong suit (I'm a software engineer by trade), so I'm not familiar with ways to properly debug this situation, especially with RouterOS.
Using WireGuard on the host while directly connected to the LAN I see both traffic going out, and then the response.
Using it to monitor via the WAN it goes out but I never see a response.
Help?
