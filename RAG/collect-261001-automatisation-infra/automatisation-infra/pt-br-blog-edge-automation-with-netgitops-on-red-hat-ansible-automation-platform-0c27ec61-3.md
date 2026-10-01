---
id: collect-261001-automatisation-infra/automatisation-infra/pt-br-blog-edge-automation-with-netgitops-on-red-hat-ansible-automation-platform-0c27ec61-3
title: "pt-br-blog-edge-automation-with-netgitops-on-red-hat-ansible-automation-platform-0c27ec61"
domain: automatisation-infra
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["cost", "nvidia"]
source: docs/RAG/collect-261001-automatisation-infra/pt-br-blog-edge-automation-with-netgitops-on-red-hat-ansible-automation-platform-0c27ec61.md
source_anchor: ""
source_lines: [166, 218]
sha256: 74795860ce20d3663de79e02a7fef8e1e44dcca699c3414a761db4f8a6b54645
---

# pt-br-blog-edge-automation-with-netgitops-on-red-hat-ansible-automation-platform-0c27ec61

  - Go back to the repo
  - Click on “Settings”
  - Click on “Webhooks”, then “Add webhook”
  - Fill with the following info:
  - Payload URL: The webhook URL from previous step (see image above)
  - Content Type: application/json
  - Secret: Use the webhook key from previous step (see image above)
  - Select the events that trigger the webhook. In our case, we selected just “push”.
  - Leave “Active” checked.
Putting everything together
Let’s have a look at the original state in the SDN dashboard. One SSID is wrong: it should be MyCompany_customers, not MyCompany_clustomers. We need to correct this typo.
This means to modify the network_vars.yml file locally, we must make a commit and push it to the repo. 
The execution can be observed in the video
Triggering Ansible automation from GitHub where we can observe how a job is triggered from GitHub in Ansible Automation Platform, which corresponds to the job template with our playbook and how it changes the environment.
As a final step, let’s go back to the SDN controller dashboard to confirm the changes.
Now the customer’s SSID has a correct name.
What’s next?
At Red Hat, we advocate a “start small, think big” approach for customers that are starting their automation journeys. Some of the most effective beginner tasks are often simpler tasks that are repeated often, and automating them helps users practice automation skills and learn lessons to build up to more complex use cases in the future.
Thinking beyond network configuration automation
We covered the steps required to build a very simple use case for SDN automation through the SD-WAN controller, for a network edge WiFi configuration at scale.
If you are beginning your automation journey, focusing on automating your network configuration tasks is a great starting point to gain consistency, scalability and reduce risk due to human-error. However, configuration automation can only solve a small portion of your IT provisioning workflow.
Frequently, IT processes require approvals, validations, ticket opening/update/closure, and a rollback strategy. There are several improvements or extensions you can make in this scenario:
- Integrate it with an ITSM, such as ServiceNow, to allow you to document every step and change made.
- Trigger a workflow job template, or more complex scenarios, to be executed via webhooks.
- Use automation services catalog in Ansible Automation Platform to launch these tasks with the required approvals.
- Perform verifications before applying changes, as a way to have a configuration drift approach.
If you want to recreate a similar scenario, the code is available in this GitHub repo, and you can sign up for a free trial of Ansible Automation Platform to recreate the steps covered.
Ansible Automation Platform workshops are also available at no-cost to learn how to build and execute network automation in a multi-vendor fashion.
If you are interested in learning more about network automation use cases, we have a good video summary here.
Acknowledgement: I’d like to thank Dafné Mendoza for her outstanding contributions to this blog.
Sobre o autor
Mais como este
O Red Hat Device Edge já está disponível para execução no NVIDIA Jetson Orin
Superando desafios: IA e Kubernetes na fonte dos dados
Untangling Networks | Compiler
Infrastructure At The Edge | Compiler
Navegue por canal
Automação
Últimas novidades em automação de TI para empresas de tecnologia, equipes e ambientes
Inteligência artificial
Descubra as atualizações nas plataformas que proporcionam aos clientes executar suas cargas de trabalho de IA em qualquer ambiente
Nuvem híbrida aberta
Veja como construímos um futuro mais flexível com a nuvem híbrida
Segurança
Veja as últimas novidades sobre como reduzimos riscos em ambientes e tecnologias
Edge computing
Saiba quais são as atualizações nas plataformas que simplificam as operações na borda
Infraestrutura
Saiba o que há de mais recente na plataforma Linux empresarial líder mundial
Aplicações
Conheça nossas soluções desenvolvidas para ajudar você a superar os desafios mais complexos de aplicações
Virtualização
O futuro da virtualização empresarial para suas cargas de trabalho on-premise ou na nuvem
