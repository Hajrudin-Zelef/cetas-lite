---
id: etape7-phaseh-webproxy/00-webproxy/overview
title: "Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways"
domain: step-7-phase-h-web-servers-reverse-proxies-api-gateways
role: deep-dive
task: reference
actors: ["Google", "Microsoft"]
dates: ["2011-11-10", "2017-07", "2024-12", "2026-01", "2026-06-08", "2026-08", "2026-09", "2026-09-04", "2026-09-17", "2026-09-22"]
keywords: ["apache", "benchmarks", "distribution", "memory", "packaging", "research", "training"]
source: docs/RAG/etape7_phaseH_webproxy.md
source_anchor: ""
source_lines: [1, 55]
section: "Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways"
sha256: 9edd3673b3e5178f3950dad6ee0df3322c246c4d2cf4d8d83f1acb84c7ee1839
---

# Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways

**Research cut-off date:** 2026-09-22
**Method:** Public web research (browser_search / browser_open), September 2026. Facts are tagged at sentence/clause level with one of: `[official]` (vendor or project documentation/releases), `[vendor-reported]` (vendor blog/press material not independently verified), `[independent]` (third-party benchmarks, analysts), `[secondary]` (community docs, homelab write-ups, wikis), `[unverified]` (single-source or unconfirmed claims).
**Provenance legend:** every factual claim below carries one of the five tags above; claims without a tag are framing prose. When sources disagree, both claims are kept and the conflict is flagged. Prices include the date observed. URLs are quoted verbatim from search results.
**Scope:** Apache httpd, nginx, Caddy, Traefik, HAProxy, Envoy, lighttpd, LiteSpeed, Nginx Proxy Manager; reverse-proxy playbooks (TLS termination, ACME/cert-manager, security headers, rate limiting, WebSocket/gRPC, SSO integration); API gateways (Kong, APISIX, Tyk, service discovery, plugins, rate limiting, OpenAPI); performance benchmarks; homelab vs enterprise patterns. Observation window: 2026-09-22.

---

## 1. Market landscape (September 2026)

- W3Techs "Usage Statistics and Market Share of Web Servers, September 2026" (updated daily; data dated 17 September 2026): nginx 30.9%, Cloudflare Server 30.7%, Apache 22.0%, LiteSpeed 14.6%, Node.js 7.2%, Microsoft-IIS 3.0%, Caddy 1.0%, Google Servers 1.0%, Envoy 0.5%, Railway edge proxy 0.1%, Tengine 0.1%, IdeaWebServer 0.1%, ArvanNginx 0.1%, Kestrel 0.1% [official — W3Techs, 2026-09-17]. URL: https://w3techs.com/technologies/overview/web_server?ref=bojana.dev
- W3Techs notes that a website may use more than one web server, so shares do not sum to 100% [official].
- W3Techs comparison page dated 4 September 2026 (nginx vs Apache vs LiteSpeed): nginx 31.3%, Apache 22.5%, LiteSpeed 14.7%; broken down by rank — nginx 27.7% top-1M, 27.3% top-100k, 29.7% top-10k, 30.9% top-1k; Apache 17.1% top-1M, 14.4% top-100k, 11.8% top-10k, 10.0% top-1k; LiteSpeed 8.9% top-1M, 6.2% top-100k, 4.8% top-10k, 1.5% top-1k [official — W3Techs, 2026-09-04]. URL: https://w3techs.com/technologies/comparison/ws-apache,ws-litespeed,ws-nginx
- Interpretation pattern across W3Techs rank slices: nginx share rises with site rank (27.3%→30.9% from top-1M to top-1k), Apache share falls (17.1%→10.0%), LiteSpeed falls (8.9%→1.5%) [official — W3Techs, 2026-09-04; analysis of the published slices].
- Secondary analysis (programming-helper.com, January 2026, citing W3Techs/Netcraft): Netcraft December 2024 survey had nginx 19.67% and Apache 17.30%; by January 2026 W3Techs showed nginx ~33.0% vs Apache ~24.2% [secondary — Netcraft data via third-party article; Netcraft December 2024 figures could not be re-verified from Netcraft directly in this research round].
- Caveat: W3Techs detects the server that answers (or is advertised) at the edge; sites behind Cloudflare report "Cloudflare Server" even when the origin runs nginx/Apache, so the Cloudflare 30.7% slice overlaps with origin-server shares [secondary; methodological caution, consistent with W3Techs methodology notes].
- LiteSpeed holds 14.6–14.7% overall, almost entirely hosting-market driven (shared hosting, WordPress, cPanel/Plesk/DirectAdmin ecosystems) [official — W3Techs share; secondary for the hosting-market characterization].
- Caddy: 1.0% of known sites (W3Techs, 2026-09-17) [official]; small but stable footprint concentrated in homelab/self-hosted and developer-tooling deployments [secondary].
- Envoy: 0.5% of known sites (W3Techs, 2026-09-17) [official]; its footprint is disproportionately in Kubernetes ingress and service mesh, under-represented in edge-server surveys [secondary].

