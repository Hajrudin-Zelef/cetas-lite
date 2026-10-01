---
id: collect-261001-ia-llm/ia-llm/benchmark-lara-les-ia-violent-l-ai-act-a-93-2026-2
title: "benchmark-lara-les-ia-violent-l-ai-act-a-93-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "EU", "Google", "Hugging Face", "Meta", "Microsoft", "Mistral", "OpenAI"]
dates: []
keywords: ["benchmark", "agents", "benchmarks", "claude", "deepseek", "gemini", "gpt-5.6", "luna", "mistral", "open-weight", "opus 4", "opus 5"]
source: docs/RAG/collect-261001-ia-llm/benchmark-lara-les-ia-violent-l-ai-act-a-93-2026.md
source_anchor: ""
source_lines: [39, 78]
sha256: e6b56fbaf88c40a65b6d1d8ee90ca816aa8838cedc8994b3e0050a9f7f3af9c8
---

# benchmark-lara-les-ia-violent-l-ai-act-a-93-2026

Pour un groupe comme Alphabet, Microsoft ou Meta, dont le chiffre d’affaires mondial annuel se compte en centaines de milliards de dollars, le plafond de 7 % représente un risque financier qui se chiffre potentiellement en dizaines de milliards d’euros. C’est cette disproportion entre le montant fixe (35 millions d’euros) et le pourcentage du chiffre d’affaires qui rend le régime de sanctions de l’AI Act particulièrement dissuasif pour les géants américains et chinois de l’IA, à la différence de startups plus modestes pour lesquelles le plafond fixe s’appliquera plus souvent.

| Palier de sanction | Montant maximal | Alternative en % du CA mondial | Infractions concernées | 
|---|---|---|---|
| Palier 1 — le plus élevé | 35 M€ | 7 % du chiffre d’affaires mondial | Pratiques interdites, manipulation, notation sociale | 
| Palier 2 | 15 M€ | 3 % du chiffre d’affaires mondial | Non-conformité systèmes à haut risque et modèles GPAI | 
| Palier 3 | 7,5 M€ | 1 % du chiffre d’affaires mondial | Informations incorrectes ou trompeuses aux autorités | 
| Obligations GPAI de base | En vigueur depuis | 2 août 2025 | Documentation, données d’entraînement, droit d’auteur | 
| Pouvoir de sanction effectif | En vigueur depuis | 2 août 2026 | Contrôle et amendes opposables aux fournisseurs GPAI | 

## Le Digital Omnibus : 16 mois de sursis pour les systèmes à haut risque

Tout ne bascule pourtant pas le 2 août 2026. Un règlement d’assouplissement, le Digital Omnibus (Regulation (EU) 2026/1744), entré en vigueur le 27 juillet 2026, reporte l’application des obligations relatives aux systèmes d’IA à haut risque autonomes, visés à l’article 6(2) et à l’annexe III de l’AI Act, du 2 août 2026 au 2 décembre 2027, soit un sursis de 16 mois. Sont concernés les usages en ressources humaines, éducation, notation de crédit et certains contextes proches des forces de l’ordre. Les systèmes à haut risque intégrés dans des produits déjà régulés par ailleurs (annexe I) bénéficient d’un délai encore plus long, jusqu’au 2 août 2028.

Ce report ne touche en revanche ni l’interdiction des pratiques prohibées, ni les obligations déjà applicables aux modèles à usage général depuis août 2025, ni les exigences de transparence de l’article 50, qui imposent par exemple d’informer un utilisateur qu’il interagit avec une IA ou qu’un contenu a été généré ou modifié artificiellement. Le Digital Omnibus centralise également la supervision : le Bureau de l’IA obtient une compétence exclusive sur les systèmes fondés sur un modèle GPAI développé par le même groupe, ainsi que sur l’IA intégrée aux très grandes plateformes en ligne et très grands moteurs de recherche désignés au titre du DSA.

Le résultat est un calendrier à double vitesse : les obligations de transparence et le contrôle des modèles fondamentaux s’appliquent dès maintenant, avec un risque de sanction réel, tandis que les usages professionnels les plus sensibles (recrutement automatisé, scoring de crédit) bénéficient d’un délai supplémentaire pour se mettre en conformité. Pour les entreprises françaises qui déploient des agents IA dans ces domaines, cela laisse une fenêtre jusqu’à fin 2027 pour corriger les manquements que des benchmarks comme LARA commencent déjà à documenter.

