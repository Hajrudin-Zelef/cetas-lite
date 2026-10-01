---
id: collect-261001-ia-llm/ia-llm/top-15-frameworks-python-pour-creer-des-applications-dia-agentique-completes-3
title: "pip install -U langchain \"langchain[openai]\""
domain: ia-llm
role: reference
task: reference
actors: ["SGLang", "vLLM"]
dates: []
keywords: ["agent", "agentic", "agents", "llama", "memory", "open source", "qwen", "research", "sglang", "vllm"]
source: docs/RAG/collect-261001-ia-llm/top-15-frameworks-python-pour-creer-des-applications-dia-agentique-completes.md
source_anchor: ""
source_lines: [265, 442]
sha256: 4314a29f4d806257de1d0fe4d71f675acf67ac49ad94a6c2dc44cecdc5199d74
---

# pip install -U langchain "langchain[openai]"

```
# pip install crewai
# Set OPENAI_API_KEY before running
from crewai import Agent, Crew, Process, Task
researcher = Agent(
    role="AI Researcher",
    goal="Find the key facts about {topic}",
    backstory="You research technical topics carefully.",
)
writer = Agent(
    role="Technical Writer",
    goal="Turn research into a clear summary",
    backstory="You explain complex topics simply.",
)
research_task = Task(
    description="Research {topic} and identify three key findings.",
    expected_output="Three concise findings.",
    agent=researcher,
)
writing_task = Task(
    description="Write a short summary using the research findings.",
    expected_output="A clear one-paragraph report.",
    agent=writer,
    context=[research_task],
)
crew = Crew(
    agents=[researcher, writer],
    tasks=[research_task, writing_task],
    process=Process.sequential,
)
result = crew.kickoff(inputs={"topic": "agentic AI"})
print(result.raw)
```
Le chercheur termine la première tâche, puis le rédacteur s’appuie sur sa sortie comme contexte pour rédiger le rapport final.

### 10. MetaGPT

MetaGPT est un framework multi-agents qui simule une entreprise logicielle via des agents spécialisés tels qu’un product manager, un architecte, un chef de projet et un ingénieur. Ces agents suivent des procédures opératoires structurées pour transformer un besoin succinct en plans, conceptions techniques, documentation et code opérationnel.

L’exemple suivant crée une équipe logicielle et lui demande de construire un simple gestionnaire de tâches en ligne de commande.

```
# pip install metagpt
# Configure a supported LLM API before running
import asyncio
from metagpt.roles import (
    Architect,
    Engineer,
    ProductManager,
    ProjectManager,
)
from metagpt.team import Team
async def startup(idea: str):
    company = Team()
    company.hire([
        ProductManager(),
        Architect(),
        ProjectManager(),
        Engineer(),
    ])
    company.invest(investment=3.0)
    company.run_project(idea=idea)
    await company.run(n_round=5)
if __name__ == "__main__":
    asyncio.run(
        startup("Build a simple command-line task manager")
    )
```
Les agents se répartissent les exigences selon leurs rôles et collaborent pour planifier et développer le projet logiciel.

### 11. AgentScope

AgentScope est un framework Python pour créer des agents et des applications multi-agents contrôlables et observables. Il propose des agents prêts à l’emploi, des outils, de la mémoire, du routage de messages, du streaming, des appels d’outils en parallèle, des workflows, du tracing et l’intervention humaine. AgentScope Studio permet également d’inspecter et de visualiser l’exécution des agents.

L’exemple suivant crée un agent ReAct capable d’écrire et d’exécuter du code Python pour réaliser un calcul.

