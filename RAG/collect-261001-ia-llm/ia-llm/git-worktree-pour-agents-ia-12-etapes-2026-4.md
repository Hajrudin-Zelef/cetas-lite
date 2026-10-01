---
id: collect-261001-ia-llm/ia-llm/git-worktree-pour-agents-ia-12-etapes-2026-4
title: "Exemple de sortie : /Users/marie/projets/api-recettes"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-ia-llm/git-worktree-pour-agents-ia-12-etapes-2026.md
source_anchor: ""
source_lines: [278, 358]
sha256: b8a4254b3b7fd0bbf37e58c448f78bbfcb2433b44a887f7aa45bff1817ecd66b
---

# Exemple de sortie : /Users/marie/projets/api-recettes

Le changement de branche classique garde son utilité pour un travail strictement séquentiel, une seule tâche à la fois, sans besoin de revenir en arrière rapidement. Dès qu’un second agent ou un second développeur entre en jeu sur le même projet au même moment, la bascule vers les worktrees change la donne : plus besoin de stasher des modifications en cours pour aller vérifier autre chose, plus de risque d’oublier un fichier non validé avant de changer de contexte.

## Tableau récapitulatif des commandes git worktree

| Commande | Rôle | Exemple | 
|---|---|---|
| worktree add | Crée un nouveau worktree, avec ou sans nouvelle branche | git worktree add ../app-feature -b feature/x | 
| worktree list | Liste tous les worktrees enregistrés | git worktree list –porcelain | 
| worktree remove | Supprime un worktree propre (sans fichiers non validés) | git worktree remove ../app-feature | 
| worktree prune | Nettoie les références orphelines après une suppression manuelle | git worktree prune | 
| worktree lock | Empêche la purge automatique d’un worktree | git worktree lock ../app-feature –reason “en cours” | 
| worktree unlock | Retire le verrou posé par lock | git worktree unlock ../app-feature | 
| worktree move | Déplace ou renomme un worktree sans casser le lien | git worktree move ../ancien ../nouveau | 
| worktree repair | Recalcule les métadonnées après un déplacement manuel | git worktree repair | 

## 5 pièges courants à éviter avec les worktrees et les agents IA

- **Oublier de copier le fichier .env :** le worktree hérite du code suivi par Git, jamais des fichiers ignorés. Sans script d’automatisation, cette étape saute une fois sur deux et l’agent se retrouve à déboguer une erreur de configuration au lieu de sa vraie tâche.
- **Laisser deux agents partager le même port ou la même base de données :** les symptômes sont trompeurs, un agent semble « ignorer » les changements de l’autre alors qu’en réalité les deux serveurs se battent pour la même ressource.
- **Travailler directement dans le dépôt principal pendant qu’un agent est actif ailleurs :** le dossier racine doit rester réservé aux fusions et à la lecture, jamais à l’édition simultanée.
- **Accumuler des worktrees jamais nettoyés :** sur un projet actif, dix worktrees oubliés en un mois ne sont pas rares, et chacun consomme de l’espace disque via ses propres node_modules ou environnements virtuels.
- **Vouloir extraire deux fois la même branche :** Git bloque cette opération par conception, mais le message d’erreur surprend souvent les nouveaux utilisateurs, qui pensent avoir trouvé un bug plutôt qu’une protection.
- **Donner des noms de branches trop proches les uns des autres :** agent/fix, agent/fix2 et agent/fix-final finissent par se confondre dans les revues de code. Préférez un identifiant de ticket ou une description courte et unique, alignée sur le nom du dossier du worktree.

## Dépannage : 8 problèmes fréquents et leurs solutions

**« fatal: ‘nom-branche’ is already checked out at … »**

Cette branche est déjà active dans un autre worktree. Utilisez `git worktree list` pour la retrouver, ou créez une nouvelle branche dédiée avec `-b` plutôt que de réutiliser l’existante.

**« fatal: ‘<chemin>’ already exists »**

