package statistics

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gnasnik/titan-explorer/config"
	"github.com/gnasnik/titan-explorer/core/dao"
	"github.com/gnasnik/titan-explorer/core/generated/model"
	"github.com/gnasnik/titan-explorer/utils"
	"github.com/oschwald/geoip2-golang"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

const maxPageSize = 100

type NodeFetcher struct {
	BaseFetcher
}

func newNodeFetcher() *NodeFetcher {
	return &NodeFetcher{BaseFetcher: newBaseFetcher()}
}

func (n *NodeFetcher) Fetch(ctx context.Context, scheduler *Scheduler) error {
	log.Info("start to fetch 【all nodes】")
	start := time.Now()
	defer func() {
		log.Infof("fetch all nodes done, cost: %v", time.Since(start))
	}()

	var total int64
	page, size := 1, maxPageSize

loop:
	offset := (page - 1) * size
	resp, err := scheduler.Api.GetNodeList(ctx, offset, size)
	if err != nil {
		log.Errorf("api ListNodes: %v", err)
		return nil
	}

	total += int64(len(resp.Data))
	page++

	var nodes []*model.DeviceInfo
	for _, node := range resp.Data {
		if node.NodeID == "" {
			continue
		}
		nodeInfo := toDeviceInfo(ctx, node)
		switch node.Status {
		case 0:
			nodeInfo.DeviceStatus = "offline"
			nodeInfo.DeviceID = node.NodeID
			// just update device status
			_ = dao.UpdateDeviceStatus(ctx, nodeInfo)
			continue
		case 1:
			nodeInfo.DeviceStatus = "online"
		default:
			nodeInfo.DeviceStatus = node.Status.String()
		}
		//nodeInfo.IpLocation = scheduler.AreaId
		nodeInfo.ActiveStatus = 1
		nodeInfo.CpuUsage = node.CPUUsage
		nodeInfo.MemoryUsage = node.MemoryUsage
		nodeInfo.BandwidthUp = float64(node.BandwidthUp)
		nodeInfo.BandwidthDown = float64(node.BandwidthDown)
		nodeInfo.DiskSpace = node.DiskSpace
		nodeInfo.DiskUsage = node.DiskUsage
		nodeInfo.CumulativeProfit = node.Profit
		nodeInfo.DeviceID = node.NodeID
		nodeInfo.DeviceName = node.NodeName
		nodeInfo.OnlineTime = float64(node.OnlineDuration)
		nodeInfo.NodeType = int64(node.Type)
		nodeInfo.UploadTraffic = float64(node.UploadTraffic)
		nodeInfo.DownloadTraffic = float64(node.DownloadTraffic)
		nodeInfo.CacheCount = node.AssetCount
		nodeInfo.RetrievalCount = node.RetrieveCount
		nodeInfo.ExternalIp = node.ExternalIP
		nodeInfo.InternalIp = node.InternalIP
		nodeInfo.BoundAt = node.FirstTime
		var loc model.Location
		if node.ExternalIP != "" {
			e := GetIpLocation(ctx, node.ExternalIP, &loc)
			if e != nil {
				log.Errorf("%v", e)
				GetGip(nodeInfo)

			} else {
				nodeInfo.NetworkInfo = loc.Isp
				nodeInfo.IpProvince = loc.Province
				continent := loc.Continent
				nodeInfo.IpCountry = loc.Country
				nodeInfo.IpCity = loc.City
				nodeInfo.IpLocation = continent + "-" + nodeInfo.IpCountry + "-" + nodeInfo.IpProvince
				if nodeInfo.IpCity != "" {
					nodeInfo.IpLocation += "-" + nodeInfo.IpCity
				}
				nodeInfo.Longitude, _ = strconv.ParseFloat(loc.Longitude, 64)
				nodeInfo.Latitude, _ = strconv.ParseFloat(loc.Latitude, 64)
			}
		}
		nodes = append(nodes, nodeInfo)
	}
	if len(nodes) < 1 {
		log.Errorf("start to fetch all nodes: nodes length is 0")
		return nil
	}
	log.Infof("handling %d/%d nodes", total, resp.Total)

	n.Push(ctx, func() error {
		e := dao.BulkUpsertDeviceInfo(ctx, nodes)
		if e != nil {
			log.Errorf("bulk upsert device info: %v", e)
		}
		if e = addDeviceInfoHours(ctx, nodes); err != nil {
			log.Errorf("add device info hours: %v", err)
		}
		return nil
	})

	if total < resp.Total {
		goto loop
	}

	// add inactive node records for statistics
	e := dao.GenerateInactiveNodeRecords(context.Background(), start)
	if e != nil {
		log.Errorf("generate inactive node records: %v", e)
	}

	return nil
}

