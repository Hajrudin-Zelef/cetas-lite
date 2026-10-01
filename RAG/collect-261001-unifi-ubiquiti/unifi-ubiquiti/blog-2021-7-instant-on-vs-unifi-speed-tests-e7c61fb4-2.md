---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-2021-7-instant-on-vs-unifi-speed-tests-e7c61fb4-2
title: "blog-2021-7-instant-on-vs-unifi-speed-tests-e7c61fb4"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["cost", "throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-2021-7-instant-on-vs-unifi-speed-tests-e7c61fb4.md
source_anchor: ""
source_lines: [35, 62]
sha256: 2524bc9f6576b379b307ddcf54d58ca8a9cecb652ee21a4287bfcd0268db8881
---

# blog-2021-7-instant-on-vs-unifi-speed-tests-e7c61fb4

The AP12, AP15, U6-LR, AC-HD, and AC-Pro are all on equal footing in this test, offering 3 or 4 spatial streams, 256-QAM, and up to 1300 Mbps data rates to my 3x3 client. The AP12, AP15 and AC-Pro fall behind with wider channels, unable to match the U6-LR and AC-HD. The AP22 punches a bit above its weight, outperforming all the other 2 stream APs.
I expected the 3x3 AP12, 4x4 AP15, and 3x3 AC-Pro to perform better here. Their high end performance was less than I expected, and the AC-Pro actually achieved the worst 80 MHz result of the group. At smaller channel widths, these APs leverage the additional spatial stream well. At 80 MHz, they can’t compete with the U6-LR and AC-HD.
All the other APs support 2 spatial streams, making them incapable of matching the highest data rates of the higher spatial stream APs. Without the advantage of a 3rd spatial stream, the AP11, AP17, AP22, and U6-Lite fell behind. Despite this, at 80 MHz, they all managed beat the AC-Pro, and traded blows with the AP12 and AP15.
I couldn’t get my MacBook Pro to associate to a 40 MHz channel on any of these APs, so I excluded those results from this test. I believe Apple uses the “fat channel intolerant” setting on their devices. As always, 5 GHz is the best option for speed, and 40 MHz channels on 2.4 GHz should be avoided in most situations.
This chart also shows how the same 20 MHz channel width (combined with a 3rd spatial stream) can push more data over 5 GHz. This is due to Wi-Fi 5’s top data rates, which use 256-QAM modulation. 2.4 GHz wasn’t changed in the Wi-Fi 5 standard, so It typically tops out at 64-QAM, resulting in a lower data rate and lower throughput.
2.4 GHz Distance Testing
For my next test, I switched back to my 2x2 Wi-Fi 6 client, and tested range by moving the APs to 3 different places in my house. I wanted to show the impact of distance from your AP on a typical Wi-Fi signal. All of the above tests were very close range, and meant to show a best-case scenario. These tests are more realistic, and the 15 feet + 1 wall results are more likely what you will see in typical use.
With every foot of free space and every obstruction, a Wi-Fi signal attenuates and gets weaker. 5 GHz signals attenuate faster, and provide around half the range of 2.4 GHz. When deciding on how many access points you need, a good general rule is don’t expect 5 GHz coverage to extend further than 2 walls or 30 feet away. 2.4 GHz signals extend this circle out a bit, but with a few walls in the way, getting low SNR links and slow performance is likely. If there is clear line of sight AP range can extend much further, but every wall imposes a dBm penalty. Wall material and quantity are usually more important than distance in a home or small business network.
Since I ran these tests for my U6-Lite and U6-LR review, I modified the location of my “30 Feet + 2 Walls” slightly, resulting in a more dramatic fall off in signal and throughput. These results show how the AP performs when it’s 2.4 GHz signal is hovering around -65 dBm RSSI and around 20-25 SNR. The lower EIRP of the U6-Lite is a limiting factor here, achieving less SNR and slightly lower modulation rates out of this less-than-ideal link. The AP22’s 2.4 GHz radio advantage lets it pulls ahead for the closer tests, but it didn’t match the U6-LR at range.
Note for International Readers
- 5 feet = 1.5 meters
- 15 feet = 4.6 meters
- 30 feet = 9.1 meters
5 GHz Distance Testing
Next, I ran the same test on the 5 GHz band with 80 MHz channels. Wider channels give you the best speeds, but also require a stronger signal for effective use. At the farthest location, the speed advantage of 5 GHz is mostly eliminated. Those results show how the AP performs when it’s 5 GHz signal is hovering around -80 dBm RSSI and around 10 SNR. From the same location 2.4 GHz connections are stronger and more stable.
When further away, you can also see the impact of beamforming from the AC-HD and U6-LR. They are able to compensate by directing transmissions towards the distant client, and the U6-LR performed the best at the furthest location. This is where the U6-LR shows it’s biggest advantage over the AP22, and where I thought the AP12 and AP15 would do better than they did. The U6-LR’s high transmit power and beamforming allow it to reach further than any of the other models I tested. A few extra dBm is enough to allow the U6-LR to effectively cover a larger area, or punch through one more wall. The AP12 and AP15 did the 2nd best at range, coming close to the U6-LR’s performance.
If I moved a bit further away, the 2.4 GHz connections would slow down, and 5 GHz connections would drop into unusable levels. Normal clients will likely roam on to the stronger 2.4 GHz at my “30 Feet + 2 Walls” range, as they should. Using Wi-Fi with an SNR around 10 dBm isn’t ideal, and the results clearly show that.
Overall Winners and Recommendations
The AP22 clearly offers the best 2.4 GHz performance. 5 GHz performance depends on a lot of factors, but the U6-LR and the AC-HD performed the best overall. The AC-Pro, AP12, and AP15 all struggled with 80 MHz channels to a single client. I don’t have a reliable way to test multi-client performance yet, but they should offer more performance in multi-client tests and realistic use.
If you’re looking for an Instant On AP, the AP22 is the obvious first choice. It offers the most performance-per-dollar. If cost is most important, the AP11 is a viable option. In some cases the AP11 beats the $149 AC-Pro. Wi-Fi 6 isn’t a massive upgrade in speeds, so spending half as much on an AP11 may be the better choice for some networks.
For UniFi, the U6-Lite and U6-LR are again the obvious choices, if you can find them in stock and around MSRP. Like a lot of companies, Ubiquiti is struggling with keeping things in stock right now. If you can’t wait around to find one of their Wi-Fi models, the AC Wave 1 and AC Wave 2 models are usually available. I didn’t cover all of those models in these tests, but models like the nanoHD and FlexHD are good options, especially if high-end Wi-Fi 6 performance doesn’t matter.
Specialty: In-Wall and Outdoor
The AP11D and AP17 are for specialty use cases. If you need an AP outside or in a wall outlet, those are fine APs with decent performance. I don’t have any UniFi In-Wall APs to test, but I’ll be looking into adding some outdoor UniFi models to future tests. UniFi offers a lot more model choices for outdoor APs, and I’m a fan of the older AC-Mesh and AC-Mesh-Pros.
AP Comparisons
The most interesting comparison for Instant On APs is the AP22 vs the AP12 and AP15. I was disappointed in the performance of the AP12 and AP15 when using wider channels. The charts above only capture single-client tests, which is not the best way to show the AP12 and AP15’s strengths. Still, they weren’t able to match up against the (much more expensive) AC-HD. I wouldn’t generally recommend the AP12 or AP15 unless you’re deploying them in a dense area with a lot of devices.
The AP22 vs. the U6-Lite and U6-LR is another interesting comparison. The U6-Lite is significantly cheaper ($99 vs. ~$165), but the AP22 has a much better 2.4 GHz radio, and outperformed the U6-Lite in basically every test. The U6-LR is more expensive, but it excels at 5 GHz, and offers the most range. The AP22 is more well-rounded, and a good default option for most networks. The software is the bigger difference, but that’s a topic for another article.
If you’re able to find any of the Wi-Fi 6 models in stock around MSRP, I think they are all good access points. The U6-Lite’s weaknesses are 2.4 GHz, range, and availability. The AP22 is better, but is also priced more in line with U6-LR, and it can’t match the U6-LR’s range.
Range Comparisons
