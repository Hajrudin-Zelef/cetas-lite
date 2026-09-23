---
id: etape7-phaseh-webproxy/00-webproxy/12-sso-integration-authelia-authentik-oauth2-proxy-keycloak
title: "12. SSO integration (Authelia, Authentik, OAuth2-Proxy, Keycloak)"
domain: step-7-phase-h-web-servers-reverse-proxies-api-gateways
role: deep-dive
task: reference
actors: []
dates: ["2025-06-27", "2026-02-05", "2026-04-08", "2026-06-16", "2026-08-20", "2026-09-02"]
keywords: ["apache", "governance", "mcp", "memory", "model context protocol", "pricing"]
source: docs/RAG/etape7_phaseH_webproxy.md
source_anchor: ""
source_lines: [244, 306]
section: "Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways"
sha256: 09909c7364d9ba41c6128d0fc9d6c1c506bb2fdb5e04633311b07b4fe6982dc5
---

# 12. SSO integration (Authelia, Authentik, OAuth2-Proxy, Keycloak)

## 12. SSO integration (Authelia, Authentik, OAuth2-Proxy, Keycloak)

### 12.1 Comparison (2026 versions)

- Observed current versions (2026): Authelia 4.39 (4.39.26), authentik 2026.8 (CalVer), Keycloak 26.7 [secondary — Stack Harbor comparison, 2026-09]. URL: https://stackharbor.com/en/knowledge-base/self-hosted-sso-keycloak-authelia/
- Forward auth: native and primary design in Authelia; authentik via proxy-provider outposts; Keycloak has none (needs oauth2-proxy or similar in front) [secondary].
- Protocols: Authelia OIDC is OpenID-certified but labeled open beta (no stability guarantee); authentik and Keycloak OIDC/SAML generally available; authentik also serves LDAP/RADIUS/SCIM/WS-Fed [secondary].
- MFA: Authelia TOTP/WebAuthn/Duo push; authentik TOTP/WebAuthn/SMS/Duo/recovery codes; Keycloak TOTP/HOTP/WebAuthn incl. passwordless [secondary].
- Footprint: Authelia <30 MB idle (argon2id reserves 64 MiB per login by default, up to 2 GiB on recommended profile); authentik documented at 2 CPU/2 GB RAM across four containers; Keycloak ~1250 MB base per pod at 10k cached sessions [secondary — same comparison; treat as indicative].
- Admin UX: Authelia has no admin UI (YAML + `authelia config validate`); authentik has web admin with visual flow designer; Keycloak has admin console + REST API [secondary].
- Storage: Authelia SQLite/PostgreSQL/MySQL + Redis session store; authentik PostgreSQL; Keycloak own DB [secondary].

### 12.2 Forward-auth integration patterns

- Traefik: `forwardAuth.address: http://authelia:9091/api/authz/forward-auth`, `trustForwardHeader: true`, `authResponseHeaders: Remote-User,Remote-Groups,Remote-Name,Remote-Email`, `maxResponseBodySize: 8192`; attach via `middlewares: authelia@docker` on routers [secondary]. URL: https://github.com/zhangqi444/open-forge/blob/HEAD/plugins/open-forge/skills/open-forge/references/projects/authelia.md
- nginx: `auth_request /authelia` subrequest pattern with `auth_request_set` for user/group headers [official — Authelia docs pattern; widely documented].
- Caddy: `forward_auth` directive to Authelia endpoint [official — Caddy docs pattern].
- Envoy: ext_authz filter [official — Envoy docs].
- Gotcha (2026 homelab report): unauthenticated redirect shape differs by client (302 to portal for browsers, 401 for non-browser `Accept` headers); tests must assert the redirect target, not just the status [secondary].
- OAuth2-Proxy: alternative for services with native OIDC; commonly paired with Keycloak or cloud IdPs; 7.x line current in 2026 (exact 2026 minor not captured — gap).
- Helm deployment (Authelia 4.39): helmforge chart with forward-auth for Traefik/nginx/Caddy/Envoy, Prometheus metrics, S3 backup, Gateway API HTTPRoute support, cert-manager TLS [secondary]. URL: https://github.com/helmforgedev/charts/blob/HEAD/charts/authelia/README.md

---

## 13. API gateways

### 13.1 Kong

