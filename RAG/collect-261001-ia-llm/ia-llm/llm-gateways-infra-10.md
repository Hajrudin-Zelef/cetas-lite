---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-10
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "Nvidia", "vLLM"]
dates: []
keywords: ["gpu", "diffusion", "embeddings", "gguf", "gpus", "inference", "lora", "nvidia", "quantization", "qwen", "reranker", "vllm"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [1200, 1348]
sha256: 51206b441266ec62ac9092cb567923478fdc3f7bd61e744c3c6090b9751f0d21
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

**Règle** : avant d'intégrer un modèle à ton infra pro, **lire la fiche licence du repo exact** (pas « la famille en général »). Les licences changent entre versions.

## 83. Spaces : tester sans installer

Les **Spaces** = démos Gradio/Streamlit hébergées gratuitement. Devant un modèle inconnu : cherche son Space, teste 10 prompts de ton domaine, décide ensuite si tu le télécharges. C'est l'équivalent d'un essai routier — 0 install, 0 GPU.

## 84. Workflow complet : « trouver et tester un modèle pour mon RAG »

1. **Besoin** : ex. « reranker FR léger pour CPU ».
2. **Hub** : filtre `text-reranking` + langue FR → candidats (`bge-reranker-v2-m3`…).
3. **Fiche** : licence OK ? dimensions ? contexte ?
4. **Space / widget** : test rapide de pertinence sur 5 paires query/doc de ton domaine.
5. **Inference Providers** : test API en 10 lignes (section 79) pour mesurer latence/format.
6. **Local** : TEI (embeddings/rerank) ou Ollama (LLM) avec le modèle exact.
7. **Éval** : tes 50 questions, métrique précision@5 (section 34).
8. **Décision** : figer `model-id` + version + hash dans ta doc (reproductibilité).

## 85. `huggingface_hub` : la CLI utile

```bash
pip install huggingface_hub
hf download BAAI/bge-m3                       # télécharger un modèle en cache local
hf download Qwen/Qwen3-32B-GGUF --include "*Q4_K_M*"  # ne prendre qu'une quantization
huggingface-cli login                         # token hf_... pour les modèles gated
```

Le cache `~/.cache/huggingface` est partagé avec TEI (volume `/data`), vLLM et transformers : **télécharge une fois, sers partout**.

## 86. Checklist Hugging Face

- [ ] Compte + token `hf_...` (lecture) créé, stocké dans le gestionnaire de secrets.
- [ ] Savoir filtrer par tâche/langue/licence sur le Hub.
- [ ] Tester un modèle via Space + Inference Providers avant tout téléchargement.
- [ ] Lire la licence du repo exact avant usage pro.
- [ ] Cache HF partagé entre TEI/vLLM/transformers (un seul volume).
- [ ] Épingler les `model-id` + révisions dans ta doc d'infra.

---

# PARTIE I — LOUER DU GPU : VAST.AI ET RUNPOD

## 87. Pourquoi louer plutôt qu'acheter (le calcul)

Acheter une RTX 4090 ≈ 1 900 €. Louer ≈ $0.34–0.69/h.
- **Seuil de rentabilité** : 1 900 € / 0.50 €/h ≈ **3 800 h** ≈ 158 jours de 24h. Si ton besoin annuel < ~1 000 h, **la location gagne**.
- La location donne aussi accès à du **H100/A100** impossible à acheter pour un particulier, et à la **facturation à la seconde** (tu paies 20 min, pas le mois).
- En revanche : pas de données sensibles (machine d'un tiers), disponibilité variable (spot), egress parfois facturé.

## 88. Vast.ai : principe

**Marketplace pair-à-pair** : des particuliers/datacenters louent leurs GPU, tu enchéris. Deux modes :
- **Interruptible (spot)** : le moins cher (~30–50 % sous l'on-demand), mais l'hôte peut reprendre la machine (préempter). Pour batch/fine-tune avec checkpoints fréquents.
- **On-demand** : prix fixe plus élevé, pas de préemption.
Facturation **à la seconde**, pas de minimum. Catalogue : de la RTX 3060 au H100 8×. Egress : ~$0.01–0.02/Go (vérifié avril 2026 — à re-vérifier, c'est le coût caché classique).

## 89. Vast.ai : prix vérifiés (ordres de grandeur, avril–sept 2026)

| GPU | Spot (interruptible) | On-demand | Usage type |
|---|---|---|---|
| RTX 3090 24 Go | ~$0.13/h | ~$0.22/h | 32B Q4 limite, embeddings |
| RTX 4090 24 Go | ~$0.29–0.34/h | ~$0.45–0.59/h | 32B Q4, LoRA 8B |
| A100 80 Go | ~$0.75/h | ~$1.10–1.45/h | 70B Q4, fine-tune sérieux |
| L40S 48 Go | ~$1.20/h | ~$1.75/h (à vérifier) | Inférence 70B |
| H100 80 Go | ~$1.65–2.60/h | ~$2.40–2.90/h | Gros fine-tune, vLLM 70B |

Prix **variables par hôte** (enchères) : ce sont des médianes constatées, pas un tarif officiel. Toujours trier par prix/fiabilité (`reliability > 98 %`, `cuda_vers` compatible).

## 90. Vast.ai : louer pas à pas (CLI)

```bash
# 1. Installer la CLI (SDK + CLI en un)
pip install vastai
vastai set api-key TA_CLE_VAST_FAKE

# 2. Chercher une 4090 fiable et pas chère
vastai search offers \
  'gpu_name=RTX 4090 num_gpus=1 reliability>0.98 cuda_max_good>=12.4 disk_space>60' \
  -o 'dph_total-'

# 3. Créer l'instance (image PyTorch de base + port SSH + Jupyter)
vastai create instance <OFFER_ID> \
  --image pytorch/pytorch:2.5.1-cuda12.4-cudnn9-devel \
  --disk 60 \
  --ssh \
  --jupyter-dir /

# 4. Suivre le démarrage
vastai show instances
vastai logs <INSTANCE_ID>

# 5. SSH (la CLI affiche la commande exacte : port haute, ex. 23456)
ssh -p 23456 root@<IP_HOTE> -i ~/.ssh/id_ed25519

# 6. Détruire à la fin (LE réflexe à automatiser — section 98)
vastai destroy instance <INSTANCE_ID>
```

**Template d'image** : pars de `pytorch/pytorch:...-devel` (toolkit CUDA complet) ou d'une image avec vLLM/Ollama préinstallé. Monte un volume persistant si tu veux garder les poids entre deux locations.

## 91. RunPod : principe

RunPod = **cloud GPU orienté devs**, deux offres :
- **Community Cloud** : machines de particuliers (comme Vast.ai), les moins chères, dispo variable.
- **Secure Cloud** : infra RunPod directe, SLA, plus chère (~1,7×).
- **Serverless** : facturation à la milliseconde, scale-to-zero — pour exposer une API d'inférence sans gérer de machine.
- **Templates** : images 1-clic (PyTorch, vLLM, Ollama, Stable Diffusion…) — le gros avantage vs Vast.ai pour démarrer vite.
- **Egress gratuit**, facturation fine (ms), API Python/GraphQL + CLI `runpodctl` de qualité.

## 92. RunPod : prix vérifiés (avril–sept 2026)

| GPU | Community | Secure Cloud | Note |
|---|---|---|---|
| RTX 4090 24 Go | ~$0.34/h | ~$0.59–0.74/h | Le standard dev |
| RTX 3090 24 Go | ~$0.22/h | n/a | Budget |
| A100 80 Go | ~$1.19/h | ~$1.89–2.49/h | 70B |
| L40S 48 Go | ~$1.90/h | ~$2.49/h | Inférence pro 70B |
| H100 80 Go | ~$2.49/h | ~$2.99–3.89/h | Haut de gamme |

Egress : **$0** (vérifié). À comparer avec Vast.ai ($0.01–0.02/Go) quand tu rapatries des checkpoints de 40–140 Go.

## 93. RunPod : louer pas à pas (UI + API)

**Via l'UI** (le plus simple) :
1. `runpod.io` → Deploy → choisir **Community** (budget) ou **Secure** (fiabilité).
2. GPU : RTX 4090. Template : **« vLLM »** ou **« PyTorch 2.x »** (1-clic).
3. Disque conteneur : 60 Go min (les poids + le cache). Volume réseau si persistance voulue.
4. Exposer : HTTP 8000 (vLLM), TCP 22 (SSH). **Mettre une clé SSH, pas de mot de passe.**
5. Deploy → attendre `RUNNING` → noter l'URL/ports.

**Via l'API Python** (reproductible) :
```python
import runpod
runpod.api_key = "RUNPOD_API_KEY_FAKE"
gpu = runpod.get_gpus()
# Choisir une 4090 community, créer le pod avec le template vLLM :
pod = runpod.create_pod(
    name="vllm-32b",
    image_name="runpod/pytorch:2.4.0-py3.11-cuda12.4.1-devel-ubuntu22.04",
    gpu_type_id="NVIDIA GeForce RTX 4090",
    cloud_type="COMMUNITY",
    volume_in_gb=60,
)
print(pod["id"], pod["desiredStatus"])
```

**Templates à connaître** : `runpod/pytorch` (base), images communautaires **vLLM**, **Ollama** (chercher « ollama » dans les templates), **Jupyter**. Un template = image + ports + variables d'env : fork-le pour figer ta config (ex. ton Modelfile pré-chargé).

## 94. Vast.ai vs RunPod : tableau comparatif

