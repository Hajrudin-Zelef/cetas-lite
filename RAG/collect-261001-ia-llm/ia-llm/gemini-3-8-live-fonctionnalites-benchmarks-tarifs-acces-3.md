---
id: collect-261001-ia-llm/ia-llm/gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces-3
title: "gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces"
domain: ia-llm
role: reference
task: reference
actors: ["Google", "OpenAI"]
dates: []
keywords: ["gemini", "agent", "agents", "astra", "gemini 3.8", "gpt-live", "transcription", "voice"]
source: docs/RAG/collect-261001-ia-llm/gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces.md
source_anchor: ""
source_lines: [131, 199]
sha256: 6a104006790bfc875dab8b1d3b0263c4af5d360c8efb2a26723f31aaeade2b34
---

# gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces

Le grounding avec Google Search est pris en charge sur les deux paliers, avec 5 000 requêtes de recherche gratuites par mois partagées sur tous les modèles Gemini 3.x, puis 14 $ par 1 000 requêtes. Google n’a pas publié de limites de débit spécifiques à ces modèles ; vérifiez la page des limites de l’API Gemini pour votre offre avant de planifier un lancement.

La disponibilité se décline en trois volets :

- **Développeurs :** les deux modèles sont en disponibilité générale dans l’API Gemini et Google AI Studio depuis le 15 septembre.
- **Entreprises :** disponibles en préversion privée dans Gemini Enterprise et bientôt dans Gemini Enterprise for Customer Experience, avec Extended Thinking également à venir pour les clients professionnels de Google Workspace.
- **Grand public :** Gemini 3.8 Live alimente Search Live, tandis qu’Extended Thinking arrive dans Gemini Live, ainsi que dans Docs pour les abonnés Google AI Pro et Ultra, et dans Gmail et Keep pour tous les abonnés Google AI.

Toutes les sorties audio des deux modèles portent un filigrane SynthID, et la fiche modèle de Google est transparente sur les limites : les modèles peuvent encore halluciner, la résistance aux jailbreaks est en amélioration, et il peut y avoir des lenteurs ou timeouts ponctuels. À lire avant de mettre un modèle devant des clients.

## Comment obtenir l’accès à Gemini 3.8 Live ?

Les identifiants de modèle sont `gemini-3.8-live` et `gemini-3.8-live-extended-thinking`, deux chaînes stables sans suffixe preview.

Vous pouvez y accéder via l’API Gemini Live avec le SDK google-genai pour Python ou JavaScript, via des WebSockets bruts, ou en interaction dans la vue Stream de Google AI Studio. Google liste également Agora, Fishjam, LiveKit, LangChain, Pipecat, Vercel et Vision Agents comme partenaires d’intégration gérant la couche de streaming média, et un chemin de streaming via l’Agent Development Kit si vous construisez déjà sur ADK.

La session Python minimale ci-dessous ouvre une connexion, envoie un tour en texte et active une transcription de sortie pour lire ce que le modèle a dit :

```
import asyncio
from google import genai
client = genai.Client()
config = {"response_modalities": ["AUDIO"], "output_audio_transcription": {}}
async def main():
    async with client.aio.live.connect(model="gemini-3.8-live", config=config) as session:
        await session.send_client_content(
            turns={"role": "user", "parts": [{"text": "Read back the claim number 7Q-4418-B."}]},
            turn_complete=True,
        )
        async for response in session.receive():
            if response.server_content and response.server_content.output_transcription:
                print(response.server_content.output_transcription.text)
asyncio.run(main())
```
Si vous migrez depuis `gemini-3.1-flash-live-preview`, changez l’identifiant de modèle et supprimez tout `thinking_level` ou `thinking_config` de votre configuration ; le cycle d’un tour reste sinon identique. L’API Live est par défaut serveur-à-serveur ; les clients navigateur et mobile nécessitent des jetons éphémères.

Pour un pas-à-pas complet sur la capture audio, la lecture et la gestion de session, consultez notre tutoriel Gemini Live API, et pour la partie texte de la même famille, notre tutoriel Gemini 3.8 Flash API couvre les niveaux de thinking et l’appel de fonctions en Python.

## Dernières réflexions

