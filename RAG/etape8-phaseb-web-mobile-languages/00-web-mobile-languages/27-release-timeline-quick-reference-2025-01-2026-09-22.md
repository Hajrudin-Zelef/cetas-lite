---
id: etape8-phaseb-web-mobile-languages/00-web-mobile-languages/27-release-timeline-quick-reference-2025-01-2026-09-22
title: "27. Release timeline quick reference (2025-01 → 2026-09-22)"
domain: step-8-phase-b-web-mobile-languages-runtimes-and-toolchains
role: deep-dive
task: reference
actors: ["OpenAI"]
dates: ["2025-03", "2025-03-18", "2025-05-06", "2025-09-15", "2025-10-07", "2025-10-16", "2025-10-31", "2025-11-11", "2025-11-17", "2025-11-20", "2025-12-18", "2025-12-23", "2025-12-25", "2025-12-27", "2026-01-20", "2026-04-08", "2026-07-23", "2026-09", "2026-09-10", "2026-09-11", "2026-09-16", "2026-09-22"]
keywords: ["embedding", "memory"]
source: docs/RAG/etape8_phaseB_web_mobile_languages.md
source_anchor: ""
source_lines: [703, 762]
section: "Step 8 — Phase B — Web & Mobile Languages, Runtimes and Toolchains"
sha256: 1742104d4c7ebc2893e18d50555b971f7800a48c4e837e0d9fba45f98e9b3e58
---

# 27. Release timeline quick reference (2025-01 → 2026-09-22)

## 27. Release timeline quick reference (2025-01 → 2026-09-22)

