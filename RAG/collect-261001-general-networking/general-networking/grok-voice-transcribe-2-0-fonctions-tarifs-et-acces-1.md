---
id: collect-261001-general-networking/general-networking/grok-voice-transcribe-2-0-fonctions-tarifs-et-acces-1
title: "grok-voice-transcribe-2-0-fonctions-tarifs-et-acces"
domain: general-networking
role: reference
task: reference
actors: ["Google", "Meta", "OpenAI", "SpaceX", "xAI"]
dates: []
keywords: ["grok", "voice", "agent", "attention", "benchmarks", "muse", "transcription"]
source: docs/RAG/collect-261001-general-networking/grok-voice-transcribe-2-0-fonctions-tarifs-et-acces.md
source_anchor: ""
source_lines: [1, 74]
sha256: be029f8b77a00fb7c17edac813044a25bd5cec5342712908da77423593ceddfa
---

# grok-voice-transcribe-2-0-fonctions-tarifs-et-acces

Cours

Le streaming speech-to-text est devenu un véritable terrain de compétition. OpenAI, Google, ElevenLabs, Meta et Deepgram proposent tous des modèles de transcription en temps réel, et Artificial Analysis classe désormais des dizaines d'entre eux selon un indice de taux d'erreurs par mot. Le tout dernier modèle de SpaceXAI, Grok Voice Transcribe 2.0, occupe désormais la première place de cet indice.

Dans cet article, nous passons en revue toutes les nouveautés de Grok Voice Transcribe 2.0 : ses fonctionnalités, les benchmarks publics et internes, ainsi que les tarifs et l'accès à l'API. Consultez aussi nos guides de l'API Grok Voice Think Fast 2.0 et de l'API GPT Live Transcribe.

## En bref

- Grok Voice Transcribe 2.0 est le nouveau modèle speech-to-text de xAI, qui remplace Grok Voice Transcribe 1.0 au même tarif horaire.
- Il se classe premier en précision parmi les modèles de streaming sur le classement public d'Artificial Analysis.
- Les plus grands gains concernent les audios difficiles : lignes téléphoniques, codes de compte et e-mails dictés, et courtes commandes multilingues.
- Le compromis, c'est la latence : le délai pour renvoyer une transcription finale est plus long que pour la 1.0 et qu'ElevenLabs Scribe v2.
- Les intégrations existantes bénéficient de la mise à niveau sans changer de code, mais indiquez explicitement l'ID du modèle en attendant le basculement du défaut.
- Si vous payez aujourd'hui un concurrent pour du streaming, cela vaut la peine de tester côte à côte sur vos propres audios.

## Qu'est-ce que Grok Voice Transcribe 2.0 ?

Grok Voice Transcribe 2.0 est le modèle speech-to-text de deuxième génération de SpaceXAI, que xAI présente comme son meilleur modèle de transcription. Il est fondé sur le modèle de base audio derrière Grok Voice, qui, selon SpaceXAI, gère déjà des dizaines de milliers d'appels de support client par jour, transcrit des millions d'heures de narration vidéo et alimente l'assistant Grok dans les véhicules Tesla.

Ce qui change par rapport à la 1.0, c'est l'entraînement, pas l'API. SpaceXAI a entraîné la 2.0 sur de l'audio en direct, bruité et multilingue, enregistré dans des environnements variés, puis l'a affinée par un post-entraînement ciblant les conditions réelles les plus difficiles :

- Lignes téléphoniques instables
- Voix concurrentes
- Accents locaux
- Numéros de téléphone ou adresses e-mail dictés

La promesse phare : la 2.0 serait deux fois plus précise que la 1.0 sur les évaluations en conditions réelles de xAI, sans changement de tarif. Nous reviendrons sur cette affirmation dans la section benchmarks, car le classement public raconte une histoire plus nuancée que les jeux internes.

## Fonctionnalités clés de Grok Voice Transcribe 2.0

La liste des fonctions reste celle de la 1.0 : xAI a concentré toute la mise à niveau sur la précision et n'a pas touché à la surface de l'API.

### Transcrivez des fichiers ou de l'audio en direct avec un seul modèle

Vous pouvez envoyer un fichier enregistré ou une URL vers l'endpoint batch, ou bien streamer de l'audio brut via WebSocket et recevoir des événements de transcription pendant que l'interlocuteur parle encore. Les deux voies utilisent le même modèle : la transcription d'un enregistrement d'appel et celle de l'appel en direct doivent donc être identiques.

