package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "xcloud",
	Short: "🌩️  xCloud CLI - Interface de linha de comando para a plataforma xCloud",
	Long: `xCloud CLI é uma ferramenta poderosa para gerenciar recursos na plataforma xCloud.
	
Recursos principais:
• Deploy de aplicações serverless
• Gerenciamento de ambientes
• Monitoramento e logs
• Configuração de serviços`,
	Version: fmt.Sprintf("%s (commit: %s, built at: %s)", version, commit, date),
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Mostra a versão do xCloud CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("xCloud CLI %s\n", rootCmd.Version)
		fmt.Printf("Commit: %s\n", commit)
		fmt.Printf("Build Date: %s\n", date)
	},
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy de aplicações na xCloud",
	Long:  "Realiza o deploy de aplicações serverless na plataforma xCloud",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 Iniciando deploy na xCloud...")
		// TODO: Implementar lógica de deploy
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Verifica o status dos recursos",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📊 Status dos recursos xCloud:")
		// TODO: Implementar verificação de status
	},
}

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Visualiza logs das aplicações",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📜 Logs das aplicações:")
		// TODO: Implementar visualização de logs
	},
}

var initCmd = &cobra.Command{
	Use:   "init [nome-do-projeto]",
	Short: "Inicializa um novo projeto xCloud",
	Long:  "Cria a estrutura básica de um novo projeto xCloud com templates e configurações",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := "my-xcloud-project"
		if len(args) > 0 {
			projectName = args[0]
		}

		template, _ := cmd.Flags().GetString("template")

		fmt.Printf("🚀 Inicializando projeto xCloud: %s\n", projectName)
		if verbose {
			fmt.Printf("📦 Template selecionado: %s\n", template)
		}

		// TODO: Implementar lógica de inicialização completa
		fmt.Println("✅ Projeto inicializado com sucesso!")
	},
}

var createCmd = &cobra.Command{
	Use:       "create [tipo] [nome]",
	Short:     "Cria recursos na xCloud",
	Long:      "Cria novos recursos como funções serverless, APIs, serviços e componentes na plataforma xCloud",
	Args:      cobra.MinimumNArgs(1),
	ValidArgs: []string{"function", "api", "service", "component"},
	Run: func(cmd *cobra.Command, args []string) {
		resourceType := args[0]
		resourceName := ""
		if len(args) > 1 {
			resourceName = args[1]
		} else {
			resourceName = "my-" + resourceType
		}

		template, _ := cmd.Flags().GetString("template")
		runtime, _ := cmd.Flags().GetString("runtime")

		fmt.Printf("✨ Criando %s: %s\n", resourceType, resourceName)
		if verbose {
			fmt.Printf("📦 Template: %s\n", template)
			fmt.Printf("⚙️  Runtime: %s\n", runtime)
		}

		// TODO: Implementar lógica de criação de recursos
		fmt.Println("✅ Recurso criado com sucesso!")
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(logsCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(createCmd)

	// Flags globais
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "arquivo de configuração (default: $HOME/.xcloud.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "saída detalhada")

	// Deploy flags
	deployCmd.Flags().StringP("env", "e", "production", "ambiente de destino")
	deployCmd.Flags().BoolP("watch", "w", false, "observar mudanças e re-deploy automaticamente")

	// Logs flags
	logsCmd.Flags().IntP("tail", "t", 100, "número de linhas para mostrar")
	logsCmd.Flags().BoolP("follow", "f", false, "seguir logs em tempo real")

	// Init flags
	initCmd.Flags().StringP("template", "t", "basic", "template do projeto (basic, serverless, api)")

	// Create flags
	createCmd.Flags().StringP("template", "t", "default", "template do recurso")
	createCmd.Flags().StringP("runtime", "r", "nodejs", "runtime para funções (nodejs, python, go)")
}

var cfgFile string
var verbose bool

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".xcloud")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil && verbose {
		fmt.Fprintln(os.Stderr, "Usando arquivo de configuração:", viper.ConfigFileUsed())
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
		os.Exit(1)
	}
}
