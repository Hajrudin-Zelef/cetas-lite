---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-17
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["ByteDance", "EU", "Google", "MiniMax", "Moonshot", "OpenAI", "Perplexity"]
dates: ["2026-07-16", "2026-07-24", "2026-09-15", "2026-09-24"]
keywords: ["agents", "apache", "gemini", "gpu", "inference", "kimi", "perplexity"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [1639, 1777]
sha256: b188f44c7cd21434f247dffa5d7c2bf3578c2d48f8e0e3ff024d8dbc840b0bed
---

# IA générative : image, vidéo, recherche

| Outil | Version confirmée | Prix confirmés | API confirmée | Licence vérifiée |
|-------|-------------------|----------------|---------------|------------------|
| Kling | 3.0 (fév. 2026) | Web : 0/6,99/25,99/64,99/127,99 $ ; API : packs 9,80-7 560 $ | Oui (`kling.ai/dev`, packs prépayés) | Commercial : plans payants uniquement |
| FLUX | FLUX.2 (pro/max/flex/dev/klein) ; FLUX.1 (pro/dev/schnell/kontext) | API BFL : 0,014-0,07 $/MP ; FLUX.1 : 0,04-0,08 $ | Oui (`api.bfl.ai`, `api.eu.bfl.ai`, `api.us.bfl.ai`) | klein 4B & schnell : Apache 2.0 ; dev : non-commercial |
| Perplexity | Offres Free/Pro/Max/Enterprise, Comet | Pro 20 $, Max 200 $, API Sonar ~1-15 $/1M + 5-14 $/1K req | Oui (Sonar, OpenAI-compatible) | Réponses : oui ; sources : leurs licences |
| Midjourney | V8.2 (défaut 24/07/2026) | 10/30/60/120 $ (GPU time) | Non (pas d'API publique) | Commercial oui ; >1 M$ → Pro/Mega |
| OpenAI image | gpt-image-2/1.5/1-mini/1, DALL-E 3 | 0,005-0,25 $/image selon modèle | Oui (Images API) | Oui (CGU) |
| Google image | Imagen 4 (Fast/Std/Ultra), Nano Banana (1/2/Pro) | 0,02-0,24 $/image ; batch -50 % | Oui (Gemini API / Vertex) | Oui (API payante) |
| Sora | API **fermée le 24/09/2026** | — | **Non (coupée)** | — |
| Veo 3.1 | 3.1 (Standard/Fast/Lite) | 0,05-0,60 $/s | Oui (Gemini/Vertex) | Oui |
| Runway | Gen-4.5 | 15/35/95 $ (crédits) | Oui (crédits) | Oui (payant) |
| Luma | Ray 3.2 / Luma Agents | 30/90/300 $ (Agents) | Oui | Oui (dès Plus) |
| Hailuo/MiniMax | H3 | 0,26 $/s forfaitaire ; abo ~14,99 $ | Oui | **Territoriale restrictive — à vérifier** |
| Seedance | 2.0/2.5 (ByteDance) | ~0,134-0,29 $/s | Oui (BytePlus + agrégateurs) | Oui (CGU) |
| Phind / You.com | Actifs en 2026 | Freemium (**à vérifier**) | Limitée/**à vérifier** | — |
| Kimi | K3 (16/07/2026, poids ouverts) | Abonnements **suspendus** | Oui (`platform.kimi.com`) | Licence Kimi K3 |
| Jev (TypeSafe) | jev-1.13.0 (stealth levé 15/09/2026) | 0,042 $/1M in, out gratuit | Oui (`api.typesafe.ai`) | CGU TypeSafe |

---

# PARTIE S — PATTERNS API COMPARÉS : le code qui marche en prod

## 162. Le pattern universel des API génératives

Toutes les API de ce guide partagent le même squelette (à l'exception des API synchrones simples) :

```
1. Authentification  → clé dans header (x-key, Bearer, ?key=)
2. Soumission        → POST avec prompt + paramètres → job_id / task_id
3. Attente           → polling GET ou webhook (préférer webhook en prod)
4. Récupération      → URL temporaire → télécharger VITE (elle expire !)
5. Journalisation    → modèle, version, seed, coût, durée
```

Les URL de téléchargement **expirent** (souvent 1h à 24h) : toujours rapatrier le fichier sur ton stockage immédiatement, jamais « le lien dans un tableur ».

## 163. Pattern : génération d'image synchrone (fal.ai, exemple)

fal.ai expose un endpoint synchrone pratique pour les prototypes (pas de polling) :

```python
import os, requests

FAL_KEY = os.environ["FAL_API_KEY"]
headers = {"Authorization": f"Key {FAL_KEY}", "Content-Type": "application/json"}

# FLUX schnell : brouillons à ~0,003 $/image
r = requests.post("https://fal.run/fal-ai/flux/schnell",
                  headers=headers,
                  json={"prompt": "armoire électrique industrielle, borniers alignés, "
                                  "photo technique nette, lumière neutre",
                        "image_size": "square_hd",
                        "num_inference_steps": 4}).json()
url = r["images"][0]["url"]
print("Image :", url)   # télécharger immédiatement (URL temporaire)
```

> Hôte et chemins illustratifs : la doc fal.ai fait foi. Le principe (clé `Key`, POST synchrone, `images[].url`) est stable.

## 164. Pattern : BFL asynchrone robuste (avec timeout et erreurs)

Version « prod » de l'exemple de la section 19 : timeouts, erreurs typées, téléchargement immédiat.

```python
import os, time, requests

BFL_KEY = os.environ["BFL_API_KEY"]
EU = "https://api.eu.bfl.ai"
H = {"x-key": BFL_KEY, "Content-Type": "application/json"}

def generer_image_bfl(prompt, width=1024, height=1024, model="flux-2-pro",
                      timeout_s=300):
    # 1. Soumission
    sub = requests.post(f"{EU}/v1/{model}", headers=H,
                        json={"prompt": prompt, "width": width, "height": height},
                        timeout=30).json()
    task_id = sub["id"]
    # 2. Attente bornée
    t0 = time.time()
    while time.time() - t0 < timeout_s:
        res = requests.get(f"{EU}/v1/get_result", headers=H,
                           params={"id": task_id}, timeout=30).json()
        st = res["status"]
        if st == "Ready":
            url = res["result"]["sample"]
            # 3. Rapatriement immédiat
            img = requests.get(url, timeout=60).content
            with open(f"bfl_{task_id}.png", "wb") as f:
                f.write(img)
            # 4. Journalisation (modèle, version, paramètres, coût estimé)
            print(f"OK {model} {width}x{height} -> bfl_{task_id}.png")
            return f"bfl_{task_id}.png"
        if st in ("Error", "Failed"):
            raise RuntimeError(f"BFL échec {task_id} : {res}")
        time.sleep(3)
    raise TimeoutError(f"BFL timeout {task_id} après {timeout_s}s")

# Exemple : coût ~0,03 $ (1 MP, flux-2-pro)
generer_image_bfl("salle serveurs, rangées de baies 19 pouces, LEDs bleues, "
                  "photo documentaire, lumière froide")
```

## 165. Pattern : Kling 3.0 via agrégateur (sans pack prépayé)

Quand tu veux de l'API Kling **sans** acheter un pack de 700 $ : passer par un agrégateur à l'usage pur. Schéma (fal.ai, tarifs constatés sept 2026 : ~0,112 $/s sans audio, ~0,168 $/s avec audio — **tarifs agrégateur, pas Kling officiel**) :

```python
import os, requests, time

FAL_KEY = os.environ["FAL_API_KEY"]
headers = {"Authorization": f"Key {FAL_KEY}", "Content-Type": "application/json"}

# Image-to-video : animer une vraie photo d'équipement
job = requests.post("https://fal.run/fal-ai/kling-video/v3.0/pro/image-to-video",
    headers=headers,
    json={
        "image_url": "https://ton-stockage/photo-onduleur.jpg",  # ton asset
        "prompt": "léger travelling avant, les voyants verts clignotent doucement, "
                  "ambiance feutrée de local technique, 10 secondes",
        "duration": "10",
        "aspect_ratio": "16:9",
    }).json()
print("Résultat :", job)  # selon l'agrégateur : synchrone ou job à suivre
```

Arbitrage : agrégateur = **0 $ d'engagement**, marge de ~20-40 % sur le tarif officiel ; API officielle = **moins cher à l'unité**, pack prépayé qui expire. Seuil de bascule typique : usage hebdomadaire régulier.

## 166. Pattern : Perplexity Sonar avec gestion du budget

```python
import os, time
from openai import OpenAI

client = OpenAI(api_key=os.environ["PERPLEXITY_API_KEY"],
                base_url="https://api.perplexity.ai")

BUDGET_REQUETES_MOIS = 500
compteur = {"n": 0}

