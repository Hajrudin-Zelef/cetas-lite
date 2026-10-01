---
id: collect-261001-cisco/cisco/api-deki-pages-12657-pdf-cisco-2bsecure-2baccess-2bintegration-2b-2bdesign-2bbes-bf436f64-1
title: "api-deki-pages-12657-pdf-cisco-2bsecure-2baccess-2bintegration-2b-2bdesign-2bbes-bf436f64"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/api-deki-pages-12657-pdf-cisco-2bsecure-2baccess-2bintegration-2b-2bdesign-2bbes-bf436f64.md
source_anchor: ""
source_lines: [1, 93]
sha256: 94a9a558e755a64db09abe35b8e664de15e6931a67491b8a797fa6dcf9c7afa2
---

# api-deki-pages-12657-pdf-cisco-2bsecure-2baccess-2bintegration-2b-2bdesign-2bbes-bf436f64

Cisco Secure Access Integration - Design Best Practices
Cisco Meraki with Secure Access integration design best practices article outlines options for deploying Meraki MX SD-
WAN hub and spoke networks with Secure Access. Following these recommended designs will help ensure your
network achieves optimal performance and security. Listed designs had been validated with scale, optimal operation
and resiliency in mind.
Overview
Today’s remote and hybrid workforce, widespread cloud adoption, and increased internet-bound traffic require organizations to deliver secure, optimized access
to applications anywhere, on any device. Traditional network models struggle to keep pace, prompting organizations to embrace Secure Access Service Edge
(SASE) architectures. By integrating Cisco Meraki SD-WAN with Cisco Secure Access, businesses benefit from unified, cloud-native networking and security,
ensuring consistent protection, simplified operations, and a scalable user experience across all locations.
Cisco Merakiand Secure Accessstand out with their centralized, cloud-based dashboards, enabling IT teams to deploy, monitor, and manage networks from
anywhere—eliminatingthe need to build moreon-premises controllers. The platform integration prioritizes simplicity, scalability, and automation, while offering
built-in analytics and security for both enterprise and distributed environments. Managed through the Cisco Meraki dashboard, Meraki MX SD-WAN leverages
AutoVPN technology to seamlessly orchestrate and provision Secure Access cloud hubs and tunnels between sites in a spoke-and-hub network.
This design best practices document outlines options for deploying Meraki MX SD-WAN with Secure Access to deliver a comprehensive SASE solution.
Following these recommended designs will help ensure your network achieves optimal performance and security.
Dual Fabric Terminology
For clearer guidance, the following special terms describe devices and traffic paths within two SD-WAN fabrics:Traditional Cisco Meraki MX SD-WAN and
Cisco SSE Secure Access.
Cloud Fabric: Transport via cloud hubs using Cisco Secure Access; cloud path or cloud transport
Cloud Spoke: Cisco Meraki spoke connected (enrolled) exclusively to the cloud
Cloud Hub: Cisco Secure Access-enhanced head-ends that provide cloud transport and are designated with the CPSC-HUB platform type in Cisco
Meraki MX Dashboard
MX (Local) Fabric: Transport via MX hubs, using traditional Cisco Meraki MX SD-WAN tunnels and routing; also referred to as the MX path or MX
transport
Local Spoke: MX in spoke mode operating solely in the MX fabric, listing local or enrolled hubs
Local Hub: MX in hub mode that appears in a spoke's hub priority list
Dual-fabric devices:
Hybrid Spoke: MX spoke that uses both cloud fabric and MX fabric for traffic transport
Hybrid (Enrolled) Hub: MX hub that forms a mesh with other MX hubs while enrolled in the cloud fabric and route peering (hub-to-hub) with the
region’s pair of cloud hubs
Further explanation on thesetermscan be found in the Secure ConnectHub Integrationdocument.
1

Cisco SASE Supported Topologies for Meraki
The following diagrams show the difference between spoke-spoke communication when enabling secure private access for users.
Example Meraki topology:
Example Meraki topology via Secure Access:
2

