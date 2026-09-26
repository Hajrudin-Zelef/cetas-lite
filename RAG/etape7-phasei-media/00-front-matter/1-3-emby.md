---
id: etape7-phasei-media/00-front-matter/1-3-emby
title: "1.3 Emby"
domain: front-matter
role: reference
task: model-release
actors: ["Apple", "Samsung"]
dates: ["2026-09-22"]
keywords: ["cost", "license", "pricing"]
source: docs/RAG/etape7_phaseI_media.md
source_anchor: ""
source_lines: [55, 89]
section: "Step 7 — Phase I: Media Servers, Transcoding and Upscaling"
sha256: 1eae9f07f82952ab5985222e11ada653ca291ff216942b4e645872ae679793bc
---

# 1.3 Emby

Where Plex still wins in 2026: app polish and setup ease, mainstream smart-TV coverage, and remote-access convenience that saves troubleshooting time `[secondary]`. Where it loses: cost, account dependency, and privacy (all streams route through Plex account infrastructure for remote access) `[independent]`.

### 1.3 Emby

Emby is the proprietary media server Jellyfin forked from in 2018 `[independent]`. Official pricing observed 2026-09-22 `[vendor-reported]`:

- Emby Premiere: **$4.99/month**, **$54/year**, or **$119 one-time**.
- Standard license terms describe a 15-device household limit `[vendor-reported]`.

Emby occupies the middle ground: cheaper lifetime than Plex, free tier retained, Premiere adds DVR, full mobile apps, and parental controls `[secondary]`. Secondary comparisons rate Emby's Apple TV/Roku clients as stronger than Jellyfin's community-built options `[secondary]`.

### 1.4 Stremio

Stremio is categorically different: a **player and catalog with an add-on system, not a personal-media server** `[independent]`. It does not ingest a folder of local files and stream them; it ingests catalog entries and resolves them through add-ons `[independent]`. The desktop client code is GPL-2.0 (`Stremio/stremio-web`); the company behind it is commercial and the model centers on the add-on catalog `[independent]`. First-party add-ons handle metadata and legitimate streaming sources; third-party add-ons extend into Debrid and torrent-streaming integrations whose legal status varies by jurisdiction `[independent]`. Stremio has no official Roku channel (screen-mirroring workarounds only) `[secondary]`. Evaluation rule: if the media library is already on disk, Stremio is the wrong tool; if the goal is discovering new content, it is genuinely useful at what it does — judge the two cases separately `[independent]`.

### 1.5 Kodi

Kodi (originally Xbox Media Center, 2002; renamed 2014) is a free, open-source **media centre player**, not a server: it manages media on the device being watched on `[independent]`. It runs on Windows, macOS, Linux, Android, and Fire TV, with the largest add-on ecosystem of the five-way comparison `[secondary]`, and handles PVR backends for live TV `[secondary]`. The recommended home-theater pattern is **Jellyfin as server + Kodi as player** via the Jellyfin-for-Kodi add-on: Kodi syncs the full library, watched status syncs back, Kodi's playback engine handles every codec natively, and Jellyfin supplies multi-user support, remote access, and transcoding `[secondary]`. The same pairing pattern exists with the Kodi Emby plugin `[independent]`.

### 1.6 Head-to-head matrix (2026)

| Dimension | Jellyfin | Plex | Emby | Stremio | Kodi |
|---|---|---|---|---|---|
| Model | Self-hosted server, FOSS, no tiers | Self-hosted server, proprietary | Self-hosted server, proprietary | Cloud catalog + add-on player | Local player |
| Cost | Free, no paywall | Free at home; Pass $6.99/mo or $749.99 lifetime `[secondary]` | Free tier; Premiere $4.99/mo, $54/yr, $119 lifetime `[vendor-reported]` | Free | Free, FOSS |
| Hardware transcoding | Free, all accelerators | Plex Pass only `[secondary]` | Premiere-tier feature set `[secondary]` | N/A | N/A (plays natively) |
| Multi-user | Yes, free | Pass-gated home mgmt `[secondary]` | Yes | Basic | No |
| Smart TV | Official Samsung + LG apps since Feb 2026 `[secondary]`; Android TV; Roku; Xbox | Broadest coverage `[secondary]` | Samsung, LG `[secondary]` | No (web/mobile/TV casting) `[secondary]` | LG webOS (v21+) `[secondary]` |
| Apple TV | Swiftfin (official, beta) + Infuse | Polished native app `[secondary]` | Stronger than Jellyfin's `[secondary]` | Limited | Via MrMC-class wrappers `[secondary]` |
| Live TV / DVR | Via tuner plugins | Pass feature `[secondary]` | Premiere `[secondary]` | No | PVR backends `[secondary]` |
| Metadata | Self-managed, plugin providers | Automatic, account-linked | Automatic | Automatic | Local |
| Privacy | Full self-control | Account-dependent | Account-dependent | Add-on-dependent | Full self-control |
| Offline | Your files, yes | Pass downloads `[secondary]` | Yes | No | Local files, yes |
| Intro skipping | Intro Skipper plugin → media segments `[independent]` | Pass feature (Skip Intro) `[secondary]` | Cinema intros `[secondary]` | No | No |

