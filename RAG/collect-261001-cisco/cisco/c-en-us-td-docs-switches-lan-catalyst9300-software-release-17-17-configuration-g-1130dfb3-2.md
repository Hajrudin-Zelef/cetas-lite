---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-1130dfb3-2
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-1130dfb3"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "latency", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-1130dfb3.md
source_anchor: ""
source_lines: [4, 51]
sha256: b7e3d3523c297bb545d1f024e7eb8639834616ead3051cc54573c9f75337a596
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-1130dfb3

                                 Apply the Access control list (ACL) filter before sending the monitored traffic on to the tunnel.
The following restrictions apply for this feature:
Truncation is supported only on IPv4 and IPv6 spanned packets and not on Layer 2 packets without an IP header.
An ERSPAN destination interface can be part of only one session. The same destination interface cannot be configured for multiple ERSPANs/SPANs.
You can configure either a list of ports or a list of VLANs as a source, but cannot configure both for a given session.
Filter IP/IPv6/MAC/VLAN access-group and filter SGT cannot be configured at the same time.
When a session is configured through the ERSPAN CLI, the session ID and the session type cannot be changed. To change them, you must use the no form of the commands to remove the session and then reconfigure it.
ERSPAN source sessions do not copy locally-sourced RSPAN VLAN traffic from source trunk ports that carry RSPAN VLANs.
ERSPAN source sessions do not copy locally-sourced ERSPAN Generic routing encapsulation (GRE)-encapsulated traffic from source ports.
Disabling the ip routing command for IPv4 connections and ipv6 unicast-routing command for IPv6 connections stops ERSPAN traffic flow to the destination port.
ERSPAN over MPLS VPN is supported on Layer 3 VPNs, Segment Routing and Seamless MPLS.
ERSPAN over MPLS VPN is not supported for L2VPN, 6PE, 6VPE, MPLS over GRE and InterAS.
You cannot configure an MPLS core switch as an ERSPAN destination. ERSPAN traffic can be transported from one Provider Edge (PE) to another Provide Edge (PE). It cannot be transported to the core switch between the two Provider Edges.
ERSPAN sessions do not capture DHCP-inject packets. DHCP-inject packets refer to DHCP packets (DISCOVER, OFFER, REQUEST, and ACK packets) which are modified by the CPU and inserted back into the network.
If a backup configuration having ERSPAN session enabled is restored to the running configuration, ERSPAN sessions are created automatically in disabled state. You must manually enable these ERSPAN sessions.
ERSPAN does not support QoS shaping and policing.
The following sections provide information about configuring ERSPAN.
The Cisco ERSPAN feature allows you to monitor traffic on ports or VLANs, and send the monitored traffic to destination ports. ERSPAN sends traffic to a network analyzer, such as a Switch Probe device or a Remote Monitoring (RMON) probe. ERSPAN supports source ports, source VLANs, and destination ports on different devices, which help remote monitoring of multiple devices across a network.
ERSPAN supports encapsulated packets of up to 9180 bytes. ERSPAN consists of an ERSPAN source session, routable ERSPAN GRE-encapsulated traffic, and an ERSPAN destination session.
You can configure an ERSPAN source session, an ERSPAN destination session, or both on a device. A device on which only an ERSPAN source session is configured is called an ERSPAN source device. A device on which only an ERSPAN destination session is configured is called an ERSPAN termination device. A device can act as both; an ERSPAN source device and a termination device. To avoid over-subscription of traffic, which can lead to drop in management traffic on the destination device, ensure that the destination session is configured and is working on the destination device, before configuring a source session on the source device.
For a source port or a source VLAN, the ERSPAN can monitor the ingress, egress, or both ingress and egress traffic. By default, ERSPAN monitors all traffic, including multicast, and Bridge Protocol Data Unit (BPDU) frames.
A device supports up to 66 sessions. A maximum of eight source sessions can be configured and the remaining sessions can be configured as RSPAN destinations sessions. A source session can be a local SPAN source session or an RSPAN source session or an ERSPAN source session. The number of source sessions decreases by the number of configured ERSPAN destination sessions.
A device can support a maximum of 50 Security Group Tag (SGT) filter per session.
An ERSPAN source session is defined by the following parameters:
A session ID.
ERSPAN flow ID.
List of source ports or source VLANs that are monitored by the session.
Optional attributes, such as, IP type of service (ToS) and IP Time to Live (TTL), related to the Generic Routing Encapsulation (GRE) envelope.
The destination and origin IP addresses. These are used as the destination and source IP addresses of the GRE envelope for the captured traffic, respectively.
| Note |  | 
The Cisco ERSPAN feature supports the following sources:
Source ports—A source port that is monitored for traffic analysis. Source ports in any VLAN can be configured and trunk ports can be configured as source ports along with nontrunk source ports.
Source VLANs—A VLAN that is monitored for traffic analysis.
A destination port is a Layer 2 or Layer 3 LAN port to which ERSPAN source sends traffic for analysis.
When you configure a port as a destination port, it can no longer receive any traffic. The port is dedicated for use only by the ERSPAN feature. An ERSPAN destination port does not forward any traffic except that required for the ERSPAN session. You can configure trunk ports as destination ports, which allows destination trunk ports to transmit encapsulated traffic.
A Security Group Tag (SGT) is a 16-bit value that the Cisco Identity Services Engine (ISE) assigns to the user or endpoint session upon login. The network infrastructure views the SGT as another attribute to assign to the session and inserts the Layer 2 tag to all traffic from that session. A platform can support a maximum of 50 SGT policies per session.
On an existing flow-based SPAN (FSPAN) or VLAN filter session, SGT filtering configurations are not allowed.
ERSPAN Timestamp is automatically enabled when the ERSPAN header is set to type III. The timestamp field is used to calculate packet latency in devices. The ERSPAN source session fills in the timestamp field with local time information when a packet is received. The destination session can hand over this timestamp to the application. ERSPAN supports all timestamps in 32-bit format. It supports 100 nanosecond (ns) granularity and the timestamp field wraparound time is around 7 minutes.
Starting with the Cisco IOS XE Bengaluru 17.5.x release, you can transport ERSPAN traffic over a Multiprotocol Label Switching (MPLS) Virtual Private Network (VPN). To enable ERSPAN over MPLS VPN you will have to enable Multiprotocol Label Switching (MPLS), Label Distribution Protocol (LDP), and Cisco Express Forwarding in your network.
You can configure the ERSPAN destination to select the source VRF for the ERSPAN traffic over MPLS VPN. You can use the vrf keyword in the ERSPAN destination session source command to configure the source VRF.
The following sections provide information about how to configure ERSPAN.
The ERSPAN source session defines the session configuration parameters and the ports or VLANs to be monitored. To define an IPv4 ERSPAN source session, complete the following procedure:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | monitor session span-session-number type erspan-source Example: Device(config)# monitor session 1 type erspan-source | Defines an ERSPAN source session using the session ID and the session type, and enters ERSPAN monitor source session configuration mode.  | 
| Step 4 | description string Example: Device(config-mon-erspan-src)# description source1 | (Optional) Describes the ERSPAN source session.  | 
