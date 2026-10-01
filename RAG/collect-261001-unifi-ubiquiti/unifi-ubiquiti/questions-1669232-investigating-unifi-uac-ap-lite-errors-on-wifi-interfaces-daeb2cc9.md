---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1669232-investigating-unifi-uac-ap-lite-errors-on-wifi-interfaces-daeb2cc9
title: "questions-1669232-investigating-unifi-uac-ap-lite-errors-on-wifi-interfaces-daeb2cc9"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1669232-investigating-unifi-uac-ap-lite-errors-on-wifi-interfaces-daeb2cc9.md
source_anchor: ""
source_lines: [1, 30]
sha256: 53be135512d61e9fd7fa2aecf8a51a1203d86624c4377563169b7dc295c86fe1
---

# questions-1669232-investigating-unifi-uac-ap-lite-errors-on-wifi-interfaces-daeb2cc9

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
We used Unifi to install Wifi in a house for a customer. Originally we covered everything with 5 UniFi AC PRO versions of the Unifi access points.
There are several VLANs, one for the office/staff, one for the reception one for another company using some rooms in the same building and one for the guests. This ran for two years relatively without issues.
However guests complained, that in the rooms there was always bad Wifi coverage. Therefore we ordered 20 UniFi AC Lite versions and installed one in each room. While that improved the coverage, we now had guests complaining, that they had connection drops and difficulties to connect to the internet at all. They would be able to associate with the AP in their room but then the connection to the internet didn't work.
To investigate the problem, I installed a monitoring system (Check MK) in their network and added all APs via SNMP. There I saw that alot of the VLAN interfaces of the access points suffered lots of errors. These happened seemingly completely random. I since then removed all VLANs from the Lite APs and only left the guest Wifi open. However I still see the errors coming in. Some have a few percent out-errors, some have up to 100%.
It's always the same interface (the one providing 5G Wifi).
What could cause these issues and what could I do to further investigate the problem?
EDIT: All APs, Switches and Controller components are on the latest version. Each IP also has got assigned a static IP in their management-network.
Look at the UniFi control panel. You should find some useful statistics under the Wi-Fi Metrics dashboard:
Anomalies
Association failures - does one failure type appear particularly often?
Access point retry rate - do specific APs have significantly more retries?
Also take a look at connected devices. See if any have particularly poor signal strength. Even consider manually forcing a reconnect (you can do that in the control panel) and see which AP they choose to connect to - is it a better one? This will inform you which steps you should take next.
High-density deployments
Sometimes, having a super-dense deployment is a bad thing. At certain levels, APs start interfering with each other. Also, devices may connect to suboptimal APs and roam incorrectly.
You should consider reducing signal strength as much as possible, to reduce interference and overlap. Ideally each AP does not overlap with more than its immediate neighbours.
You may wish to remove some APs to give them more breathing room.
You may wish to set up a minimum RSSI to kick clients that have a poor connection, in the hopes that they connect to a more appropriate AP. This helps force stubborn clients to roam to a closer AP rather than remaining connected to a poor, far-away, one.
Consider tuning 2.4 GHz transmit power as low as possible, and maybe even turning off the 2.4 GHz radio on adjacent APs, as there are very limited channels (3 non-overlapping, total!) available and 2.4 GHz signals propagate relatively far. Hopefully most client devices you're dealing with support 5 GHz.
Consider reducing bandwidth - a 20 MHz bandwidth is likely sufficient for public internet access at short distance from the AP. This wastes fewer channels than 40 or 80 MHz widths, giving you more chance at avoiding overlap/noise.
Depending on your layout, you may wish to assign channels manually so overlapping APs do not use adjacent (or, worse, same) channels. Auto-channel selection may already take care of this. Use measurement/graphing tools to find out.
All of this is described in the design and cell sizing section of the Ubiquiti document.
Use signal measurement tools
There are various tools available that can measure and graph signal levels visible to a client device. Ubiquiti themselves have WiFiman, available as both an iOS and Android app.
Use one to measure signal strength in each of your locations. Ideally you want one strong signal per non-overlapping channel, so you should tune your signal strengths to achieve this result. Situations where you have either no strong signals, or many strong signals (= noise) are suboptimal and may result in degraded experiences.
