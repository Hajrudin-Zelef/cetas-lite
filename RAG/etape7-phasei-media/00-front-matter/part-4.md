---
id: etape7-phasei-media/00-front-matter/part-4
title: "Step 7 — Phase I: Media Servers, Transcoding and Upscaling (part 4)"
domain: front-matter
role: reference
task: model-release
actors: ["Apple"]
dates: []
keywords: ["decode", "gpu"]
source: docs/RAG/etape7_phaseI_media.md
source_anchor: ""
source_lines: [138, 149]
section: "Step 7 — Phase I: Media Servers, Transcoding and Upscaling"
sha256: 9ab44ab4ad3c390a6002585fad38b5873b548c4e279c97c9bb720ab5b00c290f
---

# Step 7 — Phase I: Media Servers, Transcoding and Upscaling (part 4)

**FFmpeg `tonemap` filter** (`-vf tonemap=...`) `[official]`:

- Algorithms: `clip`, `linear`, `gamma`, **`reinhard`** (classic photographic curve, fast, desaturates highlights), **`hable`** (filmic curve approximating filmic response, the community default for HDR→SDR), `mobius`, `bt2390` `[official]`.
- Typical usage: `tonemap=hable:desat=0` (desaturation control matters for neon/LED-heavy content) `[independent]`; for OpenCL-accelerated mapping: `tonemap_opencl=hable` after `hwupload` to the OpenCL device `[official]`.
- Jellyfin 10.10.0 added **software HDR tone mapping** as a headline feature `[official]`; community 12.1 validation checklists still include explicit "forced VA-API transcoding + HDR tone mapping" verification `[independent]`.
- Filter ordering: decode → `tonemap` (software pixel format) → format conversion → `hwupload` → hardware encode `[independent]`; doing tone mapping after `hwupload` without an OpenCL/Vulkan-capable filter fails or falls back `[independent]`.
- `libplacebo` filter offers higher-quality, GPU-shader tone mapping (needs Vulkan device) and is the quality ceiling for offline HDR→SDR conversion `[official]`/`[independent]`.

**Dolby Vision:** Jellyfin 10.10.0 improved Dolby Vision handling `[official]`; on Apple TV, Infuse provides full Dolby Vision profile support, which is why it remains the recommended Apple TV 4K HDR client over Swiftfin `[secondary]`.

**Practical rules:** (1) avoid transcoding HDR whenever direct play is possible — client capability (Infuse/Kodi) beats server tone mapping; (2) if tone mapping is unavoidable, prefer hardware decode + software `tonemap=hable` + hardware encode over full-software; (3) for archival, consider keeping a separate SDR 1080p version (multi-version support exists in Jellyfin 12) rather than tone-mapping on the fly `[independent]`.

