---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-995427-can-i-have-multiple-aps-of-different-network-on-same-unifi-cont-18244d08
title: "questions-995427-can-i-have-multiple-aps-of-different-network-on-same-unifi-cont-18244d08"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-995427-can-i-have-multiple-aps-of-different-network-on-same-unifi-cont-18244d08.md
source_anchor: ""
source_lines: [1, 8]
sha256: 1d94bc08b4cfa0be79346f47a7e59f2ae8eb40c8113900de511c9bd02defe531
---

# questions-995427-can-i-have-multiple-aps-of-different-network-on-same-unifi-cont-18244d08

I have multiple AP's in a big building. But on each floor i run different network for example APn have 192.168.22.n and APz have 192.168.23.z. Can i have all those AP on one unifi controller. It would be a big headache if i have to control each network on different controller. I can switch over to network to access them or ping them, but i can't adopt them every single time.
2 Answers 2
There are actually two good ways to do what you are asking:
- Have multiple "sites" in a single Unifi Controller: If you want each floor of the building to be managed as if it was a separate location but still make it easy to manage them all from a single controller, you can create sites named "First floor", "Second floor", etc. and Adopt the access points from each floor into the site for that floor. This is the option to choose if you want, for example, to have additional administrators for each floor who cannot access other floors. Each administrator can access a single site, but you (and any other "Super Admin") can access all of the sites using the dropdown menu on the controller website (or in the mobile app). This also keeps the statistics (such as users connected, bandwidth used, etc.) for each floor separate. I have a single controller with 12 different sites in it, and I know people who have controllers with 50 sites or more in a single controller.
- Create "Wi-Fi Network Groups" (in a single site) and assign access points to them: If you want all of the Access Points to be in a single "site", you can create a separate "Wi-Fi Network Group" for each floor, named "First floor", "Second floor", etc. In each group, you create the Wi-Fi network(s) that will be available in that group. Then, in the access point settings, you can assign the access point to a group.
If I understand the question correctly, yes. You can configure multiple APs to all have multiple SSIDs (and trunk and retain separate VLANS for the different SSIDs)
- 
        Can i configure multiple APs on one Unifi Controller of different networks?Adeelhashmi145– Adeelhashmi1452019-12-13 16:04:48 +00:00Commented Dec 13, 2019 at 16:04
