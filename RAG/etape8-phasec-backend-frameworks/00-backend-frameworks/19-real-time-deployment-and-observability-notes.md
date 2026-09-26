---
id: etape8-phasec-backend-frameworks/00-backend-frameworks/19-real-time-deployment-and-observability-notes
title: "19. Real-time, deployment and observability notes"
domain: step-8-phase-c-backend-frameworks-apis
role: deep-dive
task: reference
actors: ["Lambda"]
dates: ["2026-11-10", "2026-12-31"]
keywords: ["agent", "agents", "inference", "mcp", "memory", "throughput"]
source: docs/RAG/etape8_phaseC_backend_frameworks.md
source_anchor: ""
source_lines: [493, 536]
section: "Step 8 — Phase C: Backend Frameworks & APIs"
sha256: 7355615f2ef2d298c45490a32ea03fe4b85858dc0097945b83df5000136f736d
---

# 19. Real-time, deployment and observability notes

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

