---
id: etape7-phaseh-webproxy/00-webproxy/part-4
title: "Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways (part 4)"
domain: step-7-phase-h-web-servers-reverse-proxies-api-gateways
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape7_phaseH_webproxy.md
source_anchor: ""
source_lines: [113, 118]
section: "Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways"
sha256: 67e8bd9fa06f158898f4d52ba417b7cd34b18d33e747a3deb1ddc3981760a464
---

# Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways (part 4)

- Minimal reverse proxy: `reverse_proxy localhost:8080` inside a site block; automatic HTTPS on public domains with no TLS config [official — Caddy docs].
- Common directives: `reverse_proxy` (load balancing policies: random, least_conn, round_robin, ip_hash, uri_hash, header), `handle`/`handle_path` routing, `header` manipulation, `basicauth`, `rate_limit` (third-party module), `tls` (internal CA, on_demand, dns), `encode` (gzip/zstd) [official].
- Layer-4 proxying available via caddy-l4 third-party module (not in core) [secondary].

---

