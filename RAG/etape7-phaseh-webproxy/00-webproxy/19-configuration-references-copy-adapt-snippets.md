---
id: etape7-phaseh-webproxy/00-webproxy/19-configuration-references-copy-adapt-snippets
title: "19. Configuration references (copy-adapt snippets)"
domain: step-7-phase-h-web-servers-reverse-proxies-api-gateways
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["apache"]
source: docs/RAG/etape7_phaseH_webproxy.md
source_anchor: ""
source_lines: [416, 470]
section: "Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways"
sha256: 3a3d74fcc13950b9011564b9e99422fa0b2ae5b1f1384f92d343b99a8e9de483
---

# 19. Configuration references (copy-adapt snippets)

## 19. Configuration references (copy-adapt snippets)

### 19.1 nginx reverse proxy + TLS (2026 idioms)

- Upstream with keepalive and passive health checks: `upstream app { server 10.0.0.11:8080 max_fails=3 fail_timeout=30s; server 10.0.0.12:8080; keepalive 32; }` [official].
- Location block essentials: `proxy_pass http://app; proxy_http_version 1.1; proxy_set_header Connection ""; proxy_set_header Host $host; proxy_set_header X-Real-IP $remote_addr; proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for; proxy_set_header X-Forwarded-Proto $scheme;` [official].
- TLS server block: `listen 443 ssl; ssl_certificate /etc/letsencrypt/live/example.com/fullchain.pem; ssl_certificate_key .../privkey.pem; ssl_protocols TLSv1.2 TLSv1.3; ssl_prefer_server_ciphers off;` (with TLS 1.3, cipher preference is client-driven) [official].
- Security headers block: `add_header Strict-Transport-Security "max-age=63072000; includeSubDomains; preload" always; add_header X-Content-Type-Options nosniff always; add_header X-Frame-Options DENY always; add_header Referrer-Policy "strict-origin-when-cross-origin" always;` [official directives; secondary for the value set].
- Rate limit: `limit_req_zone $binary_remote_addr zone=login:10m rate=5r/m;` then `limit_req zone=login burst=10 nodelay;` on the login location [official].
- WebSocket location: `proxy_http_version 1.1; proxy_set_header Upgrade $http_upgrade; proxy_set_header Connection "upgrade"; proxy_read_timeout 86400;` [official].
- gRPC: `location /grpc/ { grpc_pass grpc://10.0.0.11:50051; grpc_set_header X-Real-IP $remote_addr; }` [official].
- Caching: `proxy_cache_path /var/cache/nginx levels=1:2 keys_zone=static:10m max_size=1g inactive=60m;` then `proxy_cache static; proxy_cache_valid 200 10m; proxy_cache_use_stale error timeout updating;` [official].
- Compression: `gzip on; gzip_types text/css application/javascript application/json;` — brotli requires the third-party ngx_brotli module; zstd support is module-dependent [official; secondary].
- Reload without downtime: `nginx -s reload` spawns new workers with the new config and gracefully drains old ones; `nginx -T` dumps the effective config for auditing [official].
- Access log JSON: `log_format json_combined escape=json '{...}';` for structured log shipping [official].

### 19.2 Caddyfile reference

- Basic: `example.com { reverse_proxy 127.0.0.1:8080 }` — HTTPS automatic [official].
- Multi-upstream LB: `reverse_proxy 10.0.0.11:8080 10.0.0.12:8080 { lb_policy least_conn; health_uri /healthz; health_interval 10s; }` [official].
- Header control: `header { Strict-Transport-Security "max-age=31536000; includeSubDomains" X-Content-Type-Options nosniff -Server }` [official].
- TLS options: `tls { dns cloudflare {env.CLOUDFLARE_API_TOKEN}; on_demand; }` — DNS-01 for wildcards, on-demand issuance for SaaS-style custom domains (rate-limit caution) [official].
- Internal PKI: `tls internal` issues from Caddy's local CA (install with `caddy trust`) for LAN services [official].
- Logging: `log { output file /var/log/caddy/access.log; format json; }` [official].
- Handle routing: `handle /api/* { reverse_proxy 127.0.0.1:9000 }` + `handle { reverse_proxy 127.0.0.1:3000 }` [official].

