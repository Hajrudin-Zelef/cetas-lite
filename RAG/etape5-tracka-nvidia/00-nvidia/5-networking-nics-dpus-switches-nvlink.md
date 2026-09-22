---
id: etape5-tracka-nvidia/00-nvidia/5-networking-nics-dpus-switches-nvlink
title: "5. Networking: NICs, DPUs, switches, NVLink"
domain: step-5-track-a-nvidia-gpus-b200-b300-rubin-dgx-systems-netwo
role: deep-dive
task: hardware
actors: ["AWS", "Cohere", "CoreWeave", "Google", "Groq", "Mistral", "Nvidia", "OpenAI", "Perplexity"]
dates: ["2026-02", "2026-07-13", "2026-09-22"]
keywords: ["nvlink", "acquisition", "agent", "agentic", "aws", "blackwell", "cohere", "compute", "cpo", "datacenter", "disaggregated", "ethernet"]
source: docs/RAG/etape5_trackA_nvidia.md
source_anchor: ""
source_lines: [184, 269]
section: "Step 5 — Track A: NVIDIA GPUs (B200/B300/Rubin), DGX Systems & Networking"
sha256: 2a1686188f379ec12c62d40d2ed234926935ed0afd70bc7a1662536d9576ac3d
---

# 5. Networking: NICs, DPUs, switches, NVLink

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

- **Vera Rubin full-stack platform:** full production confirmed; 7 chips, 5 rack-scale systems; partner availability H2 2026 [independent/secondary]
- **Feynman architecture:** first public preview — Rosa CPU, LP40 LPU, BlueField-5, CX10, Kyber co-packaged optics, NVLink 8, Spectrum-7 [secondary]
- **Groq acquisition:** NVIDIA confirmed acquisition of Groq's team and technology (LPUs integrated into Rubin/Feynman inference strategy) [secondary — https://acaderesearch.com/nvidia-gtc-2026-the-inference-inflection-vera-rubin-and-the-ai-factory-era/]
- **NemoClaw:** open-source enterprise AI agent framework (Wired leak March 10; briefings to Salesforce, Cisco, Google, Adobe, CrowdStrike) [secondary]
- **Dynamo:** inference operating system / cluster orchestration for disaggregated inference [secondary]
- **Nemotron 3:** Super/Ultra/Nano model families; Nemotron Coalition partners (Mistral, Perplexity, Cursor, Cohere) [secondary]
- **DLSS 5:** neural rendering beyond frame generation; "1,000,000× leap" in path-tracing performance claimed [vendor-reported/secondary]
- **Physical AI:** DRIVE Hyperion partnerships (Hyundai, BYD, Nissan) for Level 4 autonomy; T-Mobile AI-RAN partnership; Nokia deploying RTX PRO 4500 Blackwell Server Edition in AI-RAN base stations [secondary]
- **Spectrum-X CPO switch:** world's first mass-produced co-packaged-optics switch [secondary]
- **AWS partnership:** 1M+ NVIDIA GPUs across AWS regions (Blackwell + Rubin, RTX PRO Blackwell Server Edition, Groq 3 LPUs) [secondary]
- **Space-1 Vera Rubin:** orbital AI data centers concept [secondary — marketing announcement, treat as ambition]
- **Demand projection:** AI infrastructure demand doubled to **$1T through 2027** [vendor-reported]
- Context: keynote followed a ~14% stock pullback after a record $68.1B revenue quarter [secondary]

---

## 7. Verification log (open items)

1. **B200 usable memory** — 192 GB vs 180 GB "usable" conflict across secondary spec tables. [unverified]
2. **B300 FP8/FP4 sparse figures** — 9/18 vs 9/20 PFLOPS (B200) and ~28–30 PFLOPS sparse (B300) vary by source. [unverified]
3. **B300 FP8 compute** — 4.5 vs ~11–13.5 PFLOPS conflict. [unverified]
4. **336B-transistor Rubin figure** — from secondary sources; not verified against NVIDIA. [unverified]
5. **"H300" rebranding of Rubin R100** at AWS/Google Cloud — single secondary source. [unverified]
6. **Rubin Ultra configuration** — NVL576/576 GPUs vs 144 vertical-tray GPUs conflict. [unverified]
7. **OpenAI Q3 2026 Rubin deployment at scale** — Bloomberg via secondary relay. [unverified]
8. **DGX Station 2026 refresh** — no verified details located. [gap]
9. **DGX Cloud 2026 pricing/availability news** — no verified details located. [gap]
10. **RTX Spark launch timing** — "weeks before launch" per eTeknix Sep 21, 2026. [unverified]
11. **HBM4 sold out through 2026** — GTC 2026 secondary reporting. [unverified]
12. **CoreWeave 10× token output claim** — secondary relay of vendor/partner claim. [unverified]
13. **Space-1 orbital datacenters** — marketing concept, no deployment details. [unverified]
14. **Groq acquisition terms** — confirmed as announced but terms/price not verified in fetched sources. [unverified]

---

## 8. Collection metadata

- **Research date:** September 22, 2026
- **Sources prioritized:** NVIDIA official (docs.nvidia.com, investor.nvidia.com, datasheets), DCD, EE Times, VideoCardz, ServeTheHome-adjacent coverage, SemiAnalysis, Tom's Hardware, LMSYS, IntuitionLabs, GPUaaS, GPUSmith
- **Conventions:** all prices in USD unless noted; cloud rates are per-GPU-hour; "street" prices are retailer snapshots dated Sep 2026; sparse vs dense FLOPS always distinguished where sources allow
