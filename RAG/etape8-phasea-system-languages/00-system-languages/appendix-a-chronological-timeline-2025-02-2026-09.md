---
id: etape8-phasea-system-languages/00-system-languages/appendix-a-chronological-timeline-2025-02-2026-09
title: "Appendix A: Chronological timeline (2025-02 → 2026-09)"
domain: step-8-phase-a-systems-languages
role: deep-dive
task: reference
actors: ["CISA", "EU", "Microsoft", "Qualcomm"]
dates: ["2025-02-11", "2025-03-05", "2025-05-21", "2025-08-12", "2025-08-19", "2025-10-11", "2025-10-12", "2025-11-26", "2026-01-01", "2026-02-10", "2026-03-28", "2026-04-13", "2026-04-30", "2026-07-12", "2026-08-07", "2026-08-11", "2026-08-16", "2026-08-18", "2026-08-19", "2026-08-20", "2026-09", "2026-09-03", "2026-09-10", "2026-09-18", "2026-09-22"]
keywords: ["cyber", "latency", "license", "memory", "research"]
source: docs/RAG/etape8_phaseA_system_languages.md
source_anchor: ""
source_lines: [674, 752]
section: "Step 8 — Phase A: Systems Languages"
sha256: 62bd85da93cafd28c51bc1f25fac8547787dc1c2316427a8c77a700f01e65feb
---

# Appendix A: Chronological timeline (2025-02 → 2026-09)

## Appendix A: Chronological timeline (2025-02 → 2026-09)

