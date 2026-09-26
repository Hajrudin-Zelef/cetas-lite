---
id: etape8-phasea-system-languages/00-system-languages/part-11
title: "Step 8 — Phase A: Systems Languages (part 11)"
domain: step-8-phase-a-systems-languages
role: deep-dive
task: reference
actors: ["Microsoft"]
dates: ["2026-09-18"]
keywords: ["acquisition", "cost"]
source: docs/RAG/etape8_phaseA_system_languages.md
source_anchor: ""
source_lines: [578, 596]
section: "Step 8 — Phase A: Systems Languages"
sha256: f38ee771a2ea7be77c2c10f23e21d2b25ba8017c4d0ab7dd2f136973a72b3cfb
---

# Step 8 — Phase A: Systems Languages (part 11)

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

