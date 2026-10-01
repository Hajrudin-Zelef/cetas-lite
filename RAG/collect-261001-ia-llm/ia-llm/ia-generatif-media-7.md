---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-7
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Google", "OpenAI"]
dates: []
keywords: ["apache", "benchmark", "chatgpt", "diffusion", "gemini", "gpu", "leaderboard", "license", "qwen"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [600, 704]
sha256: c14427bead863b6aa851da99bcf014ed41efad916f6df379af3c7233d36bc889
---

# IA générative : image, vidéo, recherche

| Modèle | Prix indicatif / image (API) | Notes |
|--------|------------------------------|-------|
| **gpt-image-2** | fourchette ~0,006-0,211 $ selon taille/qualité (sources tierces, **à vérifier** sur la page officielle) | Le plus récent, édition native, tailles libres |
| gpt-image-1.5 | ~0,009-0,200 $ | Version allégée |
| gpt-image-1-mini | ~0,005-0,052 $ | Économique |
| gpt-image-1 | 0,011 $ (low) / 0,042 $ (med) / 0,167 $ (high) en 1024² | Référence 2025 |
| **DALL-E 3** | 0,040 $ (standard) / 0,080 $ (HD) | Stable, prévisible |
| DALL-E 2 | 0,016-0,020 $ | Legacy |

Points forts : **fidélité aux instructions** (le modèle comprend des consignes complexes et littérales), **édition d'image native** (inpainting au masque, fond transparent), intégration à l'écosystème OpenAI/ChatGPT (abonnés Plus/Pro : génération incluse dans l'abonnement chat). Points faibles : moins « beau » par défaut que Midjourney, moins photoréaliste que les meilleurs Google/FLUX.2, pas d'usage local.

Exemple d'appel API (endpoint images) :

```python
import os
from openai import OpenAI
client = OpenAI(api_key=os.environ["OPENAI_API_KEY"])

img = client.images.generate(
    model="gpt-image-1",          # ou "dall-e-3", "gpt-image-2" (tarifs à vérifier)
    prompt="pictogramme vectoriel minimaliste d'un onduleur triphasé, "
           "style ligne bleue sur fond blanc, sans texte",
    size="1024x1024",
    quality="medium",
)
print(img.data[0].b64_json[:60], "...")   # base64 : décoder puis sauvegarder
```

## 39. Google : Imagen 4 et Nano Banana

Google propose **deux gammes** via la Gemini API et Vertex AI (facturation au token / à l'image, **pas de palier gratuit pour l'image en API** — l'UI gratuite d'AI Studio existe avec quotas) :

| Modèle | Prix / image (API) | Positionnement |
|--------|-------------------|----------------|
| **Imagen 4 Fast** | **0,02 $** | Volume, brouillons, A/B tests |
| **Imagen 4 Standard** | **0,04 $** | Équilibre qualité/prix (80-90 % du premium) |
| **Imagen 4 Ultra** | **0,06 $** | Photorealisme max, 2K natif, zéro artefact |
| **Nano Banana** (gemini-2.5-flash-image) | **0,039 $** | Édition rapide originale |
| **Nano Banana 2** (gemini-3.1-flash-image) | 0,045-0,067 $ selon résolution | Haut du leaderboard, rapide |
| **Nano Banana Pro** (gemini-3-pro-image) | **0,134 $** (1-2K), **0,24 $** (4K) | Photorealisme extrême, contrôles pro (lumière, focale, colorimétrie) |

- **Batch API : -50 %** sur tous ces tarifs (0,01 $ l'Imagen 4 Fast en batch : imbattable pour du volume planifiable).
- Abonnements Google AI : AI Pro (~19,99 $/mois ≈ 100 images NB Pro/jour), AI Ultra (~30 $/mois ≈ 1 000/jour) — ordres de grandeur, **à vérifier**.
- Points forts : **photoréalisme** (Imagen 4 Ultra préféré à 78 % pour les portraits lors de tests indépendants juin 2026), **rendu de texte**, intégration à l'écosystème Google (même clé/facturation que Gemini texte).
- Points faibles : peaux parfois « plastiques », raisonnement spatial perfectible, **pas de cohérence de personnage native** au niveau de FLUX Kontext, pas d'usage local.

## 40. Ideogram 4.0 : le spécialiste du texte dans l'image

**Ideogram** (startup, modèle 4.0 en 2026) s'est fait un nom sur un point précis : **le texte long et lisible intégré à l'image** (affiches, visuels marketing, miniatures). Tarifs API : **~0,03-0,10 $/image** selon qualité, ou abonnement. À évaluer quand ton besoin n°1 est « une affiche avec un vrai titre lisible » — là où Midjourney et FLUX.1 peinent encore sur les longs passages.

## 41. Adobe Firefly (Image Model 5) : l'option « juridiquement sûre »

- Abonnement **9,99 $ à 199,99 $/mois** (crédits ; les modèles partenaires/premium consomment plus).
- Argument unique : modèle entraîné sur du contenu **licencié** (Adobe Stock et domaine public), intégré à **Photoshop, Illustrator, Premiere**. Pour une entreprise qui veut minimiser le risque de contentieux sur les données d'entraînement, c'est l'option la plus défendable.
- Qualité correcte mais en retrait des meilleurs sur le photoréalisme pur ; l'intérêt est l'intégration au workflow Adobe existant, pas le benchmark.

## 42. Les poids ouverts image (rappel)

Pour mémoire (détaillés pour FLUX sections 16-27) :

| Modèle | Licence | Intérêt |
|--------|---------|---------|
| FLUX.1 [schnell] | Apache 2.0 | Rapide, commercial OK |
| FLUX.2 [klein] 4B | Apache 2.0 | Qualité 2025+, commercial OK, 12 Go VRAM |
| Stable Diffusion 3.5 (Large/Medium) | Stability Community License (gratuit < 1 M$ revenus) | Écosystème immense, 10 Go VRAM (Medium) |
| Qwen-Image 2.0 (Alibaba, 20B) | Apache 2.0 | Typographie pro, édition |
| Z-Image (Alibaba) | Ouverte (**licence à vérifier**) | Alternative légère |

## 43. Tableau comparatif image (sept 2026)

| Critère | Midjourney V8.2 | GPT Image 2 | Imagen 4 Ultra | Nano Banana Pro | FLUX.2 Pro (API) | FLUX local (klein 4B) | Ideogram 4.0 | Firefly |
|---------|----------------|-------------|----------------|-----------------|------------------|----------------------|--------------|---------|
| Qualité esthétique | ★★★★★ | ★★★★ | ★★★★ | ★★★★ | ★★★★ | ★★★ | ★★★ | ★★★ |
| Photoréalisme | ★★★★ | ★★★ | ★★★★★ | ★★★★★ | ★★★★ | ★★★ | ★★★ | ★★★ |
| Fidélité au prompt | ★★★★ | ★★★★★ | ★★★★ | ★★★★ | ★★★★ | ★★★★ | ★★★★ | ★★★ |
| Texte dans l'image | ★★★ | ★★★★ | ★★★★★ | ★★★★★ | ★★★★ | ★★★ | ★★★★★ | ★★★ |
| Cohérence personnage | ★★★★ | ★★★ | ★★ | ★★★ | ★★★★★ (multi-ref) | ★★★★ | ★★★ | ★★★ |
| Prix / image (ordre) | abo (~0,03-0,05 $ eff.) | 0,01-0,21 $ | 0,06 $ | 0,134 $ | **0,03 $** | ~0 $ (GPU) | 0,03-0,10 $ | abo |
| API | Non (officiel) | Oui | Oui | Oui | Oui | Locale (ComfyUI) | Oui | Oui (Adobe) |
| Usage local | Non | Non | Non | Non | Non | **Oui** | Non | Non |
| Usage commercial | Oui (Pro/Mega si >1 M$) | Oui (CGU API) | Oui (API payante) | Oui (API payante) | Oui (CGU API) | **Oui (Apache 2.0)** | Oui (payant) | Oui (licencié) |

Lecture : pas de gagnant absolu. **Stack recommandée pour un usage pro mixte** : FLUX (local ou API BFL) en base quotidienne + Nano Banana Pro ou Imagen 4 Ultra pour le photoréalisme critique + Ideogram pour les affiches à texte + Midjourney si l'esthétique « belle image » prime.

## 44. Image : prompts efficaces — la méthode commune

Tous les modèles image 2026 partagent la même grammaire de prompt :

```
[sujet précis] + [contexte/environnement] + [lumière] + [point de vue/cadrage]
+ [style/rendu] + [détails techniques] + [ce qu'on NE veut pas]
```

Exemple « doc technique » (à adapter par modèle) :

> « photographie technique d'un onduleur triphasé 40 kVA en armoire 19 pouces, voyants verts en façade, câbles d'arrivée par le bas, local technique propre, lumière neutre 5000K, cadrage frontal légèrement en contre-plongée, netteté maximale, style documentation constructeur, -- pas de personnes, pas de texte flou, pas de filigrane »

Notes par modèle :
- **Midjourney V8** : le langage naturel narratif marche mieux que la liste de mots-clés (« un technicien fatigué qui... » > « technicien, fatigué, néons »). Paramètres `--style raw` pour moins d'embellissement.
- **FLUX** : prompts descriptifs denses, guillemets pour le texte à afficher (« panneau avec le texte "LOCAL TGBT" »), négatif explicite utile.
- **GPT Image** : instructions littérales et structurées très bien suivies (« à gauche..., à droite..., exactement trois... »).
- **Imagen/Nano Banana** : bon avec la description d'appareil photo (« 35mm, f/2.8, lumière dorée »).

---