Why do these changes occur?
When deploying a SASE architecture, it is recommended to inspect East-West traffic between sites to maximize the security efficacy of the security services.
Secure Access integration with Meraki orgdisables the direct Spoke-Spoke communication via MX Hub (2) to ensure all traffic is inspected, and policy is applied
as intended (1).
Platform Optimization
To improve platform stability and resiliency, the following optimization is adopted at onboarding for all organizations that enable the Meraki Secure Access
integration:
• Dashboard installed routes are removed from all Spokes and Hubs.
• BGP protocol is used as a sole method of providing routes to sites
• Meraki Hubs are prevented from sharing the Spoke routes they learned toward other Spokes.
• Cloud Spokes traffic to any other sites (any Spoke or Hub) prefers the cloud path
The table below describes the Cloud vs. Local tunnel path used by end points as they communicate.
* MX Hub (Hybrid or Local) must be configured on a Spoke to enable direct local-path access to its prefixes.
Scale of Sites and Routes in the Meraki Secure Access Solution
Building sizable SD-WAN site topologies for Meraki with Secure Access requires careful consideration for routes scale on all devices. MX Hubs must stay
under10,000-route limit while Secure Access Cloud Hubs need to stay below 70,000-route limit. This allows unoptimized organizationsto support up to1000
enrolled Sites, with each advertising a maximum of10 routing prefixes to the Cloud Fabric. These MX-advertised prefixes encompass local subnets, eBGP-
learned routes, and defined static routes. In this case, Cloud and HybridSpokes will receive a default route and specific routesknown by the Cloud Hubs.
For quicker convergence with Cloud Hubs, platform can be optimized so Spoke sites can be configured to receiveonly a default route, bypassing specific iBGP
route prefixes. This simplified routing also allows sites to advertise an increased limit of up to28 prefixes with the2500 enrolled Sites. To enable this reduced
routing configuration for your regions, please contact Meraki support.
Organizations deploying below Design B at scale (Hybrid Spokes with Hybrid and Cloud Hubs) can accommodate max of2 MX Hybrid Hubs. For organizations
3

using asmall number of Sites, more MX Hybrid Hubs can be enrolled in Secure Access. The SitesUInow allows 100 hubs enrolled, validated by the
Design D (Hybrid Hubs with Cloud Hubs).
Use of templates significantly improves enrollment time ofSites. WhenSites do not use templates, theyshould be enrolled ingroups ofno more
than200Sites. Each batch should completelyenrollbefore proceeding to enroll more sites. Contact Meraki Support for assistance.
Learn more about the 2025 Platform Optimization here, implemented first with Secure Connect.
These settings, derived from insights gained from the Secure Connect product and its underlying technology, ensure faster failover during Data Center outages.
They also simplify spoke site routing, thereby reducing overall route complexity and improvingBGP convergence.
Do I need a Maintenance Window?
• Do you have the Secure AccessSecure Internet Access package only (No SPA package purchased)?
• Customers with the following packagesmay need a maintenance window:
• Secure Internet Access (SIA) Essentials
• Secure Internet Access (SIA) Advantage
• Customers with the following packageswill not need a maintenance window:
• Secure Internet Access (SIA) + Secure Private Access (SPA)
• DNS Defense is not yet supported via this integration.
• Does your organization requirespoke-spoke communication using an MX hub?
• Every Meraki organization has this capability by default, but not all customers use it.
• Some larger organizations use this to manage the scalability of their deployments when changing routing configuration.
To addressbroaderdeployment needs, increasedscale forSpokes and Hubs is beingvalidated and updated in this document.If the design you
intend to use is not documented here, please contact your account team.
If you have answeredYes to both questions above, you should schedule a maintenance window BEFORE enabling the Secure Access integration.
• In case your Meraki organizationrequires MX Hubs mesh feature to beDISABLED?
◦ Every Meraki organization has the Hub mesh enabledby default, but customers can ask Meraki
support to disable this feature.
◦ Secure Access integration onboarding re-enables MX Hub mesh toenablesupportforMX Hub
enrollment into SSE as part of Design B.
◦ If your organization has MX Hub mesh disabled, and you are onboarding to Secure Access integration,
4

