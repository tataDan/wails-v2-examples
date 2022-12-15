# README

## About

This application allows the user to search for information on books using search criteria such as title, subject, person, etc.

## Developing, Building & Distributing

You might need to use `wails dev -f` instead of `wails dev` the first time that you run the application.  Likewise, you might need to use `wails build -f` instead of `wails build` the first time that you build the application.

This application uses the API at [openlibrary](https://openlibrary.org/).

It is based on the information found on this web page at [how to](https://openlibrary.org/search/howto).

## Using the Application

Understanding my theory of how the openlibrary API operates is best done by example.

Suppose you enter "Albert" in the person box. Also suppose that the API has a document that contained the following for the person array:

Albert
Maria
Joseph Albert
Einstein

The API would iterate through the person array of each document and return any document that contains "Albert" in least one row in the array.  In this case it would return the document because the first and third rows contain "Albert".

If you had entered "Albert Einstein" instead, the API would also return the document, because at least one row contains "Albert" and at least one row contains "Einstein".

However, if the document contained the person array:

Albert
Maria
Joseph Albert

then the document would not be returned becase none of the rows contain "Einstein".

Suppose you wanted to see only documents that contained the exact string of "Albert Einstein".  In that case you could use the "Exact Person Search" option.

If the person array was

Albert
Maria
Joseph Albert
Stephanie Einstein

then the document would not be returned.

However, if the person array was

Albert
Maria
Joseph Albert
Albert Einstein

the document would be returned.

However if the person array contained

Albert
Maria
Joseph Albert
Albert Einstein (1879-1955)

then you would need to use the "Person Substring Search" option to see the document.

In order to avoid performance problems, this application imposes a limit of how many documents can be shown and warns the user if an excessive number of documents are returned by the API.  One way to reduce the number of documents retured by the API is to use a "compound" search.  So instead of entering just "Travel" in the subject box, you could enter "Travel" in the subject box and "Morocco" in the place box.  Another way is to use more precise criteria.  For example, instead of entering "Romance" in the title box, you could enter "Romance France".

