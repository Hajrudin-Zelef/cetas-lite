---
id: collect-261001-ia-llm/ia-llm/odysseus-l-ia-auto-habergae-de-pewdiepie-enfin-rapide-sur-mac-korben-1
title: "Odysseus - L'IA auto-hÃ©bergÃ©e de PewDiePie enfin rapide sur Mac"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Apple", "OpenAI", "OpenRouter"]
dates: []
keywords: ["agent", "agents", "gpu", "llama", "llama.cpp", "mai", "qwen", "research"]
source: docs/RAG/collect-261001-ia-llm/odysseus-l-ia-auto-habergae-de-pewdiepie-enfin-rapide-sur-mac-korben.md
source_anchor: ""
source_lines: [1, 81]
sha256: 1f3df5e418bd115519575a897b85bf1af75013d8685a0b6c13d6b808f9ada4f9
---

# Odysseus - L'IA auto-hÃ©bergÃ©e de PewDiePie enfin rapide sur Mac

## Ce quâil faut retenir RÃ©sumÃ© gÃ©nÃ©rÃ© par IA

1. Odysseus, le workspace IA auto-hÃ©bergÃ© de PewDiePie, regroupe chat, agents, recherche web, email, calendrier, notes, documents et gÃ©nÃ©ration d'images en local, avec un mode Compare pour tester deux modÃ¨les cÃ´te Ã cÃ´te en aveugle.
2. Sur Mac, Docker ne peut pas accÃ©der au GPU Metal et les modÃ¨les tournent sur CPU, donc il faut installer Odysseus en natif via le script start-macos.sh et connecter Ollama pour faire tourner des modÃ¨les comme Qwen3-30B-A3B sans ralentissements.
3. Le chat et l'Ã©diteur Markdown avec historique fonctionnent bien, mais le Cookbook, les agents, le calendrier et la recherche factuelle restent trop fragiles : la fonctionnalitÃ© Deep Research a inventÃ© une vidÃ©o, un numÃ©ro de marque et des statistiques sans fondement.

**Odysseus**, c'est le workspace IA que **PewDiePie** a balancÃ© sur GitHub fin mai. Chat, agents, recherche web, email, calendrier, notes et documents, tout est rÃ©uni dans un superbe cockpit auto-hÃ©bergÃ© dans lequel les modÃ¨les et l'historique peuvent rester sur VOTRE machine. Vous avez aussi un Ã©diteur Markdown avec historique de versions, une gÃ©nÃ©ration d'images intÃ©grÃ©e, et un mode Compare pour opposer deux modÃ¨les cÃ´te Ã  cÃ´te en aveugle (exclusion faite des services tiers liÃ©s Ã  la recherche web, l'email et le calendrier, Ã©videmment...).

J'ai voulu l'installer sur mon Mac, mais j'ai quand mÃªme rencontrÃ© un petit piÃ¨ge qui a failli transformer l'expÃ©rience en Ã©chec... Voici donc comment le faire tourner correctement sous Mac + quelques idÃ©es de ce que vous allez pouvoir faire avec ce truc.

## Avant de commencer

Odysseus Ã©volue vite. La branche `main` est stable et testÃ©e, mais la branche `dev` peut casser n'importe quand. PrivilÃ©giez la branche stable en la clonant explicitement :

```
git clone https://github.com/odysseus-dev/odysseus.git
cd odysseus
git checkout main
```
## Ãtape 1 : Oubliez Docker sur Mac

Sur Apple Silicon, Docker ne peut pas accÃ©der au GPU Metal. Les modÃ¨les tournent sur le CPU, et Ã§a rame comme pas permis. Du coup, sur Mac, on est obligÃ© d'installer Ã§a en natif.

## Ãtape 2 : Installer Odysseus en 2 commandes

Ouvrez un terminal et lancez Ã§a :

```
git clone https://github.com/odysseus-dev/odysseus.git
cd odysseus
./start-macos.sh
```
Le script installe alors les dÃ©pendances et dÃ©marre le serveur. Sur mon Mac, AirPlay squattait le port 7000, donc l'interface s'est ouverte automatiquement sur `http://127.0.0.1:7860`. Ensuite, au premier lancement, il faut renseigner un nouveau mot de passe pour le compte admin.

Si vous accÃ©dez Ã Odysseus depuis une autre machine (VPS ou serveur distant), ouvrez un tunnel SSH plutÃ´t que d'exposer le port directement sur Internet :

```
ssh -L 7860:127.0.0.1:7860 [email protected]
```
Ensuite, accÃ©dez simplement Ã  `http://127.0.0.1:7860` en local. C'est plus sÃ©curisÃ© que d'ouvrir le port sur le public.

## Ãtape 3 : Un modÃ¨le rapide, genre Qwen

Pour du rapidos sans que Ã§a devienne stupidos, j'ai choisi pour mes tests **Qwen3-30B-A3B**. Ce modÃ¨le n'active qu'une partie de ses paramÃ¨tres Ã  chaque rÃ©ponse, ce qui colle bien Ã  la RAM unifiÃ©e d'un gros Mac.

