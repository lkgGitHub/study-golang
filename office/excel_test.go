package office

import (
	"fmt"
	"strings"
	"study-goland/office/model"
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() *gorm.DB {
	host := "127.0.0.1"
	port := 3306
	username := "mysql"
	password := `mysql`
	dbName := "mysql"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		Logger: logger.Default.LogMode(logger.Info), // info 打印sql日志
	})
	if err != nil {
		panic(err.Error())
	}
	return db
}

func TestRead(t *testing.T) {
	db := InitDB()
	filename := `企业信息.xlsx`
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println("excelize.OpenFile error:" + err.Error())
		return
	}
	// Get value from cell by given worksheet name and axis.
	//cell, err := f.GetCellValue("Sheet1", "B2")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, row := range rows {
		if i == 0 {
			continue
		}

		company := &model.CompanyInfo{
			BusinessLicense:     "营业执照",
			LinkTel:             "1111111",
			CompanyFullName:     row[1],
			CompanyAbbreviation: row[2],
			CompanyProfession:   row[3],
			FinanceSituation:    row[4],
			//Address: row[5],
			StaffSize:               row[6],
			CompanyIntroduction:     row[7],
			IsStress:                row[9],
			SocialInsuranceType:     row[10],
			AccumulationFundPercent: row[11],
			OtherBenefits:           row[12],
			CommissionRule:          row[14],
			//CreateTime: time.Now(),
			//CreateUser: "admin",
			//UpdateTime: time.Now(),
			//UpdateUser: "admin",
		}
		if address := strings.Split(row[5], "-"); len(address) > 0 {
			company.Province = address[0]
			if len(address) > 1 {
				company.City = address[1]
			}

			if len(address) > 2 {
				company.District = address[2]
			}

			if len(address) > 3 {
				company.Street = address[3]
			}
		}

		//layout := "15-04"
		//workTime := strings.Split(row[8], "-")
		//workStartTime, err := time.Parse(layout, workTime[0])
		//if err != nil {
		//
		//}
		//workEndTime , err := time.Parse(layout, workTime[1])
		//company.WorkStartTime = workStartTime
		//company.WorkEndTime = workEndTime
		result := db.Create(company)
		if result.Error != nil {
			println("db.Create(company) error:" + result.Error.Error())
		}

		//for _, colCell := range row {
		//	fmt.Print(colCell, "\t")
		//}
		//fmt.Println()
	}

}

func TestWrite(t *testing.T) {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
