---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-1
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Microsoft", "MiniMax", "Moonshot", "OpenAI", "Perplexity", "Z.ai"]
dates: ["2026-03-24", "2026-04-26", "2026-09-24", "2026-09-27"]
keywords: ["claude", "copilot", "deepseek", "gemini", "glm", "kimi", "llama", "omni", "open source", "open weights", "perplexity", "qwen"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [1, 95]
sha256: c1d7d9a18aeb363551143f510493e535e81c3613c764537c793fa31108f724e4
---

# IA générative : image, vidéo, recherche

> Guide terrain pour Zelef — chef de service systèmes & énergies, sysadmin, constructeur de son propre RAG.
> **Périmètre :** outils de génération d'image, de vidéo et de recherche IA, avec prix, API, licences et pièges.
> **Ce guide ne duplique pas** les volumes existants (`ia_modeles_chine.md`, `ia_modeles_occident.md`, `ia_modeles_specialises.md`) : des renvois explicites sont faits quand un sujet y est déjà couvert.
> **Dates :** toutes les informations datées ont été vérifiées le **27/09/2026** sauf mention contraire. Les prix et versions marqués « à vérifier » n'ont pas pu être confirmés de source fiable à cette date.

---

## 1. Comment lire ce guide

Ce guide est un manuel d'atelier, pas une brochure. Chaque outil est traité avec la même grille :

1. **Ce que c'est** (éditeur, positionnement, version actuelle au 27/09/2026)
2. **Comment y accéder** (web, API officielle, API tierces, local)
3. **Combien ça coûte** (prix vérifiés le 27/09/2026 ou « à vérifier »)
4. **Licence et droits d'usage commercial** (le point que tout le monde oublie)
5. **Limites honnêtes** (ce que l'outil ne fait pas bien)
6. **Exemple concret** (prompt ou appel API copiable)

Conventions utilisées dans tout le document :

- **« vérifié le 27/09/2026 »** : information confirmée par au moins une source web consultée ce jour.
- **« à vérifier »** : information plausible mais non confirmée par une source fiable au 27/09/2026.
- **« RUMEUR non confirmée »** : bruit de couloir, aucune annonce officielle.
- Les prix sont en USD sauf mention contraire, hors taxes, et **changent vite** : toujours re-vérifier la page tarifaire officielle avant d'engager un budget.
- **Aucune clé API réelle** n'apparaît dans ce guide : les exemples utilisent `VOTRE_CLE_API` ou des variables d'environnement.

## 2. Carte du territoire

Le paysage fin septembre 2026 se résume en trois blocs :

| Bloc | Leaders (sept 2026) | Accès local possible ? |
|------|---------------------|------------------------|
| Image | Midjourney V8.2, GPT Image 2, Imagen 4 / Nano Banana Pro, FLUX.2, Ideogram 4.0, Adobe Firefly | Oui : FLUX (poids ouverts), SD 3.5, Qwen-Image, Z-Image |
| Vidéo | Kling 3.0, Veo 3.1, Runway Gen-4.5, Luma Ray 3.2, Hailuo H3, Seedance 2.5 | Partiel : Wan 2.x, LTX 2.3 (poids ouverts), MiniMax H3 (licence restrictive) |
| Recherche IA | Perplexity (+ API Sonar), Phind, You.com, Kimi, Copilot | Non (services cloud) |

Le point structurant de la période : **l'API vidéo Sora 2 d'OpenAI a été fermée le 24/09/2026** (annonce le 24/03/2026, appli fermée le 26/04/2026, sans remplaçant annoncé). Leçon à retenir pour tout le guide : ne jamais bâtir un pipeline de production sur une API sans plan B. Voir section 210.

## 3. Renvois vers les volumes IA existants

Pour éviter les doublons avec les trois volumes déjà livrés :

- **Modèles de langage chinois** (Qwen, DeepSeek, Kimi/GLM...) → voir `ia_modeles_chine.md`. Ce guide n'y revient que pour **Kimi** en tant qu'outil de recherche (section 104).
- **Modèles occidentaux** (GPT, Claude, Gemini, Llama...) → voir `ia_modeles_occident.md`. Ce guide n'y revient que pour leurs **capacités image/vidéo/recherche** (GPT Image, Veo, Gemini image).
- **Modèles spécialisés** → voir `ia_modeles_specialises.md`.
- **MiniMax H3** : licence territoriale restrictive signalée dans nos recherches du 27/09/2026 (voir volume Chine) — rappel section 96.

## 4. Le vocabulaire minimal avant de commencer

- **Text-to-image (T2I)** : image générée depuis un texte.
- **Image-to-image / image-to-video (I2I / I2V)** : on part d'une image existante qu'on transforme ou qu'on anime.
- **Text-to-video (T2V)** : vidéo générée depuis un texte.
- **Inpainting** : régénérer une zone masquée d'une image. **Outpainting** : étendre l'image au-delà de ses bords.
- **Upscale** : augmenter la résolution après génération (souvent facturé à part).
- **Poids ouverts (open weights)** : les fichiers du modèle sont téléchargeables. ≠ open source : la licence peut interdire l'usage commercial (cas de FLUX Dev).
- **Crédit** : unité de facturation interne d'une plateforme. Un crédit n'a pas de valeur universelle : toujours convertir en USD par génération.

## 5. La règle d'or du budget : coût par livrable accepté

Ne compare jamais « prix par image » ou « prix par seconde » entre outils. Compare le **coût par livrable accepté** :

```
coût_par_livrable = (dépenses_totales_génération) / (livrables_validés)
```

Les dépenses totales incluent : les essais ratés, les re-générations, l'upscale, le montage, et le temps humain de tri. Un outil à 0,03 $/image dont tu jettes 9 images sur 10 coûte plus cher qu'un outil à 0,13 $/image dont tu gardes 1 sur 2. Tiens un compteur simple (tableur ou log d'API) pendant tes deux premières semaines d'usage : c'est le seul chiffre qui compte pour décider.

---

# PARTIE A — KLING (Kuaishou) : la vidéo

## 6. Kling : ce que c'est

**Kling AI** est le générateur vidéo de **Kuaishou Technology** (géant chinois de la vidéo courte, concurrent de TikTok/Douyin). Lancé en accès restreint en juin 2024 (plus d'un million de demandes, 300 000 accès initiaux), ouvert à tous ensuite, il s'est imposé comme l'une des références mondiales de la vidéo générative. Revendications de l'éditeur : 60 millions de créateurs, 30 000 clients entreprise (vérifié le 27/09/2026 via la presse spécialisée — chiffres éditeur, à prendre avec recul).

Positionnement : **le meilleur rapport qualité/prix du marché pour la vidéo IA commerciale**, avec l'audio natif et la 4K native comme différenciateurs. Disponible dans le monde entier, interface en anglais (et chinois).

## 7. Kling : versions au 27/09/2026

La numérotation a beaucoup bougé ; état vérifié le 27/09/2026 :

| Version | Statut | Notes |
|---------|--------|-------|
| **Kling 3.0** | **Version phare actuelle** (sortie février 2026) | 4K native, audio natif, transfert de mouvement, multi-prompt |
| Kling 2.6 / 2.5 | Versions précédentes, encore proposées (mode Turbo économe) | Moins chères, qualité légèrement inférieure |
| **Kling O1** | Endpoint séparé « identité » | Travail sur références : verrouiller un personnage/produit |
| **Kling O3** | Endpoint « édition » | Reference-to-video, édition vidéo |
| Kling Image 3.0 / Omni | Génération d'image (moteur « Kolors ») | 1 crédit/image en text-to-image |

Règle pratique : pour un nouveau projet, pars sur **Kling 3.0** sauf contrainte budgétaire forte (auquel cas 2.5/2.6 Turbo en brouillon, puis 3.0 pour le master).

## 8. Kling 3.0 : fonctionnalités clés

