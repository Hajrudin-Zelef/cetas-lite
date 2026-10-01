---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-9
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["Google", "Huawei", "Microsoft", "Moonshot", "OpenAI", "Perplexity", "United States", "Xiaomi"]
dates: ["2026-07-16", "2026-09-15", "2026-09-27"]
keywords: ["agents", "chatgpt", "copilot", "distribution", "gemini", "gpu", "ipo", "kimi", "perplexity"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [774, 865]
sha256: e6c3cf262ebf03124e34bb50a6e8b7ddfceb71af71e0f1fbd6391e178e2fc1f0
---

# IA générative : image, vidéo, recherche

| Critère | Kling 3.0 | Veo 3.1 | Runway Gen-4.5 | Luma Ray 3.2 | Hailuo H3 | Seedance 2.5 |
|---------|-----------|---------|----------------|--------------|-----------|--------------|
| Durée max / appel | 15 s | 8 s | ~10 s | ~10 s | 15 s | **30 s** |
| Résolution max | **4K native** | 4K (Standard) | 4K (upscale) | 4K (upscale) | 2K | 1080p |
| Audio natif | Oui (+50 %) | Oui (inclus) | Limité | **Non** | Oui | Oui |
| Coût indicatif / 10 s 1080p | ~1,1-1,7 $ (API) | 4,0 $ (Standard) / 1,2 $ (Fast) | ~5-6 $ (équiv.) | ~2-3 $ (équiv.) | 2,60 $ | ~2,9 $ |
| API officielle | Oui (packs prépayés) | Oui (Gemini/Vertex) | Oui (crédits) | Oui | Oui | Oui (BytePlus) |
| Usage local | Non | Non | Non | Non | Partiel (**licence à vérifier**) | Non |
| Point fort | Qualité/prix, motion transfer | Réalisme premium | Régularité, suite pro | Atmosphérique | Vitesse, prix | Mouvement précis, 30 s |
| Usage commercial | Oui (plans payants) | Oui (API) | Oui (plans payants) | Oui (dès Plus) | Oui (**licence à vérifier**) | Oui (CGU) |

## 53. Vidéo : prompts efficaces — spécificités

Par rapport à l'image, ajouter systématiquement : **la durée ressentie** (« en 8 secondes »), **le mouvement de caméra**, **le rythme** (« lent », « dynamique »), **le son** si audio natif. Exemple « présentation d'équipement » :

> « plan séquence de 10 secondes : travelling avant lent vers un onduleur triphasé dans un local technique, LEDs vertes qui clignotent, léger flou d'arrière-plan sur les baies voisines, lumière froide industrielle, ambiance feutrée professionnelle. Audio : ronronnement grave et régulier de ventilation, pas de voix. Style documentaire corporate, photoréaliste. »

Et en **image-to-video** (souvent meilleur pour du pro) : partir d'une **vraie photo** du matériel + prompt minimal de mouvement (« léger travelling avant, les voyants clignotent doucement, 8 secondes »). Le réel ancre la crédibilité ; l'IA n'ajoute que le mouvement.

---

# PARTIE F — PANORAMA RECHERCHE IA : les alternatives à Perplexity

## 54. Pourquoi comparer

Perplexity est le mieux packagé, pas le seul. Selon le besoin (code, vie privée, personnalisation, écosystème chinois, intégration Microsoft), une alternative peut être meilleure. Toutes les suivantes existent et sont actives en sept 2026.

## 55. Phind : la recherche pour développeurs

**Phind** reste en 2026 le moteur de recherche IA **orienté code** : questions techniques, réponses avec extraits de code et citations de documentation. Pour un sysadmin qui cherche « exemple nftables rate-limit SSH » ou « script Python Netmiko sauvegarde configs Huawei », c'est souvent plus direct que Perplexity généraliste. Positionnement : moins « moteur de réponse grand public », plus « pair programmer qui cite la doc ».

## 56. You.com : la recherche personnalisable

**You.com** : plateforme de recherche IA **personnalisable** (sources préférées, apps intégrées, agents). Intérêt : construire une recherche « à ta main » (ex. prioriser la doc constructeur). En retrait de Perplexity en notoriété, mais vivant et utile pour des workflows de recherche répétitifs et paramétrables.

## 57. Kimi (Moonshot AI) : l'outsider chinois à suivre

- **Kimi**, assistant de **Moonshot AI** (Pékin) : recherche web intégrée, raisonnement profond, contexte ultra-long.
- **16/07/2026** : sortie de **Kimi K3** (modèle revendiqué à 2 800 Md de paramètres, **poids ouverts** sous licence Kimi K3 — le plus gros modèle ouvert jamais publié ; ~1,4 To de stockage, usage self-host réservé aux grosses infra).
- Fait marquant : **suspension des nouveaux abonnements** après le lancement (saturation des GPU — la demande a dépassé la capacité), IPO de Hong Kong en préparation (~3 Md$ visés selon Reuters, sept 2026).
- API : `platform.kimi.com` (compatible Chat Completions). Pertinent pour : veille sur l'écosystème chinois (docs Huawei, Xiaomi...), alternative géopolitique aux clouds US.

## 58. Microsoft Copilot : l'intégré

**Copilot** (ex-Bing Chat) : recherche IA **intégrée à l'écosystème Microsoft** (Edge, Windows, M365). Si ton poste et ta messagerie sont Microsoft, c'est la recherche IA « sans friction » : pas d'abonnement séparé, données dans ton tenant (selon ton contrat). Moins bon en citation rigoureuse que Perplexity, meilleur en intégration bureautique.

## 59. Les autres cités en 2026

**Brave Leo** (vie privée, navigateur Brave), **DuckDuckGo Search Assist** (vie privée), **Komo** (catégories ciblées), **Andi** (résultats visuels, gratuit), **ChatGPT avec recherche** et **Gemini / AI Mode de Google** (les assistants généralistes font aussi de la recherche web temps réel — voir volumes existants). La tendance : **la recherche IA devient une fonctionnalité partout**, plus un produit isolé.

## 60. Tableau comparatif recherche IA

| Critère | Perplexity | Phind | You.com | Kimi | Copilot | Google/Bing classique |
|---------|-----------|-------|---------|------|---------|----------------------|
| Citations sources | ★★★★★ | ★★★★ | ★★★ | ★★★★ | ★★★ | ★ (c'est toi qui cherches) |
| Synthèse rédigée | ★★★★★ | ★★★★ | ★★★★ | ★★★★ | ★★★★ | Non |
| Orientation code | ★★★ | ★★★★★ | ★★★ | ★★★★ | ★★★ | ★★★ (docs brutes) |
| API développeur | ★★★★★ (Sonar) | Limitée (**à vérifier**) | Oui (**à vérifier**) | Oui | Via Azure | Non (SERP payantes) |
| Vie privée | ★★ | ★★ | ★★★ | ★★ | ★★★ (tenant MS) | ★ |
| Prix entrée | 0 $ / 20 $ | Freemium (**à vérifier**) | Freemium (**à vérifier**) | Freemium / suspendu | Inclus M365 | 0 $ |
| Point fort | Recherche sourcée | Code + doc | Personnalisation | Contexte énorme, Chine | Intégration MS | Contrôle total, gratuit |

## 61. Quelle recherche pour quel besoin (pense-bête)

- **Vérifier un fait technique vite** → Perplexity (puis ouvrir la source).
- **Chercher un exemple de code/config** → Phind.
- **Veille automatisée / digest d'équipe** → Perplexity Spaces ou API Sonar.
- **Écosystème chinois / docs constructeurs CN** → Kimi.
- **Déjà dans M365, zéro friction** → Copilot.
- **Recherche obscure, opérateurs fins, gratuit** → Google/Bing classique (ça reste un outil pro).

---

# PARTIE G — JEV : le nom mystérieux, identifié

## 62. « jev » : identification (recherche du 27/09/2026)

Bonne nouvelle : **« jev » n'est pas une coquille**. C'est **Jev**, le premier modèle public de **TypeSafe AI** (San Francisco, fondée en 2024), sorti de la discrétion (**stealth**) le **15/09/2026**, avec une levée seed de **40 M$** (valorisation ~200 M$ selon Forbes). Vérifié le 27/09/2026 via la presse tech et la documentation développeur.

Point crucial : **Jev n'est PAS un outil de génération d'image ou de vidéo.** C'est un modèle de **décision** — ce que TypeSafe appelle « System One » (référence à Kahneman, *Thinking, Fast and Slow*) : là où un LLM « parle » (System 2 : lent, verbal), Jev **juge** (System 1 : rapide, pré-verbal). Il ne génère aucun texte, aucune image : il répond à des questions fermées par des **probabilités**.

## 63. Jev : ce qu'il fait, techniquement

- **Entrée** : un `state` (texte ou JSON : le contexte à juger) + une **liste de questions** nommées, évaluées **en parallèle**.
- **Trois types de questions, rien d'autre** :

| Type | Question | Retour |
|------|----------|--------|
| `noul` | jugement oui/non | P(oui) entre 0 et 1 |
| `choice` | choisir parmi ≤ 255 options | option + distribution de probabilités + confiance |
| `score` | noter selon des niveaux ordonnés | score + probabilités + confiance |

