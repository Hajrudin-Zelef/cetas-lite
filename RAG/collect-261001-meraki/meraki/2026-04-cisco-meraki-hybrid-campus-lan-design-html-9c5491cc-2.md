---
id: collect-261001-meraki/meraki/2026-04-cisco-meraki-hybrid-campus-lan-design-html-9c5491cc-2
title: "2026-04-cisco-meraki-hybrid-campus-lan-design-html-9c5491cc"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution", "license", "voice"]
source: docs/RAG/collect-261001-meraki/2026-04-cisco-meraki-hybrid-campus-lan-design-html-9c5491cc.md
source_anchor: ""
source_lines: [84, 147]
sha256: 0e6c9716921e744f693a8dd72edd0f09adf3f9bbfce1cb6233798b71c25c5b7d
---

# 2026-04-cisco-meraki-hybrid-campus-lan-design-html-9c5491cc

In Layer 3 Access design, the IP routing boundary is pushed all the way down to the access layer. Instead of a flat Layer 2 STP domain spanning the entire campus, each access switch or stack terminates VLANs locally and routes traffic upward via Layer 3 links. This eliminates the non-deterministic nature of STP and replaces it with fast, predictable routing protocol convergence.
🏆 Why Layer 3 Access is Best for Large Deployments
⚡ Fast Convergence
Routing protocols converge faster and more predictably than STP in large networks.
🔒 Better Segmentation
Routed boundaries naturally contain broadcast domains and improve security posture.
📈 Better Scalability
No STP topology changes ripple across the entire campus — faults remain local.
🎯 Deterministic
Equal-cost multipath (ECMP) routing provides predictable load sharing and failover.
In this design, all SVIs (Switch Virtual Interfaces) are created at the collapsed core layer. The access switches act as pure Layer 3 forwarders, and wireless roaming across closets requires a Layer 3 Roaming Concentrator to maintain session continuity when clients move between access switch domains.
💡 Wireless Roaming Consideration: Layer 3 Access does not natively support seamless wireless roaming across campus zones without additional configuration. A Layer 3 Roaming Concentrator (tunnel-based roaming) must be deployed to anchor wireless clients as they move between different IP subnets served by different access switch stacks.
9. Security: 802.1x, Cisco ISE & Adaptive Policy
Security is deeply embedded throughout this CVD. The architecture supports both wired and wireless 802.1x authentication via RADIUS, integrated with Cisco ISE (Identity Services Engine) for policy enforcement, posture checking, and dynamic VLAN assignment.
802.1x & MAB Authentication
Access policies are configured through the Meraki Dashboard (Switching → Configure → Access Policies). The CVD validates two modes:
🔐 802.1x — EAP Authentication
Used for corporate devices supporting EAP. Users authenticate against Cisco ISE using credentials. ISE returns dynamic VLAN assignment and Security Group Tags (SGT) via RADIUS. CoA (Change of Authorization) is enabled to push policy updates post-authentication.
🔑 MAB — MAC Authentication Bypass
Fallback method for devices that don't support 802.1x (printers, IoT devices). The device MAC address is passed to ISE as credentials. ISE policy determines appropriate VLAN and access level based on the MAC address database.
Wireless SSID Design & ISE Integration
The CVD validates two SSIDs, each mapped to a distinct VLAN and ISE policy:
Adaptive Policy (SGT — Security Group Tagging)
The CVD also validates Adaptive Policy (Cisco's Meraki implementation of TrustSec SGTs) for micro-segmentation. Groups are defined in Meraki Dashboard (Organization → Configure → Adaptive Policy), and SGT values are assigned per user group. This allows policy-based access control independent of VLANs or IP addresses — making segmentation portable and scalable.
10. Quality of Service (QoS) Design
A properly designed QoS policy is essential for supporting real-time applications such as voice, video conferencing (Webex), and critical enterprise applications on a shared campus LAN. The CVD validates QoS end-to-end across both wired and wireless infrastructure.
QoS is configured via Meraki Dashboard (Wireless → Configure → Firewall & Traffic Shaping). Key DSCP markings validated in this CVD include:
✅ Validated Scenario: Wireless roaming was tested between two campus zones and APs homed to different switch stacks while maintaining an active Webex meeting with Audio, Video, and Content Share — confirming that QoS policies hold through roaming transitions and that voice/video quality is maintained.
11. High Availability & Redundancy Design
High availability is a fundamental design requirement for any enterprise campus LAN. The CVD incorporates HA principles at every layer of the architecture:
🔁 MX250 Warm-Spare (WAN Edge)
The MX250 WAN edge appliance is deployed in warm-spare (active/passive) configuration, requiring only a single license for both appliances. In the event of the primary MX failure, the secondary takes over seamlessly with no manual intervention.
🔗 Link Aggregation (EtherChannel / LACP)
LAG (802.3ad) is deployed between access switches and the distribution/core layer, providing higher bandwidth, load sharing, and link-level redundancy without relying on STP to handle failures.
🏗️ Redundant Triangles (Access-to-Core)
Best practice design uses redundant triangle topologies rather than redundant squares. Each access stack connects to two distribution/core switches via dual equal-cost paths, enabling fast failover without spanning tree loops.
⚡ StackPower for Access Switches
Cisco StackPower shares power budgets across stacked access switches, providing power redundancy so that a single power supply failure does not bring down the entire stack. Redundant power supplies are recommended across all campus switching layers.
12. Planning & Best Practices Checklist
Use this pre-deployment checklist to ensure your Hybrid Campus LAN implementation follows the validated design guidelines from this CVD:
☁️ Cloud & Platform
Determine if access switches will run in Cloud Managed (Meraki Dashboard full control) or Cloud Monitored (visibility only) mode
Ensure Catalyst switches meet the minimum firmware version 17.3.1 before onboarding to cloud monitoring
Verify TCP port 443 outbound internet access is available from all SVIs on Catalyst switches
Confirm valid DNS configuration on all Catalyst switches to be cloud-monitored
🏗️ Design & Topology
Select the appropriate logical design option (L2 VLAN1, L2 No VLAN1, or L3 Access) based on your scale, security, and roaming requirements
Standardize on MST (802.1s) as the STP protocol across all Meraki and Catalyst platforms in the same STP domain
Deploy redundant triangles from access stacks to core, not redundant squares
If using Layer 3 Access, plan for a Layer 3 Roaming Concentrator to support seamless wireless mobility
🔐 Security
Configure 802.1x and MAB access policies in Meraki Dashboard for both wired and wireless clients
Integrate Cisco ISE for RADIUS authentication, dynamic VLAN assignment, SGT, and posture checking
Enable RADIUS CoA to allow ISE to dynamically change client authorization post-connection
Enable STP BPDU Guard on access ports to protect the STP topology from unauthorized devices
📡 QoS & Wireless
Configure DSCP markings end-to-end (EF for VoIP, AF41 for video, AF11 for critical data)
Test wireless roaming across multiple zones with an active real-time application (video call) before go-live
Deploy WiFi 6 (802.11ax) access points for high-density environments and future-proof wireless coverage
13. Conclusion
The Cisco Meraki Hybrid Campus LAN CVD represents one of the most comprehensive and practically validated design guides available for enterprise network architects. By combining the operational simplicity of the Meraki Dashboard with the raw performance and advanced security of the Catalyst 9000 series, organizations can build campus networks that are simultaneously easy to manage, secure, scalable, and ready for the future.
Whether you choose Layer 2 with MST for maximum VLAN flexibility, the security-hardened no-VLAN-1 design, or the Layer 3 Access model for deterministic convergence and scale — this CVD provides a pre-validated path with documented testing results, verified configurations, and clear design trade-offs.
The validated integration of Cisco ISE, Adaptive Policy/SGT, QoS, WiFi 6, and Cloud Monitoring makes this design a complete enterprise-grade campus blueprint — not just a reference diagram. For organizations modernizing their campus networks, this CVD is the starting point, not the destination.
Want to go deeper into Cisco Campus Design?
