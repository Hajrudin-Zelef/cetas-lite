---
id: etape8-phasec-backend-frameworks/00-backend-frameworks/part-12
title: "Step 8 — Phase C: Backend Frameworks & APIs (part 12)"
domain: step-8-phase-c-backend-frameworks-apis
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape8_phaseC_backend_frameworks.md
source_anchor: ""
source_lines: [537, 550]
section: "Step 8 — Phase C: Backend Frameworks & APIs"
sha256: 4c5ecb335ba5735c0b2d2da863f879637b3ac79102ba6af4291c6109b83ef6e6
---

# Step 8 — Phase C: Backend Frameworks & APIs (part 12)

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

