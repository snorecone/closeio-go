package closeio

import (
	"encoding/json"
)

type Lead struct {
	Name          string            `json:"name,omitempty"`
	Url           string            `json:"url,omitempty"`
	Description   string            `json:"description,omitempty"`
	StatusId      string            `json:"status_id,omitempty"`
	Status        string            `json:"status,omitempty"`
	Contacts      *[]Contact        `json:"contacts"`
	Custom        map[string]string `json:"custom,omitempty"`
	Addresses     *[]Address        `json:"addresses"`
	Opportunities *[]Opportunity    `json:"opportunities"`
}

type LeadResp struct {
	StatusId    string `json:"status_id"`
	StatusLabel string `json:"status_label"`
	//Tasks          []string          `json:"tasks"` // TODO: change this
	DisplayName    string            `json:"display_name"`
	Description    string            `json:"description"`
	Addresses      []Address         `json:"addresses"`
	Custom         map[string]string `json:"custom"`
	Name           string            `json:"name"`
	Contacts       []ContactResp     `json:"contacts"`
	Url            string            `json:"url"`
	Id             string            `json:"id"`
	Opportunities  []Opportunity     `json:"opportunities"`
	DateUpdated    string            `json:"date_updated"`
	DateCreated    string            `json:"date_created"`
	CreatedBy      string            `json:"created_by"`
	UpdatedBy      string            `json:"updated_by"`
	OrganizationId string            `json:"organization_id"`
	HtmlUrl        string            `json:"html_url"`
}
type Leads struct {
	HasMore      bool       `json:"has_more"`
	TotalResults int        `json:"total_results"`
	Data         []LeadResp `json:"data"`
}

func (c *Closeio) Leads() (l *Leads, err error) {
	resp, err := request("lead", "GET", c.Token, nil)
	if err != nil {
		return nil, err
	}
	leads := Leads{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&leads)
	if err != nil {
		return nil, err
	}
	return &leads, nil
}

func (c *Closeio) Create(lead *Lead) (l *LeadResp, err error) {
	data, err := marshal(lead)
	if err != nil {
		return nil, err
	}
	resp, err := request("lead", "POST", c.Token, data)
	if err != nil {
		return nil, err
	}
	leadresp := LeadResp{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&leadresp)
	if err != nil {
		return nil, err
	}
	return &leadresp, nil
}
