---
id: etape8-phasea-system-languages/00-system-languages/3-rust
title: "3. Rust"
domain: step-8-phase-a-systems-languages
role: deep-dive
task: reference
actors: ["AWS", "Google", "Lambda", "Meta", "Microsoft"]
dates: ["2025-12", "2026-05", "2026-08-20", "2026-09", "2026-09-03", "2026-09-18", "2026-09-22"]
keywords: ["aws", "compute", "cost", "memory", "research"]
source: docs/RAG/etape8_phaseA_system_languages.md
source_anchor: ""
source_lines: [125, 192]
section: "Step 8 — Phase A: Systems Languages"
sha256: 7afd1ebb78c8dc5a1b6bb288f76f6db69d7ed9845e49c5a4bcc54c25b0487775
---

# 3. Rust

## 3. Rust

### 3.1 Language model

- Rust's core claim is **memory safety without a garbage collector** via the ownership/borrowing model: each value has one owner, borrows are checked at compile time, and data races are rejected by the type system (`Send`/`Sync`). [secondary]
- The language has a **safe/unsafe boundary**: `unsafe` blocks allow raw pointer dereference, unchecked indexing, and FFI; the compiler makes no safety guarantees inside them, and auditing `unsafe` usage is the standard practice for soundness review. [secondary]
- Rust has no `null`; absence is expressed with `Option<T>`, forcing explicit handling. Array bounds are checked at runtime by default (panic on violation). [secondary]
- Concurrency: threads plus `async/await`; `tokio` is the de facto async runtime in 2026. There is no M:N green-thread scheduler in the standard library. [secondary]
- Compilation model: ahead-of-time native compilation via rustc (LLVM backend), with monomorphized generics (zero-cost abstractions when optimized). [secondary]

### 3.2 Releases 2025–2026 (official release history)

Observed from the official release history (rust-lang/rust `RELEASES.md`) on 2026-09-22: [official]

- **Rust 1.97**: changed the stable default symbol mangling scheme to `v0`; added Cargo warning-control behavior. [official]
  - Source: https://github.com/rust-lang/rust/blob/HEAD/RELEASES.md , https://linuxiac.com/rust-1-97-released-with-new-default-symbol-mangling-scheme/
- **Rust 1.98.0** — released **2026-08-20**. [official]
- **Rust 1.98.1** — released **2026-09-03**, fixing a vtable-generation miscompilation. [official]
- Rust follows a six-week train release model with stable/beta/nightly channels; editions (2015/2018/2021/2024) are opt-in language epochs that do not break existing code. [secondary]
- Rust 2026 project goals (rust-lang/goals) covered the 2025H1 planning cycle; the Rust-for-Linux goal tracked compiler-version policy and kernel integration work. [secondary]
  - Source: https://github.com/rust-lang/goals/blob/HEAD/src/2025h1/rfl.md

### 3.3 Toolchain: rustup, Cargo, crates.io

- **rustup** is the official toolchain installer managing stable/beta/nightly channels and cross-compilation targets. [secondary]
- **Cargo** is the integrated build system and package manager: dependency resolution, build, test, documentation, and publishing. In 2026 it was widely cited as best-in-class developer experience among systems languages (highly admired in Stack Overflow survey commentary). [secondary]
- **crates.io** hosted over 200,000 crates in 2026. Key ecosystem pieces: `tokio` (async runtime), `axum` and `actix-web` (web frameworks), `serde` (serialization), `sqlx`/`diesel`/`sea-orm` (databases), `clap` (CLI), `ratatui` (TUI), `embassy` (async embedded), `wasm-bindgen`/`wasm-pack` (WebAssembly). [secondary]
  - Source: https://github.com/dingjiu1989-hue/dingjiu1989-hue.github.io/blob/HEAD/md/en/compare/rust-go-zig-comparison.md
- **MSRV** (minimum supported Rust version) is a widely used community convention for communicating the oldest toolchain a crate supports. [secondary]
- Compiler-version policy for Rust-for-Linux pins which rustc versions the kernel supports; this was an active coordination topic in 2025–2026. [secondary]

### 3.4 Linux kernel adoption