Le mode batch accepte des fichiers jusqu'à 500 Mo dans 12 formats audio, dont WAV, MP3, FLAC, ainsi que des conteneurs MP4, plus le PCM brut et les codecs téléphoniques G.711. En streaming, le modèle peut émettre des transcriptions partielles toutes les ~500 ms et marque chaque résultat comme bloc figé ou énoncé final, ce qui permet à une interface de sous-titrage d'afficher très tôt puis de remplacer ensuite.

### Identifiez qui a dit quoi, et quand

Chaque mot est renvoyé avec une heure de début et de fin, et l'activation de la diarisation ajoute un label de locuteur à chaque mot sans coût supplémentaire. Pour l'audio centre d'appels avec l'agent sur un canal et le client sur un autre, le mode multicanal transcrit jusqu'à 8 canaux indépendamment, ce qui permet de se passer de la diarisation.

La réponse est un objet JSON unique contenant le texte complet, la langue détectée (code BCP-47), la durée audio et le tableau des mots. Il n'y a pas d'appel séparé pour les horodatages.

### Apprenez-lui votre vocabulaire

Vous pouvez transmettre jusqu'à 100 termes clés par requête, chacun jusqu'à 50 caractères, pour biaiser le modèle vers des noms de produits, de médicaments, ou tout terme prononcé propre à votre domaine et absent des dictionnaires. Le formatage du texte écrit les nombres, dates, devises, numéros de téléphone et adresses e-mail sous leur forme écrite lorsque vous définissez le paramètre de langue, que SpaceXAI prend en charge pour 25 langues.

Les mots de remplissage comme « euh » sont supprimés par défaut. C'est le bon choix pour une transcription destinée à la lecture humaine, et le mauvais pour l'analyse conversationnelle ; activez le drapeau si vous souhaitez les conserver.

### Indiquez à un agent vocal quand l'appelant a terminé

La détection intelligente de tours de parole permet à un agent vocal d'attendre la fin d'une idée plutôt que de réagir à chaque silence. Vous définissez un seuil de confiance entre 0 et 1, et le modèle ne marque l'énoncé comme terminé que lorsqu'il est suffisamment confiant que l'intervenant a fini ; SpaceXAI suggère 0,5 pour un usage équilibré et 0,7 lorsque des nombres ou adresses sont dictés.

Un délai compagnon force la fin du tour après un silence fixe, évitant qu'un appelant parti séjourne ne bloque la session. Sans détection intelligente, l'endpointing par défaut se déclenche après 400 ms de silence : bien pour de courtes commandes, pénalisant pour la dictée d'un numéro de téléphone.

### Changez de langue en plein enregistrement

Le modèle détecte automatiquement la langue et gère les changements de langue au sein d'un même enregistrement en un seul passage, sans indice de langue requis. SpaceXAI présente l'amélioration en multilingue comme la plus marquante face à la 1.0, et les chiffres sur courtes phrases présentés plus loin le confirment.

Un point d'attention : le paramètre de langue reste important pour le formatage. Renseignez-le si vous souhaitez des nombres et devises en forme écrite, et laissez-le vide si l'audio mêle des langues et que vous préférez les mots bruts.

## Comment Grok Voice Transcribe 2.0 se comporte-t-il sur les benchmarks ?

Grok Voice Transcribe 2.0 mène le classement public du streaming en précision et surpasse la 1.0 sur tous les jeux internes rapportés par SpaceXAI, mais l'ampleur du gain varie fortement selon les types d'audio.

### Précision en streaming sur le classement public

Sur l'indice streaming speech-to-text d'Artificial Analysis, Grok Voice Transcribe 2.0 affiche un WER (Word Error Rate) de 2,7 %, le plus bas des 34 modèles en streaming listés au 21 septembre 2026. Muse Voice Transcribe de Meta suit à 3,1 %, ElevenLabs Scribe v2 Realtime est à 3,6 %, et Grok Voice Transcribe 1.0 ferme la marche à 3,9 %. L'annonce de SpaceXAI comptait 32 modèles sur le même classement au lancement.

**Ce que mesure le WER :** le WER correspond à la part de mots mal transcrits : plus il est bas, mieux c'est.

**Ce que mesure l'indice speech-to-text :** l'indice d'Artificial Analysis moyenne le WER sur 3 jeux de test (AA-AgentTalk, VoxPopuli et Earnings-22). L'avance la plus large de Grok 2.0 est sur VoxPopuli : son WER de 1,4 % représente moins de la moitié de celui de la 1.0 (3,1 %).

