---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-10
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["Perplexity"]
dates: ["2026-09-15", "2026-09-27"]
keywords: ["benchmarks", "gpu", "perplexity"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [866, 985]
sha256: eceff40d3edf39658445d4b21e5d8156b418713889d7a04cad885a3eb55dd8bb
---

# IA générative : image, vidéo, recherche

- **Sortie garantie valide par construction** : 0 % de JSON malformé (le cauchemar du « demande à un LLM de classer et renvoie du JSON » disparaît).
- **Latence** : 70 à 500 ms de bout en bout. **Texte uniquement** : pas d'image, pas d'audio, pas de vidéo en entrée.
- **Endpoint** : `POST https://api.typesafe.ai/v1/systemone`, clé Bearer. SDKs **Python et JavaScript**, intégration **Vercel AI Gateway**. Aliases `jev-latest` / `jev-preview` (pointent vers la même version ; **pinner la version** `jev-1.13.0` en prod pour éviter les changements silencieux).
- **Contexte** : 64k tokens par requête (dont 32k pour `state` + la plus longue question).

## 64. Jev : prix et accès (vérifiés le 27/09/2026)

- **0,042 $ par million de tokens d'entrée ; tokens de sortie gratuits** (il ne génère pas de texte : logique).
- Rate limits affichés : 250 000 tokens/s, 1 200 req/min (ajustés selon la demande).
- Accès : **API hébergée, liste d'attente early-access**. **Pas de poids ouverts, pas de self-host, pas de nombre de paramètres publié.**

Ordre de grandeur : qualifier 1 000 tickets ou visuels coûte **quelques centimes** (les benchmarks éditeur annoncent ~0,0004 $/cas — chiffre éditeur, à prendre avec recul, mais l'ordre de grandeur « quasi-gratuit » est crédible vu le prix d'entrée).

## 65. Jev : à quoi ça sert dans TON pipeline (angle Zelef)

Jev ne génère pas tes visuels, mais il peut **garder** ton pipeline de génération :

1. **Contrôle qualité automatique** : après génération d'une série d'images (schémas, visuels de doc), envoyer chaque image *décrite en texte* (légende + prompt + métadonnées) à Jev : « le schéma montre-t-il un onduleur triphasé ? (noul) », « quel défaut potentiel ? (choice: texte illisible / mains déformées / aucun) ». Rejeter sous un seuil de confiance.
2. **Routage** : « ce ticket demande-t-il une image, une vidéo, ou une recherche ? (choice) » → aiguillage automatique vers le bon outil.
3. **Garde-fous** : « ce prompt contient-il une marque identifiable / une personne réelle ? (noul) » avant génération — filtre anti-risque juridique à 0,00004 $ l'appel.
4. **Triage de veille** : sur un digest Perplexity Sonar, scorer la pertinence de chaque item avant de l'envoyer à l'équipe.

Exemple (schéma) :

```python
import os, requests
JEV_KEY = os.environ["JEV_API_KEY"]
r = requests.post("https://api.typesafe.ai/v1/systemone",
    headers={"Authorization": f"Bearer {JEV_KEY}"},
    json={
        "model": "jev-1.13.0",          # pinner la version, pas l'alias
        "state": "Image générée pour la doc 'Remplacement module UPS' : "
                 "prompt='armoire électrique ouverte, borniers alignés', "
                 "modèle=flux-2-pro, seed=42",
        "questions": {
            "montre_borniers": {"type": "noul",
                "question": "La description indique-t-elle des borniers visibles et alignés ?"},
            "defaut": {"type": "choice",
                "question": "Quel est le défaut le plus probable de ce visuel ?",
                "options": ["texte_illisible", "cablage_incoherent", "aucun_defaut"]},
        },
    }).json()
print(r)  # ex. {"montre_borniers": {"p_yes": 0.93}, "defaut": {"choice": "aucun_defaut", "confidence": 0.88}}
```

## 66. Jev : limites honnêtes

- **Ne génère rien** : ni texte, ni image, ni vidéo, ni raisonnement visible. Pour une phrase, un LLM ; pour une décision, Jev.
- **Pas de multimodalité** : il ne « voit » pas tes images — il juge leur **description textuelle**. Le contrôle qualité visuel réel reste humain (ou un modèle vision).
- **Jeune produit** (sorti de stealth le 15/09/2026) : écosystème naissant, prix et limites susceptibles d'évoluer vite, **pas d'option on-prem**.
- **Anglais d'abord** : les autres langues sont documentées comme moins précises — écrire les `state`/`questions` en anglais en prod.
- Ne pas lui demander : calculs exacts, dates, raisonnement multi-étapes, appels d'outils — ce n'est pas son métier.

---

# PARTIE H — COÛTS RÉELS : budgéter sans se faire piéger

## 67. La méthode : partir du livrable, pas du prix facial

Pour chaque besoin, estimer dans l'ordre :

1. **Volume de livrables/mois** (ex. 40 visuels de doc, 6 clips vidéo, 1 digest de veille hebdo).
2. **Taux de rejet** réaliste (×2 à ×3 générations par livrable au début, ×1,5 quand tu maîtrises).
3. **Coût unitaire tout compris** (génération + essais + upscale éventuel).
4. **Frais fixes** (abonnements) vs **variables** (API).
5. Comparer **3 scénarios** : tout-abonnement / tout-API / hybride.

## 68. Exemple chiffré 1 : 40 visuels/mois pour de la doc technique

Hypothèses : 40 images 1024² validées, taux de rejet ×2 (80 générations), usage commercial, confidentialité souhaitée.

| Scénario | Calcul | Coût / mois |
|----------|--------|-------------|
| **FLUX local (klein 4B)** | 0 $ marginal (GPU déjà là, élec. négligeable à cette échelle) | **~0 $** |
| API BFL (FLUX.2 pro) | 80 × 0,03 $ | **2,40 $** |
| fal.ai (schnell brouillon + pro final : 60×0,003 + 20×0,03) | 0,18 + 0,60 | **0,78 $** |
| Midjourney Basic | abonnement | **10 $** (illimité relax : non, Basic n'a pas le relax) |
| Nano Banana Pro (API) | 80 × 0,134 $ | **10,72 $** |

Verdict : pour de la doc technique interne, **le local ou l'API BFL écrasent tout**. Midjourney ne se justifie que si l'esthétique « belle image » est le besoin.

## 69. Exemple chiffré 2 : 6 clips vidéo/mois (présentation d'équipement)

Hypothèses : 6 clips 10 s 1080p validés, ×3 essais (18 générations), audio natif souhaité.

| Scénario | Calcul | Coût / mois |
|----------|--------|-------------|
| Kling web **Pro** (25,99 $) | 18 gén × ~35-70 crédits (pro 10 s ± audio) ≈ 630-1 260 crédits < 3 000 inclus | **25,99 $** (marge restante) |
| Kling **API** (pack Trial 98 $/1 000 u) | 18 × 10 s × 0,168 $/s (1080p+audio) ≈ 30 $ | **~30 $** sur le pack (validité 30 j !) |
| Veo 3.1 Standard (API) | 18 × 10 s... 8 s max → 18 × 8 × 0,40 $ | **57,60 $** |
| Veo 3.1 Fast | 18 × 8 × 0,12 $ | **17,28 $** |
| Runway Standard (15 $) | 625 crédits ; 10 s Gen-4.5 ≈ 250 crédits → 2 clips max | **insuffisant** → Pro 35 $ |

Verdict : **Kling Pro web** est le sweet spot pour ce volume ; l'API Kling ne devient intéressante qu'au-delà (ou via fal.ai à l'usage pur, sans pack qui expire).

## 70. Exemple chiffré 3 : veille automatisée (API Sonar)

Hypothèses : 1 digest hebdomadaire = ~20 questions de veille (4 semaines = 80 requêtes/mois).

| Scénario | Calcul | Coût / mois |
|----------|--------|-------------|
| `sonar` | 80 × (~2 000 tokens à 1 $/1M + frais req. ~5 $/1K) | **~0,50 $** |
| `sonar-pro` | 80 × (~5 000 tokens à 3-15 $/1M + frais req. ~14 $/1K) | **~2-5 $** |
| Abonnement Pro (20 $) + Spaces | inclus | **20 $** (mais tu as tout le reste) |

Verdict : la veille automatisée via API coûte **quelques dollars par mois**. L'abonnement Pro se justifie par l'usage interactif quotidien, pas par la veille seule.

## 71. Le vrai coût caché : les packs qui expirent

Récapitulatif des « dates de péremption » (vérifié sept 2026) :

| Fournisseur | Ce qui expire | Délai |
|-------------|--------------|-------|
| Kling **API** (packs) | Unités non consommées | **30 ou 180 jours** selon le pack, sans report |
| Kling **web** (abonnement) | Crédits mensuels | Fin de période, sans report |
| Luma | Crédits mensuels | Fin de période (seuls les Top-Up durent 12 mois) |
| Midjourney | Heures GPU | Fin de période, sans report |
| BFL (API) | Crédits pré-chargés | Selon CGU (**à vérifier** : pas de durée trouvée au 27/09/2026) |
| fal.ai / Replicate | Facturation à l'usage | **Pas d'expiration** (tu paies ce que tu consommes) |

