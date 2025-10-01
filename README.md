<div align="center">

# xCloud CLI

**⚡ Interface de linha de comando para implantar e gerenciar aplicações**

[![Status](https://img.shields.io/badge/Status-Active-success)](https://github.com/PageCloudv1/xcloud-cli)
[![License](https://img.shields.io/badge/License-MIT-blue)](LICENSE)
[![CI/CD](https://github.com/PageCloudv1/xcloud-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/PageCloudv1/xcloud-cli/actions/workflows/ci.yml)

</div>

---

## 🚀 Sobre o Projeto

O **xCloud CLI** é a interface de linha de comando para implantar e gerenciar aplicações na plataforma xCloud.

---

## 🎯 Ecossistema xCloud Platform

A xCloud Platform é composta por um conjunto de repositórios projetados para trabalhar em conjunto, fornecendo uma experiência de desenvolvimento completa e integrada.

| Repositório | Descrição |
|---|---|
| **[xcloud-platform](https://github.com/PageCloudv1/xcloud-platform)** | Core da plataforma, orquestrando build, deploy e gerenciamento. |
| **[xcloud-cli](https://github.com/PageCloudv1/xcloud-cli)** | Interface de linha de comando em Go para interagir com a plataforma. |
| **[xcloud-dashboard](https://github.com/PageCloudv1/xcloud-dashboard)** | Interface web em React para gerenciamento de projetos e analytics. |
| **[xcloud-runtime](https://github.com/PageCloudv1/xcloud-runtime)** | Runtime serverless para funções em Python, Node.js e Go. |
| **[xcloud-docs](https://github.com/PageCloudv1/xcloud-docs)** | Documentação completa da plataforma. |
| **[xcloud-templates](https://github.com/PageCloudv1/xcloud-templates)** | Templates de projetos prontos para uso. |
| **[xcloud-components](https://github.com/PageCloudv1/xcloud-components)** | Marketplace de componentes de UI e integrações. |
| **[xcloud-examples](https://github.com/PageCloudv1/xcloud-examples)** | Projetos de exemplo e demonstrações. |
| **[xcloud-bot](https://github.com/PageCloudv1/xcloud-bot)** | Assistente de IA para operações DevOps. |
| **[xcloud-containers](https://github.com/PageCloudv1/xcloud-containers)** | Configurações de contêineres Podman para o ambiente de desenvolvimento. |

---

## 🤖 Automação e IA

O xCloud CLI possui um sistema de automação completo com suporte a IA para facilitar o desenvolvimento:

### 🔄 Workflows Disponíveis

#### Auto Refactor Issues
- **Trigger**: Quando uma issue é criada ou reaberta
- **Ações**: 
  - Auto-atribuição ao criador
  - Aplicação automática de labels
  - Comentário de confirmação
  - Notificação do @Copilot

#### Gemini Review
- **Trigger**: Comentários com comandos `/gemini`
- **Comandos Suportados**:
  - `/gemini review` - Análise completa do repositório
  - `/gemini analyze` - Análise detalhada de código
  - `/gemini suggest` - Sugestões de melhorias
  - `/gemini validate` - Validação de configurações
- **Ações**: 
  - Análise de código e testes
  - Validação de build
  - Relatório detalhado com métricas
  - Sugestões de otimização

#### CI/CD Pipeline
- **Lint**: Verificação de formatação e estilo
- **Test**: Testes automatizados com coverage
- **Build**: Compilação e validação
- **Quality Gate**: Validação geral de qualidade

#### Release Automation
- **Trigger**: Tags de versão (v*.*.*)
- **Ações**: Build multi-plataforma e publicação de releases

### 📚 Documentação de Automação

- **[AUTOMATION_VALIDATION.md](.github/AUTOMATION_VALIDATION.md)** - Guia completo de validação
- **[app-permissions.md](cmd/app-permissions.md)** - Configuração do GitHub App
- **[GEMINI.md](GEMINI.md)** - Contexto para integração com IA

### 🚀 Como Usar

1. **Para análise automática**: Crie uma issue e ela será automaticamente processada
2. **Para review com IA**: Comente `/gemini review` em qualquer issue ou PR
3. **Para CI/CD**: Faça push para main/develop e o pipeline executará automaticamente
4. **Para releases**: Crie uma tag de versão (ex: `v1.0.0`)

---

## 🤝 Como Contribuir

Contribuições são bem-vindas! Para começar, leia nosso **[Guia de Contribuição](CONTRIBUTING.md)** e siga nosso **[Código de Conduta](CODE_OF_CONDUCT.md)**.

## 📝 Licença

Este projeto é licenciado sob a Licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
