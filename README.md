<!--
SPDX-FileCopyrightText: The RamenDR authors
SPDX-License-Identifier: Apache-2.0
-->

# kubectl-ramen

The kubectl ramen plugin provides high level commands for managing
[Ramen](https://github.com/RamenDR/ramen).

## Commands

- `clusterset` - manage clustersets configurations
- `deploy` - deploy ramen in a clusterset
- `undeploy` - undeploy ramen in a clusterset
- `enable` - enable DR for an OCM managed workload
- `disable` - disable DR for an OCM managed workload
- `failover` - failover DR enabled workload to another cluster in the
  clusterset.
- `relocate` - relocate DR enabled workload to another cluster in the
  clusterset.
- `status` - print the status of a ramen protected workload in a
  clusterset.

## Example usage

```shell
# Add clusterset from drenv environment file
kubectl ramen clusterset add-env rdr --env-file regional-dr.yaml

# Add clusterset from kubeconfigs
kubectl ramen clusterset add-cfg rdr --hub hub.cfg --cluster1 c1.cfg --cluster2 c2.cfg

# Deploy ramen in clusterset "rdr"
kubectl ramen deploy --clusterset rdr

# Enable DR for application "busybox-sample"
kubectl ramen enable --clusterset rdr --namespace busybox-sample

# Print application "busybox-sample" status
kubectl ramen status --clusterset rdr --namespace busybox-sample

# Watch application "busybox-sample" status
kubectl ramen status --clusterset rdr --namespace busybox-sample --watch

# Failover application to the other managed cluster
kubectl ramen failover --clusterset rdr --namespace busybox-sample

# Relocate application to the other managed cluster
kubectl ramen relocate --clusterset rdr --namespace busybox-sample

# Disable DR for the busybox-sample application
kubectl ramen disable --clusterset rdr --namespace busybox-sample

# Undeploy ramen in clusterset "rdr"
kubectl ramen undeploy --clusterset rdr
```

## Status

This is work in progress; only some commands are implemented.

## License

kubectl-ramen is under the [Apache 2.0](LICENSES/Apache-2.0.txt)
license.
