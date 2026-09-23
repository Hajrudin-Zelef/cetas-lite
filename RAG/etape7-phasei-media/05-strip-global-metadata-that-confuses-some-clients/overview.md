---
id: etape7-phasei-media/05-strip-global-metadata-that-confuses-some-clients/overview
title: "Strip global metadata that confuses some clients"
domain: strip-global-metadata-that-confuses-some-clients
role: deep-dive
task: actor-profile
actors: ["Apple", "Intel", "Nvidia", "Samsung"]
dates: ["2026-09"]
keywords: ["benchmark", "decode", "gpu", "intel", "nvidia", "pricing", "throughput"]
source: docs/RAG/etape7_phaseI_media.md
source_anchor: ""
source_lines: [607, 683]
section: "Strip global metadata that confuses some clients"
sha256: 213d9dad7def1a41bc0e00e9c537702fbde0d679cd9e26436ea224cd16b726e9
---

# Strip global metadata that confuses some clients
ffmpeg -i in.mkv -map 0 -c copy -map_metadata -1 -map_chapters -1 out.mkv
```

Jellyfin's own transcode pipeline uses `-map_metadata -1 -map_chapters -1` for HLS output — chapters are served via the API/segments system instead `[independent]`.

### 14.27 Bandwidth planning for remote streaming

| Stream | Needed sustained throughput |
|---|---|
| 1080p H.264 8 Mbps | ~10 Mbps (headroom) |
| 1080p HEVC 6 Mbps | ~8 Mbps |
| 4K HEVC 40 Mbps | ~50 Mbps |
| 4K remux 70 Mbps | ~90 Mbps |

Plex's own guidance: ≥2 Mbps up for 1080p remote, 20+ Mbps for 4K `[secondary]`. Rule: remote bitrate limit = 80% of measured sustained upload; direct-play 4K remuxes remotely only on symmetric/fast fiber `[independent]`. For slow uplinks, keep a pre-transcoded 1080p/720p version (Jellyfin 12 multi-version support) instead of live-transcoding 4K `[independent]`.

### 14.28 GPU buying guide for transcode hosts (2026)

- **Intel N100/N305 mini-PC:** the default homelab pick — iGPU handles multiple 4K HEVC transcodes, sips power, cheap `[independent]`.
- **Intel Arc A380 (used/new):** AV1 encode with no artificial session cap per community validation; needs recent kernel; best value discrete option `[independent]`.
- **Used NVIDIA (e.g. GTX 1660 / RTX 2060+):** NVENC quality is mature; verify per-model session limits and 10-bit HEVC support for the specific card `[independent]`.
- **Avoid:** Intel F-suffix CPUs (no iGPU) for transcode hosts; old Kepler/Maxwell cards (no HEVC 10-bit, weak NVENC) `[official]`/`[independent]`.
- **Apple Silicon:** superb client decode; as a server it's viable via Docker/macOS but VideoToolbox server tooling is thinner than QSV/NVENC `[independent]`.

### 14.29 Glossary additions

- **fmp4-HLS** — fragmented MP4 segments over HLS; Jellyfin's default transcode output format `[independent]`.
- **Chromaprint** — acoustic fingerprinting library powering Intro Skipper analysis `[independent]`.
- **RPU** — Reference Processing Unit; Dolby Vision dynamic metadata layer `[independent]`.
- **SDH / forced subs** — SDH = subtitles for deaf/hard-of-hearing (includes sound cues); forced = subtitle track that should display even when subs are off (alien dialogue) `[independent]`.
- **VMAF** — Netflix's perceptual quality metric; SVT-AV1 v4.2.0 added a VMAF tune `[independent]`.
- **CRF** — constant rate factor; quality-targeted encoding mode (lower = better/bigger) `[official]`.
- **oneVPL / libvpl** — Intel's current Video Processing Library; successor to the deprecated legacy QSV SDK paths `[secondary]`.
- **Debrid** — paid multi-hoster/cached-torrent service used by Stremio add-ons; legal status varies `[independent]`.

### 14.30 New-build quick-start checklist (opinionated)

1. Host: Intel N100-class mini-PC or Arc A380 box; Linux, recent kernel, `/dev/dri` passed through `[independent]`.
2. Server: Jellyfin 12.x (post-upgrade scan + plugin rebuilds per §14.13) `[secondary]`.
3. Files: H.264 or HEVC rips with stereo AAC fallback track + SRT/ASS subs; avoid PGS-only anime unless clients overlay bitmaps `[independent]`.
4. Automation: Sonarr/Radarr + Prowlarr + Bazarr on one compose stack (§14.9); hardlink imports `[independent]`.
5. Library optimization: Tdarr (fleet) or Unmanic (single host), staged output, checksums, rollback plan `[independent]`.
6. Remote: Tailscale/WireGuard first; Nginx reverse proxy with the §14.10 header/cache set only if sharing beyond the VPN `[independent]`.
7. Clients: Kodi or Infuse where codecs are exotic; Swiftfin/Findroid elsewhere; verify against the codec matrix before blaming the server `[independent]`/`[secondary]`.
8. Monitoring: Playback Reporting + session watch; alert on unexpected transcode reasons (`§14.11`) `[independent]`.
9. Backup: built-in Jellyfin backup + cold volume snapshots before every major upgrade `[official]`.

### 14.31 Transcode-decision flowchart (prose)

For each playback session the server evaluates, in order `[independent]`:

1. **Container/codec support** — if the client profile lists the video codec, profile, and level as supported, video direct-plays; otherwise transcode. Unknown clients get the conservative default profile (transcode-heavy) — registering proper client profiles is high-leverage `[independent]`.
2. **Bitrate ceiling** — remote bitrate limits force transcode even when the codec is supported; this is the most common "why is my 4K remux transcoding on LTE" answer `[independent]`.
3. **Audio** — unsupported audio codec/channels → audio-only transcode (video untouched, cheap); passthrough-capable HDMI chains avoid it entirely `[independent]`.
4. **Subtitles** — bitmap subs (PGS/VobSub) or ASS on weak clients → burn-in → full video transcode (expensive); SRT/ASS on capable clients → direct render (free) `[independent]`.
5. **HDR metadata** — HDR→SDR on an SDR-only client → tone-mapping stage; if the chain can't do it in hardware, it falls back to software and throughput collapses `[independent]`.

Operator habit: read the session's stated transcode reason first, then the FFmpeg log — never tune blind `[independent]`.

### 14.32 Additional conflict and caveat notes

- HandBrake 1.11's bundled SVT-AV1 4.0.1 vs standalone SVT-AV1 4.2.0: GUI releases lag the encoder; CLI FFmpeg builds may be newer — don't assume feature parity across tools `[secondary]`/`[independent]`.
- "Jellyfin 12.1" evidence comes from a single homelab repo's September 2026 docs; treat current-patch claims as `[unverified]` until an official announcement is captured `[independent]`.
- Infuse pricing (≈ $10/year) is secondary-reported; Firecore's store listing is the authority and was not captured here `[secondary]`.
- Emby's `$119 lifetime` was observed on a Samsung-TV promo URL path; figures match standard Premiere tiers but the canonical pricing page should be re-checked at purchase time `[vendor-reported]`.
- Any benchmark comparing NVENC vs QSV vs AMF quality is generation-specific; a 2024 result does not describe 2026 silicon — re-test on target hardware `[independent]`.

### 14.33 Myth corrections

- **"More GPU = better transcodes."** Transcode throughput is usually bound by decode + filter stages (tone mapping, scaling), not raw encoder power; a balanced QSV iGPU often beats a mismatched discrete card `[independent]`.
- **"AV1 is always smaller."** At low-effort presets or on noisy sources, AV1 can lose to well-tuned HEVC; preset and source matter more than codec name `[independent]`.
- **"Plex/Jellyfin quality differs."** Given the same source file and direct play, output is bit-identical — differences appear only in transcode settings and client profiles `[independent]`.
- **"Hardware encoding looks bad."** Modern NVENC/QSV/AMF at sane bitrates rival software x264 fast/medium; the quality gap closed generations ago — bitrate starvation is the real culprit `[independent]`.
- **"You need 10-bit for SDR."** 10-bit pipelines reduce banding even for SDR content, but the file must actually be encoded 10-bit end-to-end; a 10-bit container around 8-bit decisions buys nothing `[independent]`.
- **"Burning in subtitles is fine."** It forces full video re-encode — on 4K HDR it's the single most expensive common operation; fix the subtitle format instead `[independent]`.

