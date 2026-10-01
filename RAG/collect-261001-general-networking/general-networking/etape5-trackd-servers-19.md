---
id: collect-261001-general-networking/general-networking/etape5-trackd-servers-19
title: "Step 5 — Server Vendors + AI Server Market + Data-Center Networking"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "AWS", "Alibaba", "ByteDance", "China", "Nvidia", "Samsung", "United States"]
dates: []
keywords: ["amd", "asic", "aws", "blackwell", "capex", "dram", "gpu", "gpus", "hyperscaler", "inference", "licenses", "liquid cooling"]
source: docs/RAG/collect-261001-general-networking/etape5_trackD_servers.md
source_anchor: ""
source_lines: [1227, 1276]
sha256: b109a8280ce7ab43c7d06250fc8eeae50a55b033c5cae8f609d28ac3f02f91ef
---

# Step 5 — Server Vendors + AI Server Market + Data-Center Networking

### 4.2 TrendForce — GPU vs. ASIC vs. general
- 2026: GPUs **69.7% of AI server shipments**; ASIC-based AI servers ~27–27.8% (revised slightly down mid-year); remainder FPGA/other. [official — TrendForce, Jan/Apr 2026]
- AI inference servers projected at **nearly 50% of total AI server shipments** (TrendForce, older view via InfotechLead). [secondary]
- Inference growth is lifting general-purpose server shipments (DIGITIMES: +10% in 2026) as inference pre/post-processing runs on general servers, not just AI racks. [secondary]

### 4.3 8-GPU platforms — pricing and volume signals
- **B200 8-GPU configurations**: NVIDIA list price ~**$30,000–$40,000 per GPU** in 8-GPU volume; 2026 street quotes under supply constraints: **$45,000–$55,000/unit**. [unverified — tech-insider.org citing Silicon Data] — [tech-insider.org](https://tech-insider.org/nvidia-b200-residual-value-158-percent-2026/)
- DGX B200 appliance: ~**$515,000** list; 8-GPU HGX B200 server: **$400,000–$500,000** (quotes commonly ~$450,000). [unverified]
- **B200 residual value: 158% of launch price** ~1 year after mass availability (Silicon Data, reported Sept 16, 2026) — scarcity premium; A100/H100 also holding above straight-line depreciation. [unverified]
- Cloud B200 rental: neo-cloud on-demand averaged **$5.09/GPU-hour** in Q2 2026; hyperscaler list up to **$14.24/GPU-hour** (AWS Blackwell). [unverified]
- AMD AI GPU units (UBS Global Research, Feb 2026): 2025 revised +28% to **544,500 units**; 2026 revised +59% to **816,800 units** (MI350-led); 2027 ~**1.9M units** (MI450 ramp). [secondary — via Wccftech] — [Wccftech](https://wccftech.com/amd-server-cpu-revenue-to-surge-80-percent-2026-estimated-1-9m-ai-gpus-shipping-by-2027/)
- Rack-scale platforms (GB300 NVL72): **$3–4.5M per rack** (72 Blackwell Ultra GPUs, 36 Grace CPUs, 576 HBM3E stacks, NVLink switching, liquid cooling). [secondary — wnie.online]
- Small-transaction datapoint: Quanome Technologies 8-K (Sept 21, 2026): 32 GPU server units from Compal for **$18.8M** (~$587K/unit), US data-center delivery. [official — SEC filing via Kalkine] — [Kalkine](https://kalkine.com/news/announcements/quanome-technologies-signs-188m-gpu-server-purchase-deal)

---

## 5. Regional breakdowns

### 5.1 IDC server revenue by region
**Q1 2026** [official — IDC](https://www.idc.com/resource-center/press-releases/1q26-server-tracker/):
- United States: **$79.6B (+24.1% YoY), 64.9% of global revenue**. [official]
- PRC (China): **$19.2B (+30.9% YoY)**. [official]
- APeJC: $9.7B (+45.2%); Western Europe: $7.6B (+80.6%); Japan: −16.1% (vs. strong 1Q25 base). [official]
- Canada +190.9%, Middle East & Africa +121.4%, Latin America +64.1%. [official]
- Sovereign AI programs span **>40 countries** — a new policy-driven demand layer. [official — per InfotechLead summary]

**Q2 2026** [independent — StorageReview on IDC](https://www.storagereview.com/news/idc-external-storage-tracker-q2-2026-10-3-billion-as-server-market-tops-166-billion):
- United States: **$112.2B (+54.9% YoY), 67.4% of global revenue**. [independent]
- PRC: **$26.4B (+43.4% YoY)**, "reaccelerating from recent quarters." [independent]
- APeJC: $10.9B (+31.5%); Western Europe: $9.1B (+62.7%); Central & Eastern Europe: $0.7B (+98.3% off small base); Japan +11.1%. [independent]
- Canada: **+202.6% YoY** (fastest-growing worldwide); Middle East & Africa +68.8%; Latin America +32.8%. [independent]

### 5.2 IDC AI infrastructure spending by region
- Q1 2026: United States **$67.9B (75.7% of global AI infrastructure spend, +30.3% YoY)**; China **$7.8B** (returned to growth). [secondary — QuantumRun on IDC AI Infrastructure Tracker]

### 5.3 China-specific signals
- NVIDIA H20 China shipment restrictions trimmed TrendForce's 2025 AI-server growth forecast (to ~24%). [official — TrendForce, Oct 2025]
- Jan 2026: US BIS shifted H200/MI325X-equivalent export licenses from "presumption of denial" to "case-by-case review"; ~10 Chinese firms (incl. Alibaba, Tencent, ByteDance) cleared to buy H200 with a **75,000-unit cap per customer** (per industry analysis). [unverified — Momoview]
- TrendForce (Aug 2026): Chinese CSPs accelerating deployment of domestic AI solutions for LLM services. [official]
- CXMT (ChangXin) has become the largest DRAM supplier outside Samsung/SK Hynix/Micron, though the top three still hold >80% combined share. [secondary — fiisual]

### 5.4 North America dominance in CapEx
- Nine largest CSPs' 2026 CapEx >$886.7B (+~90% YoY); **five North American hyperscalers ≈ 90%**. [official — TrendForce, Aug 2026]
- Top-five hyperscaler 2026 CapEx ≈ $685B (+~70% vs. ~$400B in 2025); 2027 forecast >$850B, some analysts >$1T. [secondary — wnie.online]
- "Big Tech combined AI-related outlays set to surpass $730B this year" (2026) per market commentary around Supermicro's guidance. [secondary — BigGo Finance]

---

## 6. Notable analyst reports and publications (2026)

