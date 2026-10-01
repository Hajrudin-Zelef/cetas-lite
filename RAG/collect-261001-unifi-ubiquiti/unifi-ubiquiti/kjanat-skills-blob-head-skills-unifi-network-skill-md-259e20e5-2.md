---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/kjanat-skills-blob-head-skills-unifi-network-skill-md-259e20e5-2
title: "kjanat-skills-blob-head-skills-unifi-network-skill-md-259e20e5"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/kjanat-skills-blob-head-skills-unifi-network-skill-md-259e20e5.md
source_anchor: ""
source_lines: [80, 96]
sha256: c73457e374428b424cd17de71ff49233141ae1947bd46dc2a541fae55aa7d602
---

# kjanat-skills-blob-head-skills-unifi-network-skill-md-259e20e5

Most UniFi questions arrive as a reachability failure, not as an API question. The failure mode tells you which layer to look at, and starting at the wrong layer is how an hour disappears:
| Symptom | What it means | 
|---|---|
| Timeout | Packets went nowhere — routing, firewall, or wrong path | 
| Connection refused | Path works, nothing is listening | 
| TLS error | You arrived; certificate or SNI is wrong | 
| HTTP 502/504 | Reverse proxy is up, the service behind it is not | 
| HTTP 404 | Everything is up; the routing rule does not match | 
Only the last two are application problems. A timeout is a network problem and no amount of inspecting the service will explain it.
references/diagnostics.md has the full triage recipe, including the traps that make a UniFi
network lie to you — a gateway that answers pings while refusing to route, a device inventory
that looks like a network inventory, and a firewall rule that filters only new connections and
therefore breaks reachability in exactly one direction.
- references/connector.md — auth, key scopes, failure signatures, connector limits
- references/endpoints.md — verified endpoints per generation, with what each returns
- references/diagnostics.md — reachability triage, UniFi-specific traps, worked example
- scripts/unifi-query.sh — connector call wrapper; resolves the host id, reports real status
