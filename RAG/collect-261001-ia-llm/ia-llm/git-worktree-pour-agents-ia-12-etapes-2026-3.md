---
id: collect-261001-ia-llm/ia-llm/git-worktree-pour-agents-ia-12-etapes-2026-3
title: "Exemple de sortie : /Users/marie/projets/api-recettes"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "claude"]
source: docs/RAG/collect-261001-ia-llm/git-worktree-pour-agents-ia-12-etapes-2026.md
source_anchor: ""
source_lines: [157, 277]
sha256: ad736449d75f8030739d90668fa410323f2b56bd1c4fbc674531cf99447c3a76
---

# Exemple de sortie : /Users/marie/projets/api-recettes

Git worktree ne partage que ce qui est suivi par Git. Les dossiers `node_modules`, les environnements virtuels Python ou les dossiers `vendor` ne le sont pas, et doivent donc être réinstallés dans chaque worktree.

```
# Dans chaque worktree Node.js
npm install
# Dans chaque worktree Python
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
```
Avec trois ou quatre worktrees actifs, cette duplication peut représenter plusieurs gigaoctets rapidement, surtout sur des projets Node.js avec de grosses arborescences de dépendances. Deux options réduisent l’impact réel. La première consiste à utiliser un gestionnaire de paquets qui mutualise le cache global, comme pnpm : les fichiers sont stockés une seule fois sur le disque et liés physiquement dans chaque `node_modules`, ce qui limite l’espace réellement consommé même si chaque worktree paraît complet. La seconde, plus radicale, consiste à ne réinstaller les dépendances que dans les worktrees réellement actifs, et à en supprimer certains dès qu’une tâche est terminée plutôt que de les laisser s’accumuler pendant des semaines.

## Étape 10 – Verrouiller, déplacer et réparer vos worktrees

Trois commandes complètent la boîte à outils pour une utilisation prolongée. `git worktree lock` empêche Git de considérer un worktree comme obsolète, utile pour une tâche longue qui dure plusieurs jours ou pour un worktree stocké sur un disque externe qui n’est pas toujours connecté :

```
git worktree lock .trees/fix-auth-401 --reason "Migration longue, ne pas purger"
git worktree unlock .trees/fix-auth-401
```
`git worktree move` renomme ou déplace un worktree existant sans casser son lien avec le dépôt principal, ce qui est plus sûr qu’un simple `mv` dans le terminal :

`git worktree move .trees/fix-auth-401 .trees/fix-auth-legacy`
Si malgré tout un worktree a été déplacé à la main, ou si les métadonnées internes se retrouvent désynchronisées après une synchronisation de dossier ou une restauration de sauvegarde, `git worktree repair` recalcule les références et remet tout en ordre :

`git worktree repair`
## Étape 11 – Nettoyer avec remove et prune, puis fusionner le travail validé

Une fois qu’un agent a terminé sa tâche et que sa branche a été fusionnée, supprimez proprement le worktree correspondant plutôt que de le laisser traîner :

`git worktree remove .trees/fix-auth-401`
Git refuse cette suppression si des fichiers non validés subsistent, sauf à ajouter `--force`, ce qui protège contre une perte de travail accidentelle. Si un worktree a été supprimé manuellement avec un simple `rm -rf` plutôt qu’avec la commande dédiée, Git conserve une référence fantôme dans ses métadonnées. La commande `git worktree prune` nettoie ces références orphelines :

```
git worktree prune
git worktree list   # confirme que le worktree fantôme a disparu
```
Prenez l’habitude de lancer `git worktree list` et `git worktree prune` en début de semaine. Sur un projet actif avec plusieurs agents, les worktrees oubliés s’accumulent plus vite qu’on ne l’imagine.

## Étape 12 – Automatiser la création de worktrees avec un script shell

Répéter les mêmes commandes à chaque nouvelle tâche devient vite fastidieux. Un petit script bash suffit à standardiser la création d’un worktree d’agent, copie du `.env` comprise :

