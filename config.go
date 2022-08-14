package tlog

type configWrapper struct {
	Log []*Config `yaml:"log"`
}

// Config 日志配置
type Config struct {
	Type  string `yaml:"type"`  // 日志类型：console/file
	Level string `yaml:"level"` // 日志等级：debug/info/warn/error/fatal

	// 以下配置仅在类型为file时生效
	Prefix       string `yaml:"prefix"`        // 文件名前缀，例如prefix为tmp，则文件名为tmp.log
	MaxAge       string `yaml:"max-age"`       // 文件最大保存时间，使用time.ParseDuration函数进行计算
	RotationTime string `yaml:"rotation-time"` // 文件切割时间间隔，使用time.ParseDuration函数进行计算
	RotationSize int64  `yaml:"rotation-size"` // 文件最大大小，单位MB
}
