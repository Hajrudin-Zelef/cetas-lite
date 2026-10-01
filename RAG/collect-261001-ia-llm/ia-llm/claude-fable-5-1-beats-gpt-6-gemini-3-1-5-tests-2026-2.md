---
id: collect-261001-ia-llm/ia-llm/claude-fable-5-1-beats-gpt-6-gemini-3-1-5-tests-2026-2
title: "claude-fable-5-1-beats-gpt-6-gemini-3-1-5-tests-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Microsoft", "Mistral", "OpenAI"]
dates: []
keywords: ["claude", "gemini", "gpt-6", "astra", "benchmark", "benchmarks", "chatgpt", "deepseek", "fable 5", "foundry", "llama", "mistral"]
source: docs/RAG/collect-261001-ia-llm/claude-fable-5-1-beats-gpt-6-gemini-3-1-5-tests-2026.md
source_anchor: ""
source_lines: [43, 82]
sha256: f2c5cee8a546ea95f84ec27135362619389952f22fd8d18d37d7808c90bcfc0b
---

# claude-fable-5-1-beats-gpt-6-gemini-3-1-5-tests-2026

GPT-6 Astra a été présenté par OpenAI comme son modèle le plus intelligent et le plus aligné à ce jour. Le déploiement s’est fait en deux temps : accès restreint aux programmes Trusted Access et Daybreak dès le 3 septembre 2026, puis ouverture aux abonnés ChatGPT Plus, Pro, Business et Enterprise ainsi qu’à l’API publique le lendemain. Cette prudence n’est pas cosmétique : selon la fiche de sécurité publiée par OpenAI, GPT-6 Astra atteint un niveau de capacité cybersécurité qualifié de “critique” au sens du cadre de préparation interne de l’entreprise, ce qui signifie qu’il peut identifier et exploiter de façon autonome des failles de sécurité jusque-là inconnues.

Sur la tarification, GPT-6 Astra affiche exactement les mêmes tarifs de base que Claude Fable 5.1 : 10 dollars par million de tokens en entrée, 50 dollars en sortie. L’entrée en cache tombe à 1 dollar par million de tokens, tandis que l’écriture en cache coûte 12,50 dollars. Sur Microsoft Foundry, la tarification se décline en paliers selon la zone géographique et la longueur de contexte : environ 10 dollars (contexte court) à 20 dollars (contexte long) par million de tokens en entrée sur la zone globale, et 11 à 22 dollars sur la zone de données américaine dédiée.

La fenêtre de contexte de GPT-6 Astra atteint 1,05 million de tokens, répartis entre environ 922 000 tokens d’entrée et 128 000 tokens de sortie maximum. Le modèle propose cinq paliers de raisonnement (low, medium, high, xhigh, max), une granularité supérieure à celle de Claude Fable 5.1, pensée pour les équipes qui veulent arbitrer finement entre coût de calcul et profondeur de résolution selon la tâche.

Côté modalités, GPT-6 Astra reste en retrait par rapport à Gemini : le modèle accepte le texte et l’image en entrée, avec une sortie textuelle enrichie par un accès à des outils (navigation web, interpréteur de code, shell hébergé, usage d’ordinateur, génération d’images). OpenAI met en avant ces capacités “d’usage d’ordinateur”, c’est-à-dire la possibilité pour le modèle de piloter une interface graphique comme le ferait un humain, un axe stratégique qui le distingue nettement des deux autres modèles de ce comparatif.

## Gemini 3.1 Pro Preview : le flagship multimodal le moins cher de Google

Contrairement à ses deux concurrents, Gemini 3.1 Pro Preview n’est pas un lancement de rentrée : le modèle est disponible depuis le 19 février 2026 et a pris la relève de Gemini 3 Pro Preview, sorti en novembre 2025. Google n’a pas annoncé de date de retrait pour l’ancien modèle, qui reste listé comme actif dans le catalogue, mais Gemini 3.1 Pro Preview est désormais présenté comme le modèle Pro de référence dans la documentation développeur et dans les guides tarifaires tiers.

L’écart le plus frappant concerne le prix. Gemini 3.1 Pro Preview facture 2 dollars par million de tokens en entrée et 12 dollars en sortie pour les requêtes jusqu’à 200 000 tokens, un tarif qui grimpe à 4 dollars en entrée et 18 dollars en sortie au-delà de ce seuil. Même dans son palier le plus cher, Gemini 3.1 Pro Preview reste nettement moins onéreux que Claude Fable 5.1 et GPT-6 Astra sur toute la plage de contexte usuelle.

