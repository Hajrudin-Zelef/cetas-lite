---
id: etape7-phasei-media/00-front-matter/14-16-intro-skipper-tuning
title: "14.16 Intro Skipper tuning"
domain: front-matter
role: reference
task: model-release
actors: []
dates: []
keywords: ["decode", "gpu", "gpus", "latency"]
source: docs/RAG/etape7_phaseI_media.md
source_anchor: ""
source_lines: [507, 556]
section: "Step 7 — Phase I: Media Servers, Transcoding and Upscaling"
sha256: d61ba5229df13dadea980db25c1546aeeb7f382fe0488174f756f7c49f9924f4
---

# 14.16 Intro Skipper tuning

### 14.16 Intro Skipper tuning

- Detection engine: Chromaprint audio fingerprinting over episode audio; fingerprints computed by `ffmpeg -f chromaprint` (jellyfin-ffmpeg includes the muxer) with `fpcalc` fallback `[independent]`.
- Detects intros, recaps, credits, and previews; results surface as Jellyfin media segments (10.10+) with per-client skip buttons `[official]`/`[independent]`.
- Analysis is CPU-heavy on first scan — schedule off-hours on large libraries; results are cached `[independent]`.
- Versioning: plugin builds are keyed to the server ABI (1.10.11.x for 10.11, 12.0.4.0 for 12.x); the manifest at `https://manifest.intro-skipper.org/manifest.json` selects per server version — installing a mismatched build breaks startup `[independent]`.
- Alternative: Chapter Segments Provider (official) exposes chapter-based segments without audio analysis `[independent]`.

### 14.17 Upscaler model notes

- **Topaz Video AI models (2026):** Proteus (general enhancement), Iris (face recovery), Nyx (denoise), Rhea (detail), Apollo/Chronos (frame interpolation), Aion (newer interpolation) — model choice dominates output quality more than settings `[secondary]`.
- **Real-ESRGAN models:** `realesr-general-x4v3` (general), `realesrgan-x4plus` (photo), `realesrgan-x4plus-anime_6B` (anime-optimized) `[independent]`.
- **Real-CUGAN:** SE (fast) vs Pro (quality) variants; 2x/3x/4x scales; strong on anime denoising+upscale `[independent]`.
- **Anime4K shaders:** `Anime4K_Upscale_CNN_L/X` + `Denoise` + `Restore` chains in mpv; quality scales with GPU — integrated GPUs handle CNN_M, discrete GPUs handle CNN_X/UL `[independent]`.
- **Offline workflow rule:** enhance a *copy*, compare frames (original vs enhanced) before batching, keep the original — upscaling artifacts (hallucinated detail, temporal flicker) are irreversible `[independent]`.

### 14.18 FFmpeg 8.x features relevant to media servers

- **Whisper filter** (`-vf whisper`): speech-to-text subtitle *generation* inside FFmpeg — enables auto-subtitle pipelines for libraries without subs `[secondary]`.
- **`colordetect` filter:** detects color range/primaries metadata — useful in automation to route HDR vs SDR files differently `[secondary]`.
- **`pad_cuda`:** GPU-side padding for hardware chains (avoids download/upload round-trips) `[secondary]`.
- **Vulkan AV1 encode / VP9 decode acceleration:** growing non-vendor-specific GPU path; `libplacebo` benefits from the same Vulkan device setup `[secondary]`.
- **D3D12 H.264/AV1 encoders (8.1):** Windows-server hardware encoding without vendor SDKs `[secondary]`.
- **JPEG-XS via libsvtjpegxs:** mezzanine codec for low-latency contribution — niche for homelab, relevant for production pipelines `[secondary]`.

### 14.19 Audio normalization and loudness

- **EBU R128 / loudnorm:** two-pass loudness normalization (`-af loudnorm`) for consistent volume across episodes — an offline Unmanic/Bazarr-adjacent step, not a real-time default `[official]`/`[independent]`.
- **Dynamic Range Compression:** FFmpeg `acompressor`/`dynaudnorm` for night-mode style playback; client-side (Kodi/Infuse) DRC settings are preferable to baking it into files `[independent]`.
- **Language/track hygiene:** set default and forced flags correctly at ingest (`-disposition`), so clients pick the right audio/subtitle without user intervention — the cheapest "transcode avoidance" there is `[independent]`.

### 14.20 Worked storage-sizing example

Assumptions: 500 movies + 40 series (20 eps each) = 1,300 titles `[independent]`:

| Profile | Avg/title | Library | +20% headroom | Notes |
|---|---|---|---|---|
| 1080p H.264 | 8 GB | 10.4 TB | ~12.5 TB | Max compatibility |
| 1080p HEVC (Tdarr) | 4.5 GB | 5.9 TB | ~7 TB | ~45% saving, workload-dependent |
| 4K HEVC remux mix | 35 GB | 45.5 TB | ~55 TB | Enthusiast 4K |
| 4K AV1 re-encode | 18 GB | 23.4 TB | ~28 TB | Offline SVT-AV1 |

All sizes `[independent]` estimates; remux collections scale linearly and dominate planning. Transcode/metadata SSD: 256–512 GB is ample for most libraries `[independent]`.

### 14.21 Jellyfin API automation snippets

API key via Dashboard → API Keys; header `X-Emby-Token` (legacy name retained) `[independent]`:

```bash
J=http://jellyfin:8096; K=$JELLYFIN_API_KEY
