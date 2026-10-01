---
id: collect-261001-ia-llm/ia-llm/top-15-frameworks-python-pour-creer-des-applications-dia-agentique-completes-4
title: "pip install -U langchain \"langchain[openai]\""
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Google", "Hugging Face", "OpenAI"]
dates: []
keywords: ["agent", "agents", "claude", "gemini", "inference", "memory", "open source", "qwen", "valuation"]
source: docs/RAG/collect-261001-ia-llm/top-15-frameworks-python-pour-creer-des-applications-dia-agentique-completes.md
source_anchor: ""
source_lines: [443, 573]
sha256: d44f3e533448ce7d4a7045b86e56bd34cdbf4b3dc671dd73474dc78f1b3e9164
---

# pip install -U langchain "langchain[openai]"

```
# pip install -U haystack-ai
# Set OPENAI_API_KEY before running
from haystack import Document, Pipeline
from haystack.components.builders import ChatPromptBuilder
from haystack.components.generators.chat import (
    OpenAIResponsesChatGenerator,
)
from haystack.components.retrievers import InMemoryBM25Retriever
from haystack.dataclasses import ChatMessage
from haystack.document_stores.in_memory import InMemoryDocumentStore
document_store = InMemoryDocumentStore()
document_store.write_documents([
    Document(content="The London AI event starts at 6:00 PM."),
    Document(content="The Berlin AI event starts at 7:00 PM."),
])
template = [
    ChatMessage.from_system(
        "Answer using the following documents:\n"
        "{% for doc in documents %}{{ doc.content }}{% endfor %}"
    ),
    ChatMessage.from_user("{{ question }}"),
]
pipeline = Pipeline()
pipeline.add_component(
    "retriever",
    InMemoryBM25Retriever(document_store),
)
pipeline.add_component(
    "prompt_builder",
    ChatPromptBuilder(template=template),
)
pipeline.add_component(
    "llm",
    OpenAIResponsesChatGenerator(model="gpt-5.5"),
)
pipeline.connect("retriever", "prompt_builder.documents")
pipeline.connect("prompt_builder", "llm")
question = "When does the London AI event start?"
result = pipeline.run({
    "retriever": {"query": question},
    "prompt_builder": {"question": question},
})
print(result["llm"]["replies"][0].text)
```
Haystack récupère le document correspondant, l’ajoute au prompt et utilise GPT-5.5 pour produire une réponse étayée.

## Frameworks légers et modèles ouverts

Ces frameworks proposent des abstractions plus simples pour créer des agents avec des modèles ouverts hébergés ou locaux.

### 15. Hugging Face smolagents

Hugging Face smolagents est un framework Python léger pour créer des agents avec très peu de code. Son CodeAgent agit en écrivant du Python, tandis que ToolCallingAgent utilise des appels d’outils structurés. Il fonctionne particulièrement bien avec les modèles Hugging Face, les Inference Providers et les modèles open source hébergés localement.

L’exemple suivant crée un agent de recherche web, puis le confie à un agent manager qui décide quand déléguer la recherche.

```
# pip install -U "smolagents[toolkit]"
# Set HF_TOKEN before running
from smolagents import (
    CodeAgent,
    InferenceClientModel,
    ToolCallingAgent,
    WebSearchTool,
)
model = InferenceClientModel(
    model_id="Qwen/Qwen3-Next-80B-A3B-Thinking"
)
web_agent = ToolCallingAgent(
    tools=[WebSearchTool()],
    model=model,
    name="web_search_agent",
    description="Searches the web and returns useful information.",
)
manager_agent = CodeAgent(
    tools=[],
    model=model,
    managed_agents=[web_agent],
)
result = manager_agent.run(
    "Who created Hugging Face, and when was it founded?"
)
print(result)
```
Le manager décide quand il a besoin d’informations à jour, délègue la recherche à l’agent web spécialiste, puis produit la réponse finale.

## Comparatif rapide des frameworks

Le tableau ci-dessous compare la principale force de chaque framework et le type d’application d’IA agentique pour lequel il est le plus adapté.

