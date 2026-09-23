---
id: etape8-phasea-system-languages/00-system-languages/9-v-secondary-coverage
title: "9. V (secondary coverage)"
domain: step-8-phase-a-systems-languages
role: deep-dive
task: reference
actors: ["CISA", "EU", "United States"]
dates: ["2026-01-01", "2026-07-12", "2026-09", "2027-12"]
keywords: ["cyber", "memory"]
source: docs/RAG/etape8_phaseA_system_languages.md
source_anchor: ""
source_lines: [363, 407]
section: "Step 8 — Phase A: Systems Languages"
sha256: 228c77f02c809336b267c2314aebeea08a3c13c41fc3712c79fa039c0a5df05c
---

# 9. V (secondary coverage)

## 9. V (secondary coverage)

### 9.1 Release status

- **V 0.5.2** — released **2026-07-12** per GitHub releases. [secondary]
  - Source: https://GitHub.Com/vlang/v/releases
- V remains **pre-1.0**; its own README states core APIs may change before 1.0. [secondary]
  - Source: https://github.com/vlang/v/blob/HEAD/README.md
- 0.5.2 highlights reported: experimental **ownership** support behind the `-ownership` flag, the **VGC** (V garbage collector work), a **V2 backend**, C and JavaScript backends, and removal of deprecated `vweb`. [secondary]

### 9.2 Language model and positioning

- V pitches itself as a simple, fast, safe compiled language with Go-like ergonomics and C-level performance; it compiles to C (transpiled) as one backend strategy. [secondary]
- Memory-safety story: VGC/autofree concepts plus the experimental ownership flag; the safety model is less formally specified than Rust's and should not be equated with it. [secondary]
- Security policy is documented in the repo (`SECURITY.md`). [secondary]
  - Source: https://github.com/vlang/v/blob/HEAD/SECURITY.md
- Adoption remained niche at the cutoff: no corroborated production deployments on the scale of Go/Rust/Zig's named users. Treat adoption claims as [unverified].

---

## 10. Memory safety: cross-language comparison

| Language | Safety model | GC? | Escape hatches | Notes |
|---|---|---|---|---|
| C | None by construction | No (manual) | N/A — everything is raw | Sanitizers/fuzzing/CHERI are mitigations, not guarantees [secondary] |
| C++ | None by default | No (manual/RAII) | Raw pointers, `reinterpret_cast` | Profiles/hardened libs in progress; C++26 adds no default safety [secondary] |
| Rust | Ownership/borrow checker, `Send`/`Sync` | No | `unsafe` blocks | Soundness depends on auditing `unsafe`; the only language here with a formal safety-by-default story [secondary] |
| Go | GC + bounds checks | Yes | `unsafe` package | GC pauses shrinking (Green Tea); memory-safe for ordinary code [official] |
| Zig | Runtime checks in safe build modes | No (explicit allocators) | Everything is explicit; safety is a build-mode choice | No type-system guarantee; Debug/ReleaseSafe vs ReleaseFast [secondary] |
| Carbon | Design goal, experimental | Planned | TBD | No production story at the cutoff [secondary] |
| Mojo | Rust-inspired concepts + 1.0 diagnostics | Planned/GC-adjacent | Python interop boundary | Vendor-reported; not equivalent to Rust's guarantees [vendor-reported] |
| Julia | GC + bounds checks | Yes | `ccall`/`Ptr` | Safe for ordinary code; C interop is raw [secondary] |
| V | VGC/autofree + experimental ownership flag | Partial | C backend boundary | Less formally specified; pre-1.0 [secondary] |

### 10.1 Policy context (2026)

- CISA/NSA and allied guidance (US, plus partner agencies) continued to recommend memory-safe languages for new code and incremental migration of high-risk legacy components. [secondary]
- Android's reported memory-safety bug share fell from **76% to 24%** as memory-safe languages displaced C in new code — a widely cited datapoint; attribute it as correlation in a large uncontrolled migration, not a controlled experiment. [secondary]
- CISA Secure by Design pledge: 296+ organizations signed by late 2025; memory-safety roadmaps targeted around 2026-01-01. Exact official wording was not re-verified; do not cite as a hard legal deadline. [secondary]
- EU Cyber Resilience Act: vulnerability-reporting obligations began September 2026; main product obligations apply December 2027. German BSI NIS-2 guidance names Rust explicitly. [secondary]
- CISA + FBI issued a 2026 Secure by Design alert specifically on buffer overflows. [secondary]
  - Sources: https://cyberpress.org/cisa-launches-guide-to-mitigate-memory-safety-vulnerabilities/ , https://www.infosecurity-magazine.com/news/cisa-fbi-buffer-overflow/ , https://siliconangle.com/2024/06/27/cisa-joint-guidance-warns-memory-safety-vulnerabilities-open-source-projects/

---

