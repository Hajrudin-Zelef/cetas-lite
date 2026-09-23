---
id: etape7-phasei-media/00-front-matter/14-1-jellyfin-internals-worth-knowing
title: "14.1 Jellyfin internals worth knowing"
domain: front-matter
role: reference
task: model-release
actors: ["Intel"]
dates: []
keywords: ["cost", "intel"]
source: docs/RAG/etape7_phaseI_media.md
source_anchor: ""
source_lines: [307, 362]
section: "Step 7 — Phase I: Media Servers, Transcoding and Upscaling"
sha256: c5b47c4e7fa9c1e9d478270791afc1f5b9dfe04904fee8838ff6cfbed4bdf0ae
---

# 14.1 Jellyfin internals worth knowing

### 14.1 Jellyfin internals worth knowing

- **Database:** 10.11 migrated to Entity Framework Core with a unified schema `[official]`; 12.0 reworks database structures again `[secondary]`. Practical consequence: upgrades from 10.10.x require passing through 10.10.7 and can take a long time on large libraries — plan maintenance windows `[secondary]`.
- **Backup/restore:** built into the server from 10.11 (Dashboard), covering config, database and metadata `[official]`; the official docs page is `docs/general/administration/backup-and-restore.md` `[official]`. Third-party scripts predate it and still exist `[independent]`.
- **Trickplay:** 10.10 made trickplay (timeline preview thumbnails) extraction faster `[official]`; the extracted tiles live under the metadata directory and multiply small-file I/O — SSD placement matters `[independent]`.
- **Multiple versions:** 12.0 supports multiple versions of TV episodes natively `[independent]`; the Merge Versions plugin bulk-merges duplicates on older releases `[independent]`.
- **Direct-play decision inputs:** container, video codec/profile/level, audio codec/channels, subtitle format, bitrate vs client limit, and HDR metadata — the server computes a transcode reason per session, visible in the dashboard `[independent]`.
- **Playback reporting / audit:** the official Playback Reporting plugin records who watched what and when — the Jellyfin answer to Tautulli (which is Plex-only) `[independent]`.
- **Web client codec matrix:** Jellyfin publishes a per-client codec-support matrix at `jellyfin.org/docs/general/clients/codec-support/` `[official]` — consult it before blaming the server for a transcode.
- **Real Jellyfin transcode command** (from a 2026 user log, Intel QSV host, HEVC→AAC HLS remux) `[independent]`:

```
"/usr/lib/jellyfin-ffmpeg/ffmpeg" -analyzeduration 200M -probesize 1G -readrate 10 \
 -fflags +genpts -f matroska -i file:"/mediateque/Series/.../Friends.S01E03.1080p.MULTi.AvALoN.mkv" \
 -map_metadata -1 -map_chapters -1 -threads 3 -map 0:0 -map 0:1 -map -0:s \
 -codec:v:0 copy -tag:v:0 hvc1 -bsf:v hevc_mp4toannexb -start_at_zero \
 -codec:a:0 libfdk_aac -ac 2 -ab 256000 -af "volume=2" \
 -copyts -avoid_negative_ts disabled -max_muxing_queue_size 2048 \
 -f hls -max_delay 5000000 -hls_time 6 -hls_segment_type fmp4 \
 -hls_fmp4_init_filename "<id>-1.mp4" -start_number 0 \
 -hls_segment_filename "/home/michael/jellyfin/metadata/images/transcodes/<id>%d.mp4" \
 -hls_playlist_type vod -hls_list_size 0 -hls_segment_options movflags=+frag_discont \
 -y "/home/michael/jellyfin/metadata/images/transcodes/<id>.m3u8"
```

Reading this pattern tells an operator: video was direct-streamed (`copy` + bitstream filter), audio transcoded to AAC (client lacked the source codec), subtitles dropped (`-map -0:s` — external/SRT handling instead), output fmp4-HLS `[independent]`.

### 14.2 Plex feature detail (2026)

- Plex Pass headline features: hardware transcoding, offline mobile downloads, Live TV/DVR with tuner, Skip Intro/Skip Credits, Plexamp (music app), multi-user Home management, trailer/extras, camera upload `[secondary]`.
- Remote access moved toward paid tiers in 2025; free-tier remote streaming limits are the mechanism Plex uses to convert free users `[secondary]`.
- Plex's relay feature routes remote streams through Plex infrastructure when direct connection fails — convenient, bandwidth-limited, and a privacy consideration `[independent]`.
- Plex HTPC/desktop apps and the web app remain free for local playback `[secondary]`.

### 14.3 Emby feature detail

- Premiere headline features: DVR with guide data, full mobile apps (free tier apps have playback limits), parental controls, offline sync, Cinema Mode intros, cover-art/fanart enhancements `[secondary]`.
- Emby's theater apps (Android TV, Fire TV) and Kodi add-on are mature; the Emby-for-Kodi plugin predates Jellyfin's and shares its sync architecture `[independent]`.

### 14.4 x264/x265 preset reference

**libx264 presets** (speed → compression), `[official]`:

| Preset | Typical use |
|---|---|
| ultrafast, superfast, veryfast | Real-time capture/transcode |
| faster, fast | Balanced real-time |
| medium | Default; offline baseline |
| slow, slower | Archival encodes |
| veryslow | Maximum-effort archival |
| placebo | Diminishing returns; rarely justified |

**libx265 presets** add `placebo`; practical streaming presets are `veryfast`–`fast`; archival `slow`–`veryslow`; 10-bit (`-pix_fmt yuv420p10le`) is recommended for HDR and reduces banding even in SDR `[official]`/`[independent]`.

**Tune options:** `zerolatency` for real-time (disables lookahead — quality cost) `[official]`; `film`, `animation`, `grain` for content-specific x264/x265 tuning `[official]`.

