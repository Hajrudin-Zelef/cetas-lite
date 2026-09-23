---
id: etape7-phaseh-webproxy/00-webproxy/31-migration-notes
title: "31. Migration notes"
domain: step-7-phase-h-web-servers-reverse-proxies-api-gateways
role: deep-dive
task: reference
actors: ["AWS"]
dates: []
keywords: ["apache", "aws", "benchmark", "pricing"]
source: docs/RAG/etape7_phaseH_webproxy.md
source_anchor: ""
source_lines: [572, 616]
section: "Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways"
sha256: 3b7590479ad79051856515b501509257dc2c14da6c505d04b6c7478586411c7b
---

# 31. Migration notes

## 31. Migration notes

- Apache → nginx: translate .htaccess rewrites to `rewrite`/location blocks; move auth to `auth_basic` or forward-auth; benchmark static vs dynamic split (nginx in front of Apache for dynamic is a classic hybrid) [secondary].
- nginx → Caddy: Caddyfile is far shorter; watch for missing knobs (no built-in rate limit in core; header control via `header` directive) [secondary].
- ingress-nginx → Gateway API: audit annotations first (auth-url, snippets, custom config); choose Envoy Gateway or Traefik for broadest annotation coverage [independent — blog.none.at].
- HAProxy 2.x → 3.2: QUIC library choice (OpenSSL 3.5 vs quictls vs AWS-LC) affects PQC and performance; Data Plane API 3.2 adds ACME [official; secondary].

## 32. Troubleshooting checklist

- `nginx -t` / `haproxy -c -f` / `apachectl configtest` / `caddy validate` / `traefik` debug log before reload [official].
- 502s: upstream down, wrong `proxy_pass` scheme, keepalive mismatch, SELinux/AppArmor blocking egress [secondary].
- TLS errors: expired cert, missing intermediate chain, SNI mismatch, clock skew; test with `openssl s_client -connect host:443 -servername host` [secondary].
- WebSocket failures: missing Upgrade headers, buffering on, short timeouts [secondary].
- ACME failures: port 80 blocked (HTTP-01), DNS propagation (DNS-01), rate limits hit — check staging first [secondary].
- Performance: check `ulimit -n`, `worker_connections`/`maxconn` saturation, upstream keepalive reuse, TLS session resumption hit rate [secondary].

*End of sections 19–32. Extended reference sections (33–44) follow.*

---

## 33. nginx beyond HTTP: stream, mail, njs

- Stream module (`--with-stream`): TCP/UDP load balancing — databases, MQTT, DNS; `upstream` with `least_conn`, `hash $remote_addr consistent`; `proxy_protocol` support; SSL termination for TCP via `ssl_preread` SNI routing without decrypting [official].
- `ssl_preread`: route TLS by SNI at L4 (`map $ssl_preread_server_name $backend`) for multi-tenant TLS passthrough [official].
- Mail proxy: IMAP/POP3/SMTP proxying with auth via `auth_http` [official].
- njs: JavaScript scripting in nginx config (subrequests, dynamic routing, JWT validation without Lua); njs 0.9.x line in 2026 [official — nginx njs docs; secondary for version].
- JavaScript (njs) vs OpenResty Lua: njs is lighter and officially maintained by F5/nginx; OpenResty LuaJIT is faster and has a larger library ecosystem [secondary].

## 34. Apache module catalog (selection, 2026)

- ACME: mod_md (managed domains) automates Let's Encrypt issuance/renewal natively — the Apache answer to Caddy's automatic HTTPS [official].
- Compression: mod_deflate (gzip), mod_brotli (Brotli, separate package on most distros) [official].
- Proxy: mod_proxy_http, mod_proxy_fcgi, mod_proxy_uwsgi, mod_proxy_ajp, mod_proxy_wstunnel, mod_proxy_hcheck, mod_proxy_balancer [official].
- Auth: mod_auth_basic/digest, mod_auth_openidc (third-party, OpenIDC), mod_auth_gssapi (Kerberos), mod_authnz_ldap [official for bundled; secondary for mod_auth_openidc maintenance].
- Security: mod_reqtimeout (Slowloris mitigation), mod_evasive (third-party DoS), mod_security (third-party WAF, OWASP CRS) [official for bundled; secondary for third-party status].
- Caching: mod_cache, mod_cache_disk, mod_cache_socache, mod_file_cache [official].
- Scripting: mod_php (prefork only for thread safety), mod_proxy_fcgi + php-fpm (recommended with event MPM), mod_perl, mod_wsgi, mod_lua [official].

## 35. HAProxy Community vs Enterprise (2026)

- Community (MPLv2): full LB/proxy core, ACLs, stick tables, QUIC/HTTP-3 (3.2+), Data Plane API, Prometheus exporter, SPOE, Lua scripting [official].
- Enterprise adds: WAF (next-gen), Bot Management + Threat Detection Engine (3.2), Global Traffic Manager, API/AI Gateway, DeviceAtlas detection, realtime dashboard, FIPS builds, SLA support [vendor-reported].
- HAProxy One: platform bundling Enterprise data plane + Fusion control plane + secure edge network; Fusion 2.0 with K8s operator, Terraform provider, Ansible playbooks [vendor-reported].
- Pricing: Enterprise quotation-only; no public list prices as of 2026 [secondary].

