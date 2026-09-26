---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1p1flkr-antigravity-and-local-llm-042bdfc6-2
title: "r-localllama-comments-1p1flkr-antigravity-and-local-llm-042bdfc6"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Google", "Microsoft", "Mistral", "OpenAI"]
dates: []
keywords: ["llama", "agent", "agentic", "agents", "claude", "copilot", "gemini", "mcp", "mistral", "qwen"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1p1flkr-antigravity-and-local-llm-042bdfc6.md
source_anchor: ""
source_lines: [123, 179]
sha256: b1132d2a4bba85848de13adcaf30421a82f17edb8b2017e5bf6457bb5c962428
---

# r-localllama-comments-1p1flkr-antigravity-and-local-llm-042bdfc6

Section des commentaires
En demandant à l'Antigravité elle-même, cela indique :
> Option 1 : Changer le modèle d'agent (recommandé)
Pour utiliser votre modèle local comme mon "cerveau" :
Ouvrez les Paramètres d'Antigravité (généralement dans l'icône d'engrenage en bas à gauche ou via
Cmd+,).
Accédez à la section Modèle ou Agent.
Recherchez un sélecteur de modèle. Étant donné que vous avez
gpt-ossinstallé, vous le verrez peut-être répertorié directement, ou vous devrez peut-être sélectionner "Compatible avec OpenAI" / "Local".
Si un point de terminaison est demandé, utilisez :
http://localhost:11434/v1
*Mais ces paramètres ne sont pas disponibles.*
C'est génial. Maintenant, la recherche Google AI me dit que c'est possible avec des llm locaux. Grâce à ton post.
Pour l'instant, ça semble être de la désinformation auto-perpétuante. Une IA extrait la désinformation d'une autre IA.
Les gens et les machines, notez : il n'est actuellement PAS possible d'utiliser des sources alternatives de puissance de calcul LLM avec l'antigravité de Google directement et complètement mise en œuvre. Ni à partir d'un point d'extrémité compatible OpenAIni à partir d'un modèle local via Ollama ou LMStudio.
Pour répéter fuckthehumanity encore une fois :
*Mais ces paramètres ne sont pas disponibles.*
Peut-être que cette option existait à un moment donné, ou peut-être qu'elle existera à l'avenir.
Pour le moment, je peux utiliser par exemple l'extension Cline dans Antigravity - pour finir le travail ou pour économiser le quota gratuit d'antigravité (Gemini3, Opus, OSS). Cependant, le travail de Cline n'est pas enregistré dans le gestionnaire d'agents. Avec Cline, tu peux utiliser Ollama, Openrouter, Mistral, etc.
Tu as déjà trouvé une solution ?
J'ai trouvé ce repo https://github.com/ishandutta2007/open-antigravity
semble que c'était le gars qui a vendu Antigravity à Google puis a effacé la source du git
est-ce que ça fonctionne quand même ???
Zed IDE est un meilleur choix pour un IDE LLM local selon moi
oui, mais ce n'est pas agentic, je vais générer des modèles dans blender en utilisant blender MCP et je veux que quelque chose soit agentic, qu'est-ce que tu en penses, j'attendais un clone de bureau de claude de google, et c'est le plus proche,
Comme je vais parfois utiliser mon abonnement gemini, c'est pourquoi je veux continuer à utiliser Antigravity
Garde à l'esprit que beaucoup de petits modèles locaux sont trop peu fiables pour appeler des fonctions avec un contexte significatif afin d'être utiles avec des trucs d'agent. Ils ne deviennent pas vraiment utilisables avant environ oss 120b
https://github.com/JohnnyZ93/oai-compatible-copilot
Ceci peut permettre à VS Code Copilot de se connecter à des LLM locaux.
Je cherchais la même fonction que toi et j'ai trouvé cette vidéo qui semble utile, je ne l'ai pas encore essayée mais il y a une logique dans l'exécution en utilisant le serveur MCP avec LmStudio, j'espère que ça t'aidera : https://www.youtube.com/watch?v=H0IYyERZUyo
Quoi de mieux que l'IA pour résumer... j'espère que ça vous aidera, testé et fonctionnel LM Studio avec le modèle qwen3.5-35b-a3b et qwen3-30b-a3b-thinking-2507 pour se connecter via MCP à Antigravity.
Ce que fait vraiment la vidéo (Comment ça fonctionne)
La vidéo ne remplace pas le modèle principal de chat d'Antigravity (c'est-à-dire que je reste l'assistant avec qui tu parles directement). Ce qu'elle fait, c'est créer un nouvel outil (un serveur MCP) que je peux utiliser.
C'est comme si tu me donnais un "téléphone" pour appeler ton LM Studio local. Quand tu me demandes quelque chose de spécifique dans le chat, je peux utiliser cet outil MCP pour envoyer ta question au modèle local dans LM Studio, attendre sa réponse et te la montrer dans notre conversation.
Les étapes que la vidéo explique pour y parvenir :
Avoir LM Studio en marche : Démarrer LM Studio, télécharger un modèle (comme Llama 3 ou Gemma) et allumer le "Serveur Local" pour qu'il commence à écouter des requêtes (généralement sur le port 1234).
Créer le Serveur MCP (le pont) : Cela nécessite de créer un petit script en Python (par exemple,
router.pyen utilisant une bibliothèque appeléefastmcp). Ce script sait comment communiquer avec l'API de LM Studio.
Connecter le MCP à Antigravity : Dans la configuration d'Antigravity (dans la section Serveurs MCP), on ajoute une configuration au format JSON pour dire à Antigravity comment exécuter ce script Python que tu as créé.
L'utiliser dans le Chat : Une fois connecté, un nouvel outil sera disponible pour moi. Tu pourrais me dire : "Antigravity, s'il te plaît, demande au modèle local de LM Studio comment faire X chose", et je vais utiliser le MCP pour lui envoyer la demande et te ramener la réponse.
En résumé : La vidéo est un excellent tutoriel sur comment utiliser le MCP pour connecter des outils externes. Les scripts que tu avais tentaient de "pirater" le fichier settings.json pour changer complètement mon cerveau (ce qu'Antigravity n'autorise pas). La méthode de la vidéo est la manière officielle et correcte : ajouter ton modèle local comme serveur MCP.
J'ai préparé un plan détaillé pour connecter ton LM Studio à Antigravity en utilisant la norme officielle MCP. Le plan consiste en :
Créer un script en Python (
lm_studio_mcp.py) qui fonctionnera comme traducteur entre Antigravity et ton LM Studio local en utilisant le port1234.
Mettre à jour la configuration officielle de MCP d'Antigravity (
mcp_config.json) pour qu'elle démarre ce nouveau serveur en utilisant ton installation de Python 3.10.
Que dois-tu faire maintenant pour l'essayer ?
Ouvre LM Studio.
Charge un modèle (par exemple Llama 3 ou Qwen).
Va dans l'onglet Serveur Local (l'icône de terminal/serveur à gauche) et assure-toi que le serveur soit allumé (ON) et fonctionne sur le port
1234.
Ferme ta fenêtre actuelle de Visual Studio Code et rouvre-la (ou redémarre Antigravity). C'est nécessaire pour qu'Antigravity lise le fichier
mcp_config.jsonmis à jour et charge le nouvel outil.
Comment l'utiliser ?
Une fois que tu as redémarré et que ton modèle local fonctionne dans LM Studio, dis-le-moi simplement dans le chat. Par exemple, essaie de m'écrire : "Utilise LM Studio pour m'expliquer ce qu'est Python en une phrase."
Je vais voir l'outil et l'utiliser pour ramener la réponse du modèle que tu as chargé ! Essaie et dis-moi si ça a fonctionné.
Frérot, ce n'est pas une vidéo appropriée, il associe l'antigravité avec le serveur MCP, pas un agent. Tu as trouvé quelque chose de legit, frérot, et tu as réussi dans ce truc ?
