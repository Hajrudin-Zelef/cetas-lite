# INDEX — Step 7 — Phase I: Media Servers, Transcoding and Upscaling

Corpus `etape7-phasei-media` · **18 fichiers** · 750 lignes source · ~9240 mots · partition exacte de `docs/RAG/etape7_phaseI_media.md`.

## Mode d'emploi

1. Filtrer dans `manifest.json` (ou les tableaux ci-dessous) sur `domain`, `task`, `actors`, `dates` ou `keywords`.
2. Ouvrir 1 à 3 fichiers ciblés ; chaque fichier est une unité thématique auto-suffisante avec un en-tête YAML.
3. Pour un événement répété dans plusieurs sections, préférer le fichier marqué `canonical_for` (voir la table Événements canoniques).

## Domaines (dossiers → fichiers)

### `00-front-matter/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [Step 7 — Phase I: Media Servers, Transcoding and Upscaling](00-front-matter/overview.md) | 1–54 | reference | model-release |
| 02 | [1.3 Emby](00-front-matter/1-3-emby.md) | 55–89 | reference | model-release |
| 03 | [2. FFmpeg and the hardware-encoder landscape](00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md) | 90–137 | reference | hardware |
| 04 | [Step 7 — Phase I: Media Servers, Transcoding and Upscaling (part 4)](00-front-matter/part-4.md) | 138–149 | reference | model-release |
| 05 | [4. Transcoder configurations (Jellyfin / Plex / Emby)](00-front-matter/4-transcoder-configurations-jellyfin-plex-emby.md) | 150–200 | reference | model-release |
| 06 | [7. Audio: passthrough and transcoding](00-front-matter/7-audio-passthrough-and-transcoding.md) | 201–261 | reference | model-release |
| 07 | [10. Decision guides](00-front-matter/10-decision-guides.md) | 262–306 | reference | model-release |
| 08 | [14.1 Jellyfin internals worth knowing](00-front-matter/14-1-jellyfin-internals-worth-knowing.md) | 307–362 | reference | model-release |
| 09 | [14.5 SVT-AV1 command patterns](00-front-matter/14-5-svt-av1-command-patterns.md) | 363–410 | reference | model-release |
| 10 | [14.9 *arr stack reference deployment](00-front-matter/14-9-arr-stack-reference-deployment.md) | 411–455 | reference | model-release |
| 11 | [14.11 Transcode troubleshooting matrix](00-front-matter/14-11-transcode-troubleshooting-matrix.md) | 456–506 | reference | model-release |
| 12 | [14.16 Intro Skipper tuning](00-front-matter/14-16-intro-skipper-tuning.md) | 507–556 | reference | model-release |

### `01-list-installed-plugins-and-versions/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [List installed plugins and versions](01-list-installed-plugins-and-versions/overview.md) | 557–559 | deep-dive | reference |

### `02-trigger-a-library-scan/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [Trigger a library scan](02-trigger-a-library-scan/overview.md) | 560–561 | deep-dive | reference |

### `03-active-sessions-see-transcode-reasons/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [Active sessions (see transcode reasons)](03-active-sessions-see-transcode-reasons/overview.md) | 562–604 | deep-dive | reference |

### `04-copy-chapters-from-source-when-remuxing/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [Copy chapters from source when remuxing](04-copy-chapters-from-source-when-remuxing/overview.md) | 605–606 | deep-dive | reference |

### `05-strip-global-metadata-that-confuses-some-clients/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [Strip global metadata that confuses some clients](05-strip-global-metadata-that-confuses-some-clients/overview.md) | 607–683 | deep-dive | actor-profile |
| 02 | [15. Line-count and QC attestation](05-strip-global-metadata-that-confuses-some-clients/15-line-count-and-qc-attestation.md) | 684–750 | deep-dive | actor-profile |

## Par tâche

