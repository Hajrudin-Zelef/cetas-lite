# NOTES — corpus `etape8-phaseb-web-mobile-languages`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-web-mobile-languages` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 8 — Phase B — Web & Mobile Languages, Runtimes and Toolchains` — 762 lignes, 12 chunks.

- How to read this file
- Contents
- 1. JavaScript language standard — ECMAScript 2025 / 2026 and TC39 proposals
- 2. TypeScript — 5.8, 5.9, and the 6.0/Go-rewrite claims
- 3. Node.js — 22/24/26 lifecycle and runtime features
- 4. Deno 2.x
- 5. Bun 1.3.x
- 6. npm / pnpm / Yarn / Corepack — JS package managers and the npm registry
- 7. Supply-chain security — Shai-Hulud and registry hardening
- 8. Python — 3.13 / 3.14, free threading, JIT
- 9. Python tooling — uv, Ruff, ty, Poetry, and the Astral acquisition
- 10. PyPI statistics and packaging standards
- 11. Java — 24 / 25 LTS and the JEP landscape
- 12. JVM tooling — GraalVM, Maven, Gradle, Spring
- 13. C# / .NET — 9 / 10 and C# 14
- 14. Kotlin — 2.2 / 2.3 / 2.4 claims and KMP
- 15. Swift — 6.2 "approachable concurrency"
- 16. Dart / Flutter
- 17. PHP 8.5
- 18. Ruby 4.0 — ZJIT, Ruby Box, Ractor
- 19. Elixir 1.19
- 20. Lua 5.5 / 5.4.9
- 21. Adoption rankings 2026 — TIOBE, Stack Overflow, GitHub Octoverse
- 22. Lifecycle / support / license comparison tables
- 23. Decision guides
- 24. Conflict and gap register
- 25. Glossary
- 26. Source index
- 27. Release timeline quick reference (2025-01 → 2026-09-22)
- 28. Cross-cutting 2026 themes (interpretive synthesis)
- 29. What changed in this phase vs. prior knowledge (delta log)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
