---
id: etape7-phasei-media/00-front-matter/14-5-svt-av1-command-patterns
title: "14.5 SVT-AV1 command patterns"
domain: front-matter
role: reference
task: model-release
actors: ["Apple", "Intel"]
dates: []
keywords: ["decode", "gpu", "intel"]
source: docs/RAG/etape7_phaseI_media.md
source_anchor: ""
source_lines: [363, 410]
section: "Step 7 — Phase I: Media Servers, Transcoding and Upscaling"
sha256: df167465aa9c6454c844732fa0fdab3edb3dea4e741c32c91eca5f4d6a1019f5
---

# 14.5 SVT-AV1 command patterns

### 14.5 SVT-AV1 command patterns

Software SVT-AV1 via FFmpeg (archival, NOT real-time) `[official]`/`[independent]`:

```
ffmpeg -i input.mkv -c:v libsvtav1 -preset 6 -crf 28 -pix_fmt yuv420p10le \
 -svtav1-params tune=0:film-grain=8 -c:a copy output.mp4
```

- `-preset 4–6`: enthusiast archival balance; `-preset 8+`: fast near-real-time `[independent]`.
- `-crf 26–32`: typical 1080p film range; animation tolerates higher CRF `[independent]`.
- `tune=0` (VQ), `tune=1` (PSNR default), `tune=3` (IQ), `tune=4` (MS-SSIM), `tune=5` (VMAF, v4.2.0+) `[independent]`.
- `--keyint`/gop: default ~5 s (`keyint -2`); for hardware-decode compatibility keep ≤ 300 `[independent]`.
- Hardware AV1 for real-time: `-c:v av1_nvenc -preset p5` / `-c:v av1_qsv` / `-c:v av1_vaapi` — preset scales differ per vendor; test on target hardware `[official]`/`[independent]`.

### 14.6 HDR tone-mapping command patterns

Software Hable (quality baseline) `[official]`/`[independent]`:

```
-vf "tonemap=hable:desat=0,zscale=t=linear,format=yuv420p10le"
```

Intel QSV full-hardware HDR→SDR with OpenCL tone mapping `[independent]`:

```
-init_hw_device qsv=hw -filter_hw_device hw \
 -vf "hwupload=extra_hw_frames=64,tonemap_opencl=hable:desat=0:format=nv12,hwdownload,format=nv12" \
 -c:v h264_qsv output.mp4
```

Higher-quality offline path: `libplacebo` with Vulkan device (`-init_hw_device vulkan`) `[official]`.

**Dolby Vision profiles:** profile 5 (streaming, IPTPQc2) vs profiles 7/8 (disc-derived, BL+EL/RPU layers) — client support varies wildly; Infuse handles all profiles on Apple TV 4K, most other clients do not `[secondary]`. When in doubt, keep an HDR10 base layer `[independent]`.

### 14.7 Tdarr operational patterns

- **Staging pattern:** input watch folder → Tdarr transcode → output folder → checksum/verification → atomic move over original; never transcode in place without a verified copy `[independent]`.
- **Flow example:** health check (error detection) → transcode H.264→HEVC (QSV/NVENC) → strip unwanted audio/subtitle tracks → set title/language metadata → move to output `[independent]`.
- **Node scaling:** one server, N workers; workers can be CPU-only boxes doing health checks while GPU workers transcode `[independent]`.
- **Scheduling:** run GPU transcode queues overnight; health checks anytime — transcode queues saturate disks and contend with evening streaming `[independent]`.

### 14.8 Unmanic operational patterns

- **Plugin pipeline:** file test (is it already compliant?) → transcode (ffmpeg wrapper with configurable args) → post-processor (move, notify) `[independent]`.
- **Typical single-host config:** watch library root, convert everything to HEVC 10-bit + AAC stereo fallback track, keep original audio, extract SRT sidecars `[independent]`.
- Remote workers exist for spreading load, but the single-host story is the product's center of gravity `[independent]`.

