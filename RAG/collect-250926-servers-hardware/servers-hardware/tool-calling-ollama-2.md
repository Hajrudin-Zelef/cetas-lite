---
id: collect-250926-servers-hardware/servers-hardware/tool-calling-ollama-2
title: "with pip"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "parameters", "tool calling"]
source: docs/RAG/clean4/tool-calling-ollama.md
source_anchor: ""
source_lines: [261, 516]
sha256: 045b9296559e2f0d2625e46a48593417916d69ce152bd847e3ce3feb715ee5f9
---

# with pip

  Args:
    city: The name of the city
  Returns:
    The current weather conditions for the city
  """
  conditions = {
    "New York": "Partly cloudy",
    "London": "Rainy",
    "Tokyo": "Sunny"
  }
  return conditions.get(city, "Unknown")
messages = [{'role': 'user', 'content': 'What are the current weather conditions and temperature in New York and London?'}]
# The python client automatically parses functions as a tool schema so we can pass them directly
# Schemas can be passed directly in the tools list as well 
response = chat(model='qwen3', messages=messages, tools=[get_temperature, get_conditions], think=True)
# add the assistant message to the messages
messages.append(response.message)
if response.message.tool_calls:
  # process each tool call 
  for call in response.message.tool_calls:
    # execute the appropriate tool
    if call.function.name == 'get_temperature':
      result = get_temperature(**call.function.arguments)
    elif call.function.name == 'get_conditions':
      result = get_conditions(**call.function.arguments)
    else:
      result = 'Unknown tool'
    # add the tool result to the messages
    messages.append({'role': 'tool',  'tool_name': call.function.name, 'content': str(result)})
  # generate the final response
  final_response = chat(model='qwen3', messages=messages, tools=[get_temperature, get_conditions], think=True)
  print(final_response.message.content)
```
```
import ollama from 'ollama'
function getTemperature(city: string): string {
  const temperatures: { [key: string]: string } = {
    "New York": "22°C",
    "London": "15°C",
    "Tokyo": "18°C"
  }
  return temperatures[city] || "Unknown"
}
function getConditions(city: string): string {
  const conditions: { [key: string]: string } = {
    "New York": "Partly cloudy",
    "London": "Rainy",
    "Tokyo": "Sunny"
  }
  return conditions[city] || "Unknown"
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
  {
    type: 'function',
    function: {
      name: 'get_conditions',
      description: 'Get the current weather conditions for a city',
      parameters: {
        type: 'object',
        required: ['city'],
        properties: {
          city: { type: 'string', description: 'The name of the city' },
        },
      },
    },
  }
]
const messages = [{ role: 'user', content: 'What are the current weather conditions and temperature in New York and London?' }]
const response = await ollama.chat({
  model: 'qwen3',
  messages,
  tools,
  think: true
})
// add the assistant message to the messages
messages.push(response.message)
if (response.message.tool_calls) {
  // process each tool call 
  for (const call of response.message.tool_calls) {
    // execute the appropriate tool
    let result: string
    if (call.function.name === 'get_temperature') {
      const args = call.function.arguments as { city: string }
      result = getTemperature(args.city)
    } else if (call.function.name === 'get_conditions') {
      const args = call.function.arguments as { city: string }
      result = getConditions(args.city)
    } else {
      result = 'Unknown tool'
    }
    // add the tool result to the messages
    messages.push({ role: 'tool', tool_name: call.function.name, content: result })
  }
  // generate the final response
  const finalResponse = await ollama.chat({ model: 'qwen3', messages, tools, think: true })
  console.log(finalResponse.message.content)
}
```
## Multi-turn tool calling (Agent loop)

An agent loop allows the model to decide when to invoke tools and incorporate their results into its replies. It also might help to tell the model that it is in a loop and can make multiple tool calls.
- Python
- JavaScript

```
from ollama import chat, ChatResponse
def add(a: int, b: int) -> int:
  """Add two numbers"""
  """
  Args:
    a: The first number
    b: The second number
  Returns:
    The sum of the two numbers
  """
  return a + b
def multiply(a: int, b: int) -> int:
  """Multiply two numbers"""
  """
  Args:
    a: The first number
    b: The second number
  Returns:
    The product of the two numbers
  """
  return a * b
available_functions = {
  'add': add,
  'multiply': multiply,
}
messages = [{'role': 'user', 'content': 'What is (11434+12341)*412?'}]
while True:
    response: ChatResponse = chat(
        model='qwen3',
        messages=messages,
        tools=[add, multiply],
        think=True,
    )
    messages.append(response.message)
    print("Thinking: ", response.message.thinking)
    print("Content: ", response.message.content)
    if response.message.tool_calls:
        for tc in response.message.tool_calls:
            if tc.function.name in available_functions:
                print(f"Calling {tc.function.name} with arguments {tc.function.arguments}")
                result = available_functions[tc.function.name](**tc.function.arguments)
                print(f"Result: {result}")
                # add the tool result to the messages
                messages.append({'role': 'tool', 'tool_name': tc.function.name, 'content': str(result)})
    else:
        # end the loop when there are no more tool calls
        break
  # continue the loop with the updated messages
```
```
import ollama from 'ollama'
type ToolName = 'add' | 'multiply'
function add(a: number, b: number): number {
  return a + b
}
function multiply(a: number, b: number): number {
  return a * b
}
const availableFunctions: Record<ToolName, (a: number, b: number) => number> = {
  add,
  multiply,
}
const tools = [
  {
    type: 'function',
    function: {
      name: 'add',
      description: 'Add two numbers',
      parameters: {
        type: 'object',
        required: ['a', 'b'],
        properties: {
          a: { type: 'integer', description: 'The first number' },
          b: { type: 'integer', description: 'The second number' },
        },
      },
    },
  },
  {
    type: 'function',
    function: {
      name: 'multiply',
      description: 'Multiply two numbers',
      parameters: {
        type: 'object',
        required: ['a', 'b'],
        properties: {
          a: { type: 'integer', description: 'The first number' },
          b: { type: 'integer', description: 'The second number' },
        },
      },
    },
  },
]
async function agentLoop() {
  const messages = [{ role: 'user', content: 'What is (11434+12341)*412?' }]
  while (true) {
    const response = await ollama.chat({
      model: 'qwen3',
      messages,
      tools,
      think: true,
    })
    messages.push(response.message)
    console.log('Thinking:', response.message.thinking)
    console.log('Content:', response.message.content)
    const toolCalls = response.message.tool_calls ?? []
    if (toolCalls.length) {
      for (const call of toolCalls) {
        const fn = availableFunctions[call.function.name as ToolName]
        if (!fn) {
          continue
        }
        const args = call.function.arguments as { a: number; b: number }
        console.log(`Calling ${call.function.name} with arguments`, args)
        const result = fn(args.a, args.b)
        console.log(`Result: ${result}`)
        messages.push({ role: 'tool', tool_name: call.function.name, content: String(result) })
      }
    } else {
      break
    }
  }
}
agentLoop().catch(console.error)
```
## Tool calling with streaming

When streaming, gather every chunk of`thinking`, `content`, and `tool_calls`, then return those fields together with any tool results in the follow-up request.
- Python
- JavaScript

```
from ollama import chat 
def get_temperature(city: str) -> str:
  """Get the current temperature for a city
  
