---
id: collect-240926-datacamp/datacamp/kimi-k3-fonctionnalites-benchmarks-api-et-5-exemples-pratiques-2
title: "kimi-k3-fonctionnalites-benchmarks-api-et-5-exemples-pratiques"
domain: datacamp
role: reference
task: reference
actors: ["Moonshot", "OpenAI", "United States"]
dates: []
keywords: ["kimi", "agent", "alignment", "cost", "decode", "license", "multimodal", "open weights", "parameters", "reasoning", "tool calling"]
source: docs/RAG/clean_en/datacamp/kimi-k3-fonctionnalites-benchmarks-api-et-5-exemples-pratiques.md
source_anchor: ""
source_lines: [133, 224]
sha256: 3f3ec57f94165751a26c8f6b66e8ecf97a3a3fd5a04de5129eba76705ca6508a
---

# kimi-k3-fonctionnalites-benchmarks-api-et-5-exemples-pratiques

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
