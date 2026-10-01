---
id: collect-261001-ia-llm/ia-llm/fin-de-la-guerre-des-prix-ia-tarifs-llm-en-hausse-2026-1
title: "fin-de-la-guerre-des-prix-ia-tarifs-llm-en-hausse-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Mistral", "Nvidia", "OpenAI", "Z.ai"]
dates: ["2026-09-09", "2026-09-10", "2026-11-21", "2026-12-31", "2027-01-01"]
keywords: ["arr", "claude", "fable 5", "gemini", "gemini 3.8", "glm", "gpt-5.6", "gpu", "mistral", "nvidia", "opus 5", "sol"]
source: docs/RAG/collect-261001-ia-llm/fin-de-la-guerre-des-prix-ia-tarifs-llm-en-hausse-2026.md
source_anchor: ""
source_lines: [1, 47]
sha256: 5dd4ec72bd8046a8391fc3357e8d1ca973b94c1a0e8d544cac8407a55d7fc1af
---

# fin-de-la-guerre-des-prix-ia-tarifs-llm-en-hausse-2026

La baisse quasi ininterrompue des prix de l’intelligence artificielle générative, qui a marqué 2024 et 2025, vient de s’arrêter net. Depuis début septembre 2026, plusieurs fournisseurs majeurs relèvent leurs tarifs par million de tokens, invoquant la pénurie de puces Nvidia, la flambée des coûts énergétiques des centres de données et l’explosion de la demande en inférence longue. Pour les entreprises françaises et européennes qui ont bâti leurs budgets 2026 sur l’hypothèse d’une baisse continue, le réveil est brutal : GLM-5.3-Flash a doublé son tarif du jour au lendemain le 9 septembre, Solar Pro 4 a suivi le 10 septembre, et Google a déjà programmé un doublement de Gemini 3.7 Flash au 1er janvier 2027. Ce retournement rebat les cartes du calcul de rentabilité de l’IA en entreprise, à un moment où les dépenses mondiales en IA doivent bondir de 47 % cette année selon Gartner.

## La fin de la guerre des prix de l’IA : ce qui vient de changer

Pendant deux ans, la règle tacite de l’industrie de l’IA générative tenait en une phrase : chaque nouvelle génération de modèle coûte moins cher que la précédente à capacité égale. Entre début 2024 et fin 2025, les prix des modèles dits “frontières” sont passés d’environ 20 dollars par million de tokens en entrée à seulement 2 à 3 dollars fin 2025, tandis que les modèles économiques (“flash” ou “mini”) sont tombés jusqu’à 0,03-0,10 dollar par million de tokens au plus bas. Cette dynamique a permis à des milliers de start-up et de PME d’intégrer des LLM dans leurs produits sans exploser leur budget cloud.

Ce cycle s’est interrompu en septembre 2026. Selon le comparatif hebdomadaire publié par Json House, daté du 14 septembre, les prix des API de LLM s’étalent désormais sur deux ordres de grandeur, de 0,10 dollar par million de tokens en entrée pour les modèles économiques comme Gemini 2.5 Flash-Lite jusqu’à 30 dollars pour les niveaux “pro” de la gamme GPT-5.5. Le site NavyaAI, dans sa mise à jour du 20 septembre, résume la situation : les tarifs des modèles frontières se regroupent désormais autour de 4 à 5 dollars par million de tokens en entrée et 20 à 25 dollars en sortie, un plancher qui n’existait pas six mois plus tôt.

Le site spécialisé BenchLM.ai, dans sa note du 18 septembre 2026 consacrée aux tendances tarifaires des LLM, constate qu’il n’y a plus de course au moins-disant sur les modèles les plus avancés : la compétition s’est déplacée vers les capacités et les fonctionnalités destinées aux entreprises, tandis que les niveaux d’entrée de gamme et les modèles ouverts continuent, eux, de baisser. Autrement dit, la guerre des prix ne disparaît pas complètement, elle se scinde en deux marchés parallèles qui obéissent à des logiques opposées.

## GLM-5.3-Flash et Solar Pro 4 : les premiers doublements de tarifs

Le signal le plus visible de ce retournement est venu de Zhipu AI, l’éditeur chinois derrière la gamme GLM. Le 9 septembre 2026, GLM-5.3-Flash est passé de 0,065 euro à 0,13 euro par million de tokens en entrée, et de 0,22 euro à 0,43 euro en sortie, soit un doublement pur et simple appliqué du jour au lendemain, sans période de transition annoncée à l’avance. Pour les développeurs qui avaient intégré ce modèle économique dans des pipelines à fort volume, la facture mensuelle a mécaniquement doublé sans le moindre changement de code.

