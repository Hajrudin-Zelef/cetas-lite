---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-26
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["datacenter", "arr", "capex", "cpo", "gpu", "incident", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [3853, 3907]
sha256: aa4e893568179cddf9442a4c0e9f84410cd056c903e3af290a7ca525f4d39f24
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

| # | Erreur | Coût estimé |
|---|---|---|
| 1 | Fabric unique IA+prod (ralentissements permanents) | 500 k€-2 M€/an de GPU idle |
| 2 | FW partiel sur fabric RoCE (−30 % pendant 4 j) | 200-800 k€ |
| 3 | Pas de test ECMP (polarisation découverte en prod) | 100-500 k€ |
| 4 | Fibre sale généralisée (chantier + pas de nettoyage) | 50-200 k€ |
| 5 | DAC trop longs (recommande en urgence ×2) | 50-150 k€ |
| 6 | Config non versionnée (rollback impossible) | 50-300 k€/incident |
| 7 | UPS réseau = UPS GPU (arrêt à l'aveugle) | 100-400 k€/incident |
| 8 | MTU incohérent (fragmentation silencieuse) | 50-200 k€ |
| 9 | Pas de spares (MTTR ×10) | 100-500 k€/an |
| 10 | Achat CAPEX seul (TCO ×2 sur 5 ans) | 1-5 M€ |

**Total** : les erreurs « humaines/process » coûtent **plus cher que le
matériel**. D'où les checklists, les templates et les tests de ce guide.

## 185. Pour aller plus loin — ressources

| Ressource | Type | Usage |
|---|---|---|
| Docs NVIDIA Networking | Doc officielle | ConnectX, Spectrum-X, DOCA |
| Docs Arista (EOS Central) | Doc officielle | EOS, CloudVision |
| UEC (ultraethernet.org) | Consortium | Spec 1.0/1.0.1 |
| SONiC (sonic-net.github.io) | Open-source | NOS whitebox |
| RFC 7938 (BGP datacenter) | RFC | Underlay |
| RFC 7432 (EVPN) | RFC | Overlay |
| ServeTheHome / Next Platform | Presse | Tests, analyses |
| Guides Zelef | Local | Onduleurs, Proxmox, Debian, NetBox, Ansible |

## 186. Revue de conception — la checklist du design review

Avant de signer un design réseau, 20 questions :

1. Le ratio est-il **calculé** (pas promis) ? 2. Backend et front-end
séparés ? 3. PCIe validé pour chaque NIC ? 4. Fibre : type, longueurs,
budgets ≥ 3 dB ? 5. Breakout validé par port ? 6. ECN/DCQCN : qui règle ?
7. PFC : par file seulement ? 8. Buffers : shallow leaf / deep spine ?
9. ECMP : hash seed aléatoire ? 10. Tests de réception chiffrés au contrat ?
11. FW/driver figés ? 12. Supervision : BER, ECN, PFC, thermique ?
13. Spares : 10 % optiques, 1 switch/plan ? 14. SLA : 4 h sur l'IA ?
15. Énergie : switch + optiques + NIC + PUE chiffrés ? 16. UPS réseau
séparé ? 17. CPO : clause de sortie du pluggable ? 18. UEC-ready exigé ?
19. Croissance : +20 % de ports, fibre surdimensionnée ? 20. **Qui opère,
avec quelles compétences, formé quand ?**

Si une réponse manque : le design n'est pas fini.

---

*Fin du guide — **186 sections numérotées**. Objectif : ≥ 4000 lignes vérifiées par `wc -l`.*

# PARTIE S — 110 DÉCISIONS RAPIDES « SI / ALORS »

## 187. Le mémo ultime — une ligne par décision