- **2025-02-11** — Go 1.24 released. [secondary]
- **2025-03-05** — Zig 0.14.0 released. [secondary]
- **2025-05-21** — Zig 0.14.1 released. [secondary]
- **2025-08-12** — Go 1.25 released, with experimental Green Tea GC and experimental `encoding/json/v2`. [official]
- **2025-08-19** — Zig 0.15.1 released. [secondary]
- **2025-10-11** — Zig 0.15.2 released (one source: 2025-10-12). [secondary]
- **2025-11-26** — Zig project migrates its repository from GitHub to Codeberg. [secondary]
- **2025-12** — Kernel Maintainers Summit: Rust for Linux deemed no longer experimental, permanent adoption path. [secondary]
- **Late 2025** — CISA Secure by Design pledge passes 296 signatory organizations. [secondary]
- **2026-01-01** — Target date for signatories' memory-safety roadmaps (reported; exact official wording unverified). [secondary]
- **2026-02-10** — Go 1.26 released; Go 1.24 reaches end of life. [official]
- **2026-03-28** — WG21 approves/finalizes the C++26 draft. [secondary]
- **2026-04-13** — Zig 0.16.0 released (current stable at the cutoff). [secondary]
- **2026-04-30** — GCC 16.1 released (first of the 16 series). [secondary]
- **2026-07-12** — V 0.5.2 released. [secondary]
- **Mid-2026** — Qualcomm acquires Modular (reported ~2 weeks before Mojo 1.0). [secondary]
- **2026-08-07** — GCC 16.2 released. [secondary]
- **2026-08-11** — Modular 26.5 / Mojo 1.0 released (production-ready, API-stable commitment). [official]
- **2026-08-16** — Julia 1.10.12 LTS released. [secondary]
- **2026-08-18** — ModCon event, San Francisco (Qualcomm plans for Mojo expected). [secondary]
- **2026-08-19** — Go 1.27 released; Go 1.25 reaches end of life. [official]
- **2026-08-20** — Rust 1.98.0 released. [official]
- **2026-09** — EU Cyber Resilience Act vulnerability-reporting obligations begin. [secondary]
- **2026-09-03** — Rust 1.98.1 released (vtable miscompilation fix). [official]
- **2026-09-10** — Julia 1.13.0 released. [secondary]
- **2026-09-18** — Microsoft designates Rust a Tier 1 internal language (reported). [secondary]
- **2026-09** — TIOBE September 2026 published (C #2, C++ #3, Rust #10, Go #12); Phoronix publishes GCC 16 vs Clang 21.1.4 comparison; Fedora 44 ships GCC 16. [independent]

## Appendix B: Toolchain support matrix (at the cutoff)

| Language | Reference toolchain | Latest verified | Standard/version supported |
|---|---|---|---|
| C | GCC 16.2 / Clang 21–22 | 2026-08-07 (GCC) | C23 (`-std=c23`); C2y drafts only [secondary] |
| C++ | GCC 16.2 / Clang 21–22 / MSVC 17.x | 2026-08-07 (GCC) | C++20 default in GCC 16; experimental C++26 (reflection, contracts) [secondary] |
| Rust | rustc 1.98.1 via rustup | 2026-09-03 | 2024 edition current; 6-week trains [official] |
| Go | Go 1.27 toolchain | 2026-08-19 | Go 1 promise (backward compatibility) [official] |
| Zig | zig 0.16.0 | 2026-04-13 | No stability guarantee pre-1.0 [secondary] |
| Carbon | experimental toolchain | n/a | Pre-0.1; evaluable toolchain planned [secondary] |
| Mojo | Modular 26.5 toolchain | 2026-08-11 | Mojo 1.0 language version; 1.x additive commitment [official] |
| Julia | Julia 1.13.0 (LLVM JIT) | 2026-09-10 | 1.10.x LTS supported alongside [secondary] |
| V | v 0.5.2 | 2026-07-12 | Pre-1.0; APIs may change [secondary] |

## Appendix C: Stability and compatibility policies compared

- **Rust editions** (2015/2018/2021/2024): opt-in language epochs; old code keeps compiling. Six-week release trains; MSRV is a community convention. [secondary]
- **Go 1 promise**: programs written for Go 1 keep compiling and running; generics (1.18+) and recent changes (1.26–1.27) stayed within it. [official]
- **C/C++ ISO standards**: committee-driven (WG14/WG21); compilers implement incrementally with `-std=` flags; ABI stability is a platform concern, not a language guarantee. [secondary]
- **Zig 0.x**: explicitly no stability; breaking changes every minor (0.15 "Writergate" is the canonical example). [secondary]
- **Mojo 1.x**: first stability commitment — additive evolution, breaking changes only with care (C++-style). Compiler still closed at the cutoff. [official]
- **Julia**: follows SemVer-like practice; LTS line (1.10.x) maintained for conservative users. [secondary]
- **V 0.x**: pre-1.0; no stability commitment. [secondary]
- **Carbon**: no released version; no policy yet. [secondary]

## Appendix D: Open research questions (post-cutoff watch list)

1. Does WG21/ISO publish C++26 as an International Standard in late 2026 or 2027, and with what final feature deltas versus the 2026-03-28 draft? [unverified]
2. Does Modular open-source the Mojo compiler in the remainder of 2026 as committed, and what license does it use? [unverified]
3. How does Qualcomm's ownership change Modular's MAX/Mojo roadmap and the ModCon announcements of 2026-08-18? [unverified]
4. Does Carbon 0.1 land in early 2027 as the 2026 slide deck suggested, or slip further? [unverified]
5. Does Zig's 0.16 async-I/O architecture stabilize enough for a 1.0 roadmap announcement? [unverified]
6. Do CISA's memory-safety roadmap expectations and the EU CRA change enterprise language choice measurably in 2027 procurement? [unverified]
7. Does Rust's Microsoft Tier-1 designation translate into measurable job-posting growth in 2026–2027? [unverified]
8. Does `encoding/json/v2` stabilize in a future Go release, and do the 1.27 press claims (UUID package, post-quantum crypto) materialize in official notes? [unverified]
9. Does Julia 1.13+ close the startup-latency gap enough to expand beyond scientific niches? [unverified]
10. Does V reach 1.0, and does any production deployment get corroborated? [unverified]

---

## Verification

- Line count and integrity checks were run on 2026-09-22 after writing.
- Target: at least 750 lines. Result: see `wc -l` output below (must be ≥ 750).
- Markdown sanity: headings nested sequentially, code fences balanced, no stray backticks — checked by inspection script.
- Provenance audit: every factual claim carries one of `[official]`, `[vendor-reported]`, `[independent]`, `[secondary]`, or `[unverified]`; claims that could not be corroborated are marked `[unverified]` and listed in Section 16.
- URLs were copied verbatim from tool results; none were invented or shortened.

*End of file — Step 8 Phase A: Systems Languages.*
