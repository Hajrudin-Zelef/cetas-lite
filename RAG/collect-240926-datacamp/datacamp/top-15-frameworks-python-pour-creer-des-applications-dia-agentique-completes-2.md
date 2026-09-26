---
id: collect-240926-datacamp/datacamp/top-15-frameworks-python-pour-creer-des-applications-dia-agentique-completes-2
title: "pip install -U langchain \"langchain[openai]\""
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "Google", "OpenAI"]
dates: []
keywords: ["agent", "agentic", "agents", "claude", "gemini", "mcp", "memory", "open source", "research"]
source: docs/RAG/clean_en/datacamp/top-15-frameworks-python-pour-creer-des-applications-dia-agentique-completes.md
source_anchor: ""
source_lines: [153, 340]
sha256: bed1293bce779ddd639bc6fd808ddc87c6555e5342f44d99063ed8f496a3a20f
---

# pip install -U langchain "langchain[openai]"

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

