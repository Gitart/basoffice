// 
// DATE        : 10.07/2016 10:00
// AUTHOR      : Savchenko Arthur
// DESCRIPTION :
//               Обновление одной таблицы в цикле
//               из другой таблицы которые связанные 
//               между собой общими полями
//
// JS
// r.db("C001").table('Matrix').forEach(function(hero) {
//  return r.db("C001").table('Log').filter({"code":hero('code')}).update({"name":hero('name'), "Direciva":hero('id'), Date:"2016-02-24" })
// })
//
func Test_Ttt(w http.ResponseWriter, h *http.Request) {
	 log.Println("Ok start")
	 var row r.Term

	TU := Mst{"NID":  row.Field("id")}                             // Обновление
	TF := Mst{"Code": row.Field("code")}                           // Фильтр
	TY := r.DB("C001").Table("Log").Filter(TF).Update(TU)
    UD := r.UpdateOpts{Durability: "soft", ReturnChanges: false, NonAtomic: true}

	r.DB("C001").
	  Table("Matrix").
	  ForEach(func(row r.Term) Mii {return TY }).Run(sessionArray[0])

	r.DB("C001").
  	  Table("Matrix").
	  ForEach( func(row r.Term) Mii {
			   return r.DB("C001").Table("Log").Filter(Mst{"code": row.Field("code")}).Update(Mst{"Directiva": row.Field("id"), "Date": r.Now()},UD)
		}).
	   Run(sessionArray[0])

	rt := Sys_Encode("Прцедура обновлена.")
	
	fmt.Fprintf(w, "Выполненна процедура.")
	fmt.Println(rt)

	// Delete all rows with the given ids
	// var response r.WriteResponse
	// Delete multiple rows by primary key
	// heroNames := []string{"0f13dc4b-63a8-47a6-9a5f-e76a6258db08", "748cc96e-f027-4c05-9591-52d1fda39b22"}
	// deleteHero := func(name r.Expr) r.Query { return r.DB("C001").Table("Log").Get(name).Delete() }
	// err := r.Expr(heroNames).ForEach(deleteHero).Run(session)
}
