---
id: collect-261001-meraki/meraki/2026-04-cisco-meraki-hybrid-campus-lan-design-html-9c5491cc-1
title: "2026-04-cisco-meraki-hybrid-campus-lan-design-html-9c5491cc"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "parameters", "throughput"]
source: docs/RAG/collect-261001-meraki/2026-04-cisco-meraki-hybrid-campus-lan-design-html-9c5491cc.md
source_anchor: ""
source_lines: [1, 83]
sha256: d02670b855a642fe480525dd8ce45c40f91dc742b3f52b746827c0195ec59621
---

# 2026-04-cisco-meraki-hybrid-campus-lan-design-html-9c5491cc

Cisco Meraki Hybrid Campus LAN Design Guide
CIsco Meraki Enterprise Networking
Architecture, Deployment Options, Cloud Monitoring & Best Practices for a Scalable, Secure Enterprise Campus Network
1. What Is a Hybrid Campus LAN?
A Local Area Network (LAN) is the networking infrastructure that provides access to communication services and resources for end users and devices spread across a single floor, building, or campus. When multiple LANs spread across a local geographic area are interconnected, the result is a Campus Network — which can scale from a single switch in a small remote office all the way to thousands of switching ports in a large multi-building enterprise complex.
A Hybrid Campus LAN refers to a network design that deliberately combines two or more hardware and software platforms — specifically Cisco's Catalyst (DNA/Catalyst Center) portfolio and the Meraki cloud-managed portfolio — into a single, unified campus fabric. This "hybrid" mixing of platforms is extremely common in enterprise environments, whether due to phased upgrades, budget constraints, or the desire to benefit from both portfolios simultaneously.
🎯 What This Design Enables
🔗 Tiered LAN Connectivity
Access, distribution, and core layers designed for performance and scale.
📡 Wired + Wireless Ready
Infrastructure prepared for multimedia, VoIP, and high-density wireless.
📦 IP Multicast
Efficient data distribution for video, streaming, and enterprise applications.
🌐 WAN & Internet Edge
Core interconnection to WAN and internet via SD-WAN capable MX appliances.
2. Why Hybrid? Catalyst + Meraki Together
Cisco offers two distinct but complementary campus LAN product lines. Understanding why organizations choose a hybrid approach requires understanding what each brings to the table:
⚙️ Cisco Catalyst Portfolio
Powered by Cisco Catalyst Center (formerly DNA Center) — an on-premises automation, assurance, and security management platform.
- Network automation & programmability
- AI/ML-driven network assurance
- SD-Access segmentation
- Advanced security with Cisco ISE
- MACSec encryption at the port level
☁️ Cisco Meraki Portfolio
Powered by the Meraki Dashboard — a cloud-first management platform that simplifies deployment, visibility, and day-to-day operations.
- Zero-touch provisioning
- Unified dashboard across switching, wireless, SD-WAN
- Built-in security (UTM) and traffic analytics
- Easy remote management and troubleshooting
- Fast deployment with minimal expertise required
💡 Key Insight: The Hybrid Campus LAN approach is ideal for organizations that want to maximize value from both portfolios — leveraging Meraki's operational simplicity for day-to-day management while benefiting from Catalyst's advanced security, routing, and scalability at the core and distribution layers. Designing for interoperability between the two requires careful planning, which is exactly what this CVD provides.
3. Cloud Management & Monitoring for Cisco Catalyst
One of the cornerstone features of this CVD is the ability to bring Cisco Catalyst switching infrastructure into the Meraki Dashboard for unified visibility. This capability — known as Cloud Monitoring for Catalyst — allows network teams to see both Meraki-managed and Catalyst-managed devices in a single pane of glass.
Supported Catalyst Platforms for Cloud Monitoring
⚠️ Critical Limitation — Monitor Only, Not Manage
Cloud Monitoring for Catalyst provides visibility only. You can see device status, topology, and selected configuration parameters from the Meraki Dashboard — but you cannot push configuration changes to Catalyst devices via the dashboard. Full management of C9500 core switches remains via Cisco Catalyst Center or CLI.
Pre-Requisites for Cloud Monitoring Onboarding
The switch must be a supported model (C9200, C9300, or C9500)
Must have minimum firmware 17.3.1 or higher
Must have an SVI or routed interface with internet access on TCP port 443
Must have a valid DNS server configured
Must have a valid Cisco DNA / Catalyst Center subscription
📌 Note: The onboarding process for C9500 core switches is out of scope for the CVD itself. Refer to the Cisco Cloud Monitoring for Catalyst documentation for a step-by-step onboarding guide.
4. Hybrid Campus LAN Architecture & Components
The reference architecture validated in this CVD combines best-of-breed hardware from both the Meraki and Catalyst portfolios. Each component is purpose-selected for its role in the campus fabric.
MR55 / MR56 / MR57 & C9166-MR Access Points
Management: Meraki Dashboard | Standard: WiFi 6 (802.11ax)
MS390-24P & C9300-24P Access Switches
Management: Meraki Dashboard | Switching Capacity: 208 Gbps
C9500-24Y4C Collapsed Core Switch
Management: Monitor Only in Meraki Dashboard | Switching Capacity: 6.4 TB
MX250 SD-WAN / UTM Appliance (Warm-Spare HA)
Management: Meraki Dashboard | Throughput: 4 Gbps FW / 2 Gbps SD-WAN
📌 Note: Catalyst –M and –MR model SKUs are pre-shipped in Cloud Managed (Meraki) mode. Non-M models can be transitioned to Meraki management mode via CLI. C9300-M uses Cloud Managed mode; C9500 uses Monitor Only mode.
5. Logical Architecture: Three Design Options
This CVD defines three validated logical architecture options, each with different characteristics in terms of VLAN flexibility, convergence behavior, and STP configuration requirements. Choosing the right option depends on your operational priorities and tolerance for complexity.
6. Option 1 — Layer 2 Access with Native VLAN 1 (STP-Based)
In this design, the Spanning Tree Protocol (STP) domain is extended all the way from the access layer to the core layer. VLANs can span across multiple access switch stacks and closets, making this the most flexible option in terms of VLAN design. The recommended STP protocol for this hybrid environment is Multiple Spanning Tree Protocol (MST / 802.1s), as it is supported on both Meraki MS390 and Catalyst platforms.
✅ Pros
- Flexible VLAN design across the campus
- Facilitates seamless wireless roaming
- Easier and consistent configuration across all switches
- Well-suited for large, flat campus deployments
❌ Cons
- Non-deterministic route failover
- Slow STP convergence if not properly tuned
- Different STP protocol support on Catalyst vs Meraki
- Risk of VLAN hopping without proper hardening
⚠️ STP Tuning is Critical: Since Catalyst platforms can run different STP variants than Meraki MS390 switches, it is essential to standardize on MST (802.1s) across your entire campus STP domain. STP must be carefully tuned — including BPDUs, port types, root bridge priority, and convergence timers — to avoid loops and slow failover in a mixed-vendor environment.
7. Option 2 — Layer 2 Access without Native VLAN 1
This option is architecturally similar to Option 1, but with one important security distinction: Native VLAN 1 is replaced with a dedicated, non-trivial management VLAN. This separation reduces the risk of VLAN hopping attacks and aligns with security hardening best practices for enterprise campus LAN design.
✅ Pros
- All VLAN flexibility of Option 1
- Minimizes risk of VLAN hopping
- Management VLAN isolated from Native VLAN
- Preferred option for security-conscious organizations
❌ Cons
- Non-deterministic route failover (same as Option 1)
- Requires careful Native VLAN consistency across all switches
- Different STP protocol support remains a challenge
📌 Important Note: When using Cisco PVST (Per-VLAN STP) in any part of the network, VLAN 1 becomes essential for backward-compatible BPDU communication between switches. If VLAN 1 must be completely eliminated, ensure MST is consistently configured across all Meraki and Catalyst platforms in your STP domain to prevent unexpected BPDU behavior.
8. Option 3 — Layer 3 Access (Recommended for Scale)
