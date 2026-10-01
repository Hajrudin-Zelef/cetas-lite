---
id: collect-261001-huawei/huawei/welcome-gemma-4-frontier-multimodal-intelligence-on-device-4
title: "Cette image montre un personnage anthropomorphe ressemblant à un lapin, vêtu d'un manteau bleu et d'un pantalon beige, se tenant sur un chemin de terre dans un paysage rural idyllique ..."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["agents", "decode", "fine-tuning", "inference", "parameters", "tool calling"]
source: docs/RAG/collect-261001-huawei/welcome-gemma-4-frontier-multimodal-intelligence-on-device.md
source_anchor: ""
source_lines: [247, 345]
sha256: 0510f022f79dfb2c125b0e9c94522519cb01ea4fde7bacf3bb1009ef27defaff
---

# Cette image montre un personnage anthropomorphe ressemblant à un lapin, vêtu d'un manteau bleu et d'un pantalon beige, se tenant sur un chemin de terre dans un paysage rural idyllique ...

```
import re 
WEATHER_TOOL = {
    "type": "function",
    "function": {
        "name": "get_weather",
        "description": "Gets the current weather for a specific location.",
        "parameters": {
            "type": "object",
            "properties": {
                "city": {"type": "string", "description": "The city name"},
            },
            "required": ["city"],
        },
    },
}
tools = [WEATHER_TOOL]
messages = [
    {"role": "user", "content": [
          {"type": "text", "text": "What is the city in this image? Check the weather there right now."},
        {"type": "image", "image": "https://huggingface.co/datasets/merve/vlm_test_images/resolve/main/thailand.jpg"},
    ]},
]
inputs = processor.apply_chat_template(
    messages,
    tools=[WEATHER_TOOL],
    tokenize=True,
    return_dict=True,
    return_tensors="pt",
    add_generation_prompt=True,
    enable_thinking=True,
).to(model.device)
output = model.generate(**inputs, max_new_tokens=1000)
input_len = inputs.input_ids.shape[-1]
generated_text_ids = output[0][input_len:]
generated_text = processor.decode(generated_text_ids, skip_special_tokens=True)
result = processor.parse_response(generated_text)
print(result["content"])
```
| Model | Output | 
|---|---|
| E2B | **Analyze the Request:** The user has provided an image and asked two questions:\n    *   "What is the city in this image?" *   "Check the weather there right now."\n\n2.**Analyze the Image:** The image displays a large, ornate, tall structure, characteristic of temple architecture, likely in Southeast Asia (given the style). It appears to be a famous landmark. 3.**Identify the Landmark/Location (Internal Knowledge/Visual Cues):** The structure strongly resembles the chedis/stupas found in places like Bangkok, Thailand (e.g., Wat Arun or similar structures).\n\n4.**Determine the City:** Based on the visual evidence, the location is highly likely to be Bangkok, Thailand.\n\n5.**Check Available Tools:** The available tool is`get_weather(city: STRING)` . 6.  **Formulate the Plan:**\n    *   Step 1: Identify the city from the image (Bangkok).\n    *   Step 2: Use the`get_weather` tool with the identified city.\n\n7.**Construct the Tool Call:**`get_weather(city="Bangkok")` \n\n8.**Generate Response:** The model should first state the identified city and then call the tool to get the weather. (Self-correction: The model*must* call the tool in the correct format if it needs external information.)call:get_weather{city:Bangkok} | 
| E4B | '1. **Analyze the Request:** The user has provided an image and asked two things:\n    a. Identify the city in the image.\n    b. Check the weather in that city right now.\n\n2.**Analyze the Image:** The image shows a very distinctive, ornate tower/spire structure, characteristic of Thai architecture (specifically, a chedi or prang).\n\n3.**Determine City Identification (Internal Knowledge/Visual Recognition):** Based on the architectural style, the structure strongly resembles prominent temples in Thailand, such as Wat Arun in Bangkok.\n\n4.**Evaluate Tool Availability:** I have one available tool:`get_weather(city: STRING)` . This tool requires a city name as input.\n\n5. **Formulate a Plan:**\n    a. I need the city name to use the`get_weather` tool.\n    b. I must first identify the city from the image.\n    c. Since my internal knowledge suggests Bangkok, I will proceed with that assumption, but I must acknowledge that I cannot*definitively* prove the location from the image alone.\n\n6.**Execute Step 1 (City Identification):** Assume the city is Bangkok based on the visual evidence of the architecture.\n\n7.**Execute Step 2 (Tool Call Generation):** Use the`get_weather` tool with "Bangkok" as the`city` argument.\n\n8.**Construct the Response:** Generate the appropriate tool call in JSON format.', 'answer': '<\|tool_call>call:get_weather{city:<\|"\|>Bangkok<\|"\|>}<tool_call\|>' | 
| 31B | The user wants to know the city in the image and the current weather there. 1. **Analyze the image:** The image shows ornate, spire-like structures (prangs) characteristic of Thai temple architecture. Specifically, these look like the prangs at Wat Phra Kaew (Temple of the Emerald Buddha) or Wat Pho in Bangkok, Thailand. The style is very distinctive of Bangkok's royal temple complexes. 2.**Identify the city:** Based on the architectural style, the city is Bangkok. 3.**Find the weather:** I need to call the`get_weather` tool with`city="Bangkok"` .call:get_weather{city:Bangkok} | 
| 26/A4B | The image shows several tall, ornate towers (prangs) typical of Thai Buddhist temple architecture. Specifically, these look like the structures within the Grand Palace or Wat Phra Kaew in Bangkok, Thailand. The style is unmistakably Thai. I'll confirm this by looking at the specific towers. These are very iconic.\n2. **Confirm the location:** The image shows elements of the Grand Palace complex in Bangkok.\n3.**Get weather information:** Once the city is identified as Bangkok, I will use the`get_weather` tool with`city='Bangkok'` .call:get_weather{city:Bangkok} | 

