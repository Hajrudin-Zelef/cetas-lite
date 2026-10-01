---
id: collect-261001-ia-llm/ia-llm/top-15-frameworks-python-pour-creer-des-applications-dia-agentique-completes-1
title: "pip install -U langchain \"langchain[openai]\""
domain: ia-llm
role: reference
task: reference
actors: ["Google", "OpenAI"]
dates: []
keywords: ["agent", "agents", "open source", "research", "valuation"]
source: docs/RAG/collect-261001-ia-llm/top-15-frameworks-python-pour-creer-des-applications-dia-agentique-completes.md
source_anchor: ""
source_lines: [1, 122]
sha256: e9b857918d001721818ddeb97d14b452758fdbf301b2612116413250c89558ff
---

# pip install -U langchain "langchain[openai]"

Cours

Créer des applications d’IA agentique est aujourd’hui bien plus simple qu’il y a quelques années. Plutôt que de concevoir de zéro des boucles d’agent, des intégrations d’outils, la mémoire, la récupération et l’orchestration, les développeurs peuvent s’appuyer sur des frameworks Python qui intègrent déjà l’essentiel de ces composants.

J’ai utilisé bon nombre des frameworks présentés dans cet article pour concevoir des applications RAG, des agents autonomes, des systèmes multi-agents et des outils d’IA déployables.

J’utilise souvent LangChain pour des workflows d’agents flexibles, LlamaIndex et Haystack pour des applications connectées aux données, ainsi que le SDK OpenAI Agents ou Google ADK lorsque je veux intégrer rapidement des modèles fournis par ces éditeurs.

Certains de ces frameworks ont tellement progressé qu’une solution sur mesure est souvent inutile. Avec une clé d’API, quelques crédits de modèles et le bon framework, vous pouvez créer un agent, le connecter à des outils ou à des données privées, le tester, puis le déployer via une API ou une interface utilisateur.

Cet article compare 15 frameworks Python répartis en cinq catégories : frameworks d’agents généralistes, SDK officiels d’agents, frameworks d’orchestration multi-agents, frameworks orientés données et RAG, et frameworks légers pour modèles ouverts.

Chaque section inclut également un court exemple Python à copier-coller pour comprendre rapidement le fonctionnement du framework.

## Frameworks d’agents généralistes

Ces frameworks fournissent des blocs de construction complets pour créer des agents, connecter des outils, gérer l’état et piloter des workflows en plusieurs étapes.

### 1. LangChain

Pour de nombreux développeurs, le développement moderne d’applications LLM a commencé avec LangChain, un framework Python open source conçu pour connecter les modèles de langage à des données externes, des outils, des API et de la logique applicative. Il a été l’un des premiers frameworks largement adoptés pour utiliser l’API OpenAI avec des sites web, des documents, des bases vectorielles et d’autres sources de données afin de créer des applications plus contextuelles.

Aujourd’hui, LangChain a évolué en un écosystème complet pour créer des applications d’IA agentique de bout en bout, incluant des systèmes RAG, des agents capables d’utiliser des outils, des workflows, des intégrations, du monitoring et de l’évaluation. Je vous recommande de consulter le parcours AI Engineering with LangChain pour aller plus loin.

L’exemple suivant crée un agent LangChain simple capable d’utiliser un outil de recherche dans la documentation pour trouver des informations pertinentes avant de répondre.

```
# pip install -U langchain "langchain[openai]"
from langchain.agents import create_agent
def search_docs(query: str) -> str:
    """Search the company documentation."""
    return f"Documentation found for: {query}"
agent = create_agent(
    model="openai:gpt-5.5",
    tools=[search_docs],
    system_prompt="Use the documentation tool when needed.",
)
result = agent.invoke({
    "messages": [{
        "role": "user",
        "content": "What is our remote-work policy?"
    }]
})
print(result["messages"][-1].content)
```
### 2. LangGraph

