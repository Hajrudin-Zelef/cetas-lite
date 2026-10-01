---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-b5bea11b-1
title: "platform-management-dashboard-administration-design-and-configure-architectures--b5bea11b"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["distribution", "throughput"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--b5bea11b.md
source_anchor: ""
source_lines: [1, 56]
sha256: b99bd9adb74d514c113c616d6fd5b5a60fb3809c7e8e99e9dfd8cdf1747a945a
---

# platform-management-dashboard-administration-design-and-configure-architectures--b5bea11b

High Density Wi-Fi Deployments
High-density Wi-Fi is a design strategy for large deployments to provide pervasive connectivity to clients when a high number of clients are expected to connect to Access Points within a small space. A location can be classified as high density if more than 30 clients are connecting to an AP. To better support high-density wireless, Cisco Meraki access points are built with a dedicated radio for RF spectrum monitoring allowing the MR to handle the high-density environments. Unless additional sensors or air monitors are added, access points without this dedicated radio have to use proprietary methods for opportunistic scans to better gauge the RF environment and may result in suboptimal performance.
Large campuses with multiple floors, distributed buildings, office spaces, and large event spaces are considered high density due to the number of access points and devices connecting. More extreme examples of high-density environments include sports stadiums, university auditoriums, casinos, event centers, and theaters.
As Wi-Fi continues to become ubiquitous, there is an increasing number of devices consuming an increasing amount of bandwidth. The increased need for pervasive connectivity can put additional strain on wireless deployments. Adapting to these changing needs will not always require more access points to support greater client density. As the needs for wireless connectivity have changed over time, the IEEE 802.11 wireless LAN standards have changed to adapt to greater density, from the earliest 802.11a and 802.11b standards in 1999 to the most recent 802.11ac standard, introduced in 2013 and the new 802.11ax standard currently being developed.
Planning
In the recent past, the process to design a Wi-Fi network centered around a physical site survey to determine the fewest number of access points that would provide sufficient coverage. By evaluating survey results against a predefined minimum acceptable signal strength, the design would be considered a success. While this methodology works well to design for coverage, it does not take into account requirements based on the number of clients, their capabilities, and their applications' bandwidth needs.
Understanding the requirements for the high density design is the first step and helps ensure a successful design. This planning helps reduce the need for further site surveys after installation and for the need to deploy additional access points over time. It is recommended to have the following details before moving onto the next steps in the design process:
- 
    Type of applications expected on the network
- 
    Supported technologies (802.11 a/b/g/n/ac/ax)
- 
    Type of clients to be supported (Number of spatial streams, technologies, etc.)
- 
    Areas to be covered
- 
    Expected number of simultaneous devices in each area
- 
    Aesthetic requirements (if any)
- 
    Cabling constraints (if any)
- 
    Power constraints (It’s best to have PoE+ capable infrastructure to support high performance APs)
Capacity Planning
Once the above mentioned details are available, capacity planning can then be broken down into the following phases:
- 
    Estimate Aggregate Application Throughput
- 
    Estimate Device Throughput
- 
    Estimate Number of APs
Calculating the number of access points necessary to meet a site's bandwidth needs is the recommended way to start a design for any high density wireless network.
Estimate Aggregate Application Throughput
Usually there is a primary application that is driving the need for connectivity. Understanding the throughput requirements for this application and any other activities on the network will provide will provide a per-user bandwidth goal. This required per-user bandwidth will be used to drive further design decisions. Throughput requirements for some popular applications is as given below:
| Application | Throughput | 
| Web Browsing | 500 kbps (kilobits) | 
| VoIP | 16 - 320 kbps | 
| Video conferencing | 1.5 Mbps | 
| Streaming - Audio | 128 - 320 kbps | 
| Streaming - Video | 768 kbps | 
| Streaming - Video HD | 768 kbps - 8mbps | 
| Streaming - 4K | 8 mbps - 20mbps | 
Note: In all cases, it is highly advisable to test the target application and validate its actual bandwidth requirements. It is also important to validate applications on a representative sample of the devices that are to be supported in the WLAN. Additionally, not all browsers and operating systems enjoy the same efficiencies, and an application that runs fine in 100 kilobits per second (Kbps) on a Windows laptop with Microsoft Internet Explorer or Firefox, may require more bandwidth when being viewed on a smartphone or tablet with an embedded browser and operating system
Once the required bandwidth throughput per connection and application is known, this number can be used to determine the aggregate bandwidth required in the WLAN coverage area. It is recommended to have an aggregate throughput for different areas such as classrooms, lobby, auditorium, etc. as the requirements for these areas might be different.
As an example, we will design a high-density Wi-Fi network to support HD video streaming that requires 3 Mbps of throughput. Based on the capacity of the auditorium, there may be up to 600 users watching the HD video stream. The aggregate application throughput can be calculated using the below given formula:
(Application Throughput) x (Number of concurrent Users) = Aggregate Application Throughput
3 Mbps x 600 users = 1800 Mbps
Note that 1.8 Gbps exceeds the bandwidth offerings of almost all internet service providers. The total application bandwidth we are estimating is a theoretical demand upper bound, which will be used in subsequent calculations.
Estimate Device Throughput
While Meraki APs support the latest technologies and can support maximum data rates defined as per the standards, average device throughput available often dictated by the other factors such as client capabilities, simultaneous clients per AP, technologies to be supported, bandwidth, etc.
Client capabilities have a significant impact on throughput as a client supporting only legacy rates will have lower throughput as compared to a client supporting newer technologies. Additionally, bands supported by the client may also have some impact on the throughput. Meraki APs have band steering feature that can be enabled to steer dual band clients to 5 GHz.
Note: A client supporting only 2.4 GHz might have lower throughput as compared to a dual band client since higher noise level is expected on the 2.4GHz as compared to 5 GHz and the client might negotiate lower data rate on 2.4GHz.
In certain cases, having dedicated SSID for each band is also recommended to better manage client distribution across bands and also removes the possibility of any compatibility issues that may arise.
To assess client throughput requirements, survey client devices and determine their wireless capabilities. It is important to identify the supported wireless bands (2.4 GHz vs 5 GHz), supported wireless standards (802.11a/b/g/n/ac), and the number of spatial streams each device supports. Since it isn’t always possible to find the supported data rates of a client device through its documentation, the Client details page on Dashboard can be used as an easy way to determine capabilities.
Example Client details listing
Wi-Fi is based on CSMA/CA and is half-duplex. That means only one device can talk at a time while the other devices connected to the same AP wait to for their turn to access the channel. Hence, simultaneous client count also has an impact on AP throughput as the available spectrum is divided among all clients connected to the AP. While Meraki has client balancing feature to ensure clients are evenly distributed across AP in an area an expected client count per AP should be known for capacity planning.
