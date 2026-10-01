---
id: collect-261001-ia-llm/ia-llm/installer-claude-code-13-etapes-50-min-2026-2
title: "La sortie doit afficher v18.x.x ou une version supérieure"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "attention", "claude", "distribution", "mai"]
source: docs/RAG/collect-261001-ia-llm/installer-claude-code-13-etapes-50-min-2026.md
source_anchor: ""
source_lines: [44, 136]
sha256: 2f91f6aadd4fdb8f172a9d6a242d2e4c7f4ab02ab8c275d7fe8997059f360d26
---

# La sortie doit afficher v18.x.x ou une version supérieure

## Prérequis techniques : versions, comptes et configuration minimale

Claude Code s’installe sur Windows, macOS et Linux. Avant de lancer la moindre commande, vérifiez les points suivants pour éviter les trois quarts des erreurs d’installation les plus courantes.

| Élément | Exigence minimale | Recommandation | 
|---|---|---|
| Système d’exploitation | Windows 10, macOS 12 ou distribution Linux récente | Dernière version stable de l’OS | 
| Node.js | Version 18 | Version 22 ou supérieure si vous installez via npm | 
| Mémoire vive | 4 Go disponibles | 8 Go ou plus sur de gros dépôts | 
| Terminal | Terminal natif, ou WSL sous Windows | Terminal Linux natif ou iTerm2 sur macOS pour un meilleur rendu | 
| Compte Anthropic | Abonnement Claude Pro (20 $/mois) ou crédits API | Claude Max pour un usage quotidien intensif | 
| Git | Recommandé, non obligatoire | Dernière version stable | 
| Espace disque | 500 Mo disponibles | 2 Go si plusieurs projets sont suivis en parallèle | 

Deux exigences méritent une explication. D’abord Node.js : Claude Code s’appuie sur le runtime Node pour fonctionner, même si vous ne développez pas vous-même en JavaScript. Une version trop ancienne ne provoque pas toujours une erreur explicite, elle se traduit parfois par des plantages silencieux au milieu d’une session. Ensuite le compte Anthropic : contrairement à un simple outil local, Claude Code a besoin d’une connexion active aux serveurs d’Anthropic pour fonctionner, ce qui exclut tout usage entièrement hors ligne, y compris pour les tâches les plus simples.

Point souvent négligé : Claude Code n’a pas de palier gratuit propre. Contrairement au chat Claude.ai, qui propose un accès sans abonnement, l’agent en ligne de commande exige au minimum un abonnement Claude Pro actif ou des crédits API valides. Gardez cette nuance en tête si vous cherchez à **utiliser Claude Code gratuitement** : ce n’est possible qu’en consommant des crédits API limités offerts ponctuellement, pas via un forfait permanent à 0 euro. La documentation officielle sur le démarrage rapide de Claude Code détaille les conditions d’accès à jour.

## Étapes 1 à 3 : préparer l’environnement de développement

Les trois premières étapes se jouent avant même d’installer quoi que ce soit. Elles évitent la majorité des blocages recensés dans la section dépannage plus bas.

- **Étape 1** : vérifiez la version de Node.js installée sur votre machine.
- **Étape 2** : mettez à jour Node.js si la version détectée est antérieure à la 18, en la téléchargeant depuis le site officiel.
- **Étape 3** : ouvrez un terminal directement dans le dossier du projet sur lequel vous comptez travailler, pas dans votre dossier utilisateur.

```
node --version
# La sortie doit afficher v18.x.x ou une version supérieure
# Exemple de sortie attendue : v22.11.0
```
Si la commande renvoie une erreur ou une version trop ancienne, téléchargez la version recommandée depuis la page officielle de téléchargement de Node.js. Sous Windows, privilégiez l’installateur MSI plutôt qu’une installation manuelle des binaires, cela simplifie la mise à jour du PATH.

La troisième étape mérite une attention particulière. Beaucoup de nouveaux utilisateurs lancent Claude Code depuis leur dossier personnel plutôt que depuis la racine d’un dépôt Git existant. L’agent fonctionne bien mieux quand il dispose d’un contexte de projet clair dès le départ.

## Étapes 4 à 6 : installer Claude Code selon trois méthodes possibles

Anthropic propose trois façons d’installer Claude Code. Elles aboutissent toutes au même résultat, mais l’une d’elles convient mieux selon votre système et vos habitudes.

