---
id: collect-261001-ia-llm/ia-llm/gemini-cli-tutoriel-gratuit-en-12-etapes-2026-4
title: "Vérifier les versions actuelles"
domain: ia-llm
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["agent", "gemini", "license", "open source", "sandbox"]
source: docs/RAG/collect-261001-ia-llm/gemini-cli-tutoriel-gratuit-en-12-etapes-2026.md
source_anchor: ""
source_lines: [269, 362]
sha256: 4ae14bd7f732f811ce7ea76b558276946d36ad154d09d28e4469ba9577f5a189
---

# Vérifier les versions actuelles

Gemini CLI n’est pas réservé aux sessions interactives. Le **drapeau `-p`** (pour *prompt*) exécute une requête unique puis rend la main, ce qui le rend parfaitement scriptable. Vous pouvez ainsi l’intégrer dans un pipeline CI/CD, un hook Git ou une tâche planifiée. La sortie s’imprime sur la sortie standard, prête à être redirigée ou traitée.

```
# Requête unique, sortie sur stdout
gemini -p "Génère un message de commit à partir du diff suivant"
# Alimenter l'agent via un pipe
git diff --staged | gemini -p "Rédige un message de commit Conventional Commits"
# Exemple de script : générer des notes de version
#!/usr/bin/env bash
set -euo pipefail
LAST_TAG=$(git describe --tags --abbrev=0)
git log "${LAST_TAG}..HEAD" --pretty=format:'%s' \
  | gemini -p "Résume ces commits en notes de version claires, en français" \
  > CHANGELOG_DRAFT.md
echo "Brouillon de changelog écrit dans CHANGELOG_DRAFT.md"
```
Ce mode headless ouvre d’immenses possibilités d’automatisation : revue de code automatique sur chaque pull request, génération de documentation à la volée, triage d’issues, ou encore rédaction de tests à partir d’une spécification. Combinez-le avec une clé API et la facturation activée pour disposer de quotas fiables en environnement d’intégration continue, où les limites du palier gratuit seraient rapidement atteintes.

Pour aller plus loin, Gemini CLI propose une intégration **GitHub Actions** via la commande `/setup-github`, qui configure des workflows de triage d’issues et de revue de pull requests directement dans votre dépôt. C’est la manière la plus simple de faire travailler l’agent en continu sur votre projet, sans intervention manuelle.

## Étape 11 – Sécurité : sandbox, dossiers de confiance et checkpointing

Un agent qui exécute des commandes shell et modifie des fichiers doit être encadré. Gemini CLI intègre trois garde-fous complémentaires. Le premier est le **sandbox** : avec le drapeau `--sandbox` (ou `-s`), les outils potentiellement dangereux s’exécutent dans un conteneur isolé (Docker ou Podman), ce qui protège votre système hôte des effets de bord. Activez-le dès que vous laissez l’agent travailler de façon autonome.

```
# Lancer l'agent avec isolation par conteneur
gemini --sandbox
# Ou via une variable d'environnement
export GEMINI_SANDBOX=true
gemini
# Passer en mode plan (lecture seule) : l'agent propose sans exécuter
> /plan
# Restaurer les fichiers à l'état d'avant la dernière action
> /restore
```
Le deuxième garde-fou est le système de **dossiers de confiance** : Gemini CLI vous demande d’autoriser explicitement un dossier avant d’y opérer, et gère des permissions persistantes. Le troisième est le **checkpointing** : avant toute modification de fichier, l’agent crée un point de restauration. Les commandes `/restore` et `/rewind` vous permettent d’annuler une action ou de revenir en arrière dans la conversation, comme un « ctrl-Z » pour votre agent.

Enfin, le **mode plan** (`/plan`) fait passer l’agent en lecture seule : il élabore et vous présente un plan d’action sans rien exécuter. C’est idéal pour valider une stratégie avant de laisser l’agent agir. À l’inverse, méfiez-vous du mode « YOLO » (auto-approbation de tous les outils) : commode pour les démonstrations, il est dangereux sur du code réel. Ne l’utilisez jamais hors d’un environnement jetable et isolé.

