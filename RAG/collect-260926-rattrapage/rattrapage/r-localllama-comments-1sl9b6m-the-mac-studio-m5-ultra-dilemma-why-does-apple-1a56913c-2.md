---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1sl9b6m-the-mac-studio-m5-ultra-dilemma-why-does-apple-1a56913c-2
title: "r-localllama-comments-1sl9b6m-the-mac-studio-m5-ultra-dilemma-why-does-apple-1a56913c"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "MiniMax", "Nvidia", "SGLang", "vLLM"]
dates: []
keywords: ["llama", "agent", "benchmarks", "blackwell", "claude", "inference", "llama.cpp", "nvidia", "qwen", "sglang", "vllm"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1sl9b6m-the-mac-studio-m5-ultra-dilemma-why-does-apple-1a56913c.md
source_anchor: ""
source_lines: [12, 56]
sha256: d425103738a7b6015a7ca6696569484df3bc94aeffd75f8e4c58b5e42e76b03f
---

# r-localllama-comments-1sl9b6m-the-mac-studio-m5-ultra-dilemma-why-does-apple-1a56913c

Maintenant, je regarde le futur Mac Studio. D'après les prix précédents, le modèle de base M5 Ultra sera probablement autour de 4 600 $. Mais voici le hic : l'Ultra de base est équipé de 96 Go. C’est la définition de "inutile mais cher". 96 Go c'est une sentence de mort pour tout ce qui dépasse 70B si vous voulez vraiment travailler pendant que le modèle fonctionne.
Si je passe à 256 Go, Apple va probablement me taxer encore 2 000 $ de plus. Ça semble être un énorme surplus, mais comme il n'y a pas de niveau 128 Go ou 192 Go pour l'Ultra, je suis coincé entre l'enclume et le marteau. C'est frustrant parce qu'un Ultra de base devrait être le meilleur compromis, mais la gestion de la mémoire par Apple fait que le Max haute gamme a l'air mieux que l'Ultra d’entrée de gamme, ce qui est juste bizarre.
Quelques questions pour les légendes ici :
- 
      Des fuites "fais-moi confiance, mec" sur les véritables niveaux de mémoire pour le M5 Ultra ? Y a-t-il un espoir pour un niveau intermédiaire de 128 Go ou 128 Go+ ?
- 
      Alternatives matérielles locales ? J'ai regardé Nvidia, mais c'est un vrai casse-tête. Les P40 et V100 sont de l'histoire ancienne. Même une configuration 3090/4090 nécessite 3 cartes pour rivaliser avec la VRAM du Mac, et à ce stade, le coût est pratiquement le même que celui du Mac, mais avec le "bonus" supplémentaire d'une facture d'électricité énorme et d'une pièce qui ressemble à un sauna.
- 
      Je suis dans l'écosystème Mac depuis plus de 15 ans - c'est une dépendance à ce stade. Comment puis-je obtenir des "tokens infinis" (ou au moins une expérience utilisable de 70B+) sans vendre un rein pour 256 Go de mémoire unifiée ?
Section des commentaires
Mais voilà le truc
On peut toujours dire quand les gens utilisent un LLM pour réfléchir à leur place lorsque le résultat est un modèle dense de 70 milliards.
et honnêtement ?
Les blagues mises à part, j'ai remarqué que j'ai commencé à utiliser ce langage aussi. Le pire ? Je pense que les LLM influencent mon style d'écriture.
Ce n'est pas juste de la mimique, c'est de l'évolution.
Mon PTSD déclenché, merci
Essayez oMLX. La persistance du cache KV sur disque fait une énorme différence. La plupart des tours se terminent avec en gros pas de temps de pré-remplissage.
Je devrais rédiger un petit guide…
Édition : Petit guide rédigé ! https://www.reddit.com/r/oMLX/comments/1slcfit/local_inference_on_apple_silicon_with_subsecond/
Est-ce que ce n'est pas un peu ce que font les checkpoints et les slots dans lcpp, ou est-ce que c'est quelque chose de différent ? Si c'est différent, y a-t-il quelque chose de similaire qui utilise lcpp ?
Un peu ? llama.cpp pré-alloue un nombre fixe de "slots" au démarrage. Chacun est un contexte d'inférence isolé avec son propre cache KV. Le slot 0 et le slot 1 peuvent exécuter exactement la même invite système et chacun maintiendra sa propre copie. Le nombre de slots est fixe au démarrage, la mémoire évolue avec le nombre de slots × la longueur du contexte, et la persistance se fait via une sauvegarde/restauration explicite des checkpoints. Vous déterminez manuellement si vous souhaitez sérialiser l'état KV d'un slot sur le disque et le recharger plus tard. Pensez à chaque slot comme à une ligne téléphonique dédiée avec sa propre mémoire et un bouton "pause/reprendre".
oMLX adopte l'approche vLLM/SGLang adaptée pour la mémoire unifiée d'Apple Silicon : un arbre préfixe partagé de blocs de cache KV où les requêtes entrantes correspondent automatiquement à des préfixes déjà calculés et ne paient que pour le suffixe nouveau. Cinq requêtes partageant une invite système le calculent une fois ; la mémoire évolue avec un contenu unique, pas avec le parallélisme. Les pages du cache entrent et sortent du NVMe selon les besoins, survivent aux redémarrages et gèrent automatiquement l'éviction par niveaux chaud/froid. En gros, il n'y a pas de capacité à faire des checkpoints manuellement.
En résumé : les slots de llama.cpp sont explicites, isolés et potentiellement persistants ; le cache de préfixe d'oMLX est automatique, partagé et sauvegardé sur disque. Pour un fan-out multi-agent avec des invites partagées, oMLX est plus efficace en mémoire. Pour "je veux une session longue durée que je peux geler et dégeler à la main", les slots de llama.cpp sont plus simples.
Edit : De mon point de vue, llama.cpp me donne plus, comme un outil pour construire un harnais autour du comportement explicite que je veux (recherche/chargement automatique des slots). oMLX est plus "piles incluses" et essaie en partie d'être un harnais pour le faire à votre place sans que vous ayez à configurer quoi que ce soit de spécial.
Faites fonctionner le modèle sur le studio 128 Go et travaillez à partir d'un MacBook 64 Go = profit (pour Apple).
Je pense que choisir 128 Go comme point idéal est un peu arbitraire.
Même pour votre flux de travail actuel, qwen-coder-next 80b est bien, mais il est surpassé par le plus petit qwen 3.5 27b dans la plupart des benchmarks. Donc, une perspective pourrait être que 96 Go est parfaitement acceptable, le décrire comme "inutile mais cher" est un peu absurde.
D'un autre côté, même si vous obtenez une machine de 128 Go, il est inévitable qu'il y aura des moments où vous regretterez de ne pas avoir 256 Go. Par exemple, pour MiniMax-M2.7 (un grand bond par rapport à qwen-coder-next, beaucoup plus proche de quelque chose comme Claude Sonnet 3.6), il peut *techniquement* tenir dans 108 Go à la quantification Q4, mais vous n'aurez pas beaucoup de marge pour une longueur de contexte plus longue ou d'autres charges des applications système. Et qui sait, peut-être que le prochain modèle ouvert de pointe sera dans la fourchette de 400b paramètres.
Je ne m'inquiéterais pas trop à ce sujet, le domaine change si rapidement qu'il n'y a pas moyen de prédire quelle quantité de mémoire est la plus optimale. Prenez juste autant de mémoire que vous pouvez vous permettre, et priorisez la bande passante mémoire (c'est-à-dire Ultra > Max) afin que vous puissiez avoir des vitesses de génération rapides peu importe le modèle que vous choisissez.
Cela dépend. Si tu aimes Sonnet/Opus et que tu utilises Claude Code, Minimax 2.7 est terrible pour le travail de codage comparé à qwen-coder-next.
La formation de Minimax est pleine de :
appels de sonde
lectures redondantes
boucles de lecture → analyse → re-lecture
traces synthétiques lourdes en JSON
schémas d'outils hallucinés
appels “glob” (
read_file("*"))
boucles de réessai
“réflexion” auto-référentielle entre les appels
Le fait que tout soit globé signifie que ça se perd et peut potentiellement faire des choses dangereuses - surtout dans le claude-code. Une fois que tu commences à écrire un harness d'agent pour gérer cela, tu réalises que tu as un problème infini d'un modèle à un trillion de paramètres à gérer.
Il ignore les invites système, ignore les invites de claude-code, il émet des appels d'outil quand on lui dit de ne pas le faire, il réessaye les mêmes appels plein de fois, il continue de sonder l'environnement et de globé et globé et globé et globé et d'halluciner des outils manquants... tu traces ses outils et tu vois tous les noms d'outils décoys et tu réalises qu'ils ont entraîné leurs paramètres d'outils sur une tonne de déchets distillés sans autre direction que ILS veulent que minimax soit l'agent - pas ton IDE/Claude Code/Cursor/OpenCode.
Pour le codage, qwen3-coder-next est meilleur que Qwen3.5 27b principalement parce que 27b n'est pas configuré pour les appels d'outils agentiques. 27b est génial pour discuter de code, mais pas nécessairement pour prospérer dans Claude Code / Open Code
Blackwell a6000x2 fait ressentir les Macs comme bon marché
On peut essayer quelque chose comme le Strix Halo, ça pourrait t'y amener....
