---
id: collect-261001-general-networking/general-networking/grok-voice-transcribe-2-0-fonctions-tarifs-et-acces-2
title: "grok-voice-transcribe-2-0-fonctions-tarifs-et-acces"
domain: general-networking
role: reference
task: reference
actors: ["Google", "Meta", "OpenAI", "OpenRouter", "SpaceX", "xAI"]
dates: []
keywords: ["grok", "voice", "agent", "agents", "grok 4", "muse", "transcription"]
source: docs/RAG/collect-261001-general-networking/grok-voice-transcribe-2-0-fonctions-tarifs-et-acces.md
source_anchor: ""
source_lines: [75, 165]
sha256: 858412b97156d20fca2e34ee1133505e70d7c69b11159cd74e4668eef69b090c
---

# grok-voice-transcribe-2-0-fonctions-tarifs-et-acces

L'écart avec la 1.0 sur cet indice est de 31 %, et non pas « fois deux » comme le laisse entendre l'annonce. Cette affirmation plus ambitieuse vient des jeux en production de SpaceXAI, que nous abordons ensuite. Le même classement mesure également le délai jusqu'à la transcription finale, et là, le tableau s'inverse.

| Modèle | Indice WER (transcription finale) | Délai jusqu'à la transcription finale | 
|---|---|---|
| Grok Voice Transcribe 2.0 | 2,7 % | 0,49 s | 
| Muse Voice Transcribe (Meta) | 3,1 % | 0,16 s | 
| ElevenLabs Scribe v2 Realtime | 3,6 % | 0,14 s | 
| Grok Voice Transcribe 1.0 | 3,9 % | 0,37 s | 
| Deepgram Flux | 7,4 % | 0,02 s | 

La latence est le coût à payer. Grok 2.0 met 0,49 s à renvoyer une transcription finale, contre 0,37 s pour la 1.0 et 0,14 s pour Scribe v2 Realtime. Pour des sous-titres, c'est invisible. Pour un agent vocal qui doit répondre à chaque tour de parole, ces dixièmes s'additionnent.

### Audio réel : téléphonie, identifiants et courtes phrases

SpaceXAI mesure le WER sur quatre jeux internes issus du trafic de production :

- Téléphonie 8 kHz de centres de support client
- Conversations avec Grok
- Identifiants dictés (codes de compte, adresses e-mail)
- Courtes commandes d'assistant vocal en 19 langues

Grok Voice Transcribe 2.0 s'améliore face à la 1.0 sur les quatre, et sur la téléphonie, SpaceXAI affirme devancer tous les modèles testés.

Le seul chiffre publié par SpaceXAI pour ces jeux concerne les courtes phrases : le WER passe de 20,6 % (1.0) à 6,8 % (2.0). En voiture, de brèves commandes laissent très peu de contexte pour identifier la langue, ce qui était justement le point faible de la 1.0 ; un gain par 3 constitue l'indicateur le plus probant derrière l'affirmation « deux fois plus précise ».

Le reste des graphiques internes, y compris la comparaison multilingue face à ElevenLabs Scribe v2 et Deepgram Nova-3, est présenté sous forme d'histogrammes sans valeurs chiffrées. Traitez « devance tous les modèles testés » sur la téléphonie comme une revendication fournisseur jusqu'à reproduction sur vos propres audios d'appel.

## Tarifs et disponibilité de Grok Voice Transcribe 2.0

Grok Voice Transcribe 2.0 coûte 0,10 $ par heure d'audio en batch et 0,20 $ par heure en streaming, identique à Grok Voice Transcribe 1.0. La diarisation, les horodatages par mot et le biais sur termes clés sont inclus dans ces tarifs, sans suppléments.

| Mode | Prix | Inclus | 
|---|---|---|
| Batch (fichier ou URL) | 0,10 $ par heure d'audio | Diarisation, horodatages, termes clés, formatage | 
| Streaming (WebSocket) | 0,20 $ par heure d'audio | Diarisation, horodatages, termes clés, formatage, tour de parole intelligent | 

Ces tarifs comptent parmi les plus bas du classement streaming. Artificial Analysis indique Grok 2.0 à 3,33 $ pour 1 000 minutes, contre 3,00 $ pour Muse Voice Transcribe de Meta et 6,50 $ pour ElevenLabs Scribe v2 Realtime comme pour Deepgram Flux. Parmi les quatre modèles les plus précis de l'indice, seul Muse est moins cher, mais avec une précision inférieure.

