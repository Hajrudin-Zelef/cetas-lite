---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-23
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "Apple", "DeepSeek", "Huawei", "Hugging Face", "Intel", "Nvidia", "OpenAI", "SGLang", "Unsloth", "vLLM"]
dates: []
keywords: ["gpu", "amd", "apache", "attention", "datacenter", "deepseek", "dpo", "embeddings", "fine-tuning", "gguf", "gpus", "inference"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [2685, 2793]
sha256: c36fe2a80dc55ce416de140bc100ce902527f03a29c9bb2645516d901058a1a1
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

### 161.5. Quelle quantification pour ton matériel : le guide express

Le navigateur de modèles de LM Studio affiche un indicateur de compatibilité, mais voici la logique derrière — à connaître pour ne pas télécharger 40 Go pour rien :

| Ta VRAM/RAM | Quantification à viser (modèle 7B–8B) | Taille fichier | Qualité |
|---|---|---|---|
| 8 Go VRAM | Q4_K_M | ~5 Go | Très bonne (défaut recommandé) |
| 12–16 Go VRAM | Q5_K_M ou Q6_K | ~6–7 Go | Proche du FP16 |
| 24 Go VRAM | Q8_0, voire 2 modèles Q4 en parallèle | ~8–9 Go | Quasi-identique au modèle d'origine |
| CPU uniquement (16 Go RAM) | Q4_K_M, modèles ≤ 8B | ~5 Go | Lente (~5–15 tok/s) mais utilisable |
| Apple Silicon unifié (24 Go+) | Q4/Q5, modèles 8B–32B selon RAM | variable | Excellente (bande passante unifiée) |

Règles : **Q4_K_M = le défaut raisonnable** dans 90 % des cas ; au-delà de Q6_K, les gains qualité sont marginaux pour un usage chat/RAG ; un 70B en Q4 demande ~40 Go — c'est du multi-GPU ou du Mac Studio, pas un laptop. LM Studio te dit « passe / limite / ne passe pas » AVANT le téléchargement : fais-lui confiance.

### 161.6. LM Link + Locally : le détail qui change tout (v0.4.16)

Le montage exact, vérifié juin 2026 :
1. Ton **Mac** (moteur de calcul — Apple Silicon recommandé) fait tourner LM Studio avec le modèle chargé.
2. **LM Link** établit un lien via un mesh **Tailscale** chiffré de bout en bout entre ton Mac et ton iPhone/iPad.
3. L'app **Locally** (iOS/iPadOS uniquement au lancement) devient un terminal : tu chattes depuis ton téléphone, l'inférence tourne sur le Mac.

Ce qui ne transite JAMAIS par le cloud : tes prompts, tes réponses, tes modèles. Ce qui touche les serveurs LM Studio : uniquement la liste de découverte des appareils (pour que le téléphone trouve le Mac). Pour un chef de service qui veut interroger ses docs depuis le terrain sans exposer de données : c'est l'architecture « private inference » la plus simple du marché — à condition d'avoir le Mac allumé au bureau.

### 161.7. Le CLI `lms` : LM Studio sans la souris

LM Studio embarque un CLI pour les allergiques à la GUI (utile pour scripter) :

```bash
# Lister / télécharger / charger un modèle
lms ls
lms get qwen3-8b              # télécharge (équivalent du clic dans Search)
lms load qwen3-8b --gpu 1.0   # charge en VRAM (gpu 0..1 = fraction offloadée)

# Démarrer le serveur OpenAI-compatible (port 1234 par défaut)
lms server start
lms server stop

# Chat direct au terminal
lms chat qwen3-8b
```

Cas d'usage : un script qui charge le modèle le matin (`lms load`), démarre le serveur, et l'éteint le soir — ton poste devient un mini-backend local sans intervention.

### 161.8. LM Studio dans la chaîne « modèle local » complète

```
Hugging Face (poids GGUF)
   → LM Studio : découverte, test interactif, comparaison de quants
   → export/décision : « ce modèle vaut le coup en prod ? »
   → Ollama (serveur léger) ou vLLM/SGLang (serveur lourd) ou NIM (clé en main)
   → ton routeur LiteLLM
```

LM Studio n'est pas le maillon de prod — c'est le **maillon d'évaluation**. Sa valeur = réduire à 10 minutes la question « ce nouveau modèle open-weight mérite-t-il une place dans ma stack ? », avec historique et comparaisons à l'appui au lieu d'impressions.

### 161.9. Modèles recommandés 2026 pour ton usage (à tester dans LM Studio)

| Usage | Modèle GGUF à chercher | Pourquoi |
|---|---|---|
| Chat technique FR généraliste | Qwen3-8B / 14B instruct (Q4_K_M) | Bon FR, tool calling correct, rapide |
| Raisonnement / debug | Qwen3-32B (si 24 Go+ VRAM) ou DeepSeek-R1-Distill 8B | Le distill R1 donne du raisonnement visible pas cher |
| RAG local : embeddings | nomic-embed-text / bge-m3 (GGUF) | Pour tester le « tout local » avant de décider |
| Vision (lire une plaque onduleur) | Qwen2.5-VL / Qwen3-VL (GGUF) | Glisser-déposer d'image dans le chat LM Studio |
| Code sysadmin | Qwen3-Coder / Devstral | Scripts bash, Ansible, debug |

Protocole d'évaluation express (30 min) : charge 2 candidats → même 5 questions métier FR → compare côte à côte (fonction native) → note qualité/vitesse/VRAM → le gagnant part en test sur Ollama/vLLM. Ce rituel t'évite d'ajouter un modèle à ta stack sur un simple effet de mode Hugging Face.

---

## 162. Unsloth : le fine-tuning efficace (2x plus rapide, 2x moins de VRAM)

**En une phrase :** une surcouche open source (Apache 2.0) à Hugging Face TRL qui rend le fine-tuning LoRA/QLoRA accessible sur une seule GPU consommatrice — grâce à des kernels Triton écrits à la main et un moteur de backprop optimisé, sans perte de précision revendiquée.

### 162.1. Principe

Le fine-tuning classique d'un LLM est gourmand : un 7B en 16 bits demande ~60 Go de VRAM pour l'entraînement complet (poids + gradients + états d'optimiseur). Unsloth attaque le problème sur deux axes :

1. **LoRA / QLoRA** : au lieu de réentraîner tous les poids, on gèle le modèle et on n'entraîne que de petites matrices d'adaptation (rank `r` = 16 typiquement, soit < 1 % des paramètres). En QLoRA, le modèle de base est en plus quantifié en 4 bits → un 7B tient dans ~5 Go + les adaptateurs.
2. **Kernels Triton manuels** : Unsloth remplace les opérations coûteuses (attention, backward des couches LoRA, RoPE, cross-entropy) par des kernels écrits à la main, avec un moteur de backpropagation manuel. Résultat revendiqué : **~2x plus rapide, ~70 % de VRAM en moins, 0 % d'erreur d'approximation** (pas de raccourci mathématique, juste du code mieux optimisé).

En pratique : fine-tuner un 7B–8B instruct sur tes données métier (tickets de maintenance, procédures onduleurs, docs Huawei) devient faisable sur une **RTX 4090 24 Go**, voire une 3090 — là où le fine-tuning naïf exigerait du datacenter.

### 162.2. Installation (vérifiée sept 2026)

```bash
# Méthode recommandée (Linux ou WSL2 — pas de support Windows natif)
pip install unsloth

# Avec vLLM (utile pour le GRPO : l'inférence rapide des rollouts)
pip install uv && uv pip install unsloth vllm

# Docker : tout pré-installé, avec Jupyter
docker run -d --gpus all -p 8888:8888 \
  -e JUPYTER_PASSWORD="mot-de-passe-fictif" \
  -v $(pwd)/work:/workspace/work \
  unsloth/unsloth
```

Prérequis : **GPU NVIDIA, capacité CUDA 7.0+** (T4, V100, RTX 20/30/40/50, A100, H100, L40), Python 3.10–3.13. Support AMD/Intel en cours (partiel — vérifie la doc pour ta carte). Sur Colab/Kaggle : fonctionne, c'est d'ailleurs le chemin le plus courant pour un premier essai gratuit (T4 gratuit = QLoRA 7B possible, lentement).

Piège classique : Unsloth gère l'attention en interne — **n'installe pas `flash-attn` manuellement** avec Unsloth, il utilise SDPA en interne sur les modèles récents et ça entre en conflit.

### 162.3. Cas d'usage pour ton RAG (et ton métier)

1. **Adaptation au domaine** : SFT (supervised fine-tuning) sur tes paires question/réponse métier — ex. 500–2000 Q/R sur les onduleurs, les CLI Huawei, les procédures de maintenance. Le modèle apprend ton vocabulaire, ton format de réponse, tes gammes de produits.
2. **Format de sortie contraint** : entraîner le modèle à toujours répondre en JSON structuré, ou avec les citations au format que ton RAG attend.
3. **GRPO / DPO** : au-delà du SFT, Unsloth supporte le reinforcement learning (GRPO, DPO, GSPO) — utile quand tu as une **fonction de récompense** plutôt que des exemples parfaits (ex. : récompenser les réponses qui citent la bonne section de doc).
4. **Export universel** : après entraînement, exporte en **GGUF** (Ollama, LM Studio — section 161), fusionne en 16 bits, ou sers via **vLLM / SGLang** (section 163). Le LoRA reste compatible avec l'écosystème transformers.

