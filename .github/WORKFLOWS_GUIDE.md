# 🤖 xCloud Workflows - Quick Reference

## 📋 Workflows Overview

### 1. Auto Refactor Issues (`auto-refactor-issues.yml`)

**Trigger**: `issues` → `opened`, `reopened`

**What it does**:
- ✅ Auto-assigns issue to creator
- ✅ Applies labels: `🤖 auto-refactored`, `👤 <username>`
- ✅ Posts confirmation comment
- ✅ Notifies @Copilot

**Example Flow**:
```
User creates issue #123
  ↓
Workflow triggers automatically
  ↓
Issue assigned to user
  ↓
Labels applied
  ↓
Bot posts comment with instructions
```

**Permissions Required**:
- `issues: write`
- `contents: read`

---

### 2. Gemini Review (`gemini-review.yml`)

**Trigger**: `issue_comment` → `created`

**Commands**:
- `/gemini review` - Complete repository analysis
- `/gemini analyze` - Code analysis
- `/gemini suggest` - Improvement suggestions
- `/gemini validate` - Configuration validation

**What it does**:
- ✅ Detects Gemini commands in comments
- ✅ Analyzes repository structure
- ✅ Runs tests, go vet, and builds
- ✅ Validates workflows
- ✅ Posts comprehensive report

**Example Usage**:
```
User comments: /gemini review
  ↓
Workflow detects command
  ↓
Runs full analysis
  ↓
Posts detailed report with:
  - Repository stats
  - Test results
  - Code quality metrics
  - Workflow validation
  - Improvement suggestions
```

**Permissions Required**:
- `issues: write`
- `pull-requests: write`
- `contents: read`
- `actions: read`

---

### 3. CI/CD Pipeline (`ci.yml`)

**Trigger**: `push`, `pull_request` → `main`, `develop`

**Jobs**:

#### Lint
- Runs `go fmt` check
- Executes `go vet`
- Runs `golangci-lint`

#### Test
- Tests on Go 1.21 and 1.22
- Runs with race detection
- Generates coverage reports
- Uploads to Codecov

#### Build
- Builds xCloud CLI binary
- Verifies binary functionality
- Uploads artifacts

#### Quality Gate
- Validates all job results
- Fails if any check fails

**Example Flow**:
```
Developer pushes code
  ↓
Lint checks run
  ↓
Tests execute on multiple Go versions
  ↓
Binary builds
  ↓
Quality gate validates
  ↓
All checks pass ✅
```

**Permissions Required**:
- `contents: read`
- `actions: read`

---

### 4. Release Automation (`release.yml`)

**Trigger**: `push` → `tags: v*.*.*` or `workflow_dispatch`

**What it does**:
- ✅ Builds for Linux (amd64, arm64)
- ✅ Builds for macOS (amd64, arm64)
- ✅ Builds for Windows (amd64, arm64)
- ✅ Creates GitHub Release
- ✅ Uploads all binaries

**Example Usage**:
```bash
# Create and push a version tag
git tag v1.0.0
git push origin v1.0.0
  ↓
Workflow triggers
  ↓
Multi-platform builds
  ↓
Release created with all binaries
```

**Permissions Required**:
- `contents: write`

---

## 🎯 Usage Examples

### Creating an Issue
1. Go to Issues tab
2. Click "New Issue"
3. Fill in details
4. Submit
5. ✅ Auto-assignment happens
6. ✅ Bot comments with instructions

### Using Gemini Commands
1. Open any issue or PR
2. Add a comment: `/gemini review`
3. Wait for analysis (2-5 minutes)
4. ✅ Receive detailed report

### Pushing Code
1. Make changes locally
2. Commit and push to main/develop
3. ✅ CI/CD runs automatically
4. Check Actions tab for results

### Creating a Release
1. Ensure all tests pass
2. Create version tag:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```
3. ✅ Release builds automatically
4. Binaries available in Releases

---

## 🔧 Configuration

### GitHub App Setup

Required for auto-assignment and bot features:

1. Create GitHub App at: Settings → Developer Settings → GitHub Apps
2. Configure permissions (see `cmd/app-permissions.md`)
3. Subscribe to events: `issues`, `issue_comment`, `pull_request`
4. Install app on repository

### Workflow Secrets

No secrets required for basic functionality. All workflows use `GITHUB_TOKEN` automatically.

---

## 📊 Monitoring

### Workflow Status

Check workflow status:
1. Go to Actions tab
2. Select workflow from left sidebar
3. View recent runs

### Job Logs

View detailed logs:
1. Click on a workflow run
2. Click on a job name
3. Expand steps to see details

### Summary Reports

Each workflow provides summary:
- View in the workflow run page
- Check "Summary" section at top

---

## 🐛 Troubleshooting

### Workflow not triggering?

**Check**:
- Is workflow file in `.github/workflows/`?
- Is YAML syntax valid?
- Are triggers configured correctly?
- Are Actions enabled in repository settings?

### Auto-assignment not working?

**Check**:
- Is GitHub App installed?
- Does app have `issues: write` permission?
- Is `issues` event subscribed?
- Check webhook deliveries in app settings

### Gemini command not responding?

**Check**:
- Is command syntax correct? (e.g., `/gemini review`)
- Is workflow file present?
- Check Actions tab for errors
- Verify `issue_comment` trigger is enabled

### Tests failing?

**Check**:
- Run tests locally: `go test ./...`
- Check go version: `go version`
- Verify dependencies: `go mod download`
- Review error logs in Actions tab

---

## 🔗 Related Documentation

- **[AUTOMATION_VALIDATION.md](AUTOMATION_VALIDATION.md)** - Complete validation guide
- **[../cmd/app-permissions.md](../cmd/app-permissions.md)** - GitHub App setup
- **[../GEMINI.md](../GEMINI.md)** - AI integration context
- **[../README.md](../README.md)** - Project overview

---

## 📈 Performance Metrics

| Workflow | Avg Duration | Timeout |
|----------|-------------|---------|
| auto-refactor-issues | ~30s | 5 min |
| gemini-review | ~3m | 10 min |
| ci | ~8m | 15 min |
| release | ~4m | N/A |

---

## ✅ Best Practices

1. **Issue Creation**: Let auto-assignment handle it
2. **Code Review**: Use `/gemini review` before merging
3. **Testing**: Ensure CI passes before merging
4. **Releases**: Use semantic versioning (v1.0.0)
5. **Monitoring**: Check Actions tab regularly

---

*Last Updated: 2024-10-01*  
*Version: 1.0*
