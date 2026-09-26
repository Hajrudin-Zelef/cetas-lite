---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1q2wm33-how-capable-is-gptoss120b-and-what-are-your-1fecf34e-3
title: "r-localllama-comments-1q2wm33-how-capable-is-gptoss120b-and-what-are-your-1fecf34e"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Apple", "Google", "OpenAI"]
dates: []
keywords: ["llama", "amd", "benchmarks", "chatgpt", "fp4", "gpu", "leaderboard", "llama.cpp", "mcp", "mxfp4", "nvidia"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1q2wm33-how-capable-is-gptoss120b-and-what-are-your-1fecf34e.md
source_anchor: ""
source_lines: [50, 72]
sha256: a9d221a5434d9cd127a6eaffe5e0d24649d3b70bbf6c20ef4e684a8d85a8af5a
---

# r-localllama-comments-1q2wm33-how-capable-is-gptoss120b-and-what-are-your-1fecf34e

Avec 96 Go de DDR4, 1x3090, 32k context, full SWA, context shift (dans llama.cpp), j'ai obtenu 12 t/s en vitesse de génération et environ 400t/s en traitement des prompts. Correct pour des conversations normales. Vous devriez être capable d'obtenir des résultats similaires avec cette configuration.
J'ai OSS GPT qui tourne sur deux 3090s pour le fun, je le connecte à mon téléphone via un logiciel que j'ai construit.
Je l'utilise comme un Google/Stack Overflow/Quora plus intelligent et beaucoup plus efficace. Et oui, de temps en temps pour coder. Je trouve ça génial. (Mes compétences en codage (et ma patience) sont faibles. YMMV)
Je ne lui fais jamais entièrement confiance, et plus tu t'enfonces dans le terrier du lapin sur un sujet de niche, moins tu devrais lui faire confiance. Mais la plupart du temps, ça te pointera dans une direction utile.
Cherche les versions ablatées. Ou les versions avec -heretic dans le nom.
C'est super pratique, j'adorerais voir un modèle avec 210b paramètres actifs de 8 à 10b comme m2.1 mais entraîné nativement en format q4 ou mxpf4 comme oss 120b, en supposant que ça fasse environ 105 à 110 Go, ce serait le top du top que je puisse faire tourner en local avec 128 Go de RAM et 16 Go de VRAM
got-oss:120b 4-bit est mon option de chat préférée.
J'ai été le 8-bit avec mes MCP dans LLM studio juste pour déconner et il y a de légers avantages, mais 4-bit est le point idéal pour mes besoins. De plus, pas assez d'amour pour gpt-oss:20b en comparaison.
J'ai 256 gb de mémoire unifiée et je dois dire que le compromis est minime pour la plupart de ce que je recherche. Bien que, lors de la recherche sur le web, j'ai eu plus d'échecs avec 20b.
Deuxième partie de ta question
Je crois qu'on va avoir un modèle MASSIVE quantifié à 2-but en provenance de Chine avec des benchmarks incroyables.
Et un changement MASSIF dans la façon dont le marché considère les puces Mac Silicon et leurs coûts de mémoire unifiée qui sont extrêmement précieux pour cette année et cette année seulement, car ils ont réduit les coûts, maintenant les choses basses.
Par exemple, quelqu'un utilisera un studio m3 Ultra 256 gb de mémoire unifiée (maintenant plus ancien) pour la mémoire et le cluster avec le modèle le moins cher qui est livré avec la puce m5 pro/max pour le traitement.
Et tout cela devient extrêmement précieux alors qu'Apple commence à lancer ses propres modèles qui nous intéressent vraiment.
(D'où leur concentration sur le fait que Siri utilise un LLM local sur votre appareil plutôt que ChatGPT finalement.)
Je ne suggérerais jamais qu'Apple aille au-delà du GPU nvidia + cuda, mais ils vont combler l'écart et aussi créer une valeur extrême.
Rappelle-toi MacBook Pro modèle de base : $10 = 1gb Mac Studio m3 ultra : $20 = 1gb 5090 : $100 = 1gb
Les Macs ne surpasseront pas, mais leur valeur est très forte.
Bien que je dirai que le DGX Spark est incroyable et très fort en valeur.
C'est pas que 8-bit et 4-bit, c'est à peu près la même taille à cause de la quantification MXFP4 originale, non ?
N'oubliez pas Stryx Halo, un truc du genre le Bosgame M5 avec 128GB est beaucoup moins cher que du matos Apple similaire. En plus, tu peux faire tourner Linux dessus. Quand il y aura un stryx halo avec 256GB de RAM, ça pourrait tout changer pour les LLMs locaux. Aussi, sur ces mini PC stryx halo, tu peux brancher deux eGPU en plus, une via oculink (adaptateur nvme) et une via usb4. Comme ça, tu peux avoir CUDA en plus d'un CPU/iGPU AMD costaud.
C'est un super modèle pour plein d'applications, mais faut vraiment pas le faire tourner en dessous de fp4.
Si tu regardes le classement de leaderboard.neurometric.ai, qui teste les modèles sur des tâches liées au travail, gpt-oss-120b était le meilleur modèle global, battant même les modèles anthropic sur de nombreuses tâches.
