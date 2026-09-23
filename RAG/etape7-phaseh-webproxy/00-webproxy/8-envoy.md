---
id: etape7-phaseh-webproxy/00-webproxy/8-envoy
title: "8. Envoy"
domain: step-7-phase-h-web-servers-reverse-proxies-api-gateways
role: deep-dive
task: reference
actors: []
dates: ["2026-02-10", "2026-02-25", "2026-07-18", "2026-09", "2026-09-17"]
keywords: ["apache", "benchmark", "cost", "latency", "memory"]
source: docs/RAG/etape7_phaseH_webproxy.md
source_anchor: ""
source_lines: [166, 243]
section: "Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways"
sha256: c68cce4ce4b0905f182dcb8620e7a7f088bee8b5f65230ebf90df44b4e0bdd5d
---

# 8. Envoy

## 8. Envoy

- Envoy is a cloud-native L4/L7 edge and service proxy (C++), the data plane of Istio/Linkerd-class meshes and of Envoy Gateway [official — envoyproxy.io; standard fact].
- 0.5% of known sites per W3Techs (2026-09-17) — edge-survey undercounts its K8s-internal footprint [official share; secondary interpretation].
- Strengths: xDS dynamic configuration APIs, sophisticated L7 (circuit breaking, outlier detection, retries, shadowing), rich observability (Prometheus stats, distributed tracing) [official; secondary comparison].
- Trade-offs: heavyweight — ~80–120 MB statically-linked binary, 150–200 MB Docker images, 45–90 min clean build; higher memory/CPU vs nginx/HAProxy for equivalent static routing [secondary — community benchmark comparison PDF; treat exact numbers as indicative, not lab-grade].
- Envoy Gateway is the Kubernetes Gateway API implementation; scored well in a 2026 ingress-nginx replacement analysis (9 green / 8 yellow / 0 red on a 17-row feature matrix) [secondary — blog.none.at, 2026-07-18].
- Common pattern: cloud L7 LB at edge → Envoy sidecar mesh internally → HAProxy for legacy TCP services [secondary — community architecture guidance].

---

## 9. lighttpd (secondary)

- lighttpd 1.4.x remains maintained; 1.5 development branch long-running [secondary; exact 2026 release numbers not captured in this round — gap].
- Niche: very low memory footprint, mod_proxy/mod_fastcgi/mod_magnet (Lua), historically strong in embedded and high-connection static serving [secondary].
- Not in W3Techs top-10 as of September 2026 (below 0.1% threshold of listed servers) [official — W3Techs table].

---

## 10. Nginx Proxy Manager (homelab)

- Nginx Proxy Manager (NPM) is the popular homelab GUI for nginx reverse proxying with built-in Let's Encrypt management [secondary — widely documented].
- Typical split in homelab designs (2026): NPM handles internal services behind the LAN; Caddy or Traefik handles external edge (e.g. behind Cloudflare Tunnel) [secondary — homelab architecture docs].
- No specific 2026 release/version data captured in this round; treat current version as [gap — verify at nginxproxymanager.com / GitHub].

---

## 11. Reverse-proxy playbooks

### 11.1 TLS termination

