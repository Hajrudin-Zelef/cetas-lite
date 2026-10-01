---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-22
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["Perplexity", "United States"]
dates: ["2026-09-24", "2026-09-27"]
keywords: ["diffusion", "gpu", "lora", "perplexity"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [2315, 2414]
sha256: e6345eaaa2382cd12b6b1109e3c76422dc18cb0d7b2f0a72ffaf2c47eddc90fa
---

# IA générative : image, vidéo, recherche

```
┌─────────────────────────────────────────────────────────┐
│  IA GÉNÉRATIVE — FICHE RÉFLEXE (v1.0, 27/09/2026)       │
├─────────────────────────────────────────────────────────┤
│ AVANT DE GÉNÉRER                                        │
│ □ Bon outil ? (pense-bête §105)  □ Brouillon pas cher   │
│ □ Rien de confidentiel (§78)     □ Budget OK (×3 essais)│
│ □ Licence commerciale OK (§74)   □ Pas de personne réelle│
│                                                         │
│ AVANT DE PUBLIER                                        │
│ □ Textes relus caractère par caractère                  │
│ □ Cohérence de la série vérifiée                        │
│ □ Pas de watermark gratuit                              │
│ □ Mention "généré par IA" si externe                    │
│ □ Prompts + seeds archivés                              │
│                                                         │
│ EN PROD (API)                                           │
│ □ Clé en variable d'env, 1 clé/usage, alertes ON        │
│ □ Version pinnée, timeouts, logs (modèle+coût)          │
│ □ Webhook > polling  □ Plan B fournisseur identifié     │
│                                                         │
│ TARIFS REPÈRES (§110) : image 0,003-0,134 $ · vidéo     │
│ 10 s 1080p 1-4 $ · Sonar ~1 $/1M · Jev 0,042 $/1M      │
│                                                         │
│ URGENCES : clé fuitée → rotation immédiate. Facture     │
│ anormale → couper la clé, lire les logs, comprendre.    │
└─────────────────────────────────────────────────────────┘
```

## 186. Mot de la fin : l'outillage ne remplace pas le métier

Ce guide t'a donné les prix, les API, les licences et les pièges de l'IA générative image/vidéo/recherche au 27/09/2026. Mais l'outil ne fait pas l'électricien, pas plus que la perceuse ne fait le menuisier :

- **L'IA propose, tu disposes** : chaque visuel, chaque clip, chaque valeur sourcée passe par ta validation. C'est toi qui signes la doc, pas le modèle.
- **Le coût se mesure, il ne se lit pas** : deux semaines de compteur valent mieux que dix comparatifs.
- **La confidentialité ne se négocie pas** : local pour le sensible, UE pour l'interne, cloud US pour le public — et jamais l'inverse.
- **Le fournisseur peut disparaître** (Sora, 24/09/2026) : interface interne, versions pinnées, plan B.

Avec ~670 $/an (section 72), un GPU correct et les réflexes de ce guide, ton service produit des docs illustrées, des clips de formation et une veille automatique — sans dépendre d'un prestataire, sans exposer tes données, et sans mauvaise surprise sur la facture. Le reste, c'est de la pratique : génère, rate, mesure, recommence.

---

# PARTIE W — APPROFONDISSEMENTS

## 187. Kling vs Veo vs Runway : qui gagne, selon le scénario

| Scénario | Gagnant | Pourquoi |
|----------|---------|----------|
| 20 clips/mois pour intranet/formation | **Kling 3.0 (Pro web)** | 25,99 $/mois, qualité suffisante, audio natif en option |
| 1 clip vitrine pour appel d'offres | **Veo 3.1 Standard** | 3,20 $ les 8 s : le réalisme se paie, une fois |
| Série avec montage, effets, voix | **Runway** | Suite complète (Aleph, voix), même si le €/seconde pique |
| Ambiance (brume, eau, feu) | **Luma Ray 3.2** | Personne ne fait mieux l'atmosphérique |
| Geste technique précis à montrer | **Seedance 2.5** | Le plus obéissant sur le mouvement |
| Valider une idée en 2 min | **Hailuo H3** | 0,70 $ les 6 s, rendu en quelques secondes |
| Pipeline 100 % API, multi-modèles | **fal.ai / Replicate** | Un compteur, zéro pack prépayé |

## 188. Dictionnaire des paramètres (ceux qui changent vraiment le résultat)

**Image (FLUX/Midjourney)** : `steps` (20-50 dev, 1-4 schnell), `guidance/CFG` (3-7 : fidélité au prompt), `seed` (reproductibilité), `aspect_ratio` (choisir **avant**, pas après — recadrer une image générée, c'est la dénaturer), `negative_prompt` (défauts à bannir), `--style raw` (Midjourney : moins d'embellissement), `--stylize` (plus ou moins de « patte » MJ).

**Vidéo (Kling/Veo/Runway)** : `duration` (5/8/10/15 s : **une action par clip**), `resolution` (720p brouillon → 1080p master → 4K exception), `audio` (natif ou non : voir le coût), `camera` (dans le prompt : travelling, fixe, drone), `start/end frame` (I2V : ancrage), `multi_prompt` (segments temporisés).

**Recherche (Perplexity/Sonar)** : `model` (sonar éco vs sonar-pro profond), `search_recency` (filtrer par fraîcheur quand l'API le permet), `max_tokens` (borner le coût), périmètre de sources (web / académique / social).

**Décision (Jev)** : `model` pinné (`jev-1.13.0`), `seuil` de probabilité/confiance (0,70-0,80 selon le risque), `state` en anglais, questions **atomiques** (une question = un jugement).

## 189. Dix erreurs de prompting et leurs corrections

1. **Trop vague** : « un onduleur » → **précis** : « onduleur triphasé 40 kVA en armoire 19 pouces, façade avec afficheur LCD bleu, voyants verts ».
2. **Deux actions** : « le technicien ouvre l'armoire puis remplace le module » → **deux clips** (ou multi-prompt).
3. **Pas de caméra** : ajouter « plan moyen, léger travelling latéral » (vidéo) / « cadrage frontal, 35mm » (image).
4. **Pas de lumière** : « lumière neutre d'atelier 5000K » change tout.
5. **Style absent** : « style documentation constructeur » vs « photoréaliste » vs « illustration épurée » — trois résultats radicalement différents.
6. **Texte long dans l'image** : réduire à 2-3 mots entre guillemets, ou passer en PAO.
7. **Négatif oublié** : ajouter « pas de texte flou, pas de filigrane, pas de mains déformées ».
8. **Contradictions** : « de nuit » + « lumière du jour éclatante » → le modèle choisit au hasard. Relire son prompt.
9. **Unités floues** : « grand » → « 2 mètres de haut » ; les modèles comprennent mal les dimensions absolues mais mieux que le vague.
10. **On ne décrit pas l'évident** : le modèle ne sait pas ce qu'est « un TGBT normal » — décrire comme à un stagiaire de première semaine.

## 190. Roadmap d'adoption 90 jours (service systèmes & énergies)
**Semaines 1-2 — Découverte (budget : ~10 $)**
- Perplexity gratuit : 5 recherches Pro/jour sur tes sujets métier ; ouvrir systématiquement les sources.
- Kling gratuit : 3-4 clips tests (I2V depuis tes photos) ; noter le taux de rejet.
- fal.ai : 20 images FLUX schnell (~0,06 $) pour sentir le prompting.

**Semaines 3-4 — Cadrage**
- Choisir 1 besoin pilote (ex. illustrer une procédure réelle).
- Tester FLUX local (si GPU) ou API BFL ; mesurer le coût/livrable réel.
- Rédiger ta fiche réflexe (section 185) adaptée à ton équipe.

**Mois 2 — Production pilote**
- Produire la doc pilote complète (visuels + relecture + archivage prompts/seeds).
- Si vidéo : Kling Pro 1 mois (25,99 $), 3-6 clips, montage, diffusion intranet.
- Mettre en place la veille Sonar hebdo (script section 182 ou Space).

**Mois 3 — Industrialisation**
- Batch ComfyUI ou script BFL pour les séries ; LoRA si besoin d'identité verrouillée.
- Revue des coûts : coût/livrable vs estimé, ajustement des abonnements.
- Partage : Space Perplexity d'équipe, bibliothèque de prompts validés, rituel trimestriel de re-vérification des tarifs.

