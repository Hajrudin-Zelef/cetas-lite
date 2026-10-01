---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-14
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Google", "OpenAI", "Perplexity"]
dates: ["2026-09-27"]
keywords: ["apache", "diffusion", "distillation", "fine-tuning", "fp8", "gemini", "lora", "open source", "perplexity", "qwen", "research", "safetensors"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [1317, 1466]
sha256: 67f1e9cdca899d446b4ca0bef0112e5aa0f296749b000c37c6ba813e8e8f59a4
---

# IA générative : image, vidéo, recherche

- [ ] Textes dans l'image **relus caractère par caractère** (tensions, étiquettes, logos)
- [ ] Cohérence vérifiée (couleurs, équipement, personnage identiques dans la série)
- [ ] Pas de watermark / pas d'export « gratuit »
- [ ] Mention « généré par IA » si diffusion externe ou doc sensible
- [ ] Métadonnées de provenance conservées (C2PA si présentes)
- [ ] Prompts + seeds + modèle **archivés** avec le livrable (reproductibilité)
- [ ] Pour la vidéo : visionnée **en entier**, audio écouté (pas de mot fantôme / bruit bizarre)

## 108. Check-list : mise en prod d'une API IA (pour ton RAG / tes scripts)

- [ ] Clé dans variable d'environnement / gestionnaire de secrets, **jamais** dans le code
- [ ] Une clé par usage (dev / prod / batch), alertes de dépense activées
- [ ] Version du modèle **pinnée** (pas d'alias flottant)
- [ ] Timeouts + retries bornés + logs (modèle, version, coût par appel)
- [ ] Webhooks plutôt que polling quand supportés
- [ ] Plan B fournisseur identifié (leçon Sora, section 87)
- [ ] Page de dépréciation du fournisseur surveillée (trimestriel)
- [ ] Données envoyées : classification OK (section 78), rien d'interdit par NDA/client

## 109. Kit de survie : commandes et liens utiles

```bash
# Suivi des dépenses (exemples, adapter)
# BFL : console https://dashboard.bfl.ai  (crédits, historique)
# Kling : https://kling.ai/dev            (packs, consommation API)
# Perplexity : console API                (usage tokens + requêtes)
# fal.ai / Replicate : dashboard           (facturation à l'usage, par modèle)

# Hygiène des clés
export BFL_API_KEY="..." PERPLEXITY_API_KEY="..." KLING_API_KEY="..."
# Ne JAMAIS : git add d'un .env ; copier une clé dans un chat ; la laisser dans l'historique shell partagé
```

## 110. Anti-sèche tarifs (au 27/09/2026 — re-vérifier trimestriellement)

```
IMAGE / génération 1024² : FLUX schnell ~0,003 $ | FLUX.2 pro 0,03 $ | Imagen 4 Std 0,04 $
  | DALL-E 3 0,04 $ | Nano Banana Pro 0,134 $ | GPT Image high 0,167 $
VIDÉO / 10 s 1080p : Kling ~1,1-1,7 $ (API) | Hailuo H3 2,60 $ | Seedance 2,5 ~2,9 $
  | Veo Fast 1,2 $ | Veo Standard 4 $ | Runway ~5-6 $
RECHERCHE : Perplexity Pro 20 $/mois | Sonar API ~1 $/1M tok | Sonar Pro ~3/15 $/1M + 5-14 $/1K req
DÉCISION : Jev 0,042 $/1M tokens entrée, sortie gratuite
```

---

# PARTIE M — GLOSSAIRE

## 111. Air-gap

Machine ou réseau **physiquement isolé** d'Internet. Seuls les modèles en local (FLUX poids ouverts) permettent un vrai workflow air-gap pour des données sensibles.

## 112. Apache 2.0

Licence open source **permissive** : usage commercial autorisé, y compris modification et redistribution, avec obligation de conserver la notice de licence. Dans ce guide : FLUX.1 [schnell], FLUX.2 [klein] 4B, VAE FLUX.2, Qwen-Image 2.0.

## 113. Batch API

Mode de traitement **différé** (file d'attente, rendu en heures) avec **~50 % de remise** (Google : -50 % sur Imagen/Nano Banana ; OpenAI : -50 % sur plusieurs API). Idéal pour le volume planifiable, inutilisable pour l'interactif.

## 114. C2PA

Standard de **provenance des contenus** (métadonnées signées : « cette image a été générée par tel outil à telle date »). De plus en plus intégré par les plateformes. À conserver : c'est ta traçabilité.

## 115. Checkpoint / poids (weights)

Fichiers contenant le modèle entraîné (ex. `flux2_dev_fp8mixed.safetensors`, ~35 Go). « Poids ouverts » = téléchargeables ; la **licence** dit ce que tu as le droit d'en faire.

## 116. CFG / guidance scale

Paramètre qui règle à quel point le modèle **suit le prompt** (valeurs typiques 3-7). Trop haut : artefacts, image « brûlée ». Trop bas : le modèle ignore ton prompt.

## 117. ComfyUI

Orchestrateur **open source** de workflows de génération d'image/vidéo par **nœuds** (interface web locale). Standard de fait pour FLUX en local : reproductible, scriptable via API, immense écosystème d'extensions.

## 118. Crédit (plateforme)

Unité de facturation interne. **Piège** : 1 crédit Kling ≠ 1 crédit Luma ≠ 1 crédit Runway. Toujours convertir en **USD par livrable**.

## 119. Deep Research

Mode de recherche IA **approfondie** : le système enchaîne des dizaines de recherches et de lectures avant de synthétiser (Perplexity, Gemini...). Puissant pour cartographier un sujet, lent (minutes), à vérifier comme le reste.

## 120. Diffusion (modèle de)

Famille de modèles génératifs qui partent du **bruit** et le « débruitent » pas à pas vers une image/vidéo (principe de FLUX, Stable Diffusion, Kling, Veo...). D'où les paramètres « steps » et « sampler ».

## 121. Distillation

Technique qui compresse un gros modèle en un **petit modèle rapide** qui l'imite (FLUX [schnell], klein 4B). Moins fidèle, beaucoup plus rapide et léger.

## 122. Fine-tuning / LoRA

Ré-entraînement (complet ou léger via **LoRA**, quelques Mo) d'un modèle sur **tes** images pour lui apprendre un style, un équipement, un personnage. Rentable au volume (catalogue de 200 visuels cohérents), overkill pour 10 images.

## 123. FP8 / quantification

Réduction de la **précision des poids** (32 bits → 8 bits) pour diviser la VRAM et accélérer, avec une perte de qualité faible. Le réglage standard 2026 pour FLUX.2 en local (24 Go au lieu de 64).

## 124. Guidance-distilled

Modèle distillé pour fonctionner avec peu de steps **sans** réglage fin de guidance (cas de FLUX.1 [dev]). En pratique : moins de boutons à tourner, qualité stable.

## 125. Inpainting / outpainting

**Inpainting** : régénérer une zone masquée d'une image (effacer un défaut, changer un détail). **Outpainting** : étendre l'image au-delà de ses bords (passer du carré au 16:9).

## 126. Kontexte (FLUX.1 Kontext)

Modèle d'**édition d'image par instruction** (« remplace X par Y ») **sans masque**, avec conservation du reste de l'image et cohérence de personnage. Le plus abouti de sa catégorie en 2026.

## 127. Lip-sync

Synchronisation **lèvres/parole** dans la vidéo générée (Kling 3.0, Veo 3.1, Seedance : audio natif avec dialogues). À tester dans la langue cible : la qualité varie selon les langues.

## 128. LoRA

Voir 122. Fichier léger (~Mo) qui **spécialise** un modèle de base (un équipement, un style). S'échange et se combine facilement dans l'écosystème ouvert.

## 129. Motion transfer

Extraire le **motif de mouvement** d'une vidéo de référence pour l'appliquer à un autre sujet (différenciateur Kling 3.0). Utile quand tu as la « bonne » gestuelle sous les yeux mais pas le bon sujet.

## 130. Multi-prompt

Découper une génération vidéo en **segments temporels**, chacun avec son prompt (Kling 3.0 : jusqu'à 15 s, 512 caractères/segment). La façon de raconter une mini-histoire en un plan continu.

## 131. Nano Banana

Surnom communautaire des modèles **image de Google** (Gemini 2.5 Flash Image = Nano Banana, 3.1 = Nano Banana 2, 3 Pro = Nano Banana Pro). Photorealisme et édition de haut niveau via l'API Gemini.

## 132. Negative prompt

Ce qu'on **ne veut pas** (« flou, déformé, filigrane, texte illisible, six doigts »). Très utile sur FLUX/Midjourney pour éliminer les défauts récurrents.

## 133. Polling / webhook

Deux façons de suivre une génération **asynchrone** : **polling** = interroger régulièrement « c'est prêt ? » ; **webhook** = le fournisseur t'appelle quand c'est prêt. Le webhook est la bonne pratique en prod.

## 134. Prompt

La **consigne** en langage naturel. En 2026, les bons prompts sont **descriptifs et structurés** (sujet + action + caméra + lumière + style), pas des listes de mots-clés — sauf en recherche où les opérateurs classiques restent rois.

## 135. Quantification

Voir 123 (FP8).

## 136. RAG (rappel)