**Étape 4 : installation native (méthode recommandée par Anthropic)**. Cette méthode télécharge et exécute un script officiel qui détecte automatiquement votre système d’exploitation.

```
# macOS, Linux ou WSL
curl -fsSL https://claude.ai/install.sh | bash
# Windows (PowerShell)
irm https://claude.ai/install.ps1 | iex
```
**Étape 5 : installation via npm**, pour les développeurs déjà habitués à l’écosystème Node.js. Le paquet reste pris en charge par Anthropic.

```
npm install -g @anthropic-ai/claude-code
# N'utilisez jamais sudo avec cette commande :
# cela corrompt les permissions du dossier global npm
```
**Étape 6 : installation via Homebrew**, réservée à macOS. Pratique si votre équipe gère déjà ses outils de développement via Homebrew.

```
brew install --cask claude-code
# Homebrew ne met pas à jour automatiquement Claude Code
# Pensez à lancer périodiquement :
brew upgrade claude-code
```
Ces trois méthodes sont documentées avec leurs variantes sur le dépôt GitHub officiel de Claude Code, utile si vous voulez suivre les changements entre versions. Depuis la mise à jour d’août 2026 du guide de démarrage rapide, la documentation officielle référence aussi des installeurs curl dédiés pour macOS, Linux, WSL, Windows PowerShell et CMD, ainsi qu’un cask Homebrew et un identifiant WinGet propre à Windows ; sur Linux, l’outil s’installe désormais directement via les gestionnaires de paquets système apt, dnf et apk, couvrant Debian, Fedora, RHEL et Alpine. Pour Homebrew, la documentation officielle du gestionnaire de paquets explique comment le remettre à niveau si vous ne l’avez jamais utilisé.

## Étapes 7 à 9 : premier lancement, authentification et vérification

**Étape 7** : vérifiez que l’installation a fonctionné en interrogeant la version installée. À titre de repère, la CLI affichait la version 2.1.145 sur le dépôt GitHub officiel à la mi-mai 2026, après une 2.1.129 diffusée quelques jours plus tôt lors de l’événement Code with Claude SF 2026 : si le numéro renvoyé par votre terminal est nettement inférieur, mieux vaut mettre à jour avant de poursuivre.

```
claude --version
# Exemple de sortie attendue :
# claude-code/1.x.x darwin-arm64 node-v22.11.0
```
**Étape 8** : placez-vous dans le dossier racine d’un projet réel (pas un dossier vide) et lancez l’agent.

```
cd mon-projet/
claude
```
**Étape 9** : au premier lancement, Claude Code ouvre une fenêtre de navigateur pour l’authentification. Connectez-vous avec le compte Anthropic associé à votre abonnement Claude Pro, Max, Team ou Enterprise. Une fois la connexion validée, le terminal affiche une invite de commande interactive prête à recevoir des instructions.

Dès ce premier lancement dans un dossier de projet, Claude Code demande l’autorisation de lire et modifier les fichiers de ce répertoire. Acceptez cette autorisation pour la racine du projet : c’est elle qui conditionne tout le reste du tutoriel.

## Étapes 10 et 11 : intégrer Claude Code à VS Code ou à un autre éditeur

**Étape 10** : ouvrez simplement le terminal intégré de votre éditeur (VS Code, JetBrains ou autre) et lancez la commande `claude` depuis cet emplacement. Claude Code détecte automatiquement certains éditeurs ouverts sur le même dossier et peut y afficher les modifications de fichiers en direct, sans extension supplémentaire à configurer manuellement dans la majorité des cas.

**Étape 11** : ajustez les réglages d’affichage si vous travaillez sur plusieurs écrans ou avec un terminal partagé. Beaucoup de développeurs gardent l’éditeur ouvert dans une moitié d’écran et le terminal Claude Code dans l’autre, ce qui permet de suivre les modifications de fichiers en temps réel pendant que l’agent travaille.

Cette configuration côte à côte change concrètement la façon de travailler. Au lieu de valider chaque suggestion ligne par ligne comme avec un auto-complétion classique, vous validez des lots de modifications plus larges, en relisant le diff proposé avant de l’accepter.

## Étapes 12 et 13 : configurer les permissions et le fichier CLAUDE.md

