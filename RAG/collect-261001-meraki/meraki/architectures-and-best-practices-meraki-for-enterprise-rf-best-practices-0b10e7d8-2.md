---
id: collect-261001-meraki/meraki/architectures-and-best-practices-meraki-for-enterprise-rf-best-practices-0b10e7d8-2
title: "architectures-and-best-practices-meraki-for-enterprise-rf-best-practices-0b10e7d8"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/architectures-and-best-practices-meraki-for-enterprise-rf-best-practices-0b10e7d8.md
source_anchor: ""
source_lines: [33, 73]
sha256: 9bc5358d01a9e3e860f44fb7dcdee96b9f148ab3a5a7474206ca6e663c2a3d71
---

# architectures-and-best-practices-meraki-for-enterprise-rf-best-practices-0b10e7d8

| Video conferencing | 1 -3 Mbps | 
| Streaming - Audio | 128 - 320 Kbps | 
| Streaming - Video | 768 Kbps | 
| Streaming - Video HD | 768 Kbps – 8Mbps | 
| Streaming - 4K | 8 – 20Mbps | 
| Streaming - 8K | 100 – 150 Mbps | 
| AR/VR | 2 - 200 Mbps | 
| Cloud service | 100 – 200 Mbps | 
Design Solutions
An Access Point represents a finite amount of potential capacity in an enterprise space. Each client joining the cell gets more or less of that cell's capacity based on its application needs. Suppose the number of clients or the requirements of the applications exceed the capacity available in a single cell. In that case, another AP/Cell is required to increase the capacity available in the same space. Each cell requires an operating channel that doesn’t interfere with other cells and the number of channels is limited. If you exceed the number of channels operated within a given space, the result is interference and a performance loss.
The following design guidelines have been developed to help you – the architect produce and configure high-performance coverage without interference within an enterprise office environment. These guidelines are applicable to an Enterprise open office environment with ceiling heights between 8-12 ft (2.4-3 m). Cisco/Meraki best practices recommend AP densities between 1.2- 2k f2 (110 to 185 m2) per Access Point for 5- and 6 GHz radio coverage. This cell size yields AP placement at 40-50 Ft./ 12-15 m between APs.
At 5 GHz, channel width can be 40 MHz for 1, 5 GHz interface or 2 interfaces at 20 MHz each without minimal interference.
For Wi-Fi 6E/6 GHz: In the ETSI plan, regions restricted to UNii 5 (500 MHz) use 40 MHz channel width; for regions with the full UNii 5-8 and 1200 MHz operating band use up to 80 MHz channels.
In all cases, the goal is consistent user experience within any cell. Cisco Meraki recommends using Manual channel width selection in both the 5 and 6 GHz bands for high client-density networks.
Please refer to the Cisco High-Density Wireless Design video to better understand channel planning for high-performance networks. Depending on the AP and client density, networks typically operate with transmit powers ranging from 4 to 15 dBm for 5GHz.
We recommend using AutoRF for Transmit power selection with a range as mentioned below in the usecases.
Note: Do not set Transmit power rage min and max at the highest level. This can cause high interference in neighboring Aps and defeat the purpose of using AutoRF
Optimally, Access points (APs) should be positioned so that they do not detect each other on the same channel with a signal strength above -82 dBm. In modern enterprise environments, getting APs closer in some locations is usually necessary.  You can review the RF spectrum page to check if there are APs detecting each other on the same channel (add screenshot and explanation of how to read it).
If you notice multiple neighbor APs on the same channel with an RSSI stronger than -80 dB, you could manage it by configuring RX-SOP at -78 dBm. The higher the RX-SOP level, the less sensitive the radio is and the smaller the receiver's cell size. However, this should be used with caution, as you can create coverage area issues if this is set too high. If using values above -78 dBm, it is recommended to test connectivity with a couple of different client types at the intended edge of the cell to ensure suitability before implementing RX-SOP in production.
Note: We recommend using custom RF profiles while using RX-SOP.
Configuring the minimum bitrate is of utmost importance to restrict the types of clients allowed to connect to your network. Allowing clients to remain connected with lower data rates can negatively impact the experience of other clients sharing the same spectrum, as they all share the same channels and Airtime. The minimum data rate also helps determine the signal strength at which a client can directly connect. At lower AP density (APs are further apart), a lower minimum bitrate would be selected; in a higher density network (more APs closer together, a higher minimum bitrate would be used to help encourage roaming and load balancing. The minimum bitrate is typically set at 12 or 24 in a typical enterprise office environment, depending on the AP density. Caution: Setting the bitrate too high can result in coverage gaps or holes in the network.
Reducing cell size ensures the clients are connected to the nearest access point using the highest possible data rates. In a high-density environment, the smaller the cell size, the better. In a dense AP environment, increasing TX Power to uphold SNR is essential. Evaluating the optimal power level becomes crucial in this context. However, in a genuinely congested environment, power levels will naturally decrease, making higher power only justified if dealing with a high ceiling.
Let us examine three distinct use cases characterized by growing client density and complexity in an office building. Determine the essential design criteria and understand the reasons for their variation. Identify the design elements requiring management, explore the employed solutions, and delve into the best practices and configuration choices, explaining the rationale behind each decision. In each scenario, our assumptions and proposed solutions are grounded in the existing hardware infrastructure, designed explicitly for 802.11ax and briefly touch upon Wi-Fi 6E deployment.
Open Office
As we mentioned earlier about a regular office space, the average cell size of an AP should be approx. 2k sq. mt. The data rate is 12-24, depending on roaming requirements. With advancements in the standards and technology, we can now
The Open Office RF profile template gives a pretty good configuration set for an office space to have an ideal RF configuration.
This recommendation does not stand true for warehouse and retail environments.
| Band vs. RF Profile recommendations | Tx power | Channel Width | Data Rate | RX-SOP | 
| 2.4GHz | 11-17 dBm | 20MHz | 12 Mbps | -78 | 
| 5GHz | 14-20 dBm | 20/40MHz | 12-24 Mbps | -78 | 
| 6GHz | 8-30 dBm | 40/80MHz | 12-24 Mbps | -78 | 
Conference /Quiet meeting rooms
Let us dissect factors such as distance, size, and the practical consideration of the expected number of occupants, leading to a certain threshold. A highly effective practice involves mounting an access point, ideally at the center of the conference room cluster, creating a substantial coverage area. This setup accommodates a sizable group of people, forming a 'bubble of noise' around the AP near them, thereby preserving the signal quality. This approach ensures consistent connectivity, preventing the usage of Airtime that would otherwise extend to your open office floor.
Opt for a higher Minimum Mandatory Data rate, like 24 Mbps, while disabling all rates below it to minimize the cell size. This approach aims to configure user traffic within the conference room to the immediate boundaries. Raising the minimum data rate necessitates client devices to be closer to the AP, enabling a higher SNR for association. This effectively diminishes the operational cell size, ensuring that all connected devices utilize the specified data rate or higher throughout their connection.
| Band vs. RF Profile recommendations | Tx power | Channel Width | Data Rate | RX-SOP | 
| 2.4GHz | 8-14 dBm | 20 MHz | 24 Mbps | -78 | 
| 5GHz | 11-17 dBm | 40 MHz | 24 Mbps | -78 | 
| 6GHz | 8-30 dBm | 80 MHz | 24 Mbps | -78 | 
Please refer to the High-Density Wi-Fi Deployments guide for more information on RF design.
Common areas
Common areas in the office include anything with open space like the atrium, lobby, and break rooms need a different Access point deployment than a regular open office deployment.
