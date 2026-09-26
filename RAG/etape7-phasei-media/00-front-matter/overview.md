---
id: etape7-phasei-media/00-front-matter/overview
title: "Step 7 — Phase I: Media Servers, Transcoding and Upscaling"
domain: front-matter
role: reference
task: model-release
actors: ["AWS", "Apple", "Google", "Nvidia", "Samsung"]
dates: ["2018-12", "2019-01", "2024-05-11", "2024-10-26", "2024-10-27", "2025-10-19", "2025-10-20", "2026-01", "2026-02", "2026-07-01", "2026-09-07", "2026-09-08", "2026-09-17", "2026-09-22"]
keywords: ["nvidia", "pricing", "research"]
source: docs/RAG/etape7_phaseI_media.md
source_anchor: ""
source_lines: [1, 54]
section: "Step 7 — Phase I: Media Servers, Transcoding and Upscaling"
sha256: 036ac04321a0d6b17324f5e6709eb64b914de072dc1cdf067365d721f82b2263
---

# Step 7 — Phase I: Media Servers, Transcoding and Upscaling

> **Document control.** Step 7 / Phase I — media and transcoding. Current through 2026-09-22 `[independent]`. English per project rule `[independent]`. Every factual claim carries one allowed provenance label only: `[official]`, `[vendor-reported]`, `[independent]`, `[secondary]`, `[unverified]`. No claim is asserted without a label. Conflicts are recorded, not silently resolved `[independent]`. One writer per file `[independent]`.

| Field | Value |
|---|---|
| File | `etape7_phaseI_media.md` |
| Coverage | Jellyfin / Plex / Emby / Stremio / Kodi; FFmpeg 7.x–8.x; hardware encoders (NVENC, QSV, VAAPI, AMF, VideoToolbox); HDR tone mapping; AV1/H.264/H.265; Tdarr; Unmanic; upscaling; audio; subtitles; homelab patterns |
| Line target | ≥ 750 lines `[independent]` |

## Provenance key

- `[official]` — project or vendor documentation, release notes, changelogs, source repositories.
- `[vendor-reported]` — vendor statements not in formal documentation (pricing pages, announcements, store listings).
- `[independent]` — reproducible community testing, homelab write-ups, third-party engineering blogs, Wikipedia.
- `[secondary]` — press/tech media reporting, often without primary confirmation.
- `[unverified]` — claims observed once or lacking a citable source; presented as uncorroborated.

## 1. Media servers

### 1.1 Jellyfin

Jellyfin is a free, open-source media server forked from Emby in December 2018 by Andrew Rabert, Joshua Boniface and other contributors after Emby closed its open-source development model `[independent]`. The name is a play on "streaming" `[independent]`. Its distinctive versioning began at 10.0.0 in January 2019 `[independent]`.

**Release chronology (recent):**

- **10.9.0** — 2024-05-11 `[independent]`.
- **10.10.0** — 2024-10-27 `[independent]` (announced 2024-10-26 in the official blog `[official]`). Headline features: media segments (chapters/intros as first-class items), faster trickplay extraction, software HDR tone mapping, improved QSV hardware transcoding, Dolby AC-4 audio, Dolby Vision handling, FFmpeg 7.0 base `[official]`/`[secondary]`.
- **10.11.0** — 2025-10-20 `[independent]` (official blog dated 2025-10-19 `[official]`). Headline features: migration to Entity Framework Core with a unified database architecture, built-in backup/restore tooling, FFmpeg 7.1, startup UI and log viewer, HEVC decoding support in Firefox; 32-bit ARM (armhf) support removed `[official]`. The 10.10.x line was dropped as upgrade prerequisite (10.10.7 required) `[secondary]`; community evidence shows the 10.11 line reached at least **10.11.11** `[independent]`.
- **12.0** — 2026-09-07 `[independent]` (secondary press reported 2026-09-08 `[secondary]`; the one-day discrepancy is recorded as a conflict, not resolved). The "10." prefix is dropped. Changes: backend on **.NET 10** `[secondary]`, transcoding stack on **FFmpeg 8.1** `[secondary]`, reworked database structures, improved books/comics support, playlists and collections rework, search and recommendation extensibility, a Modern UI, removal of legacy `/emby/` and `/mediabrowser/` endpoints, and a "Still watching?" prompt plus frame-by-frame playback and multiple TV episode versions `[independent]`/`[secondary]`. Upgrade requires backup, a supported starting version (secondary sources conflict on whether 10.10.7 or 10.10.8 is the minimum — recorded as a conflict), and a full post-upgrade library scan `[secondary]`. Plugins compiled against 10.11 must be rebuilt for 12.0 `[secondary]`.
- **12.1** — community homelab documentation updated 2026-09-17 references stable 12.1 builds and Intro Skipper 12.0.4.0 rebuilt against them `[independent]`; no official 12.1 announcement was captured, so current patch state is `[unverified]`.

