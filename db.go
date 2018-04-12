
// 
// 
// Title               : Формирование индексов для таблицы
// Date & Time         : 09.02.2016 12:00
// Описание параметров :
//      Создание множества индексов для выбранной таблицы
//      Первый параметр имя базы
//      Второй параметр имя таблицы
//      Все остальные параметры имя индексов....
//      Sys_create_indeхs("DB","TB","IDX1","IDX2","IDX3")
//
func Sys_create_indeхs(DB, TB string, IDX ...string) {
	for _, t := range IDX {
		//rk ,err:=r.DB(DB).Table(TB).IndexCreate(t).Run(sessionArray[0])
		//defer rk.Close()
		//if err!=nil{return}
		err := r.DB(DB).Table(TB).IndexCreate(t).Exec(sessionArray[0])
		if err != nil {
		   return
		}
	}
}

