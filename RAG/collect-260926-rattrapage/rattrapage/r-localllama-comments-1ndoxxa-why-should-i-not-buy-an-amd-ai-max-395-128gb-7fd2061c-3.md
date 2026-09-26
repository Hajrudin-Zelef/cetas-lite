---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1ndoxxa-why-should-i-not-buy-an-amd-ai-max-395-128gb-7fd2061c-3
title: "r-localllama-comments-1ndoxxa-why-should-i-not-buy-an-amd-ai-max-395-128gb-7fd2061c"
domain: rattrapage
role: reference
task: reference
actors: ["Apple", "Intel", "Nvidia"]
dates: []
keywords: ["llama", "benchmarks", "gpu", "intel", "llama.cpp", "lpddr5x", "nvidia", "vllm"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1ndoxxa-why-should-i-not-buy-an-amd-ai-max-395-128gb-7fd2061c.md
source_anchor: ""
source_lines: [67, 73]
sha256: 96c2843e088ceb7f0d722b28a46c2231e3c58d1549fc015815976e0ab89baa22
---

# r-localllama-comments-1ndoxxa-why-should-i-not-buy-an-amd-ai-max-395-128gb-7fd2061c

Le Strix Halo a évidemment une mémoire plus rapide qu'un PC classique avec de la RAM système et une carte graphique - mais voici quelques benchmarks d'exemple pour ktransformers sur du matériel grand public (Core i9-14900KF + DDR5-4000 MT/s en double canal) + RTX 4090 : https://github.com/kvcache-ai/ktransformers/blob/main/doc/en/AMX.md
Le matériel Apple est plutôt bon pour un seul utilisateur à la fois, mais le traitement des invites n'est pas super rapide et la concurrence (plusieurs utilisateurs ou plusieurs tâches en même temps) n'est pas facile. Voici quelques benchmarks de llama.cpp pour différents chips de la série m : https://github.com/ggml-org/llama.cpp/discussions/4167 et ce PR a quelques notes sur la vitesse de llama.cpp en mode haute capacité avec llama-batched-bench : https://github.com/ggml-org/llama.cpp/pull/14363
Donc j'adorerais essayer le Strix Halo pour quelque chose comme ktransformers (connecter un GPU via le lien PCIe x4) ou pour faire tourner vllm afin d'obtenir une haute concurrence (puisque vllm sur Apple est uniquement CPU, ne supporte pas les quants, et est très lent, environ 8 tok/sec sur llama-2-7b-fp16 bs=1). J'ai trouvé ces benchmarks sur vllm sur le Framework Desktop : https://github.com/lhl/strix-halo-testing/tree/main/vllm (92 tok/s sur le Mac le plus puissant pour la taille de lot 1 q4 comparé à environ 357 tok/s pour la taille de lot 16 avec strix halo vllm q4 - et il semble que ce soit 149 tok/s sur llama-batched-bench pour M3 Ultra sur llama-2-7b-chat-q4km à bs=16 (ou -npl 16) pour un contexte 4k).
9/10 je recommanderais fortement. Je trouve tellement d'utilisations aléatoires pour l'IA bon marché que je ne savais pas que j'avais avant
C'est aussi ma QUESTION ! Je viens d’attraper le virus LocalLaLM. En ce moment, j'utilise mon ordinateur portable ROG avec un i9-13900H de 13e génération + Nvidia 4070. C'est suffisant pour me montrer les possibilités sans gratter la démangeaison. Alors, est-ce que je dépense 2 000 $ sur l'un de ces Strix Halo, un peu moins sur un bureau avec un 5070ti (je ne peux pas me permettre d'aller plus haut). Bien sûr, j'ai même vu quelques ordinateurs portables dans ce prix-là avec des processeurs Intel. Mais, sérieusement, je suis sur le point de me décider pour le GMKtec EVO-X2 AI Mini PC Ryzen Al Max+ 395 Mini Gaming Computers, 128 Go LPDDR5X 8000 MHz (16 Go*8) 2 To SSD PCIe 4.0. Est-ce que je fais une bêtise ?
Ce post spécifique a été supprimé en utilisant Redact. La motivation pourrait être liée à la vie privée, à la sécurité, à l'opsec, ou tout simplement un choix personnel d'enlever du vieux contenu.
friendly sugar fearless doll pen attempt yam obtainable liquid rustic
