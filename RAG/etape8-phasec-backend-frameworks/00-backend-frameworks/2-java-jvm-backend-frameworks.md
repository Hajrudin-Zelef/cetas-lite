---
id: etape8-phasec-backend-frameworks/00-backend-frameworks/2-java-jvm-backend-frameworks
title: "2. Java / JVM backend frameworks"
domain: step-8-phase-c-backend-frameworks-apis
role: deep-dive
task: reference
actors: []
dates: ["2025-11-11", "2025-11-20", "2026-01-12", "2026-03-25", "2026-05-13", "2026-06-10", "2026-07-27", "2026-08-21", "2026-09", "2026-09-08", "2026-09-09", "2026-09-11", "2026-09-17", "2026-09-21", "2026-09-22", "2026-09-24", "2026-11", "2026-11-10", "2026-11-14", "2027-03-25", "2027-09", "2028-11-14"]
keywords: ["agent", "apache", "cost", "governance", "mcp", "memory"]
source: docs/RAG/etape8_phaseC_backend_frameworks.md
source_anchor: ""
source_lines: [76, 138]
section: "Step 8 — Phase C: Backend Frameworks & APIs"
sha256: 0e18aba00feb78f6a3fed508b65b9b9b6ca0fffb68f4c9ef4cc321e073f737f2
---

# 2. Java / JVM backend frameworks

## 2. Java / JVM backend frameworks

### 2.1 Spring Boot — scope asked for 3.4/3.5, current is 4.x

- The requested scope (Spring Boot 3.4/3.5) is now one major behind current: Spring Boot 4.0 shipped 2025-11-20 on Spring Framework 7, and Spring Boot 4.1 followed 2026-06-10; Spring Boot 4.2 is targeted for November 2026 `[secondary]`:
  - source: `http://dev.to/jamilxt/quarkus-4-is-coming-for-spring-boot-4-here-is-how-the-two-big-java-frameworks-actually-compare-in-8p7`
- Latest observed patch: Spring Boot 4.1.1 released 2026-08-21 `[secondary]`:
  - source: `http://www.codejava.net/frameworks/spring-boot/spring-boot-version-history`
- Spring Boot 4.1 feature highlights (from the same secondary comparison, quoting release documentation): gRPC auto-configuration, HTTP client SSRF mitigation, Kotlin 2.3 support, lazy datasource connections, async context propagation for `@Async` methods `[secondary]`.
- Breaking-change clusters in 4.0: Jakarta EE 11 baseline and the Jackson 3 move — the standard major-boundary migration tax `[secondary]`.
- For the 3.4/3.5 scope: these lines remain the migration baseline for shops not yet on 4.x; 3.x maintenance state should be verified against the Spring support policy before committing to new work on it `[unverified — 3.x EOL dates not collected in this wave; see gaps]`.
- Release cadence: roughly two releases per year (May/November) — easier to plan around than Quarkus's monthly train `[secondary]`.
- Spring Boot EOL reference: `https://www.herodevs.com/blog-posts/spring-boot-versions-eol-dates-and-latest-releases-april-2026` `[secondary]`.

### 2.2 Quarkus

- Quarkus 3.39.4 (released 2026-09-17) is the current supported release `[independent]`:
  - sources: `http://eosl.date/eol/product/quarkus-framework/`, `https://github.com/endoflife-date/endoflife.date/blob/HEAD/products/quarkus-framework.md`
- Quarkus 3.33 is the current LTS (released 2026-03-25, EOL 2027-03-25); Quarkus 3.27 LTS reaches end of support 2026-09-24 — i.e. days after this cutoff, plan upgrades now `[independent]`.
- Forward roadmap: Quarkus 3.40 LTS planned September 2026 as the last 3.x feature release (community maintenance to September 2027); Quarkus 4.0 Beta 1 in September 2026, GA targeted November 2026 `[secondary]`:
  - source: `http://dev.to/jamilxt/quarkus-4-is-coming-for-spring-boot-4-here-is-how-the-two-big-java-frameworks-actually-compare-in-8p7`
- Monthly point-release cadence; experimental features land publicly first (e.g. Signal event bus, Project Leyden AOT support landing in 3.x before Quarkus 4) `[secondary]`.
- Quarkus 3.26 milestone notes (official): Hibernate ORM 7.1, Hibernate Search 8.1, Hibernate Reactive 3.1; named persistence units/data sources in Hibernate Reactive; offline startup and dialect configuration for Hibernate ORM; Dev UI exposed as MCP functions (AI-agent integration); OIDC client filter token-refresh on 401; JFR runtime-data capture; Gradle bumped to 9.0; legacy config classes support dropped entirely `[official]`:
  - source: `https://quarkus.io/blog/quarkus-3-26-released/`
