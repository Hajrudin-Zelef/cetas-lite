---
id: collect-240926-datacamp/datacamp/top-15-frameworks-python-pour-creer-des-applications-dia-agentique-completes
title: "pip install -U langchain \"langchain[openai]\""
domain: datacamp
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Google", "Hugging Face", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["agent", "agentic", "agents", "claude", "cost", "gemini", "gpt-5.6", "guardrails", "inference", "llama", "mcp", "memory"]
source: docs/RAG/clean_en/datacamp/top-15-frameworks-python-pour-creer-des-applications-dia-agentique-completes.md
source_anchor: ""
source_lines: [1, 586]
sha256: f499e817941368749ce99b80a5108a5f7a11bc5ab2428f5b483e8de24f1bd25c
---

# pip install -U langchain "langchain[openai]"

<!-- source: https://www.datacamp.com/fr/blog/top-python-frameworks-for-agentic-ai -->

Course

Building agentic AI applications is now much simpler than it was a few years ago. Rather than designing agent loops, tool integrations, memory, retrieval, and orchestration from scratch, developers can rely on Python frameworks that already incorporate most of these components.

I have used many of the frameworks presented in this article to design RAG applications, autonomous agents, multi-agent systems, and deployable AI tools.

I often use LangChain for flexible agent workflows, LlamaIndex and Haystack for data-connected applications, as well as the OpenAI Agents SDK or Google ADK when I want to quickly integrate models provided by these vendors.

Some of these frameworks have advanced so much that a custom solution is often unnecessary. With an API key, a few model credits, and the right framework, you can create an agent, connect it to tools or private data, test it, and then deploy it via an API or a user interface.

This article compares 15 Python frameworks divided into five categories: general-purpose agent frameworks, official agent SDKs, multi-agent orchestration frameworks, data- and RAG-oriented frameworks, and lightweight frameworks for open models.

Each section also includes a short Python example to copy and paste to quickly understand how the framework works.

## General-purpose agent frameworks

These frameworks provide complete building blocks for creating agents, connecting tools, managing state, and driving multi-step workflows.

### 1. LangChain

For many developers, modern LLM application development began with LangChain, an open-source Python framework designed to connect language models to external data, tools, APIs, and application logic. It was one of the first widely adopted frameworks for using the OpenAI API with websites, documents, vector databases, and other data sources to create more contextual applications.

Today, LangChain has evolved into a complete ecosystem for building end-to-end agentic AI applications, including RAG systems, agents capable of using tools, workflows, integrations, monitoring, and evaluation. I recommend checking out the AI Engineering with LangChain path to go further.

The following example creates a simple LangChain agent capable of using a documentation search tool to find relevant information before responding.

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

If LangChain makes it easier to create agents and tools, LangGraph provides more control over how these agents work. It allows you to structure the application as a graph of connected steps, which simplifies control over loops, tool calls, shared state, human validations, and multi-agent workflows.

I find it particularly useful for autonomous agents that need to iterate: decide the next action, choose the right tool, check the result, and continue until the task is complete. LangGraph also handles persistence, streaming, checkpoints, and resumption for long-running applications. Our LangGraph tutorial provides more details.

The following example creates an agent that can repeatedly call calculation tools until it has enough information to respond:

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
The graph sends the request to the model, executes the necessary tools, then loops back to the model until a final response is produced.

### 3. Agno

Agno is a Python framework for creating agents, multi-agent teams, and structured workflows. It also includes AgentOS, which helps expose agents via APIs, store sessions and traces, and manage everything in production. Since it can run on your own infrastructure, organizations retain better control over their data and security.

The following example creates a research agent capable of browsing the web before producing a concise response.

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
The agent uses the search tool to get up-to-date information and streams a short summary to the user.

### 4. Pydantic AI

Pydantic AI is a Python-native framework, created by the team behind Pydantic, for building agents and generative AI applications. The experience is reminiscent of Pydantic models or FastAPI: agents, tools, dependencies, and outputs are defined through standard Python functions, type hints, decorators, and BaseModel classes. The framework then relies on Pydantic validation to make tool arguments and model outputs more reliable.

I recommend the Pydantic AI tutorial to discover the framework.

The following example creates an agent that returns a validated article plan as a Pydantic model.

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
The agent generates an ArticlePlan object and validates each field according to the required Python types and rules before returning it.

## Official agent SDKs

Official SDKs are maintained by AI providers and generally offer the simplest integration with their models, APIs, tools, deployment systems, and observability platforms.

### 5. OpenAI Agents SDK

The OpenAI Agents SDK is a lightweight framework, designed primarily for Python, for building single-agent and multi-agent applications. It provides a built-in agent loop, function tools, handoffs, guardrails, sessions, human validations, and tracing, without overloading with abstractions.

I particularly enjoy using the OpenAI Agents SDK because it is simple, quick to integrate, and easy to understand. I have built many applications with it without major issues. It also automatically creates traces visible in the OpenAI Dashboard, which makes it easy to inspect model responses, tool calls, handoffs, and the complete execution flow.

