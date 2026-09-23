---
id: etape7-phasei-media/00-front-matter/14-9-arr-stack-reference-deployment
title: "14.9 *arr stack reference deployment"
domain: front-matter
role: reference
task: model-release
actors: ["AMD", "Intel"]
dates: []
keywords: ["arr", "amd", "intel"]
source: docs/RAG/etape7_phaseI_media.md
source_anchor: ""
source_lines: [411, 455]
section: "Step 7 — Phase I: Media Servers, Transcoding and Upscaling"
sha256: 58d777f09ad2f95598552987a463c02cb0c75ac3156c064327cc956db2a615bd
---

# 14.9 *arr stack reference deployment

### 14.9 *arr stack reference deployment

```yaml
services:
  prowlarr:  { image: lscr.io/linuxserver/prowlarr:latest, ports: ["9696:9696"] }
  sonarr:    { image: lscr.io/linuxserver/sonarr:latest,    ports: ["8989:8989"] }
  radarr:    { image: lscr.io/linuxserver/radarr:latest,    ports: ["7878:7878"] }
  bazarr:    { image: lscr.io/linuxserver/bazarr:latest,    ports: ["6767:6767"] }
  jellyfin:
    image: jellyfin/jellyfin:latest
    ports: ["8096:8096"]
    devices: ["/dev/dri:/dev/dri"]   # Intel/AMD hardware transcoding
    group_add: ["105"]               # render group (host-specific GID)
```

- Data flow: Prowlarr → Sonarr/Radarr (indexers synced) → download client → media folder → Jellyfin library scan (inotify or scheduled) → Bazarr subtitles → Intro Skipper fingerprinting `[independent]`.
- Hardlink/atomic-move pattern: download client and *arrs share the same volume so imports are instant renames, not copies `[independent]`.
- All images above are community linuxserver.io builds; pin digests in production `[independent]`.

### 14.10 Reverse-proxy snippet (Nginx, Jellyfin)

```nginx
proxy_cache_path /var/cache/nginx/jellyfin levels=1:2 keys_zone=jellyfin:100m
                 max_size=15g inactive=30d use_temp_path=off;
server {
  location / {
    proxy_pass http://jellyfin:8096;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_set_header X-Forwarded-Host $http_host;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "upgrade";  # WebSocket for web client
  }
  location ~ /Items/(.*)/Images {
    proxy_pass http://jellyfin:8096;
    proxy_cache jellyfin; proxy_cache_revalidate on; proxy_cache_lock on;
  }
}
```

Observed in production homelab configs 2026-09 `[independent]`. Caddy/Traefik equivalents are common; the header set and WebSocket upgrade are the non-negotiable parts `[independent]`.

