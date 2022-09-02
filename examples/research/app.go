package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/gen2brain/dlgs"
)

type Reference struct {
	Id           int
	Category     string
	Keywords     string
	Date_created string
	Url          string
	Book_title   string
	Isbn         string
	Excerpt      string
}

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) CreateRef(ref Reference) {
	insertSql := `INSERT INTO reference(category, keywords, date_created, url, book_title, isbn, excerpt) VALUES (?, ?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertSql)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return
	}

	t := time.Now()
	date := t.Format("2006-01-02")
	_, err = statement.Exec(ref.Category, ref.Keywords, date, ref.Url, ref.Book_title, ref.Isbn, ref.Excerpt)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return
	} else {
		a.msgDlgOk(runtime.InfoDialog, "INFORMATION", "The record was inserted.")
	}
}

func (a *App) Update(id string, col string, newValue string) bool {
	ret, err := dlgs.Question("CONFIRM UPDATE", "Are you sure you want to update this reference?", true)
	if (err != nil) || (!ret) {
		return false
	}
	prepStr := fmt.Sprintf("UPDATE reference SET %s = ? WHERE id = ?", col)
	stmt, err := db.Prepare(prepStr)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return false
	}
	defer stmt.Close()
	res, err := stmt.Exec(newValue, id)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return false
	}
	rows_affected, err := res.RowsAffected()
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return false
	}
	if rows_affected == 1 {
		a.msgDlgOk(runtime.InfoDialog, "INFORMATION", "The record was updated.")
	}
	return true
}

func (a *App) UpdateCustom(sqlStmt string) int64 {
	tmpSqlStmt := strings.ToUpper(sqlStmt)

	if !strings.HasPrefix(tmpSqlStmt, "UPDATE") {
		a.msgDlgOk(runtime.WarningDialog, "WARNING", "Please enter valid UPDATE statement.")
		return 0
	}

	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "PREPARE ERROR", err.Error())
		return 0
	} else {
		if !strings.Contains(tmpSqlStmt, "WHERE") {
			ret, err := dlgs.Question("CONFIRM UPDATE", "Are you sure you want to update ALL the records in the database??", true)
			if (err != nil) || (!ret) {
				return 0
			}
		} else {
			ret, err := dlgs.Question("CONFIRM UPDATE", "Are you sure you want to update the specified record or records in the database?", true)
			if (err != nil) || (!ret) {
				return 0
			}
		}
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return 0
	}

	rows_affected, err := res.RowsAffected()
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return 0
	}
	return rows_affected
}

func (a *App) Delete(id string) bool {
	ret, err := dlgs.Question("CONFIRM DELETION", "Are you sure you want to delete this reference?", true)
	if (err != nil) || (!ret) {
		return false
	}
	stmt, err := db.Prepare("DELETE FROM reference WHERE id = ?")
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return false
	}
	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return false
	}
	rows_affected, err := res.RowsAffected()
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return false
	}
	if rows_affected == 1 {
		a.msgDlgOk(runtime.InfoDialog, "INFORMATION", "The record was deleted.")
	}
	return true
}

func (a *App) DeleteCustom(sqlStmt string) int64 {
	tmpSqlStmt := strings.ToUpper(sqlStmt)

	if !strings.HasPrefix(tmpSqlStmt, "DELETE") {
		a.msgDlgOk(runtime.WarningDialog, "WARNING", "Please enter a valid DELETE statement.")
		return 0
	}

	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "PREPARE ERROR", err.Error())
		return 0
	} else {
		if !strings.Contains(tmpSqlStmt, "WHERE") {
			ret, err := dlgs.Question("CONFIRM DELETION", "Are you sure you want to delete ALL the records in the database?", true)
			if (err != nil) || (!ret) {
				return 0
			}
		} else {
			ret, err := dlgs.Question("CONFIRM DELETION", "Are you sure you want to delete the record or records in the database?", true)
			if (err != nil) || (!ret) {
				return 0
			}
		}
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return 0
	}

	rows_affected, err := res.RowsAffected()
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return 0
	}
	return rows_affected
}

func (a *App) QueryAll() []Reference {
	var refs = []Reference{}

	rows, err := db.Query("SELECT * FROM reference ORDER BY id")
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return []Reference{}
	}
	defer rows.Close()

	var id int
	var category string
	var keywords string
	var date_created string
	var url string
	var book_title string
	var isbn string
	var excerpt string

	for rows.Next() {
		err := rows.Scan(&id, &category, &keywords, &date_created, &url, &book_title, &isbn, &excerpt)
		if err != nil {
			a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
			return []Reference{}
		}

		ref := Reference{
			Id:           id,
			Category:     category,
			Keywords:     keywords,
			Date_created: date_created,
			Url:          url,
			Book_title:   book_title,
			Isbn:         isbn,
			Excerpt:      excerpt,
		}
		refs = append(refs, ref)
	}
	return refs
}

func (a *App) QueryById(idIn string) Reference {
	if strings.TrimSpace(idIn) == "" {
		a.msgDlgOk(runtime.WarningDialog, "WARNING", "No id was entered")
		return Reference{}
	}

	id, err := strconv.Atoi(idIn)
	if err != nil {
		a.msgDlgOk(runtime.WarningDialog, "WARNING", "Id must be a positive integer without leading or trailing spaces")
		return Reference{}
	}

	var category string
	var keywords string
	var date_created string
	var url string
	var book_title string
	var isbn string
	var excerpt string

	row, err := db.Query("SELECT * FROM reference WHERE id = ?", id)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return Reference{}
	}
	defer row.Close()

	for row.Next() {
		err := row.Scan(&id, &category, &keywords, &date_created, &url, &book_title, &isbn, &excerpt)
		if err != nil {
			a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
			return Reference{}
		}
	}

	ref := Reference{
		Id:           id,
		Category:     category,
		Keywords:     keywords,
		Date_created: date_created,
		Url:          url,
		Book_title:   book_title,
		Isbn:         isbn,
		Excerpt:      excerpt,
	}

	return ref
}

func (a *App) QueryByCategory(category string) []Reference {
	var refs = []Reference{}

	if strings.TrimSpace(category) == "" {
		a.msgDlgOk(runtime.WarningDialog, "WARNING", "No category was entered")
		return []Reference{}
	}

	var id int
	var keywords string
	var date_created string
	var url string
	var book_title string
	var isbn string
	var excerpt string

	rows, err := db.Query("SELECT * FROM reference WHERE category = ?", category)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return []Reference{}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &category, &keywords, &date_created, &url, &book_title, &isbn, &excerpt)
		if err != nil {
			a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
			return []Reference{}
		}

		ref := Reference{
			Id:           id,
			Category:     category,
			Keywords:     keywords,
			Date_created: date_created,
			Url:          url,
			Book_title:   book_title,
			Isbn:         isbn,
			Excerpt:      excerpt,
		}
		refs = append(refs, ref)
	}

	return refs
}

func (a *App) QueryByKeywords(keywords string) []Reference {
	var refs = []Reference{}

	if strings.TrimSpace(keywords) == "" {
		a.msgDlgOk(runtime.WarningDialog, "WARNING", "No keyword was entered")
		return []Reference{}
	}

	var id int
	var category string
	var date_created string
	var url string
	var book_title string
	var isbn string
	var excerpt string

	rows, err := db.Query("SELECT * FROM reference WHERE keywords LIKE '%" + keywords + "%'")
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return []Reference{}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &category, &keywords, &date_created, &url, &book_title, &isbn, &excerpt)
		if err != nil {
			a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
			return []Reference{}
		}
		ref := Reference{
			Id:           id,
			Category:     category,
			Keywords:     keywords,
			Date_created: date_created,
			Url:          url,
			Book_title:   book_title,
			Isbn:         isbn,
			Excerpt:      excerpt,
		}
		refs = append(refs, ref)
	}
	return refs
}

func (a *App) QueryByDate(startDate string, endDate string) []Reference {
	var refs = []Reference{}

	if (strings.TrimSpace(startDate) == "") || (strings.TrimSpace(endDate) == "") {
		a.msgDlgOk(runtime.WarningDialog, "WARNING", "Both dates must be entered")
		return []Reference{}
	}

	var id int
	var category string
	var keywords string
	var date_created string
	var url string
	var book_title string
	var isbn string
	var excerpt string

	rows, err := db.Query("SELECT * FROM reference WHERE date_created BETWEEN ? AND ?", startDate, endDate)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return []Reference{}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &category, &keywords, &date_created, &url, &book_title, &isbn, &excerpt)
		if err != nil {
			a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
			return []Reference{}
		}
		ref := Reference{
			Id:           id,
			Category:     category,
			Keywords:     keywords,
			Date_created: date_created,
			Url:          url,
			Book_title:   book_title,
			Isbn:         isbn,
			Excerpt:      excerpt,
		}
		refs = append(refs, ref)
	}
	return refs
}

func (a *App) QueryByUrl(url string) []Reference {
	var refs = []Reference{}

	if strings.TrimSpace(url) == "" {
		a.msgDlgOk(runtime.WarningDialog, "WARNING", "No url was entered")
		return []Reference{}
	}

	var id int
	var category string
	var keywords string
	var date_created string
	var book_title string
	var isbn string
	var excerpt string

	rows, err := db.Query("SELECT * FROM reference WHERE url = ?", url)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return []Reference{}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &category, &keywords, &date_created, &url, &book_title, &isbn, &excerpt)
		if err != nil {
			a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
			return []Reference{}
		}

		ref := Reference{
			Id:           id,
			Category:     category,
			Keywords:     keywords,
			Date_created: date_created,
			Url:          url,
			Book_title:   book_title,
			Isbn:         isbn,
			Excerpt:      excerpt,
		}
		refs = append(refs, ref)
	}
	return refs
}

func (a *App) QueryByBookTitle(book_title string) []Reference {
	var refs = []Reference{}

	if strings.TrimSpace(book_title) == "" {
		a.msgDlgOk(runtime.WarningDialog, "WARNING", "No book title was entered")
		return []Reference{}
	}

	var id int
	var category string
	var keywords string
	var date_created string
	var url string
	var isbn string
	var excerpt string

	rows, err := db.Query("SELECT * FROM reference WHERE book_title = ?", book_title)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return []Reference{}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &category, &keywords, &date_created, &url, &book_title, &isbn, &excerpt)
		if err != nil {
			a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
			return []Reference{}
		}

		ref := Reference{
			Id:           id,
			Category:     category,
			Keywords:     keywords,
			Date_created: date_created,
			Url:          url,
			Book_title:   book_title,
			Isbn:         isbn,
			Excerpt:      excerpt,
		}
		refs = append(refs, ref)
	}
	return refs
}

func (a *App) QueryByIsbn(isbn string) []Reference {
	var refs = []Reference{}

	if strings.TrimSpace(isbn) == "" {
		a.msgDlgOk(runtime.WarningDialog, "WARNING", "No ISBN was entered")
		return []Reference{}
	}

	var id int
	var category string
	var keywords string
	var date_created string
	var url string
	var book_title string
	var excerpt string

	rows, err := db.Query("SELECT * FROM reference WHERE isbn = ?", isbn)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return []Reference{}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &category, &keywords, &date_created, &url, &book_title, &isbn, &excerpt)
		if err != nil {
			a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
			return []Reference{}
		}

		ref := Reference{
			Id:           id,
			Category:     category,
			Keywords:     keywords,
			Date_created: date_created,
			Url:          url,
			Book_title:   book_title,
			Isbn:         isbn,
			Excerpt:      excerpt,
		}
		refs = append(refs, ref)
	}
	return refs
}

func (a *App) QueryCustom(sqlStmt string) []Reference {
	var refs = []Reference{}

	tmpSqlStmt := strings.ToUpper(sqlStmt)
	if !strings.HasPrefix(tmpSqlStmt, "SELECT") {
		a.msgDlgOk(runtime.WarningDialog, "WARNING", "No SELECT statement was entered")
		return []Reference{}
	}

	rows, err := db.Query(sqlStmt)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return []Reference{}
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return []Reference{}
	}

	colNum := len(columns)

	for rows.Next() {
		ref := Reference{}
		cols := make([]interface{}, colNum)
		for i := 0; i < colNum; i++ {
			cols[i] = ReferenceCol(columns[i], &ref)
		}
		err := rows.Scan(cols...)
		if err != nil {
			a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
			return []Reference{}
		}
		refs = append(refs, ref)
	}
	return refs
}

func (a *App) GetIds() []string {
	var ids = []string{}
	var id string

	rows, err := db.Query("SELECT id FROM reference ORDER BY id")
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return []string{}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
			return []string{}
		}
		ids = append(ids, id)
	}
	return ids
}

func (a *App) GetCategories() []string {
	var categories = []string{}
	var category string

	rows, err := db.Query("SELECT DISTINCT category FROM reference ORDER BY category")
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return []string{}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&category)
		if err != nil {
			a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
			return []string{}
		}
		if len(category) > 0 {
			categories = append(categories, category)
		}
	}
	return categories
}

func (a *App) GetKeywords() []string {
	var kws = []string{}
	var keywords string

	rows, err := db.Query("SELECT keywords FROM reference")
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return []string{}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&keywords)
		if err != nil {
			a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
			return []string{}
		}

		kw := strings.Split(keywords, ",")

		for _, v := range kw {
			v = strings.TrimSpace(v)
			if (len(v) > 0) && (!contains(kws, v)) {
				kws = append(kws, v)
			}
		}
	}
	sort.Strings(kws)
	return kws
}

func (a *App) GetUrls() []string {
	var urls = []string{}
	var url string

	rows, err := db.Query("SELECT DISTINCT url FROM reference ORDER BY url")
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return []string{}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&url)
		if err != nil {
			a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
			return []string{}
		}
		if len(url) > 0 {
			urls = append(urls, url)
		}
	}
	return urls
}

func (a *App) GetBookTitles() []string {
	var book_titles = []string{}
	var book_title string

	rows, err := db.Query("SELECT DISTINCT book_title FROM reference ORDER BY book_title")
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return []string{}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book_title)
		if err != nil {
			a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
			return []string{}
		}
		if len(book_title) > 0 {
			book_titles = append(book_titles, book_title)
		}
	}
	return book_titles
}

func (a *App) GetIsbns() []string {
	var isbns = []string{}
	var isbn string

	rows, err := db.Query("SELECT DISTINCT isbn FROM reference ORDER BY isbn")
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return []string{}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&isbn)
		if err != nil {
			a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
			return []string{}
		}
		if len(isbn) > 0 {
			isbns = append(isbns, isbn)
		}
	}
	return isbns
}

func (a *App) BackupDb() {
	destDir, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Destination Folder for Backup",
	})
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return
	}
	if destDir == "" {
		return
	}
	destFile := destDir + string(os.PathSeparator) + "./reference.db_BACKUP"
	_, err = copy("./reference.db", destFile)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return
	} else {
		a.msgDlgOk(runtime.InfoDialog, "INFORMATION", "Backup file ("+destFile+") was created.")
	}
}

func (a *App) OpenCustomStmtFile() []string {
	var lines []string

	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select File with Custom SQL Statements",
	})
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return []string{}
	}
	if file == "" {
		return []string{}
	}
	f, err := os.Open(file)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return []string{}
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return []string{}
	}
	return lines
}

func (a *App) msgDlgOk(dlgType runtime.DialogType, title string, msg string) {
	_, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    dlgType,
		Title:   title,
		Message: msg,
	})
	if err != nil {
		log.Println(err)
	}
}

func ReferenceCol(colname string, ref *Reference) interface{} {
	switch colname {
	case "id":
		return &ref.Id
	case "category":
		return &ref.Category
	case "keywords":
		return &ref.Keywords
	case "date_created":
		return &ref.Date_created
	case "url":
		return &ref.Url
	case "book_title":
		return &ref.Book_title
	case "isbn":
		return &ref.Isbn
	case "excerpt":
		return &ref.Excerpt
	default:
		log.Println("unknown column " + colname)
		return ""
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
