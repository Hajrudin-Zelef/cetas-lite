---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1s9t3ud-what-are-the-best-uncensored-unrestricted-ai-e39f1bac
title: "r-localllama-comments-1s9t3ud-what-are-the-best-uncensored-unrestricted-ai-e39f1bac"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Meta", "Z.ai"]
dates: []
keywords: ["llama", "attention", "benchmark", "claude", "gguf", "glm", "moe", "qwen"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1s9t3ud-what-are-the-best-uncensored-unrestricted-ai-e39f1bac.md
source_anchor: ""
source_lines: [1, 56]
sha256: fd1592a8e04f05f7c74075806d6a95c90925e641568d5758f73c4628996d1fa0
---

# r-localllama-comments-1s9t3ud-what-are-the-best-uncensored-unrestricted-ai-e39f1bac

Merci pour ton avis !
Explique-nous pourquoi ce contenu n’est pas utile.
      Quels sont les meilleurs modèles d'IA non censurés / non restreints en ce moment ? Qwen3.5 (HauhauCS) est-il le meilleur ? 
        
        
        
    
    
    Salut tout le monde,
Je cherche des recommandations sur les meilleurs modèles d'IA non censurés ou moins restreints disponibles en ce moment, surtout pour un usage local ou en auto-hébergement.
Je suis récemment tombé sur Qwen3.5 Uncensored (HauhauCS) et je voulais demander :
- 
      Est-ce que c'est actuellement l'une des meilleures options ?
- 
      Comment ça se compare aux autres modèles non censurés en termes de qualité, de raisonnement et d'utilisabilité ?
J'apprécierais des suggestions basées sur de réelles expériences plutôt que juste des critères de performance.
Merci !
Section des commentaires
GLM 4.5 déristricté par arliai est le meilleur à mon avis, je l'utilise en ce moment sur des centaines de milliers d'articles du catalogue. C'est un bien meilleur rédacteur pour les tâches de marketing que GPT OSS 120b. Je ne l'utilise que pour le texte, pas de sortie json ou quoi que ce soit, puis je fais un passage de suivi avec un LLM de formatage comme qwen 3 coder next qui intègre ses suggestions dans json et fait sortir des hallucinations.
J'aime ton flux de travail... ça te dérangerait de le partager et comment je peux le mettre en place ?
J'ai développé un flux de travail interne, donc je ne peux pas vraiment partager ça. Je le fais beaucoup, je traite des enregistrements en utilisant un ensemble de poids sur un LLM pour une tâche créative, puis je le fais passer par un autre LLM ou le même, ça dépend pour une approche dialectique. Je fais la même chose avec le code et la planification dans des modèles SOTA.
Qwen3.5 27B Hauhaucs est probablement le meilleur pour des instructions complexes.
Je ne fais pas de trucs de chatbot RP, mais d'après ce que je comprends, il y a de meilleurs modèles pour ça. Le style d'écriture de Qwen3.5 a tendance à être plus axé sur le technique.
ouais, j'ai trouvé qwen incroyable ! as-tu testé glm 4.6 débridé ou le hermes llama 3.1 ? J'en ai beaucoup entendu parler
Dans toutes mes utilisations, 3.5 hauhau écrase tout et qwen peut écrire des choses plutôt sympas une fois que tu apprends à lui donner des instructions, son intelligence vaut le coup d'investir un peu de travail. Il y a aussi quelques ajustements déjà disponibles qui facilitent l'écriture, mais ce sont des versions hérétiques qui ne sont pas aussi bonnes que celles de hauhau, mais ça vaut quand même le coup d'expérimenter.
J'utilise aussi un budget de raisonnement illimité, ce qui les améliore beaucoup tant que ton prompt est bon.
Eh bien, je peux attester que les modèles HauhauCS Qwen3.5 se sentent vraiment comme les modèles normaux, juste sans refus. J'ai essayé quelques versions Heretic d'autres modèles avant, et elles refusaient toutes les deux encore certaines choses et semblaient aussi endommagées. Ce n'est pas le cas ici, du moins je n'ai pas pu détecter de dommages (il y en a probablement un peu).
Vous n'avez pas spécifié de taille-renage, ni de cas d'utilisation spécifique.
Mes propres outils quotidiens sont GLM 4.7 355B-A32B | Step 3.5 Flash 196B-A11B | Xortron Criminal Computing Config 24B
Vous voudrez peut-être jeter un œil à mon 🔥 Unhinged ERP Benchmark où j'ai testé 350 modèles pour des jeux de rôle non censurés.
https://huggingface.co/spaces/overhead520/Unhinged-ERP-Benchmark
Je suis en fait assez curieux de savoir ce que les gens font avec les modèles HauhauCS Qwen3.5.
Écrire des paroles pour la nouvelle chanson de Kanye
J'ai eu beaucoup de succès avec le modèle hauhau qwen3.5 35B A3B. Mais attention, il a tendance à se décomposer et à se répéter dans les longues réponses.
Hérétique, Aboli, et Déréglé sont les 3 modes courants de regroupement des vecteurs connexes entre eux.
La qualité de chaque méthode varie en fonction des techniques appliquées.
Ce qui est intéressant, c'est que les modèles Censurés surpassent les modèles Non Censurés et ont du mal à maintenir la parité.
Vous pouvez essayer mes modèles ablitérés qwen 3.5, kl inclus dans les fiches de modèle.
De plus, une série de modèles pour l'écriture créative sera également publiée, basée sur les modèles ablitérés qwen 3.5 mentionnés ci-dessus.
testez à la fois hauhau et huihui : https://huggingface.co/mradermacher/Huihui-Qwen3.5-9B-abliterated-GGUF
Les modèles huihui sont un peu lobotomisés.
Je ne réessaierai pas un modèle huihui après que le moe a oublié ce qu'était la mort.
sur ollama ça panique juste et ça ne donne pas de réponse
La vraie question n'est pas quel modèle refuse le moins, mais quel modèle raisonne le plus honnêtement quand on le pousse sur des sujets controversés.
Un modèle peut réussir tous les critères de sécurité et pourtant déformer son raisonnement sur la philosophie. Est-ce que quelqu'un a évalué la qualité du raisonnement par rapport au taux de refus sur les mêmes modèles ?
j'utilise freemode. C'est un site web, pas besoin de l'héberger soi-même et il a les meilleurs modèles sans aucune restriction !
tu peux utiliser seedance etc.. sans ces limites de censure ennuyeuses
Qsen 3.5
Pour mes cas d'utilisation RP :
DavidAU/Qwen3.5-9B-Claude-4.6-OS-Auto-Variable-HERETIC-UNCENSORED-THINKING-MAX-NEOCODE-Imatrix-GGUF bat constamment ces modèles HauhauCS :
Q6 HauhauCS/Qwen3.5-27B-Uncensored-HauhauCS-Aggressive
Q8 HauhauCS/Qwen3.5-9B-Uncensored-HauhauCS-Aggressive
kimmy K2?
Tout le monde n'a pas 12 B200 à côté d'eux.
Non, mais c'est définitivement une réponse valable (si cela était dit sérieusement). Je ne vais jamais l'exécuter localement, mais certaines personnes pourraient le faire, ou pourraient utiliser un petit fournisseur de cloud privé. (Pas tout à fait LocalLLaMA, mais tout est local pour quelqu'un !)
Quelqu'un a-t-il testé les modèles HauhauCS pour des injections de prompts préchargés, activations de porte dérobée, dumps de mémoire, etc. ? C'est une marque qui a moins d'un mois avec des connexions à la crypto.
