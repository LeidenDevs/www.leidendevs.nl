+++
title = "Mapping the Cloud with Elixir"

[extra]
speaker = "peter-ullrich"
+++


The “Cloud” sometimes feels like this abstract thing you deploy your server to and otherwise forget about. But what does it look like? What pathways connect this globally spanning network of datacenters? I used Elixir to find out.

I deployed an Erlang cluster of tiny probes across 100+ datacenters on AWS, GCP, and Azure, collecting months of traceroute data in TimescaleDB to map the hidden topology of the cloud. This talk explores how I deployed the probes efficiently and cheaply, what the data revealed, and the practical challenges of running a globally distributed Elixir system.

## Key Takeaways

* How to deploy your Elixir release using Burrito and Terraform reliably and at scale.
* How traceroute works.
* How “the cloud” looks like.

## Target Audience

Anyone who has touched the cloud

