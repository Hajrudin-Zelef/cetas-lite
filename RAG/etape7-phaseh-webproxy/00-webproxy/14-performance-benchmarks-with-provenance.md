---
id: etape7-phaseh-webproxy/00-webproxy/14-performance-benchmarks-with-provenance
title: "14. Performance benchmarks (with provenance)"
domain: step-7-phase-h-web-servers-reverse-proxies-api-gateways
role: deep-dive
task: benchmark
actors: ["AWS"]
dates: ["2026-07-18", "2026-09-22"]
keywords: ["benchmark", "benchmarks", "agents", "apache", "aws", "cost", "governance", "mcp", "memory", "model context protocol", "pricing", "throughput"]
source: docs/RAG/etape7_phaseH_webproxy.md
source_anchor: ""
source_lines: [307, 363]
section: "Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways"
sha256: 484ad16f7b33794a17d2c4bf10ccf7ff571fbf16a8a4f0232dae159fede91912
---

# 14. Performance benchmarks (with provenance)

## 14. Performance benchmarks (with provenance)

> All figures below are indicative and workload-specific; treat as directional, not absolute rankings.

- Community benchmark (kanywst/reverse-proxy-benchmark, MacBook Pro M1 Max, "hello world" workload): nginx 54,826 req/s @ 80% CPU / 11.97 MB RSS; Envoy 37,637 req/s @ 383% CPU / 43.30 MB; Pingora 30,991 req/s @ 77% CPU / 19.60 MB; Traefik 8,603 req/s @ ~60.96 MB (Go runtime overhead; CPU unmeasurable in that harness) [independent — community benchmark repo; methodology documented in-repo]. Conclusion stated by author: nginx most efficient per-request, Envoy scales across cores at higher resource cost, Pingora balanced, Traefik trades throughput for usability [secondary].
- Community proxy benchmark harness (praxis-proxy/praxis, 2026-09): 8 workloads × proxies (Envoy v1.31, nginx:alpine, haproxy:3.1/3.x, Praxis) with runs/warmup/duration documented; framework exists, per-run result numbers are repo-local [secondary — methodology only]. URL: https://github.com/praxis-proxy/praxis/blob/HEAD/docs/benchmarks.md
- Vendor-adjacent comparison (index.dev PDF, undated, older crawl): Envoy 100k+ req/s/core, nginx 10k–50k+ req/s/instance, HAProxy 40k+ concurrent connections; binary sizes Envoy 80–120 MB, nginx 1–2 MB, HAProxy 3–4 MB; base memory Envoy 10–50 MB, nginx 10–50 MB, HAProxy 2–4 MB [secondary — marketing PDF via CloudFront; low trust, figures undated].
- Feature comparison (community DevOps notes, 2026-09): NGINX = web server + reverse proxy, OSS active health checks absent, static config; HAProxy = dedicated LB, active checks, precise timeout/queue control; Envoy = dynamic xDS, circuit breaking, richest observability, heaviest [secondary]. URL: https://github.com/aadarsh-nagrath/devops-cloud/blob/HEAD/Load%20Balancing/software-load-balancers.md
- nginx E2E baseline (cnkang/nginx-markdown-for-agents, NGINX 1.28.2 stable, loopback, `ab`): passthrough 44–47k req/s small/medium payloads; module-path overhead documented per scenario [secondary — single-project baseline, not a cross-proxy comparison]. URL: https://github.com/cnkang/nginx-markdown-for-agents/blob/HEAD/docs/testing/PERFORMANCE_BASELINES.md
- HAProxy vs nginx streaming behavior (dev.to, 2026-09-22): nginx `proxy_buffering on` costs ~50 ms only in adversarial slow-client case; HAProxy kernel batching costs ~200 ms only for small-frame streams; both one-line tunables [independent — careful measurement with published configs]. URL: https://dev.to/rlnorthcutt/haproxys-200-ms-llm-token-delay-is-a-smart-default-heres-when-to-change-it-3eah
- Non-comparability warning: benchmarks across different machines, payloads, TLS settings and concurrency levels are not directly comparable; prefer same-harness numbers [independent — methodological note].

---

## 15. Homelab vs enterprise patterns

