# Development in our Kubermatic Fork

This is a fork of Loodse's kubermatic upstream repository (https://github.com/kubermatic/kubermatic).

Use `api/hack/sys11-run-api.sh` and `api/hack/sys11-run-dashboard-and-api.sh` scripts for for launching the
API or API and dashboard locally. There's also `sys11-run-controller.sh`, `sys11-run-master-controller-manager.sh`,
and `sys11-run-userclustercontroller.sh` for running the kubermatic controller manager, master controller manager,
and user cluster controller manager locally.

See https://intra.syseleven.de/confluence/display/K8s/Repository%3A+kubermatic for a general overview.

## Branching Strategy

Upstream employs a modified release flow branching model – doing feature development directly in master or creating
feature branches (named "some-feature" or fix/some-bugfix) from it for larger changes and merging them back via pull
requests when they're done, and creating release branches (named "release/vx.y) from master for each major release,
cherry-picking bugfixes from master and tagging releases (tags named vx.y.z). The release branches will be closed and
not merged back to master once a new major release starts.

In this fork, created our own integration branch, named syseleven/release-master, from upstream release/v2.7.1 (the
latest upstream release at the time when we started working on our fork). Our internal development should be done right
there. When upstream tags new releases, we merge those into syseleven/release-master in a timely manner. For our own
releases, we can just tag those in syseleven/release-master directly.

In case we get into more heavy-lift development later on, we can essentially do release flow branching in our
syseleven/release-master branch, creating syseleven/feature/some-feature or syseleven/fix/some-bugfix branches as well
as syseleven/release/vx.y release branches from syseleven/release-master. If our code is of interest to upstream, we
cherry-pick the corresponding commits into our own "some-feature" branch created from upstream master and submit a pull
request.


## Mergin with upstream and releasing new version

New branch is created
```
git co syseleven/release-master
git pull
git checkout -b <username>/merge_<upstream-release-tag>
```

Upstream is updated
```
git fetch upstream
```

Merge is initiated
```
git merge <upstream-release-tag>
```

Resolving `/vendor` folder conflicts by merging dependencies file. `Gopkg.lock` and `Gopkg.toml` and running:
```
dep ensure
```

Conflicts in `/fixtures/` and `/apiclient/` folders can be ignored because they can be regenerated using scripts at `/hack/` directory.

Most of the rest of `*.go` files need to be resolved manually. Best is to go from down up in directory structure.

To check merge use `make lint`, `make test`, `make build` commands.

**Along with this repo `machine-controller` and `kubermatic-installer` repositories need to be updated.**

Note: Kubermatic has hardcode for machine-controller version at `kubermatic/api/pkg/resources/machine-controller/deployment.go`. It needs to be adjusted with `machine-controller` fork's merge.

### Updating machine-controller

`machine-controller` needs to be updated with upstream.

New release needs to be issued and hardcode at `kubermatic/api/pkg/resources/machine-controller/deployment.go` needs to be updated accordingly.

### Update kubermatic-installer

This steps may vary from release to realse.

Things to look at:

- `addons` which manually coppied from `kubermatic` repo to adjust for our needs
- configuration `*yaml.mako` files
