# NOTES — corpus `etape7-phasec-iac`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-iac` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 7 Phase C — IaC & Platform Automation` — 750 lignes, 12 chunks.

- Landscape map (2026)
- A. Ansible ecosystem
- B. Terraform vs OpenTofu vs Pulumi (+ Crossplane)
- C. Image building: Packer, cloud-init, golden images
- D. GitOps
- E. Secrets management
- F. Policy as code
- G. State backends, TACO platforms, drift detection
- H. Adoption numbers, case studies, market signals (2026)
- I. Conflicts, ambiguities and gaps
- J. Glossary
- K. Source index (verbatim URLs consulted)
- L. Deep dives (supplement)
- M. Reference tables and timelines (supplement 2)
- N. Final supplement
- O. Version pin quick-reference (2026-09-22)
- P. Cross-references and usage notes

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
