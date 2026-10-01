---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-49
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "China", "DeepSeek", "Huawei", "Mistral", "Nvidia", "United States"]
dates: ["2025-09-01", "2026-01-01", "2026-05-14", "2027-01-01"]
keywords: ["agents", "ascend", "benchmarks", "claude", "compute", "cyber", "deepseek", "mai", "mistral", "nvidia", "valuation"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [3792, 3871]
sha256: 3d46f764dc382568d80a77e1c79ce12b42d286c2080edeecd8c043c37b3ebb8f
---

# IA — Le grand dossier

**Contexte politique** : l'UE investit parallèlement dans le « compute souverain » et soutient des
champions domestiques (ex. : **Mistral**) pour réduire la dépendance aux hyperscalers US, tout en
maintenant la sévérité réglementaire (source : globalnewsnetwork11.com, sept. 2026).

#### 2.4.2. États-Unis — pas de loi fédérale, un patchwork d'États, un bras de fer

**Le tableau fédéral :**

- **Aucune loi fédérale globale sur l'IA.** L'executive order 14110 de Biden (sécurité de l'IA) a été **abrogé le 20 janvier 2025** ; l'EO 14179 l'a remplacé par un cap pro-innovation allégeant les exigences de test/reporting pour les développeurs de modèles frontières.
- **11 décembre 2025 — EO 14365** « Ensuring a National Policy Framework for Artificial Intelligence » (Trump) : cap « minimally burdensome », création d'une **AI Litigation Task Force au DOJ** (janvier 2026, procureure générale Pam Bondi) chargée d'attaquer en justice les lois d'États jugées contraires à la politique fédérale, évaluation par le département du Commerce de toutes les lois d'États (90 jours), et conditionnement d'environ **21 Md$ de fonds broadband (BEAD)** non encore déboursés au renoncement des États à leurs régulations « onéreuses ». Le Sénat avait rejeté en juillet 2025 (99-1) un moratoire fédéral de 10 ans sur la régulation des États.
- **Juin 2026 — EO 14409** « Promoting Advanced AI Innovation and Security » : cadre *volontaire* pour les modèles frontières (accès pré-lancement pour tests gouvernementaux, benchmarks cyber via le NIST) — **explicitement pas un régime de licence**.
- **TAKE IT DOWN Act** (P.L. 119-12, signé le 19 mai 2025) : première loi fédérale ciblant les contenus générés par IA (images intimes non consenties, deepfakes) — jusqu'à 2 ans de prison (3 ans si mineur), retrait sous 48 h par les plateformes.

**Le patchwork des États (sélection vérifiée) :**

| État | Loi | Contenu / statut |
|---|---|---|
| Californie | **SB 53** « Transparency in Frontier AI Act » | En vigueur au **01/01/2026** ; transparence et reporting de sécurité pour les développeurs de modèles au-delà d'un seuil de calcul (~10^26 FLOPs) |
| New York | **RAISE Act** | Reporting des incidents de sécurité critiques |
| Texas | **TRAIGA** | En vigueur **janv. 2026** ; interdit des usages nuisibles ciblés (approche par l'intention) |
| Colorado | AI Act (SB24-205) → **abrogé avant application** ; remplacé par **SB 26-189** (signé 14/05/2026) | Nouveau texte resserré sur les décisions automatisées (ADMT) : notice/consommateur plutôt qu'évaluations d'impact ; **effectif au 01/01/2027** |
| Illinois | HB 3773 | En vigueur **janv. 2026** : usage discriminatoire de l'IA dans l'emploi = violation des droits civils |
| New York City | Local Law 144 | Audit annuel indépendant des biais des outils de recrutement automatisés |

**Lecture** : les US ont choisi la vitesse d'innovation + la régulation sectorielle/étatique, au prix
d'une complexité de conformité « 50 régimes » dénoncée par la Maison-Blanche elle-même. Pour un
éditeur européen, c'est le miroir inverse de l'UE.

#### 2.4.3. Chine — réguler par couches, pas par grande loi

**Pas de loi-cadre unique** : au 31 août 2026, aucune « loi sur l'IA » ne figure au programme
législatif de l'Assemblée nationale populaire (vérifié via le tracker du 14e NPC — source :
casrai.org). La Chine régule par **mesures administratives sectorielles** du CAC (Cyberspace
Administration of China) et des co-régulateurs — plus de 30 standards applicables.

**Instruments en vigueur (sélection) :**

| Instrument | Depuis | Contenu |
|---|---|---|
| Mesures intérimaires sur l'IA générative | Août 2023 | Évaluation de sécurité du CAC **obligatoire avant lancement** ; 796 services enregistrés en fév. 2026 ; s'applique aussi aux services étrangers visant des utilisateurs en Chine (art. 20) |
| Dispositions sur la synthèse profonde (deep synthesis) | 2023 | Étiquetage des deepfakes, enregistrement des algorithmes |
| **Mesures d'étiquetage des contenus synthétiques** | **01/09/2025** | **Double marquage obligatoire** : label visible + métadonnées implicites (filigrane) sur tout contenu IA (texte, image, audio, vidéo) ; norme nationale **GB 45438-2025** (dimensions des filigranes, format des métadonnées) |
| Lignes directrices éthique-sécurité TC260 1.0 | Mai 2026 | Cycle de vie complet de l'IA |
| **Mesures sur les services d'interaction anthropomorphique** (5 ministères) | **Effectif juil. 2026** | Règles pour agents/chatbots « humains-like » : interdiction des compagnons virtuels intimes pour mineurs, mode mineur, anti-addiction (rappels après 2 h continues) |
| Plan 2026 : 48 nouveaux standards | 2026 | Sécurité IA, sécurité des données, cloud ; norme GB/Z 185-2026 sur l'interopérabilité multi-agents en entreprise |

**Enforcement réel** : le **12 février 2026**, le CAC a sanctionné **13 421 comptes** et supprimé
**543 000 contenus** pour défaut d'étiquetage IA ; amendes jusqu'à 1 M RMB. Une loi nationale
globale est en « recherche législative » — le ministre de la Justice He Rong évoque en mars 2026
une adoption au plus tôt en **2027** (sources : inotives.github.io, casrai.org, transcend.io).

**Stratégie industrielle** : open-source d'État (DeepSeek V4 sur puces **Huawei Ascend** — réduction
de la dépendance à Nvidia), financements publics massifs (~912 Md$ de « government guidance funds »
2000-2023, « à vérifier » sur la méthodologie). La Chine régule *et* accélère : contrôle du contenu,
course aux capacités.

#### 2.4.4. Le reste du monde (bref)

- **Royaume-Uni** : approche pro-innovation, pas de loi horizontale ; l'**AI Security Institute**
  (ex-AI Safety Institute) publie des évaluations (dont l'étude d'août 2026 sur les agents).
- **Corée du Sud, Japon, Singapour, Canada, Brésil** : lois ou cadres en construction — « à vérifier »
  au cas par cas, le paysage 2026 bouge par trimestres.
- **International** : sommets (Bletchley 2023, Séoul 2024, Paris 2025, **India AI Impact Summit,
  fév. 2026**), Rapport international sur la sécurité de l'IA 2026 (Bengio, 29 nations).

**Tableau comparatif des philosophies :**

| | UE | États-Unis | Chine |
|---|---|---|---|
| Philosophie | Droits fondamentaux, approche par les risques | Innovation d'abord, patchwork sectoriel | Contrôle de l'information + souveraineté techno |
| Instrument | Une loi horizontale (AI Act) | EO + lois d'États + droit existant (FTC…) | Mesures administratives empilées |
| Point dur 2026 | Transparence applicable ; haut risque reporté | Bras de fer fédéral vs États | Étiquetage obligatoire + enregistrement des algos |
| Tendance | Simplification (Omnibus) sans renoncement | Vers un standard national ? (litiges en cours) | Vers une loi-cadre en 2027 ? |

---

## 3. Focus Claude / Anthropic

### 3.1. Histoire factuelle