| Date | Event | Tag |
|---|---|---|
| 2025-01 | Poetry 2.0 adopts PEP 621 `[project]` table | [secondary] |
| 2025-03-18 | Java 24 released (non-LTS) | [secondary] |
| 2025-05-06 | Node.js 24 "Krypton" initial release (→ Active LTS) | [secondary] |
| 2025-06 | ECMAScript 2025 ratified by TC39 | [secondary] |
| 2025-08 | TypeScript 5.9 (`import defer`, `--module node20`) | [secondary] |
| 2025-09-15/16 | Java 25 LTS (18 JEPs); exact day disputed | [secondary] |
| 2025-10-07 | Python 3.14 released | [secondary] |
| 2025-10-16 | Elixir 1.19 released | [secondary] |
| 2025-10-31 | GitHub Octoverse 2025 published (TypeScript No. 1 contributors) | [secondary] |
| 2025-11 | Shai-Hulud aggressive variant (preinstall hook) documented | [secondary] |
| 2025-11-11/12 | .NET 10 LTS GA; exact day disputed | [secondary] |
| 2025-11-17 | Ruby 4.0.0-preview2 (build tag observed) | [secondary] |
| 2025-11-20 | PHP 8.5 released (pipe operator, URI extension) | [secondary] |
| 2025-12-18 | Ruby 4.0.0-preview3 | [secondary] |
| 2025-12-23 | lua-language-server 3.16.3 upgrades to Lua 5.5 | [secondary] |
| 2025-12-25 | Ruby 4.0.0 released (ZJIT experimental, Ruby Box experimental) | [secondary] |
| 2025-12-27 | Lua 5.5.0 (per package metadata) | [secondary] |
| 2026-01-20 | lua-language-server 3.17.1 | [secondary] |
| 2026-early | OpenAI acquires Astral (report; official confirmation not observed) | [secondary] |
| 2026-02 | Flutter 3.41 window; Swift 6.2 "approachable concurrency" coverage | [secondary] |
| 2026-03 | TypeScript 6.0 claimed (unverified); TS 5.8 released March 2025 | [unverified] |
| 2026-04-08 | lua-language-server 3.18.1 (Lua 5.5 syntax support) | [secondary] |
| 2026-05 | Flutter 3.44 window | [secondary] |
| 2026-06 | Kotlin 2.4 claimed in search results (unverified) | [unverified] |
| 2026-07 | Rust enters TIOBE top 10 (first run) | [secondary] |
| 2026-07-23 | Lua 5.5.1 | [secondary] |
| 2026-08 | Flutter 3.47 window; TIOBE top-10 unchanged Aug→Sep | [secondary] |
| 2026-09-10 | Lua 5.4.9 (maintenance) | [secondary] |
| 2026-09-11 | Flutter 3.47.4 / Dart 3.13.3 observed | [secondary] |
| 2026-09-16 | TIOBE September 2026 published (Python 17.76%, Rust #10) | [secondary] |
| 2026-10 | Node.js 26 → Active LTS planned; Octoverse 2026 expected (not yet published at cutoff) | [secondary] |
| 2026-11 | .NET 9 STS support ends; .NET 10 LTS continues to Nov 2028; Flutter 3.50 targeted | [secondary] |

## 28. Cross-cutting 2026 themes (interpretive synthesis)

- **AI reshapes every ranking**: Python's No. 1 positions (TIOBE, Octoverse 2024, SO 2026 usage) are all attributed by sources to AI/ML demand; 1.1M+ public repos imported an LLM SDK by Aug 2025 (+178% YoY) [secondary].
- **Typed systems gain from AI codegen**: the Octoverse 2025 narrative (94% of LLM-generated compile errors being type-check failures per a 2025 academic study) favors TypeScript, C#, and other typed languages in AI-assisted workflows [secondary].
- **Supply-chain security is now a first-class feature**: pnpm v10+ disabling lifecycle scripts by default and adding `minimumReleaseAge` is a direct response to the Shai-Hulud class of attacks; npm's 2FA mandates and PyPI's token guidance are the same trend [secondary].
- **Rust as infrastructure language**: Rust 1.85+ required to build Ruby 4.0's ZJIT; uv and Ruff are Rust binaries; Rust entered the TIOBE top 10 — the language is becoming the default choice for new developer tooling [secondary].
- **Concurrency and parallelism everywhere**: Python free-threading (PEP 779), Ruby Ractors + Ruby Box, Swift approachable concurrency, Java scoped values — every major runtime shipped a 2025–2026 story about using more cores safely [secondary].
- **Runtime convergence on "batteries included"**: Deno (lint/fmt/test), Bun (SQL clients, bundler), uv (version manager + packager), Node (`--experimental-strip-types`, `--env-file`, `--watch`) all reduce the third-party toolchain surface [secondary].
- **Embedding and scripting stay relevant**: Lua 5.5's GC and memory work, Ruby's isolation primitives, and Elixir's compilation speedups show mature languages competing on runtime quality rather than syntax [secondary].
- **Web platform keeps absorbing the toolchain**: ES2025's import attributes, JSON modules, and Iterator helpers move work from bundlers/userland into the language [secondary].

## 29. What changed in this phase vs. prior knowledge (delta log)

- Ruby 3.5 never shipped as 3.5 — the line was renamed to Ruby 4.0 (released 2025-12-25); update any index that lists "Ruby 3.5" as a release [secondary].
- Lua 5.5 is a real released line (Dec 2025), not a hypothetical — replaces assumptions that 5.4.x was latest [secondary].
- Java 25 is the active LTS (not 21 alone); Java 24 was a short-term release [secondary].
- .NET 10 is the active LTS (Nov 2028); .NET 9 STS ends Nov 2026 [secondary].
- Node 26 is the Current line heading to LTS in Oct 2026; Node 24 is Active LTS [secondary].
- TypeScript 6.0/Go-rewrite claims are unverified — prior drafts must not assert them [unverified].
- pnpm v10+/v11 security defaults (no lifecycle scripts, `minimumReleaseAge`) are new since 2024-era comparisons [secondary].
- Flutter 3.47/Dart 3.13 are the observed stable pair at cutoff [secondary].

*End of file.*
