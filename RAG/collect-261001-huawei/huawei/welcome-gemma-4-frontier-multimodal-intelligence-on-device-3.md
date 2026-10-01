---
id: collect-261001-huawei/huawei/welcome-gemma-4-frontier-multimodal-intelligence-on-device-3
title: "Cette image montre un personnage anthropomorphe ressemblant à un lapin, vêtu d'un manteau bleu et d'un pantalon beige, se tenant sur un chemin de terre dans un paysage rural idyllique ..."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["decode", "inference", "training", "transcription"]
source: docs/RAG/collect-261001-huawei/welcome-gemma-4-frontier-multimodal-intelligence-on-device.md
source_anchor: ""
source_lines: [146, 246]
sha256: 1b7e7013c4ba445259d78b165fb9550542d6379c15f659f70cd39c1c5dc8df17
---

# Cette image montre un personnage anthropomorphe ressemblant à un lapin, vêtu d'un manteau bleu et d'un pantalon beige, se tenant sur un chemin de terre dans un paysage rural idyllique ...

```
messages = [
    {
        "role": "user",
        "content": [
            {"type": "image", "url": "https://huggingface.co/datasets/merve/vlm_test_images/resolve/main/bird.png"},
            {"type": "text", "text": "Write single detailed caption for this image."},
        ],
    },
]
inputs = processor.apply_chat_template(
    messages,
    tokenize=True,
    return_dict=True,
    return_tensors="pt",
    add_generation_prompt=True,
).to(model.device)
output = model.generate(**inputs, max_new_tokens=512)
input_len = inputs.input_ids.shape[-1]
generated_text_ids = output[0][input_len:]
generated_text = processor.decode(generated_text_ids, skip_special_tokens=True)
result = processor.parse_response(generated_text)
print(result["content"])
```
| Model | Output | 
|---|---|
| E2B ✅ | A medium shot captures a weathered seagull perched atop a stone pedestal in what appears to be a bustling European square, with a grand, classical-style building featuring ornate columns and architectural details dominating the right side of the frame. In the background, a low, light-colored building stretches across the middle ground, flanked by some green foliage, and hints of other structures, including a terracotta-roofed building, are visible. In the foreground, part of a modern bus stop or shelter is visible, with signage indicating routes, and the overall scene is bathed in bright daylight under a pale sky. | 
| E4B ✅ | A solitary seagull perches atop a grey stone pedestal in a bustling European piazza, with the historical grandeur of a classical building dominating the right side of the frame, contrasting with the more modern, low-rise architecture in the background under a bright, hazy sky. | 
| 26B/A4B ✅ | A full shot shows a white and brown seagull perched on a gray pillar in Venice, Italy. The seagull faces left, its beak pointed towards the edge of the frame. Behind it is a light gray, two-story building with several dark rectangular windows. To the right of the seagull is a large, white building with columns and intricate carvings. The sky is a clear, light blue. The lighting is bright and even, suggesting it is daytime. | 
| 31B ✅ | A high-angle, medium shot shows a white and gray seagull perched on a gray concrete pylon in Venice, Italy. The seagull is facing left, its body angled slightly towards the camera. Its white head, neck, and chest contrast with its gray wings and back. The seagull's orange beak is slightly open. The pylon is a light gray concrete cylinder with a weathered texture. In the background, there are several buildings. To the left is a long, low, white building with a series of small, rectangular windows. In the center is a cluster of green trees. To the right is a red building with a tiled roof and several small, rectangular windows. Further right is a large, ornate, white stone building with several columns and sculptures. The sky is a pale blue, and the light is bright and clear. In the lower-left corner, there is a yellow bus with the word "Ferrovia" printed on it in black letters. There are several people walking in the background, but they are small and out of focus. | 

These models are trained to answer questions about speech in audio. Music and non-speech sounds were not part of the training data.

## Inference code

```
messages = [
    {
        "role": "user",
        "content": [
            {"type": "audio", "url": "https://huggingface.co/datasets/hf-internal-testing/dummy-audio-samples/resolve/main/obama_first_45_secs.mp3"},
            {"type": "text", "text": "Can you describe this audio in detail?"},
        ],
    },
]
inputs = processor.apply_chat_template(
    messages,
    tokenize=True,
    return_dict=True,
    return_tensors="pt",
    add_generation_prompt=True,
).to(model.device)
output = model.generate(
    **inputs,
    max_new_tokens=1000,
    do_sample=False,
)
print(processor.decode(output[0], skip_special_tokens=True))
```
| Model | Output | 
|---|---|
| E2B | This audio is a personal reflection. The speaker is talking about their final farewell address to the nation, which they delivered in Chicago. They express gratitude for the conversations they've had with the American people, noting that despite not having met them face-to-face or even greeted them, these interactions in various settings like living rooms, schools, farms, factory floors, diners, and military outposts have been what has kept them going. | 
| E4B | The audio is a speech excerpt where a speaker is delivering a farewell address to the nation from Chicago. The speaker reflects on their time in office, expressing gratitude for the conversations they had with the American people across various settings like living rooms, schools, farms, factories, diners, and military outposts. The tone is reflective and appreciative, highlighting the importance of these interactions in their political journey. | 

Here is an example if you want to do transcription:

## Inference code

```
messages = [
    {
        "role": "user",
        "content": [
            {"type": "audio", "url": "https://huggingface.co/datasets/hf-internal-testing/dummy-audio-samples/resolve/main/obama_first_45_secs.mp3"},
            {"type": "text", "text": "Transcribe the audio?"},
        ],
    },
]
inputs = processor.apply_chat_template(
    messages,
    tokenize=True,
    return_dict=True,
    return_tensors="pt",
    add_generation_prompt=True,
).to(model.device)
output = model.generate(
    **inputs,
    max_new_tokens=1000,
    do_sample=False,
)
print(processor.decode(output[0], skip_special_tokens=True))
```
| Model | Output | 
|---|---|
| E2B | This week I traveled to Chicago to deliver my final farewell address to the nation following in the tradition of presidents before me It was an opportunity to say thank you whether we've seen eye to eye or rarely agreed at all my conversations with you the American people in living rooms and schools at farms and on factory floors at diners and on distant military outposts all these conversations are what have kept me honest | 
| E4B | This week I traveled to Chicago to deliver my final farewell address to the nation following in the tradition of presidents before me. It was an opportunity to say thank you. Whether we've seen eye to eye or rarely agreed at all, my conversations with you, the American people, in living rooms and schools, at farms and on factory floors, at diners and on distant military outposts, all these conversations are what have kept me honest. | 

We test the model by asking to get the weather in the place shown in the image.

## Inference code

