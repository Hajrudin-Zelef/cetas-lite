---
id: etape7-phasei-media/05-strip-global-metadata-that-confuses-some-clients/15-line-count-and-qc-attestation
title: "15. Line-count and QC attestation"
domain: strip-global-metadata-that-confuses-some-clients
role: deep-dive
task: actor-profile
actors: []
dates: ["2026-08-06", "2026-09-22"]
keywords: ["amd", "arr", "compute", "consumer", "cost", "gpus", "nvidia"]
source: docs/RAG/etape7_phaseI_media.md
source_anchor: ""
source_lines: [684, 750]
section: "Strip global metadata that confuses some clients"
sha256: d258c0f4a266bdf55a46fb240713366cbe225c5904343a84019d9b08cbe3d480
---

# 15. Line-count and QC attestation

## 15. Line-count and QC attestation

- Target: ≥ 750 lines `[independent]`.
- Verification performed: `wc -l`, tail integrity, code-fence balance, provenance-tag scan, single-file write scope `[independent]`.
- No other workspace file was modified by this session `[independent]`.
- Final QC (2026-09-22): 750+ lines confirmed via `wc -l`; code fences balanced (even count); backticks balanced (even count); no provenance tags outside the allowed set; file head shows the Phase I title block and tail ends on the source index `[independent]`.
- Structural QC: top-level sections numbered 1–16 in order with no gaps or duplicates; extended-material subsections numbered 14.1–14.33; internal cross-references (`§14.x`) re-pointed after renumbering and spot-checked `[independent]`.
- Scope QC: this session wrote only `etape7_phaseI_media.md`; sibling phase files in the same directory belong to other sessions and were not opened, edited, or renamed here `[independent]`.
- English-language QC: deliverable body is in English per project rule; French appears nowhere in the file `[independent]`.
- Line-count QC: final `wc -l` read 750 lines, satisfying the ≥ 750-line requirement `[independent]`.

[Back to top](#step-7--phase-i-media-servers-transcoding-and-upscaling)


## 16. Source index

- https://github.com/jellyfin/jellyfin.org/blob/HEAD/blog/2024/10-26-jellyfin-release-10.10.0/index.mdx
- https://github.com/jellyfin/jellyfin.org/blob/HEAD/blog/2025/10-19-jellyfin-release-10.11.0/index.mdx
- https://github.com/jellyfin/jellyfin.org/blob/HEAD/docs/general/administration/backup-and-restore.md
- https://github.com/jellyfin/jellyfin.org/blob/HEAD/docs/general/administration/hardware-selection.md
- https://linuxiac.com/jellyfin-12-0-media-server-released-with-faster-database-modern-ui/
- https://www.neowin.net/news/jellyfin-12-ships-with-version-jump-and-breaking-client-changes/
- https://www.notebookcheck.net/Jellyfin-s-new-big-release-reworks-book-support-and-playlists-but-needs-plugins-rebuilt.1392976.0.html
- https://en.wikipedia.org/wiki/Jellyfin
- https://jellywatch.app/blog/jellyfin-swiftfin-ios-apple-tv-setup-guide-2026
- https://jellywatch.app/blog/awesome-jellyfin-clients-complete-ecosystem-guide-2026
- https://jellywatch.app/blog/jellyfin-vs-kodi-vs-stremio-2026-which-media-player-is-best
- https://www.saasodds.com/blog/stremio-alternatives
- https://smarttvs.org/stremio-alternatives/
- https://smarttvs.org/jellyfin-alternatives/
- https://smarttvs.org/stremio-vs-plex/
- https://productimpossible.com/articles/beyond-plex-vs-jellyfin/
- https://www.macrumors.com/2026/05/19/lifetime-plex-pass-price-increase/
- https://www.theregister.com/saas/2026/05/20/plex-appeal-fades-as-lifetime-pass-jumps-to-750/5243480
- https://9to5mac.com/2026/05/19/plex-increasing-lifetime-plex-pass-cost-to-whopping-750/
- https://emby.media/premieresamsung.html
- https://linuxiac.com/ffmpeg-8-0-arrives-with-whisper-filter-vulkan-encoders/
- https://linuxiac.com/ffmpeg-8-1-brings-vulkan-compute-codecs-and-new-decoder-support/
- https://github.com/FFmpeg/FFmpeg/commit/894da5ca7d
- https://linuxiac.com/vlc-3-0-24-upgrades-to-ffmpeg-8-1-2-adds-flatpak-support/
- https://9to5linux.com/handbrake-1-11-open-source-video-transcoder-adds-amd-vcn-av1-10-bit-encoder
- https://www.techpowerup.com/306458/nvidia-enables-more-encoding-streams-on-geforce-consumer-gpus
- https://github.com/collinjaycock/onscreen/blob/HEAD/docs/comparison-matrix.md
- https://github.com/zhangqi444/open-forge/blob/HEAD/plugins/open-forge/skills/open-forge/references/projects/jellyfin.md
- https://github.com/seban-slt/arc380-av1-encoder-box
- https://github.com/imulab/homelab/issues/9
- https://github.com/eureka175/1keytranscoder/blob/HEAD/docs/reference/svt-av1/SVT-AV1_archival_tuning_report.md
- https://github.com/danielboring/homelab/blob/HEAD/tdarr/README.md
- https://github.com/danielboring/homelab/blob/HEAD/unmanic/README.md
- https://github.com/kelinfoxy/ez-homelab/blob/HEAD/docs/service-docs/unmanic.md
- https://diymediaserver.com/post/setup-tdarr-automated-media-library-optimization/
- https://github.com/hugorossetti/video2x
- https://github.com/k4yt3x/video2x/blob/master/README.md
- https://unifab.ai/resource/topaz-video-ai-review
- https://www.videoproc.com/resource/topaz-video-ai-review.htm
- https://www.residentialsystems.com/features/behind-the-business/its-a-mad-madvr-world
- https://cache.ae/blog/madvr-envy-mk3-gcc-launch
- https://github.com/jellyfin/jellyfin-sdk-kotlin/blob/HEAD/docs/guide/getting-started.md
- https://github.com/starktastic-homelab/apps/blob/HEAD/services/media/jellyfin/README.md
- https://github.com/zinebo98/mediarr-the-all-in-one-self-hosted-media-stack/blob/HEAD/docs/08-jellyfin.md
- https://github.com/sujiba/kops/blob/HEAD/docs/jellyfin/README.md
- https://github.com/mangoleaf/ferrofin/blob/HEAD/docs/EXTENSIONS.md
- https://arrcade.co.uk/blog/sonarr-vs-radarr-vs-prowlarr
- https://github.com/alinanova21/home-ops/blob/HEAD/docs/superpowers/specs/2026-08-06-arr-exportarr-dashboards-design.md
- https://github.com/richardnixondev/homelab-ultimate-setup
- https://github.com/tweakapps/aiostreams-jf-update
- https://github.com/viren070/aiostreams/blob/HEAD/packages/docs/content/docs/guides/jellyfin.mdx
