package steps

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
	"github.com/dimayaschhu/vocabulary/pkg/db"
	"github.com/dimayaschhu/vocabulary/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strconv"
	"time"
)

type DBStepHandler struct {
	client        *mongo.Client
	dbConfig      db.Config
	objectMatcher *utils.ObjectMatcher
}

func NewDBStepHandler(client *mongo.Client, dbConfig db.Config, objectMatcher *utils.ObjectMatcher) *DBStepHandler {
	return &DBStepHandler{
		client:        client,
		dbConfig:      dbConfig,
		objectMatcher: objectMatcher,
	}
}

func (h *DBStepHandler) RegisterSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^The next fixtures exist in "([^"]*)" table:$`, h.nextRecordsExist)
	ctx.Step(`^I see next records in "([^"]*)" table:$`, h.checkTableRecordsq)
}

func (h *DBStepHandler) nextRecordsExist(tableName string, tableVars *godog.Table) error {
	var columnNames []string

	for _, tableHeaderCell := range tableVars.Rows[0].Cells {
		columnNames = append(columnNames, tableHeaderCell.Value)
	}

	coll := h.client.Database(h.dbConfig.GetNameDB()).Collection(tableName)
	for _, tableRow := range tableVars.Rows[1:] {
		var doc bson.D
		for i, cell := range tableRow.Cells {
			doc = append(doc, bson.E{columnNames[i], cell.Value})
		}
		_, err := coll.InsertOne(context.TODO(), doc)
		if err != nil {
			panic(err.Error())
		}

	}

	return nil
}

func (h *DBStepHandler) checkTableRecords(tableName string, records *godog.Table) error {
	coll := h.client.Database(h.dbConfig.GetNameDB()).Collection(tableName)
	cursor, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}

	var resDB []struct {
		Name      string
		Translate string
	}
	if err = cursor.All(context.TODO(), &resDB); err != nil {
		return err
	}
	recordsDB := make(map[string]struct{})
	for _, raw := range resDB {
		recordsDB[raw.Name] = struct{}{}
		recordsDB[raw.Translate] = struct{}{}
	}

	recordsTest := make(map[string]struct{})
	for _, tableRow := range records.Rows[1:] {
		for _, cell := range tableRow.Cells {
			recordsTest[cell.Value] = struct{}{}
		}
	}

	for c, _ := range recordsTest {
		if _, ok := recordsDB[c]; !ok {
			return errors.New("Not found record:" + c)
		}
	}

	return nil
}

func (h *DBStepHandler) checkTableRecordsq(tableName string, records *godog.Table) error {
	var columns []string

	for _, tableCell := range records.Rows[0].Cells {
		columns = append(columns, tableCell.Value)
	}

	actualRawResult := []map[string]interface{}{}

	coll := h.client.Database(h.dbConfig.GetNameDB()).Collection(tableName)
	cursor, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &actualRawResult); err != nil {
		return err
	}

	if len(actualRawResult) != len(records.Rows[1:]) {
		return errors.New("expected table rows")
	}

	actualStringResult, err := convertMapValuesToString(actualRawResult)
	if err != nil {
		return err
	}

	expectedResult := convertGoDogTableValuesToHeadingMap(records)

	// compare
	for key, expectedValue := range expectedResult {
		if err := h.objectMatcher.Match(actualStringResult[key], expectedValue); err != nil {
			return err
		}
	}

	h.RemoveCollDB(tableName)

	return nil
}
func (h *DBStepHandler) RemoveCollDB(name string) {
	coll := h.client.Database(h.dbConfig.GetNameDB()).Collection(name)
	_, err := coll.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
}
func (h *DBStepHandler) RemoveDB() {

}

func convertGoDogTableValuesToHeadingMap(records *godog.Table) []map[string]interface{} {
	var expectedResult []map[string]interface{}

	for _, row := range records.Rows[1:] {
		expectedRowResult := make(map[string]interface{})

		for cellIndex, cell := range row.Cells {
			columnName := records.Rows[0].Cells[cellIndex].Value
			expectedRowResult[columnName] = cell.Value
		}

		expectedResult = append(expectedResult, expectedRowResult)
	}

	return expectedResult
}

func convertMapValuesToString(input []map[string]interface{}) ([]map[string]interface{}, error) {
	var actualStringResult []map[string]interface{}
	for _, actualRawRes := range input {
		actualStringMap := make(map[string]interface{})
		for key, val := range actualRawRes {
			switch val := val.(type) {
			case primitive.ObjectID:
				actualStringMap[key] = val.String()
			case string:
				actualStringMap[key] = val
			case int32:
				actualStringMap[key] = strconv.Itoa(int(val))
			case int64:
				actualStringMap[key] = strconv.Itoa(int(val))
			case time.Time:
				if val.IsZero() {
					actualStringMap[key] = ""
				} else {
					actualStringMap[key] = val.Format(time.RFC3339Nano)
				}
			case bool:
				actualStringMap[key] = strconv.FormatBool(val)
			case []uint8:
				actualStringMap[key] = string(val)
			case nil:
				actualStringMap[key] = ""
			default:
				return nil, errors.New("Unexpected value type")
			}
		}

		actualStringResult = append(actualStringResult, actualStringMap)
	}

	return actualStringResult, nil
}
