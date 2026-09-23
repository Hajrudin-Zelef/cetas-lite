---
id: etape8-phaseb-web-mobile-languages/00-web-mobile-languages/24-conflict-and-gap-register
title: "24. Conflict and gap register"
domain: step-8-phase-b-web-mobile-languages-runtimes-and-toolchains
role: deep-dive
task: reference
actors: ["Anthropic", "Microsoft", "OpenAI"]
dates: ["2025-09", "2025-09-15", "2025-09-16", "2025-11-11", "2025-11-12", "2026-03", "2026-06", "2026-10"]
keywords: ["acquisition"]
source: docs/RAG/etape8_phaseB_web_mobile_languages.md
source_anchor: ""
source_lines: [569, 614]
section: "Step 8 — Phase B — Web & Mobile Languages, Runtimes and Toolchains"
sha256: 58fbfece9cccd5ff81b71f7c6aa257dcacafd8bf969eded6dfd12105b743133e
---

# 24. Conflict and gap register

## 24. Conflict and gap register

Each row's Status column carries the provenance assessment for that item (what was verified, what conflicts, what is missing).

| # | Item | Status |
|---|---|---|
| C1 | TypeScript 6.0 shipped March 2026 / TS 7 Go rewrite in preview | Claimed by secondary sources only; not verified against Microsoft release pages — do not assert as fact |
| C2 | Node 26 Active LTS promotion exact date (October 2026) | Mirrored schedules disagree on the day; only month asserted |
| C3 | Java 25 GA date: 2025-09-15 vs 2025-09-16 | Sources disagree; only "mid-September 2025" asserted |
| C4 | .NET 10 GA date: 2025-11-11 vs 2025-11-12 | Sources disagree; only month asserted |
| C5 | npm registry size: >2.5M vs >5M vs ~1.8M | Different counting methods; no single figure asserted |
| C6 | Set methods: ES2024 vs ES2025 | Sources file them under both editions; boundary unverified |
| C7 | npm version bundled with Node 22 (v10 vs v11) | Secondary source says v10; conflicts with broader expectation of v11 |
| C8 | Kotlin 2.4 (June 2026) | Claimed in search results; not verified against JetBrains notes |
| C9 | Bun acquisition by Anthropic + Rust rewrite | Single satirical-register source; treat as rumor |
| C10 | OpenAI acquisition of Astral (early 2026) | AI-news aggregator only; no official confirmation observed |
| C11 | Stack Overflow 2026 Python 38% vs JS 36% | Secondary write-up; methodology differs from prior years; self-selected sample |
| C12 | Octoverse 2024 vs 2025 #1 (Python vs TypeScript) | Methodology break (repos vs contributors); not a pure adoption swing |
| G1 | Octoverse 2026 report | Not yet published at cutoff |
| G2 | GraalVM/Maven/Gradle/Spring 2026 versions | Not verified in this session |
| G3 | PyPI exact current counts | Not verified against official stats endpoints |
| G4 | Swift 7 / next-major plans | Not covered in this session |
| G5 | Yarn current major version | Not verified in this session |
| G6 | Python 3.15 (expected Oct 2026) | Not released by cutoff; features not asserted |
| G7 | Java 26/27 specifics | Only a sidebar mention of Java 27; not verified |

---

## 25. Glossary

- **TC39**: the Ecma technical committee that standardizes ECMAScript; features advance through stages 0–4, with Stage 4 meaning accepted for the next edition [secondary].
- **LTS**: long-term support release line with extended maintenance [secondary].
- **Free threading**: CPython builds without the GIL (global interpreter lock), enabling true multi-threaded parallelism; opt-in in 3.13/3.14 [secondary].
- **JIT**: just-in-time compiler (YJIT, ZJIT, JVM JIT, .NET JIT) [secondary].
- **AOT**: ahead-of-time compilation (native-image, .NET NativeAOT) [secondary].
- **JSR**: the JavaScript Registry, Deno's first-class package registry alongside npm [secondary].
- **KMP**: Kotlin Multiplatform — shared business logic across platforms [secondary].
- **Impeller**: Flutter's rendering backend [secondary].
- **Ractor**: Ruby's parallel-execution (actor-like) mechanism [secondary].
- **ZJIT**: Ruby 4.0's experimental method-based JIT written in Rust [secondary].
- **Ruby Box**: Ruby 4.0's experimental in-process code isolation [secondary].
- **Phantom dependency**: a package accessible at install time only because of hoisting, though not declared in `package.json` [secondary].
- **Shai-Hulud**: self-replicating npm supply-chain worm (2025) [secondary].

---

