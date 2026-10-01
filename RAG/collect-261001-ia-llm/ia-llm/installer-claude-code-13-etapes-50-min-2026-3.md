---
id: collect-261001-ia-llm/ia-llm/installer-claude-code-13-etapes-50-min-2026-3
title: "La sortie doit afficher v18.x.x ou une version supérieure"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "claude"]
source: docs/RAG/collect-261001-ia-llm/installer-claude-code-13-etapes-50-min-2026.md
source_anchor: ""
source_lines: [137, 214]
sha256: 46a68a5c6b78c7a369d990d2f0cea5d0d91d73351624ea00d3b6336ded3a8ff1
---

# La sortie doit afficher v18.x.x ou une version supérieure

**Étape 12** : passez en revue les permissions accordées à l’agent. Par défaut, Claude Code demande une confirmation avant chaque action potentiellement destructrice (suppression de fichier, exécution de commande système, écriture hors du dossier du projet). Ne désactivez pas ces confirmations tant que vous n’avez pas testé l’outil sur un projet non critique.

Ce modèle de permissions granulaire répond directement à la méfiance mesurée par Stack Overflow chez près de la moitié des développeurs interrogés. Concrètement, cela signifie qu’un agent mal configuré reste rare : la plupart des incidents rapportés par les utilisateurs viennent d’une confirmation acceptée trop vite, pas d’une action que l’outil aurait menée seul sans validation.

**Étape 13** : créez un fichier `CLAUDE.md` à la racine du projet. Ce fichier donne à l’agent un contexte permanent, relu à chaque session : conventions de code, commandes de build et de test, structure des dossiers importants. C’est l’étape la plus rentable de tout ce tutoriel pour la qualité des réponses obtenues par la suite.

```
# CLAUDE.md
## Commandes utiles
- Installer les dépendances : npm install
- Lancer les tests : npm test
- Démarrer le serveur local : npm run dev
## Conventions
- TypeScript strict, pas de "any" sans commentaire justificatif
- Tests obligatoires pour toute nouvelle route d'API
- Commits en anglais, au format impératif court
## Structure du projet
- src/routes/ : définitions des routes Express
- src/services/ : logique métier
- tests/ : tests unitaires et d'intégration (Vitest)
```
Ce fichier n’a rien d’obligatoire pour que Claude Code fonctionne, mais son absence explique une bonne partie des réponses génériques ou mal adaptées que rapportent certains utilisateurs débutants.

## Exemple de projet complet : automatiser une tâche réelle

Pour rendre ce tutoriel concret, voici un exemple de bout en bout sur un petit projet réaliste : une API Express minimale à laquelle on demande d’ajouter une route et sa couverture de tests.

Le projet de départ contient un serveur Express simple avec une route `/health`. L’objectif : ajouter une route `/users/:id` qui retourne un utilisateur fictif, avec sa validation d’entrée et ses tests. Une fois Claude Code lancé dans le dossier du projet (étape 8), il suffit de formuler la demande en langage naturel.

```
> Ajoute une route GET /users/:id qui valide que l'id est
  un entier positif, renvoie une erreur 400 sinon, et retourne
  un utilisateur fictif au format JSON. Ajoute aussi les tests
  correspondants avec Vitest, en couvrant le cas valide et le
  cas d'erreur.
```
Claude Code lit d’abord les fichiers existants dans `src/routes/` pour repérer les conventions déjà en place (celles décrites dans le fichier CLAUDE.md), puis propose un plan avant d’écrire quoi que ce soit. Le résultat typique inclut trois modifications : la nouvelle route, une fonction de validation partagée, et un fichier de test. Le développeur relit le diff proposé pour chaque fichier avant validation, exactement comme il le ferait pour une pull request d’un collègue.

Une fois les fichiers acceptés, une deuxième instruction permet de vérifier que tout fonctionne :

