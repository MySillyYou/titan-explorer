// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package model

import (
	"database/sql"
	"time"
)

type Application struct {
	ID                int64     `db:"id" json:"id"`
	UserID            string    `db:"user_id" json:"user_id"`
	Email             string    `db:"email" json:"email"`
	IpCountry         string    `db:"ip_country" json:"ip_country"`
	IpCity            string    `db:"ip_city" json:"ip_city"`
	NodeType          int32     `db:"node_type" json:"node_type"`
	Amount            int32     `db:"amount" json:"amount"`
	UpstreamBandwidth float64   `db:"upstream_bandwidth" json:"upstream_bandwidth"`
	DiskSpace         float64   `db:"disk_space" json:"disk_space"`
	Status            int32     `db:"status" json:"status"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time `db:"updated_at" json:"updated_at"`
}

type ApplicationResult struct {
	ID            int64     `db:"id" json:"id"`
	ApplicationID int64     `db:"application_id" json:"application_id"`
	UserID        string    `db:"user_id" json:"user_id"`
	DeviceID      string    `db:"device_id" json:"device_id"`
	NodeType      int32     `db:"node_type" json:"node_type"`
	Secret        string    `db:"secret" json:"secret"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

type BlockInfo struct {
	ID          int64         `db:"id" json:"id"`
	DeviceID    string        `db:"device_id" json:"device_id"`
	CarfileHash string        `db:"carfile_hash" json:"carfile_hash"`
	CarfileCid  string        `db:"carfile_cid" json:"carfile_cid"`
	Status      sql.NullInt32 `db:"status" json:"status"`
	Size        sql.NullInt32 `db:"size" json:"size"`
	CreatedTime sql.NullTime  `db:"created_time" json:"created_time"`
	EndTime     sql.NullTime  `db:"end_time" json:"end_time"`
}

type CacheEvent struct {
	ID         int64     `db:"id" json:"id"`
	DeviceID   string    `db:"device_id" json:"device_id"`
	CarfileCid string    `db:"carfile_cid" json:"carfile_cid"`
	BlockSize  float64   `db:"block_size" json:"block_size"`
	Blocks     int64     `db:"blocks" json:"blocks"`
	Time       time.Time `db:"time" json:"time"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}

type DeviceInfo struct {
	ID            int64     `db:"id" json:"id"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt     time.Time `db:"deleted_at" json:"deleted_at"`
	DeviceID      string    `db:"device_id" json:"device_id"`
	SchedulerID   string    `db:"scheduler_id" json:"scheduler_id"`
	NodeType      int32     `db:"node_type" json:"node_type"`
	DeviceRank    int32     `db:"device_rank" json:"device_rank"`
	DeviceName    string    `db:"device_name" json:"device_name"`
	UserID        string    `db:"user_id" json:"user_id"`
	SnCode        string    `db:"sn_code" json:"sn_code"`
	Operator      string    `db:"operator" json:"operator"`
	NetworkType   string    `db:"network_type" json:"network_type"`
	SystemVersion string    `db:"system_version" json:"system_version"`
	ProductType   string    `db:"product_type" json:"product_type"`
	NetworkInfo   string    `db:"network_info" json:"network_info"`
	ExternalIp    string    `db:"external_ip" json:"external_ip"`
	InternalIp    string    `db:"internal_ip" json:"internal_ip"`
	IpLocation    string    `db:"ip_location" json:"ip_location"`
	IpCountry     string    `db:"ip_country" json:"ip_country"`
	IpCity        string    `db:"ip_city" json:"ip_city"`
	MacLocation   string    `db:"mac_location" json:"mac_location"`
	NatType       string    `db:"nat_type" json:"nat_type"`
	Upnp          string    `db:"upnp" json:"upnp"`
	PkgLossRatio  float64   `db:"pkg_loss_ratio" json:"pkg_loss_ratio"`
	// Nat
	NatRatio         float64 `db:"nat_ratio" json:"nat_ratio"`
	Latency          float64 `db:"latency" json:"latency"`
	CpuUsage         float64 `db:"cpu_usage" json:"cpu_usage"`
	CpuCores         int32   `db:"cpu_cores" json:"cpu_cores"`
	MemoryUsage      float64 `db:"memory_usage" json:"memory_usage"`
	Memory           float64 `db:"memory" json:"memory"`
	DiskUsage        float64 `db:"disk_usage" json:"disk_usage"`
	DiskSpace        float64 `db:"disk_space" json:"disk_space"`
	WorkStatus       string  `db:"work_status" json:"work_status"`
	DeviceStatus     string  `db:"device_status" json:"device_status"`
	DiskType         string  `db:"disk_type" json:"disk_type"`
	IoSystem         string  `db:"io_system" json:"io_system"`
	OnlineTime       float64 `db:"online_time" json:"online_time"`
	TodayOnlineTime  float64 `db:"today_online_time" json:"today_online_time"`
	TodayProfit      float64 `db:"today_profit" json:"today_profit"`
	YesterdayProfit  float64 `db:"yesterday_profit" json:"yesterday_profit"`
	SevenDaysProfit  float64 `db:"seven_days_profit" json:"seven_days_profit"`
	MonthProfit      float64 `db:"month_profit" json:"month_profit"`
	CumulativeProfit float64 `db:"cumulative_profit" json:"cumulative_profit"`
	BandwidthUp      float64 `db:"bandwidth_up" json:"bandwidth_up"`
	BandwidthDown    float64 `db:"bandwidth_down" json:"bandwidth_down"`
	TotalDownload    float64 `db:"total_download" json:"total_download"`
	TotalUpload      float64 `db:"total_upload" json:"total_upload"`
	BlockCount       int64   `db:"block_count" json:"block_count"`
	DownloadCount    int64   `db:"download_count" json:"download_count"`
}

type DeviceInfoDaily struct {
	ID                int64     `db:"id" json:"id"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt         time.Time `db:"deleted_at" json:"deleted_at"`
	UserID            string    `db:"user_id" json:"user_id"`
	DeviceID          string    `db:"device_id" json:"device_id"`
	Time              time.Time `db:"time" json:"time"`
	Income            float64   `db:"income" json:"income"`
	OnlineTime        float64   `db:"online_time" json:"online_time"`
	PkgLossRatio      float64   `db:"pkg_loss_ratio" json:"pkg_loss_ratio"`
	Latency           float64   `db:"latency" json:"latency"`
	NatRatio          float64   `db:"nat_ratio" json:"nat_ratio"`
	DiskUsage         float64   `db:"disk_usage" json:"disk_usage"`
	UpstreamTraffic   float64   `db:"upstream_traffic" json:"upstream_traffic"`
	DownstreamTraffic float64   `db:"downstream_traffic" json:"downstream_traffic"`
	RetrieveCount     int64     `db:"retrieve_count" json:"retrieve_count"`
	BlockCount        int64     `db:"block_count" json:"block_count"`
}

type DeviceInfoHour struct {
	ID                int64     `db:"id" json:"id"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt         time.Time `db:"deleted_at" json:"deleted_at"`
	UserID            string    `db:"user_id" json:"user_id"`
	DeviceID          string    `db:"device_id" json:"device_id"`
	Time              time.Time `db:"time" json:"time"`
	HourIncome        float64   `db:"hour_income" json:"hour_income"`
	OnlineTime        float64   `db:"online_time" json:"online_time"`
	PkgLossRatio      float64   `db:"pkg_loss_ratio" json:"pkg_loss_ratio"`
	Latency           float64   `db:"latency" json:"latency"`
	NatRatio          float64   `db:"nat_ratio" json:"nat_ratio"`
	DiskUsage         float64   `db:"disk_usage" json:"disk_usage"`
	UpstreamTraffic   float64   `db:"upstream_traffic" json:"upstream_traffic"`
	DownstreamTraffic float64   `db:"downstream_traffic" json:"downstream_traffic"`
	RetrieveCount     int64     `db:"retrieve_count" json:"retrieve_count"`
	BlockCount        int64     `db:"block_count" json:"block_count"`
}

type FullNodeInfo struct {
	ID                       int64     `db:"id" json:"id"`
	TotalNodeCount           int32     `db:"total_node_count" json:"total_node_count"`
	ValidatorCount           int32     `db:"validator_count" json:"validator_count"`
	CandidateCount           int32     `db:"candidate_count" json:"candidate_count"`
	EdgeCount                int32     `db:"edge_count" json:"edge_count"`
	TotalStorage             float64   `db:"total_storage" json:"total_storage"`
	TotalUpstreamBandwidth   float64   `db:"total_upstream_bandwidth" json:"total_upstream_bandwidth"`
	TotalDownstreamBandwidth float64   `db:"total_downstream_bandwidth" json:"total_downstream_bandwidth"`
	TotalCarfile             int64     `db:"total_carfile" json:"total_carfile"`
	TotalCarfileSize         float64   `db:"total_carfile_size" json:"total_carfile_size"`
	RetrievalCount           int64     `db:"retrieval_count" json:"retrieval_count"`
	NextElectionTime         time.Time `db:"next_election_time" json:"next_election_time"`
	Time                     time.Time `db:"time" json:"time"`
	CreatedAt                time.Time `db:"created_at" json:"created_at"`
	UpdatedAt                time.Time `db:"updated_at" json:"updated_at"`
}

type LoginLog struct {
	ID            int64     `db:"id" json:"id"`
	LoginUsername string    `db:"login_username" json:"login_username"`
	IpAddress     string    `db:"ip_address" json:"ip_address"`
	LoginLocation string    `db:"login_location" json:"login_location"`
	Browser       string    `db:"browser" json:"browser"`
	Os            string    `db:"os" json:"os"`
	Status        int32     `db:"status" json:"status"`
	Msg           string    `db:"msg" json:"msg"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}

type OperationLog struct {
	ID               int64     `db:"id" json:"id"`
	Title            string    `db:"title" json:"title"`
	BusinessType     int32     `db:"business_type" json:"business_type"`
	Method           string    `db:"method" json:"method"`
	RequestMethod    string    `db:"request_method" json:"request_method"`
	OperatorType     int32     `db:"operator_type" json:"operator_type"`
	OperatorUsername string    `db:"operator_username" json:"operator_username"`
	OperatorUrl      string    `db:"operator_url" json:"operator_url"`
	OperatorIp       string    `db:"operator_ip" json:"operator_ip"`
	OperatorLocation string    `db:"operator_location" json:"operator_location"`
	OperatorParam    string    `db:"operator_param" json:"operator_param"`
	JsonResult       string    `db:"json_result" json:"json_result"`
	Status           int32     `db:"status" json:"status"`
	ErrorMsg         string    `db:"error_msg" json:"error_msg"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
}

type RetrievalInfo struct {
	ID             int64     `db:"id" json:"id"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt      time.Time `db:"deleted_at" json:"deleted_at"`
	ServiceCountry string    `db:"service_country" json:"service_country"`
	ServiceStatus  string    `db:"service_status" json:"service_status"`
	TaskStatus     string    `db:"task_status" json:"task_status"`
	FileName       string    `db:"file_name" json:"file_name"`
	FileSize       string    `db:"file_size" json:"file_size"`
	CreateTime     string    `db:"create_time" json:"create_time"`
	Cid            string    `db:"cid" json:"cid"`
	Price          float64   `db:"price" json:"price"`
	MinerID        string    `db:"miner_id" json:"miner_id"`
	UserID         string    `db:"user_id" json:"user_id"`
}

type RetrieveEvent struct {
	ID                int64     `db:"id" json:"id"`
	DeviceID          string    `db:"device_id" json:"device_id"`
	Blocks            int64     `db:"blocks" json:"blocks"`
	Time              time.Time `db:"time" json:"time"`
	UpstreamBandwidth float64   `db:"upstream_bandwidth" json:"upstream_bandwidth"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time `db:"updated_at" json:"updated_at"`
}

type Scheduler struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Group     string    `db:"group" json:"group"`
	Address   string    `db:"address" json:"address"`
	Status    int32     `db:"status" json:"status"`
	Token     string    `db:"token" json:"token"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at"`
}

