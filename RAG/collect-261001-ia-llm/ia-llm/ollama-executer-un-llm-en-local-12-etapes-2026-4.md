---
id: collect-261001-ia-llm/ia-llm/ollama-executer-un-llm-en-local-12-etapes-2026-4
title: "macOS (via Homebrew)"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Nvidia"]
dates: []
keywords: ["agent", "agents", "attention", "deepseek", "embeddings", "gpu", "llama", "nvidia", "qwen", "tool calling"]
source: docs/RAG/collect-261001-ia-llm/ollama-executer-un-llm-en-local-12-etapes-2026.md
source_anchor: ""
source_lines: [282, 411]
sha256: 59c3dbeff40f58c373f30b13f89e72f383222d6615c6ebca48c81cf9e042b354
---

# macOS (via Homebrew)

## Étape 10 – Appel d’outils (tool calling) pour des agents IA

L’**appel d’outils** (ou *function calling*) permet à un modèle d’invoquer vos fonctions : consulter une base de données, appeler une API météo, lancer un calcul. Le modèle ne s’exécute pas lui-même la fonction ; il indique laquelle appeler et avec quels arguments, à charge pour votre code de l’exécuter et de renvoyer le résultat. C’est la brique de base de tout agent IA.

La bibliothèque Python d’Ollama simplifie l’opération : passez directement vos fonctions Python dans le paramètre `tools`, et la bibliothèque génère automatiquement leur schéma à partir des annotations de type et de la docstring.

```
import ollama
def meteo(ville: str) -> str:
    """Renvoie la météo actuelle d'une ville donnée."""
    donnees = {"Paris": "18 °C, nuageux", "Lyon": "21 °C, ensoleillé"}
    return donnees.get(ville, "Ville inconnue")
reponse = ollama.chat(
    model="qwen3",
    messages=[{"role": "user", "content": "Quel temps fait-il à Lyon ?"}],
    tools=[meteo],
)
for appel in reponse["message"].get("tool_calls", []):
    if appel["function"]["name"] == "meteo":
        resultat = meteo(**appel["function"]["arguments"])
        print("Résultat de l'outil :", resultat)
        # Résultat de l'outil : 21 °C, ensoleillé
```
En 2026, Ollama prend en charge l’appel d’outils **en flux continu** : l’application peut afficher le texte généré tout en déclenchant les appels de fonction en temps réel. Tous les modèles ne savent pas appeler des outils ; Qwen 3, Llama 3.1+ et gpt-oss font partie des plus fiables sur cette tâche. La documentation dédiée au tool calling détaille le format complet et la gestion des réponses d’outils.

## Étape 11 – Modèles de raisonnement (thinking) et Ollama Cloud

Les **modèles de raisonnement** comme DeepSeek-R1 produisent une phase de « réflexion » interne avant de formuler leur réponse, ce qui améliore nettement les tâches logiques. Ollama expose un paramètre `think` qui active ou désactive cette phase, et accepte même des niveaux d’intensité : `"low"`, `"medium"`, `"high"` ou `"max"`.

```
# En session interactive
ollama run deepseek-r1:8b
# Via l'API, avec contrôle du niveau de raisonnement
curl http://localhost:11434/api/chat -d '{
  "model": "deepseek-r1:8b",
  "messages": [{ "role": "user", "content": "Combien font 17 x 24 ?" }],
  "think": "medium",
  "stream": false
}'
```
Mais que faire si votre machine ne peut pas charger un modèle de 120 ou 671 milliards de paramètres ? C’est là qu’intervient **Ollama Cloud**. Introduit pour combler le fossé matériel, ce service exécute des modèles géants sur des serveurs distants tout en conservant exactement la même interface CLI et API. Vous gardez vos petits modèles en local et basculez ponctuellement vers le cloud pour les tâches lourdes.

```
# Connexion à votre compte Ollama Cloud
ollama signin
# Exécuter un modèle géant hébergé (aucun GPU local requis)
ollama run gpt-oss:120b-cloud
ollama run deepseek-v3.1:671b-cloud
ollama run qwen3-vl:235b-cloud
```
Le suffixe `-cloud` indique que l’inférence se déroule à distance. Attention toutefois : dès lors que vous utilisez Ollama Cloud, vos données quittent votre machine, et l’argument de confidentialité du « tout local » ne s’applique plus. Réservez le cloud aux contenus non sensibles, et conservez vos données personnelles ou confidentielles sur les modèles locaux.

