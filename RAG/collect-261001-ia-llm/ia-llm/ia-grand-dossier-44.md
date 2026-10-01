---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-44
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "United States"]
dates: ["2026-09-27"]
keywords: ["agents", "diffusion", "mai"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [3511, 3559]
sha256: e01430d2c6ea0094d2df13104c2535ab68f2b6772d93e70c0141d83a8f0935ab
---

# IA — Le grand dossier

| Étude | Méthode | Résultat principal | Nuance |
|---|---|---|---|
| **Brynjolfsson, Li & Raymond, Stanford/MIT, 2023** | 5 179 agents de support client, essai contrôlé randomisé | +14 % de résolutions par heure en moyenne ; **+34 % pour les agents novices**, quasi aucun gain pour les experts | Effet de diffusion des connaissances des meilleurs vers les débutants |
| **Noy & Zhang, 2023** | 453 professionnels, tâches de rédaction | −40 % de temps, +18 % de qualité évaluée | Gains concentrés sur la rédaction |
| **Dell'Acqua et al., Harvard Business School / BCG, 2023** | 758 consultants BCG | +12,2 % de tâches terminées, +25,1 % de vitesse, +40 % de qualité — **mais −19 points sur les tâches hors « frontière »** (l'IA induit en erreur quand elle est utilisée hors de son domaine de compétence) | « Jagged frontier » : gains inégaux selon les tâches |
| **Acemoglu, MIT, « The Simple Macroeconomics of AI », 2024** | Modélisation à partir des études ci-dessus | +0,66 % de productivité totale des facteurs et +0,93 à +1,56 % de PIB **sur 10 ans** | Qualifie ces gains de « non triviaux mais modestes » ; prévient contre les extrapolations |
| **MIT Project NANDA, juillet 2025** | Enquête auprès de ~300 initiatives d'IA générative en entreprise | **~95 % des pilotes IA n'atteignent pas un impact P&L mesurable ; ~5 % réussissent** | Les auteurs précisent : chiffres directionnels, pas d'audit comptable. L'échec vient rarement de la techno (données, intégration, conduite du changement) |
| **RAND Corporation, 2024** | Revue de projets | Plus de 80 % des projets IA échouent, soit ~2× le taux d'échec des projets IT non-IA | — |
| **Cruces et al., NBER, fév. 2026** | Expérience randomisée, tâches business | L'IA réduit de **deux tiers** l'écart de performance lié au niveau d'éducation (+1,24 écart-type pour les moins diplômés vs +0,83 pour les plus diplômés) | Bémol : délégation possible au lieu d'apprentissage réel |
| **Aral & Ju, MIT Sloan, 2025** | 2 000+ participants, équipes humain-humain vs humain-IA | +60 % de productivité par travailleur avec agents IA (au prix de −23 % de messages sociaux) | Coût social de la collaboration |

#### Lecture croisée (attribuée)

- **Les optimistes** (ex. : Goldman Sachs, 2024, estimation d'un potentiel de +15 % de productivité du travail à long terme dans les économies avancées) voient dans l'IA un choc de productivité comparable à l'électrification.
- **Les prudents** (Daron Acemoglu, MIT ; Robert Gordon, Northwestern) rappellent que les gains mesurés en laboratoire ne se traduisent pas automatiquement en gains macroéconomiques : adoption lente, coûts d'intégration, tâches non automatisables.
- **Le point de consensus** : les gains existent, ils sont hétérogènes (forts pour les débutants et les tâches de rédaction/codage, faibles ou négatifs hors domaine de compétence de l'IA), et l'écart entre pilotes réussis et échecs massifs (95 % selon le MIT) tient à l'organisation, pas au modèle.

> **Pour un sysadmin.** La productivité de l'IA en production ne se mesure pas au modèle mais au
> pipeline : données propres, garde-fous, métriques. Un pilote IA sans baseline mesurée est un
> hobby, pas un investissement.

### 1.3. Emploi : les deux camps

#### Camp « substitution massive »

- **Dario Amodei, CEO d'Anthropic** — interview Axios, mai 2025, puis Davos janvier 2026 et essai de 20 000 mots « The Adolescence of Technology » (26 janvier 2026) : l'IA pourrait éliminer **~50 % des emplois de bureau débutants en 1 à 5 ans**, avec un chômage potentiel de **10 à 20 %**. Il plaide pour la préparation publique (dont revenu universel / taxation des entreprises d'IA, cadre économique publié par Anthropic en juin 2026). Nuance interne : **Peter McCrory, économiste en chef d'Anthropic**, a publié en juillet 2026 une analyse contredisant son propre CEO — chômage US à 4,2 % en juin 2026, aucune divergence d'emploi mesurable entre métiers très exposés à l'IA et les autres, et il ne s'attend pas à une hausse notable du chômage d'ici un an « à cause de l'IA ».
- **Geoffrey Hinton** (prix Nobel de physique 2024) met en avant le chômage de masse et les inégalités comme risques socio-économiques majeurs, tout en reconnaissant les bénéfices (médicaments, santé, éducation).
- **Le Forum économique mondial, Future of Jobs Report 2025** : 92 millions d'emplois déplacés mais 170 millions créés d'ici 2030 selon les employeurs interrogés (solde net positif, mais avec une recomposition massive des compétences).

#### Camp « transformation, pas disparition »

- **David Autor, MIT** (économiste du travail) : l'IA pourrait *restaurer* la valeur du travail intermédiaire en rendant les non-experts capables de tâches d'experts — l'inverse d'une polarisation totale. Il insiste sur le rôle des institutions (formation, négociation) dans le partage des gains.
- **Erik Brynjolfsson, Stanford** : l'IA augmente les travailleurs plutôt qu'elle ne les remplace dans la plupart des cas mesurés ; le risque est la *mauvaise* conception des déploiements.
- **Les sceptiques de la substitution** rappellent l'histoire : le tableur n'a pas supprimé les comptables, le CMS n'a pas supprimé les rédacteurs. La productivité détruit des *tâches*, pas des *métiers* — mais la transition peut être brutale pour les cohortes concernées.

#### Ce que montrent les données début 2026 (attribué)

- L'étude NBER de Cruces et al. (février 2026) montre que l'IA profite *davantage* aux moins qualifiés en performance de tâche — ce qui plaide pour un effet d'égalisation salariale *si* les salaires suivent la productivité (condition non garantie).
- Anthropic (juin 2026, Economic Policy Framework) note que les premiers signaux montrent déjà une croissance de l'emploi plus faible pour les débutants dans les métiers exposés à l'IA — « à vérifier » sur données longitudinales plus longues.
- **Position honnête du dossier** : personne ne sait. Les modèles économiques divergent parce que l'adoption est à ses débuts et que les effets de second ordre (nouveaux métiers, baisse des prix, demande induite) sont par nature imprévisibles.

### 1.4. Créativité et propriété intellectuelle

Le conflit est simple à énoncer, impossible à trancher proprement : les modèles sont entraînés sur
des œuvres protégées sans autorisation systématique ; les créateurs y voient un pillage, les
laboratoires une « utilisation transformative » (fair use).

#### L'état du contentieux au 27/09/2026 (faits vérifiés par recherche web)

