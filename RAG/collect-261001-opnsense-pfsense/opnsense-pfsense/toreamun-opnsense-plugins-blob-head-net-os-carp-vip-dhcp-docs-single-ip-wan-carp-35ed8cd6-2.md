---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/toreamun-opnsense-plugins-blob-head-net-os-carp-vip-dhcp-docs-single-ip-wan-carp-35ed8cd6-2
title: "ON  - backup borrows the master's internet over SYNC (run from node B)"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/toreamun-opnsense-plugins-blob-head-net-os-carp-vip-dhcp-docs-single-ip-wan-carp-35ed8cd6.md
source_anchor: ""
source_lines: [84, 147]
sha256: 08c14b6c8d5e953ed6dc65b8eb01f676af7c80272f97f01536522d0aaa848594
---

# ON  - backup borrows the master's internet over SYNC (run from node B)

The leased public address is the CARP VIP's own address (the "direct" model). On a follow the plugin rewrites the VIP and re-applies it add-before-remove, so the vhid never loses its address on that node. Point Source NAT and any address-dependent rule at the plugin-managed firewall Host alias rather than a hardcoded IP - the plugin updates the alias content live on a follow, so rules track the address without a ruleset reload. (A CARP advertisement's HMAC covers the VIP prefixes, so during a follow the two nodes briefly advertise different prefixes; if they adopted the new address more than ~3 s apart the backup could stop validating the master's adverts and promote - a transient dual-master (lab-confirmed). The keeper closes this: it passively observes the peer's DHCP ACK on the shared chaddr and follows within the same exchange, so both nodes converge well under the ~3 s CARP timeout. Falls back to independent convergence if the peer's ACK isn't visible.)
An alternative binds the public address as an IP-alias VIP on top of a CARP VIP that
carries a stable private election address (same vhid, hence the same virtual MAC, so they
fail over together). It gives textbook same-subnet CARP, but adds a second VIP, needs
a ≥/29 private WAN block, and changes nothing at L2 - inbound is answered with
the virtual MAC and egress uses the physical MAC either way (lab-verified). It is a
matter of taste, not a functional win, and the plugin's follow logic targets the
direct model - so direct is the default; reach for the alias form only if you
specifically want the election address in the node-IP subnet.
flowchart TB
    ISP["ISP<br/>gw 123.123.123.1 - one DHCP IP"]
    subgraph WANL2["WAN segment (L2)"]
        WANSW["WAN-front switch<br/>(any switch, one L2 domain)"]
    end
    ISP --> WANSW
    subgraph NA["Node A - master by default"]
        AW["WAN if<br/>link-local 169.254.255.<b>1</b>/29<br/>+ CARP VIP 123.123.123.123 (vhid 9)<br/>DHCP via virtual MAC (plugin)"]
        AS["SYNC if 10.2.2.<b>1</b>/30"]
    end
    subgraph NB["Node B - backup by default"]
        BW["WAN if<br/>link-local 169.254.255.<b>2</b>/29<br/>+ CARP VIP 123.123.123.123 (vhid 9)"]
        BS["SYNC if 10.2.2.<b>2</b>/30"]
    end
    WANSW --> AW
    WANSW --> BW
    AS <-->|"pfsync + config-sync + transit"| BS
    subgraph INT["Internal (LAN / VLANs)"]
        LANSW["Internal switch<br/>per-VLAN LAN CARP VIPs"]
    end
    AW --- LANSW
    BW --- LANSW
    classDef isp fill:#ffe8cc,stroke:#e69f00,color:#5c3d00
    classDef sw fill:#e9ecef,stroke:#868e96,color:#343a40
    classDef master fill:#cfe6f5,stroke:#0072b2,color:#00344f
    classDef backup fill:#f4e1ec,stroke:#cc79a7,color:#5c2547
    classDef sync fill:#fbf7c4,stroke:#b8a900,color:#524a00
    classDef internal fill:#cceee2,stroke:#009e73,color:#00402e
    class ISP isp
    class WANSW sw
    class AW master
    class BW backup
    class AS,BS sync
    class LANSW internal
    
Node A/B's link-local WAN IPs (169.254.255.1/2) are for CARP only. All real
outbound traffic is NAT'd out of the VIP 123.123.123.123.
sequenceDiagram
    participant ISP as ISP
    participant A as Node A (master)
    participant B as Node B (backup)
    Note over A,B: Steady state
    A->>ISP: master holds the VIP's DHCP lease (virtual MAC)
    A->>B: CARP advertisements (advskew 0) + pfsync
    B-->>A: (optional) borrows internet via SYNC, on demand
    Note over A: Node A dies (HW/crash)
    A--xB: advertisements stop
    Note over B: advertisement timeout, B becomes MASTER
    B->>ISP: takes over the VIP's lease + gratuitous ARP
    Note over B: default already points at WAN_ISP, VIP now active, route works at once (no routing change)
    B->>ISP: traffic straight out the VIP
    Note over A: Node A recovers (preempt=1)
    A->>B: advertisements with advskew 0 (lower)
    Note over A: A takes master back, B returns to backup
    