Sur la multimodalité, Gemini 3.1 Pro Preview conserve l’avantage hérité de la famille Gemini : le modèle accepte nativement le texte, le code, les images, l’audio, la vidéo et les fichiers PDF en entrée. Les spécifications techniques de Google documentent une prise en charge allant jusqu’à environ 45 minutes de vidéo avec audio (ou une heure sans audio) par requête, jusqu’à dix vidéos dans un même prompt, et un traitement de fichiers PDF avec reconnaissance optique de caractères intégrée. Cette capacité à ingérer des heures de contenu audio et vidéo en une seule requête reste, à ce jour, une spécificité que ni Claude Fable 5.1 ni GPT-6 Astra ne proposent nativement.

Le plafond de sortie de 65 536 tokens, environ deux fois plus bas que celui de ses deux concurrents, constitue la principale contrepartie. Pour générer un très long document ou un gros volume de code en une seule réponse, ce plafond peut nécessiter une segmentation de la tâche en plusieurs appels, un point à anticiper lors du choix d’architecture applicative.

## Benchmarks : ce que disent les classements indépendants

La comparaison des scores bruts entre ces trois modèles reste un exercice délicat en septembre 2026 : plusieurs organismes de référence, dont Artificial Analysis, ont révisé leur méthodologie de notation (passage de la version 4.2 à la version 4.3 de leur Intelligence Index) après des interrogations sur la fiabilité du score initial attribué à GPT-6 Astra. Dans ce contexte mouvant, Anthropic revendique que Claude Fable 5.1 occupe la première place du classement d’intelligence d’Artificial Analysis, une affirmation reprise par plusieurs médias spécialisés sans qu’un score numérique précis et stabilisé ne soit publiquement détaillé au moment de la rédaction de cet article.

Sur le plan communautaire, LMArena, la plateforme de comparaison par votes humains, avait déjà vu Claude Fable 5 s’installer en tête de son classement à la mi-juin 2026, confirmant une dynamique de forte adoption côté Anthropic avant même la sortie de la version 5.1. Pour Gemini, un tracker de prix indépendant associe à Gemini 3.1 Pro un score de 94,3 % sur le benchmark scientifique GPQA Diamond, une mesure qui positionne le modèle comme une référence sur les tâches de raisonnement scientifique multimodal, même si ce chiffre provient d’un agrégateur tiers et non d’une publication officielle de Google.

Un troisième angle de lecture vient des benchmarks orientés langue française. Un comparatif publié début septembre 2026 par un cabinet d’analyse français, portant sur 18 modèles dont GPT-5, Claude, Gemini, Mistral, Llama 4 et DeepSeek, confirme que les modèles ouverts comme DeepSeek V3 restent compétitifs sur le rapport coût-performance en français, même si les modèles propriétaires les plus récents comme Claude Fable 5.1 et GPT-6 Astra n’y figuraient pas encore au moment de la publication, ces deux modèles étant sortis quelques jours plus tard. Cette absence illustre bien la difficulté à comparer des modèles dont le rythme de sortie dépasse celui des cycles de benchmarking indépendant, un problème déjà documenté sur des sujets connexes comme le taux d’hallucination de GPT-6 Astra face à DeepSeek.

Le point commun à ces trois sources : aucune ne permet, à ce jour, d’établir un classement définitif et stabilisé entre Claude Fable 5.1, GPT-6 Astra et Gemini 3.1 Pro Preview sur un score unique et directement comparable. C’est pourquoi ce comparatif privilégie les données vérifiables (tarifs officiels, fenêtres de contexte, modalités, politiques d’accès) plutôt qu’un score composite qui resterait sujet à interprétation.

## Tarification : le tableau des prix par million de tokens

C’est sur ce terrain que l’écart entre les trois modèles est le plus net et le plus facilement vérifiable, puisqu’il s’agit de tarifs publiés directement par les éditeurs ou leurs partenaires cloud.

| Modèle | Prix entrée ($/M tokens) | Prix sortie ($/M tokens) | Cache lecture | Cache écriture | 
|---|---|---|---|---|
| Claude Fable 5.1 | 10,00 $ | 50,00 $ | 0,25 $/M | Non communiqué | 
| Claude Fable 5.1 (batch) | 5,00 $ | 25,00 $ | — | — | 
| GPT-6 Astra | 10,00 $ | 50,00 $ | 1,00 $/M | 12,50 $/M | 
| Gemini 3.1 Pro Preview (≤ 200k tokens) | 2,00 $ | 12,00 $ | ≈ 0,20 $/M | ≈ 0,375 $/M | 
| Gemini 3.1 Pro Preview (> 200k tokens) | 4,00 $ | 18,00 $ | ≈ 0,40 $/M | ≈ 0,375 $/M | 

