# NOTES — corpus `etape8-phasea-system-languages`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-system-languages` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 8 — Phase A: Systems Languages` — 752 lignes, 12 chunks.

- C, C++, Rust, Go, Zig, Carbon, Mojo, Julia (and V as secondary)
- 0. Method and provenance legend
- Table of contents
- 1. C
- 2. C++
- 3. Rust
- 4. Go
- 5. Zig
- 6. Carbon
- 7. Mojo
- 8. Julia
- 9. V (secondary coverage)
- 10. Memory safety: cross-language comparison
- 11. Build systems and package management
- 12. Benchmarks and performance signals
- 13. Rankings, adoption, and job-market signals
- 14. Learning resources
- 15. Decision matrix
- 16. Gaps and conflicts register
- 17. Glossary
- 18. Source index
- Appendix A: Chronological timeline (2025-02 → 2026-09)
- Appendix B: Toolchain support matrix (at the cutoff)
- Appendix C: Stability and compatibility policies compared
- Appendix D: Open research questions (post-cutoff watch list)
- Verification

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
