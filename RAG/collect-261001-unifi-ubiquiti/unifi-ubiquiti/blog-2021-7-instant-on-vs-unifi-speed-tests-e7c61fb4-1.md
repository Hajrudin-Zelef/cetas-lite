---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-2021-7-instant-on-vs-unifi-speed-tests-e7c61fb4-1
title: "blog-2021-7-instant-on-vs-unifi-speed-tests-e7c61fb4"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-2021-7-instant-on-vs-unifi-speed-tests-e7c61fb4.md
source_anchor: ""
source_lines: [1, 34]
sha256: b2a61b3aa2b17741ce4675910fcc2911381870c8618418c5a86a365dce76a5c5
---

# blog-2021-7-instant-on-vs-unifi-speed-tests-e7c61fb4

Wi-Fi Speed Tests — Aruba Instant On vs. UniFi
Originally Posted: July 25th, 2021
Last Edited: August 29th, 2021
TL;DR:
- Instant On offers switches and APs with basic cloud-managed software. Instant On switches can be managed locally, but their APs are cloud-only.
- UniFi offers routers, switches, APs, security cameras, access control, and VoIP phones with local or remote management. UniFi software has more features, but also more bugs.
- The early Wi-Fi 6 models (AP22, U6-Lite, U6-LR) usually perform the best, but there are situations where a higher spec’d Wi-Fi 5 AP can outperform them.
- Overall, the AP22 has the best 2.4 GHz performance, and the U6-LR and AC-HD have the best 5 GHz performance.
Wi-Fi Speed Tests — Aruba Instant On vs. UniFi
If you’re not familiar with Aruba, check out my Instant On Overview for more details. I also covered the Instant On software and my overall thoughts in my Instant On AP22 Review. This review focuses on the Wi-Fi access points and performance.
For testing, I’m comparing all available Instant On APs to UniFi’s Wi-Fi 6 U6 Lite and U6 Long Range, as well as UniFi’s older Wi-Fi 5 AC-Pro and AC-HD. I used local iPerf tests and public speed test servers at multiple locations, and with multiple devices. Let’s start by looking at the specs of all of the models, and how they compare.
Comparison of APs Tested
Wi-Fi 6 Models
Wi-Fi 5 Models
There are 6 Instant On APs to pick from, and their model numbers mostly go in order. The first number is their generation (1 or 2), and the 2nd number is their place in the line up. Generally speaking, higher number = higher specs, and higher price.
Starting at the bottom we have the $79 AP11 (AC1200, 2x2), which has a variation known as the AP11D. The AP11D features mostly the same specs, but can be mounted on a desk or in a wall outlet. The AP11 and AP11D performed similarly, so I only showed results from the AP11. The $239 AP17 is Instant On’s only outdoor AP, which is another AC1200 model like the AP11.
The higher-end Instant On APs are the $149 AP12 (AC1600, 3x3 5 GHz) and $199 AP15 (AC2100, 4x4 5 GHz). The $169 AP22 is AX1800, with two spatial streams on each band. The AP22 is the only Wi-Fi 6 model, giving it an edge over all the other Instant On APs.
For UniFi, I tested the first Wi-Fi 6 models, the U6-Lite and U6-LR. These both skipped upgrading their 2.4 GHz radios, making them N/AX1500 and N/AX3000 APs if you want to be pedantic. I also tested the older AC-Pro and AC-HD to show some Wi-Fi 5 results. See the charts below for their full spec differences.
Single Client Speed Tests and Throughput Graphs
For Wi-Fi 6 client performance, the main comparision to watch is the AP22 vs. the U6-Lite and U6-LR. The AP22 is a 2 spatial stream AP like the $99 U6-Lite. The $179 U6-LR is a bit more expensive, but has the advantage of 4 spatial streams on it’s 5 GHz radio. The U6-LR also has a slight edge in transmit power and a few other areas. The AP22 is usually available around $160, so it fits between the two models in price. The AP22 is the only AP with Wi-Fi 6 support on it’s 2.4 GHz radio, though.
For Wi-Fi 5 performance, all the other Instant On APs are compared against the UniFi AC-Pro and AC-HD. The $149 AC-Pro is an older 802.11ac Wave 1 AP, offering 3 spatial streams on it’s 2.4 GHz and 5 GHz radios. The AC-HD is a newer AC Wave 2 model, offering 4 spatial streams on both of it’s radios. This isn’t the most fair of comparisons as the AC-HD retails for $349, but it gives you an idea of what a higher-end Wi-Fi 5 AP can deliver. The $79 AP11 is the cheapest AP here, so keep that in mind when viewing the test results.
In all the speed test charts that follow, the numbers I’m showing are throughput in Mbps, averaged over five or more minute-long local iPerf TCP tests. I went over these numbers multiple times, and I tried to make them as accurate as possible. You won’t necessarily see the same results on your network with your devices, but it should give you a general idea of expected performance. Keep in mind that these numbers represent averages rather than exact measurements. I cover my testing setup at the end of the article if you’re interested in more details about how and why I test this way.
The first few tests cover an ideal scenario, with a nearby client on a clean channel. In typical use you’ll see less throughput. This is a test of the APs capability in an ideal scenario, and how much data they can deliver to a single client.
Wi-Fi 6 Speed Comparison
2.4 GHz Speed Comparison
First, I tested all of the APs on 2.4 GHz, trying both 20 MHz and 40 MHz channels. I don’t recommend using 40 MHz channels in the 2.4 GHz band, due to them overlapping with over 80% of the already-crowded spectrum. There’s only one non-overlapping 40 MHz channel in North America, and the rest of the world only has two. Like 160 MHz channels in 5 GHz, there’s just not enough available frequency for them to be reliably used in most situations.
In the best-case scenario throughput roughly doubles, but you’re more likely to have issues with interference, less range, and an inconsistent experience. Most manufacturers quoted data rates for 2.4 GHz rely on this 40 MHz channel trick. With a normal 2x2 client, you’re more likely to effectively use 20 MHz channels with 150 Mbps data rates than 40 MHz channels with 300 Mbps. Wi-Fi 6 raises that to 287/574 Mbps, but that only applies to the AP22.
Of these nine access points, the AP22 is the only one that supports Wi-Fi 6 on it’s 2.4 GHz radio. The Wi-Fi 5 standard only applied to the 5 GHz band, and the U6-Lite and U6-LR both stayed with older 2.4 GHz radios that support Wi-Fi 4 (802.11n). The results are as you’d expect, and it’s not a very close competition.
In these results you can see the impact of higher 1024-QAM modulation and the longer symbol duration of WiFi 6, resulting in a theoretical 38% increase in data rates. In most cases, the throughput difference is more like 10-20%. The difference is more extreme when you’re very close to the AP, and in an ideal scenario like I’m testing here. Since I was using a 2 stream client, the extra spatial streams in the 5 GHz radios of the AC-Pro and AC-HD didn’t come into play, besides possibly improving beamforming.
The AP22 is the only AP which can leverage Wi-Fi 6, resulting in a significantly higher max data rate (574 vs 300 Mbps when using 2 spatial streams on a 40 MHz channel), and significantly higher throughput. All the other APs are stuck at lower modulation rates, and performed similarly. The difference in real world performance isn’t going to be as big as my results would lead you to believe, but the AP22 is undeniably the 2.4 GHz champ.
5 GHz Speed Comparison
Next, I did the same test using 20, 40, and 80 MHz channels in 5 GHz. At 80 MHz, the Wi-Fi 5 models maxed out at a typical 867 Mbps data rate, while the AP22, U6-Lite, and U6-LR top out at 1200 Mbps. You can see the impact of Wi-Fi 6 on all three channel widths, with the biggest difference being at 80 MHz. At this width, Wi-Fi 6 closes in on the gigabit barrier. It’s possible to get up to near gigabit speeds with 80 MHz channels, but throughput over 1 Gbps usually requires 160 MHz width, or a 3rd spatial stream. It also requires near-ideal conditions and short range like I’m showing here.
Also worth noting: The AC-Pro beat the AP11 with 20 MHz width, but struggled with wider channels. The AP12 and AP15 couldn’t match the AC-HD at 80 MHz. Both of those are signs of things to come.
For the next test, I switched over to my MacBook Pro and it’s 3 spatial stream Wi-Fi 5 radio. This is an interesting test because it shows the impact of an additional spatial stream, and removes the highest-end modulation (1024-QAM) and longer symbol duration of Wi-Fi 6. This is a more even playing field, and a chance for the 3x3 and 4x4 APs to show their strength.