- Standard pattern: terminate TLS at the edge proxy (nginx/HAProxy/Traefik/Caddy), forward plain HTTP or re-encrypted HTTPS to backends; backend re-encryption (`proxy_pass https://`) protects east-west traffic at CPU cost [official docs; secondary for cost guidance].
- Cipher guidance (2026): TLS 1.2 + 1.3 only; TLS 1.0/1.1 long deprecated; prefer AEAD suites (AES-GCM, ChaCha20-Poly1305); disable 3DES/RC4/compression [official — Mozilla SSL Configuration Generator guidance; secondary for 2026 consensus].
- OCSP stapling reduces handshake latency and privacy leakage; session resumption via tickets or session cache [official docs].
- Post-quantum: Caddy 2.10 enables x25519mlkem768 by default (2026-06) [secondary — release notes]; OpenSSL 3.5 ships native QUIC API but PQC on QUIC paths depends on library choice (quictls fork lacks OpenSSL 3.5's PQC key exchange) [secondary].
- Certificate lifecycle automation is the operational norm: ACME (Let's Encrypt) for public, internal CA (step-ca, smallstep, Vault PKI) or self-signed for LAN [secondary].

### 11.2 ACME, Let's Encrypt, cert-manager

- cert-manager is the de-facto Kubernetes certificate controller: watches Certificate/Issuer/ClusterIssuer CRs, provisions via ACME (Let's Encrypt), Vault, Venafi, private CA [secondary].
- 2026 versions observed: v1.19.4 (helm chart, deployed 2026-02-25), v1.19.3 (kURL/Replicated add-on, 2026-02-10), v1.19.2, v1.19.1 in active homelab/cluster use [secondary]. URLs: https://github.com/cert-manager/cert-manager/releases/, https://github.com/thorix/kubernetes-deploy/blob/HEAD/cert-manager/README.md
- v1.19.0 feature notes: ACME profiles extension support, IPv6 default network-policy rules, global nodeSelector in helm chart, `certmanager_certificate_challenge_status` Prometheus metric, CAInjectorMerging promoted to beta [secondary — GitHub release notes].
- Install pattern: `helm install cert-manager jetstack/cert-manager --namespace cert-manager --create-namespace --version v1.19.x --set crds.enabled=true` (or OCI: `oci://quay.io/jetstack/charts/cert-manager`) [secondary].
- Gateway API support: `config.enableGatewayAPI=true` helm value enables Gateway API certificate integration [secondary].
- Traefik native ACME: `certificatesResolvers.<name>.acme` with httpChallenge/tlsChallenge/dnsChallenge, storage in acme.json; DNS-01 required for wildcards [official — Traefik docs].
- Caddy native ACME: automatic on public domains; DNS-01 via dns provider modules; ECH and ACME profiles in 2.10+ [secondary].
- HAProxy Data Plane API 3.2: native `acme` section with crt-store and runtime renew/status; HTTP-01 challenge is Enterprise-only [official].
- Let's Encrypt 6-day certificate profiles (ACME profiles draft) are the direction for short-lived certs; Caddy 2.10 added ACME profiles support [secondary].

### 11.3 Security headers

- Baseline set: `Strict-Transport-Security` (HSTS, includeSubDomains, preload), `X-Content-Type-Options: nosniff`, `X-Frame-Options: DENY` or `frame-ancestors`, `Referrer-Policy`, `Permissions-Policy`, `Content-Security-Policy` [secondary — OWASP Secure Headers Project consensus].
- nginx: `add_header` / `more_set_headers` (headers-more module); Traefik: headers middleware (customRequestHeaders, customResponseHeaders, stsSeconds, contentSecurityPolicy…); Caddy: `header` directive; HAProxy: `http-response set-header` [official docs].
- Remove/replace `Server` tokens: nginx `server_tokens off`, Apache `ServerTokens Prod`; Caddy 2.10+ sets `Via` instead of duplicate `Server` [official; secondary].

### 11.4 Rate limiting & abuse controls

- nginx: `limit_req_zone $binary_remote_addr zone=one:10m rate=10r/s` + `limit_req zone=one burst=20 nodelay`; `limit_conn_zone`/`limit_conn` for concurrency [official].
- Traefik: rateLimit middleware (average/burst, sourceCriterion by IP or header) [official].
- HAProxy: stick-tables + `http-request track-sc0`/`deny` for global rate limiting; Enterprise adds bot management/TDE [official; vendor-reported for TDE].
- Caddy: core has no built-in rate limit; community module `mholt/caddy-ratelimit` used (pinned commit in audited 2026 builds) [secondary].
- Kong/APISIX/Tyk: plugin-based rate limiting (fixed window, sliding window, token bucket; Redis-backed for distributed) [official].

### 11.5 WebSocket, SSE, gRPC proxying

- WebSocket: nginx needs `proxy_http_version 1.1`, `Upgrade`/`Connection` header mapping, long `proxy_read_timeout`; Traefik/Caddy handle upgrade transparently; HAProxy `mode http` with `timeout tunnel` [official docs].
- gRPC: nginx `grpc_pass` (HTTP/2 upstream, `grpc_set_header`); Traefik native h2c support; Caddy `reverse_proxy` with h2c transport; HAProxy mode http with ALPN h2 [official docs].
- SSE/streaming: disable proxy buffering (`proxy_buffering off` / X-Accel-Buffering: no) for event streams; HAProxy tunnel timeouts as noted in §7.2 [official; secondary].

### 11.6 Mutual TLS (mTLS)

- nginx: `ssl_verify_client on`, `ssl_client_certificate` CA bundle, `ssl_verify_depth`; optional `ssl_crl` [official].
- HAProxy: `bind ... ssl crt ... ca-file ... verify required`; `crt-list` with per-cert verify options [official].
- Traefik: `serversTransport` / entrypoint `clientAuth` with CA files [official].
- Use cases: service-to-service auth, IoT fleets, admin endpoints; often combined with SPIFFE/SPIRE in mesh environments [secondary].

---

