---
id: collect-261001-ia-llm/ia-llm/tutoriel-ollama-2026-llm-local-en-13-etapes-1
title: "Vérifier la version du pilote NVIDIA"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "Apple", "Microsoft", "Nvidia", "OpenAI"]
dates: []
keywords: ["nvidia", "agent", "amd", "attention", "benchmark", "distribution", "flash attention", "gguf", "gpu", "llama", "llama.cpp", "mai"]
source: docs/RAG/collect-261001-ia-llm/tutoriel-ollama-2026-llm-local-en-13-etapes.md
source_anchor: ""
source_lines: [1, 46]
sha256: 556e9777cd6020bc03d16b86a9da0d9fad6bc92a9daae3098efeac4288bbf936
---

# Vérifier la version du pilote NVIDIA

**Publié le 29 avril 2026 — Mise à jour : Ollama v0.32.1 (juillet 2026), patchée en v0.32.14 (août 2026)**

Ollama est devenu en 2026 l’outil de référence pour exécuter des grands modèles de langage (LLM) en local, sans envoyer une seule requête vers le cloud. Selon **TechCrunch** (juillet 2026), Ollama revendique désormais **8,9 millions de développeurs actifs par mois** dans le monde et une adoption chez **85 % des entreprises du Fortune 500** — une dynamique confirmée par le blog officiel d’Ollama qui, en août 2026, annonce à son tour **8,9 millions de développeurs** et un total de **88 millions de dollars levés**. Avec plus de **27 100 recherches mensuelles en France** selon DataForSEO (avril 2026), le runtime cumule désormais le support natif de **Llama 3.3 70B**, **Qwen 3.6 27B** — annoncé à **77,2 % sur le benchmark SWE-bench** en juillet 2026 comme nouvelle référence de la gamme Qwen —, **Gemma 3** (1B, 4B, 12B, 27B) et l’API compatible OpenAI sur le port `11434`. Le classement PromptQuorum de juillet 2026 place d’ailleurs Qwen3.6 27B aux côtés de **Llama 4 Scout 17B** et **Phi-4 14B** parmi les modèles Ollama les plus recommandés du moment. Ce tutoriel pas-à-pas vous guide de l’installation jusqu’au déploiement d’un assistant RAG complet, avec 13 étapes testées sur Ubuntu 24.04 LTS, macOS Sonoma et Windows 11.

Contrairement aux API payantes (OpenAI à 5 $/Mtok input, Anthropic à 3 $/Mtok), exécuter Llama 3.3 70B en local coûte uniquement votre électricité — même si Ollama a lancé en juillet 2026 un **tier cloud hybride facturé au temps GPU**, qui revendiquait déjà **8,9 millions d’utilisateurs actifs mensuels** selon Open Source For You. Ce virage commercial suit une levée de **65 millions de dollars en Series B** bouclée en juillet 2026, portant le total levé par Ollama à **88 millions de dollars** d’après TechCrunch. Depuis la v0.32.0, sortie le **11 juillet 2026** selon la mise à jour de juillet 2026 de PromptQuorum, lancer la commande `ollama` sans argument ouvre directement un **agent interactif** plutôt qu’un simple message d’aide, et la commande `ollama launch` propose désormais une **fenêtre de session de codage de 5 heures** à usage étendu, selon le blog Ollama (mai 2026). Le rythme de publication reste soutenu : la v0.32.1 a suivi le **16 juillet 2026** d’après LocalAIMaster, qui situe désormais la dernière version stable à **v0.33.2**, publiée le **27 août 2026**. Cette cadence accompagne une croissance rapide de la base d’utilisateurs actifs, passée d’environ **1,5 million en mai 2024** à environ **5,0 millions en mai 2026** selon le rapport écosystème de Presenc AI. Pour un développeur français soumis au RGPD ou un consultant qui manipule des données clients, c’est l’option par défaut en 2026.

## Qu’est-ce qu’Ollama et pourquoi l’adopter en 2026

Ollama est un runtime open source écrit en Go qui simplifie radicalement l’exécution locale de modèles de langage au format GGUF. Là où llama.cpp exige de compiler les sources, télécharger manuellement les poids quantifiés et configurer chaque paramètre, Ollama propose une expérience proche de Docker : `ollama pull llama3.3`, `ollama run llama3.3`, et le modèle répond. Cette philosophie « batteries incluses » explique son adoption massive auprès des développeurs européens qui doivent composer avec les contraintes du RGPD et de l’AI Act.

