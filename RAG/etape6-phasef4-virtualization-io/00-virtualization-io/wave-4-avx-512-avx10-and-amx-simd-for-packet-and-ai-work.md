---
id: etape6-phasef4-virtualization-io/00-virtualization-io/wave-4-avx-512-avx10-and-amx-simd-for-packet-and-ai-work
title: "Wave 4 — AVX-512, AVX10, and AMX: SIMD for packet and AI work"
domain: phase-f4-i-o-virtualization-cpu-acceleration-extensions
role: deep-dive
task: reference
actors: ["AMD", "Intel", "TSMC"]
dates: []
keywords: ["18a", "amd", "claude", "consumer", "fp8", "inference", "intel", "latency", "licenses", "memory", "research", "throughput"]
source: docs/RAG/etape6_phaseF4_virtualization_io.md
source_anchor: ""
source_lines: [186, 234]
section: "Phase F4 — I/O Virtualization & CPU Acceleration Extensions"
sha256: 9b43139b47a2b0e8c1917e4841d4f417f9d3c44ec76a0add123f40ec3c84a462
---

# Wave 4 — AVX-512, AVX10, and AMX: SIMD for packet and AI work

## Wave 4 — AVX-512, AVX10, and AMX: SIMD for packet and AI work

### 4.1 AVX-512 generational support on Xeon Scalable
- AVX-512 debuted on Skylake-SP/Cascade Lake (Xeon Scalable 1st/2nd gen) with "light" vs "heavy" 512-bit frequency licenses — heavy AVX-512 measured cutting a Xeon Silver 4116 from 2.1 GHz base to ~1.4 GHz all-core (~33%), the origin of the "AVX-512 throttles" folklore [secondary].
  Source: https://github.com/sunstoneinstitute/horndb/blob/HEAD/crates/simd/simd-research-findings.md
- Ice Lake-SP (3rd gen, Sunny Cove): AVX-512 F/CD/VL/DQ/BW/IFMA/VBMI/VBMI2/VNNI/VPOPCNTDQ/BITALG/GFNI/VAES/VPCLMULQDQ — no BF16, no FP16, no AMX; Intel progressively relaxed the light/heavy licensing distinction [secondary].
  Source: https://github.com/adaworldapi/ndarray/blob/HEAD/.claude/knowledge/td-simd-cpu-dispatch-matrix.md
- Sapphire Rapids (4th gen, family 6 model 143, e.g. Xeon Gold 5412U): full AVX-512 superset incl. BF16 + AMX (INT8/BF16 tile ops) [secondary].
  Source: https://github.com/adaworldapi/ndarray/blob/HEAD/.claude/knowledge/td-simd-cpu-dispatch-matrix.md
- Emerald Rapids (5th gen, family 6 model 0xCF): die-shrink of Sapphire Rapids on Raptor Cove, same SIMD/AMX ISA — "Dispatch handled by the SapphireRapids profile" in SIMD runtimes; up to 64 cores, 5 MB L3/core, DDR5-5600, 80× PCIe 5.0 lanes [secondary].
  Source: https://github.com/adaworldapi/ndarray/blob/HEAD/.claude/knowledge/td-simd-cpu-dispatch-matrix.md
  Source: https://en.wikipedia.org/wiki/Emerald_Rapids
- Granite Rapids (Xeon 6, Redwood Cove, Intel 3): AVX-512 superset + **AMX-FP16** (TDPFP16PS, CPUID.07H.1H:EAX bit 21); up to 1024 BF16/FP16 + 2048 INT8 FLOPS per core per cycle via AMX; Xeon 6900P up to 128 P-cores [secondary].
  Source: https://github.com/adaworldapi/ndarray/blob/HEAD/.claude/knowledge/td-simd-cpu-dispatch-matrix.md
- Phoronix on Emerald Rapids (Platinum 8592+): AVX-512 workloads showed doubled/tripled performance with "no major power usage or heat generation issues" — the throttling folklore does not transfer to current parts [independent].
  Source: https://www.cryptopolitan.com/intels-5th-gen-xeon-avx-512-boosts/
- AMD side: EPYC Genoa has 256-bit AVX-512 datapath (double-pumped); **Zen 5 exposes full 512-bit AVX-512 on desktop CPUs** — noted as pressure on Intel to restore wide vectors client-side [secondary].
  Source: https://videocardz.com/newz/intel-documents-confirm-avx10-support-on-next-gen-nova-lake
- [unverified] Granite Rapids AVX-512-FP16 specifics and exact AMX throughput figures above come from community SIMD-dispatch research, not Intel ARK — treat as indicative.