| Framework | Atout principal | À choisir lorsque | 
| LangChain | Large écosystème d’intégrations | Vous avez besoin d’intégrations étendues, d’outils et de flexibilité | 
| LangGraph | Orchestration stateful par graphe | Vous avez besoin de boucles contrôlées, d’exécutions durables et de workflows complexes | 
| Agno | Plateforme d’agents complète | Vous voulez créer, auto-héberger et gérer des services d’agents | 
| Pydantic AI | Développement Python avec typage sûr | Vous avez besoin d’outils validés, de dépendances et de sorties structurées | 
| OpenAI Agents SDK | Développement d’agents et tracing simplifiés | Vous utilisez principalement les modèles OpenAI et voulez intégrer rapidement | 
| Google ADK | Développement natif pour Gemini | Vous utilisez Gemini, Google Cloud ou des workflows multi-agents | 
| Claude Agent SDK | Outils pour ordinateur, fichiers et commandes | Vous concevez des agents de code, de recherche ou longue durée avec Claude | 
| CrewAI | Équipes d’agents basées sur les rôles | Plusieurs agents spécialisés doivent collaborer sur un projet | 
| MetaGPT | Simulation d’entreprise logicielle | Vous voulez que des agents planifient et génèrent des projets logiciels complets | 
| AgentScope | Agents observables et contrôlables | Vous avez besoin d’agents traçables ou d’applications multi-agents | 
| CAMEL-AI | Recherche et simulation multi-agents | Vous étudiez le role-play, les sociétés d’agents ou de grands systèmes d’agents | 
| LlamaIndex | Applications connectées aux données et sensibles au contexte | Vos agents doivent récupérer et raisonner sur des documents privés | 
| Haystack | Pipelines RAG modulaires et transparents | Vous voulez contrôler la récupération, le prompting et le flux de données | 
| smolagents | Agents légers pour modèles ouverts | Vous souhaitez des abstractions minimales, des code agents ou le support de modèles locaux | 

## Derniers conseils

Si vous débutez avec l’IA agentique, commencez par un SDK officiel comme l’OpenAI Agents SDK, Google ADK ou le Claude Agent SDK. Ils fonctionnent de près avec leurs modèles et API, donc la mise en route est généralement simple. Ajoutez une clé d’API, créditez votre compte, et commencez à construire.

J’apprécie ces SDK car ils facilitent la création d’agents, la connexion d’outils et l’inspection de ce qui s’est passé pendant une exécution. Avec quelques lignes de Python, vous pouvez créer un agent utile, voire un petit système multi-agents.

Quand j’ai besoin de plus de flexibilité, je me tourne généralement vers LangChain ou LangGraph. LangChain est utile lorsque j’ai besoin de nombreuses intégrations, tandis que LangGraph m’offre plus de contrôle sur l’état, les boucles et les workflows longs.

Pour des équipes d’agents aux rôles variés, CrewAI est une bonne option. Pour le RAG, les documents privés et les applications connectées aux données, je privilégie LlamaIndex ou Haystack.

Enfin, smolagents est un excellent choix quand je veux quelque chose de léger, simple et facile à utiliser avec des modèles ouverts ou locaux.

Mon conseil est de démarrer avec le framework le plus simple qui réponde à votre besoin. Créez une version minimale, testez-la correctement, et n’ajoutez de complexité que lorsque c’est réellement nécessaire.

## FAQs

### Combien coûte l’exécution d’un agent d’IA par rapport à un prompt LLM standard ?

Si les frameworks Python eux-mêmes sont open source et gratuits, exécuter des agents peut coûter sensiblement plus cher que des appels API LLM standards. Les agents fonctionnent en boucle, renvoyant sans cesse au modèle le system prompt, les descriptions d’outils, les sorties d’outils précédentes et les étapes de raisonnement (comme ReAct) : la consommation de tokens s’accumule très vite. Une seule requête utilisateur peut déclencher cinq ou six appels LLM avant que l’agent n’aboutisse à une réponse finale. Pour maîtriser les coûts, les développeurs utilisent souvent des modèles plus petits et moins chers pour les tâches simples de routage, et réservent les modèles plus puissants au raisonnement complexe.

### Comment évaluer si mon agent fonctionne réellement bien ?

Tester des agents non déterministes demande plus que des tests unitaires classiques. Les développeurs recourent à des frameworks d’évaluation spécialisés de type « LLM-as-a-judge », comme **Ragas**, **TruLens** ou **DeepEval**. Ces outils exécutent votre agent sur un jeu de questions de test et notent les sorties selon des métriques spécifiques, par exemple :

