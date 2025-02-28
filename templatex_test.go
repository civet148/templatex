package templatex

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"
)

const (
	builtinJsonData = `{
        "total_count": 105,
        "done_count": 28,
        "diff_count": -77,
        "audit_date": "2025-02-25",
        "total_num": 3409.67,
        "total_weight": 3310.73,
        "total_audit_num": 201.78,
        "total_audit_weight": 1347.74,
        "total_num_diff": 3245.13,
        "total_weight_diff": 4366.99,
        "audit_no": "202502250001",
        "audit_type_text": "日结",
        "audit_nature_text": "称重",
        "shift_type_text": "晚班",
        "storages": "原料仓,配件仓",
        "auditors": "张学友,刘德华",
        "list": [
            {
                "product_name": "原料",
                "product_no": "YL7010002",
                "product_specs_value": "无",
                "num": 3012.89,
                "net_weight": 3012.98,
                "audit_num": 14,
                "audit_net_weight": 17.07,
                "num_diff": 2998.89,
                "weight_diff": 2995.91,
                "sort_no": 1
            },
            {
                "product_name": "废料",
                "product_no": "FL10004",
                "product_specs_value": "",
                "num": 69.17,
                "net_weight": 69.17,
                "audit_num": 4,
                "audit_net_weight": 667,
                "num_diff": 69.07,
                "weight_diff": 734.07,
                "sort_no": 2
            },
            {
                "product_name": "K金",
                "product_no": "KJ2010001",
                "product_specs_value": "无",
                "num": 44.61,
                "net_weight": 44.61,
                "audit_num": 2,
                "audit_net_weight": 1,
                "num_diff": 42.61,
                "weight_diff": 43.61,
                "sort_no": 3
            }
        ]
    }`
)

type InventoryAuditRecord struct {
	ProductName       string  `json:"product_name"`
	ProductNo         string  `json:"product_no"`
	ProductSpecsValue string  `json:"product_specs_value"`
	Num               float64 `json:"num"`
	NetWeight         float64 `json:"net_weight"`
	AuditNum          float64 `json:"audit_num"`
	AuditNetWeight    float64 `json:"audit_net_weight"`
	NumDiff           float64 `json:"num_diff"`
	WeightDiff        float64 `json:"weight_diff"`
	SortNo            int32   `json:"sort_no"`
}

type InventoryAuditOverview struct {
	TotalCount       int32                   `json:"total_count"`
	DoneCount        int32                   `json:"done_count"`
	DiffCount        int32                   `json:"diff_count"`
	AuditDate        string                  `json:"audit_date"`
	TotalNum         float64                 `json:"total_num"`
	TotalWeight      float64                 `json:"total_weight"`
	TotalAuditNum    float64                 `json:"total_audit_num"`
	TotalAuditWeight float64                 `json:"total_audit_weight"`
	TotalNumDiff     float64                 `json:"total_num_diff"`
	TotalWeightDiff  float64                 `json:"total_weight_diff"`
	AuditNo          string                  `json:"audit_no"`
	AuditTypeText    string                  `json:"audit_type_text"`
	AuditNatureText  string                  `json:"audit_nature_text"`
	ShiftTypeText    string                  `json:"shift_type_text"`
	Storages         string                  `json:"storages"`
	Auditors         string                  `json:"auditors"`
	List             []*InventoryAuditRecord `json:"list"`
}

//go:embed overview_template.html
var builtinTemplate string

func TestGenerate(t *testing.T) {
	var data InventoryAuditOverview
	err := json.Unmarshal([]byte(builtinJsonData), &data)
	if err != nil {
		t.Fatal(err)
	}
	var strTempFilePath = "overview_template.html"
	var strHtml string
	strHtml, err = Generate(strTempFilePath, data, "./test/result.html")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("generate from file %s, template [%s]\n\n", strTempFilePath, strHtml)

	strHtml, err = Generate(builtinTemplate, data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("generate from builtin template [%s]\n\n", strHtml)
}

