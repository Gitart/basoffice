
// Description status with customer
type Descriptstaus struct{
	Id     int
	Title  string     
	Type   string 
	Eng    string 
}

type Status struct {
     Status []Descriptstaus
     Title  string
}


// 
func StatusDescript(rw http.ResponseWriter, req *http.Request){
    Param:= Sti(req.URL.Path[len("/status/"):])-1
    // fmt.Fprintln(rw, Param)

    var (
	    dt   Status
	    data Mst
	    form string
	 )
    form="Данные для статуса \n Id....... %v\n Title ....... %s\n Type ....... %s\n  Eng ....... %s\n"

	s:=`{"Title"  :"Описание мепинга статусов для работы с заказчиком",
         "Status" :[{"Id": 1,  "Title": "Відкрита",              "Type": "Dev",  "Eng": "Open"},
	                {"Id": 2,  "Title": "Оцінка",                "Type": "Dev",  "Eng": "Estimated"},
	                {"Id": 3,  "Title": "Погоджено",             "Type": "Dev",  "Eng": "Agree"},
	                {"Id": 4,  "Title": "Розробка",              "Type": "Dev",  "Eng": "Develop"},
	                {"Id": 5,  "Title": "Уточнення",             "Type": "Dev",  "Eng": "Refinement"}, 
	                {"Id": 6,  "Title": "Уточнення замовником",  "Type": "Cust", "Eng": "Clarification by the customer"},
	                {"Id": 7,  "Title": "Анаіз",                 "Type": "Dev",  "Eng": "Analys"},
	                {"Id": 8,  "Title": "Розробка",              "Type": "Dev",  "Eng": "Develop"},
	                {"Id": 9,  "Title": "Внутрешне тестування",  "Type": "Dev",  "Eng": "Ext test"},
	                {"Id": 10, "Title": "Доробка",               "Type": "Dev",  "Eng": "Finalization"}, 
	                {"Id": 11, "Title": "Оброблено розробником", "Type": "Dev",  "Eng": "Processed by the developer"},
	                {"Id": 12, "Title": "Тестування замовником", "Type": "Cust", "Eng": "Customer test"},
	                {"Id": 13, "Title": "Передано на продуктів", "Type": "Cust", "Eng": "Prod"},
	                {"Id": 14, "Title": "Оброблена",             "Type": "Cust", "Eng": "Processed"}, 
	                {"Id": 15, "Title": "Закрита",               "Type": "Cust", "Eng": "Closed"}]}`


	eru := json.Unmarshal([]byte(s), &data)
	Err(eru, "Error unmarshaling.")

    eru = json.Unmarshal([]byte(s), &dt)
	Err(eru, "Error unmarshaling.")


    if Param > len(dt.Status) || Param < 0 {
       fmt.Fprintf(rw, "bad") 	
       return
    }
    
    fmt.Println(data["Status"].([]interface{})[Param].(map[string]interface{})["Title"])  
	fmt.Fprintf(rw, form, dt.Status[Param].Id, dt.Status[Param].Title, dt.Status[Param].Type, dt.Status[Param].Eng)
}
