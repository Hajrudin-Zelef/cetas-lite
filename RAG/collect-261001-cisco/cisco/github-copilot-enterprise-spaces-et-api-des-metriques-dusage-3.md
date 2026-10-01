---
id: collect-261001-cisco/cisco/github-copilot-enterprise-spaces-et-api-des-metriques-dusage-3
title: "github-copilot-enterprise-spaces-et-api-des-metriques-dusage"
domain: cisco
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2025-10-01", "2026-03-10", "2026-05-14"]
keywords: ["copilot", "agent", "benchmarks"]
source: docs/RAG/collect-261001-cisco/github-copilot-enterprise-spaces-et-api-des-metriques-dusage.md
source_anchor: ""
source_lines: [298, 552]
sha256: 5ddd22d22b6d628a90eff0bc98baad2cd26ebb20cbbc99fa90552d93b1c080c4
---

# github-copilot-enterprise-spaces-et-api-des-metriques-dusage

```
{
  "type": "object",
  "title": "Copilot Metrics 1 Day Report",
  "description": "Links to download the Copilot usage metrics report for an enterprise/organization for a specific day.",
  "properties": {
    "download_links": {
      "type": "array",
      "items": {
        "type": "string",
        "format": "uri"
      },
      "description": "The URLs to download the Copilot usage metrics report for the enterprise/organization for the specified day."
    },
    "report_day": {
      "type": "string",
      "format": "date",
      "description": "The day of the report in YYYY-MM-DD format."
    }
  },
  "required": [
    "download_links",
    "report_day"
  ]
}
```
#### Rapport organisationnel sur 28 jours

Le rapport sur 28 jours met en évidence les schémas d’adoption et les évolutions à plus long terme. Les commandes sont quasi identiques, en pointant vers l’API 28 jours.

Exemple de requête :

```
curl -L \
 -H "Accept: application/vnd.github+json" \
 -H "Authorization: Bearer <YOUR_TOKEN>" \
-H “X-GitHub-Api-Version: 2026-03-10” \
https://api.github.com/enterprises/ENTERPRISE/copilot/metrics/reports/enterprise-28-day/latest
```
Vous obtiendrez une réponse similaire, avec toutefois `response_start_day` et `response_end_day`.

#### Structure des rapports organisationnels

Les rapports JSON sur 1 jour et 28 jours au niveau organisation peuvent ressembler à ceci :

```
[
  {
    "user_id": 1001,
    "user_login": "octocat",
    "day": "2026-05-14",
    "organization_id": "999",
    "team_id": 42,
    "slug": "frontend"
  },
  {
    "user_id": 1001,
    "user_login": "octocat",
    "day": "2026-05-14",
    "organization_id": "999",
    "team_id": 43,
    "slug": "backend"
  },
  {
    "user_id": 1002,
    "user_login": "hubot",
    "day": "2026-05-14",
    "organization_id": "999",
    "team_id": 42,
    "slug": "frontend"
  }
]
```
Vous obtenez ainsi une vue d’ensemble des utilisateurs d’une organisation, de leurs équipes et de leurs tags d’équipe.

### Endpoints au niveau utilisateur

Les rapports au niveau utilisateur offrent une visibilité plus granulaire de l’adoption. Vous pouvez comprendre, à un niveau élevé, comment chaque personne utilise Copilot.

Endpoints courants :

- 
`users-1-day`
- 
`users-28-day`
- 
`user-teams-1-day`

Ces rapports aident les administrateurs à identifier :

- Les utilisateurs très actifs
- Les équipes à faible adoption
- Les besoins de formation
- Les tendances d’usage par département

Ces requêtes ressemblent beaucoup aux rapports 1 jour et 28 jours au niveau organisation, en pointant simplement vers un autre endpoint.

#### Rapport utilisateur sur une journée

Exemple d’appel `users-1-day` :

```
curl -L \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer <YOUR-TOKEN>" \
  -H "X-GitHub-Api-Version: 2026-03-10" \
  "https://api.github.com/enterprises/ENTERPRISE/copilot/metrics/reports/users-1-day?day=DAY"
```
#### Rapport utilisateur sur 28 jours

Exemple d’appel `users-28-day` :

```
curl -L \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer <YOUR-TOKEN>" \
  -H "X-GitHub-Api-Version: 2026-03-10" \
   https://api.github.com/enterprises/ENTERPRISE/copilot/metrics/reports/users-28-day/latest
```
#### Rapport user-teams sur une journée

Un endpoint `user-teams-1-day` existe également et associe chaque utilisateur à ses équipes. Il ne contient pas de métriques d’usage : il sert de clé de jointure pour agréger des données par équipe.

#### Structure des rapports au niveau utilisateur

Le niveau de détail est bien plus élevé, puisqu’il s’agit de l’usage d’un utilisateur donné :

