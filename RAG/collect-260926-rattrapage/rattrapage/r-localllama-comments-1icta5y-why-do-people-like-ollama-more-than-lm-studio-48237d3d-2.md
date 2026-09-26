---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1icta5y-why-do-people-like-ollama-more-than-lm-studio-48237d3d-2
title: "r-localllama-comments-1icta5y-why-do-people-like-ollama-more-than-lm-studio-48237d3d"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft", "OpenAI"]
dates: []
keywords: ["llama", "gguf", "llama.cpp", "open source", "vllm"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1icta5y-why-do-people-like-ollama-more-than-lm-studio-48237d3d.md
source_anchor: ""
source_lines: [9, 38]
sha256: 133fc13af746c47f200f061aea8339535a4fbb95293676d664bfc07187aa94fc
---

# r-localllama-comments-1icta5y-why-do-people-like-ollama-more-than-lm-studio-48237d3d

    Je suis juste curieux. Je vois plein de gens discuter d'Ollama, mais en tant qu'utilisateur de LM Studio, je ne vois pas beaucoup de gens en parler.
Mais LM Studio me semble tellement mieux. [ÉDITÉ] Il a une interface graphique vraiment sympa, pas des commandes obscures en mode headless. Si je veux essayer un nouveau modèle, c'est super facile de le chercher, de le télécharger, de l'essayer, de le jeter ou de le proposer à AnythingLLM pour du RAG ou du foldering.
(Avant que tu ne mentionnes KoboldCPP, oui, absolument KoboldCPP, il ne tourne juste pas sur ma machine.)
Alors pourquoi cette obsession pour Ollama sur ce forum ? Aide-moi à comprendre.
[ÉDITÉ] - J'avais à l'origine mal compris que Ollama nécessite son propre format de fichier de modèle par rapport à l'utilisation des GGUF. Je ne comprenais pas que tu pouvais récupérer des modèles qui n'étaient pas dans l'index d'Ollama, mais les gens dans ce fil ont corrigé l'erreur. Malgré tout, ce fil est un débat très utile sur le sujet de 'full app' contre 'mostly headless API.'
Section des commentaires
Que dire de SillyTavern, Llama-cpp, LlamaFiles, Oogabooga... ?
SillyTavern est un frontend, pas un backend. Il n'exécute pas de modèles. LM Studio et Ollama peuvent lui servir.
Llama.cpp est le back-backend pour tous ces outils. LM Studio est llama.cpp, Ollama est llama.cpp, etc. Je ne l'utilise juste pas parce qu'il faut être un vrai codeur, mais je comprends totalement pourquoi les vrais codeurs le font.
Je ne sais pas ce qu'est LlamaFiles.
Oogabooga et KoboldCPP ne fonctionnent pas sur ma machine pour des raisons que je ne comprends pas, mais peu importe, fais comme tu veux. J'aime bien ceux-là. Je suis surtout intéressé par Ollama parce que je pense que son étrange et limité éventail de modèles dans un format propriétaire, c'est vraiment un désavantage.
KOBOLDCPP GANG LEVEZ-VOUS !
Qu'est-ce qui empêche KoboldCpp de fonctionner sur ta machine ? Si LMStudio fonctionne, KoboldCpp devrait sûrement fonctionner. On peut probablement t'aider à le faire marcher.
Lmstudio est optimisé pour faire fonctionner des modèles localement sur votre ordinateur portable/ de bureau.
Ollama + Openwebui vous permet d'avoir votre propre service d'IA auquel vous pouvez accéder à distance et de le partager avec vos amis/ famille.
Ça a du sens. Je n'ai pas d'amis, donc je n'y ai pas pensé.
LM Studio a également une option de serveur, donc vous pouvez l'utiliser pour servir des points de terminaison dans un format compatible avec l'API OpenAI.
J'ai commencé à utiliser ollama beaucoup plus récemment, et ce que j'ai découvert, c'est que même si l'interface utilisateur est un terminal minimaliste, c'est en fait un excellent serveur de point d'API. Toute interface web avec le support d'Ollama te permet de changer de modèles à la volée, contrairement au rechargement complet à la main comme dans KoboldCPP. L'allocation des couches est automatique, ce qui élimine la difficulté de régler divers paramètres avant de l'utiliser. Et la plus importante de toutes, tout modèle que tu n'utilises pas pendant plus de 5 minutes se décharge de la VRAM, économisant énormément d'électricité et de puissance de calcul. Contrairement à koboldcpp, cela signifie que tu peux toujours la garder en marche en permanence, avec seulement une très légère utilisation de RAM, contrairement à d'autres back-ends où c'est toujours complètement allumé ou complètement éteint.
En ce qui concerne les fichiers propriétaires, ollama est en fait juste un wrapper de llama.cpp, et il utilise des fichiers .GGUF. Le seul problème, c'est que c'est très ennuyeux de configurer la longueur du contexte et d'autres trucs depuis le terminal. C'est pourquoi ils utilisent un fichier modèle, que tu importes ensuite, mais honnêtement, c'est assez ennuyeux aussi. En passant, LM Studio utilise aussi une structure de fichiers bizarre même si c'est juste un wrapper de llama.cpp.
En termes de design d'interface utilisateur, Open webUI est bien au-dessus de toute autre interface web en matière de fonctionnalités et de performances, même si elle pourrait définitivement bénéficier de choix UX plus rationnels. Elle est spécifiquement conçue pour ollama, et tu peux même obtenir une image docker fournie avec Ollama, bien qu'elle puisse également être utilisée avec tout autre point d'API. Elle résout également le problème de configuration des fichiers modèles. Je te recommande absolument de l'essayer, car c'est probablement le meilleur pour les cas d'utilisation au travail.
La plupart des raisons pour lesquelles les gens n'utilisent pas LM Studio, c'est qu'ils ont tendance à être soucieux de leur vie privée, et il y a beaucoup de personnes très techniques ici. Lorsque nous nous soucions de notre vie privée, nous ne voulons pas utiliser une application à code source fermé comme LM Studio, car nous n'avons aucune idée de ce qu'elle peut faire avec nos données. Nous plaidons tous fermement pour des modèles open source, donc ce serait aussi hypocrite d'utiliser un logiciel d'inférence à code source fermé. KoboldCPP est la version la plus conviviale de la version barebones de llama.cpp. C'est un exe à un clic, et ce n'est pas encombré de manière quelconque. Son interface utilisateur est nulle, mais la plupart des gens l'utilisent avec SillyTavern ou une autre interface. Je suppose que tu as un Mac et que tu ne veux pas rester là à compiler des binaires toi-même, donc l'option suivante la plus facile était LM Studio. Cela dit, je te suggérerais fortement d'essayer open webUI avec Ollama, ça a une petite courbe d'apprentissage, mais je pense que tu seras agréablement surpris une fois que tu auras appris à l'utiliser.
Pour moi, c'est l'utilisation directe depuis la ligne de commande. Ils ont ajouté la commande à LMS, mais d'après ce que je comprends, tu dois quand même lancer l'interface graphique au moins une fois. J'ai de meilleures performances sur certains matériels/drivers par rapport à LM studio. J'aime LMS et j'utilise différentes choses comme lms, ollama, kobold, vllm, etc. Sur quelle machine te trouves-tu ?
Je suis sur un Microsoft Surface Laptop 7 avec Snapdragon. J'ai utilisé LM Studio, Ollama, AnythingLLM. J'ai utilisé le mode serveur de LM Studio pour servir à SillyTavern. KoboldCPP ne fonctionne pas.
Je comprends ce que tu dis. Je ne suis pas une personne de ligne de commande. Genre, je fais servir LM Studio à AnythingLLM, mais je veux voir ce qui se passe.
À propos des outils GUI, en plus de ComfyUI et Automatic1111, qu'est-ce que les gens utilisent pour la génération d'images ?
Je peux seulement parler pour moi-même, mais je suis aveugle, LLM Studio est difficile à naviguer avec un lecteur d'écran, et OLlama est facile à installer en plus d'avoir une jolie API locale.
LM Studio n'est pas open source, donc cela ne m'intéresse pas.
(Résumé, sans IA : Je découvre que mon cas d'utilisation de tout faire sur mon ordinateur portable n'est en fait pas courant ici, et qu'Ollama fait un bon serveur pour les personnes qui essaient de servir leurs modèles à une variété d'appareils autour de leur maison/entreprise.)
La vérité, c'est que les deux ont leurs forces et leurs faiblesses :
Ollama est une option backend simple mais efficace. Cependant, je n'aime pas ses conventions de nommage (vois comment les gens sont confus à propos des modèles R1 distillés et quels quantis ils utilisent) et je ne pense pas que l'approche semblable à Docker ait vraiment été nécessaire ici. Ça donne parfois une impression de décalage et contribue à cette sensation d'opacité que tu sembles décrire. Bonne expérience quand c'est associé à OpenWebUI.