Le dossier cible existe déjà, souvent parce qu’un worktree précédent n’a pas été supprimé proprement. Vérifiez avec `git worktree list`, puis `git worktree remove` avant de recréer.

**Un worktree supprimé avec rm -rf apparaît encore dans git worktree list**

Git conserve une référence administrative tant qu’on ne lance pas `git worktree prune`. C’est le comportement attendu, pas un bug.

**« fatal: cannot remove a locked working tree »**

Le worktree a été verrouillé volontairement. Retirez le verrou avec `git worktree unlock <chemin>` avant de le supprimer.

**npm run dev échoue avec « port already in use »**

Deux worktrees tentent d’utiliser le même port. Attribuez un port distinct à chaque worktree via la variable PORT ou un fichier .env spécifique.

**L’agent IA ne voit pas les derniers commits de la branche principale**

Un worktree ne se synchronise pas automatiquement avec origin/main. Lancez `git fetch origin` puis `git rebase origin/main` ou `git merge origin/main` depuis le worktree concerné.

**« fatal: is not a working tree » après une synchronisation cloud (Dropbox, iCloud, OneDrive)**

Les outils de synchronisation cloud perturbent les métadonnées internes de Git. Excluez le dossier .git et les dossiers .trees de la synchronisation, ou lancez `git worktree repair` pour recalculer les références.

**git worktree move échoue avec une erreur de verrouillage**

Fermez tout processus qui a un fichier ouvert dans ce worktree (éditeur, serveur de développement, terminal actif) avant de tenter le déplacement.

**Une pull request d’agent contient des fichiers inattendus (node_modules, .env)**

Le .gitignore du dépôt principal ne couvre pas toujours les dossiers créés dans .trees. Ajoutez explicitement .trees/ à la racine du .gitignore pour éviter tout commit accidentel.

**Deux agents proposent des correctifs qui se contredisent sur le même fichier**

Ce n’est pas un problème de worktree à proprement parler, mais un problème de découpage des tâches. Vérifiez que les tickets assignés à chaque agent touchent des fichiers ou des modules suffisamment distincts avant de lancer les sessions en parallèle, sans quoi la fusion finale demandera un arbitrage manuel de toute façon.

## Astuces avancées pour les équipes européennes

Pour les équipes qui travaillent sur des dépôts contenant des données sensibles ou soumises au RGPD, gardez à l’esprit qu’un worktree copie les fichiers suivis sur le disque local, comme n’importe quel checkout Git. Si votre politique de sécurité interdit certaines données de configuration en dehors d’un environnement contrôlé, appliquez les mêmes règles de chiffrement de disque et de gestion des secrets à tous les worktrees, pas seulement au dépôt principal. Un secret oublié dans un `.env` copié à la hâte dans `.trees/` reste un secret exposé.

Côté productivité, un alias Git simplifie la vie au quotidien. Ajoutez ceci dans votre `~/.gitconfig` :

```
[alias]
    wt = worktree
    wtl = worktree list
```
Vous gagnez alors la possibilité de taper `git wtl` au lieu de la commande complète, un détail qui compte quand vous vérifiez l’état de vos worktrees une dizaine de fois par jour. Sur VS Code, l’extension native de gestion de dépôts affiche désormais les worktrees dans la barre latérale Source Control, ce qui évite de revenir sans cesse au terminal juste pour lister les sessions actives. Dans les IDE JetBrains (IntelliJ, PyCharm, WebStorm), chaque worktree s’ouvre simplement comme un projet distinct, avec sa propre fenêtre et son propre index de fichiers, sans configuration supplémentaire.

Enfin, pour les pipelines d’intégration continue, évitez de faire tourner des builds directement dans un dossier `.trees` partagé avec les développeurs. Préférez un checkout dédié dans l’environnement CI, qui reste indépendant des sessions d’agents locales et évite tout effet de bord entre une exécution automatisée et un poste de travail humain.

