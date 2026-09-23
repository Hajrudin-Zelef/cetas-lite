---
id: etape7-phasei-media/00-front-matter/4-transcoder-configurations-jellyfin-plex-emby
title: "4. Transcoder configurations (Jellyfin / Plex / Emby)"
domain: front-matter
role: reference
task: model-release
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["accelerator", "amd", "cost", "gpu", "gpus", "intel", "licenses", "nvidia", "pricing"]
source: docs/RAG/etape7_phaseI_media.md
source_anchor: ""
source_lines: [150, 200]
section: "Step 7 — Phase I: Media Servers, Transcoding and Upscaling"
sha256: b8017e5034f982082f3617f5b9f4efa4b226b1cd5663ee8215c434312ac9c167
---

# 4. Transcoder configurations (Jellyfin / Plex / Emby)

## 4. Transcoder configurations (Jellyfin / Plex / Emby)

### 4.1 Jellyfin hardware transcoding setup

1. Linux: expose the GPU to the container (`--device /dev/dri` for Intel/AMD; NVIDIA Container Toolkit for NVIDIA) and add the service user to `render`/`video` groups `[independent]`.
2. Dashboard → Playback → Transcoding: select the acceleration API (QSV, VAAPI, NVENC, AMF, VideoToolbox), enable hardware decoding for the needed codecs, and enable hardware encoding `[independent]`.
3. Prefer QSV on Intel unless testing shows native VAAPI faster on the specific Arc/iGPU + kernel combo `[independent]`.
4. Verify with an intentionally incompatible client profile (force transcode), then check the active session's transcode reason and read the actual FFmpeg transcode log — the dashboard label alone is insufficient `[independent]`.
5. `jellyfin-ffmpeg` exposes the patched filter set (including `tonemap_opencl`, chromaprint for Intro Skipper) — prefer it over distro FFmpeg for parity with logs `[independent]`.

### 4.2 Plex / Emby notes

- Plex: hardware transcoding is a Plex Pass feature `[secondary]`; select the accelerator in Settings → Transcoder; Plex ships its own FFmpeg-derived transcoder (`Plex Transcoder`) `[independent]`.
- Emby: hardware acceleration options in Transcoding settings; Premiere unlocks the full feature set `[vendor-reported]`/`[secondary]`.
- Both transcode to HLS/DASH segments similarly to Jellyfin's fmp4-HLS pipeline (Jellyfin CLI evidence shows `-f hls -hls_segment_type fmp4` style output) `[independent]`.

## 5. Library automation: Tdarr vs Unmanic

**Tdarr** — distributed transcode automation: server + node/worker model, queued jobs across hosts, library health checks (corrupt-file detection), per-library transcode templates, plugin/flow system `[independent]`. Best when: multiple machines can share the queue, or health-checking a large existing library `[independent]`. Operational requirements: staging/validation before replacing originals, free-space monitoring, checksums, tested rollback — both tools can replace originals in place `[independent]`.

**Unmanic** — plugin-oriented, simpler single-host continuous library optimization; remote-worker support also reported `[independent]`. Best when: one host, "watch folder and keep the library uniformly encoded" workflow `[independent]`.

**Decision rule:** Tdarr for fleets and health-checking; Unmanic for simplicity on one box `[independent]`. Neither replaces a backup: a transcode bug that corrupts files is a data-loss event, so automation output directories should sit under the same backup/snapshot policy as the library (see Step 7 Phase D and E) `[independent]`.

**Compression expectations:** "40–60% H.264→HEVC savings" style claims are workload-dependent secondary guidance, not guarantees `[secondary]`. Always spot-check: transcode a representative sample (grain, animation, dark scenes, sports), measure VMAF/SSIM if quality matters, and only then unleash the queue `[independent]`.

## 6. Upscaling and frame interpolation

Scope discipline: **playback-time upscaling** (madVR, FSRCNNX, Anime4K shaders) and **offline enhancement** (Topaz, Real-ESRGAN, Video2X) are different tools; frame interpolation (RIFE/SVP) is a third axis `[independent]`.

### 6.1 Playback-time (real-time) upscaling

- **FSRCNNX** — fast convolutional neural network shader for mpv/MPC; the quality-per-cost sweet spot for real-time anime/live-action upscaling on modest GPUs `[independent]`.
- **Anime4K** — shader set optimized for anime/cartoon content; runs in mpv and forks `[independent]`. Video2X 6.x also supports Anime4K shaders and custom GLSL `[independent]`.
- **madVR / madVR Envy** — madVR is the reference-quality Windows DirectShow renderer; **madVR Envy** is the hardware product line (2026: MK3/Core MK2 support HDMI 2.1, 4K120, VRR, tone mapping, subtitle handling, higher-resolution AI upscaling depending on model/package; Core MK1 reported at $4,995; high-end pricing dealer-only) `[secondary]`. Envy is a luxury home-theater appliance, not a homelab tool — include for completeness only `[independent]`.

### 6.2 Offline enhancement

- **Topaz Video AI** — commercial; secondary sources for Aug/Sep 2026 report Personal at $59/month, $39/month with annual commitment, or $299/year prepaid; Pro ≈ $699/year prepaid; new perpetual licenses reportedly discontinued in late 2025 — all pricing claims are `[secondary]` (official store text not captured). Cloud-credit mechanics likewise `[secondary]`.
- **Real-ESRGAN** — open-source GAN upscaler; strong on photographic/real-world content, weaker than Anime4K-class models on line art `[independent]`. ncnn/Vulkan accelerated; Video2X 6.x bundles it `[independent]`.
- **Real-CUGAN** — anime-specialized upscaler/denoiser; Video2X 6.x bundles it `[independent]`.
- **Video2X 6.x** — rewritten C/C++ architecture supporting Anime4K + custom GLSL shaders, Real-ESRGAN, Real-CUGAN, and **RIFE** frame interpolation, with Vulkan/ncnn acceleration on Windows and Linux `[independent]`.
- **RIFE** — open-source frame interpolation (e.g. 24→60 fps); quality is content-dependent and it multiplies output size — use deliberately, not by default `[independent]`.

### 6.3 Guidance

- Real-time: FSRCNNX/Anime4K in mpv (or Kodi with shaders) for daily watching `[independent]`.
- Archival enhancement: offline Topaz/Real-ESRGAN jobs on copies, never in place without checksums and rollback `[independent]`.
- Frame interpolation: RIFE for specific content (sports, animation); avoid for filmic 24p where soap-opera effect is unwanted `[independent]`.
- Server-side "upscaling transcodes" (e.g. 1080p→4K on the fly) are almost never worth it: bandwidth and GPU cost rise while quality barely beats client-side upscaling `[independent]`.

