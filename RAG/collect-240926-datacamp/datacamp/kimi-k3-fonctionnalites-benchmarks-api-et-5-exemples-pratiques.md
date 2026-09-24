---
id: collect-240926-datacamp/datacamp/kimi-k3-fonctionnalites-benchmarks-api-et-5-exemples-pratiques
title: "kimi-k3-fonctionnalites-benchmarks-api-et-5-exemples-pratiques"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "China", "Moonshot", "OpenAI", "United States"]
dates: ["2026-07-16", "2026-07-17"]
keywords: ["benchmark", "kimi", "agent", "alignment", "claude", "context window", "cost", "decode", "license", "multimodal", "open weights", "parameters"]
source: docs/RAG/clean_en/datacamp/kimi-k3-fonctionnalites-benchmarks-api-et-5-exemples-pratiques.md
source_anchor: ""
source_lines: [1, 224]
sha256: 53781fbb3dd35a340b4d25151ff5dde003a9561d8b76125f5e9beaf3ada25675
---

# kimi-k3-fonctionnalites-benchmarks-api-et-5-exemples-pratiques

<!-- source: https://www.datacamp.com/fr/tutorial/kimi-k3-tutorial -->

Cursus

The race for open models moved again on July 16, 2026, when Moonshot AI launched Kimi K3, a 2.8-trillion-parameter model with a one-million-token context window and native vision. It is the largest open model released by Moonshot, far beyond Kimi K2 in size, and the first they describe as belonging to the 3-trillion-parameter class.

If you are looking for the launch story, the architecture analysis, the benchmark charts, the comparisons with Claude, GPT, and the other Chinese labs, as well as the list of Moonshot's limitations, our blog post on Kimi K3 covers all of that. This tutorial is the practical counterpart: how to access it and how it behaves in use. I go through five small examples, four via the API where I show actual token usage and cost, and two in the kimi.com web app. Together, they show how K3 handles:

- Tool calling and returning strict JSON
- Loading a tool definition on the fly
- Reducing the cost of long contexts through automatic caching
- Reading a screenshot and fixing the layout
- Building an interactive dashboard from a single prompt

The four API examples were run on July 17, 2026 with the `kimi-k3` model and cost about 11 cents cold, or a few cents after caching was enabled.

## How to access Kimi K3

The fastest way to try the model is kimi.com, where the web app and mobile apps use Kimi K3 for general agent tasks with no setup.

For heavier work like reports and dashboards, there is Kimi Work, a desktop application.

If you live in the terminal, Kimi Code is a coding agent installed via npm under `@moonshot-ai/kimi-code`, and you choose the model with the `/model` command. Using K3 in Kimi Code requires a paid subscription, and the one-million-token window requires a higher tier.

This tutorial focuses on the raw API and the web app, but the command-line agent is available if you need it.

K3 does not replace its predecessors, however. The table below shows how the current lineup breaks down.

| **Model** | **Context window** | **Ideal for** | 
| `kimi-k3` | 1,48,576 tokens | Flagship work: long coding, vision, knowledge | 
| `kimi-k2.7-code` | 262,144 tokens | Dedicated coding, with a faster high-speed option | 
| `kimi-k2.6` | 262,144 tokens | General text, image, and video conversation | 

In short, K3 is the model to favor when a task mixes code, tools, documents, and images, or when you genuinely need the one-million-token window. For pure code generation where speed matters more than context, `kimi-k2.7-code` remains the smarter choice: do not assume the latest model is always the best for your case.

## Setting up the Kimi K3 API

The API is compatible with the OpenAI SDK, so if you have already used it, almost none of what follows will surprise you. You need Python 3.9 or higher and an API key.

### Step 1: generate an API key

Start by logging in to the Kimi platform and opening the API Keys page in the console. Create a key, copy it once, and store it somewhere safe, because you will not see it again. Also plan for a small credit on the account to make calls; a few dollars is more than enough for this entire tutorial.

*Creating a Kimi K3 API key. Image by the author.*

### Step 2: install the SDK

Next, install the OpenAI SDK in your environment. A single command is enough.

`python -m pip install --upgrade "openai>=1.0"`
This downloads the client library used by the rest of the examples, with no Kimi-specific installation.

### Step 3: store the key and initialize the client

It is better to read the key from an environment variable than to paste it into your code. Set `MOONSHOT_API_KEY` in your shell or a `.env` file, then point the client to Moonshot's base URL.

