package scanner

import (
	"log"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

type testModel struct {
	ID           *int64           `json:"ID,omitempty" sql:"id"`
	Name         *string          `json:"name,omitempty" sql:"name"`
	CompanyIDs   []*int64         `json:"companyIDs,omitempty" sql:"company_ids"`
	TestEmdedded []*embeddedModel `json:"testEmbedded,omitempty" sql:"test_embedded"`
	IsTest       *bool            `json:"isTest,omitempty" sql:"is_test"`
}

type embeddedModel struct {
	ID      int64  `json:"ID,omitempty" sql:"id"`
	Content string `json:"content,omitempty" sql:"content"`
	IsBool  bool   `json:"isBool,omitempty" sql:"is_bool"`
}

type testData struct {
	query string
	cols  []string
	value interface{}
}

var testCases = []testData{
	{ // 1. Test string scan
		query: "SELECT name FROM t_test WHERE id = 1",
		cols:  []string{"name"},
		value: "testName",
	},
	{ // 2 Test integer scan
		query: "SELECT id FROM t_test WHERE id = 1",
		cols:  []string{"id"},
		value: 1,
	},
	{ // 3 Test boolean scan
		query: "SELECT is_test FROM t_test WHERE id = 1",
		cols:  []string{"is_test"},
		value: true,
	},
	{ // 4 Test int array scan
		query: "SELECT company_ids FROM t_test WHERE id = 1",
		cols:  []string{"company_ids"},
		value: []int{1, 2, 3},
	},
	{ // 5 Test nested struct scan
		query: "SELECT test_embedded FROM t_test WHERE id = 1",
		cols:  []string{"test_embedded"},
		value: []byte(`{"foo":1,"content":"something","isBool:true}`),
	},
}

func TestScan(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("Enexpected error while opening mock database")
		return
	}
	defer db.Close()

	for _, testCase := range testCases {
		rows := sqlmock.NewRows(testCase.cols).AddRow(testCase.value)
		mock.ExpectQuery(testCase.query).WillReturnRows(rows)
		testRows, err := db.Query(testCase.query)
		if err != nil {
			log.Fatal("Error while quering rows")
		}
		var m testModel
		for testRows.Next() {
			if err := Scan(&m, testRows, "sql"); err != nil {
				log.Fatal(err.Error())
			}
		}

		assert.NotEmpty(t, m)
		assert.NoError(t, err)
	}

	log.Print("End of tests")
}
