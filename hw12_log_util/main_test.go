package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ = func() bool {
	testing.Init()
	return true
}()

func TestToLevel(t *testing.T) {
	testCases := []struct {
		want   level
		desc   string
		input1 string
	}{
		{
			desc:   "check info",
			input1: "INFO",
			want:   Info,
		},
		{
			desc:   "check trace",
			input1: "Trace",
			want:   Trace,
		},
		{
			desc:   "check debug",
			input1: "debug",
			want:   Debug,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			logLevel := toLevel(tC.input1)
			assert.Equal(t, tC.want, logLevel)
		})
	}
}

func TestGetLogLevel(t *testing.T) {
	testCases := []struct {
		want level
		desc string
	}{
		{
			desc: "check default level",
			want: Info,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			logLevel := getLogLevel()
			assert.Equal(t, tC.want, logLevel)
		})
	}
}

func TestGetStrEnv(t *testing.T) {
	testCases := []struct {
		want  string
		desc  string
		input string
		env   string
	}{
		{
			desc:  "check default file",
			want:  "default.log",
			input: defaultFile,
			env:   "LOG_ANALYZER_FILE",
		},
		{
			desc:  "check default outp",
			want:  "stdout",
			input: defaultOutput,
			env:   "LOG_ANALYZER_OUTPUT",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := getStrEnv(tC.env, tC.input)
			assert.Equal(t, tC.want, result)
		})
	}
}

func TestParseLine(t *testing.T) {
	testCases := []struct {
		want      string
		desc      string
		input     string
		level     level
		errorTest bool
	}{
		{
			desc:  "check default info",
			want:  "04110",
			input: "19:11:30.405 info 04110 Option 'snapshot' has been set to the value ''",
			level: Info,
		},
		{
			desc:  "check default trace",
			want:  "04106",
			input: "19:11:30.409 trace 04106 Log Messages file 'RouterServer.lms' successfully loaded",
			level: Trace,
		},
		{
			desc:  "check default debug",
			want:  "20001",
			input: "19:11:32.584 debug 20001 interaction 0091036872e759ae is started",
			level: Debug,
		},
		{
			desc:      "check default err",
			want:      "",
			input:     "19:11:32.584 0091036872e759ae is started",
			level:     Info,
			errorTest: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			conf.logLevel = tC.level
			result, err := parseLine(tC.input)
			if tC.errorTest && err == nil {
				t.Errorf("error")
			}
			assert.Equal(t, tC.want, result)
		})
	}
}

func TestReadLog(t *testing.T) {
	testCases := []struct {
		want      int
		desc      string
		input     string
		level     level
		path      string
		errorTest bool
	}{
		{
			desc:  "check default info",
			want:  122,
			input: "04500",
			level: Info,
		},
		{
			desc:  "check default trace",
			want:  8,
			input: "04503",
			level: Trace,
		},
		{
			desc:  "check default debug",
			want:  24,
			input: "20001",
			level: Debug,
		},
		{
			desc:      "check default err",
			want:      0,
			input:     "19:11:32.584 0091036872e759ae is started",
			level:     Info,
			errorTest: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			conf.logLevel = tC.level

			if tC.errorTest {
				conf.logFile = "./default23.log"
				_, err := ReadLog()
				if err == nil {
					t.Errorf("error")
				}
			} else {
				conf.logFile = "./default.log"
				result, _ := ReadLog()
				assert.Equal(t, tC.want, result[tC.input])
			}
		})
	}
}

func TestWriteResult(t *testing.T) {
	testCases := []struct {
		want      error
		desc      string
		input     map[string]int
		outp      string
		errorTest bool
	}{
		{
			desc: "check default info",
			want: nil,
			input: map[string]int{
				"04106": 1,
				"4v4":   6,
			},
			outp: "outp.txt",
		},
		{
			desc: "check default trace",
			want: nil,
			input: map[string]int{
				"04106":  1,
				"234234": 6,
			},
			outp: defaultOutput,
		},
		{
			desc: "check default debug",
			want: nil,
			input: map[string]int{
				"dfgfdgb": 1,
				"234":     6,
			},
			outp: "outp.txt",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			conf.statOutput = "outp.txt"
			err := WriteResult(tC.input)
			assert.Equal(t, tC.want, err)
		})
	}
}
