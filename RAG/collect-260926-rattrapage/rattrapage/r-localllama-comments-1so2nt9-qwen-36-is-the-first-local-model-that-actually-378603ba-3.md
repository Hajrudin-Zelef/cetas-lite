---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1so2nt9-qwen-36-is-the-first-local-model-that-actually-378603ba-3
title: "r-localllama-comments-1so2nt9-qwen-36-is-the-first-local-model-that-actually-378603ba"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Unsloth"]
dates: []
keywords: ["qwen", "claude"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1so2nt9-qwen-36-is-the-first-local-model-that-actually-378603ba.md
source_anchor: ""
source_lines: [51, 57]
sha256: d4ef0e95b38fff96063724f2a996aa108ff22d4c0d83d97352bd326191dbd1de
---

# r-localllama-comments-1so2nt9-qwen-36-is-the-first-local-model-that-actually-378603ba

2025 a été l'année des LLM locaux, où les sauts de qualité étaient visibles chaque trimestre. C'est bon de voir que ça ne semble pas ralentir encore. Maintenant, nous sommes déjà dans une situation où des modèles locaux de gamme inférieure ou intermédiaire peuvent gérer certaines choses mieux que les modèles SOTA grâce au contrôle accru que vous avez sur eux. Une large sélection de différents modèles, chacun configuré pour cette tâche spéciale sur un disque NVMe, et vous pouvez déjà remplacer les modèles SOTA avec très peu de compromis.
Ils n'ont sorti que le 35B ? Je pensais que le 27b avait remporté le vote ? Pas intéressé par le 35b…
Je suis content que les gens avec des cartes plus petites puissent vivre ça maintenant, je pense qu'on y est depuis environ 6 mois, mais avec les modèles plus grands. On va bien se régaler à partir de maintenant !
Mec, je ne peux plus me permettre de faire ça 😭 au moins la gemma 26B s'adapte à mes 16 Go de vram, j'suis jaloux
J'exécute Qwen 3.6 localement sur un Mac. Modèle impressionnant, mais je veux vérifier ma propre expérience par rapport au battage médiatique. Mon installation : Sonnet produit des vidéos courtes pour moi en production. Le fichier de compétence, le flux de travail, la structure du dossier de projet, tout a été testé en conditions réelles. Il livre une sortie finie de manière cohérente. J'ai remis le même dossier à Qwen. Je lui ai dit de suivre la compétence et de continuer le flux de travail. Il lit tout, reconnaît les étapes, puis produit une sortie qui ne correspond pas au brief. La structure dérive, le ton dérive, des étapes de compétence sont sautées. Pas utilisable sans un lourd nettoyage manuel. Question sincère aux utilisateurs avancés de Qwen ici : est-ce que je fais quelque chose de mal ? Est-ce que vous le poussez différemment que vous ne le feriez avec Sonnet ? Structure de l'invite système différente, manière différente de référencer les fichiers de compétence, fenêtres de contexte plus petites, paramètres de sampler spécifiques ? Je suis heureux qu'on me dise que je le tiens mal. Parce que sur papier, ça devrait le faire. En pratique, sur ma machine, ce n'est pas encore ça.
Je me retrouve toujours dans des boucles de réflexion avec Qwen depuis la 3.5. Les paramètres sont les mêmes qu'avec Unsloth, mais ça continue de se répéter et je ne sais honnêtement pas comment le résoudre. Pendant ce temps, Gemma4 répond presque instantanément et gère bien les appels d'outils.
L'appel de fonction dans opencode n'a pas échoué une seule fois pour l'instant (gemma a eu du mal). Modifier des pages html m'a donné des résultats étonnamment décents. Cependant, je l'ai vu halluciner lorsqu'on lui a demandé de comparer les performances par rapport aux modèles gpt et claude. Q8_K_XL
