---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-34
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Cerebras", "DeepSeek", "Google", "Groq", "LongCat", "Microsoft", "MiniMax", "Mistral", "Moonshot", "OpenAI", "OpenRouter", "United States", "Z.ai"]
dates: []
keywords: ["agent", "agents", "apache", "astra", "benchmarks", "cost", "deepseek", "embedding", "embeddings", "glm", "gpt-6", "gpu"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [2415, 2498]
sha256: fe211a523201c8b66611b26821e360b458b75bbb2596bbf21c620d02f08c7e91
---

# IA — Le grand dossier

1. **Spend logs** (Postgres) : chaque appel loggue modèle, provider, tokens in/out, coût calculé, latence, clé virtuelle. Requêtes SQL types pour ton pilotage :
   ```sql
   -- Coût par modèle et par jour (table LiteLLM_SpendLogs)
   SELECT "model", date("startTime") AS jour,
          SUM("spend") AS usd, COUNT(*) AS requetes,
          AVG("endTime"-"startTime") AS latence_moy
   FROM "LiteLLM_SpendLogs"
   GROUP BY 1, 2 ORDER BY 2 DESC, 3 DESC;
   ```
2. **UI** (`:4000/ui`) : dashboard d'usage, génération de **clés virtuelles** (une par appli/équipe, avec `max_budget`, `models` autorisés, `tpm`/`rpm`), test en direct.
3. **Prometheus** (`/metrics`) : latence, erreurs, tokens/s par deployment — à scraper dans ton Zabbix/Prometheus existant (tu as déjà les guides).
4. **Alertes budget** : webhook/Email quand une clé dépasse X % de son budget — indispensable avant de brancher un agent autonome (boucle infinie = facture infinie).
5. **Health** : `GET /health/readiness` (sans auth) pour tes sondes ; `/health/liveliness` pour le détail par deployment.

### 4.7. Routage « intelligent » : au-delà du round-robin

- **Par coût** : `cost-based-routing` + groupes ordonnés (pas cher d'abord). Idéal pour l'indexation RAG (millions de tokens, qualité « suffisante »).
- **Par qualité** : un **juge** — une première passe bon marché, puis escalade vers le modèle frontier uniquement si un score de confiance est bas. Pattern : `triage` répond + auto-évalue (`"confiance": 0.0-1.0` en JSON) ; si < 0,7 → rappel vers `raisonnement` avec le contexte. Coût moyen divisé par 3 à 10 selon les benchmarks internes rapportés (à vérifier sur tes données).
- **Par latence** : `latency-based-routing` + TTFT (time-to-first-token) comme métrique : Groq/Cerebras pour l'interactif, Together/DeepInfra pour le batch.
- **Par compétence** : routage **sémantique** en amont (un petit classifieur local : « code → qwen3-coder », « résumé → luna », « raisonnement → opus »). Des routeurs dédiés existent (ex. **Not Diamond**, routeur qualité-prix) mais un classifieur maison de 8B fait l'affaire et ne coûte rien.
- **Garde-fous** : `max_tokens` par groupe, `timeout` par appel, **budget dur par clé** (la requête est rejetée au-delà), allowlist de modèles par équipe (la compta n'a pas accès à GPT-6 Astra à 50 $/1M en sortie).

### 4.8. Schéma d'architecture cible (PME/ETI)

```
[Apps / Agents / RAG] ──1 clé virtuelle par app──▶ [LiteLLM :4000]
                                                        ├─▶ triage ──▶ Groq ──┐
                                                        │              ├─▶ Cerebras ├─ fallback ─▶ raisonnement
                                                        │              └─▶ OpenRouter ──┘
                                                        ├─▶ raisonnement ─▶ Anthropic ─┐
                                                        │                   └─▶ OpenAI ─┴─ fallback ─▶ triage
                                                        ├─▶ local ─▶ Ollama (LAN, données sensibles)
                                                        └─▶ embeddings ─▶ text-embedding-3-small
                                                              │
                                              [Postgres: spend logs] [Prometheus: métriques]
                                                              │
                                              [Alertes budget] ──▶ [Zabbix / mail]
```

**Règle d'or** : le code applicatif ne connaît que les noms logiques (`triage`, `raisonnement`, `local`, `embeddings`). Changer de provider = éditer le YAML, pas redéployer les apps. C'est exactement le même principe qu'un reverse proxy ou qu'un DNS : **l'indirection est la résilience**.

---

## 5. L'IA locale en profondeur

### 5.1. Pourquoi faire tourner des modèles chez soi ? (les 4 vraies raisons)

| Raison | Détail | Exemple concret (métier de Zelef) |
|---|---|---|
| **Coût** | Après l'achat du GPU, le coût marginal ≈ l'électricité (~0,04–0,06 $/h pour un GPU 350 W). Au-delà d'un seuil de volume, c'est imbattable (section 8) | Indexer 10M de pages de doc technique : en API ~200–2 000 $ ; en local ~quelques $ d'élec |
| **Confidentialité** | Les données ne quittent jamais le site. Aucune clause de « zero retention » à négocier, aucun DPA à faire signer | Logs d'onduleurs, plans électriques, contrats de maintenance : restent dans la baie |
| **Offline / résilience** | Fonctionne sans internet, sans quota, sans « modèle délisté du jour au lendemain » | Site isolé, intervention terrain, coupure fibre : le RAG répond quand même |
| **Souveraineté** | Aucune dépendance à un lab US/CN, aucune télémétrie, reproductibilité (même poids, même réponse, dans 5 ans) | Appel d'offres public exigeant l'hébergement FR/UE des données |

**Et les 3 mauvaises raisons** (autant être honnête) : « c'est gratuit » (le GPU a un coût d'amortissement), « c'est aussi bon » (un 8B local ≠ un frontier ; il faut choisir les tâches adaptées), « c'est simple » (c'est de l'exploitation : drivers, VRAM, mises à jour de poids — ton métier, justement).

### 5.2. Les modèles ouverts : familles, tailles, licences (état sept. 2026)

**Rappel juridique** : « open-weight » ≠ « open source » (OSI). La plupart publient les **poids** mais ni le code d'entraînement ni les données. Pour un usage pro, la **licence** compte autant que la qualité :

| Licence | Ce qu'elle permet | Modèles concernés (sept. 2026) |
|---|---|---|
| **Apache 2.0** | Usage commercial libre, brevet inclus ; la plus sûre juridiquement | Qwen3/3.5 (toutes tailles), GPT-OSS-20B/120B, Gemma 4, Tencent Hunyuan 3, IBM Granite, Phi (MIT, équivalent) |
| **MIT** | Comme Apache, sans clause brevet explicite ; très permissive | DeepSeek V4/V3.2/R1 (+ distills), GLM-4.5/4.7/5.2, Kimi K2.6 (vérifier MAU), MiniMax M2.7, MiMo-V2.x Pro, LongCat-2.0 |
| **MIT modifiée / custom à seuils** | Commercial **sous conditions** (plafond de revenus ou de MAU) | Mistral Medium 3.5 (plafond revenus), Kimi K2.5 (100M MAU), MiniMax M2.5 (territoriale), Liquid LFM (10M $ revenus) |
| **Llama Community License** | Commercial OK mais **restrictions** : pas d'entraînement de concurrents au-delà de 700M MAU, clauses d'usage acceptable | Llama 3.3/4 (Scout, Maverick) |
| **Licences maison** | À lire au cas par cas | GLM-5.3 (abandon de MIT), Kimi K3, Nemotron OML, Gemma (custom Google, usage responsable) |

**Les familles à connaître par taille** (tailles « utiles » en local, pas les monstres 1T+) :

| Famille | Tailles locales pertinentes | Points forts | Licence |
|---|---|---|---|
| **Qwen3 / Qwen3.5** | 0.6B → 32B (+ Coder, VL) | Le meilleur généraliste ; thinking mode ; tool calling solide | Apache 2.0 |
| **Qwen2.5-Coder** | 7B, 14B, 32B | Toujours une référence code à petit prix | Apache 2.0 |
| **Llama 3.3** | 8B, 70B | Écosystème immense (tutos, quantifs) | Llama Community |
| **Mistral Small / Ministral** | 3B → 24B | Efficacité, français correct | Apache 2.0 (vérifier par variante) |
| **Devstral 2** | 24B (Small 4 ensuite) | Code agentique | MIT modifiée — à vérifier |
| **Gemma 4** | 2B → 27B | Propre, multimodal (image+audio sur MLX) | Apache 2.0 |
| **Phi-4** | 3.8B, 14B | Raisonnement compact, MIT | MIT |
| **DeepSeek-R1-Distill** | 8B, 14B, 32B | Raisonnement (chaîne de pensée) | MIT |
| **GPT-OSS** | 20B, 120B | Raisonnement OpenAI ouvert ; 20B tient sur 16 Go RAM | Apache 2.0 |
| **Nomic Embed** | 137M | **Embeddings** locaux (274 Mo) | Apache 2.0 |
| **Qwen3-Embedding** | 0.6B → 8B | Embeddings multilingues (dont français) | Apache 2.0 |

