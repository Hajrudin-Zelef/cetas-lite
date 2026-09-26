---
id: collect-260926-mikrotik/mikrotik/beginner-s-question-about-config-dual-wan-failover-warp
title: "beginner-s-question-about-config-dual-wan-failover-warp"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/beginner-s-question-about-config-dual-wan-failover-warp.md
source_anchor: ""
source_lines: [1, 83]
sha256: c2a7ca3cd52b31616a3c587c7f4b705759e105c39fc464eb82765e23d422165a
---

# beginner-s-question-about-config-dual-wan-failover-warp

@jaclaz you are right, I accidentally removed the firewall rules (any suggestions are more than welcome)

```
 /ip firewall address-list
add address=xxxx/24 list=Trusted_IPs
/ip firewall filter
add action=accept chain=input dst-port=xxxx protocol=udp
add action=drop chain=input comment="Drop Invalid Connections" \
    connection-state=invalid
add action=add-src-to-address-list address-list=port_scanners \
    address-list-timeout=1d chain=input comment="Detect port scanners" \
    protocol=tcp psd=21,3s,3,1
add action=drop chain=input comment="Drop port scanners" src-address-list=\
    port_scanners
add action=accept chain=input comment="Allow Established and Related" \
    connection-state=established,related
add action=accept chain=input comment="Allow Limited ICMP" limit=5,10 \
    protocol=icmp
add action=drop chain=input comment="Drop Excess ICMP" protocol=icmp
add action=accept chain=input comment="Allow WinBox from LAN" dst-port=xxxx \
    protocol=tcp src-address=xxxx
add action=accept chain=input comment="Allow SSH from LAN" dst-port=xxxx \
    protocol=tcp src-address=xxxx
add action=accept chain=input comment="Allow ICMP (Ping) from LAN" protocol=\
    icmp src-address=xxxx
add action=log chain=input comment="Log Unmatched" log-prefix=\
    "Firewall Drop: "
add action=drop chain=input comment="Drop Everything Else (Default Deny)"
/ip firewall mangle
add action=mark-routing chain=prerouting new-routing-mark=To-Cloudflare \
    passthrough=no src-address=xxxx
	
/ip service
set www disabled=yes
set ssh port=xxxx
set winbox port=xxxx
```

And the bridge definition

```
/interface bridge
add name=LAN port-cost-mode=short
add disabled=yes name=WLAN port-cost-mode=short
```

In regard to your recommendation on the interfaces and the use of lists, I will definitely adopt it as it is much cleaner indeed!

@anav Thank you for your comments! Much appreciated!

**1.** Understood

**2.** Understood, although the ““ are coming out from the config.rsc and not something I wrote.

**3.** I am having the VPN temporarily to see how it goes, and haven’t decided if I’ll keep it or disable it. I will definitely note down your recommendation and adjust the config accordingly, depending on what I choose.

**4.** Same as above

**5.** I know the comments seem confusing, but it was the easiest way for me to understand some basic things. There are not all of the Route B, and if you look closely, you’ll see

- ISP1 - Default Route A
- ISP1 - Default Route B
- ISP2 - Default Route A
- ISP2 - Default Route B
- ISP1 - Recursive Route A
- ISP1 - Recursive Route B
- ISP2 - Recursive Route A
ISP2 - Recursive Route B

The reason I have made these many Recursive routes is to cover the case that if a public DNS of the Recursive Route A is down ex. 1.1.1.1 Then check the  public DNS of the Recursive Route B 8.8.8.8 to make sure ISP1 does not have Internet before switching to ISP2 (and I have applied the same logic to ISP2).

Regarding your suggestion “NOT to use or set default routes in IP DHCP client and simply use the manual routes below.” Can you elaborate on this, please?

1. 
Noted and will try it
2. 
Check my reply to Jaclaz above. The DNS I use is the following:

```
/ip dns
set allow-remote-requests=yes servers=\
    1.1.1.1,8.8.8.8,8.8.4.4,9.9.9.9,1.0.0.1,208.67.220.220
```
