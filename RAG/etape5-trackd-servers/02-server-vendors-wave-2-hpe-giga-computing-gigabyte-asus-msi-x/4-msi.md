---
id: etape5-trackd-servers/02-server-vendors-wave-2-hpe-giga-computing-gigabyte-asus-msi-x/4-msi
title: "4. MSI"
domain: server-vendors-wave-2-hpe-giga-computing-gigabyte-asus-msi-x
role: deep-dive
task: reference
actors: ["AMD", "Baidu", "China", "Huawei", "Intel", "Nvidia"]
dates: ["2019-05-16", "2026-03", "2026-06"]
keywords: ["agentic", "alignment", "amd", "ascend", "blackwell", "compute", "diffusion", "ethernet", "gpu", "gpus", "inference", "intel"]
source: docs/RAG/etape5_trackD_servers.md
source_anchor: ""
source_lines: [615, 662]
section: "Server Vendors — Wave 2: HPE + Giga Computing (Gigabyte) + ASUS + MSI + xFusion"
sha256: e434cf22a1c7e89a5acb7fa42c61f4d05dbd7face673436180b14826607fddaa
---

# 4. MSI

## 4. MSI

### 4.1 AI server offerings & 2026 launches

**CG681-S6093 — 6U liquid-cooled AI server (NVIDIA MGX) — unveiled COMPUTEX 2026 (June 2026)**
- Dual **AMD EPYC** processors; up to **8× NVIDIA RTX PRO 6000 Blackwell Server Edition (liquid-cooled)** GPUs; positioned for large-scale AI inference, agentic AI, physical AI, simulation/graphics/video [official — https://www.msi.com/news/detail/MSI-Accelerates-Enterprise-AI-Strategy-with-Cloud-to-Edge-Ecosystem-at-COMPUTEX-2026-148988].
- **NVIDIA ConnectX-8 SuperNICs**, up to 8× 400 Gbps Ethernet; rack-scale: up to **4× CG681-S6093 systems in a 48RU configuration**; networking anchored by **NVIDIA Spectrum-4 SN5600** Ethernet + SN2201 out-of-band switches [official — same].
- COMPUTEX 2026 theme: cloud-to-edge AI continuum — also **NVIDIA DGX Station** and **NVIDIA DGX Spark** deskside AI supercomputers [official — same].

**AMD EPYC 9006 SP7/SP8 platform lineup (2026)**
- MSI launched server platforms on **6th Gen AMD EPYC (SP7 and SP8 sockets)** for AI/cloud/HPC/virtualization modernization; SP7 for high-density/high-performance, SP8 for balanced mainstream [secondary — https://finance.biggo.com/news/0d245922-df28-40e6-8c8e-0f6badfdbb25]. ⚠️ MSI-specific model numbers and ship dates were not detailed in sources consulted — gap (Giga Computing's schedule is above; MSI's was not).

**AI-vRAN (MWC 2026, March 2026)**
- Unified AI-vRAN platform for O-RAN/private 5G/vRAN with NVIDIA AI Aerial; GPU server solutions integrated into network operations [official — https://www.prnewswire.co.uk/news-releases/msi-unveils-scalable-ai-ran-with-nvidia-ai-aerial-solutions-to-accelerate-5g-and-beyond-at-mwc-2026-302695369.html].

**Earlier scalable AI rack (2025, in-market)**
- NVIDIA Enterprise Reference Architecture rack: four-node MGX scalable unit, each node **8× NVIDIA H200 NVL**, **NVIDIA Spectrum-X** networking; expandable to 32 server systems / **256× H200 NVL GPUs** per deployment [secondary — https://itbrief.com.au/story/msi-launches-scalable-ai-server-solutions-with-nvidia-technology].

### 4.2 Gaps for MSI
- No 2026 HGX B200/B300 8-GPU flagship or GB200/GB300 NVL72 rack-scale MSI product was confirmed in this research — possible gap in coverage (MSI's public emphasis in 2026 appears to be MGX inference + EPYC platforms).
- No customer wins or AI-server revenue figures found — gap.
- No public pricing — gap.

---

## 5. xFusion

### 5.1 Company status (post-Huawei spinoff)
- **Spun off from Huawei in 2021** as Huawei restructured under U.S. sanctions; Huawei remains on the U.S. blacklist [secondary — https://www.tekedia.com/huawei-spin-off-xfusion-lines-up-ipo-adviser-as-chinas-ai-stock-boom-draws-in-server-heavyweights/; https://www.sdxcentral.com/news/huawei-server-spin-out-xfusion-looking-to-go-public-in-china-report/].
- Shareholders reportedly include **China Telecom Group Investment** and **China Mobile Capital Holding** (regional reporting, not from a formal prospectus) [secondary — tekedia; https://www.prismnews.com — prismnews notes the CSRC filing did not disclose a definitive owner register].
- **IPO in preparation**: hired **CITIC Securities** as listing tutor (CSRC filing); **no listing venue, timetable, or offering size disclosed** in the filing [secondary — https://www.folio3.ai/ai-pulse/xfusion-hires-investment-bank-potential-ipo-as-china-accelerates-ai-listings/; https://www.prismnews.com article]. Valuation cited at **~$9B in 2023** (Greatwall Strategy Consultants, via Henan government website) [secondary — tekedia/financeneoteric].
- Light Reading: sells general-purpose, AI, and mission-critical servers; reportedly **China's biggest liquid-cooled server supplier**; took strategic investment from a **SASAC-run fund**; Henan government calls it its **"No.1 digital economy project"** [secondary — https://www.lightreading.com/finance/huawei-server-spinoff-xfusion-heads-for-ipo]. The company's global website describes a Singapore company and does **not** identify Chinese state ownership or Huawei-sourced technology [secondary — Light Reading].
- ⚠️ China AI-stock context 2026: Biren Technology (HK debut +76%), Moore Threads (~+400%), MetaX (~+700%); CSI AI Index +67% in 2025; Baidu Kunlunxin filed for HK listing Jan 2, 2026 [secondary — folio3/tekedia]. This context supports the IPO push but says nothing about xFusion fundamentals.

### 5.2 AI server products / 2026 news
- ⚠️ **Thin verifiable product news**: the only product-adjacent 2026 item surfaced was an artificialintelligence-news.com article (~85 days before Sept 22, 2026) describing a "FusionXpark" compliance appliance and "TokenBox" token-generation appliance with NVIDIA DGX OS — this reads as low-quality/synthetic content with unverified product claims **[unverified — do not use without corroboration]**.
- No verifiable 2026 xFusion AI server model launches (e.g., Kunpeng/Ascend-based FusionServer successors) were found in the sources consulted — **flagged as a research gap**; the company's main website is not accessible outside China [secondary — Light Reading].
- Huawei (separate from xFusion) showcased **CloudMatrix 384** as a GB200 NVL72 rival (more memory/bandwidth, ~559 kW power draw — ~4× NVIDIA's) [secondary — https://www.sdxcentral.com/news/huawei-server-spin-out-xfusion-looking-to-go-public-in-china-report/].

### 5.3 Sanctions / export-control implications
- **Huawei** (and 68+ affiliates) has been on the BIS Entity List since **May 16, 2019** (superseding indictment re: Iran sanctions violations); 46 more affiliates added Aug 2019; licensing under presumption of denial; foreign-direct-product rule expanded [independent — BIS Entity List FAQ, https://media.bis.gov/media/documents/dec-3-2020-update-huawei-entity-listing-faqs-posting.pdf].
- ⚠️ **No source found confirming xFusion itself is on the U.S. Entity List** as of Sept 2026 — flagged as unverified/unknown. xFusion was created precisely as a sanctions-mitigation spinoff vehicle, but its post-2021 technology supply relationships (e.g., access to Intel x86 server chips, which sanctions blocked for Huawei [secondary — folio3]) are not documented in the sources found. xFusion reportedly relies on domestic supply chains and state-backed investment instead.
- 2026 policy context: U.S. AI diffusion framework replacement circulated then withdrawn; lawmakers released the **MATCH Act** (Multilateral Alignment of Technology Controls on Hardware) to tighten controls on advanced compute/chipmaking exports to China [secondary — https://fusionlabs.ghost.io/content/files/2026/04/Export-Controls-State-of-Play-April-2026.pdf]. Any xFusion IPO venue/timetable will be shaped by these regimes [secondary — prismnews].

---

## 6. Cross-vendor notes & uncertainties

