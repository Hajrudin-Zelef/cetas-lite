---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-17
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Hugging Face", "OpenRouter"]
dates: ["2026-07-24", "2026-09-09", "2026-09-10", "2026-09-11", "2026-09-14", "2026-09-27"]
keywords: ["agents", "benchmark", "deepseek", "dpo", "fine-tuning", "lora", "moe", "open-weight", "pricing", "rlhf", "sonnet 5", "training"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [2148, 2237]
sha256: e5a2b3043331865a43665f6f33f59fa15af8dc082ab1a7c009ca6e433efffadb
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

| Code | Signification | Réaction |
|---|---|---|
| 401 | Clé invalide/absente | Vérifier la clé, sa révocation |
| 402 | Crédits épuisés | Recharger (OpenRouter) ou changer de provider |
| 404 | Modèle inconnu | ID exact ? Modèle déprécié ? |
| 429 | Rate limit / quota | **Backoff exponentiel**, failover (c'est le régime normal du gratuit) |
| 500/502/503 | Provider en panne | Failover auto (sections 9, 102) |
| 529 | Surcharge (Anthropic) | Retry + fallback |

Le 429 n'est pas une erreur, c'est un **signal de pilotage** : il dit « change de provider ».

## 155. Backoff exponentiel : l'implémentation propre

```python
import random, time
def appel_avec_retry(fn, essais=5):
    for i in range(essais):
        try:
            return fn()
        except RateLimitError:          # 429 — adapter à ton client HTTP
            if i == essais - 1: raise
            delai = min(2 ** i + random.uniform(0, 1), 60)
            time.sleep(delai)           # 1s, 2s, 4s, 8s, 16s + jitter
        except ServerError:             # 5xx
            time.sleep(5)
```

Jitter obligatoire : sans lui, N clients re-tentent en même temps et re-saturent le provider (thundering herd).

## 156. Ce que ce guide ne couvre pas (périmètre honnête)

- L'entraînement complet (pre-training) — hors de portée d'un labo perso.
- Le fine-tuning avancé (DPO, RLHF) — aperçu LoRA seulement (section 145).
- La conformité réglementaire (AI Act UE, etc.) — à traiter avec ton DPO/juriste avant usage pro des modèles.
- Les détails d'optimisation CUDA/Triton — niveau ingénierie ML.
- Les agents multimodaux complexes — DeepSeek Harness en preview (section 49).

## 157. Vérification finale avant mise en prod (checklist)

- [ ] Éval 30+ questions : score ≥ baseline, 0 hallucination critique.
- [ ] Fallback testé : couper le provider principal, vérifier la bascule < 5 s.
- [ ] 429 simulé : le retry/backoff fonctionne, pas de boucle infinie.
- [ ] Budgets : limites par clé, alertes de coût, dashboard hebdo.
- [ ] Secrets : aucune clé dans Git, rotation documentée.
- [ ] Données : classification (sensible → local/ZDR uniquement).
- [ ] Santé : `/health` de chaque service + alerte down.
- [ ] Backup : configs + script de réindexation testés (restauration essayée une fois).
- [ ] Coût projeté : calculé sur 3 mois avec les logs réels (pas au doigt mouillé).
- [ ] Doc : `INFRA.md` à jour (endpoints, quotas, dates de vérification).

## 158. Le mot de la fin : ta doctrine

1. **Mesure d'abord** (tokens, coûts, qualité) — tout le reste en découle.
2. **Gratuit + local** couvrent 95 % d'un usage perso/étude — le payant est un scalpel, pas un marteau.
3. **Un seul endpoint** dans ton code (LiteLLM, routeur maison, ou FreeLLMAPI) — jamais de provider en dur.
4. **Éval continue** : les modèles changent tous les mois, tes 30 questions sont ton ancre.
5. **Données** : le vrai coût du gratuit — traite tes prompts comme des données de prod.

*Fin de la partie principale du guide — 158 sections. La section 159 « À venir » suit ci-dessous ; son contenu est daté et doit être recontrôlé à chaque changement de mois.*
## 159. À venir — annonces vérifiées au 27/09/2026

Cette section est la seule du guide qui parle du futur, et elle le fait avec des règles strictes : **uniquement des sorties officiellement annoncées**, avec la source et la date d'annonce. Tout le reste est marqué « rien d'officialisé » ou « RUMEUR non confirmée ». Relis cette section à chaque début de mois : les annonces ci-dessous seront déjà périmées pour la plupart.

### 159.1. Règles de lecture (à appliquer avant de citer quoi que ce soit)

1. **ANNONCÉ** = le nom du modèle/produit apparaît dans une communication officielle de l'éditeur (changelog, pricing page, communiqué, thread de lancement officiel) ou dans les propos d'un membre identifié de l'équipe technique cités par plusieurs sources.
2. **RUMEUR** = information rapportée par la presse ou des fuites, non reprise par l'éditeur. Marquée explicitement « RUMEUR non confirmée ».
3. **RIEN D'OFFICIALISÉ** = aucune annonce trouvée au 27/09/2026 après recherche web. Ce n'est pas « ça n'existe pas », c'est « l'éditeur ne l'a pas annoncé ».
4. Un prix, un benchmark ou une date qui ne viennent pas d'une source vérifiable restent « à vérifier ».
5. Ne cite jamais un nom de modèle futur dans une config de prod (LiteLLM, FreeLLMAPI, scripts) : un nom annoncé n'est pas un endpoint qui répond.

### 159.2. DeepSeek V4.1 Pro — ANNONCÉ, non sorti (vérifié le 27/09/2026)

Le cas d'école de l'annonce qui existe mais du produit qui n'existe pas encore.

| Point | État au 27/09/2026 |
|---|---|
| Annonce | **Le 09/09/2026**, un membre de l'équipe technique DeepSeek (pseudonyme `@tianyi`) a expliqué que, une fois DeepSeek V4.1 Flash officiel, les requêtes adressées à `deepseek-v4-pro` seraient reroutées vers V4.1 Flash aux tarifs Flash « jusqu'au lancement de DeepSeek V4.1 Pro ». Le nom « V4.1 Pro » a donc été employé **une seule fois** côté vendeur. |
| Sortie | **Non sorti.** Le plan de reroutage du 14/09/2026 a été **retiré** sous la pression des développeurs : DeepSeek a ensuite indiqué sur sa page pricing que le service V4 Pro continue. |
| Model card / poids open-weight | **Aucun.** Le dépôt le plus récent de l'organisation `deepseek-ai` sur Hugging Face au 27/09/2026 est `DeepSeek-V4.1-Flash` (créé le 10/09/2026). Rien nommé « V4.1 Pro » n'existe, ni sur le Hub ni ailleurs. |
| Prix | **Aucun publié.** Pour l'échelle : V4.1 Flash = $0.15/$0.60 par million de tokens hors pic ($0.30/$1.20 en pic), V4 Pro = $0.66/$1.98 hors pic ($1.32/$3.96 en pic) — prix vérifiés le 27/09/2026 via les pages pricing DeepSeek relayées par des tiers. Un successeur serait logiquement attendu près de la colonne Flash, c'est précisément pour ça que le plan de retrait existait. |
| Endpoints API | **Aucun.** L'API DeepSeek expose exactement deux noms : `deepseek-flash` (V4.1 Flash) et `deepseek-v4-pro` (V4 Pro). Aucun endpoint V4.1 Pro, aucune limite de débit publiée pour un tel modèle. |
| Ce qui est sorti à la place | **DeepSeek V4.1 Flash**, le 10/09/2026 : architecture Causal Encoder-Decoder (552B MoE, 8B actifs en entrée / 16B en sortie), entrée image native, contexte 1M tokens, jusqu'à ~384K tokens de sortie, identifiant API `deepseek-flash`. Les anciens IDs `deepseek-v4-flash` et `deepseek-v4-flash-vision-exp` sont retirés mais toujours acceptés et routés vers V4.1 Flash. Les alias historiques `deepseek-chat` et `deepseek-reasoner` sont arrivés à leur date de retrait le 24/07/2026. |

Sources (vérifiées le 27/09/2026) : changelog officiel DeepSeek (`api-docs.deepseek.com`, note du 10/09/2026), page pricing DeepSeek (continuité du service V4 Pro), synthèse d'orcarouter.ai (« DeepSeek V4.1 Pro Leak: A Window That Closes Sept 30 »), note de recherche codex-router du 11/09/2026, benchlm.ai (historique des versions).

**Leçon infra pour Zelef :** c'est exactement pour ce genre de situation que la section 142 (versions épinglées) existe. Le 10/09/2026, `deepseek-v4-pro` a failli changer de modèle sous-jacent du jour au lendemain, puis ne l'a pas fait. Si ton routeur maison utilise des alias vendeur au lieu d'IDs datés, tu subis ce genre de décision. Épingle, loggue le modèle réel retourné, et mets une alerte sur les changements de changelog.

### 159.3. Anthropic — Sonnet 5.5 et Haiku 5.5 ANNONCÉS (sortie « dans les prochaines semaines »)

