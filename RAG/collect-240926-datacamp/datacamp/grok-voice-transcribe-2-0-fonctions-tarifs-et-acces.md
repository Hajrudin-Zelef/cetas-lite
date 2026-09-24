---
id: collect-240926-datacamp/datacamp/grok-voice-transcribe-2-0-fonctions-tarifs-et-acces
title: "grok-voice-transcribe-2-0-fonctions-tarifs-et-acces"
domain: datacamp
role: reference
task: reference
actors: ["Google", "Meta", "OpenAI", "OpenRouter", "SpaceX", "xAI"]
dates: ["2026-09-21"]
keywords: ["grok", "voice", "agent", "agents", "benchmarks", "consumer", "cost", "grok 4", "latency", "leaderboard", "muse", "pricing"]
source: docs/RAG/clean_en/datacamp/grok-voice-transcribe-2-0-fonctions-tarifs-et-acces.md
source_anchor: ""
source_lines: [1, 170]
sha256: e8d108918ea592624c2019f3e5399c512f648ae48bef46a6137555e3f5676eac
---

# grok-voice-transcribe-2-0-fonctions-tarifs-et-acces

<!-- source: https://www.datacamp.com/fr/blog/grok-voice-transcribe-2-0 -->

Course

Streaming speech-to-text has become a real competitive arena. OpenAI, Google, ElevenLabs, Meta, and Deepgram all offer real-time transcription models, and Artificial Analysis now ranks dozens of them according to a word error rate index. SpaceXAI's latest model, Grok Voice Transcribe 2.0, now holds the top spot on that index.

In this article, we review all the new features of Grok Voice Transcribe 2.0: its capabilities, the public and internal benchmarks, as well as pricing and API access. Also see our guides to the Grok Voice Think Fast 2.0 API and the GPT Live Transcribe API.

## In brief

- Grok Voice Transcribe 2.0 is xAI's new speech-to-text model, replacing Grok Voice Transcribe 1.0 at the same hourly rate.
- It ranks first in accuracy among streaming models on Artificial Analysis's public leaderboard.
- The biggest gains are on difficult audio: phone lines, account codes and dictated emails, and short multilingual commands.
- The trade-off is latency: the time to return a final transcription is longer than for 1.0 and ElevenLabs Scribe v2.
- Existing integrations benefit from the upgrade without changing code, but explicitly specify the model ID while waiting for the default to switch over.
- If you currently pay a competitor for streaming, it's worth testing side by side on your own audio.

## What is Grok Voice Transcribe 2.0?

Grok Voice Transcribe 2.0 is SpaceXAI's second-generation speech-to-text model, which xAI presents as its best transcription model. It is based on the audio foundation model behind Grok Voice, which, according to SpaceXAI, already handles tens of thousands of customer support calls per day, transcribes millions of hours of video narration, and powers the Grok assistant in Tesla vehicles.

What changes compared to 1.0 is the training, not the API. SpaceXAI trained 2.0 on live, noisy, multilingual audio recorded in varied environments, then fine-tuned it with post-training targeting the most difficult real-world conditions:

- Unstable phone lines
- Overlapping voices
- Regional accents
- Dictated phone numbers or email addresses

The flagship promise: 2.0 would be twice as accurate as 1.0 on xAI's real-world evaluations, with no change in pricing. We'll return to this claim in the benchmarks section, because the public leaderboard tells a more nuanced story than the internal evaluations.

## Key features of Grok Voice Transcribe 2.0

The feature list remains that of 1.0: xAI focused the entire upgrade on accuracy and did not touch the API surface.

### Transcribe files or live audio with a single model

You can send a recorded file or a URL to the batch endpoint, or stream raw audio via WebSocket and receive transcription events while the speaker is still talking. Both paths use the same model: transcribing a call recording and transcribing the live call should therefore be identical.

Batch mode accepts files up to 500 MB in 12 audio formats, including WAV, MP3, FLAC, as well as MP4 containers, plus raw PCM and G.711 telephone codecs. In streaming, the model can emit partial transcriptions every ~500 ms and marks each result as a finalized chunk or final utterance, which allows a captioning interface to display very early and then replace afterward.

### Identify who said what, and when

