---
id: etape8-phasec-backend-frameworks/00-backend-frameworks/11-api-contracts-and-auth-openapi-oauth-2-0-oidc-jwt
title: "11. API contracts and auth: OpenAPI, OAuth 2.0/OIDC, JWT"
domain: step-8-phase-c-backend-frameworks-apis
role: deep-dive
task: reference
actors: ["Google", "Microsoft"]
dates: ["2026-04", "2026-08"]
keywords: ["alignment", "throughput"]
source: docs/RAG/etape8_phaseC_backend_frameworks.md
source_anchor: ""
source_lines: [297, 321]
section: "Step 8 — Phase C: Backend Frameworks & APIs"
sha256: 1e9067dd532e9bf6f96aa62d7f7748d953f747d0cd4b256fec39f71886a0674f
---

# 11. API contracts and auth: OpenAPI, OAuth 2.0/OIDC, JWT

- REST: 70%+ of developer job listings (APIScout, April 2026); 93% among API teams vs 14% gRPC (Postman 2025 State of the API, via Refonte Learning summary August 2026); ~80% of business APIs (Optimum Web). Verdict: the default for public APIs; OpenAPI 3.1 + HTTP/3 strengthened its position `[secondary]`:
  - sources: `https://apiscout.dev/guides/rest-vs-graphql-vs-grpc-vs-trpc-2026`, `https://tech-insider.org/grpc-vs-rest-2026/`, `https://www.optimum-web.com/blog/api-development-services-2026-rest-graphql-grpc-guide/`
- GraphQL: ~25% enterprise adoption (down from a ~40% peak), concentrated in orgs with complex frontend data needs and multiple client platforms `[secondary]`. The "GraphQL is dying" narrative overstates it: it retreated from hype-cycle ubiquity to its actual sweet spot (BFF aggregation, multi-platform data) — N+1 discipline and persisted queries/APQ remain the operational taxes.
- gRPC: dominant for internal service-to-service (Google, Netflix, Uber scale: billions of RPCs/day; 7–10x throughput gains over JSON REST on serialization-heavy workloads per vendor claims); ~14% in the Postman survey; requires gRPC-Web proxy for browsers, opaque to curl — the debugging tax is real `[secondary]`.
- tRPC: ~15% of TypeScript job postings and climbing; 37,000+ GitHub stars; the default API layer for the T3 stack (Next.js + Prisma + tRPC); TypeScript-monorepo-only by design — kills adoption in mixed-language teams `[secondary]`.
- The 2026 consensus pattern is hybrid: REST/OpenAPI for public APIs, tRPC or GraphQL for the frontend data layer, gRPC for internal microservices `[secondary]`:
  - source: `http://dev.to/wantsvibes/api-architecture-comparison-rest-vs-graphql-vs-trpc-vs-grpc-for-cloud-native-backends-2m0o`
- Decision rule of thumb (synthesized): public + third-party → REST; TypeScript monorepo SaaS → tRPC; multi-platform dynamic UI → GraphQL; internal polyglot microservices → gRPC; bidirectional real-time → gRPC streams or GraphQL subscriptions.

---

## 11. API contracts and auth: OpenAPI, OAuth 2.0/OIDC, JWT

- OpenAPI 3.1 is the current contract standard in 2026: ASP.NET Core 10 generates OpenAPI 3.1 documents natively `[secondary]`; FastAPI/Litestar/NestJS/Hono emit 3.x documents by default. JSON Schema alignment (3.1) is the reason 3.1 won over 3.0 — one schema dialect for validation and documentation.
- OpenAPI 3.2 status: not confirmed in collected sources — mark as `[unverified]`; do not assume tooling support (see gaps).
- Codegen economy: OpenAPI Generator / Orval / Hey API / oRPC-adjacent tooling turn 3.1 documents into typed clients; the contract-first workflow (design in OpenAPI, generate server stubs + clients) is standard in enterprise Java/.NET shops, while the TypeScript world increasingly prefers inferred contracts (tRPC, Hono RPC client, FastAPI→client codegen).
- AuthN/AuthZ stack in 2026:
  - OAuth 2.0 + OIDC remain the delegation/SSO standards (Authorization Code + PKCE for public clients; client credentials for service-to-service).
  - JWT (JWS compact) remains the access-token format; opaque reference tokens persist where revocation matters (introspection endpoint).
  - Framework support is table stakes: Quarkus OIDC (with client-side token refresh on 401 as of 3.26 `[official]`), Spring Security OAuth2, ASP.NET Core JWT Bearer + OIDC handlers, NestJS Passport strategies, `@hono/jwt` + `@hono/oauth-providers`, Echo/Fiber JWT middleware, Laravel Passport/Sanctum, Symfony Security.
  - 2026 hardening notes: short-lived access tokens + rotating refresh tokens; DPoP/mTLS sender-constraining for high-assurance APIs; large IdP JWTs are a real operational concern (Fiber documents explicit `ReadBufferSize` sizing for Clerk-style JWTs `[secondary]`).
- SSE vs WebSockets vs gRPC streams for push: SSE won the "simple server→client push" slot (ASP.NET Core 10 added first-class SSE `[secondary]`; Hono/Fastify/Axum all support it); WebSockets for bidirectional; gRPC streams for service mesh internals.

---

