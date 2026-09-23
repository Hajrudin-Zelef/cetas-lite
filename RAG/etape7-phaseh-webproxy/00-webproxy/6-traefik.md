---
id: etape7-phaseh-webproxy/00-webproxy/6-traefik
title: "6. Traefik"
domain: step-7-phase-h-web-servers-reverse-proxies-api-gateways
role: deep-dive
task: reference
actors: ["AWS"]
dates: ["2026-01-29", "2026-08-07", "2026-09"]
keywords: ["agent", "aws", "benchmark", "latency", "open source", "pricing", "throughput"]
source: docs/RAG/etape7_phaseH_webproxy.md
source_anchor: ""
source_lines: [119, 165]
section: "Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways"
sha256: d8dff52d3b5198ce289aba66e97c5bb8d4ef3d8c042d3f6df4da91b3ae49e8c6
---

# 6. Traefik

## 6. Traefik

### 6.1 Release status (2026)

- Traefik Proxy v3.7.x is the current line as of September 2026: v3.7.10 (helm chart 41.2.0, 2026-08-07); v3.7.11+ adds collision-safe naming (safeNaming option); v3.7.7 changed wildcard `Host(*)` to catch-all; v3.7.9 rejects HTTP/1 CONNECT with 501 [secondary — official changelog/helm chart mirrors]. URLs: https://github.com/traefik/traefik-helm-chart/blob/HEAD/traefik/Changelog.md, https://github.com/traefik/traefik/blob/HEAD/CHANGELOG.md
- Version support policy updated starting with v3.6 [secondary — traefik/traefik PR #13627].
- Security: v3.7.8 fixed GHSA-8rxv-jg7p-wvg3; v3.6.7 fixed CVE-2026-22045 (2026-01-29, NixOS backport) [secondary].
- Traefik Hub v3.20 ships with Traefik Proxy v3.7: Gateway API v1.5.1 (TLSRoute now Standard channel), Knative provider v1.20.0, wildcard Host/HostSNI matchers, provider precedence, service-level middlewares for K8s CRD, ForwardAuth options (authSigninURL, maxResponseBodySize), ACME certificateTimeout, cipherSuites on ServersTransport, OTel-conformant trace attributes for OTLP access logs [secondary — official Traefik Hub docs]. URL: https://doc.traefik.io/traefik-hub/api-gateway/release-notes
- Traefik v2.10 remains the legacy fallback for older Docker environments [secondary].

### 6.2 Architecture & routing model

- Written in Go; dynamic configuration discovery via providers: Docker, Kubernetes Ingress/CRD/Gateway API, Consul, etcd, file, Nomad, ECS, etc. [official — Traefik docs].
- Core concepts: EntryPoints → Routers (rules: Host, Path, Headers, Method…) → Middlewares → Services (load-balanced servers) [official].
- Middlewares: rateLimit, basicAuth/digestAuth, forwardAuth, headers (HSTS/security headers), compress, circuitBreaker, retry, stripPrefix/replacePath, ipAllowList, inFlightConn, buffering, errors pages, encodedCharacters (v3.7+) [official].
- ACME: built-in Let's Encrypt support (HTTP-01, TLS-ALPN-01, DNS-01 with 100+ providers), `certificatesResolvers`, `certificateTimeout` (v3.7+) [official].
- Dashboard/API: `api@internal` router; insecure mode off by default in production guidance [official].
- Forward-auth pattern with Authelia: `forwardAuth.address: http://authelia:9091/api/authz/forward-auth`, `trustForwardHeader: true`, `authResponseHeaders: Remote-User,Remote-Groups,...` — widely used homelab SSO pattern [secondary]. URL: https://github.com/zhangqi444/open-forge/blob/HEAD/plugins/open-forge/skills/open-forge/references/projects/authelia.md

---

## 7. HAProxy

### 7.1 Release status (2026)

- HAProxy 3.2 is the current long-term support release of the open-source project (per HAProxy Technologies, KubeCon Amsterdam announcement 2026) [vendor-reported]. URL: https://ott.einnews.com/pr_news/901566349/haproxy-technologies-launches-haproxy-fusion-2-0-and-haproxy-unified-gateway-1-0-at-kubecon-amsterdam
- HAProxy Enterprise 3.2 (announced ~Nov 2025): Threat Detection Engine (TDE) in the Bot Management Module (app DDoS, brute force, scrapers, vuln scanners), QUIC/UDP/HTTP LB, API/AI gateway, K8s routing, SSL offload, global rate limiting, next-gen WAF; positioned as "world's fastest software load balancer + edge security layer" [vendor-reported]. URL: https://www.dbta.com/Editorial/News-Flashes/HAProxy-Enterprise-32-Reinforces-Next-Gen-Security-Intelligence-172203.aspx
- HAProxy One: commercial platform unifying data plane, control plane (Fusion) and secure edge; HAProxy Fusion 2.0 adds Kubernetes-native deployment (Fusion Operator), official Terraform Provider, Ansible playbooks, Fusion API v2 [vendor-reported].
- HAProxy Unified Gateway 1.0 (free, open source, KubeCon Amsterdam 2026): brings Kubernetes Gateway API standard to HAProxy; HTTP/HTTPS/TLS routes; config via CRDs; dynamic scaling from K8s API; Prometheus metrics; Helm install; built on HAProxy 3.2 [vendor-reported].
- HAProxy Data Plane API 3.2: ACME support (`acme` section, `crt-store`, runtime `acme renew`/`acme status` APIs; HTTP-01 Enterprise-only) [official]. URL: http://cdn.haproxy.com/documentation/haproxy-data-plane-api/release-notes/v3.2/
- QUIC/HTTP/3: as of HAProxy 3.2, QUIC builds against multiple TLS libraries — OpenSSL 3.5 (native QUIC API), quictls, WolfSSL, AWS-LC; HAProxy's 2025 benchmarking recommends AWS-LC for best performance; quictls has no release since its OpenSSL 3.3 fork (no PQC key exchange on QUIC path) [secondary — technical blog, 2026-09]. URL: https://instatunnel.my/blog/the-enterprise-edge-load-balancing-migration-transitioning-from-localhost-tunnels-to-global-infrastructure
- Typical QUIC termination config: `bind quic4@:443 ssl crt ... alpn h3` + `cluster-secret` for stateless reset tokens [secondary].
- Community image tags: `haproxy:3.1`, `haproxy:latest` used in 2026 benchmark harnesses [secondary].
- HAProxy Enterprise pricing is quotation-based, not publicly disclosed [secondary — spotsaas.com 2026].
- G2 Spring 2026 Container Networking Grid: HAProxy customer satisfaction score 98 (top in category); average time-to-ROI 4 months vs 9-month category average [vendor-reported via press release].

### 7.2 Configuration model

- Sections: global, defaults, frontend, backend, listen, resolvers, peers, cache, mailers, http-errors [official].
- Load-balancing algorithms: roundrobin, static-rr, leastconn, first, source (IP hash), uri, url_param, hdr, rdp-cookie [official].
- Health checks: active TCP/HTTP checks with `check`, `httpchk`, rise/fall thresholds, agent-check; on-the-fly server state via stats socket / Data Plane API [official].
- ACLs: L4–L7 request routing (`hdr(host)`, `path_beg`, `ssl_fc`, method matchers), maps for large ACL datasets [official].
- SSL/TLS: `bind :443 ssl crt`, ALPN h2/http1.1, crt-list with SNI, OCSP stapling, `ssl-min-ver`, dynamic cert update via Data Plane API [official].
- Noted behavioral default (Sept 2026 technical analysis): HAProxy batches body writes via kernel (favors throughput; can add ~200 ms on small-frame streaming workloads such as LLM token streams unless tuned) while nginx defaults `tcp_nodelay on` (favors streaming latency); both are one-config-line tunables [secondary — dev.to analysis, 2026-09]. URL: https://dev.to/rlnorthcutt/haproxys-200-ms-llm-token-delay-is-a-smart-default-heres-when-to-change-it-3eah

---

