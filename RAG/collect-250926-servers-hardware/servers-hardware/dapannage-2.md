---
id: collect-250926-servers-hardware/servers-hardware/dapannage-2
title: "DÃ©pannage"
domain: servers-hardware
role: reference
task: reference
actors: ["Anthropic", "Google", "OpenAI"]
dates: []
keywords: ["gemini", "kimi"]
source: docs/RAG/clean4/dapannage.md
source_anchor: ""
source_lines: [128, 178]
sha256: 6b712255c8df8cb2ff730bf53c531bd8562376c9b4677bf2a89cd374f0feead2
---

# DÃ©pannage

1. VÃ©rifiez que vous Ãªtes authentifiÃ© auprÃ¨s du fournisseur
2. VÃ©rifiez que le nom du modÃ¨le dans votre configuration est correct
3. Certains modÃ¨les peuvent nÃ©cessiter un accÃ¨s ou des abonnements spÃ©cifiques

Si vous rencontrez `ProviderModelNotFoundError`, vous avez probablement tort
faire rÃ©fÃ©rence Ã  un modÃ¨le quelque part.
Les modÃ¨les doivent Ãªtre rÃ©fÃ©rencÃ©s comme suitÂ : `<providerId>/<modelId>`

ExemplesÂ :

- `openai/gpt-4.1`
- `openrouter/google/gemini-2.5-flash`
- `opencode/kimi-k2`

Pour dÃ©terminer Ã  quels modÃ¨les vous avez accÃ¨s, exÃ©cutez `opencode models`

Si vous rencontrez une ProviderInitError, vous avez probablement une configuration non valide ou corrompue.

Pour rÃ©soudre ce problÃ¨meÂ :

1. 
Tout dâabord, vÃ©rifiez que votre fournisseur est correctement configurÃ© en suivant le guide du fournisseur
2. 
Si le problÃ¨me persiste, essayez dâeffacer votre configuration stockÃ©eÂ :

Sur Windows, appuyez sur `WIN+R` et supprimezÂ : `%USERPROFILE%\.local\share\opencode`

1. RÃ©-authentifiez-vous auprÃ¨s de votre fournisseur Ã  lâaide de la commande `/connect` dans le TUI.

Si vous rencontrez des erreurs dâappel API, cela peut Ãªtre dÃ» Ã des packages de fournisseurs obsolÃ¨tes. opencode installe dynamiquement les packages du fournisseur (OpenAI, Anthropic, Google, etc.) selon les besoins et les met en cache localement.

Pour rÃ©soudre les problÃ¨mes liÃ©s au package du fournisseurÂ :

1. 
Videz le cache du package du fournisseurÂ :

Sur Windows, appuyez sur `WIN+R` et supprimezÂ : `%USERPROFILE%\.cache\opencode`

1. RedÃ©marrez opencode pour rÃ©installer les derniers packages du fournisseur

Cela forcera opencode Ã tÃ©lÃ©charger les versions les plus rÃ©centes des packages du fournisseur, ce qui rÃ©sout souvent les problÃ¨mes de compatibilitÃ© avec les paramÃ¨tres du modÃ¨le et les modifications de API.

Les utilisateurs de Linux doivent disposer de lâun des utilitaires de presse-papiers suivants installÃ©s pour que la fonctionnalitÃ© copier/coller fonctionneÂ :

**Pour les systÃ¨mes X11Â :**

**Pour les systÃ¨mes WaylandÂ :**

**Pour les environnements sans tÃªteÂ :**

opencode dÃ©tectera si vous utilisez Wayland et prÃ©fÃ©rez `wl-clipboard`, sinon il essaiera de trouver les outils du presse-papiers dans lâordreÂ : `xclip` et `xsel`.
