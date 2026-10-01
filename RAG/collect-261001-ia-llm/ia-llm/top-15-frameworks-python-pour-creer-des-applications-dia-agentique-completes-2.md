---
id: collect-261001-ia-llm/ia-llm/top-15-frameworks-python-pour-creer-des-applications-dia-agentique-completes-2
title: "pip install -U langchain \"langchain[openai]\""
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "OpenAI"]
dates: []
keywords: ["agent", "agents", "claude", "gemini", "gpt-5.6", "mcp", "open source", "sol"]
source: docs/RAG/collect-261001-ia-llm/top-15-frameworks-python-pour-creer-des-applications-dia-agentique-completes.md
source_anchor: ""
source_lines: [123, 264]
sha256: 08e8e874519add9cf523b83bf44cf0298ac8b883771667f830a4d7ee29c4dfc4
---

# pip install -U langchain "langchain[openai]"

```
# pip install -U pydantic-ai
from pydantic import BaseModel, Field
from pydantic_ai import Agent
class ArticlePlan(BaseModel):
    title: str
    key_points: list[str]
    reading_time_minutes: int = Field(ge=1)
agent = Agent(
    "openai:gpt-5.6-sol",
    output_type=ArticlePlan,
    instructions="Create concise article plans for technical readers.",
)
result = agent.run_sync(
    "Create a short article plan about building AI agents in Python."
)
print(result.output)
```
L’agent génère un objet ArticlePlan et valide chaque champ selon les types et règles Python requis avant de le retourner.

## SDK officiels d’agents

Les SDK officiels sont maintenus par les fournisseurs d’IA et offrent généralement l’intégration la plus simple avec leurs modèles, API, outils, systèmes de déploiement et plateformes d’observabilité.

### 5. OpenAI Agents SDK

Le SDK OpenAI Agents est un framework léger, pensé d’abord pour Python permettant de créer des applications mono-agent et multi-agents. Il fournit une boucle d’agent intégrée, des outils de fonction, des handoffs, des garde-fous, des sessions, des validations humaines et du tracing, sans surcharger d’abstractions.

J’apprécie particulièrement utiliser l’OpenAI Agents SDK car il est simple, rapide à intégrer et facile à comprendre. J’ai construit de nombreuses applications avec, sans problème majeur. Il crée également automatiquement des traces visibles dans le Dashboard OpenAI, ce qui facilite l’inspection des réponses du modèle, des appels d’outils, des handoffs et du flot d’exécution complet.

L’exemple suivant crée un agent d’histoire capable d’appeler une fonction Python pour récupérer un fait historique surprenant.

```
# pip install openai-agents
import asyncio
from agents import Agent, Runner, function_tool
@function_tool
def history_fun_fact() -> str:
    """Return a surprising historical fact."""
    return "The first computer programmer, Ada Lovelace, lived in the 1800s."
agent = Agent(
    name="History Assistant",
    instructions=(
        "Answer history questions clearly and briefly. "
        "Use history_fun_fact when it is helpful."
    ),
    tools=[history_fun_fact],
)
async def main():
    result = await Runner.run(
        agent,
        "Tell me something surprising about the history of computing.",
    )
    print(result.final_output)
if __name__ == "__main__":
    asyncio.run(main())
```
Le SDK exécute l’agent, appelle l’outil d’histoire si nécessaire, renvoie son résultat au modèle et produit la réponse finale. L’exécution complète peut aussi être inspectée via le trace viewer.

### 7. Google Agent Development Kit

L’Agent Development Kit de Google, ou ADK, est un framework open source pour créer, évaluer et déployer des agents individuels, des assistants utilisant des outils, des workflows en graphe et des systèmes multi-agents. Il fonctionne particulièrement bien avec les modèles Gemini, tout en prenant en charge d’autres fournisseurs.

ADK est devenu l’un de mes frameworks favoris, car il est simple et s’intègre naturellement avec Gemini.

Je l’ai utilisé pour créer plusieurs applications sans problème majeur. Même si je trouve encore le SDK OpenAI Agents plus simple à intégrer et plus flexible en matière de contrôles et de fonctionnalités natives, ADK est une excellente alternative, particulièrement compétitive pour des applications basées sur Gemini.

L’exemple suivant crée un assistant d’événements capable d’appeler une fonction Python pour récupérer l’heure de début d’un événement.

```
# pip install google-adk
from google.adk.agents.llm_agent import Agent
def get_event_time(city: str) -> dict:
    """Return the event starting time for a city."""
    return {
        "status": "success",
        "city": city,
        "time": "6:00 PM",
    }
root_agent = Agent(
    model="gemini-flash-latest",
    name="event_assistant",
    description="Provides information about events in different cities.",
    instruction=(
        "Answer event questions clearly. "
        "Use the get_event_time tool when the starting time is requested."
    ),
    tools=[get_event_time],
)
```
L’agent utilise la fonction comme un outil dès qu’il doit récupérer l’horaire d’un événement pour une ville donnée.

### 8. Claude Agent SDK

Le Claude Agent SDK est le framework Python et TypeScript d’Anthropic pour créer des agents autonomes en s’appuyant sur la même boucle d’agent, les mêmes outils et la même gestion du contexte que ceux de Claude Code. Il est particulièrement utile pour des agents devant lire des fichiers, éditer du code, exécuter des commandes, se connecter à des outils MCP et mener à bien des tâches longues.

Anthropic, OpenAI et Google intègrent étroitement leurs modèles à leurs propres frameworks d’agents, ce qui accélère la mise en place des outils, des sessions, du tracing et d’autres fonctionnalités avancées.

Cependant, le Claude Agent SDK est conçu spécifiquement pour Claude plutôt que pour des modèles open source ou hébergés localement. Vous pouvez en savoir plus dans notre tutoriel Claude Agent SDK.

L’exemple suivant crée un agent de code qui passe en revue un fichier Python, identifie les bugs et les corrige automatiquement.

```
# pip install claude-agent-sdk
import asyncio
from claude_agent_sdk import (
    AssistantMessage,
    ClaudeAgentOptions,
    ResultMessage,
    query,
)
async def main():
    async for message in query(
        prompt="Review app.py for errors that could cause crashes and fix them.",
        options=ClaudeAgentOptions(
            allowed_tools=["Read", "Edit", "Glob"],
            permission_mode="acceptEdits",
        ),
    ):
        if isinstance(message, AssistantMessage):
            for block in message.content:
                if hasattr(block, "text"):
                    print(block.text)
                elif hasattr(block, "name"):
                    print(f"Tool used: {block.name}")
        elif isinstance(message, ResultMessage):
            print(f"Completed: {message.subtype}")
if __name__ == "__main__":
    asyncio.run(main())
```
Le SDK exécute la boucle de l’agent pendant que Claude lit le fichier, sélectionne les outils requis, modifie le code et diffuse sa progression jusqu’à la fin de la tâche.

## Frameworks d’orchestration multi-agents

Les frameworks multi-agents coordonnent plusieurs agents spécialisés, chacun étant assigné à un rôle, un objectif, un ensemble d’outils ou une étape d’un workflow plus large.

### 9. CrewAI

CrewAI est un framework Python populaire pour créer des équipes multi-agents. Les développeurs définissent des agents avec différents rôles, objectifs et tâches, puis les combinent au sein d’une équipe qui collabore de manière autonome. Les tâches peuvent s’enchaîner séquentiellement ou suivre un processus hiérarchique où un manager coordonne et délègue le travail aux agents les plus adaptés.

L’exemple Python ci-dessous crée un chercheur et un rédacteur qui collaborent pour produire un court rapport.