The following example creates a history agent capable of calling a Python function to retrieve a surprising historical fact.

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
The SDK runs the agent, calls the history tool if necessary, returns its result to the model, and produces the final response. The complete execution can also be inspected via the trace viewer.

### 7. Google Agent Development Kit

Google's Agent Development Kit, or ADK, is an open-source framework for building, evaluating, and deploying individual agents, tool-using assistants, graph workflows, and multi-agent systems. It works particularly well with Gemini models, while supporting other providers.

ADK has become one of my favorite frameworks because it is simple and integrates naturally with Gemini.

I have used it to build several applications without major issues. Although I still find the OpenAI Agents SDK simpler to integrate and more flexible in terms of controls and native features, ADK is an excellent alternative, particularly competitive for Gemini-based applications.

The following example creates an event assistant capable of calling a Python function to retrieve an event's start time.

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
The agent uses the function as a tool whenever it needs to retrieve the time of an event for a given city.

### 8. Claude Agent SDK

The Claude Agent SDK is Anthropic's Python and TypeScript framework for building autonomous agents by relying on the same agent loop, tools, and context management as Claude Code. It is particularly useful for agents that need to read files, edit code, run commands, connect to MCP tools, and carry out long tasks.

Anthropic, OpenAI, and Google tightly integrate their models with their own agent frameworks, which accelerates the setup of tools, sessions, tracing, and other advanced features.

However, the Claude Agent SDK is designed specifically for Claude rather than for open source or locally hosted models. You can learn more in our Claude Agent SDK tutorial.

The following example creates a coding agent that reviews a Python file, identifies bugs, and fixes them automatically.

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
The SDK runs the agent loop while Claude reads the file, selects the required tools, modifies the code, and streams its progress until the task is complete.

## Multi-agent orchestration frameworks

Multi-agent frameworks coordinate several specialized agents, each assigned to a role, a goal, a set of tools, or a step in a larger workflow.

### 9. CrewAI

CrewAI is a popular Python framework for building multi-agent teams. Developers define agents with different roles, goals, and tasks, then combine them into a team that collaborates autonomously. Tasks can be chained sequentially or follow a hierarchical process in which a manager coordinates and delegates work to the most suitable agents.

The Python example below creates a researcher and a writer who collaborate to produce a short report.

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
The researcher completes the first task, then the writer uses its output as context to draft the final report.

### 10. MetaGPT

MetaGPT is a multi-agent framework that simulates a software company through specialized agents such as a product manager, an architect, a project manager, and an engineer. These agents follow structured operating procedures to transform a brief requirement into plans, technical designs, documentation, and working code.

The following example creates a software team and asks it to build a simple command-line task manager.

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
The agents divide the requirements according to their roles and collaborate to plan and develop the software project.

### 11. AgentScope

AgentScope is a Python framework for building controllable and observable agents and multi-agent applications. It offers ready-to-use agents, tools, memory, message routing, streaming, parallel tool calls, workflows, tracing, and human intervention. AgentScope Studio also makes it possible to inspect and visualize agent execution.

The following example creates a ReAct agent capable of writing and executing Python code to perform a calculation.

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
The agent decides to call the Python execution tool, runs the generated code, and uses the result to respond to the user.

### 12. CAMEL-AI

CAMEL-AI is an open-source Python framework for building agents, multi-agent societies, and role-playing simulations. Highly research-oriented, it offers components for collaboration between agents, tools, memory, retrieval, data generation, and world simulation. It supports cloud models and locally hosted models via platforms such as Ollama, vLLM, and SGLang.

The following example creates an agent capable of searching the web before answering a question.

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
The agent uses DuckDuckGo search when it needs up-to-date information, then returns the final answer.

## Data-oriented and RAG frameworks

These frameworks connect LLMs and agents to documents, databases, APIs, and other private data sources to retrieve relevant context before responding.

### 13. LlamaIndex

LlamaIndex is a Python framework for building context-aware applications on private or business data. It offers connectors, indexes, retrievers, query engines, agents, and workflows for use cases such as RAG, enterprise search, and document assistants.

LlamaIndex was one of my favorite frameworks when I started developing RAG applications. It greatly simplifies loading, indexing, and querying data, and I often find its code simpler and shorter than the equivalent with LangChain. You can follow the LlamaIndex course to learn more.

The following example loads documents from a folder, creates a queryable index, and answers a question based on the retrieved context.

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
LlamaIndex indexes the files in the data folder and retrieves the relevant information before generating the response.

### 14. Haystack

Haystack is an open-source Python framework for building production-ready RAG pipelines, semantic search systems, agents, and data-centric AI applications. Its modular pipeline structure provides clear control over information retrieval, its insertion into the prompt, and sending it to the model.

I have used Haystack to build RAG and multi-agent applications without major issues. Although it is general-purpose, it is particularly well suited to applications that connect LLMs to documents, text, search systems, and other data sources.

The following example stores a few documents, retrieves the most relevant one, and uses it to answer a question.

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
Haystack retrieves the matching document, adds it to the prompt, and uses GPT-5.5 to produce a grounded answer.

## Lightweight frameworks and open models

These frameworks offer simpler abstractions for building agents with hosted or local open models.

