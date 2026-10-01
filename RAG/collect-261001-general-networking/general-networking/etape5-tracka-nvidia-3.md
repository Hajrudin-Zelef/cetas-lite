---
id: collect-261001-general-networking/general-networking/etape5-tracka-nvidia-3
title: "Step 5 — Track A: NVIDIA GPUs (B200/B300/Rubin), DGX Systems & Networking"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Nvidia"]
dates: ["2026-02", "2026-07-13"]
keywords: ["gpu", "gpus", "nvidia", "rubin", "agentic", "blackwell", "cpo", "datacenter", "decode", "ethernet", "fp4", "fp8"]
source: docs/RAG/collect-261001-general-networking/etape5_trackA_nvidia.md
source_anchor: ""
source_lines: [162, 229]
sha256: e598228283801acd629b18491992ba779f47f2ab765107077df63f7a1edc41e4
---

# Step 5 — Track A: NVIDIA GPUs (B200/B300/Rubin), DGX Systems & Networking

### 4.1 DGX Spark (GB10)
- **Specs:** NVIDIA GB10 Grace Blackwell Superchip (20-core Arm: 10× Cortex-X925 + 10× Cortex-A725); Blackwell GPU with 6,144 CUDA cores, 5th-gen Tensor, 4th-gen RT; up to **1 PFLOP (sparse FP4)** AI performance [secondary — https://peterfalkingham.com/2026/09/18/academic-tech-gigabyte-ai-top-nvidia-dgx-spark-review/]
- **Memory:** 128 GB LPDDR5x unified, 273 GB/s; model capacity ~200B params FP4 inference (405B with two units over ConnectX-7) [secondary]
- **Storage:** 4 TB NVMe Gen5, self-encrypting; **Networking:** 10 GbE RJ-45, ConnectX-7 @ 200 Gbps, WiFi 7; **Power:** 240W PSU (GB10 TDP 140W; 40–60W observed at wall under load) [secondary]
- **OS:** NVIDIA DGX OS (Ubuntu 24.04 base) [secondary]
- **Pricing:**
  - Launch MSRP: $3,999 → raised to **$4,699** (Feb 23, 2026) [secondary]
  - Sep 2026 street: NVIDIA marketplace $4,699; Amazon $4,679; retailer markups $4,999.99 (Micro Center via Amazon) to $6,155.99 (Best Buy/Newegg); UK £4,900–£5,485; MSI EdgeXpert variant spiked to $7,342 [secondary — https://www.eteknix.com/rtx-spark-prices-are-already-surging-weeks-before-launch/]
  - Price history: lowest $3,979.19 (Jan 30, 2026), highest $7,346.22 (Sep 15, 2026), average $4,757.26 [secondary — pricehistory.app]
- **Performance (independent):** LMSYS Org — GPT-OSS 20B in Ollama: 2,053 tok/s prefill, 49.7 tok/s decode (~1/4 of RTX Pro 6000 Blackwell WE; slower than single RTX 5090); Llama 3.1 8B @ batch 32: 368 tok/s decode; Sebastian Raschka: single-sample PyTorch inference "roughly on par with the 6× more expensive H100" [independent — https://gpusmith.com/articles/en/pdfs/nvidia-dgx-spark-review-specs-performance.pdf]
- **Issues:** thermal and power-delivery problems on early units — John Carmack publicly reported units capping near 100W of rated 240W and rebooting under sustained load; NVIDIA dev forum moderators confirmed "a known, documented problem across the platform" [independent/secondary]
- **Reception:** Tom's Hardware — "a well-rounded toolkit for local AI... but a pricey platform if you don't intend to use its features to the fullest" [independent]
- **Alternatives (Sep 2026):** Strix Halo 128GB ($2,600–3,600), OEM GB10 boxes (often below FE price), Radeon PRO W7900 ($3,700–3,999), Mac Studio M5 Max (from $2,499), RTX 5090 ($3,700–5,000 street) [secondary — https://memeburn.com/dgx-spark-alternatives/]
- **RTX Spark:** per eTeknix (Sep 21, 2026), "RTX Spark Mini PCs make their market debut" weeks out — prices already surging [secondary — treat launch timing as [unverified]]

### 4.2 DGX Station / DGX Cloud
- No verified 2026 DGX Station (desktop/workstation) refresh details were located in the fetched sources — **gap flagged** [unverified]
- DGX Cloud: NVIDIA's AI supercomputing cloud offering; no standalone 2026 pricing/availability news located in this research pass — **gap flagged** [unverified]
- Context: DGX Cloud was the vehicle for NVIDIA's "AI factory" enterprise push alongside Nemo/Dynamo software at GTC 2026 [secondary]

---

## 5. Networking: NICs, DPUs, switches, NVLink

### 5.1 ConnectX NICs
- **ConnectX-7:** 200 Gb/s; ships in DGX Spark as SmartNIC; still the volume NIC for H100/H200/B200-era deployments [secondary]
- **ConnectX-8 SuperNIC:** up to 800 Gb/s total (2× 400 Gb/s ports); PCIe Gen6 (up to 48 lanes), backward-compatible with Gen5/4/3; advanced QoS and congestion control; performance isolation for multi-tenant AI factories [official — Spectrum-X datasheet]
- **ConnectX-9 SuperNIC:** firmware v82.48.1000 GA **February 2026** [official — docs.nvidia.com]; up to 800 Gb/s per port over InfiniBand and Ethernet; up to **1.6 Tb/s throughput to Rubin GPUs**; programmable IO and intelligent congestion control; pairs with Spectrum-X Ethernet and Quantum-X800 [official]
- **ConnectX-10 (CX10):** roadmap — paired with Feynman (2028) [secondary]

### 5.2 BlueField DPUs
- **BlueField-3:** current-gen networking/security/storage offload; BlueField-3 SuperNIC delivers up to 400GbE RoCE with GPUDirect RDMA; HHHL form factor, sub-75W envelope (B3140H) [official/secondary]
- **BlueField-4:** announced at GTC 2026 as part of the Vera Rubin full-stack platform; extends infrastructure offload into agentic-AI storage via STX/CMX context-memory systems [secondary — https://github.com/hczhu/stock-research/blob/HEAD/memos/2026-07-13-nvidia-hardware-lineup-ai-factory-ecosystem.md]
- **BlueField-5:** future roadmap step paired with Feynman/Rosa-era infrastructure [secondary]
- Deployment note: in Spectrum-X reference topologies, BF3 supports single-plane operation only; north-south (management/storage) traffic handled by BlueField DPUs [secondary]

### 5.3 Spectrum switches and Spectrum-X
- **Spectrum-4 (SN5600/SN5600D/SN5610):** 64 ports of 800GbE in 2U; **51.2 Tb/s total throughput**; smart-leaf/spine/superspine and rail-optimized designs; 10–800GbE connectivity [official — Spectrum-X datasheet]
- **Spectrum-X800 platform:** SN5600 800Gb/s switch + BlueField-3 SuperNIC; optimized for multi-tenant genAI clouds with performance isolation per tenant [official — investor.nvidia.com]
- **Spectrum-6:** **102.4 Tb/s** switch for the Vera Rubin generation (2× Spectrum-4) [secondary]
- **Spectrum-X with co-packaged optics (CPO):** announced at GTC 2026 as "the world's first mass-produced CPO switch" — 5× optical power efficiency vs pluggable optics, 2 Tb/s per-port bandwidth, 10× network reliability [secondary — https://medium.com/@strategycheatsheet/key-takeaways-from-jensen-huangs-nvidia-gtc-2026-keynote-speech-aefbae9de255]
- **Spectrum-7:** roadmap with Feynman (2028) [secondary]
- Vera Rubin reference: Grace Blackwell = Spectrum-4 + ConnectX-8 (800 Gb/s); Vera Rubin = Spectrum-6 + ConnectX-9 (1.6 Tb/s) [secondary]

### 5.4 InfiniBand: Quantum-X800
- Quantum-X800 platform: Quantum Q3400 switch + ConnectX-8 SuperNIC; end-to-end 800 Gb/s; 5× bandwidth capacity and 9× SHARPv4 In-Network Computing (14.4 TFLOPS) vs previous generation; SHARPv4 supports FP8 [official — investor.nvidia.com]
- Rubin NVL144 CPX networking: Quantum-X800 InfiniBand or Spectrum-X Ethernet with ConnectX-9 SuperNICs [secondary]

### 5.5 NVLink scale-up fabric
| Architecture | NVLink gen | Per-GPU bandwidth | NVL72 aggregate |
|---|---|---|---|
| Hopper | NVLink 4 | 900 GB/s | n/a |
| Blackwell | NVLink 5 | 1.8 TB/s | 130 TB/s |
| Rubin | NVLink 6 | 3.6 TB/s | 260 TB/s |
| Feynman | NVLink 8 | (not disclosed) | (not disclosed) |
[secondary]
- NVIDIA adding optical scale-up alongside copper so NVLink domains can expand beyond a single rack (Kyber networking architecture) [secondary]
- NVLink Switch trays: 9 per Vera Rubin NVL72 rack [secondary]

### 5.6 200/400/800 GbE adoption
- 800GbE is the current datacenter standard for AI fabrics (Spectrum-4 64× 800GbE; ConnectX-8 800Gb/s) [official]
- ConnectX-9 pushes per-port 800 Gb/s and 1.6 Tb/s aggregate to Rubin GPUs [official]
- CPO switches (2 Tb/s per port) mark the transition toward 1.6T/3.2T-class optics with the Feynman generation [secondary]

---

## 6. GTC 2026 (March 16–19, San Jose) — key announcements

