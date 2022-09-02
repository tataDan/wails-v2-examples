# README

## Developing, Building & Distributing

This application uses a sqlite database to store the user's data.

It uses the go package [go-sqlite3](https://github.com/mattn/go-sqlite3) as a driver for the database.  That package requires that a gcc compiler be on the developer's PATH. On my Linux computer gcc was already installed. On Windows I installed the MinGW-w64 gcc compiler by using the instructions at [Msys2](https://www.msys2.org/). Executing the command `pacman -S mingw-w64-x86_64-gcc` in step 6 is critical. After the installation, I added `C:\msys64\mingw64\bin` to my PATH as a System Variable.

Other gcc compilers (such as TDM-GCC) might also work.

NOTE: This application was not tested on MacOS, but it should not be too dificult to install gcc if it is not already installed.

The sqlite database is named "reference.db", and it has one table named "reference". The database provided with the application contains no records in the table.

The table was created with a sqlite command-line application using this SQL command:

`CREATE TABLE reference (
		id integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		category TEXT,
		keywords TEXT,
		date_created TEXT,
		url TEXT,
		book_title TEXT,
		isbn TEXT,
		excerpt TEXT
	  );`
	  
NOTE: Installing a sqlite command-line application is completely optional, but it could be usefull in some circumstances.

For developing, use `wails dev`. If that does not work, try `wails dev -f`.

To build, use `wails build`.

One simple way for distributing is as follows:

Create a folder with a suitable name of your choice (e.g. "research").
Copy `build\bin\research.exe` to that folder if using Windows; otherwise copy `build/bin/research` to that folder.
Copy `research.db` to that folder.
Copy that folder to a different computer if desired.

## Using the Application

The application allows the user to store references while doing research.  Here "reference" means excerpts of information found in a book or on a web site.

As the database is initially empty, you need to create at least one reference by using the "Create References" tab in order for the application to be useful.

To find, modify, or remove a reference, you will probable want to use the "Query, Update & Delete" tab. Using this tab, you are only allowed to update one column of one record at a time, and you are only allowed to delete one record at a time.

If you are comfortable using SQL statements, then you might want to use the "Custom Select", "Custom Update", or "Custom Delete" tabs. In addition to allowing you to enter and execute a SQL statement, they allow you to retrieve a SQL statement from a text file.  You can choose your own name and folder for this file. This file should contain one valid SQL statement per line.  They can be SELECT, UPDATE, or DELETE statements.

The "Backup Database" tab allows you to make a backup of the database.  You can choose the folder where the backup file will be placed, but the file name will always be "reference.db_BACKUP".

