---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1kg5j75-what-is-the-best-local-ai-model-for-coding-519de930
title: "r-localllama-comments-1kg5j75-what-is-the-best-local-ai-model-for-coding-519de930"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "OpenAI", "Z.ai"]
dates: []
keywords: ["llama", "arr", "claude", "deepseek", "gemini", "glm", "gpu", "mai", "mcp", "moe", "qwen", "reasoning"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1kg5j75-what-is-the-best-local-ai-model-for-coding-519de930.md
source_anchor: ""
source_lines: [1, 52]
sha256: f7b810b8586fa1bc91355fc4dae30435bbf676bade8f13e0942ba9e86fb3be3e
---

# r-localllama-comments-1kg5j75-what-is-the-best-local-ai-model-for-coding-519de930

Merci pour ton avis !
Explique-nous pourquoi ce contenu n’est pas utile.
      Quel est le meilleur modèle d'IA local pour le codage ? 
        
        
        
    
    
    Je cherche surtout pour Javascript/Typescript.
Et Frontend (HTML/CSS) + Backend (Node) s'il y en a de bons spécifiquement pour Tailwind.
Y a-t-il un modèle qui est au top maintenant ? J'ai lu un fil il y a 3 mois qui disait Qwen 2.5-Coder-32B mais Qwen 3 vient de sortir, donc je pensais que je devrais le télécharger directement.
Mais ensuite, j'ai vu dans LMStudio qu'il n'y a pas encore de Coder Qwen 3. Donc des alternatives pour l'instant ?
Section des commentaires
Peut-être juste attendre un peu pour Qwen 3 Coder. :)
https://x.com/ggerganov/status/1918373399891513571
Super ! Je me demandais à ce sujet et je n'avais rien vu à ce propos. Merci de partager !
J'espère qu'ils feront une version codeur des 30b MoE aussi, car l'inférence rapide fonctionnerait parfaitement pour la complétion dans l'IDE.
https://nitter.net/ggerganov/status/1918373399891513571
Oh c'est sympa. Tu penses que ça sortira quand ? Des rumeurs ? Ou des prédictions ?
Dans les derniers Radar Trends (mai 2025)...
"Pour ceux d'entre nous qui aiment garder notre IA à la maison, il y a maintenant DeepCoder, un modèle de 14B qui se spécialise dans le codage et qui prétend avoir des performances similaires à celles du o3-mini de OpenAI. Les ensembles de données, le code, les journaux d'entraînement et les optimisations système sont tous ouverts. https://www.together.ai/blog/deepcoder "
Oh j’adore ça. Aujourd’hui j’ai appris à propos de Radar Trends donc merci pour ça aussi. C'est tellement utile.
As-tu utilisé ce modèle ? J'avais entendu parler de DeepCoder mais j'ai oublié, comme j'utilise surtout des modèles en ligne, mais ouais, la plupart des problèmes peuvent être résolus localement, comme je fais beaucoup d'OCR sur des images pour attraper rapidement du texte (et non, les outils d'OCR ne fonctionnent pas puisque j'ai parfois besoin de texte dans un format spécifique que les outils d'OCR ne peuvent pas fournir)
Je n'ai pas essayé DeepCoder, mais j'ai essayé DeepScaleR, leur modèle mathématique de 1,8B. DeepScaleR est totalement légitime. C'est terrible à tout sauf en mathématiques, mais il peut résoudre la plupart des problèmes de mathématiques de niveau honneur au lycée (et quelques problèmes de physique) assez bien. Et c'est rapide, évidemment.
Donc, l'équipe derrière DeepCoder est apparemment douée pour les réglages fins très spécialisés.
GLM-4-0414 32B et Qwen 3 32B sont bons pour leur taille dans les tâches de développement web
Je vois beaucoup GLM récemment. Je vais jeter un œil.
GLM 4 est super pour le développement web. J'ai expérimenté avec et je peux attester que c'est génial. Il génère du code complet et verbeux, parfois au niveau d'un sonnet de Claude.
https://huggingface.co/collections/deepcogito/cogito-v1-preview-67eb105721081abe4ce2ee53
https://huggingface.co/collections/Tesslate/uigen-t15-reasoning-model-67e0fc3605add0af7c427c75
Merde, je viens de cliquer sur les liens et de jeter un œil. J'ai trouvé des perles sous-estimées. Je vais tester à quel point elles sont bonnes. Celui de l'UI (maintenant je comprends le nom) avait vraiment l'air pas mal.
Merci pour les liens. Ils sont les meilleurs pour quoi ? C'est la première fois que je les vois.
Par poids ouvert ? C'est encore Deepseek R1 / V3
Par quelque chose que tu pourrais réalistiquement faire tourner localement sans être riche en GPU ? Probablement Qwen3-32B. QwQ peut parfois comprendre des choses que Qwen3 ne peut pas, mais c'est presque inutile en tant qu'outil de codage en attendant tellement de tokens à générer
Est-ce que QwQ et Qwen sont différents ? Je pensais qu'ils étaient pareils. Je ne me suis pas trop intéressé aux trucs locaux, donc je ne sais pas.
La date limite de connaissances pour la plupart est quelque part en 2024 au mieux, donc la version la plus récente de Tailwind (4.x) n'est souvent pas incluse. Peut-être que les nouveaux modèles gemma (3) / qwen l'incluent ?
Ouais, la plupart des trucs autour de Tailwind v4 qui sont importants, c'est la transition de
tailwind.config.ts, donc je peux le faire manuellement, donc surtout j'ai juste besoin de ceux qui ont besoin de utilitaires, ce qui est probablement le cas de tous.
J'ai commencé à utiliser le Context7 MCP pour cette raison exacte. C'est essentiellement RAG alimenté par des appels d'outils sur des docs à jour pour une tonne de bibliothèques. https://context7.com/
J'ai traversé cette lutte pendant la semaine dernière. La syntaxe de Tailwind v4.x n'est pas la réponse par défaut pour les LLM, même si tu utilises une intégration éditeur avec quelque chose comme continue.dev et que tu passes ton fichier CSS. J'ai dû jongler entre tous mes suspects habituels et simplement faire un rondo jusqu'à obtenir une réponse qui aide vraiment. Assure-toi de préciser dans ton contexte que tu utilises Tailwind v4.1, ou quelle que soit la version que tu as.
Aujourd'hui, Llama 4 Maverick (via OR) fonctionnait bien pour moi. 0 % de réussite en zero shot, mais pratiquement 100 % après un commentaire de feedback. Claude Sonnet 3.7 (OR) s'est révélé étonnamment inutile. Même Gemini 2.5 Pro Preview a un peu bloqué, mais au final, c'est ça qui a le plus aidé.
Concernant les modèles locaux, je n'arrêtais pas de passer de qwen3-32b à glm4-32b, et parfois je revenais à qwen2.5-coder-32b pour essayer de faire quelque chose de désespéré.
Peut-être que ça ne sera pas aussi difficile pour toi, parce que tu intègres Tailwind de manière plus normale, contrairement à moi (Rust & Sycamore/Trunk). Mais j'ai été honnêtement choqué de voir à quel point il était difficile pour moi d'obtenir une aide de l'IA sur certaines de ces choses, en tant que personne qui touche rarement à la création web front-end et qui s'occupe généralement de choses plus bas niveau.
(Et oui, j'attends avec impatience qwen3 coder...)
Haha, pour moi c'était facile puisque j'utilise principalement React + Tailwind, qui est rempli d'exemples sur le web.
Lequel des modèles que tu as listés est le plus petit ? Je veux le meilleur local + petite taille puisque ma M4 n'a que 16 Go de mémoire.
GLM-4-0414 32B, c'était dans mon test html le meilleur et même mieux que o4… donc si html et js sont une chose, je ferais un essai.
Je suis juste en train de lire à ce sujet. Je vais jeter un œil.
GLM
Je fais surtout tourner DeepSeek V3 (le quant UD-Q4_K_XL) et parfois R1. J'ai aussi l'intention d'essayer R1T Chimera une fois que j'aurai terminé de le télécharger.
Pour des tâches simples à de complexité moyenne, le modèle Qwen3 30B fonctionne bien, j'ai trouvé particulièrement intéressant la version A6B qui a été ajustée pour utiliser une quantité double d'experts (contrairement à forcer le modèle à utiliser plus d'experts, ce qui ne produit généralement pas d'amélioration) - il reste rapide mais semble mieux s'en sortir avec des choses plus complexes. Cependant, la différence par rapport à A3B est subtile et je ne suis pas exactement sûr qu'il soit vraiment meilleur en général puisqu'il n'y a eu que des tests limités jusqu'à présent.
Quel modèle utiliser dépend de votre matériel, par exemple si vous avez un seul GPU et de la RAM en double canal, alors Qwen3 30B A3B pourrait être le plus intéressant. Vous pouvez aussi essayer l'ajustement A6B mentionné ci-dessus, ou pour l'inférence uniquement sur CPU (ou avec très peu de VRAM / une carte lente) https://huggingface.co/DavidAU/Qwen3-30B-A1.5B-High-Speed pourrait intéresser puisqu'il utilise seulement la moitié des experts par rapport à la version standard A3B, donc deux fois plus rapide, même si au prix de perdre un peu de qualité.
