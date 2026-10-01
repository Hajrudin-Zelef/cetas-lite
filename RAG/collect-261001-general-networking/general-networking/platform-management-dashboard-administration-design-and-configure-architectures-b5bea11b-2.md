---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-b5bea11b-2
title: "platform-management-dashboard-administration-design-and-configure-architectures--b5bea11b"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Microsoft"]
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--b5bea11b.md
source_anchor: ""
source_lines: [57, 119]
sha256: cb292f956a5213368731719b93949e80ee79969ee621fe3e36a17e018e2c12db
---

# platform-management-dashboard-administration-design-and-configure-architectures--b5bea11b

Note: In order to ensure quality of experience it is recommended to have around 25 clients per radio or 50 clients per AP in high-density deployments.
Starting 802.11n, channel bonding is available to increase throughput available to clients but as a result of channel bonding the number of unique available channels for APs also reduces. Due to the reduced channel availability, co-channel interference can increase for bigger deployments as channel reuse is impacted causing a negative impact on overall throughput.
Note:In a high-density environment, a channel width of 20 MHz is a common recommendation to reduce the number of access points using the same channel.
Client devices don’t always support the fastest data rates. Device vendors have different implementations of the 802.11ac standard. To increase battery life and reduce size, most smartphone and tablets are often designed with one (most common) or two (most new devices) Wi-Fi antennas inside. This design has led to slower speeds on mobile devices by limiting all of these devices to a lower stream than supported by the standard. In the chart below, you can see the maximum data rates for single stream (433 Mbps), two stream (866 Mbps), and three stream (1300 Mbps). No devices on the market today support 4 spatial streams or wider 160 MHz channels, but these are often advertised as optional "Wave 2" features of the 802.11ac standard.
| Streams | 20 MHz Channel Width | 40 MHz Channel Width | 80 MHz Channel Width | 
| 1 Stream | 87 Mbps | 200 Mbps | 433 Mbps | 
| 2 Streams | 173 Mbps | 400Mbps | 866 Mbps | 
| 3 Streams | 289 Mbps | 600 Mbps | 1300 Mbps | 
The actual device throughput is what matters to the end user, and this differs from the data rates. Data rates represent the rate at which data packets will be carried over the medium. Packets contain a certain amount of overhead that is required to address and control the packets. The actual throughput is payload data without the overhead. Based on the advertised data rate, next estimate the wireless throughput capability of the client devices. A common estimate of a device's actual throughput is about half of the data rate as advertised by its manufacturer. As noted above, it is important to also reduce this value to the data rate for a 20 MHz channel width. Below are the most common data rates and the estimated device throughput (half of the advertised rate). Given the multiple factors affecting performance it is a good practice to reduce the throughput further by 30%
| Protocol | Data rate (Mbps) | Estimated Throughput (1/2 advertised rate) | Throughput w/Overhead | 
|---|---|---|---|
| 802.11a or 802.11g | 54 Mbps | 27 Mbps | ~19 Mbps | 
| 1 stream 802.11n | 72 Mbps | 36 Mbps | ~25 Mbps | 
| 2 stream 802.11n | 144 Mbps | 72 Mbps | ~50 Mbps | 
| 3 stream 802.11n | 216 Mbps | 108 Mbps | ~76 Mbps | 
| 1 stream 802.11ac | 87 Mbps | 44 Mbps | ~31 Mbps | 
| 2 stream 802.11ac | 173 Mbps | 87 Mbps | ~61 Mbps | 
| 3 stream 802.11ac | 289 Mbps | 145 Mbps | ~102 Mbps | 
| 1 stream 802.11ax | 143 Mbps | 72 Mbps | ~50 Mbps | 
| 2 stream 802.11ax | 287 Mbps | 144 Mbps | ~101 Mbps | 
| 3 stream 802.11ax | 430 Mbps | 215 Mbps | ~151 Mbps | 
Estimate the Number of APs
It's important to document and review the requirements and assumptions and confirm they are reasonable. Changing one assumption will significantly impact the number of access points and the costs. If you assumed just 1.5 Mbps for HD video chat (as recommended by Microsoft Skype and Cisco Spark) you would need half the number of access points. If you assumed 5 Mbps was required for HD video streaming (as recommended by Netflix) you would need more access points. If you were designing to support 600 1 stream devices instead of 600 3 stream laptops, you would need roughly 3 times the number of access points. For this example, we now have the following requirements and assumptions:
- 
    Video streaming requires 3 Mbps for HD quality video
