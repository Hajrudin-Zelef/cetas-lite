---
id: collect-261001-cisco/cisco/etape6-tracka-cisco-juniper-8
title: "Step 6 Track A — Cisco + Juniper (Enterprise Data-Center Networking), 2026"
domain: cisco
role: reference
task: reference
actors: ["AMD", "Broadcom", "China", "Huawei", "Nvidia", "Oracle", "TSMC", "United States"]
dates: ["2026-03", "2026-05", "2026-06", "2026-09-22"]
keywords: ["acquisition", "amd", "coherent optics", "compute", "cpo", "dci", "ethernet", "helios", "hyperscaler", "nand", "nvidia", "optics"]
source: docs/RAG/collect-261001-cisco/etape6_trackA_cisco_juniper.md
source_anchor: ""
source_lines: [368, 430]
sha256: afdb501ffd343fc70f009b699a1d283ec11bcbd649aa700f8e0a7928954e6ba3
---

# Step 6 Track A — Cisco + Juniper (Enterprise Data-Center Networking), 2026

- **Market data (IDC, Q2 2026)** via Next Platform (Sept 20, 2026) and kad8:
  - Global Ethernet switch revenue **$18.9B** in Q2 2026 (+64.5% y/y for the data-center segment, $12.3B). [independent — IDC via Next Platform/kad8]
  - **Cisco: $5.43B, +36.2%** (28.7% global share); ~$2.2B from data-center switches; router business $1.6B (+31.6%). Largest vendor overall. [independent]
  - **NVIDIA: ~$2.5B (per kad8, +181% y/y; ~20% of DC Ethernet) — but Next Platform cites $3.86B growing 2.8×.** The two figures appear to use different scopes; **flagged as unresolved.** [independent — discrepancy noted]
  - **Arista: $2.53B, +37.6%** (13.4% global; 18.7% DC). Arista launched a modular router at >115 Tbps in early 2026 aimed at AI. [independent — IDC via Next Platform; Arista launch via Think Insights]
  - **HPE (incl. Juniper): $1.15B Ethernet, +6.1%** — "a lot of incremental business for HPE, whose Ethernet business was largely wireless under Aruba." Huawei $1.55B (+28.6%); ODMs ~$3.56B (+2.43×). [independent — Next Platform's IDC augmentation]
  - **Reading:** HPE-Juniper is #4 by Ethernet revenue (behind Cisco, NVIDIA/Arista, Huawei) but is the only vendor with full campus-to-DC-to-routing plus compute/storage; its growth lags the AI-cluster leaders (NVIDIA, Arista) in DC switching share.
- **Silicon positioning:** Juniper claims Express 5 (28.8 Tb/s) beats Cisco P100 (19.2 Tb/s) by ~33% throughput — a vendor claim Cisco declined to confirm or deny. [vendor-reported]
- **Analyst commentary:**
  - Futurum (Discover 2026 recap): "the networking story is the strongest, most defensible part of HPE's hand"; honest tension is the two coexisting management planes (HPE Mist vs Aruba Central); key tests = self-driving delivery in brownfield, security story above the network, NVIDIA dependency. [independent]
  - ACG's Ray Mota: "HPE delivers the architectural and operational foundation service providers need to fully participate in the AI value chain." [independent quote]
  - EMA's Shamus McGillicuddy: Juniper was showing more value than anyone else; HPE's acquisition centers on that. [independent]
  - Think Insights industry analysis: HPE-Juniper is a new full-stack competitor "explicitly positioned to challenge Cisco and Nvidia in AI-native networking"; Arista retains hyperscale/DC advantage; white-box/SONiC competes on price. [independent]
  - ETR AI survey cited by Futurum: ~21% of orgs building AI apps use HPE for hosting/deployment (and plan to stay), 14% considering — "most of the field is still open." [independent]
  - HPE's own claim (Feb 2026, MWC): "I don't think you could find one company who has a leadership in compute, storage, security and networking across the board" — the full-stack pitch vs Cisco/Arista. [vendor-reported]
- **Customer wins vs Cisco (2026):** Oracle's gigawatt-scale HPE Juniper win is the marquee 2026 reference; Wi-Fi 7 AP sales 7× and campus record orders suggest campus displacement momentum; no 2026-dated named hyperscale DC wins against Cisco surfaced. [official on Oracle; gaps flagged]

## 6. Optics and co-packaged optics (CPO), 2026

- **No Juniper-specific CPO product announcement found in 2026.** This is a notable absence given Broadcom's 2026 CPO push. [explicit gap]
- **Adjacent CPO facts (context):**
  - Broadcom's **Tomahawk 6** (102.4 Tbps) ships with a **native CPO variant** (built on TSMC photonics tech); Juniper builds on Tomahawk 6 for its AI switches and endorsed the chip's "UEC-compliant architecture" and CPO direction. [official — Broadcom PR; vendor-reported Juniper quote]
  - Industry context: TSMC's first true CPO implementation (COUPE photonics) entering production in 2026; NVIDIA's Spectrum-X Ethernet Photonics in production alongside Vera Rubin; 1.6T optical modules ramping in 2026. [secondary]
- **Juniper optics that did move in 2026:** **800G ZR+ coherent optics** on PTX/MX for distributed AI-factory DCI (HPE AI Grid / GTC 2026 announcements; also the AI Factory Grid architecture with NVIDIA). [official/secondary]
- Broadcom's 800G **Thor Ultra NIC** is positioned with the "HPE Juniper Networking AI-optimized fabric" — end-to-end fabric messaging rather than a Juniper transceiver product. [official — Broadcom PR]

## 7. Key uncertainties & gaps

1. **UEC membership:** no 2026 source confirms Juniper in the Ultra Ethernet Consortium; only that Juniper endorses UEC-compliant merchant silicon. [unverified]
2. **CPO product:** no Juniper CPO product announcement; strategy appears to ride Broadcom's CPO roadmap. [gap]
3. **Apstra 2026 version:** no discrete 2026 Apstra release found. [gap]
4. **Junos OS 2026 release:** no new named Junos version for 2026 found; latest AI-ML docs are 24.4-era. [gap]
5. **ACX 2026:** no announcements found. [gap]
6. **DCQCN/RoCEv2 2026 news:** none found (only pre-existing Junos EVO RoCE/AI-load-balancing docs). [gap]
7. **NVIDIA IDC revenue discrepancy** ($2.5B vs $3.86B for Q2 2026) unresolved. [flagged]
8. **Deal value** cited as both $14B and $13.4B (likely gross vs net of cash); used "$14B (some outlets $13.4B)". [flagged]
9. **Instant On divestiture** details (buyer, timing) rest on a single buyer-guide source; DOJ-approved buyer unknown. [unverified]

## 8. Sources consulted (all via web search/fetch, Sept 2026)
- CRN (deal close); The Stack (close/analyst quotes); ainvest (Aug 2026 ruling, Judge Pitts); appliedantitrust.com (HPE court filing, Dec 2025); networkdevicesinc.com (buyer guide); NAND Research (Discover 2026 recap); Futurum Group (Discover 2026 analysis, 24 June 2026); SiliconANGLE (Discover Barcelona, Dec 2025); The Register (723H AP, 8 May 2026); Tech Field Day (Wi-Fi 7 updates); Network World (PTX12000, Express 5); Data Center Knowledge (PTX, MWC 2026); Fierce Network (MWC AI); HostingJournalist (PTX); SDxCentral (Express 5 vs Cisco P100/Nokia FP5); stocktitan (HPE MWC PR); HPE Store UK (QFX5240 SKUs); Juniper.net docs (Junos EVO AI-ML, GLB/DLB); blocksandfiles.com (Q2 FY26); Fool (Q2 FY26 earnings transcript); control.vg (Q2 FY26); Constellation Research (Q2 & Q3 FY26); ConvergeDigest (Q3 FY26, Oracle win); Zacks (Q3 FY26); marketbusinessnews (Q3 FY26); skn-finance (Q3 FY26); dividendinformer (Q3 FY26); Next Platform (IDC Q2 2026 Ethernet data, 20 Sept 2026); kad8 (IDC Q2 2026); Think Insights (industry analysis); ad-hoc-news.de (analyst coverage); Broadcom PR via GlobeNewswire (Tomahawk 6); cxotoday.com (AMD Helios, Dec 2025); businesswire (HPE AI Grid with NVIDIA, 17 Mar 2026); business-news-today/expresscomputer (HPE-NVIDIA sovereign AI); CRN/stocktitan/chartmill/barchart/nomios/saudishopper (SRX400, RSA March 2026); markets.financialcontent.com/urdupure.com/github silicon-photonics research (CPO context).

---

# PART 3 — CROSS-VENDOR SYNTHESIS

*(Assembler-added section; research as of September 22, 2026.)*

---

## 9. Vendor positioning in AI data-center networking (2026)

### 9.1 Market snapshot — Ethernet switch revenue, Q2 2026

| Vendor | Q2 2026 switch revenue | Global Ethernet share | Data-center share | Notable 2026 move |
|---|---|---|---|---|
| Cisco | $5.43B (+36.2% y/y) | 28.7% (#1) | ~$2.2B (~18%) | Silicon One G300 102.4T; FY26 $9.3B hyperscaler AI infra orders |
| NVIDIA | ~$2.5B (+181% y/y; Next Platform cites $3.86B — scope discrepancy flagged) | — | 20.4% (#1 branded DC, IDC) | Spectrum-X CPO + Spectrum-X Ethernet Photonics in production with Vera Rubin |
| Arista | $2.53B (+37.6% y/y) | 13.4% | 18.7% | Tomahawk 6-based 7060XE7 up to 1.6T; full-year 2026 revenue guided >$10B |
| HPE (incl. Juniper) | $1.15B (+6.1%) | — | ~#4 | Oracle gigawatt-scale win (QFX scale-out + PTX scale-across); Networks-for-AI $2.2B cumulative orders |
| Huawei | $1.55B (+28.6%) | 8.2% | — | SP strength, China/emerging markets |

[independent — IDC via Next Platform (Sept 20, 2026) and kad8; NVIDIA figure discrepancy flagged in both drafts]

