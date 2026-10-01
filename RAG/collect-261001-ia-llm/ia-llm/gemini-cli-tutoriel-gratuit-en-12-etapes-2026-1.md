---
id: collect-261001-ia-llm/ia-llm/gemini-cli-tutoriel-gratuit-en-12-etapes-2026-1
title: "Vérifier les versions actuelles"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Microsoft"]
dates: []
keywords: ["agent", "apache", "claude", "copilot", "gemini", "mcp", "open source", "sandbox"]
source: docs/RAG/collect-261001-ia-llm/gemini-cli-tutoriel-gratuit-en-12-etapes-2026.md
source_anchor: ""
source_lines: [1, 63]
sha256: 05db64169dc4f26897fc4dae5249f273f8a5fe9937506e8929fad19514f71f79
---

# Vérifier les versions actuelles

Et si l’assistant de codage le plus puissant de 2026 tenait dans votre terminal ? C’est la promesse de **Gemini CLI**, l’agent IA open source lancé par Google le 25 juin 2025 sous licence Apache 2.0. Un an plus tard, jour pour jour, le projet a publié sa version **v0.49.0** sur GitHub (25 juin 2026) et s’est imposé comme l’une des portes d’entrée les plus suivies vers l’IA agentique : une fenêtre de contexte d’un million de tokens, un accès à la famille de modèles Gemini 3, et un quota qui a longtemps atteint 1 000 requêtes gratuites par jour avec un simple compte Google – avant que Google ne fasse évoluer les conditions d’accès mi-2026, comme nous le détaillons plus loin.

Ce tutoriel vous guide pas à pas, en 12 étapes et environ 40 minutes, pour installer, authentifier, configurer et exploiter Gemini CLI comme un développeur professionnel. Vous apprendrez à connecter des serveurs MCP, à écrire un fichier de contexte `GEMINI.md`, à automatiser vos scripts et à sécuriser l’exécution avec le sandbox. À la fin, vous disposerez d’un projet complet : un agent capable d’auditer et de documenter n’importe quel dépôt de code. Guide à jour au 03 juin 2026.

## Qu’est-ce que Gemini CLI et pourquoi l’adopter en 2026 ?

Gemini CLI est un **agent IA en ligne de commande** qui apporte la puissance des modèles Gemini directement dans votre terminal. Contrairement à un simple chatbot, il fonctionne selon une boucle agentique de type « raisonner puis agir » : il lit vos fichiers, exécute des commandes shell, interroge le web, écrit du code et vérifie ses propres résultats, le tout sous votre supervision. C’est cette autonomie encadrée qui distingue un agent d’un autocomplète classique.

Trois arguments expliquent son adoption fulgurante. D’abord, le **code source est ouvert** (Apache 2.0) : n’importe quel développeur peut inspecter le fonctionnement de l’agent, auditer sa sécurité et contribuer. Ensuite, le **quota généreux** a longtemps été le plus large du marché : dès le lancement, la connexion avec un compte Google personnel donnait 60 requêtes par minute et 1 000 requêtes par jour, sans carte bancaire – un accès gratuit qui a pris fin le 18 juin 2026, Google orientant désormais les comptes personnels vers les formules payantes Google AI Pro et Ultra, la première facturée environ 19,99 $/mois pour 1 500 requêtes par jour selon Comparly.ai (juillet 2026). Enfin, la **fenêtre de contexte d’un million de tokens** permet de charger des dépôts entiers, une documentation volumineuse ou plusieurs fichiers de logs sans découpage manuel.

Pour les équipes françaises et européennes, Gemini CLI présente un intérêt supplémentaire : la possibilité de le brancher sur **Vertex AI** et de choisir des régions de traitement au sein de l’UE, un point clé pour les organisations soumises au RGPD ou à des exigences de souveraineté des données. Là où un abonnement propriétaire impose son infrastructure, la nature ouverte de l’outil laisse le contrôle au développeur.

Face à la concurrence, Gemini CLI se positionne comme l’option « terminal-first » de Google, aux côtés d’alternatives comme Claude Code d’Anthropic ou GitHub Copilot. Le tableau ci-dessous résume ce positionnement. Chaque outil a ses forces : le choix dépend de votre modèle de facturation, de votre écosystème et de vos exigences de confidentialité.

