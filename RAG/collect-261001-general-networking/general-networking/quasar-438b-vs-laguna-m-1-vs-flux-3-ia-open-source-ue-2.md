---
id: collect-261001-general-networking/general-networking/quasar-438b-vs-laguna-m-1-vs-flux-3-ia-open-source-ue-2
title: "quasar-438b-vs-laguna-m-1-vs-flux-3-ia-open-source-ue"
domain: general-networking
role: reference
task: reference
actors: ["Hugging Face", "Mistral", "Poolside", "Stability AI", "Z.ai"]
dates: []
keywords: ["agents", "apache", "benchmarks", "diffusion", "glm", "mistral", "multimodal", "open source", "open-weight", "valuation"]
source: docs/RAG/collect-261001-general-networking/quasar-438b-vs-laguna-m-1-vs-flux-3-ia-open-source-ue.md
source_anchor: ""
source_lines: [29, 64]
sha256: 7113a9536c09073631e1d1b39958fb1d7535c1a67a4f9a302f0cee203584c4c2
---

# quasar-438b-vs-laguna-m-1-vs-flux-3-ia-open-source-ue

Contrairement aux deux précédents, Black Forest Labs ne construit pas de grand modèle de langage. Fondée en 2024 par Robin Rombach, Andreas Blattmann, Patrick Esser et Dominik Lorenz, quatre chercheurs qui ont contribué à la création de Stable Diffusion chez Stability AI, l’entreprise allemande a choisi de se concentrer sur la génération d’images, de vidéo et, plus récemment, d’actions robotiques. Elle a levé plus de 450 millions de dollars au total, dont 300 millions de dollars lors d’une série B bouclée en décembre 2025, qui la valorise à 3,25 milliards de dollars.

Le 24 juillet 2026, Black Forest Labs a dévoilé FLUX 3, sa troisième génération de modèles génératifs. La rupture technique tient dans l’architecture : au lieu d’enchaîner des modules séparés pour l’image, le son et la vidéo, FLUX 3 traite ces modalités, plus les actions robotiques, dans un seul réseau neuronal unifié. Le produit se décline en quatre variantes. FLUX 3 Video génère des séquences allant jusqu’à 20 secondes avec un audio synchronisé nativement, gère les transitions par images-clés, le doublage multilingue avec synchronisation labiale, et le rendu de texte à l’intérieur même de la vidéo. FLUX 3 Image suit une ouverture progressive. FLUX 3 Action cible la robotique, avec un accès restreint à des partenaires de recherche. FLUX 3 Dev, enfin, est annoncé comme une future version open-weight, sans date précise, pensée pour un déploiement sur site.

Sur la performance, Black Forest Labs revendique des taux de préférence élevés dans des évaluations à l’aveugle menées en interne : FLUX 3 serait préféré dans 77 % des comparaisons face à Runway Gen-4.5, et dans 93 % des cas face à Luma Ray 3.2. Ces chiffres, publiés par l’entreprise elle-même, n’ont pas été vérifiés par un évaluateur indépendant comme Artificial Analysis, contrairement aux scores de Quasar 438B ou de Laguna S 2.1. Le premier cas d’usage concret documenté concerne le constructeur automobile Audi, qui teste “FLUX-mimic” pour convertir des démonstrations vidéo réalisées par des humains en séquences d’actions exécutables par des robots, sans programmation manuelle.

Aucun tarif public n’a été communiqué pour FLUX 3 Video ou FLUX 3 Action à la date de lancement, les deux variantes restant réservées à des partenaires sous conditions non divulguées. Sur les trois projets comparés ici, Black Forest Labs est le seul dont le siège social, la recherche et la propriété intellectuelle restent entièrement localisés en Europe, sans participation d’un investisseur américain majeur ni dépendance à un modèle chinois sous-jacent.

## Le contexte réglementaire : pourquoi Bruxelles pousse l’IA open source

Ces trois lancements ne surgissent pas dans le vide. La Commission européenne a multiplié en 2026 les textes destinés à réduire la dépendance du continent aux fournisseurs américains et chinois d’intelligence artificielle. Le 3 juin 2026, elle a publié sa Communication sur la souveraineté technologique européenne, accompagnée d’une stratégie IA open source et d’un règlement instaurant un cadre de souveraineté pour le cloud et l’IA baptisé CADA, structuré autour de quatre niveaux d’assurance de souveraineté. Le 7 juillet 2026, un second texte, le plan d’action sur la cybersécurité et l’intelligence artificielle, a ajouté un volet consacré à l’expansion des “capacités souveraines en IA, incluant les modèles frontières, la puissance de calcul et l’expertise en cybersécurité”.

