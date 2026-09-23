---
id: etape8-phasec-backend-frameworks/00-backend-frameworks/15-gaps-conflicts-and-unverified-claims
title: "15. Gaps, conflicts and unverified claims"
domain: step-8-phase-c-backend-frameworks-apis
role: deep-dive
task: reference
actors: ["Lambda"]
dates: ["2026-03-24", "2026-05", "2026-07", "2026-09", "2026-09-22", "2026-09-24", "2026-11-10", "2026-12-31"]
keywords: ["advisory", "agent", "agents", "benchmarks", "consumer", "inference", "mcp", "memory", "throughput"]
source: docs/RAG/etape8_phaseC_backend_frameworks.md
source_anchor: ""
source_lines: [447, 550]
section: "Step 8 — Phase C: Backend Frameworks & APIs"
sha256: af03982193cde053ad9a188b3a258b3ac2faa0ed6047cd22ef4162d103770f8e
---

# 15. Gaps, conflicts and unverified claims

## 15. Gaps, conflicts and unverified claims

- **TechEmpower Round 24 does not exist** (repo archived 2026-03-24; Round 23/Feb 2025 is last) — any Round 24 numbers encountered elsewhere are fabricated or mislabeled `[secondary]`.
- **NestJS 12**: preparation PRs and September 2026 articles exist, but no official GA release confirmed — do not cite as current `[unverified]`.
- **NestJS latest patch**: 11.1.24 (May 2026) per one secondary source vs 11.1.27/28 seen in project snapshots (June–July 2026) — patch-level conflict, minor `[unverified]`.
- **Phoenix 2026 patch level**: no official 2026 release note captured; 1.8.x asserted from release history `[unverified]`.
- **OpenAPI 3.2**: status not confirmed in collected sources `[unverified]`.
- **Spring Boot 3.4/3.5 EOL dates**: not collected in this wave; verify against the Spring support policy before planning maintenance on 3.x `[gap]`.
- **Jakarta Data**: mentioned as the 2026 standards-track repository spec; adoption status not verified `[unverified]`.
- **Axum exact 2026 patch**: 0.8.x confirmed as line, exact latest patch not pinned `[gap]`.
- **Adoption percentages** (REST 70–93%, GraphQL ~25%, gRPC ~14%, tRPC ~15%): all `[secondary]` syntheses of job-listing/survey data with differing methodologies — use for ranking, not precision.
- **Laravel 13 "AI-native" framing** and **Symfony 25%-of-Packagist math**: vendor/marketing claims `[vendor-reported]` — directionally informative, not audited.
- **FastAPI 0.141.1, Flask 3.1.3, Hono 4.11.7, Prisma 7.9.0/8.0.0-rc.10, Drizzle 0.45.2**: versions observed via secondary mirrors/skill repos, not the projects' own release pages — patch-level details should be re-verified against official tags before procurement use.

---

## 16. Key takeaways for backend selection (2026-09-22)

1. **Runtime consolidation is over; the lines are drawn.** Python = FastAPI/Litestar (async) with Django for monoliths; JVM = Spring Boot 4.x vs Quarkus 3.39 vs Micronaut 5.2 (Java 25 baseline); .NET = ASP.NET Core 10 LTS; Go = Gin/Echo (net/http) vs Fiber v3 (fasthttp); Rust = Axum (momentum) vs Actix-web (throughput); TS = NestJS (enterprise) vs Fastify (perf) vs Hono (edge).
2. **Upgrade pressure is immediate on three fronts:** .NET 8/9 support ends 2026-11-10 (move to .NET 10 LTS); Quarkus 3.27 LTS support ends 2026-09-24; Fastify 4 is EOL and Laravel 11 is out of security support.
3. **The ORM layer is absorbing AI/document workloads:** EF Core 10 (JSON columns, vector search, LeftJoin/RightJoin), Prisma 8 (contract-first, vector extensions), SQLAlchemy 2.0.54 (free-threaded Python) — the "just use Postgres + your ORM" stack keeps getting wider.
4. **API style is settled as hybrid:** REST/OpenAPI 3.1 at the edge, tRPC or GraphQL for frontends, gRPC inside the mesh. Framework choice should follow the consumer, not the hype cycle.
5. **Benchmarks inform architecture, not purchasing:** TechEmpower is archived at Round 23; use it for engine-level signals (fasthttp vs net/http, Kestrel vs uvicorn, native AOT vs JIT), never for "X% faster" claims — and re-check every version number here against official tags before acting on it.

---

---

## 17. 2026 security advisory roundup (backend-relevant)

