---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-2
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Cerebras", "DeepSeek", "Google", "Groq", "Huawei", "OpenAI", "OpenRouter", "xAI"]
dates: []
keywords: ["gpu", "claude", "deepseek", "gemini", "glm", "gpt-6", "grok", "llama", "mai", "memory", "nvidia", "open-weight"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [103, 240]
sha256: 130617c2e6bd02ffcd3dd6a57e116685922f71a675e822d1227fb76b65500d07
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

- Prépaiement par carte : tu charges des crédits ($5 min constaté — à vérifier), la conso décrémente.
- **Zéro Data Retention (ZDR)** disponible au niveau organisation : aucun entraînement sur tes données, utile en contexte pro.
- **Garde-fous budgétaires** : limites de dépense par clé (quotidienne/hebdo/mensuelle, reset auto) — la limite journalière sert de coupe-circuit anti-dérapage.
- Chaque réponse API inclut la génération et le coût ; le dashboard détaille par modèle.
- Marge OpenRouter : petite surcouche vs prix direct fournisseur — à volume élevé, comparer avec l'API directe.
- BYOK : tu peux apporter tes propres clés fournisseurs via OpenRouter (1M requêtes/mois gratuites puis +5 %) — utile pour centraliser sans changer de facturation.

## 9. OpenRouter : routage et providers

Par défaut, OpenRouter route vers le fournisseur le plus pertinent (vitesse/prix/disponibilité). Contrôle fin via le champ `provider` :

```json
{
  "model": "meta-llama/llama-3.1-70b-instruct",
  "messages": [{"role": "user", "content": "Bonjour"}],
  "provider": {
    "order": ["Groq", "Together", "DeepInfra"],
    "allow_fallbacks": true,
    "require_parameters": true,
    "data_collection": "deny"
  }
}
```

- `order` : fournisseurs préférés dans l'ordre.
- `allow_fallbacks: true` : si le premier tombe, bascule auto sur le suivant — **le cœur du pattern fallback**.
- `require_parameters: true` : n'utilise que des fournisseurs supportant les paramètres demandés (JSON mode, tools…).
- `data_collection: "deny"` : exclut les fournisseurs qui entraînent sur les données.
- `GET /api/v1/models/<id>/endpoints` : liste les fournisseurs réels derrière un modèle, avec prix et latence — indispensable avant de figer un choix.

## 10. OpenRouter : exemple d'appel (curl)

```bash
curl https://openrouter.ai/api/v1/chat/completions \
  -H "Authorization: Bearer sk-or-v1-FAKEFAKEFAKEFAKE" \
  -H "Content-Type: application/json" \
  -H "HTTP-Referer: https://mon-lab.local" \
  -H "X-Title: test-rag" \
  -d '{
    "model": "z-ai/glm-5.2:free",
    "messages": [{"role": "user", "content": "Explique le routage OSPF en 3 phrases."}],
    "provider": {"allow_fallbacks": true}
  }'
```

`HTTP-Referer` / `X-Title` : optionnels, servent au classement public des apps. En Python, le client `openai` suffit : `base_url="https://openrouter.ai/api/v1"`.

## 11. OpenRouter : exemple Python (comparer 3 modèles)

```python
from openai import OpenAI

client = OpenAI(
    base_url="https://openrouter.ai/api/v1",
    api_key="sk-or-v1-FAKEFAKEFAKEFAKE",
)

MODELES = [
    "z-ai/glm-5.2:free",
    "nvidia/nemotron-3.5-lightning:free",
    "deepseek/deepseek-v4.1-flash",
]
question = "Donne la commande VRP Huawei pour créer un VLAN 10."

for m in MODELES:
    r = client.chat.completions.create(
        model=m,
        messages=[{"role": "user", "content": question}],
        provider={"allow_fallbacks": True},
    )
    print("=" * 60)
    print(m, "->", r.choices[0].message.content[:300])
```

Même code, trois moteurs : c'est exactement le cas d'usage « comparer des modèles » pour ton RAG (qualité des réponses générées à partir de tes chunks).

## 12. OpenRouter : cas d'usage recommandés

1. **Banc d'essai de modèles** : tester 10 modèles sur 50 questions de ton domaine (réseau, onduleurs) et mesurer qualité/coût avant de figer.
2. **Fallback production** : modèle principal + 2 secours via `allow_fallbacks` — ton app survit à une panne fournisseur.
3. **Routage par complexité** : petit modèle gratuit pour le triage/classification, gros modèle payant pour la synthèse finale.
4. **Veille modèles** : l'API `/models` te dit qui sort quoi et à quel prix, sans créer 15 comptes.

## 13. OpenRouter : limites et pièges

1. **Latence ajoutée** : un saut réseau de plus vs l'API directe (quelques dizaines de ms) — négligeable sauf voix temps réel.
2. **Prix variables** : le même `model` peut coûter différemment selon le provider routé et l'heure (cf. DeepSeek peak). Fixe `provider.order` si tu veux de la prédictibilité.
3. **Disponibilité des `:free`** : les modèles gratuits sont les premiers délestés en surcharge ; prévois un fallback payant pas cher.
4. **Contexte effectif** : le contexte annoncé est celui du modèle, mais certains providers le tronquent — vérifier via `/info`.
5. **Données** : par défaut, certains providers peuvent logger ; `data_collection: "deny"` + ZDR pour du contenu sensible.
6. **50 req/jour** en gratuit : c'est un quota d'essai, pas une infra — pour du volume gratuit, empile avec Groq/Cerebras/Gemini (partie E).

## 14. Checklist OpenRouter (mise en route en 10 min)

- [ ] Créer un compte, générer une clé (`sk-or-v1-...`).
- [ ] Charger $10 de crédits (débloque 1 000 req/jour sur les `:free`).
- [ ] Tester `GET /api/v1/models` et filtrer `pricing.prompt == "0"`.
- [ ] Lancer le script de comparaison (section 11) sur 5 questions de ton domaine.
- [ ] Fixer `provider.order` + `allow_fallbacks` pour ton modèle de prod.
- [ ] Poser une limite de dépense quotidienne par clé (garde-fou).
- [ ] Activer ZDR si données pro.

---

# PARTIE B — GROQ

## 15. Groq : présentation (ne pas confondre avec Grok de xAI)

Groq (groq.com, fondée 2016) ne fabrique pas de modèles : elle conçoit des **puces d'inférence**, les **LPU (Language Processing Unit)**, et vend l'inférence dessus via GroqCloud. Positionnement : **la vitesse**. Là où un GPU fait ~30–80 tok/s sur un 70B, les LPU annoncent **275 à 1 000 tok/s** selon le modèle.
API 100 % compatible OpenAI : `https://api.groq.com/openai/v1`. Inscription sans carte bancaire, **free tier permanent** (rate-limité, jamais facturé : au-delà, erreur 429, pas de facture surprise).

## 16. Groq : pourquoi c'est rapide (LPU)

- Architecture **déterministe** : le compilateur connaît le graphe du modèle à l'avance, pas d'ordonnanceur dynamique comme sur GPU → latence ultra-faible et **débit prévisible**.
- Mémoire SRAM énorme sur la puce, bande passante interne très élevée → le décodage auto-régressif (memory-bound sur GPU) s'envole.
- Conséquence pratique : **time-to-first-token** de l'ordre de la centaine de ms et génération à 500–1 000 tok/s sur les petits modèles.
- Revers : catalogue limité aux modèles **open-weight** compilés pour LPU (pas de GPT-6 ni de Claude ici) ; fenêtres de contexte parfois réduites vs l'original.

## 17. Groq : modèles servis vérifiés (sept 2026)

État relevé fin août 2026 (le catalogue tourne vite — vérifier sur `console.groq.com/docs/models`) :

| Model ID exact | Contexte | Complétion max | Vitesse constatée | Prix payant /1M (vérifié mai–août 2026) |
|---|---|---|---|---|
| `openai/gpt-oss-120b` | 131 072 | 65 536 | ~500 tok/s | $0.15 / $0.60 |
| `openai/gpt-oss-20b` | 131 072 | 65 536 | ~1 000 tok/s | $0.075 / $0.30 |
| `qwen/qwen3.6-27b` | 131 072 | 16 384 | à vérifier | $0.60 / $3.00 |
| `qwen/qwen3.8-27b` | 131 042 | 16 384 | à vérifier | $0.80 / $4.00 (preview) |
| `openai/gpt-oss-safeguard-20b` | 131 072 | 65 536 | à vérifier | Modération de contenu |
| `groq/compound` | 131 072 | 8 192 | système agentique | Système, pas un simple chat |
| `groq/compound-mini` | 131 072 | 8 192 | système agentique | Version légère |
| `whisper-large-v3` | audio | — | ~200× temps réel | $0.111 / heure transcrite |
| `whisper-large-v3-turbo` | audio | — | ~220× temps réel | $0.04 / heure transcrite |
| `canopylabs/orpheus-v1-english` | 4 000 | 50 000 (audio) | TTS | preview |

**Dépréciés** (ne plus utiliser dans du neuf) : `llama-3.3-70b-versatile`, `llama-3.1-8b-instant`, Mixtral 8x7B (déc. 2025). Remplacements conseillés par Groq : `openai/gpt-oss-120b` et `qwen/qwen3.6-27b`.
Réductions cumulables (vérifié 2026) : **input caché −50 %**, **batch −50 %**, **tier développeur −25 %**.

