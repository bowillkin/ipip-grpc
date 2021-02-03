package handler

import (
	"context"

	"github.com/ipipdotnet/ipdb-go"

	"github.com/bowillkin/proto/ipip"
)

type IpIp struct {
	DB *ipdb.City
}

func (i *IpIp) GetAreaDataByIp(ctx context.Context, req *ipip.GetAreaDataByIpReq) (*ipip.GetAreaDataByIpResp, error) {
	res := new(ipip.GetAreaDataByIpResp)
	data, err := i.DB.FindMap(req.RemoteIp, "CN")
	if err != nil {
		return nil, err
	}
	var cityName, countryName, regionName string
	countryName = data["country_name"]
	regionName = data["region_name"]
	if regionName == "" {
		regionName = countryName
	}
	cityName = data["city_name"]
	if cityName == "" {
		cityName = regionName
	}
	res.City = cityName
	res.Country = countryName
	res.Province = regionName
	return res, nil
}
