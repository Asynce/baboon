# Contribution guidelines

Thank you for your interest in contributing to baboon!. We strive to review and merge all the PRs in a timely manner, but please note the following:

To ensure that your efforts are not wasted, please discuss any proposed changes on the issues section. While you're welcome to submit PRs directly, they may be rejected if they do not align with the required standards.

Baboon is a HTML library for Go web frontend projects or Go Webassembly projects. So, all changes must adhere to the responsiveness & web standards.Submitted changes must be compatible with the existing code. They should not break existing structure or alter any runtime behaviourr in ways that could cause issues for library users.

## Development

While creating directories one must follow snake case naming structure instead of cammel case.
for e.g. For "Internal Models" the name will be `internal_models`.

Before opening up a pull request, there are a few things you need to do make tthe code review process easier for other developers. Bad qualiryt pull requests are a bad usage of reviewer's time.
* No warnings/error.
* Spelling should be correct and variable names should make sense.
* Consistent naming convention and coding style.
* Test are written & they pass.
* If you commented out code then, it should be removed.
* If yor code needs to be documented using comments, please do it.

When you open a pull request, please fill out the pull request template with as much details as necessary to make it clear what the purpose is, what changed.

## Git strategy
[GitFlow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow) is the branching strategy we use.

The project will have a stable branch for releases. The stable branch & the development branch are protected. Direct merging/force merging is not allowed.

Before working on issues/feature please refer to the issue created. create branching according to the issue identifier.

Checkout a new feature branch off the develop branch under this project for issue-1 (`git checkout -b feature/ISSUE-1`)

If working on bugs then checkout a new bugfix branch off develop under this project respective of the isue (`git checkout -b bugfix/ISSUE-2`)

### Git commit strategy


## Git Commit Strategy

### Summary

Git commit strategy provides you a standard way of handling the commit messages for your Git repository. There are some strict set of rules defined in order to maintain the code standard of the repository. The convention defined with this strategy uses [Semantic Versioning](https://semver.org/). This strategy also takes inspiration from [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

### The Commit

The conventional commit message should be as follows:

---

```bash
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

---

The commit message should use below structure, to maintain the standard of repository in informative way.

- **fix[<module_name>]:** A commit of this type can be used when there’s a patch fix or bug fix in program. A `fix` tag is used to define a generic fix. `<module_name>` is the name of module which is affected by this commit.

Example:

```bash
git commit -m "fix[style]: aspect-ratio style fixed"
```

- **feat[<module_name>]:** A commit of this type can be used when there’s a new feature that needed to be added to repository. `feat` commit message only handles new feature implementation.`<module_name>` is the name of module which is affected by this commit.

Example:

```bash
git commit -m "feat[component]: new button component added"
```

- **!** : A commit with `!` is considered as important. In case of new MAJOR feature/bugfix of any Breaking changes which holds importance This operator can be used. Operator can be used with both `feat[<module_name>]` & `fix[<module_name>]`

Example:

```bash
git commit -m "fix[helper]!: helper for styling updated to dynamic handling"
```

- **ci**: A commit with this tag is applied when there’s changes to CI/CD files for e.g. push.yml.

Example: 

```bash
git commit -m "ci: push.yml updated for release task"
```

- **doc[<module_name>]**: A commit with this tag is applied when there are changes to code documentation.

Example

```bash
git commit -m "doc[style]: documentation over styling added"
```

- **refactor[<module_name>]**: A commit with this tag is applied when the refactoring of code is done to the project.

Example:

```bash
git commit -m "refactor[component]: scroll component reusable component added"
```

- **test**: A commit with this tag applied when there are test cases added or fixed existing test cases.

Example: 

```bash
git commit -m "test[controller]: controller tests added"
```

- **revert**: A commit with this tag is applied when there’s commit which needed to be reverted. For e.g. rolling back a feature or fix.

Example:

```bash
git commit -m "revert: rolling back feature1 due to unstability"
```

### Specification:

The documentation specification is defined using the reference of [RFC 2119](https://www.ietf.org/rfc/rfc2119.txt).

1. A commit message MUST begin with commit tag defined above. Followed by optional mark/module name `!` & `[<module_name>]`.
2. The correct tag MUST be used for the commit. A `feat` tag is only applied to feature commit. A `fix` commit MUST NOT be applied to a `feat` commit
3. A scope or module name MAY be empty depending on commit. A scope/module MUST contained within `[ ]`.
4. commit message MUST be in lowercase. Case rule is not applicable for file names, class names, variable names etc.
5. Any breaking change MUST contain `!` & the message following MUST indicate the change done. If `!` is used then the tag MUST contain scope/module affected by this change.
6. `[<module_name>]` will be the name of module affected by commit. A module MAY contain multiple files affected by this commit. 
7. A `final` tag MUST be used when all the checks of PR is passed & the pull request is ready to merge to develop. This tag finalises the commit to PR.