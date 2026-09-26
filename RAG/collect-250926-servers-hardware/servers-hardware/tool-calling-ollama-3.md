---
id: collect-250926-servers-hardware/servers-hardware/tool-calling-ollama-3
title: "with pip"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/clean4/tool-calling-ollama.md
source_anchor: ""
source_lines: [517, 662]
sha256: b349fa4528d4fc842b83e295341006371f5630a9cff278bf2959112f5bb9e779
---

# with pip

  Args:
    city: The name of the city
  Returns:
    The current temperature for the city
  """
  temperatures = {
    'New York': '22°C',
    'London': '15°C',
  }
  return temperatures.get(city, 'Unknown')
messages = [{'role': 'user', 'content': "What is the temperature in New York?"}]
while True:
  stream = chat(
    model='qwen3',
    messages=messages,
    tools=[get_temperature],
    stream=True,
    think=True,
  )
  thinking = ''
  content = ''
  tool_calls = []
  done_thinking = False
  # accumulate the partial fields
  for chunk in stream:
    if chunk.message.thinking:
      thinking += chunk.message.thinking
      print(chunk.message.thinking, end='', flush=True)
    if chunk.message.content:
      if not done_thinking:
        done_thinking = True
        print('\n')
      content += chunk.message.content
      print(chunk.message.content, end='', flush=True)
    if chunk.message.tool_calls:
      tool_calls.extend(chunk.message.tool_calls)
      print(chunk.message.tool_calls)
  # append accumulated fields to the messages
  if thinking or content or tool_calls:
    messages.append({'role': 'assistant', 'thinking': thinking, 'content': content, 'tool_calls': tool_calls})
  if not tool_calls:
    break
  for call in tool_calls:
    if call.function.name == 'get_temperature':
      result = get_temperature(**call.function.arguments)
    else:
      result = 'Unknown tool'
    messages.append({'role': 'tool', 'tool_name': call.function.name, 'content': result})
```
```
import ollama from 'ollama'
function getTemperature(city: string): string {
  const temperatures: Record<string, string> = {
    'New York': '22°C',
    'London': '15°C',
  }
  return temperatures[city] ?? 'Unknown'
}
const getTemperatureTool = {
  type: 'function',
  function: {
    name: 'get_temperature',
    description: 'Get the current temperature for a city',
    parameters: {
      type: 'object',
      required: ['city'],
      properties: {
        city: { type: 'string', description: 'The name of the city' },
      },
    },
  },
}
async function agentLoop() {
  const messages = [{ role: 'user', content: "What is the temperature in New York?" }]
  while (true) {
    const stream = await ollama.chat({
      model: 'qwen3',
      messages,
      tools: [getTemperatureTool],
      stream: true,
      think: true,
    })
    let thinking = ''
    let content = ''
    const toolCalls: any[] = []
    let doneThinking = false
    for await (const chunk of stream) {
      if (chunk.message.thinking) {
        thinking += chunk.message.thinking
        process.stdout.write(chunk.message.thinking)
      }
      if (chunk.message.content) {
        if (!doneThinking) {
          doneThinking = true
          process.stdout.write('\n')
        }
        content += chunk.message.content
        process.stdout.write(chunk.message.content)
      }
      if (chunk.message.tool_calls?.length) {
        toolCalls.push(...chunk.message.tool_calls)
        console.log(chunk.message.tool_calls)
      }
    }
    if (thinking || content || toolCalls.length) {
      messages.push({ role: 'assistant', thinking, content, tool_calls: toolCalls } as any)
    }
    if (!toolCalls.length) {
      break
    }
    for (const call of toolCalls) {
      if (call.function.name === 'get_temperature') {
        const args = call.function.arguments as { city: string }
        const result = getTemperature(args.city)
        messages.push({ role: 'tool', tool_name: call.function.name, content: result } )
      } else {
        messages.push({ role: 'tool', tool_name: call.function.name, content: 'Unknown tool' } )
      }
    }
  }
}
agentLoop().catch(console.error)
```
## Using functions as tools with Ollama Python SDK

The Python SDK automatically parses functions as a tool schema so we can pass them directly. Schemas can still be passed if needed.```
from ollama import chat
def get_temperature(city: str) -> str:
  """Get the current temperature for a city
  
  Args:
    city: The name of the city
  Returns:
    The current temperature for the city
  """
  temperatures = {
    'New York': '22°C',
    'London': '15°C',
  }
  return temperatures.get(city, 'Unknown')
available_functions = {
  'get_temperature': get_temperature,
}
# directly pass the function as part of the tools list
response = chat(model='qwen3', messages=messages, tools=available_functions.values(), think=True)
```
