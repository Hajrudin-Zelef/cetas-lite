---
id: collect-261001-automatisation-infra/automatisation-infra/mcp-lsp-protocoles-13
title: "MCP, LSP, code-server, websearch — Guide pratique"
domain: automatisation-infra
role: reference
task: reference
actors: ["Google", "Perplexity"]
dates: ["2025-08-11", "2027-01-01"]
keywords: ["mcp", "agent", "agents", "arr", "open source", "perplexity"]
source: docs/RAG/collect-261001-automatisation-infra/mcp_lsp_protocoles.md
source_anchor: ""
source_lines: [2073, 2228]
sha256: 6efe4876444232cde7e8dbbe484398f22fa667c4e37c17d0c7c588d88e1dcc7e
---

# MCP, LSP, code-server, websearch — Guide pratique

- **Bing Search API** : **retirée le 11/08/2025**. Les tutos qui l'utilisent
  sont morts.
- **Google Programmable Search (CSE)** : fermée aux nouveaux clients,
  **arrêt total le 01/01/2027**. Ne construis rien dessus.
- **DuckDuckGo « API »** : pas de vraie API — que des réponses instantanées
  ou du scraping non officiel. À rejeter pour un agent sérieux.

⚠️ Si un tuto 2024 te propose Bing ou CSE : change de tuto.

## 106. SearxNG auto-hébergé : l'option sysadmin

**SearxNG** = métamoteur open source à héberger toi-même (agrège Google, Bing,
Brave, etc.) :

```bash
docker run -d --name searxng --restart unless-stopped \
  -p 127.0.0.1:8080:8080 \
  -v ./searxng:/etc/searxng \
  searxng/searxng:latest
# API : GET http://127.0.0.1:8080/search?q=...&format=json
```

- **Coût** : 0 € hors serveur. **Limites** : pas de SLA, les gros moteurs
  rate-limitent les IP de datacenters, maintenance à ta charge, qualité
  variable.
- 💡 Pour Zelef : parfait en **labo/dev** et en **fallback** si l'API
  payante tombe. Pas pour de la prod critique.

## 107. Tableau comparatif complet (sept 2026, 📌 prix à vérifier)

| API | 1k req (ordre) | Free tier | Index propre | Contenu LLM-ready | MCP officiel | Verdict |
|---|---|---|---|---|---|---|
| Tavily | 8-16 $ | 1k crédits/mois | Non | ✅ excellent | ✅ | Agent exigeant |
| Exa | 7-15 $ | 1k/mois + 10 $ | ✅ | ✅ | ✅ (150/j) | RAG sémantique |
| Brave | 5 $ | ~1k/mois (5 $) | ✅ | Partiel | ✅ | **Meilleur rapport Q/P** |
| Serper | 0,3-1 $ | 2,5k one-shot | Non | Non | Non | Volume pas cher |
| SerpAPI | 25 $+ | 250/mois | Non | Non | ✅ | Multi-moteurs |
| Perplexity | 5 $ + tokens | Limité | ✅ | ✅ (réponse) | Non | Synthèse directe |
| Firecrawl | ~16 $+/mois | 1k crédits/mois | Non | ✅ | ✅ | Search+scrape |
| SearxNG | 0 $ | ∞ | Non | Non | Non | Labo/fallback |

## 108. Exemple d'intégration : outil `web_search` dans un agent Python

