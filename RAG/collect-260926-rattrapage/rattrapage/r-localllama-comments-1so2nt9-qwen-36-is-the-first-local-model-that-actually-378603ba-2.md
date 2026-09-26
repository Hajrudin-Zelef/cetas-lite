---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1so2nt9-qwen-36-is-the-first-local-model-that-actually-378603ba-2
title: "r-localllama-comments-1so2nt9-qwen-36-is-the-first-local-model-that-actually-378603ba"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Unsloth", "Z.ai"]
dates: []
keywords: ["qwen", "agent", "arr", "benchmarks", "datacenter", "gguf", "glm", "gpu"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1so2nt9-qwen-36-is-the-first-local-model-that-actually-378603ba.md
source_anchor: ""
source_lines: [9, 50]
sha256: 59ba5050663c5e9c3fcc3410168b2fc1669043275e3161ab8eb75405a5259563
---

# r-localllama-comments-1so2nt9-qwen-36-is-the-first-local-model-that-actually-378603ba

    J'ai passé du temps hier après le travail à essayer le nouveau modèle qwen3.6-35b-a3b, et au moins pour moi, c'est la première fois que j'ai vraiment l'impression qu'un modèle local n'était pas plus pénible à utiliser qu'il n'en valait la peine.
J'utilise des LLM dans mes projets personnels/jetables depuis quelques mois, pour le genre de code que je n'ai pas vraiment envie d'écrire (la plupart des XML UI dans Avalonia, C++ pour les systèmes embarqués), et j'avais Sonet et Opus gratuitement grâce au programme étudiant de Github, mais ils ont arrêté ça. J'ai essayé des modèles locaux pendant un certain temps aussi, mais j'avais surtout l'impression jusqu'à présent qu'ils étaient soit trop idiots pour faire le travail, soit qu'ils pouvaient le terminer mais que je passerais tellement de temps à corriger/ajuster/formater/réécrire le code que j'aurais autant pu le faire moi-même.
Qwen3.6 semble enfin avoir changé cela, du moins sur mon système et mes projets. Avec un 5090 + 4090, je peux charger le modèle Q8 avec un contexte complet de 260k, obtenir environ 170 tokens par seconde en fait l'un des modèles les plus rapides que j'ai essayés. Et contrairement à tous les autres modèles que j'ai essayés récemment, y compris Gemma 4, il peut réellement accomplir des tâches et nécessite uniquement des conseils ou des corrections mineurs à la fin. 9 fois sur 10, il suffit de lui demander de revoir ses propres changements une fois qu'il est "terminé" pour qu'il capte et corrige tout ce qui n'allait pas.
Je suis assez impressionné et c'est vraiment cool de voir les modèles locaux enfin commencer à atteindre ce niveau. Cela me donne de l'espoir pour un avenir où cette technologie n'est pas limitée à des centres de données massifs et des services d'abonnement, mais plutôt optimisée au point où même les ordinateurs de milieu de gamme peuvent en tirer parti.
Section des commentaires
chaque sortie, mêmes publications
Je suppose que c'est ce qui se passe quand chaque nouvelle sortie est meilleure que la dernière...
Ouais, mais si cette chose est meilleure que la dense Gemma 4 31B, comme le suggèrent les benchmarks que j'ai vus, ça déchire. La Gemma 4 est le premier modèle à franchir ce seuil pour moi, donc faire ça mais beaucoup plus vite, ça semble être un rêve devenu réalité.
Édition : Je viens d'essayer le plus petit Q4 GGUF d'Unsloth, et putain, ça tient vraiment la route par rapport à l'engouement. Utilisation parfaite d'agent, t/s vraiment rapide, et des performances de codage qui semblent meilleures que celles de la Gemma 4 dense.
Oui, alors les APIs payantes s'améliorent beaucoup. Ensuite, les modèles ouverts s'améliorent beaucoup.
Ensuite, les APIs payantes s'améliorent beaucoup. Ensuite, les modèles ouverts s'améliorent beaucoup.
Ensuite, les APIs payantes s'améliorent beaucoup. Ensuite, les modèles ouverts s'améliorent beaucoup.
[NaN lignes répétées cachées] /s
Ouais ? Ça papote et ça tourne en rond avec toi aussi ?
ça fait pour moi. 80k jetons jusqu'à présent puis j'ai abandonné.
N'a pas fait jusqu'à présent. J'ai utilisé LM Studio et OpenCode avec des réglages quasiment par défaut.
Pour moi, Qwen 3.5 27b est bien meilleur pour exécuter des tâches et résoudre des problèmes.
Si vous avez suffisamment de RAM et un 5090 + 4090, pourquoi ne pas faire tourner le GLM 4.7 358B A32B à IQ4XS ou IQ3XXS ?
La différence entre GLM 4.7 358B A32B et Qwen 3.6 35b A3B sera incroyablement grande.
Je considère Qwen 3.6 35b A3b et Gemma 4 26b a4b comme des modèles vraiment légers, proches de 9-12b de densité.
J'ai essayé GLM 4.7 (j'oublie quel quant mais probablement Q3) et j'ai rencontré beaucoup de problèmes avec l'arrêt soudain en plein milieu de la tâche, ou en ignorant de larges parties de la tâche.
Édition : c'était Q2, ce qui est probablement pourquoi ça a autant galéré.
C'est plutôt bien jusqu'à présent. Cependant, cela peut *vraiment* se retrouver coincé dans une boucle quand il réfléchit. Au point de remplir tout le contexte et de ne pas répondre. Essayez ces prompts à la suite :
Qu'est-ce qui est marron et collant ?
Très bien. Quelles sont d'autres blagues simples basées sur des jeux de mots comme celle-ci ?
Y en a-t-il qui sont un peu moins enfantines, et plus osées/adultes ?
Cela envoie qwen dans une spirale, itérant sans fin sur les mêmes trois ou quatre blagues nulles et décidant qu'elles ne sont pas drôles et pas adultes. Il se reconnaît même qu'il est dans une boucle plusieurs fois, mais il échoue à en sortir.
Donc, si la version 3.6 est aussi bonne et peut essentiellement remplacer les énormes chatbots de datacenter pour discuter normalement, comme pour le grand public. Cela ne semblerait-il pas raisonnable que les versions 3.7 ou 3.8 soient si performantes que nous puissions arrêter de construire des datacenters ?!? Comme si ceux-ci pouvaient fonctionner sur nos ordinateurs, quel est l'intérêt de continuer à construire ces énormes datacenters qui consomment toute la RAM et les SSD ?
Je ne veux pas justifier le bâtiment de datacenter actuel et le nettoyage du matériel, mais 3.6 n'est toujours pas au niveau des modèles frontier. Et même si je le considère « assez bon » pour la plupart des tâches, il vous faut toujours un matériel très coûteux pour le faire fonctionner à une vitesse raisonnable. Je suis vraiment content de ça, mais j'ai aussi 6 000 $ de GPU dans mon ordinateur donc... voilà.
Regarder Hermes-Agent travailler avec une quantité illimitée de jetons à >100tk/s avec ce modèle est un peu effrayant...
Des trucs comme chatjimmy à 16k jetons par seconde font vraiment peur aussi. Avoir ça localement serait fou.
J'ai seulement lu les posts et c'est probablement l'un des plus divisifs que j'ai suivis après la sortie. Les gens l'adorent ou le détestent.
Je ne l'ai pas encore essayé. Mais le truc le plus étrange dans le récit pour moi, c'est qu'une mise à jour de 0.1 d'un modèle qui n'a pas été publié 'il n'y a pas si longtemps' provoque une telle réaction passionnée parmi les utilisateurs. Je comprendrais si des problèmes de sur-analyse avaient été ajustés, mais d'après ce que j'ai vu dans les discussions, ce ne semble pas être le cas. Ce qui m'intrigue le plus, c'est de savoir si les gens qui adorent ou détestent vraiment utilisé 3.5 assez pour comparer les deux.
Est-ce que quelqu'un fait tourner ça sur un système avec 4 Go de VRAM et 32 Go de RAM ? Juste une question pour un ami (vous n'avez pas besoin de me rappeler que je suis pauvre).
Ça tourne sur mon portable 3050 4 Go. Si tu veux de la vision sur le GPU, tu auras très peu de place pour le contexte.
Sans vision, (soit pas chargé ou sur CPU), tu peux faire tourner
Q4_K_M quant
131k de contexte à la mémoire cache Q8 KV
taille de lot 512
et ça va à peine tenir. Si tu veux un peu de place et de la stabilité, tu peux opter pour moins de contexte. ~25tk/s
Avec mon PC qui a plus de 8 ans, équipé d'une 2080 Ti (11 Go de VRAM) et de 64 Go de RAM système, je peux obtenir 29 t/s avec Q6_K_XL et le contexte complet. C'est assez impressionnant, étant donné la complexité des tâches techniques qu'il peut gérer.
Ils se complètent bien avec Gemma, car Gemma a un avantage en écriture créative, ce qui en fait un meilleur interlocuteur général. C'est bien pour le brainstorming ou simplement pour réfléchir.
