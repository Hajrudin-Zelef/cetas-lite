---
id: collect-261001-ia-llm/ia-llm/gemini-cli-tutoriel-gratuit-en-12-etapes-2026-2
title: "Vérifier les versions actuelles"
domain: ia-llm
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["agent", "exploit", "gemini"]
source: docs/RAG/collect-261001-ia-llm/gemini-cli-tutoriel-gratuit-en-12-etapes-2026.md
source_anchor: ""
source_lines: [64, 157]
sha256: 171f6c4ef83b0e1f9e371e17768d8e0c9f2e670f5a7c094d2ca7e5089437643d
---

# Vérifier les versions actuelles

```
# Méthode 1 – Sans installation (essai immédiat)
npx @google/gemini-cli
# Méthode 2 – Installation globale via npm (usage quotidien)
npm install -g @google/gemini-cli
# Méthode 3 – Homebrew (macOS / Linux)
brew install gemini-cli
# Méthode 4 – MacPorts
sudo port install gemini-cli
# Vérifier l'installation
gemini --version
```
Une fois installé globalement, la commande `gemini` devient disponible partout dans votre terminal. Si `gemini --version` renvoie un numéro de version, l’installation est réussie. Gemini CLI propose trois **canaux de publication** : *stable* (recommandé pour la majorité des utilisateurs), *preview* (fonctionnalités en avance de phase) et *nightly* (build quotidien, pour tester les nouveautés). Par défaut, npm installe le canal stable.

**Astuce professionnelle :** si vous rencontrez une erreur de permission `EACCES` lors de l’installation globale sous Linux ou macOS, ne la contournez pas avec `sudo`. Utilisez plutôt `nvm` (étape 1), qui installe les paquets dans votre répertoire utilisateur et supprime définitivement ce type d’erreur. C’est le piège numéro un des débutants.

## Étape 3 – Premier lancement et authentification Google

Lancez maintenant l’agent en tapant simplement `gemini` dans un dossier de projet. Au premier démarrage, l’assistant vous propose de choisir un thème visuel, puis une méthode d’authentification. Sélectionnez **« Se connecter avec Google »** (Login with Google) : votre navigateur s’ouvre automatiquement, vous vous authentifiez avec votre compte Google, et la CLI stocke localement un jeton OAuth qu’elle rafraîchit ensuite toute seule.

```
# Se placer dans un projet, puis lancer l'agent
cd mon-projet
gemini
# Au premier lancement :
#  1. Choix du thème (ex. Default Dark)
#  2. Choix de l'authentification -> "Login with Google"
#  3. Le navigateur s'ouvre pour valider le compte
#  4. Retour au terminal : l'agent est prêt
```
Jusqu’au 18 juin 2026, cette méthode débloquait un **palier entièrement gratuit** de 60 requêtes par minute et 1 000 requêtes par jour. Depuis cette date, Google réserve ce niveau de service aux abonnés **Google AI Pro** (19,99 $/mois pour 1 500 requêtes par jour) ou **Ultra**, avec accès à la famille de modèles Gemini (routage automatique vers Gemini 3, principalement Flash pour les tâches courantes) et à la fenêtre d’un million de tokens. Pour un usage individuel intensif, l’abonnement Pro reste largement suffisant. Aucune configuration de clé n’est nécessaire dans ce mode, seule la connexion à un compte Google abonné.

Vous voyez apparaître l’invite de l’agent. Testez-la immédiatement avec une question simple, par exemple : *« Résume l’arborescence de ce projet et identifie le point d’entrée principal. »* L’agent va lister les fichiers, lire les plus pertinents, puis vous répondre. Si vous travaillez sur une machine distante en SSH sans navigateur, l’ouverture automatique échouera : dans ce cas, utilisez plutôt l’authentification par clé API décrite à l’étape suivante.

## Étape 4 – Authentification par clé API et Vertex AI

La connexion Google convient à la plupart des cas, mais trois situations imposent une **clé API** : les serveurs sans navigateur (CI/CD, machines distantes), le besoin de quotas payants plus élevés, et l’accès garanti aux modèles Pro. Rendez-vous sur Google AI Studio pour générer une clé, puis exportez-la dans la variable d’environnement `GEMINI_API_KEY`.

