<!--
SPDX-FileCopyrightText: The RamenDR authors
SPDX-License-Identifier: Apache-2.0
-->

# kubectl-ramen

The kubectl ramen plugin provide high level commands for managing ramen.

## Commands

- `clusterset` - manage clustersets configurations
- `deploy` - deploy ramen in a clusterset
- `undeploy` - undeploy ramen in a clusterset
- `enable` - enable DR for an OCM managed workload
- `disable` - disable DR for an OCM managed workload
- failover - failover DR enabled workload to another cluster in the
  clusterset.
- relocate - relocate DR enabled workload to another cluster in the
  clusterset.
- `status` - print the status of a ramen protected workload in a
  clusterset.

## Status

This is work in progress.

## License

kubectl-ramen is under the [Apache 2.0](LICENSES/Apache-2.0.txt)
license.
