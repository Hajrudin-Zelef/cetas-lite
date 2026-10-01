---
id: collect-261001-meraki/meraki/architectures-and-best-practices-meraki-for-enterprise-rf-best-practices-0b10e7d8-1
title: "architectures-and-best-practices-meraki-for-enterprise-rf-best-practices-0b10e7d8"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "latency", "throughput"]
source: docs/RAG/collect-261001-meraki/architectures-and-best-practices-meraki-for-enterprise-rf-best-practices-0b10e7d8.md
source_anchor: ""
source_lines: [1, 32]
sha256: 5039df23086246018970a6eaf9d62488f2a1e052ae07f1f0639074393cddc772
---

# architectures-and-best-practices-meraki-for-enterprise-rf-best-practices-0b10e7d8

Meraki Wireless for Enterprise Best Practices - RF Design
Click 日本語 for Japanese
Challenge Statement/Requirements
Mobile devices now outnumber fixed devices accessing the internet. It should be apparent that network connectivity has become a critical infrastructure for your business and your customers. Most people are carrying 2-3 devices at the same time. All these devices need to be connected; not having network connectivity is as bad as the power going out, and no work will be done. So, the role of the network has become even more critical. So, how does one begin to think about designing a reliable network that has the capacity to meet current and future business needs? This paper will cover some of the thought processes and give recommendations on how to be successful. These guidelines have been curated from the world’s largest and most complex Wi-Fi installations.
When architecting a Wi-Fi network solution, nothing equals success, like having a solid understanding of the requirements. In Wi-Fi, building for the worst and hoping for the best is a tried-and-true mantra, and through the years, this approach has served the industry well. Technology has significantly improved in the last 4-5 years for both the infrastructure and the client devices. While the newer technologies certainly make it much easier to get fantastic results, it is crucial to understand where these technologies will and won’t provide benefits. This document will give examples of requirements and solutions for a typical enterprise office environment by breaking down the “Office” into functional areas with different coverage needs.
Understanding Client Devices
Thinking about how to design for different client densities and types has evolved with the standards. The rules of thumb continue to evolve; for instance - we once considered using 40 MHz channels as a waste* because of poor client support for 40 MHz channel widths. This changed with the adoption and proliferation of Wi-Fi 5 clients. Wi-Fi 5/802.11ax language stated that a client “shall” support all channel bandwidths up to 80 MHz, whereas 802.11n said a client “may” support up to a 40 MHz channel width. For 802.11n, only the very high-end clients supported 40 MHz, a small percentage of most user populations.
Another example would be Spatial Streams – most client implementations opted for Single Spatial streams to reduce cost and complexity in 802.11n days. Today, with Wi-Fi 5 and 6 being 98% or better of the average client base, most clients support 2 spatial streams. MU-MIMO and beamforming in the standards added the necessity, and most mainstream clients support it today.
What if all clients don’t support MU-MIMO or multiple Spatial Streams? The target goal is to improve airtime efficiency. Even if every client doesn’t support it, every client you get on and off the air faster leaves more Airtime for those needing more time. Everyone gets a better experience with more capacity.
While most users carry between 2-3 devices at a time, it is infrequent that all 3 are being used simultaneously. This means that from a network resource standpoint – things like IP addresses, DHCP, and Auth need a 1 for 1 capacity. Airtime is usually not a factor, though, as the user typically uses just one with the others sitting near idle.
Mobile (and therefore wireless) Client devices in a typical office are broken down into 3 primary categories: Laptops, Smartphones, and Tablets. Then there are “Others”; we’ll discuss them too.
Laptops
Laptops have replaced desktops, with users preferring to take the experience with them throughout the office. As the primary platform, these devices tend to have higher and more consistent bandwidth usage during active work periods. They contribute to network contention, especially when engaged in data-intensive activities like file transfers or video conferencing. Laptops vary depending on the make and model and most today support at least Wi-Fi 5, with many supporting Wi-Fi 6 and 6E. Most commonly used laptops support at least 2 spatial streams on 5GHz.
Laptops typically have larger form factors, allowing for better antenna placement and the potential for larger, higher gain antennas. They provide robust connectivity during data-intensive tasks.
Smartphones
Smartphones are designed for on-the-go communication and quick access to information. While their individual impact on network contention is modest, the cumulative effect can be notable during peak times.
The antennas in smartphone devices are the victims of physics, as a smartphone is a function designed as a small device. Less real estate requires smaller antennas with modest gains averaging between -2 to 0 dBi. These devices are likely to have the worst radio sensitivity on the network and constitute most users' first impressions.
Tablets
Tablets may have variable network usage. They may contribute to contention during active work periods, especially when used for collaborative tasks or streaming. However, during casual use, their impact is generally lower. Tablets are also used for video conferences in the workplace.
Tablets, with a larger form factor than phones, can accommodate slightly more advanced antennas. However, their design still prioritizes portability, and antenna gains remain in the -2 to 0 dBi range.
Other Devices
Apart from these individual user-based wireless devices, another category of “OTHER” devices exists. Some of these can be intense bandwidth hogs, such as high-resolution “Collaboration” endpoints. Printers and increasingly, IoT devices are also common. Things that don’t move should be wired. Squaring that with an all “Wireless” policy might seem to some to be at odds. Wireless is for “Mobility” but there are still a lot of wires involved in a “Wireless” network. For high-resolution collaboration endpoints, It is advised to keep them on the wired network as they are latency-sensitive and bandwidth-hungry and typically do not move in the course of a business day.
IoT and Printers can go either way depending on demand. IoT increasingly comes with only wireless as an option and most printers can be either. Definitely connect them to the network to ensure they are part of the infrastructure network and not running as an independent wireless network. This allows the network to coordinate “ALL” users as a single network.
Some applications should raise concerns for a network admin. One such application is Miracast, a wireless display standard that allows for streaming audio and video content from one device to another, typically from a mobile device or computer to a television or monitor. Miracast uses Wi-Fi Direct, which operates on the 2.4 GHz or 5 GHz bands. Wi-Fi Direct will run outside your network security and operate as an independent network in the same spectrum as your corporate network. It is “Interference” to the primary coverage. Whether or not this causes a problem for the network's clients depend entirely on how much Airtime is left over when both need it.
Bandwidth needs, ranging from 1 Mbps to 4 Mbps, depend on factors like compression efficiency in SD and HD. The challenge arises as it operates as an independent network within the corporate network, causing interference for connected clients.
Applications and Typical Requirements
A client device's utility comes down to its applications. Typically, there will be core “Business Critical” applications that all users require to run the business. These include collaboration, email, file access, web-enabled applications and browsing.
Understanding the clients' Bandwidth and Latency tolerance for critical applications can vary widely depending on the type of application, the nature of the data being transferred, and the required user experience.
| Application | Throughput | 
| Web Browsing | 500 Kbps (kilobits) | 
| Office Applications | 100 Kbps | 
| VoIP | 16 - 320 Kbps | 
