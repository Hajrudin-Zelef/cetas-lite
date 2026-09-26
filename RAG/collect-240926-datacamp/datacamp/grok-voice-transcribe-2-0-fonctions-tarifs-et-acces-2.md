---
id: collect-240926-datacamp/datacamp/grok-voice-transcribe-2-0-fonctions-tarifs-et-acces-2
title: "grok-voice-transcribe-2-0-fonctions-tarifs-et-acces"
domain: datacamp
role: reference
task: reference
actors: ["Google", "Meta", "OpenAI", "OpenRouter", "SpaceX", "xAI"]
dates: ["2026-09-21"]
keywords: ["grok", "voice", "agent", "agents", "consumer", "cost", "grok 4", "latency", "leaderboard", "muse", "pricing", "sandbox"]
source: docs/RAG/clean_en/datacamp/grok-voice-transcribe-2-0-fonctions-tarifs-et-acces.md
source_anchor: ""
source_lines: [98, 170]
sha256: 69f5ac59cdbf9054871da3e0da968a037e79c30961bf686584e8dcd9e01c4096
---

# grok-voice-transcribe-2-0-fonctions-tarifs-et-acces

Grok Voice Transcribe 2.0 improves over 1.0 on all four, and on telephony, SpaceXAI claims to lead all models tested.

The only figure SpaceXAI published for these sets concerns short phrases: WER drops from 20.6% (1.0) to 6.8% (2.0). In a car, brief commands leave very little context to identify the language, which was precisely 1.0's weak point; a 3x gain is the most convincing indicator behind the "twice as accurate" claim.

The rest of the internal charts, including the multilingual comparison against ElevenLabs Scribe v2 and Deepgram Nova-3, are presented as histograms without numerical values. Treat "leads all models tested" on telephony as a vendor claim until reproduced on your own call audio.

## Pricing and availability of Grok Voice Transcribe 2.0

Grok Voice Transcribe 2.0 costs $0.10 per hour of audio in batch and $0.20 per hour in streaming, identical to Grok Voice Transcribe 1.0. Diarization, word-level timestamps, and key term biasing are included in these rates, with no surcharges.

| Mode | Price | Included |
|---|---|---|
| Batch (file or URL) | $0.10 per hour of audio | Diarization, timestamps, key terms, formatting |
| Streaming (WebSocket) | $0.20 per hour of audio | Diarization, timestamps, key terms, formatting, smart turn-taking |

These rates are among the lowest on the streaming leaderboard. Artificial Analysis lists Grok 2.0 at $3.33 per 1,000 minutes, versus $3.00 for Meta's Muse Voice Transcribe and $6.50 for ElevenLabs Scribe v2 Realtime as well as Deepgram Flux. Among the four most accurate models on the index, only Muse is cheaper, but with lower accuracy.

The model is generally available via the SpaceXAI API, with no waitlist or regional restrictions mentioned in the announcement or documentation. There is no consumer offering, as this is an API product, and neither the announcement nor the documentation mentions a free tier for speech-to-text.

The default is in transition. SpaceXAI states that 2.0 will soon become the default for the Speech-to-Text API and that 1.0 will be deprecated in the coming weeks. Until then, set the model ID explicitly.

## How do you get access to Grok Voice Transcribe 2.0?

The model ID is `grok-voice-transcribe-2.0`, passed as a form field on the REST endpoint or as a query parameter on the WebSocket endpoint. To stay on the old model during the deprecation period, specify `grok-voice-transcribe-1.0`.

The model runs on two surfaces: the SpaceXAI API, with batch at `https://api.x.ai/v1/stt` and streaming at `wss://api.x.ai/v1/stt`, and the SpaceXAI console, which offers a live voice transcription sandbox. It does not appear in the OpenRouter catalog as of September 21, 2026.

Here is the minimal batch call that returns a transcript with diarization:

```
import os
import requests
response = requests.post(
    "https://api.x.ai/v1/stt",
    headers={"Authorization": f"Bearer {os.environ['XAI_API_KEY']}"},
    data=[("model", "grok-voice-transcribe-2.0"), ("diarize", "true")],
    files={"file": open("standup.wav", "rb")},
)
print(response.json()["text"])
```
For voice agents built on xAI's WebSocket voice APIs, see our tutorial on the Grok Voice Think Fast 2.0 API, which covers the speech-to-speech model, as well as our Grok Voice Agent API guide. If you also call Grok's text models from the same code, our Grok 4.6 API tutorial covers keys and tool orchestration.

## Conclusion

Grok Voice Transcribe 2.0 is today the cheapest way to get the most accurate streaming transcription on the public leaderboard, a rare combination. xAI is sending the signal that its voice stack is aiming at OpenAI, Google, and ElevenLabs, not just competing on text models.

I would switch if my audio is telephone-quality, multilingual, or full of dictated numbers: that is where 2.0 stands out from 1.0 and from most rivals. I would wait, however, for a voice agent that is very latency-sensitive, until xAI closes the gap with Scribe v2 on final transcription delay.

If you want to build transcription pipelines yourself, we recommend our Spoken Language Processing in Python course, which goes from raw audio files to transcription and call classification.

## Grok Voice Transcribe 2.0: FAQ

### How does Grok Voice Transcribe 2.0 compare to Grok Voice Transcribe 1.0?

Grok Voice Transcribe 2.0 is more accurate than 1.0 at the same price and through the same API. On Artificial Analysis's streaming WER index, it scores 2.7% versus 3.9% for 1.0, and on xAI's internal set of short sentences, the error rate drops from 20.6% to 6.8%. It is slower to return a final transcription: 0.49 s versus 0.37 s for 1.0.

### How much does Grok Voice Transcribe 2.0 cost?

Batch transcription costs $0.10 per hour of audio and streaming $0.20 per hour, identical to Grok Voice Transcribe 1.0. Speaker diarization, word-level timestamps, and key-term bias are included in these rates. xAI does not publish a free tier for speech-to-text.

### How do you access Grok Voice Transcribe 2.0?

Call xAI's Speech-to-Text API with the model ID `grok-voice-transcribe-2.0`, either as a multipart form field on the REST endpoint or as a query parameter on the WebSocket streaming endpoint. You can also try it in the speech-to-text playground in the xAI console.

### What languages does Grok Voice Transcribe 2.0 support?

The model transcribes dozens of languages, automatically detects the language, and follows language changes within the same recording. Written-form formatting of numbers, dates, and currencies is available in 25 languages when you set the language parameter, including English, Spanish, German, French, Japanese, Hindi, and Arabic.

### Does Grok Voice Transcribe 2.0 handle speaker diarization and real-time streaming?

Yes. Diarization adds a speaker label to each word at no extra cost, and multichannel mode transcribes up to 8 independent audio channels. Streaming is done via WebSocket with partial transcriptions every ~500 ms, plus intelligent turn detection so a voice agent waits for the end of an idea rather than every pause.

**Senior Data Science Editor at DataCamp |** **I am passionate about forecasting and development using APIs.**
