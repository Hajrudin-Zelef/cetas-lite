---
id: etape8-phasec-backend-frameworks/00-backend-frameworks/23-version-and-support-quick-reference-2026-09-22
title: "23. Version and support quick reference (2026-09-22)"
domain: step-8-phase-c-backend-frameworks-apis
role: deep-dive
task: reference
actors: []
dates: ["2025-06-30", "2026-02-19", "2026-03-12", "2026-03-17", "2026-05", "2026-05-29", "2026-07-29", "2026-08-21", "2026-09-02", "2026-09-11", "2026-09-15", "2026-09-16", "2026-09-17", "2026-09-21", "2026-09-22", "2026-09-24", "2026-11-10", "2026-12-31", "2027-03-25", "2027-10-10", "2028-11-14"]
keywords: ["agent", "alignment", "guardrails", "memory", "throughput"]
source: docs/RAG/etape8_phaseC_backend_frameworks.md
source_anchor: ""
source_lines: [638, 699]
section: "Step 8 — Phase C: Backend Frameworks & APIs"
sha256: 94efff0885ccdef28a2a1fccbdfb29acb005a662517d4221e95e489765ffd7d4
---

# 23. Version and support quick reference (2026-09-22)

  @Post()
  @UsePipes(new ValidationPipe({ whitelist: true }))
  create(@Body() dto: CreateItemDto) {
    return this.itemsService.create(dto); // DI-injected, validated, documented via @nestjs/swagger
  }
}
```

---

## 23. Version and support quick reference (2026-09-22)

| Product | Current (cutoff) | Previous / LTS | Support note | Provenance |
|---|---|---|---|---|
| Django | 6.1.1 (2026-09-02) | 5.2.17 LTS, security to 2028-04 | Stay on 5.2.x for LTS | `[independent]` |
| Flask | 3.1.3 (2026-02-19) | 3.0.x | Patch for CVE-2026-27205 | `[secondary]` |
| FastAPI | 0.141.1 (2026-07-29) | 0.11x/0.12x | Rolling minor line | `[secondary]` |
| Litestar | 2.23.0 (2026-05-29) | 2.x | v3 DI overhaul in prep | `[official]` |
| Spring Boot | 4.1.1 (2026-08-21) | 3.4/3.5 (verify EOL) | 4.2 targeted Nov 2026 | `[secondary]` |
| Quarkus | 3.39.4 (2026-09-17) | 3.33 LTS (EOL 2027-03-25); 3.27 LTS EOS 2026-09-24 | 4.0 GA targeted Nov 2026 | `[independent]` |
| Micronaut | 5.2.3 (2026-09-21) | 4.10.28 / 3.10.12 still patched | Java 25 baseline on 5.x | `[official]`/`[independent]` |
| .NET / ASP.NET Core | 10.0.12 LTS (to 2028-11-14) | 9.0.20 STS (EOL 2026-11-10); 8.x LTS (EOL 2026-11-10) | Migrate 8/9 → 10 now | `[official]` |
| Laravel | 13.32 (2026-09-16) | 12.x (maintained); 11.x security-EOL 2026-03-12 | 13 released 2026-03-17 | `[secondary]` |
| Symfony | 8.1 (May 2026) | 7.4 LTS (to 2029); 6.4 LTS (security to 2027-11) | 8.2 targeted Nov 2026; PHP 8.4+ for v8 | `[official]`/`[secondary]` |
| Rails | 8.1.3.1 (2026-07-29) | 8.0.5.1 / 7.2.3.2 backports | 8.1 security to 2027-10-10 | `[secondary]`/`[independent]` |
| Gin | 1.12.0 | 1.11.x | Rolling minor line | `[official]` |
| Echo | v5.3.1 | v4 (security to 2026-12-31) | Migrate v4 → v5 this year | `[official]` |
| Fiber | v3.x (3.1.0+ patched) | v2.52.11 (security) | v3 released Feb 2026 | `[official]`/`[independent]` |
| Actix-web | 4.13.0 | 4.x | actix-http MSRV Rust 1.88 | `[official]`/`[secondary]` |
| Axum | 0.8.x | 0.7.x | Exact latest patch unverified | `[secondary]` |
| Phoenix | 1.8.x | 1.7.x | 2026 patch level unverified | `[unverified]` |
| NestJS | 11.1.x | 10.x | Node ≥ 20; v12 unverified | `[secondary]` |
| Express | 5.2.1 (Dec 2025) | 4.x (legacy) | Default since 2024 | `[secondary]` |
| Hono | 4.11.7 (2026-09-22) | 4.x | Very active patch train | `[secondary]` |
| Fastify | 5.12.4 (2026-09-11) | 4.x EOL 2025-06-30 | Min recommended 5.8.5+ | `[secondary]` |
| SQLAlchemy | 2.0.54 (2026-09-15) | 2.0.x | 2.1 in RC | `[official]` |
| Prisma | 7.9.0 stable; 8.0.0-rc.10 | 6.x | 8 is contract-first rewrite | `[secondary]` |
| Drizzle | 0.45.2 (npm latest) | 1.0.0-rc.3 (rc tag) | Pin rc for v1 experiments | `[secondary]` |
| EF Core | 10 (LTS to 2028-11) | 8.x / 9.x | Ships with .NET 10 | `[secondary]` |
| Hibernate ORM | 7.1 (via Quarkus 3.26) | 6.x | Reactive 3.1; Search 8.1 | `[official]` |
| OpenAPI | 3.1 (current) | 3.0.x | 3.2 status unverified | `[secondary]` |

---

## 24. Selection decision guide (synthesis)

- **Choose Django** when the product is a content/admin-heavy monolith and the team knows Python; choose **FastAPI** for typed JSON APIs and ML serving; choose **Litestar** when you want FastAPI's ergonomics with more opinionated structure (controllers, guards, DTOs).
- **Choose Spring Boot** for large Java teams with existing Spring investment; **Quarkus** for Kubernetes-native Java with fast startup and native-image ambitions; **Micronaut** for compile-time-DI microservices where memory footprint and build-time safety matter most.
- **Choose ASP.NET Core** when raw throughput plus enterprise integration (Azure, Windows, EF Core) matters; Minimal APIs for new JSON services, controllers only when you need their filters/conventions.
- **Choose Laravel** for full-stack PHP products where the surrounding platform (Forge/Vapor/Octane/Herd) reduces ops burden; **Symfony** (or API Platform) for component reuse, complex domains, and enterprise PHP.
- **Choose Rails** for small-team full-stack velocity with Kamal-deployed monoliths; accept mid-pack throughput as the trade.
- **Choose Gin/Echo** for Go APIs that must stay on net/http (ecosystem compatibility, HTTP/2); **Fiber v3** when maximum JSON throughput justifies the fasthttp trade-offs (no HTTP/2, middleware incompatibilities).
- **Choose Axum** for new Rust services (tower ecosystem, Tokio alignment); **Actix-web** when squeezing the last plaintext req/s or maintaining an existing Actix codebase.
- **Choose Phoenix** when real-time collaboration (LiveView/Channels) is the product's core; **NestJS** for enterprise TypeScript modular monoliths; **Fastify** for performance-conscious Node APIs; **Hono** for edge/serverless TypeScript; **Express 5** for maintaining the existing Node estate.
- **Choose SQLAlchemy** (Python default), **Prisma** (TS schema-DSL + guardrails + agent skills) vs **Drizzle** (TS SQL transparency + edge), **EF Core 10** (.NET, no contest), **Hibernate 7.x** (JPA shops).
- **API style**: public → REST + OpenAPI 3.1; TS monorepo frontend → tRPC; dynamic multi-platform UI → GraphQL; internal polyglot mesh → gRPC; simple push → SSE.
- **Auth**: OIDC for SSO, OAuth2 Authorization Code + PKCE for public clients, short-lived JWTs + rotating refresh tokens for APIs, mTLS/DPoP for high assurance; never roll your own token format.

---

---

