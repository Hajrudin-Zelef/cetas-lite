---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-21
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["perplexity"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [2170, 2314]
sha256: 03f9c96d39a02d457b0d6779b797daa1cf4bda16ecb421b64533fdd1b2189855
---

# IA générative : image, vidéo, recherche

journal = []
for nom, prompt, seed in JOBS:
    try:
        path = generer(nom, prompt, seed)
        journal.append([datetime.now().isoformat(), nom, "flux-2-pro", seed,
                        "OK", COUT_PAR_IMAGE, path, prompt[:80]])
        print(f"OK  {nom} -> {path}")
    except Exception as e:   # on continue le batch malgré un échec
        journal.append([datetime.now().isoformat(), nom, "flux-2-pro", seed,
                        f"ECHEC: {e}", 0, "", prompt[:80]])
        print(f"KO  {nom} : {e}")

with open("journal_batch.csv", "a", newline="") as f:
    csv.writer(f).writerows(journal)
total = sum(float(l[5]) for l in journal)
print(f"Terminé : {len(journal)} jobs, coût estimé {total:.2f} $ — voir journal_batch.csv")
# Rappel : relire chaque visuel (textes !) avant usage — piège n°11.
```

## 182. Script : digest de veille multi-sources (Sonar + déduplication)

```python
#!/usr/bin/env python3
"""Digest hebdo : interroge Sonar sur N sujets, déduplique les URLs déjà vues.
Usage : PERPLEXITY_API_KEY=... python3 veille.py
État : urls_vues.txt conserve l'historique (ne jamais alerter 2x sur la même URL).
"""
import os, re
from openai import OpenAI

client = OpenAI(api_key=os.environ["PERPLEXITY_API_KEY"],
                base_url="https://api.perplexity.ai")
VU = "urls_vues.txt"
vues = set(open(VU).read().splitlines()) if os.path.exists(VU) else set()

SUJETS = [
    "nouveautés onduleurs triphasés 10 à 120 kVA",
    "batteries lithium-ion pour onduleurs : sécurité et rappels",
    "failles CVE équipements d'énergie et building management connectés",
    "normes électriques NF C 15-100 / NF C 18-510 / IEC 62040 : évolutions",
]

def ask(s):
    r = client.chat.completions.create(
        model="sonar",
        messages=[{"role": "user",
                   "content": f"Actualités des 7 derniers jours : {s}. "
                              "Format : une puce par info, avec titre, 1 ligne de "
                              "résumé, et l'URL source entre chevrons <url>. "
                              "Maximum 6 puces. Si rien de notable : 'RAS'."}])
    return r.choices[0].message.content

nouvelles_urls = set()
sortie = []
for s in SUJETS:
    texte = ask(s)
    urls = set(re.findall(r"<(https?://[^>]+)>", texte))
    fraiches = urls - vues
    nouvelles_urls |= urls
    if fraiches:
        sortie.append(f"## {s}\n" + "\n".join(
            l for l in texte.splitlines()
            if not any(u in l for u in (urls - fraiches)) or "RAS" in l))
    else:
        sortie.append(f"## {s}\nRAS ou déjà signalé cette semaine.")

with open(VU, "a") as f:
    f.write("\n".join(sorted(nouvelles_urls - vues)) + "\n")

print("VEILLE ÉNERGIE — " + __import__("datetime").date.today().isoformat())
print("\n\n".join(sortie))
print(f"\n---\nURLs déjà connues ignorées : {len(vues)} en historique.")
# À brancher sur cron le lundi 7h + envoi mail (voir section 103).
```

## 183. Script : contrôle qualité d'un dossier de visuels (Jev)

```python
#!/usr/bin/env python3
"""Passe chaque visuel d'un dossier au crible Jev (métadonnées textuelles).
Usage : JEV_API_KEY=... python3 qc_visuels.py ./livrables
Principe : Jev ne voit pas l'image, il juge sa fiche (prompt, modèle, relecture).
"""
import json, os, sys, requests

JEV = "https://api.typesafe.ai/v1/systemone"
H = {"Authorization": f"Bearer {os.environ['JEV_API_KEY']}",
     "Content-Type": "application/json"}
SEUIL = 0.70

def check(fiche):
    r = requests.post(JEV, headers=H, timeout=20, json={
        "model": "jev-1.13.0",
        "state": fiche,
        "questions": {
            "licence_ok": {"type": "noul", "question":
                "La licence du modèle cité autorise-t-elle l'usage commercial ?"},
            "risque": {"type": "choice", "question":
                "Quel est le risque principal de ce visuel ?",
                "options": ["aucun", "texte_douteux", "personne_reconnaissable",
                            "marque_visible", "contenu_sensible"]},
        }}).json()
    lic = r["licence_ok"]["p_yes"] >= SEUIL
    risq = r["risque"]
    ok = lic and risq["choice"] == "aucun" and risq["confidence"] >= SEUIL
    return ok, {"licence_p": round(r["licence_ok"]["p_yes"], 3),
                "risque": risq["choice"],
                "confiance": round(risque_conf(risque), 3)}

def risque_conf(x):
    return x.get("confidence", 0)

dossier = sys.argv[1]
# Convention : chaque visuel a sa fiche JSON {prompt, modele, licence, seed, relecture}
for fn in sorted(os.listdir(dossier)):
    if not fn.endswith(".json"):
        continue
    fiche = json.load(open(os.path.join(dossier, fn)))
    texte = (f"Visuel {fn}: prompt='{fiche['prompt']}', modèle={fiche['modele']} "
             f"(licence {fiche['licence']}), seed={fiche['seed']}. "
             f"Relecture humaine : {fiche['relecture']}")
    ok, det = check(texte)
    print(("PASS " if ok else "REVIEW ") + fn, det)
# Coût : ~0,0004 $/visuel. Les REVIEW repartent en relecture humaine.
```

## 184. Matrice de décision fournisseur (à remplir pour TON cas)

Recopie ce tableau, mets tes chiffres, décide en 10 minutes :

| Critère (poids) | Option A : ______ | Option B : ______ | Option C : ______ |
|-----------------|-------------------|-------------------|-------------------|
| Coût / livrable accepté (×3) | | | |
| Qualité sur mon besoin (×3) | | | |
| Confidentialité des données (×2) | | | |
| Vitesse / file d'attente (×1) | | | |
| Intégration API (×2) | | | |
| Licence commerciale claire (×2) | | | |
| Risque fournisseur (×1) | | | |
| **Score pondéré** | | | |

Règle : le « coût / livrable accepté » se **mesure** (2 semaines de test avec compteur), il ne se lit pas sur une page tarifaire. Le gagnant du tableau n'est valable que pour **ton** besoin : refaire l'exercice par besoin (doc, vidéo, veille), pas une fois pour tout.

## 185. Fiche réflexe imprimable (1 page — à afficher)

