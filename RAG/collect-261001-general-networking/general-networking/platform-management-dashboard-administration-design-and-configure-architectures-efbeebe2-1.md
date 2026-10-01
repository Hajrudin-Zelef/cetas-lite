---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-efbeebe2-1
title: "platform-management-dashboard-administration-design-and-configure-architectures--efbeebe2"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "throughput"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--efbeebe2.md
source_anchor: ""
source_lines: [1, 80]
sha256: 6285082555a86a329e491175077afa11e90ef57c2326a759e8ed58dae5499fb2
---

# platform-management-dashboard-administration-design-and-configure-architectures--efbeebe2

Auto VPN Hub Deployment Recommendations
About This Document
This document provides recommendations for AutoVPN hub deployments.
These recommendations and the suggested deployment configurations have been collected across the Meraki MX install base (covering hundreds of thousands of AutoVPN sites) and have been vetted by the Meraki MX product team.
While this document provides a high level overview and emphasizes important considerations, please refer to the online documentation for specific details on how to implement the suggested configurations. Meraki’s 24x7 Support is also available to assist as needed.
The Meraki MX AutoVPN technology is versatile and supports many configuration options that are used to address different use cases - many of these are not mentioned here.
This document covers the most popular, common and robust AutoVPN hub deployment options.
Datacenter Hub Deployment
Summary
- 
    Use a hub and spoke topology.
- 
    Use OSPF if dynamic routing is required. (an additional router can be used for BGP redistribution)
- 
    Turn off all non-VPN features. (e.g. Security features, Traffic Analytics)
- 
    It is always better to re-IP than to use NAT translation of any sort.
Overview
In order to achieve the maximum possible scale for a Meraki AutoVPN deployment, there is really only one topographical choice - Hub and Spoke (H&S). This is due to the large number of tunnels a full mesh solution would incur. The formulae for working out the likely total tunnel count and individual MX tunnel count for both support topologies are as follows:
Hub and Spoke - Total Tunnel Count
Where H is the number of hubs, S is the number of spokes, and L is the number of uplinks the MX has (LH for the hubs, LS for the spokes). If each MX has a different number of uplinks, then a sum series, as opposed to a multiplication, will be required.
For example, if all MXs have 2 uplinks (both WAN1 and WAN2 active), and if we have 4 hubs and 100 spokes, then the total number of VPN tunnels in the organization would be 24 + 1600 = 1624.
H = 4, S = 100, LH= 2, and LS = 2
All appliances in this example have two uplinks, so LH = LS = 2. For the hubs, this works out to (4 x (4-1) x 22)/2 = 24. The number of tunnels for all 100 spokes is (4 x 100 x 2 x 2) = 1600.
Hub and Spoke - Hub Tunnel Count
To continue our example, each hub would have a total of 12 tunnels to the other hubs and 400 tunnels to the spokes for a total of 412 tunnels per hub MX.
Hub and Spoke - Spoke Tunnel Count
To complete our example, each MX spoke will have 4 AutoVPN tunnels established to each MX hub for a total of 16 tunnels. That is, each spoke has 4 tunnels to each hub: WAN1-WAN1, WAN1-WAN2, WAN2-WAN1 and WAN2-WAN2, and for four hubs that is 16 tunnels per spoke.
Full Mesh - Total Tunnel Count
Where H is the number of MXs and L is the number of uplinks each MX has.
For example, if all MXs have 2 uplinks and there are 50 MXs, then the total number of VPN tunnels would be 4900.
Full Mesh - Tunnel Count per MX
To complete the example every MX would have to be able to support 196 tunnels, in this case, we would need around 50 MX100s.
There are, however, multiple ways in which we can architect the H&S network such that we achieve greater flexibility. We will illustrate each of these models below.
Standard Hub & Spoke
How it works
Utilizing the standard Meraki AutoVPN registry to ascertain how the VPN tunnels configured need to form (i.e. via public address space or via private interface address space) as described in Configuring Site-to-site VPN over MPLS.
When should it be used?
Whenever possible is the short answer. It is strongly recommended that this model is the 1st, 2nd and 3rd option when designing the topology. Only if the customer has an exceptionally strong requirement should one of the following H&S derivatives be considered.
Note: AutoVPN hubs should not be added to templates at all. It is not possible to configure an MX as a spoke with exit hub that is part of a template.
Configuration - VPN Concentrator
Whilst the high-level configuration on a VPN is relatively straightforward, there are a number of potential pitfalls that will be covered here. If more information is required please refer to the definitive guide - VPN Concentrator Deployment Guide.
Global Configuration
It is strongly recommended that all MX AutoVPN hubs are dedicated hubs. As such, the Addressing & VLANs page should look like this:
VPN Configuration
From the Site-to-site VPN page, we need to set the type to Hub (Mesh), as shown below:
Hub means form a VPN tunnel to everyone who is also a Hub and any spoke that has you configured as a hub. Whereas Spoke means to just VPN to the MXs you have configured as Hubs.
Exit Hub
When an MX is configured as a Hub, then an additional config option becomes available: ‘Exit hub’. This allows you to bind a default route (0/0) to the IPSec security association of that hub in a similar fashion to the ‘Default Route’ option for Spoke MXs.
Local Network Advertisement
It should be known that networks that are accessible from the concentrator MX in the data center and need to be advertised to other hubs and/or spokes MXs need to be defined and advertised. As shown below:
In the rare instances that the locally available subnets are not globally unique in the IP schema of the AutoVPN domain, then VPN translation can be enabled such that the entire locally available range can be translated to a unique range, as shown below:
These options are only to be used in emergencies, as the best solution is always to re-IP the offending range such that duplicates do not exist.
Best practice - it is generally recommended to summarize continuous addressing blocks whenever possible (e.g. 10.0.0.0/8).
For reference, below are the RFC1918 private address blocks:
- Class A - 10.0.0.0/8
- Class B - 172.16.0.0/12
- Class C - 192.168.0.0/16
Any additional, more specific subnets contained within these supernets that are available via the advertising hub can/should also be advertised too to affect prioritization among routes. This also extends to non-RFC1918 traffic that is publicly routable that is accessible via the AutoVPN domain.
In the case where more complex routing is needed, please refer to the MX routing behavior document for more information.
If, as per the above, more than one hub is advertising the same subnet or supernet address ranges, then the priority in which those routes are used by other hub MXs is configured in the Organization-wide settings section, as per the below:
Note: On MX-Z devices, traffic for the following services/tools will adhere to the route priority outlined in our MX Routing Behavior article
- 
    Advanced Malware Protection registration
- 
    Meraki Cloud Authentication
- 
    Meraki Cloud Communication on TCP ports 80, 443, and 7734
- 
    Ping and Dashboard Throughput Live Tools
- 
    List Updates for the following services 
  - 
        Content Filtering
  - 
        IDS/IPS Rule Updates
  - 
        Geo-IP Lists for Layer 7 Country-Based Firewall Rules
- 
        
