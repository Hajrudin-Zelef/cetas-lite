---
id: collect-261001-ia-llm/ia-llm/ollama-vs-code-ia-locale-en-12-etapes-2026-5
title: "macOS et Linux"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Apple", "Microsoft", "Nvidia", "OpenAI"]
dates: []
keywords: ["agent", "amd", "apache", "arr", "copilot", "embeddings", "gpu", "nvidia", "open source"]
source: docs/RAG/collect-261001-ia-llm/ollama-vs-code-ia-locale-en-12-etapes-2026.md
source_anchor: ""
source_lines: [256, 311]
sha256: f8760c7988664365a03e828d5075a20bd2a31e97eb43df2243e18bb215adfa6b
---

# macOS et Linux

| Symptôme | Cause probable | Solution | 
|---|---|---|
| “ollama: command not found” | Le binaire n’est pas dans le PATH après installation | Redémarrez le terminal, ou réinstallez avec le script officiel | 
| “connection refused” sur le port 11434 | Le service Ollama n’est pas démarré | Lancez `ollama serve` dans un terminal dédié | 
| Le téléchargement du modèle s’arrête ou expire | Connexion instable sur un fichier de plusieurs Go | Relancez `ollama pull` , qui reprend le téléchargement interrompu | 
| Continue n’affiche aucun modèle disponible | Fichier config.yaml mal formé ou Ollama non démarré | Vérifiez l’indentation YAML et testez l’API avec curl avant de rouvrir VS Code | 
| L’autocomplétion est très lente | Modèle trop volumineux ou calcul sur CPU | Passez à un modèle plus petit ou vérifiez la détection GPU avec `ollama ps` | 
| Erreur de mémoire ou crash au chargement | VRAM insuffisante pour le modèle choisi | Utilisez une quantification plus agressive ou un modèle plus petit | 
| Le GPU n’est pas détecté | Pilotes graphiques obsolètes ou absents | Mettez à jour les pilotes Nvidia/AMD/Apple avant de relancer Ollama | 
| Erreur de syntaxe dans config.yaml | Indentation incorrecte (les espaces comptent en YAML) | Repartez de l’exemple de l’étape 6 et validez avec un linter YAML | 
| @codebase ne trouve rien dans le dépôt | Indexation par embeddings pas encore terminée ou modèle d’embeddings manquant | Vérifiez que nomic-embed-text est bien téléchargé et attendez la fin de l’indexation | 
| VS Code ne détecte pas l’extension après installation | Redémarrage de l’éditeur non effectué | Rechargez la fenêtre via la palette de commandes (“Reload Window”) | 

## Conseils avancés pour aller plus loin

Une fois l’environnement de base stable, plusieurs réglages permettent d’en tirer davantage. Séparer les rôles entre deux modèles différents, un modèle léger et rapide pour l’autocomplétion et un modèle plus large pour le chat, comme configuré à l’étape 6, donne le meilleur équilibre entre réactivité et qualité des réponses : c’est la logique déjà retenue par la configuration de ce tutoriel.

Continue permet aussi de créer des commandes personnalisées, appelées “slash commands”, pour des actions répétitives comme générer des tests unitaires ou rédiger un message de commit à partir du diff en cours. Ces commandes se définissent dans le même fichier `config.yaml`, sous une section `prompts` dédiée. Pour les équipes qui veulent partager une configuration commune, ce fichier peut être versionné dans un dépôt central et distribué à chaque poste, en laissant le champ `apiBase` configurable par variable d’environnement si certains développeurs préfèrent un serveur Ollama partagé sur le réseau local plutôt qu’une instance par machine.

Enfin, pensez à réévaluer le choix du modèle tous les deux ou trois mois. Le rythme de sortie des modèles ouverts spécialisés dans le code reste soutenu, et un modèle jugé secondaire aujourd’hui peut devenir le meilleur choix pour votre profil matériel après une nouvelle quantification ou une nouvelle version, sans qu’aucune autre partie de cette configuration n’ait besoin de changer.

Si plusieurs développeurs de la même équipe travaillent sur des machines aux capacités différentes, évitez d’imposer un seul modèle pour tout le monde. Un plus gros modèle sur la machine d’un développeur senior ne sert à rien si trois autres membres de l’équipe tournent sur des portables sans GPU dédié : dans ce cas, laissez chaque poste choisir son propre modèle d’autocomplétion selon le tableau de l’étape 3, et réservez un modèle commun, plus large, hébergé sur un serveur Ollama partagé sur le réseau local, pour les tâches de chat et d’agent qui demandent davantage de puissance. Cette approche hybride évite d’acheter du matériel neuf pour toute l’équipe alors qu’un seul serveur suffit à mutualiser l’accès aux modèles les plus gourmands.

## Foire aux questions

### Ollama est-il gratuit ?

Oui, Ollama est un logiciel open source et gratuit. Le seul coût réel est celui du matériel nécessaire pour faire tourner les modèles dans de bonnes conditions, en particulier la VRAM si vous visez des modèles plus grands que 7 milliards de paramètres.

### Quelle est la différence entre Ollama et LM Studio ?

Les deux servent des modèles en local, mais Ollama fonctionne d’abord en ligne de commande avec une API HTTP, pensée pour s’intégrer à des outils tiers comme Continue. LM Studio propose une interface graphique complète pour parcourir, télécharger et discuter avec des modèles, avec en plus un serveur compatible API OpenAI activable en un clic.

### Continue.dev fonctionne-t-il encore après son rachat par Cursor ?

Oui. La version 2.0.0, sortie le 19 juin 2026, reste installable et pleinement fonctionnelle avec un modèle local via Ollama. Le dépôt GitHub est en lecture seule et n’accueillera plus de nouvelles fonctionnalités officielles, mais le code Apache 2.0 continue de tourner normalement, y compris pour la configuration décrite dans ce tutoriel.

### Quel modèle choisir si je n’ai pas de carte graphique dédiée ?

Partez sur qwen2.5-coder:7b en version quantifiée, qui tourne sur processeur avec 8 Go de RAM disponibles. Les temps de réponse seront plus longs qu’avec un GPU, mais l’autocomplétion et le chat restent utilisables pour un usage ponctuel.

### Mon code envoyé à Ollama quitte-t-il mon ordinateur ?

Non, tant que la configuration pointe vers `http://localhost:11434` comme indiqué à l’étape 6. Le modèle tourne entièrement sur votre machine et aucune requête ne part vers un serveur externe, sauf si vous ajoutez volontairement un modèle cloud dans la même configuration Continue.

### Ollama peut-il remplacer complètement GitHub Copilot ?

Pour l’autocomplétion et le chat contextuel du quotidien, oui, dans la plupart des cas. Sur des tâches d’architecture très complexes ou des raisonnements multi-étapes sur un très gros dépôt, un modèle cloud haut de gamme garde souvent l’avantage. Beaucoup d’équipes utilisent les deux en parallèle, en réservant le local au code sensible.

### Comment mettre à jour un modèle déjà téléchargé ?

Relancez simplement la commande `ollama pull` suivie du même nom de modèle. Ollama ne télécharge que les blocs modifiés par rapport à la version déjà présente sur le disque, ce qui accélère nettement la mise à jour par rapport à un premier téléchargement.

### Puis-je utiliser Ollama avec JetBrains ou Neovim au lieu de VS Code ?

Oui. Continue propose un plugin JetBrains qui reprend la même logique de configuration que celle décrite ici pour VS Code. Pour Neovim, plusieurs intégrations communautaires consomment directement l’API Ollama sur le port 11434, avec un principe de fonctionnement identique.
