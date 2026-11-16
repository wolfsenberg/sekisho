# Contributing to Sekisho

Sekisho is at the foundation stage. Read [PRODUCT.md](PRODUCT.md) before expanding scope. Small improvements to a reliable Codex-first workflow are more useful than broad provider coverage at this stage.

## Local workflow

1. Install Go 1.26+, Node.js 24, pnpm 9.12.2, and jq.
2. Run `make setup`, then `make dev`.
3. Write a failing test for a behavior change; implement the change and rerun the test.
4. Run `make check` before submitting your work.
5. Update relevant documentation and every supported UI dictionary when behavior changes.

See the [Development section of the README](README.md#development) for commands.

## Change boundaries

- Keep changes focused on one problem and describe the concrete before/after behavior.
- Use synthetic fixtures. Do not include auth files, subscription identifiers, real request payloads, or personal information in tests, screenshots, logs, or issue reports.
- Add SQLite migrations rather than rewriting migrations that have shipped.
- Keep provider SDK types behind their integration boundary.
- Avoid adding dependencies for small tasks already handled by the standard library or installed packages.
- Document validation actually performed, plus any untested client or platform.

## Interface changes

Follow [DESIGN.md](DESIGN.md). Check loading, empty, error, and success states where applicable. Verify keyboard navigation, narrow screens, light/dark themes, and English copy. Use the existing semantic tokens and component conventions.

## Pull requests

Explain the problem, resulting behavior, relevant tradeoffs, and verification. Attach sanitized UI screenshots to the PR only when they help reviewers; do not commit screenshots solely as PR evidence. Commit screenshots only when current documentation uses them. Include migration notes when persisted data changes. Do not include unrelated cleanup or local planning documents.

No contribution agreement or public support process is configured yet. Contributions must be compatible with the project's AGPL-3.0-only license and preserve third-party attribution.

## Documentation changes

Update `README.md` and the other documents listed in [the documentation index](docs/README.md) in the same change as the behavior they describe.

## Publishing a version

Choose an unused release tag. Before creating it, use `git-cliff 2.14.2` to run `make changelog-check` and generate the versioned changelog below, replacing `vX.Y.Z` with the chosen tag:

```sh
git-cliff --config cliff.toml --offline --no-exec --use-branch-tags --tag vX.Y.Z --output CHANGELOG.md
```

Review the new version section and comparison link, then commit the updated `CHANGELOG.md` and merge it into `main` through the normal PR flow. The generator labels the planned release without creating a Git tag.

On clean, synchronized `main` containing that changelog, run `make publish-check TAG=vX.Y.Z` and `make publish TAG=vX.Y.Z`. Wait for the Release workflow to succeed and verify the published assets before announcing completion. The workflow generates GitHub Release notes but does not update the repository changelog.