En avril 2026, Ollama supporte officiellement trois plateformes : **Linux** (Ubuntu, Debian, Fedora, Arch), **macOS** (Apple Silicon avec Metal 3) et **Windows 10/11**. Le moteur d’inférence s’appuie sur llama.cpp pour le calcul, mais ajoute un serveur HTTP exposant deux surfaces d’API : l’API native (`/api/generate`, `/api/chat`, `/api/embed`, `/api/tags`) et une couche de compatibilité OpenAI (`/v1/chat/completions`) qui permet de réutiliser les SDK existants sans modification.

Les versions récentes ont apporté plusieurs améliorations clés, portées par un rythme de publication soutenu qui a franchi la barre des **216 releases cumulées sur GitHub** dès mai 2026. Ce rythme se reflète aussi dans la bibliothèque de modèles elle-même, qui comptait déjà plus de **1 200 modèles référencés** dans la registry Ollama en mai 2026, selon les données d’usage compilées par Presenc AI. La v0.6.2 introduit Flash Attention v2.7 pour AMD ROCm 6.3+, ce qui débloque enfin les Radeon RX 7900 XTX comme alternative crédible aux RTX 4090. Côté vitesse pure, la **v0.23.1** (5 mai 2026) avait déjà introduit le décodage spéculatif pour **Gemma 4 31B**, doublant son débit d’inférence selon Fazm.ai. Sur Apple Silicon, la préversion MLX livrée avec la v0.19 (mars 2026) exploite désormais nativement la mémoire unifiée d’Apple et vise environ **2× le débit d’inférence**, ce qui permet de faire tenir Llama 3.3 70B sur un MacBook Pro M3 Max 128 Go. Côté contexte, la variable `OLLAMA_CONTEXT_LENGTH` permet de définir une fenêtre par défaut (exemple : `OLLAMA_CONTEXT_LENGTH=8192 ollama serve`) sans toucher au Modelfile.

## Prérequis matériels et logiciels (versions exactes)

Avant d’installer Ollama, vérifiez que votre machine satisfait aux exigences. Les besoins en RAM varient drastiquement selon la taille du modèle et le niveau de quantification choisi. Un modèle 7B en Q4_K_M occupe environ 4,5 Go de RAM, contre 40 Go pour un 70B en Q4_K_M et plus de 140 Go pour le même modèle non quantifié en FP16. Pour 90 % des usages français (rédaction, RAG documentaire, génération de code), un GPU de 8 à 12 Go de VRAM suffit pour faire tourner les modèles 7B-13B avec des performances correctes.

| Composant | Minimum | Recommandé (13B) | Idéal (70B) | 
|---|---|---|---|
| OS | Ubuntu 22.04 / macOS 12 / Win 10 | Ubuntu 24.04 / macOS Sonoma / Win 11 | Ubuntu 24.04 LTS | 
| RAM système | 8 Go | 16 Go | 64 Go | 
| VRAM GPU | 4 Go (3B-7B) | 12 Go (RTX 4070) | 48 Go (RTX 6000 Ada) | 
| CPU (fallback) | 4 cœurs AVX2 | 8 cœurs AVX-512 | 16 cœurs Threadripper | 
| Stockage SSD | 20 Go | 100 Go NVMe | 500 Go NVMe Gen4 | 
| Connexion (pull initial) | 10 Mb/s | 100 Mb/s | 1 Gb/s | 

Côté logiciel, vous aurez besoin de Python 3.11 ou supérieur pour les scripts d’intégration (Python 3.12 recommandé en 2026), de `curl` pour les installeurs, et optionnellement de Docker 25.0+ si vous préférez containeriser le serveur. Pour les utilisateurs NVIDIA, les pilotes CUDA 12.4 ou supérieurs sont obligatoires depuis la v0.5.0. Sur AMD ROCm, la version 6.3 minimum est requise pour activer Flash Attention v2.7. Vérifiez votre version GPU avant de continuer.

```
# Vérifier la version du pilote NVIDIA
nvidia-smi | grep "Driver Version"
# Attendu : Driver Version: 550.xx ou supérieur
# Vérifier CUDA
nvcc --version
# Attendu : Cuda compilation tools, release 12.4+
# Vérifier Python
python3 --version
# Attendu : Python 3.11.x ou supérieur
# Vérifier l'espace disque libre
df -h ~ | tail -1
```
## Étape 1 : Installer Ollama sur Linux, macOS et Windows

L’installation d’Ollama varie selon votre système. Sur Linux, le script officiel détecte automatiquement votre distribution et installe le binaire dans `/usr/local/bin`, ainsi qu’un service systemd qui démarre Ollama au boot. Sur macOS, vous avez le choix entre l’application native (.dmg) qui intègre une icône dans la barre de menus, ou la formule Homebrew. Sur Windows 11, l’installeur PowerShell est devenu la méthode officielle depuis la v0.5.0.

