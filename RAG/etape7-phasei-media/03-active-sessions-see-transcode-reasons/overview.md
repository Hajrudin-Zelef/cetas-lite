---
id: etape7-phasei-media/03-active-sessions-see-transcode-reasons/overview
title: "Active sessions (see transcode reasons)"
domain: active-sessions-see-transcode-reasons
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["embedding"]
source: docs/RAG/etape7_phaseI_media.md
source_anchor: ""
source_lines: [562, 604]
section: "Active sessions (see transcode reasons)"
sha256: e4b6b6e44a2e5cb05d11ad4574f53e0618843cf1dad1772160a8755858444b8d
---

# Active sessions (see transcode reasons)
curl -s -H "X-Emby-Token: $K" "$J/Sessions" | python3 -m json.tool | less
```

Useful for monitoring (session watch → alert on unexpected transcodes) and for plugin-version audits before upgrades `[independent]`.

### 14.22 Kodi + Jellyfin add-on modes

- **Add-on mode (default):** Kodi shows the Jellyfin library through the add-on's virtual paths; metadata/artwork come from Jellyfin; playback uses Kodi's engine — maximum codec compatibility, slightly slower browsing on huge libraries `[independent]`.
- **Native/direct-path mode:** Kodi accesses files directly via SMB/NFS paths; faster, but path mapping must be exact or items won't play `[independent]`.
- **Kodi Sync Queue plugin** (server side) pushes library deltas to Kodi so add-on mode stays fresh without full rescans `[independent]`.

### 14.23 Jellyfin server performance tuning

- **Library scans:** scheduled scans during idle hours; real-time monitoring (inotify) for small libraries, periodic scans for network mounts where inotify is unreliable `[independent]`.
- **Parallel image extraction:** chapter/trickplay extraction is I/O-bound — SSD metadata dir matters more than CPU `[independent]`.
- **Transcode thread control:** Jellyfin passes `-threads N` per job; on many-core hosts cap concurrent transcodes rather than threads to avoid context-switch storms `[independent]`.
- **Database:** keep `jellyfin.db` on fast local storage; network-mounted databases (NFS) risk locking issues and slow queries — local disk + backup is the supported pattern `[independent]`.
- **Reverse-proxy caching:** image endpoints (`/Items/*/Images`) are highly cacheable; HLS segments must NOT be cached aggressively (they're per-session) — cache only images and static web assets `[independent]`.
- **Log retention:** transcode logs accumulate under the logs dir; rotate or they fill small boot volumes `[independent]`.

### 14.24 Migrating between servers (Plex/Emby → Jellyfin)

- **Watched state:** no native importer; community scripts map Plex/Emby watch history into Jellyfin via API — verify counts after migration `[independent]`.
- **Metadata:** NFO/sidecar artwork largely re-reads; embedded tags carry over; collections/playlists need rebuilding or scripted import `[independent]`.
- **Clients:** audit every household device against the codec-support matrix *before* cutover — the #1 migration regret is a client that direct-played on Plex but transcodes on Jellyfin due to profile differences `[independent]`.
- **Parallel run:** run Jellyfin alongside the old server pointed at the same (read-only) library for a week before decommissioning `[independent]`.
- **Intro Skipper / Tautulli equivalents:** reinstall Playback Reporting + Intro Skipper on day one so feature parity is visible `[independent]`.

### 14.25 Subtitle conversion commands

Extract embedded PGS to sidecar (then OCR externally if text needed) `[official]`/`[independent]`:

```bash
ffmpeg -i input.mkv -map 0:s:0 subs.srt          # text subs -> SRT sidecar
ffmpeg -i input.mkv -map 0:s:0 -c:s copy subs.sup # bitmap subs -> sidecar (still bitmap)
```

Batch-extract first SRT/ASS track from a library with Tdarr/Unmanic flows or a find loop; Bazarr then manages fetching missing languages `[independent]`. Prefer keeping one full SDH track + one forced track per language `[independent]`.

### 14.26 Chapter and metadata embedding

```bash
