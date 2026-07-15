package logger

import (
	"encoding/json"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type WsLogger struct {
	file    *zap.Logger
	console *zap.Logger
	seen    sync.Map
}

// @inject
func NewWsLogger() *WsLogger {
	consoleCfg := zap.NewDevelopmentEncoderConfig()
	consoleCfg.TimeKey = ""
	consoleCfg.LevelKey = ""
	consoleCfg.CallerKey = ""

	return &WsLogger{
		file: NewDailyFileZapLogger("logs/ws/"),
		console: zap.New(zapcore.NewCore(
			zapcore.NewConsoleEncoder(consoleCfg),
			zapcore.AddSync(os.Stdout),
			zapcore.DebugLevel,
		)),
	}
}

func (this *WsLogger) Connected(connID, channel, url string) {
	if _, loaded := this.seen.LoadOrStore(connID, struct{}{}); loaded {
		return
	}
	this.file.Info("ws connected",
		zap.String("conn_id", connID),
		zap.String("channel", channel),
		zap.String("url", url),
	)
}

func (this *WsLogger) Disconnected(connID, channel string, err error) {
	this.seen.Delete(connID)
	fields := []zap.Field{
		zap.String("conn_id", connID),
		zap.String("channel", channel),
	}
	if err != nil {
		fields = append(fields, zap.Error(err))
	}
	this.file.Info("ws disconnected", fields...)
}

func (this *WsLogger) MessageReceived(connID string, payload []byte) {
	if isPing(payload) {
		return
	}
	this.file.Info("ws recv", zap.String("conn_id", connID), zap.String("payload", prettyJSON(payload)))
	this.console.Sugar().Infof("[ws recv] [%s] %s", connID, prettyJSON(payload))
}

func (this *WsLogger) MessageSent(connID string, payload []byte) {
	if isPing(payload) {
		return
	}
	this.file.Info("ws send", zap.String("conn_id", connID), zap.String("payload", prettyJSON(payload)))
	this.console.Sugar().Infof("[ws send] [%s] %s", connID, prettyJSON(payload))
}

func (this *WsLogger) HandlerError(connID, channel, event string, err error) {
	this.file.Error("ws handler error",
		zap.String("conn_id", connID),
		zap.String("channel", channel),
		zap.String("event", event),
		zap.Error(err),
	)
	this.console.Sugar().Errorf("[ws error] [%s] %s: %v", connID, event, err)
}

func (this *WsLogger) SendError(connID string, err error) {
	this.file.Error("ws send error", zap.String("conn_id", connID), zap.Error(err))
	this.console.Sugar().Errorf("[ws send error] [%s] %v", connID, err)
}

func (this *WsLogger) DialError(url, channel string, err error) {
	this.file.Error("ws dial error",
		zap.String("url", url),
		zap.String("channel", channel),
		zap.Error(err),
	)
	this.console.Sugar().Errorf("[ws dial error] [%s] %v", url, err)
}

func (this *WsLogger) Info(msg string, fields ...zap.Field) {
	this.file.Info(msg, fields...)
}

func (this *WsLogger) Error(msg string, fields ...zap.Field) {
	this.file.Error(msg, fields...)
	this.console.Error(msg, fields...)
}

func (this *WsLogger) UpgradeError(channel, remoteAddr string, err error) {
	this.file.Error("ws upgrade error",
		zap.String("channel", channel),
		zap.String("remote_addr", remoteAddr),
		zap.Error(err),
	)
	this.console.Sugar().Errorf("[ws upgrade error] [%s] %s: %v", channel, remoteAddr, err)
}

func isPing(payload []byte) bool {
	var msg struct {
		EventType string `json:"event_type"`
	}
	return json.Unmarshal(payload, &msg) == nil && msg.EventType == "ping"
}

func prettyJSON(payload []byte) string {
	var v any
	if err := json.Unmarshal(payload, &v); err != nil {
		return string(payload)
	}
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return string(payload)
	}
	return string(out)
}
