---
id: collect-261001-ia-llm/ia-llm/ollama-executer-un-llm-en-local-12-etapes-2026-2
title: "macOS (via Homebrew)"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Google", "Meta", "Microsoft", "OpenAI"]
dates: []
keywords: ["attention", "deepseek", "embeddings", "gpu", "llama", "qwen"]
source: docs/RAG/collect-261001-ia-llm/ollama-executer-un-llm-en-local-12-etapes-2026.md
source_anchor: ""
source_lines: [51, 151]
sha256: 0fa0079a2058017f7e4ca04cf0219b9193602aafb8c7eaba4878002979830c2b
---

# macOS (via Homebrew)

L’installation d’Ollama tient en une commande. Choisissez la ligne correspondant à votre système d’exploitation. Sur macOS, Homebrew est le plus propre ; sur Linux, le script officiel détecte automatiquement votre GPU et installe le service systemd ; sur Windows, le gestionnaire de paquets `winget` évite tout téléchargement manuel.

```
# macOS (via Homebrew)
brew install ollama
# Linux (script officiel – relisez-le avant exécution)
curl -fsSL https://ollama.com/install.sh | sh
# Windows (via winget, dans PowerShell)
winget install Ollama.Ollama
```
Sur Linux, le script crée un utilisateur système `ollama` et démarre un service en arrière-plan qui écoute sur le port 11434. Sur macOS et Windows, une application de barre des tâches – actuellement en version **3.7.4**, publiée le 18 août 2026 après la 3.7.0 du 27 juillet 2026 selon le changelog Ollaman – lance le serveur au démarrage. Vérifiez immédiatement que tout est en place avec la commande de version.

```
ollama --version
# Sortie attendue :
# ollama version is 0.30.11
```
Si vous préférez vérifier que le serveur répond, une simple requête HTTP confirme qu’Ollama écoute. Ce point sera central dès l’étape consacrée à l’API.

```
curl http://localhost:11434
# Réponse : Ollama is running
```
Sous Linux, si la commande renvoie une erreur de connexion, démarrez le service avec `sudo systemctl start ollama` puis activez son lancement automatique via `sudo systemctl enable ollama`. Vous pouvez télécharger directement les binaires depuis la page officielle de téléchargement si votre environnement bloque les scripts.

## Étape 2 – Télécharger et lancer votre premier modèle

Place à la pratique. La commande `ollama run` télécharge le modèle s’il est absent, le charge en mémoire puis ouvre une session de chat interactive. Pour un premier essai léger et rapide, Llama 3.2 (3 milliards de paramètres) est idéal : moins de 2,5 Go à télécharger et des réponses quasi instantanées même sans GPU.

```
ollama run llama3.2
# Le modèle se télécharge (barre de progression), puis :
# >>> Bonjour, peux-tu te présenter en une phrase ?
# Je suis un assistant IA exécuté localement via Ollama,
# prêt à répondre à vos questions sans connexion externe.
#
# >>> /bye   (pour quitter la session)
```
Dans la session interactive, plusieurs commandes spéciales commencent par une barre oblique : `/?` affiche l’aide, `/bye` quitte, `/clear` efface le contexte de la conversation, et `/show info` détaille le modèle chargé. Vous pouvez aussi coller un texte multi-lignes en l’entourant de triples guillemets.

Si vous voulez seulement télécharger un modèle sans l’exécuter – par exemple pour préparer plusieurs modèles à l’avance – utilisez `pull`. Précisez une étiquette (tag) après les deux-points pour choisir une taille précise.

```
ollama pull qwen3:8b          # Qwen 3, version 8 milliards
ollama pull deepseek-r1:8b    # Modèle de raisonnement
ollama pull gemma3:4b         # Google Gemma 3, compact
ollama pull nomic-embed-text  # Modèle d'embeddings (pour le RAG)
```
Les modèles sont stockés dans `~/.ollama/models` sur macOS et Linux, et dans le profil utilisateur sous Windows. Vous pouvez déplacer cet emplacement avec la variable d’environnement `OLLAMA_MODELS`, utile si votre disque système est petit.

## Étape 3 – Maîtriser les commandes CLI essentielles d’Ollama

La ligne de commande d’Ollama est volontairement minimaliste, y compris dans la version **0.33.2 du 27 août 2026** – la plus récente recensée par LocalAIMaster au moment d’écrire ces lignes. Une poignée de verbes couvre 95 % des usages quotidiens. Le tableau ci-dessous regroupe les commandes que vous utiliserez le plus, de la gestion des modèles à la supervision de la mémoire.