- **Django 5.x**: Fedora 43 python-django5 advisory stream active through 2026 — keep 5.2.x patched; 5.2.17 is the current patch `[independent]`.
- **Flask**: CVE-2026-27205 fixed in Flask 3.1.3 (Vary: Cookie) — upgrade Pallets trio together (Flask/Werkzeug/Jinja2) `[secondary]`.
- **Rails**: CVE-2026-66066 (Active Storage) fixed in 8.1.3.1 / 8.0.5.1 / 7.2.3.2 — critical, patch immediately if on an affected line `[secondary]`.
- **Fiber**: v3.1.0 fixed static-middleware path traversal (Windows) and CVE-2026-25899 (flash-cookie msgpack DoS, CVSS 7.5); v2 users take 2.52.11 for CVE-2025-66630 (predictable UUIDs, CVSS 9.4) `[independent]`.
- **Fastify**: 5.12.2 fixed four High advisories (header validation bypass, request validation bypass, auth bypass via malformed URLs, body replacement); practitioner minimum is 5.8.5+, and 4.x is EOL `[secondary]`.
- **Quarkus**: 3.27.5 LTS maintenance release — Vert.x redirect/cookie CVEs (CVE-2026-15075/15076), Jackson-databind `@JsonIgnore`/`@JsonView` bypasses (CVE-2026-59888/59889), Netty 4.1.136.Final upgrade covering ~12 HTTP/HTTP2 memory-exhaustion and header-neutralization CVEs `[official-vendor]`.
- **Micronaut**: 5.1.4 fixed WebSocket fragmentation DoS, raw exception-message leakage, gzip/deflate decompression-bomb DoS, and tenant-identifier / Sort-Pageable SQL injection; 5.1.5 pulled Netty 4.2.18.Final `[official]`.
- **actix-http**: 3.12.1 fixed request smuggling; 3.12.0 raised MSRV to Rust 1.88 `[official]`.
- **Hono**: v4.2.7-era fix for restricted directory traversal in `serveStatic` with Deno (referenced in a TanStack Router renovate PR) — keep 4.x patched; 4.11.7 is current `[secondary]`.
- **Prisma**: 7.9.0's AI safety checkpoint (blocking `migrate reset` from the MCP server, guarding `db push --accept-data-loss`) is a supply-chain-adjacent hardening worth noting for agent-driven workflows `[secondary]`.
- **Cross-cutting 2026 themes**: decompression bombs (gzip/deflate), HTTP request smuggling variants, header neutralization, JWT-adjacent issues, and deserialization flaws dominate the advisory stream — validate/sanitize at the edge, keep Netty/Vert.x-class HTTP engines patched, and treat "latest minor" as the only supported target on fast-moving lines (Quarkus monthly, Micronaut 5.1.x, Fastify 5.12.x, Hono 4.11.x).

---

## 18. Major-upgrade migration notes (2025–2026 generation changes)

- **Spring Boot 3.x → 4.x**: Jakarta EE 11 baseline; Jackson 3 migration (behavioral changes in databind); Spring Framework 7; plan for the May/November cadence (4.2 targeted Nov 2026). Testcontainers-based integration suites catch most breakage early `[secondary]`.
- **Fiber v2 → v3**: handler signature change (`*fiber.Ctx` → `fiber.Ctx` value); `Ctx` now implements `context.Context`; sessions become middleware; `app.Static()`/`Filesystem` removed; middleware data via `FromContext()`; stricter prefix matching; default redirect 302→303. Use the CLI migrator (`fiber migrate --to v3`) then run the full test suite — behavior changes (redirect codes, session lifecycle) bite in production, not compile time `[official]`.
- **Echo v4 → v5**: public API changes documented in API_CHANGES_V5.md; v4 security support ends 2026-12-31 — the migration is time-boxed `[official]`.
- **Symfony 7.4 → 8.0**: same features, zero deprecations allowed — the upgrade is "resolve every deprecation notice" then bump; PHP 8.4+ required `[official]`.
- **Laravel 12 → 13**: mostly additive per secondary upgrade notes; breakage concentrates in removed deprecated helpers and minimum-PHP bumps — verify against the official upgrade guide, not blog summaries `[secondary]`.
- **Express 4 → 5**: async error handling without `next(err)` wrappers; stricter routing; removed deprecated APIs. Express 5 has been the default since 2024; remaining 4.x apps are now two years behind `[secondary]`.
- **.NET 8/9 → 10**: LTS-to-LTS (8→10) or STS-to-LTS (9→10) before 2026-11-10; EF Core 10's JSON/complex-type mapping may change generated migrations — review migration diffs; Native AOT requires trimming-safe code (no unannotated reflection) `[official][secondary]`.
- **Prisma 7 → 8**: contract-first rewrite — PSL compiles to versioned JSON contracts, generated client replaced by composable DSL; 8.0.0-rc.10 is pre-release, so production stays on 7.9.0 while validating 8 in staging; agent-skill and MCP-surface changes affect AI-assisted workflows `[secondary]`.
- **Litestar 2.x → 3 (forthcoming)**: DI overhaul previewed in 2.23.0 (`NamedDependency`, `SkipValidation`, type-based DI coming) — adopt the new markers now to reduce the v3 migration surface `[official]`.
- **Quarkus 3.x → 4.0**: Beta 1 Sept 2026, GA Nov 2026; severe breaking changes (Jackson 3) flagged months ahead — the monthly 3.x train is the on-ramp `[secondary]`.
- **General rule**: on lines with monthly/quarterly cadence (Quarkus, Micronaut, Fastify, Hono), automate minor upgrades in CI (renovate/dependabot with test gates); reserve manual effort for majors.

