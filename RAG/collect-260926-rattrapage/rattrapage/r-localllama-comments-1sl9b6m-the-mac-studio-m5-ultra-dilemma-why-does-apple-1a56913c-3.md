---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1sl9b6m-the-mac-studio-m5-ultra-dilemma-why-does-apple-1a56913c-3
title: "r-localllama-comments-1sl9b6m-the-mac-studio-m5-ultra-dilemma-why-does-apple-1a56913c"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Meta", "Nvidia"]
dates: []
keywords: ["turboquant"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1sl9b6m-the-mac-studio-m5-ultra-dilemma-why-does-apple-1a56913c.md
source_anchor: ""
source_lines: [57, 81]
sha256: 81d099bc6459188e8d626537cc28719857f6abced3dd458a73b3038098338e61
---

# r-localllama-comments-1sl9b6m-the-mac-studio-m5-ultra-dilemma-why-does-apple-1a56913c

C'est presque comme si l'IA n'était pas vraiment fonctionnelle étant donné les contraintes actuelles en matière d'ordinateurs et d'électricité, choquant, je sais.
Turboquant pourrait aider, qui sait.
3a. Mac a toujours été à prix premium, pour être honnête, c'est l'option la plus intéressante pour ton argent maintenant, ce qui est fou, pour dire le moins.... c'est un problème qui te concerne
3b. Tu peux essayer de trouver quelqu'un pour souder la mémoire pour toi ou utiliser une autre machine pour l'IA tournant sous Linux avec les Blackwells ci-dessus ou le Strix Halo et utiliser un MacBook Air ou un équivalent comme ta machine de travail..
Si tu veux vraiment de bonnes réponses et de l'aide pour ça, tu devrais utiliser les forums Level1Techs, ils sont beaucoup mieux adaptés aux trucs pro-consommateur LLM que cette communauté, à mon avis..
Parce qu'ils ne le construisent pas pour un LLM local
Je pense que 256 et 384 sont des points idéaux étant donné la prévalence de bons modèles récemment dans la gamme de taille ~370B. Et 512 pourrait même ne pas être proposé sur m5 ultra pendant un certain temps à cause de la rampocalypse.
Le préremplissage est censé avoir été beaucoup amélioré par l'architecture m5. Je ne sais pas quoi te dire.
Tout ce qui est en dessous de 512 Go n'est pas rentable. Même 512 Go c'est un peu faible. S'ils proposaient 768 Go, je prendrais comme Reksio prend du jambon.
Mec, tout mon disque dur principal local est un SSD de 128 Go
Ce post n'a pas bien vieilli... les prix sont pires.
Il se passe TELLEMENT de choses en termes d'efficacité des modèles locaux. Tu devrais juste essayer une autre variante du modèle.
voilà un modèle mlx 4 bits qui pourrait fonctionner
https://huggingface.co/mlx-community/Qwen3-Next-80B-A3B-Instruct-4bit
regarde un peu autour. voici quelqu'un qui fait tourner un modèle 80b sur un m1 à 35tps
https://www.reddit.com/r/LocalLLaMA/comments/1ni2chb/qwen3next_80b_mlx_mac_runs_on_latest_lm_studio/
La vitesse de traitement des prompts du M5 Max est très correcte. J'ai tout ce qu'il faut sous la main que tu peux exécuter localement, et le M5 Max est probablement l'un des meilleurs après le RTX 5090 et 6000 pour le traitement des prompts. Pourquoi penses-tu qu'Ultra va être tellement mieux ?
Je pense que ton problème, c'est que tu ne mets pas en cache les prompts correctement. Parce que je peux tourner sur mes Sparks QWEN3.5 122b et 397 et ils font le traitement des prompts à peut-être environ 1500, je n'ai pas les chiffres en tête, mais ça semble assez rapide parce que le caching est adéquat. Donc, je te conseille juste de bien gérer le caching.
Si tu as besoin de beaucoup de RAM, alors le Mac Pro avait 1 To de RAM je crois xd
tu peux acheter le Mac Studio de 512 Go ou même attendre la version de 1 To
utilise-le pour un setup de bureau et comme serveur quand tu veux utiliser ton MBP
« Développeur heavy utilisant l'IA, qui vit presque dans son IDE » correspond étonnamment mal au contenu technique du post. Sans scénarios concrets sur ce que l'on souhaite réellement faire avec la configuration, quels sont les fenêtres de contexte nécessaires et à quoi ressemble le flux de travail réel, les affirmations concernant les besoins en mémoire semblent assez générales. Avant de présenter 128 Go comme un minimum à peine suffisant, tu devrais peut-être revoir les bases du contexte, du cache KV et de la quantification.
Ce n'est pas appelé le Mac LLM. C'est difficile à imaginer mais les gens l'utilisent pour des choses autres que l'IA locale.
Vous pouvez toujours construire une machine threadripper à 20 000 $ qui a également 96 Go.
Toujours, ouais, tu peux toujours faire ça.
