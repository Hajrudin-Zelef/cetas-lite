---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-48
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "EU", "Google", "Meta"]
dates: ["2025-02-02", "2025-08-02", "2026-01-26", "2026-08-02", "2026-09-27", "2026-12-02", "2027-08-02", "2027-12-02", "2028-08-02"]
keywords: ["agi", "arr", "reasoning", "valuation"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [3733, 3791]
sha256: b17d9c0baf305b4d630bb66cdaa8b79450b8b962d73a7fca6cdf5591da001093
---

# IA — Le grand dossier

| Voix | Position | Citation / fait marquant |
|---|---|---|
| **Geoffrey Hinton** (Nobel 2024, ex-Google, parti en 2023 pour alerter) | Risque existentiel pris au sérieux : ~10-20 % de chance d'extinction humaine liée à l'IA dans les prochaines décennies (Noël 2024) ; les systèmes pourraient développer des sous-objectifs (recherche de contrôle, résistance à l'arrêt, tromperie). Seul le **régulateur** peut forcer les labs à investir en sécurité, pas le profit. | « Just leaving it to the profit motive of large companies is not going to be sufficient to make sure they develop it safely. » |
| **Yoshua Bengio** (Turing 2018) | Président du **Rapport international sur la sécurité de l'IA 2026** (29 nations, ONU, OCDE, UE, 100+ experts — la plus grande collaboration scientifique sur le sujet). Craint que l'IA ne déjoue la supervision humaine ; **contre l'open-sourcing des modèles les plus puissants** ; pour un investissement massif en recherche sécurité et une gouvernance forte. | A co-signé (avec Hinton, Russell, Lessig, août 2024) le soutien au projet de loi californien SB 1047, qualifié de « bare minimum » de régulation. |
| **Stuart Russell** (Berkeley, auteur du manuel de référence) | Le problème de l'**alignement** (s'assurer que les objectifs de l'IA restent compatibles avec les valeurs humaines) est le problème central ; plaide pour des IA « provably beneficial ». | Co-signataire de la lettre ouverte du Future of Life Institute (mars 2023, pause de 6 mois — non suivie d'effet). |
| **Dario Amodei** (CEO Anthropic) | Alerte emploi (50 % des cols blancs débutants, chômage 10-20 %) + appel à la préparation publique ; publie un cadre de politique économique (juin 2026) avec paliers d'intervention jusqu'au revenu universel. | « The Adolescence of Technology », 26/01/2026. Contredit publiquement par son propre économiste en chef (McCrory, juil. 2026). |
| **Yann LeCun** (Turing 2018, quitte Meta en nov. 2025 pour fonder **AMI Labs**, 1,03 Md$ — plus grosse seed européenne) | **Le dissident** : les LLM actuels sont « plus bêtes qu'un chat », une impasse pour l'AGI sans modèles du monde (JEPA) ; l'intelligence n'implique pas la motivation — la peur de la domination est une projection anthropomorphique ; la sécurité doit être **conçue dans l'architecture** (« controllable by design »), pas rafistolée après coup. | « Large language models cannot reach human-level reasoning without an internal world model. » |
| **Eliezer Yudkowsky** (MIRI) | Position maximaliste : l'alignement est un problème quasi insoluble à temps ; a appelé à des mesures extrêmes (dont frappes sur data centers voyous — proposition condamnée même par d'autres alarmistes). | Figure clivante : utile comme borne du spectre, pas comme boussole. |
| **Emily M. Bender & Timnit Gebru** (chercheuses, critique du « stochastic parrots ») | Critique *socio-technique* plutôt qu'existentialiste : les dangers sont **présents et concrets** (biais, exploitation des travailleurs de l'annotation, concentration du pouvoir, dommages environnementaux), pas hypothétiques. Dénoncent le discours « extinction » comme un écran de fumée servant les labs. | « Stochastic Parrots » (2021) ; Gebru a fondé le DAIR Institute. |
| **Bruce Schneier** (cryptographe) | Les risques majeurs sont la **concentration du pouvoir** (surveillance, manipulation) et l'érosion de la confiance sociale ; technophile prudent, pas alarmiste existentiel. | Plaide pour des contre-pouvoirs démocratiques et techniques (transparence, interopérabilité). |
| **Fei-Fei Li** (Stanford HAI) | Centrée humain : l'IA doit amplifier l'humain ; critique d'une régulation qui étoufferait la recherche académique au profit des seuls géants. | Voix influente pour une « IA centrée humain » pragmatique. |

