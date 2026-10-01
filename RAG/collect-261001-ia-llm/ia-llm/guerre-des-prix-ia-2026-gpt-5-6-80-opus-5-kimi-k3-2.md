---
id: collect-261001-ia-llm/ia-llm/guerre-des-prix-ia-2026-gpt-5-6-80-opus-5-kimi-k3-2
title: "Estimation simplifiee du cout mensuel (500M tokens entree, 100M tokens sortie)"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Moonshot", "OpenAI"]
dates: []
keywords: ["agents", "attention", "chatgpt", "claude", "deepseek", "fable 5", "gemini", "gpt-5.6", "kimi", "luna", "mai", "open source"]
source: docs/RAG/collect-261001-ia-llm/guerre-des-prix-ia-2026-gpt-5-6-80-opus-5-kimi-k3.md
source_anchor: ""
source_lines: [35, 78]
sha256: 14fe2b7d52055095477043edb9415b00ba895ef56d843a016e7c189e3b51b332
---

# Estimation simplifiee du cout mensuel (500M tokens entree, 100M tokens sortie)

Google DeepMind a ouvert le bal des annonces de l’été en dévoilant, le 21 juillet 2026, Gemini 3.6 Flash et Gemini 3.5 Flash-Lite. Gemini 3.6 Flash est facturé 0,75 $ par million de tokens en entrée et 3,75 $ en sortie, avec un tarif réduit à 0,075 $ pour les tokens mis en cache, selon la documentation officielle de tarification Gemini. Sa fenêtre de contexte atteint 1 048 576 tokens en entrée et 65 536 tokens en sortie, avec des paliers batch qui divisent le tarif par deux, à 0,75 $/3,75 $.

Gemini 3.5 Flash-Lite, lancé le même jour, vise directement le segment ultra-économique où évoluent DeepSeek et les petits modèles open source : 0,30 $ par million de tokens en entrée et 2,50 $ en sortie, pour une fenêtre de contexte d’environ 1 million de tokens. Avec cette double sortie, Google couvre simultanément le segment agentique haut de gamme (Flash) et le segment volume (Flash-Lite), une stratégie de tenaille qui laisse peu d’espace tarifaire aux nouveaux entrants. Voir aussi notre article sur Gemini 3.5 Pro et l’accès restreint de GPT-5.6.

## Tableau comparatif : les prix des principaux LLM en août 2026

Le tableau ci-dessous résume les tarifs officiels en vigueur au 18 août 2026, exprimés en dollars par million de tokens, pour les modèles cités dans cette analyse.

| Modèle | Éditeur | Date de sortie / MAJ | Entrée ($/M tokens) | Sortie ($/M tokens) | Contexte | 
|---|---|---|---|---|---|
| GPT-5.6 Luna | OpenAI | 30 juillet 2026 (baisse) | 0,20 $ | 1,20 $ | ~200K | 
| GPT-5.6 Terra | OpenAI | 30 juillet 2026 (baisse) | 2,00 $ | 12,00 $ | ~200K | 
| GPT-5.6 Sol | OpenAI | 9 juillet 2026 | 5,00 $ | 30,00 $ | ~200K | 
| Claude Opus 5 | Anthropic | 24 juillet 2026 | 5,00 $ | 25,00 $ | 1M | 
| Claude Fable 5 | Anthropic | mi-juin 2026 | 10,00 $ | 50,00 $ | 1M+ | 
| Gemini 3.6 Flash | Google DeepMind | 21 juillet 2026 | 1,50 $ | 7,50 $ | 1 048 576 | 
| Gemini 3.5 Flash-Lite | Google DeepMind | 21 juillet 2026 | 0,30 $ | 2,50 $ | ~1M | 
| Kimi K3 (poids ouverts) | Moonshot AI | 27 juillet 2026 | 3,00 $ | 15,00 $ | 1M | 
| DeepSeek V4-Flash 0731 | DeepSeek | 31 juillet 2026 | 0,14 $ | 0,28 $ | 128K+ | 

Ce tableau illustre un écart de prix qui atteint désormais un facteur 71 entre le modèle le plus économique (DeepSeek V4-Flash 0731, à 0,14 $ en entrée) et le modèle le plus premium (Claude Fable 5, à 10 $ en entrée). Un tel écart n’existait pas il y a encore dix-huit mois, quand les modèles frontières se négociaient tous dans une fourchette resserrée entre 3 $ et 15 $ par million de tokens en entrée.

## Pourquoi les prix s’effondrent : la mécanique économique

