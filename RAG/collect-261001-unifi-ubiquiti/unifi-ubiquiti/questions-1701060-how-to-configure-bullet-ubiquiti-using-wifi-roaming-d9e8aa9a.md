---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1701060-how-to-configure-bullet-ubiquiti-using-wifi-roaming-d9e8aa9a
title: "questions-1701060-how-to-configure-bullet-ubiquiti-using-wifi-roaming-d9e8aa9a"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1701060-how-to-configure-bullet-ubiquiti-using-wifi-roaming-d9e8aa9a.md
source_anchor: ""
source_lines: [1, 15]
sha256: 1bdf2ce9f1b4ad436c133e1d4a6294ef8406e55e3bc80b36a286c81e756a99db
---

# questions-1701060-how-to-configure-bullet-ubiquiti-using-wifi-roaming-d9e8aa9a

I am using a Bullet 2HP and a Bullet M2. The 2HP has the following options for wireless mode:
- Access Point
- Access Point WDS
- Station
- Station WDS
The M2 has the following options for wireless mode:
- Access Point
- Station
- AP repeater
What I want to accomplish is both bullets working as Access Points, and I'd like to be able to do wifi roaming with them. Both bullets do not need to have access to the world wide web (Internet), they use static IP addresses, and while moving a laptop connected to one access point, the laptop should switch to the strongest signal seamlessly, without the user doing anything manually.
Furthermore, the laptop should be able to connect to any device connected to any of the two bullet's Access Points.
I thought I would accomplish this wifi roaming, by using the same SSID and password in both Access Points and configuring the bullet M2 as a AP-Reapeater and the bullet 2HP as an Access Point WDS. But this is not working.
Here are the configs I am using for the bullet 2HP:
Here are the configs I am using for the bullet M2:
A device connected to one bullet's Acess Point can not connect to the device connected in the other Acess Point (it can not even ping the other bullet's IP address). What am I doing wrong?