- Homelab (2026, observed across homelab repos): Caddy or NPM on a 1 vCPU/256 MB LXC as LAN edge; Traefik with Docker labels for container estates; Authelia forward-auth for SSO on 5–15 services; Let's Encrypt via built-in ACME; Cloudflare Tunnel for external ingress without port forwarding; Jellyfin/media behind the same proxy [secondary].
- Reference architectures seen: `Cloudflare edge → Tunnel → Caddy :80 → Traefik → app container` with `header_up Host {host}` preservation [secondary — homelab docs, 2026-08]. URL: https://github.com/duresa7/homelab/blob/HEAD/Platforms/Caddy/README.md
- Enterprise (2026): F5 NGINX Plus or HAProxy Enterprise/One at edge with WAF + bot management; cert-manager or Venafi for K8s PKI; Kong Konnect or APISIX for API governance; OIDC via Keycloak/Authentik/Entra ID; mTLS east-west; Gateway API (Envoy Gateway / HAProxy Unified Gateway / Traefik) replacing Ingress [secondary; vendor-reported for commercial components].
- ingress-nginx migration (2026-07 analysis): Envoy Gateway and F5 NGINX Ingress Controller scored highest on a 17-row feature matrix for ingress-nginx replacement; Gateway-API-only options (NGINX Gateway Fabric, HAProxy Unified Gateway) lag on snippets/rate-limit/external-auth — maturity gap, not design flaw [independent — blog.none.at study]. URL: https://blog.none.at/pdf/2026/2026-07-18-ingress-nginx-candidates.pdf
- Selection guidance: nginx/Caddy for simplicity and static efficiency; HAProxy when LB precision and TCP matter; Traefik for Docker/K8s auto-discovery; Envoy for mesh/dynamic; Kong/APISIX when API governance (portals, monetization, AI gateway) is required [secondary — synthesis of community guidance].

---

## 16. Gaps, conflicts & unverified claims

- [conflict] nginx "1.29.7 with MPTCP" blog claim vs official release train (1.29.x EOL at 1.29.8, no MPTCP in CHANGES): the blog appears to describe a hypothetical/custom build; do not cite MPTCP as an nginx feature.
- [conflict] Kong pricing: legacy Konnect consumption pricing (~$105/service/month + $34.25/1M requests, from 2023–2025 PDF) vs 2026 modular Plus pricing ($25/$200/$500 control planes + $200/additional 1M requests); both circulate in 2026 analyses — verify against the live pricing page before quoting.
- [gap] LiteSpeed Enterprise 2026 list pricing not captured.
- [gap] Tyk 2026 version and commercial pricing not captured.
- [gap] Nginx Proxy Manager current version/release status not captured.
- [gap] lighttpd 2026 release numbers not captured.
- [gap] OAuth2-Proxy 2026 minor version not captured.
- [gap] Envoy 2026 release train numbers not captured in this round.
- [gap] NGINX Plus R37+ release notes beyond the CVE-2026-42945 fix mention not reviewed.
- [unverified] AWS F1/FPGA claims from earlier phases are out of scope here; no new data.
- [method] W3Techs detects edge-advertised servers; Cloudflare-fronted origins are counted under Cloudflare, understating nginx/Apache origin share.

---

## 17. Glossary

- **Reverse proxy:** server that sits in front of backends, terminating client connections (TLS, routing, caching) and forwarding to upstreams.
- **ACME:** Automated Certificate Management Environment (RFC 8555); protocol used by Let's Encrypt et al.
- **MPM:** Multi-Processing Module (Apache); connection-handling model (event/worker/prefork).
- **SNI:** Server Name Indication; TLS extension selecting the certificate by hostname.
- **ALPN:** Application-Layer Protocol Negotiation; negotiates h2/http1.1/h3 during TLS handshake.
- **OCSP stapling:** server attaches certificate revocation status to the handshake.
- **ECH:** Encrypted ClientHello; encrypts the SNI/hostname in the TLS handshake.
- **PQC:** post-quantum cryptography (e.g. ML-KEM key exchange).
- **Forward auth:** proxy delegates the auth decision to an external service via subrequest (auth_request / forwardAuth / ext_authz).
- **xDS:** Envoy's dynamic configuration APIs (LDS/CDS/EDS/RDS).
- **Gateway API:** Kubernetes SIG-Network standard (successor to Ingress) for L4–L7 routing.
- **MCP:** Model Context Protocol; 2026 API-gateway proxying target (Kong AI MCP Proxy, APISIX mcp-bridge).

---