- Security posture (Red Hat build stream): 3.27.5 LTS maintenance release fixes Vert.x redirect/cookie CVEs, Jackson-databind `@JsonIgnore`/`@JsonView` bypasses, Jansi buffer overflow, LangChain4j SQL injection, and a Netty 4.1.136.Final upgrade covering a dozen HTTP/HTTP2 memory-exhaustion and neutralization CVEs `[official-vendor]`:
  - source: `https://quarkus.io/blog/quarkus-3-27-5-released/`
- Core differentiator remains build-time AOT processing + GraalVM native-image first-class support: sub-50ms startup and low RSS for serverless/container density, at the cost of longer builds and reflection-registration discipline.

### 2.3 Micronaut

- Micronaut Framework 5.2.3 (released 2026-09-21) is the current supported release — one day before cutoff `[independent]`:
  - sources: `http://eosl.date/eol/product/micronaut/`, `https://endoflife.date/micronaut` (page updated 2026-09-22)
- Micronaut 5.0 released 2026-05-13 with a Java 25 baseline, Kotlin 2.3 support and Apache Groovy 5 support `[official]`:
  - sources: `https://micronaut.io/category/release-announcements/`, `https://micronaut.io/blog/`
- Micronaut 5.1.0 (2026-07-27) added Open DI support plus updates across Core, Data, AI, cloud, messaging and APIs `[official]`.
- Patch train is active: 5.1.1 → 5.1.5 within August–September 2026, mostly Netty/security driven — 5.1.5 (2026-09-11) pulls Netty 4.2.18.Final with several security fixes; 5.1.4 (2026-09-09) fixed WebSocket fragmentation DoS, raw exception-message leakage in default error responses, gzip/deflate decompression-bomb DoS, and SQL injection via unvalidated tenant identifiers and Sort/Pageable property names `[official]`:
  - source: `https://micronaut.io/2026/09/09/micronaut-framework-5-1-4/`
- Older lines still patched: 4.10.28 and 3.10.12 both released 2026-09-11 (active maintenance) `[independent]`.
- Governance note: Micronaut announced plans to join the Commonhaus Foundation (2026-01-12) `[official]`.
- Technical identity: compile-time dependency injection (no runtime reflection), AOT-friendly, first-class GraalVM native image, Micronaut Data (compile-time repository generation, incl. R2DBC reactive), low memory footprint — the "Spring-compatible concepts without the runtime cost" pitch.
- Java 25 baseline is the headline 2026 bet: virtual threads, structured concurrency and scoped values become first-class framework citizens rather than bolt-ons `[official]`.

---

## 3. .NET — ASP.NET Core 9/10 and Minimal APIs

- .NET 10.0.12 (patch, 2026-09-08) is the current LTS patch; .NET 10 LTS support runs to 2026-11-14... correction: LTS support ends 2028-11-14 `[official]`:
  - source: `https://dotnet.microsoft.com/en-us/download/dotnet?initial-os=linux`
- Support matrix as of cutoff `[official]`:
  - .NET 10: LTS, GA 2025-11-11, latest 10.0.12, supported to 2028-11-14.
  - .NET 9: STS, latest 9.0.20 (2026-09-08), support ends 2026-11-10 — weeks after cutoff; migrate to 10 now.
  - .NET 8: LTS, support ends 2026-11-10 — same migration urgency.
  - .NET 11: release candidate in September 2026 (STS line forming).
  - source: `https://dotnet.microsoft.com/platform/support/policy/dotnet-core/`
- Language: C# 14 ships with .NET 10 `[official]`.
- ASP.NET Core 10 web-stack highlights (secondary summary of release notes): validation and JSON Patch support for Minimal APIs, server-sent events (SSE) support, and OpenAPI 3.1 document generation `[secondary]`.
- Minimal APIs remain the recommended shape for new JSON APIs: `MapGet`/`MapPost` route handlers with typed results, built-in validation, `TypedResults`, OpenAPI annotations via `WithOpenApi()` — the ceremony-free alternative to controllers that closed most of the gap with FastAPI/Hono-style developer experience.
- EF Core 10 ships with .NET 10 (see 12.4): complex types, first-class LeftJoin/RightJoin in LINQ, JSON column improvements, vector search — the ORM keeps pace with the API layer.
- Performance: ASP.NET Core (Kestrel) is perennially at the top of the TechEmpower composite/plaintext rankings among mainstream full frameworks; Native AOT compilation is production-viable for APIs in .NET 8+ and improves cold start and container density.
- Migration guidance: .NET 8 → 10 is the LTS-to-LTS jump teams should execute before 2026-11-10; .NET 9 STS users have no choice but 10 (or the 11 STS line).

---

