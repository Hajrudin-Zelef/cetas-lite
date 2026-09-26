---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1rubyij-looking-for-a-100-free-ai-agent-that-can-control-65d574a7-3
title: "r-localllama-comments-1rubyij-looking-for-a-100-free-ai-agent-that-can-control-65d574a7"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["agent", "agents", "mcp", "open source", "qwen"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1rubyij-looking-for-a-100-free-ai-agent-that-can-control-65d574a7.md
source_anchor: ""
source_lines: [55, 59]
sha256: 39564e71a9be8af0a20a337e7fbccd0f0dc3705e782e2232c3aec3934d460703
---

# r-localllama-comments-1rubyij-looking-for-a-100-free-ai-agent-that-can-control-65d574a7

Excellente question. Il est important de noter que "gratuit" cache souvent des coûts matériels significatifs. Les agents de navigateur utilisant des modèles de vision (Qwen, Gemma) ont besoin de 16 Go de VRAM ou plus pour être utilisables, et peuvent consommer entre 20 000 et 200 000 jetons par affichage de page. Pour des tâches complexes, même les modèles de plus de 30 milliards de paramètres ont du mal. Une approche plus pratique : utiliser des arbres d'accessibilité au lieu de captures d'écran - cela permet aux petits modèles (~7 milliards) de travailler avec des données DOM structurées, réduisant ainsi de manière spectaculaire les besoins en calcul et l'utilisation des jetons. Envisagez également des approches hybrides : local pour la navigation simple, API cloud uniquement pour le raisonnement complexe afin d'équilibrer coût et capacité.
Que dire de anythingLLM et deepagents/langchain ? Ou goose ai ou smolagents avec l'outil de navigateur par défaut.
OpenCode + Navigateur Stealth MCP
l'utilisation du navigateur est probablement l'option open source la plus activement maintenue en ce moment... elle s'associe à n'importe quel modèle local compatible openai via ollama, donc vous pouvez l'exécuter entièrement gratuitement si vous avez le matériel... playwright gère le contrôle réel du navigateur en dessous
PageAgent - The GUI Agent vivant dans votre page Web
