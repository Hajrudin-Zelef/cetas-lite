---
id: etape8-phasec-backend-frameworks/00-backend-frameworks/4-php-backend-frameworks
title: "4. PHP backend frameworks"
domain: step-8-phase-c-backend-frameworks-apis
role: deep-dive
task: reference
actors: []
dates: ["2025-11", "2026-01-18", "2026-02", "2026-02-11", "2026-03-12", "2026-03-17", "2026-05", "2026-07-29", "2026-09-16", "2026-11", "2026-12-31", "2027-01", "2027-10-10"]
keywords: ["benchmarks", "claude", "embedding", "energy", "throughput"]
source: docs/RAG/etape8_phaseC_backend_frameworks.md
source_anchor: ""
source_lines: [139, 219]
section: "Step 8 — Phase C: Backend Frameworks & APIs"
sha256: 88465198b2d66a767017fa8ca37fc48c5932e1e0f5496b4491d61b150aa5aa18
---

# 4. PHP backend frameworks

## 4. PHP backend frameworks

### 4.1 Laravel — scope asked for 11/12, current is 13

- Laravel 13 released 2026-03-17; latest observed patch 13.32 (2026-09-16) `[secondary]`:
  - sources: `https://laravel-news.com/tag/releases`, `https://dev.to/marco_odev/laravel-13-what-actually-breaks-when-you-upgrade-and-what-doesnt-13d6`
- Laravel 11 is out of security support since 2026-03-12 `[secondary]`; Laravel 12 remains the relevant migration/maintenance baseline for the requested scope `[secondary]`.
- Upgrade notes for 12→13 (secondary, verify against official upgrade guide): mostly additive; breaking changes concentrate in dropped deprecated helpers and minimum PHP bumps — the linked dev.to piece documents what actually breaks `[secondary]`.
- "AI-native" framing around Laravel 13 appears in secondary coverage; treat vendor-marketing claims as `[vendor-reported]` until confirmed in official release notes `[unverified for specifics]`.
- Laravel's 2026 moat: the surrounding product surface (Forge, Vapor, Nova, Herd, Octane for long-lived workers via FrankenPHP/Swoole/RoadRunner) matters more than framework-version deltas for TCO discussions.

### 4.2 Symfony

- Symfony 8.1 is the current stable line (released May 2026; community support indicated to January 2027) `[secondary lifecycle matrix]`:
  - source: `https://www.cmarix.com/blog/different-versions-of-symfony-framework/`
- Symfony 8.0 and 7.4 LTS released simultaneously end of November 2025; 8.0 contains no deprecated features vs 7.4 — upgrade path is "resolve all deprecations" `[official]`:
  - source: `https://github.com/symfony/symfony/blob/HEAD/UPGRADE-8.0.md`
- Symfony v8 requires PHP 8.4+ `[official]`.
- Symfony 8.1 FrameworkBundle changelog highlights: validation `property_metadata_existence_check`, JsonStreamer value-object handling, per-pool cache marshallers, `lock://` semaphore DSN, `createFormFlowBuilder` in AbstractController, mocking non-shared services in tests, `ConsoleCommandAssertionsTrait`, HTML sanitizer defaults, decoration-stack in `debug:container`, `MicroKernelTrait::getAllowedEnvs()` `[official]`:
  - source: `https://github.com/symfony/symfony/blob/HEAD/src/Symfony/Bundle/FrameworkBundle/CHANGELOG.md`
- Symfony 8.2 in development (target November 2026); weekly development notes show PropertyInfo attribute configurability, RFC 9213 Cache-Control in HttpFoundation, Messenger parallel-processing concurrency option, Mailer rate-limited transports, and a newly announced Symfony AI Core Team `[official-vendor blog]`:
  - source: `https://symfony.com/blog/a-week-of-symfony-1024-august-10-16-2026`
- New components introduced across the 8.x generation: Emoji, JsonPath, JsonStreamer, ObjectMapper, TypeInfo `[secondary — keynote recap]`:
  - source: `https://github.com/gnugat/knowledge/blob/HEAD/talks-recapped/2026-01-symfony-online/001-20-years-of-symfony-whats-next.md`
- Scale signal: Symfony packages account for roughly 25% of all PHP package installs on Packagist (~25M downloads/day, noisy with CI) — the components are the de-facto PHP standard library, embedded in Laravel, API Platform, Composer itself `[vendor-reported keynote]`; treat the exact percentage as marketing math, the embedding claim as broadly true.
- Runtime story: the Runtime component decouples Symfony from PHP-FPM; FrankenPHP (with worker mode), Swoole and RoadRunner are the modern deployment targets — relevant because long-lived workers change the PHP performance conversation vs classic php-fpm.

---

## 5. Ruby — Rails 8.x

- Rails 8.1.3.1 (2026-07-29) is the current patch release `[secondary]`:
  - source: `https://github.com/fmanimashaun/claude-skills/blob/HEAD/skills/rails-8/SKILL.md`
- 8.1.3.1 was a critical security release: CVE-2026-66066 in Active Storage, with backports 8.0.5.1 and 7.2.3.2 `[secondary]`.
- Rails 8.1 security support indicated until 2027-10-10 `[independent]`:
  - source: `http://eosl.date/eol/product/rails/`
