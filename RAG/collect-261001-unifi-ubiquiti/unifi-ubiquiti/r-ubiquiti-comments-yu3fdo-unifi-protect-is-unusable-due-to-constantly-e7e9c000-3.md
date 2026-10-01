---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-yu3fdo-unifi-protect-is-unusable-due-to-constantly-e7e9c000-3
title: "r-ubiquiti-comments-yu3fdo-unifi-protect-is-unusable-due-to-constantly-e7e9c000"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["ethernet", "latency"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-yu3fdo-unifi-protect-is-unusable-due-to-constantly-e7e9c000.md
source_anchor: ""
source_lines: [55, 69]
sha256: a5fff2ec6955ab33da6980f1bef560bb2660db1feebeea61b27a3ab7d3731732
---

# r-ubiquiti-comments-yu3fdo-unifi-protect-is-unusable-due-to-constantly-e7e9c000

Did you bond the two network cables?
Try unplugging one. GBE is plenty for several cameras
I simply have my UNVR ethernet plugged into port 1 on the UDMP. Then I created a new VLAN for my Cameras (192.168.2.0/24) in the Network Settings of my UDMP. Then in the Port Settings of the UDMP I selected port 1 and set it to that Cameras VLAN.
I'll switch to 10GB when I setup more cameras, just wanted to get the ball rolling on 1GB first.
Have you enabled low latency video in the protect settings? How are you accessing it? If your using the UniFi.ui.com relay, IMO it’s not great, and direct access helps a lot with performance.
I am in fact accessing it through https://unifi.ui.com when viewing in Chrome on my Macbook Pro. Otherwise, I'm using the Protect app on my iPhone. I can't seem to directly connect to the UNVR's IP address. I can directly connect to my UDMP's IP. Odd.
And yes I do have "Low Latency Video" and "Insights" checked in the Protect Settings. I don't have "Geofencing" enabled though. I guess I should A/B test with those disabled. Good catch! What in the world does "Low Latency Video" do anyway? Why would you ever want high latency video!? LOL. I found this thread where alot of folks asked with no response -> https://community.ui.com/questions/Protect-Low-Latency-Video-setting-has-anyone-from-Ubiquiti-explained-what-it-does/4825cc24-fe57-4f8c-84bf-8b6eb3e12b57
The 10gb version is explicitly not supported so I am guessing your 16tb version isn't either.
https://help.ui.com/hc/en-us/articles/360037340954-UniFi-HDD-Requirements-and-Compatibility
It is the first thing that popped into my head as I had a hell of a time when I didn't follow the QVL when I tested mine with an SMR drive.
I thought similarly, but for the 10TB it says "Does not fit the drive tray". Mine definitely fits in the tray. Also, I've seen someone in here last week or two post the same drives in their NVR. Neither of which is a clear indicator the drives aren't the issue. You do have a point. I've got a support ticket open with Ubiquiti on this. I will report back!
Running into the same issues also it seems network is also struggling to the point where it crashed yesterday. (Running on a UCK-G2-Plus)
Commentaire supprimé par le membre
Yeah it sucks. I’m just testing with a single 2k G4 Instant. Makes me want to return everything and go Reolink or something else.
What Seagate Skyhawk HDD are you using? Maybe it’s our drives?
