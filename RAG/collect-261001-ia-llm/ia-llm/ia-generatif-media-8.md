---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-8
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "ByteDance", "Google", "MiniMax", "OpenAI"]
dates: ["2026-03-24", "2026-04-26", "2026-09-24", "2026-09-27"]
keywords: ["agents", "arr", "attention", "gemini", "open weights"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [705, 773]
sha256: 6121fcb7d56681eb46172c7f40ab639245110dacfbd0769da04dd33007dad9b2
---

# PARTIE E — PANORAMA VIDÉO : les autres acteurs

## 45. Sora : l'avertissement (API fermée le 24/09/2026)

Rappel factuel, vérifié le 27/09/2026 :

- **24/03/2026** : OpenAI annonce l'arrêt (6 mois de préavis, politique standard de dépréciation).
- **26/04/2026** : fermeture de l'appli et du web Sora.
- **24/09/2026** : **fin de la Videos API** — modèles `sora-2`, `sora-2-pro` et tous les snapshots datés **retirés de l'API, sans remplaçant annoncé** (colonne « remplacement recommandé » : vide).
- Les contenus devaient être exportés avant la fermeture (suppression définitive après la fenêtre d'export).

**Ce que ça t'apprend** : (1) pinner les versions d'API et suivre les pages de dépréciation ; (2) isoler l'appel vidéo derrière une interface interne pour pouvoir changer de fournisseur ; (3) ne jamais laisser des assets uniquement chez le fournisseur. Les alternatives de migration citées par la presse : **Veo 3.1** (même écosystème Google si tu y es déjà) et **Kling 3.0** (qualité/prix).

## 46. Google Veo 3.1 : le haut de gamme

- Accès : **Gemini API** et **Vertex AI** (même SDK/facturation que le reste de Google — gros avantage si tu es déjà client GCP).
- Points forts : **réalisme cinématographique**, **audio natif synchronisé**, dialogues (12 langues de script selon les tests tiers), contrôles de caméra. Le choix « premium » quand le budget le permet.
- Tarifs API vérifiés (sept 2026, par seconde) :

| Variante | 720p / 1080p | 4K |
|----------|--------------|-----|
| Standard (+audio) | **0,40 $/s** | 0,60 $/s |
| Fast | 0,10 $ / 0,12 $/s | 0,30 $/s |
| Lite | 0,05 $ / 0,08 $/s | — |

Un clip 8 s 1080p : **3,20 $** (Standard), 0,96 $ (Fast), 0,64 $ (Lite). Les variantes sont des modèles différents, pas des paliers de qualité équivalente : Fast/Lite sont visiblement en retrait.
- Clips de **8 s** max. Pas d'offre gratuite API (l'accès passe par les abonnements Google AI côté grand public).
- Idéal pour : le clip « vitrine » (présentation d'entreprise, vidéo d'appel d'offres), là où la qualité justifie 3 $ les 8 secondes.

## 47. Runway Gen-4.5 : la plateforme des pros de l'image

- Positionnement : **suite de production** plus que simple générateur — Gen-4.5 (génération), **Aleph** (montage/édition IA), **Act-Two** (transfert de jeu d'acteur), **Motion Brush 2.0**, Director Mode, voix custom.
- Tarifs (vérifiés sept 2026) : **Standard 15 $/mois** (625 crédits), **Pro 35 $/mois** (2 250 crédits), **Unlimited 95 $/mois** (2 250 crédits + générations relax illimitées sur certains outils) ; annuel ≈ -20 % (12/28/76 $).
- Consommation : **Gen-4.5 ≈ 25 crédits/seconde** → le plan Standard ne produit qu'**~25 secondes** de Gen-4.5 par mois. Un clip 10 s Gen-4 + upscale 4K ≈ 140 crédits. **C'est le plus cher du panorama à qualité égale** — le prix de la plateforme et de la régularité (le moins de re-rolls nécessaires selon les comparatifs).
- **API développeur** : existe (modèles Runway + modèles tiers comme Veo 3.1 au catalogue, facturation au crédit). Attention : changer de modèle dans l'API ne rend pas les paramètres interchangeables — chaque modèle a ses formats/durées propres.
- Plan gratuit : 125 crédits **one-shot** (non renouvelés), et Gen-4 exclu du gratuit (on teste avec Gen-3 Alpha Turbo : trompeur pour évaluer le produit actuel).

## 48. Luma Ray 3.2 (Dream Machine) : l'atmosphérique

- Points forts : **effets atmosphériques** (fumée, feu, eau, brouillard) — le meilleur du marché sur ce créneau selon les comparatifs 2026 ; mouvement naturel ; **Brainstorm mode** ; start/end frames historiques.
- Points faibles : **pas d'audio natif** (à ajouter à part), plus de variance d'un prompt à l'autre (davantage de re-rolls), cohérence de personnage en retrait de Kling.
- Tarifs 2026 (le produit a été restructuré : gamme **Luma Agents** + legacy Dream Machine) : **Plus 30 $/mois** (10 000 crédits), **Pro 90 $/mois** (40 000), **Ultra 300 $/mois** (150 000) ; legacy Dream Machine : Lite 9,99 $, Plus 29,99 $, Unlimited 94,99 $. **Usage commercial à partir de Plus** ; le gratuit/pas cher = watermark + usage personnel uniquement.
- Repère de coût : ~**170 crédits pour 5 s**, ~800 crédits pour 10 s (variable selon résolution/modèle). **Les crédits mensuels n'expirent pas en report** : non consommés = perdus (seuls les packs Top-Up, dès 4 $/1 200 crédits, sont valables 12 mois).
- Luma Agents regroupe sous un même compteur **Veo 3.1, Kling 3.0, Seedance et ElevenLabs** : intéressant si tu veux jongler entre moteurs sans multiplier les abonnements.

## 49. Hailuo / MiniMax H3 : le brouillon rapide et pas cher

- **MiniMax** (éditeur chinois, produit **Hailuo**) : le positionnement « **rapide et économique** ».
- Version 2026 : **H3 (Hailuo 3.0)** — tarif API officiel **forfaitaire : 0,26 $/seconde** (5 s = 1,30 $, 15 s = 3,90 $), résolution 2K, durées 5-15 s ; mode ref2v : +0,08 $/s par image de référence au-delà de 5. Génération en 5-15 s chrono : le plus rapide du marché.
- Abonnement : **Standard 14,99 $/mois** (des offres à 9,99 $/mois ont existé — **à vérifier**).
- Usage : **itération rapide**, storyboards animés, brouillons avant de masteriser sur Kling/Veo. La communauté l'utilise en « draft » et Kling/Runway en « final ».
- **Licence** : nos recherches du 27/09/2026 (volume Chine) signalent pour MiniMax H3 une **licence territoriale restrictive** — **à vérifier** dans le texte de licence avant tout usage commercial hors périmètre autorisé. Versions précédentes (2.3) : 512p-1080p, 6-10 s, ~0,045 $/s via agrégateurs.

## 50. Seedance (ByteDance) : le précis

- **Seedance** est le modèle vidéo de **ByteDance** (maison-mère de TikTok). Version marquante : **2.0 (février 2026)**, puis **2.5**.
- Points forts : **précision du mouvement** (danse, action, expressions — l'ADN TikTok), **entrée multimodale** (texte + images + vidéo + audio combinés, jusqu'à 12 fichiers), **clips longs : jusqu'à 30 s** en un appel (contre 15 s chez Kling), audio natif.
- Tarifs API (2.5, sept 2026) : **0,134 $/s à 480p**, **~0,29 $/s à 720p** (10 s 720p ≈ 2,90 $ ; 30 s ≈ 8,70 $). Moins « beau » que Kling/Veo à réglages égaux, mais **plus obéissant** sur le mouvement demandé.
- Accès : API BytePlus/Seedance + agrégateurs (fal.ai : ~0,30 $/s en 720p). Écosystème API plus jeune que Kling/Runway/Veo.

## 51. Les poids ouverts vidéo : Wan, LTX

- **Wan 2.x (Alibaba)** : famille **open weights** (2.5/2.6/2.7), 1080p, 10-15 s, avec audio sur certaines versions. Le choix « FLUX de la vidéo » pour du local/confidentiel. Via agrégateurs : dès ~0,03-0,07 $/s.
- **LTX 2.3 (Lightricks)** : ouvert, rapide, 20 s max, pensé pour le temps réel et l'itération.
- **MiniMax H3** : évoqué comme partiellement ouvert par certaines sources (**licence restrictive — à vérifier**, voir section 49).
- En local, la vidéo générative reste **très gourmande** (24 Go VRAM minimum pour du court, temps de rendu en minutes) : le cloud reste la norme, le local l'exception motivée (confidentialité, volume).

## 52. Tableau comparatif vidéo (sept 2026)