```
# pip install agentscope
# Set DASHSCOPE_API_KEY before running
import asyncio
import os
from agentscope.agent import ReActAgent
from agentscope.formatter import DashScopeChatFormatter
from agentscope.memory import InMemoryMemory
from agentscope.message import Msg
from agentscope.model import DashScopeChatModel
from agentscope.tool import Toolkit, execute_python_code
async def main():
    toolkit = Toolkit()
    toolkit.register_tool_function(execute_python_code)
    agent = ReActAgent(
        name="Nova",
        sys_prompt="You are a helpful Python assistant named Nova.",
        model=DashScopeChatModel(
            model_name="qwen-max",
            api_key=os.environ["DASHSCOPE_API_KEY"],
            stream=True,
        ),
        formatter=DashScopeChatFormatter(),
        toolkit=toolkit,
        memory=InMemoryMemory(),
    )
    await agent(
        Msg(
            name="user",
            content="Use Python to calculate the sum of numbers from 1 to 100.",
            role="user",
        )
    )
if __name__ == "__main__":
    asyncio.run(main())
```
L’agent décide d’appeler l’outil d’exécution Python, lance le code généré et utilise le résultat pour répondre à l’utilisateur.

### 12. CAMEL-AI

CAMEL-AI est un framework Python open source pour créer des agents, des sociétés multi-agents et des simulations de rôles. Très orienté recherche, il offre des composants pour la collaboration entre agents, les outils, la mémoire, la récupération, la génération de données et la simulation de mondes. Il prend en charge les modèles cloud et les modèles hébergés localement via des plateformes comme Ollama, vLLM et SGLang.

L’exemple suivant crée un agent capable de rechercher sur le web avant de répondre à une question.

```
# pip install "camel-ai[web_tools]"
# Set OPENAI_API_KEY before running
from camel.agents import ChatAgent
from camel.models import ModelFactory
from camel.toolkits import SearchToolkit
from camel.types import ModelPlatformType, ModelType
model = ModelFactory.create(
    model_platform=ModelPlatformType.OPENAI,
    model_type=ModelType.GPT_5_5,
    model_config_dict={"temperature": 0.0},
)
agent = ChatAgent(
    system_message="You are a helpful AI research assistant.",
    model=model,
    tools=[SearchToolkit().search_duckduckgo],
)
response = agent.step(
    "What are the main uses of multi-agent AI systems?"
)
print(response.msgs[0].content)
```
L’agent utilise la recherche DuckDuckGo lorsqu’il a besoin d’informations à jour, puis renvoie la réponse finale.

## Frameworks orientés données et RAG

Ces frameworks connectent les LLM et les agents à des documents, bases de données, API et autres sources de données privées pour récupérer un contexte pertinent avant de répondre.

### 13. LlamaIndex

LlamaIndex est un framework Python pour créer des applications sensibles au contexte sur des données privées ou métier. Il propose des connecteurs, index, retrieveurs, moteurs de requêtes, agents et workflows pour des cas tels que le RAG, la recherche d’entreprise et les assistants de documents.

LlamaIndex a été l’un de mes frameworks préférés lorsque j’ai commencé à développer des applications RAG. Il simplifie énormément le chargement, l’indexation et l’interrogation des données, et je trouve souvent son code plus simple et plus court que l’équivalent avec LangChain. Vous pouvez suivre le cours LlamaIndex pour en savoir plus.

L’exemple suivant charge des documents depuis un dossier, crée un index interrogeable et répond à une question en s’appuyant sur le contexte récupéré.

```
# pip install llama-index
# Set OPENAI_API_KEY before running
from llama_index.core import SimpleDirectoryReader, VectorStoreIndex
documents = SimpleDirectoryReader("data").load_data()
index = VectorStoreIndex.from_documents(documents)
query_engine = index.as_query_engine()
response = query_engine.query(
    "What are the main findings in these documents?"
)
print(response)
```
LlamaIndex indexe les fichiers du dossier data et récupère les informations pertinentes avant de générer la réponse.

### 14. Haystack

Haystack est un framework Python open source pour créer des pipelines RAG prêts pour la production, des systèmes de recherche sémantique, des agents et des applications d’IA centrées sur la donnée. Sa structure modulaire en pipeline offre un contrôle clair sur la récupération d’information, son insertion dans le prompt et l’envoi au modèle.

J’ai utilisé Haystack pour construire des applications RAG et multi-agents sans problème majeur. Bien qu’il soit généraliste, il est particulièrement adapté aux applications qui relient les LLM à des documents, du texte, des systèmes de recherche et d’autres sources de données.

L’exemple suivant stocke quelques documents, récupère le plus pertinent et s’en sert pour répondre à une question.

