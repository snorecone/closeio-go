package closeio

import (
	"testing"
	"log"
)

func TestList(t *testing.T) {

	closeAPI := New("")
	leads, err := closeAPI.Leads()
	if err != nil {
		log.Println(err)
	}
	log.Println(leads.Data[0])

}

func TestCreate(t *testing.T) {
	closeAPI := New("")
	lead := Lead{
			Name: "Bluth Company",
			Url: "http://thebluthcompany.tumblr.com/",
			Description: "Best. Show. Ever.",
			//StatusId: "stat_1ZdiZqcSIkoGVnNOyxiEY58eTGQmFNG3LPlEVQ4V7Nk",
			Contacts: []Contact{
				Contact{
					Name: "Gob",
					Title: "Sr. Vice President",
					Emails: []Email{
						Email{
							Type: "office",
							Email: "gob@example.com",
						},
					},
					Phones: []Phone{
						Phone{
							Type: "office",
							Phone: "8004445555",
						},
					},
				},
			},
			Custom: map[string]string{
				"Source": "Website contact form",
				"Transportation": "Segway",
			},
			Addresses: []Address{
				Address{
					Label: "business",
					Address1: "747 Howard St",
					Address2: "Room 3",
					City: "San Francisco",
					State: "CA",
					Zipcode: "94103",
					Country:"US",
				},
			},
		}
		leadresp, err := closeAPI.Create(&lead)
		if err != nil {
			log.Println(err)
		}
		log.Println(leadresp)

}
