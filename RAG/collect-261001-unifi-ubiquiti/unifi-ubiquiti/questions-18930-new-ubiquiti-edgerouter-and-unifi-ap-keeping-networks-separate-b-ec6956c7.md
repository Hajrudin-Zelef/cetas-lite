---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-18930-new-ubiquiti-edgerouter-and-unifi-ap-keeping-networks-separate-b-ec6956c7
title: "questions-18930-new-ubiquiti-edgerouter-and-unifi-ap-keeping-networks-separate-b-ec6956c7"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-18930-new-ubiquiti-edgerouter-and-unifi-ap-keeping-networks-separate-b-ec6956c7.md
source_anchor: ""
source_lines: [1, 51]
sha256: fb3451ffece4e83653e0dcdc1b9a8b72f09a2beb71e4e473c16790e88b853f75
---

# questions-18930-new-ubiquiti-edgerouter-and-unifi-ap-keeping-networks-separate-b-ec6956c7

I've got a Ubiquiti EdgeRouter PoE (5 port) which I'm setting up just now. I'm trying to achieve a solution where the single UniFi AP can hand out IP addresses in the 10.0.2.0/23 space, while devices in the office which require a cable connection can get IPs in 10.0.0.0/23. So far I have set up:
- Port 0 with address 192.168.1.1, patched directly to a port in the office so that I can still log into manage the router regardless of the other port settings.
- Port 1 connected to the internet line using DHCP
- Port 2, "LAN", configured with an address of 10.0.0.1/23
- Port 3, "WLAN", configured with an address of 10.0.2.1/23
I've set up two DHCP pools - one for 10.0.0.0/23, which correctly hands out IPs when you connected to the LAN port, and one for 10.0.2.0/23, which also correctly hands out IPs when you connect to the WLAN port.
The LAN port is connected to a switch, managed by the company above us who own the racks.
So now, I've connected the AP to the WLAN port, and I can see in the EdgeRouter UI that it has obtained a DHCP lease for 10.0.1.10 (yes, I would prefer to manually assign this outside of the DHCP range, but I don't think it affects this). However, when I run the UniFi controller on my laptop, connected via the switch to the LAN port, it can't see the AP.
How can I get the controller to see the AP, keep the LAN and WLAN DHCP pools separate, and still maintain routing between the two? I'm reading a lot about UniFi having issues when not connected to the same switch, which it isn't, but I don't see anything concrete about how to solve the problem! I'm wondering if I need to connect everything through the switch and use VLANs, but I've not used VLANs before and I don't know if that would just amount to the same setup!
Edit:
I've actually made some progess, in that I went with Jim G's "Option 3", and I've at least got some WiFi working. I have set up two SSIDs, one which applies the VLAN tag and one which leaves the traffic untagged. Untagged traffic through the AP successfully gets an IP in the 10.0.0.0 network and can access the internet. However traffic on the tagged SSID, which is the setup I really want, sees the same problem - an IP is handed out but there is no network access.
Some investigation with ping led me to see that initially after the DHCP lease is assigned, I can ping the router and use the internet on the VLAN. However the response times from ping are around 5-7 seconds, and then they start timing out and only come back occasionally (but they do come back for the odd response). Given that they do come back, the timing of just having connected may be insignificant though.
Edit - here's the ping times and pattern:
PING 10.0.2.1 (10.0.2.1): 56 data bytes
Request timeout for icmp_seq 0
Request timeout for icmp_seq 1
Request timeout for icmp_seq 2
Request timeout for icmp_seq 3
Request timeout for icmp_seq 4
64 bytes from 10.0.2.1: icmp_seq=0 ttl=64 time=5696.117 ms
64 bytes from 10.0.2.1: icmp_seq=1 ttl=64 time=4987.305 ms
Request timeout for icmp_seq 7
Request timeout for icmp_seq 8
64 bytes from 10.0.2.1: icmp_seq=2 ttl=64 time=7390.786 ms
64 bytes from 10.0.2.1: icmp_seq=3 ttl=64 time=6985.669 ms
64 bytes from 10.0.2.1: icmp_seq=4 ttl=64 time=6236.806 ms
64 bytes from 10.0.2.1: icmp_seq=5 ttl=64 time=5840.661 ms
64 bytes from 10.0.2.1: icmp_seq=6 ttl=64 time=5108.968 ms
64 bytes from 10.0.2.1: icmp_seq=7 ttl=64 time=4186.371 ms
64 bytes from 10.0.2.1: icmp_seq=8 ttl=64 time=3535.087 ms
64 bytes from 10.0.2.1: icmp_seq=9 ttl=64 time=2627.659 ms
64 bytes from 10.0.2.1: icmp_seq=10 ttl=64 time=1683.275 ms
64 bytes from 10.0.2.1: icmp_seq=11 ttl=64 time=776.689 ms
64 bytes from 10.0.2.1: icmp_seq=18 ttl=64 time=1878.685 ms
Request timeout for icmp_seq 20
Request timeout for icmp_seq 21
Request timeout for icmp_seq 22
Request timeout for icmp_seq 23
Request timeout for icmp_seq 24
Request timeout for icmp_seq 25
Request timeout for icmp_seq 26
Request timeout for icmp_seq 27
Request timeout for icmp_seq 28
Request timeout for icmp_seq 29
Request timeout for icmp_seq 30
Request timeout for icmp_seq 31
Request timeout for icmp_seq 32
^C
--- 10.0.2.1 ping statistics ---
34 packets transmitted, 13 packets received, 61.8% packet loss
round-trip min/avg/max/stddev = 776.689/4379.544/7390.786/2045.686 ms
