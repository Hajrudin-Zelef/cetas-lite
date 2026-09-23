---
id: etape8-phasec-backend-frameworks/00-backend-frameworks/13-benchmarks-techempower-and-methodology-limits
title: "13. Benchmarks — TechEmpower and methodology limits"
domain: step-8-phase-c-backend-frameworks-apis
role: deep-dive
task: benchmark
actors: ["Falcon"]
dates: ["2025-02", "2026-02-19", "2026-03-24", "2026-05", "2026-05-29", "2026-07-29", "2026-08-21", "2026-09-11", "2026-09-16", "2026-09-17", "2026-09-21", "2026-09-22", "2028-11-14"]
keywords: ["benchmark", "benchmarks", "embedding", "latency", "memory", "throughput"]
source: docs/RAG/etape8_phaseC_backend_frameworks.md
source_anchor: ""
source_lines: [392, 446]
section: "Step 8 — Phase C: Backend Frameworks & APIs"
sha256: 5a876146e37e8fb7a7146eae0478d981b911176a6482761ce80311c6ece8c66e
---

# 13. Benchmarks — TechEmpower and methodology limits

## 13. Benchmarks — TechEmpower and methodology limits

### 13.1 TechEmpower FrameworkBenchmarks status (critical correction)

- TechEmpower Round 23 (published February 2025) appears to be the last published round `[secondary]`.
- The TechEmpower FrameworkBenchmarks repository was archived on 2026-03-24 — there is **no published Round 24** in collected sources; earlier task assumptions of a Round 24 were wrong `[secondary]`:
  - sources: `https://dev.to/kaliumhexacyanoferrat/techempower-framework-benchmarks-are-now-archived-whats-next-3l0a`, `https://github.Com/TechEmpower/FrameworkBenchmarks/issues`
- This file therefore reports benchmark posture from Round 23 and framework characteristics, not Round 24 numbers — inventing Round 24 results would be fabrication.

### 13.2 What Round 23-era results say (directional, not absolute)

