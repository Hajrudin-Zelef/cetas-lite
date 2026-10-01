---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-15
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["Google", "OpenAI", "Perplexity"]
dates: ["2026-03-24", "2026-09-15", "2026-09-24", "2026-09-27"]
keywords: ["apache", "diffusion", "fp8", "gemini", "gpu", "perplexity", "pricing", "text-to-video"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [1467, 1571]
sha256: 4a8c47e68abe90523f0aa73fa72778a4a416456737d1a851c1e9bd41c4ceedbf
---

# IA générative : image, vidéo, recherche

**Retrieval-Augmented Generation** : ton pipeline — une base documentaire interrogée par un LLM qui cite ses sources. Ce guide t'aide à choisir les **outils de recherche** (Perplexity/Sonar) et de **génération de visuels** qui gravitent autour de ton RAG, pas à le reconstruire (voir tes volumes existants).

## 137. Rectified flow

Technique d'entraînement des modèles FLUX (v2) : trajectoire de « débruitage » plus directe que la diffusion classique → **moins de steps, meilleure fidélité**. Détail technique, mais c'est ce qui rend FLUX rapide à qualité égale.

## 138. Relax mode

File d'attente **lente et illimitée** (Midjourney Standard/Pro/Mega, Luma Unlimited) : tu ne consommes pas ton quota, tu attends. Idéal pour le volume non urgent.

## 139. Seed

**Graine aléatoire** : à prompt et modèle identiques, même seed = même image. La base de la **reproductibilité** (noter le seed de chaque visuel validé).

## 140. Sonar

Famille de modèles **API de Perplexity** : recherche web temps réel + génération avec citations, facturée au token + à la requête. Le pont entre « je cherche à la main » et « ma veille tourne toute seule ».

## 141. Stealth mode

Générations **privées** (non visibles par la communauté) — Midjourney Pro/Mega uniquement. Sans ça, tes prompts et images sont publics par défaut. Pour du pro : indispensable ou rédhibitoire.

## 142. Steps (pas de diffusion)

Nombre d'itérations de débruitage. Plus = mieux (jusqu'à un plateau) mais plus lent. Repères : FLUX.1 schnell 1-4, FLUX.1/2 dev 20-50.

## 143. System One / System Two

Référence à Kahneman reprise par **TypeSafe** : System 2 = le LLM qui raisonne en parlant (lent) ; **System One = Jev**, qui juge vite par probabilités. Deux outils complémentaires, pas concurrents.

## 144. Text-to-video / image-to-video

**T2V** : vidéo depuis un texte. **I2V** : animation d'une image existante — souvent **meilleur pour le pro** (le réel ancre la crédibilité, l'IA n'ajoute que le mouvement).

## 145. Token

Unité de facturation des API de langage/recherche (~0,75 mot en anglais). Les API Sonar facturent **tokens + requêtes** : surveiller les deux compteurs.

## 146. Upscale

Augmentation de la **résolution après génération** (ex. 1024 → 4096). Souvent facturée à part (Runway, BFL au mégapixel). À n'appliquer que sur les visuels **validés**.

## 147. VAE

**Auto-encodeur variationnel** : le module qui compresse/décompresse l'image dans l'espace latent du modèle. Celui de FLUX.2 est **Apache 2.0** : tu peux l'utiliser librement dans tes pipelines.

## 148. VRAM

Mémoire de la carte **GPU**. Le facteur limitant n°1 du local : 12 Go (klein 4B), 16 Go (FLUX.1 dev FP8), 24 Go (FLUX.2 dev FP8).

## 149. Watermark (filigrane)

Marque incrustée par les **offres gratuites** (Kling, Luma...). Bannie des livrables : passer au payant ou ne pas utiliser.

## 150. Weights

Voir 115 (checkpoint).

---

# PARTIE N — QUIZ : 10 questions pour valider

## 151. Questions

**Q1.** Quelle est la version phare de Kling au 27/09/2026, et quelles sont ses deux différenciateurs majeurs face à la concurrence ?
**Q2.** Tu génères des visuels en local avec FLUX.2 [dev] pour une documentation client facturée. Problème ?
**Q3.** Quelle est la différence de facturation entre l'abonnement web Kling et l'API Kling ?
**Q4.** Cite trois choses qui expirent / ne se reportent pas chez les fournisseurs de ce guide.
**Q5.** Pourquoi faut-il « pinner » la version d'un modèle dans un pipeline de production ? (Deux raisons.)
**Q6.** L'API Sora 2 d'OpenAI : que s'est-il passé le 24/09/2026, et quelle leçon d'architecture en tires-tu ?
**Q7.** Dans quel cas l'audio natif de Kling 3.0 est-il un mauvais calcul économique ?
**Q8.** Perplexity te répond avec des citations sur une valeur de réglage (ex. tension de floating). Que fais-tu avant de l'utiliser dans une note de calcul ?
**Q9.** Qu'est-ce que Jev (TypeSafe), et pourquoi ne peut-il pas générer tes visuels ?
**Q10.** Calcule le coût API de 200 images 1024² en FLUX.2 [pro] via l'API BFL, et compare avec 200 images Nano Banana Pro via l'API Gemini.

