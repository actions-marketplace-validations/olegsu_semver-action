# Semver Bump

Github action that bumps semantic version

# Example
```yaml
    steps:
    - uses: olegsu/semver-action@main
      id: version
      with:
        version: "1.1.1" # required
        bump: patch # [patch,minor,major] (patch is default)
    - name: Get the new version
      run: echo "${{ steps.version.outputs.version }}" # 1.1.2
```