Google mise sur le prix plus que sur la qualité brute. Extended Thinking devance GPT-Live-1 Astra d’un ou deux points sur chaque tableau, mais Gemini 3.8 Live à 0,84 $ par heure d’audio en entrée coûte près de 7 fois moins que le modèle d’OpenAI, et c’est ce chiffre qui oriente le trafic voix en production.

Mon avis : si vous faites tourner quoi que ce soit sur Gemini 3.1 Flash Live, mettez à jour cette semaine : le prix est identique et l’appel d’outils asynchrone vaut à lui seul la bascule. Si vous êtes sur la pile temps réel d’OpenAI, le modèle de base mérite un essai pour les flux de tri et de recherche, et Extended Thinking en mérite un partout où vos outils prennent des secondes. Les 30,1 % du modèle de base sur τ-Voice sont la raison de ne pas l’utiliser pour de l’agentique.

Pour construire correctement la partie appel d’outils d’un agent vocal, nous recommandons notre cours Building AI Agents with Google ADK, qui connecte Gemini à un agent de support client avec outils, garde-fous et délégation.

## FAQ

### Quelle est la différence entre Gemini 3.8 Live et Gemini 3.8 Live Extended Thinking ?

Gemini 3.8 Live est le modèle speech-to-speech à faible latence de Google pour les tâches vocales directes, avec un raisonnement entrelacé et sans réglage de niveau de thinking. Gemini 3.8 Live Extended Thinking ajoute un raisonnement de fond configurable (low, medium ou high) et annonce des mises à jour d’avancement pendant qu’il exécute des outils asynchrones. Extended Thinking obtient 68,6 % sur τ-Voice contre 30,1 % pour le modèle de base : choisissez-le pour le travail agentique multi-étapes.

### Combien coûte Gemini 3.8 Live ?

Les deux modèles partagent une seule grille de prix sur l’offre payante : 3,00 $ par 1 million de tokens audio en entrée (environ 0,005 $ par minute) et 12,00 $ par 1 million de tokens audio en sortie, tokens de thinking inclus (environ 0,018 $ par minute). Le texte coûte 0,75 $ par 1 million de tokens en entrée et 4,50 $ par 1 million en sortie. Une offre gratuite couvre les deux modèles, avec les données du palier gratuit utilisées pour améliorer les produits de Google.

### Où accéder à Gemini 3.8 Live ?

Les développeurs peuvent utiliser `gemini-3.8-live` et `gemini-3.8-live-extended-thinking` via l’API Gemini Live et dans Google AI Studio, où les deux sont en disponibilité générale. Les entreprises y accèdent en préversion privée dans Gemini Enterprise. Côté grand public, on retrouve Gemini 3.8 Live dans Search Live et Extended Thinking dans Gemini Live, ainsi que dans Docs, Gmail et Keep pour les abonnés Google AI.

### Comment Gemini 3.8 Live se compare-t-il à Gemini 3.1 Flash Live ?

Gemini 3.8 Live remplace la préversion 3.1 Flash Live au même prix, avec l’appel de fonctions asynchrone activé par défaut et de meilleurs scores sur les tableaux de Google : 76,0 contre 71,5 sur le Speech to Speech Index d’Artificial Analysis. Migrer consiste à changer l’identifiant de modèle et à supprimer tout `thinking_level` ou `thinking_config` de la configuration de session. Google recommande à tous les utilisateurs de 3.1 Flash Live de passer à 3.8 Live.

### Gemini 3.8 Live prend-il en charge l’appel de fonctions et l’entrée vidéo ?

Oui. Les deux modèles prennent en charge l’appel de fonctions, et les appels d’outils s’exécutent en arrière-plan pendant que le modèle continue de parler. Gemini 3.8 Live accepte aussi des frames vidéo et des images en plus de l’audio et du texte, pour que l’agent réagisse à ce que l’utilisateur regarde. Les sessions audio+vidéo sont limitées à 2 minutes et les sessions audio seules à 15 minutes, sauf si vous utilisez la gestion de session pour les prolonger.

**Rédacteur en chef Data Science chez DataCamp |** **Je suis passionné par la prévision et le développement à l'aide d'API.**
