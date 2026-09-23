---
id: etape8-phaseb-web-mobile-languages/00-web-mobile-languages/22-lifecycle-support-license-comparison-tables
title: "22. Lifecycle / support / license comparison tables"
domain: step-8-phase-b-web-mobile-languages-runtimes-and-toolchains
role: deep-dive
task: licenses
actors: ["Apple"]
dates: ["2025-10-07", "2025-10-16", "2025-11-20", "2025-12-25", "2026-07-23", "2026-09-10", "2026-09-11", "2026-09-22", "2026-10", "2027-04-30", "2028-04-30"]
keywords: ["license", "acquisition", "alignment", "apache", "benchmark", "inference", "licenses"]
source: docs/RAG/etape8_phaseB_web_mobile_languages.md
source_anchor: ""
source_lines: [498, 568]
section: "Step 8 — Phase B — Web & Mobile Languages, Runtimes and Toolchains"
sha256: edd5f4855d5245da9dedad146a3bfd121bbefebc1500dbc6134706ceee9f3040
---

# 22. Lifecycle / support / license comparison tables

## 22. Lifecycle / support / license comparison tables

### 22.1 Runtime release lines at cutoff

| Runtime | Current line (2026-09-22) | LTS line | Support/EOL note | Tag |
|---|---|---|---|---|
| Node.js | 26 (Current) | 24 "Krypton" (Active) | 22 Maintenance LTS → EOL 2027-04-30; 24 → EOL 2028-04-30; 26 → Active LTS Oct 2026 | [secondary] |
| Deno | 2.x | n/a (rolling) | No LTS scheme; follows semver with 2.x major line | [secondary] |
| Bun | 1.3.x | n/a (rolling) | 1.3.14 described as final Zig release; Rust-rewrite reports unverified | [secondary] |
| Python | 3.14 | n/a (per-version bugfix windows) | 3.14 released 2025-10-07; 3.13 experimental free-threading/JIT | [secondary] |
| Java | 25 (LTS) | 25 | 6-month cadence; 21 and 25 are the active LTS lines | [secondary] |
| .NET | 10 (LTS) | 10 | .NET 9 STS through Nov 2026; .NET 10 LTS through Nov 2028 | [secondary] |
| Ruby | 4.0 | n/a (per-version maintenance) | 4.0.0 released 2025-12-25; 3.x line in maintenance | [secondary] |
| PHP | 8.5 | n/a (per-version: 2y active + 1y security) | 8.5 released 2025-11-20 | [secondary] |
| Lua | 5.5 | n/a (per-release) | 5.5.1 on 2026-07-23; 5.4.9 maintenance 2026-09-10 | [secondary] |
| Elixir | 1.19 | n/a (per-release) | 1.19 released 2025-10-16 | [secondary] |
| Flutter/Dart | Flutter 3.47 / Dart 3.13 | n/a (rolling stable) | 3.47.4 / 3.13.3 observed 2026-09-11; 3.50 targeted Nov 2026 | [secondary] |
| Swift | 6.2 | n/a (Xcode-bundled) | "Approachable concurrency" release line | [secondary] |
| Kotlin | 2.3 (2.4 claimed) | n/a (rolling) | 2.2 and 2.3 verified via secondary refs; 2.4 unverified | [secondary] |

### 22.2 Licenses (as commonly reported; verify per redistribution use)

| Language/runtime | License | Tag |
|---|---|---|
| Node.js | MIT | [secondary] |
| Deno | MIT | [secondary] |
| Bun | MIT (core) / proprietary for some components — verify before redistribution | [unverified] |
| Python (CPython) | PSF License (permissive) | [secondary] |
| Java (OpenJDK) | GPLv2 with Classpath Exception | [secondary] |
| .NET / C# | MIT (runtime, Roslyn) | [secondary] |
| Kotlin | Apache 2.0 | [secondary] |
| Swift | Apache 2.0 | [secondary] |
| Dart / Flutter | BSD-style (BSD 3-clause) | [secondary] |
| PHP | PHP License (permissive, non-copyleft) | [secondary] |
| Ruby | Ruby License / BSD dual | [secondary] |
| Elixir | Apache 2.0 | [secondary] |
| Lua | MIT | [secondary] |
| TypeScript | Apache 2.0 | [secondary] |

---

## 23. Decision guides

### 23.1 Choosing a JavaScript runtime (2026)

- **Node.js 24 LTS**: default choice for production services; Active LTS, predictable EOL, largest ecosystem compatibility [secondary].
- **Node.js 26**: for teams that want newest V8/features and accept the Current line until October 2026 LTS promotion [secondary].
- **Deno 2.x**: for security-first services (permissions model), or teams wanting built-in TypeScript/lint/format/test without a toolchain [secondary].
- **Bun 1.3.x**: for speed-sensitive builds/tests and all-in-one DX; weigh the unverified acquisition/rewrite reports before betting a multi-year platform on it [secondary].

### 23.2 Choosing a JS package manager (2026)

- **pnpm**: default recommendation for monorepos and CI speed; strict isolation prevents phantom dependencies; v10+ security defaults (no lifecycle scripts, `minimumReleaseAge`) directly address the Shai-Hulud class of attacks [secondary].
- **npm**: zero-config default bundled with Node; fine for small projects; enforce `package-lock.json` and audit in CI [secondary].
- **Yarn**: viable alternative with strong workspace support; pin via `packageManager` + Corepack whichever you choose [secondary].
- Never mix two package managers in one repo [secondary].

### 23.3 Choosing a Python stack (2026)

- **New projects**: uv (package/project/Python-version management) + Ruff (lint/format) + a type checker (ty or mypy); `uv.lock` for reproducibility [secondary].
- **AI/ML workloads**: CPython + PyTorch/TensorFlow/JAX; watch free-threaded builds (3.13/3.14 opt-in) for CPU-parallel inference/data pipelines, but expect single-thread regressions and extension-module incompatibilities — benchmark your own workload [secondary].
- **Existing Poetry/PDM/Hatch projects**: no forced migration; Poetry 2.0's PEP 621 alignment keeps it standards-compatible [secondary].

### 23.4 Mobile/cross-platform (2026)

- **Flutter 3.47 / Dart 3.13**: single-codebase iOS+Android+web+desktop; Impeller maturing; check minimum-OS bumps before upgrading [secondary].
- **Kotlin Multiplatform**: share logic, keep native UIs; Swift export improving for iOS interop [secondary].
- **Native**: Swift 6.2 (iOS) with approachable concurrency; Kotlin 2.3 (Android) targeting Java 25 bytecode [secondary].

---

