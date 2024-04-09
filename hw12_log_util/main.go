package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type level int8

const (
	Info level = iota
	Trace
	Debug
)

const (
	defaultOutput = "stdout"
	defaultFile   = "default"
)

func (e level) String() string {
	switch e {
	case Info:
		return "info"
	case Trace:
		return "trace"
	case Debug:
		return "debug"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type Config struct {
	logFile    string
	logLevel   level
	statOutput string
}

func New() *Config {
	config := &Config{
		logFile:    getStrEnv("LOG_ANALYZER_FILE", defaultFile),
		logLevel:   getLogLevel(),
		statOutput: getStrEnv("LOG_ANALYZER_OUTPUT", defaultOutput),
	}
	return config
}

func getLogLevel() level {
	value := getStrEnv("LOG_ANALYZER_LEVEL", "info")
	return toLevel(value)
}

func toLevel(value string) level {
	switch {
	case strings.EqualFold(value, Trace.String()):
		return Trace
	case strings.EqualFold(value, Debug.String()):
		return Debug
	case strings.EqualFold(value, Info.String()):
		return Info
	default:
		log.Fatal("wrong level type, utillity stops")
		return 0
	}
}

func getStrEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

var conf *Config

func init() {
	log.Println("- Start")
	var levelStr string
	conf = New()
	flag.StringVar(&conf.logFile, "file", conf.logFile, "log file name")
	flag.StringVar(&conf.statOutput, "output", conf.statOutput, "output file name")
	flag.StringVar(&levelStr, "level", conf.logLevel.String(), "log level")
	flag.Parse()
	conf.logLevel = toLevel(levelStr)
	fmt.Printf("- log file:\"%s\";\n- log level:\"%s\";\n- stat output:\"%s\".\n",
		conf.logFile,
		conf.logLevel,
		conf.statOutput)
}

func main() {
	if conf.logFile == defaultFile {
		log.Fatal("flag file is mandatory, utillity stops")
	}
	result, err := ReadLog()
	if err != nil {
		log.Fatal("Some error:", err)
	}
	WriteResult(result)
	log.Println("- Done")
}

func ReadLog() (map[string]int, error) {
	file, openErr := os.Open(conf.logFile)
	if openErr != nil {
		log.Println(openErr)
		return nil, openErr
	}
	defer file.Close()
	duplicatesCount := make(map[string]int)
	scanner := bufio.NewScanner(file) // default line capacity is 64K
	for scanner.Scan() {
		if val, err := parseLine(scanner.Text()); err == nil {
			duplicatesCount[val]++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return duplicatesCount, err
	}
	return duplicatesCount, nil
}

func parseLine(line string) (string, error) {
	re := regexp.MustCompile(`[0-2][0-9]\:[0-6][0-9]\:[0-6][0-9]\.[0-9]{3}\s` +
		regexp.QuoteMeta(conf.logLevel.String()) +
		`\s(?P<ID>[0-9]{5})\s`)
	arr := re.FindAllStringSubmatch(line, -1)
	if arr != nil {
		return arr[0][1], nil
	}
	return "", errors.New("empty string")
}

func WriteResult(input map[string]int) error {
	var result strings.Builder
	resultFormat(&result, input)
	if conf.statOutput == defaultOutput {
		fmt.Println(result.String())
		return nil
	}
	file, _ := os.Create(conf.statOutput)
	defer file.Close()
	_, err := file.Write([]byte(result.String()))
	if err != nil {
		log.Printf("failed to write: %v", err)
		return err
	}
	return nil
}

func resultFormat(result *strings.Builder, input map[string]int) {
	result.WriteString("log stat is:\n")
	for k, v := range input {
		result.WriteString(fmt.Sprintf("Event code[%s] count match[%d]\n", k, v))
	}
}