---

## 19. Real-time, deployment and observability notes

- **Server-sent events (SSE)**: first-class in ASP.NET Core 10 `[secondary]`; supported in Hono, Fastify, Axum, Actix-web, Spring (SseEmitter), Quarkus/Micronaut reactive — SSE is the default for server→client push in 2026 unless bidirectional messaging is needed.
- **WebSockets**: Phoenix Channels/LiveView (BEAM, highest connection density), Spring WebSocket/STOMP, Quarkus/Micronaut reactive websockets, NestJS gateways (Socket.io/ws), Fastify (`@fastify/websocket`), Hono (runtime adapters), Axum (native), Django Channels, Rails Action Cable (Solid Cable, DB-backed — no Redis needed).
- **gRPC transport**: grpc-go, grpc-java, grpc-dotnet, tonic (Rust), grpcio (Python), `@grpc/grpc-js`; Spring Boot 4.1 gRPC auto-configuration `[secondary]`; Quarkus/Micronaut/NestJS first-class gRPC modules. HTTP/2 (or HTTP/3 where supported) is assumed — note Fiber/fasthttp has no HTTP/2, which rules out gRPC on the same port.
- **Containers/Kubernetes**: Quarkus and Micronaut optimize for k8s (fast startup, low RSS, health/readiness probes, Dev Services for local deps); Spring Boot has Docker/Buildpacks + actuator probes; ASP.NET Core has first-class container publishing (`dotnet publish` container images, no Dockerfile needed); all others follow the standard distroless/chainguard image pattern.
- **Native AOT / ahead-of-time**: GraalVM native image is first-class in Quarkus and Micronaut (sub-50ms startup); .NET Native AOT is production-viable for ASP.NET Core APIs; Go/Rust are AOT by nature. AOT matters for serverless scale-to-zero and dense k8s packing, not for steady-state throughput.
- **Serverless/edge**: Hono is the TypeScript edge default (Workers/Bun/Deno/Lambda); Prisma's Rust-free engine and Drizzle's zero-dep design target the same runtimes; Python on Lambda favors FastAPI+Mangum or Django+Zappa-style adapters (cold starts remain the tax); Quarkus/Micronaut native images target Knative/Cloud Run scale-to-zero.
- **Observability**: OpenTelemetry is the standard — ASP.NET Core, Spring Boot (Micrometer→OTel bridge), Quarkus, Micronaut, NestJS, Fastify, Django, Rails (via gems) all ship OTel instrumentation; Prometheus/Grafana remains the metrics backend default; structured logging (slog in Go/Echo v5, Serilog in .NET, structlog in Python) is table stakes.
- **Background jobs**: Rails Solid Queue, Django + Celery/Dramatiq, Laravel Queues/Horizon, Symfony Messenger (+2026 parallel-processing option `[official-vendor]`), Spring + ShedLock/Quartz, Quarkus Scheduler, Hangfire/Quartz.NET, BullMQ (Node), Oban (Elixir), litestar-queues 0.6 (own-process workers `[secondary]`), saq/arq (Python async).

---

## 20. Testing and API tooling notes

