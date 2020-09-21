package main

import (
    "errors"
    "fmt"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "io/ioutil"
    "os"
)

func main() {
    levels()
    setGlobalLevel()
    structured()
    logError()
    sublogger()
    fileOutput()
    prettyConsole()
}

func levels() {
    // Panic of fatal messages stop the execution flow
    // log.Panic().Msg("This is a panic message")
    // log.Fatal().Msg("This is a fatal message")
    log.Error().Msg("This is an error message")
    log.Warn().Msg("This is a warning message")
    log.Info().Msg("This is an information message")
    log.Debug().Msg("This is a debug message")
    log.Trace().Msg("This is a trace message")
}

func setGlobalLevel() {
    zerolog.SetGlobalLevel(zerolog.DebugLevel)

    log.Debug().Msg("Debug message is displayed")
    log.Info().Msg("Info Message is displayed")

    zerolog.SetGlobalLevel(zerolog.InfoLevel)
    log.Debug().Msg("Degug message is no longer displayed")
    log.Info().Msg("Info message is displayed")
}

func structured() {
    log.Info().Str("mystr","this is a string").Msg("")
    log.Info().Int("myint",1234).Msg("")
    log.Info().Int("myint",1234).Str("str","some string").Msg("And a regular message")
}

func logError() {
    err := errors.New("there is an error")

    log.Error().Err(err).Msg("this is the way to log errors")
}

func sublogger() {

    mainLogger := zerolog.New(os.Stderr).With().Logger()
    mainLogger.Info().Msg("This is the output from the main logger")

    subLogger := mainLogger.With().Str("component","componentA").Logger()
    subLogger.Info().Msg("This is the the extended output from the sublogger")
}

func fileOutput() {
    // create a temp file
    tempFile, err := ioutil.TempFile(os.TempDir(),"deleteme")
    if err != nil {
        // Can we log an error before we have our logger? :)
        log.Error().Err(err).Msg("there was an error creating a temporary file four our log")
    }

    fileLogger := zerolog.New(tempFile).With().Logger()
    fileLogger.Info().Msg("This is an entry from my log")

    fmt.Printf("The log file is allocated at %s\n", tempFile.Name())
}

func prettyConsole() {
    log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

    log.Error().Msg("This is an error message")
    log.Warn().Msg("This is a warning message")
    log.Info().Msg("This is an information message")
    log.Debug().Msg("This is a debug message")
}