Si LangChain facilite la création d’agents et d’outils, LangGraph apporte davantage de contrôle sur la façon dont ces agents fonctionnent. Il permet de structurer l’application comme un graphe d’étapes connectées, ce qui simplifie la maîtrise des boucles, des appels d’outils, de l’état partagé, des validations humaines et des workflows multi-agents.

Je le trouve particulièrement utile pour des agents autonomes qui doivent itérer : décider de la prochaine action, choisir le bon outil, vérifier le résultat et poursuivre jusqu’à ce que la tâche soit terminée. LangGraph gère aussi la persistance, le streaming, les checkpoints et la reprise pour les applications longues. Notre tutoriel LangGraph donne plus de détails.

L’exemple suivant crée un agent qui peut appeler à plusieurs reprises des outils de calcul jusqu’à disposer de suffisamment d’informations pour répondre :

```
# pip install -U langgraph langchain langchain-openai
from langchain.chat_models import init_chat_model
from langchain.tools import tool
from langgraph.graph import MessagesState, StateGraph, START
from langgraph.prebuilt import ToolNode, tools_condition
model = init_chat_model("openai:gpt-5.5", temperature=0)
@tool
def add(a: int, b: int) -> int:
    """Add two numbers."""
    return a + b
@tool
def multiply(a: int, b: int) -> int:
    """Multiply two numbers."""
    return a * b
tools = [add, multiply]
model_with_tools = model.bind_tools(tools)
def call_model(state: MessagesState):
    response = model_with_tools.invoke(state["messages"])
    return {"messages": [response]}
builder = StateGraph(MessagesState)
builder.add_node("agent", call_model)
builder.add_node("tools", ToolNode(tools))
builder.add_edge(START, "agent")
builder.add_conditional_edges("agent", tools_condition)
builder.add_edge("tools", "agent")
agent = builder.compile()
result = agent.invoke({
    "messages": [{
        "role": "user",
        "content": "Add 12 and 8, then multiply the result by 3."
    }]
})
print(result["messages"][-1].content)
```
Le graphe envoie la requête au modèle, exécute les outils nécessaires, puis reboucle vers le modèle jusqu’à produire une réponse finale.

### 3. Agno

Agno est un framework Python pour créer des agents, des équipes multi-agents et des workflows structurés. Il inclut aussi AgentOS, qui aide à exposer les agents via des API, stocker les sessions et traces, et gérer le tout en production. Comme il peut tourner sur votre propre infrastructure, les organisations gardent un meilleur contrôle sur leurs données et leur sécurité.

L’exemple suivant crée un agent de recherche capable d’explorer le web avant de produire une réponse concise.

```
# pip install -U agno ddgs openai
from agno.agent import Agent
from agno.models.openai import OpenAIResponses
from agno.tools.duckduckgo import DuckDuckGoTools
agent = Agent(
    name="AI Research Assistant",
    model=OpenAIResponses(id="gpt-5.5"),
    tools=[DuckDuckGoTools()],
    instructions="Search when needed and keep the answer concise.",
)
agent.print_response(
    "Find one recent development in open-source AI and summarize it.",
    stream=True,
)
```
L’agent utilise l’outil de recherche pour obtenir des informations à jour et diffuse un court résumé à l’utilisateur.

### 4. Pydantic AI

Pydantic AI est un framework Python-native, créé par l’équipe derrière Pydantic, pour construire des agents et des applications d’IA générative. L’expérience rappelle celle des modèles Pydantic ou de FastAPI : agents, outils, dépendances et sorties sont définis via des fonctions Python standards, des hints de types, des décorateurs et des classes BaseModel. Le framework s’appuie ensuite sur la validation Pydantic pour fiabiliser les arguments des outils et les sorties des modèles.

Je vous recommande le tutoriel Pydantic AI pour découvrir le framework.

L’exemple suivant crée un agent qui renvoie un plan d’article validé sous forme de modèle Pydantic.

