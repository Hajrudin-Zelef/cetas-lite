---
id: etape7-phasei-media/00-front-matter/10-decision-guides
title: "10. Decision guides"
domain: front-matter
role: reference
task: model-release
actors: ["AMD", "Apple", "Intel", "Nvidia", "Samsung"]
dates: ["2026-07", "2026-09-07", "2026-09-08", "2026-09-22"]
keywords: ["accelerator", "amd", "arr", "benchmarks", "consumer", "cost", "decode", "gpu", "gpus", "intel", "nvidia", "pricing"]
source: docs/RAG/etape7_phaseI_media.md
source_anchor: ""
source_lines: [262, 306]
section: "Step 7 — Phase I: Media Servers, Transcoding and Upscaling"
sha256: 8367686ac7c8e0771c20e67c0b575a486edfedf49fa9adfc11824611901b71b3
---

# 10. Decision guides

## 10. Decision guides

**Pick the server:** want free + full control → Jellyfin `[independent]`; want polish + zero tinkering → Plex (pay the Pass) `[secondary]`; want middle ground → Emby Premiere lifetime $119 `[vendor-reported]`; own no files → Stremio (mind the add-on legality) `[independent]`; HTPC playback only → Kodi, ideally fronting Jellyfin `[independent]`.

**Pick the accelerator:** new build → Intel (N100/A380-class, QSV or VAAPI) `[independent]`; GPU on hand → NVIDIA NVENC (verify per-model session limits) `[independent]`; Apple TV 4K playback → Infuse + direct play, skip server transcode `[secondary]`.

**Pick the codec workflow:** maximum compatibility → H.264 `[independent]`; 4K remote bandwidth → HEVC `[independent]`; archival re-encode → SVT-AV1 slow presets offline `[independent]`; never → real-time software AV1 for streaming `[independent]`.

## 11. Conflict log

1. Jellyfin 12.0 release date: 2026-09-07 `[independent]` vs 2026-09-08 `[secondary]` — recorded unresolved.
2. Minimum upgrade base for 12.0: 10.10.7 vs 10.10.8 `[secondary]` — recorded unresolved.
3. Current Jellyfin patch (12.1 / 10.11.11) attested only by community homelab docs `[independent]` — `[unverified]` as official state.
4. NVIDIA consumer session limits: 2, 3, or 5 concurrent sessions reported across generations/drivers — no universal 2026 number; verify per model/driver `[independent]`.
5. "40–60% H.264→HEVC savings": workload-dependent secondary guidance, not a guarantee `[secondary]`.
6. VAAPI vs QSV speed on Intel: one 2026 test matrix favors VAAPI on Arc A770 — workload-specific, not universal `[independent]`.
7. Topaz Video AI pricing (2026): secondary-only, official store text not captured — `[secondary]`.
8. Plex 2026 pricing: secondary-only (MacRumors/The Register/9to5Mac), official Plex announcement not captured — `[secondary]`.
9. Emby pricing: official pricing page observed 2026-09-22 `[vendor-reported]`; the "emby.media/premieresamsung.html" URL is a Samsung-TV promo path, not the main pricing page — note the oddity, price figures match the standard Premiere tiers `[vendor-reported]`.
10. Jellyfin smart-TV apps (Feb 2026) and Swiftfin 1.4 details: secondary (jellywatch.app, smarttvs.org) — `[secondary]`.

## 12. Gaps (not researched / not claimed)

- Official Jellyfin 12.0 release notes text and the exact current 12.x patch as of 2026-09-22 `[unverified]`.
- Official Plex pricing announcement and current Plexamp/Plexamp-tier details `[unverified]`.
- Per-model NVIDIA NVENC support-matrix numbers for 2026 GPUs `[unverified]`.
- Measured VMAF/PSNR shootouts between SVT-AV1 4.2.0, x265 and hardware AV1 on identical content `[unverified]`.
- Jellyfin-on-Apple-Silicon server benchmarks `[unverified]`.
- Current *arr release versions beyond the July 2026 secondary snapshot `[unverified]`.

## 13. Glossary

- **Direct play** — client receives the file untouched; no server CPU/GPU cost `[independent]`.
- **Remux** — container/codec repackaging without re-encoding video (cheap) `[independent]`.
- **Transcode** — full decode→filter→re-encode (expensive; HDR tone mapping and subtitle burn-in force it) `[independent]`.
- **Tone mapping** — converting HDR luminance/color volume into SDR range; algorithms include Hable (filmic) and Reinhard (photographic) `[official]`/`[independent]`.
- **Media segments** — Jellyfin's first-class intro/credit/chapter markers, fed by Intro Skipper's chromaprint analysis `[independent]`.
- **NVENC/QSV/VAAPI/AMF/VideoToolbox** — vendor hardware encode APIs for NVIDIA/Intel/Linux-AMD+Intel/Windows-AMD/Apple `[official]`.
- **SVT-AV1** — Intel's production AV1 software encoder; presets 0–13 (lower = slower/better), PSY fork adds perceptual tuning `[independent]`.
- **RIFE** — open-source frame-interpolation model `[independent]`.
- **FSRCNNX / Anime4K** — real-time upscaling shader sets for mpv-class players `[independent]`.
- **AV1 / HEVC / E-AC-3 / TrueHD / DTS-HD** — codecs as used in §7 `[independent]`.

## 14. Extended reference material

