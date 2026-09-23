# NOTES — corpus `etape8-phasec-backend-frameworks`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-backend-frameworks` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 8 — Phase C: Backend Frameworks & APIs` — 751 lignes, 10 chunks.

- Research header
- 1. Python backend frameworks
- 2. Java / JVM backend frameworks
- 3. .NET — ASP.NET Core 9/10 and Minimal APIs
- 4. PHP backend frameworks
- 5. Ruby — Rails 8.x
- 6. Go backend frameworks
- 7. Rust backend frameworks
- 8. Elixir — Phoenix
- 9. JavaScript / TypeScript backend frameworks
- 10. API styles: REST, GraphQL, gRPC, tRPC
- 11. API contracts and auth: OpenAPI, OAuth 2.0/OIDC, JWT
- 12. ORMs and data-access layers
- 13. Benchmarks — TechEmpower and methodology limits
- 14. Comparative matrix (as of 2026-09-22)
- 15. Gaps, conflicts and unverified claims
- 16. Key takeaways for backend selection (2026-09-22)
- 17. 2026 security advisory roundup (backend-relevant)
- 18. Major-upgrade migration notes (2025–2026 generation changes)
- 19. Real-time, deployment and observability notes
- 20. Testing and API tooling notes
- 21. GraphQL and gRPC server implementations per stack
- 22. Minimal code sketches (idiomatic 2026 style)
- 23. Version and support quick reference (2026-09-22)
- 24. Selection decision guide (synthesis)
- 25. Ecosystem and community health signals (2026)
- 26. Source index (all URLs cited, verbatim)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