**Architecture:** server written in .NET (C#), currently .NET 10 `[secondary]`; EF Core unified database from 10.11 `[official]`; bundled `jellyfin-ffmpeg` binary (a Jellyfin-patched FFmpeg) drives all transcoding `[official]`; SQLite-backed item metadata and library database (see Step 7 Phase F for database internals) `[independent]`; optional DLNA server via plugin and Chromecast sending support `[independent]`. Clients communicate over a REST API plus WebSockets for events like library changes `[official]`.

**Client matrix (current 2026):**

- Official: Jellyfin Web `[independent]`; Jellyfin Media Player (Windows/macOS/Linux) `[independent]`; Android phone/tablet `[official]`; Android TV (also covers NVIDIA Shield, Amazon Fire TV, Google TV) `[official]`; Roku, Xbox `[independent]`; LG webOS `[independent]`; Kodi add-on `[independent]`; DLNA via plugin `[independent]`.
- Official Apple-native: **Swiftfin** (iOS/iPadOS/tvOS), built natively in Swift/SwiftUI `[secondary]`. **Swiftfin 1.4** (January 2026) was a major milestone: navigation/routing overhaul, full Jellyfin 10.11 support, revamped media-player manager, improved library browsing, better subtitle handling, tvOS improvements `[secondary]`. Swiftfin remains in beta `[secondary]`.
- Third-party: Findroid (native Android) `[official]`; Homedia (Jetpack Compose Android TV) `[official]`; Streamyfin `[independent]`; Infuse (tvOS/iOS/macOS) — community-rated as the best 4K HDR/Dolby Vision client on Apple TV 4K with full Dolby Vision profile support and Dolby Atmos passthrough; pricing reported ≈ $10/year `[secondary]` (official store price not captured in this research). InfuseSync plugin on the server improves library sync for Infuse `[independent]`.
- Exotic: Sailfin (Sailfish OS), FinVideo/FinMusic (HarmonyOS), Switchfin (Nintendo Switch, beta) `[secondary]`.
- 2026 status note: Jellyfin landed official apps on Samsung and LG smart TVs in February 2026, closing a long-standing platform gap vs Plex/Emby `[secondary]`.

**Plugin ecosystem:** plugins install from Dashboard → Plugins → Catalog or custom manifests `[independent]`. Bundled metadata providers: TMDb, OMDb, Studio Images, AudioDB, MusicBrainz `[independent]`. Commonly deployed third-party plugins: Intro Skipper (chromaprint audio-fingerprint intro/credit detection exposed as media segments; manifest `https://manifest.intro-skipper.org/manifest.json`; version keyed to server ABI, e.g. 12.0.4.0 for 12.x) `[independent]`; Chapter Segments Provider (official) `[independent]`; Playback Reporting (official) `[independent]`; Trakt scrobbling (official) `[independent]`; TMDb Box Sets `[independent]`; Merge Versions `[independent]`; Kodi Sync Queue (official) `[independent]`; LDAP-Auth (official) `[independent]`; Webhook (official; commonly wired to ntfy for "item added" notifications) `[independent]`; Open Subtitles (official; often skipped when Bazarr owns subtitles) `[independent]`; Subtitle Extract (official, conditional) `[independent]`; File Transformation `[independent]`.

### 1.2 Plex

Plex is the proprietary, commercially operated media server that anchors the category; the client list above is benchmarked against its device coverage `[independent]`. Key 2026 economics (all `[secondary]` unless noted; official Plex pricing announcement text was not captured in this research):

- Plex Pass: **$6.99/month** or **$69.99/year** `[secondary]`.
- Lifetime Plex Pass: raised from $249.99 to **$749.99**, effective 2026-07-01; existing lifetime holders grandfathered `[secondary]`.
- Hardware transcoding remains a Plex Pass feature `[secondary]`.
- Plex moved remote access toward paid tiers during 2025 `[secondary]`.
- At lifetime pricing the payback vs monthly is roughly nine years, which is why secondary press frames it as aimed at committed users rather than casual ones `[secondary]`.