## 152. Réponses

**R1.** **Kling 3.0** (sortie février 2026). Différenciateurs : le **motion transfer** (appliquer le mouvement d'une vidéo de référence à un autre sujet) et le **rendu de texte à l'écran** le plus lisible du marché, plus la 4K native et l'audio natif avec lip-sync multilingue.
**R2.** **Oui, problème** : FLUX.2 [dev] est en **licence non-commerciale**. Pour du commercial en local : FLUX.2 [klein] 4B ou FLUX.1 [schnell] (Apache 2.0), ou licence commerciale BFL.
**R3.** Ce sont **deux facturations totalement séparées** : l'abonnement web (crédits mensuels, interface) ne crédite pas l'API, et les packs API prépayés (30/180 jours de validité) ne servent pas sur le web.
**R4.** Au choix : crédits mensuels Kling web (fin de période), unités des packs API Kling (30/180 j), crédits Luma mensuels, heures GPU Midjourney, crédits mensuels non consommés en général.
**R5.** (1) Éviter les **changements silencieux** de comportement quand l'alias ou le snapshot bouge ; (2) garantir la **reproductibilité** et le débogage (savoir exactement quelle version a produit quoi).
**R6.** **Fermeture définitive** de la Videos API (sora-2, sora-2-pro, snapshots), annoncée le 24/03/2026, **sans remplaçant**. Leçon : isoler chaque appel IA derrière une **interface interne** (fonction à toi), exporter les assets, suivre les dépréciations, toujours un fournisseur B.
**R7.** Quand le son n'apporte rien au livrable (ex. clip muet d'intranet, ambiance ajoutée de toute façon en montage) : l'audio natif peut **multiplier le coût par 2,5 à 5** pour zéro valeur ajoutée.
**R8.** **Ouvrir les sources citées** et vérifier la valeur dans la **source primaire** (doc constructeur, norme). Perplexity est un pointeur, pas une preuve — surtout pour une valeur d'ingénierie.
**R9.** Jev est un modèle de **décision** (TypeSafe AI, « System One ») : il répond à des questions fermées (oui/non, choix, score) par des **probabilités**, en ~100-500 ms, pour ~0,042 $/1M tokens. Il **ne génère rien** (ni texte, ni image, ni vidéo) : il juge, il ne crée pas. Usage : contrôle qualité, routage, garde-fous.
**R10.** FLUX.2 [pro] : 200 × 0,03 $ = **6 $**. Nano Banana Pro : 200 × 0,134 $ = **26,80 $**. Soit **~4,5× moins cher** avec FLUX.2 pro à qualité souvent comparable — d'où son intérêt pour le volume documentaire.

---

# PARTIE O — POUR ALLER PLUS LOIN

## 153. Documentations officielles (faire foi en cas de conflit avec ce guide)

- API Black Forest Labs : `https://docs.bfl.ai` — tarifs : `https://bfl.ai/pricing` (**à vérifier** : URLs de documentation, je n'ai pas ouvert chaque page le 27/09/2026 ; la page pricing est citée par la presse spécialisée).
- Kling développeurs : `https://kling.ai/dev` (packs, modèles, polling/webhooks).
- Perplexity API : `https://docs.perplexity.ai` — tarifs : `https://docs.perplexity.ai/docs/getting-started/pricing`.
- OpenAI (images) : `https://platform.openai.com/docs` (grille tarifaire officielle).
- Google (Imagen / Gemini image / Veo) : `https://ai.google.dev/gemini-api/docs/pricing` et console Vertex AI.
- Midjourney : `https://docs.midjourney.com` (modèles, paramètres, CGU).
- TypeSafe Jev : `https://docs.typesafe.ai` (**à vérifier** — produit sorti de stealth le 15/09/2026, la doc évolue vite).
- ComfyUI : `https://docs.comfy.org` ; templates FLUX.2 dans l'appli.

## 154. Pour creuser côté sysadmin

