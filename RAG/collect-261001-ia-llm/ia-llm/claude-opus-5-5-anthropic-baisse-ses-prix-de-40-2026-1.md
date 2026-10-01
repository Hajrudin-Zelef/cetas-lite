---
id: collect-261001-ia-llm/ia-llm/claude-opus-5-5-anthropic-baisse-ses-prix-de-40-2026-1
title: "claude-opus-5-5-anthropic-baisse-ses-prix-de-40-2026"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "Microsoft", "OpenAI"]
dates: []
keywords: ["claude", "agents", "arr", "astra", "benchmark", "benchmarks", "fable 5", "foundry", "gpt-6", "mythos 5", "opus 5", "sonnet 5"]
source: docs/RAG/collect-261001-ia-llm/claude-opus-5-5-anthropic-baisse-ses-prix-de-40-2026.md
source_anchor: ""
source_lines: [1, 44]
sha256: 706023a435a9ded8b80885c69b933ba2f4d5bd2712f4ab4aca9849f35e5927f3
---

# claude-opus-5-5-anthropic-baisse-ses-prix-de-40-2026

Anthropic a mis en ligne Claude Opus 5.5 le 22 septembre 2026, deux jours avant la publication de cet article. Le nouveau modèle coûte 40 % de moins à exploiter que son prédécesseur Opus 5 pour un niveau de performance jugé comparable à celui de Claude Fable 5.1, le modèle phare de la maison lancé trois semaines plus tôt. La sortie intervient dans un contexte particulier : dix jours après un texte du patron d’Anthropic, Dario Amodei, appelant l’industrie à ralentir le rythme des lancements de modèles. L’entreprise, elle, continue d’accélérer, tout en préparant une introduction en bourse qui pourrait intervenir avant la fin de l’année.

## Qu’est-ce que Claude Opus 5.5, exactement ?

Claude Opus 5.5 porte l’identifiant technique `claude-opus-5-5` et inaugure ce qu’Anthropic appelle sa famille de modèles 5.5, dont devraient suivre des déclinaisons Sonnet 5.5 et Haiku 5.5 dans les prochains mois. Contrairement à d’autres lancements récents passés par une phase de préversion, Opus 5.5 est arrivé directement en disponibilité générale, accessible dès le jour de l’annonce sur l’API Claude ainsi que sur Amazon Web Services, Google Cloud et Microsoft Azure via Microsoft Foundry.

Le modèle cible en priorité les usages agentiques de longue durée : agents de codage capables de travailler plusieurs heures sur une tâche, automatisation d’ordinateur (computer use) et travail de connaissance complexe. Il embarque une fenêtre de contexte d’un million de tokens par défaut et peut produire jusqu’à 128 000 tokens de sortie en une seule réponse, avec un mode de raisonnement adaptatif activé en permanence plutôt que sur demande.

Opus 5.5 est proposé aux abonnés Claude Pro, Max, Team et Enterprise. Anthropic a également relevé les limites d’usage sur cinq heures pour ces formules au moment du lancement, et offert une réinitialisation promotionnelle des quotas utilisable jusqu’au 22 octobre 2026.

## Une baisse de prix qui tranche avec la tendance du secteur

Le fait le plus commenté de ce lancement reste son tarif. Sur la page officielle du produit, Anthropic précise que “les tokens d’entrée et de sortie coûtent 4 et 20 dollars par million, soit 20 % de moins que Opus 5” (Anthropic, page de lancement Claude Opus 5.5). Opus 5 facturait jusqu’ici 5 dollars par million de tokens en entrée et 25 dollars en sortie. La baisse ne s’arrête pas au tarif catalogue : Anthropic ajoute que “les lectures de cache, qui représentent l’essentiel des coûts du travail agentique et du codage, coûtent 0,20 dollar par million de tokens, soit 60 % de moins que Opus 5” (Anthropic, page de lancement Claude Opus 5.5).

Sur la fiche produit générale du modèle Opus, Anthropic résume l’écart en une phrase : “Opus 5.5 coûte 4 dollars par million de tokens d’entrée et 20 dollars par million de tokens de sortie, 20 % de moins qu’Opus 5” (Anthropic, page produit Opus). L’entreprise va plus loin en affirmant qu’un projet agentique typique coûte au total 40 % de moins à exécuter avec Opus 5.5, le modèle consommant aussi moins de tokens pour arriver au même résultat, et non pas seulement facturé moins cher à l’unité.

Ce mouvement va à contre-courant de ce que le marché de l’IA générative a connu ces derniers mois en Europe, où plusieurs fournisseurs ont plutôt relevé leurs tarifs face à la saturation de la demande en calcul. Anthropic choisit ici la voie inverse sur son modèle le plus performant, un pari qui vise autant les développeurs indépendants que les grands comptes qui font tourner des agents en continu.

