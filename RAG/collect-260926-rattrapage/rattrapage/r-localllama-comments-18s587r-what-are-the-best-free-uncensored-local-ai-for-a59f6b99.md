---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-18s587r-what-are-the-best-free-uncensored-local-ai-for-a59f6b99
title: "r-localllama-comments-18s587r-what-are-the-best-free-uncensored-local-ai-for-a59f6b99"
domain: rattrapage
role: reference
task: reference
actors: ["Cohere", "Google", "Microsoft", "Mistral", "OpenAI"]
dates: []
keywords: ["chatgpt", "cohere", "gguf", "gptq", "gpu", "mistral", "research"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-18s587r-what-are-the-best-free-uncensored-local-ai-for-a59f6b99.md
source_anchor: ""
source_lines: [1, 52]
sha256: 153193c7e6ccf83ef12d4547e21d1c3b8d2aa856712c535f7834c96ee4301d2a
---

# r-localllama-comments-18s587r-what-are-the-best-free-uncensored-local-ai-for-a59f6b99

Merci pour ton avis !
Explique-nous pourquoi ce contenu n’est pas utile.
      Quels sont les meilleures IA gratuites locales non censurées pour la génération d'idées/écriture ? 
        
        
        
    
    
    Je trouve que la plupart des choses que je veux faire nécessitent une idée ou une histoire de quelque sorte, mais le truc, c'est que j'ai du mal avec les idées et la planification, et tout ça.
J'ai utilisé des IA comme chatgpt auparavant pour ce genre de choses, mais je remarque un thème commun : en termes d'écriture/idées, c'est plutôt clichéd, mais aussi super fade et banal. Donc, je suis à la recherche d'une IA d'histoire gratuite qui soit pas mal et ce serait bien si elle était de 7b ou autour de ça, locale et non censurée si possible.
Section des commentaires
Téléchargez LM Studio. Recherchez GGUF dans le programme et regardez les fiches des modèles. Pour tout ce qui est listé, RP écrira des histoires.
est-ce bizarre que j'utilise Kobold Lite parce que ça fonctionne mieux, et aussi tu dis que l'IA de jeu de rôle écrit de bonnes histoires. Comment ça marche ?
Les deux suivent une narration. Certains des LLM RP 7B les mieux notés répondent bien aux incitations d’histoire. Je donne des incitations par segment d’histoire, et l’IA semble mieux appréhender l’information que pour une incitation plus longue.
Quelques conseils gratuits, car il y a plusieurs cas d'utilisation liés au contexte de ta question.
Défini "meilleur" pour signifier bon à _____. Une fois que tu ajoutes plus d'une tâche spécifique, tu passes à la généralisation.
Défini "gratuit" pour signifier que tu n'as pas à payer pour utiliser ou que tu possèdes le résultat du modèle. Le type de licence devient important ici.
Défini "non censuré" pour signifier quel type de garde-fous assouplis tu recherches. Devrait-il utiliser un langage profane ou ne pas avoir de biais culturel, socio-économique ou politique ?
Cherche d'autres solutions d'écriture créative dont les gens ont discuté ici comme points de départ.
Une chose que tu découvriras probablement, c'est que planifier une histoire et écrire des segments de celle-ci sont des tâches différentes. Trouve un cadre d'écriture d'histoire à utiliser pour installer tes prompts. Les modèles de personnages et le mapping d'histoire sont deux autres exemples de tâches différentes. Bonne chance et n'oublie pas de partager ce que tu as fini par créer et comment ça s'est passé !
Tu peux essayer mon modèle sur collab, il écrit avec un style unique, rien de trop fancy, mais TRES différent de chatGPT:
https://colab.research.google.com/drive/1G_XXGrjhUirt0Ffws_ayzH8Q5E3hERIx?usp=drive_link
modèle complet pour les personnes avec un bon GPU :
SicariusSicariiStuff/Tenebra_30B_Alpha01_4BIT
SE CREA IMAGENES?
Cela fait 2 ans. putain.
Un de mes plotbots ?
https://huggingface.co/FPHam/Plot_BOT_V3_13b_GPTQ
https://huggingface.co/FPHam/PlotBot_13B-GPTQ-V2
https://huggingface.co/FPHam/Plot_BOT_13b_GPTQ
et pour parler d'écriture : https://huggingface.co/FPHam/Writing_Partner_Mistral_7B
Si seulement il y avait, je ne sais pas, un tas de ces posts littéralement chaque jour que quelqu'un pourrait parcourir.
Pour écrire, vérifie absolument Koboldcpp avec quelque chose comme Tiefighter. L'UI Lite intégrée de Koboldcpp a des modes d'écriture dédiés ainsi que des améliorations qui aident avec les tâches d'écriture. Tiefighter, en tant que modèle, a des données de roman supplémentaires ajoutées.
Mise à jour : J'ai vu dans les commentaires que tu es déjà familier avec KoboldAI Lite, dans ce cas tu sais que niveau interface, c'est très adapté à l'écriture et il te faut juste un bon modèle comme celui que j'ai suggéré ci-dessus. Les solutions alternatives recommandées par les gens sont toutes très orientées sur le chat et moins adaptées à l'écriture continue.
Pour ceux qui veulent l'essayer sur colab : https://koboldai.org/colabcpp avec les paramètres par défaut.
Tu peux essayer Infermatic.ai avec le modèle Noromaid, c'est pour l'écriture d'histoires et le jeu de rôle non censuré NSFW, tu peux aussi trouver Mixtral, Solar et Dolphin.
Salut, bonjour, c'est gratuit ?
D'accord.
Deux choses qui vont aider :
GPT4All a un client de chat local, qui prend des modèles que vous lui donnez au format GGUF, et il peut faire des 'localdocs'. Vous déposez vos documents de création de monde dans un dossier, vous les intégrez dans GPT4All, et il fera référence à ceux-ci lors de ses réponses.
Mythomax est un modèle assez bon, et il est non censuré.
GPT4all peut fonctionner avec votre RAM plutôt qu'avec votre VRAM, donc il sera beaucoup plus accessible pour des modèles légèrement plus grands, selon votre système. Ça ne sera peut-être pas ultra rapide, mais ça a son utilité.
Ensuite, vous devez configurer un bon prompt système (ce qui est donné au LLM avant la conversation, définissant en gros les termes), pour un assistant d'écriture. Revenez et je vous donnerai plus d'infos quand vous serez prêt, si vous le souhaitez.
Je suis un grand fan de Mistral-Medium, qui est gratuit sur poe.com :
Bien que ce ne soit pas complètement non censuré, la censure n'est pas du tout comparable à OpenAI, Google ou Microsoft :
https://poe.com/Mistral-Medium
J'utilise Faraday, ça rend super facile de télécharger différents modèles à tester. En créant un "personnage", fais de ce personnage un éditeur et de toi un écrivain. Demande-lui des listes d'idées pour quelque chose que tu essaies d'écrire.
Tous les modèles que j'ai essayés ont le problème dont tu parles. Ils génèrent soit des absurdités, soit les idées les plus clichés possibles. À mon avis, ils ne sont utiles pour écrire que si tu proposes toutes les idées toi-même et que tu les utilises pour rédiger du contenu en suivant une description détaillée de la scène que tu veux qu'il écrive. Dans Faraday, tu peux mettre en pause la génération de texte et réécrire des parties avant de relancer la génération pour rester sur la bonne voie.
je comprends totalement cette galère lol, je me retrouvais souvent coincé dans cette même boucle de réponses d'IA sans saveur... je joue avec plein de modèles locaux mais honnêtement, la plupart d'entre eux nécessitent beaucoup de travail sur les prompts ou hallucinent juste des trucs farfelus. pour info, j'ai été en train d'esquisser des idées brutes avec des complétions de chat puis je les envoie dans Walter Writes pour les humaniser un peu... ça fait beaucoup moins robotique, parfois même passable, aussi si tu es dans les trucs locaux, les gens semblent apprécier KoboldCPP ou OpenHermes sur ollama
Cohere AI est génial parce qu'il n'a pas de limites et presque pas de censure.
Chat | Cohere
Salut