- Rails 8 generation identity (background, 2024–2025): Kamal 2 deployment, Thruster (HTTP/2 proxy + caching + X-Sendfile), Solid Queue/Cache/Cable (database-backed job queue, cache, websockets — no Redis required), Propshaft asset pipeline, no-Node default frontend, `rails new` authentication generator. The 8.x line's pitch is "one person, one VPS, one framework" — deployment simplicity as a feature.
- 2026 maintenance reality: the framework is in harvest mode — security patches and minor refinements rather than architectural churn; most innovation energy is in the Hotwire/Stimulus/Turbo front-end layer and hosting (Kamal) rather than the Ruby core.
- Performance: Rails remains mid-pack on TechEmpower; production Rails performance in 2026 is dominated by database/caching architecture (Solid Cache, read replicas) and YJIT (now default in modern CRuby) rather than framework internals.

---

## 6. Go backend frameworks

### 6.1 Gin

- Gin 1.12.0 is the current release line per the official README `[official]`:
  - source: `https://github.com/gin-gonic/gin/blob/HEAD/README.md`
- A secondary source dates the 1.12.0 release to February 2026 `[secondary]`.
- Identity: httprouter-derived radix-tree routing, middleware chains, binding/validation, ~the default Go HTTP framework by adoption; JSON rendering via encoding/json (with sonic/go-json drop-ins common in performance-sensitive shops).
- Go 1.24+ context: Gin benefits from the ongoing runtime improvements (PGO, range-over-func) without framework changes; the main 2026 decision is Gin vs stdlib+chi vs Fiber, not Gin version.

### 6.2 Echo

- Echo v5 is the current major line (v5 as of 2026-01-18); the official site announces "Echo v5.3.1 — now released" `[official]`:
  - sources: `https://github.com/labstack/echo/blob/HEAD/README.md`, `https://echo.labstack.com/`
- Echo v4 remains supported with security updates and bug fixes until 2026-12-31 `[official]`.
- Architecture: built on Go's standard `net/http` (interop via `echo.WrapHandler`/`echo.WrapMiddleware`), fast radix-tree router with zero dynamic allocation per request, request binding (pluggable validator), 25+ built-in middleware, centralized error handling, automatic TLS via Let's Encrypt, HTTP/2 support `[official]`.
- Adoption signal: 32.7k GitHub stars `[official site]`; the Encore sponsorship banner reflects real commercial backing.
- v5 migration notes live in API_CHANGES_V5.md; README instructs `go get github.com/labstack/echo/v5` `[official]`.
- Trade-off vs Fiber: Echo stays on net/http (full ecosystem compatibility, HTTP/2), Fiber uses fasthttp (raw speed, no HTTP/2, middleware incompatibilities) — see 6.3.

### 6.3 Fiber

- Fiber v3 released February 2026 (official "What's new in Fiber v3" blog, 2026-02-11) — the first major since v2 (2021) `[official]`:
  - source: `https://github.com/gofiber/docs/blob/HEAD/blog/2026-02-11-whats-new-in-fiber-v3.md`
- Headline v3 changes:
  - `fiber.Ctx` natively implements Go's standard `context.Context` — handlers take `c fiber.Ctx` and pass `c` directly to repositories/gRPC clients/tracing; the v2 `.UserContext()` ritual is gone `[official][secondary explainer]`.
  - Sessions re-architected as middleware (register once via `app.Use(session.New())`, access via `session.FromContext(c)`) with storage backends and automatic lifecycle `[official]`.
  - Stricter middleware prefix matching (slash-boundary), `app.Static()`/`Filesystem` middleware removed in favor of the static middleware, middleware data via `FromContext()` instead of `c.Locals()` string keys `[official]`.
  - Cookie/RFC hardening: auto-Secure for SameSite=None, CHIPS partitioned-cookie support, RFC 6266/8187 attachment filenames, default redirect 302→303 `[official]`.
  - Compression middleware gains zstd alongside gzip/deflate/brotli; storage adapters gain `WithContext` variants; client package rebuilt (cookie jar, hooks, retries, proxy, debug); new Services feature for app dependencies `[official]`.
  - CLI migration helper: `go install github.com/gofiber/cli/fiber@latest && fiber migrate --to v3` `[official]`.
- Security: Fiber v3.1.0 patched a static-middleware path-traversal (Windows) and a flash-cookie msgpack deserialization flaw (CVE-2026-25899, 10-char cookie → up to 85GB allocation attempt, CVSS 7.5); Fiber v2.52.11 fixed CVE-2025-66630 (crypto/rand failure → predictable UUIDs in sessions/CSRF, CVSS 9.4) `[independent]`:
  - source: `https://app.opencve.io/cve/?vendor=gofiber`
- Real-world caveat from a production user (secondary): choosing Fiber means choosing fasthttp — no HTTP/2, third-party net/http middleware doesn't drop in (adaptor exists but compatibility must be checked per library), reused fixed-size buffers require explicit `ReadBufferSize` sizing for large JWTs `[secondary]`:
  - source: `https://medium.com/devsameday/go-fiber-v3-is-here-breaking-changes-shiny-new-superpowers-and-why-you-should-finally-upgrade-24bac84fac97`
- Net: Fiber v3 is the fastest mainstream Go framework on plaintext benchmarks and a legitimate 2026 default for high-throughput JSON APIs — provided the team accepts the fasthttp trade-offs.

---

