---
id: etape8-phasec-backend-frameworks/00-backend-frameworks/overview
title: "Step 8 — Phase C: Backend Frameworks & APIs"
domain: step-8-phase-c-backend-frameworks-apis
role: deep-dive
task: reference
actors: ["Anthropic", "Google", "Microsoft", "vLLM"]
dates: ["2026-02-19", "2026-05-29", "2026-06", "2026-07", "2026-07-25", "2026-07-29", "2026-09-02", "2026-09-22", "2028-04"]
keywords: ["advisory", "agent", "benchmark", "claude", "copilot", "gemini", "inference", "packaging", "research", "throughput", "vllm"]
source: docs/RAG/etape8_phaseC_backend_frameworks.md
source_anchor: ""
source_lines: [1, 75]
section: "Step 8 — Phase C: Backend Frameworks & APIs"
sha256: 2d57471032130b7d2901df1e067de7ff9d6c4ebf68ebc889f5d3cd8abb42922f
---

# Step 8 — Phase C: Backend Frameworks & APIs

## Research header

- **Scope:** backend web frameworks and API tooling, developer angle: versions, features, performance and adoption.
- **Cutoff date:** 2026-09-22. All statements are current as of this date unless dated otherwise.
- **Method:** web research conducted on 2026-09-22 across official release notes, changelogs, vendor documentation, security advisories and community reporting; conflicting or unverified claims are flagged inline.
- **Provenance legend (every factual claim carries one):**
  - `[official]` — vendor or project documentation / release notes / changelog / README.
  - `[vendor-reported]` — vendor blog or announcement, not independently audited.
  - `[independent]` — third-party measurement or audit (EOL trackers, security databases).
  - `[secondary]` — press, blogs, community digests, skill repos; plausible but not authoritative.
  - `[unverified]` — claim seen in project snapshots or preparation PRs without a confirming official release.
- **Rule:** URLs, versions, prices and identifiers are copied verbatim from sources; none are invented. Missing data is reported as a gap, not filled in.
- **Scope of this file:** Python (Django 5.x, Flask, FastAPI 0.115+, Litestar, async ecosystem), Java (Spring Boot 3.4/3.5, Quarkus, Micronaut), .NET (ASP.NET Core 9/10, Minimal APIs), PHP (Laravel 11/12, Symfony), Ruby (Rails 8.x), Go (Gin, Echo, Fiber), Rust (Actix-web, Axum), Elixir (Phoenix), JavaScript/TypeScript (NestJS, Express, Hono, Fastify), ORMs (SQLAlchemy, Prisma, Drizzle, Entity Framework Core, Hibernate), API styles (REST, GraphQL, gRPC, OpenAPI, OAuth 2.0/OIDC, JWT), and TechEmpower Round 23/24 benchmark context with methodology limits.

---

## 1. Python backend frameworks

### 1.1 Django

- Django 5.2.17 is the current Django 5.x line release observed as the latest patch on the 5.x branch `[independent]`:
  - source: `http://eosl.date/eol/product/django/`
- Django 5.2 is the long-term support (LTS) release line, with security support indicated until April 2028 `[independent]`:
  - source: `http://eosl.date/eol/product/django/`
- Django 6.1.1 was released 2026-09-02, so Django 5.2 remains the requested LTS scope but is no longer the latest feature-release line `[independent]`:
  - source: `https://linuxsecurity.com/advisories/fedora/fedora-43-python-django5-2026-6b28b4e483`
- Consequences for greenfield vs maintenance:
  - New projects that want the LTS safety net stay on 5.2.x and take patch upgrades; projects that want the newest ORM/forms features track 6.x.
  - Security advisories in 2026 (e.g. Fedora 43 python-django5 advisory stream) continue to reference the 5.x packaging line `[secondary]`.
- Django's async story in the 5.x era: ASGI support, async views, async ORM query paths (QuerySet async methods), async middleware and async-safe session/auth primitives — matured through 4.x and 5.x `[official docs, prior knowledge — treat as background, not a 2026 change]`.
- Performance profile: Django is not a top-end throughput contender; its strength is batteries-included productivity (ORM, admin, migrations, auth) with horizontal scaling via WSGI/ASGI workers (gunicorn/uvicorn), caching, and read replicas.

### 1.2 Flask

- Flask 3.1.3 released 2026-02-19 is the current stable line `[secondary]`:
  - sources: `https://archlinux.org/packages/extra/any/python-flask/`, `https://github.com/cloudbrain/htm-challenge/commit/afd0de4bf47f15a342a667215e878001056eb1a3`
