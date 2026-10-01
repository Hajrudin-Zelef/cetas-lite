---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-41541-ubiquiti-unifi-wireless-bridging-over-two-floors-8f2fd80e
title: "questions-41541-ubiquiti-unifi-wireless-bridging-over-two-floors-8f2fd80e"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-41541-ubiquiti-unifi-wireless-bridging-over-two-floors-8f2fd80e.md
source_anchor: ""
source_lines: [1, 15]
sha256: 3e23217cd581b85caa159705161f0510af9005fd0ac38132be04d7b0ef2fae88
---

# questions-41541-ubiquiti-unifi-wireless-bridging-over-two-floors-8f2fd80e

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
The short answer is no. There is no way to configure an Ubiquiti WAP in Unifi to act as a client. The main gigabit port is to be connected to the network infrastructure, the secondary port is just a pass-through port, and the wireless component acts as an access point as it's configured via Unifi.
Even when you SSH into these devices about the only configuration change you can make to them is to set the which Unifi server they should contact for configuration information.
If your desired physical layer diagram is:
Router <-wire-> Device <-radio-> Device <-wire-> PC
Then you're looking for transparent bridging.
The appropriate inexpensive Ubiquiti devices for that are the Ubiquiti Airmax line; you can get all AC (NOT compatible with standard 802.11ac) or all M (802.11n compatible) equipment for ideal matches; or, with an AC line as the access point, M devices can act at stations if the AP is configured that way. Definitely choose the 5Ghz options.
I personally use elements of this product line to punch through walls and floors indoors. Note that the narrower the beam/higher the gain, the more sensitive they are to aim.
Make sure to crank the power way down if at all possible. The newest generation Airmax devices do provide spectrum analysis tools to help you pick the best frequency and bandwidth.
You can set up most (recent) Unify APs to act as clients now via their Wireless Uplink feature. The linked article walks through how to do it, and the supported hardware. I have done it with an old UAP-AC-lite acting as a client with a UAP-AC-PRO as the uplink AP.
I know this is an old question, but tech changes, and this is still a high google result.
I can confirm that the AP wireless up-link feature will allow what you're trying to attempt. You should know that the ports on the AP are bridged to the port it's up-linking to.
I have(all unifi):
USG-4p - wire - 16portPOE switch - wire - AP-HD - Radio - AP-AC-Pro - WireToPoeIn - 8portPOE switch - wired entertainment devices
The 8port switch lets me set a vlan for each of the devices, where as before everything would show as if it was directly connected to the AP-HD's port on the switch using the ALL/native vlan. I also overrode the 5ghz wlans on the AP-AC-Pro so they would not broadcast, but left the 2.4ghz on. Get ~350/400mbs speed test from a device plugged into the 8port.
