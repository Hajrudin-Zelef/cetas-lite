---
id: etape6-phasea-vendors-dc/01-round-5-complementary-research-pass-2026-09-22-deep-datashee/part-4
title: "Round-5 complementary research pass (2026-09-22) — deep datasheet and radix detail (part 4)"
domain: round-5-complementary-research-pass-2026-09-22-deep-datashee
role: deep-dive
task: reference
actors: ["Broadcom", "Nvidia"]
dates: ["2026-03", "2026-09-22"]
keywords: ["compute", "distribution", "ethernet", "gpu", "nvidia"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1937, 1945]
section: "Round-5 complementary research pass (2026-09-22) — deep datasheet and radix detail"
sha256: 9c3b05eae1e1382289c89c207da218c7b7485858eeabe489e8655638dfd866a8
---

# Round-5 complementary research pass (2026-09-22) — deep datasheet and radix detail (part 4)

- Dell support KB 000228560 (Minimum/Recommended/Latest code versions, crawl ~Jul 2026) shows **Enterprise SONiC 4.6.0 as the latest release** on both Broadcom (30-June-2026) and NVIDIA Spectrum (SN-4.6.0, 11-June-2026); **recommended remains 4.5.3 (05-June-2026)**. No 4.7.x train was located as of 2026-09-22 — 4.6.0 is the current top-of-train [official].
  Source: https://www.dell.com/support/kbdoc/en-ph/000228560/minimum-recommended-and-latest-code-versions-for-networking-products
- The same KB documents Dell Enterprise SONiC 4.6.0 availability for **EdgeCore white-box models** — AS4630-54PE, AS5835-54T, AS7326-56X, AS7712-32X, AS7726-32X, AS7816-64X, AS9716-32D (latest 4.6.0 released 30-June-2026, recommended 4.5.3; note: EdgeCore products are supported by EdgeCore, Dell supports only the Enterprise SONiC distribution) [official].
- **SN5601** appears in the KB as a Spectrum platform carrying SN-4.6.0 (SN2201, SN4700, SN5600, SN5601) — a model name not previously itemized in this file [official].
- Dell InfoHub white paper **H04658** "Dell Technologies AI Fabrics Overview: NVIDIA Spectrum with Dell SONiC" (March 2026) and PoC technical brief **HO4617** "AI Fabrics NVIDIA Spectrum with Dell SONiC" (March 2026): describe the GPU-cluster back-end fabric, compute multi-tenant front-end fabric, and out-of-band management network using Spectrum switches with Dell SONiC [official].
  Sources: http://infohub.delltechnologies.com/static/media/client/7phukh/DAM_60dba377-fc13-4251-9dd6-c81409b4095d.pdf ; http://infohub.delltechnologies.com/static/media/client/7phukh/DAM_215fe847-2754-45d2-9f66-b54fe3bfbd6a.pdf
- Dell blog (2026): "Open Ethernet for AI: NVIDIA Spectrum-X with Dell SONiC" — RoCEv2 integrated, adaptive routing for congestion avoidance, ECN-marking + PFC congestion management integrated with Spectrum-X telemetry, native in-band telemetry with Grafana integrations, compliant with the NVIDIA Cloud Partner Reference Architecture [official].
  Source: https://www.dell.com/en-us/blog/open-ethernet-for-ai-nvidia-spectrum-x-with-dell-sonic/

