---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-20
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["EU", "Google", "Perplexity"]
dates: []
keywords: ["gpu", "lora", "perplexity", "text-to-video"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [2037, 2169]
sha256: e68cfe607cc92005c8958b433c031dc0b6a32ece745960f76ee81484f56b0266
---

# IA générative : image, vidéo, recherche

**V10 — Teaser événement (Runway, Aleph + Gen-4.5)**
> Générer 3 plans (Gen-4.5), puis assembler et rythmer dans l'éditeur (Aleph) avec titrage.
*Runway se justifie quand tu restes dans sa suite de bout en bout.*

---

# PARTIE U — FAQ : 20 questions qu'on pose vraiment

## 177. Questions générales

**F1. Par quoi commencer si je n'ai jamais rien généré ?**
Par **Perplexity gratuit** (recherche, zéro risque) puis **FLUX.1 [schnell] en local ou via fal.ai** (~0,003 $/image) pour l'image, et le **plan gratuit de Kling** (66 crédits/jour) pour la vidéo. Tu auras tout testé pour moins d'1 $.

**F2. Lequel choisir entre abonnement et API ?**
Usage **manuel et régulier** → abonnement web. Usage **programmé / en volume / irrégulier** → API (officielle si volume, agrégateur si ponctuel). Voir sections 67-73.

**F3. C'est vraiment utilisable pour du travail sérieux, ou c'est du gadget ?**
Les deux. Gadget : générer « pour voir ». Sérieux : **doc technique illustrée, clips de formation, veille automatisée, maquettes avant travaux** — les 4 cas de la Partie K sont du travail réel, pas de la démo.

**F4. Combien de temps pour être autonome ?**
Une **demi-journée** par outil pour les bases (prompts + réglages + export), une **semaine d'usage réel** pour les réflexes (taux de rejet, coûts, limites). ComfyUI en local : compter une journée d'installation et de prise en main.

**F5. Et si je n'ai pas de GPU ?**
Image : API (BFL dès 0,014 $/image, fal.ai dès 0,003 $). Vidéo : 100 % cloud de toute façon. Recherche : 100 % cloud. Le GPU local n'est **obligatoire pour rien** — c'est un choix de confidentialité et de coût au volume.

## 178. Questions image

**F6. Midjourney ou FLUX pour mes docs techniques ?**
**FLUX** : plus neutre/documentaire, texte mieux géré (v2), local possible, 10× moins cher. Midjourney si tu veux du « beau » (com, affiches) plutôt que du « juste ».

**F7. Pourquoi mes textes générés sont-ils en charabia ?**
Parce que tu demandes un **long texte fin** à un modèle qui le « dessine » au lieu de l'écrire. Solutions : texte court entre guillemets, modèle fort en typo (Ideogram, Nano Banana Pro, FLUX.2), ou **texte ajouté en PAO après** (la vraie solution pro).

**F8. Comment garder le même équipement d'une image à l'autre ?**
FLUX.2 **multi-références** (3-6 photos de l'équipement en entrée), description verrouillée, seed noté. Au-delà de ~20 visuels : **LoRA** (section 173).

**F9. Les images gratuites sont-elles utilisables au travail ?**
**Non** : watermarks + interdiction d'usage commercial sur les plans gratuits (Kling, Luma...). Le gratuit sert à **évaluer**, pas à produire.

**F10. Quelle résolution pour quel usage ?**
Web/intranet : 1024² suffit. Présentation projetée : 2048². Print (affiche A3) : générer grand ou **upscaler** (section 172). Ne jamais générer en 4K « au cas où » (coût ×3-5).

## 179. Questions vidéo

**F11. Text-to-video ou image-to-video ?**
**Image-to-video** dès que tu as une bonne photo de départ (équipement réel, site réel) : plus crédible, moins d'aléas. Text-to-video pour l'imaginaire (ambiances, concepts).

**F12. Pourquoi ma vidéo est-elle « bizarre » (mains, visages, physique) ?**
Limites connues : gestes fins, mains, foules, textes longs, mouvements imposés au millimètre. Contourner : plans larges, actions simples, **une action par clip**, I2V depuis le réel, Seedance pour le geste précis.

**F13. L'audio natif vaut-il le coup ?**
Pour un dialogue ou une ambiance impossible à refaire : oui. Sinon **non** : +50 % à ×5 du coût pour un son que tu referas mieux en montage (voix off, bruitages libres).

**F14. Quelle durée par clip ?**
5-10 s : le standard (un geste, un plan). 15 s : mini-scénario (multi-prompt). 30 s : Seedance en un appel, ou **montage** de 3-4 clips (mieux : rythme contrôlé, plans variés).

**F15. Où monter les clips ?**
Pas dans le générateur : **DaVinci Resolve** (gratuit, pro), Premiere, ou même CapCut pour du simple. Le générateur fait les **plans**, le monteur fait le **film** (titres, voix, musique, rythme).

## 180. Questions recherche & API

**F16. Perplexity peut-il remplacer Google ?**
**Non** : il le **complète**. Perplexity dégrossit vite avec des sources ; Google reste roi pour la recherche obscure, les opérateurs fins (`site:`, dates, forums) et le gratuit. Les pros utilisent les deux.

**F17. L'API Sonar est-elle rentable pour une petite structure ?**
Oui : une veille hebdo coûte **~0,50 $/mois** en `sonar` (section 103). Le coût n'est jamais le problème ; la **qualité du prompt de veille** et la **relecture** le sont.

**F18. Jev remplace-t-il un LLM dans mon pipeline ?**
**Non** : il le **complète**. Le LLM produit (texte, code, résumé), Jev **vérifie/route/score** en 100-500 ms pour des fractions de centime. Pattern : cascade vérifiée (section 168).

**F19. Que faire quand un fournisseur change ses prix ?**
Re-calculer ton **coût par livrable** (section 5), comparer les 2-3 alternatives de ce guide, et migrer si l'écart dépasse ~30 %. Ne jamais migrer « parce que c'est moins cher au tarif facial » sans mesurer le taux de rejet.

**F20. Comment rester à jour sans y passer mes soirées ?**
Rituel trimestriel (1 h) : re-vérifier les 4-5 tarifs que tu utilises, lire les notes de version des modèles, parcourir la section « À venir » de ce guide. Le reste du temps : **utiliser** les outils, pas les suivre.

---

# PARTIE V — SCRIPTS COMPLETS PRÊTS À ADAPTER

## 181. Script : batch d'images doc via l'API BFL (avec suivi des coûts)

```python
#!/usr/bin/env python3
"""Génère une série de visuels documentaires via l'API BFL (UE), avec journal des coûts.
Usage : BFL_API_KEY=... python3 batch_bfl.py
Coût estimé : 0,03 $/image en flux-2-pro (1 MP) — voir section 19.
"""
import csv, os, time, requests
from datetime import datetime

BFL_KEY = os.environ["BFL_API_KEY"]
EU = "https://api.eu.bfl.ai"   # résidence UE
H = {"x-key": BFL_KEY, "Content-Type": "application/json"}
COUT_PAR_IMAGE = 0.03          # flux-2-pro, ~1 MP : adapter au modèle

JOBS = [
    # (nom_fichier, prompt, seed)
    ("tgbt_vue_ensemble",
     "photographie technique documentaire : armoire TGBT industrielle ouverte, "
     "disjoncteurs modulaires alignés, borniers, lumière neutre d'atelier, "
     "cadrage frontal, netteté maximale, style documentation constructeur", 201),
    ("tgbt_detail_borniers",
     "macrophotographie : borniers à vis sur rail DIN, câbles repérés par "
     "étiquettes, cuivre net, lumière latérale d'atelier, style documentation "
     "constructeur", 202),
    ("local_energie_isometrique",
     "illustration technique épurée en isométrique : local énergie avec onduleur "
     "40 kVA, armoire TGBT, climatisation de précision, fond clair, lignes fines "
     "bleu marine, sans texte, style manuel d'ingénierie", 203),
]

def generer(nom, prompt, seed, timeout_s=300):
    sub = requests.post(f"{EU}/v1/flux-2-pro", headers=H, timeout=30,
                        json={"prompt": prompt, "width": 1024, "height": 1024,
                              "seed": seed}).json()
    tid = sub["id"]
    t0 = time.time()
    while time.time() - t0 < timeout_s:
        res = requests.get(f"{EU}/v1/get_result", headers=H,
                           params={"id": tid}, timeout=30).json()
        if res["status"] == "Ready":
            url = res["result"]["sample"]
            data = requests.get(url, timeout=60).content
            path = f"{nom}.png"
            with open(path, "wb") as f:
                f.write(data)
            return path
        if res["status"] in ("Error", "Failed"):
            raise RuntimeError(f"Échec {tid} : {res}")
        time.sleep(3)
    raise TimeoutError(tid)

