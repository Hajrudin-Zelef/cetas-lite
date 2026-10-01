---
id: collect-261001-ia-llm/ia-llm/ai-act-2026-chatgpt-claude-gemini-sous-nouvelles-regles-1
title: "ai-act-2026-chatgpt-claude-gemini-sous-nouvelles-regles"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "EU", "Google", "Mistral", "OpenAI"]
dates: []
keywords: ["chatgpt", "claude", "gemini", "attention", "benchmark", "benchmarks", "fable 5", "gpt-5.6", "luna", "mistral", "opus 4", "opus 5"]
source: docs/RAG/collect-261001-ia-llm/ai-act-2026-chatgpt-claude-gemini-sous-nouvelles-regles.md
source_anchor: ""
source_lines: [1, 37]
sha256: bfeb6f92a85da4ef275704968d40a45248231c56988a4087d1d549f81340693f
---

# ai-act-2026-chatgpt-claude-gemini-sous-nouvelles-regles

Le 2 août 2026, l’Union européenne a franchi une étape que peu d’éditeurs de chatbots avaient anticipée à cette échelle : l’entrée en vigueur des obligations de transparence de l’AI Act pour les systèmes d’intelligence artificielle à usage général. ChatGPT, Claude, Gemini et une dizaine d’autres assistants conversationnels doivent désormais signaler clairement à leurs utilisateurs qu’ils interagissent avec une machine, et étiqueter tout contenu généré artificiellement. Le calendrier tombe pile pour les fournisseurs : la bascule réglementaire coïncide avec la sortie coup sur coup de Claude Opus 5, GPT-5.6, Gemini 3.6 Flash et Mistral Large 3, quatre modèles qui redessinent le classement mondial des LLM en l’espace de six semaines. Pour les développeurs et les entreprises françaises et européennes qui bâtissent leurs produits sur ces API, la question n’est plus seulement « quel modèle choisir », mais « quel modèle reste conforme, disponible et prévisible sur le marché européen ».

## L’AI Act entre en application : ce qui change concrètement le 2 août 2026

Le règlement européen sur l’intelligence artificielle avance par paliers depuis février 2025, mais le 2 août 2026 marque l’activation d’un bloc d’obligations qui touche directement les modèles à usage général (les fameux GPAI, pour *general-purpose AI*) : Claude, GPT, Gemini, Mistral et tous les systèmes reposant sur ces briques. Concrètement, les fournisseurs doivent désormais informer explicitement les utilisateurs qu’ils dialoguent avec un système automatisé, apposer un marquage lisible par machine sur les contenus synthétiques (texte, image, audio, vidéo) lorsque cela est techniquement possible, et documenter les données d’entraînement utilisées, selon l’analyse détaillée publiée par CNET France.

Ce n’est pas la première vague de l’AI Act. Les interdictions sur les pratiques d’IA jugées « inacceptables » (notation sociale, manipulation subliminale, reconnaissance des émotions au travail) s’appliquaient déjà depuis février 2025. Ce qui change en août 2026, c’est que l’obligation de transparence sort du registre des bonnes intentions pour devenir une exigence opérationnelle assortie de sanctions. La Commission européenne a par ailleurs publié le 16 juillet 2026 son propre modèle ouvert, l’EU Institutional LLM v1, une démonstration de conformité par l’exemple plutôt qu’un simple texte réglementaire. Le calendrier complet, disponible sur le site officiel de la Commission européenne dédié au cadre réglementaire de l’IA, prévoit encore une échéance en 2027 pour les obligations les plus strictes sur les systèmes à haut risque.

Pour un développeur qui intègre l’API d’un de ces modèles dans une application destinée au marché français, l’impact est direct : il devient responsable, en tant que déployeur, d’afficher les mentions de transparence dans son interface, même si le modèle sous-jacent appartient à OpenAI, Anthropic ou Google. Ignorer cette obligation n’est plus une zone grise contractuelle, c’est une non-conformité documentée.

## Claude Opus 5 prend la tête du classement mondial des LLM

