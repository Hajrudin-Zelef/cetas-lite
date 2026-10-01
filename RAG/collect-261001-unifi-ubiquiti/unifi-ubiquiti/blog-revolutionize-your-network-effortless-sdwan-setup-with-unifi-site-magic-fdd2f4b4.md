---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-revolutionize-your-network-effortless-sdwan-setup-with-unifi-site-magic-fdd2f4b4
title: "blog-revolutionize-your-network-effortless-sdwan-setup-with-unifi-site-magic-fdd2f4b4"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "latency"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-revolutionize-your-network-effortless-sdwan-setup-with-unifi-site-magic-fdd2f4b4.md
source_anchor: ""
source_lines: [1, 54]
sha256: 809897b814b3a750e3202823e14b2eebf5ae7400acaab06cfb4dc727a89acc2f
---

# blog-revolutionize-your-network-effortless-sdwan-setup-with-unifi-site-magic-fdd2f4b4

Hey tech enthusiasts, it's Juan David here, your Tech Support Lead & UniFi NetworkingExpert at Flytec. Today, we’re diving into an exciting topic that’s transforming how businesses manage multi-location networking: **effortless** **SD-WAN setup with UniFi’s Site Magic feature****.** Buckle up as I explain how this innovative technology simplifies VPN connections and delivers high-performance routing between multiple sites. Whether you're managing branch offices, cloud integrations, or hybrid networks, this guide will help you harness the full potential of UniFi’s SD-WAN technology

**What is UniFi Site Magic?**

Site Magic is UniFi’s cutting-edge **SD-WAN** solution that simplifies creating a seamless VPN mesh network across multiple UniFi Gateways. Managed through the UniFi Site Manager at unifi.ui.com, this feature enables businesses to connect branch offices, share resources, and deploy applications effortlessly. It’s like giving your network superpowers!

The core of Site Magic is its VPN mesh topology, where every UniFi Gateway forms a secure tunnel with every other Gateway in the network. Leveraging **Open Shortest Path First (OSPF)** for high-performance routing, it ensures redundancy, speed, and resilience across sites.

**Architecture and Benefits**

Site Magic uses a **VPN mesh topology**, meaning every Gateway connects to every other Gateway in the group. The routing is managed by OSPF, providing:

- **High Performance:** OSPF dynamically selects the shortest and most efficient path for data.
- **Redundancy:** If one route fails, traffic is rerouted without interruption.
- **Scalability:** Current support is limited to 15 sites, but UniFi plans to expand this in the future.

**Why Site Magic is a Game-Changer**

- **Ease of Use:** No need for complex manual configurations. The entire process is managed through the UniFi Site Manager interface.
- **Reliability:** Site Magic connections persist even during cloud service interruptions.
- **Dynamic IP Handling:** If a Gateway’s WAN IP changes, the system automatically updates the connection, ensuring uninterrupted VPN functionality.
- **Enhanced Performance:** The more Gateways with public IPs, the better the performance and resilience of the network.

**Practical Applications**

1. **Multi-Location Offices**
Share files, applications, and resources between branch offices without worrying about latency or disconnections.
2. **Hybrid Cloud Integrations**
Connect your on-premises network with cloud environments like AWS or Azure for seamless operations.
3. **Disaster Recovery**
Ensure redundancy and failover capabilities by leveraging multiple Gateways with public IPs.
4. **Retail Chains**
Centralize POS systems and inventory management across multiple store locations

**Requirements for Using Site Magic**

Setting up Site Magic requires the following:

1. **UniFi Gateways**  - A **UniFi Cloud Gateway** or a**Next-Generation Gateway** managed with a Cloud Key or official UniFi Hosting. Examples include UDM-Pro (UniFi Dream Machine Pro), Enterprise Fortress Gateway (EFG), or UXG-Pro (UniFi Next-Generation Gateway Pro), depending on your network size and requirements.
  - For connecting third-party gateways or cloud providers like AWS, Azure, or GCP, use OpenVPN or IPsec protocols.
2. A 

1. **Public IP Address**  - At least one Gateway must have a public IP address to establish connectivity.
2. **Same UI Account Owner**  - All participating gateways must be linked to the same UI Account.

**Final Thoughts**

UniFi’s Site Magic takes **SD-WAN** deployment to the next level by providing an intuitive and powerful solution for multi-site connectivity. As a Tech Support Lead and UniFi Networking Expert at Flytec, I can confidently say that this feature can transform your network’s performance, reliability, and scalability.

Ready to set up Site Magic? Need help with configurations? At Flytec, we’re here to guide you every step of the way. Let us help you design, deploy, and optimize your network with ease.

Stay connected,

Juan David, Tech Support Lead & UniFi Networking Expert