- Rust in the Linux kernel was described as **no longer experimental after the December 2025 Kernel Maintainers Summit**, moving toward permanent adoption for kernel code. [secondary]
  - Source: https://www.devclass.com/development/2025/12/15/rust-boosted-by-permanent-adoption-for-linux-kernel-code/1725322
- Android 16 devices on Linux 6.12 reportedly shipped Rust-written **ashmem** support. [secondary]
- Stronger 2026 claims about specific kernel 7.x series details appeared only in weaker articles and should be treated as [unverified].
- In-tree Rust usage remained concentrated in drivers and subsystems rather than core scheduler/MM rewrites. [secondary]

### 3.5 Enterprise and OS adoption (2026 signals)

- **Microsoft**: on **2026-09-18**, reporting described Microsoft elevating Rust to **Tier 1** internal language status alongside C++, C#, and TypeScript, with dedicated toolchain investment, SDL compliance requirements, and deep platform integration. Microsoft already used Rust in the Win32k kernel component and font-parsing code, and collaborated with Google on a Rust-based image decoder for Chromium. [secondary]
  - Source: https://pbxscience.com/microsoft-designates-rust-as-a-tier-1-language-placing-it-alongside-c-c-and-typescript/
- **Google**: Rust in Android (new code memory-safety push), ChromeOS, Fuchsia; Google reported a 52% drop in new memory-safety bugs in Android as Rust replaced C in new code. Treat the 52% figure as a vendor/community-reported number, not a controlled experiment. [secondary]
- **Amazon/AWS**: Firecracker (the microVM monitor behind Lambda and Fargate) and Bottlerocket are Rust; internal Rust build-system investment was driven by team demand rather than top-down mandate. [secondary]
- **Cloudflare**: Pingora (next-generation HTTP proxy replacing nginx for most traffic) and BoringTun (Rust userspace WireGuard, MIT-licensed, 2019; underpins WARP) are Rust. [secondary]
- **Meta**: Buck2 build system and source-control tooling in Rust. [secondary]
- **Others in production**: Discord (real-time messaging backend; Read States service rewritten from Go to Rust for performance), Dropbox (storage backend/sync engine), Fastly Compute@Edge, npm registry backend. [secondary]
- Adoption barriers repeatedly cited in 2026: talent pool and hiring difficulty, learning curve (borrow checker), migration cost in existing codebases, compile times, and ecosystem gaps in some niches. [secondary]

### 3.6 Developer sentiment and jobs

- Rust has been the **most admired** language on the Stack Overflow Developer Survey for **9 years running** (over 80–83% of users wanting to continue), but only ~12–14% of developers reported using it. [secondary]
  - Sources: https://stackoverflow.co/advertising/resources/stack-overflow-developer-survey-for-employer-branding/insight-2/ , https://github.com/memgrafter/anti-alecto/blob/HEAD/digests/2026-04-19_rust-2026-400k-salaries-java-ai-why-it-s-not-everywhere-yet-jon-gjengset-explain_36bde278.md
- Reported 2025–2026 median salary figures (secondary, survey-based, methodology varies): Rust ~$115k, Zig ~$105k, Go ~$95k. Do not treat as controlled compensation data. [secondary]
- Rust sat around **1% / #10** on TIOBE September 2026 (1.34%), trailing far behind Java, Go, TypeScript, and Python in usage and job-posting volume per 2026 commentary — the "most loved but not most used" gap. [independent]
- Hard job-board counts for 2026 were not corroborated during this research; the jobs section below notes this as a gap. [unverified]
- A May 2026 Medium claim that "nearly half of all companies now use Rust in production" is a weak single-source survey claim; treat as [unverified].

### 3.7 Rust limitations and honest trade-offs

- Steep learning curve: ownership/borrowing is unlike anything in GC languages; "fighting the borrow checker" is the standard rite of passage. [secondary]
- Compile times remain a pain point on large codebases versus Go. [secondary]
- `unsafe` is unavoidable at FFI and low-level boundaries; safety depends on correct encapsulation and auditing. [secondary]
- Async Rust ergonomics (Pin, lifetimes in async, executor fragmentation) improved but remained a documented rough edge. [secondary]

---

