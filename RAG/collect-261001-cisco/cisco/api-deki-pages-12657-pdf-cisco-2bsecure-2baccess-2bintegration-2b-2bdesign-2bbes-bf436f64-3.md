---
id: collect-261001-cisco/cisco/api-deki-pages-12657-pdf-cisco-2bsecure-2baccess-2bintegration-2b-2bdesign-2bbes-bf436f64-3
title: "api-deki-pages-12657-pdf-cisco-2bsecure-2baccess-2bintegration-2b-2bdesign-2bbes-bf436f64"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "licenses", "throughput"]
source: docs/RAG/collect-261001-cisco/api-deki-pages-12657-pdf-cisco-2bsecure-2baccess-2bintegration-2b-2bdesign-2bbes-bf436f64.md
source_anchor: ""
source_lines: [182, 251]
sha256: 9159e15712d96d0df0bcca695288bf97a02e4278c0153d81b293dcde2814aea3
---

# api-deki-pages-12657-pdf-cisco-2bsecure-2baccess-2bintegration-2b-2bdesign-2bbes-bf436f64

Is this Design for me? - Customer Traffic Profile
This design provides SIA protection for enrolled Hybrid Spokes with direct MX transport to maintain the performance benefits of the Local Fabric. Furthermore,
this architecture is future-proof; organizations can transition tocloud-based east-west transport after adding anSPA userlicense. Cloud Spokes are excluded
from this design as they lack local fabric connectivity.
Configuration and Connectivity
To achieve optimal connectivity for Spoke-to-Hub application access, configure the relevant Local Hubs on all Hybrid Spokes.
• Local Hub Reachability:ALocal Spoke can communicate with applications hosted on the configured MX Hub;
however, it cannot reach CloudSpokes or endpoints enrolled in the SSE.
Spoke-to-Spoke Communication
When you enableSpoke-to-Spoke communication via a Local MX Hub, it impacts the routing behavior of all Spokes, as traffic paths will prioritize local MX
Hubs.
• Fabric Compatibility: Hybrid and Local Spokes that share the same group of MX Hubs can communicate
effectively. Maintainingaconsistent group of local Hubs across Hybrid Spokes ensuresadirect path of connectivity.
• Cloud Limitations: Cloud Spokes are not part of the Local Fabric and, therefore, cannot communicate with other
Spokes in this configuration.
Remote Access (RAVPN and ZTNA)
Remote workers using RAVPN or ZTNA have specific limitations in this design:
• Local Hubs: Remote workers cannot access subnets located behind Local Hubs.
• Hybrid Spokes: Networks behind Hybrid Spokes are treated as cloud-enrolled resources and are fully reachable by
remote workers.
10

Transport Prioritization
Resources connected via cloud transport maintain full communication protected by the cloud firewall. Hybrid Spokes always prioritize a direct path to Local Hubs
to reach those applications and do not use the Cloud path. Reaching other Spokes with "Spoke-to-Spoke via MX Hub" isensured by sharingthe same group of
configured Local Hubs.
Design C - Pros
• Simple deployment for customers focused exclusively on the SIA use case
• Spoke‑to‑hub and spoke‑to‑spoke traffic stays on the MX transport for full MX platform performance
• Familiar SIG‑style architecture with MX transport preferred over cloud transport
Design C - Cons
• Spoke-to-spoke and spoke-to-hub traffic is not inspected by the cloud firewall
• Hybrid Spokes must maintain consistent Hub priority lists
Support for additionalMeraki designs isbeing validated now thatMeraki with Secure Access solution reachedGeneral Availability.
Design D: Hybrid Hubs with Cloud Hubs (No Spokes)
This design is ideal for organizations that prioritize ahub mesh local path for all site-to-site communication, while gaining the benefit of Secure Private Access
(SPA) for remote workers into Hybrid Hub applications. It offers the flexibility to enable a default route on any hub to support Secure Internet Access (SIA) use
cases. This architecture preserves high-performance interconnectivity over Meraki SD-WAN while exempting private site-to-site traffic from cloud-basedsecurity
services inspection.
11

Is this Design for me? – Customer Traffic Profile
• Remote Access: Provides SPA protection for remote workers accessing enrolled Hybrid Hubs.
• Performance: Hub-to-Hub traffic utilizes direct MX transport (local path) to maintain the full performance benefits of
the Local Fabric.
• Future-Proofing: Organizations can seamlessly add SIA transport in the future by obtaining SIA user licenses and
enabling a default route on the Hybrid Hubs.
Hub-to-Hub Communication
By default,hub mesh is enabled within the Meraki organization. This allows all MX Hubs to establish AutoVPN tunnels and iBGP peering to exchange routes
automatically. This direct, site-to-site communication ensures the lowest latency and highest performance for inter-site traffic.
Remote Access (RAVPN and ZTNA)
Remote workers connecting via RAVPN or ZTNA gain secure access to all Hybrid Hub applications. Their traffic travels through secure tunnels and reaches the
destination applications via AutoVPN on the back end.
Scale of Design D
Cisco has validated this design to a scale of100 Hybrid Hubs (100 MX Hubs enrolled in Secure Access via AutoVPN), with each MX hub managing90
prefixes.
• Route Density: This scenario drives approximately 9,000 routes per hub, leaving adequate headroom for stable
operations.
• Note: Always keep the10,000-route limit of the MX appliance in mind when adding additional routes from other
Secure Access enrolled entities.
Design D – Pros
Pros
• Simplicity: A streamlined deployment for customers focused exclusively on the SPA hub mesh use case.
• Peak Performance: Leverages native MX transport for the best possible site-to-site throughput.
• Familiarity: Utilizes the standard Meraki Hub mesh architecture with the flexibility of an optional SIA cloud path for
Internet-bound traffic.
Design D – Cons
Cons
• Limited Inspection: Hub-to-Hub traffic is not inspected by cloud security services.
12