- Flask 3.1.3 fixed CVE-2026-27205 (Vary: Cookie handling issue) `[secondary]`:
  - source: `https://github.com/cloudbrain/htm-challenge/commit/afd0de4bf47f15a342a667215e878001056eb1a3`
- Companion releases in the Pallets ecosystem around the same date: Werkzeug 3.1.3 and Jinja2 3.1.6 `[secondary]`.
- Positioning: micro-framework — routing, request/response, Jinja templating, Werkzeug WSGI — with extensions (Flask-SQLAlchemy, Flask-Login, Flask-Migrate) carrying the batteries; async support exists (Flask 2.x+ async views) but the ecosystem is predominantly sync WSGI.
- Typical production shape: gunicorn (sync or gevent workers) or waitress behind nginx; FastAPI/Litestar now absorb most greenfield async Python API work that Flask used to get.

### 1.3 FastAPI

- FastAPI 0.141.1, dated 2026-07-29, is the latest version observed in 2026 `[secondary]`:
  - source: `https://github.com/ku5ic/dotfiles/blob/HEAD/claude/skills/fastapi-patterns/SKILL.md`
- Core architecture unchanged since the 0.11x line: Starlette for HTTP routing/middleware, Pydantic for validation/serialization, automatic OpenAPI + JSON Schema generation from type hints `[secondary]`.
- Accepts both `def` and `async def` handlers; sync handlers run in a threadpool, async handlers on the event loop — mixing is routine and documented `[secondary]`.
- Ecosystem anchor: the de-facto standard for Python AI/ML model serving (vLLM, text-generation-inference frontends, agent tool servers) and typed internal APIs, largely because OpenAPI generation + Pydantic validation are free by default.
- Performance: top-tier within Python async frameworks on TechEmpower-style JSON/plaintext tests relative to Django/Flask; behind compiled-language frameworks in absolute numbers (see section 14).

### 1.4 Litestar

- Litestar 2.23.0 released 2026-05-29 is the current 2.x release `[official]`:
  - source: `https://docs.litestar.dev/latest/release-notes/changelog.html`
- 2.23.0 introduced generic request-body markers: `JSONBody[T]`, `MsgPackBody[T]`, `MultipartBody[T]`, `URLEncodedBody[T]`, replacing verbose `Annotated[T, Body(media_type=...)]` forms `[official]`.
- 2.23.0 deprecated `params.Dependency` in favor of `di.NamedDependency`, and introduced `params.SkipValidation`, as preparation for an overhauled dependency-injection system planned for version 3 (type-based DI alongside name-based DI, with a `TypeDependency` counterpart) `[official]`.
- Litestar identity: "light, flexible and extensible ASGI framework | Built to scale" — opinionated API framework with class-based controllers, msgspec-powered data parsing, plugin system, ORM support, DTOs, OpenAPI, dependency injection, guards and middleware `[official]`:
  - source: `https://github.com/litestar-org`
- Companion ecosystem in 2026: `advanced-alchemy` (SQLAlchemy companion library, 788 stars, updated June 2026) and `sqlspec` ("a Query Mapper for Python", updated July 2026) show the Litestar org doubling down on the SQLAlchemy/asyncpg data layer `[secondary]`.
- Agent-tooling angle: the Litestar org publishes first-party agent skills (`litestar-skills`) targeting Claude Code, Gemini CLI, Codex CLI, Cursor, OpenCode and VS Code/Copilot — the framework is explicitly positioning for AI-assisted development `[secondary]`.
- Background-task queues: `litestar-queues` 0.6.0 (2026-07-25) moved queue workers into their own process started by the Litestar CLI, with `placement` (`server`/`asgi`/`external`) replacing `run_in_app` — a breaking change that signals production-hardening of the background-job story `[secondary]`:
  - source: `https://github.com/hasansezertasan/litestar-queues/blob/HEAD/docs/changelog.rst`

### 1.5 Python async ecosystem notes

- The Python backend landscape in 2026 is firmly async-first for APIs: FastAPI, Litestar, Starlette and Django's async paths dominate new API work; sync Flask/Django WSGI remains the legacy maintenance bulk.
- ASGI servers: uvicorn and Granian are the standard production choices; hypercorn persists for HTTP/2/3 experiments.
- Data layer: SQLAlchemy 2.0.x (see 12.1), asyncpg for PostgreSQL, and msgspec (used by Litestar) for fast serialization are the performance-critical pieces; Pydantic v2 (Rust core) underpins FastAPI validation throughput.
- Free-threaded Python (3.13t/3.14t, no-GIL builds) is now a real deployment consideration — SQLAlchemy 2.0.54 explicitly fixed Cython extension compatibility for free-threaded builds `[official]` (see 12.1).

---

