---
id: etape6-phaseb-smb-networking/06-supplementary-complementary-research-pass-2-september-22-202/part-8
title: "Supplementary / Complementary Research Pass #2 — September 22, 2026 (part 8)"
domain: supplementary-complementary-research-pass-2-september-22-202
role: deep-dive
task: reference
actors: ["EU", "United States"]
dates: ["2026-06", "2026-07"]
keywords: ["advisory", "agents", "claude", "license"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1156, 1176]
section: "Supplementary / Complementary Research Pass #2 — September 22, 2026"
sha256: 0027bf3f7821a4024346522f8c89fde90eda3a9c837b1d83bae0b9800f29454f
---

# Supplementary / Complementary Research Pass #2 — September 22, 2026 (part 8)

- The "EFG Core" previewed at UWC is officially the **Enterprise Firewall Core (EF-Core)** and is **launched and orderable on the UI Store**: **$3,499.00** ($3,775.00 surcharge incl.); UK store **£3,005.00** (£3,606.00 VAT incl.). **[official]** https://store.ui.com/us/en/category/cloud-gateways-enterprise-scale/products/ef-core?utm_source=web&utm_medium=whatsnew&utm_campaign=1carousel&utm_content=ef-core · https://uk.store.ui.com/uk/en/products/ef-core
- Specs (official + vendor-reported): 24 hyperscale-class Arm Neoverse N2 cores @ 2.5 GHz (ARM v9), 32 GB RAM, 128 GB SSD; ports = 4× 100G QSFP28 + 4× 25G SFP28 + 8× 10GbE RJ45 + 2× 1GbE RJ45 + 1 management + 1 console (default WAN = 1× 100G, 1× 25G, 1× 10GbE); 2× hot-swappable 550W AC/DC PSUs; 1U rackmount; managed through UniFi Network; **79 Gbps IDS/IPS**, 61 Gbps full SSL inspection, 22,500+ clients, 10M concurrent sessions, 120K new sessions/sec, 5,000+ concurrent IPsec/WireGuard tunnels, 38 Gbps aggregate IPsec; shadow-mode/VRRP HA; license-free SD-WAN. **[official]** https://store.ui.com/us/en/category/cloud-gateways-enterprise-scale/products/ef-core?utm_source=web&utm_medium=whatsnew&utm_campaign=1carousel&utm_content=ef-core · **[vendor-reported]** https://news.asbis.com/news/suppliers/ubiquiti-introduces-enterprise-firewall/ · https://aceperipherals.com/products/ubiquiti-ef-core-enterprise-firewall-core-100gbps-cloud-gateway-with-high-availability
- Timeline: reviewed as "available now" ~mid-June 2026 **[independent]** https://dongknows.com/ubiquiti-enterprise-firewall-core-ef-core-review/; ASBIS (distributor) announced availability ~70 days before Sep 22, 2026 (i.e., ~mid-July 2026) **[vendor-reported]** https://news.asbis.com/news/suppliers/ubiquiti-introduces-enterprise-firewall/; shown on the show floor at UWC London 2026 **[independent]** https://www.youtube.com/watch?v=wBk8fXy78Y0
- Paid add-ons at UI Store checkout: CyberSecure Enterprise (Proofpoint threat intelligence) $499.00/unit billed annually; 5-year coverage/replacement $699.00/unit; 90 days of professional phone support included in US/CA/EU/UK. **[official]** https://store.ui.com/us/en/category/cloud-gateways-enterprise-scale/products/ef-core?utm_source=web&utm_medium=whatsnew&utm_campaign=1carousel&utm_content=ef-core
- Availability Sep 2026: **sold out** at Baltic Networks ($4,969.00). **[vendor-reported]** https://www.balticnetworks.com/en-ca/collections/ubiquiti-networks?limit=288

#### A.4. UniFi software, Sep 10 → Sep 22, 2026 (window check)

- **No Network Application release found after 10.6.106 (Sep 10, 2026).** Third-party package trackers (Home Assistant add-ons, FreeBSD net-mgmt/unifi10 port) still list **10.6.106** as current as of ~4–12 days before Sep 22, 2026. **[secondary]** https://github.com/zglate/addon-unifi/blob/HEAD/README.md · https://postgoo.com/posts/6aae49d13acbd737ce601ad4?type=article · https://github.com/hassio-addons/repository-edge/blob/HEAD/unifi/CHANGELOG.md
- **No 10.7 or 11.0** Network release notes or early-access builds found in the Sep 10–22 window. **[unverified — absence of evidence]**
- No UniFi OS release found dated Sep 10–22, 2026. Most recent security baselines cited by third parties in this period: UniFi OS **5.1.12** for most supported devices, **5.1.11** for UDM-Beast, **5.1.10** for UNAS family, UniFi OS Server **5.0.8** (per Vanuatu CERT Advisory 152, citing Ubiquiti Security Bulletin 064). **[independent]** https://cert.gov.vu/images/publications/2026/Advisory_152_Ubiquiti%20UniFi%20OS%20Improper%20Input%20Validation%20Vulnerability.pdf
- 10.6 feature highlights circulating in the period (for context): Topology Spotlight, Time Machine for radios + ports, Channel AI nightly automation, port locking, expanded SafeOps, HA readiness score; Port Locking + Multicast Suppressor noted as "coming soon to Early Access" (requires UAP 8.8+ / switch firmware 7.6+). **[secondary]** https://www.youtube.com/shorts/HZd4vj0d-1Q · https://www.youtube.com/watch?v=mMUARWRLx_U

#### A.5. Homelab consensus picks & trends, 2026

- Switch shortlist (ComputingForGeeks roundup, checked live June 2026, updated ~20h before Sep 22): best overall **MikroTik CRS310-8G+2S+IN** (8× 2.5GbE + 2× 10G SFP+, $210–230); best value/easy **TP-Link Omada ES210X-M2** (8× 2.5G + 2× 10G SFP+, fanless, ~$100); PoE pick **TP-Link SG3210XHP-M2** (8× 2.5G PoE+, 240W budget, ~$385); fanless managed **TP-Link SG3210X-M2** (~$230); "best if you already run UniFi" **UniFi Flex 2.5G PoE** (~$260); pure-fiber aggregation **MikroTik CRS305-1G-4S+IN** (4× 10G SFP+, fanless, $130–150). **[independent]** https://computingforgeeks.com/best-25gbe-10gbe-managed-switch-homelab/
- Trend (2026 GitHub homelab docs): **UCG-Fiber gateway + MikroTik switching** is a recurring combo (e.g., UCG-Fiber + MikroTik CRS310, VLAN-segmented, Talos Kubernetes cluster) **[secondary]** https://github.com/reptambe/digital-garden/blob/HEAD/content/Portfolio/Homelab/4%20-%202026%20Homelab/1%20-%20Architecture.md; Minisforum MS-01 / N5 Pro Proxmox nodes common; Tailscale tailnets for remote access; `home.arpa` local DNS naming; MikroTik CRS812/CRS310 and UniFi UDM-Pro Max appear in larger builds **[secondary]** https://github.com/lushanoperera/homelab/blob/HEAD/hardware-purchases/CLAUDE.md · https://github.com/tyrion70/multica-agents/blob/HEAD/skills/homelab/SKILL.md
- ServeTheHome / LTT-specific 2026 consensus: no dedicated STH or LTT roundup surfaced in this search pass; treat the ComputingForGeeks list + r/homelab-adjacent GitHub builds as the proxy evidence. **[unverified for STH/LTT specifically]**

#### A.6. Supply / availability / surcharges, 2026

