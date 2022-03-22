package ingress

import (
	"github.com/eolinker/apinto-ingress-controller/pkg/config"
	"github.com/eolinker/eosc/log"
	"github.com/eolinker/eosc/log/filelog"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func InitLogger(cfg config.LogConfig) {

	level, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		dief("parse log level: %s", err)
	}
	lineFormatter := &log.LineFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	var t *log.Transporter
	switch cfg.LogOutput {
	case "stderr":
		t = log.NewTransport(os.Stderr, level)
	case "stdout":
		t = log.NewTransport(os.Stdout, level)
	default:
		writer := filelog.NewFileWriteByPeriod()
		writer.Set(filepath.Dir(cfg.LogOutput), filepath.Base(cfg.LogOutput), parsePeriod(cfg.LogPeriod), parseExpire(cfg.LogExpire))
		writer.Open()
		t = log.NewTransport(writer, level)
	}
	t.SetFormatter(lineFormatter)
	log.Reset(t)
}

// 解析过期时间，以天为单位
func parseExpire(errorLogExpire string) time.Duration {
	d, err := strconv.Atoi(errorLogExpire)
	if err != nil {
		return 7 * 24 * time.Hour
	}
	return time.Duration(d) * time.Hour
}

func parsePeriod(period string) filelog.LogPeriod {
	p, err := filelog.ParsePeriod(period)
	if err != nil {
		p = filelog.PeriodDay
	}
	return p
}
