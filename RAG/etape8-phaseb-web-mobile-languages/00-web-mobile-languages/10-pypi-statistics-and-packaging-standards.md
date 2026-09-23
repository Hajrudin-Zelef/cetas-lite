---
id: etape8-phaseb-web-mobile-languages/00-web-mobile-languages/10-pypi-statistics-and-packaging-standards
title: "10. PyPI statistics and packaging standards"
domain: step-8-phase-b-web-mobile-languages-runtimes-and-toolchains
role: deep-dive
task: reference
actors: ["Oracle"]
dates: ["2025-03-18", "2025-09", "2025-09-15", "2025-09-16", "2025-11-11", "2025-11-12", "2025-11-13", "2026-03", "2026-09", "2026-11", "2028-11"]
keywords: ["packaging", "memory"]
source: docs/RAG/etape8_phaseB_web_mobile_languages.md
source_anchor: ""
source_lines: [237, 297]
section: "Step 8 — Phase B — Web & Mobile Languages, Runtimes and Toolchains"
sha256: a444e0a114a43927b576824f4acee0910410e82843b929d5ae7adfa5284e5cb5
---

# 10. PyPI statistics and packaging standards

## 10. PyPI statistics and packaging standards

- PyPI was reported at "over 500,000 packages" with thousands of uploads daily in a supply-chain analysis [secondary]: https://forum.gnoppix.org/t/the-next-wave-of-supply-chain-attacks-npm-pypi-and-docker-hub-incidents-set-the-stage-for-2026/2983.
- ⚠️ Exact current project/file/download counts were not verified against PyPI's official statistics endpoints (e.g. `pypistats`) in this session — treat any single PyPI size figure as `[unverified]`.
- PyPA does not endorse a specific project manager; the standardized configuration format is `pyproject.toml` with PEP 621 `[project]` metadata [secondary].
- uv's ~13% share of PyPI downloads is a notable ecosystem signal that the download API reflects client diversity, not just pip [secondary].

---

## 11. Java — 24 / 25 LTS and the JEP landscape

### 11.1 Release status at cutoff

- Java 24: released 2025-03-18 (non-LTS, 6-month cadence release) [secondary].
- Java 25: LTS release, September 2025 — sources vary between 2025-09-15 and 2025-09-16, so only "mid-September 2025" is asserted here [secondary]. ⚠️ Date conflict preserved.
- Java 25 ships 18 JEPs according to several secondary sources [secondary].
- Java 26 (March 2026) and Java 27 (September 2026) exist on the 6-month cadence; a Linux news sidebar mentioned "Java 27 Released with G1 as the Default Garbage Collector" — noted as `[unverified]` since it was not verified against an official release page [unverified].

### 11.2 Java 25 notable items (preview vs. final matters)

- Compact source files and instance `main` methods (simplified entry points) [secondary].
- Scoped values (structured-concurrency-adjacent state sharing) [secondary].
- Flexible constructor bodies [secondary].
- Compact object headers (memory-layout work) [secondary].
- Ahead-of-time (AOT) profiling support [secondary].
- JFR method timing and tracing [secondary].
- Generational Shenandoah GC [secondary].
- ⚠️ Several listed features remain preview/incubator in Java 25 — do not describe preview features as final language features [secondary].
- Sources: https://dev.to/gaston_herrlein_3baa4d62e/from-java-8-to-java-25-the-language-you-think-you-know-no-longer-exists-4ocg, https://github.com/anupam0401/evolvdb/blob/HEAD/docs/learning/phase-14-java25-migration.md, https://www.hifitoolkit.com/tech-news/whats-new-in-java-25-features-lts/ [secondary].

---

## 12. JVM tooling — GraalVM, Maven, Gradle, Spring

- GraalVM (Oracle): native-image and polyglot JIT for the JVM; specific 2026 GraalVM release versions tied to JDK 24/25 were not verified in this session — treat current GraalVM numbering as `[unverified]` [unverified].
- Maven and Gradle remain the two dominant JVM build tools; Gradle is the default for Android and Kotlin-heavy builds, Maven for traditional enterprise Java [secondary].
- Spring Framework's ecosystem position is tied to Java LTS adoption (17 → 21 → 25); exact Spring release-to-JDK mapping was not verified in this session [unverified].
- ⚠️ Gap: no primary-source verification of 2026 GraalVM/Maven/Gradle/Spring versions was performed; downstream RAG steps should not cite version numbers from this section [unverified].

---

## 13. C# / .NET — 9 / 10 and C# 14

### 13.1 .NET lifecycle

- .NET 9: standard-term support (STS), supported through November 2026 [secondary].
- .NET 10: LTS release, supported through November 2028; GA reported on 2025-11-11 or 2025-11-12 depending on the source [secondary]. ⚠️ Day conflict preserved; month/year are consistent.
- Sources: https://github.com/funkysi1701/funkysi1701.github.io/blob/HEAD/content/posts/2026/dotnet-5-to-10-features.md, https://github.com/tokawa-ms/dailyazureupdatesgenerator/blob/HEAD/updates_en/azure-updates-2025-11-13.md [secondary].

### 13.2 C# 14 features (with .NET 10)

- Field-backed (field) properties: auto-properties with access to the compiler-generated backing field [secondary].
- Extension blocks / extension members: new syntax surface for extension declarations [secondary].
- Null-conditional assignment [secondary].
- Unbound generic `nameof` [secondary].
- Improved lambda and `ref` syntax [secondary].
- .NET 10 runtime topics: JIT and AOT compilation work, JSON and WebSocket APIs, SIMD, and security hardening [secondary].
- Source: https://dev.to/iron-software/whats-new-in-net-10-and-c-14-109n [secondary].

---

