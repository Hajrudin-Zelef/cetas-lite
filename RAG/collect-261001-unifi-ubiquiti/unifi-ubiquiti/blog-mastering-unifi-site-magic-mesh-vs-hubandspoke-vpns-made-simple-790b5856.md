---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-mastering-unifi-site-magic-mesh-vs-hubandspoke-vpns-made-simple-790b5856
title: "blog-mastering-unifi-site-magic-mesh-vs-hubandspoke-vpns-made-simple-790b5856"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-mastering-unifi-site-magic-mesh-vs-hubandspoke-vpns-made-simple-790b5856.md
source_anchor: ""
source_lines: [1, 94]
sha256: 5a14ce23ceaaf7ef6117f377597bf81320796eefe5eb0ee271535878bdb993e2
---

# blog-mastering-unifi-site-magic-mesh-vs-hubandspoke-vpns-made-simple-790b5856

UniFi made site-to-site VPNs so much easier with **Site Magic**, and trust me, if you've ever struggled with complex VPN setups, this is a game-changer. Whether you're an MSP connecting branch offices, or a business owner linking a warehouse to your HQ, UniFi Site Magic simplifies the whole process.

I’m Juan David, Lead Tech Support at Flytec and a UniFi Certified Trainer, and in this guide, I’ll break down everything you need to know to choose between **Mesh** and **Hub-and-Spoke** topologies, explain what gear is required. Plus, I’ll share real-world insights from deployments we’ve supported right here at Flytec.

**What Is UniFi Site Magic?**

**UniFi Site Magic** is Ubiquiti’s approach to building scalable, secure, and easy-to-manage site-to-site VPNs. It allows you to link multiple sites—like offices, stores, and data centers—without needing to configure subnets, NAT, or firewalls manually.

It supports two topologies:

- **Mesh** : All sites talk directly to each other. Great for shared resources like NAS, printers, or VoIP between branches.
- **Hub-and-Spoke** : Centralized model where remote sites connect to a main hub. Perfect for when all traffic flows through a central data center or internet breakout.

**Site Magic Requirements**

Before diving in, make sure you meet the following:

- All gateways must be under the **same UniFi account owner** .
- At least one site must have a **public IP** .
- Devices must run **UniFi Network 9.0.108+** and**UniFi OS 4.1.3+** .
- Compatible devices include **EFG** ,**UDM-SE, UDM-Pro-Max, UDW, UCG-Ultra, UCG-Max** ,**UX** , and more.

**Mesh VPN Setup: Simple and Decentralized**

The Mesh mode is ideal for up to **20 sites** that need to talk directly to each other.

**Use Cases:**

- Retail chains that share printers or NAS.
- Co-managed branches where local devices need peer-to-peer access.

**Steps to Set Up Mesh:**

1. Log into **unifi.ui.com** .
2. Go to **Site Magic > Mesh** and click**Get Started** .
3. Name your mesh network.
4. Select the sites you want to connect (2–20).
5. Choose which local networks (LANs) should be routable.
6. Click **Connect** .




Once connected, you'll see a green link icon across all sites. Test connectivity using tools like ping to verify NAS, printers, or shared devices are accessible.

**Firewall Tip:** Want to limit access between sites? Use the **Zone-Based Firewall** to:

- Block inter-site access by default.
- Allow access to specific IPs or ports only (e.g., just your NAS at 192.168.10.133).

**Hub-and-Spoke VPN Setup: Scalable and Centralized**

Hub-and-Spoke supports **up to 1,000 sites**, ideal for franchises, national networks, and ISPs.

**Use Cases:**

- All internet or cloud traffic routes through a central hub.
- Only certain resources (like an ERP or VoIP server) need to be accessed.

**Steps to Set Up Hub-and-Spoke:**

1. Go to **Site Magic > Hub-and-Spoke** and click**Get Started** .
2. Select a hub site (must have a public IP).
3. Add remote sites as **spokes** .
4. Define which **networks at the hub** are accessible.
5. Choose between:
6. **Failover** : Basic tunnel with backup.
7. **Redundant** : Dual WAN link protection.
8. **Highly Redundant** : Full multi-tunnel HA setup.


**Firewall Tip:** Spoke-to-hub access can be tightly controlled using network rules. Allow traffic to a NAS or server, while blocking everything else.

**Mesh vs. Hub-and-Spoke: Which Should You Use?**

Choosing between **Mesh** and **Hub-and-Spoke** VPN topologies depends on your network needs.

- **Mesh Topology:** Ideal for small to medium networks where all sites need direct communication. Great for environments with shared resources like printers or NAS devices. Offers redundancy, so if one link fails, others stay active.
- **Hub-and-Spoke Topology:** Best for larger setups where a central hub manages all traffic. Suitable for centralizing internet breakout or isolating remote sites (spokes). Often used in franchises or office networks that require controlled access to central resources.

**When to Use:**

- **Mesh:** When every site needs to communicate directly.
- **Hub-and-Spoke:** When a central site manages most traffic and spokes remain isolated.

With **UniFi Site Magic**, switching between topologies is straightforward, offering flexibility as your network grows.

**Final Thoughts**

Whether you're managing 3 sites or 300, **UniFi Site Magic** makes it easy to connect them all securely. Mesh is perfect for small networks with shared resources. Hub-and-Spoke is the way to go when scaling across regions or controlling traffic centrally.

And remember, no more confusing subnetting or complex VPN configs. UniFi Site Magic does the heavy lifting.

Still not sure which setup works best for your business? Get in touch with us at Flytec. We’re here to help you plan and deploy your next multi-site UniFi solution.
