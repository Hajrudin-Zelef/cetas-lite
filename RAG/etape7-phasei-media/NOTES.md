# NOTES — corpus `etape7-phasei-media`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

La source contient 6 titres H1 : un dossier par document H1 (un entête éventuel devient `00-front-matter`).

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 7 — Phase I: Media Servers, Transcoding and Upscaling` — 750 lignes, 16 chunks.

- Provenance key
- 1. Media servers
- 2. FFmpeg and the hardware-encoder landscape
- 3. HDR and tone mapping
- 4. Transcoder configurations (Jellyfin / Plex / Emby)
- 5. Library automation: Tdarr vs Unmanic
- 6. Upscaling and frame interpolation
- 7. Audio: passthrough and transcoding
- 8. Subtitles
- 9. Homelab media patterns
- 10. Decision guides
- 11. Conflict log
- 12. Gaps (not researched / not claimed)
- 13. Glossary
- 14. Extended reference material
- 15. Line-count and QC attestation
- 16. Source index

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