### 19.3 Traefik static + dynamic config

- Static (CLI/file): entrypoints web (:80) / websecure (:443), providers.docker=true, certificatesResolvers.letsencrypt.acme with email/storage/httpChallenge [official].
- Docker labels: `traefik.enable=true`, `traefik.http.routers.app.rule=Host(`app.example.com`)`, `traefik.http.routers.app.entrypoints=websecure`, `traefik.http.routers.app.tls.certresolver=letsencrypt`, `traefik.http.services.app.loadbalancer.server.port=8080`, `traefik.http.routers.app.middlewares=sec-headers,rate-limit` [official].
- File provider middleware example: `[http.middlewares.sec-headers.headers] stsSeconds=63072000 stsIncludeSubdomains=true contentTypeNosniff=true frameDeny=true` [official].
- TCP routing (databases, non-HTTP): `traefik.tcp.routers.*.rule=HostSNI(*)` with `tls.passthrough=true` for SNI passthrough [official].
- Dashboard: protect `api@internal` router with basicAuth or forwardAuth middleware; never expose insecurely [official guidance].

### 19.4 HAProxy config reference

- Global: `global` → `nbthread 4`, `maxconn 100000`, `ssl-default-bind-options ssl-min-ver TLSv1.2`, `tune.ssl.default-dh-param 2048` [official].
- Frontend/backend split: `frontend https_in` binds `:443 ssl crt /etc/haproxy/certs/ alpn h2,http/1.1`; ACLs `acl is_api path_beg /api`; `use_backend api_nodes if is_api`; `default_backend web_nodes` [official].
- Backend: `balance leastconn`, `option httpchk GET /healthz`, `http-check expect status 200`, `server web1 10.0.0.11:8080 check inter 5s rise 2 fall 3`, `cookie SERVERID insert indirect nocache` for sticky sessions [official].
- Rate limiting via stick-tables: `stick-table type ip size 100k expire 30s store http_req_rate(10s)` + `http-request track-sc0 src` + `http-request deny deny_status 429 if { sc_http_req_rate(0) gt 100 }` [official].
- Header security: `http-response set-header Strict-Transport-Security "max-age=63072000; includeSubDomains"` [official].
- Stats: `listen stats` with `stats enable`, `stats uri /haproxy-stats`, `stats auth admin:...`, `stats socket /var/lib/haproxy/stats` for runtime API [official].
- Runtime changes without reload: Data Plane API or `socat` to the stats socket (`set server web/web1 state drain`) [official].

### 19.5 Apache reverse-proxy reference

- Enable: `a2enmod proxy proxy_http proxy_balancer proxy_hcheck ssl http2 cache` (Debian) [official].
- Balancer: `<Proxy "balancer://app"> BalancerMember http://10.0.0.11:8080 route=w1 hcmethod=GET hcuri=/healthz BalancerMember http://10.0.0.12:8080 route=w2 ProxySet lbmethod=byrequests stickysession=JSESSIONID </Proxy>` then `ProxyPass "/app" "balancer://app"` [official].
- TLS: `SSLEngine on`, `SSLCertificateFile/KeyFile`, `SSLProtocol -all +TLSv1.2 +TLSv1.3`, `SSLCipherSuite` / `SSLHonorCipherOrder` [official].
- HTTP/2: `Protocols h2 http/1.1` [official].
- Caching: `CacheEnable disk /`, `CacheRoot /var/cache/apache`, `CacheMaxExpire` [official].
- Security: `Header always set Strict-Transport-Security ...`, `TraceEnable Off`, `ServerTokens Prod` [official].

---

