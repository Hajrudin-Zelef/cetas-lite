---
id: etape6-phasea-vendors-dc/00-front-matter/r3-e-meraki-dashboard-api-v1-74-0-september-2026-automation-
title: "R3-E. Meraki Dashboard API v1.74.0 (September 2026) — automation surface"
domain: front-matter
role: reference
task: reference
actors: ["AMD", "Nvidia", "Oracle"]
dates: ["2025-05", "2026-09"]
keywords: ["amd", "gpu", "latency", "nvidia", "nvlink", "optics", "pricing"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [741, 789]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: 26ccd385436552c2535081358c867c5ebec2e0c415154555bf5fd9133ca2b1ae
---

# R3-E. Meraki Dashboard API v1.74.0 (September 2026) — automation surface

### R3-E. Meraki Dashboard API v1.74.0 (September 2026) — automation surface

Base §4.2 noted AI-powered support cases (Aug 2026). The September 2026 API release adds [official — ciscodevnet meraki-portal changelog, documents/2026.md]:

- **API v1.74.0** (Sep 2026): 6 new operations, 38 enhanced, **894 total operations** — new Campus Gateway SSID mDNS update endpoint; wireless radio overrides get/update endpoints (device + org level); campusGateway clusters tunneling batchUpdate + byCluster/byNetwork reads; administered licensing subscription entitlements gained an optional `subscriptionType` query parameter (relevant to co-term/PDL/subscription licensing reporting) [official].
- Context: the API continues to grow as the primary automation path for Meraki (Dashboard API + Terraform provider), alongside the Catalyst 9500/C9610 device-config expansion noted in base §4.2 (first modular switching in Dashboard; 10-slot C9610 chassis, up to 480 ports/page, Cloud CLI) [official — community.meraki.com].

### R3-F. Dell N3200-ON / E3200-ON — OS support matrix (new)

Base §1.3 listed N-series only by name. The 2026 installation guide gives the full OS matrix [official — Dell N3200-ON/E3200-ON Series Installation Guide]:

| Model | OS6 | OS10 | SONiC (Edge standard) | Notes |
|---|---|---|---|---|
| N3208PX-ON | Y | N | N | 8-port compact PoE |
| N3224T-ON | Y | N | N | 24× 1G |
| N3224F-ON / E3224F-ON | Y / N | Y / Y | N / N | 24× 1G fiber |
| N3224P-ON | Y | N | N | 24× 1G PoE+ |
| N3224PX-ON | Y | N | N | 24× 1G PoE++ 90W |
| N3248TE-ON | Y | Y | Y | 48× 1G + 4× 10G SFP+; only model supporting all three NOS options |
| N3248P-ON / E3248P-ON | Y / N | N / N | Y / Y | 48× 1G PoE |
| N3248X-ON | Y | N | N | 48× multi-gig (1/2.5/5/10G) + 4× 25G SFP28 + 2× 100G QSFP28 |
| N3248PXE-ON / E3248PXE-ON | Y / N | N / N | Y / Y | 48× multi-gig PoE++ 90W |

- **Only OS6 supports stacking** on this family [official]. N3200-ON = 1G/2.5G/5G/10G multi-gig, PoE 30W/60W/90W tiers; **E3200-ON = SONiC Edge-standard-only SKUs** (no OS6/OS10) — the campus disaggregation on-ramp [official].
- Dell store strategy note: "N series for access, S and Z series for aggregation" (wired-LAN portfolio split) [secondary — ALSO AS Dell portal].

### R3-G. Dell switch street-price snapshots 2026 (new)

Base §1.7 found only optics prices and "quote-only" for switches. 2026 reseller/retailer evidence found (partially closing base §7 item 13 and the S-series price gap):

- **S5448F-ON** (48× 100G SFP56-DD + 8× 400G QSFP56-DD): reseller **$10,011.00** (shopiluxe.com); Czech retailer **CZK 3,572,227 incl. VAT** (CZK 2,952,254 ex-VAT, stofcom.cz) [secondary]. Note Dell Canada's official store page keeps the S5448F-ON as "Shop Now" current (Trident4-X9, 16 Tbps, 82 MB buffer, 1,135 ns) [official — dell.com].
- **Z9664F-ON** (64× 400G QSFP56-DD, Tomahawk-4): **$8,495.00 refurbished** (expresscomputersystems.com; stock 1, RAF airflow variant) — refurbished snapshot only, not new-unit pricing [secondary].
- New-unit Dell PowerSwitch pricing remains quote/configure-based on dell.com stores in 2026; the figures above are single-retailer snapshots, not list prices [caveat].

### R3-H. Aruba CX 10040 — 2026 status (finding: no new 2026 news located)

The CX 10040 (AMD Pensando DPU, "doubles the scale and performance" of CX 10000) was announced May 2025 [secondary — convergedigest, techpowerup]. No 2026 hardware refresh or new SKU was located; the platform continues as HPE's distributed-services flagship alongside the CX 10000 in the 2026 portfolio. Reseller SKU families (e.g. R8P13A/R8P14A bundles for the CX 10000) remain the price references [unverified — status quo assumed; 2026 GA/pricing detail for CX 10040 not found].

### R3 verification log (new open items)

1. Cisco FY2026 figures — from official release via trade summaries; Cisco's own press release page not opened (secondary relay) [secondary].
2. S5448F-ON $10,011 — single reseller (shopiluxe); not Dell list; authenticity of stock unconfirmed [secondary].
3. Z9664F-ON $8,495 — refurbished unit, single reseller, stale page metadata ("Last Updated" implausible) [unverified].
4. NVLink 6 bandwidth/latency claims — NVIDIA official technical blog (vendor-claimed); 3× latency / 10× packet-rate figures not independently benchmarked [vendor-reported].
5. NVLink roadmap to 1,152-GPU domains + optical scale-up — announced direction, not shipped [vendor-reported].
6. Aruba CX 10040 2026 pricing/GA — not located [gap].
7. Meraki MS355 refurbished prices — secondary market only; new-unit pricing still quote-based [secondary].
8. HPE Oracle multi-gigawatt deal — size/value terms undisclosed [unverified].

