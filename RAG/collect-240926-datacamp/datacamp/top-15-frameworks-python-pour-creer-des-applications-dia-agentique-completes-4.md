---
id: collect-240926-datacamp/datacamp/top-15-frameworks-python-pour-creer-des-applications-dia-agentique-completes-4
title: "pip install -U langchain \"langchain[openai]\""
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "Google", "OpenAI"]
dates: []
keywords: ["agent", "agentic", "agents", "claude", "cost", "gemini", "open source", "reasoning", "research"]
source: docs/RAG/clean_en/datacamp/top-15-frameworks-python-pour-creer-des-applications-dia-agentique-completes.md
source_anchor: ""
source_lines: [536, 586]
sha256: 50ce09ed17827df0dc4ce55ebc691db5becbcd19b5e6c977708daf01b95d414b
---

# pip install -U langchain "langchain[openai]"

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
