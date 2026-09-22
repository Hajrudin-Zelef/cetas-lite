---
id: etape5-trackb-amd/00-amd/overview
title: "Step 5 — Track B: AMD Hardware (Instinct GPUs + EPYC/Ryzen CPUs)"
domain: step-5-track-b-amd-hardware-instinct-gpus-epyc-ryzen-cpus
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "Meta", "Microsoft", "Nvidia", "Oracle"]
dates: ["2023-12", "2026-02-01", "2026-04-28", "2026-09", "2026-09-22"]
keywords: ["amd", "gpu", "gpus", "aws", "benchmark", "benchmarks", "compute", "cost", "datacenter", "hbm3", "helios", "hyperscaler"]
source: docs/RAG/etape5_trackB_amd.md
source_anchor: ""
source_lines: [1, 51]
section: "Step 5 — Track B: AMD Hardware (Instinct GPUs + EPYC/Ryzen CPUs)"
sha256: 64c626626016159fe2855c57c40aaa6debffa5495cc906f2aa3eeb159daca37e
---

# Step 5 — Track B: AMD Hardware (Instinct GPUs + EPYC/Ryzen CPUs)

## Research report (English) — collected September 22, 2026

**Project:** RAG data-collection, Step 5 (Servers & hardware)
**Coverage window:** February 1, 2026 → September 22, 2026 (with launch-context for context)
**Scope:** Instinct MI300X/MI325X, MI350X/MI355X, MI400 series (MI455X/MI430X), Helios rack platform, EPYC Turin 9005 + Venice 9006, Ryzen client AI (Gorgon Point / Strix Halo), AMD AI partnerships 2026
**Status:** Research snapshot. Cloud $/hr prices and street prices are dated snapshots (mostly Sep 2026) and move constantly; re-verify before use.

