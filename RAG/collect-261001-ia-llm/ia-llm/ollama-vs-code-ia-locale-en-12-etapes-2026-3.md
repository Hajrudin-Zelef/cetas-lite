---
id: collect-261001-ia-llm/ia-llm/ollama-vs-code-ia-locale-en-12-etapes-2026-3
title: "macOS et Linux"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Apple", "Nvidia"]
dates: []
keywords: ["agent", "attention", "embeddings", "flash attention", "gpu", "llama", "llama.cpp", "mai", "nvidia", "qwen"]
source: docs/RAG/collect-261001-ia-llm/ollama-vs-code-ia-locale-en-12-etapes-2026.md
source_anchor: ""
source_lines: [129, 216]
sha256: a938166e047efd868eb6c822e069c057564565afb521f112c2ba490f5123dd2f
---

# macOS et Linux

Ouvrez le fichier de configuration de Continue. Depuis la version 1.x, il s’agit d’un fichier YAML situé dans `~/.continue/config.yaml` (accessible aussi depuis l’icône d’engrenage du panneau Continue dans VS Code). Remplacez son contenu par une configuration qui pointe vers votre serveur Ollama local :

```
name: Assistant Local
version: 1.0.0
schema: v1
models:
  - name: Chat Qwen Coder
    provider: ollama
    model: qwen2.5-coder:7b
    apiBase: http://localhost:11434
    roles:
      - chat
      - edit
  - name: Autocomplétion Qwen Coder
    provider: ollama
    model: qwen2.5-coder:7b
    apiBase: http://localhost:11434
    roles:
      - autocomplete
  - name: Embeddings Local
    provider: ollama
    model: nomic-embed-text
    apiBase: http://localhost:11434
    roles:
      - embed
context:
  - provider: codebase
  - provider: file
  - provider: terminal
  - provider: diff
```
Enregistrez le fichier. Continue recharge automatiquement la configuration et affiche les modèles disponibles dans le sélecteur du panneau de chat, accessible par le raccourci Ctrl+L (Cmd+L sur macOS). La documentation officielle de Continue reste accessible malgré le passage du dépôt en lecture seule et détaille l’ensemble des champs disponibles dans ce schéma, notamment pour ajouter des modèles supplémentaires ou changer un rôle sans tout réécrire. Le champ `roles` est celui qui détermine où chaque modèle intervient : un même modèle peut très bien cumuler plusieurs rôles, ou au contraire être dédié à un seul, ce qui devient utile dès que vous ajoutez un deuxième modèle plus gros réservé au chat.

## Étape 7 : Configurer l’autocomplétion en temps réel

Le rôle `autocomplete` défini à l’étape précédente active déjà la complétion automatique. Pour l’ajuster à votre matériel, ajoutez une section dédiée dans le même fichier `config.yaml` :

```
tabAutocompleteOptions:
  debounceDelay: 300
  maxPromptTokens: 1024
  disableInFiles:
    - "*.md"
    - "*.txt"
```
Le paramètre `debounceDelay` fixe le délai, en millisecondes, entre la fin de la frappe et l’envoi de la requête au modèle. Sur une machine sans GPU dédié, montez cette valeur à 500 ou 800 pour éviter de saturer le processeur à chaque caractère tapé. Le paramètre `maxPromptTokens` limite la taille du contexte envoyé au modèle pour l’autocomplétion : une valeur plus basse accélère la réponse au prix d’un contexte plus restreint sur le fichier en cours.

## Étape 8 : Configurer le chat, le contexte et les embeddings

Le modèle d’embeddings `nomic-embed-text` ajouté à l’étape 6 permet à Continue d’indexer votre dépôt de code pour répondre à des questions qui portent sur l’ensemble du projet, pas seulement sur le fichier ouvert. Téléchargez-le s’il n’est pas déjà présent :

`ollama pull nomic-embed-text`
Une fois ce modèle en place, ouvrez le panneau de chat de Continue et tapez `@codebase` suivi de votre question pour interroger l’ensemble du projet plutôt qu’un seul fichier. Les fournisseurs de contexte ajoutés dans la configuration (`@file`, `@terminal`, `@diff`) permettent respectivement d’attacher un fichier précis, la sortie du terminal actif, ou les modifications en cours dans Git à votre question. Sur un premier import, l’indexation d’un gros dépôt peut prendre plusieurs minutes : Continue affiche une barre de progression dans la barre d’état de VS Code pendant ce calcul.

