---
id: collect-261001-general-networking/general-networking/wireless-design-and-configure-deployment-guides-mesh-deployment-guide-024d5c2c-2
title: "wireless-design-and-configure-deployment-guides-mesh-deployment-guide-024d5c2c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "throughput"]
source: docs/RAG/collect-261001-general-networking/wireless-design-and-configure-deployment-guides-mesh-deployment-guide-024d5c2c.md
source_anchor: ""
source_lines: [52, 74]
sha256: 17f3dc41f3a4d1753e6ba08a21330c290d92d2d63c959eb530433fc42d41042f
---

# wireless-design-and-configure-deployment-guides-mesh-deployment-guide-024d5c2c

The gateway access point may be connected to a trunk port and trunk SSIDs to different VLANs. Repeater access points will broadcast SSIDs trunked on different VLANs too. Only one SSID and associated VLAN, however, may be configured to bridge wired clients across a mesh link on a repeater access point's Ethernet port. A mix of wired clients and Meraki access points attached to one repeater access point Ethernet port is not a supported deployment configuration. Meraki access points use auto detection mechanisms to infer when they should function as a gateway or a repeater, which is why a mix of wired clients and Meraki access points is not allowed.
Mesh wired access may be treated like a traditional point-to-point link with a router on the remote site. Meraki access points may be connected to the repeater side when a Layer 3 device is separating the broadcast domains. For further discussion on this design, see Extending the LAN with a Wireless Mesh Link.
Wired clients are not subject to the same authentication requirements that wireless clients are subject to. Wired clients will bypass authentication methods such as PSK and RADIUS and gain network connectivity as though they had associated to the SSID.
Wireless Mesh Data Rate
Wireless data rate selection is an important mechanism for effective use of the available RF spectrum. The data rate can affect the throughput of clients. Throughput is an important metric used by industry publications to evaluate vendor devices.
Dynamic Rate Adaptation (DRA) introduces a process to estimate the optimal rate for packet transmissions. Correctly selecting rates is important as a too high rate leads to packet transmissions failing, which leads to communication failure. If the rate is too low, the available channel bandwidth is not used efficiently, creating the potential for network congestion and collapsed links. Meraki access points use a customized, fully-automatic DRA algorithm when establishing mesh links.
Frequency/Radio Usage
Any channel that a Meraki access point is permitted to operate on can be used for a mesh link. The channel availability for a particular access point model is subject to regulatory domain restrictions and certification. Meraki access points do not prefer one band over another, as described in the Wireless Mesh Networking "Meraki Mesh Algorithm" section. Both radios on a Meraki access point may be used concurrently for meshing while also serving wireless clients.
Any given mesh backhaul link will only use one radio, either 2.4 or 5 GHz, but not both concurrently. The gateway and/or repeater access point can serve clients and provide mesh connectivity. It is recommended not to use mesh links/radios to serve clients but rather have them remain dedicated for mesh to help maximize performance.
An access point in repeater mode may not always honor manual channel settings. More details on manually influencing the mesh channel can be found in Manually Changing Channels in a Mesh Network.
DFS Recommendation
It is recommended to avoid Dynamic Frequency Selection (DFS) channels using the "Exclude DFS channels" option on the Wireless > Configure > Radio settings page for greatest reliability. A DFS event will cause access points to silence communication on the affected channel and temporarily move to another channel as described in Dynamic Frequency Selection.
Mesh Convergence Time
The Meraki mesh algorithm is designed for stationary Meraki access points with variable links. The algorithm sends out periodic discovery frames of varying sizes on the 2.4 GHz and 5 GHz data radios to discover mesh neighbors. More discussion on neighbor discovery and gateway selection can be found in Wireless Mesh Networking. After one neighbor is selected as the gateway, new routes are passively identified on the same operating channel. A mesh link may take a couple minutes to establish and may not always be used for every data flow if a new route with better link metrics becomes available.
The Meraki mesh algorithm is not optimized for access points that are moving.
Multi-VLAN Support Over Mesh
Meraki APs will allow traffic from multiple VLANs over mesh links. This feature can be enabled by contacting support. There are a couple of conditions that must be met to support this functionality:
1. Clients wired directly into Meraki access points needs to be enabled and configured for a specific SSID where multiple VLANs are used. This option is found on the Network Wide > Configure > General page.
2. SSID configuration has to use Bridge mode. This option is found on the Wireless > Configure > Access Control page, Client IP assignment section.
Multi-VLAN support over Mesh is supported with MR 28.1 and higher firmware versions.
If you plan to have one or more Meraki access points behind a repeater access point, a layer 3 device needs to separate them as stated in Extending the LAN with a Wireless Mesh Link.
Multicast Over Mesh
Multicast protocols, including CDP, LLDP, VTP, etc., are not supported over a wireless mesh link.
