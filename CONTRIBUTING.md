
> [!WARNING]  
> This project currently does not accept any contribution since it will be graded for school. \

## Initial setup
Currently for linux-only

clone this repo
```bash
clone --recurse-submodules https://github.com/chrxer/safe-chrx-proto.git
cd safe-chrx-proto
```

(optional) open in VSC
```bash
code safe-chrx-proto/chrxer.code-workspace
```

Install deps
```bash
sudo scripts/deps.sh
```

## Patching guidelines
1. Move as much as possible to @chromium`chrxer/*`
2. Minimize modifying files
3. Don't delete any files within @chromium

## Build
Apply patches
```
scripts/patch.py
```
build
```
scipts/build.py
```

## Develop patches
Run (after your edits)
```
scripts/diff.py
```

The diff for modified files can currently be found at [os_crypt.patch](os_crypt.patch) and added files (tree-preserving) at [patch](patch/)

