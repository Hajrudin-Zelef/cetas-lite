---
id: collect-260926-mikrotik/mikrotik/ipsec-ikev2-tunnel-tuning
title: "ipsec-ikev2-tunnel-tuning"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/ipsec/ipsec-ikev2-tunnel-tuning.md
source_anchor: ""
source_lines: [1, 22]
sha256: f0d7a2f7dac4f223e3a15d90da78d1d6987cd742759a30b3bed5b926c32b2a42
---

# ipsec-ikev2-tunnel-tuning

I think I will try your suggestion and move it to raw.


and

So that the firewall of that (server) side will have more active role, managing both every day routing and IPsec.


are mutually exclusive. If you move the rules necessary to prevent fasttrack from shadowing IPSec policy to raw->prerouting, you effectively prevent the stateful firewall (connection tracking) from working. So you can do that at the “client” side, but not at the “server” side. At server side, keep those rules in filter->forward.

One last question. How to check if tunnel packets are indeed encrypted? Use Wireshark for example?


Yes, tools → packet sniffer → set file name and file size, then start, then generate some supposed-to-be-encrypted traffic, then stop, then download the file and open it with Wireshark. But as @acruhl has pointed out, your picture is not too informative regarding the logical topology of the network. From what you wrote I suppose that you have no VLANs there so the LAN traffic and the PPPoE traffic are mixed up in a single LAN and your switches are not manageable. I assume you have your reasons why you cannot conect the modem to one Ethernet port of Mikrotik and the switch and thus the rest of the LAN to another one.

In this case, I would look for L2 filtering capabilities of the modems and use them to prevent all packets with ethertype different from the two PPPoE ones from being forwarded to the ISP. The point is that otherwise LAN broadcasts packets like ARP packets, SSDP etc. as well as some unicast packets leak out to the uplink which consumes the uplink bandwidth and discloses information about devices in your LAN to the ISP.

If the switches are actually manageable, it may be easier to set these rules at the switches.

But back to the verification that the packets are encrypted - the point is that packet sniffing directly at Mikrotik exhibits some funny behaviour. Timestamps of decrypted received packets are for some reason ahead of those of their carrying IPSec (ESP) packets, and in some cases both the encrypted and decrypted version of the very same packet is captured at the same interface. So once your LAN and WAN logical interfaces share the same media, you have to look carefully at ethertype value (to tell normal IP over Ethernet packets from IP over PPPoE) and at MAC addresses of the sender and recipient. You may find the same packet there three times - first the unpacked one on its way from Mikrotik to the LAN device, with source MAC address of the Mikrotik and destination MAC address of the real recipient, after that the ESP packet encapsulated in PPPoE with source MAC address of your ISP’s PPPoE server and destination MAC address of the Mikrotik, and after that the decrypted version of that packet with same source and destination MAC addresses as before. While the actual order of these packets is 2,3,1 (first the encrypted one coming in from the internet, then the decrypted version of it, and last the decrypted version on its way to the LAN destination).

Edit: oops, @acruhl’s comment below has made me realize that I’ve mixed things up a bit. Yes, I had in mind to capture at the ethernet interface between the Mikrotik and the switch, so you’ll have both the LAN traffic and the PPPoE in the same file, and you’ll see what leaks out in parallel to PPPoE, but in this case, you should see in the capture only ethernet:pppoe:ip:udp:esp packets between the ISP and the Mikrotik and the ethernet:ip:whatever_l4_protocol packets between the Mirotik and the LAN device (talking about IPSec related packets here, you’ll likely see some ntp and dns traffic unencrypted as well). The decrypted version of the IPSec packets still on their way from the modem could be visible only if capturing at the PPPoE interface, where the PPPoE headers are already stripped down. For the IPSec packets, the in-interface is the PPPoE client one, not the etherX through which they actually come in from the PPPoE server.
