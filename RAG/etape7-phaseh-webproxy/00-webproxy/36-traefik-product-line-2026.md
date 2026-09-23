---
id: etape7-phaseh-webproxy/00-webproxy/36-traefik-product-line-2026
title: "36. Traefik product line (2026)"
domain: step-7-phase-h-web-servers-reverse-proxies-api-gateways
role: deep-dive
task: reference
actors: ["AWS"]
dates: []
keywords: ["apache", "aws", "distribution", "governance", "license", "mcp"]
source: docs/RAG/etape7_phaseH_webproxy.md
source_anchor: ""
source_lines: [617, 661]
section: "Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways"
sha256: 531d852076cc6e013eeb1c1a0032cf14695b31e52210c86ad348086266a5de15
---

# 36. Traefik product line (2026)

## 36. Traefik product line (2026)

- Traefik Proxy (OSS, MIT): the core reverse proxy/LB described in §6 [official].
- Traefik Hub (SaaS): API gateway/management layer on top of Proxy v3.7 — APIAuth, token rate limit & quota, MCP registry (tech preview), Gateway API v1.5.1 [secondary — official Hub docs].
- Traefik Enterprise (historical): commercial distribution with distributed rate limiting, OIDC, advanced clustering — superseded in messaging by Hub as of 2024–2026 [secondary].
- Version support policy updated from v3.6 (2026) — check the official support matrix before pinning [secondary].

## 37. Caddy module ecosystem (xcaddy, 2026)

- Build: `xcaddy build --with github.com/caddyserver/transform-encoder ...`; reproducible builds pin module commits (audited builds pin e.g. caddy-ratelimit@5625512, cache-handler@v0.16.0) [secondary].
- DNS providers: ~70+ modules (cloudflare, route53, digitalocean, gandi, ovh, hetzner, porkbun…) for DNS-01 [secondary].
- Storage: Redis, Consul, etcd, S3, Vault, PostgreSQL/MySQL/MongoDB cert-storage backends for clustered deployments [secondary].
- L4: caddy-l4 (layer-4 proxying, not in core) [secondary].
- Auth: caddy-security (OAuth2/OIDC portal), forward_auth with Authelia [secondary].
- Note: every extra module increases binary size and supply-chain surface; pin versions in production [secondary].

## 38. API gateway comparison matrix (2026)

- Language/core: Kong — Lua/OpenResty (nginx); APISIX — Lua/OpenResty (nginx); Tyk — Go; Traefik Hub — Go (Traefik Proxy); KrakenD — Go (Lura); Gloo — Go (Envoy) [official].
- License: Kong Gateway OSS Apache-2.0; APISIX Apache-2.0; Tyk Gateway MPL-2.0/BSL mix (verify per version); KrakenD Apache-2.0 [official; secondary].
- Config store: Kong — Postgres/DB-less YAML; APISIX — etcd; Tyk — Redis; Traefik — providers [official].
- Control plane SaaS: Kong Konnect (Plus $25/$200/$500/mo + usage, 2026); Tyk Cloud (quote); API7 Cloud for APISIX [secondary].
- Developer portal: Kong Konnect (paid), Tyk (commercial), APISIX (dashboard, OSS) [official; secondary].
- AI/MCP 2026: Kong AI MCP Proxy (Enterprise), APISIX mcp-bridge (OSS 3.13+) [secondary].
- Best fit: Kong — enterprise governance + AI gateway budget; APISIX — OSS-first, etcd shops; Tyk — Go shops, simpler ops; KrakenD — ultra-light aggregation [secondary].

## 39. Protocol support matrix (2026)

- HTTP/1.1: all (nginx, Apache, Caddy, Traefik, HAProxy, Envoy, lighttpd) [official].
- HTTP/2 (h2): all; Apache via mod_http2; nginx `http2` directive; HAProxy ALPN; Caddy/Traefik automatic [official].
- HTTP/3/QUIC: nginx (experimental QUIC build), Caddy (quic-go, mature), HAProxy 3.2 (OpenSSL 3.5/quictls/WolfSSL/AWS-LC), LiteSpeed (lsquic, mature), Apache (not in 2.4.x core — gap), Traefik (Go quic-go) [official; secondary].
- WebSocket: all L7 proxies [official].
- gRPC: nginx `grpc_pass`, Traefik native, Caddy h2c transport, HAProxy mode http + ALPN h2, Envoy first-class [official].
- TCP/UDP passthrough: HAProxy (native), nginx stream, Caddy (caddy-l4 module), Traefik TCP routers, Envoy listeners [official].
- ECH: Caddy 2.10+ [secondary]; others — not confirmed in core as of 2026 (gap).

## 40. TLS feature matrix (2026)

- Auto-ACME: Caddy (built-in), Traefik (built-in resolvers), Apache (mod_md), HAProxy (Data Plane API 3.2), nginx (certbot/external or njs/acme modules) [official].
- Wildcard certs: via DNS-01 on all with ACME support [official].
- OCSP stapling: nginx, Apache, HAProxy, Caddy [official].
- mTLS client auth: nginx, Apache, HAProxy, Traefik, Caddy (`client_auth`) [official].
- PQC key exchange: Caddy 2.10 default x25519mlkem768 [secondary]; nginx/HAProxy depend on OpenSSL 3.5+ build [secondary].
- TLS 1.3: all current releases [official].

