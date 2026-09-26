---
id: collect-240926-datacamp/datacamp/dagster-vs-airflow-comparaison-des-meilleurs-outils-dorchestration-de-donnees-pour-des-dat-1
title: "dagster-vs-airflow-comparaison-des-meilleurs-outils-dorchestration-de-donnees-pour-des-data-stacks-m"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "benchmark", "cost", "license", "open source"]
source: docs/RAG/clean_en/datacamp/dagster-vs-airflow-comparaison-des-meilleurs-outils-dorchestration-de-donnees-pour-des-data-stacks-m.md
source_anchor: ""
source_lines: [1, 116]
sha256: 623a01d134dc0c91c29a936b2aa506fca23d432e789728556216a744672546cc
---

# dagster-vs-airflow-comparaison-des-meilleurs-outils-dorchestration-de-donnees-pour-des-data-stacks-m

<!-- source: https://www.datacamp.com/fr/blog/dagster-vs-airflow -->

Course

Every time you scroll through LinkedIn, it feels like a new tool appears in an already crowded data landscape. As data teams stack up the building blocks of their "data stack," they need a way to manage and connect these disparate tools.

Data orchestration refers to designing, executing, and monitoring processes that combine, transform, and organize data across the entire ecosystem. To do this, data engineers rely on data orchestration tools. Here, we review two of the most popular solutions: Apache Airflow and Dagster. Let's dive in!

## What is Airflow?

Airflow is the market benchmark for building, running, and monitoring data pipelines as code. Initially developed in 2014 by Airbnb's Data Engineering team, the project has since been adopted by the Apache Software Foundation and has become the most popular offering under their license.

Each month, Airflow is downloaded more than thirty million times and benefits from a highly active community of users and contributors. In Airflow, data pipelines are called DAGs, for directed acyclic graphs.

To learn more, check out our tutorial on getting started with Apache Airflow.

## What is Dagster?

Like Airflow, Dagster is an orchestration tool that allows you to define data pipelines as code. Open source, it was launched in 2019. Dagster uses Python to define "assets," the basic building blocks of a data pipeline.

Dagster has established itself as a competitor to Airflow in a relatively uncrowded space. The project is maintained by its open source community and benefits from commercial support from Dagster Labs.

## Connecting a modern data stack

So, where do Airflow and Dagster fit into a modern data stack? Let's start with the architecture diagram below.

Architecture diagram showing a sample of common tools in a modern data stack.

This diagram illustrates what many consider a standard data stack. It consists of source systems, data warehouse and transformation tools, as well as downstream "destinations" like Tableau and Looker.

Here, Airflow and Dagster are the arrows that connect this data ecosystem. Without an orchestration tool, it is difficult to move data from one tool to another.

Tools like Airflow and Dagster also add a layer of observability, to easily understand where, how, and when data is routed from source to destination, and how it is transformed along the way.

## Dagster vs Airflow: key features

As you may have guessed, Airflow and Dagster share many commonalities in their usage and role in a modern data stack. Among them:

- Ability to define pipelines as code
- Native integrations with the modern data stack
- Local development experience

However, each has its own strengths and unique features.

### Airflow

Let's start with Airflow's main features:

#### Understanding Airflow and DAGs

In Airflow, data pipelines are called DAGs — for directed acyclic graphs. Think of a DAG as a set of tasks connected in a specific order. The most basic component of a DAG is the task. For example, in an extract, transform, load (ETL) pipeline, the "transform" step is a task in its own right.

These tasks are generally defined via Airflow operators. But, as you will see, there is another way to define tasks thanks to the TaskFlow API.

#### Scheduling DAGs in Airflow

Scheduling a pipeline in Airflow is remarkably flexible. Need to run it every day? No problem. The first Friday of each month? Easy too. And if your manager asks you to trigger a DAG when a dataset is updated? Airflow can do that as well.

Airflow offers many customization options for DAG scheduling. With CRON, Timetables, and data-aware scheduling, you can run your DAGs whenever you need.

#### TaskFlow API

To make writing DAGs more accessible to data professionals, the Airflow community introduced the TaskFlow API. Rather than using only operators, the TaskFlow API allows you to define tasks simply by decorating functions. This greatly facilitates sharing data between tasks and defining dependencies.

Moreover, tools like the Astro SDK have been built on top of the TaskFlow API, further expanding its possibilities.

### Dagster

Now let's see how Dagster positions itself:

#### Dagster's asset-centric approach

Dagster takes an asset-based approach to building pipelines. In Dagster, any persistently stored data object — file, table, etc. — is an asset.

These assets are defined in code via Python functions. When they run, Dagster automatically creates dependencies and materializes the asset. This asset-centric approach makes tracking the production and consumption of data within a pipeline particularly simple.

#### Understanding ops in Dagster

Another key concept in Dagster: "ops," quite similar to tasks in Airflow. Ops represent individual steps of a pipeline, with inputs and outputs that can be assets. Even better, these inputs and outputs can be typed to clearly specify the data being processed.

What is the relationship between ops and assets? Ops correspond to the steps (extract, transform, load) that operate on assets, that is, the data objects themselves.

#### Dagster's type system

To clarify the data flows of a pipeline, Dagster relies on a robust type system to validate the inputs and outputs of each op. Airflow supports Python types, but Dagster has made it a pillar of writing pipelines.

You can type not only ops, but also asset definitions, to ensure that the data produced by an asset is correct and validated.

## Airflow vs Dagster: Developer Experience

When choosing an orchestration tool, one often evaluates feature richness, total cost of ownership, and scalability. That said, it is crucial to consider the developer experience. An intuitive, efficient, and robust experience promotes rapid iteration and anticipates issues like unit testing and dependency management.

### Airflow

#### The Airflow Community

Have an Airflow question or bug you can't resolve? There's a good chance the answer already exists. Airflow has one of the largest Slack communities among Apache projects and has more than three thousand contributors. Thousands of blogs, tutorials, and courses have been dedicated to Airflow. In addition to helping users, the community builds and maintains the project itself.

Astronomer, the commercial vendor behind Apache Airflow, has elevated the developer experience to another level. The astro CLI adds a layer of abstraction and features to Airflow's utilities. Astronomer also offers its registry and the "Ask Astro" chatbot to answer any Airflow questions a developer might have.

#### Developing and testing locally

One of Airflow's major strengths is the ability to develop and test locally. With the airflow CLI, a developer can launch an instance in seconds. They then modify the code locally and test via the UI or CLI. The result: rapid iteration and shortened delivery cycles, especially for heavily Airflow-oriented teams.

#### New tools for developing DAGs

For a time, the only way to define an Airflow DAG was through traditional operators. The TaskFlow API changed that by offering an alternative. Now, tools like dag-factory and gusty make writing and running DAGs easier than ever.

### Dagster

Dagster's local development experience is very similar to Airflow's, with the ability to develop, test, and run pipelines locally. Both tools support common enterprise SDLC practices, such as CI/CD.

However, it is difficult to compare the scale of the Airflow community to that of Dagster. Where Airflow has thousands of community-generated resources, Dagster still has to close the gap. When developing and troubleshooting pipelines with Dagster, practitioners may struggle to quickly find the resources needed to implement a solution.

#### Easy to test

