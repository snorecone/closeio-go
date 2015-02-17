package closeio

import (
	"encoding/json"
	"bytes"
	"errors"
	"io/ioutil"
	"fmt"
	"net/http"
)

const baseURL = "https://app.close.io/api"
const version = "v1"

type Closeio struct {
	Token string
}

type Lead struct {
	Name        string            `json:"name"`
	Url         string            `json:"url"`
	Description string            `json:"description"`
	StatusId    string            `json:"status_id,omitempty"`
	Contacts    []Contact         `json:"contacts"`
	Custom      map[string]string `json:"custom"`
	Addresses   []Address         `json:"addresses"`
}

type LeadResp struct {
	StatusId       string            `json:"status_id"`
	StatusLabel    string            `json:"status_label"`
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
type Opportunity struct {
}
type Contact struct {
	Name   string  `json:"name"`
	Title  string  `json:"title"`
	Emails []Email `json:"emails"`
	Phones []Phone `json:"phones"`
}
type ContactResp struct {
	Name           string      `json:"name"`
	Title          string      `json:"title"`
	Emails         []EmailResp `json:"emails"`
	Phones         []PhoneResp `json:"phones"`
	CreatedBy      string      `json:"created_by"`
	UpdatedBy      string      `json:"updated_by"`
	Id             string      `json:"id"`
	OrganizationId string      `json:"organization_id"`
	DateCreated    string      `json:"date_created"`
	DateUpdated    string      `json:"date_updated"`
}
type Email struct {
	Type  string `json:"type"`
	Email string `json:"email"`
}
type EmailResp struct {
	Type       string `json:"type"`
	Email      string `json:"email"`
	EmailLower string `json:"email_lower"`
}
type Phone struct {
	Type  string `json:"type"`
	Phone string `json:"phone"`
}
type PhoneResp struct {
	Type           string `json:"type"`
	Phone          string `json:"phone"`
	PhoneFormatted string `json:"phone_formatted"`
}
type Address struct {
	Label    string `json:"label"`
	Address1 string `json:"address_1"`
	Address2 string `json:"address_2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zipcode  string    `json:"zipcode"`
	Country  string `json:"country"`
}
type Leads struct {
	HasMore bool `json:"has_more"`
	TotalResults int `json:"total_results"`
	Data []LeadResp `json:"data"`
}

func New(token string) *Closeio {
	return &Closeio{token}
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

func (c *Closeio) Create(lead *Lead)(l *LeadResp, err error) {
	data, err := marshal(lead)
	if err != nil {
		return nil, err
	}
	resp, err := request("lead","POST", c.Token, data)
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
func marshal(data interface{}) (jsonD []byte, err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
func request(urlPart string,  reqType string, key string, data []byte) (resp *http.Response, err error) {
	client := &http.Client{}
	url := baseURL + "/"+version + "/"+ urlPart + "/"
	fmt.Println(url)
	body := bytes.NewBuffer(data)
	req, err := http.NewRequest(reqType, url, body)
	req.SetBasicAuth(key, "")

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		bod, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(bod))
	}
	return resp, nil
}
