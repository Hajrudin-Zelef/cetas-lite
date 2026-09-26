---
id: etape7-phaseh-webproxy/00-webproxy/part-2
title: "Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways (part 2)"
domain: step-7-phase-h-web-servers-reverse-proxies-api-gateways
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["apache", "memory"]
source: docs/RAG/etape7_phaseH_webproxy.md
source_anchor: ""
source_lines: [50, 55]
section: "Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways"
sha256: dbb213e40fb3ad9dae61efc315ae3d82ea6f68b867757b9b45ffbb1b0ee62c96
---

# Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways (part 2)

- Apache remains default on shared hosting and legacy environments; nginx is the default for reverse proxy, load balancing and high-concurrency serving [secondary — medium.com comparison, 2026-01].
- Architecture contrast: Apache process/thread-per-connection model is flexible but more memory-hungry under load; nginx event-driven async model handles thousands of connections per worker with low memory [secondary].
- Apache's modular ecosystem and .htaccess remain advantages in multi-tenant hosting; nginx's static-config + reload model and lower resource use dominate high-traffic, container and edge deployments [secondary].

---

