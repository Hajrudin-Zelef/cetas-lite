---
id: collect-261001-ia-llm/ia-llm/tutoriel-ollama-2026-llm-local-en-13-etapes-2
title: "Vérifier la version du pilote NVIDIA"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Meta", "Poolside"]
dates: []
keywords: ["benchmark", "gpu", "llama", "mistral", "moe", "open source", "parameters", "qwen"]
source: docs/RAG/collect-261001-ia-llm/tutoriel-ollama-2026-llm-local-en-13-etapes.md
source_anchor: ""
source_lines: [47, 135]
sha256: 15390909fdc60c62c2739c8bd9e49af16d783b0769e5652f0e1b57097800ea77
---

# Vérifier la version du pilote NVIDIA

```
# --- LINUX (Ubuntu, Debian, Fedora, Arch) ---
curl -fsSL https://ollama.com/install.sh | sh
# Vérifier l'installation
ollama --version
# Attendu : ollama version is 0.6.2
# Activer le service systemd
sudo systemctl enable ollama
sudo systemctl start ollama
# --- macOS ---
# Option A : Homebrew
brew install ollama
# Option B : Application native
# Télécharger https://ollama.com/download/Ollama-darwin.zip
# --- WINDOWS 11 (PowerShell admin) ---
irm https://ollama.com/install.ps1 | iex
```
Après installation, le serveur Ollama écoute par défaut sur `http://localhost:11434`. Pour vérifier qu’il fonctionne, lancez `curl http://localhost:11434` : la réponse doit être « Ollama is running ». Si vous voyez une erreur de connexion, démarrez manuellement le serveur avec `ollama serve` dans un terminal. Sur Linux, vérifiez le statut systemd avec `systemctl status ollama` : le service doit être actif et activé.

**Piège fréquent #1** : sur Windows, l’antivirus Defender peut bloquer le téléchargement du binaire CUDA pendant l’installation. Désactivez temporairement la protection en temps réel ou ajoutez `%LOCALAPPDATA%\Programs\Ollama` aux exclusions. Sur Linux, si vous utilisez ZFS ou Btrfs pour `/usr`, le service peut échouer à démarrer avec une erreur de permission ; déplacez le binaire vers `/opt/ollama` et ajustez le fichier service.

## Étape 2 : Télécharger votre premier modèle avec ollama pull

Une fois le serveur opérationnel, l’étape suivante consiste à télécharger un modèle depuis la bibliothèque officielle. Ollama utilise un système de tags inspiré de Docker : `llama3.3:70b` spécifie la version 70 milliards de paramètres, tandis que `llama3.3:70b-instruct-q4_K_M` précise la quantification. Les nouveautés se propagent vite : le tag `qwen3.8:27b`, sorti mi-août 2026, a déjà cumulé environ **501 800 pulls** lors de sa première semaine de disponibilité, selon les chiffres publiés par MorphLLM — un signe de l’appétit des développeurs pour la gamme Qwen, dont la version **Qwen3.6 27B** atteint **84 % sur le benchmark MMLU** tout en tenant dans environ **17 Go de RAM** en quantification Q4_K_M, d’après le classement de juillet 2026 de PromptQuorum. D’autres arrivées confirment cette dynamique : le modèle MoE **Poolside Laguna XS 2.1** (33B, avec un routage actif de 3B) a atteint **70,9 % sur SWE-bench** dès son intégration à Ollama le **2 juillet 2026**, toujours selon PromptQuorum. La commande `ollama pull` télécharge le modèle, vérifie son hash SHA-256, puis le décompresse dans `~/.ollama/models` (Linux/macOS) ou `%USERPROFILE%\.ollama\models` (Windows).

```
# Modèles populaires en 2026 (vérifiés sur ollama.com/library)
# Pour débuter — 4,7 Go, 8 Go RAM suffisent
ollama pull llama3.1:8b
# Modèle français-friendly — 4,1 Go
ollama pull mistral:7b-instruct-v0.3
# Code & raisonnement — 4,1 Go
ollama pull qwen3:7b
# Petit, rapide, multilingue — 815 Mo
ollama pull gemma3:1b
# Top tier — 40 Go, GPU 48 Go requis
ollama pull llama3.3:70b
# Vérifier les modèles installés
ollama list
# NAME              ID            SIZE      MODIFIED
# llama3.1:8b       42182419e950  4.7 GB    2 minutes ago
# mistral:7b        2ae6f6dd7a3d  4.1 GB    5 minutes ago
```
Le téléchargement peut être long sur une connexion ADSL : prévoyez 8 à 15 minutes pour un modèle 7B sur une fibre 100 Mb/s, et jusqu’à 2 heures pour Llama 3.3 70B (40 Go). Ollama gère le téléchargement par couches comme Docker, ce qui signifie que si vous tirez plusieurs variantes du même modèle de base, les couches communes ne sont téléchargées qu’une fois. Cela économise typiquement 30 à 50 % d’espace disque pour qui expérimente avec plusieurs quantifications.

