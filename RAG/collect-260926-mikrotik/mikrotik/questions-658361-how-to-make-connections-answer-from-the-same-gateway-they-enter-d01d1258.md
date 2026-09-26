---
id: collect-260926-mikrotik/mikrotik/questions-658361-how-to-make-connections-answer-from-the-same-gateway-they-enter-d01d1258
title: "This rule is disabled because, when enabled, users cannot browse Internet"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-658361-how-to-make-connections-answer-from-the-same-gateway-they-enter-d01d1258.md
source_anchor: ""
source_lines: [1, 48]
sha256: c59b4607160fd3558df16d082704d9158631278b1945dddb6a41d8528bbe245d
---

# This rule is disabled because, when enabled, users cannot browse Internet

I have a MikroTik RouterOS 6.23 device, and my network is as follows:
Router
  |
  |-- bridge1_LAN (wlan1 + ether1) (192.168.0.210) -- LAN (192.168.0.0/24)
  |   Here is where computers are. Those include some servers and some users.
  |   Users should be able to navigate always, and servers should
  |   be reachable online always.
  |
  |-- ether2_ADSL (192.168.2.2) -- ADSL router (192.168.2.1) -- WAN
  |   Users should navigate through here because there is no traffic limit.
  |   Incoming traffic should work exactly as with ether3_3G, as a temporary
  |   backup solution in case it fails.
  |
  |-- ether3_3G (192.168.3.2) -- 3G router (192.168.3.1) -- WAN
      This connection has a traffic limit, but faster upload rate, so it's
      mainly for incoming traffic. In case ether2_ADSL fails, this should be
      used as a temporary backup connection for outgoing traffic.
Now, the relevant configuration:
/ip firewall mangle
# This rule is disabled because, when enabled, users cannot browse Internet
add action=mark-routing chain=prerouting connection-mark=no-mark disabled=yes \
    in-interface=ether2_ADSL new-routing-mark=to_ether2_ADSL passthrough=no
# This marks all traffic coming from ether3_3G to get out through there too
add action=mark-routing chain=prerouting in-interface=ether3_3G \
    new-routing-mark=to_ether3_3G passthrough=no
/ip firewall nat
add action=masquerade chain=srcnat out-interface=ether2_ADSL
add action=masquerade chain=srcnat out-interface=ether3_3G
# This is just an example web server listening in port 8069, for testing purposes
add action=dst-nat chain=dstnat comment="Test server" dst-port=8069 \
    in-interface=ether2_ADSL protocol=tcp to-addresses=192.168.0.156 \
    to-ports=8069
add action=dst-nat chain=dstnat comment="Test server" dst-port=8069 \
    in-interface=ether3_3G protocol=tcp to-addresses=192.168.0.156 \
    to-ports=8069
/ip route
# Outgoing traffic by routing-mark
add check-gateway=ping distance=10 gateway=192.168.3.1 routing-mark=\
    to_ether3_3G
add check-gateway=ping distance=10 gateway=192.168.2.1 routing-mark=\
    to_ether2_ADSL
# Outgoing traffic by default
add check-gateway=ping distance=20 gateway=192.168.2.1
add check-gateway=ping distance=30 gateway=192.168.3.1
With this configuration, all traffic goes out by ether3_3G only if ether2_ADSL fails, and by ether2_ADSL otherwise (most of the time).
Now the problem is that incoming connections only work through ether2_ADSL. Connections incoming from ether3_3G get always stuck in syn received state.
It seems to me that incoming connections from ether3_3G reach the target server, but the response goes out through ether2_ADSL and that's why the TCP handshake never completes. In fact, if I physically unplug the ether2_ADSL cable, then all connections to/from ether3_3G start working OK.
How can I fix that?
