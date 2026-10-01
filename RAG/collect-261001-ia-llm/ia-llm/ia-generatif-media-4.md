---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-4
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["EU", "Fireworks AI", "Nvidia", "Together AI"]
dates: []
keywords: ["apache", "decode", "diffusion", "fp8", "gguf", "gpu", "inference", "mistral", "nvidia", "safetensors", "text-to-image"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [301, 439]
sha256: 91baa7114a573ef06abdf8fa619260ed06f1b9bee373034231987a7abd6e7240
---

# IA générative : image, vidéo, recherche

FLUX.1 (forfait) : 1.1 [pro] **0,04 $**, Ultra **0,06 $**, Kontext [pro] **0,04 $**, Kontext [max] **0,08 $**, Fill [pro] **0,05 $**.

Ordre de grandeur : une image 1024×1024 en FLUX.2 [pro] = **3 centimes**. C'est ~4× moins cher que Nano Banana Pro à qualité comparable sur beaucoup d'usages (voir section 78). Pour du volume (centaines d'images de doc), l'écart est massif.

Exemple d'appel (schéma officiel BFL) :

```python
import os, time, requests

BFL_KEY = os.environ["BFL_API_KEY"]
EU = "https://api.eu.bfl.ai"   # résidence UE : privilégier pour données européennes

# 1. Soumettre la génération (FLUX.2 Pro, ~1 MP)
r = requests.post(f"{EU}/v1/flux-2-pro",
    headers={"x-key": BFL_KEY, "Content-Type": "application/json"},
    json={"prompt": "photographie documentaire d'une armoire électrique industrielle "
                    "ouverte, borniers et disjoncteurs modulaires alignés, "
                    "étiquettes lisibles, lumière d'atelier neutre, "
                    "style photo technique, très net",
          "width": 1024, "height": 1024}).json()
task_id = r["id"]

# 2. Attendre le résultat (en prod : webhook)
while True:
    res = requests.get(f"{EU}/v1/get_result",
                       headers={"x-key": BFL_KEY},
                       params={"id": task_id}).json()
    if res["status"] == "Ready":
        print("Image :", res["result"]["sample"])   # URL temporaire de l'image
        break
    time.sleep(2)
```

## 20. FLUX via API tierces : Replicate, fal.ai et les autres

BFL n'est pas le seul guichet. FLUX est exposé chez **Replicate, fal.ai, Together AI, Fireworks AI, Runware, DeepInfra, Cloudflare**, avec leur propre grille (marge variable). Intérêts :

- **Un seul compte, une seule clé, une seule facture** pour FLUX + Kling + Veo + autres (fal.ai et Replicate hébergent des dizaines de modèles).
- Facturation **à l'usage pur**, sans crédit BFL à pré-charger.
- Exemples de tarifs constatés (sept 2026, **tarifs agrégateurs, pas BFL**) : FLUX [schnell] ~**0,003 $/image** sur fal.ai (idéal pour les brouillons jetables), Nano Banana ~0,04 $, Nano Banana Pro ~0,15-0,20 $.

Quand choisir quoi :
- **Volume régulier, un seul modèle** → API officielle BFL (moins cher, relation directe, endpoint UE).
- **Multi-modèles, prototypes, volumes irréguliers** → fal.ai / Replicate (simplicité > prix unitaire).
- Toujours comparer le **coût total à volume égal** avant de t'engager, et noter le fournisseur dans ton suivi (les prix bougent indépendamment).

## 21. FLUX en local : pourquoi c'est un cas à part

C'est **le** différenciateur stratégique de FLUX pour un sysadmin :

- **Confidentialité totale** : plans, photos d'installations clients, schémas réseau ne quittent jamais ta machine. Le seul flagship 2026 qui le permet avec des poids ouverts.
- **Coût marginal nul** après l'investissement GPU : des milliers d'images sans compteur.
- **Reproductibilité** : seed fixe, version de modèle figée, pas de « mise à jour silencieuse » côté fournisseur.
- **Air-gap** : possible (télécharger les poids une fois, travailler déconnecté).

Inconvénients : investissement GPU, consommation électrique, maintenance (drivers, CUDA), vitesse inférieure au cloud, et **licence** (voir section 18 : klein 4B ou schnell en Apache 2.0 pour du commercial).

## 22. FLUX en local : prérequis matériels (vérifiés sept 2026)

| Modèle | VRAM min (pratique) | Config type | Disque |
|--------|---------------------|-------------|--------|
| FLUX.2 [dev] 32B (FP8) | **24 Go** | RTX 3090 / 4090 / 5090 | ~75 Go |
| FLUX.2 [dev] 32B (BF16) | ~44-64 Go | RTX 6000 Ada / pro | ~90 Go |
| FLUX.2 [dev] (GGUF Q4_K_S) | ~19-24 Go | RTX 4090 | ~40 Go |
| FLUX.2 [klein] 9B (FP8) | **16 Go** | RTX 4080+ | ~30 Go |
| FLUX.2 [klein] 4B | **12 Go** | RTX 4070 / 3060 12 Go | ~15 Go |
| FLUX.1 [dev]/[schnell] 12B (FP8) | **16 Go** | RTX 4090 confortable | ~30 Go |

La quantification **FP8** (optimisée avec NVIDIA pour FLUX.2) divise la VRAM par ~2 avec une perte de qualité faible : c'est le réglage standard en 2026. Sans GPU correct, la location d'un GPU cloud (RunPod, etc.) reste l'option réaliste — pas de chemin CPU viable pour un modèle 32B.

## 23. FLUX en local via ComfyUI : installation pas à pas

**ComfyUI** est l'orchestrateur de workflows de génération d'image le plus utilisé (interface par nœuds, reproductible, scriptable via API). Procédure pour FLUX.2 [dev] en FP8 (sept 2026) :

```bash
# 1. Installer ComfyUI (portable)
git clone https://github.com/comfyanonymous/ComfyUI.git
cd ComfyUI
python -m venv venv && source venv/bin/activate
pip install torch torchvision --index-url https://download.pytorch.org/whl/cu128
pip install -r requirements.txt

# 2. Récupérer les poids (ex. ~71 Go au total en FP8/BF16 mixte)
pip install -U "huggingface_hub[cli]"
hf download Comfy-Org/flux2-dev split_files/diffusion_models/flux2_dev_fp8mixed.safetensors \
  --local-dir models/diffusion_models        # ~35,5 Go : le modèle (UNET)
hf download Comfy-Org/flux2-dev split_files/text_encoders/mistral_3_small_flux2_fp8.safetensors \
  --local-dir models/text_encoders           # version FP8 du CLIP : moitié moins lourde
hf download black-forest-labs/FLUX.2-small-decoder full_encoder_small_decoder.safetensors \
  --local-dir models/vae                     # ~1 Go : le VAE

# 3. Lancer
python main.py --listen 127.0.0.1 --port 8188
# Interface : http://127.0.0.1:8188
```

Dans l'interface : charger le **template officiel FLUX.2** (menu Templates de ComfyUI), vérifier que les trois chargeurs pointent vers les bons fichiers, régler **20-30 steps** pour les tests (jusqu'à 50 pour le master), sampler `euler`, guidance ~3,5, résolution 1024×1024 pour commencer (monter prudemment : la VRAM explose au-delà).

**Astuce VRAM** : le nœud « purge VRAM » entre les grosses générations, et le weight streaming (déchargement partiel vers la RAM système) activé par les builds récents rendent FLUX.2 utilisable même quand ça coince.

## 24. FLUX en local : premier workflow et variante légère

Workflow minimal text-to-image dans ComfyUI (nœuds) :

1. `Load Diffusion Model` → `flux2_dev_fp8mixed.safetensors`
2. `Load CLIP` → `mistral_3_small_flux2_fp8.safetensors`
3. `Load VAE` → le VAE FLUX.2
4. `CLIP Text Encode (Prompt)` → ton prompt (positif) ; un second pour le négatif (`floou, déformé, filigrane, texte illisible`)
5. `Empty Latent Image` → 1024×1024
6. `KSampler` → steps 25, cfg 3.5, sampler `euler`, scheduler `simple`, seed fixe (ex. 42) pour reproductibilité
7. `VAE Decode` → `Save Image`

**Variante légère (klein 4B, Apache 2.0, usage commercial OK)** : même workflow, modèle 4B, ~8-12 Go de VRAM, génération en ~1 s sur RTX 5090. Parfait pour : itérer sur les compositions, faire tourner un service interne de visuels, ou équiper un poste modeste. La qualité est inférieure à dev/pro mais largement suffisante pour des schémas, illustrations de doc et maquettes.

## 25. FLUX via Diffusers (Python, sans ComfyUI)

Pour intégrer la génération dans un script ou un pipeline (ex. générer les visuels d'une doc automatiquement) :

```python
import torch
from diffusers import FluxPipeline   # FLUX.1 ; voir la doc diffusers pour FLUX.2

pipe = FluxPipeline.from_pretrained(
    "black-forest-labs/FLUX.1-schnell",   # Apache 2.0 : commercial OK
    torch_dtype=torch.bfloat16,
).to("cuda")

image = pipe(
    "schéma isométrique d'un local technique avec deux onduleurs 40 kVA, "
    "armoire électrique et climatisation de précision, style illustration "
    "technique épurée, fond clair",
    height=1024, width=1024,
    num_inference_steps=4,                 # schnell : 1 à 4 steps suffisent
    guidance_scale=0.0,
    generator=torch.Generator("cuda").manual_seed(7),
).images[0]
image.save("local_technique.png")
```