Anthropic a lancé Claude Opus 5 le 24 juillet 2026, et le modèle a immédiatement pris la première place du classement fondé sur l’Intelligence Index v4.1 d’Artificial Analysis, mesuré au 1er août 2026. Avec un score d’intelligence de 61, une fenêtre de contexte d’un million de tokens et un tarif de 5 dollars par million de tokens en entrée et 25 dollars en sortie, Claude Opus 5 devance de peu son cousin Claude Fable 5, crédité d’un score de 60 mais facturé deux fois plus cher (10 dollars en entrée, 50 dollars en sortie), selon le classement des meilleurs LLM d’août 2026 publié par Digitiz. Nous avions déjà documenté cette bascule dans notre article sur Claude Opus 5 en tête du classement LLM.

Ce qui frappe dans ce nouveau palmarès, c’est la densité de sorties concentrées sur juillet-août 2026. Le site Blog du Modérateur recense un top 20 des modèles d’IA les plus performants où l’on retrouve, dans les dix premières places, pas moins de cinq variantes différentes de la famille Claude (Fable 5, Opus 4.6 High, Opus 4.7 High, Opus 5 High, Opus 5 Max), signe d’une stratégie de segmentation tarifaire agressive d’Anthropic plutôt qu’une simple montée en version linéaire.

Face à cette offensive, OpenAI a répliqué avec trois variantes de GPT-5.6 (Sol, Terra, Luna), positionnées sur des segments de prix très différents : de 0,20 dollar par million de tokens en entrée pour Luna, la version économique, jusqu’à 5 dollars pour Sol, la version haut de gamme. Les scores s’échelonnent de 55 à 59 sur l’Intelligence Index v4.1, un cran en dessous des meilleurs Claude, mais avec un avantage tarifaire net sur l’entrée de gamme. Cette guerre des prix, déjà amorcée avant l’été, fait l’objet de notre analyse sur la guerre des prix entre GPT-5.6 et Opus 5.

## Tableau comparatif : les LLM de pointe après la vague de juillet-août 2026

| Modèle | Éditeur | Date de sortie | Score Intelligence Index v4.1 | Contexte | Prix entrée / sortie (par 1M tokens) | 
|---|---|---|---|---|---|
| Claude Opus 5 | Anthropic | 24 juillet 2026 | 61 | 1M tokens | 5 $ / 25 $ | 
| Claude Fable 5 | Anthropic | Juillet 2026 | 60 | 1M tokens | 10 $ / 50 $ | 
| GPT-5.6 Sol | OpenAI | Été 2026 | 59 | 1M tokens | 5 $ / — (haut de gamme) | 
| GPT-5.6 Terra | OpenAI | Été 2026 | 57 | 1M tokens | Intermédiaire | 
| GPT-5.6 Luna | OpenAI | Été 2026 | 55 | 1M tokens | 0,20 $ / — (économique) | 
| Gemini 3.6 Flash | Google DeepMind | 21 juillet 2026 | 50 | 1M tokens | 1,50 $ / 7,50 $ | 
| Mistral Large 3 | Mistral AI | Juillet 2026 | Non communiqué (benchmark interne comparable) | 1M tokens | Non communiqué | 
| Claude Sonnet 4.6 | Anthropic | 3 juillet 2026 | ~92,5 (échelle distincte) | 1M tokens | 3 $ / 15 $ | 

Attention à la lecture croisée de ce tableau : les scores « Intelligence Index v4.1 » d’Artificial Analysis et les scores de benchmarks internes publiés par certains cabinets de conseil ne partagent pas la même échelle. C’est pour cette raison que Mistral Large 3 et Claude Sonnet 4.6 apparaissent avec des chiffres non directement comparables aux autres lignes. Le message à retenir n’est pas le score brut, mais la tendance : quatre éditeurs différents ont livré un modèle de rupture en moins de six semaines, un rythme de sortie qu’on n’avait plus observé depuis le printemps 2025.

## Gemini 3.6 Flash : Google mise sur le rapport performance-prix

Google DeepMind a dévoilé trois nouveaux modèles le 21 juillet 2026, avec Gemini 3.6 Flash en tête d’affiche. Présenté comme une IA « plus puissante et moins chère » que sa génération précédente, il affiche un score de 50 sur l’Intelligence Index v4.1, en retrait par rapport aux meilleurs Claude et GPT, mais avec un positionnement tarifaire agressif à 1,50 dollar en entrée et 7,50 dollars en sortie par million de tokens, selon Frandroid. La stratégie de Google n’est pas de viser le sommet du classement, mais de capter les cas d’usage à haut volume où le coût par requête pèse plus lourd que le dernier point de score sur un benchmark académique.

