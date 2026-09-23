---
id: etape7-phaseh-webproxy/00-webproxy/3-nginx
title: "3. nginx"
domain: step-7-phase-h-web-servers-reverse-proxies-api-gateways
role: deep-dive
task: reference
actors: []
dates: ["2025-04-23", "2025-06-24", "2026-03", "2026-03-24", "2026-04", "2026-04-07", "2026-04-08", "2026-04-14", "2026-05", "2026-05-13", "2026-05-27", "2026-06", "2026-06-06", "2026-07-06", "2026-07-15", "2026-07-18", "2026-08", "2026-08-19", "2026-08-30", "2026-09", "2026-09-15"]
keywords: ["advisory", "apache", "benchmark", "packaging", "pricing", "research"]
source: docs/RAG/etape7_phaseH_webproxy.md
source_anchor: ""
source_lines: [56, 118]
section: "Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways"
sha256: 971f670515fc1d812a12fe3946a9018bf5ed768e6030497dc5e5b37808c1124f
---

# 3. nginx

## 3. nginx

### 3.1 Release status (2026)

- nginx release cadence in 2026 per endoflife.date tracking of nginx.org CHANGES: mainline 1.31 cycle released 2026-05-13, latest 1.31.4 (2026-08-19); stable 1.30 cycle released 2026-04-14, latest 1.30.4 (2026-07-15); 1.29 cycle 2025-06-24 → EOL 2026-05-13 (latest 1.29.8, 2026-04-07); 1.28 cycle 2025-04-23 → EOL 2026-04-14 (latest 1.28.3, 2026-03-24) [secondary — endoflife.date mirror of official changelog]. URL: https://github.com/maxfalstein/endoflife.date/blob/HEAD/products/nginx.md
- Community package-compatibility matrix (2026-09) lists stable 1.30.4 and 1.28.3 with deb/rpm packages for debian12 and almalinux9, amd64/arm64, glibc/musl [secondary].
- Container images: `nginx:1.27-alpine`, `nginx:alpine` remain common defaults in benchmark harnesses; homelab images (nfrastack/tiredofit) moved from 1.27/1.28 to 1.29.x by early 2026 [secondary].
- Security: CVE-2026-42945 ("NGINX Rift") — heap buffer overflow in the rewrite module reachable via `rewrite` with `?` in replacement followed by `set` with capture references; confirmed in-the-wild exploitation reported late May 2026; fixed in 1.28.1 (stable) and 1.29.0 (mainline); NGINX Plus R37 includes the fix; NGINX Ingress Controller 5.5.0 patched; F5 security portal advisory K000160932 [secondary — researcher write-up citing DepthFirst Research, Akamai, Qualysec, The Hacker News, May 2026]. URL: https://github.com/gary23w/garrettstimpson.ca/blob/HEAD/_posts/2026-05-27-cve-2026-42945-nginx-rift-heap-overflow-rce.md
- CVE-2026-1642 — SSL upstream injection vulnerability; fixed in 1.28.2 (stable) and 1.29.5 (mainline); ingress-nginx v1.14.3 bundled nginx 1.27.1 (affected) as of the report [secondary — kubernetes/ingress-nginx issue #14528]. URL: https://github.com/kubernetes/ingress-nginx/issues/14528
- Blog claim (low-trust single source): "nginx 1.29.7" with MPTCP support, released March 2026 — the version number is inconsistent with the endoflife.date release train (1.29.x EOL'd at 1.29.8 in April 2026; MPTCP support in upstream nginx is not documented in the official CHANGES); treat MPTCP claims as [unverified].

### 3.2 F5 / NGINX business status

- F5 acquired NGINX in 2019; as of 2026 the commercial product line is marketed as "F5 NGINX" with NGINX Plus, NGINX WAF, NGINX Amplify, NGINX Controller / Instance Manager [vendor-reported].
- G2 pricing page (updated 2026-04-08): NGINX Plus Single Instance starts at $2,500/year; NGINX WAF $2,000/year; free trial available; 3 pricing editions starting at $2,000; rated 4.6/5 [secondary — G2, pricing supplied by vendor]. URL: https://www.g2.com/products/f5-nginx/pricing
- Reviewer guidance (G2, 2026): OSS nginx suits SMB/startups for LB/reverse proxy; NGINX Plus targets banking, telecom, retail, manufacturing with SLA/support/monitoring [secondary].
- nginx open-source development continues on nginx.org with regular mainline/stable releases through August 2026 [official — via endoflife.date mirror].
- Notable ecosystem: ingress-nginx (Kubernetes) remains widely deployed but lags upstream versions (bundled 1.27.1 vs 1.28.2+ fix trains) [secondary]; alternatives NGINX Gateway Fabric and F5 NGINX Ingress Controller scored in 2026 migration analyses [secondary — blog.none.at ingress-nginx replacement study, 2026-07-18]. URL: https://blog.none.at/pdf/2026/2026-07-18-ingress-nginx-candidates.pdf

### 3.3 Config patterns, caching, load balancing

- Load-balancing algorithms (OSS nginx): round_robin (default), least_conn, ip_hash; NGINX Plus adds least_time [official — nginx docs].
- Active health checks are Plus-only; OSS has passive health checks (max_fails/fail_timeout) [official — nginx docs; also noted in community comparison tables].
- Caching: `proxy_cache` with cache zones, `proxy_cache_valid`, `proxy_cache_use_stale`, `proxy_cache_lock`; `fastcgi_cache`/`uwsgi_cache` for app backends; cache purging is Plus-only (or via third-party module in OSS) [official — nginx docs].
- Config style: declarative directive blocks (events/http/server/location/upstream); static config requiring reload (`nginx -s reload`, zero-downtime via new worker processes) [official].
- WebSocket proxying: `proxy_http_version 1.1` + `Upgrade`/`Connection` headers; gRPC via `grpc_pass` (HTTP/2 upstream) [official — nginx docs].
- Rate limiting: `limit_req_zone`/`limit_req` (leaky bucket), `limit_conn_zone`/`limit_conn` [official].
- TLS termination patterns: `ssl_certificate`/`ssl_certificate_key`, `ssl_protocols TLSv1.2 TLSv1.3`, `ssl_ciphers`, `ssl_prefer_server_ciphers`, OCSP stapling, session tickets; HTTP/3 requires the QUIC build (`listen 443 quic reuseport`) — QUIC support still ships as experimental in mainline builds as of 2026 [official — nginx docs; secondary for packaging variance].

---

## 4. LiteSpeed (secondary focus)

- LiteSpeed Web Server is a commercial drop-in replacement for Apache: reads Apache config files, honors .htaccess, integrates with cPanel/Plesk/DirectAdmin [secondary — medium.com guide, 2026-09].
- Architecture: event-driven worker pool handling thousands of concurrent connections vs Apache process/thread-per-connection [secondary].
- W3Techs September 2026: 14.6–14.7% of known sites; growth built on hosting providers benchmarking against Apache and switching [official share; secondary for the driver].
- OpenLiteSpeed is the open-source edition (GPLv3) with feature parity gaps vs Enterprise (notably .htaccess in some contexts, ESI cache, QUIC/HTTP/3 maturity historically) [secondary].
- QUIC/HTTP/3: LiteSpeed (lsquic library) was among the earliest production HTTP/3 stacks, widely deployed in the WordPress hosting segment [secondary].
- Pricing: LiteSpeed Enterprise is licensed per-worker/per-domain tiers via LiteSpeed Technologies; exact 2026 list pricing was not captured in this research round [gap — needs vendor price page check].

---

## 5. Caddy

### 5.1 Release status (2026)

- Caddy v2.11.4 observed in production homelab deployments as of 2026-08-30 (upgraded from 2.11.2/2.11.3); requires Go >= 1.25.1; verified build environment Go 1.26.7 with xcaddy v0.4.7 per one audited deployment [secondary]. URLs: https://github.com/pironex9/homelab/blob/HEAD/docs/hosts/caddy.md, https://github.com/satoshiportal/recoverbull-server/blob/HEAD/deploy/caddy/README.md
- Caddy 2.10 (June 2026) headline features per release notes: Encrypted ClientHello (ECH) support; post-quantum key exchange x25519mlkem768 enabled by default; ACME profiles support (experimental draft; e.g. Let's Encrypt 6-day certificate profiles); reverse proxy now sets `Via` header instead of duplicate `Server` header; global DNS provider option (`dns` global option for ACME DNS challenges and ECH); wildcards used by default for subdomains when present [secondary — releasebot.io summary of official release notes]. URL: https://releasebot.io/updates/caddy
- W3Techs September 2026: Caddy at 1.0% of known sites [official].
- Project model: single static Go binary, JSON or Caddyfile config, automatic HTTPS via CertMagic (ACME, default ZeroSSL with Let's Encrypt fallback historically) [official — caddyserver.com docs; widely documented].
- Ecosystem: caddy-dns providers (e.g. caddy-dns-cloudflare v0.2.15, 2026-07-06, tracking Caddy 2.11.x) enable DNS-01 challenges for wildcard certs [secondary]; xcaddy build tool for custom module builds (cache-handler, caddy-ratelimit, storage backends) [secondary].
- Known deployment profile: single LXC/VM, 1 vCPU, 256 MB RAM, reverse proxy for all LAN services — typical homelab edge [secondary — homelab docs, 2026].
- K8s ingress: caddy K8s ingress controller & Helm chart (caddy-dns-cloudflare repo v0.2.14, 2026-06-06) [secondary].
- caddy-proxy-manager v1.11.3 (2026-09-15) — third-party management UI project [secondary].

### 5.2 Caddyfile reverse-proxy essentials

- Minimal reverse proxy: `reverse_proxy localhost:8080` inside a site block; automatic HTTPS on public domains with no TLS config [official — Caddy docs].
- Common directives: `reverse_proxy` (load balancing policies: random, least_conn, round_robin, ip_hash, uri_hash, header), `handle`/`handle_path` routing, `header` manipulation, `basicauth`, `rate_limit` (third-party module), `tls` (internal CA, on_demand, dns), `encode` (gzip/zstd) [official].
- Layer-4 proxying available via caddy-l4 third-party module (not in core) [secondary].

---

