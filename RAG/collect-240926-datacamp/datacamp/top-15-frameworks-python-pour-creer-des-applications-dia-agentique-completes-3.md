---
id: collect-240926-datacamp/datacamp/top-15-frameworks-python-pour-creer-des-applications-dia-agentique-completes-3
title: "pip install -U langchain \"langchain[openai]\""
domain: datacamp
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["agent", "agentic", "agents", "inference", "llama", "memory", "open source", "qwen", "research", "sglang", "vllm"]
source: docs/RAG/clean_en/datacamp/top-15-frameworks-python-pour-creer-des-applications-dia-agentique-completes.md
source_anchor: ""
source_lines: [341, 535]
sha256: 0ee7561283f5c4fe9aa8671738aba0a7c6df5accb433dac6a0cf4899958648f4
---

# pip install -U langchain "langchain[openai]"

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

