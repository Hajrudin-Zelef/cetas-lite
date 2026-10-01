---
id: collect-261001-huawei/huawei/welcome-gemma-4-frontier-multimodal-intelligence-on-device-2
title: "Cette image montre un personnage anthropomorphe ressemblant à un lapin, vêtu d'un manteau bleu et d'un pantalon beige, se tenant sur un chemin de terre dans un paysage rural idyllique ..."
domain: huawei
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["decode", "gemini", "inference"]
source: docs/RAG/collect-261001-huawei/welcome-gemma-4-frontier-multimodal-intelligence-on-device.md
source_anchor: ""
source_lines: [67, 145]
sha256: 9cf9732af9388c4342192909b098fea610c071eb910ae5ecd8f21d3dd166bfa8
---

# Cette image montre un personnage anthropomorphe ressemblant à un lapin, vêtu d'un manteau bleu et d'un pantalon beige, se tenant sur un chemin de terre dans un paysage rural idyllique ...

We test Gemma 4 on GUI element detection and pointing across different sizes, with the following image and text prompt: "What's the bounding box for the "view recipe" element in the image?"

With this prompt, the model natively responds in JSON format with the detected bounding boxes - no need for specific instructions or grammar-constrained generation. We found the coordinates refer to an image size of 1000x1000, relative to the input dimensions.

We visualize the outputs below for your convenience. We parse the bounding boxes from the returned JSON: `json\n[\n  {"box_2d": [171, 75, 245, 308], "label": "view recipe element"}\n]\n`

We test models to detect everyday objects, here we ask them to detect the bike and compare different model outputs. As in the previous case, we parse the bounding box from the json and translate to image space coordinates.

We asked Gemma 4 to write HTML code to reconstruct a page we made with Gemini 3. Below you can find the code to do this, we enable thinking and ask each model to generate up to 4000 new tokens, to make it foolproof.

## Inference code

```
messages = [
    {
        "role": "user",
        "content": [
            {
                "type": "image",
                "image": "https://huggingface.co/datasets/merve/vlm_test_images/resolve/main/landing_page.png",
            },
            {"type": "text", "text": "Write HTML code for this page."},
        ],
    }
]
inputs = processor.apply_chat_template(
    messages,
    tokenize=True,
    return_dict=True,
    return_tensors="pt",
    add_generation_prompt=True,
    enable_thinking=True,
).to(model.device)
output = model.generate(**inputs, max_new_tokens=4000)
input_len = inputs.input_ids.shape[-1]
generated_text_ids = output[0][input_len:]
generated_text = processor.decode(generated_text_ids, skip_special_tokens=True)
result = processor.parse_response(generated_text)
print(result["content"])
```
Smaller Gemma 4 models can take in videos with audio while larger ones can take in videos without audio. While the models are not explicitly post-trained on videos, they can understand videos both with and without audio. The model is particularly strong in audios

## Inference code

```
messages = [
    {
        "role": "user",
        "content": [
            {"type": "video", "url": "https://huggingface.co/datasets/merve/vlm_test_images/resolve/main/concert.mp4"},
            {"type": "text", "text": "What is happening in the video? What is the song about?"},
        ],
    },
]
inputs = processor.apply_chat_template(
    messages,
    tokenize=True,
    return_dict=True,
    return_tensors="pt",
    add_generation_prompt=True,
    load_audio_from_video=True, # disable this for larger models
).to(model.device)
output = model.generate(**inputs, max_new_tokens=200)
input_len = inputs.input_ids.shape[-1]
generated_text_ids = output[0][input_len:]
generated_text = processor.decode(generated_text_ids, skip_special_tokens=True)
print(result["content"])
```
| Model | Output | 
|---|---|
| E2B (hallucinated for the audio) | Based on the video provided, here is an analysis: **What is happening in this video?** The video captures a live musical performance, likely at a large outdoor festival or concert. ***Setting:** The scene is a large stage set up outdoors, indicated by the extensive lighting rig, large speakers, and the audience visible in the foreground. ***Performers:** There are several musicians on stage. The focal point seems to be a lead singer/guitarist in a blue shirt and white pants who is actively performing, playing an acoustic or electric guitar. There are also other musicians, including drummers and bassists, visible in the background. ***Atmosphere:** The atmosphere is energetic and dramatic, characterized by bright stage lights, intense blue and white spotlights, and a significant amount of stage fog/smoke, which adds to the concert mood. ***Audience:** The foreground is filled with the backs of audience members, showing they are close to the stage and engaged in watching the performance.**What is the song about?** The lyrics provided below the video are: > "I don't care 'bout street > From what I cherish > Could it's time to protest | 
| E4B ✅ | The video captures a live performance of a band on a stage with bright lights and smoke effects, surrounded by an audience. The band performs and sings to the crowd. The song appears to be about the struggles and disillusionment of modern life, specifically the feeling of being stuck and the difficulties in finding one's place in a chaotic and overwhelming world. Lyrics mention themes of the "concrete street," "broken promises," "storm of the day," and the difficulty of finding oneself amidst societal pressures. | 
| 26/A4B (no audio) ✅ | This video shows a live music performance at an outdoor concert or festival at night. A solo male artist, dressed in a blue shirt and white pants, is the central focus. He is playing an acoustic guitar and singing into a microphone on a large, lit stage. The performance is accompanied by a band, including a drummer and other musicians visible in the background and on large video screens. The stage is filled with atmospheric blue lighting, smoke, and bright spotlights. The video captures different perspectives, including shots from the crowd looking up at the stage and close-up views of the performers projected onto large screens. The crowd is visible in the foreground, watching the show. One of the screens also displays the "Estrella Damm" brand name. | 
| 31B (no audio) ✅ | A musician is performing on stage during a concert, playing an acoustic guitar and singing into a microphone. He is wearing a blue shirt and white pants. The stage is filled with atmospheric smoke and illuminated by bright blue and white lights. Other band members, including a drummer and a keyboardist, are also visible on stage. The video shows the musician from the perspective of the audience, with some shots focusing on him and others showing the large screen on the side of the stage, which displays close-ups of the performer and the crowd. | 

We have tested all models on captioning. All checkpoints perform very well and accurately capture nuance in complex scenerios. Here's the image we use, with the prompt "Write single detailed caption for this image.".

## Inference code