## Étape 12 – Projet complet : un assistant RAG local sur vos documents

Assemblons tout dans un projet concret et utile : un assistant **RAG** (Retrieval-Augmented Generation) qui répond à vos questions à partir de vos propres documents, sans qu’aucun fichier ne quitte votre poste. Le principe : on découpe les documents en morceaux, on les transforme en vecteurs avec un modèle d’embeddings, puis pour chaque question on récupère les passages les plus pertinents et on les fournit au LLM comme contexte.

Prérequis : créez un dossier `documents/` contenant quelques fichiers `.txt` ou `.md`, et installez les dépendances. Nous utilisons `nomic-embed-text` pour les embeddings et `llama3.2` pour la génération.

```
pip install ollama numpy
ollama pull nomic-embed-text
ollama pull llama3.2
mkdir documents   # placez-y vos fichiers .txt / .md
```
Voici le script complet. Il indexe vos documents au démarrage, puis ouvre une boucle de questions-réponses. La recherche utilise une similarité cosinus sur des vecteurs maintenus en mémoire avec NumPy – pas besoin de base vectorielle externe pour un volume modeste.

```
import os
import numpy as np
import ollama
DOSSIER = "./documents"
MODELE_EMBED = "nomic-embed-text"
MODELE_CHAT = "llama3.2"
def charger_documents(dossier):
    morceaux = []
    for nom in os.listdir(dossier):
        if nom.endswith((".txt", ".md")):
            chemin = os.path.join(dossier, nom)
            with open(chemin, encoding="utf-8") as f:
                texte = f.read()
            # Découpage en morceaux d'environ 1000 caractères
            for i in range(0, len(texte), 800):
                morceaux.append(texte[i:i + 1000])
    return morceaux
def indexer(morceaux):
    vecteurs = []
    for m in morceaux:
        rep = ollama.embed(model=MODELE_EMBED, input=m)
        vecteurs.append(rep["embeddings"][0])
    return np.array(vecteurs)
def rechercher(question, morceaux, matrice, k=3):
    rep = ollama.embed(model=MODELE_EMBED, input=question)
    q = np.array(rep["embeddings"][0])
    scores = matrice @ q / (
        np.linalg.norm(matrice, axis=1) * np.linalg.norm(q) + 1e-10
    )
    indices = scores.argsort()[-k:][::-1]
    return [morceaux[i] for i in indices]
def repondre(question, contexte):
    invite = (
        "Réponds uniquement à partir du contexte ci-dessous. "
        "Si l'information est absente, dis-le clairement.\n\n"
        "Contexte :\n" + "\n---\n".join(contexte) +
        "\n\nQuestion : " + question
    )
    flux = ollama.chat(
        model=MODELE_CHAT,
        messages=[{"role": "user", "content": invite}],
        stream=True,
    )
    for partie in flux:
        print(partie["message"]["content"], end="", flush=True)
    print()
if __name__ == "__main__":
    print("Indexation des documents locaux...")
    morceaux = charger_documents(DOSSIER)
    matrice = indexer(morceaux)
    print(f"{len(morceaux)} morceaux indexés. Posez vos questions (Ctrl+C pour quitter).")
    while True:
        q = input("\nVous : ")
        contexte = rechercher(q, morceaux, matrice)
        print("Assistant : ", end="")
        repondre(q, contexte)
```
Lancez le script avec `python assistant.py`. Posez une question dont la réponse figure dans vos documents : l’assistant retrouve les passages pertinents et répond en s’appuyant dessus, en français. Vous disposez d’un moteur de recherche documentaire intelligent, 100 % local et conforme au RGPD par conception. Pour l’orchestrer dans des flux plus complexes, voyez notre tutoriel n8n auto-hébergé, qui sait appeler Ollama comme nœud d’IA.

## Optimiser les performances : GPU, quantification et contexte

Une fois Ollama opérationnel, plusieurs leviers améliorent la vitesse et la stabilité. Le premier est l’**accélération GPU** : Ollama la détecte automatiquement, mais vérifiez avec `ollama ps` que la colonne « PROCESSOR » indique bien « GPU ». Sur une machine multi-GPU NVIDIA, la variable `CUDA_VISIBLE_DEVICES` permet de cibler une carte précise.