- **actor-profile** — [Strip global metadata that confuses some clients](05-strip-global-metadata-that-confuses-some-clients/overview.md), [15. Line-count and QC attestation](05-strip-global-metadata-that-confuses-some-clients/15-line-count-and-qc-attestation.md)
- **hardware** — [2. FFmpeg and the hardware-encoder landscape](00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md)
- **model-release** — [Step 7 — Phase I: Media Servers, Transcoding and Upscaling](00-front-matter/overview.md), [1.3 Emby](00-front-matter/1-3-emby.md), [Step 7 — Phase I: Media Servers, Transcoding and Upscaling (part 4)](00-front-matter/part-4.md), [4. Transcoder configurations (Jellyfin / Plex / Emby)](00-front-matter/4-transcoder-configurations-jellyfin-plex-emby.md), [7. Audio: passthrough and transcoding](00-front-matter/7-audio-passthrough-and-transcoding.md), [10. Decision guides](00-front-matter/10-decision-guides.md), [14.1 Jellyfin internals worth knowing](00-front-matter/14-1-jellyfin-internals-worth-knowing.md), [14.5 SVT-AV1 command patterns](00-front-matter/14-5-svt-av1-command-patterns.md), [14.9 *arr stack reference deployment](00-front-matter/14-9-arr-stack-reference-deployment.md), [14.11 Transcode troubleshooting matrix](00-front-matter/14-11-transcode-troubleshooting-matrix.md), [14.16 Intro Skipper tuning](00-front-matter/14-16-intro-skipper-tuning.md)
- **reference** — [List installed plugins and versions](01-list-installed-plugins-and-versions/overview.md), [Trigger a library scan](02-trigger-a-library-scan/overview.md), [Active sessions (see transcode reasons)](03-active-sessions-see-transcode-reasons/overview.md), [Copy chapters from source when remuxing](04-copy-chapters-from-source-when-remuxing/overview.md)

## Par acteur

- **AMD** (4) — [00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md](00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md), [00-front-matter/4-transcoder-configurations-jellyfin-plex-emby.md](00-front-matter/4-transcoder-configurations-jellyfin-plex-emby.md), [00-front-matter/10-decision-guides.md](00-front-matter/10-decision-guides.md), [00-front-matter/14-9-arr-stack-reference-deployment.md](00-front-matter/14-9-arr-stack-reference-deployment.md)
- **AWS** (1) — [00-front-matter/overview.md](00-front-matter/overview.md)
- **Apple** (9) — [00-front-matter/overview.md](00-front-matter/overview.md), [00-front-matter/1-3-emby.md](00-front-matter/1-3-emby.md), [00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md](00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md), [00-front-matter/part-4.md](00-front-matter/part-4.md), [00-front-matter/7-audio-passthrough-and-transcoding.md](00-front-matter/7-audio-passthrough-and-transcoding.md), [00-front-matter/10-decision-guides.md](00-front-matter/10-decision-guides.md), [00-front-matter/14-5-svt-av1-command-patterns.md](00-front-matter/14-5-svt-av1-command-patterns.md), [00-front-matter/14-11-transcode-troubleshooting-matrix.md](00-front-matter/14-11-transcode-troubleshooting-matrix.md), [05-strip-global-metadata-that-confuses-some-clients/overview.md](05-strip-global-metadata-that-confuses-some-clients/overview.md)
- **Google** (1) — [00-front-matter/overview.md](00-front-matter/overview.md)
- **Intel** (7) — [00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md](00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md), [00-front-matter/4-transcoder-configurations-jellyfin-plex-emby.md](00-front-matter/4-transcoder-configurations-jellyfin-plex-emby.md), [00-front-matter/10-decision-guides.md](00-front-matter/10-decision-guides.md), [00-front-matter/14-1-jellyfin-internals-worth-knowing.md](00-front-matter/14-1-jellyfin-internals-worth-knowing.md), [00-front-matter/14-5-svt-av1-command-patterns.md](00-front-matter/14-5-svt-av1-command-patterns.md), [00-front-matter/14-9-arr-stack-reference-deployment.md](00-front-matter/14-9-arr-stack-reference-deployment.md), [05-strip-global-metadata-that-confuses-some-clients/overview.md](05-strip-global-metadata-that-confuses-some-clients/overview.md)
- **Nvidia** (5) — [00-front-matter/overview.md](00-front-matter/overview.md), [00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md](00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md), [00-front-matter/4-transcoder-configurations-jellyfin-plex-emby.md](00-front-matter/4-transcoder-configurations-jellyfin-plex-emby.md), [00-front-matter/10-decision-guides.md](00-front-matter/10-decision-guides.md), [05-strip-global-metadata-that-confuses-some-clients/overview.md](05-strip-global-metadata-that-confuses-some-clients/overview.md)
- **Samsung** (5) — [00-front-matter/overview.md](00-front-matter/overview.md), [00-front-matter/1-3-emby.md](00-front-matter/1-3-emby.md), [00-front-matter/10-decision-guides.md](00-front-matter/10-decision-guides.md), [00-front-matter/14-11-transcode-troubleshooting-matrix.md](00-front-matter/14-11-transcode-troubleshooting-matrix.md), [05-strip-global-metadata-that-confuses-some-clients/overview.md](05-strip-global-metadata-that-confuses-some-clients/overview.md)

## Par date

