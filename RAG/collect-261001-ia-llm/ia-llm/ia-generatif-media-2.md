---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-2
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["Google"]
dates: ["2026-09-27"]
keywords: ["sol", "text-to-image", "text-to-video"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [96, 202]
sha256: 3f039e2f787c99019429352c9864abc2e0953ba9a757c2c081f457702637529b
---

# IA générative : image, vidéo, recherche

- **Text-to-video et image-to-video** : 3 à 15 secondes par génération.
- **Résolution jusqu'à 4K native** (le plafond le plus haut du marché grand public en sept 2026).
- **Audio natif** : génération du son en même temps que l'image, avec **lip-sync multilingue** et dialogues multi-personnages. L'audio coûte environ +50 % sur le tarif API (voir section 11).
- **Motion transfer** (transfert de mouvement) : tu fournis une vidéo de référence, Kling extrait le motif de mouvement et l'applique à un autre sujet. Différenciateur unique début 2026, toujours distinctif en septembre.
- **Multi-prompt** : jusqu'à 15 secondes découpées en segments temporisés, chacun avec son propre prompt (512 caractères max par segment), rendus en un seul plan continu. Idéal pour les mini-scénarios.
- **Start frame + end frame** : tu imposes la première et la dernière image ; Kling 3.0 remplit l'intervalle. (Pas d'images de référence intermédiaires sur 3.0 : pour le verrouillage d'identité, passer par Kling O1.)
- **AI Director** : planification automatique de l'enchaînement des plans.
- **Rendu de texte à l'écran** : le meilleur du marché pour garder lisibles pancartes, logos, étiquettes de prix dans la vidéo — point décisif pour la vidéo produit/e-commerce.
- **Cohérence du sujet** : jusqu'à 4 images de référence pour verrouiller un personnage (selon le mode/endpoint).

Limites : Kling 3.0 prend **plus de libertés créatives** que certains concurrents (Seedance 2.x) : moins fiable quand le brief exige un mouvement précis et prédéterminé. Files d'attente > 30 min aux heures de pointe sur l'interface web (l'API priorise mieux).

## 9. Kling : prise en main (interface web)

1. Aller sur `klingai.com`, créer un compte (e-mail ou Google).
2. Le plan **gratuit** donne **66 crédits/jour** (sans report), sorties **watermarkées**, **usage non commercial uniquement**.
3. Choisir le mode : Text-to-Video ou Image-to-Video.
4. Régler : modèle (3.0 / 2.6...), mode (Standard / Professional), durée (5 ou 10 s, extensible), résolution (720p / 1080p / 4K), audio natif ou non.
5. Écrire le prompt (voir sections 113-122 pour la méthode), lancer, attendre, télécharger.

Le plan gratuit suffit pour évaluer la qualité, pas pour produire : watermark + interdiction d'usage commercial (voir section 10).

## 10. Kling : offres et prix web (vérifiés le 27/09/2026)

| Plan | Prix mensuel | Prix « remisé » 1er abonnement | Crédits/mois | Notes |
|------|--------------|-------------------------------|--------------|-------|
| Free | 0 $ | — | 66/jour | Watermark, **non commercial** |
| Standard | 10 $ | **6,99 $** | 660 | Sans watermark, modèles standard |
| Pro | 37 $ | **25,99 $** | 3 000 | File prioritaire, tous modèles |
| Premier | 92 $ | **64,99 $** | 8 000 | Mode professionnel |
| Ultra | 180 $ | **127,99 $** | 26 000 | Priorité max |

- Facturation annuelle : **-34 %** environ sur Standard/Pro/Premier. Ultra : pas d'option annuelle (vérifié mi-2026).
- **Les crédits payants expirent en fin de période, sans report.** C'est le piège n°1 : ne surdimensionne pas ton plan.
- Coûts de génération constatés (interface web) :
  - 5 s : **10 crédits** (Standard) / **35 crédits** (Professional)
  - 10 s : **20 crédits** (Standard) / **70 crédits** (Professional)
  - Avec audio natif : coût multiplié par ~2,5 à 5 selon version (ex. 5 s : 50 crédits Standard / 100 crédits Pro sur la grille v2.6 ; sur 3.0 compter **6 à 12 crédits/seconde** en résolution standard, **30 crédits/seconde** en 4K native).
- Génération d'image (Kolors) : 1 crédit/image en text-to-image, 2 en restyle, 5 en essayage virtuel.

**Droit d'usage commercial :** tous les plans payants l'autorisent. Le plan gratuit l'interdit explicitement (mention sur la page tarifaire officielle).

## 11. Kling : l'API développeur (oui, elle existe)

Accès : portail développeur **`kling.ai/dev`**. Points essentiels, vérifiés le 27/09/2026 :

- **Facturation totalement séparée de l'abonnement web.** Un plan Pro web ne donne aucun crédit API. Les packs API ne servent pas sur l'interface web.
- **Packs prépayés** (pas d'abonnement API, pas de facturation à l'usage pur) :

| Pack | Prix | Unités | Validité |
|------|------|--------|----------|
| Trial (petit) | 9,80 $ | 100 | 30 jours |
| Trial (grand) | 98 $ | 1 000 | 30 jours |
| Standard 1 | 700 $ | 5 000 | 180 jours |
| Standard 2 | 2 100 $ | 15 000 | 180 jours |
| Standard 3 | 4 200 $ | 30 000 | 180 jours |
| Large | 7 560 $ | 60 000 | 180 jours |

- **Sans report, sans extension** à l'expiration. Un pack de 700 $ non consommé en 180 jours = 700 $ perdus. Pour un usage ponctuel, le pack Trial à 9,80 $ est le seul rationnel.
- Tarif facial : **0,14 $/unité**. Tarifs par seconde constatés (sept 2026, page développeur) :
  - Kling 3.0, 720p, sans audio : **0,084 $/s** ; avec audio : **0,126 $/s**
  - Kling 3.0, 1080p, sans audio : **0,112 $/s** ; avec audio : **0,168 $/s**
  - Kling 3.0 Turbo + audio : 0,112 $/s (720p) / 0,140 $/s (1080p)
- Fonctionnement : **asynchrone**. Tu POSTes la demande → tu reçois un `job_id` → tu interroges (polling) jusqu'à `completed` (compter ~2 min par clip en pratique). Les **callbacks/webhooks** sont supportés : préfère-les au polling en production (moins d'appels, même coût, meilleure hygiène).
- **Accès tiers** : Kling 3.0 est aussi exposé via des agrégateurs (**fal.ai**, Replicate, Krea...) avec leur propre facturation (ex. fal : ~0,112 $/s audio off, 0,168 $/s audio on, 0,196 $/s avec contrôle vocal — tarifs fal, pas tarifs Kling officiels). Intérêt : un seul compte/clé pour plusieurs modèles, facturation à l'usage sans pack prépayé. Inconvénient : marge de l'agrégateur, conditions propres.

## 12. Kling : exemple d'appel API (schéma)

```python
import os, time, requests

API_KEY = os.environ["KLING_API_KEY"]          # jamais en dur, jamais commitée
BASE = "https://api.klingai.com"               # hôte illustratif : voir la doc dev officielle

headers = {"Authorization": f"Bearer {API_KEY}", "Content-Type": "application/json"}

# 1. Lancer une génération text-to-video (Kling 3.0, 10 s, 1080p, sans audio)
job = requests.post(f"{BASE}/v1/videos/text-to-video", headers=headers, json={
    "model": "kling-3.0",
    "prompt": "travelling avant lent dans une salle serveurs, rangées de baies "
              "19 pouces avec LEDs bleues, sol technique gris, lumière froide, "
              "ambiance data center professionnel, photoréaliste",
    "duration": 10,
    "resolution": "1080p",
    "audio": False,
}).json()
job_id = job["job_id"]

# 2. Polling (en prod : préférer le webhook / callback)
while True:
    st = requests.get(f"{BASE}/v1/videos/{job_id}", headers=headers).json()
    if st["status"] == "completed":
        print("Vidéo prête :", st["video_url"])
        break
    if st["status"] == "failed":
        raise RuntimeError(f"Échec génération : {st.get('error')}")
    time.sleep(10)
```

> Note d'honnêteté : les chemins d'endpoint exacts évoluent ; **la doc officielle `kling.ai/dev` fait foi**. Le schéma ci-dessus (job asynchrone + polling) est le pattern stable de toutes les API vidéo.

## 13. Kling : prompts vidéo efficaces — méthode

Un bon prompt vidéo = **sujet + action + caméra + lumière + ambiance + style**, dans cet ordre. Exemple décortiqué :

> « [sujet] un technicien en tenue de travail bleu marine [action] remplace un module d'onduleur dans une armoire électrique ouverte [caméra] plan moyen, léger travelling latéral gauche-droite [lumière] lumière d'atelier froide, reflets métalliques [ambiance] calme professionnel, gestes précis [style] photoréaliste, 24 i/s, netteté documentaire »

