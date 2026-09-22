---
id: etape5-trackd-servers/03-data-center-networking-for-ai-2026/8-data-center-interconnect-800g-adoption-1-6t-timeline
title: "8. Data center interconnect: 800G adoption, 1.6T timeline"
domain: data-center-networking-for-ai-2026
role: deep-dive
task: reference
actors: ["Broadcom", "Google", "Meta", "Nvidia", "Oracle"]
dates: ["2026-06", "2026-07-02", "2026-07-14", "2026-07-16", "2026-08-26", "2026-09-06", "2026-09-17"]
keywords: ["compute", "cpo", "dci", "ethernet", "gpu", "gpus", "hyperscaler", "latency", "lpo", "nvidia", "optics", "research"]
source: docs/RAG/etape5_trackD_servers.md
source_anchor: ""
source_lines: [933, 1007]
section: "Data-Center Networking for AI (2026)"
sha256: 68d4684a84653c46ea815e18282df29c9e622c71c93942989765782115a1e8d2
---

# 8. Data center interconnect: 800G adoption, 1.6T timeline

## 8. Data center interconnect: 800G adoption, 1.6T timeline

- **800G is the mainstream AI-fabric generation:** "the vast majority of Ethernet switch shipments and revenues in AI backend networks" in Q1 2026 were 800G (Dell'Oro). [independent]
- **Optical volume:** 800G+ shipments 24M units (2025) → ~63M (2026) (TrendForce). [secondary]
- **1.6T timeline:** 1600G switches "only beginning to sample," expected to "ramp in the second half of 2026" (Dell'Oro, June 2026). 1.6T optical shipments 2.5M (2025) → 20M+ by end-2026 (OFC-2026 analyst confirmation); "1.6T is shipping now" with NVIDIA and Google integrating, Meta and Oracle slated next (eetimes, Sept 2026). [independent] [secondary]
- IEEE 802.3dj (200G/lane, covering 200G/400G/800G/1.6T) on track for completion late 2026; early 200G/lane products expected during 2026; 400G/lane project next. [independent] (networkworld)
- **DCI / scale-across:** emerging as a distinct revenue layer — NVIDIA Spectrum-XGS (available now, SW/FW upgrade), Cisco P200 silicon for scale-across, Arista 1.6T scale-across platforms. [official] [vendor-reported]
- **3.2T:** development underway around 400G-per-lane designs (Kozlov, eetimes); Spectrum-X3200/Quantum-X3200 on NVIDIA's roadmap (TrendForce). [secondary]

---

## 9. Key uncertainties and explicit flags

1. **Arista FY2026 guide conflict:** $11.5B (Aug 2026 coverage) vs $12.6B (remio.ai, Sept 19, 2026). Unresolved. [unverified]
2. **UEC 1.0.3 currency:** the "1.0.3 (July 16, 2026) current" claim comes from a community wiki, not ultraethernet.org. Verify before citing. [unverified]
3. **UEC production adoption:** no confirmed hyperscaler production deployment of UEC/UET found; 2026 evidence is interop demos, NICs, switches, test gear. [unverified]
4. **NVIDIA silicon-photonics/CPO shipping:** roadmap language (Spectrum-X Ethernet Photonics H2 2026; Quantum-X early 2026) vs SemiAnalysis yield concerns; Shainer's "shipping and ramping H2 2026" is a vendor claim. Large-scale CPO volume realistically 2028–2030 per Yole. [unverified]
5. **Quantum-3:** no evidence located; roadmap shows Quantum-2 → Quantum-X800 → Quantum-X1600. Do not assert a Quantum-3 product. [unverified]
6. **Spectrum-XGS announcement date:** "available now" release; Hot Chips mention suggests Sept 2025 but date not verified from NVIDIA's page. [unverified]
7. **Market-size figures conflict:** transceiver TAM estimates differ widely by firm (LightCounting ~$23.8B 2025 vs MarketsandMarkets ~$9.2B DC 2025). Treat as directional. [secondary]
8. **DriveNets claims** (6% better than IB, 15% better than Spectrum-X on NCCL): single vendor blog using SemiAnalysis test data, unaudited. [vendor-reported]
9. **Quantum-X800 latency/bandwidth table figures** (ascentoptics: <100 ns, 115.2 Tb/s vs 51.2 Tb/s): secondary product-page comparison, internally inconsistent with other latency quotes; quote with caution. [secondary]

---

## 10. Sources (URLs as found)

- https://investor.nvidia.com/news/press-release-details/2025/NVIDIA-Introduces-Spectrum-XGS-Ethernet-to-Connect-Distributed-Data-Centers-Into-Giga-Scale-AI-Super-Factories/default.aspx [official]
- http://nvidianews.nvidia.com/news/spectrum-x-ethernet-networking-xai-colossus [official]
- https://cxotoday.com/hardware/nvidia-debuts-spectrum-6-switches-to-power-ai-networks-of-the-future/ [vendor-reported]
- https://www.techpowerup.com/340218/nvidia-links-data-centers-into-a-unified-supercomputer-with-spectrum-xgs-ethernet [secondary]
- https://www.techpowerup.com/337615/broadcom-ships-tomahawk-6-switch-chip-series-with-102-4-tbps [secondary]
- https://pr.wvcjournal.com/article/Broadcom-Now-Shipping-Worlds-First-1024-Tbps-Switch-in-Production-Volume/69b31c1b3971457808b0979e [official] (GlobeNewswire)
- https://www.stocktitan.net/news/AVGO/broadcom-now-shipping-world-s-first-102-4-tbps-switch-in-production-abi415f14rou.html [secondary]
- https://temperature2.com/p/2026-09-06-guide-infiniband-vs-ethernet/ [secondary] (Sept 2026, independent-ish analysis)
- http://gpusmith.com/articles/en/pdfs/infiniband-vs-spectrum-x-vs-roce-vs-ethernet-ai-clusters.pdf [secondary] (Dell'Oro Q1 2026 data)
- https://rdp.in/gpu-mart/knowledge-base/infiniband-vs-spectrum-x-vs-ethernet-for-ai-clusters/ [secondary]
- https://ascentoptics.com/blog/nvidia-quantum-x800/ [secondary]
- https://www.trendforce.com/insights/infiniband-vs-ethernet [secondary] (TrendForce roadmap table)
- https://www.communicationstoday.co.in/over-800g-optical-transceiver-shipments-to-soar-2-6x-by-2026/ [secondary] (TrendForce figures)
- https://www.eetimes.com/ai-demand-reshapes-optical-connectivity-and-photonics-roadmaps/ [independent] (Kozlov commentary, Sept 2026)
- https://markets.financialcontent.com/kelownadailycourier/article/bizwire-2026-3-16-keysight-advances-ai-networking-with-ultra-ethernet-llr-and-cbfc-interoperability-demonstration-at-ofc-2026 [official] (Keysight)
- https://github.com/abuabdurahman82/llm-systems-wiki/blob/HEAD/AI-Factory-Networking/30-ultra-ethernet-consortium.md [secondary] (community wiki, 2026-08-26)
- https://www.networkworld.com/article/4113364/ethernet-groups-keep-2026-focus-on-higher-bandwidth-ai-demands.html [independent]
- https://itbrief.co.nz/story/semtech-sets-200g-lpo-targets-for-data-centre-links [independent] (Semtech LPO targets)
- https://dataintelo.com/report/linear-pluggable-optics-lpo-market [secondary] (market research estimate)
- https://www.eedesignit.com/why-hpc-chip-designers-are-turning-to-linear-pluggable-optics/ [secondary]
- https://www.gminsights.com/industry-analysis/optical-transceiver-market [secondary] (market research estimate)
- https://markets.financialcontent.com/pentictonherald/article/marketminute-2026-3-25-the-16t-supercycle-optical-and-networking-stocks-surge-as-hyperscalers-ignite-gigascale-ai-infrastructure [secondary]
- https://s21.q4cdn.com/861911615/files/doc_news/Arista-Networks-Inc--Reports-Second-Quarter-2026-Financial-Results-2026.pdf [official] (Arista Q2 2026)
- https://www.gurufocus.com/news/9003286/arista-networks-inc-reports-second-quarter-2026-financial-results [secondary]
- https://cryptobriefing.com/arista-networks-3b-quarter-ai-demand/ [secondary]
- https://www.tikr.com/blog/arista-networks-just-posted-its-first-3-billion-quarter-the-ai-networking-buildout-is-accelerating [secondary]
- https://www.zacks.com/commentary/2990188/bull-of-the-day-arista-networks-inc-anet [secondary] (Sept 16, 2026)
- https://www.remio.ai/post/arista-networks-outlook-jumps-as-ai-demand-meets-an-nvidia-challenge [secondary] (Sept 19, 2026 — $12.6B claim, conflicting)
- https://infotechlead.com/networking/cisco-revenue-jumps-18-to-17-3-bn-as-ai-orders-hit-9-3-bn-and-networking-demand-surges-97684 [secondary] (Cisco FY2026)
- https://www.ainvest.com/news/cisco-delivered-ai-infrastructure-results-wanted-margins-telling-2608/ [secondary]
- https://www.ainvest.com/news/cisco-ai-supercycle-architecturally-real-drop-earnings-wrong-read-2608/ [secondary]
- https://financefeeds.com/cisco-booked-4-billion-of-ai-orders-in-one-quarter-beat-revenue-and-eps-and-the-stock-still-fell/ [secondary]
- https://www.trefis.com/stock/csco/articles/615687/is-cisco-stock-priced-for-orders-that-are-not-yet-revenue/2026-09-17 [secondary] (Sept 17, 2026)
- https://www.techradar.com/pro/xais-colossus-supercomputer-cluster-uses-100-000-nvidia-hopper-gpus-and-it-was-all-made-possible-using-nvidias-spectrum-x-ethernet-networking-platform [secondary]
- https://www.sdxcentral.com/news/xai-to-double-colossus-compute-capacity-reveals-cluster-uses-nvidia-spectrum-x-ethernet/ [independent]
- https://measuredai.substack.com/p/xai-colossus-data-center-cluster [secondary]
- https://www.ainvest.com/news/arista-networks-anet-ai-driven-data-center-infrastructure-playbook-2026-2512/ [secondary]
- https://www.fxempire.com/forecasts/article/arista-networks-huge-ai-sales-upcoming-earnings-lift-shares-1614084 [secondary]
- https://github.com/47tzp4ydc9-cmyk/life-os/blob/HEAD/investment-os/narrative/research/silicon-photonics.md [secondary] (investor research note, 2026-07-02)
- https://github.com/roachx92/equity-watch/blob/HEAD/tickers/AAOI/reports/2026-07-14.md [secondary]
- https://medium.com/teradata-labs/the-ultra-ethernet-endgame-cec7242a243b [secondary] (advocacy piece)
- https://drivenets.com/blog/why-infiniband-falls-short-of-ethernet-for-ai-networking/ [vendor-reported]


---

## §4 AI server market figures

