---
id: collect-261001-ia-llm/ia-llm/github-copilot-agent-mode-14-etapes-55-min-2026-5
title: ".github/copilot-instructions.md"
domain: ia-llm
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["copilot", "agent", "diffusion", "mcp"]
source: docs/RAG/collect-261001-ia-llm/github-copilot-agent-mode-14-etapes-55-min-2026.md
source_anchor: ""
source_lines: [308, 371]
sha256: 54ad946af8368ade04767eb374e5f98e594325511b7773fcf5a06448f55b806e
---

# .github/copilot-instructions.md

```
gh extension install github/gh-copilot
# Exemple d'usage : demander une commande shell
gh copilot suggest "trouver tous les fichiers Python modifiés cette semaine"
# Exemple d'usage : expliquer une commande existante
gh copilot explain "find . -mtime -7 -name '*.py'"
```
Cette voie en ligne de commande complète bien l’Agent Mode intégré à VS Code : elle convient aux scripts d’intégration continue, aux environnements sans interface graphique, ou tout simplement aux développeurs qui passent le plus clair de leur temps dans un terminal plutôt que dans un éditeur. Copilot CLI est passé en disponibilité générale pour tous les abonnés Copilot le 25 février 2026 ; chaque invite y consomme désormais une requête premium, facturée 0,04 $ au-delà du quota inclus dans votre palier, exactement le même mécanisme de facturation que pour l’Agent Mode intégré à VS Code.

## Cinq erreurs fréquentes avec l’Agent Mode (et comment les éviter)

La plupart des déconvenues rapportées par les développeurs qui débutent avec l’Agent Mode viennent des mêmes habitudes mal ajustées. En voici six à corriger dès le départ.

- **Un prompt trop vague.** Écrire « corrige le bug » sans préciser lequel, ni le comportement attendu, pousse l’agent à deviner. Décrivez le symptôme observé, le fichier concerné si vous le connaissez, et le résultat souhaité.
- **Valider les changements sans les relire.** Cliquer sur Keep par réflexe, sans ouvrir le diff, revient à donner un accès en écriture non supervisé à votre code. Prenez systématiquement dix secondes pour parcourir les fichiers modifiés.
- **Laisser filer des commandes destructrices.** Une commande comme une suppression de fichiers ou une migration de base de données mérite une lecture attentive avant validation, jamais une approbation automatique.
- **Oublier de committer avant une longue session.** Sans point de restauration Git propre, revenir en arrière après une séquence de modifications ratée devient pénible. Committez avant de lancer une tâche agentique ambitieuse.
- **Ignorer le compteur de crédits.** Une session qui boucle plusieurs fois sur le même problème consomme des crédits sans forcément progresser. Surveillez le tableau de bord, surtout sur le palier Pro.
- **Utiliser un modèle premium pour tout, même le trivial.** Renommer une variable ne nécessite pas le modèle le plus coûteux disponible. Réservez les modèles premium aux tâches qui le justifient réellement.

## Dépannage : 9 problèmes courants et leurs solutions

La majorité de ces blocages se résolvent en moins de cinq minutes une fois la cause identifiée. Voici les blocages les plus signalés par les développeurs qui configurent l’Agent Mode pour la première fois, avec une piste de résolution pour chacun.

| Problème | Cause probable | Solution | 
|---|---|---|
| Agent absent du sélecteur de mode | VS Code ou l’extension sont trop anciens | Mettez à jour VS Code (1.99+) et réinstallez l’extension Copilot | 
| Compte non reconnu après connexion | Synchronisation lente entre GitHub et l’extension | Déconnectez-vous puis reconnectez-vous, patientez quelques minutes | 
| L’agent boucle sans converger | Prompt trop large ou tâche mal délimitée | Reformulez en une tâche plus petite, redémarrez une session propre | 
| Commandes terminal qui échouent silencieusement | Environnement virtuel non activé ou permissions insuffisantes | Activez l’environnement attendu avant de lancer l’agent | 
| Instructions personnalisées ignorées | Chemin ou nom de fichier incorrect | Vérifiez l’emplacement exact : .github/copilot-instructions.md à la racine | 
| Serveur MCP qui ne se connecte pas | Commande de démarrage ou variable d’environnement mal configurée | Consultez les journaux MCP dans la vue de sortie de VS Code | 
| Crédits épuisés en cours de tâche | Session agentique longue sur un modèle premium | Basculez sur un modèle standard ou patientez le renouvellement mensuel | 
| Modifications appliquées sur la mauvaise branche | Branche Git active non vérifiée avant de lancer l’agent | Contrôlez systématiquement votre branche avec git status avant de démarrer | 
| Réponses lentes ou latence élevée | Charge réseau ou modèle premium très sollicité | Vérifiez votre connexion, testez avec un modèle standard pour comparer | 

