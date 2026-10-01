---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-18
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["apache", "attention", "gemini", "gpu", "lora"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [1778, 1939]
sha256: 0447931b6bc641891bf6303e6829f2ceba017db87f796cf1e30d983555677d75
---

# IA générative : image, vidéo, recherche

def sonar_ask(question, model="sonar", max_tokens=800):
    """Interroge Sonar avec garde-fous : modèle économe par défaut, compteur."""
    if compteur["n"] >= BUDGET_REQUETES_MOIS:
        raise RuntimeError("Budget mensuel Sonar atteint — augmenter ou attendre.")
    t0 = time.time()
    r = client.chat.completions.create(
        model=model,   # "sonar" par défaut ; "sonar-pro" si besoin de profondeur
        messages=[{"role": "system",
                   "content": "Réponds en français, de façon concise, avec les sources "
                              "(titre + URL + date) après chaque affirmation importante."},
                  {"role": "user", "content": question}],
        max_tokens=max_tokens,
    )
    compteur["n"] += 1
    usage = r.usage
    print(f"[Sonar] req {compteur['n']}/{BUDGET_REQUETES_MOIS} | "
          f"{usage.prompt_tokens} in / {usage.completion_tokens} out | "
          f"{time.time()-t0:.1f}s")
    return r.choices[0].message.content

print(sonar_ask("Quelle est la durée de vie typique des condensateurs du bus DC "
                "d'un onduleur double conversion ? Cite les sources."))
```

## 167. Pattern : Gemini image (Nano Banana Pro) via l'API

```python
import os
from google import genai   # SDK google-genai (nom du package : à vérifier à l'install)

client = genai.Client(api_key=os.environ["GEMINI_API_KEY"])

# Nano Banana Pro : ~0,134 $/image en 1-2K
result = client.models.generate_content(
    model="gemini-3-pro-image-preview",
    contents=["Photographie produit d'un onduleur triphasé 60 kVA en armoire, "
              "fond gris clair studio, lumière douce, ultra net, style catalogue "
              "constructeur"],
)
# Sauvegarder les parties image de la réponse (structure selon SDK : à vérifier)
for i, part in enumerate(result.parts):
    if getattr(part, "as_bytes", None):
        with open(f"nbpro_{i}.png", "wb") as f:
            f.write(part.as_bytes)
        print(f"Image sauvegardée : nbpro_{i}.png (~0,134 $)")
```

> Le SDK `google-genai` évolue vite : vérifier le nom du modèle (`gemini-3-pro-image-preview`) et la structure de réponse dans la doc au moment de l'usage.

## 168. Pattern : Jev en garde-fou de pipeline (cascade vérifiée)

Le pattern recommandé par l'écosystème Jev : **LLM qui produit, Jev qui vérifie** (verified cascade).

```python
import os, requests

JEV_KEY = os.environ["JEV_API_KEY"]
JEV_URL = "https://api.typesafe.ai/v1/systemone"
JH = {"Authorization": f"Bearer {JEV_KEY}", "Content-Type": "application/json"}

def jev_check(state, questions, model="jev-1.13.0", seuil=0.75):
    """Retourne (ok: bool, détails). ok=False si une proba < seuil."""
    r = requests.post(JEV_URL, headers=JH, timeout=20,
                      json={"model": model, "state": state,
                            "questions": questions}).json()
    verdicts = {}
    for nom, q in questions.items():
        rep = r[nom]
        if q["type"] == "noul":
            p = rep["p_yes"]; ok = p >= seuil
        else:  # choice / score : on seuille la confiance
            p = rep.get("confidence", 0); ok = p >= seuil
        verdicts[nom] = {"ok": ok, "p": round(p, 3), "brut": rep}
    return all(v["ok"] for v in verdicts.values()), verdicts

# Exemple : valider un visuel avant publication (coût : une fraction de centime)
ok, details = jev_check(
    state="Visuel pour doc 'TGBT atelier' : prompt='armoire électrique, borniers', "
          "model=flux-2-klein-4b, seed=11. Relecture humaine : textes OK, câblage OK.",
    questions={
        "commercial_ok": {"type": "noul",
            "question": "La licence du modèle (Apache 2.0) autorise-t-elle l'usage commercial ?"},
        "risque": {"type": "choice",
            "question": "Quel est le risque résiduel de ce visuel ?",
            "options": ["aucun", "texte_douteux", "personne_visible", "marque_visible"]},
    })
print("PUBLIER :" if ok else "BLOQUER :", details)
```

## 169. Gestion d'erreurs commune : ce qu'il faut toujours coder

- **429 (rate limit)** : backoff exponentiel + jitter, jamais de retry bourrin (risque de bannissement, cf. leçon support.huawei.com de ton pipeline : ne pas bourriner).
- **Timeouts** : toujours bornés (30 s connexion, 300 s génération) ; un job vidéo peut durer plusieurs minutes, c'est normal.
- **Idempotence** : générer un `idempotency_key` par livrable pour ne pas payer deux fois en cas de retry.
- **Journal** : chaque appel loggue `{fournisseur, modèle, version, prompt_hash, seed, coût_estimé, statut}`. C'est ce journal qui permet la revue mensuelle des coûts (section 73, règle 10).
- **Secrets** : rotation immédiate si une clé apparaît dans un log, un chat ou un commit.

---

# PARTIE T — COMFYUI AVANCÉ : du prototype au batch

## 170. Piloter ComfyUI par API (batch de visuels)

ComfyUI expose une API HTTP : tu peux envoyer des workflows en JSON et récupérer les images — idéal pour générer une série de visuels documentaires sans cliquer.

```python
import json, os, requests, uuid

COMFY = "http://127.0.0.1:8188"

def queue_prompt(workflow: dict):
    r = requests.post(f"{COMFY}/prompt",
                      json={"prompt": workflow,
                            "client_id": str(uuid.uuid4())}).json()
    return r["prompt_id"]

# Exemple : on charge un workflow exporté (format API) et on change le prompt + seed
with open("workflow_flux2_api.json") as f:
    wf = json.load(f)

visuels = [
    ("Vue d'ensemble du local technique avec deux onduleurs 40 kVA", 101),
    ("Détail des borniers d'arrivée, câbles repérés", 102),
    ("Façade onduleur, voyants verts, afficheur", 103),
]
for texte, seed in visuels:
    wf["6"]["inputs"]["text"] = (  # nœud CLIP Text Encode : adapter l'id à ton workflow
        f"photographie technique documentaire : {texte}, lumière neutre "
        f"d'atelier, netteté maximale, style documentation constructeur")
    wf["3"]["inputs"]["seed"] = seed          # nœud KSampler
    pid = queue_prompt(wf)
    print(f"En file : {texte[:40]}... (job {pid}, seed {seed})")
# Récupérer ensuite via /history/{prompt_id} puis /view
```

Méthode : construire le workflow une fois dans l'UI, l'exporter en **format API** (menu), puis le paramétrer en Python. Chaque visuel : seed noté, prompt archivé → reproductibilité totale.

## 171. Workflow inpainting : corriger un défaut sans tout refaire

Le schéma est bon mais une étiquette est en charabia ? Pas besoin de régénérer :

1. Charger l'image dans ComfyUI (`Load Image`).
2. Peindre un **masque** sur la zone à corriger (clic droit sur l'image → « Open in MaskEditor »).
3. Brancher : image + masque → nœud d'inpainting du modèle (FLUX.1 [fill] ou le pipeline inpaint de ton modèle).
4. Prompt ciblé : « étiquette métallique vierge, brossé, sans texte » (le texte sera ajouté en DAO après — piège n°11).
5. Denoise ~0,7-0,85 : assez pour corriger, pas assez pour dénaturer le reste.

En local c'est gratuit et instantané comparé à une re-génération complète — et ça préserve tout ce qui était déjà bon.

## 172. Workflow upscale propre (pas juste « agrandir »)

Chaîne recommandée pour passer un visuel validé en haute résolution :

1. `Load Image` (le visuel validé en 1024²).
2. `Upscale Image (using Model)` → modèle `4x-UltraSharp.pth` (ou équivalent ESRGAN, ~67 Mo).
3. (Option) second passage avec un upscaler « créatif » à faible denoise (0,2-0,3) pour restaurer les micro-détails — attention : à fort denoise, le modèle **invente** des détails (dangereux sur un schéma technique : rester en upscale « fidèle »).
4. `Save Image`.

Règle : **upscaler seulement le visuel validé**, jamais les brouillons (l'upscale coûte du temps GPU et ne sauve pas une mauvaise composition).

## 173. LoRA : spécialiser FLUX sur tes équipements

