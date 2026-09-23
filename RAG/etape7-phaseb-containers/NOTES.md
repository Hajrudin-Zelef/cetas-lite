# NOTES — corpus `etape7-phaseb-containers`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-containers` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 7 — Phase B: Containers, Orchestration & Sandbox Runtimes` — 796 lignes, 12 chunks.

- Provenance legend
- 1. Container runtimes
- 2. Desktop and local development tooling
- 3. Podman
- 4. Compose and build tooling
- 5. Kubernetes — upstream releases
- 6. Kubernetes distributions
- 7. Packaging and deployment: Helm, Kustomize, operators
- 8. System containers: LXC, LXD, Incus
- 9. Sandbox and microVM runtimes
- 10. Container registries
- 11. Kubernetes networking: CNI, load balancing
- 12. Storage: CSI drivers
- 13. Service mesh
- 14. GPUs in containers
- 15. Ecosystem, adoption, and operations
- 16. Conflicts, gaps, and unverified claims
- 17. Glossary
- 18. Source index (verbatim URLs)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
