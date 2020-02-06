package targets

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/grafana/loki/pkg/promtail/scrape"
	"github.com/prometheus/common/model"
	"github.com/stretchr/testify/require"
)

type line struct {
	labels model.LabelSet
	entry  string
}

type clientRecorder struct {
	recorded []line
}

func (c *clientRecorder) Handle(labels model.LabelSet, time time.Time, entry string) error {
	c.recorded = append(c.recorded, line{labels: labels, entry: entry})
	return nil
}

func Test_newReaderTarget(t *testing.T) {
	tests := []struct {
		name    string
		in      io.Reader
		cfg     scrape.Config
		want    []line
		wantErr bool
	}{
		{
			"no newlines",
			bytes.NewReader([]byte("bar")),
			scrape.Config{},
			[]line{
				{nil, "bar"},
			},
			false,
		},
		{
			"empty",
			bytes.NewReader([]byte("")),
			scrape.Config{},
			nil,
			false,
		},
		{
			"newlines",
			bytes.NewReader([]byte("\nfoo\r\nbar")),
			scrape.Config{},
			[]line{
				{nil, "foo"},
				{nil, "bar"},
			},
			false,
		},
		{
			"default config",
			bytes.NewReader([]byte("\nfoo\r\nbar")),
			defaultStdInCfg,
			[]line{
				{model.LabelSet{"job": "stdin", "hostname": model.LabelValue(hostName)}, "foo"},
				{model.LabelSet{"job": "stdin", "hostname": model.LabelValue(hostName)}, "bar"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := &clientRecorder{}
			got, err := newReaderTarget(tt.in, recorder, tt.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("newReaderTarget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			<-got.ctx.Done()
			require.Equal(t, tt.want, recorder.recorded)
		})
	}
}

type mockShutdownable struct {
	called chan bool
}

func (m *mockShutdownable) Shutdown() {
	m.called <- true
}

type fakeStdin struct {
	io.Reader
	os.FileInfo
}

func newFakeStin(data string) *fakeStdin {
	return &fakeStdin{
		Reader: strings.NewReader(data),
	}
}

func (f fakeStdin) Stat() (os.FileInfo, error) { return f.FileInfo, nil }

func Test_Shutdown(t *testing.T) {
	stdIn = newFakeStin("line")
	appMock := &mockShutdownable{called: make(chan bool, 1)}
	recorder := &clientRecorder{}
	manager, err := newStdinTargetManager(appMock, recorder, []scrape.Config{{}})
	require.NoError(t, err)
	require.NotNil(t, manager)
	called := <-appMock.called
	require.Equal(t, true, called)
	require.Equal(t, []line{{labels: nil, entry: "line"}}, recorder.recorded)
}

func Test_StdinConfigs(t *testing.T) {

	// should take the first config
	require.Equal(t, scrape.DefaultScrapeConfig, getStdinConfig([]scrape.Config{
		scrape.DefaultScrapeConfig,
		scrape.Config{},
	}))
	// or use the default if none if provided
	require.Equal(t, defaultStdInCfg, getStdinConfig([]scrape.Config{}))
}

type mockFileInfo struct{}

func (mockFileInfo) Name() string       { return "" }
func (mockFileInfo) Size() int64        { return 1 }
func (mockFileInfo) Mode() os.FileMode  { return 1 }
func (mockFileInfo) ModTime() time.Time { return time.Now() }
func (mockFileInfo) Sys() interface{}   { return nil }
func (mockFileInfo) IsDir() bool        { return false }

func Test_isPipe(t *testing.T) {
	fake := newFakeStin("line")
	fake.FileInfo = &mockFileInfo{}
	stdIn = fake
	require.Equal(t, true, isStdinPipe())
	stdIn = os.Stdin
	require.Equal(t, false, isStdinPipe())
}