```
import os
from openai import OpenAI
client = OpenAI(
    api_key=os.environ["MOONSHOT_API_KEY"],
    base_url="https://api.moonshot.ai/v1",
)
```
The only two differences from a standard OpenAI setup are the `base_url` and the model name, namely `kimi-k3`. With that in place, you are ready to make a first call.

### Step 4: make your first call

Let's move on to a first request. I asked the model to introduce itself in one sentence, which produced a small moment of candor.

```
completion = client.chat.completions.create(
    model="kimi-k3",
    messages=[{"role": "user", "content": "Introduce Kimi K3 in one sentence."}],
    max_completion_tokens=800,
)
print(completion.choices[0].message.content)
```
The response was a polite refusal to guess: the model said it did not have reliable information about Kimi K3, having been trained before its own release, and pointed me to Moonshot's announcements. A useful reminder: a model does not know itself. The API call I just made costs about seven tenths of a cent. Note the `max_completion_tokens` cap that I set on every call in this tutorial to prevent verbose outputs from driving up the bill.

*First response from the Kimi K3 API. Image by the author.*

## Example 1: streaming reasoning and final answer

K3 reasons systematically, and the API returns that reasoning on a channel separate from the answer. In streaming mode, each chunk can contain `reasoning_content`, the final `content`, or both, which lets you display the thinking and the answer separately.

```
stream = client.chat.completions.create(
    model="kimi-k3",
    messages=[{"role": "user", "content": "A bat and a ball cost $1.10 together. The bat costs $1.00 more than the ball. How much is the ball?"}],
    max_completion_tokens=1200,
    stream=True,
    stream_options={"include_usage": True},
)
for chunk in stream:
    if not chunk.choices:
        continue
    delta = chunk.choices[0].delta
    reasoning = getattr(delta, "reasoning_content", None)
    if reasoning:
        print(reasoning, end="", flush=True)
    if delta.content:
        print(delta.content, end="", flush=True)
```
The model first streamed its reasoning: it recognized the bat-and-ball question as the classic Cognitive Reflection Test, flagged the intuitive wrong answer of $0.10, then set up the algebra to arrive at $0.05 and verify that $1.05 plus $0.05 does indeed equal $1.10. The separation is valuable: in a real application, you show the `content` to your users and keep the `reasoning_content` for the logs, because displaying raw reasoning in production is rarely desirable. This call used 488 output tokens and cost less than a cent.

*Streaming reasoning followed by the final answer. Image by the author.*

## Example 2: tool calling with structured output

Kimi K3 is the model in the lineup that supports `tool_choice="required"`, which forces at least one tool call during a turn. Handy when you want the model to fetch data before answering rather than guessing. Here, I gave it two fake tools, a price lookup and a stock check, forced a tool call, ran the tools locally, then requested the result in strict JSON via `response_format`.

```
first = client.chat.completions.create(
    model="kimi-k3",
    messages=messages,
    tools=TOOLS,
    tool_choice="required",
    max_completion_tokens=2500,
)
assistant_message = first.choices[0].message
messages.append(assistant_message)
for tool_call in assistant_message.tool_calls or []:
    args = json.loads(tool_call.function.arguments)
    messages.append({"role": "tool", "tool_call_id": tool_call.id, "content": run_tool(tool_call.function.name, args)})
```
The model called both tools with the correct product code, then returned a clean order summary in JSON: five mechanical keyboards at $89 each, a total of $445, and a stock indicator set to true. Two details make the difference in practice: you must reinsert the full assistant message into the conversation before adding the tool results, and you should only parse the `content` for the JSON, never the reasoning field. The pair of calls cost less than a cent in total.

Tool calls and structured JSON output. Image by the author.

## Example 3: loading tools dynamically

If you have dozens of tools, sending all their definitions with every request wastes tokens and clutters the prompt. Kimi K3 lets you inject a tool definition mid-conversation via a `system` message that contains a `tools` field and no `content`. The tool becomes available from that point on, which keeps large tool catalogs out of your cached prefix until a tool is actually needed.

```
messages = [
    {"role": "user", "content": "Convert 100 US dollars to euros at a rate of 0.92."},
    {"role": "system", "tools": [{
        "type": "function",
        "function": {
            "name": "convert_currency",
            "description": "Convert an amount from one currency to another",
            "parameters": {
                "type": "object",
                "properties": {"amount": {"type": "number"}, "rate": {"type": "number"}},
                "required": ["amount", "rate"],
            },
        },
    }]},
]
completion = client.chat.completions.create(model="kimi-k3", messages=messages)
print(completion.choices[0].message.tool_calls)
```
K3 took the freshly loaded tool into account and called `convert_currency` with an amount of 100 and a rate of 0.92, as expected. Keep in mind that the server does not retain that definition for you: resend the system message in subsequent requests if you want the tool to remain available. This was the cheapest call in the series, at about two-tenths of a cent.


