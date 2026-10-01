---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-3
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["Mistral", "Stripe", "United States"]
dates: ["2026-09-27"]
keywords: ["apache", "attention", "diffusion", "fine-tuning", "mistral"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [203, 300]
sha256: 9edbf5f2c188816ee6df10d0158b206cf3c23c989b4cf31340b77b0c7d1c7195
---

# IA générative : image, vidéo, recherche

Règles :
- **Une action principale par clip** (5-10 s). Deux actions = deux clips.
- Décrire le **mouvement de caméra** explicitement (travelling, panoramique, zoom, plan fixe, drone).
- L'audio natif : décrire le son souhaité (« Audio : ronronnement grave de ventilation, clics métalliques ») — Kling 3.0 le génère avec l'image.
- Le **multi-prompt** pour les mini-scénarios : un prompt par segment temporel, avec continuité du sujet (« même technicien, même armoire »).
- Ce que Kling fait mal : texte long à l'écran (au-delà de quelques mots), gestes ultra-précis (mains), chorégraphies exactes.

## 14. Kling : cas d'usage pro (angle Zelef)

1. **Vidéo de présentation d'équipement** (onduleur, baie, groupe électrogène) : image-to-video depuis une photo réelle du matériel + prompt de caméra lente + audio natif (ronronnement). Usage : formation interne, page intranet, réponse à appel d'offres.
2. **Clip de sensibilisation sécurité** : scénario multi-prompt (ex. consignation électrique en 3 plans de 5 s). Bien moins cher qu'un tournage.
3. **Visite virtuelle de site** : enchaîner 4-6 clips image-to-video depuis des photos du local technique → montage simple → visite guidée pour un client qui ne se déplace pas.
4. **Contenu formation** : animer des schémas (plan unifilaire, synoptique) pour expliquer un cheminement électrique. Attention : le texte dans l'image source peut se déformer à l'animation — tester d'abord.

## 15. Kling : limites honnêtes

- **Files d'attente** sur le web aux heures de pointe (jusqu'à 30+ min constatées). L'API est plus prévisible.
- **Générations ratées facturées** : un clip raté consomme les crédits quand même, sur le web comme (selon l'agrégateur) via API. Budgéter 2 à 3 essais par clip final.
- **Liberté créative** : Kling interprète plus qu'il n'exécute au pied de la lettre. Pour un mouvement imposé, Seedance ou le motion transfer sont meilleurs.
- **Personnes** : visages et mains perfectibles, surtout en 720p. Pour un visage récurrent, passer par l'endpoint O1 (références).
- **Données** : tout ce que tu uploades (photos d'installations, plans) transite par les serveurs de Kuaishou. Pas d'option on-prem. Pour du confidentiel : flouter/anonymiser avant envoi, ou ne pas utiliser.
- **Pas de contrôle fin au photogramme** : ce n'est pas un logiciel de montage. Le montage final se fait ailleurs (DaVinci, Premiere...).

---

# PARTIE B — FLUX (Black Forest Labs) : l'image

## 16. FLUX : ce que c'est

**FLUX** est la famille de modèles d'image de **Black Forest Labs (BFL)**, startup allemande fondée par d'anciens de l'équipe Stable Diffusion. Positionnement : **le meilleur compromis du marché entre qualité, coût et déploiement local**. Particularité stratégique : BFL publie des **poids ouverts** pour une partie de la gamme, ce qui en fait le seul acteur « flagship » qu'on peut faire tourner **on-premise / air-gap** — argument décisif pour du travail confidentiel (plans, installations client).

Architecture (rappel technique) : transformeur de diffusion à flux rectifié (rectified flow), ~12 Md de paramètres pour FLUX.1, **32 Md** pour FLUX.2 Dev (associé à un modèle vision-langage Mistral-3), double encodeur texte T5 + CLIP sur la génération 1.

## 17. FLUX : la famille au 27/09/2026

### FLUX.2 (génération actuelle, sortie novembre 2025)

| Variante | Positionnement | Licence / accès |
|----------|---------------|-----------------|
| **FLUX.2 [pro]** | Meilleure qualité, hébergé | API BFL + partenaires |
| **FLUX.2 [max]** | Qualité maximale, le plus cher | API BFL |
| **FLUX.2 [flex]** | Paramètres réglables (steps, guidance) : compromis vitesse/qualité | API BFL + Playground |
| **FLUX.2 [dev]** | **Poids ouverts téléchargeables** (32 Md), qualité proche du pro | **Non-commercial** ; licence commerciale à demander à BFL |
| **FLUX.2 [klein] 4B** | Petit, rapide (~1 s sur RTX 5090), **Apache 2.0** | **Usage commercial autorisé**, déploiement on-prem |
| **FLUX.2 [klein] 9B** | Meilleur ratio qualité/latence | Poids ouverts, **non-commercial** |

Le VAE (auto-encodeur variationnel) de FLUX.2 est publié sous **Apache 2.0** : il définit l'espace latent commun, ce qui permet l'interopérabilité entre tes pipelines internes et les API BFL.

Points forts de FLUX.2 : **multi-références** (jusqu'à ~6-10 images de référence pour la cohérence d'un personnage/objet sans fine-tuning), **rendu de texte** très amélioré (infographies, UI, multilingue), contrôle de pose direct, sortie jusqu'à **4 mégapixels**.

### FLUX.1 (génération précédente, toujours pertinente)

| Variante | Positionnement | Licence |
|----------|---------------|---------|
| FLUX.1 [pro] / 1.1 | API, haute qualité | Propriétaire (API) |
| FLUX.1 [pro] Ultra | 4 MP (2048×2048+) | API |
| **FLUX.1 [dev]** | Poids ouverts 12 Md, 20-50 steps | **Non-commercial** |
| **FLUX.1 [schnell]** | 1-4 steps, ultra-rapide | **Apache 2.0** (commercial OK) |
| FLUX.1 [fill] | Inpainting/outpainting | API (+ variantes ouvertes communautaires) |
| FLUX.1 [kontext] (Max/Pro/Dev) | **Édition d'image sans masque** par instruction texte | Dev : poids ouverts 12 Md non-commerciaux |
| FLUX.1 [canny]/[depth], [redux] | ControlNet, variations | Écosystème ouvert |

**FLUX.1 Kontext** mérite une mention spéciale : c'est l'éditeur « sans masque » le plus abouti — tu décris la modification en langage naturel (« remplace le câble rouge par un câble bleu ») et le modèle édite en conservant le reste, avec une excellente **cohérence de personnage** entre scènes sans ré-entraînement.

## 18. FLUX : les licences, sans détour

C'est le point le plus mal compris de l'écosystème. Tableau de vérité :

| Variante | Poids téléchargeables ? | Usage commercial ? |
|----------|------------------------|--------------------|
| FLUX.2 [klein] **4B** | Oui | **Oui (Apache 2.0)** |
| FLUX.1 [schnell] | Oui | **Oui (Apache 2.0)** |
| FLUX.2 VAE | Oui | **Oui (Apache 2.0)** |
| FLUX.2 [dev] 32B | Oui | **Non — licence commerciale à obtenir auprès de BFL** |
| FLUX.2 [klein] 9B | Oui | **Non (non-commercial)** |
| FLUX.1 [dev] 12B | Oui | **Non (non-commercial)** |
| FLUX.1 [kontext-dev] | Oui | **Non (non-commercial)** |
| Versions [pro]/[max]/[flex] | Non (API uniquement) | Oui (via l'API, selon les CGU) |

**Piège classique** : générer avec FLUX.2 [dev] en local « parce que c'est gratuit » puis utiliser les visuels dans une doc client ou un appel d'offres → **violation de licence**. Pour un usage pro local : **klein 4B (Apache 2.0)** ou licence commerciale BFL pour dev. Pour un usage pro via API : les CGU de l'API couvrent l'usage commercial.

## 19. FLUX : l'API officielle BFL

- Endpoints : `https://api.bfl.ai` (global, bascule auto), `https://api.eu.bfl.ai` (**résidence UE — à privilégier pour des données européennes / RGPD**), `https://api.us.bfl.ai` (résidence US).
- Auth : clé API (`x-key`), compte en self-service sur `dashboard.bfl.ai`, paiement Stripe, **aucune entité juridique requise**.
- **Modèle d'appel asynchrone** : POST → `id` de tâche → GET `/v1/get_result` jusqu'au résultat ; **webhooks** supportés pour la prod.
- Facturation : **1 crédit = 0,01 $**, au **mégapixel** (entrée + sortie comptent) pour FLUX.2 ; au forfait par image pour FLUX.1.

Tarifs API BFL vérifiés le 27/09/2026 (1er mégapixel / mégapixels suivants) :

| Modèle | Endpoint | 1er MP | +MP | ~1 MP T2I |
|--------|----------|--------|-----|-----------|
| FLUX.2 [klein] 4B | `/v1/flux-2-klein-4b` | 1,4 ¢ | 0,1 ¢ | **0,014 $** |
| FLUX.2 [klein] 9B | `/v1/flux-2-klein-9b` | 1,5 ¢ | 0,2 ¢ | **0,015 $** |
| FLUX.2 [pro] | `/v1/flux-2-pro` | 3 ¢ | 1,5 ¢ | **0,030 $** |
| FLUX.2 [flex] | `/v1/flux-2-flex` | 5 ¢ | 5 ¢ | **0,050 $** |
| FLUX.2 [max] | `/v1/flux-2-max` | 7 ¢ | 3 ¢ | **0,070 $** |

