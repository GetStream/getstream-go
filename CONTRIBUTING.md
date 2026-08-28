# :recycle: Contributing

Contributions to this project are very much welcome, please make sure that your code changes are tested and that they follow Go best-practices.

## Getting started

### Required environmental variables

The tests require at least two environment variables: `STREAM_KEY` and `STREAM_SECRET`. There are multiple ways to provide that:
- simply set it in your current shell (`export STREAM_KEY=xyz`)
- you could use [direnv](https://direnv.net/)
- if you debug the tests in VS Code, you can set up an env file there as well: `"go.testEnvFile": "${workspaceFolder}/.env"`.

### Code formatting & linter

We enforce code formatting with [`gofumpt`](https://github.com/mvdan/gofumpt) (a stricter `gofmt`). If you use VS Code, it's recommended to set this setting there for auto-formatting:

```json
{
    "editor.formatOnSave": true,
    "gopls": {
        "formatting.gofumpt": true
    },
    "go.lintTool": "golangci-lint",
    "go.lintFlags": [
        "--fast"
    ]
}
```

Gofumpt will mostly take care of your linting issues as well.

### Pre-commit hook

We provide a pre-commit hook that runs the linter automatically before each commit. To enable it, run:

```bash
make setup-hooks
```

This only needs to be done once per local clone.

## Commit message convention

Since we're autogenerating our [CHANGELOG](./CHANGELOG.md), we need to follow a specific commit message convention.
You can read about conventional commits [here](https://www.conventionalcommits.org/). Here's how a usual commit message looks like for a new feature: `feat: allow provided config object to extend other configs`. A bugfix: `fix: prevent racing of requests`.

## Release (for Stream developers)

Releases are driven by [release-please](https://github.com/googleapis/release-please).

- Merge PRs to `main` with conventional-commit titles. The PR title becomes the commit subject and is what determines the next version, so a non-conventional title ships nothing.
- release-please keeps a Release PR open with the version bump and the generated changelog. Review it.
- The Release PR is opened by `github-actions[bot]`, so its CI runs are held at "action required". If the PR is behind `main`, clicking **Update branch** also releases them, because that commit is attributed to you; otherwise click **Approve and run**. It needs a code-owner approval like any other PR.
- Merge the Release PR. That creates the tag and the GitHub Release, and `proxy.golang.org` picks the tag up. There is no separate publish step.

Only `feat`, `fix`, `perf` and breaking changes produce a release (`revert` may also). A window of only `chore`, `ci`, `docs`, `test`, `refactor`, `style` or `build` commits produces no Release PR, which is intended.

### Forcing a specific version

Land a commit on `main` whose subject is conventional and whose body carries a `Release-As` footer. The footer also overrides the rule above, so it works even for a hidden commit type:

```bash
git commit --allow-empty -m "chore: republish as 5.3.1" -m "Release-As: 5.3.1"
```

If direct pushes to `main` are blocked, merge a PR and type `Release-As: 5.3.1` into the **commit message box** in the squash dialog. That box is blank by default and the PR description is never copied into it, so a footer left in the description is silently dropped.

### Hotfix while a release is pending

`main`'s next release includes everything merged since the last tag, so when `main` carries something you are not ready to ship, you cannot cut a patch from it. Release from a maintenance branch instead. The workflow triggers on any `*.x` branch and releases against that branch, so nothing needs editing per hotfix.

1. Branch from the last released tag and push it: `git switch -c 5.x v5.3.0 && git push -u origin 5.x`.
2. If that tag predates release-please, copy `release-please-config.json` and `.release-please-manifest.json` onto the branch and set the manifest to that tag's version.
3. Cherry-pick the `fix:` commit onto `5.x` and push.
4. release-please opens a Release PR against `5.x`. Merge it to cut the patch.
5. Forward-port the fix to `main`.

The pending Release PR on `main` is untouched throughout.

### Major versions

Go resolves a v2+ module only if `go.mod` carries the matching `/vN` suffix, and release-please does not rewrite it. The release is tagged at the Release PR's merge commit, so the migration must already be in that commit's history:

1. Land the `feat!` change. release-please opens a Release PR for the next major.
2. Migrate `go.mod` and every self-import to `/vN` in a separate PR and merge it to `main`. The Release workflow fails on those pushes while the two disagree, which is expected and blocks nothing else.
3. Click **Update branch** on the major Release PR, then merge it. The tag is created against a tree that already has the migrated `go.mod`.

Never merge a major Release PR before step 2. If that happens the tag is blocked, and the fix is manual: land the migration, remove the `autorelease: pending` label from the merged Release PR, then tag by hand at the migrated commit.
