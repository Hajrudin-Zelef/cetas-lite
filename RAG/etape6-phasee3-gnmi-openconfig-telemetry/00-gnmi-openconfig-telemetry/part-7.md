---
id: etape6-phasee3-gnmi-openconfig-telemetry/00-gnmi-openconfig-telemetry/part-7
title: "Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability (part 7)"
domain: step-6-phase-e3-gnmi-openconfig-model-driven-telemetry-progr
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["optics"]
source: docs/RAG/etape6_phaseE3_gnmi_openconfig_telemetry.md
source_anchor: ""
source_lines: [250, 254]
section: "Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability"
sha256: a4f5a272067dd907ff4a91ec8e310eecedb82fd3c85d686e818238e175af8047
---

# Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability (part 7)

- **Version skew:** devices in the field expose OpenConfig models years behind HEAD (e.g. Junos support-portal table lists BGP 2.0.1/2.1.1-era models for 16.1R/17.1R; the gnmic issue dump shows interfaces 2.3.0 while 2026 revisions exist) — multi-version fleets are the norm [official](https://supportportal.juniper.net/sfc/servlet.shepherd/document/download/0693c00000LXblkAAD/?operationContext=S1) [secondary](https://github.com/openconfig/gnmic/issues/451).
- **Deviations:** vendors publish deviation modules (e.g. Arista `arista-bfd-deviations`, `arista-exp-eos-*`) where OpenConfig doesn't fit the implementation [secondary](https://files.botwerks.net/presentations/20220922-mnnug-openconfig.pdf).
- **Coverage holes:** platform pipeline counters, MACsec, terminal-device optics, and power-state are 2025–2026 additions — recent enough that not all vendors implement them [official — dated revisions in Waves 1–2].
- **gNMI `Capabilities` as the discovery mechanism:** clients must query supported models/encodings per device rather than assuming a uniform OpenConfig surface [secondary](https://files.botwerks.net/presentations/20220922-mnnug-openconfig.pdf).

