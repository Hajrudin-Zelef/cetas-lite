---
id: etape7-phaseh-webproxy/00-webproxy/25-apisix-plugin-catalog-2026
title: "25. APISIX plugin catalog (2026)"
domain: step-7-phase-h-web-servers-reverse-proxies-api-gateways
role: deep-dive
task: reference
actors: []
dates: ["2026-07-18"]
keywords: ["apache", "cost", "governance", "mcp", "open source", "pricing"]
source: docs/RAG/etape7_phaseH_webproxy.md
source_anchor: ""
source_lines: [521, 571]
section: "Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways"
sha256: 1707653283ba77d97d37cd49a1affb0216e805a36488f9b55907c31d49b4ccc0
---

# 25. APISIX plugin catalog (2026)

## 25. APISIX plugin catalog (2026)

- Auth: key-auth, jwt-auth, basic-auth, openid-connect, ldap-auth, hmac-auth, wolf-rbac, forward-auth [official — APISIX docs].
- Traffic: limit-req, limit-conn, limit-count (Redis cluster support), proxy-cache, proxy-rewrite, redirect, fault-injection, traffic-split [official].
- AI/MCP (2026): mcp-bridge (3.13+), plus AI-adjacent proxy plugins [official].
- Observability: prometheus, opentelemetry, skywalking, datadog, error-log-logger, http-logger, kafka-logger [official].
- etcd as config center; Admin API (REST) for CRUD; dashboard (manager-api) optional [official].
- Plugin development in Lua with hot reload; Java/Python/Go via ext-plugin runners [official].

## 26. Tyk OSS (2026)

- Tyk Gateway (Go, open source): API definitions via REST API or dashboard; middleware in JavaScript/Go/gRPC (Python via plugin compiler historically) [official — tyk.io docs; standard fact].
- Features: rate limiting/quotas (Redis-backed), key management, URL rewriting, versioning, mocking, circuit breaker, request signing [official].
- Commercial: Tyk Dashboard, Developer Portal, MDCB (multi-data-center bridge); pricing quote-based [official; gap on 2026 list prices].
- Positioning: lighter governance footprint than Kong Konnect; strong in hybrid/multi-cloud [secondary].

## 27. cert-manager YAML reference

- ClusterIssuer (Let's Encrypt prod, DNS-01 via Cloudflare): `apiVersion: cert-manager.io/v1, kind: ClusterIssuer` with `acme: server: https://acme-v02.api.letsencrypt.org/directory, privateKeySecretRef, solvers: dns01: cloudflare: apiTokenSecretRef` [official — cert-manager docs pattern].
- Certificate: `secretName`, `dnsNames`, `issuerRef: {name: letsencrypt-prod, kind: ClusterIssuer}`; auto-renewal 30 days before expiry by default [official].
- Ingress annotation flow: `cert-manager.io/cluster-issuer: letsencrypt-prod` on Ingress + `tls:` block; or Gateway API `config.enableGatewayAPI=true` [official; secondary].
- Private PKI alternative: step-ca or Vault issuer for internal services; `selfsigned` issuer for bootstrap [official].

## 28. OAuth2-Proxy reference

- Role: adds OIDC/OAuth2 login in front of apps without native SSO (complements Keycloak/Authentik; alternative to Authelia forward-auth) [secondary].
- Typical flags: `--provider=oidc`, `--oidc-issuer-url`, `--client-id/--client-secret`, `--cookie-secret`, `--email-domain`, `--upstream=http://app:8080`, `--http-address=0.0.0.0:4180` [official — oauth2-proxy docs pattern].
- nginx integration: `auth_request /oauth2/auth; error_page 401 = /oauth2/start?rd=$request_uri;` with oauth2-proxy location blocks [official docs pattern].
- 7.x line current in 2026 (exact minor not captured — gap).

## 29. Kubernetes: Ingress → Gateway API (2026)

- Implementations (2026): Envoy Gateway, Traefik (Gateway API provider), HAProxy Unified Gateway 1.0, NGINX Gateway Fabric, Istio, Cilium [official; secondary].
- 2026-07 migration study (17-row matrix, ingress-nginx replacements): Envoy Gateway 9 green/8 yellow/0 red; F5 NGINX Ingress Controller 8/9/0; NGINX Gateway Fabric 8/5/4; HAProxy KIC 6/10/1; HAProxy Unified Gateway 6/5/6 — Gateway-API-only options lack snippet mechanisms, rate limiting (HUG), external auth [independent — blog.none.at]. URL: https://blog.none.at/pdf/2026/2026-07-18-ingress-nginx-candidates.pdf
- Traefik Hub v3.20: Gateway API v1.5.1, TLSRoute in Standard channel [secondary — official docs].
- cert-manager Gateway API support via `config.enableGatewayAPI=true` [secondary].
- Practical guidance: new clusters should prefer Gateway API; existing ingress-nginx fleets should inventory annotations (snippets, custom templates) before migrating — they are the main migration cost [secondary].

## 30. Decision matrix (2026)

- Static website / shared hosting: Apache (htaccess ecosystem) or LiteSpeed (performance on same ecosystem) [secondary].
- General reverse proxy + LB, maximum docs/community: nginx OSS [secondary].
- Homelab / automatic HTTPS priority: Caddy [secondary].
- Docker auto-discovery: Traefik [secondary].
- TCP/LB precision, active health checks, low resource: HAProxy [secondary].
- Kubernetes dynamic routing / mesh: Envoy (Gateway) [secondary].
- API governance + developer portal + monetization: Kong (Konnect/Enterprise) [secondary].
- API gateway self-hosted, etcd-native, low cost: APISIX [secondary].
- API gateway Go-native, Redis-backed: Tyk [secondary].
- AI/LLM gateway features (2026): Kong AI Gateway (Enterprise), Cloudflare AI Gateway (free tier), Azure AI Gateway (consumption) [secondary].

