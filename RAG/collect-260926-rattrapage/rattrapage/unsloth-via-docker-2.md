---
id: collect-260926-rattrapage/rattrapage/unsloth-via-docker-2
title: "Install Unsloth via Docker"
domain: rattrapage
role: reference
task: reference
actors: ["Hugging Face", "Unsloth"]
dates: []
keywords: ["diffusion", "embedding", "embeddings", "fine-tuning", "gguf", "gpu", "memory", "parameters", "quantization", "training"]
source: docs/RAG/lot-rattrapage/fine-tuning/Unsloth via Docker.md
source_anchor: ""
source_lines: [156, 238]
sha256: 954709b4e360fb87b838d006e81f4672f7def60cad1425fd11ec3f6ebfc2a863
---

# Install Unsloth via Docker

docker logs -f unsloth
```
Look for **Unsloth container ready** for links and password details. Press **Ctrl+C** to stop following the logs; the container keeps running.
{% endstep %}
{% step %}
**Sign in**
Open [`http://localhost:8000`](http://localhost:8000) and sign in as `unsloth`. On first setup, use the generated password and choose a new one when prompted. Otherwise, use your existing password.
{% endstep %}
{% step %}
**Forgot your password?**
```bash
docker exec unsloth unsloth studio reset-password --username unsloth
```
This generates a new password, signs out existing sessions and revokes API keys. No restart is needed. These commands assume your container is named `unsloth`. If needed, run `docker ps` to find its name or ID.
{% hint style="info" %}
If Docker runs on another machine (like a cloud GPU or a server), `localhost` means your own computer. You must forward the ports before you open the links. Run `ssh -L 8000:localhost:8000 -L 8888:localhost:8888 user@your-server`. See the [Security Notes](#security-notes) section for details.
{% endhint %}
{% endstep %}
{% endstepper %}
#### **Chat with a model**
{% hint style="info" %}
Unsloth Studio and JupyterLab share the same GPU. A model loaded in Unsloth keeps its GPU memory until you unload it. If you run out of memory, unload the model in Unsloth or stop the notebook kernel.
{% endhint %}
{% stepper %}
{% step %}
**Choose and download a model**
Open **Select model** at the top of the page. Browse **Recommended**, or enter a model name and click **Search Hub** to search Hugging Face. You can also browse models in **Model hub**.
For GGUF models, choose a quantization that fits your available RAM and VRAM. Avoid variants marked **OOM**.
Wait for the model to finish downloading and loading before sending your first message. You can select downloaded models again from **On Device**.
{% hint style="info" %}
Downloaded models are stored in `/workspace/.cache/huggingface`. With the Quickstart volume mount, this cache is shared with notebooks and stays on your computer even if you remove the container.
{% endhint %}
{% endstep %}
{% step %}
**Send a message**
Once the model is ready, type a message and press **Enter**. Try:
> Explain how Docker containers differ from virtual machines.
See the [Unsloth Studio Chat](/docs/new/studio/chat.md) guide for web search and more chat settings.
{% endstep %}
{% endstepper %}
#### **Train a model**
You can train models across text, [vision](https://unsloth.ai/docs/basics/vision-fine-tuning) and [audio](https://unsloth.ai/docs/basics/text-to-speech-tts-fine-tuning), with support for [embeddings](https://unsloth.ai/docs/basics/embedding-finetuning) and image diffusion too. All from the same Docker image.
**Choose a model and dataset**
Select **Train** and open **Configure**. Choose a model and training method, then select a Hugging Face dataset or upload your own.
**Start training**
Review **Parameters** in **Simple** or **Advanced** mode, then click **Start Training**. Track progress, loss and GPU usage in **Current Run**.
**Try your trained model**
When training finishes, click **Compare in Chat**. Test the original and fine-tuned models with the same prompts to see how their responses differ.
See the [Studio training guide](https://unsloth.ai/docs/new/studio/start) for more detail, or use [Data Recipes](https://unsloth.ai/docs/new/studio/data-recipe) to prepare your dataset.
#### **Generate images**
{% stepper %}
{% step %}
**Open Images and choose a model**
Select **Images** in the sidebar. The **Create** workflow opens by default.
Open **Select image model** and choose from **Recommended**, such as Z-Image-Turbo. For GGUF models, pick a quantization that fits your device. **TIGHT** may run slowly; **OOM** is unlikely to fit.
{% endstep %}
{% step %}
**Enter a prompt and generate**
Describe the image you want to create. For example:
> Top-down shot of a koi pond in a Japanese garden, dozens of orange, white and black koi swimming in tight formation, water clear enough to see the stone bottom, red maple leaves floating on the surface. Bright midday light, saturated colour, crisp reflections. Realistic photo, 50mm.
Leave the image settings at their defaults and click **Generate**.
Open your image in the gallery and click **Download** to save it, or **Recipe** to reuse its prompt and settings.
**Transform**, **Inpaint** and **Edit** depend on the loaded model. See the [image diffusion guide](https://unsloth.ai/docs/basics/diffusion-image) for supported workflows and settings.
{% endstep %}
{% endstepper %}
#### **Access JupyterLab**
{% columns %}
{% column %}
The Unsloth Docker image includes JupyterLab and ready-to-use notebooks.
Open [`http://localhost:8888`](http://localhost:8888) and sign in. Use the `JUPYTER_PASSWORD` you chose when starting the container. If you did not set one, use the generated password printed in the container log.
{% hint style="info" %}
JupyterLab passwords are set when the container is created. To change it, remove and recreate the container with a new password.
{% endhint %}
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
After signing in, you’ll see **Unsloth Notebooks**, with folders grouped by model and task.
These folders contain shortcuts to the notebook files in `/workspace/unsloth-notebooks`. Each time the container starts, the notebooks are refreshed from GitHub without overwriting
\
your edits.
Double-click **01 Main Notebooks** to browse examples, or choose a category such as vision, speech or reinforcement learning.