## Les benchmarks : Opus 5.5 dépasse Fable 5.1 sur le terrain agentique

Anthropic publie ses propres scores d’évaluation interne, repris par plusieurs médias spécialisés dès le 22 septembre. Sur Terminal-Bench 4.0, un test qui mesure la capacité d’un modèle à exécuter des tâches longues dans un terminal, Opus 5.5 atteint 66,4 %, contre 55,8 % pour Fable 5.1 et 52,3 % pour Opus 5, selon les chiffres relayés par MarkTechPost (MarkTechPost, 22 septembre 2026). Sur ce même test, GPT-6 Astra d’OpenAI se situe à 57,9 %, derrière Opus 5.5 mais devant Fable 5.1.

Le tableau se complique sur Terminal-Bench-Science 0.1, orienté recherche scientifique agentique : Opus 5.5 y obtient 58,7 %, en retrait par rapport aux 64,6 % de GPT-6 Astra, mais nettement devant les 29,0 % d’Opus 5. Anthropic ne domine donc pas sur tous les fronts, et l’entreprise elle-même ne revendique une parité qu'”au niveau de Claude Fable 5.1 sur la plupart des tâches”, une formulation qui laisse de côté les cas où l’écart se creuse dans un sens ou dans l’autre.

| Benchmark | Claude Opus 5.5 | Claude Fable 5.1 | Claude Opus 5 | GPT-6 Astra | 
|---|---|---|---|---|
| Terminal-Bench 4.0 | 66,4 % | 55,8 % | 52,3 % | 57,9 % | 
| Terminal-Bench-Science 0.1 | 58,7 % | 52,6 % | 29,0 % | 64,6 % | 
| FrontierCode v1.1 Main | 54,4 % | 50,3 % | 48,0 % | non communiqué | 
| CursorBench 4.0 | 57,8 % | 51,8 % | 46,6 % | non communiqué | 
| GDPval-AA v2.1 (score Elo) | 1 846 | 1 735 | 1 708 | non communiqué | 

Sur la vitesse, Anthropic indique qu’Opus 5.5 génère ses réponses plus de 30 % plus vite qu’Opus 5, un gain qui compte particulièrement pour les agents de codage qui enchaînent des dizaines d’appels successifs. Côté sécurité, l’entreprise avance qu’Opus 5.5 est 85 % moins susceptible de tenter de contourner les limites qui lui sont fixées, comparé à Opus 5 et à Mythos 5.1, sans détailler la méthodologie exacte de ce test interne.

## Le paradoxe Amodei : ralentir en paroles, accélérer dans les faits

Le calendrier de ce lancement interroge. Le 12 septembre 2026, Dario Amodei a publié un texte appelant l’ensemble de l’industrie à ralentir le rythme d’amélioration des capacités des modèles d’IA, en proposant un plan en trois étapes conçu, selon ses propres mots, sans “sacrifier l’avantage commercial ni le leadership américain en matière d’IA” (CNBC, 14 septembre 2026). Dix jours plus tard, Anthropic livrait justement un nouveau modèle plus performant et moins cher que le précédent.

Les analyses publiées entre-temps notent que l’appel d’Amodei visait le secteur dans son ensemble plutôt qu’une annonce de changement du calendrier de lancement propre à Anthropic. Reste que la contradiction visuelle est difficile à ignorer pour les observateurs du secteur : une entreprise qui demande publiquement de freiner la course aux capacités tout en sortant, dix jours après, un modèle qui bat son propre modèle phare sur plusieurs benchmarks agentiques. Ce grand écart illustre une tension structurelle propre à l’industrie de l’IA générative, où les discours sur la prudence coexistent avec une pression concurrentielle qui ne laisse guère de place à la pause.

## Anthropic vise 100 milliards de dollars de revenu annualisé avant son introduction en bourse

Ce lancement s’inscrit dans une séquence financière chargée pour Anthropic. L’entreprise serait en passe de dépasser 100 milliards de dollars de revenu annualisé d’ici la fin de l’année 2026, un chiffre rapporté à la fois par le New York Times et par plusieurs agrégateurs financiers spécialisés à la mi-septembre, dont Gulf News (Gulf News, 19 septembre 2026). Une partie de cette accélération tiendrait aux revenus du deuxième trimestre 2026, rapportés à plus de 11,5 milliards de dollars sur la période, contre 787 millions de dollars un an plus tôt selon des chiffres préliminaires cités par la presse spécialisée en investissement.

