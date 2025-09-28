# xCloud CLI - Context for Gemini AI

## 🌩️ Project Overview

**xCloud CLI** é uma ferramenta de linha de comando escrita em Go para gerenciar recursos na plataforma xCloud. É um CLI moderno, cross-platform, focado em performance e segurança.

## 🏗️ Architecture

### Core Structure
```
cmd/xcloud/           # Main CLI application
├── main.go          # Entry point with Cobra setup
└── main_test.go     # Core tests
go.mod               # Dependencies management
```

### Tech Stack
- **Language:** Go 1.21+
- **CLI Framework:** Cobra + Viper
- **Build Targets:** Windows & Linux (amd64, arm64)
- **Distribution:** Single optimized binaries

## 📋 Code Standards & Best Practices

### Go Guidelines
- Follow idiomatic Go patterns
- Use `gofmt` for consistent formatting
- Implement proper error handling with context
- Prefer interfaces for testability
- Keep packages focused and cohesive

### CLI UX Standards  
- Consistent command structure: `xcloud <command> [subcommand] [flags]`
- Clear, actionable error messages
- Helpful usage examples in help text
- Progress indicators for long operations
- Graceful interrupt handling (Ctrl+C)

### Cross-Platform Requirements
- Use `filepath.Join()` for path operations
- Handle Windows/Unix path differences appropriately
- Test file permissions compatibility
- Account for different shells (PowerShell, Bash, Zsh)
- Ensure consistent exit codes across platforms

### Security Priorities
- Validate all user inputs rigorously
- Sanitize file paths and prevent directory traversal
- Use secure defaults for API communications
- Implement proper credential handling
- Avoid exposing sensitive information in logs/errors

### Performance Optimization
- Minimize startup time (target: <50ms)
- Lazy load heavy dependencies
- Use connection pooling for API calls
- Implement efficient caching strategies
- Keep binary size optimized (<15MB)

## 🎯 Commands & Features

### Current Commands
- `xcloud version` - Show CLI version and build info
- `xcloud deploy` - Deploy applications to xCloud
- `xcloud status` - Check resource status
- `xcloud logs` - View application logs

### Planned Enhancements
- Auto-completion for all shells
- Configuration file management
- Interactive mode for complex operations
- Plugin system for extensibility
- Real-time monitoring integration

## 🧪 Testing Strategy

### Unit Tests
- Test all command logic thoroughly
- Mock external dependencies
- Cover error scenarios comprehensively
- Benchmark performance-critical paths

### Integration Tests
- Test actual CLI commands end-to-end
- Verify cross-platform compatibility
- Test with real (staging) xCloud environment
- Validate build artifacts functionality

### Performance Tests
- Startup time benchmarks
- Memory usage profiling
- Concurrent operation handling
- Large dataset processing tests

## 🔧 Build & Release Process

### Local Development
```bash
go run ./cmd/xcloud version        # Test locally
go test ./...                      # Run all tests  
go build -o xcloud ./cmd/xcloud   # Build binary
```

### Cross-Platform Builds
```bash
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o xcloud-linux ./cmd/xcloud
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o xcloud.exe ./cmd/xcloud
```

### Release Criteria
- All tests pass on CI/CD
- Performance benchmarks within acceptable range
- Security scans show no critical issues
- Manual testing on Windows & Linux
- Documentation updated appropriately

## 🚨 Common Pitfalls to Watch

### Go-Specific
- Avoid goroutine leaks in CLI commands
- Handle context cancellation properly
- Don't ignore errors (use proper error wrapping)
- Be careful with global state in CLI apps

### CLI-Specific  
- Don't make breaking changes to command structure
- Maintain backward compatibility in configuration
- Provide migration paths for config changes
- Test with different terminal widths/capabilities

### Cross-Platform
- Windows path length limitations
- File locking differences between OS
- Case sensitivity variations
- Different default shell behaviors

## 🎨 Code Review Focus Areas

When reviewing xCloud CLI code, prioritize:

1. **Security**: Input validation, path safety, credential handling
2. **Performance**: Startup time, memory usage, binary size
3. **Usability**: Clear help text, intuitive commands, good error messages
4. **Maintainability**: Clean abstractions, testable code, proper documentation
5. **Compatibility**: Windows/Linux differences, shell variations

## 📊 Success Metrics

### Performance Targets
- Startup time: <50ms
- Binary size: <15MB  
- Memory usage: <50MB for typical operations
- Command execution: <2s for most operations

### Quality Targets
- Test coverage: >80%
- Zero critical security issues
- Zero blocking cross-platform issues
- User satisfaction: >90% (from surveys)

## 🔮 Future Vision

The xCloud CLI should become the go-to tool for xCloud platform management, known for:
- Lightning-fast performance
- Intuitive, developer-friendly interface  
- Rock-solid reliability across platforms
- Extensible architecture for future needs
- Best-in-class security practices

---

**Remember**: This CLI represents the xCloud brand and developer experience. Every interaction should be fast, secure, and delightful! 🚀