## Étape 12 – Projet complet : auditer et documenter un dépôt

Assemblons maintenant tout ce que vous avez appris en un **projet complet et fonctionnel** : transformer Gemini CLI en agent d’audit et de documentation pour n’importe quel dépôt. L’objectif : obtenir un rapport de qualité de code, un fichier de contexte, et une documentation générée, en une seule session encadrée.

```
# 1. Se placer dans le dépôt et lancer l'agent en sandbox
cd mon-depot
gemini --sandbox
# 2. Générer le contexte projet
> /init
# 3. Passer en mode plan pour cadrer l'audit (lecture seule)
> /plan
> Audite ce dépôt : dépendances obsolètes, code mort, risques de sécurité,
  couverture de tests. Ne modifie rien, propose un plan priorisé.
# 4. Valider le plan, puis demander l'exécution ciblée
> Applique uniquement les points 1 et 3 du plan. Confirme chaque écriture.
# 5. Générer la documentation
> @src/  Rédige un fichier docs/ARCHITECTURE.md décrivant les modules
  et leurs interactions, en français.
# 6. Vérifier le résultat, puis committer
> !npm test
> git diff --staged | gemini -p "Rédige un message de commit détaillé"
```
Voici un extrait de la sortie attendue lors de la phase d’audit. Remarquez comment l’agent enchaîne les outils – lecture de l’arborescence, analyse des dépendances, recherche web pour vérifier une CVE – avant de produire un rapport structuré et priorisé.

```
✦ Audit du dépôt en cours...
  [outil] ReadFolder   src/           -> 27 fichiers
  [outil] ReadFile     package.json
  [outil] GoogleSearch "CVE express 4 vulnérabilité 2026"
  [outil] Shell        npm outdated
  Rapport d'audit (priorisé)
  ──────────────────────────
  1. [CRITIQUE] 2 dépendances avec CVE connues -> mettre à jour
  2. [MAJEUR]   Couverture de tests : 41 % (cible conseillée : 80 %)
  3. [MINEUR]   3 fonctions non utilisées dans src/utils/
  4. [INFO]     Aucun fichier LICENSE détecté
  Souhaitez-vous que je traite les points 1 et 3 ? [O/N]
```
Vous disposez désormais d’un workflow réutilisable, transférable d’un projet à l’autre. En sauvegardant votre `GEMINI.md` et votre `settings.json` dans un dépôt de configuration, vous reproduisez ce comportement partout. C’est là toute la force d’un agent open source et scriptable : il devient une brique d’infrastructure, pas seulement un gadget de démonstration.

## 5 pièges courants à éviter avec Gemini CLI

La plupart des frustrations rencontrées avec Gemini CLI proviennent d’une poignée d’erreurs évitables. Les connaître à l’avance vous fera gagner des heures.

1. **Version de Node.js trop ancienne.** En dessous de la 20, l’installation ou le lancement échoue avec des erreurs obscures. Vérifiez toujours`node -v` et passez par`nvm` pour une version LTS.
2. **Contourner les permissions npm avec sudo.** L’erreur`EACCES` pousse les débutants à installer en root, ce qui casse ensuite les mises à jour. Utilisez`nvm` ou`npx` à la place.
3. **Saturer le contexte avec `@`.** Injecter`@./` sur un énorme dépôt consomme des centaines de milliers de tokens d’un coup, ralentit l’agent et épuise le quota. Ciblez des dossiers précis et utilisez`/compress` .
4. **Committer sa clé API.** Un fichier`.env` ou`.gemini/.env` versionné expose votre clé. Ajoutez-les au`.gitignore` dès le départ.
5. **Laisser l’agent agir sans sandbox ni confirmation.** Le mode auto-approbation (« YOLO ») sur du code de production peut supprimer ou écraser des fichiers. Gardez les confirmations actives et activez`--sandbox` .

## Dépannage : 8 problèmes fréquents et leurs solutions

Voici les incidents les plus signalés par la communauté et la façon de les résoudre rapidement. Conservez ce tableau à portée de main.

