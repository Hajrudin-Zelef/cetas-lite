---
id: etape8-phasea-system-languages/00-system-languages/overview
title: "Step 8 — Phase A: Systems Languages"
domain: step-8-phase-a-systems-languages
role: deep-dive
task: reference
actors: ["Qualcomm"]
dates: ["2026-09-22"]
keywords: ["acquisition", "benchmark", "benchmarks", "memory", "research"]
source: docs/RAG/etape8_phaseA_system_languages.md
source_anchor: ""
source_lines: [1, 53]
section: "Step 8 — Phase A: Systems Languages"
sha256: f37804e2afea4cfbf6317ae665bf5919187ccf1d664c7ca16acbb6a1eaf1286a
---

# Step 8 — Phase A: Systems Languages
## C, C++, Rust, Go, Zig, Carbon, Mojo, Julia (and V as secondary)

**Observation cutoff:** 2026-09-22
**Scope:** Systems programming languages and their 2026 toolchains, releases, adoption, performance, memory safety, rankings, learning resources, and job-market signals.
**Languages covered (primary):** C, C++, Rust, Go, Zig, Carbon, Mojo, Julia.
**Languages covered (secondary):** V.

---

## 0. Method and provenance legend

- Research was conducted on 2026-09-22 by web search plus fetching of official documentation and release notes. [secondary]
- Every factual claim in this file carries one of the following tags:
  - `[official]` — observed in vendor/official documentation or release notes fetched during this session.
  - `[vendor-reported]` — stated by a vendor or project team in news/announcements; not independently verified here.
  - `[independent]` — reported by an independent third party (benchmark labs, independent press) using its own methodology.
  - `[secondary]` — reported by secondary/community sources (blog posts, GitHub research notes, community docs).
  - `[unverified]` — a claim seen in weak or single sources that could not be corroborated before the cutoff; do not rely on it.
- Dates are given in ISO format (YYYY-MM-DD). Exact URLs cited were returned verbatim by search tooling on 2026-09-22; where a page could not be fetched, that is noted.
- Version numbers are as reported at the cutoff. The languages below move fast; versions released after 2026-09-22 are out of scope.
- **Benchmark caveat (applies to the whole file):** published language/toolchain performance comparisons use different compilers, flags, CPUs, workloads, and baselines; they are not comparable to each other. Each number is reported with its own methodology and must not be generalized. [secondary]
- **Ranking caveat:** popularity indices (TIOBE, PYPL, Stack Overflow surveys) measure different signals — search traffic, tutorial queries, respondent sentiment — not code volume, production use, or hiring. [independent]

---

## Table of contents

1. [C](#1-c) — C23/C2y, compilers, tooling, safety
2. [C++](#2-c) — C++26, WG21, compilers, ecosystem
3. [Rust](#3-rust) — releases 1.97–1.98, ownership model, tooling, kernel/OS/enterprise adoption
4. [Go](#4-go) — releases 1.24–1.27, generics trajectory, GC, adoption
5. [Zig](#5-zig) — 0.14–0.16, pre-1.0 churn, build system, adoption
6. [Carbon](#6-carbon) — experimental C++ successor, roadmap status
7. [Mojo](#7-mojo) — 1.0, licensing, Qualcomm acquisition
8. [Julia](#8-julia) — 1.13, scientific computing, HPC
9. [V](#9-v-secondary-coverage) — 0.5.2, pre-1.0 status
10. [Memory safety: cross-language comparison](#10-memory-safety-cross-language-comparison)
11. [Build systems and package management](#11-build-systems-and-package-management)
12. [Benchmarks and performance signals](#12-benchmarks-and-performance-signals)
13. [Rankings, adoption, and job-market signals](#13-rankings-adoption-and-job-market-signals)
14. [Learning resources](#14-learning-resources)
15. [Decision matrix](#15-decision-matrix)
16. [Gaps and conflicts register](#16-gaps-and-conflicts-register)
17. [Glossary](#17-glossary)
18. [Source index](#18-source-index)
19. [Appendix A: Chronological timeline](#appendix-a-chronological-timeline-2025-02--2026-09)
20. [Appendix B: Toolchain support matrix](#appendix-b-toolchain-support-matrix-at-the-cutoff)
21. [Appendix C: Stability policies](#appendix-c-stability-and-compatibility-policies-compared)
22. [Appendix D: Open research questions](#appendix-d-open-research-questions-post-cutoff-watch-list)

---

