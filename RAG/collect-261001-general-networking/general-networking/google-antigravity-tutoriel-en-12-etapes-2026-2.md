---
id: collect-261001-general-networking/general-networking/google-antigravity-tutoriel-en-12-etapes-2026-2
title: "macOS (Homebrew)"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "Google", "OpenAI", "xAI"]
dates: []
keywords: ["agent", "agents", "attention", "chatgpt", "claude", "gemini", "grok", "grok 4", "mai", "open-weight", "opus 4"]
source: docs/RAG/collect-261001-general-networking/google-antigravity-tutoriel-en-12-etapes-2026.md
source_anchor: ""
source_lines: [36, 109]
sha256: e120b549520b84d57368e421bd59dbe2a4616412900ff5fdd3bc0936aa9a74e2
---

# macOS (Homebrew)

| Prérequis | Version minimale | Recommandé | Rôle | 
|---|---|---|---|
| Windows | Windows 10 64 bits | Windows 11 | Système hôte | 
| macOS | Monterey 12 | Sonoma 14+ | Système hôte | 
| Linux | 64 bits, glibc 2.28+ | Ubuntu 22.04+ | Système hôte | 
| Compte Google | Compte personnel | Compte personnel | Authentification, aperçu gratuit | 
| RAM | 8 Go | 16 Go+ | Agents et éditeur | 
| Espace disque | 4 Go libres | 10 Go+ | Installation et projets | 
| Python (démo) | 3.11 | 3.12+ | API du projet | 
| Navigateur | Chrome/Chromium récent | Chrome à jour | Sous-agent navigateur | 

Points d’attention pour l’Europe : l’aperçu gratuit requiert un **compte Google personnel** (les comptes Workspace scolaires ou d’entreprise peuvent être bloqués par votre administrateur). Aucune carte bancaire n’est demandée pour la version gratuite. Sous Linux, la page de téléchargement d’Antigravity liste toujours, pour la version 2.11.0 mise en ligne en août 2026, les mêmes contraintes `glibc 2.28+` et `glibcxx 3.4.25+` qui excluent les distributions très anciennes ; une Ubuntu 22.04 LTS ou plus récente ne pose aucun problème. Enfin, prévoyez une connexion stable : les agents dialoguent en continu avec les serveurs de Google.

Pour le projet complet de ce tutoriel, installez Python 3.11 ou supérieur et vérifiez sa présence avec `python3 --version`. Nous n’aurons besoin d’aucune base de données externe : l’application de gestion de tâches utilisera SQLite, embarqué avec Python. Le navigateur Chrome est requis uniquement pour l’étape de vérification automatisée par le sous-agent navigateur.

## Étapes 1 et 2 : télécharger et installer Google Antigravity

**Étape 1 – Téléchargement.** Rendez-vous sur antigravity.google et cliquez sur « Get started » (Commencer). La page détecte votre système d’exploitation et propose l’installeur adapté : `.exe` pour Windows, `.dmg` pour macOS, `.deb` ou archive `.tar.gz` pour Linux. Vous pouvez aussi passer par un gestionnaire de paquets, ce qui simplifie les mises à jour. Voici les commandes selon votre plateforme :

```
# macOS (Homebrew)
brew install --cask antigravity
# Windows (winget, PowerShell)
winget install Google.Antigravity
# Linux (Debian/Ubuntu, après téléchargement du .deb)
sudo apt install ./antigravity_latest_amd64.deb
# Linux (archive générique)
tar -xzf antigravity-linux-x64.tar.gz -C ~/apps/
~/apps/antigravity/antigravity
```
**Étape 2 – Installation.** Lancez l’installeur téléchargé et suivez les invites standard. Sous macOS, glissez l’application dans le dossier *Applications* puis autorisez son exécution dans *Réglages Système > Confidentialité et sécurité* si Gatekeeper la bloque. Sous Windows, si SmartScreen affiche un avertissement, cliquez sur « Informations complémentaires » puis « Exécuter quand même » (l’aperçu public n’est pas encore signé par tous les certificats). Au premier lancement, l’application vérifie les mises à jour ; laissez-la se mettre à jour vers la dernière version disponible.

Sortie attendue au premier démarrage dans le terminal (installation Linux via archive) :

```
[antigravity] starting…
[antigravity] checking for updates… up to date
[antigravity] no workspace open – showing welcome screen
[antigravity] sign-in required to access model providers
```
Si l’application se ferme immédiatement sous Linux, c’est presque toujours une question de dépendances graphiques manquantes (`libnss3`, `libgbm1`). Nous traitons ce cas précis dans la section dépannage. Une fois l’écran d’accueil affiché, vous êtes prêt pour la connexion.