Each word is returned with a start and end time, and enabling diarization adds a speaker label to each word at no additional cost. For call center audio with the agent on one channel and the customer on another, multichannel mode transcribes up to 8 channels independently, which makes it possible to do without diarization.

The response is a single JSON object containing the full text, the detected language (BCP-47 code), the audio duration, and the word array. There is no separate call for timestamps.

### Teach it your vocabulary

You can pass up to 100 key terms per request, each up to 50 characters, to bias the model toward product names, medication names, or any term pronounced that is specific to your domain and absent from dictionaries. Text formatting writes numbers, dates, currencies, phone numbers, and email addresses in their written form when you set the language parameter, which SpaceXAI supports for 25 languages.

Filler words like "uh" are removed by default. This is the right choice for a transcription intended for human reading, and the wrong one for conversational analysis; enable the flag if you want to keep them.

### Tell a voice agent when the caller has finished

Intelligent turn detection allows a voice agent to wait for the end of an idea rather than reacting to every silence. You set a confidence threshold between 0 and 1, and the model only marks the utterance as finished when it is sufficiently confident that the speaker has finished; SpaceXAI suggests 0.5 for balanced use and 0.7 when numbers or addresses are being dictated.

A companion delay forces the end of the turn after a fixed silence, preventing a caller who has left the session idle from blocking it. Without intelligent detection, the default endpointing triggers after 400 ms of silence: good for short commands, penalizing for dictating a phone number.

### Switch languages mid-recording

The model automatically detects language and handles language switches within a single recording in one pass, with no language hint required. SpaceXAI presents the multilingual improvement as the most notable versus 1.0, and the short-phrase figures presented below confirm this.

One point to note: the language parameter remains important for formatting. Set it if you want numbers and currencies in written form, and leave it empty if the audio mixes languages and you prefer raw words.

## How does Grok Voice Transcribe 2.0 perform on benchmarks?

Grok Voice Transcribe 2.0 leads the public streaming leaderboard in accuracy and outperforms 1.0 on all internal sets reported by SpaceXAI, but the size of the gain varies widely by audio type.

### Streaming accuracy on the public leaderboard

On Artificial Analysis's streaming speech-to-text index, Grok Voice Transcribe 2.0 shows a WER (Word Error Rate) of 2.7%, the lowest among the 34 streaming models listed as of September 21, 2026. Meta's Muse Voice Transcribe follows at 3.1%, ElevenLabs Scribe v2 Realtime is at 3.6%, and Grok Voice Transcribe 1.0 trails at 3.9%. SpaceXAI's announcement counted 32 models on the same leaderboard at launch.

**What WER measures:** WER is the share of incorrectly transcribed words: the lower, the better.

**What the speech-to-text index measures:** Artificial Analysis's index averages WER across 3 test sets (AA-AgentTalk, VoxPopuli, and Earnings-22). Grok 2.0's largest lead is on VoxPopuli: its WER of 1.4% is less than half that of 1.0 (3.1%).

The gap with 1.0 on this index is 31%, not "twice" as the announcement suggests. That more ambitious claim comes from SpaceXAI's production sets, which we address next. The same leaderboard also measures the delay until final transcription, and there, the picture reverses.

| Model | WER index (final transcription) | Delay until final transcription |
|---|---|---|
| Grok Voice Transcribe 2.0 | 2.7% | 0.49 s |
| Muse Voice Transcribe (Meta) | 3.1% | 0.16 s |
| ElevenLabs Scribe v2 Realtime | 3.6% | 0.14 s |
| Grok Voice Transcribe 1.0 | 3.9% | 0.37 s |
| Deepgram Flux | 7.4% | 0.02 s |

Latency is the price to pay. Grok 2.0 takes 0.49 s to return a final transcription, versus 0.37 s for 1.0 and 0.14 s for Scribe v2 Realtime. For subtitles, that's invisible. For a voice agent that must respond at every turn of speech, those tenths add up.

### Real-world audio: telephony, identifiers, and short phrases

SpaceXAI measures WER on four internal sets drawn from production traffic:

- 8 kHz telephony from customer support centers
- Conversations with Grok
- Dictated identifiers (account codes, email addresses)
- Short voice assistant commands in 19 languages

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
