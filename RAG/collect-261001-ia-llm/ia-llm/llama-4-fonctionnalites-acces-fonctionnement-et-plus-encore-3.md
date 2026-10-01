---
id: collect-261001-ia-llm/ia-llm/llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore-3
title: "llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore"
domain: ia-llm
role: reference
task: reference
actors: ["Meta", "Microsoft", "vLLM"]
dates: []
keywords: ["llama", "benchmarks", "gpu", "llama.cpp", "multimodal", "open-weight", "scout", "vllm"]
source: docs/RAG/collect-261001-ia-llm/llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore.md
source_anchor: ""
source_lines: [134, 162]
sha256: b1c30525e7a4de6ae7a0e3ccb9a1c26bfd7774d85e2f48f26721787fdf05527a
---

# llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore

### Puis-je ajuster finement Llama 4 Scout ou Maverick sur mes propres données ?

Oui, les deux modèles sont open-weight et peuvent être ajustés finement.

### Puis-je faire tourner les modèles Llama 4 en local ?

Vous pouvez exécuter Scout en local si vous avez accès à un GPU haut de gamme (comme un A100 ou un H100). Maverick est nettement plus grand et requiert généralement plusieurs GPU ou une infrastructure distribuée. Pour des tests légers, des versions quantifiées de Scout peuvent convenir sur du matériel grand public avec des outils comme llama.cpp ou vLLM.

### Quelles sont les exigences matérielles pour Llama 4 Scout ?

Scout est conçu pour tenir sur un seul GPU H100. Cela dit, selon la longueur de contexte et la taille des lots, vous pourriez faire tourner des versions plus petites ou des modèles quantifiés sur des GPU de gamme inférieure comme l’A100 ou même une RTX 4090, avec des performances réduites.

### Llama 4 est-il multilingue ?

Oui — Maverick et Behemoth affichent de très bons résultats sur des benchmarks multilingues comme Multilingual MMLU. Bien que Meta n’ait pas publié de détails par langue, les premiers benchmarks suggèrent de bonnes performances sur les principales langues non anglaises.

### Puis-je utiliser Llama 4 dans des produits commerciaux ?

Oui, sauf si votre entreprise ou produit dépasse 700 millions d’utilisateurs actifs mensuels, auquel cas vous devrez obtenir une licence spéciale auprès de Meta. Pour la plupart des startups, chercheurs et développeurs individuels, la licence standard s’applique.

### Puis-je distiller mon propre modèle à partir de Llama Behemoth ?

Pas encore. Behemoth n’a pas été publié, et rien n’indique quand Meta le rendra public. Cela dit, Meta a utilisé Behemoth en interne pour distiller Scout et Maverick : s’il est publié, il pourrait servir de base à d’autres distillations.

### Quelle est la différence entre Llama 3.1, Llama 3.3 et Llama 4 ?

Llama 3.1 et 3.3 étaient des modèles denses avec un support multimodal limité ou absent. Llama 4 passe à une architecture mixture-of-experts et ajoute un entraînement multimodal natif. Scout et Maverick intègrent aussi des fenêtres de contexte plus longues et des techniques de post-entraînement améliorées.

Je suis rédacteur et écrivain et je couvre les blogs, les tutoriels et les actualités sur l'IA, en m'assurant que tout est conforme à une stratégie de contenu solide et aux meilleures pratiques en matière de référencement. J'ai rédigé des cours de science des données sur Python, les statistiques, les probabilités et la visualisation des données. J'ai également publié un roman primé et je consacre mon temps libre à l'écriture de scénarios et à la réalisation de films.
