---
id: collect-261001-ia-llm/ia-llm/github-copilot-agent-mode-14-etapes-55-min-2026-4
title: ".github/copilot-instructions.md"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Microsoft", "Mistral", "OpenAI", "xAI"]
dates: []
keywords: ["copilot", "agent", "benchmarks", "gemini", "grok", "grok 4", "mistral", "opus 4"]
source: docs/RAG/collect-261001-ia-llm/github-copilot-agent-mode-14-etapes-55-min-2026.md
source_anchor: ""
source_lines: [178, 307]
sha256: e8691f28d5740ffb9cc4d6bca0748e2dfe2e7941516abd28753bbe1f05bdc7d4
---

# .github/copilot-instructions.md

Cette possibilité de changer de modèle à la volée est l’un des arguments que GitHub met en avant face à ses concurrents, aux côtés de son intégration native avec les dépôts GitHub. Si vous voulez comparer les modèles disponibles en détail, notamment leurs performances respectives sur des tâches de programmation, nos comparatifs Opus 4.8 vs GPT-5.5 vs Mistral Large 3 et Grok 4.5 vs Opus 4.8 vs Gemini 3.1 Pro détaillent les écarts de score sur les benchmarks de développement logiciel.

Une règle simple pour débuter : laissez le modèle par défaut pour l’essentiel de votre usage quotidien, et ne basculez sur un modèle premium que pour les tâches où vous avez déjà vu le modèle standard échouer ou tourner en boucle. Cela préserve vos crédits mensuels pour les moments où ils font vraiment la différence.

Un exemple concret pour fixer les idées : demander à l’agent d’ajouter un endpoint REST qui suit un schéma déjà présent dans le projet est une tâche que le modèle par défaut traite très bien, car le contexte à imiter est déjà là. À l’inverse, lui demander de repenser la stratégie de cache d’une application entière, avec des compromis à évaluer entre cohérence des données et latence, sollicite un raisonnement plus long où un modèle premium montre un net avantage. La différence se voit autant dans la qualité du plan proposé que dans le nombre d’itérations nécessaires pour y arriver.

## Étape 10 à 13 – Construisez un projet complet : une API météo testée de bout en bout

Passons à la pratique avec un projet complet, assez simple pour se terminer en une vingtaine de minutes, mais assez complet pour mobiliser tout ce qui vient d’être expliqué : lecture de contexte, écriture multi-fichiers, exécution de commandes, génération de tests et correction automatique. L’objectif est une petite API qui renvoie la météo actuelle d’une ville, en s’appuyant sur l’API publique et gratuite Open-Meteo, qui ne nécessite aucune clé d’API.

### Étape 10 – Initialiser le projet

Créez un dossier de projet, initialisez un dépôt Git et un environnement virtuel Python :

```
mkdir api-meteo && cd api-meteo
git init
python3 -m venv venv
source venv/bin/activate
pip install fastapi uvicorn httpx pytest
```
Ouvrez ce dossier dans VS Code, faites un premier commit vide ou avec un fichier `.gitignore`, puis ouvrez la vue Chat en Agent Mode.

### Étape 11 – Générer l’API avec l’Agent Mode

Soumettez le prompt suivant à l’agent :

```
Crée une API FastAPI dans main.py avec un endpoint GET /meteo/{ville}.
Cet endpoint doit : 1) appeler l'API de géocodage Open-Meteo
(https://geocoding-api.open-meteo.com/v1/search) pour trouver les
coordonnées de la ville, 2) appeler l'API de prévisions Open-Meteo
(https://api.open-meteo.com/v1/forecast) avec current_weather=true,
3) renvoyer un JSON avec ville, pays, temperature_c, vent_kmh et
heure_mesure. Utilise httpx en mode asynchrone. Renvoie une erreur
404 explicite si la ville n'est pas trouvée.
```
L’agent lit le prompt, annonce son plan, puis génère un fichier proche de celui-ci :