- **2018-12** — [00-front-matter/overview.md](00-front-matter/overview.md)
- **2019-01** — [00-front-matter/overview.md](00-front-matter/overview.md)
- **2024-05-11** — [00-front-matter/overview.md](00-front-matter/overview.md)
- **2024-10-26** — [00-front-matter/overview.md](00-front-matter/overview.md)
- **2024-10-27** — [00-front-matter/overview.md](00-front-matter/overview.md)
- **2024-10-28** — [00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md](00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md)
- **2025-02-18** — [00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md](00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md)
- **2025-07-24** — [00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md](00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md)
- **2025-10-19** — [00-front-matter/overview.md](00-front-matter/overview.md)
- **2025-10-20** — [00-front-matter/overview.md](00-front-matter/overview.md)
- **2026-01** — [00-front-matter/overview.md](00-front-matter/overview.md)
- **2026-01-13** — [00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md](00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md)
- **2026-02** — [00-front-matter/overview.md](00-front-matter/overview.md), [00-front-matter/14-11-transcode-troubleshooting-matrix.md](00-front-matter/14-11-transcode-troubleshooting-matrix.md)
- **2026-03** — [00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md](00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md)
- **2026-03-16** — [00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md](00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md)
- **2026-03-22** — [00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md](00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md)
- **2026-03-23** — [00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md](00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md)
- **2026-05-04** — [00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md](00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md)
- **2026-07** — [00-front-matter/7-audio-passthrough-and-transcoding.md](00-front-matter/7-audio-passthrough-and-transcoding.md), [00-front-matter/10-decision-guides.md](00-front-matter/10-decision-guides.md)
- **2026-07-01** — [00-front-matter/overview.md](00-front-matter/overview.md)
- **2026-07-14** — [00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md](00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md)
- **2026-08-06** — [05-strip-global-metadata-that-confuses-some-clients/15-line-count-and-qc-attestation.md](05-strip-global-metadata-that-confuses-some-clients/15-line-count-and-qc-attestation.md)
- **2026-09** — [00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md](00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md), [05-strip-global-metadata-that-confuses-some-clients/overview.md](05-strip-global-metadata-that-confuses-some-clients/overview.md)
- **2026-09-07** — [00-front-matter/overview.md](00-front-matter/overview.md), [00-front-matter/10-decision-guides.md](00-front-matter/10-decision-guides.md)
- **2026-09-08** — [00-front-matter/overview.md](00-front-matter/overview.md), [00-front-matter/10-decision-guides.md](00-front-matter/10-decision-guides.md)
- **2026-09-17** — [00-front-matter/overview.md](00-front-matter/overview.md)
- **2026-09-22** — [00-front-matter/overview.md](00-front-matter/overview.md), [00-front-matter/1-3-emby.md](00-front-matter/1-3-emby.md), [00-front-matter/10-decision-guides.md](00-front-matter/10-decision-guides.md), [05-strip-global-metadata-that-confuses-some-clients/15-line-count-and-qc-attestation.md](05-strip-global-metadata-that-confuses-some-clients/15-line-count-and-qc-attestation.md)

## Carte de couverture (lignes source)

| plage | fichier |
|---|---|
| 1–54 | etape7-phasei-media/00-front-matter/overview.md |
| 55–89 | etape7-phasei-media/00-front-matter/1-3-emby.md |
| 90–137 | etape7-phasei-media/00-front-matter/2-ffmpeg-and-the-hardware-encoder-landscape.md |
| 138–149 | etape7-phasei-media/00-front-matter/part-4.md |
| 150–200 | etape7-phasei-media/00-front-matter/4-transcoder-configurations-jellyfin-plex-emby.md |
| 201–261 | etape7-phasei-media/00-front-matter/7-audio-passthrough-and-transcoding.md |
| 262–306 | etape7-phasei-media/00-front-matter/10-decision-guides.md |
| 307–362 | etape7-phasei-media/00-front-matter/14-1-jellyfin-internals-worth-knowing.md |
| 363–410 | etape7-phasei-media/00-front-matter/14-5-svt-av1-command-patterns.md |
| 411–455 | etape7-phasei-media/00-front-matter/14-9-arr-stack-reference-deployment.md |
| 456–506 | etape7-phasei-media/00-front-matter/14-11-transcode-troubleshooting-matrix.md |
| 507–556 | etape7-phasei-media/00-front-matter/14-16-intro-skipper-tuning.md |
| 557–559 | etape7-phasei-media/01-list-installed-plugins-and-versions/overview.md |
| 560–561 | etape7-phasei-media/02-trigger-a-library-scan/overview.md |
| 562–604 | etape7-phasei-media/03-active-sessions-see-transcode-reasons/overview.md |
| 605–606 | etape7-phasei-media/04-copy-chapters-from-source-when-remuxing/overview.md |
| 607–683 | etape7-phasei-media/05-strip-global-metadata-that-confuses-some-clients/overview.md |
| 684–750 | etape7-phasei-media/05-strip-global-metadata-that-confuses-some-clients/15-line-count-and-qc-attestation.md |

