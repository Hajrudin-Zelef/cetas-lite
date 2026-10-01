---
id: collect-261001-ia-llm/ia-llm/hallucination-ia-2026-deepseek-94-vs-gpt-6-astra-51-2
title: "Comparaison basique du taux de réponses \"je ne sais pas\" entre deux fournisseurs"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Microsoft", "Mistral", "OpenAI", "xAI"]
dates: []
keywords: ["agents", "agi", "astra", "benchmark", "benchmarks", "claude", "deepseek", "fable 5", "gemini", "gpt-5.6", "gpt-6", "grok"]
source: docs/RAG/collect-261001-ia-llm/hallucination-ia-2026-deepseek-94-vs-gpt-6-astra-51.md
source_anchor: ""
source_lines: [27, 64]
sha256: 245f50450f2a661584d742a035363dd7f760973d7428b5910db0cae43c381055
---

# Comparaison basique du taux de réponses "je ne sais pas" entre deux fournisseurs

| Modèle | Éditeur | Sortie | Contexte | Indice AA-Omniscience | Taux d’hallucination (max effort) | Prix entrée/sortie ($/M tokens) | 
|---|---|---|---|---|---|---|
| GPT-6 Astra | OpenAI | 3 sept. 2026 | Non précisé publiquement | Non isolé (Index Intelligence : 53) | 51 % (contre 92 % pour Sol) | 10 $ / 50 $ | 
| GPT-5.6 Sol | OpenAI | 2026 | Non précisé publiquement | Index Intelligence : 47 | 92 % | 4 $ / 20 $ | 
| Claude Opus 5 | Anthropic | 2026 | 1 M tokens | 37 (précision 61 %) | 61 % | 5 $ / 25 $ | 
| Claude Fable 5.1 | Anthropic | été 2026 | 1 M tokens | 43,5 (précision 67,2 %) | Non isolé publiquement | 10 $ / 50 $ | 
| Gemini 3.1 Pro |  | 2026 | ~1 M tokens | 32 (précision 55 %) | 51 % | 2 $ / 12 $ (jusqu’à 200K) | 
| Gemini 3 Deep Think |  | 2026 | 1 M tokens | Précision 41 % | Non isolé publiquement | 2 $ / 12 $ (jusqu’à 200K) | 
| DeepSeek V4 Pro | DeepSeek | avril 2026 | 1 M tokens | -10 | 94 % | 0,435 $ / 0,87 $ | 
| DeepSeek V4 Flash | DeepSeek | avril 2026 (V4.1 le 10 sept.) | 1 M tokens | -23 | 96 % | 0,14 $ / 0,28 $ | 
| Mistral Large 3 | Mistral AI (France) | 2 déc. 2025 | ~262K tokens | Non testé publiquement | Non testé publiquement | 0,50 $ / 1,50 $ | 
| Grok 4.6 | xAI | 12 août 2026 | 500K tokens | 30,5 | Non isolé publiquement | 2 $ / 6 $ (jusqu’à 200K) | 
| Qwen3.8-Max | Alibaba | 3 août 2026 | ~1 M tokens | Non noté (“still unscored”) | Non isolé publiquement | 2 $ / 6 $ | 

Trois lectures ressortent immédiatement de ce tableau. D’abord, les modèles les moins chers ne sont pas nécessairement les moins fiables sur le plan du raisonnement pur, mais ils affichent systématiquement les taux d’hallucination les plus élevés du panel : DeepSeek V4 Pro et V4 Flash dépassent 90 %. Ensuite, l’indice AA-Omniscience et le prix par token n’évoluent pas dans le même sens : Claude Fable 5.1 obtient l’indice le plus élevé du tableau (43,5) tout en étant l’un des modèles les plus chers. Enfin, Mistral Large 3, seul représentant européen de ce comparatif, n’apparaît tout simplement pas dans les benchmarks d’hallucination indépendants consultés à la date de rédaction, un vide statistique qui mérite un traitement à part plus loin dans cet article.

## GPT-6 Astra : le chiffre d’OpenAI qui a changé trois fois en une semaine

GPT-6 Astra a été lancé le 3 septembre 2026 avec un argumentaire centré sur la fiabilité factuelle, comme le détaille notre actualité sur les conditions d’accès restreint à GPT-6 Astra. La page de lancement d’OpenAI annonçait un taux d’hallucination interne de 4,2 %, contre 12,2 % pour GPT-5.6 Sol, soit une division par près de trois. Le modèle affiche par ailleurs des scores impressionnants sur les benchmarks de raisonnement pur, avec 97,6 % sur FrontierMath Tier 4 et 99,9 % sur ARC-AGI-3 selon les chiffres relayés par plusieurs analystes techniques.