Le modèle est généralement disponible via l'API SpaceXAI, sans liste d'attente ni restriction régionale mentionnées dans l'annonce ou la documentation. Il n'y a pas d'offre grand public, car il s'agit d'un produit API, et ni l'annonce ni la documentation ne mentionnent de palier gratuit pour le speech-to-text.

Le défaut est en cours de transition. SpaceXAI indique que la 2.0 deviendra bientôt le défaut de l'API Speech-to-Text et que la 1.0 sera dépréciée dans les prochaines semaines. D'ici là, fixez l'ID du modèle explicitement.

## Comment obtenir l'accès à Grok Voice Transcribe 2.0 ?

L'ID du modèle est `grok-voice-transcribe-2.0`, transmis comme champ de formulaire sur l'endpoint REST ou comme paramètre de requête sur l'endpoint WebSocket. Pour rester sur l'ancien modèle pendant la période de dépréciation, indiquez `grok-voice-transcribe-1.0`.

Le modèle fonctionne sur deux surfaces : l'API SpaceXAI, avec le batch sur `https://api.x.ai/v1/stt` et le streaming sur `wss://api.x.ai/v1/stt`, et la console SpaceXAI, qui propose un bac à sable de transcription vocale en direct. Il n'apparaît pas dans le catalogue OpenRouter au 21 septembre 2026.

Voici l'appel batch minimal qui renvoie une transcription avec diarisation :

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
Pour les agents vocaux construits sur les APIs vocales WebSocket de xAI, consultez notre tutoriel sur l'API Grok Voice Think Fast 2.0, qui couvre le modèle speech-to-speech, ainsi que notre guide de l'API Grok Voice Agent. Si vous appelez aussi les modèles texte de Grok depuis le même code, notre tutoriel API Grok 4.6 couvre les clés et l'orchestration d'outils.

## Conclusion

Grok Voice Transcribe 2.0 est aujourd'hui la façon la moins chère d'obtenir la transcription streaming la plus précise du classement public, une combinaison rare. xAI envoie le signal que sa pile vocale vise OpenAI, Google et ElevenLabs, pas seulement la compétition sur les modèles texte.

Je basculerais si mon audio est de qualité téléphonique, multilingue ou plein de nombres dictés : c'est là que la 2.0 se démarque de la 1.0 et de la plupart des rivaux. J'attendrais en revanche pour un agent vocal très sensible à la latence, le temps que xAI comble l'écart avec Scribe v2 sur le délai de transcription finale.

Si vous souhaitez construire vous-même des pipelines de transcription, nous vous recommandons notre cours Spoken Language Processing in Python, qui va des fichiers audio bruts jusqu'à la transcription et la classification d'appels.

## Grok Voice Transcribe 2.0 : FAQ

### Comment Grok Voice Transcribe 2.0 se compare-t-il à Grok Voice Transcribe 1.0 ?

Grok Voice Transcribe 2.0 est plus précis que la 1.0 au même prix et via la même API. Sur l'indice de WER en streaming d'Artificial Analysis, il obtient 2,7 % contre 3,9 % pour la 1.0, et sur le jeu interne de courtes phrases de xAI, le taux d'erreur passe de 20,6 % à 6,8 %. Il est plus lent à renvoyer une transcription finale : 0,49 s contre 0,37 s pour la 1.0.

### Combien coûte Grok Voice Transcribe 2.0 ?

La transcription batch coûte 0,10 $ par heure d'audio et le streaming 0,20 $ par heure, identiques à Grok Voice Transcribe 1.0. La diarisation des locuteurs, les horodatages au niveau des mots et le biais sur termes clés sont inclus dans ces tarifs. xAI ne publie pas de palier gratuit pour le speech-to-text.

### Comment accéder à Grok Voice Transcribe 2.0 ?

Appelez l'API Speech-to-Text de xAI avec l'ID de modèle `grok-voice-transcribe-2.0`, soit comme champ de formulaire multipart sur l'endpoint REST, soit comme paramètre de requête sur l'endpoint de streaming WebSocket. Vous pouvez également l'essayer dans l'aire de test speech-to-text de la console xAI.

### Quelles langues Grok Voice Transcribe 2.0 prend-il en charge ?

Le modèle transcrit des dizaines de langues, détecte automatiquement la langue et suit les changements de langue au sein d'un même enregistrement. Le formatage en forme écrite des nombres, dates et devises est disponible dans 25 langues lorsque vous définissez le paramètre de langue, notamment l'anglais, l'espagnol, l'allemand, le français, le japonais, l'hindi et l'arabe.

### Grok Voice Transcribe 2.0 gère-t-il la diarisation des locuteurs et le streaming en temps réel ?

