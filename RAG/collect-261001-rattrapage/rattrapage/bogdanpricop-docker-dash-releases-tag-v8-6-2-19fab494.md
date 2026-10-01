---
id: collect-261001-rattrapage/rattrapage/bogdanpricop-docker-dash-releases-tag-v8-6-2-19fab494
title: "bogdanpricop-docker-dash-releases-tag-v8-6-2-19fab494"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-05-27"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/bogdanpricop-docker-dash-releases-tag-v8-6-2-19fab494.md
source_anchor: ""
source_lines: [1, 14]
sha256: 7601c1e20a4703acf704c611d40043746c475abe12dda4ffec934e2fc14c340d
---

# bogdanpricop-docker-dash-releases-tag-v8-6-2-19fab494

v8.6.2 — UniFi Network Application template
New built-in template: UniFi Network Application
A self-hosted Ubiquiti UniFi controller, ready to deploy from Containers → Templates. Built on the LinuxServer.io image (the community-maintained successor since Ubiquiti discontinued the official Docker controller).
What's in the stack
- lscr.io/linuxserver/unifi-network-application:latest — the controller.
- mongo:7.0 — dedicated MongoDB sidecar (pinned per LSIO guidance; 7.0 is the highest version supported by UniFi <9.0).
- Inline mongo init script via Compose configs: withcontent: (Compose 2.23+) — auto-creates theunifi DB user withdbOwner onunifi andunifi_stat on first start. No host file to bind-mount.
- Healthcheck-gated startup: the controller's depends_on: { unifi-db: { condition: service_healthy } } + amongosh ping healthcheck solves the LSIO-documented "restart container after Mongo is ready" footgun.
- All 9 standard ports declared with role comments (4 required: 8443 admin / 8080 device / 3478 STUN / 10001 AP discovery; 5 optional commented with their purpose).
- Named volumes for /config and/data/db .
Before first deploy
The header comment in the template spells it out: replace changeme in BOTH MONGO_PASS AND the init script pwd: fields — they must match (the init script runs ONCE on first Mongo start; rotating later means dropping the unifi-db volume or updating via mongosh).
Verified
Registered in BUILTIN_VERIFICATION (2026-05-27); YAML round-trip parsed via the project's yaml lib; template-tests + full suite green (1444).
