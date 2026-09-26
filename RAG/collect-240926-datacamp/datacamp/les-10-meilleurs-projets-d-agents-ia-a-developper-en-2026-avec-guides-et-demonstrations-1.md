---
id: collect-240926-datacamp/datacamp/les-10-meilleurs-projets-d-agents-ia-a-developper-en-2026-avec-guides-et-demonstrations-1
title: "les-10-meilleurs-projets-d-agents-ia-a-developper-en-2026-avec-guides-et-demonstrations"
domain: datacamp
role: reference
task: reference
actors: ["Alibaba", "Google", "Mistral", "OpenAI"]
dates: []
keywords: ["agent", "agents", "agentic", "gguf", "inference", "llama", "memory", "mistral", "qwen", "reasoning", "research"]
source: docs/RAG/clean_en/datacamp/les-10-meilleurs-projets-d-agents-ia-a-developper-en-2026-avec-guides-et-demonstrations.md
source_anchor: ""
source_lines: [1, 102]
sha256: cfe01b7a439935356f1d4f1ab44d8620b0a3066dc316e6ca634107eedc765b8d
---

# les-10-meilleurs-projets-d-agents-ia-a-developper-en-2026-avec-guides-et-demonstrations

<!-- source: https://www.datacamp.com/fr/blog/top-ai-agent-projects -->

Course

AI agents are advanced, goal-oriented systems that perceive context and execute multi-step workflows with minimal supervision, unlike basic chatbots. They can call APIs, use software, query data, and leverage memory and feedback loops.

These agents have become a central part of the AI ecosystem, and knowing about them will help you start your career in this field or improve your current skills. The best way to understand a new concept is to put it into practice.

In this article, we will look at 10 AI agent projects useful for all skill levels:

- Beginners: Build quickly using low-code tools like Langflow, Flowise, and Make AI.
- Intermediate: Use frameworks like LangGraph, Mistral Agents, and Qwen-Agent to build custom agents.
- Advanced: Design multi-agent systems with Haystack, ADK, and CrewAI.

## AI Agent Projects for Beginners

Simple projects use GUI tools or low-code agent tools. You can drag and drop and visually connect components, add your LLM API key, and run the pipeline. These builders let you create stateful agent prototypes, integrate data sources and APIs, and chain prompts and tool calls without standard code.

### 1. Language Tutor with Langflow

Language Tutor with Langflow is a small agentic system that generates short reading passages tailored to the learner's current vocabulary. It runs on Langflow with a Postgres database (via Docker) and uses psycopg2 to read/write words.

You upload your vocabulary from a CSV file, add new words using a tool in the chat, and a storytelling tool retrieves the words you have saved to ask an LLM to write a story in the language of your choice. The main agent routes your request, "add a word" or "create a story," and returns the result.

Guide: Langflow: A guide with a demo project

### 2. Data Analyst AI Agent with Flowise

The Data Analyst AI Agent with Flowise is a workflow that lets you ask questions about a database and get answers with the exact SQL code used.

You connect Flowise to a SingleStore database, add a custom code block to read the table schema, then integrate it into a prompt that asks an LLM (via an LLM chain with OpenAI) to generate a SQL query.

The query is stored, cleaned, executed via another custom code block, and the results along with the query are formatted by a final prompt and LLM chain.

Guide: Flowise: A guide with a demo project

### 3. Customer Service Automation with Make AI

The Customer Service AI Agent with Make automatically responds to rental inquiries submitted via a Tally form.

When a person fills out a Tally form, Make AI retrieves the rental details from a Google Doc. Then an OpenAI module drafts a response using that information and the question asked by the person. Finally, the email module sends the response directly to the email address provided in the form.

Guide: Building an AI: A guide with practical examples

## Intermediate AI Agent Projects

Intermediate AI agent projects focus on building complete workflows using modern agent frameworks and APIs. At this stage, you are not just connecting APIs, you are also designing simple user interfaces so that users can interact directly with the agents.

### 4. Nutrition Coach with Mistral Agents

Nutrition Coach with Mistral Agents is an AI agent project that logs your meals, estimates calories, and suggests a next healthy meal, accompanied by an image.

The application includes the Mistral Agents API, which has a web search agent to obtain calorie estimates. It also has a fallback estimator, a logger to record meals, calories, and timestamps, and an image generation agent to visualize the suggested dish.

Users can enter their meals and preferences, and the application will search for or estimate the calorie content, log the entry, suggest a complementary meal, and display an automatically generated image of the dish, along with a clear summary of the tools used.

Guide: Mistral Agents API: A guide with a demo project

### 5. Deep Research Assistant with Jan-v1

The Deep Research Assistant AI agent project with Jan‑v1 is an application that turns a topic into a polished research report using local Jan‑v1 inference, asynchronous web search, and strict report formatting. It generates intelligent queries, extracts sources, and synthesizes a clear, professional report.

It uses Streamlit for the user interface, llama-cpp to run a GGUF Jan-v1 model locally, Serper for web search, and a set of helper functions for chunking and cleaning.

Users enter a topic, choose the level of detail, focus, time period, and format; the application generates queries, performs asynchronous searches, compiles notes, and asks Jan‑v1 to produce a structured report, then displays the sources and lets you export to TXT/JSON. Progress bars track the steps, and session state avoids reloading the model between runs.

Guide: January-V1: A guide with a demo project

### 6. Real-Time Web Summarizer Extension with Qwen-Agent

The Real-Time Web Summarizer extension is a Chrome add-on that captures the visible text of any page and streams a clear, concise summary in real time, powered locally by Qwen3 via Ollama and a FastAPI backend.

Users can click "Summarize" in the popup window, as shown below. The extension retrieves the page text, sends it to http://127.0.0.1:7864/summarize_stream_status, and displays the response continuously.

To set it up, you need to download the Qwen model using Ollama, start the FastAPI server, load the unpacked extension, and you can then get instant summaries on any page, presented in a clear, editor-like format.

Guide: Qwen Agent: A Guide with a Demo Project

### 7. Real-Time Analytics with LangGraph

Real-time Analytics is a LangGraph-powered assistant that answers questions, performs web searches, and executes Python code. It combines Mistral Medium 3 for reasoning, Tavily for web search, and a Python REPL to execute code based on user prompts.

When users ask a question, the agent determines whether to perform a search, execute code, or use both methods. It then provides the final answer along with information about the tools used. The setup process includes adding the Tavily API key, installing the necessary packages, initializing the LLM and tools, creating the LangGraph agent, and invoking it with user messages.

Guide: Mistral Medium 3 Tutorial: Developing Agentic Applications

## Advanced AI Agent Projects

Advanced AI agent projects bring together multiple agents working with various tools in a single workflow. These projects use powerful agent frameworks such as Haystack, ADK, and CrewAI, allowing you to design complex collaborative systems in which agents can coordinate, specialize, and accomplish more sophisticated tasks.

### 8. RAG and Web AI Agents with Haystack AI

Haystack Agentic RAG and Web Access is an assistant that answers questions using a private knowledge base and, if necessary, real-time web results. It first routes queries to retrieval-augmented generation, then switches to web search to obtain recent, real-time information.

It uses Haystack pipelines and agents, an in-memory document store with OpenAI integrations, a custom RAG tool, and a custom Tavily web search tool integrated as a ComponentTool. The agent is powered by GPT 4.1 Mini with a system prompt that guides tool selection.

Users can ask questions, and the agent retrieves information from the knowledge base or performs a Tavily search on current topics. It then provides an answer along with the tools used.

Guide: Haystack AI Tutorial: Building Agentic Workflows

### 9. Travel Planner with ADK and A2A

