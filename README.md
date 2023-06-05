To publish new module changes:
1. `go mod init $MODULE_NAME` - initalize new module
2. Use Semantic Versioning (MAJOR.MINOR.PATCH): `git tag v1.0.0`
3. commit & push
4. `git push --tags`
5. 

```
In case if you want to change MAJOR version, make sure you passed version at the end of module.

For example: module github.com/Vleasikss/testmoduleproducer/v5
Every major version (starting from 2nd) must be a separate module.
```

See also:
- https://semver.org/ - semantic versioning