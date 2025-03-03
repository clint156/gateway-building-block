# Conventions
This document contains a description of the conventions which should be followed when contributing to this repository. They are not final and suggestions are welcome.

## Formatting
- Follow [Effective Go](https://go.dev/doc/effective_go) conventions.
- Use snake-case everywhere in the database and JSON objects.

## Changelog
We should keep the changelog up to date as this is part of the open source platform. Each issue should be added to the top of the appropriate [verb](https://keepachangelog.com/en/1.0.0/#how) in the `[Unreleased]` section of the changelog in the corresponding issue branch in the format `{issue name} [#{issue-ID}]({issue url})` (eg. [Unreleased] Added - Set up project skeleton [#1](https://github.com/rokwire/core-building-block/issues/1))

## API Documentation
When implementing an API:
- Use [swag](https://github.com/swaggo/swag) annotations when defining a new API handler function
- Run `make swagger` to generate the `swagger.yaml`, `swagger.json`, and `docs.go` files stored in the `docs/` folder. To run this command, you will need to install [swag](https://github.com/swaggo/swag). This command will automatically generate Open API 3 documentation for all functions using the annotations. Please do not change any of the files in the `docs/` folder manually.
- Test you API via the documentation - Open http://localhost/gateway/api/doc/ui/ , choose "Local server" from the "Servers" combobox and run your API. This is an alternative to Postman. Make sure to set the correct value in the `HOST` environment variable (eg. http://localhost/gateway) before running the service to access the docs.

## Pull Requests
Pull requests should be linked to the associated issue with a [keyword](https://docs.github.com/en/issues/tracking-your-work-with-issues/creating-issues/linking-a-pull-request-to-an-issue#linking-a-pull-request-to-an-issue-using-a-keyword) in the description (eg. `Resolves #{issue number}`). This will close the issue automatically when the PR is merged. 

## Unit Tests
Whenever a new interface is created, a unit test should be created for each function it exposes. The purpose of these unit tests is primarily to ensure that the contract with consumers established by the interfaces are not unintentionally broken by future implementation changes. With this in mind, test cases should include all common usage, as well as any edge cases for which consistency is important. 

When updating or changing existing implementations, run the associated unit tests to ensure that they still pass. If they do not, the implementation changes likely changed the interface as well. If the change to the interface was intentional, update the unit tests as needed to make them pass and document the [Breaking Change](#breaking-changes). If the change was not intentional, rework your implementation changes to keep the interface consistent and ensure all tests pass.

## Releases
Whenever a new release is made, the following process should be followed.

1. Make a pull request from `develop` into `main` named `Release vX.X.X` (eg. `Release v1.1.7`)
2. Review the changes included in the update to ensure they are all production ready.
3. Update the "Unreleased" version in the [CHANGELOG](CHANGELOG.md#unreleased) to `X.X.X - YYYY-MM-dd` (eg. `[1.1.7] - 2022-06-08`) on the `develop` branch.
4. Update [SECURITY.md](SECURITY.md) to reflect the latest supported and unsupported versions on the `develop` branch.
5. Update the latest version in any docs or source code as needed on the `develop` branch. 
6. Make any changes needed to document [breaking changes](#breaking-changes) and [deprecations](#deprecations).
7. Merge the pull request using "Create a merge commit"
8. Create a new tag from the `main` branch called `vX.X.X` (eg. `v1.1.7`)
9. **RECOMMENDED** - Publish a new [GitHub Release](https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository#creating-a-release) from this tag with the title `vX.X.X` (eg. `v1.1.7`). Include the contents from the [CHANGELOG](CHANGELOG.md) for this latest version in the release notes, as well as a link to the whole [CHANGELOG](CHANGELOG.md) on the `main` branch. For libraries this is highly recommended.

## Breaking Changes
Breaking changes should be avoided when possible, but will sometimes be necessary. In the event that a breaking change does need to be made, this change should be documented clearly for developers relying on the functionality. This includes the following items:
* Create and apply a "breaking" label to the associated issue in GitHub
* Add a "BREAKING:" prefix to the associated line in the CHANGELOG
* Document upgrade instructions in the README in the `Upgrading > Migration steps > Unreleased > Breaking changes` section. These should explain the changes that were made, as well as all changes the developer will need to make to handle the breaking change. Examples should be provided where appropriate.

When a release including the breaking change is created, the following steps must be taken:
* Update the MAJOR version number to indicate that incompatible interface changes have occurred (see [Semantic Versioning](https://semver.org/))
* Update the `Upgrading > Migration steps > Unreleased` section in the README to the latest version (eg. `Upgrading > Migration steps > v1.1.0`)
* Add a "BREAKING" warning to the release notes
* Include a copy of the upgrade instructions from the README in the release notes

## Deprecations
In some cases when [Breaking Changes](#breaking-changes) need to be made, the existing functionality must be maintained to provide backwards compatibility. To do so, the new component (function, type, field, package...) should be created and the old component should be maintained and flagged as deprecated. This will give time for developers relying on the component to make the necessary updates before it becomes unavailable. In these cases, the following process should be followed:
* Add a "DEPRECATED:" prefix to the associated line in the CHANGELOG
* Add a "Deprecated:" comment to the component and provide information about the deprecation and replacement. See the [Godoc](https://go.dev/blog/godoc) documentation for more information.
* Document upgrade instructions in the README in the `Upgrading > Migration steps > Unreleased > Deprecations` section. These should explain the changes that were made, as well as all changes the developer will need to make to replace the deprecated component. Examples should be provided where appropriate. If known, include a timeline for when the deprecated components will be removed.

When a release including the deprecation is created, the following steps must be taken:
* Update the `Upgrading > Migration steps > Unreleased` section in the README to the latest version (eg. `Upgrading > Migration steps > v1.1.0`)
* Include a copy of the upgrade instructions from the README in the release notes

When the deprecated components are finally removed, follow the process to document this as a [Breaking Change](#breaking-changes). 