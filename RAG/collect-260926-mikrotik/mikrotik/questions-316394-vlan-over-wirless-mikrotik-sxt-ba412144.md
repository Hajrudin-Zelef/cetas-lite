---
id: collect-260926-mikrotik/mikrotik/questions-316394-vlan-over-wirless-mikrotik-sxt-ba412144
title: "questions-316394-vlan-over-wirless-mikrotik-sxt-ba412144"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-316394-vlan-over-wirless-mikrotik-sxt-ba412144.md
source_anchor: ""
source_lines: [1, 22]
sha256: a9b8f56fc2d4e59bce4a3ddfe956a9934015573892a6123f40c6829337854840
---

# questions-316394-vlan-over-wirless-mikrotik-sxt-ba412144

I stumbled across a big problem, when I tried to make a VLAN work over a wireless connection between two MikroTik SXTs, and a Wireless AP at the end of this connection.
The setup is like that:
DHCP-Server ------- SXT ====== SXT ---- AP
LAN  ---
WLAN ===
In this Setup multiple Access Points should be serverd by a single DHCP Server and are interconnected through a VLAN backbone. Without the Wireless MikroTek Bridge, by only using patch cables, everything is fine, but when I use the wireless bridge, the DHCP Reply is blocked at any of the MicroTek devices.
First I thought it was the missing VLAN configuration on the MicroTeks, but when I add tmen nothing changed, I used the following configuration on both mikroteks:
interface vlan add name=vlan1 interface=bridge vlan-id=1001 disable=no
ip address add address=192.168.160.12/24 interface=vlan0
The bridge and wireless config is like follows (just for documentation):
system reset-configuration
interface bridge add name=bridge
interface bridge port add bridge=bridge interface=wlan1
interface bridge port add bridge=bridge interface=ether1 
ip address add interface=bridge address=192.168.151.14/24
ip address edit number=0 value-name=address
  192.168.88.1/24 -> 192.168.88.2/24
interface wireless set country="germany 5.8 fixed p-p" mode=station-pseudobridge ssid=XXX numbers=wlan1
interface wireless security-profiles add name=p1 authentication-types=wpa2-psk wpa2-pre-shared-key=xxx
interface wireless set security-profile=p1 numbers=wlan1
interface wireless enable numbers=wlan1
Has anyone an idea what I could have missed? Thanks!