- Plaintext/JSON microbenchmark leaders: Rust (Actix-web historically #1 in plaintext), Go (Fiber/gofiber, Echo, Gin near the top), .NET (ASP.NET Core Kestrel, consistently top-tier among full frameworks), Java (lightweight stacks over Spring MVC).
- Python async (FastAPI, Litestar, Starlette, Django async) clusters far below the compiled frameworks on raw req/s — the gap is real but mostly irrelevant outside >10k req/s/core workloads.
- Node frameworks (Fastify, Hono on Bun/Node, NestJS) sit mid-pack; Hono-on-Bun/edge numbers are not directly comparable to Node-server numbers.
- PHP (Laravel Octane/FrankenPHP, Symfony Runtime) improved vs php-fpm baselines but remains throughput-modest; Rails similar.

### 13.3 Methodology limits (read before citing any number)

- TechEmpower tests are **microbenchmarks**: JSON serialization, single-query DB, multi-query, fortunes (HTML), plaintext — they measure framework + driver + tuning, not application performance.
- Results depend on: hardware (the official runs use specific dedicated servers), protocol (HTTP/1.1 vs h2), server (Kestrel vs uvicorn vs fasthttp vs net/http), database driver (Npgsql vs pgjdbc vs asyncpg vs node-postgres), ORM vs raw SQL, connection pooling, PGO/JIT warmup, and per-framework tuning effort by the contributor.
- Database tests conflate framework overhead with driver efficiency — a "framework" win is often a driver win.
- Real applications are dominated by: query design, N+1s, caching, serialization choices, network topology, and tail latency — none of which TechEmpower measures.
- Framework versions in Round 23 are frozen at February 2025 — they do not reflect 2026 releases (Fiber v3, Hono 4.11, Fastify 5.12, Quarkus 3.39, .NET 10).
- Verdict: use TechEmpower for **architectural signals** (event-loop vs thread-per-request vs virtual threads; fasthttp vs net/http; native AOT vs JIT) and **never** for procurement-grade "framework X is N% faster" claims.

---

## 14. Comparative matrix (as of 2026-09-22)

| Framework | Current version (cutoff) | Min runtime | Concurrency model | Built-in ORM/data | OpenAPI | GraphQL/gRPC | Maturity / recommended use |
|---|---|---|---|---|---|---|---|
| Django | 5.2.17 LTS (sec. to 2028-04) `[independent]`; 6.1.1 latest feature `[independent]` | Python 3.10+ (3.12/3.13 typical) | Sync WSGI + async ASGI paths | Django ORM + migrations | drf-spectacular / Ninja | graphene / strawberry; grpcio | Batteries-included monoliths, admin-heavy apps, content platforms |
| Flask | 3.1.3 (2026-02-19) `[secondary]` | Python 3.9+ | Sync WSGI (async views opt-in) | Extensions (Flask-SQLAlchemy) | flasgger / apispec | Extensions | Small services, legacy APIs, embedding in larger apps |
| FastAPI | 0.141.1 (2026-07-29) `[secondary]` | Python 3.8+ (3.10+ realistic) | Async (Starlette) + threadpool for sync | SQLAlchemy (external) | Native (auto) | strawberry; grpcio | Typed JSON APIs, ML/AI model serving, internal platforms |
| Litestar | 2.23.0 (2026-05-29) `[official]` | Python 3.8+ | Async ASGI | advanced-alchemy / SQLAlchemy | Native (auto) | strawberry plugin | Opinionated FastAPI alternative; controllers, guards, DTOs |
| Spring Boot | 4.1.1 (2026-08-21) `[secondary]`; scope asked 3.4/3.5 | Java 17+ (21/25 typical) | Servlet threads / WebFlux reactive / virtual threads (21+) | Spring Data JPA (Hibernate) | springdoc-openapi | Netflix DGS / Spring GraphQL; grpc-spring-boot-starter (4.1 auto-config `[secondary]`) | Enterprise Java default; massive ecosystem; heaviest runtime |
| Quarkus | 3.39.4 (2026-09-17) `[independent]`; 3.33 LTS | Java 17+ | Reactive (Vert.x/Mutiny) + imperative; virtual threads | Hibernate ORM 7.1 / Panache `[official]` | smallrye-openapi | smallrye-graphql; quarkus-grpc | Cloud-native Java, native-image serverless, k8s density |
| Micronaut | 5.2.3 (2026-09-21) `[independent]` | Java 25 baseline (5.x) `[official]` | Netty event loop; virtual threads first-class | Micronaut Data (compile-time) | micronaut-openapi | micronaut-graphql; micronaut-grpc | Compile-time-DI microservices; low memory; GraalVM native |
| ASP.NET Core | .NET 10.0.12 LTS (to 2028-11-14) `[official]` | .NET 10 | Async/await + Kestrel; Native AOT option | EF Core 10 `[secondary]` | Native 3.1 gen (Minimal APIs) `[secondary]` | Hot Chocolate; grpc-dotnet | Top throughput among full frameworks; enterprise + high-perf APIs |
| Laravel | 13.32 (2026-09-16) `[secondary]`; scope asked 11/12 | PHP 8.2+ (8.4 for latest) | php-fpm / Octane long-lived workers | Eloquent | Scribe / Scramble | Lighthouse; protoc-gen | Full-stack PHP products; surrounding platform (Forge/Vapor/Octane) |
| Symfony | 8.1 (May 2026) `[secondary]`; 7.4 LTS | PHP 8.4+ (v8) `[official]` | php-fpm / FrankenPHP workers | Doctrine ORM | NelmioApiDocBundle | API Platform (GraphQL); grpc | Component ecosystem; API Platform; enterprise PHP |
| Rails | 8.1.3.1 (2026-07-29) `[secondary]` | Ruby 3.2+ (YJIT) | Puma threads / Falcon async | Active Record | rswag | graphql-ruby; grpc | Solo/small-team full-stack; Kamal-deployed monoliths |
| Gin | 1.12.0 `[official]` | Go 1.21+ | Goroutines + net/http | GORM/sqlx (external) | swaggo | gqlgen; grpc-go | Default Go API framework; broad adoption |
| Echo | v5.3.1 `[official]` | Go (last 4 majors) `[official]` | Goroutines + net/http | GORM/sqlx (external) | swaggo / oapi-codegen | gqlgen; grpc-go | net/http-compatible Go APIs; middleware depth |
| Fiber | v3.x (v3.1.0+ patched) `[official]` | Go 1.21+ | Goroutines + fasthttp | GORM/sqlx (external) | swaggo | fiber contrib; grpc-go | Max-throughput Go JSON APIs (accept fasthttp trade-offs) |
| Actix-web | 4.13.0 `[secondary]` | Rust 1.88 (MSRV) `[official]` | Tokio async | Diesel/SQLx (external) | utoipa / aide | async-graphql; tonic | Raw throughput; control over allocations |
| Axum | 0.8.x `[secondary]` | Recent stable Rust | Tokio async (tower) | SQLx (typical) | utoipa / aide | async-graphql; tonic | New Rust services; tower ecosystem |
| Phoenix | 1.8.x `[unverified]` | Elixir/OTP | BEAM processes (preemptive) | Ecto | open_api_spex | Absinthe; grpc-elixir | Real-time (LiveView/Channels); connection density |
| NestJS | 11.1.x (11.1.24 obs.) `[secondary]` | Node ≥ 20 `[secondary]` | Node event loop (Express/Fastify adapter) | TypeORM/Prisma/MikroORM | @nestjs/swagger | @nestjs/graphql; @nestjs/microservices (gRPC) | Enterprise TypeScript; modular monoliths |
| Express | 5.2.1 (Dec 2025) `[secondary]` | Node 18+ | Node event loop | External | swagger-jsdoc | External | Legacy/middleware-rich Node services |
| Hono | 4.11.7 (2026-09-22) `[secondary]` | Any JS runtime | Runtime event loop (edge) | Drizzle/Prisma (external) | hono-openapi `[secondary]` | External | Edge/serverless TypeScript APIs (Workers/Bun/Deno) |
| Fastify | 5.12.4 (2026-09-11) `[secondary]` | Node ≥ 20 `[secondary]` | Node event loop | Prisma/Drizzle (external) | @fastify/swagger (native schema) | mercurius; grpc-js | Performance-conscious Node APIs; JSON-schema-first |

---

