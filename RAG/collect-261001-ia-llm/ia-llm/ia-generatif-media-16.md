---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-16
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["ByteDance", "Moonshot", "OpenAI", "Perplexity"]
dates: ["2026-03-24", "2026-04-26", "2026-07-16", "2026-07-20", "2026-09-03", "2026-09-15", "2026-09-23", "2026-09-24", "2026-09-25", "2026-09-27"]
keywords: ["agents", "apache", "attention", "fp8", "gpu", "ipo", "kimi", "perplexity"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [1572, 1638]
sha256: 7f683cf404b25d1690dc139e0a866ee9d0a6abf4f5918af719b9f13497095278
---

# IA générative : image, vidéo, recherche

- **Automatiser la génération** : ComfyUI expose une **API HTTP** (file de prompts en JSON) — piloter tes batches d'images depuis un script, avec seeds et logs.
- **RAG + visuels** : indexer tes visuels générés (avec prompts/seeds en métadonnées) dans ton RAG : « retrouve le visuel de la procédure X » devient une recherche.
- **Supervision des coûts** : un export hebdo des consoles (BFL, Kling, Perplexity, fal.ai) vers un tableur ; alerte si coût/livrable dérive de >20 %.
- **GPU local** : pour FLUX.2 dev en FP8, une RTX 4090/5090 24 Go d'occasion ou un GPU cloud à l'heure (RunPod & co) selon ton volume — calculer le seuil de rentabilité (section 68).

## 155. Pour creuser côté métier (systèmes & énergies)

- **Banque de visuels d'équipements** : constituer ta bibliothèque de références (photos réelles + prompts validés) par famille (onduleurs, TGBT, groupes, clim) — c'est un actif qui s'amortit à chaque doc.
- **Gabarits vidéo** : un template de montage (intro 3 s, cartons, outro) réutilisé pour chaque équipement — la génération IA ne fait que les plans.
- **Veille réglementaire** : Space Perplexity « Normes & énergie » partagé à l'équipe (section 34) + digest API hebdo (section 103).
- **Formation** : clips courts (Kling/Hailuo) pour illustrer les gestes sécurité — plus mémorables qu'un paragraphe, moins chers qu'un tournage.

---

# PARTIE P — À VENIR : annonces vérifiées au 27/09/2026

## 156. Annonces officielles récentes (sourcées)

- **24/09/2026 — OpenAI ferme la Videos API (Sora 2).** Annonce initiale le 24/03/2026 (6 mois de préavis), appli fermée le 26/04/2026, API coupée le 24/09/2026, **aucun remplaçant annoncé**. Source : page de dépréciation officielle OpenAI (developers.openai.com), reprise par la presse tech le 24-25/09/2026.
- **15/09/2026 — TypeSafe AI sort Jev de stealth.** Premier modèle public `jev-1.13.0` (« System One »), levée seed 40 M$ (Forbes), API `api.typesafe.ai`. Source : presse tech (ADTmag 23/09/2026, The Rundown) et documentation développeur.
- **16/07/2026 — Moonshot AI publie Kimi K3** (poids ouverts, 2 800 Md de paramètres revendiqués), puis **suspend les nouveaux abonnements** (saturation GPU, 19-20/07/2026, Reuters). IPO de Hong Kong en préparation (~3 Md$ visés, Reuters 03/09/2026).
- **Février 2026 — Kling 3.0** (4K native, audio natif, motion transfer) et **Seedance 2.0** (ByteDance) : les deux lancements qui ont structuré le marché vidéo 2026.
- **Février 2026 — Perplexity coupe sa régie publicitaire** (lancée octobre 2025) : modèle 100 % abonnement/API.
- **Novembre 2025 — FLUX.2** (Black Forest Labs) : génération actuelle, avec klein 4B Apache 2.0.

## 157. Rumeurs (non confirmées — à ne pas prendre pour des faits)

- **RUMEUR non confirmée** : évolutions futures de la numérotation Kling (au-delà de 3.0), de Midjourney (au-delà de V8.2) ou de Veo (au-delà de 3.1) — **rien d'officialisé au 27/09/2026**.
- **RUMEUR non confirmée** : une API publique Midjourney grand public (seule une offre Enterprise sur mesure est évoquée par des sources tierces).
- **RUMEUR non confirmée** : réouverture des abonnements Kimi (dépend des capacités GPU, aucune date annoncée).
- Principe : tant qu'il n'y a pas d'annonce officielle (blog éditeur, page de dépréciation, presse vérifiée), c'est une rumeur. Les prix et versions de ce guide sont un **instantané du 27/09/2026**, pas une promesse.

## 158. Signaux faibles à surveiller (pas des annonces)

- **FLUX 3 vidéo** : une documentation tierce (sept 2026) fait état d'une offre vidéo BFL au compteur (~0,17 $/s en HD text/image-to-video) — **non confirmé par une annonce officielle BFL au 27/09/2026, à vérifier**. Si ça se confirme, BFL deviendrait un acteur vidéo à part entière.
- **Agrégateurs multi-modèles** (fal.ai, Replicate, Krea, Luma Agents) : la tendance est au **guichet unique** — un seul compteur pour Kling + Veo + Seedance + FLUX. À terme, ça peut changer l'équation « API officielle vs tiers ».
- **Contentieux éditeurs vs IA** (Perplexity et autres) : l'issue (accords de licence ? retraits de sources ?) impactera la qualité des moteurs de réponse. À suivre trimestriellement.

---

# PARTIE Q — NOMS NON IDENTIFIÉS

## 159. Verdict d'identification (recherche du 27/09/2026)

Deux noms étaient incertains dans la demande. Après recherche web sérieuse :

| Nom demandé | Verdict |
|-------------|---------|
| **« plexplxity »** | **Coquille pour Perplexity** — traité en Partie C (sections 28-34). Aucun outil distinct trouvé sous ce nom. |
| **« jev »** | **Identifié : Jev, modèle de décision de TypeSafe AI** (« System One »), sorti de stealth le 15/09/2026 — traité en Partie G (sections 62-66). **Attention** : ce n'est pas un outil d'image/vidéo. |

**Aucun nom n'est resté non identifié.** Il n'y a donc pas de section « NON TROUVÉ » à proprement parler — les deux pistes ont abouti. (Si tu rencontres d'autres noms obscurs, la méthode est la même : recherche web ciblée + vérification de l'éditeur avant d'inventer quoi que ce soit.)

---

# PARTIE R — SOURCES & DATES DE VÉRIFICATION

## 160. Comment ce guide a été vérifié

- **Date de vérification unique : 27/09/2026.** Recherche web ciblée par outil (existence, version, features, prix, API, licence).
- **Règle appliquée** : tout nom introuvable après recherche sérieuse aurait été marqué « NON TROUVÉ (sept 2026) » ; **aucune spec, aucun prix, aucune version n'a été inventé** — les points non confirmés sont marqués « à vérifier » ou « RUMEUR non confirmée ».
- **Limite assumée** : les pages tarifaires officielles changent vite et certaines n'ont été consultées qu'à travers la presse spécialisée et des documentations tierces récentes (sept 2026). **Avant d'engager un budget, re-vérifier la page officielle.**
- **Aucune clé API réelle** dans ce guide : tous les exemples utilisent des variables d'environnement.
- Volumes à croiser : `ia_modeles_chine.md`, `ia_modeles_occident.md`, `ia_modeles_specialises.md` (modèles de langage et spécialisés — non dupliqués ici).

## 161. Journal des vérifications (27/09/2026)

