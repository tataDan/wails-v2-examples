<script>
  import { fix_and_outro_and_destroy_block } from "svelte/internal";
  import { SearchDocuments } from "../wailsjs/go/main/App";

  let docs = [];
  let queryParams = {};
  let showResults;
  let searchHasRun = false;
  let title = "";
  let publisher = "";
  let author = "";
  let subject = "";
  let person = "";
  let place = "";
  let isbn = "";
  let language = "";
  let searchType = "normalSearch";
  let hideAuthor = false;
  let hidePublisher = false;
  let hideSubject = false;
  let hidePerson = false;
  let hidePlace = false;
  let hideIsbn = false;
  let hideLanguage = false;

  /**
   * @param {string} query
   */
  function configQuery(query) {
    query = query.trim();
    return query.replaceAll(" ", "+");
  }

  function hideAll() {
    hideAuthor = true;
    hidePublisher = true;
    hideSubject = true;
    hidePerson = true;
    hidePlace = true;
    hideIsbn = true;
    hideLanguage = true;
  }

  function unhideAll() {
    hideAuthor = false;
    hidePublisher = false;
    hideSubject = false;
    hidePerson = false;
    hidePlace = false;
    hideIsbn = false;
    hideLanguage = false;
  }

  function clearEntries() {
    title = "";
    author = "";
    subject = "";
    person = "";
    place = "";
    publisher = "";
    isbn = "";
    language = "";
  }

  async function searchDocuments() {
    queryParams = {
      title: title,
      author: author,
      publisher: publisher,
      person: person,
      place: place,
      subject: subject,
      isbn: isbn,
      language: language,
    };
    showResults = false;
    searchHasRun = true;
    await SearchDocuments(queryParams, searchType).then(
      (result) => (docs = result)
    );
    showResults = true;
  }
</script>

