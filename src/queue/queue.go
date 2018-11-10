package queue

import (
	mysqldatabase "github.com/Electra-project/electrapay-api/src/database"
	"github.com/Electra-project/electrapay-api/src/helpers"
	"github.com/satori/go.uuid"
	"time"
)

type Queue struct {
	Id                  int64     `json: "id"`
	Reference           string    `json: "reference"`
	Category            string    `json: "category"`
	APIURL              string    `json: "apiurl"`
	APIType             string    `json: "apitype"`
	Version             string    `json: "version"`
	Parameters          string    `json: "parameters"`
	Status              string    `json: "status"`
	RequestDate         time.Time `json: "request_date"`
	RequestInfo         string    `json: "request_info"`
	ResponseDate        time.Time `json: "response_date"`
	ResponseInfo        string    `json: "response_info"`
	NodeId              string    `json: "node_id"`
	NodeRequestDate     time.Time `json: "node_request_date"`
	NodeResponseDate    time.Time `json: "node_response_date"`
	ResponseCode        string    `json: "response_code"`
	ResponseDescription string    `json: "response_description"`
}

func QueueProcess(queue Queue) (Queue, error) {
	uuid, _ := uuid.NewV4()
	queue.Reference = uuid.String()
	queue, error := queueAdd(queue)
	if error != nil {
		return queue, error
	}
	queue, error = queueWaitResponse(queue)
	if error != nil {
		return queue, error
	}
	if queue.Status != "COMPLETED_PROCESSING" {
		queue.ResponseCode = "Q001"
		queue.ResponseDescription = "No Response"
	}
	queue, error = queueClose(queue)
	if error != nil {
		return queue, error
	}
	return queue, error
}

func queueAdd(queue Queue) (Queue, error) {

	db := mysqldatabase.GetQueueDatabase()
	queue.ResponseCode = "00"
	queue.ResponseDescription = "SUCCESS"
	stmt, err := db.Prepare("INSERT INTO " +
		"queue(" +
		"reference, " +
		"category, " +
		"api_url, " +
		"api_type, " +
		"version, " +
		"parameters, " +
		"request_date, " +
		"request_info, " +
		"response_code, " +
		"response_description)  " +
		"VALUES(?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		helpers.LogErr("Queue : Cannot update Queue to process - " + err.Error())
		return Queue{}, err
	}
	res, err := stmt.Exec(queue.Reference, queue.Category, queue.APIURL, queue.APIType,
		queue.Version, queue.Parameters, time.Now().UTC(), queue.RequestInfo, queue.ResponseCode, queue.ResponseDescription)
	if err != nil {
		helpers.LogErr("Queue : Cannot insert into Queue to process - " + err.Error())
		return Queue{}, err
	}
	rowCnt, err := res.RowsAffected()
	helpers.LogInfo("Queue : New Job created in queue " + string(rowCnt) + "Row(s)")

	err1 := db.QueryRow("SELECT id "+
		"FROM queue "+
		"WHERE reference = ? "+
		"and status = 'PENDING'", queue.Reference).Scan(
		&queue.Id)
	if err1 != nil {
		helpers.LogErr("Queue : Cannot verify Queue to process - " + err.Error())
		return Queue{}, err
	}
	helpers.LogInfo("Queue : Id " + string(queue.Id) + " Processing request: " + queue.RequestInfo)
	return queue, nil

}

func queueWaitResponse(queue Queue) (Queue, error) {
	queueLoop := 1
	queuestatus := ""
	db := mysqldatabase.GetQueueDatabase()
	for queueLoop < 10 {
		err := db.QueryRow("SELECT status "+
			"FROM queue "+
			"WHERE id = ? ", queue.Id).Scan(
			&queuestatus)
		if err != nil {
			helpers.LogErr("Queue : Cannot verify Queue to process - " + err.Error())
			return Queue{}, err
		}
		if queuestatus != "COMPLETED_PROCESSING" {
			helpers.LogErr("Queue : Still Processing")
		} else {
			queueLoop = 10
		}
		queueLoop = queueLoop + 1
		time.Sleep(200 * time.Millisecond)
	}
	queueinfo, err := queueGet(queue.Id)
	return queueinfo, err
}

func queueGet(queueid int64) (Queue, error) {

	queue := Queue{}
	db := mysqldatabase.GetQueueDatabase()
	err := db.QueryRow("SELECT id,"+
		"reference, "+
		"category, "+
		"api_url, "+
		"api_type, "+
		"version, "+
		"parameters, "+
		"status, "+
		"request_date, "+
		"request_info, "+
		"response_date, "+
		"response_info, "+
		"node_id, "+
		"node_request_date, "+
		"node_response_date, "+
		"response_code, "+
		"response_description "+
		"from queue where id=?", queueid).Scan(
		&queue.Id,
		&queue.Reference,
		&queue.Category,
		&queue.APIURL,
		&queue.APIType,
		&queue.Version,
		&queue.Parameters,
		&queue.Status,
		&queue.RequestDate,
		&queue.RequestInfo,
		&queue.ResponseDate,
		&queue.ResponseInfo,
		&queue.NodeId,
		&queue.NodeRequestDate,
		&queue.NodeResponseDate,
		&queue.ResponseCode,
		&queue.ResponseDescription)
	if err != nil {
		helpers.LogErr("Queue : Cannot find Queue details - " + err.Error())
		return Queue{}, err
	}
	return queue, nil

}

func queueClose(queue Queue) (Queue, error) {

	db := mysqldatabase.GetQueueDatabase()
	stmt, err := db.Prepare("UPDATE queue " +
		"set response_code = ?," +
		"response_description = ?, " +
		"status = ?, " +
		"response_date = ? " +
		"WHERE id = ?")
	if err != nil {
		helpers.LogErr("Queue : Cannot update Queue to Close it - " + err.Error())
		return Queue{}, err
	}
	res, err := stmt.Exec(queue.ResponseCode, queue.ResponseDescription, "COMPLETED", time.Now().UTC(), queue.Id)
	if err != nil {
		helpers.LogErr("Queue :  Cannot update Queue to Close it - " + err.Error())
		return Queue{}, err
	}
	rowCnt, err := res.RowsAffected()
	helpers.LogInfo("Queue : Queue marked closed " + string(rowCnt) + "Row(s)")

	return queue, nil

}