| Critère | Gemini CLI | Claude Code | GitHub Copilot | 
|---|---|---|---|
| Éditeur |  | Anthropic | GitHub / Microsoft | 
| Licence | Open source (Apache 2.0) | Propriétaire | Propriétaire | 
| Interface | Terminal | Terminal | IDE + terminal (Copilot CLI) | 
| Palier gratuit | 1 000 requêtes/jour (compte Google) | Limité / essai | Palier gratuit plafonné | 
| Contexte | 1 M tokens | Étendu | Selon le modèle | 
| Support MCP | Oui, natif | Oui | Oui | 

Si vous hésitez encore entre écosystèmes, notre comparatif Claude Code vs Cursor et notre tutoriel Claude Code en 12 étapes apportent un éclairage complémentaire. Gemini CLI reste toutefois le point d’entrée le plus accessible pour tester l’IA agentique sans engagement financier.

## Prérequis et versions nécessaires

Avant de lancer l’installation, assurez-vous de réunir les prérequis suivants. Gemini CLI est distribué sous forme de paquet npm ; il exige donc un environnement Node.js à jour. La contrainte la plus stricte, inscrite dans le fichier `package.json` officiel, est **Node.js version 20.0.0 ou supérieure**. Nous recommandons une version LTS active (Node 22 ou Node 24) pour bénéficier des correctifs de sécurité et de meilleures performances.

| Composant | Version minimale | Recommandé (juin 2026) | Rôle | 
|---|---|---|---|
| Node.js | 20.0.0 | 22 LTS ou 24 LTS | Environnement d’exécution | 
| npm | 10.x | Fourni avec Node LTS | Installation du paquet | 
| Système | Windows 10, macOS 12, Linux | Windows 11 / WSL 2, macOS 14+, Ubuntu 22.04+ | Plateforme hôte | 
| Compte Google | Gratuit | Compte personnel ou Workspace | Authentification OAuth | 
| Docker / Podman | Optionnel | Dernière version stable | Sandbox (isolation des outils) | 
| Git | Optionnel | 2.40+ | Filtrage des fichiers ignorés | 

Côté authentification, deux chemins s’offrent à vous : un **compte Google** (méthode par défaut, la plus simple) ou une **clé API Gemini** obtenue sur Google AI Studio. Nous couvrirons les deux. Aucune carte bancaire n’est requise pour démarrer. Enfin, pour le mode sandbox – que nous verrons à l’étape 11 – un moteur de conteneurs comme Docker ou Podman est nécessaire ; il reste facultatif tant que vous ne demandez pas d’isolation renforcée. Si vous travaillez déjà avec des modèles locaux, notre guide Ollama : LLM en local montre une approche 100 % hors-ligne complémentaire.

## Étape 1 – Installer Node.js et vérifier l’environnement

Gemini CLI refusera de démarrer avec une version de Node.js antérieure à la 20. Commencez donc par contrôler votre installation. Ouvrez un terminal et exécutez les commandes de vérification. Si `node -v` renvoie une version inférieure à 20 (ou une erreur « command not found »), installez une version LTS via le gestionnaire de versions `nvm`, qui évite les problèmes de permissions liés à l’installation globale.

```
# Vérifier les versions actuelles
node -v    # doit afficher v20.x ou supérieur
npm -v
# Installer nvm (Linux / macOS), puis Node 22 LTS
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.1/install.sh | bash
# Recharger le shell, puis :
nvm install 22
nvm use 22
nvm alias default 22
# Contrôle final
node -v    # v22.x.x
```
Sous Windows, deux options fiables : installer Node.js LTS depuis le site officiel nodejs.org, ou travailler dans **WSL 2** (sous-système Windows pour Linux) et suivre les instructions Linux ci-dessus. WSL 2 est souvent l’expérience la plus fluide, car Gemini CLI exécute des commandes shell POSIX. Une fois `node -v` confirmé à la version 22 ou 24, votre environnement est prêt.

## Étape 2 – Installer Gemini CLI (npx, npm, Homebrew)

Il existe quatre façons d’installer ou de lancer Gemini CLI. La plus rapide pour un premier essai est `npx`, qui exécute le paquet sans installation permanente. Pour un usage quotidien, préférez l’installation globale via npm, ou Homebrew sur macOS et Linux. Choisissez la méthode adaptée à votre plateforme.