- **Contract testing**: OpenAPI 3.1 documents feed Prism (mock servers), Schemathesis/Dredd (property-based API fuzzing), and generated clients (openapi-generator, Orval, Hey API). FastAPI/Litestar/NestJS/Hono generate the document from code — the cheapest contract-testing setup in 2026.
- **Python**: pytest + httpx AsyncClient (FastAPI/Litestar TestClient), pytest-django, factory_boy, freezegun; Locust/k6 for load.
- **JVM**: JUnit 5 + Testcontainers (the 2026 default for integration tests against real Postgres/Kafka/Redis), RestAssured/WebTestClient, ArchUnit for architecture rules; Quarkus Dev Services auto-provision containers in dev/test.
- **.NET**: xUnit + WebApplicationFactory (in-memory TestServer), Testcontainers.NET, Alba for API assertions, NBomber/k6 for load.
- **JS/TS**: Vitest/Jest + supertest (Express/Fastify/Hono), @nestjs/testing, Playwright for E2E; tRPC's end-to-end type inference makes API contract tests partially redundant (compiler-checked).
- **Go**: stdlib testing + httptest, testify, Testcontainers-go; k6 (Go-written) or Fortio for load.
- **Rust**: tokio::test + axum-test / actix test utils; criterion for microbenchmarks.
- **Elixir/PHP/Ruby**: ExUnit + Phoenix.ConnTest; PHPUnit/Pest + Laravel HTTP tests; RSpec/Minitest + Rack::Test.
- **Load testing**: k6 is the 2026 default for API load tests across all stacks; Gatling for JVM-centric shops; always test through the real server (Kestrel/uvicorn/Netty), not the test client, before believing numbers.
- **AI-assisted development**: first-party agent skills now ship from Prisma (7.9.0+), Hono (honojs/skills + CLI), and Litestar (litestar-skills) — framework choice in 2026 increasingly includes "how well do coding agents work with this stack" as a criterion `[secondary]`.

---

## 21. GraphQL and gRPC server implementations per stack

- **Python**: Strawberry (code-first, dataclass-based, the 2026 default for new GraphQL APIs) and Ariadne (schema-first); `strawberry-graphql-django` for Django integration; `grpcio` + `grpcio-tools` for gRPC with async support via `grpc.aio`.
- **Java**: Spring GraphQL (annotation-based controllers, the Boot-native path) and Netflix DGS (schema-first, federation-ready); SmallRye GraphQL in Quarkus (code-first with MicroProfile); `micronaut-graphql`; gRPC via `grpc-spring-boot-starter` (auto-configuration new in Boot 4.1 `[secondary]`), `quarkus-grpc` (Mutiny reactive stubs), `micronaut-grpc`.
- **.NET**: Hot Chocolate (the dominant GraphQL server — code-first/schema-first, federation, subscriptions); `graphql-dotnet` (lower-level alternative); gRPC via `grpc-dotnet` (first-class, codegen from `.proto`, JSON transcoding available for REST bridges).
- **PHP**: Lighthouse (Laravel-native, schema-first, the ecosystem default); `webonyx/graphql-php` as the underlying engine; API Platform's GraphQL support for Symfony/API-Platform apps; gRPC via `grpc` PECL extension + `protobuf` (niche in PHP — REST dominates).
- **Ruby**: `graphql-ruby` (code-first classes, interpreter runtime, the standard); gRPC via `grpc` gem (uncommon in Rails shops).
- **Go**: `gqlgen` (schema-first codegen, the standard); `graphql-go/graphql` (code-first alternative); gRPC via `grpc-go` (official, the most mature gRPC implementation of any language here) — Go is arguably the best gRPC citizen.
- **Rust**: `async-graphql` (the standard, code-first/schema-merged, high performance); `tonic` (the standard gRPC implementation, tower-native, pairs with Axum).
- **Elixir**: Absinthe (the GraphQL toolkit, schema-notation macros, subscriptions via Phoenix channels); `grpc-elixir`/`elixir-grpc` for gRPC (smaller ecosystem).
- **TypeScript**: `@nestjs/graphql` (code-first and schema-first, Apollo Server or Mercurius driver); Mercurius (Fastify-native GraphQL adapter); `graphql-yoga` (envelop-based, framework-agnostic, works with Hono/Express/Fastify); tRPC (not GraphQL — the inferred alternative); gRPC via `@grpc/grpc-js` or `nice-grpc` (the modern TS wrapper).
- **Federation note**: Apollo Federation / Hive / WunderGraph gateways sit above these servers when GraphQL spans teams; subgraph compatibility (Hot Chocolate, DGS, gqlgen, Ariadne all support federation) is the 2026 checklist item, not raw query capability.
- **Subscriptions**: GraphQL subscriptions ride WebSockets (or SSE in newer servers); for pure event push, most 2026 teams choose SSE/WebSockets directly rather than GraphQL subscriptions — simpler operations, same UX.

---

