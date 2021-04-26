package logutil

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sort"
)

var DefaultZapLoggerConfig = zap.Config{
	Level: zap.NewAtomicLevelAt(zap.InfoLevel),

	Development: true,
	Sampling: &zap.SamplingConfig{
		Initial: 100,
		Thereafter: 100,
	},

	Encoding: "json",

	EncoderConfig: zapcore.EncoderConfig{
		TimeKey: "ts",
		LevelKey: "level",
		NameKey: "logger",
		CallerKey: "caller",
		MessageKey: "msg",
		StacktraceKey: "stacktrace",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	},

	OutputPaths: []string{"stderr"},
	ErrorOutputPaths: []string{"stderr"},
}

func MergeOutputPaths(cfg zap.Config) zap.Config {
	outputs := make(map[string]struct{})
	for _, v := range cfg.OutputPaths {
		outputs[v] = struct{}{}
	}

	outputSlice := make([]string, 0)
	if _, ok := outputs["/dev/null"]; ok {
		outputSlice = []string{"/dev/null"}
	}else {
		for k := range outputs {
			outputSlice = append(outputSlice, k)
		}
	}
	cfg.OutputPaths = outputSlice
	sort.Strings(cfg.OutputPaths)

	errOutputs := make(map[string]struct{})
	for _, v := range cfg.ErrorOutputPaths {
		errOutputs[v] = struct{}{}
	}
	errOutputSlice := make([]string, 0)
	if _, ok := errOutputs["/dev/null"]; ok {
		errOutputs["/dev/null"] = struct{}{}
	}else {
		for k := range errOutputs {
			errOutputSlice = append(errOutputSlice, k)
		}
	}
	cfg.ErrorOutputPaths = errOutputSlice
	sort.Strings(cfg.ErrorOutputPaths)

	return cfg
}