Trois forces convergent pour expliquer cette guerre des prix IA. D’abord, les gains d’efficacité matérielle et logicielle : chaque nouvelle génération de puces, combinée à des techniques d’inférence plus économes comme la quantification et la mise en cache agressive des préfixes, réduit mécaniquement le coût de calcul par token. Ensuite, la concurrence directe des modèles à poids ouverts chinois, DeepSeek et Kimi en tête, qui livrent des performances proches de la frontière à une fraction du prix, forçant les laboratoires américains à réagir sous peine de perdre les usages à fort volume. Enfin, la sensibilité croissante des entreprises clientes au retour sur investissement : de nombreuses équipes techniques françaises et européennes ont désormais des tableaux de bord précis du coût par requête, et arbitrent activement entre fournisseurs selon le rapport qualité-prix, ce qui accroît la pression concurrentielle sur les marges.

Plusieurs analyses de coûts publiées cet été estiment que les prix des modèles frontières ont chuté d’environ 45 à 80 % entre 2025 et 2026, selon la catégorie de modèle considérée. Sur un horizon plus long, remontant à la sortie de GPT-4 mi-2023, la baisse cumulée dépasserait 85 à 95 % pour un niveau de capacité comparable, un effondrement de prix qui rappelle la courbe d’apprentissage observée sur le stockage cloud ou la bande passante réseau dans les décennies précédentes.

## L’avis de l’Autorité de la concurrence française sur la concentration du marché

Cette guerre des prix ne se joue pas uniquement sur le terrain économique : elle attire aussi l’attention des régulateurs. En juillet 2026, l’Autorité de la concurrence française a adopté un avis établissant, sur la base de données d’usage Sensor Tower datées de mai 2026, qu’OpenAI, Google et Anthropic contrôlent ensemble plus de 84 % du marché mondial des agents IA. L’avis ne détaille pas la répartition individuelle entre ces trois acteurs, mais souligne un risque de verrouillage : des agents IA peu coûteux et fortement intégrés à un écosystème propriétaire pourraient, selon l’Autorité, concentrer davantage l’économie numérique autour d’un nombre restreint d’entreprises si aucune mesure d’interopérabilité n’est prise.

Le même avis rappelle que ChatGPT reste l’assistant IA le plus utilisé en France, avec 21,6 millions de visiteurs uniques enregistrés en septembre 2025 selon Médiamétrie. Autrement dit, la baisse des prix profite en premier lieu aux acteurs déjà dominants, qui peuvent absorber une marge plus faible sur un volume plus important, un mécanisme classique de consolidation par les coûts que les autorités de concurrence surveillent de près dans d’autres secteurs numériques.

## Impact sur les entreprises françaises et européennes

Pour les équipes techniques basées en France, cette guerre des prix change concrètement les arbitrages budgétaires. Les startups qui traitaient auparavant des volumes limités avec des modèles premium peuvent désormais migrer certaines charges de travail vers des paliers d’entrée de gamme comme GPT-5.6 Luna ou Gemini 3.5 Flash-Lite, tout en réservant les modèles haut de gamme comme Claude Opus 5 ou GPT-5.6 Sol aux tâches qui exigent un raisonnement plus poussé. Cette segmentation par cas d’usage, déjà courante chez les éditeurs de logiciels, devient un réflexe pour les équipes qui gèrent des budgets d’infrastructure IA en forte croissance.

L’ouverture des poids de Kimi K3 et la disponibilité de DeepSeek V4-Flash offrent par ailleurs une option d’hébergement souverain aux entreprises soumises à des contraintes réglementaires strictes, notamment dans les secteurs de la santé, de la finance ou du secteur public, où le RGPD et les exigences de localisation des données limitent le recours direct aux API américaines. Plusieurs fournisseurs cloud européens ont d’ailleurs commencé à proposer l’hébergement de modèles à poids ouverts sur leur propre infrastructure, une tendance qui devrait s’accélérer à mesure que ces modèles gagnent en maturité.

## Comparaison historique : de 30 $ à moins de 3 $ par million de tokens

Pour mesurer l’ampleur de cette chute, il faut remonter à mi-2023, quand GPT-4 facturait environ 30 $ par million de tokens en entrée pour un niveau de capacité qui semblait alors à la pointe de l’état de l’art. Trois ans plus tard, des modèles frontières comme Gemini 3.6 Flash ou Kimi K3 livrent des performances supérieures pour 1,50 $ à 3 $ par million de tokens en entrée, et les modèles les plus économes comme DeepSeek V4-Flash descendent sous les 0,15 $. Cette trajectoire rappelle la baisse continue du coût du calcul observée depuis les débuts de l’informatique, mais à un rythme nettement plus rapide, porté par une concurrence directe entre laboratoires américains, chinois et, dans une moindre mesure, européens.

