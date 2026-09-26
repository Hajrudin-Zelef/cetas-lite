---
id: etape7-phasei-media/00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape
title: "2. FFmpeg and the hardware-encoder landscape"
domain: front-matter
role: reference
task: hardware
actors: ["AMD", "Apple", "Intel", "Nvidia"]
dates: ["2024-10-28", "2025-02-18", "2025-07-24", "2026-01-13", "2026-03", "2026-03-16", "2026-03-22", "2026-03-23", "2026-05-04", "2026-07-14", "2026-09"]
keywords: ["accelerator", "amd", "compute", "consumer", "decode", "distribution", "gpu", "intel", "nvidia", "research"]
source: docs/RAG/etape7_phaseI_media.md
source_anchor: ""
source_lines: [90, 137]
section: "Step 7 — Phase I: Media Servers, Transcoding and Upscaling"
sha256: 6a086c1c79a27fc26d6c2d2b5e1b3826af83f1a78924b672d6ff37c0a3f7b35a
---

# 2. FFmpeg and the hardware-encoder landscape

## 2. FFmpeg and the hardware-encoder landscape

### 2.1 FFmpeg release chronology (relevant to media servers)

- FFmpeg 7.0 — bundled by Jellyfin 10.10.0 `[official]`; 7.1 by Jellyfin 10.11.0 `[official]`.
- **FFmpeg 8.0 "Huffman"**: Whisper filter integration, Vulkan AV1 encoding and VP9 acceleration, ProRes RAW decode, `pad_cuda`, `colordetect`, `scale_d3d11` filters, animated JPEG XL encoding, new codecs/formats `[secondary]`.
- **FFmpeg 8.1 "Hoare"** — reported released 2026-03-16 `[secondary]`: D3D12 H.264 and AV1 encoders, ProRes and DPX Vulkan acceleration, Rockchip H.264/HEVC encoder, D3D12 scale/motion-estimation/deinterlace filters, JPEG-XS via `libsvtjpegxs`, MPEG-H 3D Audio decoding, Vulkan `swscale` work and compute optimizations `[secondary]`. Jellyfin 12.0 moved its transcoding stack to FFmpeg 8.1 `[secondary]`.
- **FFmpeg 8.1.1** — reported 2026-05-04 `[secondary]`; 8.1.2 was in use by September 2026 releases such as VLC 3.0.24 `[secondary]`.
- **HandBrake 1.11** (March 2026) bundles FFmpeg 8.0.1, SVT-AV1 4.0.1 and oneVPL 2.16.0, and adds AMD VCN AV1 10-bit encoding, ProRes and DNxHR encoders/presets, and FFV1 preservation presets `[secondary]`; 1.11.1 followed 2026-03-22 `[secondary]`.

### 2.2 Hardware encode/decode paths

**NVIDIA (NVENC/NVDEC):** FFmpeg exposes `h264_nvenc`, `hevc_nvenc`, `av1_nvenc` encoders and `cuda`/`cuvid` decoders `[official]`. Containers require host NVIDIA drivers plus the NVIDIA Container Toolkit `[independent]`. Session limits vary by generation and driver: historical NVIDIA documents show two- or three-concurrent-session limits, later reporting showed some consumer GeForce limits raised to five `[independent]`/`[secondary]` — there is **no single universal 2026 number**; verify against the current model/driver-specific NVIDIA support matrix `[independent]`. Low-end NVIDIA models may lack an encoder entirely (Jellyfin hardware-selection documentation warns of this) `[official]`.

**Intel (QSV and native VAAPI):** two distinct FFmpeg paths exist even on Intel hardware — QSV (`h264_qsv`, `hevc_qsv`, `av1_qsv`) and native VAAPI (`h264_vaapi`, `hevc_vaapi`, `av1_vaapi`) `[official]`. On Linux both need `/dev/dri` access with appropriate `render`/`video` group membership `[independent]`. Intel Arc and 12th/13th-gen iGPUs need sufficiently recent Linux kernels `[official]`; F-suffix CPUs have no iGPU `[official]`; older Intel QSV SDK (libmfx/legacy) paths are being deprecated in favor of oneVPL/libvpl `[secondary]`. An independent 2026 test matrix found native VAAPI **faster than QSV** on the tested Arc A770 — workload-specific, not universal `[independent]`. Intel Arc A380 community validation reported AV1 encode with no artificial session cap — independent project evidence, not generalizable hardware claims `[independent]`.

**AMD (AMF/VAAPI):** `h264_amf`, `hevc_amf`, `av1_amf` on Windows; VAAPI on Linux (`h264_vaapi`, `hevc_vaapi`, `av1_vaapi`) `[official]`. HandBrake 1.11's AMD VCN AV1 10-bit encoder confirms current-generation AMD VCN AV1 encode support `[secondary]`.

**Apple (VideoToolbox):** `h264_videotoolbox`, `hevc_videotoolbox`, and AV1 paths exist in FFmpeg for Apple Silicon/macOS `[official]`. In server practice Apple hardware is uncommon as a Jellyfin/Plex host; VideoToolbox matters most on the client side (Swiftfin/Infuse decode) `[independent]`.

