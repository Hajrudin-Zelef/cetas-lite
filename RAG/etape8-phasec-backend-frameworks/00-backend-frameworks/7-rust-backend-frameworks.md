---
id: etape8-phasec-backend-frameworks/00-backend-frameworks/7-rust-backend-frameworks
title: "7. Rust backend frameworks"
domain: step-8-phase-c-backend-frameworks-apis
role: deep-dive
task: reference
actors: ["AWS", "Google", "Lambda"]
dates: ["2025-01", "2025-06-30", "2025-12", "2026-01-27", "2026-05", "2026-07", "2026-07-30", "2026-09", "2026-09-11", "2026-09-22"]
keywords: ["agent", "agents", "aws", "compute", "cost", "latency", "throughput"]
source: docs/RAG/etape8_phaseC_backend_frameworks.md
source_anchor: ""
source_lines: [220, 296]
section: "Step 8 — Phase C: Backend Frameworks & APIs"
sha256: 188da62e6b97217dfc49778f493c147a27f2cf8227bf6bf23bb23a12520a8cc0
---

# 7. Rust backend frameworks

## 7. Rust backend frameworks

### 7.1 Actix-web

- Actix-web 4.13.0 observed in 2026 project usage `[secondary]`:
  - source: `https://github.com/lx-industries/rmcp-actix-web/blob/HEAD/CHANGELOG.md`
- actix-http 3.12.1 contains a request-smuggling security fix; 3.12.0 raised the MSRV to Rust 1.88 `[official]`:
  - source: `https://github.com/actix/actix-web/blob/HEAD/actix-http/CHANGES.md`
- Identity: actor-framework heritage, Tokio-based, historically the TechEmpower plaintext champion; the 4.x line is stable and maintenance-oriented in 2026.
- The MSRV bump to 1.88 matters for enterprise Rust shops pinning toolchains — actix-http now assumes a relatively fresh compiler.

### 7.2 Axum

- Axum 0.8.x is the current line (0.8.0 referenced in 2026 comparisons) `[secondary]`.
- Identity: Tokio team's framework — tower/hyper-based, macro-free routing with typed extractors, first-class WebSocket/SSE/streaming support; the default recommendation for new Rust HTTP services in 2026.
- Axum vs Actix-web in 2026: Axum wins on ecosystem momentum (tower middleware universe, tokio integration, axum-extra); Actix-web retains raw-throughput bragging rights in some TechEmpower categories. Both are production-grade; the choice is ecosystem, not capability.

---

## 8. Elixir — Phoenix

- Phoenix 1.8.x (with LiveView) is the current generation `[unverified — no official 2026 release note captured in this wave; version asserted from release history, patch level not confirmed]`.
- Architecture recap for backend evaluation: Phoenix Channels + LiveView give server-rendered real-time UI over WebSockets with minimal JavaScript; PubSub, Presence and Ecto provide the distributed-systems primitives; the BEAM VM gives preemptive scheduling and fault isolation that maps well to long-lived connection workloads (chat, dashboards, IoT ingestion, multiplayer).
- Performance profile: not a TechEmpower plaintext leader, but connection-density (millions of concurrent WebSockets per node class) is the metric Phoenix wins; LiveView trades per-interaction latency for radically less client code.
- 2026 fit: the strongest choice when the product is real-time-collaborative by default; overkill for plain CRUD JSON APIs where Elixir's smaller hiring pool is a cost.

---

## 9. JavaScript / TypeScript backend frameworks

### 9.1 NestJS

- NestJS 11 is the current major line (shipped January 2025) with the default HTTP adapter switched from Express 4 to Express 5 `[secondary]`:
  - source: `https://tech-insider.org/fastify-vs-express-vs-nestjs-2026/`
- Latest observed patch: `@nestjs/core` 11.1.24 (published May 2026) `[secondary]`; project snapshots from June–July 2026 show 11.1.27/11.1.28 in use — patch-level drift, treat exact latest as `[unverified]`.
- NestJS 11 requires Node.js ≥ 20 (Node 16/18 support dropped); `@nestjs/platform-fastify` updated alongside the Express adapter change `[secondary]`.
- NestJS 12: preparation PRs and September 2026 articles describe it, but no official GA release was confirmed in collected sources — treat as `[unverified]`/pre-release.
- State of JS 2025 (reported 2026): Express remains the usage leader; NestJS "progressing at a nice clip," eating a growing share of new enterprise projects `[secondary]`.
- Identity: Angular-inspired DI, decorators, modules, first-class TypeScript; GraphQL (code-first/schema-first), microservices transports (gRPC, Kafka, Redis, NATS, MQTT), WebSockets, OpenAPI via `@nestjs/swagger` — the enterprise TypeScript default.

