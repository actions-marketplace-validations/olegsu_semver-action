# Semver Bump

Github action that bumps semantic version

# Example
```yaml
    steps:
    - uses: olegsu/semver-action@v1
      id: version
      with:
        version: "1.1.1" # required
        version_file: "./VERSION" # load version from this file instead of using explicit version
        bump: patch # [patch,minor,major] (patch is default)
    - name: Get the new version
      run: echo "${{ steps.version.outputs.version }}" # 1.1.2
```