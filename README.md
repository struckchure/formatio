[![Docker Image CI](https://github.com/Overal-X/api.formatio/actions/workflows/container-ci.yaml/badge.svg)](https://github.com/Overal-X/api.formatio/actions/workflows/container-ci.yaml)

## Formatio

Formatio aims to imitate the behavior of virtual machines, but using docker containers in place of VMs.
This helps with fast setup of environments and portability.

With Formatio, you should be able to deploy your applications with little effort and optimal configurations at a minimum cost.
Surely, there would be support for automatic deployments from Git supported code respositories, like Github, Gitlab, etc.

> Formatio is a latin word for `formation`

# Run tests

```bash
$ go test -v ./tests/unit/... # unit tests
```

# Docker Build

For production

```bash
$ docker build -t api.formatio . \
  --target production \
  --build-arg infisical_token=<infisical_token> \
  --build-arg infisical_project_id=<infisical_project_id> \
  --build-arg infisical_env=<infisical_env>
```
