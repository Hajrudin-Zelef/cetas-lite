---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-5546542486551-traffic-policy-management-in-unifi-f01e55ec
title: "hc-en-us-articles-5546542486551-traffic-policy-management-in-unifi-f01e55ec"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["cyber", "latency"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-5546542486551-traffic-policy-management-in-unifi-f01e55ec.md
source_anchor: ""
source_lines: [1, 31]
sha256: 0168ab91ebe307253acfc6d2fe19483f676ec4a646beeb2bf20070c0a5b92242
---

# hc-en-us-articles-5546542486551-traffic-policy-management-in-unifi-f01e55ec

Traffic & Policy Management in UniFi
UniFi provides a unified Policy Engine for managing traffic shaping, routing, and security policies across your network. Whether you're creating firewall rules, routing traffic through a VPN, applying QoS, or blocking malicious content, all major policy types can now be configured from one centralized interface.
For a full overview of UniFi's Network and Cyber Security capabilities, see here.
Centralized Control with the Policy Engine
The Policy Engine consolidates all traffic and policy features into a single interface, making it easier to manage, review, and apply consistent rules across your deployment.
- Zone-Based Firewall – Define security policies to block or allow traffic between your local networks, VPNs, and the internet.
- Application Filtering – Quickly block or allow specific applications or entire categories of applications.
- Policy-Based Routing – Orchestrate traffic through specific WAN interfaces, or force it through a VPN tunnel.
- QoS – Prioritize critical traffic and optimize network efficiency with flexible features like traffic shaping and WiFi speed limits.
- ProAV Traffic Optimization – Streamline professional audio-visual workflows with advanced traffic rules to minimize latency.
- Access Control Lists (ACLs) – Block or allow traffic directly on switches for flexible, low-latency control.
- Content & Domain Filtering – Block explicit, malicious, or unwanted domains across your networks with pre-packaged filters and custom rules.
- NAT Configuration and Port Forwarding – Flexibly configure access to internal services from external networks.
- DNS Records & Local Hostnames – Simplify access to devices and services in your network with custom DNS records and hostname creation.
Object Manager: Outcome-Based Policy Creation (Network 9.4)
Although all policies can be configured from the Master Table, UniFi’s Object Manager offers a more streamlined, outcome-driven workflow for managing firewall, routing, and QoS policies. With Object Manager, you simply define your desired outcome—such as client isolation—and UniFi automatically applies the necessary ACLs and firewall rules. This removes the complexity of manual configuration and ensures consistent, policy-driven results across your network.
To create a policy:
- Go to Policy Engine > Objects.
- Click Create.
- Select Devices, Device Groups, or Networks.
- Choose Secure, Route, and/or QoS outcomes.
  - 
Secure:
    - Internet: Streamlines Firewall and Application Filtering rules by specifying Blocklists, Allowlists, or complete Internet isolation.
    - Local: Streamlines local traffic management of internal networks and/or clients by creating firewall rules (at the gateway) and/or ACL rules (at the switches).
  - Route: Applies policy based routing rules for the specific traffic to use the specified interface.
  - QoS: Prioritize and/or Limit specific traffic across specified interfaces.
- 
Secure:
- (Optional) Set Schedule: Apply a time-based schedule for when the rule(s) should apply.
Rules are generated automatically, and UniFi determines whether to enforce them via firewall or ACL based on traffic direction and deployment type.