type TaskInfo struct {
	ID             int64     `db:"id" json:"id"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt      time.Time `db:"deleted_at" json:"deleted_at"`
	UserID         string    `db:"user_id" json:"user_id"`
	MinerID        string    `db:"miner_id" json:"miner_id"`
	DeviceID       string    `db:"device_id" json:"device_id"`
	FileName       string    `db:"file_name" json:"file_name"`
	IpAddress      string    `db:"ip_address" json:"ip_address"`
	Cid            string    `db:"cid" json:"cid"`
	BandwidthUp    string    `db:"bandwidth_up" json:"bandwidth_up"`
	BandwidthDown  string    `db:"bandwidth_down" json:"bandwidth_down"`
	TimeNeed       string    `db:"time_need" json:"time_need"`
	Time           time.Time `db:"time" json:"time"`
	ServiceCountry string    `db:"service_country" json:"service_country"`
	Region         string    `db:"region" json:"region"`
	Status         string    `db:"status" json:"status"`
	Price          float64   `db:"price" json:"price"`
	FileSize       float64   `db:"file_size" json:"file_size"`
	DownloadUrl    string    `db:"download_url" json:"download_url"`
}

type User struct {
	ID        int64     `db:"id" json:"id"`
	Uuid      string    `db:"uuid" json:"uuid"`
	Username  string    `db:"username" json:"username"`
	PassHash  string    `db:"pass_hash" json:"pass_hash"`
	UserEmail string    `db:"user_email" json:"user_email"`
	Address   string    `db:"address" json:"address"`
	Role      int32     `db:"role" json:"role"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at"`
}