La premiÃ¨re fonctionnalitÃ© assez cool dans cet outil, c'est le **Cookbook** qui permet de tÃ©lÃ©charger des modÃ¨les adaptÃ©s Ã  notre config. Ce dernier a correctement reconnu mon M4 Max et ses 128 Go mais s'est ensuite plantÃ© comme une merde sur llama.cpp avec cette option `--flash-attn auto` invalide. AprÃ¨s avoir retirÃ© le paramÃ¨tre, le serveur restait aux abonnÃ©s absents tandis que l'interface l'affichait toujours comme actif. Bref, j'ai laissÃ© tombÃ© llama.cpp.

*Le Cookbook reconnaÃ®t bien le M4 Max et propose des modÃ¨les adaptÃ©s. Le lancement, lui, a plantÃ©.*

L'autre option qui s'est offerte Ã moi, Ã§a a Ã©tÃ© Ollama ( je vous ai montrÃ© comment l'installer par ici ). Vous rÃ©cupÃ©rez le modÃ¨le qui vous intÃ©resse avec et vous lancez le serveur ouvert sur le rÃ©seau local :

```
ollama pull qwen3:30b-a3b
OLLAMA_HOST=0.0.0.0:11434 ollama serve
```
Dans les rÃ©glages d'Odysseus, ajoutez ensuite `http://localhost:11434/v1`. Et hop, Ã§a fonctionne !! La version 1.0.2 propose aussi maintenant MLX dans le Cookbook, mais Ollama reste le chemin qui a vraiment tenu bon durant mes tests. Vous aurez peut-Ãªtre plus de chances que moi.

*L'ajout manuel de l'endpoint Ollama a Ã©tÃ© le chemin le plus fiable.*

**Bonus :** Si vous n'avez pas de GPU ou que vous prÃ©fÃ©rez la fiabilitÃ©, vous pouvez aussi connecter une API externe comme **OpenAI** ou **OpenRouter**. Dans les rÃ©glages, ajoutez simplement votre clÃ© API et Odysseus basculera sur ces modÃ¨les cloud. C'est moins gratuit que du local, mais Ã§a marche sans galÃ¨re.

## 3 cas d'usages

Alors, cet outil fait plein de choses et on peut l'utiliser vraiment pour ce qu'on veut, mais pour ma part, je vous ai isolÃ© trois cas d'usage compatibles avec mon activitÃ©.

Premier scÃ©nario, la tournÃ©e de veille du matin. Je lui donne les liens reÃ§us durant la nuit, il les classe par sujet, repÃ¨re les doublons et propose trois angles courts. Je garde les sources dans un document local, avec les objections et les points Ã vÃ©rifier. Et comme vous l'avez vu sur mes captures d'Ã©cran, le chat sait dÃ©jÃ faire du brainstorming puisqu'il m'a sorti trois angles exploitables en quelques secondes...

*Trois angles d'articles en 29 secondes, avec un brouillon versionnÃ© ouvert Ã  droite.*

DeuxiÃ¨me scÃ©nario, un dossier par article. J'y range le brief, les liens primaires, les citations, les captures et les versions du texte. L'Ã©diteur Markdown rÃ©cupÃ¨re le titre depuis le premier H1 et conserve l'historique. Avec deux modÃ¨les configurÃ©s, le mode Compare permet d'opposer deux intros ou deux plans cÃ´te Ã cÃ´te sans envoyer le brouillon chez un fournisseur.

TroisiÃ¨me scÃ©nario, le secrÃ©taire de rÃ©daction local. Il surveillerait une boÃ®te dÃ©diÃ©e aux alertes de sÃ©curitÃ©, prÃ©parerait la liste des sujets urgents et poserait des rappels dans le calendrier. Sur le papier, c'est exactement le type de corvÃ©es chiantes qu'un agent devrait pouvoir absorber mais en pratique, jamais de la vie, je ne lui laisse faire des actions automatiques comme envoyer des mails, ou gÃ©rer mon calendrier. Pour moi, c'est pas encore assez safe.

J'ai demandÃ© Ã  l'agent de retrouver un document et de le rÃ©sumer. Il a dÃ©pensÃ© prÃ¨s de 1 800 jetons Ã  deviner comment appeler `manage_documents`, sans jamais lancer l'outil de recherche. De son cÃ´tÃ©, le calendrier a terminÃ© sur un `Could not extract JSON` puis la recherche web a trouvÃ© le bon dÃ©pÃ´t GitHub, mais le modÃ¨le en a choisi un faux.

Le pompon, c'est la fonctionnalitÃ© Deep Research. AprÃ¨s 4 minutes de boulot, Ã§a m'a gÃ©nÃ©rÃ© 13 pages issues de 6 sources retenues. L'outil a donnÃ© la bonne conclusion au document, mais ensuite, a totalement inventÃ© une vidÃ©o, un numÃ©ro de marque et plusieurs statistiques. Donc pour l'utiliser dans un cadre journalistique, sans contre-vÃ©rification intÃ©grale, Ã§a ne sera pas possible.

*Le rapport affirme que le site est vide alors qu'il contient un guide complet. Jolie hallucination, prÃ©sentÃ©e comme une preuve.*