func toDeviceInfo(ctx context.Context, v interface{}) *model.DeviceInfo {
	var deviceInfo model.DeviceInfo
	data, err := json.Marshal(v)
	if err != nil {
		log.Errorf("marshal device info: %v", err)
		return nil
	}
	dataMap := make(map[string]interface{})
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(data, &deviceInfo)
	if err != nil {
		return nil
	}
	deviceInfo.CpuUsage = utils.ToFixed(deviceInfo.CpuUsage, 2)
	deviceInfo.MemoryUsage = utils.ToFixed(deviceInfo.MemoryUsage, 2)
	//deviceInfo.BandwidthUp = utils.ToFixed(deviceInfo.BandwidthUp/gigaBytes, 2)
	deviceInfo.BandwidthUp = utils.ToFixed(deviceInfo.BandwidthUp, 2)
	deviceInfo.BandwidthDown = utils.ToFixed(deviceInfo.BandwidthDown, 2)
	deviceInfo.DiskSpace = utils.ToFixed(deviceInfo.DiskSpace, 2)
	deviceInfo.DiskUsage = utils.ToFixed(deviceInfo.DiskUsage, 2)
	deviceInfo.ActiveStatus = 1
	var loc model.Location
	if deviceInfo.ExternalIp != "" {
		err = GetIpLocation(ctx, deviceInfo.ExternalIp, &loc)
		if err != nil {
			log.Errorf("%v", err)
			GetGip(&deviceInfo)

		} else {
			deviceInfo.NetworkInfo = loc.Isp
			deviceInfo.IpProvince = loc.Province
			continent := loc.Continent
			deviceInfo.IpCountry = loc.Country
			deviceInfo.IpCity = loc.City
			deviceInfo.IpLocation = continent + "-" + deviceInfo.IpCountry + "-" + deviceInfo.IpProvince
			if deviceInfo.IpCity != "" {
				deviceInfo.IpLocation += "-" + deviceInfo.IpCity
			}
			deviceInfo.Longitude, _ = strconv.ParseFloat(loc.Longitude, 64)
			deviceInfo.Latitude, _ = strconv.ParseFloat(loc.Latitude, 64)
		}
	}
	return &deviceInfo
}

var _ Fetcher = &NodeFetcher{}

func GetGip(deviceInfo *model.DeviceInfo) *model.DeviceInfo {
	db, err := geoip2.Open("city.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	if deviceInfo.ExternalIp == "" {
		return deviceInfo
	}
	ip := net.ParseIP(deviceInfo.ExternalIp)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	if len(record.Subdivisions) > 0 {
		deviceInfo.IpProvince = record.Subdivisions[0].Names["en"]
		continent := record.Continent.Names["en"]
		deviceInfo.IpCountry = record.Country.Names["en"]
		deviceInfo.IpCity = record.City.Names["en"]
		deviceInfo.IpLocation = continent + "-" + deviceInfo.IpCountry + "-" + deviceInfo.IpProvince
		if deviceInfo.IpCity != "" {
			deviceInfo.IpLocation += "-" + deviceInfo.IpCity
		}
	}
	deviceInfo.Longitude = record.Location.Longitude
	deviceInfo.Latitude = record.Location.Latitude

	return deviceInfo
}

func GetIpLocation(ctx context.Context, ip string, Loc *model.Location) error {
	// get info from databases
	err := dao.GetLocationInfoByIp(ctx, ip, Loc)
	if err != nil {
		return err
	}
	if Loc.Ip != "" {
		return nil
	}
	// get info from api
	err = httpGetApi(ctx, ip, Loc)
	if err != nil {
		log.Errorf("api get loction err:%v", err)
		return err
	}
	return nil
}

func httpGetApi(ctx context.Context, ip string, Loc *model.Location) error {
	url := fmt.Sprintf("%s?ip=%s&key=%s", config.Cfg.IpUrl, ip, config.Cfg.IpKey)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var rsp LocationInfoRes
	err = json.Unmarshal(body, &rsp)
	if err != nil {
		return err
	}
	*Loc = rsp.Data.Location
	err = dao.UpsertLocationInfo(ctx, Loc)
	if err != nil {
		log.Errorf("update location err:%v", err)
	}
	return nil
}

type LocationInfoRes struct {
	Code int    `json:"code"`
	Data Data   `json:"data"`
	Msg  string `json:"msg"`
}

type Data struct {
	Location model.Location `json:"location"`
}
