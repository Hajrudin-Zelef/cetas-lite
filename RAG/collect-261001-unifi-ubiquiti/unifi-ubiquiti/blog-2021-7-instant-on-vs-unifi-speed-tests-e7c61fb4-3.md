---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-2021-7-instant-on-vs-unifi-speed-tests-e7c61fb4-3
title: "blog-2021-7-instant-on-vs-unifi-speed-tests-e7c61fb4"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple", "Intel"]
dates: []
keywords: ["ethernet", "intel", "research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-2021-7-instant-on-vs-unifi-speed-tests-e7c61fb4.md
source_anchor: ""
source_lines: [63, 98]
sha256: 733ea0b16ba14ba38f6eb1252f5202bc910b01ccdf67049d224761334b4f1938
---

# blog-2021-7-instant-on-vs-unifi-speed-tests-e7c61fb4

The range difference between the AP22/U6-Lite and the U6-LR is big enough that it could be the difference between needing one or two APs to cover an area. In some situations one U6-LR is better, in others one or two U6-Lite/AP22s may be better. It’s hard to make general conclusions.
In my house, one U6-LR on each floor is more than I need. With the AP22 I would still stick to one per floor, but I would consider a 2nd on the main floor for more consistent 5 GHz coverage. Rather than having one AP22 centrally located, I could have two AP22’s on either side. There wouldn’t be much difference in overall performance, so that decision comes down to where you have Ethernet cabling, and how much you care about 5 GHz coverage and speed. For larger homes or coverage extending outdoors, two APs on either side of the house would be the better option.
If you have a lower-power AP now and are considering upgrading, the AP22 should offer similar coverage. The U6-LR would likely be a step up in range, competing with all but the highest-power mesh kits.
While these test results show the difference in performance, the bigger difference is software. Fully comparing Instant On vs. UniFi software is a big enough topic it deserves it’s own article, but I’ll briefly go over the biggest differences and what each is good for.
Instant On software is reliable, cloud-only (besides switches), and basic. You can only choose between a few switches and APs. There are no dedicated Instant On routers, or anything else. The cloud portal is good enough for basic tasks, but doesn’t offer anywhere near the depth of what UniFi does. This is either good or bad, depending on what you want. Instant On does the basics well, but that’s about it.
UniFi software offers more features and is more flexible, but it is also buggier. Sometimes the bugs are harmless, sometimes they are funny, but they are almost always there. If you can deal with them and don’t mind doing a little research before installing an update, UniFi is a unique and flexible ecosystem. UniFi routers are generally basic and miss some important features, but UniFi switches and APs match up well against Instant On. Instant On doesn’t have anything like UniFi Protect, Access, or Talk.
AP Recommendations
If you’re considering a new Instant On network, I have no issues recommending the AP22. The AP12 or AP15 should be better for high-density networks, but I still think the AP22 is the best overall. When the Wi-Fi 6 replacements for the AP12 and AP15 arrive, I’ll be testing those. I’m also excited to test out the UniFi U6-Pro, which is currently in early access.
If you’re considered a new UniFi network and they’re in stock, the U6-Lite and U6-LR are both good options. Their crummy 2.4 GHz radios are a bummer, but for most people that’s not relevant. Fast Wi-Fi requires 5 GHz, and a nearby AP. The LR lives up to it’s Long Range claim, but the U6-Lite is good enough for smaller areas, or part of a multi-AP network. You’ll have the decide for yourself which is the better option.
Beyond Wi-Fi 6, there’s also the promise of Wi-Fi 6E on the horizon, which is a more meaningful upgrade than the 10-20% speed improvement you can expect from upgrading from a good Wi-Fi 5 AP. Wi-Fi 6 and 6E are only relevant when you have clients that support them. Wi-Fi 6E devices are just starting to roll out in 2021, and it will be a while until it’s common for most home users.
If you are happy with your Wi-Fi network, it could be a good idea to hold off on upgrading. If you’re looking for an upgrade now, Instant On and UniFi both have good options.
Network Equipment and Firmware Versions
- All Instant On APs used the latest firmware 2.3.1
- UniFi Dream Machine, running firmware version 1.10.0 
  - UniFi Network Controller version 6.2.26
  - All UniFi settings at defaults, besides channel width and transmit power. Wi-Fi AI was disabled.
- UniFi 6 Lite, firmware version 5.60.9
- UniFi 6 Long Range, firmware version 5.60.9
- UniFi AC-Pro, firmware version 4.3.28
- UniFi AC-HD, firmware version 4.3.28
- UniFi Switch Lite 8 PoE, firmware version 5.43.23
- iPerf server: Qotom mini desktop running pfSense, or Mac Mini connected via Ethernet
Wireless Clients
- Windows 10 PC with Intel AX200 Wi-Fi adapter (2x2 Wi-Fi 6)
- 2020 13-inch MacBook Pro (3x3 Wi-Fi 5)
- iPhone 12 (2x2 Wi-Fi 6)
- 2020 M1 Mac Mini (2x2 Wi-Fi 6)
iPerf details
To test only the speed of the Wi-Fi connection between the client and the AP, my iPerf server was connected over gigabit Ethernet to my Switch Lite 8. I primarily used my dedicated pfSense box for this, but I also used my PC and laptop depending on what I was testing.
To specify which AP and which band was being used, I used the settings offered in the UniFi network controller or Instant On portal, and swapped them in and out as needed. I then stepped through the different channel widths and bands, letting the connection stabilize before beginning my tests.
I ran all of my tests with multiple TCP streams, and occasionally reversed the direction as a point of comparison. These tests ran for 60 seconds, so a typical client command would look like:
iperf3 -c 172.25.10.5 -P 8 -R -t 60
For more details consult the iPerf documentation.
Comments (2)
Thank you for the detailed comparisons!
Great job mate!