```python
"""websearch_tool.py — outil de recherche web pour agent, provider Tavily.
Variable d'env : TAVILY_API_KEY. Cache + budget intégrés (§109).
"""
import os, time, hashlib, json
import urllib.request

API_KEY = os.environ["TAVILY_API_KEY"]
_CACHE: dict[str, tuple[float, dict]] = {}
CACHE_TTL = 3600          # 1 h
BUDGET_PAR_JOUR = 200     # garde-fou coût
_compteur = {"n": 0, "date": time.strftime("%Y-%m-%d")}


def _post(payload: dict) -> dict:
    req = urllib.request.Request(
        "https://api.tavily.com/search",
        data=json.dumps(payload).encode(),
        headers={"Content-Type": "application/json",
                 "Authorization": f"Bearer {API_KEY}"},
    )
    with urllib.request.urlopen(req, timeout=30) as r:
        return json.load(r)


def web_search(query: str, max_results: int = 5) -> list[dict]:
    """Recherche web pour agent. Retourne [{titre, url, contenu, score}]."""
    # --- garde-fou budget ---
    today = time.strftime("%Y-%m-%d")
    if _compteur["date"] != today:
        _compteur.update({"n": 0, "date": today})
    if _compteur["n"] >= BUDGET_PAR_JOUR:
        raise RuntimeError("Budget web_search quotidien atteint")

    # --- cache ---
    key = hashlib.sha256(query.encode()).hexdigest()
    if key in _CACHE and time.time() - _CACHE[key][0] < CACHE_TTL:
        return _CACHE[key][1]["results"]

    data = _post({"query": query, "search_depth": "advanced",
                  "max_results": max_results, "include_answer": False})
    _compteur["n"] += 1
    results = [{"titre": r.get("title"), "url": r.get("url"),
                "contenu": r.get("content", "")[:2000],
                "score": r.get("score")}
               for r in data.get("results", [])]
    _CACHE[key] = (time.time(), {"results": results})
    return results


if __name__ == "__main__":
    for r in web_search("MikroTik Netinstall error PXE boot"):
        print(f"- {r['titre']}\n  {r['url']}\n")
```

Adaptation Brave (GET + `X-Subscription-Token`) ou Exa (POST + `x-api-key`) :
même squelette, seuls l'URL, les headers et le parsing changent.

## 109. Garde-fous coût : cache, budget, déduplication

1. **Cache** (TTL 1-24 h) : les agents reposent souvent les mêmes questions.
2. **Budget quotidien** avec exception explicite (pas de fail silencieux).
3. **Déduplication** : normalise la query (minuscules, trim) avant hash.
4. **Profondeur adaptative** : `basic` d'abord, `advanced` si insuffisant.
5. **`max_results` petit** (3-5) : chaque résultat = tokens.
6. **Alerting** : loggue chaque appel (coût estimé) vers ton Loki ;
   alerte si ×3 vs la moyenne.
7. **Clé par environnement** : clé dev plafonnée, clé prod supervisée.

💡 Règle : le coût websearch doit rester **< 10 %** de ta facture LLM,
sinon c'est ton découpage (trop d'appels, pas de cache) qui est mauvais.

## 110. Critères de choix : arbre de décision

```
Besoin d'une API de recherche ?
├─ Usage < 1k req/mois, dev/lab → Brave (free tier) ou SearxNG local
├─ Agent qui répond avec sources fiables → Tavily (advanced)
├─ RAG / « trouve des pages sur X » → Exa (contenu riche)
├─ Gros volume, requêtes simples, petit budget → Serper (⚠️ ToS)
├─ Synthèse directe avec réponse → Perplexity Sonar
└─ Search + extraction en un appel → Firecrawl
Puis : vérifie rate limits du plan, SDK Python, MCP officiel.
```

💡 Pour Zelef : **Brave** en quotidien (qualité/prix), **Tavily** quand
l'agent doit produire une réponse sourcée propre (docs, procédures).
Les deux ont un MCP officiel : zéro code à écrire.

## 111. Pièges websearch : coût, snippets, rate limits, fraîcheur

1. **Le coût caché des retries** : un snippet pourri → l'agent relance →
   ×3 appels. Mieux vaut 1 appel « advanced » que 3 « basic ».
2. **Snippets vs contenu** : un snippet de 200 caractères ne suffit jamais
   pour une procédure technique → prévois l'**extraction** (`/extract`,
   Firecrawl, Jina).
3. **Rate limits** : ~1-50 req/s selon l'offre. Un agent parallèle qui
   lance 20 recherches d'un coup se fait 429 → **file d'appels + backoff**.
4. **Fraîcheur** : précise `time_range` (jour/semaine/mois) pour l'actu
   (CVE, sorties). Sans ça, tu récupères du 2023.
5. **SEO-spam** : les premiers résultats ne sont pas les meilleurs —
   demande à l'agent de **croiser 2+ sources** avant d'affirmer.
6. **Boucles de recherche** : « cherche, pas trouvé, reformule, re-cherche »
   ×10 → plafonne les itérations (ex : 5) puis rends la main à l'humain.
7. **Données personnelles** : une recherche web peut remonter des données
   à caractère personnel → ne les stocke pas dans ton RAG sans base légale.

---

# PARTIE VIII — TRANSVERSE

## 112. 18 pièges transverses (MCP + LSP + code-server + websearch)