### 9.2 Express

- Express 5.2.1 is the current stable npm release (published December 2025), the default install target since 2024; no newer major line as of cutoff `[secondary]`:
  - source: `https://tech-insider.org/fastify-vs-express-vs-nestjs-2026/`
- Express 5 changes (background): promise/async-error handling without next(err) wrappers, stricter routing, removed deprecated APIs — a long-overdue modernization that keeps the massive middleware ecosystem relevant.
- 2026 reality: Express is the legacy incumbent — biggest install base, biggest middleware catalog, slowest innovation. Fine for existing apps and simple services; new performance-sensitive or edge work goes to Fastify/Hono.

### 9.3 Hono

- Hono v4.11.7 released 2026-09-22 (release tagged ~10 hours before cutoff; "Latest" marker observed) `[secondary — GitHub release mirror]`:
  - source: `https://cve.imfht.com/ti_screenshot/2026-01-27/webpage-2026-01-27T20-06-08.090Z-01b5b3615630d00c7807.pdf`
- Hono v4 line: sub-14KB, zero dependencies, built on Web Standard APIs (`Request`/`Response`/`fetch`) — one codebase deploys to Cloudflare Workers, Fastly Compute, Deno, Bun, Vercel, AWS Lambda, Lambda@Edge and Node.js `[secondary]`:
  - sources: `https://dev.to/ottoaria/honojs-in-2026-the-fastest-web-framework-for-cloudflare-workers-and-why-its-going-mainstream-2aap`, `https://sadiqueali.medium.com/hono-in-2026-the-web-framework-that-runs-everywhere-26b83f4868a3`
- 2026 ecosystem maturity: `@hono/zod-validator`, `@hono/jwt`, `@hono/oauth-providers` (GitHub/Google/Discord OAuth in 3 lines), `@hono/swagger-ui`, `@hono/rate-limiter`, `@hono/sentry`, `hono-openapi` — "covers 95% of real-world use cases" per community assessment `[secondary]`.
- Project velocity signals: honojs org ships official agent skills and a CLI "made for AI coding agents," plus agent-dx DX measurement — the framework is optimizing for AI-assisted API development `[secondary]`:
  - source: `https://github.com/honojs`
- Positioning: the default choice for new TypeScript APIs on edge/serverless runtimes in 2026 (Cloudflare Workers, Bun); for traditional Node servers with heavy Express middleware, migration cost rarely justifies the switch `[secondary]`.

### 9.4 Fastify

- Fastify 5.12.4 (tagged 2026-09-11) is the current release `[secondary]`:
  - source: `https://tech-insider.org/fastify-vs-express-vs-nestjs-2026/`
- Fastify 5.11.0 (2026-07-30) added the RFC 10008 HTTP QUERY method, improved response-serialization performance (cached content types), and route/validation fixes `[secondary]`:
  - source: `https://releasebot.io/updates/fastify`
- Security: Fastify 5.12.2 fixed four High-severity advisories (header validation bypass, request validation bypass, auth bypass via malformed URLs, request body replacement) `[secondary]`:
  - source: `https://oday-bakkour.com/blog/dev-stack-release-audit-september-9-2026`
- Lifecycle: Fastify 4 reached EOL 2025-06-30 (no security fixes); Fastify 5 requires Node ≥ 20 (target Node 22/24 — Node 20 is EOL); a 6.0.0-alpha.0 exists on the `next` tag (pre-release, do not target) `[secondary]`:
  - source: `https://github.com/anantbhandarkar/make-it-right/blob/HEAD/skills/mir-backend-node-fastify/SKILL.md`
- Identity: schema-first (JSON Schema validation/serialization with Ajv), plugin architecture, the performance-conscious Node framework — the standard Node choice when Express is too slow and NestJS is too heavy; also NestJS's optional high-performance adapter.
- Minimum patched version guidance from practitioners: 5.8.5+, not just "5.x" (CVE chain) `[secondary]`.

---

## 10. API styles: REST, GraphQL, gRPC, tRPC

Adoption figures below are `[secondary]` (blog/vendor syntheses of job-listing and survey data, April–September 2026). They disagree in absolute numbers; the ranking is consistent.

