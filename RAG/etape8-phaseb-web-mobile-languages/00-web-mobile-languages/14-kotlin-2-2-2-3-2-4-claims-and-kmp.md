---
id: etape8-phaseb-web-mobile-languages/00-web-mobile-languages/14-kotlin-2-2-2-3-2-4-claims-and-kmp
title: "14. Kotlin — 2.2 / 2.3 / 2.4 claims and KMP"
domain: step-8-phase-b-web-mobile-languages-runtimes-and-toolchains
role: deep-dive
task: reference
actors: ["Apple"]
dates: ["2026-06", "2026-09-11"]
keywords: ["research"]
source: docs/RAG/etape8_phaseB_web_mobile_languages.md
source_anchor: ""
source_lines: [298, 357]
section: "Step 8 — Phase B — Web & Mobile Languages, Runtimes and Toolchains"
sha256: 8fb638c29d41883b728bc1a493545bca476873f762b108c56427ab880ed58971
---

# 14. Kotlin — 2.2 / 2.3 / 2.4 claims and KMP

## 14. Kotlin — 2.2 / 2.3 / 2.4 claims and KMP

### 14.1 Kotlin 2.2

- K2 compiler became the default for kapt [secondary].
- Experimental FIR incremental compilation [secondary].
- Swift export (Kotlin/Native → Swift interop) advanced [secondary].
- Kotlin/JS BigInt options [secondary].
- Kotlin/Native hardening [secondary].
- Shared `webMain` source set for multiplatform web targets [secondary].
- Source: https://github.com/heapy/kortex/blob/HEAD/plugins/kortex/skills/modern-kotlin/references/kotlin-2.2.md [secondary].

### 14.2 Kotlin 2.3

- Support for Java 25 bytecode as a compilation target [secondary].
- Language-version floors raised (older language versions deprecated/removed) [secondary].
- Ant build support removed [secondary].
- Swift export progress continued; Kotlin/JS and Kotlin/Wasm target work; AGP 9 migration [secondary].
- Source: https://github.com/heapy/kortex/blob/HEAD/plugins/kortex/skills/modern-kotlin/references/kotlin-2.3.md [secondary].

### 14.3 Kotlin 2.4 claim

- Search results claimed a Kotlin 2.4 release in June 2026 [secondary]; this was NOT verified against JetBrains' official release notes in this session — treat Kotlin 2.4 details as `[unverified]` [unverified].

### 14.4 Kotlin Multiplatform (KMP) context

- KMP shares business logic across Android, iOS, desktop, and web while keeping native UIs; the 2.2/2.3 releases incrementally improved iOS (Swift export) and web (`webMain`, Wasm) targets [secondary].
- Source: https://github.com/gdg-nantes/devfestnantesmobile/blob/HEAD/.planning/research/STACK.md [secondary].

---

## 15. Swift — 6.2 "approachable concurrency"

- Swift 6.2 is themed "approachable concurrency" [secondary].
- Optional default main-actor isolation: types and functions can opt into main-actor isolation by default rather than annotating everything [secondary].
- Caller-context async behavior: async functions inherit more caller context, reducing annotation burden [secondary].
- Explicit `@concurrent` attribute for opting out into concurrent execution [secondary].
- Broader Swift 6-era context: strict concurrency and data-race checking, `Sendable`, macros, parameter pack iteration, tuple conformances, typed throws [secondary].
- Sources: https://mjtsai.com/blog/2025/11/03/swift-6-2-approachable-concurrency/, https://github.com/durellwilson/swift-2026-course/blob/HEAD/book/src/research-2025-2026.md, https://medium.com/@dhavaljasoliya8/swift-6-2-released-whats-new-and-exciting-for-ios-developers-2a73815d2c37 [secondary].

---

## 16. Dart / Flutter

### 16.1 Release windows (2026)

- The official Flutter website repository lists 2026 release windows: 3.41 (February), 3.44 (May), 3.47 (August), 3.50 targeted for November [secondary].
- A current stack reference reported Flutter 3.47.4 / Dart 3.13.3 on 2026-09-11 [secondary].
- Sources: https://github.com/flutter/website/blob/HEAD/sites/docs/src/content/install/archive.md, https://github.com/artificialorctelligence/orclab/blob/HEAD/skills/stack-flutter/SKILL.md [secondary].

### 16.2 3.47 / Dart 3.13 topics

- Impeller (the rendering backend) expansion [secondary].
- Minimum platform-version changes [secondary].
- Package separation/treeshaking discussions in the ecosystem [secondary].
- ⚠️ Details need careful provenance; only the version numbers and dates above are asserted with confidence [secondary].
- Source: https://github.com/fluttercook/fluttercook.github.io/blob/HEAD/src/content/blog/flutter-introduction-2026.md [secondary].

---