## Astuces avancées pour aller plus loin

Une fois les bases maîtrisées, quelques réglages supplémentaires changent nettement l’expérience au quotidien.

- **Passez par le mode Plan avant l’Agent Mode** sur les tâches ambitieuses. Plan génère une feuille de route détaillée que vous pouvez amender avant de laisser l’agent exécuter quoi que ce soit, ce qui réduit les allers-retours inutiles.
- **Utilisez des instructions ciblées par dossier** avec la syntaxe applyTo plutôt qu’un seul fichier d’instructions générique, surtout sur un monorepo qui mélange plusieurs technologies.
- **Essayez le mode Cloud** pour les tâches longues, comme une migration de dépendances sur un gros projet, qui peuvent s’exécuter en arrière-plan sans monopoliser votre machine locale.
- **Combinez Copilot CLI avec vos pipelines existants** pour automatiser des vérifications ou des suggestions de commandes directement dans un script shell.
- **Créez des modes de discussion personnalisés** pour les tâches récurrentes de votre équipe, comme la revue de sécurité d’une pull request ou la génération de documentation, afin de ne pas réécrire le même prompt à chaque fois.

## GitHub Copilot Agent Mode face à Cursor et Windsurf

Trois outils dominent aujourd’hui la conversation sur la programmation assistée par agent : GitHub Copilot, Cursor et Windsurf. Leur approche technique diverge sur un point central. Copilot reste une extension qui s’installe sur n’importe quelle version de VS Code, tandis que Cursor et Windsurf sont des forks complets de l’éditeur, distribués comme des applications à part entière. Le tableau ci-dessous résume les différences qui comptent le plus au moment de choisir.

| Critère | GitHub Copilot | Cursor AI | Windsurf | 
|---|---|---|---|
| Base technique | Extension pour VS Code, JetBrains, Neovim | Fork complet de VS Code | Fork de VS Code (sous Cognition AI) | 
| Entrée de gamme | Gratuit (2 000 complétions/mois) | Gratuit (Hobby) | Gratuit | 
| Palier individuel courant | 10 $/mois (Pro) | 20 $/mois (Pro) | environ 20 $/mois (Pro) | 
| Palier le plus élevé | 100 $/mois (Max) | 200 $/mois (Ultra) | 200 $/mois (Max) | 
| Palier équipe | 19 $/utilisateur/mois (Business) | 40 $/utilisateur/mois | 40 $/utilisateur/mois | 
| Support MCP | Oui, natif | Oui, natif | Partiel selon les sources | 
| Société éditrice | Microsoft / GitHub | Anysphere | Cognition AI | 

Le principal avantage de Copilot tient à sa diffusion : il s’installe sur l’éditeur que vous utilisez déjà, sans migration, et couvre aussi JetBrains et Neovim en plus de VS Code. Cursor et Windsurf misent au contraire sur une intégration plus poussée entre l’éditeur et l’agent, au prix d’un changement d’outil complet. Si vous voulez approfondir la configuration de Cursor spécifiquement, notre tutoriel Cursor AI en 13 étapes couvre l’installation et la prise en main de cet éditeur.

Dans les faits, le choix se joue souvent sur des critères plus organisationnels que techniques. Une équipe déjà installée sur VS Code, JetBrains ou Neovim a peu d’intérêt à migrer vers un éditeur entièrement nouveau juste pour l’Agent Mode, alors que Copilot s’ajoute en quelques minutes à un poste existant. À l’inverse, une équipe qui démarre un projet de zéro et cherche l’intégration agentique la plus poussée dès l’installation peut légitimement préférer un éditeur conçu autour de l’agent dès le départ.