## Étape 9 : Test grandeur nature du projet complet

Il est temps de vérifier que toute la chaîne fonctionne ensemble. Créez un nouveau dossier de projet, ouvrez-le dans VS Code, et créez un fichier `calculatrice.py` vide. Commencez à taper une signature de fonction :

`def diviser(a, b):`
Après une courte pause, une suggestion grisée doit apparaître automatiquement, proposant par exemple une gestion de la division par zéro. Appuyez sur Tab pour l’accepter. Ouvrez ensuite le panneau de chat (Ctrl+L) et posez une question qui mobilise le contexte du projet entier : “@codebase quelles fonctions de ce dossier ne gèrent pas les exceptions ?”. Le modèle doit lister les fonctions concernées en se basant sur l’index construit à l’étape précédente. Terminez en sélectionnant un bloc de code dans l’éditeur, clic droit, puis “Continue : Edit”, et demandez une refactorisation, par exemple l’ajout d’une vérification de type. Ajoutez un dernier test pour valider la génération de tests unitaires : demandez au chat “@file calculatrice.py génère des tests unitaires avec pytest pour la fonction diviser, y compris le cas de la division par zéro”. Le modèle doit produire un fichier de test complet avec au moins deux cas, le cas normal et le cas d’exception. Si ces quatre interactions, complétion, chat contextuel, édition et génération de tests, fonctionnent sans message d’erreur, l’environnement complet est opérationnel et prêt pour un usage quotidien.

## Étape 10 : Optimiser les performances (GPU, quantification, Flash Attention)

Si les réponses restent lentes malgré un GPU compatible, vérifiez d’abord qu’Ollama détecte bien la carte graphique :

```
ollama ps
# La colonne PROCESSOR doit indiquer "100% GPU"
# et non "100% CPU"
```
Sur les GPU Nvidia de génération récente (RTX 5090, 5080, 5070 notamment), Ollama active automatiquement l’accélération CUDA et le Flash Attention. Sur du matériel plus ancien, cette optimisation peut nécessiter une activation manuelle via une variable d’environnement :

`OLLAMA_FLASH_ATTENTION=1 ollama serve`
Si la VRAM disponible reste juste, privilégiez une version quantifiée plus agressive du même modèle (suffixe `q4_K_M` plutôt que `q8_0`, par exemple), au prix d’une légère perte de précision sur les réponses les plus complexes. Sur les puces Apple Silicon, Ollama s’appuie sur le moteur MLX en complément de llama.cpp : introduit en aperçu dans la version 0.19 dès mars 2026, ce moteur est passé en version stable avec la 0.30 le 13 mai 2026, doublant la vitesse d’inférence sur Apple Silicon selon Andrew.ooo, avant d’être encore affiné dans la 0.30.10 ; les notes de version du 14 août 2026 signalent par ailleurs, selon ReleaseBot, une nouvelle accélération de Qwen3.5 sur GPU Apple grâce à des moteurs MLX et llama.cpp mis à jour.

## Étape 11 : Sécuriser l’environnement et les données de code

Un environnement local n’est pas automatiquement un environnement sûr. Trois réglages simples réduisent le risque de fuite accidentelle. D’abord, ajoutez le dossier de configuration de Continue à votre `.gitignore` global pour éviter de committer par erreur des clés d’API si vous configurez aussi des modèles cloud en complément :

```
# Dans ~/.gitignore_global ou le .gitignore du projet
.continue/
```
Ensuite, si votre machine est accessible sur un réseau partagé, n’exposez pas le port 11434 au-delà de `localhost` sans authentification : par défaut, Ollama n’écoute que sur l’interface locale, ce qui suffit pour l’usage décrit dans ce tutoriel. Enfin, si votre configuration mélange un modèle local pour le code sensible et un modèle cloud pour d’autres tâches, vérifiez à chaque session que le rôle `chat` par défaut pointe bien vers le modèle local avant d’ouvrir un fichier confidentiel, pour éviter d’envoyer par erreur du code protégé vers un fournisseur externe.

## Étape 12 : Automatiser avec le mode agent

