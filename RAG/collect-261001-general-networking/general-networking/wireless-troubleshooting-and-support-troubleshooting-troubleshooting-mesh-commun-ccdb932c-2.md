---
id: collect-261001-general-networking/general-networking/wireless-troubleshooting-and-support-troubleshooting-troubleshooting-mesh-commun-ccdb932c-2
title: "wireless-troubleshooting-and-support-troubleshooting-troubleshooting-mesh-commun-ccdb932c"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["omni", "regulation"]
source: docs/RAG/collect-261001-general-networking/wireless-troubleshooting-and-support-troubleshooting-troubleshooting-mesh-commun-ccdb932c.md
source_anchor: ""
source_lines: [62, 88]
sha256: 679fe972719161159519bfeb3ad07ac42cd6c0dd754a2a3ff9a7162e5e541f12
---

# wireless-troubleshooting-and-support-troubleshooting-troubleshooting-mesh-commun-ccdb932c

Figure 9: Two APs manually configured to use channel 36 with a transmit power of 5dBm. The transmit power level could be an issue if the nodes are too far apart, but if they are close enough they should become mesh neighbors.
Note: Dashboard will calculate the distance between neighboring APs when placed on Google map. The distance is based on the lat/long (from Dashboard) of a mesh repeater and the neighbors it sees.This distance is then reported in the Dashboard mesh neighbors table.
Dynamic Frequency Selection (DFS)
In certain regulatory domains, 5GHz WiFi channels are shared with radar. DFS regulation in these environments require that an AP change it's channel whenever it detects a radar pulse.
Channel changes, including those due to DFS, will bring down a mesh link. Therefore if 5GHz is being used for mesh, non-DFS channels are preferred. However, this may not be an option depending on the country of operation (regulatory domain). Therefore, always check for DFS as a possible cause of mesh link instability.
The Channel planning section of the Radio settings page has the regulatory domain of the network, which is based on the selected Country, and an option to allow DFS channels.
Figure 10: A network that allows DFS channels. APs in this network could be subject to automatic channel switching.
DFS channel change events are reported in the Dashboard Event log. If you suspect DFS could be causing the mesh link to drop, the Event log is a good place to check.
RF spectrum
If an AP has a dedicated Air Marshal radio, it can report real-time RF information in Dashbord under Wireless > Monitor > RF spectrum. Check this page to find any indications of high channel utilization. If there is high utilization on certain channels, a manual channel adjustment may be necessary.
Physical Inspection
If an offline repeater is not showing up as a mesh neighbor of nearby APs, there may be a physical issue that requires on-site troubleshooting. In particular, it may be necessary to spot check mesh APs for signs that something in the deployment has changed or was installed incorrectly.
Check the LED
Monitor the LED of the offline node and note the color and activity patterns. The MR installation guide has the LED code definitions.
- If there are no LED, does the AP have a working power source?
- If there are no LED, try resetting the offline unit. After a reset, if there is still no LED activity, verify the offline AP has a working power source and check the cabling and hardware for damage.
- If the LED shows the offline node in a scanning state, it may be unable to find a path to the gateway AP. This could be attributed to antenna placement or orientation which should be checked on both ends of the link.
Check the Antenna (outdoor APs)
Outdoor MR units use an externally attached antenna. If installed incorrectly, this could be the cause of poor/unreliable mesh links.
- Confirm the antenna type and radio band being used is correct on both ends of the link.
- Confirm antennas are attached securely to the correct radios, i.e. the 5GHz antenna is connected to the 5GHz radio posts or that an antenna is installed at all.
- Confirm antennas are oriented correctly based on their coverage patterns. Meraki antenna coverage patterns are in the datasheets for each antenna model. For example, if a directional antenna is being used, make sure it is aiming at its target with direct line of sight. If an omni antenna is used, verify it is at the correct elevation.
Do not attempt to establish a mesh link between outdoor APs without external antennas installed, even if they are a few feet apart, the communication may not be reliable.
Packet Captures
When troubleshooting a mesh link, packet captures may be required in order to analyze traffic between APs. These promiscuous-mode packet captures must be taken using external monitoring stations to effectively capture all nearby wireless traffic.
A monitor should be placed near each AP, and configured to capture packets on the channels being used for the mesh link. Once the packet captures been gathered, they can be used to identify the issue, or provide helpful troubleshooting information for Cisco Meraki Support.
Figure 11: In this example, the gateway and repeater on are channel 1. Therefore, a monitor is placed near each AP and configured to capture on channel 1.
