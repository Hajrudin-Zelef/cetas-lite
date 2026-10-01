---
id: collect-261001-general-networking/general-networking/questions-12752-how-to-configure-internet-port-and-public-ip-settings-on-the-cis-d1cc143e
title: "questions-12752-how-to-configure-internet-port-and-public-ip-settings-on-the-cis-d1cc143e"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-12752-how-to-configure-internet-port-and-public-ip-settings-on-the-cis-d1cc143e.md
source_anchor: ""
source_lines: [1, 17]
sha256: 169af40b0d7d97c2a18a75bfad8663f6b0de7b05eb644119ecc55c8381d2d132
---

# questions-12752-how-to-configure-internet-port-and-public-ip-settings-on-the-cis-d1cc143e

What exactly do the Internet Port and Public IP settings mean on the CISCO Meraki Z1 gateway?
Context
Note: this is a detailed explanation with images - I apologize for its length but I'm striving for quality.
Objective:
I'm setting up the Z1 for my father to connect to a VPN with a cloud-hosted service from his work. It seems like the configuration should be easy, but I'm getting some grief from the tech people at the cloud hosting service but not really getting their help in figuring out the answers.
Network config:
Internet <--> DSL Modem-Router <--> Meraki Z1, Wireless Router, etc.
Problem:
The specific problem I'm having is that I've been told that the Internet Port and the Public IP settings for the Z1 should be the same. I don't think I can set up the Internet Port to match our public IP since that would mean that the actual endpoint on the Internet would be the Z1 - and that's not possible since we need the modem.
My solution so far:
I checked and followed the Meraki Z1 installation documentation:
It tells me that the only necessary WAN configuration for the device is to tell it how to connect to the Internet, which it can do through DHCP or statically (I've set it statically to be sure it's the correct way to do it).
I've also set the public IP address on the configuration page, so that it knows our IP on the Internet.
The current overview on the device looks like this (all green):
Finally
This looks like it's ready to go. So, in summary: have I made an incorrect assumption? Should I do something differently?
Thanks a lot!
