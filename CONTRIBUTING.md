# Contributing to Ourmanga

The [Open Source Guides](https://opensource.guide/) website has a collection of resources for individuals, communities, and companies who want to learn how to run and contribute to an open source project. Contributors and people new to open source alike will find the following guides especially useful:

- [How to Contribute to Open Source](https://opensource.guide/how-to-contribute/)
- [Building Welcoming Communities](https://opensource.guide/building-community/)

## Our Development Process

Ourmanga uses [GitHub](https://github.com/djamaile/ourmanga.community) as its source of truth. The core team will be working directly there. All changes will be public from the beginning.

When a change made on GitHub is approved, it will be checked by our continuous integration system, Github Actions.

### Branch Organization

Ourmanga has one primary branch `main` and we use feature branches with deploy previews to deliver new features with pull requests.

## Bugs

We use [GitHub Issues](https://github.com/djamaile/ourmanga.community/issues) for our public bugs. If you would like to report a problem, take a look around and see if someone already opened an issue about it. If you are certain this is a new, unreported bug, you can submit a [bug report](#reporting-new-issues).

You can also file issues as [feature requests or enhancements](https://github.com/djamaile/ourmanga.community/labels/feature%20request). If you see anything you'd like to be implemented, create an issue with [feature template](https://raw.githubusercontent.com/djamaile/ourmanga.community/main/.github/ISSUE_TEMPLATE/feature_request.md)

## Reporting New Issues

When [opening a new issue](https://github.com/djamaile/ourmanga.community/issues/new/choose), always make sure to fill out the issue template. **This step is very important!** Not doing so may result in your issue not managed in a timely fashion. Don't take this personally if this happens, and feel free to open a new issue once you've gathered all the information required by the template.

- **One issue, one bug:** Please report a single bug per issue.
- **Provide reproduction steps:** List all the steps necessary to reproduce the issue. The person reading your bug report should be able to follow these steps to reproduce your issue with minimal effort.

## Pull Requests

### Your First Pull Request

So you have decided to contribute code back to upstream by opening a pull request. You've invested a good chunk of time, and we appreciate it. We will do our best to work with you and get the PR looked at.

Working on your first Pull Request? You can learn how from this free video series:

[**How to Contribute to an Open Source Project on GitHub**](https://egghead.io/courses/how-to-contribute-to-an-open-source-project-on-github)

### Proposing a Change

If you would like to request a new feature or enhancement but are not yet thinking about opening a pull request, you can also file an issue with the [feature template](https://github.com/djamaile/ourmanga.community/issues/new?template=feature.md).

If you're only fixing a bug, it's fine to submit a pull request right away but we still recommend filing an issue detailing what you're fixing. This is helpful in case we don't accept that specific fix but want to keep track of the issue.

### Sending a Pull Request

Small pull requests are much easier to review and more likely to get merged. Make sure the PR does only one thing, otherwise please split it. It is recommended to follow this [commit message style](#semantic-commit-messages).

Please make sure the following is done when submitting a pull request:

1. Fork [the repository](https://github.com/djamaile/ourmanga.community) and create your branch from `main`.
2. Add the copyright notice to the top of any code new files you've added.
3. Make sure your Jest tests pass (`yarn test`).

All pull requests should be opened against the `main` branch.

#### Breaking Changes

When adding a new breaking change, follow this template in your pull request:

```md
### New breaking change here

- **Who does this affect**:
- **How to migrate**:
- **Why make this breaking change**:
- **Severity (number of people affected x effort)**:
```

#### Copyright Header for Source Code

Copy and paste this to the top of your new file(s):

```js
/**
 * Copyright (c) Djamale Rahamat.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */
```

## Style Guide

[Prettier](https://prettier.io) will catch most styling issues that may exist in your code. You can check the status of your code styling by simply running `yarn prettier`.

However, there are still some styles that Prettier cannot pick up.

## Semantic Commit Messages

See how a minor change to your commit message style can make you a better programmer.

Format: `<type>(<scope>): <subject>`

`<scope>` is optional

## Example

```
feat: allow overriding of webpack config
^--^  ^------------^
|     |
|     +-> Summary in present tense.
|
+-------> Type: chore, docs, feat, fix, refactor, style, or test.
```

The various types of commits:

- `feat`: (new feature for the user, not a new feature for build script)
- `fix`: (bug fix for the user, not a fix to a build script)
- `docs`: (changes to the documentation)
- `style`: (formatting, missing semi colons, etc; no production code change)
- `refactor`: (refactoring production code, eg. renaming a variable)
- `test`: (adding missing tests, refactoring tests; no production code change)
- `chore`: (updating grunt tasks etc; no production code change)

Use lower case not title case!

### Code Conventions

#### General

- **Most important: Look around.** Match the style you see used in the rest of the project. This includes formatting, naming files, naming things in code, naming things in documentation.
- "Attractive"

### Documentation

- Do not wrap lines at 80 characters - configure your editor to soft-wrap when editing documentation.

## License

By contributing to Ourmanga, you agree that your contributions will be licensed under its MIT license.