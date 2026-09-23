---
id: etape8-phasea-system-languages/00-system-languages/5-zig
title: "5. Zig"
domain: step-8-phase-a-systems-languages
role: deep-dive
task: reference
actors: ["Google"]
dates: ["2025-03-05", "2025-05-21", "2025-08-19", "2025-10-11", "2025-10-12", "2025-11-26", "2025-12", "2026-04-13"]
keywords: ["claude", "memory"]
source: docs/RAG/etape8_phaseA_system_languages.md
source_anchor: ""
source_lines: [245, 304]
section: "Step 8 — Phase A: Systems Languages"
sha256: 0689ad451969b5be9d107f3647123eff774981b09490a3238d3b12dab7fb47ea
---

# 5. Zig

## 5. Zig

### 5.1 Project status: pre-1.0

- Zig remains **pre-1.0** at the cutoff; the 0.x series carries **no stability guarantee** and routinely ships breaking changes. The 0.15.1 release notes themselves warn that "working on a non-trivial project using Zig may require participating in the development process." [secondary]
- On **2025-11-26**, the Zig project announced migration of its repository from **GitHub to Codeberg**; the GitHub repository is now a read-only mirror, and `codeberg.org/ziglang/zig` holds the standard library source. Nightly master builds paused during the migration and resumed in December 2025. [secondary]
  - Source: https://github.com/redhat-et/ripwire/blob/HEAD/prompts/help-wanted/zig-language.md (quoting https://ziglang.org/news/migrating-from-github-to-codeberg/)
- Consequence for researchers: ziglang stopped tagging GitHub releases after 0.15.2, so the GitHub releases API under-reports; the authoritative list is https://ziglang.org/download/index.json. [secondary]

### 5.2 Release history (corrected dates)

Dates below follow the community correction against ziglang.org/download; mark 0.15.0 as **retracted/not officially released**: [secondary]

- **0.14.0** — 2025-03-05 [secondary]
- **0.14.1** — 2025-05-21 [secondary]
- **0.15.1** — 2025-08-19 [secondary]
- **0.15.2** — 2025-10-11 (one source says 2025-10-12; one-day conflict noted) [secondary]
- **0.16.0** — **2026-04-13** (current stable at the cutoff per community plugin sources) [secondary]
  - Source: https://github.com/jkingston/zig_guide/commit/39fa4aae23106505550002b6379741505622b5ac , https://github.com/vinnie357/claude-skills/blob/HEAD/plugins/languages/zig/skills/sources.md
- Official release notes: https://ziglang.org/download/0.15.1/release-notes.html and https://ziglang.org/download/0.16.0/release-notes.html (URLs as quoted in secondary sources; pages not directly fetched in this session). [secondary]

### 5.3 Language and toolchain characteristics

- **comptime**: Zig's compile-time execution model — generics are functions evaluated at `comptime`, types are values; enables metaprogramming without a macro system. [secondary]
- **Explicit allocators**: no hidden allocation; every allocating API takes an allocator, making memory behavior explicit and auditable. [secondary]
- **Safety modes**: bounds checking and other runtime safety checks in Debug/ReleaseSafe, strippable in ReleaseFast/ReleaseSmall — safety is a build-mode choice, not a type-system guarantee like Rust. [secondary]
- **C interoperability**: first-class `@cImport`/`@cInclude` (0.16 deprecates `@cImport` in favor of `b.addTranslateC()`); the Zig toolchain ships C/C++ compilation via Clang/LLVM integration, and `zig cc` is used as a drop-in cross-compiler. [secondary]
- **Cross-compilation**: a headline feature — targeting dozens of platforms from one toolchain invocation without external sysroots in the common cases. [secondary]
- **Self-hosted backends**: 0.15.1 notes describe the new self-hosted x86_64 backend as the Debug default, with LLVM still selectable via `-fllvm`. [secondary]
- **"Writergate" (0.15)**: complete overhaul of `std.io.Reader`/`std.io.Writer` → `std.Io` with explicit caller-provided buffering ("Please use buffering! And don't forget to flush!"); HTTP client/server reworked to depend on I/O streams; `std.ArrayList` became unmanaged by default; `usingnamespace` and old `async`/`await` removed; `root_source_file` removed from build options in favor of `root_module`; `addStaticLibrary`/`addSharedLibrary` replaced by `addLibrary`; `{f}` format specifier for custom format methods. [secondary]
  - Source: https://github.com/raymond-w-ko/dullahan/blob/HEAD/docs/zig-0.15-notes.md , https://dev.to/bkataru/zig-0151-io-overhaul-understanding-the-new-readerwriter-interfaces-30oe
- **build.zig / build.zig.zon**: Zig's build system is written in Zig itself; 0.15 made the package `name` an enum literal and added a required `fingerprint` field for package identity. [secondary]

### 5.4 Adoption

- Named production users: **TigerBeetle** (financial transactions database, overwhelmingly Zig), **Ghostty** (terminal emulator, mostly Zig), **Bun** (JavaScript runtime, uses Zig for parts), **ZLS** (Zig language server, written in Zig). [secondary]
- Zig ranked near Rust in "most loved"/"most admired" in survey commentary with only ~4% of respondents using it — typical enthusiast skew for a young language. Reported median salary ~$105k (secondary). [secondary]
- When NOT to pick Zig (community consensus): need a mature library ecosystem; team lacks C/C++ experience; shipping production software needing GC-managed memory safety; need to hire from a standard talent pool. [secondary]
  - Source: https://github.com/dingjiu1989-hue/dingjiu1989-hue.github.io/blob/HEAD/md/en/compare/rust-go-zig-comparison.md

---

## 6. Carbon

### 6.1 What it is

- Carbon is Google's **experimental successor language for C++**, designed for **bidirectional interoperability** with existing C++ codebases — the strategy is migration, not a clean break. [secondary]
- It is not production-ready and has no stability commitment at the cutoff. [secondary]
- Source: https://en.wikipedia.org/wiki/Carbon_(programming_language) [secondary]

### 6.2 Roadmap status (conflicted)

- A 2026 Chandler Carruth slide deck (NDC Toronto Carbon update) states **Carbon 0.1** aims to deliver an evaluable compiler/toolchain, CMake/Make/Bazel integration, and most C++ interop; release timing was phrased as "**maybe early next year**" — not a confirmed date. [secondary]
  - Source: https://github.com/chandlerc/blog/blob/HEAD/content/slides/2026-ndc-toronto-carbon-update/graduation.md
- Wikipedia/search snippets predicting a late-2026 Carbon 0.1 **conflict** with the newer slide deck; treat any "Carbon 0.1 released" claim as [unverified] unless corroborated.
- An unsupported Medium claim about "Carbon 0.8" should be rejected — no such release was corroborated. [unverified]
- Memory-safety posture: Carbon's design goals include memory safety, but as an experimental language it offered no production safety story at the cutoff. [secondary]

---