**Piège fréquent #2** : si votre dossier `~/.ollama` est sur un disque chiffré (LUKS, FileVault), les performances de chargement initial peuvent chuter de 40 à 60 %. Configurez `OLLAMA_MODELS` pour pointer vers un SSD NVMe non chiffré dédié aux modèles : `export OLLAMA_MODELS=/mnt/nvme/ollama`. Cette variable doit être définie avant le démarrage du service, pas dans une session interactive.

## Étape 3 : Lancer votre premier prompt avec ollama run

La commande `ollama run` charge le modèle en mémoire (RAM ou VRAM selon votre GPU) et ouvre une session de chat interactive. C’est la manière la plus rapide de tester un modèle. Pour un usage scripté, on utilisera plutôt l’API HTTP que nous verrons à l’étape 5. La session interactive supporte plusieurs commandes spéciales préfixées par `/` : `/bye` pour quitter, `/clear` pour réinitialiser le contexte, `/show` pour afficher les paramètres, et `/set` pour ajuster température, top_p, system prompt en cours de session.

```
# Lancer une session interactive
ollama run llama3.1:8b
>>> Bonjour, présente-toi en français en 2 phrases.
Je suis Llama 3.1, un modèle de langage open source développé par Meta.
Je peux répondre à des questions, rédiger du texte et discuter en plusieurs langues, dont le français.
>>> /set parameter temperature 0.3
Set parameter 'temperature' to '0.3'
>>> /show parameters
Model defined parameters:
temperature                    0.3
top_k                          40
top_p                          0.9
num_ctx                        4096
>>> /bye
# Mode "one-shot" : prompt direct depuis la ligne de commande
ollama run llama3.1:8b "Écris un haïku sur la pluie à Paris"
# Sortie : Pluie sur les toits / Le métro file en silence / Avril déguisé
```
Sur un RTX 4090 (24 Go VRAM), Llama 3.1 8B en Q4_K_M génère typiquement 80 à 110 tokens/seconde après le warm-up initial. Sur un MacBook Pro M3 Max, comptez 45 à 65 tokens/seconde avec le backend Metal 3 classique, mais un benchmark publié en mars 2026 avec la préversion MLX rapporte un débit passant d’**environ 58 à environ 112 tokens/seconde** sur le même type de modèle, selon Local AI Master. Sur un Ryzen 9 7950X sans GPU, le même modèle plafonne à 8-12 tokens/seconde, ce qui reste utilisable pour des prompts courts mais devient pénible pour de la génération longue. Pour mesurer votre propre débit, utilisez la commande `ollama run llama3.1:8b --verbose` qui affiche le temps de génération en fin de chaque réponse.

## Étape 4 : Comprendre les commandes CLI essentielles

Ollama propose une dizaine de commandes principales que tout utilisateur sérieux doit connaître. Au-delà de `pull` et `run`, la gestion quotidienne s’appuie sur `list`, `ps`, `rm`, `show` et `cp`. La commande `ollama ps` est particulièrement utile : elle affiche quels modèles sont actuellement chargés en mémoire et combien de VRAM ils consomment. Par défaut, Ollama décharge un modèle après 5 minutes d’inactivité ; vous pouvez ajuster ce timeout avec la variable `OLLAMA_KEEP_ALIVE`.

| Commande | Description | Exemple | 
|---|---|---|
| `ollama serve` | Démarre le serveur API sur le port 11434 | `OLLAMA_HOST=0.0.0.0:11434 ollama serve` | 
| `ollama pull` | Télécharge un modèle depuis la registry | `ollama pull qwen3:14b` | 
| `ollama run` | Lance une session interactive ou un prompt | `ollama run mistral "Résume ce texte..."` | 
| `ollama list` ou`ls` | Liste les modèles téléchargés | `ollama list` | 
| `ollama ps` | Affiche les modèles chargés en RAM/VRAM | `ollama ps` | 
| `ollama show` | Affiche les détails (paramètres, Modelfile) | `ollama show llama3.1:8b` | 
| `ollama rm` | Supprime un modèle local | `ollama rm mistral:7b` | 
| `ollama cp` | Crée une copie alias d’un modèle | `ollama cp llama3.1 mon-llama` | 
| `ollama create` | Crée un modèle à partir d’un Modelfile | `ollama create assistant -f Modelfile` | 
| `ollama push` | Publie un modèle sur la registry | `ollama push user/mon-modele` | 

