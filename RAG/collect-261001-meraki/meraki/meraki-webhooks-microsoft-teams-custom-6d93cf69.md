---
id: collect-261001-meraki/meraki/meraki-webhooks-microsoft-teams-custom-6d93cf69
title: "meraki-webhooks-microsoft-teams-custom-6d93cf69"
domain: meraki
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/meraki-webhooks-microsoft-teams-custom-6d93cf69.md
source_anchor: ""
source_lines: [1, 53]
sha256: 3be19d6d96b8c5598c72635c0c10e17f7dc2654d61ba409bedaebb2359ace631
---

# meraki-webhooks-microsoft-teams-custom-6d93cf69

Post a webhook alert into a Microsoft Teams channel.
You will first need to WebHook enable the channel. This process creates the HTTP server URL to enter into the Meraki Dashboard.
https://your-domain.atlassian.net/rest/api/2/issue/createmeta
body.liquid
Copy
{
    "@type": "MessageCard",
    "@context": "https://schema.org/extensions",
    "summary": "{{organizationName}}",
    "sections": [{
        "activityTitle": "{{organizationName}}",
        "activitySubtitle": "{{networkName}}",
        {% if alertData.imageUrl %}"images": [{"image":"{{alertData.imageUrl}}"}],{% endif %}
        "facts": [
    {
            "name": "When",
            "value": "{{occurredAt}}"
    },
        {
            "name": "Device",
            "value": "{{deviceName}}"
        },
        {
            "name": "Alert",
            "value": "{{alertType}} - {{alertLevel}}"
    }{% unless alertData.imageUrl %},
    {
            "name": "",
            "value": "{{alertData | json_markdown}}"
    }{% endunless %}],
    "markdown": true
  }],
    "potentialAction": [{
        "@type": "OpenUri",
        "name": "Org",
        "targets": [{"os": "default", "uri" : "{{organizationUrl}}"}]
    },
    {
        "@type": "OpenUri",
        "name": "Network",
        "targets": [{"os": "default", "uri" : "{{networkUrl}}"}]
    },
    {
        "@type": "OpenUri",
        "name": "Device",
        "targets": [{"os": "default", "uri" : "{{deviceUrl}}"}]
    }]
}
Lorsque vous consultez un site Web, celui-ci peut enregistrer ou récupérer des informations présentes dans votre navigateur, le plus souvent sous la forme de témoins. Ces informations peuvent vous concerner, ainsi que vos préférences ou votre périphérique, et servent principalement à faire fonctionner le site selon vos attentes. Généralement, ces informations ne permettent pas de vous identifier directement, mais elles sont utilisées afin de vous offrir une expérience Web personnalisée. Parce que nous respectons votre droit à la vie privée, vous pouvez choisir de ne pas autoriser certains types de témoins. Cliquez sur les différents en-têtes de catégorie pour en apprendre davantage et modifier nos paramètres par défaut. Toutefois, bloquer certains types de témoins peut avoir une incidence sur votre expérience du site et les services que nous proposons.
Ces cookies sont indispensables au bon fonctionnement du site web et ne peuvent pas être désactivés au niveau de nos systèmes. Ils ne résultent généralement que d'actions que vous avez effectuées et qui correspondent à une demande de service, comme lorsque vous définissez vos préférences en matière de confidentialité, que vous vous connectez ou que vous remplissez des formulaires. Vous pouvez configurer votre navigateur pour qu'il bloque ces cookies ou vous avertisse de leur présence, mais certaines parties du site ne seront alors plus opérationnelles. Ces cookies n'enregistrent aucune information d'identification personnelle.
Ces cookies fournissent des indicateurs relatifs aux performances et à la facilité d'utilisation de notre site. Ils visent principalement à récolter des informations sur la façon dont vous interagissez sur le site, notamment le temps de chargement des pages, les temps de réponse, les messages d'erreur et le déroulement des interactions pour chaque visiteur. Nous pouvons ainsi examiner et analyser le comportement des visiteurs pour améliorer l'ergonomie et les fonctionnalités de notre site. Ces cookies nous permettent également de comptabiliser les visites et les sources de trafic afin de mesurer et d'améliorer les performances du site. Ils nous aident à savoir quelles pages sont les plus et les moins populaires, et à voir la manière dont les visiteurs naviguent sur le site. Si vous n'autorisez pas ces cookies, nous ne saurons pas si vous avez déjà visité le site et ne pourrons pas en mesurer les performances.
Ces cookies peuvent être intégrés à notre site par nos partenaires publicitaires. Ils peuvent être utilisés par ces entreprises pour établir un profil de vos centres d'intérêt et vous présenter des publicités pertinentes sur d'autres sites. Si ces cookies n'enregistrent pas directement vos informations personnelles, ils identifient de manière unique votre navigateur et votre terminal d'accès à Internet. Si vous n'autorisez pas ces cookies, des publicités moins ciblées s'afficheront.
Ces cookies permettent d’améliorer et de personnaliser les fonctionnalités du site Web. Ils peuvent être activés par nos équipes, ou par des tiers dont les services sont utilisés sur les pages de notre site Web. Si vous n'acceptez pas ces cookies, une partie ou la totalité de ces services risquent de ne pas fonctionner correctement.