## Étapes 3 et 4 : connexion Google et import de votre config VS Code

**Étape 3 – Connexion.** Sur l’écran d’accueil, cliquez sur « Sign in with Google ». Le navigateur par défaut s’ouvre pour l’authentification OAuth. Utilisez un **compte Google personnel** : c’est lui qui débloque l’accès gratuit à la gamme complète de modèles pendant l’aperçu public. Après validation, une redirection ramène automatiquement vers Antigravity. Vous voyez alors apparaître votre quota de départ et la liste des fournisseurs de modèles disponibles.

**Étape 4 – Import de votre configuration VS Code.** Comme Antigravity est un fork de VS Code, il propose d’importer vos extensions, thèmes et raccourcis lors du premier lancement. Acceptez si vous venez de VS Code : vos réglages (linters, formateurs, palette de commandes) sont conservés. Vous pouvez aussi choisir un thème et une disposition. Si vous utilisez déjà Claude Code, sachez qu’un même fichier de contexte projet est réutilisable d’un outil à l’autre, ce qui évite de tout reconfigurer.

Créez maintenant le dossier de travail du projet et ouvrez-le dans Antigravity (menu *File > Open Folder*). C’est dans ce dossier que l’agent va écrire tout le code de l’application de gestion de tâches.

```
mkdir -p ~/projets/gestionnaire-taches
cd ~/projets/gestionnaire-taches
git init
# Puis dans Antigravity : File > Open Folder > ~/projets/gestionnaire-taches
```
Initialiser Git dès maintenant est une bonne pratique : Antigravity crée des points de contrôle (checkpoints) au fil des modifications, et un dépôt Git vous donne un filet de sécurité supplémentaire pour revenir en arrière si un agent part dans une mauvaise direction. Nous ferons un premier commit après la génération du projet.

## Étape 5 : choisir le bon modèle IA

Le choix du modèle conditionne la vitesse, la qualité et la consommation de quota. Google Antigravity expose un sélecteur de modèles (icône du modèle dans la barre de l’agent, ou *Settings > Model providers*). Le moteur par défaut appartient à la famille Gemini 3, lancée par Google pour les développeurs en même temps qu’Antigravity, et depuis le 20 mai 2026, au lendemain de l’ouverture de Google I/O 2026, **Gemini 3.5 Flash** est disponible en accès général aussi bien dans Antigravity que via l’API Gemini – l’une des quelque 100 annonces faites par Google à cette occasion –, présenté par Google comme jusqu’à 4 fois plus rapide que les autres modèles de pointe – et Google est allé plus loin en intégrant, dès août 2026, **Gemini 3.6 Flash** comme nouveau modèle rapide par défaut, selon Emergent. De son côté, **Gemini 3.1 Pro** offre désormais une fenêtre de contexte pouvant atteindre 1 000 000 de tokens, précieuse sur les gros projets. Vous pouvez même assigner un modèle différent à chaque agent d’une même mission. Voici la gamme disponible d’après la documentation officielle et la fiche Wikipédia à jour en 2026.

| Modèle | Éditeur | Idéal pour | Coût en quota | 
|---|---|---|---|
| Gemini 3 Pro |  | Raisonnement, planification, gros contexte | Élevé | 
| Gemini 3 Flash |  | Itérations rapides, édition ciblée | Faible | 
| Claude Sonnet 4.6 | Anthropic | Code robuste, refactorisation | Moyen | 
| Claude Opus 4.6 | Anthropic | Tâches complexes, architecture | Élevé | 
| GPT-OSS-120B | OpenAI (open-weight) | Autonomie, coût maîtrisé | Faible à moyen | 

Notre recommandation pratique, mise à jour après le passage à Antigravity 2.0 : utilisez **Gemini 3.1 Pro**, avec sa fenêtre de contexte de 1 000 000 de tokens, pour la phase de planification sur les gros projets (l’agent doit décomposer le projet), puis basculez sur **Gemini 3.6 Flash** – la version rapide mise à jour en août 2026, qui succède à Gemini 3.5 Flash lancé à l’I/O et annoncé jusqu’à 4 fois plus rapide – ou **Claude Sonnet 4.6** pour les itérations d’écriture et de correction, moins coûteuses en quota. Réservez **Claude Opus 4.6** aux décisions d’architecture délicates. Pour comparer plus finement les modèles Gemini face à la concurrence, consultez notre analyse Grok 4 vs ChatGPT vs Gemini 3.

### Activer Claude avec votre propre clé API

