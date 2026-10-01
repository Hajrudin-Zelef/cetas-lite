---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-b5bea11b-4
title: "platform-management-dashboard-administration-design-and-configure-architectures--b5bea11b"
domain: general-networking
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["latency", "throughput"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--b5bea11b.md
source_anchor: ""
source_lines: [156, 191]
sha256: 92135de214e3e405fdc92323bf48562c4cc5bf50b0055ec873c7e896371d1101
---

# platform-management-dashboard-administration-design-and-configure-architectures--b5bea11b

If the client devices require 2.4 GHz, enable 'Dual-band with band steering' to enable client devices to use both 2.4 GHz channels and 5 GHz. Devices will be steered to use the 5 GHz band. For more details refer to the Band Steering Overview article. With a dual-band network, client devices will be steered by the network. If 2.4 GHz support is not needed, it is recommended to use “5 GHz band only”. Testing should be performed in all areas of the environment to ensure there are no coverage holes.
Set Minimum Bitrate
Using RF Profiles, minimum bit rate can be set on a per band or a per SSID basis. For high-density networks, it is recommended to use minimum bit rates per band. If legacy 802.11b devices need to be supported on the wireless network, 11 Mbps is recommended as the minimum bitrate on 2.4 GHz. Adjusting the bitrates can reduce the overhead on the wireless network and improve roaming performance. Increasing this value requires proper coverage and RF planning. An administrator can improve the performance of clients on the 2.4 GHz and 5 GHz band by disabling lower bitrates. Management frames will be sent out at the lowest selected rate. Clients must use either the lowest selected rate or a faster one. Selecting a Minimum bitrate of 12Mbps or greater will prevent 802.11b clients from joining and will increase the efficiency of the RF environment by sending broadcast frames at a higher bitrate.
Note: As per standards, 6 Mbps, 12 Mbps and 24 Mbps are the mandatory data rates. Cisco's San Francisco office uses 18 Mbps as the Minimum bitrate.
Auto Power Reduction
Every second the access point's radios samples the signal-to-noise (SNR) of neighboring access points. The SNR readings are compiled into neighbor reports which are sent to the Meraki Cloud for processing. The Cloud aggregates neighbor reports from each AP. Using the aggregated data, the Cloud can determine each AP's direct neighbors and how by much each AP should adjust its radio transmit power so coverage cells are optimized. For determining the changes in TX power the cloud tries to ensure that there are at least 3 heard by each in the area. The calculations are done every 20 minutes and once complete, the Cloud instructs each AP to decrease or increase the transmit power.TX power can be reduced by 1-3 dB per iteration and is increased in 1 dB iterations.
AutoRF tries to reduce the TX power uniformly for all APs within a network but in complex high density network it is necessary to limit the range and the values for the AP to use. To better support complex environments, minimum and maximum TX power settings can be configured in RF profiles.
Note: For 2.4 GHz, Auto Power reduction algorithm allows TX power to go down only up to 5 dBm. For 5 GHz, Auto Power reduction algorithm allows TX power to go down only up to 8 dBm. If lower TX power is needed, APs can be statically set to lower power.
Auto Channel selection
Adding additional access points on the same channel with overlapping coverage does not increase capacity. To prevent access points nearby from sharing the same channel, Cisco Meraki access points automatically adjusts the channels of the radios to avoid RF interference (Both 802.11 and non-802.11) and develop a channel plan for the Wireless Network. Channels can be selectively assigned to be used with each RF profile. By using channels selectively, network administrators can control the co-channel interference more effectively.
Default Channel Width
- 
    In moving towards 40-Mhz or 80-Mhz channels, you are effectively halving (if selecting 40-MHz) or quartering (80-MHz) the number of non-overlapping 5GHz channels by doubling or quadrupling the channel width due to channel bonding. This, in turn, increases the distance at which access points must be placed if co-channel interference (CCI) and adjacent channel interference (ACI) are to be kept to a minimum.
- 
    While using 40-MHz or 80-Mhz channels might seem like an attractive way to increase overall throughput, one of the consequences is reduced spectral efficiency due to legacy (20-MHz only) clients not being able to take advantage of the wider channel width resulting in the idle spectrum on wider channels. Depending on the RF environment, even clients capable of 40 and 80 MHz may only use the 20 MHz base channel and is often observed in highly contentious RF environments.
- 
    Due to the mix of clients usually seen in high-density deployments (such as laptops, mobile phones, and tablets etc.) the capabilities of clients in such environments also vary (some will support 20-Mhz, some will support 40-MHz and some will support 80-Mhz channels). Due to this, it is better to have each client communicating at the lowest common channel width, giving each client equal access to the network. It is better to have 4 clients communication at 20-MHz with 4 access points, rather than 4 clients of mixed capability communicating with 1 access points at 80-MHz resulting in idle.
DFS Channels and Channel Reuse
For an example deployment with DFS channels enabled and channel reuse is not required, the below grid shows 12 access points with no channel reuse. As there are 19 channels in the US, when you reach 20 access points in the same space, the APs will need to reuse a channel.
For a deployment example where DFS is disabled and channel reuse is required, the below diagram shows 4 channels being reused in the same space. When channel reuse cannot be avoided, the best practice is to separate the access points on the same channel as much as possible.
RX-SOP
| 802.11 Band | High Threshold | Medium Threshold | Low Threshold | 
| 5 GHz | -76 dBm | -78 dBm | -80 dBm | 
| 2.4 GHz | -79 dBm | -82 dBm | -85 dBm | 
Note: RX-SOP is supported on 802.11 ac Wave 2 and 802.11 ax APs i.e. MR30H/33/42/52/53/74/84/42E/53E/45/55
Client Balancing
Roaming in High Density
Client roaming between access points
Enable Fast Roaming
Cisco Meraki MR access points support a wide array of fast roaming technologies. For a high-density network, roaming will occur more often, and fast roaming is important to reduce the latency of applications while roaming between access points. All of these features are enabled by default, except for 802.11r.
- 802.11r (Fast BSS Transition) - 802.11r allows encryption keys to be stored on all of the APs in a network. This way, a client doesn't need to perform the full re-authentication process to a RADIUS server every time it roams to a new access point within the network. This feature can be enabled from the Configure > Access control page under Security > 802.11r. If this option does not appear, a firmware update may be required.
- Opportunistic Key Caching (OKC) - 802.11r and OKC accomplish the same goal of reducing roaming time for clients, the key difference being that 802.11r is standard while OKC is proprietary. Client support for both of these protocols will vary but generally, most mobile phones will offer support for both 802.11r and OKC.
- 802.11i (PMKID caching) - PMK Caching, defined by IEEE 802.11i, is used to increase roaming performance with 802.1X by eliminating the RADIUS exchange that occurs. From a high-level perspective, this occurs by the client sending a PMKID to the AP which has that PMKID stored. If it’s a match the AP knows that the client has previously been through 802.1X authentication and may skip that exchange.
- 802.11k (Neighbor BSS) -802.11k reduces the time required to roam by allowing the client to more quickly determine which AP it should roam to next and how. The AP the client is currently connected to will provide it with information regarding neighboring APs and their channels.
Traffic Shaping
Define Traffic Shaping Rules
