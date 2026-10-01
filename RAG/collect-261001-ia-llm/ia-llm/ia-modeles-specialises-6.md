---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-6
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Google", "MiniMax", "xAI"]
dates: ["2026-09-24"]
keywords: ["apache", "arr", "benchmark", "benchmarks", "gemini", "grok", "omni", "revenue"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [577, 700]
sha256: 5f436680a91d31ab9b7dd8f75814c75e375426c800de14ac821db9e954ee67bb
---

# Encyclopédie des modèles IA — Volume 3

L'endpoint « Spicy » (version non censurée) est un signal du marché 2026 : la modération devient un paramètre produit, avec des offres différenciées. À manier avec les garde-fous habituels selon ton usage.

## 41. Kuaishou Kling 3.0

Sorti le 5 février 2026. Le champion chinois historique de la vidéo, toujours dans la course.

| Champ | Valeur |
|---|---|
| Clips | Jusqu'à 15 s ; **4K/60 fps natif revendiqué** (marketing repris par la presse — non vérifié auprès de Kuaishou) |
| Audio | Natif + lip-sync 5 langues |
| Cohérence | Sur 6 plans |
| Variantes | **Kling 3.0 Omni** (édition multi-image par référence) ; endpoint **Pro Motion Control** (mouvements caméra cinématiques complexes) |
| Prix (secondaires) | ~$0,084–0,168/s selon options ; ~$0,075–0,10/s via revendeurs ; Pro Motion Control $0,16/s ; Omni ~$0,10 (jusqu'à $0,20 avec audio) |
| Accès | API Kling, MuAPI, intégré aux plans Runway |

Kling a une particularité stratégique : il est **distribué via Runway** (section 42). Un modèle concurrent direct (Kling vs Gen-4.5) vendu dans la boutique du rival — le marché de la vidéo IA ressemble de plus en plus à un marché d'API interchangeables.

## 42. Runway Gen-4.5 / Gen-4 Aleph

La famille Gen-4 de Runway (2026, dates exactes NON VÉRIFIÉES). Runway est à la fois un labo (ses propres modèles) et une plateforme (il revend Veo 3/3.1, Kling 3.0 Pro, Seedance 2.0 via la même API).

| Champ | Valeur |
|---|---|
| Génération | Texte-vers-vidéo et image-vers-vidéo ; clips 5–10 s (chaînables, extension ×3) ; jusqu'à 1080p (4K en export sur plans Pro) ; 24–30 fps |
| **Aleph / Aleph 2.0** | Couche d'**édition vidéo-par-vidéo** : tu donnes une vidéo existante + une instruction (« passe en nuit », « change le décor »), il la réédite. La référence la plus citée hors benchmarks d'édition |
| **Act-Two** | Transfert de performance sur personnage (ton jeu d'acteur → personnage généré) |
| Benchmark | Elo 1247 revendiqué (#1 mondial, doc client Q1 2026) — indépendant NON VÉRIFIÉ |
| Prix (docs secondaires concordantes) | 1 crédit = $0,01 : Gen-4.5 = 12 crédits/s (**$0,12/s**) ; Gen-4 Turbo = 5 ($0,05/s) ; Aleph = 15 ($0,15/s) ; Act-Two = 5 ($0,05/s) |
| Abonnements | Standard $12/mois (625 crédits), Pro $28, Unlimited $76–95 |
| Accès | dev.runwayml.com (crédits API sans expiration, top-up $10 min), app Runway |

Le modèle économique est limpide : 1 crédit = 1 centime, 10 secondes de Gen-4.5 = 120 crédits = $1,20. L'édition (Aleph) coûte 25 % plus cher que la génération — logique : l'édition doit comprendre la vidéo source avant de la transformer.

## 43. Alibaba Wan — lignée 2.2 → 2.6 → 2.7

Le cas d'école des **poids ouverts partiels** : la licence change en cours de lignée, il faut suivre quelle version est vraiment ouverte.

| Version | Date | Statut poids |
|---|---|---|
| Wan 2.2 | Juillet 2025 | **Apache 2.0** — self-host possible (14B ≈ 65–80 Go VRAM ; TI2V-5B tient sur RTX 4090) |
| Wan 2.5 | Sept. 2025 | API-only fermé (plateforme Bailian) |
| Wan 2.6 | Déc. 2025 | API-only fermé |
| Wan 2.7 | Mars 2026 | **Contradiction** : Apache 2.0 partiel (bases ouvertes, premium en API) selon une source vs API-only comme 2.5/2.6 selon une autre — non tranché |
| Wan 3.0 | Attendu mi-2026 | — |

Specs 2.6/2.7 (secondaires) : T2V/I2V/V2V, 1080p, 5–15 s, synchronisation audio-visuelle, variante Flash 720p 5–10 s. Prix secondaires : fal ~$0,10/s flat (2.7), Atlas Cloud « dès $0,07/s » (2.6) — NON VÉRIFIÉS.

Seul fait solide : **Wan ≤ 2.2 = Apache 2.0**. Si tu veux de la vidéo en local et en licence permissive, c'est Wan 2.2 — en acceptant une génération d'écart avec l'état de l'art.

## 44. Luma Ray3.2 / Ray3.14 (Dream Machine)

2026, dates exactes NON VÉRIFIÉES. Le positionnement « beau rendu » de Luma.

| Champ | Valeur |
|---|---|
| Clips | ~9 s par génération ; 1080p ; références image en entrée (selon comparatifs) |
| Prix (secondaires) | fal : 540p $0,03/s, 720p $0,06/s — NON VÉRIFIÉ |
| Accès | API Luma, fal.ai |

Luma a historiquement joué la carte de l'esthétique (Dream Machine était réputé pour ses rendus « cinéma »). ~9 s par clip, c'est court : usage typique = plans d'illustration, pas narration longue.

## 45. MiniMax Hailuo Video-02 / 2.3

2026, dates exactes NON VÉRIFIÉES. MiniMax (aussi connu pour la musique IA et les LLM) décline en Std / Fast / Pro.

| Champ | Valeur |
|---|---|
| Clips | 6–10 s ; 1080p ; génération par clip (~$0,50/vidéo historique sur 02) |
| Prix (secondaires) | fal Fast ~$0,03/s par clip — NON VÉRIFIÉ |
| Accès | API MiniMax, fal.ai |

## 46. PixVerse V6

2026, date exacte NON VÉRIFIÉE. Le positionnement prix agressif.

| Champ | Valeur |
|---|---|
| Clips | 1080p ; ~5–8 s ; tarification variable selon résolution |
| Prix (secondaires) | $0,033 (360p, sans audio) – $0,150 (1080p, avec audio) /s — NON VÉRIFIÉ |
| Accès | MuAPI et agrégateurs |

La grille « par résolution × avec/sans audio » est le standard 2026 : le prix suit le nombre de pixels calculés et le passage audio. Retenir l'ordre de grandeur : diviser la résolution par 3 (1080p → 360p) divise le prix par ~4,5.

## 47. Lightricks LTX 2.3 (poids ouverts)

2026, date exacte NON VÉRIFIÉE. Le seul modèle vidéo **à poids ouverts** récent du comparatif — avec une licence à plafond.

| Champ | Valeur |
|---|---|
| Clips | Jusqu'à 20 s ; 4K50 (selon un comparatif) |
| Vitesse | 18× plus rapide que WAN 2.2 en inférence (benchmarks communautaires cités) |
| Faiblesses signalées | Cohérence produit inter-trames, rendu de texte |
| Prix (secondaires) | fal : Pro $0,06/s, Fast $0,04/s (1080p) ; first-party $0,04–0,06/s — NON VÉRIFIÉ |
| Licence | Poids ouverts ; **plafond de 10M$ d'ARR** — au-delà, licence commerciale à négocier |
| Accès | Self-host, Replicate, fal.ai, LTX Studio |

La licence à plafond d'ARR (annual recurring revenue) est un modèle hybride 2026 : ouvert pour les petits acteurs, payant quand tu scales. Pour un usage perso, c'est transparent. 20 s de clip en poids ouverts, c'est le plus long du lot open.

## 48. Mentions courtes (vidéo)

- **Vidu Q3 Pro** (ShengShu) : ~$0,04/génération, 1080p, 8–16 s, 7 images de référence ; **Vidu S2** en variante.
- **Grok Imagine 1.5-preview** (xAI) : 15 s / 1080p, x.ai $0,08 (480p)–0,25/s (1080p) — NON VÉRIFIÉ.
- **Pika 2.5**, **Adobe Firefly Video** : repérés comme concurrents, fiches non ouvertes — NON VÉRIFIÉ.
- Comparatifs utiles pour creuser : les pages « sora-alternatives » de Kling AI, et les comparatifs Sora/Veo/Kling/Seedance (lushbinary).

## 49. Tableau comparatif vidéo

| Modèle | Durée max | Résolution | Audio natif | Prix ordre de grandeur | Accès |
|---|---|---|---|---|---|
| Sora 2 / Pro | 20–25 s | 720p / 1080p | Oui | Historique $0,10–0,70/s | **API fermée 24/09/2026** |
| Seedance 2.0 | 15 s | 720p (1080p/4K hébergeurs) | Oui + lip-sync | $0,01–0,68/s selon chemin | API partenaires |
| Seedance 2.5 | 30 s | 4K | Oui | ~$0,05–0,25/s (?) | Agrégateurs |
| Kling 3.0 | 15 s | 4K/60 revendiqué | Oui + lip-sync 5 langues | ~$0,08–0,17/s (?) | API + Runway |
| Runway Gen-4.5 | 10 s (extensible) | 1080p | — | $0,12/s | API Runway |
| Runway Aleph | Édition | 1080p | — | $0,15/s | API Runway |
| Wan 2.6/2.7 | 15 s | 1080p | Oui | ~$0,07–0,10/s (?) | API Alibaba / fal |
| Wan 2.2 | — | — | — | Gratuit (self-host) | **Poids ouverts Apache 2.0** |
| Luma Ray3.2 | ~9 s | 1080p | — | ~$0,03–0,06/s (?) | API Luma / fal |
| Hailuo 2.3 | 10 s | 1080p | — | ~$0,03/s (?) | API MiniMax / fal |
| PixVerse V6 | 8 s | 1080p | En option | $0,033–0,150/s (?) | Agrégateurs |
| LTX 2.3 | 20 s | 4K50 (?) | — | $0,04–0,06/s (?) | **Poids ouverts** (plafond ARR) |

## 50. Veo 3.1 et Gemini Omni Flash

