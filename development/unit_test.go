package restoration

import (
	"github.com/alioth-center/infrastructure/logger"
	"github.com/alioth-center/infrastructure/trace"
	"github.com/alioth-center/infrastructure/utils/generate"
	"testing"
)

func TestRestoration(t *testing.T) {
	log, err := NewLogger("127.0.0.1:28081", logger.LevelInfo, "restoration_test")
	if err != nil {
		t.Error(err)
		return
	}

	tid := generate.UUID()
	t.Log(tid)

	type TestStruct struct {
		Key   string `json:"key"`
		Value int    `json:"value"`
	}

	log.Info(logger.NewFields(trace.NewContextWithTid(tid)).WithMessage("test message").WithData(&TestStruct{Key: "test", Value: 1}).WithField("test_extra_field", &TestStruct{Key: "extra", Value: 114514}))
}

func BenchmarkTestRestoration(b *testing.B) {
	log, err := NewLogger("127.0.0.1:28081", logger.LevelInfo, "restoration_test")
	if err != nil {
		b.Error(err)
		return
	}

	tid := generate.UUID()
	b.Log(tid)

	type TestStruct struct {
		Key   string `json:"key"`
		Value int    `json:"value"`
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Info(logger.NewFields(trace.NewContextWithTid(tid)).WithMessage("test message").WithData(&TestStruct{Key: "test", Value: 1}).WithField("test_extra_field", &TestStruct{Key: "extra", Value: 114514}))
	}
}