### 4.2 AVX10: the convergence answer (Hot Chips 2026 / Diamond Rapids)
- **AVX10.2** replaces AVX10.1 (which was P-core-only): converged 128/256/512-bit vector execution on **both P-cores and E-cores**, fixing the hybrid mismatch that forced AVX-512 off on Alder Lake-class client parts [secondary].
  Source: https://www.techtimes.com/articles/325660/20260826/diamond-rapids-disclosed-intel-xeon-7-packs-256-cores-gigabyte-cache-drops-smt.htm
- Diamond Rapids (Xeon 7, disclosed Hot Chips Aug 2026): up to 256 cores, 1.28 GB LLC, 16 memory channels (DDR5-8000 / MRDIMM 12800, up to 1.6 TB/s), 128× PCIe 6.0/CXL 3.0/UPI 3 lanes, Intel 18A-P — first server product shipping **AVX10.2 + APX + enhanced AMX** together [secondary].
  Source: https://www.techtimes.com/articles/325660/20260826/diamond-rapids-disclosed-intel-xeon-7-packs-256-cores-gigabyte-cache-drops-smt.htm
  Source: https://www.eweek.com/news/intel-diamond-rapids-256-core-xeon-amd-venice/
- **APX (Advanced Performance Extensions):** doubles GPRs 16→32 (r16–r31), three-operand encodings; Intel claims −10% memory loads / −20% stores for recompiled code, full binary backward compatibility [vendor-reported].
  Source: https://www.techtimes.com/articles/325660/20260826/diamond-rapids-disclosed-intel-xeon-7-packs-256-cores-gigabyte-cache-drops-smt.htm
- **Enhanced AMX** on Diamond Rapids adds **FP8** support for on-CPU ML workloads (inference and training matrix math), building on SPR's AMX [secondary].
  Source: https://www.theregister.com/hpc/2026/08/25/intel-diamond-rapids-xeon-7-cpu-deep-dive/5292427
- Diamond Rapids drops SMT (256 cores / 256 threads) vs AMD EPYC Venice (256 cores / 512 threads expected, TSMC N2) — the ISA story is Intel's differentiator [secondary].
  Source: https://www.eweek.com/news/intel-diamond-rapids-256-core-xeon-amd-venice/
- Linux enablement: KVM patches expose AVX10.2, expanded AMX (new formats/memory ops), MOVRS, and AVX10_VNNI_INT to guest VMs (KVM previously topped out at AVX10.1) [secondary].
  Source: https://videocardz.com/newz/intel-documents-confirm-avx10-support-on-next-gen-nova-lake (via embedded WorldOfSoftware report)
- Intel ISA Programming Reference Rev 060 (Nov 2025) officially ties AVX10.1/10.2 + APX to upcoming processors incl. Nova Lake client parts — first official sign of wide-vector return to consumer CPUs [official].
  Source: https://videocardz.com/newz/intel-documents-confirm-avx10-support-on-next-gen-nova-lake
- [unverified] AMD has not announced AVX10/APX adoption plans; software must dispatch AVX-512 vs AVX10 paths for the foreseeable future.

### 4.3 AMX for AI and networking-adjacent work
- AMX (Advanced Matrix Extensions): tile-based matrix engine (8× 1 KB tiles), TMUL unit; SPR: INT8/BF16; GNR adds FP16; DMR adds FP8 + new memory ops [secondary].
  Source: https://github.com/adaworldapi/ndarray/blob/HEAD/.claude/knowledge/td-simd-cpu-dispatch-matrix.md
  Source: https://www.theregister.com/hpc/2026/08/25/intel-diamond-rapids-xeon-7-cpu-deep-dive/5292427
- Relevance to networking: AVX-512/AVX-VNNI accelerate crypto (AES-GCM, SHA via VAES/VPCLMULQDQ), compression, packet classification, and 5G L1 (FlexRAN) — Emerald Rapids "-N" SKUs target network/5G/edge with high-TPT/low-latency tuning [secondary].
  Source: https://en.wikipedia.org/wiki/Emerald_Rapids
- oneDNN added early AVX10.2 512-bit support for "future Intel Core" CPUs — the AI-inference library path is being ported ahead of silicon [secondary].
  Source: https://videocardz.com/newz/intel-documents-confirm-avx10-support-on-next-gen-nova-lake
- SKU-selection note for buyers: not all Xeon SKUs expose the full ISA equally in practice (frequency licenses historically, power envelopes); verify `cpuinfo` flags on the actual SKU stepping before committing vectorized data-plane code [secondary].

---
