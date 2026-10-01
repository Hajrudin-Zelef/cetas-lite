---
id: collect-261001-meraki/meraki/architectures-and-best-practices-meraki-for-enterprise-rf-best-practices-0b10e7d8-3
title: "architectures-and-best-practices-meraki-for-enterprise-rf-best-practices-0b10e7d8"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["training"]
source: docs/RAG/collect-261001-meraki/architectures-and-best-practices-meraki-for-enterprise-rf-best-practices-0b10e7d8.md
source_anchor: ""
source_lines: [74, 92]
sha256: 86474054a294dca0c5c73543b391b9d5fb8a5c9ed714bb00ba9341080db5402b
---

# architectures-and-best-practices-meraki-for-enterprise-rf-best-practices-0b10e7d8

Operating traffic at a higher data rate reduces Airtime and diminishes the interference radius. Achieving a higher SNR is imperative for successful signal demodulation, emphasizing that below this threshold, the signal is essentially treated as noise."
Caveat: Consider the number of users and devices expected in the open space. High-density areas may require additional access points to handle increased device density and demand. Consider 2000 sq ft per AP for low to moderate client density, but if you plan to have a gathering, plan on temporary augmented coverage.
You can use a CW9166/9164i in the Lobby and break room areas to get optimum coverage and 6Ghz coverage in allowed regions.
Things to consider while designing a break room are, again, client traffic will increase during snack time and people walking to the break room to make phone calls who would expect a good Wi-Fi connection.
At the same time, the lobby is another scenario where one won’t expect to have a lot of issues until some executive enters the building and connects to the Wi-Fi while on a call, only to experience bad audio calls.
Now, talking about open spaces in the building like an atrium, more focused directional antennas can solve many inter-floor interference problems. This is a primary problem noticed by customers with glass buildings with atriums across all the floors and bad Wi-Fi experience closer to the atrium. The primary reason for this is glass walls cause unpredictable patterns, which leads to too much interference for all the neighboring APs. An option to use a directional antenna AP like CW-9166D-MR. The example below is of the Cisco Cafeteria
| Band vs. RF Profile recommendations | Tx power | Channel Width | Data Rate | RX-SOP | 
| 2.4GHz | 5-8 dBm | 20MHz | 12 Mbps | -78 | 
| 5GHz | 10-12 dBM | 20/40MHz | 18-24 Mbps | -78 | 
| 6GHz | 12-14 dBm | 40/80MHz | 18-24 Mbps | -78 | 
Training Room/ Classsrooms
To accommodate 400 users in a single corporate environment, we adopt a standard High Client Density approach with a user-to-AP (Access Point) ratio of 50 users per radio interface. This ratio is based on practical experience, taking into account the expected types of client applications and activities. It's a prudent figure, factoring in various background operations like virus scanning and backup, as well as routine IT maintenance that initiates when connecting to our (Cisco's) internal networks. These processes, although not heavily bandwidth-intensive, can impact bandwidth usage, especially on computers that are infrequently used in the office. However, individual experiences may differ.
During meetings, attendees often simultaneously watch a Webex broadcast and the live speaker to engage with remote participants. This dual participation surprisingly does not significantly increase bandwidth demands. The 50-users-per-interface ratio also accounts for additional people standing along walls and aisles, which is common during compelling presentations.
Thus, for 400 users, the design calls for 8 radio interfaces operating at 5 GHz, calculated as follows: 400 users divided by 50 users per interface equals 8 interfaces.
| Use Cases vs. RF Profile Recommendations for 5GHz | Tx power | Channel Width | Data Rate | RX-SOP | 
| Classroom/Conference Rooms | Min= 7dBm Max= 30dBm | 40 MHz * | 24 Mbps | -78 | 
| Corporate Training Room/ Medium Auditorium | Min= 7dBm Max= 30dBm | 40 MHz * | 24 Mbps | -78 | 
| Big Events/ Keynotes | Min= 10dBm Max= 12dBm | 20 MHz | 24 Mbps | -78 | 
* A channel bandwidth of 40 MHz is considered in the context of channel plans and regulations. See Typical RF Design and Spatial Reuse guidelines above.