## Pourquoi la nouvelle génération de modèles change (un peu) la donne

Un détail technique mérite d’être souligné : le benchmark LARA a testé Claude Opus 4.1 et Gemini 3.1 Pro, deux générations déjà dépassées par les lancements de l’été 2026. Anthropic a fait passer Claude Sonnet 5 en modèle par défaut le 30 juin 2026, puis a porté Claude Opus 5 en disponibilité générale le 24 juillet 2026, avec une fenêtre de contexte étendue à 1 million de tokens. Google a de son côté déployé trois nouveaux modèles Gemini le 21 juillet 2026, dont Gemini 3.6 Flash, présenté comme réduisant la consommation de tokens jusqu’à 17 % par rapport à Gemini 3.5 Flash. OpenAI a lancé la famille GPT-5.6 (déclinée en Sol, Terra et Luna) le 9 juillet 2026, et DeepSeek a fait passer son modèle V4 de la préversion à la disponibilité générale le 20 juillet 2026, avec une variante V4-Pro à 1 600 milliards de paramètres en mixture d’experts et 49 milliards de paramètres actifs par requête.

Aucun de ces modèles de dernière génération n’a encore été évalué publiquement par Aithos sur le protocole LARA. Cela signifie que les scores de 54 % pour Claude et 10 % pour Gemini reflètent un état antérieur des modèles d’Anthropic et de Google, pas nécessairement leurs versions actuellement commercialisées. Les fournisseurs ont un argument de défense légitime : leurs modèles évoluent plus vite que les cycles de benchmark indépendant. Mais cet argument a ses limites, car rien ne garantit que les versions les plus récentes progressent réellement sur les critères de conformité légale plutôt que sur les seuls benchmarks de raisonnement ou de codage, sur lesquels la compétition entre laboratoires reste focalisée.

## Comparatif : GPT-5.6, Claude Opus 5, Gemini 3.6 Flash et DeepSeek V4 face à l’enjeu réglementaire

Au-delà du seul enjeu de conformité, la génération de modèles actuellement commercialisée en Europe illustre des stratégies de positionnement différentes, qui auront un impact direct sur la manière dont chaque fournisseur gère le risque réglementaire.

| Modèle | Fournisseur | Lancement | Fenêtre de contexte | Particularité pertinente pour la conformité | 
|---|---|---|---|---|
| Claude Opus 5 | Anthropic | 24 juillet 2026 (GA) | 1 000 000 tokens | Meilleur score de départ sur LARA (génération précédente à 54 %) | 
| GPT-5.6 (Sol/Terra/Luna) | OpenAI | 9 juillet 2026 | Variable selon variante | Score individuel non publié par Aithos | 
| Gemini 3.6 Flash |  | 21 juillet 2026 | Contexte long, -17 % de tokens vs 3.5 Flash | Génération précédente (3.1 Pro) au plus bas du panel, ≈ 10 % | 
| DeepSeek V4-Pro | DeepSeek | 20 juillet 2026 (GA) | 1 000 000 tokens | Open-weight sous licence MIT, hors juridiction directe UE | 
| Mistral AI (gamme Large/Le Chat) | Mistral AI | En continu 2026 | Variable selon modèle | Seul fournisseur de frontière domicilié dans l’UE, soumis aux mêmes règles | 

Ce tableau met en évidence un paradoxe intéressant pour le marché européen : DeepSeek, fournisseur chinois dont les modèles sont diffusés en open-weight sous licence MIT sur Hugging Face, échappe en partie au même niveau de contrôle direct que les fournisseurs américains disposant d’une présence commerciale et de bureaux en Europe, tout en étant massivement utilisé par des développeurs européens via des API tierces ou des déploiements locaux. Mistral AI, à l’inverse, reste le seul fournisseur de modèles de frontière domicilié dans l’Union européenne, ce qui en fait à la fois le champion de la souveraineté numérique française et le fournisseur le plus directement exposé à la supervision du Bureau de l’IA sur son propre sol.

## Le contexte historique : du RGPD 2018 à l’AI Act 2024-2026

