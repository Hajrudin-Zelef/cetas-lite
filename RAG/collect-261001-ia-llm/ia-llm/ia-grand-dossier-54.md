---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-54
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Hugging Face", "Meta", "OpenAI", "vLLM"]
dates: []
keywords: ["agent", "agentic", "agents", "awq", "benchmarks", "chatgpt", "claude", "embeddings", "mcp", "open source", "qwen", "vllm"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [4309, 4432]
sha256: 6cea186df29cf51764f950bb26380da63f84593d7bcdece654b03756ba8b8299
---

# IA — Le grand dossier

messages = [{"role": "system", "content": "Tu es un technicien réseau. Tu ne peux agir que via run_diag."},
            {"role": "user", "content": "La connexion internet semble coupée, diagnostique."}]
for _ in range(6):  # borne anti-boucle infinie
    r = llm.chat.completions.create(model="Qwen/Qwen3-8B-AWQ", messages=messages,
                                    tools=TOOLS, tool_choice="auto", temperature=0.1)
    m = r.choices[0].message
    messages.append(m)
    if not m.tool_calls: break
    for tc in m.tool_calls:
        args = json.loads(tc.function.arguments)
        # GARDE-FOU : validation humaine avant toute action non lecture-seule
        print(f"[AGENT veut exécuter] {tc.function.name}{args} — (auto-approuvé: lecture seule)")
        out = run_diag(args["cmd"])
        messages.append({"role": "tool", "tool_call_id": tc.id,
                         "content": out[:2000]})
print(messages[-1].content)
```

**Exemple réel — serveur MCP minimal (expose tes runbooks à Claude Desktop / tout client MCP) :**

```python
# pip install "mcp[cli]"
from mcp.server.fastmcp import FastMCP
mcp = FastMCP("runbooks-zelef")

@mcp.tool()
def get_runbook(equipement: str) -> str:
    """Retourne le runbook d'un équipement (kyocera, onduleur, proxmox...)."""
    try:
        return open(f"/srv/runbooks/{equipement}.md", encoding="utf-8").read()[:4000]
    except FileNotFoundError:
        return "Runbook inconnu."

@mcp.tool()
def ups_status() -> str:
    """Lit le statut de l'onduleur via NUT (lecture seule)."""
    import subprocess
    return subprocess.run(["upsc", "ups1"], capture_output=True, text=True, timeout=10).stdout

if __name__ == "__main__":
    mcp.run()  # stdio : branché en 2 lignes dans claude_desktop_config.json
```

**Les 4 règles d'or des agents en prod** (synthèse OWASP Agentic 2026 + incidents 2.1.1) :

1. **Whitelist d'outils, jamais de shell libre** — l'agent n'exécute que ce que tu as codé.
2. **Humain dans la boucle** pour tout ce qui est irréversible (cf. exemple : lecture seule auto, le reste = approbation).
3. **Bornes** : N itérations max, timeout, budget tokens, périmètre réseau.
4. **Audit** : chaque action loggée avec qui/quoi/quand — c'est aussi une exigence réglementaire qui vient (traçabilité, AI Act art. 12 pour le haut risque).

### 4.5. Mettre les stacks ensemble : ton architecture cible

```
                    ┌─────────────────────────────────────────┐
                    │           TON APP RAG (perso)             │
                    │  Guides .md (25k+ lignes) + PDF + liens  │
                    └───────────────┬─────────────────────────┘
                                    │
        ┌───────────────────────────┼───────────────────────────┐
        ▼                           ▼                           ▼
┌───────────────┐           ┌───────────────┐           ┌───────────────┐
│  Ingestion    │           │  Retrieval    │           │  Génération   │
│  scripts      │           │  pgvector     │           │  vLLM local   │
│  Python       │           │  (embeddings  │           │  ou API cloud │
│  (chunking    │           │  + BM25 FR    │           │  (Claude /    │
│  sémantique)  │           │  + RRF +      │           │  GPT / local) │
└───────────────┘           │  rerank)      │           └───────────────┘
                            └───────────────┘
        ┌───────────────────────────┼───────────────────────────┐
        ▼                           ▼                           ▼
┌───────────────┐           ┌───────────────┐           ┌───────────────┐
│  Éval continue│           │  Sécurité     │           │  Agents (v2)  │
│  50 questions │           │  ACL/doc,     │           │  MCP runbooks │
│  de référence │           │  logs,        │           │  + NUT + GLPI │
│  + alertes    │           │  pas de       │           │  (lecture     │
│  dérive       │           │  secrets      │           │  seule d'abord)│
└───────────────┘           └───────────────┘           └───────────────┘
```

Commence par la colonne du milieu (retrieval de qualité), pas par le modèle le plus brillant :
dans un RAG, **80 % de la qualité vient du retrieval, 20 % du générateur**.

---

## 5. Les communautés (« slack » = là où ça se passe)

> L'IA bouge par semaines. Aucun dossier, même de 6 000 lignes, ne remplace un bon fil d'actu.
> Voici la carte des lieux où l'info circule *vraiment*, avec pour chacun : à quoi ça sert,
> comment l'utiliser en sysadmin, et le niveau de bruit.

### 5.1. Carte générale

| Lieu | Type | Pour quoi | Bruit | Accès |
|---|---|---|---|---|
| **r/LocalLLaMA** | Reddit | Modèles open source, quants, hardware, benchmarks — LE forum des auto-hébergeurs | Moyen-élevé | Public |
| **r/MachineLearning** | Reddit | Recherche, papiers, discussions techniques (modération stricte) | Faible | Public |
| **r/artificial** | Reddit | Actu générale IA | Élevé | Public |
| **r/ClaudeAI**, **r/ChatGPT**, **r/Anthropic** | Reddit | Usages, astuces, bugs par produit | Moyen | Public |
| **Hacker News** (news.ycombinator.com) | Forum | Actu tech + commentaires d'experts ; fils IA quotidiens | Faible-moyen | Public |
| **Hugging Face** (huggingface.co + forums) | Plateforme + forum | Modèles, datasets, Spaces ; discussions par modèle | Faible | Public |
| **GitHub** | Code | Repos des frameworks (vLLM, LangChain, Ollama…) : issues = documentation vivante | Faible | Public |
| **Discord Anthropic** | Discord | Communauté Claude, annonces, support | Moyen | Public |
| **Discord r/LocalLLaMA / Ollama / LangChain** | Discord | Entraide temps réel, hardware, déploiements | Moyen | Public |
| **Slack** | Slack | Slacks d'entreprise/communautés pro (ex. : MLOps Community) — plus pro, moins hype | Faible | Sur invitation |
| **X / Twitter** | Réseau | Annonces en avant-première des chercheurs et labs | Très élevé | Public |
| **LinkedIn** | Réseau | Retours d'expérience entreprise, réglementation | Moyen | Public |

### 5.2. Où suivre l'actu (concret)

**Newsletters (noms réels, actives en 2026) :**

| Newsletter | Auteur / org | Fréquence | Angle |
|---|---|---|---|
| **The Rundown AI** | Rowan Cheung | Quotidienne | Actu condensée, grand public pro |
| **TLDR AI** | TLDR | Quotidienne | Actu + papiers + outils, format dense |
| **The Neuron** | — | Quotidienne | Actu + prompts, ton accessible |
| **Import AI** | **Jack Clark** (cofondateur d'Anthropic !) | Hebdo | Analyse technique et politique, la plus respectée du milieu |
| **The Batch** | DeepLearning.AI (Andrew Ng) | Hebdo | Pédagogique, rigoureux |
| **AI Breakfast** | — | Quotidienne | Brief matinal |
| **Last Week in AI** | — | Hebdo | Revue de la semaine |
| **The Decoder** | — | Quotidienne | Actu Europe/international |

**Comptes et voix à suivre (vérifiés comme réels et actifs) :**

