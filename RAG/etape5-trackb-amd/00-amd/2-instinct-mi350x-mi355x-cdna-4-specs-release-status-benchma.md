---
id: etape5-trackb-amd/00-amd/2-instinct-mi350x-mi355x-cdna-4-specs-release-status-benchma
title: "2. INSTINCT MI350X / MI355X (CDNA 4) — specs, release status, benchmarks vs B200, pricing, customers"
domain: step-5-track-b-amd-hardware-instinct-gpus-epyc-ryzen-cpus
role: deep-dive
task: hardware
actors: ["AMD", "DeepSeek", "Meta", "Nvidia", "OpenAI", "Oracle", "TSMC"]
dates: ["2025-06-12", "2026-04-28"]
keywords: ["benchmark", "benchmarks", "pricing", "accelerator", "agent", "agi", "amd", "blackwell", "compute", "deepseek", "fp4", "fp8"]
source: docs/RAG/etape5_trackB_amd.md
source_anchor: ""
source_lines: [52, 98]
section: "Step 5 — Track B: AMD Hardware (Instinct GPUs + EPYC/Ryzen CPUs)"
sha256: 719900cbc9b2c1c43306d9360d9bc5b7101a11b7d4d78e43ace0337f65b6314e
---

# 2. INSTINCT MI350X / MI355X (CDNA 4) — specs, release status, benchmarks vs B200, pricing, customers

## 2. INSTINCT MI350X / MI355X (CDNA 4) — specs, release status, benchmarks vs B200, pricing, customers

