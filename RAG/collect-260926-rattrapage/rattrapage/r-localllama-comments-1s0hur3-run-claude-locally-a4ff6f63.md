---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1s0hur3-run-claude-locally-a4ff6f63
title: "r-localllama-comments-1s0hur3-run-claude-locally-a4ff6f63"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple"]
dates: []
keywords: ["claude", "llama", "agent", "agents", "qwen", "reasoning", "sol"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1s0hur3-run-claude-locally-a4ff6f63.md
source_anchor: ""
source_lines: [1, 35]
sha256: d2465b889a2d616d895aaf806d04f7a85e9f3527a2d439a340369ed8b19f3d23
---

# r-localllama-comments-1s0hur3-run-claude-locally-a4ff6f63

Merci pour ton avis !
Explique-nous pourquoi ce contenu n’est pas utile.
      Exécuter Claude localement ? 
        
        
        
    
    
    Ma question peut paraître un peu bête, désolé.
Je sais que Sonnet et Opus sont des LLM, mais je ne comprends toujours pas vraiment ce qu'est Claude Code et j'essaie de le découvrir. Au début, je pensais que c'était un peu comme ClawdBot, qui permet au modèle d'IA de fonctionner en dehors de la simple fenêtre de chat ?
Encore une fois, il est probablement évident que je n'y connais rien ;)
Bref, ma question est la suivante : est-il possible d'exécuter l'un ou l'autre de ces modèles, voire tous, en local ? J'ai entendu dire que Claude est bien meilleur que les autres modèles, notamment pour la programmation, et j'espérais en apprendre davantage à ce sujet.
Merci d'avance !
Section des commentaires
le plus proche que vous pourriez être de « faire tourner claude » est d'utiliser un opus distill comme : https://huggingface.co/Jackrong/Qwen3.5-9B-Claude-4.6-Opus-Reasoning-Distilled-v2 https://huggingface.co/Jackrong/Qwen3.5-27B-Claude-4.6-Opus-Reasoning-Distilled-v2
mais le modèle réel sur le site web n'est pas disponible en téléchargement
Parfait, merci ! Dans ce cas, je pourrais prendre un autre modèle ! Merci pour ta réponse en tout cas
Sonnet et Opus sont les modèles, Claude Code est la fonctionnalité qui "fait des trucs" toute seule en fonction des suggestions du LLM. En gros, quand tu discutes juste avec le LLM et qu'il dit "voilà le fichier HTML, mets-le dans ce dossier, fais ça dans ta config nginx et démarre le serveur", etc., il fait ces choses tout seul. Tu ne crées pas de fichiers, tu ne copies-colles pas le contenu ni ne démarres les serveurs, etc., Claude Code s'en charge pour toi. Ou si tu veux faire tourner quelque chose dans un conteneur Docker, il crée le fichier Docker, construit le conteneur et le démarre. C'est comme si tu utilisais les mêmes commandes, mais tu n'as pas besoin de le faire manuellement. Tu peux aussi utiliser Claude Code avec un modèle local.
Réponse de Claude lui-même :
D'accord, merci beaucoup d'avoir demandé, Claude !
Donc, pour résumer - Les LLM ne peuvent pas être exécutés parce qu'ils sont lourds (mais de toute façon, parce qu'ils sont en code source fermé), mais je peux faire tourner quelque chose comme Claude Code, qui permet aux IA d'agir comme des agents en leur donnant accès à mon système de fichiers ou en exécutant des commandes.
Mon résumé est-il correct ?
J'ai vu que Claude était particulièrement bon, mais je ne sais pas vraiment si j'ai entendu quelque chose de spécial ou beaucoup sur Claude Code. Ma question est : si je finis par utiliser Qwen2.5-Coder, devrais-je utiliser Claude Code pour lui permettre d'agir comme un agent ou devrais-je utiliser autre chose de mieux, compte tenu du fait que je n'utilise pas de LLMs de Claude ?
Commentaire supprimé par le membre
Merci ! J'ai une 9070 XT, un Ryzen 5 7500F et 32 Go de RAM DDR5.
Si tu trouves un moyen de voler des poids claude et que tu possèdes le matériel, alors tu as claude en local.
Ce dépôt github créé par ce gars https://github.com/nicedreamzapp
https://github.com/nicedreamzapp/claude-code-local.git
Beaucoup de choses ont changé depuis que ce dépôt a été une nuit de "est-ce que je peux faire fonctionner Claude Code sur Ollama." C'est maintenant un ensemble complet d'IA locale : un serveur Anthropic natif MLX d'environ 1000 lignes, réutilisation du cache de prompts, parsing natif des appels d'outils Gemma / Llama / Qwen, mode code (élimine automatiquement le prompt de harnais de 10K tokens de Claude Code pour les modèles locaux), l'agent de navigateur, le mode narration, un pipeline iMessage, redémarrage du lanceur conscient du modèle, et — la pièce que je pense être la plus importante — une boucle vocale entièrement sans mains sur l'appareil (Apple SFSpeechRecognizer + TTS à voix clonée) qui vit dans le projet frère NarrateClaude. Bien au-delà de ce que couvre le tableau "Le Voyage" ci-dessus.
As-tu signé pour la liste blanche de Dario ?
Non, désolé, qu'est-ce que c'est ?
Commentaire supprimé par un membre de l’équipe de modération
D'accord, parfait, merci ! Je peux télécharger Claude Code depuis ollama ainsi qu'OpenClaw et Codex depuis openai
Ils semblent tous faire la même chose, non ?
Lequel devrais-je prendre selon toi ?