Le lendemain, le 10 septembre, Upstage a appliqué le même traitement à Solar Pro 4, avec une hausse de tarif également de l’ordre du doublement selon le comparatif publié par digital-m.fr. Ces deux annonces coup sur coup, en l’espace de 48 heures, ont suffi à convaincre plusieurs analystes que le mouvement dépassait le cas isolé et traduisait une tendance de fond du secteur.

Du côté d’Anthropic, la situation est plus nuancée. L’entreprise avait programmé une hausse de Claude Sonnet 5, de 2 dollars à 3 dollars par million de tokens en entrée et de 10 à 15 dollars en sortie, à compter du 1er septembre 2026. Cette hausse a finalement été annulée : Claude Sonnet 5 reste facturé 2 dollars en entrée et 10 dollars en sortie. C’est l’un des rares exemples où un fournisseur a fait marche arrière sur une hausse déjà annoncée, ce qui illustre à quel point la pression concurrentielle reste forte sur le segment intermédiaire, même quand elle a disparu sur le haut de gamme.

## Gemini 3.8 Flash : le piège du tarif gelé

Google a choisi une stratégie différente avec Gemini 3.8 Flash, lancé le 2 septembre 2026. Le tarif affiché, 0,75 euro par million de tokens en entrée et 3,75 euros en sortie, reste identique à celui de la génération précédente, Gemini 3.7 Flash. Mais ce gel n’est que temporaire : Google a déjà publié le calendrier de la hausse à venir. À compter du 1er janvier 2027, le tarif d’entrée de Gemini 3.7 Flash doublera, passant à 1,50 euro, et le tarif de lecture du cache montera à 7,50 euros, contre 3,75 euros actuellement selon le guide développeur publié à la mi-septembre.

Pour les équipes techniques françaises qui budgétisent leurs coûts d’inférence sur douze mois, cette annonce change la donne : ce qui ressemble aujourd’hui à un tarif stable est en réalité une fenêtre promotionnelle avec date d’expiration. Digital-m.fr souligne dans son analyse du marché que huit remises tarifaires doivent expirer d’ici janvier 2027, ce qui va mécaniquement alourdir la facture cloud IA de nombreuses entreprises sans qu’elles n’aient rien changé à leur usage.

OpenAI applique une logique comparable avec GPT-5.6 Sol, dont le tarif promotionnel de 4 dollars en entrée et 20 dollars en sortie par million de tokens court jusqu’au 21 novembre 2026. Une fois cette date passée, le prix devrait se rapprocher de la norme désormais établie du segment frontière, autour de 5 dollars en entrée et 25 dollars en sortie, au niveau de Claude Opus 5.

## Tableau comparatif : tarifs des principaux modèles en septembre 2026

Le tableau ci-dessous rassemble les tarifs par million de tokens des modèles les plus utilisés en Europe à la date du 22 septembre 2026, avec leur statut (stable, en promotion temporaire, ou déjà relevé).

| Modèle | Entrée / 1M tokens | Sortie / 1M tokens | Statut tarifaire | 
|---|---|---|---|
| Claude Fable 5 | 10 $ | 50 $ | Niveau raisonnement premium, stable | 
| Claude Opus 5 | 5 $ | 25 $ | Ancre du segment frontière | 
| GPT-5.6 Sol | 4 $ (promo) | 20 $ (promo) | Promo jusqu’au 21/11/2026 | 
| Claude Sonnet 5 | 2 $ | 10 $ | Hausse prévue le 01/09 annulée | 
| Gemini 3.8 Flash | 0,75 € | 3,75 € | Gelé jusqu’au 31/12/2026 | 
| Gemini 3.7 Flash | 0,75 € → 1,50 € (01/01/2027) | 3,75 € → 7,50 € | Doublement programmé | 
| GLM-5.3-Flash | 0,065 € → 0,13 € | 0,22 € → 0,43 € | Doublé le 09/09/2026 | 
| Solar Pro 4 | doublement | doublement | Relevé le 10/09/2026 | 
| Mistral Large 3 | 0,50 $ | 1,50 $ | Stratégie volume, prix bas maintenu | 
| Claude Haiku 4.5 | 1 $ | 5 $ | +25 % vs Haiku 3.5 (0,80 $/4 $) | 

## Pourquoi les prix repartent à la hausse : trois explications

La première explication, la plus citée par les fournisseurs eux-mêmes, tient au coût matériel. Les tensions persistantes sur l’approvisionnement en puces Nvidia H200 et B100 ont fait grimper les tarifs de location de GPU dans le cloud tout au long de 2026. Plusieurs rapports sur les dépenses d’entreprise en IA, dont celui publié par Rezolve AI, pointent ces contraintes d’approvisionnement en accélérateurs comme un facteur direct de la remontée des coûts d’inférence facturés aux clients finaux.

