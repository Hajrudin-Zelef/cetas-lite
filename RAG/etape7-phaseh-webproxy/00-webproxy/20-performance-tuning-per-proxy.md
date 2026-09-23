---
id: etape7-phaseh-webproxy/00-webproxy/20-performance-tuning-per-proxy
title: "20. Performance tuning per proxy"
domain: step-7-phase-h-web-servers-reverse-proxies-api-gateways
role: deep-dive
task: reference
actors: []
dates: ["2026-05-12"]
keywords: ["apache", "backlog", "cost", "latency", "mcp"]
source: docs/RAG/etape7_phaseH_webproxy.md
source_anchor: ""
source_lines: [471, 520]
section: "Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways"
sha256: 83074da8285ab245cc283e6b4b2134dce0b7299c80fe52ed06567acd9fb45506
---

# 20. Performance tuning per proxy

## 20. Performance tuning per proxy

- Linux baseline for any proxy: `net.core.somaxconn=4096+`, `net.ipv4.tcp_max_syn_backlog`, `fs.file-max` / `ulimit -n` (≥ 65535), `net.ipv4.ip_local_port_range`, disable `net.ipv4.tcp_tw_recycle` (removed in modern kernels; use `tcp_tw_reuse` where appropriate) [secondary — standard SRE guidance].
- nginx: `worker_processes auto; worker_connections 4096; worker_rlimit_nofile 65535;` `multi_accept on;` `sendfile on; tcp_nopush on;` `keepalive_timeout 65; keepalive_requests 1000;` `open_file_cache` for static [official].
- HAProxy: `nbthread` = CPU cores; `maxconn` sized to RAM (~tens of KB per connection); `tune.bufsize`, `tune.maxrewrite`; prefer `mode http` with `option http-keep-alive` over `http-tunnel` [official].
- Apache event MPM: `StartServers`, `MinSpareThreads/MaxSpareThreads`, `ThreadsPerChild 25`, `MaxRequestWorkers 400+` sized to RAM (each thread ~MBs); `KeepAlive On`, `MaxKeepAliveRequests 100` [official].
- Caddy/Traefik (Go): `GOMAXPROCS` auto; tune via OS limits; Go GC is the main latency lever (`GOGC`) under extreme load [secondary].
- TLS cost: ECDSA (P-256) certs are cheaper per handshake than RSA-2048; session resumption (tickets or shared cache) cuts handshakes; TLS 1.3 1-RTT/0-RTT [official docs; secondary].
- Benchmarking tools: `wrk`, `wrk2` (latency-accurate), `hey`, `vegeta`, `k6`; always test with TLS on when comparing edge configs, warm up, and report p50/p99 not just RPS [secondary].

---

## 21. Observability per proxy

- nginx: access/error logs; `stub_status` module; NGINX Plus API/dashboard; Prometheus via `nginx-prometheus-exporter` (nginxinc) or VTS module [official; secondary].
- HAProxy: stats page/socket (CSV + HTML), `option httplog`, Prometheus exporter built into 2.x+ (`http-request use-service prometheus-exporter if { path /metrics }`), structured log formats [official].
- Traefik: access logs (CLF/JSON/OTLP), metrics (Prometheus/Datadog/StatsD/InfluxDB), tracing (OTel/Jaeger/Zipkin), dashboard [official].
- Caddy: JSON access logs, `debug` global option, Prometheus via `caddy-prometheus` module or native metrics endpoint (v2.6+) [official; secondary].
- Apache: `LogFormat` combined/JSON, `mod_status` (server-status), `mod_log_forensic` [official].
- Key metrics to alert on: 5xx rate, p99 latency, upstream error rate, TLS handshake failures, cert expiry days, connection saturation [secondary — SRE consensus].

---

## 22. High availability & edge patterns

- Active-passive proxy HA: keepalived/VRRP floating VIP between two proxy nodes; or BGP anycast for multi-site [secondary].
- Active-active: DNS round-robin (cheap, uneven), ECMP/BGP anycast (best), or cloud LB in front [secondary].
- Let's Encrypt in HA: shared storage for acme.json/certs (NFS/object) or DNS-01 with centralized issuance; avoid duplicate HTTP-01 issuance (rate limits) [secondary].
- Let's Encrypt rate limits (2026): 50 certificates per registered domain per week; 5 duplicate certificates per week; 300 new orders per account per 3 hours; 10 accounts per IP per 3 hours [official — letsencrypt.org rate limits docs; standard documented fact].
- Blue-green/canary at proxy layer: nginx `split_clients`, Traefik weighted services, HAProxy `weight`, Kong/APISIX traffic-split plugins [official].
- Health-check-driven failover: HAProxy `backup` servers, nginx `backup` parameter, Traefik healthcheck failover [official].

---

## 23. OpenResty (Kong/APISIX substrate)

- OpenResty: nginx + LuaJIT platform; Kong Gateway and Apache APISIX are both built on OpenResty/nginx cores, which is why they inherit nginx-class request performance with Lua programmability [official — openresty.org; secondary].
- APISIX 3.13 corrected its OpenResty version reference (1.27.1.1) in changelog fixes [secondary — apisix changelog PR, 2026-05-12].
- ingress-nginx also uses OpenResty, which constrains how fast it can track upstream nginx releases (compatibility with the OpenResty fork) [secondary — kubernetes/ingress-nginx issue #14528].
- Lua plugin hot-reload and `lua-resty-*` libraries (redis, mysql, http, jwt) are the extension mechanism shared by Kong/APISIX [official docs].

## 24. Kong plugin catalog (2026)

- Auth: key-auth, basic-auth, jwt, oauth2, ldap-auth, hmac-auth, mtls-auth, openid-connect (Enterprise), vault-auth [official — Kong plugin hub].
- Traffic: rate-limiting, response-ratelimiting, request-size-limiting, proxy-cache, request-transformer, response-transformer, correlation-id, redirect [official].
- AI (2026): ai-proxy (LLM provider routing), ai-rate-limiting, ai-prompt-guard, ai-prompt-decorator, ai-request-transformer, ai-response-transformer, ai-mcp-proxy (3.12+, Enterprise) [official — Kong docs; secondary for Enterprise gating].
- Observability: prometheus, opentelemetry, datadog, statsd, zipkin, file-log, http-log, tcp-log [official].
- decK: declarative YAML (`deck sync`) for GitOps management of Kong config [official].
- Konnect: SaaS control plane; data planes self-managed or Kong-managed cloud; service catalog, mesh manager, developer portal [official].