<main>
  <br />
  <div class="textSearch">
    <input
      type="radio"
      id="normalSearch"
      class="radio"
      value="normalSearch"
      bind:group={searchType}
    />
    <label for="normalSearch">Normal Search</label>
    &nbsp;&nbsp;
    <input
      type="radio"
      id="exactTitleSearch"
      class="radio"
      value="exactTitleSearch"
      bind:group={searchType}
    />
    <label for="exactTitleSearch">Exact Title Search</label>
    &nbsp;&nbsp;
    <input
      type="radio"
      id="subStringTitleSearch"
      class="radio"
      value="subStringTitleSearch"
      bind:group={searchType}
    />
    <label for="subStringTitleSearch">Substring Title Search</label>
    &nbsp;&nbsp;
    <input
      type="radio"
      id="exactAuthorSearch"
      class="radio"
      value="exactAuthorSearch"
      bind:group={searchType}
    />
    <label for="exactAuthorSearch">Exact Author Search</label>
    &nbsp;&nbsp;
    <input
      type="radio"
      id="subStringAuthorSearch"
      class="radio"
      value="subStringAuthorSearch"
      bind:group={searchType}
    />
    <label for="subStringAuthorSearch">Substring Author Search</label>
    <br />
    &nbsp;&nbsp;
    <input
      type="radio"
      id="exactPersonSearch"
      class="radio"
      value="exactPersonSearch"
      bind:group={searchType}
    />
    <label for="exactPersonSearch">Exact Person Search</label>
    &nbsp;&nbsp;
    <input
      type="radio"
      id="subStringPersonSearch"
      class="radio"
      value="subStringPersonSearch"
      bind:group={searchType}
    />
    <label for="subStringPersonSearch">Substring Person Search</label>
    &nbsp;&nbsp;
    <input
      type="radio"
      id="exactPlaceSearch"
      class="radio"
      value="exactPlaceSearch"
      bind:group={searchType}
    />
    <label for="exactPlaceSearch">Exact Place Search</label>
    &nbsp;&nbsp;
    <input
      type="radio"
      id="subStringPlaceSearch"
      class="radio"
      value="subStringPlaceSearch"
      bind:group={searchType}
    />
    <label for="subStringPlaceSearch">Substring Place Search</label>
    <br />
    &nbsp;&nbsp;
    <input
      type="radio"
      id="exactSubjectSearch"
      class="radio"
      value="exactSubjectSearch"
      bind:group={searchType}
    />
    <label for="exactSubjectSearch">Exact Subject Search</label>
    &nbsp;&nbsp;
    <input
      type="radio"
      id="subStringSubjectSearch"
      class="radio"
      value="subStringSubjectSearch"
      bind:group={searchType}
    />
    <label for="subStringSubjectSearch">Substring Subject Search</label>
    &nbsp;&nbsp;
    <input
      type="radio"
      id="exactPublisherSearch"
      class="radio"
      value="exactPublisherSearch"
      bind:group={searchType}
    />
    <label for="exactPublisherSearch">Exact Publisher Search</label>
    &nbsp;&nbsp;
    <input
      type="radio"
      id="subStringPublisherSearch"
      class="radio"
      value="subStringPublisherSearch"
      bind:group={searchType}
    />
    <label for="subStringPublisherSearch">Substring Publisher Search</label>
  </div>
  <br />
  <div class="hideOptions">
    <input
      type="checkbox"
      id="hideAuthor"
      class="checkbox"
      value="hideAuthor"
      bind:checked={hideAuthor}
    />
    <label for="hideAuthor">Hide Author:</label>
    &nbsp;&nbsp;
    <input
      type="checkbox"
      id="hidePublisher"
      class="checkbox"
      value="hidePublisher"
      bind:checked={hidePublisher}
    />
    <label for="hidePublisher">Hide Publisher</label>
    &nbsp;&nbsp;
    <input
      type="checkbox"
      id="hideSubject"
      class="checkbox"
      value="hideSubject"
      bind:checked={hideSubject}
    />
    <label for="hideSubject">Hide Subject:</label>
    &nbsp;&nbsp;
    <input
      type="checkbox"
      id="hidePerson"
      class="checkbox"
      value="hidePerson"
      bind:checked={hidePerson}
    />
    <label for="hidePerson">Hide Person:</label>
    &nbsp;&nbsp;
    <input
      type="checkbox"
      id="hidePlace"
      class="checkbox"
      value="hidePlace"
      bind:checked={hidePlace}
    />
    <label for="hidePlace">Hide Place</label>
    &nbsp;&nbsp;
    <input
      type="checkbox"
      id="hideIsbn"
      class="checkbox"
      value="hideIsbn"
      bind:checked={hideIsbn}
    />
    <label for="hideIsbn">Hide ISBN:</label>
    &nbsp;&nbsp;
    <input
      type="checkbox"
      id="hideLanguage"
      class="checkbox"
      value="hideLanguage"
      bind:checked={hideLanguage}
    />
    <label for="hideLanguage">Hide Language:</label>
  </div>
  <br />
  <div class="input-values">
    <label for="title">Title:</label>
    <input
      autocomplete="off"
      bind:value={title}
      class="input"
      id="title"
      type="text"
    />
    &nbsp;
    <label for="author">Author:</label>
    <input
      autocomplete="off"
      bind:value={author}
      class="input"
      id="author"
      type="text"
    />
    &nbsp;
    <label for="publisher">Publisher:</label>
    <input
      autocomplete="off"
      bind:value={publisher}
      class="input"
      id="publisher"
      type="text"
    />
    <br />
    &nbsp;
    <label for="person">Person:</label>
    <input
      autocomplete="off"
      bind:value={person}
      class="input"
      id="person"
      type="text"
    />
    &nbsp;
    <label for="place">Place:</label>
    <input
      autocomplete="off"
      bind:value={place}
      class="input"
      id="place"
      type="text"
    />
    &nbsp;
    <label for="subject">Subject:</label>
    <input
      autocomplete="off"
      bind:value={subject}
      class="input"
      id="subject"
      type="text"
    />
    <br />
    <label for="isbn">ISBN:</label>
    <input
      autocomplete="off"
      bind:value={isbn}
      class="input"
      id="isbn"
      type="text"
    />
    &nbsp;
    <label for="language">Language (3-letter code) :</label>
    <input
      autocomplete="off"
      bind:value={language}
      class="input"
      id="language"
      type="text"
    />
    &nbsp; See https://www.loc.gov/standards/iso639-2/php/code_list.php
    <br /><br />
    <div class="buttons">
      <button class="btn" on:click={hideAll}>Hide All</button>
      <button class="btn" on:click={unhideAll}>Unhide All</button>
      <button class="btn" on:click={clearEntries}>Clear Entries</button>
      <button class="btn" on:click={searchDocuments}>Search Documents</button>
    </div>
  </div>

  {#if showResults}
    <div>
      {#if docs.length > 0}
        <ul>
          {#each docs as doc, i}
            <h3>Document #{i + 1}</h3>
            {#if doc.Title !== null}
              <li>Title: {doc.Title}</li>
            {:else}
              <li>Title:</li>
            {/if}
            {#if doc.Subtitle !== null}
              <li>Subtitle: {doc.Subtitle}</li>
            {:else}
              <li>Subtitle:</li>
            {/if}
            {#if !hideAuthor}
              {#if doc.Author_name !== null}
                {#each doc.Author_name as aut}
                  <li>Author: {aut}</li>
                {/each}
              {:else}
                <li>Author:</li>
              {/if}
            {/if}
            {#if !hidePublisher}
              {#if doc.Publisher !== null}
                {#each doc.Publisher as pub}
                  <li>Publisher: {pub}</li>
                {/each}
              {:else}
                <li>Publisher:</li>
              {/if}
            {/if}
            {#if !hideSubject}
              {#if doc.Subject !== null}
                {#each doc.Subject as sub}
                  <li>Subject: {sub}</li>
                {/each}
              {:else}
                <li>Subject:</li>
              {/if}
            {/if}
            {#if !hidePerson}
              {#if doc.Person !== null}
                {#each doc.Person as per}
                  <li>Person: {per}</li>
                {/each}
              {:else}
                <li>Person:</li>
              {/if}
            {/if}
            {#if !hidePlace}
              {#if doc.Place !== null}
                {#each doc.Place as pla}
                  <li>Place: {pla}</li>
                {/each}
              {:else}
                <li>Place:</li>
              {/if}
            {/if}
            {#if !hideIsbn}
              {#if doc.Isbn !== null}
                {#each doc.Isbn as isb}
                  <li>ISBN: {isb}</li>
                {/each}
              {:else}
                <li>ISBN:</li>
              {/if}
            {/if}
            {#if !hideLanguage}
              {#if doc.Language !== null}
                {#each doc.Language as lan}
                  <li>Language: {lan}</li>
                {/each}
              {:else}
                <li>Language:</li>
              {/if}
            {/if}
            <br />
            Copy and paste into browser address window: {"https://openlibrary.org" +
              doc.Key}
            <br />
          {/each}
        </ul>
        <p>
          If the search returned more than 1,000 documents, only the first 1,000
          documents will be shown.
        </p>
      {:else}
        <br />
        <p>No documents were found.</p>
      {/if}
    </div>
  {:else if searchHasRun}
    <p>Loading...</p>
  {/if}
</main>

<style>
  .checkbox {
    outline: 1px solid white;
  }

  .radio {
    outline: 1px solid white;
  }

  .btn:focus {
    border-width: 3px;
  }

  #language {
    width: 2.5em;
  }

  #isbn {
    width: 8em;
  }
</style>