### Provenance legend
- **[official]** — AMD's own announcements, press releases, keynotes, published specs, or SEC filings.
- **[vendor-reported]** — figures claimed by AMD (benchmarks, uplift claims) without independent audit.
- **[independent]** — reputable third-party press (Reuters, The Register, ServeTheHome, Tom's Hardware, SemiEngineering, EE Times, DIGITIMES) or independent measurement.
- **[secondary]** — lower-tier press, blogs, aggregators, community trackers (incl. GitHub research repos); useful but unverified.
- **[unverified]** — single-source, conflicting, or rumor claims; treat as uncertain.

### ROCm note
ROCm software stack details (ROCm 10.0.0, TheRock Core SDK, ROCm.AI, PyTorch support matrix) were covered in **Step 4 Track C** (`etape4_trackC_cuda_rocm_pytorch.md`) and are **not repeated here** — only hardware-side implications are referenced below.

---

## 1. INSTINCT MI300X / MI325X — 2026 status: mature workhorse, pricing trends, deployments

### 1.1 Product context (orientation, outside window)
- **MI300X** launched December 2023 (CDNA 3, 192 GB HBM3, 5.3 TB/s). **MI325X** launched late 2024 (CDNA 3, 256 GB HBM3E, 6 TB/s). Both are OAM 750 W parts **[secondary — https://www.bentoml.com/blog/amd-data-center-gpus-mi250x-mi300x-mi350x-and-beyond; https://github.com/regevba/fittracker2/blob/HEAD/docs/research/2026-04-28-hadf-signature-expansion.md]**.
- MI300X was the first AMD datacenter GPU adopted at meaningful hyperscaler scale (Microsoft, Meta, Oracle) as an H100 alternative **[secondary — https://www.nasdaq.com/articles/oracle-recently-delivered-incredible-news-advanced-micro-devices-amd-stock-investors]**.

### 1.2 2026 pricing and availability trends
- AMD does **not publicly disclose Instinct GPU list prices**; enterprise pricing is structured through OEM/hyperscaler deals (Dell, HPE, Supermicro, Lenovo channel) **[secondary — https://github.com/redhat-et/physical-ai-platform-intel/blob/HEAD/deliverables/intel/companies/amd-deep-dive.md]**. ⚠️ Treat all public "street prices" as community-sourced estimates.
- **Cloud $/GPU-hr (September 2026 snapshots):**
  | Provider | GPU | Price | Date | Provenance |
  |---|---|---|---|---|
  | Oracle Cloud (OCI) | MI300X | ~$6.00/hr | Sep 2026 | [secondary — https://susiloharjo.web.id/nvidia-h200-vs-amd-mi300x-vs-google-tpu-v6e-the-2026-ai-chip-benchmark-battle/] |
  | Azure | MI300X | $6.00–$7.86/hr | Sep 2026 | [secondary — same] |
  | Vultr | MI300X on-demand | $3.99/hr | 2025 collection | [secondary — https://www.bentoml.com/blog/amd-data-center-gpus-mi250x-mi300x-mi350x-and-beyond] |
  | DigitalOcean | MI300X | $1.99/hr on-demand | 2025 collection | [secondary — same] |
  | Hot Aisle | MI300X VM | $2.99/hr new allocations (was $1.99) | **Jul 14, 2026** | [secondary — https://github.com/hotaisle/hotaisle-website/blob/HEAD/src/content/blog/why-we-raised-our-mi300x-price.md] |
  | Spheron | MI300X on-demand | $3.59/hr (H100 SXM5 at $2.65/hr on the same marketplace) | **Sep 22, 2026** | [independent-ish marketplace — https://www.spheron.network/blog/triton-on-amd-rocm-vs-nvidia-cuda-same-kernel-different-perf/] |
- **Trend signal:** Hot Aisle raised newly allocated MI300X VM pricing from **$1.99 → $2.99/GPU-hr** (existing allocations held at $1.99; bare metal $3.39) on **Jul 14, 2026**, explicitly citing **full capacity** and MI300X's value position vs NVIDIA alternatives **[secondary — Hot Aisle blog]**. On Spheron's marketplace (Sep 22, 2026), MI300X at $3.59/hr was **more expensive than H100 at $2.65/hr** — the old "AMD is the cheap alternative" assumption is breaking down on some marketplaces **[independent — Spheron blog]**.
- ⚠️ Oracle Cloud is the mainstay for first-party MI300X cloud instances; AWS p5-class MI300X instances were still not generally available as of early 2026 **[secondary — https://susiloharjo.web.id/nvidia-h200-vs-amd-mi300x-vs-google-tpu-v6e-the-2026-ai-chip-benchmark-battle/]**.
- One real-world engineering-cost note (early 2026): teams report **40–80 engineering hours** on average to migrate production workloads from CUDA to ROCm — a TCO offset to the hourly savings **[secondary — same comparison]**. ⚠️ Anecdotal, treat as directional.

### 1.3 Key 2026 deployments
- **Oracle OCI:** announced **Oct 14, 2025** that OCI would deploy **50,000 AMD AI GPUs starting calendar Q3 2026** (expansion planned into 2027), beginning with MI300X-powered shapes and extending to MI355X GA on OCI Compute / "zettascale OCI Supercluster" **[independent/secondary — https://www.mitrade.com/insights/news/live-news/article-3-1194409-20251015; https://www.btcc.com/en-CA/square/Global%20Cryptocurrency/1069079; https://en.tmtpost.com/post/7721864]**. This is one of AMD's largest disclosed cloud wins to date.
- **Oracle 30,000 MI355X order:** on Oracle's **Mar 10, 2025** investor call, chairman Larry Ellison said Oracle placed a multibillion-dollar order for **30,000 AMD MI355X GPUs** — before the chips shipped in meaningful volume (mid-year 2025 target) **[secondary — https://www.nasdaq.com/articles/oracle-recently-delivered-incredible-news-advanced-micro-devices-amd-stock-investors]**.
- **IBM + Zyphra:** announced **Oct 28, 2025** (context): multiyear agreement for IBM to host a large **MI300X cluster on IBM Cloud** for Zyphra (open-source AI research company, SF), with Pensando Pollara 400 AI NICs — presented as one of the largest AMD-hardware genAI training clusters on a commercial cloud, used to train multimodal foundation models for Zyphra's "Maia" superagent **[secondary — https://techsabado.com/2025/10/28/tech-news-amd-ibm-zyphra-forge-new-path-for-ai-infrastructure/]**.
- **Meta:** continues as an MI300/MI350-series customer; the 2026 relationship is dominated by the new 6 GW deal (see §7).
- **MI325X cloud pricing (2025 collection, background):** Vultr $4.615/hr on-demand; DigitalOcean $1.69/hr (12-month) **[secondary — https://www.bentoml.com/blog/amd-data-center-gpus-mi250x-mi300x-mi350x-and-beyond]**.

---

