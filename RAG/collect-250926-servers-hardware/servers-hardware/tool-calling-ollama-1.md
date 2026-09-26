---
id: collect-250926-servers-hardware/servers-hardware/tool-calling-ollama-1
title: "with pip"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters", "tool calling"]
source: docs/RAG/clean4/tool-calling-ollama.md
source_anchor: ""
source_lines: [1, 260]
sha256: 00de38e6fd833d963ec61b49ae59192728fe53d5466a75b739c1230adf72332d
---

# with pip

## Calling a single tool

Invoke a single tool and include its response in a follow-up request. Also known as “single-shot” tool calling.
- cURL
- Python
- JavaScript

```
curl -s http://localhost:11434/api/chat -H "Content-Type: application/json" -d '{
  "model": "qwen3",
  "messages": [{"role": "user", "content": "What is the temperature in New York?"}],
  "stream": false,
  "tools": [
    {
      "type": "function",
      "function": {
        "name": "get_temperature",
        "description": "Get the current temperature for a city",
        "parameters": {
          "type": "object",
          "required": ["city"],
          "properties": {
            "city": {"type": "string", "description": "The name of the city"}
          }
        }
      }
    }
  ]
}'
```
**Generate a response with a single tool result**

```
curl -s http://localhost:11434/api/chat -H "Content-Type: application/json" -d '{
  "model": "qwen3",
  "messages": [
    {"role": "user", "content": "What is the temperature in New York?"},
    {
      "role": "assistant",
      "tool_calls": [
        {
          "type": "function",
          "function": {
            "index": 0,
            "name": "get_temperature",
            "arguments": {"city": "New York"}
          }
        }
      ]
    },
    {"role": "tool", "tool_name": "get_temperature", "content": "22°C"}
  ],
  "stream": false
}'
```
Install the Ollama Python SDK:

```
# with pip
pip install ollama -U
# with uv
uv add ollama    
```
```
from ollama import chat
def get_temperature(city: str) -> str:
  """Get the current temperature for a city
  
  Args:
    city: The name of the city
  Returns:
    The current temperature for the city
  """
  temperatures = {
    "New York": "22°C",
    "London": "15°C",
    "Tokyo": "18°C",
  }
  return temperatures.get(city, "Unknown")
messages = [{"role": "user", "content": "What is the temperature in New York?"}]
# pass functions directly as tools in the tools list or as a JSON schema
response = chat(model="qwen3", messages=messages, tools=[get_temperature], think=True)
messages.append(response.message)
if response.message.tool_calls:
  # only recommended for models which only return a single tool call
  call = response.message.tool_calls[0]
  result = get_temperature(**call.function.arguments)
  # add the tool result to the messages
  messages.append({"role": "tool", "tool_name": call.function.name, "content": str(result)})
  final_response = chat(model="qwen3", messages=messages, tools=[get_temperature], think=True)
  print(final_response.message.content)
```
Install the Ollama JavaScript library:

```
# with npm
npm i ollama
# with bun
bun i ollama
```
```
import ollama from 'ollama'
function getTemperature(city: string): string {
  const temperatures: Record<string, string> = {
    'New York': '22°C',
    'London': '15°C',
    'Tokyo': '18°C',
  }
  return temperatures[city] ?? 'Unknown'
}
const tools = [
  {
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
  },
]
const messages = [{ role: 'user', content: "What is the temperature in New York?" }]
const response = await ollama.chat({
  model: 'qwen3',
  messages,
  tools,
  think: true,
})
messages.push(response.message)
if (response.message.tool_calls?.length) {
  // only recommended for models which only return a single tool call
  const call = response.message.tool_calls[0]
  const args = call.function.arguments as { city: string }
  const result = getTemperature(args.city)
  // add the tool result to the messages
  messages.push({ role: 'tool', tool_name: call.function.name, content: result })
  // generate the final response
  const finalResponse = await ollama.chat({ model: 'qwen3', messages, tools, think: true })
  console.log(finalResponse.message.content)
}
```
## Parallel tool calling

- cURL
- Python
- JavaScript

Request multiple tool calls in parallel, then send all tool responses back to the model.

```
curl -s http://localhost:11434/api/chat -H "Content-Type: application/json" -d '{
  "model": "qwen3",
  "messages": [{"role": "user", "content": "What are the current weather conditions and temperature in New York and London?"}],
  "stream": false,
  "tools": [
    {
      "type": "function",
      "function": {
        "name": "get_temperature",
        "description": "Get the current temperature for a city",
        "parameters": {
          "type": "object",
          "required": ["city"],
          "properties": {
            "city": {"type": "string", "description": "The name of the city"}
          }
        }
      }
    },
    {
      "type": "function",
      "function": {
        "name": "get_conditions",
        "description": "Get the current weather conditions for a city",
        "parameters": {
          "type": "object",
          "required": ["city"],
          "properties": {
            "city": {"type": "string", "description": "The name of the city"}
          }
        }
      }
    }
  ]
}'
```
**Generate a response with multiple tool results**```
curl -s http://localhost:11434/api/chat -H "Content-Type: application/json" -d '{
  "model": "qwen3",
  "messages": [
    {"role": "user", "content": "What are the current weather conditions and temperature in New York and London?"},
    {
      "role": "assistant",
      "tool_calls": [
        {
          "type": "function",
          "function": {
            "index": 0,
            "name": "get_temperature",
            "arguments": {"city": "New York"}
          }
        },
        {
          "type": "function",
          "function": {
            "index": 1,
            "name": "get_conditions",
            "arguments": {"city": "New York"}
          }
        },
        {
          "type": "function",
          "function": {
            "index": 2,
            "name": "get_temperature",
            "arguments": {"city": "London"}
          }
        },
        {
          "type": "function",
          "function": {
            "index": 3,
            "name": "get_conditions",
            "arguments": {"city": "London"}
          }
        }
      ]
    },
    {"role": "tool", "tool_name": "get_temperature", "content": "22°C"},
    {"role": "tool", "tool_name": "get_conditions", "content": "Partly cloudy"},
    {"role": "tool", "tool_name": "get_temperature", "content": "15°C"},
    {"role": "tool", "tool_name": "get_conditions", "content": "Rainy"}
  ],
  "stream": false
}'
```
```
from ollama import chat
def get_temperature(city: str) -> str:
  """Get the current temperature for a city
  
  Args:
    city: The name of the city
  Returns:
    The current temperature for the city
  """
  temperatures = {
    "New York": "22°C",
    "London": "15°C",
    "Tokyo": "18°C"
  }
  return temperatures.get(city, "Unknown")
def get_conditions(city: str) -> str:
  """Get the current weather conditions for a city
  