```
#!/usr/bin/env bash
# nouveau-worktree-agent.sh
set -e
TACHE=$1
if [ -z "$TACHE" ]; then
  echo "Usage : ./nouveau-worktree-agent.sh nom-de-la-tache"
  exit 1
fi
RACINE=$(git rev-parse --show-toplevel)
CIBLE="$RACINE/.trees/$TACHE"
git worktree add "$CIBLE" -b "agent/$TACHE"
cp "$RACINE/.env" "$CIBLE/.env"
echo "Worktree prêt : $CIBLE (branche agent/$TACHE)"
```
Rendez-le exécutable une fois avec `chmod +x nouveau-worktree-agent.sh`, puis lancez `./nouveau-worktree-agent.sh corrige-pagination` pour obtenir un worktree prêt à l’emploi en une seule commande, avec une convention de nommage cohérente sur toute l’équipe.

## Projet complet : orchestrer 3 agents IA en parallèle sur un vrai dépôt

### Le scénario : une API de recettes à améliorer

Prenons un exemple concret et complet, du premier commit à la fusion finale. Le projet est une petite API REST Node.js, `api-recettes`, avec trois tickets ouverts : un bug d’authentification, une fonctionnalité de pagination, et un manque de couverture de tests. Plutôt que de les traiter l’un après l’autre, l’objectif est de les faire avancer en parallèle grâce à trois worktrees et trois agents différents.

### Mise en place des trois worktrees

```
cd ~/projets/api-recettes
git pull origin main
./nouveau-worktree-agent.sh fix-auth-401
./nouveau-worktree-agent.sh pagination-v2
./nouveau-worktree-agent.sh tests-integration
git worktree list
# ~/projets/api-recettes                          [main]
# ~/projets/api-recettes/.trees/fix-auth-401       [agent/fix-auth-401]
# ~/projets/api-recettes/.trees/pagination-v2      [agent/pagination-v2]
# ~/projets/api-recettes/.trees/tests-integration  [agent/tests-integration]
```
Chaque worktree reçoit ensuite ses dépendances et son port dédié, puis son agent :

```
# Terminal 1
cd .trees/fix-auth-401 && npm install && PORT=3001 npm run dev &
claude
# Terminal 2
cd .trees/pagination-v2 && npm install
cursor .
# Terminal 3
cd .trees/tests-integration && npm install
codex
```
### Résultat : trois pull requests indépendantes en une matinée

Claude Code corrige la validation des tokens expirés et committe sur `agent/fix-auth-401`. Cursor construit la pagination sur `agent/pagination-v2`, avec son propre serveur de développement sur le port 3002, sans jamais voir les fichiers modifiés par Claude Code. Codex CLI écrit une suite de tests d’intégration sur `agent/tests-integration`, en lisant le code de la branche principale, indépendamment des deux autres chantiers.

Une fois chaque tâche validée, la fusion se fait normalement depuis le dépôt principal, une branche à la fois :

```
cd ~/projets/api-recettes
git checkout main
git merge agent/fix-auth-401
git worktree remove .trees/fix-auth-401
git branch -d agent/fix-auth-401
```
Répétez pour les deux autres branches une fois leurs pull requests respectives approuvées. Le résultat concret : trois chantiers qui, traités en série, auraient occupé toute une journée, avancent en parallèle et convergent en fin de matinée, chacun avec son propre historique de commits, propre et traçable.

## Git worktree face aux alternatives : branches, clones et subtrees

Git worktree n’est pas la seule façon d’isoler du travail parallèle. Voici comment elle se compare aux approches historiques, sur les critères qui comptent réellement pour un usage avec des agents IA.

| Approche | Isolation des fichiers | Espace disque | Historique partagé | Adapté aux agents IA | 
|---|---|---|---|---|
| Git worktree | Complète (dossier séparé) | Faible (objets partagés) | Oui, instantané | Oui, conçu pour ça | 
| Changement de branche classique | Aucune (même dossier) | Minimal | Oui | Non, un seul état à la fois | 
| Clone séparé du dépôt | Complète | Élevé (historique dupliqué) | Non, nécessite un push/pull | Partiel, mais lourd | 
| Git subtree / submodule | Partielle | Variable | Complexe à synchroniser | Non, pensé pour un autre usage | 

Le clone séparé reste utile dans un cas précis : si vous devez travailler sur deux dépôts distants différents, ou tester une version totalement isolée sans même partager la configuration Git locale. Pour tout le reste, et en particulier pour plusieurs agents sur le même projet, le worktree est la solution la plus légère et la plus rapide à mettre en place.

