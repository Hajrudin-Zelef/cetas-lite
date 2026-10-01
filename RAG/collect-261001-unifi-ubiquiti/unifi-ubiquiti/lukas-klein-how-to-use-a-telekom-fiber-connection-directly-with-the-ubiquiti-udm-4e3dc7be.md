---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/lukas-klein-how-to-use-a-telekom-fiber-connection-directly-with-the-ubiquiti-udm-4e3dc7be
title: "lukas-klein-how-to-use-a-telekom-fiber-connection-directly-with-the-ubiquiti-udm-4e3dc7be"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/lukas-klein-how-to-use-a-telekom-fiber-connection-directly-with-the-ubiquiti-udm-4e3dc7be.md
source_anchor: ""
source_lines: [1, 15]
sha256: 5ee8c10408c7cc26cf6076b5a41cb8b81b8f1570d7eb176bc1e15a83a800e290
---

# lukas-klein-how-to-use-a-telekom-fiber-connection-directly-with-the-ubiquiti-udm-4e3dc7be

How to use a Telekom fiber connection directly with the Ubiquiti UDM Pro
We recently got a second fiber connection built to our house (the first is from Vodafone who somehow manage to send a DOCSIS cable signal via fiber), this time “real” fiber from Deutsche Telekom. They also sent me their own fiber router, which I would ideally not like to use, since I have a Ubiquiti UDM Pro anyways and would like to skip using another device just as a modem.
Since the UDM Pro has a WAN SFP port, I thought it cannot be that hard to use it directly with the new fiber connection. Turns out, it isn’t, once you know how to set it up.
Here’s what I did as a brief reference for anyone wanting to do this as well:
I got myself a ZYXEL SFP GPON module as well as a LC/APC to SC/UPC fiber cable to connect it. I installed it in my UDM Pro, used the Telekom-provided setup link to enter my modem ID (can be found on the Zyxel module) as well as the Telekom ID that’s written on the Telekom-installed fiber socket and shut the UDM Pro down, unplugged power for a bit (a regular reboot doesn’t power cycle the SFP module) and started the UDM pro again.
Get Lukas Klein’s stories in your inbox
Join Medium for free to get updates from this writer.
In the Unifi settings, I configured WAN1 to be port 10 (the SFP one) and changed the settings to manual. Here’s what I needed to change:
- Enable VLAN, set ID to 7
- set IPv4 Connection to PPPoE
 — Your username is: ”Connection identifier (Anschlusskennung)”+”access number (formerly T-Online Number) (Zugangsnummer (vormals T-Online Nummer))”+”0001"+”@t-online.de”, e.g. 001852485618413361714302#0001@t-online.de
 — Your password is your personal password you should’ve received from Telekom
- I also set my DNS to be static servers (8.8.8.8 and 8.8.4.4 in my case), but that’s up to your preference
I then power-cycled everything and it was working perfectly out of the box.
I’m still trying to get IPv6 working and will update this post once I know more.
