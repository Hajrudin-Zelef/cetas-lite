---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-32
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "AWS", "Alibaba", "Anthropic", "Cerebras", "CoreWeave", "Crusoe", "DeepSeek", "Fireworks AI", "Google", "Groq", "Lambda", "Meta", "MiniMax", "Mistral", "Moonshot", "Nebius", "Nscale", "Nvidia", "OpenAI", "OpenRouter", "United States"]
dates: ["2026-08-26"]
keywords: ["agents", "agi", "amd", "apache", "arr", "aws", "backlog", "blackwell", "cost", "deepseek", "diffusion", "gpu"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [2180, 2264]
sha256: a7ff424ff8d63f84e77301db7a3964de46af5954c5e0e8fbee2d11676b24405f
---

# IA — Le grand dossier

| 37 | **IBM** | **watsonx** : Granite (Apache 2.0) + gouvernance | $/1M tokens ; licences entreprise | Secteurs régulés, gouvernance IA |
| 38 | **DigitalOcean** | **Gradient AI** : Inference Router (70+ modèles, une clé), GPU dédiés | $/1M tokens (0,10–1,05 $) ; dédié dès 2,59 $/h/GPU | PME : simplicité, une facture |
| 39 | **GitHub Models** | Marketplace de modèles dans GitHub (15 RPM / 150 RPD en free) | Gratuit (limites) puis Azure | Prototyper sans infra, côté code |
| 40 | **RunPod** | Pods GPU + serverless + clusters ; facturation à la seconde | H100 2,69–3,29 $/h ; 4090 ~0,44–0,74 $/h ; **egress 0 $** | Le meilleur rapport qualité/prix GPU pour un indépendant |
| 41 | **Lambda Labs** | Instances GPU bare-metal + 1-Click Clusters ; stack ML optimisée | H100 2,99–3,99 $/h ; **B200 6,69 $/h (le moins cher publié)** ; egress 0 $ | Entraînement/inférence sérieux sans hyperscaler |
| 42 | **Vast.ai** | **Marketplace P2P** de GPU (particuliers + DC) ; spot | 4090 dès ~0,14 $/h (spot) ; H100 ~1,47–2,60 $/h ; egress 0 $ | Le moins cher du marché si on accepte l'aléa |
| 43 | **CoreWeave** | **Neocloud** entreprise : Kubernetes natif, InfiniBand, Platinum ClusterMAX | H100 ~6,16 $/h on-demand (~2,46 $ spot) ; contrats pluriannuels | Prod critique ; backlog 99 Md$ (T1 2026) ; clients Meta, OpenAI, Anthropic |
| 44 | **Nebius** | AI cloud (ex-Yandex N.V.) : prix publiés, B300/B200/H200/H100 | H100 3,85 $/h on-demand ; B300 publié (seul à le faire) ; ARR 3,0 Md$ | L'alternative européenne au trio US ; 2 Md$ investis par NVIDIA (mars 2026) |
| 45 | **Nscale** | Infra IA full-stack, campus gigawatt (UK/US) | Contrats (ex. 45 Md$ sur 6 ans avec Anthropic, août 2026 — à vérifier) | Capacité souveraine à très grande échelle |
| 46 | **Crusoe** | AI cloud + data centers (énergie) ; seul avec **AMD MI300X/MI355X** au catalogue | H100 ~3,90 $/h ; H200 ~4,29 $/h (le moins cher publié) | Diversifier hors NVIDIA ; 4,9 GW contractés |
| 47 | **FluidStack** | Clusters GPU frontier-scale ; partenaire infra d'Anthropic | H100 ~2,80 $/h ; 4090 ~0,40 $/h | Gros clusters réservés |
| 48 | **Genesis Cloud** | Cloud GPU **UE** (Islande), 100 % renouvelable, conforme GDPR | H100 ~2,19 $/h ; H200 ~2,80 $/h ; 4090 ~0,55 $/h | Souveraineté UE + écologie + prix |
| 49 | **GMI Cloud** | GPU cloud (US/Taïwan) : H100/H200/B200/GB200 | Tarifs publiés agressifs | Alternative Asie/US aux neoclouds |
| 50 | **Scaleway** | Cloud **français** (Iliad) : instances GPU H100, Inference dédiée | €/h (H100) ; endpoints managés | Souveraineté FR/UE, facturation en euros |

### 3.2. Mentions honorables (hors top 50 mais réels)

- **SambaNova** (US) : puces RDU propriétaires, SambaNova Cloud — inférence rapide sur gros modèles.
- **FriendliAI** (KR/US) : moteur d'inférence, déploiements dédiés.
- **Canopywave** : routage/agrégation low-cost (DeepSeek V3.2 à 0,182 $/1M in via leur routage, mars 2026).
- **Inception Labs** : lab des **dLLM** (modèles de diffusion, Mercury-2.5) — nouveau paradigme latence.
- **MiniMax** : API propre en plus des poids (M3, M2.7).
- **DekaLLM** : hébergeur repéré sur OpenRouter (ex. Gemma-4-26B à 0,06/0,33 $).
- **FastPivot / FerryAPI** : agrégateurs low-cost émergents (Qwen, DeepSeek, Kimi) — vérifier avant prod.
- **Voltage Park** : cloud H100/Hopper/Blackwell, hardware en propre.
- **Spheron / io.net / Akash** : DePIN — marketplaces décentralisées de GPU (les moins chères, les moins garanties ; ex. Akash ~2,10 $/h H100 en médiane d'enchères, avril 2026).
- **OVHcloud** (FR) : instances GPU (A100/H100 selon périodes) — vérifier la disponibilité.
- **Cerebrium / Beam.cloud / Inferless** : serverless GPU alternatifs à Modal (crédits gratuits, scale-to-zero).
- **fal.ai** : média génératif (FLUX, Wan) à prix cassé — utile pour les pipelines multimodaux.
- **Nex AGI** : lab émergent (modèles N2.5 gratuits sur OpenRouter fin sept. 2026, offre limitée).

### 3.3. Lire le tableau comme un acheteur

1. **Trois métiers, trois facturations** : les *labs* vendent du token propriétaire ; les *hébergeurs d'open-weight* (Together, Fireworks, DeepInfra…) vendent du token mutualisé ; les *clouds GPU* (RunPod, Lambda, Vast…) vendent de l'heure GPU. Le token est simple, l'heure GPU devient rentable au-delà d'un seuil d'utilisation (section 8).
2. **Le même modèle, dix prix** : GPT-OSS-120B va de ~0,09 $/1M in (DeepInfra/OpenRouter) à 0,35 $/1M (Cerebras) — l'écart paie la vitesse (190 vs 3 000 tok/s). D'où le routage.
3. **Le free tier n'est jamais gratuit** : quotas (RPM/TPD), contextes rognés (Cerebras), données réutilisables (Google), modèles « enterprise-only » du jour au lendemain (Groq a basculé des Llama en « contact sales » le 26/08/2026 — prévoir un fallback).
4. **Souveraineté** : si les données doivent rester en UE, le shortlist est court — Mistral La Plateforme, Genesis Cloud, Scaleway, Nebius (UE), Azure/AWS régions UE avec zéro-rétention contractuelle.

---

## 4. Les passerelles et le routage intelligent

### 4.1. Pourquoi une passerelle ?

Sans passerelle, chaque provider = un SDK, une clé, une facturation, un format d'erreur, une limite de débit différente. Avec une passerelle, ton code parle **un seul dialecte** (l'API OpenAI : `POST /v1/chat/completions`) et c'est la passerelle qui :

1. **Route** la requête vers le bon (modèle, provider) selon la stratégie choisie ;
2. **Bascule (fallback)** automatiquement si le provider principal échoue (429, 500, timeout, modèle délisté) ;
3. **Équilibre la charge** entre plusieurs clés/endpoints du même modèle ;
4. **Observe** : logs, coût par requête, latence par provider, alertes budget ;
5. **Contrôle** : budgets par équipe/clé, rate limits, allowlist de modèles.

Deux philosophies : **le routeur managé** (OpenRouter : tu ne gères rien, tu paies la commission) et **la gateway auto-hébergée** (LiteLLM : tu gères, tu contrôles tout, coût = ton serveur).

### 4.2. OpenRouter — le routeur managé

**Principe** : une clé API, un endpoint (`https://openrouter.ai/api/v1`), 400+ modèles adressés par `fournisseur/modèle` ou par nom logique. Exemple d'appel (SDK OpenAI, seule la `base_url` change) :

```python
from openai import OpenAI

client = OpenAI(
    base_url="https://openrouter.ai/api/v1",
    api_key="${OPENROUTER_API_KEY}",  # variable d'environnement, jamais en clair
)
resp = client.chat.completions.create(
    model="deepseek/deepseek-v4.1-flash",   # modèle + variante
    messages=[{"role": "user", "content": "Résume ce log en 3 lignes."}],
    extra_body={"provider": {"order": ["deepseek", "together", "fireworks"],
                             "allow_fallbacks": True}},  # routage par provider
)
print(resp.choices[0].message.content)
```

**Options de routage natives** (vérifiées sept. 2026) :

| Option | Effet | Cas d'usage |
|---|---|---|
| `provider.order` | Liste ordonnée de providers à essayer | Préférer le moins cher, tolérer les autres |
| `provider.allow_fallbacks` | Bascule auto si le 1er échoue | Robustesse sans code |
| `provider.require_parameters` | N'expose que les providers supportant les paramètres demandés (ex. tools) | Agents avec function calling |
| `provider.data_collection` | `deny` = refuser les providers qui logguent les prompts | Données sensibles |
| `models: ["a", "b"]` (route `/chat/completions` multi-modèles) | OpenRouter choisit le moins cher / le plus dispo | Batch non critique |
| Variantes `:free` | Routage vers les endpoints gratuits (20 RPM ; 50 req/jour, 1000 après 10 $ de crédits) | Dev, tests, CI |

