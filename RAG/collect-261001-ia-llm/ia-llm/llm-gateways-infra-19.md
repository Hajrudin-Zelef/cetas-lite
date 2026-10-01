---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-19
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Meta", "Moonshot", "OpenAI", "Poolside", "Z.ai", "xAI"]
dates: ["2026-09-01", "2026-09-02", "2026-09-03", "2026-09-05", "2026-09-06", "2026-09-09", "2026-09-10", "2026-09-12", "2026-09-14", "2026-09-16", "2026-09-18", "2026-09-22", "2026-09-23", "2026-09-25", "2026-09-27"]
keywords: ["astra", "claude", "deepseek", "fable 5", "gemini", "gemini 3.8", "gemini 4", "glm", "gpt-6", "kimi", "llama", "luna"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [2299, 2398]
sha256: 74e97c0c6ef2e09b7d32b8d80d6628a2ba89588e2dc779a00362c26dda7df373
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

| Nom | Statut au 27/09/2026 | Source de l'annonce |
|---|---|---|
| DeepSeek V4.1 Pro | **ANNONCÉ, non sorti** (09/09/2026, équipe DeepSeek) | api-docs.deepseek.com (news 10/09/2026), pricing page |
| Claude Sonnet 5.5 | **ANNONCÉ, « dans les prochaines semaines »** | Anthropic (via digest EveryDev.ai 19–25/09/2026) |
| Claude Haiku 5.5 | **ANNONCÉ, « dans les prochaines semaines »** | Anthropic (via digest EveryDev.ai 19–25/09/2026) |
| Meta — plus haut tier de raisonnement | **ANNONCÉ, retenu jusqu'aux tests de sécurité** | wowtale.net 05/09/2026 |
| OpenAI — annonces « niveau DevDay » | **Teaser officiel**, contenu à vérifier | Équipe OpenAI (digest 16/09/2026) |
| Sonnet 4.8 | Rien d'officialisé | — |
| Haiku 5 (sans .5) | Rien d'officialisé (seul Haiku 5.5 évoqué) | — |
| Gemini 4 | Rien d'officialisé — **ne pas ajouter** | — |
| Veo 4 | Rien d'officialisé | — |
| GPT-6 Terra | Rien d'officialisé — **ne pas ajouter** | — |
| Llama 4 Behemoth | Rien d'officialisé | — |
| Poolside Malibu | Rien d'officialisé (RUMEUR) | — |
| Qwen / Kimi / GLM / Doubao (générations suivantes) | Rien d'officialisé | — |

### 159.11. Comment suivre les annonces sans te faire piéger (méthode)

```bash
# 1. Les changelogs officiels d'abord (exemples à bookmarker)
#    - https://api-docs.deepseek.com/news/
#    - https://platform.openai.com/docs/changelog  (ou le blog OpenAI)
#    - https://docs.anthropic.com/en/release-notes
#    - https://cloud.google.com/vertex-ai/generative-ai/docs/release-notes (Gemini)
#    - https://docs.x.ai/docs/changelog
# 2. Les pages pricing : une annonce y apparaît AVANT les blogs tiers
#    (cf. DeepSeek : la continuité de V4 Pro s'est lue sur la page pricing,
#     pas dans un communiqué).
# 3. Règle des 3 sources : un nom de modèle n'entre dans ton routeur que si
#    (a) l'éditeur le nomme, (b) un endpoint ou un ID d'API existe,
#    (c) le prix est publié. Il en manque un ? C'est « à vérifier ».
# 4. Dans ton RAG : taggue chaque fiche modèle avec "vérifié le <date>" et
#    une date d'expiration de 30 jours. Au-delà, la fiche passe en
#    « à recontrôler » automatiquement (script section 128).
```

Script de veille minimal (à adapter) :

```python
# watch_models.py — alerte quand un changelog officiel bouge
import hashlib, urllib.request, json, os
WATCH = {
    "deepseek_news": "https://api-docs.deepseek.com/news/news260910",
    "anthropic_releases": "https://docs.anthropic.com/en/release-notes",
}
state_file = os.path.expanduser("~/.cache/model_watch.json")
state = json.load(open(state_file)) if os.path.exists(state_file) else {}
changed = []
for name, url in WATCH.items():
    try:
        data = urllib.request.urlopen(url, timeout=20).read()
        h = hashlib.sha256(data).hexdigest()
        if state.get(name) and state[name] != h:
            changed.append(name)
        state[name] = h
    except Exception as e:
        print(f"[{name}] erreur de lecture : {e} (à vérifier manuellement)")
json.dump(state, open(state_file, "w"))
print("CHANGÉ :", changed if changed else "rien — relire manuellement quand même")
```

### 159.12. Checklist anti-rumeur (à coller dans ton INFRA.md)

- [ ] Le nom apparaît-il dans une communication **de l'éditeur** (pas un blog tiers) ?
- [ ] Un **ID d'API ou un endpoint** existe-t-il (pas seulement un nom) ?
- [ ] Le **prix** est-il publié sur une page officielle ?
- [ ] La **date d'annonce** est-elle connue et notée ?
- [ ] Si une seule de ces cases est vide : le modèle est « à vérifier », pas « disponible ».
- [ ] Interdiction de coder un nom de modèle futur en dur dans LiteLLM/FreeLLMAPI/scripts.
- [ ] Toute fiche RAG sur un modèle futur porte « vérifié le … » + expiration 30 jours.

### 159.13. Ce que ça change pour ton RAG (Zelef)

1. **Les 30 questions d'éval (section 127)** doivent inclure 2–3 questions « piège temporel » du type : « DeepSeek V4.1 Pro est-il disponible ? » — la bonne réponse au 27/09/2026 est « annoncé le 09/09/2026, non sorti, aucun prix ni endpoint publiés ». Ça teste si ton RAG date ses connaissances.
2. **Le tag « vérifié le »** (section 159.11) est la seule parade contre les guides qui pourrissent : un prix de septembre cité en décembre sans re-vérification est un mensonge poli.
3. **Ne réindexe pas sur rumeur** : ajoute une fiche « à venir » seulement pour les ANNONCÉ, jamais pour les RUMEUR.

### 159.14. Chronologie officielle de septembre 2026 (le mois le plus dense de l'année)

Pour situer les annonces « à venir » dans leur contexte : septembre 2026 a vu quatre labs sortir un flagship en quatre jours, puis une seconde vague trois semaines plus tard.

| Date | Événement | Source de l'annonce |
|---|---|---|
| 01/09/2026 | Anthropic lance **Claude Fable 5.1** (grand public) et **Claude Mythos 5.1** (institutions vérifiées, cybersécurité/sciences de la vie) | Anthropic ; relayé par startupfortune.com et wowtale.net |
| 02/09/2026 | Meta sort **Muse Spark 1.3** | Meta ; relayé par startupfortune.com |
| 03/09/2026 | OpenAI sort **GPT-6 Astra** (« système le plus capable et le plus aligné » selon OpenAI) ; Google sort **Gemini 3.8 Flash** (+ variante cybersécurité via le programme Fairwind) | OpenAI / Google ; relayés par startupfortune.com, wowtale.net |
| 06/09/2026 | CNBC parle de « model fatigue » : les acheteurs n'arrivent plus à suivre le tableau des scores. Le PDG de Runpod (Zhen Lu) : le marché est devenu si bruyant que les entreprises doivent faire du bruit pour exister. Sam Altman : les labs « passent tous à des cadences plus rapides ». | CNBC via startupfortune.com |
| 09/09/2026 | Un membre de l'équipe DeepSeek (@tianyi) nomme **DeepSeek V4.1 Pro** comme successeur à venir (voir 159.2) | Équipe DeepSeek ; relayé par orcarouter.ai |
| 10/09/2026 | DeepSeek sort **V4.1 Flash** (nouvelle architecture Causal Encoder-Decoder, 552B MoE, entrée image native, 1M de contexte) | Changelog officiel DeepSeek (news260910) |
| 12/09/2026 | Dario Amodei (Anthropic) écrit que l'industrie de l'IA **doit ralentir** ; Sam Altman (OpenAI) et Elon Musk soutiennent cet appel | daylila.com 23/09/2026 |
| 14/09/2026 | Date prévue du reroutage `deepseek-v4-pro` → V4.1 Flash — **plan retiré** sous la pression des développeurs ; le service V4 Pro continue | Page pricing DeepSeek ; benchlm.ai |
| 18/09/2026 | Reuters : Anthropic « envisagerait » un nouveau modèle pour répondre à GPT-6 Astra (**RUMEUR**, voir 159.9). Le même jour, **quatre clients poursuivent les labs** au sujet du pacte de ralentissement | Reuters via daylila.com |
| 22/09/2026 | Anthropic sort **Claude Opus 5.5** ($4/$20/M, cache $0.20/M, +30 % de vitesse) ; OpenAI sort **GPT-6 Sol et GPT-6 Luna** (moitié prix de la série 5.6) | Anthropic / OpenAI ; relayés par EveryDev.ai, TechTarget |

**Pourquoi c'est dans ton guide infra :** un mois comme celui-là casse les hypothèses figées. Le 1er septembre, ton routeur avait un « meilleur modèle » ; le 22, il en a trois autres, avec des prix divisés par deux à qualité comparable. Sans éval continue (section 127) et sans versions épinglées (section 142), tu paies le prix fort pour le modèle d'hier.

### 159.15. Le « capability-tiered gating » : la vraie annonce structurelle de septembre 2026

Au-delà des noms de modèles, septembre 2026 a officialisé une doctrine d'accès qui change la façon de consommer les API (wowtale.net, 05/09/2026) :