```
# Exporter la clé pour la session courante
export GEMINI_API_KEY="votre_cle_api_ici"
# Persistance recommandée : un fichier .env chargé automatiquement
# Ordre de priorité : ./.gemini/.env  >  ~/.gemini/.env  >  ~/.env
mkdir -p .gemini
echo 'GEMINI_API_KEY=votre_cle_api_ici' > .gemini/.env
# Variante entreprise : Vertex AI (régions UE possibles)
export GOOGLE_API_KEY="votre_cle_api_ici"
export GOOGLE_GENAI_USE_VERTEXAI=true
```
Gemini CLI charge automatiquement les fichiers `.env` selon un ordre de priorité précis : d’abord `./.gemini/.env` (spécifique au projet), puis `~/.gemini/.env`, puis `~/.env`. Cette hiérarchie permet d’isoler des clés différentes par projet. Le palier gratuit de la clé API reste plus restreint que l’ancien accès Google – environ 10 requêtes par minute et 250 par jour, limité aux modèles Flash – mais l’activation de la facturation débloque des débits bien supérieurs : Gemini 3.1 Pro Preview est ainsi facturé environ 2,00 $ par million de tokens en entrée (jusqu’à 200k tokens) depuis juillet 2026 selon Tembo, tandis que la formule Code Assist Enterprise grimpe à 54 $ par utilisateur et par mois pour 2 000 requêtes quotidiennes.

| Méthode | Débit gratuit | Quota gratuit / jour | Modèles | Cas d’usage | 
|---|---|---|---|---|
| Connexion Google (OAuth) | 60 req/min | 1 000 requêtes | Famille Gemini (routage auto) | Poste de travail individuel | 
| Clé API – palier gratuit | 10 req/min | 250 requêtes | Flash uniquement | Serveurs, scripts légers | 
| Clé API – facturation activée | Élevé | À l’usage (par token) | Gemini 3 Pro & Flash | Production, CI/CD | 
| Vertex AI | Selon quota GCP | Selon contrat | Gemini 3 Pro & Flash | Entreprise, RGPD / UE | 

**Sécurité :** ne validez jamais un fichier `.env` contenant votre clé dans Git. Ajoutez `.gemini/.env` et `.env` à votre `.gitignore` dès la création du projet. Une clé API exposée sur un dépôt public peut être exploitée en quelques minutes.

## Étape 5 – Vos premières commandes et le mode agent

Vous dialoguez avec Gemini CLI en langage naturel. Mais trois syntaxes spéciales décuplent sa puissance. Le symbole `@` injecte le contenu d’un fichier dans votre requête (avec filtrage des fichiers ignorés par Git). Le point d’exclamation `!` exécute une commande shell ou bascule en mode shell interactif – un mode devenu activé par défaut depuis la version **v0.9.0**, publiée le 15 octobre 2025. Enfin, les **commandes slash** (préfixées par `/`) pilotent la session elle-même.

```
# Poser une question en langage naturel
> Explique le rôle du fichier src/server.js et repère les failles potentielles.
# Injecter un fichier précis dans le contexte avec @
> @src/auth.js  Réécris cette fonction pour utiliser async/await.
# Injecter tout un dossier
> @src/  Dresse la liste des dépendances circulaires.
# Exécuter une commande shell avec !
> !npm test
# Basculer en mode shell (les commandes suivantes sont shell)
> !
# Lister les outils disponibles
> /tools
```
Lorsque l’agent décide d’utiliser un outil susceptible de modifier votre système – écrire un fichier, lancer une commande – il vous demande une **confirmation** avant d’agir. Vous pouvez approuver une fois, approuver pour toute la session, ou refuser. Ce garde-fou est essentiel : il vous garde aux commandes. Voici un exemple de sortie typique lorsqu’on demande à l’agent de créer un fichier.

```
> Crée un fichier README.md décrivant ce projet.
✦ Je vais analyser le projet puis rédiger le README.
  [outil] ReadFolder  .          -> 14 éléments
  [outil] ReadFile    package.json
  [outil] WriteFile   README.md   (nouveau fichier)
  Confirmer l'écriture de README.md ? [O]ui / [T]oujours / [N]on
> O
✦ README.md créé (48 lignes). Souhaitez-vous que j'ajoute un badge de build ?
```
Prenez le réflexe de taper `/help` pour afficher l’aide contextuelle et `/tools` pour voir la liste des outils intégrés : recherche Google, lecture et écriture de fichiers, exécution shell, récupération de pages web et mémoire. Ces outils sont le cœur de l’autonomie de l’agent.

## Étape 6 – Créer un fichier de contexte GEMINI.md

