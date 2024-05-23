/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	service "github.com/stas727/gbot/cmd/services"
	storage "github.com/stas727/gbot/cmd/storages"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"log"
	"os"
	"time"
)

const MessageHi = "Цей бот на базі штучного інтелекту володіє великою базою даних і здатен аналізувати великі обсяги інформації у реальному часі. Він може надати користувачам різноманітну інформацію і рекомендації залежно від їх потреб і запитань.\n\nКорисність: AI Vision може бути корисним для студентів, дослідників, бізнесменів, маркетологів та будь-яких інших людей, які потребують швидкого і точного аналізу даних. Він може допомогти у прийнятті рішень, проведенні досліджень, аналітиці даних, а також в автоматичному оновленні інформації."
const MessageVersion = "Hello, I am AI %s!"

var (
	TeleToken = os.Getenv("TELE_TOKEN")
	AiToken   = os.Getenv("AI_TOKEN")
	// MetricsHost exporter host:port
	MetricsHost = os.Getenv("METRICS_HOST")
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

var (
	storages = storage.NewStorages()
	services = service.NewServices(storages)
)

// Initialize OpenTelemetry
func initMetrics(ctx context.Context) {

	// Create a new OTLP Metric gRPC exporter with the specified endpoint and options
	exporter, _ := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithEndpoint(MetricsHost),
		otlpmetricgrpc.WithInsecure(),
	)

	// Define the resource with attributes that are common to all metrics.
	// labels/tags/resources that are common to all metrics.
	newResource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(fmt.Sprintf("gbot_%s", appVersion)),
	)

	// Create a new MeterProvider with the specified resource and reader
	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(newResource),
		sdkmetric.WithReader(
			// collects and exports metric data every 10 seconds.
			sdkmetric.NewPeriodicReader(exporter, sdkmetric.WithInterval(10*time.Second)),
		),
	)

	// Set the global MeterProvider to the newly created MeterProvider
	otel.SetMeterProvider(mp)
}

// gbotCmd represents the gbot command
var gbotCmd = &cobra.Command{
	Use:     "gbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Print(fmt.Sprintf("Gbot %s started", appVersion))

		ctx := context.Background()

		//new AI Client
		client := services.AIService.NewClient(ctx, AiToken)

		//new Bot
		newBot, err := services.BotService.NewBot(ctx, TeleToken, 10*time.Second, "")
		if err != nil {
			logger.Print(fmt.Sprintf("Telebot error. %s", err.Error()))

			return
		}

		//conversation
		e := services.BotService.Handle(ctx, func(message string) string {
			switch message {
			case "/start", "hello", "Hello", "hi", "hey", "Hi", "Hey", "привіт", "Привіт":
				return MessageHi
			case "version":
				return fmt.Sprintf(MessageVersion, appVersion)
			default:
				response, err := services.AIService.Response(ctx, message, client)

				if err != nil {
					logger.Print("Error AI response. Err : " + err.Error())
					return ""
				}
				return *response
			}
		}, newBot)

		if e != nil {
			return
		}

		newBot.Client.Start()
	},
}

func init() {
	ctx := context.Background()
	initMetrics(ctx)
	rootCmd.AddCommand(gbotCmd)
}