Le 4 septembre 2026, soit un jour après le lancement, Fortune a publié une enquête montrant que ce chiffre de 4,2 % avait été modifié à deux reprises sur la page officielle d’OpenAI dans les heures suivant l’annonce. Selon les captures d’archive consultées par le magazine, le taux est brièvement passé à 2 % avant de revenir à 4,2 %, en même temps que quatre autres métriques de la fiche produit étaient ajustées. Ce type de révision post-lancement, sur un indicateur aussi central que le taux d’hallucination, a alimenté les critiques sur la fiabilité des chiffres auto-déclarés par les éditeurs de modèles.

Le test indépendant AA-Omniscience d’Artificial Analysis, publié le 9 septembre 2026, apporte un éclairage différent. Sur ce benchmark, GPT-6 Astra fait chuter son taux d’hallucination à effort maximal de 92 % (le score de GPT-5.6 Sol) à 51 %, une amélioration réelle et mesurée par un tiers indépendant, mais qui reste très éloignée du chiffre de 4,2 % mis en avant par OpenAI. Artificial Analysis précise explicitement dans son article que cette progression s’accompagne d’une hausse de précision de 4 points sur le même test. Autrement dit : Astra hallucine deux fois moins souvent que son prédécesseur sur cette tâche exigeante, mais il continue de se tromper avec assurance plus d’une fois sur deux lorsqu’il ne connaît pas la réponse et choisit malgré tout de répondre.

Pour un lecteur français qui doit choisir un modèle pour un usage professionnel, la leçon de l’épisode Astra n’est pas que GPT-6 Astra serait un mauvais modèle : ses progrès sur AA-Omniscience sont réels et documentés par un tiers. La leçon est méthodologique : ne jamais retenir uniquement le chiffre mis en avant sur la page de lancement d’un éditeur, et toujours croiser avec au moins une source indépendante avant de fonder une décision d’architecture sur un taux de fiabilité annoncé.

## Claude Opus 5 et Claude Fable 5.1 : Anthropic mise sur la prudence, à quel prix

Anthropic communique historiquement moins sur des chiffres d’hallucination internes que OpenAI, préférant mettre en avant des scores de raisonnement et de sécurité, une position que l’on retrouve déjà dans notre analyse de Claude Opus 5 en tête du classement LLM. Sur AA-Omniscience, Claude Opus 5 obtient un indice de 37 avec une précision de 61 % et un taux d’hallucination de 61 %, un résultat honorable mais qui ne le place pas en tête du classement de ce comparatif : il reste moins performant que GPT-6 Astra (51 %) et que Gemini 3.1 Pro (51 % également) sur ce test précis.

Claude Fable 5.1, la déclinaison orientée agents et raisonnement long d’Anthropic, obtient en revanche l’indice AA-Omniscience le plus élevé de tout ce comparatif : 43,5, avec une précision de 67,2 %, ce qui en fait le modèle le mieux calibré du panel sur ce test de connaissances selon les données disponibles. Son taux d’hallucination propre n’a pas été isolé publiquement dans les sources consultées. Seule une configuration antérieure, Claude Fable 5 avec mécanisme de repli (fallback), affiche 64 % d’hallucination pour 65 % de précision, un chiffre qui ne doit pas être confondu avec celui de la version 5.1.

Sur le plan tarifaire, Anthropic facture Claude Opus 5 à 5 $ par million de tokens en entrée et 25 $ en sortie selon la page de tarification officielle d’Anthropic, tandis que Claude Fable 5.1 grimpe à 10 $ et 50 $. Les deux modèles disposent d’une fenêtre de contexte de 1 million de tokens, avec une sortie maximale d’environ 128 000 tokens, et bénéficient d’une lecture en cache à 0,25 $ par million de tokens pour Fable 5.1, une réduction de 75 % par rapport à la génération précédente. Pour une entreprise qui traite de gros volumes de documents identiques (contrats types, FAQ internes), ce mécanisme de cache change sensiblement le calcul de coût réel par rapport au prix catalogue brut.

## Gemini 3.1 Pro et Gemini 3 Deep Think : la rigueur de Google a un prix caché

Gemini 3.1 Pro affiche sur AA-Omniscience un indice de 32, avec une précision de 55 % et un taux d’hallucination de 51 %, strictement identique à celui de GPT-6 Astra sur ce même benchmark. C’est une donnée intéressante pour les équipes qui arbitrent entre les deux écosystèmes cloud : à fiabilité factuelle comparable sur ce test précis, le choix peut alors se faire sur d’autres critères, notamment l’intégration existante avec Google Workspace ou Microsoft Azure.