**Lecture du spectre** : de LeCun (« le danger est exagéré, la techno actuelle est limitée ») à
Yudkowsky (« catastrophe quasi certaine sans arrêt »), en passant par le centre Hinton/Bengio/Russell
(« risques réels, régulation nécessaire ») et la critique socio-technique Bender/Gebru (« les vrais
dégâts sont déjà là »). **Aucune de ces positions n'est un consensus scientifique** : le Rapport
international 2026 documente précisément les zones d'accord (capacités, mésusages) et de désaccord
(probabilités de catastrophe, timelines).

### 2.4. La régulation : état des lieux au 27/09/2026

> **Méthode.** Tout ce qui suit a été vérifié par recherche web le 27/09/2026. Les dates et
> chiffres portent leurs sources. Le droit bouge vite : revérifier avant toute décision de
> conformité.

#### 2.4.1. Union européenne — l'AI Act et l'AI Omnibus

**Le cadre.** Le règlement (UE) 2024/1689 (AI Act), entré en vigueur le **1er août 2024**, est la
première loi horizontale mondiale sur l'IA, à approche **par les risques** : pratiques interdites
(article 5), IA à haut risque (annexes I et III, obligations de conformité), obligations de
transparence (article 50), régime spécifique pour les modèles d'IA à usage général / GPAI
(articles 51-56).

**Ce qui a changé en 2026 : l'AI Omnibus.** Le « Digital Omnibus on AI » (règlement (UE) 2026/1744),
signé le **8 juillet 2026**, entré en vigueur le **27 juillet 2026**, a **décalé** les obligations
les plus lourdes — parce que les normes harmonisées et les organismes notifiés n'étaient pas prêts,
et que les États membres tardaient à désigner leurs autorités compétentes.

**Calendrier consolidé au 27/09/2026** (sources : Goodwin, juillet-août 2026 ; digitalapplied.com ; GitHub karunmehta-aigp) :

| Date | Obligation | Statut au 27/09/2026 |
|---|---|---|
| 02/02/2025 | Pratiques interdites (art. 5 : scoring social, manipulation, etc.) | **En vigueur** |
| 02/02/2025 | Obligation de « AI literacy » (formation des personnels) | **En vigueur** |
| 02/08/2025 | Obligations GPAI (transparence, documentation, évaluation pour les modèles à risque systémique) | **En vigueur** |
| 02/08/2026 | **Transparence (art. 50)** : tout chatbot/assistant doit déclarer qu'il est une IA ; tout contenu généré (texte, image, audio, vidéo) diffusé dans l'UE doit être marqué comme synthétique | **En vigueur depuis le 2 août 2026 — amendes possibles dès maintenant** |
| 02/08/2026 | Pouvoirs d'exécution de la Commission sur les GPAI | En vigueur |
| 02/12/2026 | Filigrane machine-readable (art. 50 §2) pour les systèmes déjà sur le marché avant le 02/08/2026 ; **nouvelle interdiction** : images intimes non consenties générées par IA + CSAM (nouvel art. 5) | **À venir** |
| 02/08/2027 | Bacs à sable réglementaires nationaux (art. 57) | À venir |
| **02/12/2027** | **Haut risque, systèmes autonomes (annexe III)** : gestion des risques, gouvernance des données, documentation technique, supervision humaine, évaluation de conformité | **Reporté de 16 mois** (était le 02/08/2026) |
| **02/08/2028** | **Haut risque embarqué dans des produits réglementés** (annexe I : machines, dispositifs médicaux, jouets…) | **Reporté de 12 mois** (était le 02/08/2027) |

**Ce que ça change pour toi, concrètement :**

- Si tu déploies un chatbot (même interne) accessible dans l'UE : **il doit se présenter comme une IA** — c'est du droit applicable aujourd'hui, pas un projet.
- Si ton RAG génère des documents diffusés : marquage « généré par IA » requis.
- Les obligations lourdes (évaluation de conformité des systèmes à haut risque : recrutement, notation de crédit, etc.) sont repoussées fin 2027 — du temps pour se préparer, pas pour oublier.
- Sanctions : jusqu'à **35 M€ ou 7 % du chiffre d'affaires mondial** pour les pratiques interdites ; 15 M€ / 3 % pour les autres manquements (montants du texte d'origine — « à vérifier » sur d'éventuels ajustements Omnibus).

