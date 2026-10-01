---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-50
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "DeepSeek", "Google", "Meta", "Microsoft", "Moonshot", "OpenAI", "United States", "xAI"]
dates: []
keywords: ["agent", "agents", "agi", "asl", "aws", "bedrock", "benchmark", "benchmarks", "claude", "copyright", "deepseek", "distribution"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [3872, 3956]
sha256: 4133af471c991fd2b2edf84eadfc3ab3ece9ac1cfdf35dffb2266f18708dbe45
---

# IA — Le grand dossier

| Date | Événement | Source |
|---|---|---|
| **Janvier 2021** | Fondation d'**Anthropic** à San Francisco par **Dario Amodei** (CEO), **Daniela Amodei** (présidente) et 5 autres ex-chercheurs seniors d'OpenAI, après des désaccords sur la direction sécurité/commerciale d'OpenAI | Wikipedia / presse |
| 2021 | Statut de **Public Benefit Corporation** (société à mission) — l'entreprise a une mission d'intérêt public inscrite dans ses statuts | — |
| Avril 2022 | Levée de 580 M$, dont **500 M$ de FTX** (Sam Bankman-Fried) | Wikipedia |
| Mars 2023 | **Claude 1** (accès restreint) ; juillet 2023 : **Claude 2** (public) | — |
| Sept. 2023 → nov. 2024 | **Amazon** investit au total **8 Md$** (1,25 + 2,75 + 4) ; AWS devient le cloud principal ; Claude disponible sur Bedrock | Wikipedia |
| Oct. 2023 → mars 2025 | **Google** investit 500 M$ + 1,5 Md$ engagés, puis 1 Md$ de plus (mars 2025) ; Claude sur Vertex AI / Gemini ? (non : sur **Google Cloud Vertex AI**) | Wikipedia |
| Mars 2024 | Famille **Claude 3** (Haiku, Sonnet, Opus) — Opus premier modèle à battre GPT-4 sur certains benchmarks | — |
| Juin 2024 | **Claude 3.5 Sonnet** — référence code durable | — |
| Mai 2025 | **Claude 4** (Opus 4, Sonnet 4) ; lancement de **Claude Code** (80,9 % à SWE-bench Verified — « à vérifier » sur la version exacte du benchmark) | — |
| Mars 2025 | Série E : **3,5 Md$ à 61,5 Md$** de valorisation (Lightspeed en tête) | Wikipedia |
| Sept. 2025 | Série F : **13 Md$ à 183 Md$** (Iconiq, Fidelity, Lightspeed ; Qatar Investment Authority) | Wikipedia |
| Déc. 2025 | Term sheet **10 Md$ à 350 Md$** (Coatue, GIC) | Wikipedia |
| 12 fév. 2026 | Série G : **30 Md$ à 380 Md$** de valorisation | Wikipedia |
| Mai 2026 | Série H : **65 Md$ à 965 Md$** (Altimeter, Dragoneer, Greenoaks, Sequoia…) — dépasse alors OpenAI (852 Md$) ; revenus en run-rate ~40-47 Md$ ; 8 du Fortune 10 clients ; IPO visée fin 2026 | siliconvalleyinvestclub / Clare Capital, juin 2026 |
| Janv. 2026 | Famille **Claude 4.5** (Opus, Sonnet, Haiku) + plateforme « Labs » + Agent SDK | ai-research-arm |
| Fév. 2026 | **Claude Opus 4.6**, puis Sonnet 4.6 | Wikipedia |
| Avril 2026 | Annonce de **Mythos** (non public, ~40 organisations) | Wikipedia / vol. IA |
| 22 sept. 2026 | **Claude Opus 5.5** : comparable au haut de gamme « Fable 5.1 », **40 % moins cher** | Wikipedia |
| Juil. 2026 | Accord copyright **1,5 Md$** approuvé (voir 1.4) — plus gros accord de droit d'auteur US connu | AP/Reuters |

**Positionnement revendiqué** : « AI safety and research company » — le labo qui veut prouver qu'on
peut être à la fois *le plus sûr* et *parmi les plus capables*. C'est un positionnement commercial
autant que moral : la sécurité est leur différenciateur marketing face à OpenAI et Google.

### 3.2. L'axe sécurité : Constitutional AI, expliqué simplement

**Le problème de départ.** La méthode standard d'alignement (RLHF — apprentissage par renforcement
avec feedback humain) a des limites : elle dépend d'annotateurs humains coûteux, elle apprend au
modèle à *plaire à l'évaluateur* (sycophancy : le modèle dit ce que tu veux entendre), et elle
transmet mal des principes abstraits (« sois honnête »).