---

## 2. Apache httpd

### 2.1 Release status (2026)

- Apache httpd 2.4.68 released 2026-06-08 — latest release from the 2.4.x stable branch, described by the project as "the best available version of Apache HTTP Server" [official]. URL: https://httpd.apache.org/?ref=tryootech.com
- Download page lists 2.4.68 (released 2026-06-08) as the current stable release; distribution files `httpd-2.4.68.tar.gz` (9.6M) and `httpd-2.4.68.tar.bz2` (7.4M) mirrored with `CURRENT-IS-2.4.68` marker [official]. Mirror index: http://ftp.jaist.ac.jp/pub/apache/httpd/
- Only the latest patch release of the 2.4.x branch is security-supported; vulnerabilities that exist only in unreleased branches (e.g. trunk) are handled as normal bug reports [official — SECURITY.md]. URL: https://github.com/apache/httpd/blob/HEAD/SECURITY.md
- Apache httpd 2.2 is end-of-life; final release 2.2.34 (July 2017); no further patches [official].
- trunk is the development line toward 2.6 (labeled 2.5.x/2.6): as of SVN r1935140 (2026-06-08) it is `2.5.1-dev` with MMN `20211221:31`, vs 2.4.x `2.4.68-dev` with MMN `20120211:142`; branches diverged at r1179239 (2011-11-10); ~12,560 trunk-only commits vs ~11,446 2.4.x-only commits [secondary — community trunk-vs-2.4 analysis, 2026-06-08]. URL: https://github.com/apache/httpd/blob/HEAD/httpd-trunk.md
- Trunk-only analysis found 730 Apache directives in trunk vs 637 in 2.4.x (count at r1935140); the majority of 2.4.x changes are back-ports of trunk work [secondary].
- Distro packaging example (August 2026): AlmaLinux EL10 appstream ships httpd 2.4.63-13.el10_2.5 with APR 1.7.5, mod_http2 2.0.29, mod_lua bundled [secondary — community Linux training repo, 2026-08].
- openEuler docker images track httpd tags 2.4.51 through 2.4.66 on various openEuler LTS releases [secondary].

### 2.2 MPMs, architecture, modules

- Apache's headline features per the project: flexible configuration (.htaccess, virtual hosts, DSO module loading), TLS/auth/access control security, performance via MPMs/HTTP/2/caching/reverse proxy, 100+ modules (rewrite, proxy, load balancing, scripting), portability (Linux, Windows, macOS, Unix) [official].
- MPMs: `event` (default and recommended for 2.4), `worker`, `prefork` (required for non-thread-safe modules such as mod_php), plus `mpm_netware`/`mpmt_os2` legacy [official — httpd docs; standard documented fact].
- `mpm_event` is the thread-per-connection hybrid model suitable for keep-alive heavy workloads; `prefork` is process-per-connection and the safe choice for mod_php [official — httpd docs].
- mod_proxy family: mod_proxy, mod_proxy_http, mod_proxy_ajp, mod_proxy_balancer (balancer algorithms: byrequests, bytraffic, bybusyness, heartbeat), mod_proxy_hcheck (health checks), mod_proxy_fcgi, mod_proxy_wstunnel (WebSocket), mod_proxy_http2, mod_proxy_uwsgi [official — httpd docs].
- mod_rewrite remains the de-facto URL manipulation engine; mod_ssl (OpenSSL), mod_http2, mod_cache/mod_cache_disk/mod_socache_*, mod_security is third-party (SpiderLabs/Trustwave, now community) [official for bundled modules; secondary for ModSecurity maintenance status].
- TLS 1.3 requires Apache 2.4.43+ with OpenSSL 1.1.1+ [official].
- `.htaccess` vs central config: per-directory .htaccess allows unprivileged users to override config but costs a filesystem lookup per request; production guidance is to disable `AllowOverride` and move rules into the main server config for performance [official — httpd htaccess docs; widely accepted operational fact].
- Hardening checklist (community/OWASP-aligned): disable server tokens (`ServerTokens Prod`, `ServerSignature Off`), disable unused modules, run dedicated user/group, TLS 1.2+ only with strong ciphers, `TraceEnable Off`, limit request sizes (`LimitRequestBody`), use mod_security CRS where applicable, keep 2.4.x patched to latest [secondary].

### 2.3 Apache vs nginx positioning

- Apache remains default on shared hosting and legacy environments; nginx is the default for reverse proxy, load balancing and high-concurrency serving [secondary — medium.com comparison, 2026-01].
- Architecture contrast: Apache process/thread-per-connection model is flexible but more memory-hungry under load; nginx event-driven async model handles thousands of connections per worker with low memory [secondary].
- Apache's modular ecosystem and .htaccess remain advantages in multi-tenant hosting; nginx's static-config + reload model and lower resource use dominate high-traffic, container and edge deployments [secondary].

---

