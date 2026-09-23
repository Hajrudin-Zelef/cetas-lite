---
id: etape7-phasei-media/00-front-matter/14-11-transcode-troubleshooting-matrix
title: "14.11 Transcode troubleshooting matrix"
domain: front-matter
role: reference
task: model-release
actors: ["Apple", "Samsung"]
dates: ["2026-02"]
keywords: ["decode"]
source: docs/RAG/etape7_phaseI_media.md
source_anchor: ""
source_lines: [456, 506]
section: "Step 7 — Phase I: Media Servers, Transcoding and Upscaling"
sha256: 4e5c585848529eda3e48d1082a6fa65bdc166f3a0e5cb834d2e8e657e9ad7365
---

# 14.11 Transcode troubleshooting matrix

### 14.11 Transcode troubleshooting matrix

| Symptom | Likely cause | Fix |
|---|---|---|
| Everything transcodes, nothing direct-plays | Client profile too strict / bitrate limit low | Raise remote bitrate limit; check client codec matrix `[independent]` |
| 4K HDR stutters, SDR fine | Tone-mapping in software chain | Enable HW tone mapping or direct-play via Infuse/Kodi `[independent]` |
| Subtitle selection triggers transcode | PGS/VobSub bitmap subs | Extract to SRT (Bazarr/Subtitle Extract); use ASS/SRT `[independent]` |
| Transcode starts then buffers | Disk I/O on transcode dir | Move `transcodes` dir to SSD; check free space `[independent]` |
| HW accel option missing | `/dev/dri` not passed / wrong groups / old kernel | Verify device, GID, kernel version per §4.1 `[independent]` |
| Green/pink output | Pixel-format mismatch in filter chain | Insert `format=` conversions before `hwupload` `[independent]` |
| Audio only, no video | Client lacks video codec | Check codec-support matrix; remux to H.264 `[official]`/`[independent]` |
| High CPU on "direct play" | Actually remuxing + audio transcode | Read the FFmpeg log; add fallback AAC track to files `[independent]` |

### 14.12 Remote-access hardening checklist

1. Prefer VPN (Tailscale/WireGuard/ZeroTier) for admin and family; reverse proxy only if you must share with non-technical users `[independent]`.
2. TLS everywhere on the proxy; HSTS; no plain-HTTP dashboard on WAN `[independent]`.
3. Strong unique Jellyfin admin password; disable unused user accounts; review active sessions periodically `[independent]`.
4. fail2ban on proxy 401/403 patterns as defense-in-depth `[independent]`.
5. Keep Jellyfin and `jellyfin-ffmpeg` patched; 12.x removed legacy endpoints — verify old client bookmarks after upgrade `[secondary]`.
6. Backups per §14.1 before every major upgrade; test restore `[official]`.

### 14.13 Jellyfin 10.11 → 12.x upgrade playbook

1. Confirm current version and read the release notes for every major in the path; 10.11 requires starting from 10.10.7, and 12.0 requires a supported 10.11/10.10 base (sources conflict on 10.10.7 vs 10.10.8 — verify against the official notes at upgrade time) `[secondary]`.
2. Run the built-in backup (10.11+) or a cold file-level backup of config, data, and metadata directories `[official]`.
3. Snapshot the host/VM or image the container volumes — database migrations are one-way `[independent]`.
4. Upgrade during a maintenance window: EF Core migrations on large libraries are slow `[official]`.
5. After upgrade: run a **full library scan**, re-verify hardware transcoding with a forced-transcode test, re-check HDR tone mapping, and reinstall/rebuild every plugin against the new ABI (Intro Skipper ships per-ABI builds, e.g. 12.0.4.0) `[secondary]`/`[independent]`.
6. Confirm clients still connect: 12.0 removed legacy `/emby/` and `/mediabrowser/` endpoints, which breaks old hardcoded clients/bookmarks `[secondary]`.
7. Keep the pre-upgrade backup until the new version has survived at least one full scan + one week of playback `[independent]`.

### 14.14 Client-by-client transcode behavior notes

- **Web browsers:** Chrome/Edge/Firefox direct-play H.264/AAC reliably; HEVC in Firefox arrived with Jellyfin 10.11 `[official]`; Safari plays HEVC/H.264 but balks at many audio codecs — expect audio transcodes `[independent]`.
- **Android TV / Fire TV / Shield:** ExoPlayer-based clients direct-play H.264/HEVC/VP9; AV1 on recent hardware; DTS-HD/TrueHD usually need audio transcode unless passthrough to a receiver `[independent]`.
- **Roku:** historically the pickiest client — H.264 + AAC is the safe profile; HEVC support varies by model `[independent]`.
- **Xbox:** official Jellyfin app exists; codec support mirrors the platform's media stack `[independent]`.
- **LG webOS / Samsung Tizen:** official Jellyfin apps since February 2026 `[secondary]`; smart-TV SoCs decode H.264/HEVC well but struggle with PGS subtitles (burn-in risk) and lossless audio `[independent]`.
- **Kodi + Jellyfin add-on:** two modes — *add-on mode* (Kodi browses the Jellyfin library natively, best codec support) and *native/direct-path mode*; add-on mode is recommended for codec-problem clients `[independent]`.
- **Infuse:** direct-plays nearly everything including Dolby Vision all profiles and Atmos passthrough on Apple TV 4K; InfuseSync plugin keeps the Jellyfin library state coherent `[secondary]`/`[independent]`.
- **Chromecast:** supported via cast targets; codec support follows the Chromecast model — often forces transcode for HEVC/ass subtitles `[independent]`.

### 14.15 Social and library features compared

- **SyncPlay (Jellyfin):** synchronized group watching built into the server since 10.6 `[independent]`; Plex Watch Together and Emby's equivalent are account-tied `[secondary]`.
- **Users/profiles:** Jellyfin multi-user with per-user parental controls and library restrictions is free `[independent]`; Plex Home management is Pass-gated `[secondary]`; Emby parental controls are Premiere `[secondary]`.
- **Music:** Jellyfin handles music libraries (MusicBrainz/AudioDB metadata) with clients like Fintunes/Jellify `[independent]`; Plexamp is the category's best music client but needs Plex Pass `[secondary]`.
- **Books/comics:** Jellyfin 12.0 reworked book support (EPUB reading dates back to 10.6); JellyBook (Android) and Fintunes-class readers cover mobile `[independent]`/`[secondary]`.
- **Live TV/DVR:** Jellyfin via tuner plugins (HDHomeRun etc.) `[independent]`; Plex Pass DVR `[secondary]`; Emby Premiere DVR `[secondary]`; Channels DVR ($8/month) is the specialist alternative `[secondary]`.

