---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-23
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["Nvidia", "vLLM"]
dates: ["2026-09-27"]
keywords: ["datacenter", "gpu", "benchmarks", "containment", "fp4", "hbm3", "incident", "liquid cooling", "memory", "mlperf", "nvlink", "pricing"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [3460, 3619]
sha256: d426a1e3efc3597f4b94abf842d61e3c7fdc01a364a9d1d65a05c64e882346e1
---

# Les GPU datacenter / IA — présent et futur vérifié

Forces : 288 Go, 8 To/s, FP4 natif, MLPerf à 92–104 % du B300.
Faiblesses : 1 400 W (355X), écosystème jeune, prix non publics.
Verdict : le pari ouvert — POC de 4 semaines avant tout engagement.

### 193.11. MI300A — le HPC

Forces : 128 Go unifiés CPU+GPU, zéro copie.
Faiblesses : pas de carte PCIe standard, écosystème HPC uniquement.
Verdict : supercalculateurs, pas datacenter IA généraliste.

## 194. Tableau maître : tout en une page (imprimable)

```
GPU            | VRAM     | BW      | TDP   | FP8d    | Format | Prix ~   | Usage
---------------|----------|---------|-------|---------|--------|----------|------------------
RTX 4090       | 24 GDDR6X| 1,0 To/s| 450 W | 0,33 PF | PCIe   | 2 400 $  | Lab
RTX 5090       | 32 GDDR7 | 1,8 To/s| 575 W | 0,42 PF | PCIe   | 2 600 $  | Lab+
L40S           | 48 GDDR6 | 0,9 To/s| 350 W | 0,73 PF | PCIe   | 10 000 $ | RAG/embed
RTX PRO 6000   | 96 GDDR7 | 1,8 To/s| 600 W | 0,80 PF | PCIe   | 12 000 $ | 70B 1 GPU
H100           | 80 HBM3  | 3,4 To/s| 700 W | 1,98 PF | SXM    | 30 000 $ | Training
H200           |141 HBM3e | 4,8 To/s| 700 W | 1,98 PF | SXM    | 35 000 $ | 70B+ inf
B200           |192 HBM3e | 8,0 To/s|1000 W | 4,50 PF | SXM    | 40 000 $ | Frontier
B300           |288 HBM3e | 8,0 To/s|1400 W | 7,00 PF | SXM    | n.c.     | Frontier+
MI300X         |192 HBM3  | 5,3 To/s| 750 W | 2,61 PF | OAM    | devis    | Alt H100
MI325X         |256 HBM3e | 6,0 To/s|1000 W | 2,61 PF | OAM    | devis    | Alt H200
MI350X         |288 HBM3e | 8,0 To/s|1000 W | ~4,6 PF | OAM    | devis    | Alt B200
MI355X         |288 HBM3e | 8,0 To/s|1400 W | ~5,0 PF | OAM    | devis    | Alt B200+
MI300A         |128 HBM3u | 5,3 To/s| 760 W | ~2,0 PF | APU    | devis    | HPC
```

## 195. Checklists récapitulatives (toutes en une page)

**Achat** : réf exacte §87 → 3 devis §75 → BOM §34-36 → élec §57 → froid §26 →
réseau §91 → checklist §83 → commande.
**Mise en service** : checklist §84 → burn-in §139 → DCGM §107 → runbook §171.
**Exploitation** : monitoring §170 → maintenance §62 → PRA §63 → tableau §172.
**Fin de vie** : §94 (revente/recyclage) → effacement §93 → renouvellement §169.

## 196. Derniers rappels énergie (pour Zelef)

- Un nœud GPU est une **charge industrielle** : 8–15 kW, harmoniques, appels
  de courant — pas un « gros PC ».
- Le poste froid+élec représente **30–45 % du budget** d'un projet B200+ :
  le chiffrer en même temps que les GPU, pas après.
- Le PUE se pilote : free cooling, DLC à haute température d'eau, récupération
  de chaleur — 3 leviers qui paient le surcoût DLC en 2–4 ans.
- La donnée la plus chère du datacenter IA n'est pas le GPU : c'est le
  **MWh**. Tout le reste (choix du GPU, du refroidissement, du site) en découle.
## 197. Grille d'évaluation POC (4 semaines, notée /100)

| Critère | Poids | Comment mesurer |
|---|---|---|
| tok/s sur vos modèles | 25 | vLLM, vos prompts, batch réaliste |
| Qualité (vos benchmarks) | 20 | Vos jeux de test métier, pas MMLU seul |
| Stabilité 7 j 24/7 | 15 | 0 incident bloquant, Xid = 0 |
| Facilité d'intégration | 15 | Temps pour porter votre stack |
| $/token mesuré | 15 | Facture réelle / tokens réels |
| Support vendeur | 10 | Délai de réponse sur 2 tickets tests |

**Seuil** : ≥ 70/100 pour passer en prod. Un POC qui n'échoue jamais est un
POC mal conçu — pousser jusqu'aux limites (batch max, ctx max).

## 198. Lexique français-anglais (pour les devis)

| Français | Anglais | Note |
|---|---|---|
| Carte graphique | Graphics card / GPU | — |
| Mémoire vidéo | VRAM / Device memory | Ne pas confondre avec RAM système |
| Refroidissement liquide | Liquid cooling / DLC | — |
| Plaque froide | Cold plate | — |
| Onduleur | UPS | — |
| Groupe électrogène | Genset / Generator | — |
| Disjoncteur | Circuit breaker | Courbe D pour les PSU |
| Armoire / baie | Rack / Cabinet | 19" standard, 21" OCP |
| Allée chaude/froide | Hot/cold aisle | — |
| Confinement | Containment | — |
| Rendement | Efficiency | PSU 80+ Titanium |
| Facteur de puissance | Power factor | PFC actif |
| Harmoniques | Harmonics (THD) | — |
| Mise en service | Commissioning | — |
| Réception | Acceptance | Avec PV de recette |
| Dossier des ouvrages exécutés | As-built documentation | Exiger le DOE |

## 199. Journal de bord d'exploitation (modèle)

```
Date : __/__/____    Nœud : ______    Opérateur : ______
□ Relecture alertes (DCGM, PUE, réseau)
□ Températures max GPU : ___°C (seuil 85)
□ Incidents 24 h : ___
□ Travaux prévus : ___
□ Filtres à air (si mensuel) : □ OK □ remplacés
□ Observations : ___
Signature : ______
```

Un journal papier/numérique tenu chaque semaine vaut mieux qu'un monitoring
que personne ne lit — les deux ensemble, c'est l'idéal.

## 200. Pour aller plus loin (prochaines étapes du dossier)

- **Dossier 2** : le réseau IA en détail (InfiniBand vs UEC, topologies, câblage).
- **Dossier 3** : le stockage pour l'IA (NVMe-oF, WEKA, VAST, checkpoints).
- **Dossier 4** : Kubernetes pour le GPU (operators, MIG, multi-tenant).
- **Dossier 5** : la sécurité des infrastructures IA (firmware, isolation, audits).
- **Dossier 6** : l'économie du token (pricing, mutualisation, revente de capacité).

---

*Fin du guide — 200 sections, données vérifiées le 27/09/2026.*
*Zelef : ce fichier est ta référence d'achat et d'exploitation. Imprime le §194
(tableau maître) et le §173 (arbre de décision) — le reste vit dans ton RAG.*
## 201. Questions à poser aux OEM (script d'appel)

1. « La référence exacte du GPU, avec le format (SXM/PCIe/OAM) ? »
2. « Le vBIOS est-il celui de votre marque, et quelle version ? »
3. « Les bridges NVLink sont-ils inclus ? Combien ? »
4. « Quelle est la température d'air maximale garantie en entrée ? »
5. « Le serveur est-il validé pour [ma charge : training 24/7 / inférence] ? »
6. « Délai ferme ? Clause de pénalité si dépassement ? »
7. « Garantie : durée, NBD ou 4 h, pièces incluses ? Les GPU sont-ils couverts ? »
8. « Quel est le prix avec et sans le support 3 ans ? »
9. « Reprise de mon ancien matériel possible ? »
10. « Pouvez-vous fournir un POC de 2 semaines sur site ou en lab ? »

Si le commercial hésite sur les questions 1–3, changer d'interlocuteur.

## 202. Erreurs de devis réels (vues en 2025-2026)

- **Devis « 8× H100 » sans préciser SXM ou PCIe** : écart de 40 000 $ entre
  les deux configs — le devis le moins cher n'était pas le moins cher.
- **Bridges NVLink facturés 4×** : 12 000 $ de « petit matériel » oublié au budget.
- **Garantie 1 an au lieu de 3** : le comparatif à prix égal cachait 2 ans
  de garantie en moins (45 000 $ de valeur).
- **Câbles IB en DAC pour 15 m** : ne fonctionne pas — 8 000 $ de câbles à jeter.
- **Onduleur non redondant** : un seul module 60 kVA pour 55 kW de charge —
  0 % de marge, refusé par le bureau de contrôle.

## 203. Tableau de suivi projet (modèle)

| Lot | Responsable | Budget | Dépensé | Avancement | Risque |
|---|---|---|---|---|---|
| GPU + serveurs | ___ | ___ | ___ | ___% | Délai |
| Réseau | ___ | ___ | ___ | ___% | Câblage |
| Élec (TGBT, onduleur) | ___ | ___ | ___ | ___% | Bureau de contrôle |
| Froid (DLC) | ___ | ___ | ___ | ___% | Fuites |
| Logiciel / MLOps | ___ | ___ | ___ | ___% | Compétences |
| Formation | ___ | ___ | ___ | ___% | — |

Revue hebdomadaire jusqu'à la mise en service, mensuelle ensuite.

## 204. Bureau de contrôle : ce qu'il va demander

- Note de calcul électrique (bilan de puissance signé).
- Plan d'implantation avec charges au sol (kg/m²).
- Attestation de conformité des onduleurs (CEI 62040).
- Rapport du test d'étanchéité DLC (si applicable).
- Plan d'évacuation et extincteurs adaptés (poudre ABC, pas d'eau sur l'élec).
- Carnet de maintenance prévisionnel.

