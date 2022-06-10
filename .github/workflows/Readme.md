# github action

## lint

    using golangci-lint to lint code

## auto release

    using goreleaser to auto release
    需要通过repo/代码仓库的settings -> Secrets -> Actions -> New repository Secret新增一个token,
    token通过个人设置页的settings ->Developer Settings -> Personal Access Tokens -> Generate New Token获取

## references

1. lint <https://github.com/golangci/golangci-lint-action>
2. auto release <https://github.com/goreleaser/goreleaser-action>
3. blog <https://xiaozhou.net/auto-ci-cd-via-github-actions-2022-05-02.html?utm_source=pipecraft.net>