**L'idée de Constitutional AI (Anthropic, 2022).** Au lieu de dire au modèle « fais plaisir aux
humains », on lui donne une **Constitution** : une liste écrite de principes (inspirés de la
Déclaration universelle des droits de l'homme, de règles de bon sens, des conditions d'utilisation
d'Anthropic). Puis on demande au modèle de **se critiquer et se corriger lui-même** à la lumière
de cette Constitution :

1. **Phase supervisée** : le modèle génère une réponse, puis génère une *critique* de sa réponse
   (« en quoi cela viole-t-il le principe X ? »), puis une version *révisée*. On l'entraîne sur
   les versions révisées.
2. **Phase RL « RLAIF »** : au lieu d'humains, c'est le modèle lui-même (guidé par la Constitution)
   qui note les réponses ; on optimise par renforcement sur ces notes automatiques.

**Pourquoi c'est malin (et ses limites) :**

- ✅ Scalable : pas besoin d'armées d'annotateurs pour chaque principe.
- ✅ Explicite : les principes sont *lisibles* (la Constitution de Claude est publique — « à vérifier »
  sur la version en vigueur), donc auditables et débattables — contrairement à des préférences
  enfouies dans des millions de comparaisons humaines.
- ✅ Moins servile : le modèle vise des principes, pas l'approbation de l'évaluateur.
- ⚠️ Limite 1 : une Constitution écrite par Anthropic reste un choix *politique* privé — qui décide
  des principes ? (Critique récurrente, y compris interne au milieu académique.)
- ⚠️ Limite 2 : l'auto-critique peut rater ce que le modèle ne voit pas lui-même (angles morts
  partagés entre le juge et le jugé — c'est la même tête).
- ⚠️ Limite 3 : ça n'empêche pas les jailbreaks (voir 2.1.2) — l'alignement n'est pas la sécurité.

**Autres piliers sécurité d'Anthropic (factuel) :**

- **Interpretability** : équipe de recherche sur la « mechanistic interpretability » (comprendre les
  circuits internes des modèles — travaux publiés sur les « features » de Claude).
- **Responsible Scaling Policy** : seuils de capacités (ASL — AI Safety Levels) déclenchant des
  mesures de sécurité renforcées ; red-teaming interne et externe.
- **Bug bounty** : programme public — les CVE Claude Code 2025-2026 listées en 2.1.1 montrent qu'il
  est actif (ex. : 100 $ d'Anthropic pour « Comment and Control », CVSS 9.4 — montant modeste
  qui a fait débat dans la communauté sécu).

### 3.3. Ce qui distingue Claude des concurrents (factuel, vérifié)

| Dimension | Claude / Anthropic | Concurrents |
|---|---|---|
| **Positionnement** | « Le labo de la sécurité » ; entreprise à mission (PBC) ; discours public sur les risques (Amodei) | OpenAI : « l'AGI pour tous », virage commercial assumé (PBC en 2025, levées géantes) ; Google : intégration produit ; xAI : vitesse ; Meta : open weights |
| **Code** | **Point fort historique** : Claude 3.5/4 Sonnet sont devenus la référence des développeurs ; Claude Code (mai 2025) + Agent SDK (janv. 2026) en font un standard des agents de dev | GPT-5.x / Codex (OpenAI), Gemini (Google) : concurrence frontale, écarts faibles et mouvants |
| **Fenêtre de contexte** | 200K en standard, **1M en bêta** sur les générations récentes (4.6+, 5) | Comparable chez Google (Gemini 1M-2M) ; variable chez OpenAI |
| **MCP** | **Anthropic a créé le Model Context Protocol** (fin 2024) — devenu le standard ouvert de connexion outils/données, adopté par OpenAI, Google, Microsoft | Les concurrents *adoptent* le standard d'Anthropic — fait notable |
| **Distribution** | API directe + **AWS Bedrock + Google Cloud Vertex AI + Microsoft Azure + Databricks** (triple ancrage cloud, fait rare) | OpenAI : Azure d'abord ; Google : son cloud |
| **Structure financière** | ~965 Md$ (mai 2026), ~144 Md$ levés au total (ordre de grandeur, « à vérifier ») ; Amazon (8 Md$) et Google (~3,5 Md$+) au capital — dépendance stratégique aux deux hyperscalers | OpenAI ~852 Md$ (2026), Microsoft au capital ; xAI adossé à X/Tesla |
| **Transparence recherche** | Publications sécurité/interprétabilité régulières ; « system cards » par modèle | Pratiques variables ; OpenAI critiqué sur l'opacité croissante |

**Ce que Claude n'est pas** (honnêteté du dossier) : ni le moins cher (l'open source chinois —
DeepSeek, Qwen, Kimi — casse les prix, cf. vol. IA), ni le plus multimodal grand public, ni le
plus rapide à dégainer des features virales. Son pari : **la confiance des entreprises** (8 du
Fortune 10, 1 000+ comptes à 1 M$/an — chiffres 2026, « à vérifier » sur définitions).

