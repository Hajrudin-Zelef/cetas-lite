# NOTES — corpus `etape6-tracka-cisco-juniper`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**. La source contient
6 titres H1 : un entête de document (`00-front-matter`) puis trois rapports et leurs
séparateurs. Les séparateurs `# PART n — …` (titre seul, sans texte) sont **rattachés au
rapport suivant** : dossiers `01-part-1-cisco`, `02-part-2-juniper-networks-post-hpe-acquisition`,
`03-part-3-cross-vendor-synthesis`.

Source sans ancre HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

« Step 6 — Track A : Cisco + Juniper (réseau datacenter entreprise) », recherche au
22 septembre 2026.

- Cisco : Silicon One (G300 et roadmap), Nexus 9000, AI PODs, résultats financiers
  (commandes/revenus infra IA), lancements 2026, acquisitions et partenariats,
  portefeuille 400G/800G et optique/CPO, positionnement vs concurrents, chronologie,
  lacunes, sources ;
- Juniper Networks (post-rachat par HPE) : statut de l'acquisition, lancements 2026,
  revenus réseau IA et clients, stratégie, positionnement vs Cisco, optique/CPO,
  incertitudes, sources ;
- synthèse inter-constructeurs et chronologie maîtresse.

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