### 15. Hugging Face smolagents

Hugging Face smolagents is a lightweight Python framework for building agents with very little code. Its CodeAgent acts by writing Python, while ToolCallingAgent uses structured tool calls. It works particularly well with Hugging Face models, Inference Providers, and open source models hosted locally.

The following example creates a web search agent, then hands it off to a manager agent that decides when to delegate the search.

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
The manager decides when it needs up-to-date information, delegates the search to the specialized web agent, then produces the final answer.

## Quick comparison of frameworks

The table below compares the main strength of each framework and the type of agentic AI application it is best suited for.

| Framework | Main strength | Choose it when | 
| LangChain | Large ecosystem of integrations | You need extensive integrations, tools, and flexibility | 
| LangGraph | Stateful graph orchestration | You need controlled loops, durable executions, and complex workflows | 
| Agno | Complete agent platform | You want to build, self-host, and manage agent services | 
| Pydantic AI | Type-safe Python development | You need validated tools, dependencies, and structured outputs | 
| OpenAI Agents SDK | Simplified agent development and tracing | You mainly use OpenAI models and want to integrate quickly | 
| Google ADK | Native development for Gemini | You use Gemini, Google Cloud, or multi-agent workflows | 
| Claude Agent SDK | Tools for computer, files, and commands | You are designing code, research, or long-running agents with Claude | 
| CrewAI | Role-based agent teams | Multiple specialized agents need to collaborate on a project | 
| MetaGPT | Software company simulation | You want agents to plan and generate complete software projects | 
| AgentScope | Observable and controllable agents | You need traceable agents or multi-agent applications | 
| CAMEL-AI | Multi-agent research and simulation | You are studying role-play, agent societies, or large agent systems | 
| LlamaIndex | Data-connected and context-aware applications | Your agents need to retrieve and reason over private documents | 
| Haystack | Modular and transparent RAG pipelines | You want to control retrieval, prompting, and data flow | 
| smolagents | Lightweight agents for open models | You want minimal abstractions, code agents, or support for local models | 

## Final tips

If you are new to agentic AI, start with an official SDK like the OpenAI Agents SDK, Google ADK, or the Claude Agent SDK. They work closely with their models and APIs, so getting started is usually simple. Add an API key, credit your account, and start building.

I appreciate these SDKs because they make it easy to create agents, connect tools, and inspect what happened during an execution. With a few lines of Python, you can create a useful agent, or even a small multi-agent system.

When I need more flexibility, I usually turn to LangChain or LangGraph. LangChain is useful when I need many integrations, while LangGraph gives me more control over state, loops, and long workflows.

```

For teams of agents with varied roles, CrewAI is a good option. For RAG, private documents, and data-connected applications, I prefer LlamaIndex or Haystack.

Finally, smolagents is an excellent choice when I want something lightweight, simple, and easy to use with open or local models.

My advice is to start with the simplest framework that meets your needs. Create a minimal version, test it properly, and only add complexity when it's truly necessary.

## FAQs

### How much does running an AI agent cost compared to a standard LLM prompt?

While the Python frameworks themselves are open source and free, running agents can cost significantly more than standard LLM API calls. Agents work in a loop, constantly sending the model the system prompt, tool descriptions, previous tool outputs, and reasoning steps (like ReAct): token consumption adds up very quickly. A single user request can trigger five or six LLM calls before the agent arrives at a final answer. To control costs, developers often use smaller, cheaper models for simple routing tasks, and reserve more powerful models for complex reasoning.

### How do I evaluate whether my agent is actually working well?

Testing non-deterministic agents requires more than classic unit tests. Developers turn to specialized "LLM-as-a-judge" evaluation frameworks, such as **Ragas**, **TruLens**, or **DeepEval**. These tools run your agent on a set of test questions and score the outputs according to specific metrics, for example:

- **Factual grounding:** did the agent hallucinate, or is the answer based solely on the retrieved context?
- **Tool selection accuracy:** did the agent choose the appropriate tool?
- **Answer relevance:** does the final answer actually address the user's request?

### What is the difference between a standard RAG pipeline (LlamaIndex/Haystack) and an "Agentic" RAG system?

A standard RAG pipeline follows a static, hard-coded path: it takes the user query, queries a vector database, injects the results into a prompt, and generates an answer. It runs once and then stops. An **Agentic RAG** system gives the LLM autonomy over the retrieval process. The agent can decide *whether* or not to search, generate multiple different search queries to broaden the context, evaluate the usefulness of the retrieved documents, and choose to search again if the information is incomplete, before finally answering the user.

As a certified data scientist, I am passionate about using cutting-edge technologies to create innovative machine learning applications. With a strong background in speech recognition, data analysis and reporting, MLOps, conversational AI, and NLP, I have honed my skills in developing intelligent systems that can have a real impact. In addition to my technical expertise, I am also a skilled communicator, adept at distilling complex concepts into clear and concise language. As a result, I have become a sought-after blogger in the field of data science, sharing my ideas and experiences with a growing community of data professionals. Currently, I focus on content creation and editing, working with large language models to develop powerful and engaging content that can help businesses and individuals get the most out of their data.
