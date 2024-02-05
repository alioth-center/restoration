package restoration

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alioth-center/infrastructure/logger"
	"github.com/alioth-center/infrastructure/trace"
	"github.com/alioth-center/infrastructure/utils/timezone"
	"github.com/alioth-center/restoration/proto/pbg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

var (
	timeFormat    = "2006.01.02-15:04:05.000Z07:00"
	levelValueMap = map[logger.Level]int{
		logger.LevelDebug: 0,
		logger.LevelInfo:  1,
		logger.LevelWarn:  2,
		logger.LevelError: 3,
		logger.LevelFatal: 4,
		logger.LevelPanic: 5,
	}
)

func NewLogger(endpoint string, level logger.Level, selfIdentification string) (logger.Logger, error) {
	conn, dialErr := grpc.DialContext(trace.NewContext(), endpoint, grpc.WithCredentialsBundle(insecure.NewBundle()))
	if dialErr != nil {
		return nil, &DialRestorationServiceError{Endpoint: endpoint, ErrorOccurred: dialErr}
	}

	return Logger{cli: pbg.NewRestorationServiceClient(conn), level: level, self: selfIdentification}, nil
}

type Logger struct {
	cli   pbg.RestorationServiceClient
	level logger.Level
	self  string
}

func (r Logger) parseFields(fields logger.Fields) (ctx context.Context, request *pbg.RestorationCollectionRequest) {
	entry := fields.Export()

	// parse call time
	callTime, parseErr := time.Parse(timeFormat, entry.CallTime)
	if parseErr != nil {
		// unable to parse call time, use current time
		callTime = timezone.NowInZeroTime()
	} else {
		// parse call time success, convert to zero time
		callTime = callTime.In(timezone.GetFixedTimeZone(timezone.LocationUTC))
	}

	// marshal entry data
	data, marshalErr := json.Marshal(entry.Data)
	if marshalErr != nil {
		// unable to marshal data field, use default marshal
		data = []byte(fmt.Sprintf("%+v", entry.Data))
	}

	// marshal entry extra
	extra := map[string][]byte{}
	if entry.Extra != nil && len(entry.Extra) > 0 {
		for k, v := range entry.Extra {
			data, marshalErr := json.Marshal(v)
			if marshalErr != nil {
				// unable to marshal extra field, use default marshal
				extra[k] = []byte(fmt.Sprintf("%+v", v))
			} else {
				extra[k] = data
			}
		}
	}

	request = &pbg.RestorationCollectionRequest{
		File:           entry.File,
		Level:          entry.Level,
		CallTime:       callTime.UnixMilli(),
		Service:        entry.Service,
		TraceId:        entry.TraceID,
		Message:        entry.Message,
		Data:           data,
		Extra:          extra,
		Identification: r.self,
	}

	return trace.NewContext(), request
}

func (r Logger) Debug(fields logger.Fields) {
	r.Log(logger.LevelDebug, fields)
}

func (r Logger) Info(fields logger.Fields) {
	r.Log(logger.LevelInfo, fields)
}

func (r Logger) Warn(fields logger.Fields) {
	r.Log(logger.LevelWarn, fields)
}

func (r Logger) Error(fields logger.Fields) {
	r.Log(logger.LevelError, fields)
}

func (r Logger) Fatal(fields logger.Fields) {
	r.Log(logger.LevelFatal, fields)
}

func (r Logger) Panic(fields logger.Fields) {
	r.Log(logger.LevelPanic, fields)
}

func (r Logger) Log(level logger.Level, fields logger.Fields) {
	if levelValueMap[level] >= levelValueMap[r.level] {
		// log level is greater than or equal to the specified level, send log to restoration service
		ctx, request := r.parseFields(fields)
		request.Level = string(level)
		_, e := r.cli.RestorationCollection(ctx, request)
		if e != nil {
			// log collection failed, print error
			fmt.Println("restoration log collection failed:", e)
		}
	}
}

func (r Logger) Logf(level logger.Level, fields logger.Fields, format string, args ...any) {
	r.Log(level, fields.WithMessage(fmt.Sprintf(format, args...)))
}

func (r Logger) Debugf(fields logger.Fields, format string, args ...any) {
	r.Logf(logger.LevelDebug, fields, format, args...)
}

func (r Logger) Infof(fields logger.Fields, format string, args ...any) {
	r.Logf(logger.LevelInfo, fields, format, args...)
}

func (r Logger) Warnf(fields logger.Fields, format string, args ...any) {
	r.Logf(logger.LevelWarn, fields, format, args...)
}

func (r Logger) Errorf(fields logger.Fields, format string, args ...any) {
	r.Logf(logger.LevelError, fields, format, args...)
}

func (r Logger) Fatalf(fields logger.Fields, format string, args ...any) {
	r.Logf(logger.LevelFatal, fields, format, args...)
}

func (r Logger) Panicf(fields logger.Fields, format string, args ...any) {
	r.Logf(logger.LevelPanic, fields, format, args...)
}
