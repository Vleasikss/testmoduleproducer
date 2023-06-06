To publish new module changes:
1. `go mod init $MODULE_NAME` - initalize new module
2. Use Semantic Versioning (MAJOR.MINOR.PATCH): `git tag v1.0.0`
3. commit & push
4. `git push --tags`
5. 

```
In case if you want to change MAJOR version, make sure you passed version at the end of module.

If you are using a version that is v2 or higher you must include the version number into the import path of your package, for example github.com/awesomerepo/pkg/v2. You will need to use the version in the import path for both go.mod imports (require github.com/awesomerepo/pkg/v2 v2.0.0), when importing the package through command line (go get github.com/awesomerepo/pkg/v2@v2.0.0) and when importing your package in your source files (import "github.com/awesomerepo/pkg/v2 v2.0.0")

If the module version is v0 or v1 do not include it in the import path of your package
```

See also:
- https://semver.org/ - semantic versioning