---
id: etape7-phasei-media/00-front-matter/7-audio-passthrough-and-transcoding
title: "7. Audio: passthrough and transcoding"
domain: front-matter
role: reference
task: model-release
actors: ["Apple"]
dates: ["2026-07"]
keywords: ["acquisition", "arr", "cost", "decode"]
source: docs/RAG/etape7_phaseI_media.md
source_anchor: ""
source_lines: [201, 261]
section: "Step 7 — Phase I: Media Servers, Transcoding and Upscaling"
sha256: 88a67bb2c5690174a02999723874f637cdf5809c7a7713ced002bb4af620bd8e
---

# 7. Audio: passthrough and transcoding

## 7. Audio: passthrough and transcoding

**Passthrough vs transcode:** if the client and HDMI chain support the codec, bitstream it untouched; only transcode when the client cannot decode it `[independent]`. Jellyfin/Emby/Plex all expose per-client audio capability profiles; the server falls back to transcoding (typically to AAC/AC3/Opus stereo or multichannel) when passthrough is impossible `[independent]`.

**Codec notes:**

- **E-AC-3 (Dolby Digital Plus)** — common on streaming rips; widely supported for passthrough on modern clients, but older devices need transcode to AC-3/AAC `[independent]`.
- **TrueHD / TrueHD Atmos** — lossless; passthrough only over HDMI eARC/direct HDMI to capable receivers; browsers and most smart-TV apps cannot handle it → transcode to E-AC-3 or AAC `[independent]`. Infuse on Apple TV 4K supports Atmos passthrough `[secondary]`.
- **DTS-HD MA / DTS:X** — same passthrough logic as TrueHD; falls back to DTS core or AAC transcode `[independent]`.
- **FLAC/PCM** — multichannel FLAC usually direct-plays on good clients; transcode to AAC/Opus for weak ones `[independent]`.
- **AC-4** — Jellyfin 10.10.0 added Dolby AC-4 support (broadcast/streaming codec) `[official]`.
- **Downmixing:** multichannel→stereo uses FFmpeg's `aresample`/pan filters with DRC optional; volume normalization (`loudnorm`) is an offline Unmanic-style step, not a real-time default `[independent]`.

**Rule of thumb:** keep the best audio track in the file, add a compatible stereo AAC track as fallback — this eliminates most audio transcodes entirely `[independent]`.

## 8. Subtitles

- **SRT (text)** — cheapest: direct-rendered by nearly all clients, no transcode `[independent]`.
- **ASS/SSA (styled text)** — direct-rendered by mpv-class players, Kodi, and modern apps (Swiftfin 1.4 improved subtitle handling `[secondary]`); weak clients trigger **burn-in**, which forces a full video transcode — the most expensive subtitle path `[independent]`.
- **PGS/VobSub (bitmap)** — cannot be rendered as text; any client that can't overlay bitmaps forces burn-in + full transcode `[independent]`. This is the #1 surprise-transcode cause in anime libraries `[independent]`.
- **Burn-in cost:** video must be fully re-encoded; on 4K HDR this combines with tone mapping into the worst-case pipeline `[independent]`. Mitigations: prefer SRT/ASS tracks, extract subtitles to sidecar SRT with the Subtitle Extract plugin or Bazarr `[independent]`, choose clients with bitmap-sub overlay support (Kodi, Infuse, mpv) `[independent]`.
- **Bazarr** (see §9) owns subtitle acquisition: downloads missing subtitles from providers, keeps them as sidecars, and avoids the Open Subtitles plugin duplication `[independent]`.

## 9. Homelab media patterns

### 9.1 The *arr stack (roles)

| App | Role | Notes |
|---|---|---|
| Sonarr (4.x) | TV series management | Default port 8989 `[secondary]` |
| Radarr (6.x) | Movie management | Default port 7878 `[secondary]` |
| Prowlarr (2.x) | Centralized indexer config for Sonarr/Radarr | Default port 9696; syncs indexers to the *arrs `[secondary]` |
| Bazarr | Subtitle management | Downloads/extracts subtitles; owns what the Open Subtitles plugin would do `[independent]` |
| Overseerr / Seerr | Request management UI | Family/friends request flow `[independent]` |
| Tautulli | Plex analytics | Plex-oriented; Jellyfin has Playback Reporting instead `[independent]` |
| exportarr | Prometheus metrics for the *arrs | Native Prometheus metrics are not guaranteed; exportarr fills the gap `[independent]` |

(Versions/ports from a secondary July 2026 comparison; verify against current releases `[secondary]`.)

### 9.2 Storage sizing

- Rule: separate **hot/transcode** (fast NVMe, small — transcode segments, trickplay images) from **cold library** (bulk HDD, optionally Ceph/mergerfs — see Step 7 Phase E) `[independent]`.
- Bitrate planning table (typical streaming rips; actual varies) `[independent]`:

| Content | Typical bitrate | 2h movie size |
|---|---|---|
| 1080p H.264 SDR | 8–12 Mbps | 7–11 GB |
| 1080p HEVC | 5–8 Mbps | 4.5–7 GB |
| 4K HEVC SDR | 25–40 Mbps | 22–36 GB |
| 4K HEVC HDR/DV remux | 50–80 Mbps | 45–72 GB |
| 4K AV1 (re-encode) | 15–30 Mbps | 14–27 GB |

- Formula: `library_TB ≈ titles × avg_GB × (1 + versions + growth)`; keep ≥15–20% free for transcode scratch, Tdarr/Unmanic staging, and snapshot headroom `[independent]`.
- Trickplay/thumbnail extraction multiplies small-file I/O — put Jellyfin's `metadata`/`transcodes` dirs on SSD `[independent]`.

### 9.3 Remote access: reverse proxy and VPN

- **Reverse proxy (recommended for sharing):** Nginx/Caddy/Traefik in front of Jellyfin with TLS, HTTP/2, and image-path caching (`/Items/*/Images` cacheable; community configs use ~100m keys zone, 15g max) `[independent]`. Required headers: `Host`, `X-Real-IP`, `X-Forwarded-For`, `X-Forwarded-Proto`, `X-Forwarded-Host` `[independent]`. WebSocket support is mandatory for the web client `[independent]`.
- **VPN (recommended for admin/family):** Tailscale/WireGuard/ZeroTier overlay — no ports exposed, Jellyfin bound to LAN + tailnet; this is the lowest-risk remote pattern `[independent]` (see Step 7 Phase G for VPN detail).
- Do not expose the Jellyfin dashboard port directly to the internet without TLS + strong auth; fail2ban/geo rules are defense-in-depth, not a plan `[independent]`.

