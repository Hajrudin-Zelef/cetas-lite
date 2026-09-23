---
id: etape8-phasea-system-languages/00-system-languages/13-rankings-adoption-and-job-market-signals
title: "13. Rankings, adoption, and job-market signals"
domain: step-8-phase-a-systems-languages
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Apple", "Meta", "Microsoft", "Nvidia"]
dates: ["2026-09", "2026-09-18"]
keywords: ["amd", "aws", "cost", "energy", "gpu", "nvidia"]
source: docs/RAG/etape8_phaseA_system_languages.md
source_anchor: ""
source_lines: [461, 506]
section: "Step 8 — Phase A: Systems Languages"
sha256: 7f0114d41f29aaab71265d3361040a88d510614071c2511d509e69b7157dbcf0
---

# 13. Rankings, adoption, and job-market signals

## 13. Rankings, adoption, and job-market signals

### 13.1 TIOBE Index — September 2026

Methodology reminder: TIOBE measures search-engine visibility (queries containing "<language> programming"), not usage, code volume, jobs, or quality. [independent]

| Language | Rank | Rating |
|---|---|---|
| C | #2 | 10.28% |
| C++ | #3 | 8.67% |
| Rust | #10 | 1.34% |
| Go | #12 | 1.10% |

- Python held #1 (exact figure not captured in this session's results). Julia did not appear in the fetched top-20; exact rank is a gap. [independent]
- TIOBE's search-visibility methodology systematically favors languages with large tutorial/Q&A footprints (C, C++, Java, Python) and underweights newer or infrastructure languages. [independent]

### 13.2 Developer sentiment (Stack Overflow survey 2025–2026)

- **Rust**: most admired language, 9 years running; 80–83%+ of users want to continue; ~12–14% of respondents use it. [secondary]
- **Zig**: near Rust in "most loved"/admired with ~4% usage — enthusiast skew. [secondary]
- **Go**: ~13% usage; "most loved" rank ~8. [secondary]
- Survey salary medians (self-reported, methodology varies): Rust ~$115k, Zig ~$105k, Go ~$95k. [secondary]
- Sources: https://stackoverflow.co/advertising/resources/stack-overflow-developer-survey-for-employer-branding/insight-2/ , https://github.com/dingjiu1989-hue/dingjiu1989-hue.github.io/blob/HEAD/md/en/compare/rust-go-zig-comparison.md [secondary]

### 13.3 Job-market signals (2026)

- 2026 commentary consistently reported Rust trailing Java, Go, TypeScript, and Python in real-world usage and **job-posting volume**, despite topping admiration rankings — the "won the argument, not the hiring" gap. Barriers cited: talent pool, learning curve, migration cost. [secondary]
  - Source: https://medium.com/rustaceans/rust-won-the-argument-so-why-are-most-companies-still-choosing-the-wrong-language-af17ee9d6138
- Counter-signal: Microsoft's **Tier-1 designation** (2026-09-18) and Linux-kernel permanent adoption are institutional demand signals that typically precede hiring waves, but no 2026 job-board counts were corroborated. [secondary]
- **Gap**: hard 2026 job-posting numbers (LinkedIn/Indeed/devjob boards) per language were not collected in this session; the decision matrix below reflects sentiment and adoption evidence, not posting counts. [unverified]

### 13.4 Adoption snapshot (named production users)

| Language | Named production deployments |
|---|---|
| C | OS kernels, embedded firmware, language runtimes everywhere [secondary] |
| C++ | Game engines, browsers, HPC, quant finance, large legacy estates [secondary] |
| Rust | Linux kernel (Rust-for-Linux), Android, Windows (Win32k, font parsing), AWS Firecracker/Bottlerocket, Cloudflare Pingora/BoringTun, Meta Buck2, Discord, Dropbox, Fastly, Firefox components [secondary] |
| Go | Docker, Kubernetes, Terraform, Prometheus, etcd; Cloudflare-adjacent edge and DevOps tooling broadly [secondary] |
| Zig | TigerBeetle, Ghostty, Bun (parts), ZLS [secondary] |
| Mojo | Modular MAX platform and Modular Cloud (vendor's own stack); custom GPU kernels on NVIDIA/AMD/Apple [official] |
| Julia | SciML/DifferentialEquations.jl ecosystem; climate/energy modeling; finance quant work [secondary] |
| V | No corroborated large-scale production deployments [unverified] |

---