| Commande | Rôle | 
|---|---|
| `ollama run MODELE` | Lancer une session interactive (télécharge si besoin) | 
| `ollama pull MODELE` | Télécharger un modèle sans l’exécuter | 
| `ollama list` | Lister les modèles installés et leur taille | 
| `ollama ps` | Voir les modèles actuellement chargés en mémoire | 
| `ollama show MODELE` | Afficher paramètres, template et licence | 
| `ollama stop MODELE` | Décharger un modèle de la mémoire | 
| `ollama rm MODELE` | Supprimer un modèle du disque | 
| `ollama cp SRC DST` | Copier ou dériver un modèle existant | 
| `ollama serve` | Démarrer le serveur manuellement (premier plan) | 

Deux commandes méritent une attention particulière. `ollama ps` indique non seulement quels modèles sont chargés, mais aussi s’ils tournent sur CPU ou GPU et combien de mémoire ils consomment – précieux pour diagnostiquer une lenteur. `ollama show`, lui, révèle le *template* de prompt et la fenêtre de contexte par défaut, deux paramètres déterminants pour la qualité des réponses.

```
ollama list
# NAME              ID            SIZE      MODIFIED
# llama3.2:latest   a80c4f17acd5  2.0 GB    il y a 2 minutes
# qwen3:8b          500a1f067a9f  5.2 GB    il y a 5 minutes
ollama ps
# NAME              PROCESSOR     CONTEXT   UNTIL
# llama3.2:latest   100% GPU      4096      4 minutes restantes
```
Par défaut, Ollama décharge un modèle de la mémoire après cinq minutes d’inactivité pour libérer des ressources. Nous verrons à l’étape performances comment ajuster ce comportement avec `OLLAMA_KEEP_ALIVE`.

## Étape 4 – Choisir le bon modèle : Llama, DeepSeek, Qwen, Gemma, gpt-oss

La bibliothèque officielle d’Ollama propose des centaines de modèles. Inutile de tous les essayer : quelques familles couvrent la quasi-totalité des besoins. Le choix dépend de votre tâche (conversation, code, raisonnement, vision, embeddings) et de votre budget mémoire.

| Modèle | Éditeur | Point fort | Tag conseillé | 
|---|---|---|---|
| Llama 3.2 / 3.1 | Meta | Généraliste polyvalent | llama3.2 / llama3.1:8b | 
| DeepSeek-R1 | DeepSeek | Raisonnement étape par étape | deepseek-r1:8b | 
| Qwen 3 | Alibaba | Multilingue, appel d’outils | qwen3:8b | 
| Gemma 3 |  | Compact et efficace | gemma3:4b | 
| Phi-4 | Microsoft | Excellent rapport taille/qualité | phi4 | 
| gpt-oss | OpenAI | Modèle ouvert, agentique | gpt-oss:20b | 
| Qwen2.5-Coder | Alibaba | Génération de code | qwen2.5-coder | 
| LLaVA | Communauté | Vision (analyse d’images) | llava | 
| nomic-embed-text | Nomic | Embeddings pour RAG | nomic-embed-text | 

Quelques repères pour décider. Pour un **assistant généraliste** francophone, Llama 3.2 ou Qwen 3 en 8B offrent le meilleur équilibre : Llama 3.1, disponible en 8B, 70B et 405B avec une fenêtre de contexte de 128 000 tokens et l’appel d’outils (état d’août 2026), ne pèse qu’environ 4,7 Go en quantification Q4 pour sa version 8B et se contente de 8 à 16 Go de RAM, ce qui en fait mi-2026 l’un des modèles compacts les plus recommandés. Pour des tâches qui exigent de la **logique** (mathématiques, planification, débogage), un modèle de raisonnement comme DeepSeek-R1 « réfléchit » avant de répondre et gagne nettement en exactitude. Pour le **code**, Qwen2.5-Coder rivalise avec des assistants payants, et la nouvelle famille Qwen 3.6 (27B et 35B, avec un contexte étendu à 256 000 tokens) s’impose début août 2026 comme l’option la plus récente pour la génération de code. Et pour analyser des **images**, LLaVA accepte un chemin de fichier directement dans le prompt.

Notez le suffixe d’étiquette : `:8b`, `:14b`, `:latest`. Sans étiquette, Ollama télécharge la version par défaut, souvent un format intermédiaire. Précisez toujours la taille pour contrôler votre empreinte mémoire.

## Étape 5 – Interroger l’API REST d’Ollama (port 11434)