**Filter-ordering rule:** HDR-to-SDR tone-mapping chains typically require mapping/filtering **before** hardware upload (`hwupload`), e.g. `tonemap` in software then `hwupload=derive_device=...` — getting this wrong is the most common reason a hardware chain silently falls back to software `[independent]`.

### 2.3 Codec and preset guidance

**H.264/H.265 (software):** `libx264` presets `ultrafast…veryslow`; `libx265` presets `ultrafast…placebo` `[official]`. For real-time transcoding, `veryfast`/`superfast` are the practical ceiling on most CPUs; offline archival encodes use `slow`/`veryslow` `[independent]`.

**SVT-AV1:** latest stable **v4.2.0 (2026-07-14)** `[independent]`. Version line: 2.3.0 (2024-10-28) → 3.0.0 (2025-02-18, presets repositioned, max distinct preset M10, API break; do not carry tuned CRF/preset pairs across the 2.x→3.x boundary) → 3.1.0 (2025-07-24) → 4.0.0 (2026-01-13, tune 3 IQ / tune 4 MS-SSIM, `--ac-bias`, `--adaptive-film-grain`, `--enable-intrabc`) → 4.1.0 (2026-03-23) → 4.2.0 (tune 5 VMAF ≈ 15% VMAF BD-rate gain at minimal PSNR loss, `--cqp`, `--enable-kf-tf`, `initial_display_delay` for A/V seek sync) `[independent]`. Preset semantics: lower = more efficient, slower; 1–3 maximum efficiency for distribution, 4–6 the home-enthusiast balance (preset 6 ≈ x265 slow per community testing), 7–13 fast/real-time; rule of thumb "use the lowest preset that is tolerable" `[independent]`. The SVT-AV1-PSY fork adds perceptual features, presets -2/-3 for research, and subjective-quality tunes `[independent]`. Film-grain-heavy sources pair with `film-grain` / `film-grain-denoise` options (community archival presets use `film-grain=12–18:film-grain-denoise=1`) `[independent]`.

**AV1 (hardware):** `av1_nvenc` (NVIDIA), `av1_qsv` / `av1_vaapi` (Intel), `av1_amf` (AMD) `[official]`. Real-time AV1 hardware encode is practical on current generations; software SVT-AV1 is **not** a real-time streaming encoder at sane presets — it is an archival/VOD tool `[independent]`. Distinguish the two roles explicitly in any design: offline library optimization → SVT-AV1 or slow x265; live transcode → hardware AV1/HEVC/H.264 `[independent]`.

**HEVC vs AV1 vs H.264 for libraries:** community archival guidance reports ≈30–50% size savings AV1-over-H.264 and meaningful HEVC-over-H.264 savings, but treat every such number as workload-dependent secondary guidance, never a guarantee `[secondary]`. H.264 remains the maximum-compatibility transcode target (every client plays it) `[independent]`; HEVC halves bandwidth for 4K remotes where the client supports it `[independent]`; AV1 is the forward path for archival where clients (2024+ devices, browsers) decode it `[independent]`.

### 2.4 Accelerator shootout: Intel Arc/QSV vs NVIDIA vs Apple VideoToolbox

| Dimension | Intel Arc / QSV (e.g. A380, A770, N100 iGPU) | NVIDIA (NVENC, consumer GeForce) | Apple VideoToolbox (Apple Silicon) |
|---|---|---|---|
| Linux server fit | Native; `/dev/dri` passthrough trivial in Docker `[independent]` | Needs proprietary driver + container toolkit `[independent]` | Rare as server; macOS/Linux-ARM niche `[independent]` |
| AV1 encode | Yes, current Arc/iGPU; A380 validated no artificial cap `[independent]` | `av1_nvenc` on current generations `[official]` | Present in FFmpeg; server use uncommon `[official]`/`[independent]` |
| HEVC 10-bit / HDR | Yes, with tone-mapping filter chains `[independent]` | Yes `[official]` | Yes on client decode; encode uncommon in servers `[independent]` |
| Session limits | No artificial cap observed `[independent]` | Generation/driver-dependent (2–5 reported) — verify per model `[independent]` | N/A |
| Transcode perf/watt | Strong for homelab (N100/A380 class) `[independent]` | Strong; idle power higher on desktop cards `[independent]` | Efficient on client devices `[independent]` |
| Software maturity | Needs recent kernel; legacy QSV SDK deprecated `[official]` | Mature, best-documented `[independent]` | Mature on macOS/iOS `[independent]` |
| Homelab verdict | Default recommendation for a new Jellyfin/Plex build `[independent]` | Best if a GPU already exists or for AI + transcode sharing `[independent]` | Client-side decode champion (Swiftfin/Infuse) `[independent]` |

## 3. HDR and tone mapping

HDR-to-SDR tone mapping is the most expensive common real-time filter stage: it forces a software filter (`tonemap`, `libplacebo`) into the chain and is the usual reason 4K HDR transcodes stutter while SDR ones fly `[independent]`.

