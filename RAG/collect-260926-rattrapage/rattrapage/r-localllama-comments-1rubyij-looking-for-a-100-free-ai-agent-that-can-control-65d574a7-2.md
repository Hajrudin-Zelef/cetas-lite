---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1rubyij-looking-for-a-100-free-ai-agent-that-can-control-65d574a7-2
title: "r-localllama-comments-1rubyij-looking-for-a-100-free-ai-agent-that-can-control-65d574a7"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Google", "Microsoft", "Moonshot", "OpenRouter"]
dates: []
keywords: ["agent", "agents", "claude", "kimi", "mcp", "model context protocol", "open source", "qwen"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1rubyij-looking-for-a-100-free-ai-agent-that-can-control-65d574a7.md
source_anchor: ""
source_lines: [9, 54]
sha256: e1bf98718f8a119ee0d1b0f4872bf533f2faca99913cd66ae01df8849bcea496
---

# r-localllama-comments-1rubyij-looking-for-a-100-free-ai-agent-that-can-control-65d574a7

    Salut tout le monde.
J'essaie de trouver un agent IA complètement gratuit qui peut contrôler un navigateur et effectuer des tâches sur des sites web.
Exemples : • ouvrir des sites web • rechercher sur Google • cliquer sur des boutons • remplir des formulaires • naviguer sur des pages • automatiser les tâches normales du navigateur
Quelque chose de similaire à des outils comme Claude Computer Use ou d'autres agents IA pour navigateurs.
Je recherche quelque chose de totalement gratuit, de préférence open source ou pouvant fonctionner localement.
Quelqu'un connaît des bons outils ou projets pour ça ?
Merci.
Section des commentaires
BrowserOS est ce que tu cherches… c'est open source et gratuit. Il suffit d'avoir stepfun 3.5 d'Open Router, de le connecter et ça devrait le faire. Ils offrent aussi Kimi 2.5 gratuitement avec certaines limites. Pas besoin de payer quoi que ce soit.
https://www.browseros.com/
J'en ai entendu parler, mais je ne suis pas sûr, en fait
L'utilisation de Browser est probablement la plus proche de ce que vous voulez. C'est open source, fonctionne avec des modèles locaux via Ollama, et gère le clic, le remplissage de formulaires, la navigation par défaut. Vous le pointez sur une tâche en anglais simple et il comprend les interactions du navigateur.
Si vous voulez quelque chose de plus léger, Stagehand (déjà mentionné) ou Playwright MCP valent le coup d'œil. Playwright MCP connecte n'importe quel LLM à un navigateur via le Model Context Protocol, donc vous pouvez le coupler avec le modèle local que vous utilisez.
Pour du local totalement, le principal goulot d'étranglement est le modèle de vision. Les agents de navigation doivent comprendre ce qui est affiché à l'écran, et les petits modèles ont du mal avec ça. Qwen 2.5 VL ou un modèle Gemma récent gèrent ça pas mal pour des tâches simples, mais les flux complexes à plusieurs étapes rencontrent des difficultés avec tout ce qui a moins de 30 milliards de paramètres d'après mon expérience.
l'utilisation du navigateur est immédiatement dirigée vers l'abonnement/le paiement lors de la première invite...
Tu devrais jeter un œil à Dassi. C'est une extension Chrome qui agit comme un assistant d'automatisation de navigateur axé sur la confidentialité.
Ce qui est cool pour ton utilisation, c'est que c'est 100 % local pour le contrôle du navigateur (fonctionne nativement via l'extension Chrome) et permet d'intégrer tes propres modèles via BYOK (Apporte Ta Propre Clé) ou de l'utiliser entièrement gratuit avec des modèles locaux en le configurant avec des points de terminaison compatibles OpenRouter/Ollama.
Caractéristiques qui se démarquent en fonction de ce que tu demandes :
Nécessite aucun code pour l'utiliser - tu discutes juste avec
Automatisation en mode double - Utilise à la fois l'Arbre d'Accessibilité (mode ref) et les coordonnées en mode visuel
Sécurité avant tout - Nécessite ta permission explicite pour des choses comme partager des infos, acheter des choses ou télécharger des fichiers.
Partage de contexte - si tu as plusieurs onglets dans un groupe, il se souvient de la conversation à travers eux.
C'est une extension, pas une application de bureau séparée, donc ça vit complètement dans ton navigateur.
Merci, je vais jeter un œil !
Skales fait exactement ça. Agent de navigateur intégré qui navigue sur les sites web, clique sur des boutons, remplit des formulaires, extrait du contenu et contourne les bannières de cookies. S'exécute localement avec Ollama, totalement gratuit.
Pas de Docker, pas de configuration de terminal. Télécharge, installe, connecte Ollama et dis-lui "va sur [site web] et [fais ça]".
https://github.com/skalesapp/skales
Windows, macOS, Linux. Plus de 650 étoiles.
Quel LLM utilise-t-il ? Faut-il payer pour quoi que ce soit
Je pense que CopilotKiwi est ce que tu cherches. Un agent de navigation AI gratuit pour Chrome.
ça ne fonctionne pas, tu cliques et ça ne fait rien
Travailler sur un navigateur de bureau Mac qui expose des points de terminaison pour qu'un LLM puisse "voir / contrôler tout", mais il est aussi destiné à un usage humain. L'objectif est d'atteindre une parité totale avec Firefox. Vidéos de démonstration de certaines choses amusantes qu'il peut faire :
https://www.youtube.com/@wkdomains
Tout est open source :
https://github.com/wkdomains/macos-app
Je suis sur Windows 11 pourtant et je me demandais s'il y a un projet de porter wkdomains sur Windows, ou si vous considéreriez le faire ? Je suis développeur travaillant avec des agents de codage et ce serait utile pour mon flux de travail. Si vous envisagez une version Windows et avez besoin d'aide, je serais heureux de contribuer.
J'aime bien utiliser le navigateur et les oies. Prends le plus gros modèle qwen3.5 que tu peux faire tourner et profite-en. Mais à moins que ce ne soit le modèle 122B, je crains que tu n'en tires pas grand-chose.
Tu peux même essayer le MCP playwright avec le même modèle, selon la complexité des tâches que tu t'attends à accomplir.
Si tu es à l'aise avec les API (peut-être que tu as un crédit openrouter ?), essaie GLM5. Rien de mieux que le source ouvert, à mon avis.
Si tu veux quelque chose de totalement local et gratuit, regarde l'utilisation d'un navigateur avec un modèle local. Ça fonctionne avec Playwright en arrière-plan. Le hic, c'est que les modèles capables de vision sont lourds, donc tu auras besoin d'au moins 16 Go de VRAM pour quelque chose de fiable. Une alternative plus légère : utilise l'arbre d'accessibilité au lieu de captures d'écran, que les modèles plus petits gèrent bien.
Qwen 3.5 (le plus grand que vous pouvez exécuter) + pinchtab
Utilise le serveur Playwright MCP avec un modèle local. Tu dois lancer Chrome avec WebSockets activés pour le Protocole Chrome DevTools (CDP), ajouter le serveur MCP avec l'URL WebSocket (unique à chaque lancement), et ensuite tu pourras utiliser un LLM pour contrôler un navigateur.
Alors, Selenium ?
Vous pourriez probablement laisser une IA écrire un script Selenium pour faire ça.
Le régisseur fait ça, je l'ai essayé. Code source ouvert. Avis mitigés. https://opensourcedisc.substack.com/p/opensourcediscovery-98-browserbase
Quelques points. 1. Vous allez avoir besoin de quelque chose comme un serveur MCP basé sur Serp API ou utilisez Crawl4AI pour les recherches Google. Les bots sont en grande partie interdits de réaliser des recherches autrement. 2. J'ai testé plusieurs modèles pour cela, Minimax 2.5 est le premier modèle à bien fonctionner de manière cohérente et à ne pas être bloqué sur différentes parties des sites web et fonctionne bien avec le MCP Playwright. Donc, vous n'avez pas besoin d'un modèle de vision pour naviguer sur le web, mais bien sûr, si la tâche implique de récupérer des éléments d'images, vous aurez évidemment besoin d'un modèle de vision. J'ai un serveur MCP personnalisé pour le crawling des résultats de recherche avec SerpAPI (je pense que Crawl4AI devrait aussi être une bonne solution pour les recherches). Maintenant, en fait, naviguer sur les pages avec Playwright consomme une tonne de jetons, donc vous voulez avoir le modèle dans la VRAM. Une vue de page web renvoie très souvent entre 20k et 200k jetons qui doivent être traités, donc si ce n'est pas dans la VRAM, je pense que vous allez avoir besoin d'un long horizon temporel pour accomplir quoi que ce soit. La conclusion... c'est à chacun de décider si un modèle nécessitant 128 Go de VRAM est gratuit à utiliser mais fonctionne bien et est rapide quand vous l'avez :D.
