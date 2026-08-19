+++
title = "Reliability of our Streaming Pipelines - How we took Flink CDC completeness from 92% to 99.78%"

[extra]
speaker = "abhishek-jain"
+++

In data infrastructure, streaming pipelines capture every database change (inserts, updates, deletes) in real time to keep downstream systems in sync.

We thought ours were 92% reliable, until we realized our measurement was double-counting some changes, hiding actual data loss. Global Transaction Identifiers (GTIDs), unique tags for each database transaction, we achieved 99.78% completeness.

This talk covers the pitfalls of duplicate counting, how GTIDs exposed the gaps, and the steps we took to harden our [Apache Flink](https://flink.apache.org/) pipelines.
