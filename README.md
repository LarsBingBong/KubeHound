# KubeHound

A Kubernetes attack graph tool

Full documentation available on confluence: https://datadoghq.atlassian.net/wiki/spaces/ASE/pages/2871592134/KubeHound+1.0

## Run

### Prerequisites - setup the infrastructure

To run the application, you can use docker image with the compose. First create and populate a .env file with the required variables:

```bash
cp deployments/kubehound/.env.tpl deployments/kubehound/.env
```
Then, edit the variables (datadog env `DD_*` related and `KUBEHOUND_ENV`):

* `KUBEHOUND_ENV`: `dev` or `prod`
* `DD_API_KEY`: api key you created from https://app.datadoghq.com/ website

Note:
* `KUBEHOUND_ENV=prod` will use prebuilt image from ghcr.io (:rotating_light: currently NOT supported :rotating_light:)
* `KUBEHOUND_ENV=dev` will build the images locally

To target a specific cluster there are 2 options:
* Select the targeted cluster via `kubectx` (need to be installed separately)     
* Use a specific kubeconfig file by exporting the env variable: `export KUBECONFIG=/your/path/to/.kube/config`

### Run Kubehound - Automated way

To run kubehound the easy way, just run:

```bash
make kubehound
```

This will do the following:
* Start the backend services via docker compose (wiping any existing data)
* Compile the kubehound binary from source
* Run the kubehound binary using the default configuration

### Run Kubehound - Manual way

To replicate the automated command and run KubeHound step-by-step. First build the application:

```bash
make build
```

Next spawn the backend infrastructure

```bash
make backend-up
```

Finally run the KubeHound binary

```bash
make run
```

Remember the targeted cluster must be set via `kubectx` or setting the `KUBECONFIG` environment variable. Additional functionality for managing the application can be found via:

```bash
make help
```

### Using KubeHound Data

To query the KubeHound graph data requires using the [Gremlin](https://tinkerpop.apache.org/gremlin.html) query language via an API call or dedicated graph query UI. A number of graph query UIs are availble, but we recommend [gdotv](https://gdotv.com/). To access the KubeHound graph using `gdotv`:

+ Download and install the application from https://gdotv.com/
+ Create a connection to the local janusgraph instance by following the steps here https://docs.gdotv.com/connection-management/ and using `hostname=localhost`
+ Navigate to the query editor and enter a sample query e.g `g.V().count()`. See detailed instructions here: https://docs.gdotv.com/query-editor/#run-your-query
+ See the provided [cheatsheet](./pkg/kubehound/graph/CHEATSHEET.md) for examples of useful queries for various use cases.


## Build

Build the application via:

```bash
make build
```

All binaries will be output to the [bin](./bin/) folder

## Unit Testing

The full suite of unit tests can be run locally via:

```bash
make test
```

## System Testing

The repository includes a suite of system tests that will do the following:
+ create a local kubernetes cluster
+ collect kubernetes API data from the cluster
+ run KubeHound using the file collector to create a working graph database
+ query the graph database to ensure all expected vertices and edges have been created correctly

The cluster setup and running instances can be found under [test/setup](./test/setup/)

If you need to manually access the system test environement with kubectl and other commands, you'll need to set (assuming you are at the root dir):
```bash
cd test/setup/ && export KUBECONFIG=$(pwd)/.kube-config
```

### Requirements

+ Kind: https://kind.sigs.k8s.io/docs/user/quick-start/#installing-with-a-package-manager
+ Kubectl: https://kubernetes.io/docs/tasks/tools/

#### Environment variable:
- `DD_API_KEY` (optional): set to the datadog API key used to submit metrics and other observability data.

### Setup

Setup the test kind cluster (you only need to do this once!) via:

```bash
make local-cluster-deploy
```

Then run the system tests via:

```bash
make system-test
```

To cleanup the environment you can destroy the cluster via:

```bash
make local-cluster-destroy
```

To list all the available commands, run:

```bash
make help
```

Note: if you are running on Linux but you dont want to run `sudo` for `kind` and `docker` command, you can overwrite this behavior by editing the following var in `test/setup/.config`:
* `DOCKER_CMD="docker"` for docker command
* `KIND_CMD="kind"` for kind command 

### CI Testing

System tests will be run in CI via the [system-test](./.github/workflows/system-test.yml) github action 

### Sample Graph

To view a sample graph demonstrating attacks in a very, very vulnerable cluster you can generate data via the system tests:

```bash
make local-cluster-deploy && make system-test
```

Then use a graph visualizer of choice (we recommend [gdotv](https://gdotv.com/)) to connect to localhost and view and query the sample data. **NOTE** you must change the port to `8183` (rather than default `8182`)