Concrètement, ce cadre se traduit par des financements directs. Le programme GenAI4EU alloue 50 millions d’euros pour faire progresser les modèles d’IA open source, tandis qu’un appel d’offres lancé le 30 juillet 2026 vise à financer jusqu’à sept “gigafactories” IA en Europe pour entraîner la prochaine génération de modèles. La Commission cite explicitement Mistral AI comme exemple de fournisseur de modèles open-weight jouant le rôle d'”alternative souveraine aux systèmes propriétaires”, et soutient le consortium openEuroLLM comme initiative phare pour bâtir des modèles fondateurs “véritablement ouverts”. Depuis le 2 août 2026, la Commission exerce en outre pleinement ses pouvoirs de supervision de l’AI Act sur les modèles d’IA à usage général, y compris ceux présentant des risques systémiques de cybersécurité.

C’est dans ce climat politique que Quasar 438B, les modèles Laguna et FLUX 3 se présentent chacun comme une pièce du puzzle de la souveraineté numérique européenne. Mais comme le montre la suite de ce comparatif, l’IA open source revendiquée par ces trois projets ne se vaut pas : un modèle peut être développé par une entreprise européenne tout en s’appuyant sur une architecture ou un capital extra-européens, ce qui complique l’évaluation que doivent faire les acheteurs publics et privés soumis aux nouvelles règles de Bruxelles.

## Tableau comparatif : caractéristiques techniques

Ce tableau réunit les caractéristiques vérifiables des trois projets, ainsi que deux points de référence externes (Mistral Large 3 et GLM-5.2, le modèle source de Quasar) pour resituer les échelles.

| Modèle | Entreprise / pays | Paramètres | Type | Licence | Contexte | Langues | Accès | Date de sortie | 
|---|---|---|---|---|---|---|---|---|
| Quasar 438B | Multiverse Computing, Espagne | 438 Md | Raisonnement / agents | Propriétaire (API seule) | 1M tokens | Anglais, espagnol | API CompactifAI | 2 septembre 2026 | 
| Laguna M.1 | Poolside, France/États-Unis | 225 Md | Planification agentique | Apache 2.0 | Non précisé publiquement | Multilingue (non détaillé) | API + hébergeurs tiers | Juillet 2026 | 
| Laguna S 2.1 | Poolside, France/États-Unis | 118 Md | Code | OpenMDW-1.1 (open-weight) | Non précisé publiquement | Multilingue (non détaillé) | Hugging Face, auto-hébergement | Juillet 2026 | 
| FLUX 3 Video/Image | Black Forest Labs, Allemagne | Non communiqué | Génératif multimodal | Propriétaire (Dev annoncé open-weight) | Vidéos jusqu’à 20 s | Doublage multilingue avec lip-sync | Accès partenaires / rollout progressif | 24 juillet 2026 | 
| FLUX 3 Action | Black Forest Labs, Allemagne | Non communiqué | Actions robotiques | Propriétaire | N/A | N/A | Partenaires recherche uniquement | 24 juillet 2026 | 
| GLM-5.2 (source de Quasar) | Zhipu AI / Z.ai, Chine | 753 Md (~40 Md actifs) | LLM généraliste / code | MIT (open-weight) | 1M tokens | Multilingue | Hugging Face, ModelScope, API | 13 juin 2026 | 
| Mistral Large 3 (référence) | Mistral AI, France | Non communiqué officiellement | LLM généraliste | Open-weight (termes Mistral) | 256K-262K tokens | Multilingue | API, Le Chat, auto-hébergement | 2 décembre 2025 | 

Ce tableau met en évidence un premier constat : sur les six lignes de modèles européens ou revendiqués comme tels, seul GLM-5.2 (qui n’est pas européen) et Laguna S 2.1 offrent un accès complet aux poids sous une licence permissive. Quasar 438B et les deux variantes principales de FLUX 3 restent verrouillées derrière une API propriétaire, ce qui limite fortement les possibilités d’auto-hébergement pour les entreprises soumises à des contraintes strictes de résidence des données.

## Benchmarks : que disent vraiment les chiffres

Comparer des benchmarks entre trois produits aussi différents impose une prudence méthodologique. Un score de raisonnement général n’a pas de sens face à un taux de préférence sur la génération vidéo. Ce tableau isole donc chaque résultat avec sa source et son contexte de mesure exacts, plutôt que de forcer une comparaison directe entre des indices incompatibles.

