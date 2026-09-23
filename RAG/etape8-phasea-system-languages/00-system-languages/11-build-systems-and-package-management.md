---
id: etape8-phasea-system-languages/00-system-languages/11-build-systems-and-package-management
title: "11. Build systems and package management"
domain: step-8-phase-a-systems-languages
role: deep-dive
task: reference
actors: ["AMD"]
dates: ["2026-09"]
keywords: ["amd", "attention", "benchmark", "benchmarks", "cost", "latency", "nvfp4", "prefill", "research"]
source: docs/RAG/etape8_phaseA_system_languages.md
source_anchor: ""
source_lines: [408, 460]
section: "Step 8 — Phase A: Systems Languages"
sha256: 6a940ba0d1cda263953cd7170878a3f852b1457757199aaeb8c4023a9fd89680
---

# 11. Build systems and package management

## 11. Build systems and package management

### 11.1 C/C++ build landscape (2026)

- **CMake 4.x**: current generation; 4.3 added CPS package import/export and instrumentation; 4.4 in development with diagnostics-state tracking and presets schema v12. [official]
- **Meson**: 1.10.1 current per September 2026 toolchain metadata. [secondary]
- **Bazel**: 8.x with bzlmod as the module system; 8.6.0 primary tested version. [secondary]
- **Ninja**: 1.13.2 current; the near-universal fast backend for CMake/Meson. [secondary]
- **GNU Make**: 4.4.1 referenced as current in September 2026 metadata. [secondary]
- **vcpkg / Conan**: the two main C++ package managers; 2026 feature deltas not corroborated (gap). [unverified]

### 11.2 Language-integrated toolchains

| Language | Build + package tool | Notes |
|---|---|---|
| Rust | **Cargo** (rustup for toolchains) | Integrated build/test/docs/publish; MSRV convention; crates.io 200k+ crates [secondary] |
| Go | **go** tool (modules, workspaces) | `go fix` modernizers (1.26–1.27), module proxy + checksum DB, bootstrap 1.24.6+ for 1.26 [official] |
| Zig | **build.zig / build.zig.zon** | Build logic in Zig itself; 0.15 added package fingerprints; `zig cc` cross-compiler [secondary] |
| Julia | **Pkg** (Project/Manifest.toml) | Reproducible per-project environments; registry-based [secondary] |
| Mojo | **modular / max CLI, `uv pip install mojo`** | Install via `uv pip install --upgrade mojo`; MAX package manager integration [secondary] |
| V | **v** CLI (vpm) | Modules via vpm; pre-1.0 churn [secondary] |
| Carbon | CMake/Make/Bazel integration planned for 0.1 | No stable toolchain story yet [secondary] |

---

## 12. Benchmarks and performance signals

### 12.1 Hard rule for this section

Do not rank languages by performance from the numbers below. Each result is bound to its own compiler version, flags, CPU, and workload. Cross-benchmark comparisons are invalid. [secondary]

### 12.2 Compiler shootouts (independent, methodology-bound)

- **GCC 16.1 vs GCC 15.2 vs Clang 21.1.4** (Phoronix, September 2026): AMD Threadripper 9980X, Fedora 44, `-O3 -march=native` across a C/C++ benchmark set. Result: GCC 16 generated faster binaries overall in that setup. This says nothing about Rust/Go/Zig/Mojo/Julia, and nothing about other CPUs or flag sets. [independent]
  - Source: https://www.phoronix.com/review/gcc-16-vs-clang-22
- Older GCC-vs-Clang deltas on different hardware (e.g., EPYC Turin results referenced in community research notes) are separate experiments and must not be merged with the above. [secondary]
  - Source: https://github.com/tamnd/rucc/blob/HEAD/spec/cross-compile/01-research-2026.md

### 12.3 Vendor/official performance claims (workload-bound)

- **Go Green Tea GC**: 10–40% reduction in GC overhead in real-world GC-heavy programs (Go team via InfoWorld). Workload-dependent; not an application-speed claim. [vendor-reported]
- **Go 1.27 allocator**: up to 30% lower small-object allocation cost; ~1% overall improvement for allocation-heavy programs. [official]
- **Go 1.26 cgo**: ~30% lower baseline cgo call overhead. [official]
- **Mojo/Modular**: marketing-grade performance claims (e.g., flash-attention prefill ~1.4x TTFT at seq 8192 on M5; FLUX.2-klein ~1.45x end-to-end via int8 W8A8; NVFP4 requant ~2.56x on FLUX.2-dev) are vendor-reported, hardware- and workload-specific. [vendor-reported]
- **Rust vs C++**: secondary sources describe performance parity with C++ in most scenarios (zero-cost abstractions, no GC); no single 2026 head-to-head benchmark was corroborated here — treat parity claims as community consensus, not measured fact. [secondary]
- **Julia**: performance pitch is C/Fortran-class numerics via LLVM JIT specialization; "time to first X" latency is the known tax, progressively reduced. [secondary]

### 12.4 What is missing

- No 2026-published, methodologically clean, multi-language shootout (same machine, same workload class, current toolchains) was found during this research. The Computer Language Benchmarks Game covers many of these languages but its workload selection and flag choices are themselves contested; no 2026 snapshot was corroborated. [unverified]

---