`> Lance la suite de tests et corrige les éventuels échecs`
L’agent exécute alors la commande définie dans le fichier CLAUDE.md (`npm test`), lit la sortie, et corrige lui-même les erreurs de syntaxe ou d’assertions ratées si le premier passage échoue. Ce cycle lire, écrire, tester, corriger, sans repasser par un copier-coller manuel entre le terminal et l’éditeur, résume assez bien ce qui distingue un agent comme Claude Code d’un simple auto-complétion.

Le même projet permet d’illustrer un second usage, plus proche de la maintenance que de l’ajout de fonctionnalité. Sur un dépôt existant depuis plusieurs années, une instruction du type “recense les dépendances npm qui n’ont pas été mises à jour depuis plus d’un an et propose un plan de migration prudent” pousse l’agent à lire le fichier de verrouillage des dépendances, croiser les versions installées avec les versions publiées, puis rédiger un plan classé par risque plutôt que d’appliquer directement les mises à jour. Cette prudence par défaut, proposer avant d’agir sur un périmètre large, revient dans la plupart des tâches de maintenance et explique pourquoi beaucoup d’équipes gardent Claude Code pour ce type de travail plutôt que pour l’écriture de fonctionnalités entièrement nouvelles.

## À quoi ressemble une session Claude Code : exemples de sortie

Voici un exemple simplifié de ce qu’affiche le terminal pendant une session type, utile pour vérifier que votre propre installation se comporte normalement.

```
$ claude
Claude Code v1.x — connecté en tant que [email protected]
Dossier de travail : ~/projets/mon-api
> Ajoute une route GET /users/:id ...
● Lecture de src/routes/index.js
● Lecture de CLAUDE.md
● Plan proposé :
  1. Créer src/routes/users.js
  2. Ajouter la validation dans src/services/validation.js
  3. Créer tests/users.test.js
Accepter ce plan ? (o/n)
```
Si votre terminal affiche une sortie très différente de ce modèle (pas de proposition de plan, aucune lecture de fichiers avant écriture, absence de demande de confirmation), reportez-vous à la section dépannage : cela indique souvent une installation partielle ou une version ancienne restée active en parallèle.

## Les erreurs courantes à éviter avec Claude Code

Ces pièges reviennent le plus souvent chez les nouveaux utilisateurs, d’après les retours recensés sur les forums de développeurs et la documentation communautaire. La plupart se règlent en quelques secondes une fois identifiés, mais ils suffisent à décourager un premier essai si personne ne les a signalés à l’avance.

- **Lancer l’agent depuis le mauvais dossier.** Démarrer Claude Code depuis le dossier utilisateur plutôt que depuis la racine d’un projet prive l’agent du contexte dont il a besoin, et peut l’amener à modifier des fichiers hors du périmètre voulu.
- **Installer avec les droits root ou sudo.** Sous macOS et Linux, forcer l’installation npm avec sudo corrompt fréquemment les permissions du dossier global, ce qui provoque des erreurs à chaque mise à jour ultérieure.
- **Ignorer la vérification de version de Node.js.** Une version trop ancienne provoque des erreurs cryptiques au lancement plutôt qu’un message clair.
- **Ne jamais créer de fichier CLAUDE.md.** Sans ce fichier, l’agent doit redécouvrir les conventions du projet à chaque session, ce qui allonge les échanges et dégrade la pertinence des réponses.
- **Accepter toutes les actions sans relire les diffs.** Désactiver les confirmations de sécurité pour aller plus vite expose à des suppressions ou modifications non désirées, surtout sur un dépôt partagé.
- **Confondre Claude Code et Claude.ai.** Le chat web Claude.ai et l’agent en ligne de commande Claude Code partagent le même modèle mais pas le même mode d’accès ni la même tarification, une confusion qui explique une partie des questions sur le tarif ou la gratuité de l’outil.
- **Oublier de surveiller la consommation de crédits API.** En facturation à l’usage, une session longue sur un dépôt volumineux peut consommer plus de jetons que prévu si aucune alerte de budget n’est configurée côté console Anthropic.

## Dépannage : les problèmes les plus fréquents et leurs solutions

