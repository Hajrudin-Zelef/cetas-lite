---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1056888-opnsense-wan-failover-causes-disruption-when-non-active-wan-is-9a93c71c
title: "OPNsense WAN failover causes disruption when non-active WAN is down"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1056888-opnsense-wan-failover-causes-disruption-when-non-active-wan-is-9a93c71c.md
source_anchor: ""
source_lines: [1, 41]
sha256: 3d2bf9abe0b4983e63d48491213dc1cc990f1ab4f062aacf8eddc21a2ed45f93
---

# OPNsense WAN failover causes disruption when non-active WAN is down

*Score : 3 | Source : https://serverfault.com/questions/1056888/opnsense-wan-failover-causes-disruption-when-non-active-wan-is-down*

I have the latest version of OPNsense set up in a VM on ESXi 7. OPNsense is very similar to pfSense, and I suspect the solution would apply to both. All the NICs are PCI passthrough devices:
WAN 1 and WAN 2 are set up as gateways. WAN 1, being the connection I want to use all the time unless it's down or has high loss, has a priority of 254, and WAN 2 is 255, thus preferring WAN 1. There are no other enabled gateways. Both are upstream gateways. There are no ping target IPs so instead they ping their respective gateway IPs as provided by the ISPs.
I then made a gateway group with both the WANs as members. Tier 1 has WAN 1 only, and Tier 2 has WAN 2 only. The "Trigger Level" is "Packet Loss or High Latency." My understanding is having two tiers like this means Tier 1 is used 100% of the time unless it fails, then Tier 2 is used 100% of the time until Tier 1 comes back.
Under normal circumstances with both WANs up and reporting no loss, this works fine. All traffic seems to be sent through WAN 1 and WAN 2 is ignored, aside from health checks and the firewall blocking random script kiddies probing it.
However, when WAN 2 starts to report packet loss and OPNsense considers it to be down, which happens usually a handful of times per day for a few minutes at a time, it is disruptive despite WAN 1 being healthy. Some devices that try to get to the internet suddenly can't until WAN 2 is healthy again. Alternatively, if I just disable the WAN 2 gateway entirely, the connection is stable, presuming of course WAN 1 also stays up. When this happens, OPNsense is still saying WAN 1 is healthy.
The logs suggest that the routes aren't changing at all from WAN 1 when WAN 2 fails, which makes this more confusing. Here are log entries around the time of a WAN 2 failure:
2021-03-10T10:31:44 kernel  pflog0: promiscuous mode disabled    
2021-03-10T10:31:44 opnsense[38519] /usr/local/etc/rc.filter_configure: ROUTING: keeping current default gateway '(WAN 1's IP)'  
2021-03-10T10:31:44 opnsense[41442] /system_gateways.php: The LAN_DHCP monitor address is empty, skipping.   
2021-03-10T10:31:44 opnsense[41442] /system_gateways.php: Choose to bind WAN 1 on (WAN 1's IP) since we could not find a proper match.   
2021-03-10T10:31:44 opnsense[41442] plugins_configure monitor (execute task : dpinger_configure_do())    
2021-03-10T10:31:44 opnsense[41442] plugins_configure monitor ()     
2021-03-10T10:31:44 opnsense[41442] /system_gateways.php: ROUTING: keeping current default gateway '(WAN 1's IP)'    
2021-03-10T10:31:44 opnsense[41442] /system_gateways.php: ROUTING: setting IPv4 default route to (WAN 1's IP)    
2021-03-10T10:31:44 opnsense[41442] /system_gateways.php: ROUTING: IPv4 default gateway set to opt2  
2021-03-10T10:31:44 opnsense[41442] /system_gateways.php: ROUTING: entering configure using defaults     
2021-03-10T10:31:43 configctl[59484]    event @ 1615397503.05 exec: system event config_changed  
2021-03-10T10:31:43 configctl[59484]    event @ 1615397503.05 msg: Mar 10 10:31:43 OPNsense.localdomain config[41442]: config-event: new_config /conf/backup/config-1615397503.052.xml
And here are gateway log entries around the same time:
2021-03-10T10:31:03 dpinger[59862]  GATEWAY ALARM: WAN 2 (Addr: ******* Alarm: 1 RTT: 8702us RTTd: 7796us Loss: 22%)     
2021-03-10T10:31:03 dpinger[77013]  WAN 2 *******: Alarm latency 8702us stddev 7796us loss 22%   
2021-03-10T10:30:05 dpinger[79106]  GATEWAY ALARM: WAN 2 (Addr: ******* Alarm: 0 RTT: 9164us RTTd: 14634us Loss: 11%)
The most confusing entry to me is keeping current default gateway which to me explicitly is saying "WAN 2 failed but I don't care because WAN 1 is the default gateway", but maybe I'm misinterpreting it.
So, in this dual WAN failover scenario, why does the failover of the second, supposedly inactive WAN cause the active WAN to be unusable?

---

### Reponse — score 2

Try going into firewall>settings>advanced and tick "Disable State Killing on Gateway Failure".
I think I had the same issue as you. I can only conclude that a gateway failure was killing all states of all gateways - in your case killing the states of a number of stateful connections established via WAN1 despite the fact that it was WAN2 that went down.

---

### Reponse — score 1

I had similar problem and solve this problem by going to System --> Settings --> General and enabling "Allow dedault gateway switching" it goes to another gateway if default gateway is the problem.