Calling a dynamic currency conversion tool. Image by the author.

## Example 4: reducing the cost of long contexts with caching

This is where the one-million-token window becomes truly practical. Context caching is automatic: no cache ID or time-to-live to manage. You send a long prefix, keep it strictly identical across subsequent requests, and the repeated part is billed at the cache hit rate rather than the cache miss rate. To make the gap visible, I used a knowledge base of about 33,000 tokens and asked a question about it.

```
knowledge = Path("knowledge_base.md").read_text(encoding="utf-8")
completion = client.chat.completions.create(
    model="kimi-k3",
    messages=[
        {"role": "system", "content": knowledge},
        {"role": "user", "content": "What is the rated payload of the Atlas robot?"},
    ],
    max_completion_tokens=600,
)
```
The first time, nothing was cached and the request cost about 9.9 cents for some 33,000 input tokens. Once the prefix had been seen, the same request hit the cache on all 32,512 prefix tokens and cost about 1.1 cents, nearly a factor of nine. The reason: the price gap; cached input is billed at $0.30 per million tokens versus $3.00 uncached. One point to note: cache writes are asynchronous, so the hit does not show up on an immediately consecutive call. It appears on a later request; running the script twice a minute apart shows the miss first, then the hit.

Cache miss vs cache hit cost. Image by the author.

## Example 5: Detecting layout bugs in a screenshot

Vision is native in K3, and the API offers a simple way to use it, even though it doesn't accept public image URLs. You send the image as a base64 data URL and make `content` an array of objects, one part for the image, one for the text. I rendered a small dashboard with a few deliberate layout bugs, saved a screenshot, and asked K3 what was wrong.

*The dashboard with deliberate layout bugs. Image by the author.*

```
import base64
from pathlib import Path
image_data = base64.b64encode(Path("broken_dashboard.png").read_bytes()).decode()
completion = client.chat.completions.create(
    model="kimi-k3",
    messages=[{
        "role": "user",
        "content": [
            {"type": "image_url", "image_url": {"url": f"data:image/png;base64,{image_data}"}},
            {"type": "text", "text": "List the layout and alignment problems you can see, and give a short CSS fix for each."},
        ],
    }],
    max_completion_tokens=3500,
)
print(completion.choices[0].message.content)
```
K3 read the image well. It spotted the card lower than the row and overlapping its neighbor, the badge sitting on a number (it even misread the 3,910 hidden as 5,910, proof of the bug), the irregular spacing before the last card, the bar overflowing toward the card above, as well as the tooltip sitting on the bars, and suggested a short CSS fix for each, such as grouping the cards into a grid. On the other hand, it missed the very low-contrast subtitle: vision catches what jumps out more than subtle details. The call cost about two hundred.

## Limitations of Kimi K3

The API examples went well, but a few rough edges are worth pointing out to avoid surprises. I ran into most of them directly.

- 
Only `reasoning_effort="max"` is available for now, so you can't yet reduce reasoning to save money.
- 
Sampling parameters are frozen. Values like `temperature`, `top_p`, and penalties are locked; don't include them in requests instead of trying to adjust them.
- 
Output can become long and expensive. Limit `max_completion_tokens`, as in the examples, and validate any agent loop.
- 
Public image URLs are not supported via the API; plan for base64 or uploaded files for vision.

None of this is blocking, but these points influence how the model is used. Output cost is the one I would monitor first.

## Conclusion

Over the course of my tests, two things stand out. Tool calling and structured output worked without retries, and caching mattered more than expected: reusing the same long prefix made repeating a large request inexpensive. For repository-scale analysis, repeated long-context calls, or multimodal engineering, K3 is a reasonable default choice; for fast, low-cost chat or fine-grained control of sampling, a smaller model will be simpler. Details about the open weights and license, which I noted above, should be clarified after the July 27 release.

To go deeper into the patterns used in these examples, our Developing AI Systems with the OpenAI API course covers function calling and connecting models to external tools in Python.

I'm a data engineer and community builder. I work on data pipelines, cloud, and AI tools, while writing practical, impactful tutorials for DataCamp and emerging developers.
