# Development in our Kubermatic Fork

This is a fork of Loodse's kubermatic upstream repository (https://github.com/kubermatic/kubermatic).

Use `api/hack/sys11-run-api.sh` and `api/hack/sys11-run-dashboard-and-api.sh` scripts for for launching the
API or API and dashboard locally.

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