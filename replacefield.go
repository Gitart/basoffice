/*******************************************************************************************************************************************************
 *
 *  Преобразование строкового поля в числовой
 *  Если длинное число то округляет последние цифры меняет на 0
 *  Максимальное количество знаков 16 все остальные цифры начиная с 17 будут забиты 0
 *  Пример после обработки :  6572025760642398638 => 6572025760642398000
 *  Внимание :  Данный метод не подходит при преобразовании длинных числе свыше 16 разрядов !
 *  ReplaceFieldToDate ("test","dt","Date")
 *  ReplaceFieldToDate ("A_3565510")
 *
 *******************************************************************************************************************************************************/
func ReplaceFieldToDate(Tablename string) {
	// Описание структуры расчета потребностей
	type User struct {
		ID_DOC        string // Расчетное количество
		ID_BUSINESS   int    // Код организации
		DOC_NAME      string // Номер документа
		DOC_DATE_TIME string // Время документа
		ID            string // Расчетное количество
	}

	// Инициализация переменных
	var user User

	// Формирование услови едля фильтра
	// res, err := r.DB("HO").Table("Docs").Get("-1000036044994982706").Run(session)

	// Описание таблицы для сканирования - первые 100 записей
	// res, err := r.DB("HO").Table("Docs").Limit(100).Run(session)

	// Описание таблицы для сканирования
	res, err := r.DB("HO").Table(Tablename).Run(sessionArray[0])

	//err = res.All(&user)

	// Обработка ошибок
	if err != nil {
		log.Println(err)
	}

	n := 1
	// Цикл по записям в базе
	for res.Next(&user) {
		n++
		r.DB("HO").Table(Tablename).Get(user.ID).Update(map[string]interface{}{"DOC_DATE_TIME": FDT(user.DOC_DATE_TIME)}).RunWrite(sessionArray[0])
		// fmt.Println(user.ID + " " + FDT(user.DOC_DATE_TIME) + " " + user.DOC_DATE_TIME)
		// println(FDT(user.DOC_DATE_TIME))
		// println(n)
	}

	fmt.Println("Ready Change Date Field ...")
}

/*******************************************************************************************************************************************************
 *
 *	 Конвертация даты из одного из форматов в формат  2014-01-02 12:02:23
 *	 Преобразование дат разного вида к виду
 *	 конверитрует в дату без части времени из строки по форматам:
 *	 "dd.mm.yyyy hh:MM:ss" ->  "yyyy-mm-dd hh:MM:ss"
 *	 "dd-mm-yyyy hh:MM:ss" ->  "yyyy-mm-dd hh:MM:ss"
 *
 *******************************************************************************************************************************************************/
func FDT(DataStr string) string {
	st := strings.Split(DataStr, "-")

	if len(st[0]) != 4 {
	   n := strings.Split(DataStr, " ")
	   s := strings.Split(n[0], ".")
	   d := strings.Split(n[1], ":")
	   return fmt.Sprintf("%04s-%02s-%02s %02s:%02s:%02s", s[2], s[1], s[0], d[0], d[1], d[2])
	} else {
	   n := strings.Split(DataStr, " ")
	   s := strings.Split(n[0], "-")
	   d := strings.Split(n[1], ":")
	   return fmt.Sprintf("%04s-%02s-%02s %02s:%02s:%02s", s[0], s[1], s[2], d[0], d[1], d[2])
	}
}


/*****************************************************************************************************************************************************
 *
 *	Конверитрует в дату без части времени из строки по форматам:
 *	"dd.mm.yyyy" + "dd.mm.yyyy hh:MM:ss"
 *	"dd-mm-yyyy" + "dd-mm-yyyy hh:MM:ss"
 *	"yyyy-mm-dd" + "yyyy-mm-dd hh:MM:ss"
 *	"yyyy.mm.dd" + "yyyy.mm.dd hh:MM:ss"
 *
 ******************************************************************************************************************************************************/
func ParseDate(sDateTime string) (t time.Time, err error) {
	sDateTime = strings.TrimSpace(sDateTime)
	i := strings.Index(sDateTime, " ")

	if i > 0 {
		sDateTime = sDateTime[0:i]
	}
	t, err = time.Parse("02.01.2006", sDateTime)

	if err != nil {
		tt, ee := time.Parse("02-01-2006", sDateTime)
		if ee == nil {
			return tt, nil
		}
		tt, ee = time.Parse("2006-01-02", sDateTime)
		if ee == nil {
			return tt, nil
		}
		tt, ee = time.Parse("2006.01.02", sDateTime)
		if ee == nil {
			return tt, nil
		}
	}
	return
}


/******************************************************************************************************************************************************
 *
 *	Удаление одного документа в таблице по его ID  c записью в лог файл
 *	func DeleteDocumentById(NameTable string, IdDoc int64) {
 *
 ******************************************************************************************************************************************************/
func DeleteDocumentById(NameTable string, IdDoc string) {
	rd := r.DeleteOpts{Durability: "soft", ReturnChanges: false}
	rk, err := r.DB(DBN).Table(NameTable).Get(IdDoc).Delete(rd).Run(sessionArray[0])
	if err != nil {
		return
	}
	defer rk.Close()

	go RecLog("Delete Documents ", string(IdDoc), "Cahanged Documents in Database")
}
