---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-13
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "OpenAI"]
dates: []
keywords: ["agent", "agents", "claude", "guardrails", "mcp", "memory"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [1911, 2095]
sha256: 18b69ba7958b0b25b61eabf84804de586561aba6502c44fe72d9674da4329701
---

# Concepts : agents IA, agentic, autonomie

Quand passer à un framework (partie H) : quand tu as **la preuve** que l'une de
ces limites te bloque — pas avant. Un framework ne rend pas un mauvais design bon ;
il rend un bon design plus scalable.

## 92. Résumé partie G : ce que tu sais faire maintenant

- [ ] écrire une boucle ReAct complète avec function calling (API OpenAI-compatible) ;
- [ ] écrire un outil shell **sécurisé** (liste blanche, pas de shell=True, timeouts) ;
- [ ] confiner les outils fichiers à un périmètre ;
- [ ] imposer des budgets durs (étapes/tokens/temps) ;
- [ ] demander une approbation humaine fail-closed ;
- [ ] logger chaque exécution en JSONL ;
- [ ] tester un agent (scénarios nominaux + attaques) ;
- [ ] brancher ton RAG comme mémoire de l'agent.

Si tu coches tout : tu comprends mieux les agents que 90 % des gens qui en parlent.
Le framework, maintenant, ne sera qu'une formalisation de ce que tu as déjà construit.

---

# PARTIE H — Version framework 2026 : LangGraph

## 93. Pourquoi LangGraph (et pas un autre) en 2026

État du marché relevé fin septembre 2026 (comparatifs multiples convergents) :

| Framework | Philosophie | Quand le choisir |
|---|---|---|
| **LangGraph** | graphes d'états explicites, checkpoints natifs | **prod** : contrôle, audit, reprise, human-in-the-loop |
| CrewAI | équipes multi-agents par rôles | proto rapide, démo, équipes « métier » |
| AutoGen (v0.4+) | conversation async entre agents | charges async/event-driven |
| Google ADK | agents natifs Vertex AI | si tu es full Google Cloud |
| OpenAI Agents SDK | handoffs + guardrails simples | chemin court vers la prod OpenAI |
| Claude Agent SDK | MCP natif, computer use | écosystème Anthropic |
| LlamaIndex | agents centrés données/RAG | ton cas RAG, si tu veux du clé en main |
| Pydantic AI | agents typés Python | si tu veux du typage strict |

Pour ton profil (sysadmin, prod, auditabilité), **LangGraph est le choix par défaut** :
c'est le seul où le graphe d'exécution est explicite, checkpointé et interruptible —
exactement ce qu'exige la doctrine de la partie C. Les autres sont de bons choix
dans leurs niches ; aucun ne remplace la compréhension de la partie G.

## 94. Installation et concepts LangGraph

```bash
pip install langgraph langchain-openai langchain-core
# à vérifier : versions compatibles au jour de l'install (langgraph évolue vite)
```

Concepts en 30 secondes :
- **State** : dictionnaire typé partagé par tout le graphe (ex : `messages`) ;
- **Node** : une fonction `(state) -> dict` — un agent, un outil, un vérificateur ;
- **Edge** : transition entre nœuds ; **conditional edge** : routage selon l'état
  (ex : « l'agent veut-il appeler un outil ? ») ;
- **Checkpointer** : sauvegarde l'état après chaque nœud → reprise, audit, time-travel ;
- **Interrupt** : pause avant/après un nœud → human-in-the-loop natif.

## 95. Exemple 1 : ReAct en 15 lignes (create_react_agent)

Le plus court chemin entre ton `agent.py` et LangGraph :

```python
import subprocess
from langchain_core.tools import tool
from langchain_openai import ChatOpenAI
from langgraph.prebuilt import create_react_agent

@tool
def disk_usage() -> str:
    """Espace disque de la machine (df -h). Lecture seule."""
    p = subprocess.run(["df", "-h"], capture_output=True, text=True, timeout=15)
    return (p.stdout + p.stderr)[-4000:]

@tool
def memory_usage() -> str:
    """RAM disponible (free -h). Lecture seule."""
    p = subprocess.run(["free", "-h"], capture_output=True, text=True, timeout=15)
    return (p.stdout + p.stderr)[-2000:]

model = ChatOpenAI(model="gpt-4.1-mini", temperature=0)  # nom à vérifier
agent = create_react_agent(model, tools=[disk_usage, memory_usage])

result = agent.invoke(
    {"messages": [("user", "Bilan rapide : disque et RAM, puis un avis.")]},
    config={"configurable": {"thread_id": "srv-web-01"}},
)
print(result["messages"][-1].content)
```

Ce que le framework t'offre déjà ici : boucle ReAct, gestion des tool calls,
historique des messages, `thread_id` pour des sessions séparées. Ce qu'il ne
t'offre pas : les budgets durs ni la liste blanche — à ajouter toi-même
(voir section 97).

## 96. Exemple 2 : graphe explicite avec vérificateur et checkpoints

Là où LangGraph bat l'agent maison : un graphe où un **vérificateur externe**
valide chaque proposition d'action avant exécution, avec reprise sur crash.

```python
from typing import Annotated, TypedDict
from langgraph.graph import StateGraph, START, END, add_messages
from langgraph.prebuilt import ToolNode, tools_condition
from langgraph.checkpointing.memory import InMemorySaver  # à vérifier selon version
from langchain_openai import ChatOpenAI

class State(TypedDict):
    messages: Annotated[list, add_messages]
    approved: bool  # le vérificateur a-t-il validé ?

model = ChatOpenAI(model="gpt-4.1-mini", temperature=0)
tools = [disk_usage, memory_usage]           # tes @tool (lecture seule ici)
llm_with_tools = model.bind_tools(tools)

def agent_node(state: State) -> dict:
    resp = llm_with_tools.invoke(state["messages"])
    return {"messages": [resp]}

def verifier_node(state: State) -> dict:
    """Vérificateur EXTERNE au modèle : inspecte le dernier appel d'outil."""
    last = state["messages"][-1]
    ok = True
    if getattr(last, "tool_calls", None):
        for tc in last.tool_calls:
            if tc["name"] not in {"disk_usage", "memory_usage"}:
                ok = False  # outil inattendu -> bloqué
    return {"approved": ok}

def route_verifier(state: State) -> str:
    return "tools" if state.get("approved") else "end_refused"

g = StateGraph(State)
g.add_node("agent", agent_node)
g.add_node("verifier", verifier_node)
g.add_node("tools", ToolNode(tools))
g.add_node("end_refused", lambda s: {"messages": [
    ("assistant", "Action refusée par le vérificateur : outil non autorisé.")]})
g.add_edge(START, "agent")
g.add_edge("agent", "verifier")
g.add_conditional_edges("verifier", route_verifier,
                        {"tools": "tools", "end_refused": "end_refused"})
g.add_edge("tools", "agent")
g.add_edge("end_refused", END)

app = g.compile(checkpointer=InMemorySaver())

# exécution avec thread : l'état est checkpointé à chaque nœud
out = app.invoke({"messages": [("user", "Bilan disque et RAM.")]},
                 config={"configurable": {"thread_id": "audit-001"}})
print(out["messages"][-1].content)
# crash en cours de route ? relance le même invoke : ça repart du checkpoint.
```

Points clés :
- `tools_condition` (utilisé dans les exemples canoniques) route vers `"tools"`
  si le modèle a émis des tool calls, sinon vers `END`. Ici on a intercalé un
  **vérificateur déterministe** (pas un modèle) : c'est le pattern
  « le code garde-fou n'est pas négociable par le modèle ».
- le checkpointer rend chaque exécution **reproductible et auditable** :
  tu peux inspecter l'état après chaque nœud (`app.get_state(config)`).

## 97. Human-in-the-loop natif : les interrupts

```python
# pause AVANT chaque exécution d'outils : un humain valide
app_hitl = g.compile(
    checkpointer=InMemorySaver(),
    interrupt_before=["tools"],
)

config = {"configurable": {"thread_id": "maintenance-042"}}
out = app_hitl.invoke({"messages": [("user", "...")]}, config=config)
# -> l'exécution s'interrompt avant "tools" : affiche l'état, demande à l'humain
state = app_hitl.get_state(config)
print("En attente de validation :", state.next)   # ('tools',)
# l'humain valide (ou modifie l'état !), puis reprise :
# from langgraph.types import Command            # à vérifier selon version
# out = app_hitl.invoke(Command(resume="go"), config=config)
```

C'est l'implémentation framework de la section 53 : la boucle se fige, l'humain
voit **exactement** l'action proposée (outil + paramètres), il approuve ou non,
et l'exécution reprend. En prod, c'est ce mécanisme (pas un `input()` bricolé)
qu'il faut utiliser.

