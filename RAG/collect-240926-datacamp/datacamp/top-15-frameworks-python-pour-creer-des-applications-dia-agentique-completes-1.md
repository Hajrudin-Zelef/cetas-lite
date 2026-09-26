---
id: collect-240926-datacamp/datacamp/top-15-frameworks-python-pour-creer-des-applications-dia-agentique-completes-1
title: "pip install -U langchain \"langchain[openai]\""
domain: datacamp
role: reference
task: reference
actors: ["Google", "OpenAI"]
dates: []
keywords: ["agent", "agentic", "agents", "gpt-5.6", "guardrails", "memory", "research", "sol"]
source: docs/RAG/clean_en/datacamp/top-15-frameworks-python-pour-creer-des-applications-dia-agentique-completes.md
source_anchor: ""
source_lines: [1, 152]
sha256: bbdf64678e4afbe984cf8302d5fdb57d6e5cdfa5f0abf3aa043ed4df92eb9f9d
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

