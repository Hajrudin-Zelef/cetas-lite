---
id: collect-261001-meraki/meraki/r-meraki-comments-smtytr-issues-with-connectivity-to-meraki-wireless-mr46-7a328301
title: "r-meraki-comments-smtytr-issues-with-connectivity-to-meraki-wireless-mr46-7a328301"
domain: meraki
role: reference
task: reference
actors: ["Intel", "Qualcomm"]
dates: []
keywords: ["intel"]
source: docs/RAG/collect-261001-meraki/r-meraki-comments-smtytr-issues-with-connectivity-to-meraki-wireless-mr46-7a328301.md
source_anchor: ""
source_lines: [1, 20]
sha256: 368763b2a84a8eb4947a752d6457af67dae3d3142d0767782a4ab656582a27e8
---

# r-meraki-comments-smtytr-issues-with-connectivity-to-meraki-wireless-mr46-7a328301

Issues with connectivity to Meraki Wireless MR46 and MR46e 
        
    We recently switched from Cisco 2602 WAPs to Meraki MR46 and MR46e devices.
We are seeing intermittent issues with connectivity by a variety of devices (Panasonic laptop, HP desktop, android, windows imbedded.)
After watching the dashboard and monitoring a device that was having connectivity issues I noticed that dropped packets started either when the device switched bands or APs. The device will believe it has a connection, but it can't even reach intranet devices.
I've got an open case with Meraki support but after being on a call with them for 2 hours yesterday and the engineer saying he: wasn't sure what was happening, didn't have anyone to escalate to, team lead wasn't responsive, and he would give me a call back, and never receiving a call back.. I'm not confident they're going to be helpful.
We swapped the AP that had the most problematic issues and it seemed to help at first but today I'm seeing issues again. I've now switched from WPA3 transition to WPA2 just to remove a possible variable. Looking at disabling 802.11w next.
Does anyone have any advice on things to look for or has anyone experienced a similar issue and found a solution?
Section des commentaires
I've switched from 80hz to 40hz bandwidth and it seems to have helped. I don't see any drops in the last 40 minutes and ping times have less jitter.. I'll report in again if I see more failures.
Hmm, not passing traffic after roaming sounds vaguely familiar. I remember reading this bug either in a client driver (Intel wireless) or in another Qualcomm-powered WiFi vendor’s release notes.
Is your firmware up to date on your APs and also are the wireless drivers up to date on the laptops?
Do you have 802.11r fast transitions enabled?
802.11r disabled.
Devices have various different NICs we've tried updating them all.
Roaming results are not consistent, sometimes it's fine then suddenly 90% packet loss and the device believes its connected to the AP but it seems to be after switching APs or bands.
We enabled band steering to try to reduce this behavior but it's difficult to say how much impact that's had since the issue has always been intermittent.
We are having the same problems on MR46 AP's. Just like you we called support and got no where. Meraki will not admit anything is wrong. We ran fine on 27.7.1 or 27.1.1 forgot the code. But as soon as we went to 28. we had this trouble. Wishing we never moved now... :/
Are you on 80hz? We eliminated maybe 90% of our disassociations by changing to 40hz. I'm not really sure if it's expected we needed to do so as there's not many wifi signals in our area.
Still not happy with the status but at least production continues