- 
    There will be 600 concurrent users streaming video to their laptop
- 
    Every user has an Apple MacBook Pro or similar
- 
    All laptops support 802.11ac and are capable of 3 spatial streams
- 
    The network will be configured to use 20 MHz channels
- 
    Each access point can provide up to 101 Mbps of wireless throughput
We can now calculate roughly how many APs are needed to satisfy the application capacity. Round to the nearest whole number.
Number of Access Points based on throughput = (Aggregate Application Throughput) / (Device Throughput)
Number of Access Points based on throughput = 1800 Mbps/101Mbps = ~18 APs
In addition to the number of APs based on throughput, it is also important to calculate the number of APs based on clients count. To determine number of APs, first step is to estimate the clients per band. With newer technologies, more devices now support dual band operation and hence using proprietary implementation noted above devices can be steered to 5 GHz.
Note: A common design strategy is to do a 30/70 split between 2.4 GHz and 5 GHz
For this example, we now have the following requirements and assumptions:
- 
    There will be 600 concurrent users streaming video to their laptop
- 
    Concurrent 2.4 GHz clients = 600 * 0.3 = 180
- 
    Concurrent 5 GHz clients = 600 * 0.7 = 420
We can now calculate roughly how many APs are needed to satisfy the client count. Round to the nearest whole number.
Number of Access Points based on client count = (Concurrent 5 GHz clients) / 25
Number of Access Points based on client count = 420 / 25 = ~17 APs
Now the Number of APs required can be calculated by using the higher of the two AP counts.
Number of Access Points = Max (Number of Access Points based on throughput, Number of Access Points based on client count)
Number of Access Points = Max (18,17) = 18 APs
Site Survey and Design
Performing an active wireless site survey is a critical component of successfully deploying a high-density wireless network and helps to evaluate the RF propagation in the actual physical environment. The active site survey also gives you the ability to actively transmit data and get data rate coverage in addition to the range.
In addition to verifying the RF propagation in the actual environment, it is also recommended to have a spectrum analysis done as part of the site survey in order to locate any potential sources of RF interference and take steps to remediate them. Site surveys and spectrum analysis are typically performed using professional grade toolkits such as Ekahau Site Survey or Fluke Networks Airmagnet. Ensure a minimum of 25 dB SNR throughout the desired coverage area. Remember to survey for adequate coverage on 5GHz channels, not just 2.4 GHz, to ensure there are no coverage holes or gaps. Depending on how big the space is and the number of access points deployed, there may be a need to selectively turn off some of the 2.4GHz radios on some of the access points to avoid excessive co-channel interference between all the access points.
Note: It is recommended to have complete coverage for both bands.
Note: Read our guide on Conducting Site Surveys with MR Access Points for more help on conducting an RF site survey.
Mounting Access Points
The two main strategies for mounting Cisco Meraki access points are ceiling mounted and wall mounted. Each mounting solution has advantages.
Ceiling mounted MR, Cisco San Francisco
Ceiling mounted access points are placed on a ceiling tile, T-bar, roof, or conduit extending down from the roof. This brings advantages such as a clear line-of-sight to the user devices below and flexibility in where to place the access point. Access points can be easily placed with even spacing in a grid and at the intersection of hallways. The disadvantage is the ceiling height and the height of the access point could negatively impact the coverage and capacity.
- 