```
# main.py
from fastapi import FastAPI, HTTPException
import httpx
app = FastAPI(title="API Météo")
GEOCODING_URL = "https://geocoding-api.open-meteo.com/v1/search"
FORECAST_URL = "https://api.open-meteo.com/v1/forecast"
@app.get("/meteo/{ville}")
async def get_meteo(ville: str):
    async with httpx.AsyncClient() as client:
        geo = await client.get(GEOCODING_URL, params={
            "name": ville, "count": 1, "language": "fr"
        })
        geo_data = geo.json()
        if not geo_data.get("results"):
            raise HTTPException(status_code=404, detail=f"Ville '{ville}' introuvable")
        lieu = geo_data["results"][0]
        forecast = await client.get(FORECAST_URL, params={
            "latitude": lieu["latitude"],
            "longitude": lieu["longitude"],
            "current_weather": True
        })
        meteo = forecast.json()["current_weather"]
        return {
            "ville": lieu["name"],
            "pays": lieu.get("country"),
            "temperature_c": meteo["temperature"],
            "vent_kmh": meteo["windspeed"],
            "heure_mesure": meteo["time"]
        }
```
### Étape 12 – Générer et exécuter les tests automatisés

Enchaînez avec un second prompt, dans la continuité de la même session :

```
Ajoute un fichier test_main.py avec pytest et TestClient de FastAPI.
Teste le cas d'une ville valide (Paris) et le cas d'une ville
inexistante. Exécute ensuite la suite de tests et corrige le code
si un test échoue.
```
L’agent crée le fichier de test, propose la commande `pytest -v`, l’exécute après votre validation, puis lit la sortie :

```
# test_main.py
from fastapi.testclient import TestClient
from main import app
client = TestClient(app)
def test_meteo_paris():
    response = client.get("/meteo/Paris")
    assert response.status_code == 200
    data = response.json()
    assert data["ville"] == "Paris"
    assert "temperature_c" in data
def test_meteo_ville_inconnue():
    response = client.get("/meteo/Xyzabc123")
    assert response.status_code == 404
```
### Étape 13 – Vérifier, corriger et déployer localement

Si tout se passe bien, le terminal affiche une sortie de ce type :

```
$ pytest -v
============================= test session starts ==============================
collected 2 items
test_main.py::test_meteo_paris PASSED                                    [ 50%]
test_main.py::test_meteo_ville_inconnue PASSED                          [100%]
============================== 2 passed in 1.84s ================================
```
Si un test échoue à la première tentative, typiquement parce que la clé retournée par l’API de géocodage ne correspond pas exactement à celle attendue dans le code, l’agent lit la trace d’erreur, propose une correction ciblée, puis relance les tests automatiquement. C’est le cycle décrit à l’étape 5, observable en conditions réelles. Lancez enfin le serveur local pour un dernier test manuel :

```
$ uvicorn main:app --reload
$ curl -s http://127.0.0.1:8000/meteo/Lyon | python3 -m json.tool
{
    "ville": "Lyon",
    "pays": "France",
    "temperature_c": 21.4,
    "vent_kmh": 11.2,
    "heure_mesure": "2026-06-15T14:00"
}
```
Vous disposez maintenant d’un projet complet, testé, généré en grande partie par l’agent mais entièrement relu et validé par vos soins à chaque étape, exactement comme le veut le flux de travail recommandé.

## Étape 14 – Suivre vos crédits et automatiser via Copilot CLI

Le tableau de bord de facturation, accessible depuis les paramètres de votre compte sur GitHub, affiche la consommation de crédits en temps quasi réel, ventilée par modèle utilisé. Prenez l’habitude de le consulter une fois par semaine si vous êtes sur le palier Pro, dont l’enveloppe de 15 $ de crédits se consomme plus vite qu’on ne l’imagine sur des sessions agentiques intensives.

Pour les tâches répétitives ou scriptables, GitHub propose aussi une extension en ligne de commande, indépendante de VS Code : Copilot CLI, entrée en préversion publique dès septembre 2025 comme agent de codage IA nativement conçu pour le terminal, disponible sur l’ensemble des paliers. Installez-la via GitHub CLI :

