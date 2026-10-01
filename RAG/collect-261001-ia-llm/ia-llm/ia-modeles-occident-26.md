---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-26
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Apple", "Broadcom", "Glasswing", "Google", "Microsoft", "Mistral", "Nvidia", "OpenAI", "Poolside", "xAI"]
dates: ["2025-10-31", "2026-02-24", "2026-03-09", "2026-03-18", "2026-05-15", "2026-06-15", "2026-06-25", "2026-06-30", "2026-07-28", "2026-07-31", "2026-08-31", "2026-09-29", "2026-10-15", "2026-11-02", "2026-11-24", "2027-02-04"]
keywords: ["arr", "chatgpt", "claude", "copilot", "cyber", "gemini", "gpt-5.6", "grok", "grok 4", "luna", "mistral", "mythos 5"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [2003, 2061]
sha256: cb27590df37b95572641e17ac7ced654529f38c7595738830bfde99a8bdc4f12
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Modèle(s) | Date de retrait / EOL | Remplacement recommandé | Source |
|---|---|---|---|
| Gemini 3 Pro Preview | **retiré 09/03/2026** | `gemini-3.1-pro-preview` (slug auto-redirigé) | changelog Gemini |
| Grok 3 / 3 Mini, Grok Code Fast 1, Grok 4.1 Fast, `grok-imagine-image-pro` | **retirés 15/05/2026 12h00 PT** | slugs → `grok-4.3` | modeldeprecations.dev |
| Sonnet 4 + Opus 4 (originels) | **retirés 15/06/2026** | Sonnet 5 / Opus 4.6+ | endoflife.date |
| `magistral-medium/small-2509` | dépréciés (fenêtre → 31/07/2026) | Mistral Medium 3.5 / Small 4 | docs Mistral |
| `devstral-2512` | déprécié (→ 31/07/2026) | Mistral Medium 3.5 | docs Mistral |
| `open-mistral-nemo` | déprécié (→ 31/07/2026) | Ministral 3 8B | docs Mistral |
| `magistral-medium/small-2506`, `devstral-medium-latest` | dépréciés depuis 31/10/2025 | — | docs Mistral |
| Gemini 3 Flash | GA → migration recommandée vers `gemini-3.8-flash` | `gemini-3.8-flash` | guide migration Google |
| Gemini 3.5 Flash | « legacy pipelines only » (pas de shutdown daté) | 3.6/3.7/3.8 Flash | guides tiers |
| Gemini 3.7 Flash | migration recommandée vers 3.8 (pas de shutdown daté) | `gemini-3.8-flash` | guides tiers |
| Nano Banana (original, `gemini-2.5-flash-image`) | legacy | Nano Banana 2 Lite | guides Google |
| IDs preview `gemini-3.1-flash-image-preview`, `gemini-3-pro-image-preview` | **dépréciés 25/06/2026** | IDs stables | guide migration |
| `grok-2-image-1212` | déprécié 24/02/2026 | `grok-imagine-image-2.0` (calendrier : 02/11/2026) | calendrier xAI |
| GPT-5.4 (Codex via login ChatGPT) | **retiré 31/08/2026** | GPT-5.6 Terra/Luna | annonce OpenAI |
| Claude Sonnet 4.5 | **EOL 29/09/2026** | Sonnet 5 | endoflife.date |
| Claude Haiku 4.5 | **EOL 15/10/2026** | Haiku 5.5 (annoncé) | endoflife.date |
| Claude Opus 4.5 | **EOL 24/11/2026** | Opus 5.5 | endoflife.date |
| Veo 2 / Veo 3 (via Gemini API) | **arrêtés 30/06/2026** | Veo 3.1 | innfactory.ai |
| Laguna XS.2 / M.1 (free tiers) | free tiers fermés 09/07 et 28/07/2026 | XS 2.1 / S 2.1 | Poolside |

Enseignement : en 2026, un modèle a une durée de vie moyenne de 4 à 8 mois côté API. Seule exception : **GPT-5.3-Codex LTS** (garanti jusqu'au 04/02/2027) — voir § 157.

## 157. Dossier : LTS et longévité — le seul engagement du marché

- **GPT-5.3-Codex**, désigné **LTS le 18/03/2026**, garanti jusqu'au **4 février 2027** pour Business/Enterprise Copilot : **premier et seul modèle LTS** documenté dans ce volume. C'est la réponse d'OpenAI aux DSI qui ne peuvent pas requalifier leur stack tous les trimestres.
- **Anthropic** publie des dates EOL explicites (Sonnet 4.5 : 29/09/2026 ; Haiku 4.5 : 15/10/2026 ; Opus 4.5 : 24/11/2026) via endoflife.date — prévisibilité sans engagement LTS.
- **Mistral** documente ses fenêtres de dépréciation (docs officielles) avec remplacements nommés — la plus propre des politiques européennes.
- **Google** recommande des migrations (3 Flash → 3.8, 3.5 → 3.6/3.7/3.8) mais **sans date de shutdown** annoncée — prévisibilité faible.
- **xAI** retire par vagues datées (15/05/2026 12h00 PT) avec redirection de slugs — brutal mais lisible.
- Conseil : en entreprise, exiger par écrit la politique de dépréciation du fournisseur **avant** de construire dessus. Le LTS de Codex est l'exception qui confirme l'absence de règle.

## 158. Dossier : tokenizers et coûts cachés — la facture réelle

Quatre mécanismes font diverger le prix affiché du prix payé :

1. **Nouveaux tokenizers** (Claude Opus 4.7, Opus 4.8, Sonnet 5) : 1,0–1,35× tokens vs génération précédente → **+35–47 % de coût réel** mesuré par la communauté à prix affiché inchangé. Migrer de Sonnet 4.6 à Sonnet 5 sans re-mesurer, c'est accepter une hausse invisible.
2. **Paliers long-contexte** : OpenAI facture ×2 (in) / ×1,5 (out) au-delà de **272K** tokens de prompt ; xAI double Grok 4.6 au-delà de **200K** ($2/$6 → $4/$12) ; Anthropic facturait $10/$37,50 au-delà de 200K sur Opus 4.6 (bêta 1M). Le « 1M de contexte » a un prix à deux étages.
3. **Seuils de prompt caching** : le cache ne s'active qu'au-delà d'un minimum — 512 (Opus 5, Fable), 1 024 (Opus 4.8, Sonnet 5), 2 048 (Opus 4.7, Mythos Preview), 4 096 (Opus 4.6, Haiku 4.5). En dessous, plein tarif. Et l'écriture du cache coûte 1,25× (Anthropic) — le cache-write pricing d'OpenAI (depuis 5.6) va dans le même sens.
4. **Sortie max vs sortie facturée** : la sortie max (128K quasi partout, 64K chez Gemini Flash, 16K chez Sonnet 4.5) est un plafond, pas une promesse de débit — et chaque token de sortie coûte 3 à 10× le token d'entrée.

Checklist avant mise en production : mesurer tokens/requête sur **vos** prompts (pas les exemples du vendor), activer le caching et vérifier les hit rates, tester l'effort `low` vs `high` sur un échantillon, et re-mesurer à chaque changement de modèle — surtout de tokenizer.

## 159. Dossier : les programmes d'accès restreint comparés

| | Project Glasswing (Anthropic) | Daybreak / Daybreak Red (OpenAI) | Fairwind Program (Google) |
|---|---|---|---|
| Modèles | Mythos, Mythos 5, Mythos 5.1 | GPT-5.4-Cyber, 5.5-Cyber, 5.6-Cyber | Gemini 3.5 / 3.8 Flash Cyber |
| Bénéficiaires | ~40 orgs vérifiées (Amazon, Microsoft, Google, Apple, CrowdStrike, Cisco, JPMorgan, Linux Foundation, Nvidia, Broadcom, Palo Alto) + chercheurs bio | trusted access (détails non publics) | gouvernements, opérateurs d'infrastructures critiques, mainteneurs |
| Cadre | NDA, « Responsible Deployment Protocol v3 » (selon une source) | accès contrôlé (détails non publics) | accès restreint (détails non publics) |
| Dotation | **$100M de crédits + $4M de dons** open-source sécu | non vérifié | non vérifié |
| Prix | $25/$125 par MTok (Mythos Preview, **source unique**) | non vérifié | **non publié** |
| Philosophie | ne pas publier le modèle le plus puissant | publier le flagship, restreindre les variantes cyber | variante assouplie du Flash |

Fait notable : Anthropic est le seul lab à avoir **refusé de commercialiser** son meilleur modèle (Mythos) — les deux autres restreignent seulement des variantes. C'est aussi le seul à chiffrer publiquement sa dotation ($100M). Pour un décideur : ces programmes sont des **canaux d'accès**, pas des produits — il faut candidater, signer un NDA, et accepter un cadre d'usage.

## 160. Dossier : le guide des licences open-weight (Occident, sept. 2026)

