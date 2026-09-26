---
id: etape6-phasee1-netbox-nautobot/00-netbox-nautobot/part-12
title: "Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File) (part 12)"
domain: phase-e1-netbox-nautobot-network-source-of-truth-research-fi
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["apache", "lean"]
source: docs/RAG/etape6_phaseE1_netbox_nautobot.md
source_anchor: ""
source_lines: [513, 521]
section: "Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File)"
sha256: 4071dd4c948bb9bec7fd54785683b10bee463cebca5804d8b3bda63751f6d9fc
---

# Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File) (part 12)

- Official introduction docs state three design tenets:
  1. **Replicate the real world** — strict data model (e.g., IP addresses assigned to interfaces, not devices; an interface may hold multiple IPs).
  2. **Serve as a "Source of Truth"** — NetBox holds the **desired state**, not operational state; automated import of live network state is **strongly discouraged**; all data should be human-vetted before entry so downstream tools can populate with high confidence.
  3. **Keep it simple** — the 80% solution is favored over a complete but complex one; low learning curve, lean codebase.
- Official application stack: HTTP service (nginx or Apache) → WSGI (gunicorn or uWSGI) → Django/Python → **PostgreSQL 15+** → task queue **Redis/django-rq**. NetBox does not talk to network nodes directly; it makes data available programmatically to automation, monitoring, and assurance tools (separation of duties; swap tools without changing the data authority) `[official]`. Source: https://github.com/netbox-community/netbox/blob/HEAD/docs/introduction.md
- README (2026): "successor to legacy IPAM and DCIM applications... central source of truth for the modern network" `[official]`. Source: https://github.com/netbox-community/netbox/blob/HEAD/README.md

---