Gemma 4 comes with day-0 support for many open-source inference engines, and is ideal for tool calling and agents! We also release ONNX checkpoints that can run on many hardware backends, allowing use cases on edge devices or in browser!

Gemma 4 comes with first-class transformers support from the get-go 🤗. This integration allows using the model with other libraries like bitsandbytes, PEFT and TRL. Make sure to install the latest version of transformers.

```
pip install -U transformers
```
The easiest way to infer with the small Gemma 4 models is through the `any-to-any` pipeline. You can initialize it as follows.

```
from transformers import pipeline
pipe = pipeline("any-to-any", model="google/gemma-4-e2b-it")
```
You can then pass in images and text as follows.

```
messages = [
    {
        "role": "user",
        "content": [
            {
                "type": "image",
                "image": "https://huggingface.co/datasets/merve/vlm_test_images/resolve/main/thailand.jpg",
            },
            {"type": "text", "text": "Do you have travel advice going to here?"},
        ],
    }
]
output = pipe(messages, max_new_tokens=100, return_full_text=False)
output[0]["generated_text"]
# Based on the image, which appears to show a magnificent, ornate **Buddhist temple or pagoda**, likely in Southeast Asia (such as Thailand, Myanmar, or Cambodia), here is some general travel advice..
```
When inferring with videos, you can include the audio track using the `load_audio_from_video` argument.

```
messages = [
    {
        "role": "user",
        "content": [
            {
                "type": "video",
                "image": "https://huggingface.co/datasets/merve/vlm_test_images/resolve/main/rockets.mp4",
            },
            {"type": "text", "text": "What is happening in this video?"},
        ],
    }
]
pipe(messages, load_audio_from_video=True)
```
Going a level lower, you can load Gemma 4 using the `AutoModelForMultimodalLM` class, especially useful for fine-tuning. The built-in chat template takes care of formatting the inputs correctly, please make sure you use it to prevent subtle mistakes when building the prompt manually.

## Inference code

