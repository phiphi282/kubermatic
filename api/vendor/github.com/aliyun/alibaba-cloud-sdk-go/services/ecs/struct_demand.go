package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Demand is a nested struct in ecs response
type Demand struct {
	ZoneId             string      `json:"ZoneId" xml:"ZoneId"`
	DemandTime         string      `json:"DemandTime" xml:"DemandTime"`
	InstanceTypeFamily string      `json:"InstanceTypeFamily" xml:"InstanceTypeFamily"`
	InstanceType       string      `json:"InstanceType" xml:"InstanceType"`
	InstanceChargeType string      `json:"InstanceChargeType" xml:"InstanceChargeType"`
	Period             int         `json:"Period" xml:"Period"`
	PeriodUnit         string      `json:"PeriodUnit" xml:"PeriodUnit"`
	StartTime          string      `json:"StartTime" xml:"StartTime"`
	EndTime            string      `json:"EndTime" xml:"EndTime"`
	DemandStatus       string      `json:"DemandStatus" xml:"DemandStatus"`
	TotalAmount        int         `json:"TotalAmount" xml:"TotalAmount"`
	AvailableAmount    int         `json:"AvailableAmount" xml:"AvailableAmount"`
	UsedAmount         int         `json:"UsedAmount" xml:"UsedAmount"`
	DeliveringAmount   int         `json:"DeliveringAmount" xml:"DeliveringAmount"`
	SupplyInfos        SupplyInfos `json:"SupplyInfos" xml:"SupplyInfos"`
}