- Kong Gateway: Lua/OpenResty-based API gateway; Kong Gateway OSS is Apache-2.0 (`Kong/kong`); Enterprise plugins and Konnect control plane are proprietary [secondary — API management comparison, 2026-09]. URL: https://github.com/robakid/apimanagementtools/blob/HEAD/comparisons/api-gateways.md
- Kong Gateway 3.12+ ships AI MCP Proxy plugin (REST↔MCP conversion, four modes: passthrough-listener, conversion-listener, conversion-only, listener) and AI MCP OAuth2 plugin; vendor states the AI MCP Proxy is "only available as part of our AI Gateway Enterprise offering" [secondary].
- Konnect Plus (modular SaaS pricing, observed 2026-09-02): $25/month serverless control plane (up to 5), $200/month hybrid (up to 2), $500/month Dedicated Cloud (one region); 1M API requests/month included, $200 per additional 1M/month, public max 10M/month on Plus; Enterprise custom annual [secondary]. URL: https://github.com/robakid/apimanagementtools/blob/HEAD/comparisons/api-gateways.md
- Older Konnect consumption pricing (still cited in 2026 analyses): ~$105/month per Gateway Service, ~$34.25 per 1M API requests, ~$720/month base for dedicated cloud gateway instances ($1/hour), bandwidth $0.15/GB, advanced analytics +$20/1M requests, mesh zones $4,166/zone/month [secondary — konghq.com pricing PDF Sept 2023–Mar 2025; treat as legacy tier]. URL: https://assets.prd.mktg.konghq.com/files/2025/03/67ca4460-kong-konnect-pricing-september2023-march2025.pdf
- AI Gateway framing (2026): prompt routing, semantic caching, PII sanitization; mid-size enterprise deployments commonly exceed $50k/year with support/governance [secondary — tech-insider.org, 2026-09].
- Plugin model: Lua plugins via PDK; external Go/JavaScript plugins; declarative config via decK YAML or Admin API; Konnect Dev Portal (Plus: up to 2 portals; Enterprise: unlimited); no portal in OSS gateway [secondary].
- Deployment: self-hosted OSS/Enterprise, Konnect SaaS control plane with self-managed or cloud data planes, Kubernetes ingress [secondary].

### 13.2 Apache APISIX

- Apache APISIX 3.18.0 released 2026-08-20 (latest, active security support); 3.17.0 (2026-06-16); 3.16.0 (2026-04-08); 3.15.0 (2026-02-05); each minor is supported until the next minor releases (~2–4 month cadence) [secondary — eosl.date]. URL: http://eosl.date/eol/product/apache-apisix/
- GitHub releases page (2026-06-16): 3.17.0 latest at that crawl; release train visible: 3.17.0, 3.16.0, 3.15.0, 3.14.1, 3.14.0, 3.13.0 (2025-06-27), 3.12.0 [secondary]. URL: https://github.com/apache/apisix/releases
- 3.13.0 highlights (2025-06-27): `mcp-bridge` and `lago` plugins, standalone Admin API (HTTP PUT/GET, JSON/YAML, stateless in-memory config), status health-check endpoint, L4 proxy health check; `server-info` plugin deprecated (etcd write churn) [official — APISIX blog]. URL: https://apisix.apache.org/blog/2025/06/27/release-apache-apisix-3.13.0/
- Architecture: OpenResty/nginx + etcd data plane; decoupled data-plane mode warns on etcd writes (to be disallowed); standalone API-driven mode for ingress controller use [official].
- APISIX Ingress Controller 2.x requires APISIX 3.13+ for standalone API-driven mode [official]. URL: https://github.com/apache/apisix-ingress-controller/blob/HEAD/docs/en/latest/overview.md
- Install: RPM/DEB repos (repos.apiseven.com), Docker, Helm; etcd required (auto-installed with Docker/Helm) [official]. URL: https://apisix.apache.org/docs/apisix/installation-guide/
- Vendor ecosystem: API7 (company behind APISIX) publishes Konnect pricing comparisons and gateway guides [secondary]. URL: https://api7.ai/blog/kong-konnect-pricing

### 13.3 Tyk

- Tyk: open-source API gateway (Go); Gateway OSS + commercial Dashboard/Developer Portal; cloud and self-hosted [official — tyk.io; standard fact].
- Pricing: quote-based for commercial tiers; no 2026 list prices captured in this round [gap].
- Noted in 2026 pricing comparisons alongside Kong OSS and APISIX as the self-hostable alternatives to Konnect [secondary].

### 13.4 Cross-gateway capabilities

- Service discovery: Kong (DNS/SRV, Consul, K8s), APISIX (etcd watch, DNS, Nacos/Eureka/Consul), Tyk (DNS, Consul), Traefik Hub (K8s, Consul) [official docs; secondary].
- Rate limiting plugins: fixed-window, sliding-window, token-bucket; Redis-backed distributed limiting for multi-instance gateways [official docs].
- Auth plugins: key-auth, JWT, OAuth2/OIDC, HMAC, LDAP, mTLS, AI-provider keys (Kong AI Gateway) [official docs].
- OpenAPI: import OpenAPI 3.x specs to generate routes (Kong Insomnia/ decK, APISIX dashboard import, Tyk OAS API definitions) [official docs].
- Observability: Prometheus metrics, OpenTelemetry tracing, access logs; Kong Konnect analytics with 14-month retention on paid tiers [official; secondary].
- MCP/AI: Kong AI MCP Proxy (Enterprise), APISIX `mcp-bridge` plugin (OSS, 3.13+) — both target Model Context Protocol server proxying in 2026 [secondary].

---