```
[{
  "code_acceptance_activity_count": 1,
  "code_generation_activity_count": 1,
  "day": "2025-10-01",
  "enterprise_id": "1",
  "loc_added_sum": 8,
  "loc_deleted_sum": 0,
  "loc_suggested_to_add_sum": 10,
  "loc_suggested_to_delete_sum": 0,
  "totals_by_cli": {
    "last_known_cli_version": {
      "cli_version": "1.0.8",
      "sampled_at": "2025-10-01T00:01:43.000Z"
    },
    "prompt_count": 2,
    "request_count": 2,
    "session_count": 2,
    "token_usage": {
      "avg_tokens_per_request": 4400.0,
      "output_tokens_sum": 5000,
      "prompt_tokens_sum": 3800
    }
  },
  "totals_by_feature": [{
    "code_acceptance_activity_count": 1,
    "code_generation_activity_count": 1,
    "feature": "code_completion",
    "loc_added_sum": 8,
    "loc_deleted_sum": 0,
    "loc_suggested_to_add_sum": 10,
    "loc_suggested_to_delete_sum": 0,
    "user_initiated_interaction_count": 0
  }],
  "totals_by_ide": [{
    "code_acceptance_activity_count": 1,
    "code_generation_activity_count": 1,
    "ide": "vscode",
    "last_known_ide_version": {
      "ide_version": "1.85.0",
      "sampled_at": "2025-10-01T00:00:02.000Z"
    },
    "last_known_plugin_version": {
      "plugin": "",
      "plugin_version": "",
      "sampled_at": "2025-10-01T00:00:02.000Z"
    },
    "loc_added_sum": 8,
    "loc_deleted_sum": 0,
    "loc_suggested_to_add_sum": 10,
    "loc_suggested_to_delete_sum": 0,
    "user_initiated_interaction_count": 0
  }],
  "totals_by_language_feature": [{
    "code_acceptance_activity_count": 1,
    "code_generation_activity_count": 1,
    "feature": "code_completion",
    "language": "unknown",
    "loc_added_sum": 8,
    "loc_deleted_sum": 0,
    "loc_suggested_to_add_sum": 10,
    "loc_suggested_to_delete_sum": 0
  }],
  "totals_by_language_model": [],
  "totals_by_model_feature": [],
  "used_agent": false,
  "used_chat": false,
  "used_cli": true,
  "user_id": 1,
  "user_login": "login1",
  "user_initiated_interaction_count": 0,
  "etl_id": "green",
  "day_partition": "2025-10-01",
  "entity_id_partition": 1
}]
```
Ces métriques sont surtout utiles comme signaux d’adoption au niveau équipe. Les taux d’acceptation et les volumes d’usage sont des indicateurs opérationnels, pas des mesures de qualité des développeurs.

Pour consulter l’ensemble potentiel des métriques exposées, reportez-vous à la documentation des données de métriques d’usage GitHub la plus à jour.

Les rapports au niveau utilisateur incluent les interactions CLI. Si vos équipes utilisent Copilot en ligne de commande, notre GitHub Copilot CLI Tutorial couvre la mise en place et les workflows courants.

## Mettre en place un workflow de reporting Copilot

Appeler l’API à la main est utile pour expérimenter et comprendre le schéma. Pour passer à l’action, mieux vaut automatiser.

Les équipes qui tirent le plus de valeur de Copilot Enterprise bâtissent généralement des pipelines légers de reporting qui croisent la télémétrie d’usage avec leurs métriques d’ingénierie internes.

### Les indicateurs clés pour démontrer le ROI

Toutes les métriques Copilot ne se valent pas. Les plus utiles incluent :

- Progression du nombre d’utilisateurs actifs
- Tendance des taux d’acceptation
- Code proposé vs code conservé
- Réduction des temps de cycle des PR
- Fréquence d’usage des IDE

GitHub a publié des benchmarks tels que :

- Achèvement des tâches 55 % plus rapide
- 88 % de code conservé

Ces chiffres indiquent des gains de productivité significatifs. Vos résultats varieront selon les équipes et workflows, d’où l’intérêt de l’API des métriques d’usage. Une équipe backend infrastructure n’utilise pas Copilot comme une équipe frontend de prototypage.

### Des données brutes à un dashboard équipe

Un workflow de reporting léger ressemble souvent à ceci :

1. Appel API planifié
2. Stockage des réponses dans une base ou un tableur
3. Transformation en tables de reporting
4. Visualisation dans votre plateforme de BI

La pile technique compte moins que la régularité.

Même un simple enchaînement de scripts Python planifiés et d’exports CSV peut offrir une visibilité opérationnelle utile.

Architecture type :

GitHub API

↓

Script Python planifié

↓

