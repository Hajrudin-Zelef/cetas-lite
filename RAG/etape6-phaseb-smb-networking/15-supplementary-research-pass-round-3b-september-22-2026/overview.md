---
id: etape6-phaseb-smb-networking/15-supplementary-research-pass-round-3b-september-22-2026/overview
title: "Supplementary Research Pass — Round 3b — September 22, 2026"
domain: supplementary-research-pass-round-3b-september-22-2026
role: deep-dive
task: reference
actors: ["Google", "United States"]
dates: ["2026-02-22", "2026-04-13", "2026-05", "2026-09-22"]
keywords: ["research", "license", "licenses", "packaging", "pricing"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [2322, 2383]
section: "Supplementary Research Pass — Round 3b — September 22, 2026"
sha256: f7954017c1352d416d23a25ae012bc889516c5263ad2238fdef33e956c3122a4
---

# Supplementary Research Pass — Round 3b — September 22, 2026

**Scope of this pass:** strictly items NOT covered by the base report or any earlier supplementary wave (S1–S18, R1–R10, T1–T8, U1–U13, V-waves). In this pass, Omada hardware-controller pricing, EnGenius 2026 L3 launches, and TRENDnet 2026 were all verified as already covered (T8/U9/T1) and are therefore **not repeated**. New in this pass: UniFi Network Application standalone-EOL / UniFi OS Server transition (2026), UniFi Protect 7.1 feature detail, Zyxel ZyWALL ATP 2026 status. No existing section was altered.

Provenance tags: [official] / [vendor-reported] / [independent] / [secondary] / [unverified] — applied per-fact.

---

## V1. UniFi Network Application standalone — 2026 EOL transition to UniFi OS Server (was: not researched)

This is the 2026 platform discontinuity most relevant to homelab/SMB self-hosted UniFi users; it was not covered in earlier waves (only OS Server version numbers appeared, §5/line 1166).

- **The standalone UniFi Network Application is approaching end-of-life.** Ubiquiti is transitioning to **UniFi OS Server**; the Home Assistant community add-on maintainer (Frenck) added an explicit EOL notice (PR #634, surfaced in a community thread dated **April 13, 2026**), stating there is **no upgrade path** from the standalone application to UniFi OS Server [independent — community repo documentation citing Ubiquiti's transition].
- **What UniFi OS Server is:** a Podman-based orchestration system running roughly **half a dozen separate containers** providing different elements of the UniFi service for both networking equipment and NVR products. Ubiquiti **does not directly publish the images** used for these components, so container/Docker re-packaging requires extracting images from a running install [independent — LinuxServer.io notice, February 22, 2026].
- **Impact on self-hosted/docker support:** LinuxServer.io announced (Feb 2026) it is **not currently considering** a UniFi OS Server image due to administrative overhead, but will keep supporting its existing UniFi Network Application image **for as long as Ubiquiti continues to make the install packages available to download**, after which it will move to deprecate [independent].
- **Practical guidance given to users:** the HA add-on README advises users planning long-term to migrate to a **dedicated machine or VM** running UniFi OS Server; the app "will continue to work as long as Ubiquiti ships the standalone application" [independent].
- **Implication for the homelab consensus (U10):** self-hosted UniFi controller setups on Docker/Home Assistant are on a deprecation clock — relevant when comparing UniFi against Omada software-controller and Grandstream GDMS (free cloud) options for SMB.
- Sources:
  - https://github.com/hassio-addons/app-unifi/blob/HEAD/README.md
  - https://github.com/hassio-addons/repository/blob/HEAD/unifi/README.md
  - https://info.linuxserver.io/issues/2026-02-22-unifi-network-application/
  - https://community.home-assistant.io/t/unifi-network-application-app-end-of-life/1004023

## V2. UniFi Protect 7.1 (May 2026) — feature detail (was: only security-fix versions mentioned)

Earlier waves cite Protect security-fix versions (7.1.83 CVE fix, 7.2.105 SAB-067) but no 7.1 feature content. This adds the 2026 feature layer as SMB ecosystem context [secondary sources; treat feature claims as vendor-reported-via-press]:

- **Protect 7.1 (announced ~May 2026):** video walls in Site Manager, **second-generation UniFi NVR** with doubled camera capacity and integrated ViewPort, Edge AI (vector search, person re-identification) computed **locally** with no recurring fees, PTZ movement tracking with vehicle recognition, native 360° camera downloads, expanded ONVIF (audio, motion events) [secondary].
- **Protect 7.0.x / 7.x (current, Aug 2026):** custom camera layouts, offsite archiving to Google Drive/OneDrive/Dropbox, a consolidated Intelligence section, PTZ movement tracking, Vantage Point multi-site view, AI Port improvements, storage I/O warnings [secondary].
- **No mandatory subscription** for core Protect functions; G6 camera line (4K + AI) from $199 per 2026 buying guides [independent].
- Sources:
  - https://globaltechonline.net/unifi-protect-7-1-video-walls-ai-nvr/
  - https://ifeeltech.com/blog/unifi-protect-cctv-guide

## V3. Zyxel ZyWALL ATP — 2026 status (was: absent)

No ATP content existed in any earlier wave; S4 covered USG FLEX H (2026 current line). Findings:

- **Explicit finding: no 2026 ATP refresh found in sources searched September 22, 2026.** The USG FLEX H series (100H/200H/500H/700H) is Zyxel's current SMB security-gateway line in 2026; ATP positioning is superseded in current portfolio materials [independent assessment; new-launch absence is evidence of a quiet line, not an EOL bulletin].
- **Launch-era positioning (2019, for baseline):** MSRP **ATP800 $1,999.99**, **ATP500 $849.99**, **ATP200 $599.99** — positioned between USG (entry) and ATP's sandboxing/evolving-threat features [secondary].
- **2026 street snapshots** [independent]:
  - ATP100: **$420.98** (pcnation, US, current listing)
  - ATP200: **£737.28 ex VAT** (comms-express, UK, 5–7 days lead)
  - ATP500 1-year gold security license: **£381**; 2-year: **£660** (ITPro UK)
- **Note for buyers:** ATP licenses are subscription add-ons (Gold = full bundle); ATP hardware is still listed at resellers but Zyxel's 2026 launches/press focus on USG FLEX H and XS/GS switches — treat ATP as channel-available/legacy-adjacent for new deployments [independent assessment].
- Sources:
  - https://www.pcnation.com/web-details.aspx?item=Z97-1033
  - https://comms-express.com/products/zyxel-atp200-eu0102f-atp200-firewall/
  - https://www.itpro.com/zyxel-networks/zyxel-zywall-atp500

## V4. Round-3b verification log

1. Omada OC200/OC220/OC300/OC400 hardware-controller pricing: already covered (T8, line ~1281; S1/S13). Not re-added.
2. EnGenius ECS6824F/ECS8830F/ECS8854F 2026 core/aggregation launch: already covered (U9, line ~1509). Not re-added. (Supplementary detail: ECS6824F launch ~Jul 14, 2026 per PRNewswire release metadata; Cloud-Lite series added in 2026 — not covered elsewhere; flagged, not researched further.)
3. TRENDnet 2026 portfolio: already covered (T1, line ~1580). Not re-added.
4. GWN7711: price $20.00 already in T4 table (line ~1252); official press-release detail (L2 Lite, 8× GbE, GDMS, GWN7711P with PoE) adds marginal context — flagged, not re-added to avoid duplication.
5. Protect 7.x: feature detail was absent (only fix-version mentions); added as V2.
6. ZyWALL ATP: wholly absent; added as V3.
7. UniFi Network Application EOL / OS Server transition: wholly absent; added as V1.
8. Currency conversions not applied; all prices are raw regional snapshots.

---
