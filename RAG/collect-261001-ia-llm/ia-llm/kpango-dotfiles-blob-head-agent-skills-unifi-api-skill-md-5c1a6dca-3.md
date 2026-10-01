---
id: collect-261001-ia-llm/ia-llm/kpango-dotfiles-blob-head-agent-skills-unifi-api-skill-md-5c1a6dca-3
title: "Restart a device"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-ia-llm/kpango-dotfiles-blob-head-agent-skills-unifi-api-skill-md-5c1a6dca.md
source_anchor: ""
source_lines: [239, 250]
sha256: cda9e94e6a26273e5322faf64fb6e820c62ea2661cab4fbd59450853e4dff50f
---

# Restart a device

tunnel's underlying physical WAN interface was passing normally while the tunnel interface's
UBIOS_IN_GEOIP/UBIOS_OUT_GEOIP chains showed the full configured country list applied. This
is a UniFi product gap, not a misconfiguration to "fix" through the UI. Verify on your own device
before assuming the gap applies: ip6tables -L UBIOS_IN_GEOIP -n (does it reference the physical
interface at all, vs. only being hooked into the DS-Lite tunnel's chains — check project memory
for which physical interface backs your DS-Lite WAN) and test actual reachability from a
GeoIP-blocked country's address over each path independently. The maintained workaround is a
standalone systemd path/timer unit scoped to that physical interface (path-triggered on
udapi-net-cfg.json changes) that mirrors the country list/direction into a separate iptables
chain that UniFi's own config-apply pipeline never touches — it must not reuse or write into
UniFi's own UBIOS_IN_GEOIP/UBIOS_OUT_GEOIP chains or ipsets, since udapi-server's own
reconciliation would then contend with it and could drop the supplemental rules on next apply.
