---
id: etape8-phasea-system-languages/00-system-languages/14-learning-resources
title: "14. Learning resources"
domain: step-8-phase-a-systems-languages
role: deep-dive
task: reference
actors: ["Microsoft", "Qualcomm"]
dates: ["2025-10-11", "2025-10-12", "2025-12", "2026-03-28", "2026-09", "2026-09-18", "2026-09-22"]
keywords: ["acquisition", "benchmark", "cost", "gpu", "latency", "memory", "research", "throughput"]
source: docs/RAG/etape8_phaseA_system_languages.md
source_anchor: ""
source_lines: [507, 596]
section: "Step 8 — Phase A: Systems Languages"
sha256: 642d44e64d9d140357ca04f37ef98d7d4aafc841a83393fbd318f27c487fd458
---

# 14. Learning resources

## 14. Learning resources

Resources below are named from 2026 community consensus; URLs are given only where they appeared verbatim in tool results during this session. Others are named without URLs rather than guessed. [secondary]

- **Rust**: *The Rust Programming Language* ("the Book", official); Rustlings exercises; `exercism.io` Rust track; rust-analyzer LSP; docs.rs for crate documentation. [secondary]
- **Go**: official tour and `go.dev` documentation; the 1.26/1.27 release notes are the authoritative feature references (https://go.dev/doc/go1.26 , https://go.dev/doc/go1.27). [official]
- **C++**: cppreference.com (community reference); WG21 papers for C++26 features; compiler conformance pages (GCC/Clang/MSVC). [secondary]
- **C**: C23 standard text (ISO, paid); compiler documentation for GCC/Clang extensions. [secondary]
- **Zig**: official release notes (https://ziglang.org/download/0.15.1/release-notes.html , https://ziglang.org/download/0.16.0/release-notes.html — URLs as quoted in secondary sources); ZLS language server; community migration guides for 0.14→0.15→0.16. [secondary]
- **Mojo**: official release notes https://mojolang.org/releases/v1.0.0/ (as quoted in the Modular 26.5 release); Modular docs at docs.modular.com; "Mojo AI Skills" marked 1.0-ready for project creation, GPU programming, and porting. Install via `uv pip install --upgrade mojo`. [official]
- **Julia**: official docs and Discourse forum (http://discourse.julialang.org/t/this-month-in-julia-world-2026-07/138558 for the 2026-07 newsletter); SciML tutorials for scientific workloads. [secondary]
- **V**: vlang/v GitHub repository documentation and release notes (https://GitHub.Com/vlang/v/releases). [secondary]
- **Carbon**: Chandler Carruth's 2026 NDC Toronto slide deck (https://github.com/chandlerc/blog/blob/HEAD/content/slides/2026-ndc-toronto-carbon-update/graduation.md) is the freshest roadmap source; no stable tutorial corpus exists yet. [secondary]

---

## 15. Decision matrix

Practical guidance synthesized from 2026 community consensus; tag as [secondary] throughout — this is judgment, not measurement.

| Goal / context | Recommended |
|---|---|
| Maximum performance + memory safety, no GC | **Rust** |
| Microservices / API layer / DevOps tooling | **Go** |
| CLI tools / single static binaries | **Go** or **Zig** |
| Replacing a C/C++ codebase incrementally | **Zig** or **Rust** |
| Embedded / IoT / `no_std` | **Rust** or **Zig** |
| Team with mixed skill levels, ship fast | **Go** |
| Browser engine / game engine / kernel module | **Rust** (or C++ for legacy) |
| Real-time / hard latency deadlines | **Rust** or **Zig** (no GC) |
| Scientific computing / numerical work | **Julia** (or Python + Rust/Julia extensions) |
| AI kernels / heterogeneous accelerators, Python-adjacent team | **Mojo** (watch compiler open-sourcing) |
| Learning systems programming fresh | **Go** or **C**, then **Rust**/**Zig** |
| WASM targeting browsers | **Rust** |

When NOT to pick each (community consensus): [secondary]
- **Not Rust**: tiny team that can't absorb the learning curve; must ship in a week; compile times dominate; domain is CRUD APIs. [secondary]
- **Not Go**: need maximum per-request throughput; writing kernels/game engines; binary size or memory footprint is critical; hard real-time. [secondary]
- **Not Zig**: need a mature library ecosystem; team has no C/C++ experience; need GC-managed memory safety; need a standard hiring pool. [secondary]
- **Not Mojo**: need a fully open-source toolchain today; need async/pattern-matching/tagged-unions now; betting production on one vendor's roadmap. [secondary]
- **Not Julia**: need general-purpose enterprise ecosystem; startup latency is unacceptable; hiring pool matters. [secondary]
- **Not Carbon**: anything production before 0.1 exists — it is a research vehicle. [secondary]
- **Not V**: anything production before 1.0 — pre-release churn. [secondary]

---

## 16. Gaps and conflicts register

1. **C++26 "release" wording**: WG21 finalized the draft 2026-03-28; final ISO publication is a separate step. Sources saying "C++26 released" may mean either. [secondary]
2. **C++26 memory safety**: claims of broad memory-safety guarantees in C++26 are uncorroborated; do not repeat. [unverified]
3. **MSVC 2026 specifics**: no corroborated 2026 MSVC version/feature claims; gap. [unverified]
4. **Carbon 0.1 timing**: slide deck says "maybe early next year" (i.e., ~2027); older snippets predicted late 2026. Conflict preserved; no confirmed date. [secondary]
5. **"Carbon 0.8"**: unsupported single-source claim; rejected. [unverified]
6. **Mojo "first beta of 1.0"**: one InfoWorld result called the August release a beta while Modular's own announcement calls it production-ready 1.0. Prefer the vendor's production-ready framing with the caveat that the compiler stays closed. [secondary]
7. **Mojo compiler open-sourcing**: promised "in 2026", unfulfilled at 2026-09-22; the key adoption-risk variable. [secondary]
8. **Qualcomm/Modular**: acquisition reported mid-2026 (~2 weeks pre-1.0) by secondary press; terms and roadmap impact unconfirmed. [secondary]
9. **Zig 0.15.2 date**: 2025-10-11 vs 2025-10-12 across sources; one-day conflict. [secondary]
10. **Zig release-note URLs**: quoted from secondary sources, not fetched directly in this session; verify before citing as [official]. [secondary]
11. **Go 1.27 extras**: native UUID package, JSON v2 stabilization, experimental SIMD, post-quantum crypto claims appeared in press summaries but not in fetched official notes. [unverified]
12. **Linux kernel 7.x Rust details**: appeared only in weak 2026 articles; the corroborated story ends at the December 2025 Maintainers Summit decision + Android 16 ashmem. [unverified]
13. **"Half of companies use Rust in production"**: single weak survey claim; do not cite. [unverified]
14. **2026 job-board counts**: no hard per-language posting numbers collected; gap. [unverified]
15. **2026 multi-language benchmark**: no clean same-machine/same-workload shootout found; performance comparisons across sources are non-comparable. [unverified]
16. **C2y**: next-C-standard work described only as in-progress drafts; no ratified content at the cutoff. [unverified]
17. **TIOBE URL**: one fetched result had malformed escaping; the TechRepublic URL was captured as the valid reference. Exact September 2026 figures (C 10.28%, C++ 8.67%, Rust 1.34%, Go 1.10%) are from search snippets; re-verify against tiobe.com if precision matters. [independent]
18. **V adoption**: no corroborated production deployments; V 0.5.2 feature list is from release-page reporting, not fetched release notes. [unverified]

---

## 17. Glossary

- **ABI** — Application Binary Interface; the low-level calling/data-layout contract. C's ABI is the cross-language interop standard. [secondary]
- **Borrow checker** — Rust's compile-time analysis enforcing ownership/borrowing rules. [secondary]
- **Bzlmod** — Bazel's module/dependency system (Bazel 8+ primary). [secondary]
- **comptime** — Zig's compile-time code execution model. [secondary]
- **CPS** — Common Package Specification; a build-system-neutral package description format adopted by CMake 4.3. [official]
- **FFI** — Foreign Function Interface; calling code written in another language. [secondary]
- **Green Tea GC** — Go's garbage collector introduced experimentally in 1.25, default in 1.26, focused on small-object locality and CPU scalability. [official]
- **LSP** — Language Server Protocol; editor intelligence protocol (Mojo 1.0 shipped a rewritten server). [official]
- **MLIR** — Multi-Level Intermediate Representation; the compiler framework Mojo builds on for heterogeneous targets. [official]
- **MSRV** — Minimum Supported Rust Version; community convention for crate compatibility floors. [secondary]
- **M:N threading** — many green threads multiplexed onto fewer OS threads (Go's goroutine model). [secondary]
- **RAII** — Resource Acquisition Is Initialization; C++'s deterministic resource-management idiom. [secondary]
- **SARIF** — Static Analysis Results Interchange Format; GCC 16 can emit diagnostics in SARIF/HTML. [secondary]
- **Tier 1 language (Microsoft)** — internal designation granting a language the same investment, tooling, SDL compliance, and platform integration as C++/C#/TypeScript; Rust received it 2026-09-18 per reporting. [secondary]
- **Writergate** — community nickname for Zig 0.15's `std.io` → `std.Io` Reader/Writer overhaul. [secondary]
- **Zero-cost abstractions** — the design claim (Rust, C++, Zig) that high-level constructs compile to code as efficient as hand-written low-level code. [secondary]

---