### 2.1 Launch and architecture
- **Announced June 12, 2025** at AMD's "Advancing AI" event in San Jose — MI350X and MI355X launched alongside **[independent — https://www.theregister.com/special-features/2025/06/12/amd-shines-a-light-on-its-helios-rack-scale-compute-platform/1352958; https://www.techpowerup.com/337940/amd-instinct-mi355x-draws-up-to-1-400-watts-in-oam-form-factor]**.
- **CDNA 4 architecture**, **185 billion transistors**, TSMC **N3P** process for the 8 XCD compute chiplets; dual I/O dies on TSMC 6nm; 128 HBM3E channels, 256 MB Infinity Cache, 4th-gen Infinity Fabric links **[official via press — https://www.tweaktown.com/news/105766/amd-launches-instinct-mi350-series-ai-chips-185-billion-transistors-288gb-hbm3e-memory/index.html; https://github.com/amd-agi/primus-turbo/blob/HEAD/agent/skills/kernel-optimize/knowledge/hardware/gfx950/overview.md]**.
- **256 active compute units** (8 XCD × 32 CU), **16,384 stream processors** — fewer CUs than MI325X/MI300X (304 CU), offset by ~2× per-CU low-precision throughput **[vendor-reported/secondary — https://www.tweaktown.com/news/105766/amd-launches-instinct-mi350-series-ai-chips-185-billion-transistors-288gb-hbm3e-memory/index.html; https://github.com/amd-agi/primus-turbo/blob/HEAD/agent/skills/kernel-optimize/knowledge/hardware/gfx950/overview.md]**.
- Memory: **288 GB HBM3E per GPU**, **8 TB/s** sustained bandwidth (12-Hi stacks, 36 GB each) **[official via press — same sources]**. AMD positions 288 GB vs **~60% more than NVIDIA B200's ~180–192 GB** and above MI325X's 256 GB, fitting up to ~520B-parameter models on a single chip **[vendor-reported — https://www.crn.com/news/components-peripherals/2025/amd-instinct-mi350-gpus-use-memory-edge-to-best-nvidia-s-fastest-ai-chips]**.
- Precision support: **FP16, FP8 (OCP), FP6, FP4 + MXFP8/MXFP6/MXFP4 block-scaled formats** — FP6/FP4 are new in CDNA 4 **[official via press — https://github.com/amd-agi/primus-turbo/blob/HEAD/agent/skills/kernel-optimize/knowledge/hardware/gfx950/overview.md]**.
- Form factor / power: **MI350X** = OAM, **1,000 W**, air-coolable; **MI355X** = OAM, **1,400 W**, direct-liquid-cooling target — the same 1,400 W envelope as NVIDIA GB300 "Grace Blackwell Ultra" **[official via press — https://www.techpowerup.com/337940/amd-instinct-mi355x-draws-up-to-1-400-watts-in-oam-form-factor]**.
- Peak theoretical throughput (per GPU, dense; ×2 with 2:4 sparsity):

| dtype | MI350X | MI355X |
|---|---|---|
| FP4 / FP6 / MX | 18.4–18.45 PFLOPS | 20.1 PFLOPS |
| FP8 | 9.2 PFLOPS | 10.1 PFLOPS |
| FP16/BF16 | 4.6 PFLOPS | 5.0 PFLOPS |
| FP64 | 72 TFLOPS | 78.6–79 TFLOPS |

**[official via press — https://www.tweaktown.com/news/105766/amd-launches-instinct-mi350-series-ai-chips-185-billion-transistors-288gb-hbm3e-memory/index.html; https://www.tomshardware.com/pc-components/gpus/amd-announces-mi350x-and-mi355x-ai-gpus-claims-up-to-4x-generational-gain-up-to-35x-faster-inference-performance; https://github.com/amd-agi/primus-turbo/blob/HEAD/agent/skills/kernel-optimize/knowledge/hardware/gfx950/overview.md]**

- Scale-up: 7 external Infinity Fabric links at 38.4 Gbps 16-lane (153.6 GB/s per link), **~1,075 GB/s** peer aggregate per device; host attach PCIe Gen 5 x16 **[secondary — https://github.com/openfabric-systems/simllm/blob/HEAD/docs/papers/amd-gpu-fabric.md]**.
- AMD's "UBB8" open AI-infrastructure standard was introduced with MI350 for faster air/liquid node deployment **[official via press — https://www.tweaktown.com/news/105766/amd-launches-instinct-mi350-series-ai-chips-185-billion-transistors-288gb-hbm3e-memory/index.html]**.

### 2.2 Release/GA status in 2026
- MI350X/MI355X began shipping in meaningful volume **mid-2025** (H2 2025 GA); by 2026 the series is in full production and is AMD's flagship shipping AI accelerator **[independent/secondary — https://www.tomshardware.com/pc-components/gpus/amd-announces-mi350x-and-mi355x-ai-gpus-claims-up-to-4x-generational-gain-up-to-35x-faster-inference-performance; https://github.com/regevba/fittracker2/blob/HEAD/docs/research/2026-04-28-hadf-signature-expansion.md]**.
- **Enterprise PCIe variant "MI350P"** is priced through channel partners (Dell, HPE, Supermicro, Lenovo) — no public list price **[secondary — https://github.com/redhat-et/physical-ai-platform-intel/blob/HEAD/deliverables/intel/companies/amd-deep-dive.md]**.
- **Cloud pricing (Sep 2026 snapshot):** Oracle Cloud MI355X at **$8.60/GPU-hr** **[secondary — https://susiloharjo.web.id/nvidia-h200-vs-amd-mi300x-vs-google-tpu-v6e-the-2026-ai-chip-benchmark-battle/]**.

### 2.3 AMD's competitive claims vs NVIDIA B200/GB200 (all vendor-reported)
At launch AMD claimed for MI355X **[vendor-reported — https://www.crn.com/news/components-peripherals/2025/amd-instinct-mi350-gpus-use-memory-edge-to-best-nvidia-s-fastest-ai-chips]**:
- **FP6:** 2× GB200 performance (FP6 is the format AMD positions for next-gen training).
- **FP4:** on par with GB200, ~10% faster than B200.
- **FP8:** on par with GB200, ~10% faster than B200.
- **FP16:** on par with GB200, ~10% faster than B200.
- **FP64:** 2× GB200 and B200.
- Inference throughput: **~20% better on DeepSeek R1** and **~30% better on Llama 3.1 405B** than B200; AMD claims MI355X "delivers the highest inference throughput" for large models **[vendor-reported — same]**.
- Generational: up to **4× vs MI300X-class (CDNA 3)** performance; up to **35× faster inference** than prior-gen in AMD's framing **[vendor-reported — https://www.tomshardware.com/pc-components/gpus/amd-announces-mi350x-and-mi355x-ai-gpus-claims-up-to-4x-generational-gain-up-to-35x-faster-inference-performance]**.
- ⚠️ Independent third-party head-to-heads at scale remain thin; TechPowerUp's launch coverage cautioned that "the question remains whether real-world benchmarks will match these ambitious specifications" **[independent — https://www.techpowerup.com/337940/amd-instinct-mi355x-draws-up-to-1-400-watts-in-oam-form-factor]**. Treat the B200 comparisons as vendor claims.

### 2.4 Customers and traction 2026
- **Oracle:** 30,000-unit MI355X order (multibillion-dollar, Mar 10, 2025 call); MI355X GA coming to OCI Compute starting calendar Q3 2026 as part of the 50,000-GPU AMD deployment **[secondary — https://www.nasdaq.com/articles/oracle-recently-delivered-incredible-news-advanced-micro-devices-amd-stock-investors; https://www.mitrade.com/insights/news/live-news/article-3-1194409-20251015]**.
- **Meta:** MI300/MI350 series installed base; 2026 procurement shifts toward the MI450/Helios generation (see §7) **[secondary — https://github.com/redhat-et/physical-ai-platform-intel/blob/HEAD/deliverables/intel/companies/amd-deep-dive.md]**.
- **OpenAI:** partnership "began with the MI300X and continued with the MI350X series" en route to MI450 deployments **[secondary — https://en.tmtpost.com/post/7721864]**.
- Rack-scale traction: Pegatron preparing **128-GPU racks** (EIA-compliant 96-GPU variant offers ~27 TB memory); full AMD rack vision: 96–128 MI350-series GPUs, up to ~36 TB HBM3E, ~2.6 exaflops FP4 / ~1.3 exaflops FP8, paired with EPYC "Turin" and Pollara 400 NICs **[secondary — https://webpronews.com (via search result); https://www.tweaktown.com/news/105766/amd-launches-instinct-mi350-series-ai-chips-185-billion-transistors-288gb-hbm3e-memory/index.html]**. ⚠️ The 30,000-unit figure and rack claims come from secondary aggregators.

---

