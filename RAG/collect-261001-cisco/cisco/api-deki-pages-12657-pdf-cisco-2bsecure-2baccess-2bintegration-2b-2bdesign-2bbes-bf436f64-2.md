---
id: collect-261001-cisco/cisco/api-deki-pages-12657-pdf-cisco-2bsecure-2baccess-2bintegration-2b-2bdesign-2bbes-bf436f64-2
title: "api-deki-pages-12657-pdf-cisco-2bsecure-2baccess-2bintegration-2b-2bdesign-2bbes-bf436f64"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-cisco/api-deki-pages-12657-pdf-cisco-2bsecure-2baccess-2bintegration-2b-2bdesign-2bbes-bf436f64.md
source_anchor: ""
source_lines: [94, 181]
sha256: 02deee7469174e9ce3ff4877be2d1a19d8d38dd89ed6600a1408bde9c7287c1b
---

# api-deki-pages-12657-pdf-cisco-2bsecure-2baccess-2bintegration-2b-2bdesign-2bbes-bf436f64

Enabling Secure Access Meraki Integration in a Maintenance Window
Once you are in your scheduled maintenance window, complete the following steps:
1. Turn on the integration by completingthese steps
2. Onboard the hubs & spokes to SSE by completingthese steps
3. If you have SIA onlyand desire Spoke-Spoke communication,call support to have them disable the feature after you have onboarded your hubs &
spokes to Secure Access.
• You can reference this document or tell them that you “desire Spoke-Spoke communication” reinstated on your organization.
4.Optional: Organizations managing a large number of sites (over 500) enrolled in a single Secure Access region can enable theonly-default-route for Spokes
feature described in the Platform Optimization section of this document. Customerscan request to enable this feature when they contact Meraki support.Learn
more here.
Design A - Cloud Spokes and Cloud Hubs
This design supports both Secure Internet Access (SIA) and Secure Private Access (SPA) use cases. It provides a straightforward, cloud‑based SD‑WAN
transport model in which all MX networks operate as Cloud Spokes.
All customer branch spokes are enrolled in Cisco Secure Access, and no MX hubs are deployed. All traffic-spoke‑to‑spoke, remote worker to site, and endpoint
to IPSec‑attached applications-flows through the cloud path. Because MX hubs are removed from the design, it can scale to a very large number of sites,
making it particularly suitable for environments such as retail deployments.
please ensure to contact support to intervene after onboarding and again disabletheMX Hub mesh.
Design C is the only scenario that supports disabled MX Hub mesh. If necessary, you can do so in the
maintenance window.
UsetheCall Me Nowfeature in the Meraki dashboard to request an immediate phone call.
Phone calls guarantee faster response times and can deploy these changes during your maintenance window.
5

Is this Design for me? - Customer Traffic Profile
Organizations whose traffic is primarily north‑south toward the internet align well with this architecture. East‑west traffic is also carried through the cloud fabric
and inspected by the cloud firewall, consolidating traffic into a single security and routing domain.
Organizations with 500 or more sites within a region can benefit significantly from this approach by reducing the number of hub connections required and
minimizing route scale at spokes. This design also appeals to customers who prefer not to purchase, host, or maintain hub hardware.
All MX appliances function as Cloud Spokes connected to the cloud fabric. Remote workers can reach any spoke‑advertised prefixes through the cloud. Each
Spoke can select the nearest Cloud Hub as its primary, and all Cloud Hubs automatically form a mesh within Secure Access.
Design A - Pros
• Simplified SD‑WAN architecture with minimal configuration
• Supports both SIA and SPA use cases
• Small routing tables at spokes, with the cloud fabric providing full prefix reachability
• No MX hub hardware to deploy or maintain
• Cloud firewall protection for east‑west traffic
• DIA breakout supported, allowing spokes to send trusted SaaS traffic directly to the internet
• Ideal for customers with 500+ sites
• Broad regional coverage through Secure Access integration
6

Design A - Cons
• No support for MX Local Hubs or Hybrid Hubs
• Less flexibility to alter traffic paths
• No customer‑owned hubs to provide backup for cloud transport-traffic relies entirely on the cloud fabric as the backbone
Design B - Any Spokes with Hybrid and Cloud Hubs
This design supports both Secure Internet Access (SIA) and Secure Private Access (SPA) use cases. It allows a mix of Cloud Spokes, Hybrid Spokes, and Local
Spokes, creating a flexible architecture that can adapt to varying connectivity and application access needs.
In the SPA use case, Cloud Spokes use the cloud path for application access, while Hybrid and Local Spokes can use the MX path to reach applications directly
through their enrolled MX hubs. In private access scenarios, an MX‑based traffic path (dotted green) is used byHybrid Spokes to access Hybrid Hub resources.
For internet‑bound traffic, Secure Internet Access (SIA) is used for most sites. Even Hybrid Hubs can enable a default route via SSE to send internet traffic
through the cloud path for SIA.
7

Is this Design for me? - Customer Traffic Profile
This design supports a broader range of traffic patterns by leveraging both the Meraki SD‑WAN fabric and the cloud fabric. Customers can selectively route
east‑west traffic through the Meraki SD‑WAN fabric when needed, bypassing the cloud fabric. This flexibility increases capability but also introduces more
complexity compared to simpler designs.
All traffic directed into the cloud fabric is subject to a unified cloud firewall policy. East‑west flows-such as spoke‑to‑hub, remote‑worker‑to‑SD‑WAN, and
spoke‑to‑spoke-can be governed by this common security layer. When Hybrid Spokes communicate directly with Hybrid Hubs through the MX fabric, these flows
bypass the cloud firewall.
Cloud Spokes can still be used for the majority of branch sites, maintaining simple routing and configuration. Hybrid Spokes provide additional flexibility by
enabling direct communication to selected Hybrid Hubs when performance or latency advantages exist. Remote workers (RAVPN or ZTNA) must be explicitly
permitted in policy to access SD‑WAN resources. Applications behind SD‑WAN must also be defined and allowed as policy destinations.
Resources connected through the cloud transport can communicate freely, except when Hybrid Spokes are configured with priority MX Hubs. Assigning a Hybrid
Hub in the hub‑priority list ensures direct MX‑tunnel connectivity from that Hybrid Spoke.
Design B - Pros
• Supports both SIA and SPA use cases
• Spoke‑to‑spoke traffic can traverse the cloud with a common security policy
• Hybrid Spokes allow direct spoke‑to‑hub access through the MX fabric to improve performance or latency
• Cloud Spokes maintain simple configuration while providing full access to cloud resources
• Provides a balanced approach between security and direct connectivity
8

• Faster convergence improves onboarding speed and redundancy
• Additional redundancy from Local and Enrolled Hubs ensures east‑west traffic continues even if the cloud path
experiences issues
Design B - Cons
• Requires additional MX hardware for Local and Enrolled Hubs
• Dual‑backbone routing increases complexity and troubleshooting difficulty
• Routing defaults to the cloud path, with direct MX‑based paths used only when configured on Hybrid Spokes
• Hybrid Spoke-to-Enrolled Hub traffic bypasses the cloud firewall
Design C: Any Spokes with Local and Cloud Hubs
This design aligns best with the organizations that purchase only theSecure Internet Access (SIA) package. It is centered on maintaining the traditional SIG
(Secure Internet Gateway) model, where the MX SD‑WAN fabric handles east‑west traffic. Itpreserves high‑performance interconnectivity over MX WAN links
while avoiding cloud‑based inspection for private traffic. In this model, Hybrid Spokes list both cloud and MX hubs, with MX hubs prioritized over cloud transport.
9

