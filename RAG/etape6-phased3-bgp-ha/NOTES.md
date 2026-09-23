# NOTES — corpus `etape6-phased3-bgp-ha`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-bgp-ha` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Phase D3 — BGP underlay and high availability in the data center` — 755 lignes, 13 chunks.

- Provenance legend
- Wave 1 — BGP as the data-center underlay: design patterns, addressing, and tuning
- Wave 2 — Route policy in the DC fabric: prefix lists, route maps, communities, RPKI
- Wave 3 — Multi-chassis LAG: per-vendor architectures, control planes, and failure behavior
- Wave 4 — First-hop redundancy, anycast gateway, ISSU/NSF, and convergence
- Wave 5 — File verification and coverage audit
- Wave 6 — BGP session security, dynamic peering, and EVPN overlay control-plane design
- Wave 7 — EVPN multihoming vs MC-LAG, STP interplay, and L3 over MC-LAG
- Wave 8 — Detection extras, FHRP tracking, control-plane protection, and GR procedures
- Wave 9 — BGP PIC, best-path determinism, RR scaling, and DCI border design
- Wave 10 — Monitoring, telemetry, and the DC BGP troubleshooting checklist
- Appendix — Consolidated open items (all waves)
- Wave 11 — Reference configs, timer tables, and failure walkthroughs
- Wave 12 — BGP in AI fabrics, document control, and close
- Wave 13 — Standards quick reference and session-scale guidance

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
