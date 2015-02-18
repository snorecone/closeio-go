package closeio

import (
)

type Task struct {
	LeadId string `json:"lead_id"`
	AssignedTo string `json:"assigned_to"`
	Text string `json:"text"`
	DueDate string `json:"due_date"`
	IsComplete bool `json:"is_complete"`
}
func (c *Closeio) CreateTask(task *Task) (error) {
	data, err := marshal(task)
	if err != nil {
		return err
	}
	_, err = request("task/", "POST", c.Token, data)
	if err != nil {
		return err
	}
	return nil
}
