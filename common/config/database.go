package config

import "time"

type Database struct {
	Driver       string `mapstructure:"driver" json:"driver" yaml:"driver"`
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	Database     string `mapstructure:"database" json:"database" yaml:"database"`
	UserName     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Charset      string `mapstructure:"charset" json:"charset" yaml:"charset"`                      // 编码格式
	MaxIdleConns int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"` // 空闲连接池中连接的最大数量
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"` // 打开数据库连接的最大数量

	ConnTimeout     time.Duration `mapstructure:"conn_timeout" json:"conn_timeout" yaml:"conn_timeout"`                   //链接过期超时时长
	ReadTimeout     time.Duration `mapstructure:"read_timeout" json:"read_timeout" yaml:"read_timeout"`                   //读取过期时长
	WriteTimeout    time.Duration `mapstructure:"write_timeout" json:"write_timeout" yaml:"write_timeout"`                //写入过期时长
	ConnMaxLifeTime time.Duration `mapstructure:"conn_max_life_time" json:"conn_max_life_time" yaml:"conn_max_life_time"` //连接可重用的最大时间长度  连接越短，连接终止的频率就越高，因此，需要从头开始创建连接的频率就越高内存消耗就大

	LogMode             string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`                                           //日志级别
	EnableFileLogWriter bool   `mapstructure:"enable_file_log_writer" json:"enable_file_log_writer" yaml:"enable_file_log_writer"` // 禁用彩色打印
	LogFilename         string `mapstructure:"log_filename" json:"log_filename" yaml:"log_filename"`                               //日志文件名称
}
