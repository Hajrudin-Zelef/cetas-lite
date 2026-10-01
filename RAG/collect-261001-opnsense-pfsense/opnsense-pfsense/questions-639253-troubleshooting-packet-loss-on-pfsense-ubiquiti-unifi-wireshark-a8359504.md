---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-639253-troubleshooting-packet-loss-on-pfsense-ubiquiti-unifi-wireshark-a8359504
title: "questions-639253-troubleshooting-packet-loss-on-pfsense-ubiquiti-unifi-wireshark-a8359504"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel", "memory"]
source: docs/RAG/collect-261001-opnsense-pfsense/questions-639253-troubleshooting-packet-loss-on-pfsense-ubiquiti-unifi-wireshark-a8359504.md
source_anchor: ""
source_lines: [1, 23]
sha256: 4293061d292818c9ecae79a5c090a9b979b73c2deb9eb28c3dc296f099463b41
---

# questions-639253-troubleshooting-packet-loss-on-pfsense-ubiquiti-unifi-wireshark-a8359504

I'm dealing with an annoying situation in a small network in my church, which I'm the primary volunteer IT caretaker, of about 20 PCs, give-or-take.
We're in Chattanooga, home of Gigabit internet, so we have plenty of bandwidth (100mb connection).
The pfSense hardware is, according to the pfSense dashboard:
Intel(R) Atom(TM) CPU D525 @ 1.80GHz
4 CPUs
Both NIC's (WAN + LAN) are gigabit ports. This thing has 2GB RAM.
We have a computer lab / after school tutoring program, so I'm using pfSense for content filtering with Squid and Squidguard.
A week & a half ago, unbeknownst to me, another IT guy came in and rearranged a bunch of IT equipment and mounted some things to the wall in the network closet without talking to me first.
That happened to be the same weekend a big storm blew through town.
Since then, internet has been spotty. Multiple times throughout the day, internet begins to slow down, and keeps slowing down until its unusable, and then most (if not all) of the folks report that its downright broken with no access to the outside world.
As I'm not on-site very often, its hard for me to actually troubleshoot the problem when it actually happens. The solution (which I don't really like, but it gets the job doen) has been to pull the power from everything in the network closet (pfSense + 1 of the Ubiquiti APs, the Cisco SG-100, and the ISP's equipment), plug everything back in, and everything comes back up at full speed.
However, at times that I have been able to be on-site, I've noticed that I'm unable to ping the gateway (pfSense) whenever internet goes down (10.0.0.1), while I am able to ping other internal devices, such as the printer located at 10.0.0.2.
Reviewing the pfSense dashboard, I've never seen traffic become saturated. We have a 100mb connection, so have plenty of bandwidth. No servers and no high-bandwidth applications are on-site.
Symptoms to me sound exactly like a Spanning Tree issue (we don't have any smart switches, although I do have a Cisco SG-100 at the core of the network.
I checked all our switches (we only have 3 throughout the building - none with more than 8 ports), and hand-traced all of the cables to make sure there are no physical loops, and make sure switches aren't plugged into each other multiple times.
So then I upgraded the pfSense hardware from 2.1.3 to 2.1.5, and upgraded the firmware on all 4 of our Ubiquiti UniFi wireless APs. I also didn't have a wireless controller running continuously, so I installed the software onto one of the staff PCs that is almost always on, so that the controller remains present.
(If you know anything about Ubiquiti UniFi, you don't have to have the controller running continuously, but I figured it wouldn't hurt)
Running a lot of pings earlier today from my own PC (Ubuntu) when the internet was slow, I saw a LOT of packet losses. I noticed that as I would run a ping to a particular external IP address, there would be a lot of packet losses at the beginning, but the longer I let ping run, the faster ping responses would be (and the more consistent / reliable).
Reviewing the Proxy Filter config on the firewall, I noticed in the Cache Management of the Proxy Server section that the Memory Cache Size was 32mb while the Maximum Object Size in RAM was set to 64mb. Realizing this could cause a problem, I increased the Memory Cache Size up to 256mb, and I turned off completely the Hard Disk Caching.
I'm hoping that'll help, but we'll watch the network over the next 24-48 hours or so.
(Update: This didn't seem to help. 5 minutes after I left, I got a call that the internet was down. So I came back and swapped out the pfSense device with a temporary Cisco Linksys router, and we'll see what happens).
Are there any other suggestions or things I should look into in troubleshooting this ongoing issue? 1 thought I do have is that the guy who moved all of the network equipment without asking me first could have pinched a cable. I replaced the cable from the pfSense device going to LAN, but that didn't help. Another thought I have is that there could have been a surge of some sort because of the storm, but everything in the network closet is behind a APC Surge Protector. Regardless, that was when the issues started.
I have WireShark, but I'm not entirely sure what to look for in a Packet Capture. Perhaps some pointers on what to do with a packet capture would be helpful too.
