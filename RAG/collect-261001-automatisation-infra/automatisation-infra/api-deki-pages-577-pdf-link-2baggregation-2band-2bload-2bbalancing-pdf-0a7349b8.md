---
id: collect-261001-automatisation-infra/automatisation-infra/api-deki-pages-577-pdf-link-2baggregation-2band-2bload-2bbalancing-pdf-0a7349b8
title: "api-deki-pages-577-pdf-link-2baggregation-2band-2bload-2bbalancing-pdf-0a7349b8"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution", "throughput"]
source: docs/RAG/collect-261001-automatisation-infra/api-deki-pages-577-pdf-link-2baggregation-2band-2bload-2bbalancing-pdf-0a7349b8.md
source_anchor: ""
source_lines: [1, 60]
sha256: 819abf348e98da52c6e9b7717ae97673c37160fc5b9e7ee4779fcdaed2885274
---

# api-deki-pages-577-pdf-link-2baggregation-2band-2bload-2bbalancing-pdf-0a7349b8

Link Aggregation and Load Balancing
Link aggregation is a term used to describe various implementations and underlying technologies. In general, link aggregation combines multiple network
connections in parallel to increase throughput and provide redundancy. While there are many approaches, this article aims to highlight the differences
in terminology.
Link bonding
Link bonding, also known as teaming or bundling depending on the vendor, is generally implemented using two or more links between two logical devices. This
could be two servers, two switches, a server to a switch, or various other combinations. Using standards such as Link Aggregation Control Protocol (LACP), the
two links are combined into a single logical link, with traffic spread across them evenly. Because link bonding operates at Layer 2, failure detection and isolation
can be done quickly, limiting the impact of a link failure. Link bonding is also useful for increasing available throughput between two devices without purchasing
more expensive hardware, for example, two 1 Gbps links versus one 10 Gbps link. Refer to Figure 1 below.
Figure 1
Load balancing
Load balancing can also be used to describe link bonding. The term is reserved for Layer 3+ operations. While application load balancers can distribute load
across an array of devices for a particular application or purpose, this article focuses on Layer 3. In that context, load balancing is defined as an even distribution
of IP traffic across two or more links. This can be done by providing a device multiple equal cost routes to the same destination over equal sized links. Refer to
Figure 2 below.
Figure 2
Load sharing
Load sharing is defined as spreading network traffic across two or more equal or unequal links or paths. Load sharing can exist between a 10 Mbps WAN link
and a 100 Mbps WAN link. While load sharing often provides the slowest recovery time, dependent on implementation and failure, it is the easiest to implement,
most flexible, and still provides levels of redundancy that link bonding and load balancing cannot. As an example, load sharing allows the use of two different
ISPs with different link speeds when Network Address Translation (NAT) is implemented. If NAT is not used with other vendor products, this becomes more
1

difficult and may outbound traffic traverse one link while inbound traffic for the same conversation uses another. Refer to Figure 3 below.
Figure 3
Implementation by Cisco Meraki
MS/CS Series
Cisco Meraki MS switches use the open standard LACP to provide Layer 2 link aggregation in the form of link bonding as described above. The MS LACP
hashing algorithm uses traffic source and destination IP, MAC, and port to determine which bonded link to use. This provides highly resilient and equal load
distribution across two or more links, between two logical devices, with rapid failure detection.
Refer to the Switch Ports for details on how to implement.
When configuring LACP port-channels on uplinks between remote switches or stacks, always configure the remote or downstream side of the port-channel first.
Once the configuration has been applied, configure the LACP port-channel on the upstream switch. This insures the downstream switch is not left without
connectivity.
The default behavior of Catalyst Switches (C9K and MS390) is to disable a LACP member interface for 30 seconds if it detects a regular trunk port on the other
end of a configured port-channel. To mitigate this, ensure "Port Channel Fallback (Standalone Mode Behavior)" is set to "Enable" on the Switch->Settings page.
This adds 'no port-channel standalone-disable' to the port-channel configuration on all port-channels and aggregate ports across the organization. This feature is
primarily needed on the "upstream"/core switch, not the downstream switch or stack. This allows you to configure the AGG port in any order and/or replace the
downstream device without having to reconfigure the core.
Link Aggregation member ports must operate at the same speed and duplex. It is recommended to use member ports with similar characteristics, including
media type where possible, for example, SFP with SFP or copper with copper.
For Catalyst switches Port Channel Fallback (Standalone Mode) behavior is enabled by default on all new networks deployed on IOS-XE 17.15.5 or
later. For existing networks, this feature is not enabled automatically and must be enabled manually.
For Meraki MS Switches, this Port Channel Fallback behavior is the default behavior.
Refer to the Port-Channel Failback Mode.
2

MX WAN appliance
Cisco Meraki WAN appliances use a proprietary algorithm to provide load balancing across two Layer 3 links (if configured). This can be customized to use
different ratios and specific rules for outbound traffic. As NAT is used, flows that are part of a particular conversation remain on the link they are placed on.
Refer to MX Load Balancing and Uplink Preferences for details on how to implement.
Configuring link aggregation between MS and Cisco Switches
You may want to set up and configure a bonded link between your Meraki MS switch and a Cisco switch. This is often referred to as link aggregation, link
bonding or EtherChannel.
To configure two or more ports, up to eight, as a port aggregate, navigate to Switching > Monitor > Switch ports and select the target ports, then choose
"Aggregate". It is recommended that the target ports physically connected to anything during this step.
On your Cisco switch, enable LACP by setting the EtherChannel mode to active or passive depending on the behavior required. For further information, refer to
Cisco's documentation on Configuring EtherChannel (this document is for the Catalyst 3000 series).
3
